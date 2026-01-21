package futures

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/PuerkitoBio/goquery"
)

// FuturesSpotBasis 现货基差数据
type FuturesSpotBasis struct {
	Date              string  `json:"date"`                // 日期
	Symbol            string  `json:"symbol"`              // 品种代码
	SpotPrice         float64 `json:"spot_price"`          // 现货价格
	NearContract      string  `json:"near_contract"`       // 临近交割合约
	NearContractPrice float64 `json:"near_contract_price"` // 临近交割合约结算价
	DomContract       string  `json:"dom_contract"`        // 主力合约
	DomContractPrice  float64 `json:"dom_contract_price"`  // 主力合约结算价
	NearBasis         float64 `json:"near_basis"`          // 临近合约基差
	DomBasis          float64 `json:"dom_basis"`           // 主力合约基差
	NearBasisRate     float64 `json:"near_basis_rate"`     // 临近合约基差率
	DomBasisRate      float64 `json:"dom_basis_rate"`      // 主力合约基差率
}

// FuturesSpotBasisPrevious 现货基差数据(历史格式)
type FuturesSpotBasisPrevious struct {
	Commodity        string  `json:"commodity"`          // 商品
	SpotPrice        float64 `json:"spot_price"`         // 现货价格
	DomContract      string  `json:"dom_contract"`       // 主力合约代码
	DomContractPrice float64 `json:"dom_contract_price"` // 主力合约价格
	DomBasis         float64 `json:"dom_basis"`          // 主力合约基差
	DomBasisRatio    float64 `json:"dom_basis_ratio"`    // 主力合约变动百分比
	BasisHigh180     float64 `json:"basis_high_180"`     // 180日内主力基差最高
	BasisLow180      float64 `json:"basis_low_180"`      // 180日内主力基差最低
	BasisAvg180      float64 `json:"basis_avg_180"`      // 180日内主力基差平均
}

// FuturesSpotPrice 获取指定交易日大宗商品现货价格及相应基差
//
// 数据源: https://www.100ppi.com/sf/
//
// 参数:
//   - date: 交易日，格式 "20240430"
//   - varsList: 合约品种列表，如 []string{"RB", "AL"}，为空则返回所有
//
// 返回:
//   - []FuturesSpotBasis: 现货基差数据
//   - error: 错误信息
func FuturesSpotPrice(date string, varsList []string) ([]FuturesSpotBasis, error) {
	d, err := time.Parse("20060102", date)
	if err != nil {
		return nil, fmt.Errorf("日期格式错误: %w", err)
	}

	if d.Before(time.Date(2011, 1, 4, 0, 0, 0, 0, time.Local)) {
		return nil, fmt.Errorf("数据源开始日期为 20110104")
	}

	// 尝试获取数据
	urls := []string{
		fmt.Sprintf("https://www.100ppi.com/sf/day-%s.html", d.Format("2006-01-02")),
		"https://www.100ppi.com/sf/",
	}

	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
	}

	var doc *goquery.Document
	for _, url := range urls {
		resp, err := utils.GetWithHeaders(url, nil, headers)
		if err != nil {
			continue
		}

		doc, err = goquery.NewDocumentFromReader(strings.NewReader(string(resp.Body())))
		if err != nil {
			continue
		}

		// 验证日期
		dateText := doc.Find("table").First().Find("tr").Eq(1).Find("td").Eq(1).Text()
		re := regexp.MustCompile(`\d+`)
		nums := re.FindAllString(dateText, -1)
		if len(nums) >= 4 {
			dateStr := strings.Join(nums[0:4], "")
			if len(dateStr) >= 8 && dateStr[0:8] == date {
				break
			}
		}
		doc = nil
	}

	if doc == nil {
		return nil, fmt.Errorf("无法获取 %s 的数据", date)
	}

	// 解析数据表格
	var result []FuturesSpotBasis
	chineseRe := regexp.MustCompile(`[\p{Han}]+`)
	digitRe := regexp.MustCompile(`\d+`)

	doc.Find("table").Eq(1).Find("tr").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			return // 跳过表头
		}

		tds := s.Find("td")
		if tds.Length() < 7 {
			return
		}

		commodityText := strings.TrimSpace(tds.Eq(0).Text())

		// 提取中文名
		chineseNames := chineseRe.FindAllString(commodityText, -1)
		if len(chineseNames) == 0 {
			// 可能是英文品种如 PTA
			chineseNames = []string{strings.TrimSpace(commodityText)}
		}

		name := strings.Join(chineseNames, "")
		if name == "" || name == "商品" || name == "价格" {
			return
		}

		// 跳过交易所名称
		skipNames := map[string]bool{
			"上海期货交易所": true,
			"郑州商品交易所": true,
			"大连商品交易所": true,
			"广州期货交易所": true,
			"暂无数据":    true,
		}
		if skipNames[name] {
			return
		}

		// 转换为英文代码
		symbol := ChineseToEnglish(name)
		if symbol == "" {
			symbol = name
		}

		// 检查是否在筛选列表中
		if len(varsList) > 0 {
			found := false
			for _, v := range varsList {
				if v == symbol {
					found = true
					break
				}
			}
			if !found {
				return
			}
		}

		// 解析价格数据
		spotPrice := utils.MustParseFloat(tds.Eq(1).Text())
		nearContract := strings.TrimSpace(tds.Eq(2).Text())
		nearContractPrice := utils.MustParseFloat(tds.Eq(3).Text())
		domContract := strings.TrimSpace(tds.Eq(5).Text())
		domContractPrice := utils.MustParseFloat(tds.Eq(6).Text())

		// 特殊品种单位转换
		switch symbol {
		case "JD": // 鸡蛋：现货元/公斤，期货元/500千克
			spotPrice *= 500
		case "FG": // 玻璃：现货元/平方米，期货元/吨
			spotPrice *= 80
		case "LH": // 生猪：现货元/公斤，期货元/吨
			spotPrice *= 1000
		}

		// 提取合约月份
		nearMonth := digitRe.FindString(nearContract)
		domMonth := digitRe.FindString(domContract)

		// 构建合约代码
		nearContractCode := formatContractCode(symbol, nearMonth)
		domContractCode := formatContractCode(symbol, domMonth)

		// 计算基差
		nearBasis := nearContractPrice - spotPrice
		domBasis := domContractPrice - spotPrice
		var nearBasisRate, domBasisRate float64
		if spotPrice != 0 {
			nearBasisRate = nearContractPrice/spotPrice - 1
			domBasisRate = domContractPrice/spotPrice - 1
		}

		result = append(result, FuturesSpotBasis{
			Date:              date,
			Symbol:            symbol,
			SpotPrice:         spotPrice,
			NearContract:      nearContractCode,
			NearContractPrice: nearContractPrice,
			DomContract:       domContractCode,
			DomContractPrice:  domContractPrice,
			NearBasis:         nearBasis,
			DomBasis:          domBasis,
			NearBasisRate:     nearBasisRate,
			DomBasisRate:      domBasisRate,
		})
	})

	return result, nil
}

// FuturesSpotPriceDaily 获取指定时间段内大宗商品现货价格及相应基差
//
// 数据源: https://www.100ppi.com/sf/
//
// 参数:
//   - startDay: 开始日期，格式 "20210201"
//   - endDay: 结束日期，格式 "20210208"
//   - varsList: 合约品种列表，如 []string{"RB", "AL"}，为空则返回所有
//
// 返回:
//   - []FuturesSpotBasis: 现货基差数据
//   - error: 错误信息
func FuturesSpotPriceDaily(startDay, endDay string, varsList []string) ([]FuturesSpotBasis, error) {
	start, err := time.Parse("20060102", startDay)
	if err != nil {
		return nil, fmt.Errorf("开始日期格式错误: %w", err)
	}

	end, err := time.Parse("20060102", endDay)
	if err != nil {
		return nil, fmt.Errorf("结束日期格式错误: %w", err)
	}

	var result []FuturesSpotBasis
	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		dateStr := d.Format("20060102")
		data, err := FuturesSpotPrice(dateStr, varsList)
		if err != nil {
			continue // 跳过非交易日或获取失败的日期
		}
		result = append(result, data...)
	}

	return result, nil
}

// FuturesSpotPricePrevious 获取具体交易日大宗商品现货价格及相应基差(历史格式)
//
// 数据源: https://www.100ppi.com/sf2/
//
// 参数:
//   - date: 交易日，格式 "20240430"
//
// 返回:
//   - []FuturesSpotBasisPrevious: 现货基差数据
//   - error: 错误信息
func FuturesSpotPricePrevious(date string) ([]FuturesSpotBasisPrevious, error) {
	d, err := time.Parse("20060102", date)
	if err != nil {
		return nil, fmt.Errorf("日期格式错误: %w", err)
	}

	if d.Before(time.Date(2011, 1, 4, 0, 0, 0, 0, time.Local)) {
		return nil, fmt.Errorf("数据源开始日期为 20110104")
	}

	url := fmt.Sprintf("https://www.100ppi.com/sf2/day-%s.html", d.Format("2006-01-02"))
	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
	}

	resp, err := utils.GetWithHeaders(url, nil, headers)
	if err != nil {
		return nil, fmt.Errorf("请求数据失败: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(resp.Body())))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	var result []FuturesSpotBasisPrevious

	// 解析主表格
	doc.Find("table").Eq(1).Find("tr").Each(func(i int, s *goquery.Selection) {
		if i <= 1 {
			return // 跳过表头
		}

		tds := s.Find("td")
		if tds.Length() < 9 {
			return
		}

		// 检查是否有有效数据(通过百分号判断)
		percentText := strings.TrimSpace(tds.Eq(4).Text())
		if !strings.HasSuffix(percentText, "%") {
			return
		}

		result = append(result, FuturesSpotBasisPrevious{
			Commodity:        strings.TrimSpace(tds.Eq(0).Text()),
			SpotPrice:        utils.MustParseFloat(tds.Eq(1).Text()),
			DomContract:      strings.TrimSpace(tds.Eq(2).Text()),
			DomContractPrice: utils.MustParseFloat(tds.Eq(3).Text()),
			DomBasis:         utils.MustParseFloat(strings.TrimSuffix(tds.Eq(4).Text(), "%")),
			DomBasisRatio:    utils.MustParseFloat(strings.TrimSuffix(tds.Eq(5).Text(), "%")),
			BasisHigh180:     utils.MustParseFloat(tds.Eq(6).Text()),
			BasisLow180:      utils.MustParseFloat(tds.Eq(7).Text()),
			BasisAvg180:      utils.MustParseFloat(tds.Eq(8).Text()),
		})
	})

	return result, nil
}

// formatContractCode 格式化合约代码
func formatContractCode(symbol, month string) string {
	if month == "" {
		return symbol
	}

	// 上期所和大商所使用小写
	shfeDceSymbols := map[string]bool{
		"CU": true, "AL": true, "ZN": true, "PB": true, "NI": true, "SN": true,
		"AU": true, "AG": true, "RB": true, "WR": true, "HC": true, "SS": true,
		"BU": true, "RU": true, "SP": true, "FU": true, "SC": true, "NR": true,
		"LU": true, "BC": true, "AO": true, "BR": true,
		"C": true, "CS": true, "A": true, "B": true, "M": true, "Y": true,
		"P": true, "FB": true, "BB": true, "JD": true, "RR": true, "L": true,
		"V": true, "PP": true, "J": true, "JM": true, "I": true, "EG": true,
		"EB": true, "PG": true, "LH": true,
	}

	// 郑商所使用3位月份
	czceSymbols := map[string]bool{
		"WH": true, "PM": true, "RI": true, "JR": true, "LR": true, "RS": true,
		"CF": true, "CY": true, "SR": true, "RM": true, "OI": true, "TA": true,
		"MA": true, "FG": true, "SF": true, "SM": true, "ZC": true, "AP": true,
		"CJ": true, "UR": true, "SA": true, "PF": true, "PK": true, "SH": true,
		"PX": true,
	}

	if shfeDceSymbols[symbol] {
		return strings.ToLower(symbol) + month
	}

	if czceSymbols[symbol] && len(month) == 4 {
		// 郑商所使用3位月份，去掉年份的第一位
		return symbol + month[1:]
	}

	return symbol + month
}
