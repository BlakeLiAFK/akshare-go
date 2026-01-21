package energy

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
)

// EnergyCarbonDomesticItem 碳交易网-行情信息项
type EnergyCarbonDomesticItem struct {
	Date     string  `json:"date"`     // 日期
	Price    float64 `json:"price"`    // 成交价
	Volume   float64 `json:"volume"`   // 成交量
	Amount   float64 `json:"amount"`   // 成交额
	Location string  `json:"location"` // 地点
}

// EnergyCarbonBJItem 北京市碳排放权公开交易行情项
type EnergyCarbonBJItem struct {
	Date     string  `json:"date"`      // 日期
	Volume   float64 `json:"volume"`    // 成交量
	AvgPrice float64 `json:"avg_price"` // 成交均价
	Amount   float64 `json:"amount"`    // 成交额
	Unit     string  `json:"unit"`      // 成交单位
}

// EnergyCarbonSZItem 深圳碳排放交易所-国内碳情项
type EnergyCarbonSZItem struct {
	Date     string  `json:"date"`      // 交易日期
	Open     float64 `json:"open"`      // 开盘价
	High     float64 `json:"high"`      // 最高价
	Low      float64 `json:"low"`       // 最低价
	AvgPrice float64 `json:"avg_price"` // 成交均价
	Close    float64 `json:"close"`     // 收盘价
	Volume   float64 `json:"volume"`    // 成交量
	Amount   float64 `json:"amount"`    // 成交额
}

// EnergyCarbonHBItem 湖北碳排放权交易中心-现货交易数据项
type EnergyCarbonHBItem struct {
	Date   string  `json:"date"`   // 日期
	Price  float64 `json:"price"`  // 成交价
	Volume float64 `json:"volume"` // 成交量
	Latest float64 `json:"latest"` // 最新
	Change float64 `json:"change"` // 涨跌
}

// EnergyCarbonGZItem 广州碳排放权交易中心-行情信息项
type EnergyCarbonGZItem struct {
	Date      string  `json:"date"`       // 日期
	Product   string  `json:"product"`    // 品种
	Open      float64 `json:"open"`       // 开盘价
	Close     float64 `json:"close"`      // 收盘价
	High      float64 `json:"high"`       // 最高价
	Low       float64 `json:"low"`        // 最低价
	Change    float64 `json:"change"`     // 涨跌
	ChangePct float64 `json:"change_pct"` // 涨跌幅
	Volume    float64 `json:"volume"`     // 成交数量
	Amount    float64 `json:"amount"`     // 成交金额
}

// EnergyCarbonDomestic 碳交易网-行情信息
// symbol: 地区，可选 {"湖北", "上海", "北京", "重庆", "广东", "天津", "深圳", "福建"}
func EnergyCarbonDomestic(symbol string) ([]EnergyCarbonDomesticItem, error) {
	url := "http://k.tanjiaoyi.com:8080/KDataController/getHouseDatasInAverage.do"
	params := map[string]string{
		"lcnK":  "53f75bfcefff58e4046ccfa42171636c",
		"brand": "TAN",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("获取碳交易行情失败: %w", err)
	}

	bodyStr := string(resp.Body())
	// 解析JSONP格式: callback({...})
	startIdx := strings.Index(bodyStr, "(")
	endIdx := strings.LastIndex(bodyStr, ")")
	if startIdx == -1 || endIdx == -1 {
		return nil, fmt.Errorf("解析响应格式失败")
	}
	jsonStr := bodyStr[startIdx+1 : endIdx]

	result := gjson.Parse(jsonStr)
	dataArr := result.Get(symbol).Array()

	var items []EnergyCarbonDomesticItem
	for _, item := range dataArr {
		arr := item.Array()
		if len(arr) >= 6 {
			items = append(items, EnergyCarbonDomesticItem{
				Price:    arr[0].Float(),
				Volume:   arr[2].Float(),
				Location: arr[3].String(),
				Amount:   arr[4].Float(),
				Date:     arr[5].String(),
			})
		}
	}

	return items, nil
}

// EnergyCarbonBJ 北京市碳排放权电子交易平台-北京市碳排放权公开交易行情
// https://www.bjets.com.cn/article/jyxx/
func EnergyCarbonBJ() ([]EnergyCarbonBJItem, error) {
	baseURL := "https://www.bjets.com.cn/article/jyxx/"

	// 获取第一页确定总页数
	resp, err := utils.GetWithHeaders(baseURL, nil, map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
	})
	if err != nil {
		return nil, fmt.Errorf("获取北京碳交易行情失败: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(resp.Body())))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	// 解析总页数
	scriptText := doc.Find("table script").Text()
	re := regexp.MustCompile(`=\s*"?(\d+)"?`)
	matches := re.FindStringSubmatch(scriptText)
	totalPage := 1
	if len(matches) > 1 {
		fmt.Sscanf(matches[1], "%d", &totalPage)
	}

	var items []EnergyCarbonBJItem

	// 解析表格数据的辅助函数
	parseTable := func(doc *goquery.Document) {
		doc.Find("table tr").Each(func(i int, s *goquery.Selection) {
			if i == 0 {
				return // 跳过表头
			}
			tds := s.Find("td")
			if tds.Length() >= 4 {
				dateStr := strings.TrimSpace(tds.Eq(0).Text())
				volumeStr := strings.TrimSpace(tds.Eq(1).Text())
				avgPriceStr := strings.TrimSpace(tds.Eq(2).Text())
				amountStr := strings.TrimSpace(tds.Eq(3).Text())

				// 解析成交额和单位
				unit := ""
				if strings.Contains(amountStr, "(") {
					parts := strings.Split(amountStr, "(")
					amountStr = parts[0]
					if len(parts) > 1 {
						unit = strings.Trim(parts[1], ")")
					}
				}
				amountStr = strings.ReplaceAll(amountStr, ",", "")

				item := EnergyCarbonBJItem{
					Date: dateStr,
					Unit: unit,
				}
				fmt.Sscanf(volumeStr, "%f", &item.Volume)
				fmt.Sscanf(avgPriceStr, "%f", &item.AvgPrice)
				fmt.Sscanf(amountStr, "%f", &item.Amount)

				items = append(items, item)
			}
		})
	}

	// 解析第一页
	parseTable(doc)

	// 解析其他页
	for page := 2; page <= totalPage; page++ {
		pageURL := fmt.Sprintf("%s?%d", baseURL, page)
		resp, err := utils.GetWithHeaders(pageURL, nil, map[string]string{
			"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		})
		if err != nil {
			continue
		}
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(resp.Body())))
		if err != nil {
			continue
		}
		parseTable(doc)
	}

	// 按日期排序
	sort.Slice(items, func(i, j int) bool {
		return items[i].Date < items[j].Date
	})

	return items, nil
}

// EnergyCarbonSZ 深圳碳排放交易所-国内碳情
// http://www.cerx.cn/dailynewsCN/index.htm
func EnergyCarbonSZ() ([]EnergyCarbonSZItem, error) {
	return fetchCarbonSZData("http://www.cerx.cn/dailynewsCN/index")
}

// EnergyCarbonEU 深圳碳排放交易所-国际碳情
// http://www.cerx.cn/dailynewsOuter/index.htm
func EnergyCarbonEU() ([]EnergyCarbonSZItem, error) {
	return fetchCarbonSZData("http://www.cerx.cn/dailynewsOuter/index")
}

// fetchCarbonSZData 获取深圳碳排放交易所数据的通用函数
func fetchCarbonSZData(baseURL string) ([]EnergyCarbonSZItem, error) {
	url := baseURL + ".htm"
	resp, err := utils.GetWithHeaders(url, nil, map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
	})
	if err != nil {
		return nil, fmt.Errorf("获取深圳碳交易行情失败: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(resp.Body())))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	// 获取总页数
	totalPage := 1
	doc.Find(".pagebar option").Each(func(i int, s *goquery.Selection) {
		var page int
		fmt.Sscanf(s.Text(), "%d", &page)
		if page > totalPage {
			totalPage = page
		}
	})

	var items []EnergyCarbonSZItem

	// 解析表格数据
	parseTable := func(doc *goquery.Document) {
		doc.Find("table tr").Each(func(i int, s *goquery.Selection) {
			if i == 0 {
				return // 跳过表头
			}
			tds := s.Find("td")
			if tds.Length() >= 8 {
				item := EnergyCarbonSZItem{
					Date: strings.TrimSpace(tds.Eq(0).Text()),
				}
				fmt.Sscanf(strings.TrimSpace(tds.Eq(1).Text()), "%f", &item.Open)
				fmt.Sscanf(strings.TrimSpace(tds.Eq(2).Text()), "%f", &item.High)
				fmt.Sscanf(strings.TrimSpace(tds.Eq(3).Text()), "%f", &item.Low)
				fmt.Sscanf(strings.TrimSpace(tds.Eq(4).Text()), "%f", &item.AvgPrice)
				fmt.Sscanf(strings.TrimSpace(tds.Eq(5).Text()), "%f", &item.Close)
				fmt.Sscanf(strings.TrimSpace(tds.Eq(6).Text()), "%f", &item.Volume)
				fmt.Sscanf(strings.TrimSpace(tds.Eq(7).Text()), "%f", &item.Amount)

				items = append(items, item)
			}
		})
	}

	// 解析第一页
	parseTable(doc)

	// 解析其他页
	for page := 2; page <= totalPage; page++ {
		pageURL := fmt.Sprintf("%s_%d.htm", baseURL, page)
		resp, err := utils.GetWithHeaders(pageURL, nil, map[string]string{
			"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		})
		if err != nil {
			continue
		}
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(resp.Body())))
		if err != nil {
			continue
		}
		parseTable(doc)
	}

	// 按日期排序
	sort.Slice(items, func(i, j int) bool {
		return items[i].Date < items[j].Date
	})

	return items, nil
}

// EnergyCarbonHB 湖北碳排放权交易中心-现货交易数据-配额-每日概况
// http://www.hbets.cn/
func EnergyCarbonHB() ([]EnergyCarbonHBItem, error) {
	url := "https://www.hbets.cn/"
	resp, err := utils.GetWithHeaders(url, nil, map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
	})
	if err != nil {
		return nil, fmt.Errorf("获取湖北碳交易行情失败: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(resp.Body())))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	// 提取script中的JSON数据
	scriptText := ""
	doc.Find("div.threeLeft script").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		if strings.Contains(text, "cjj = '[") {
			scriptText = text
		}
	})

	if scriptText == "" {
		return nil, fmt.Errorf("未找到数据")
	}

	// 解析JSON数组
	startIdx := strings.Index(scriptText, "cjj = '[") + 7
	endIdx := strings.Index(scriptText[startIdx:], "cjj =")
	if endIdx == -1 {
		endIdx = len(scriptText) - startIdx
	}
	jsonStr := scriptText[startIdx : startIdx+endIdx]
	jsonStr = strings.TrimRight(jsonStr, "'; \n\t\r")

	result := gjson.Parse(jsonStr)
	dataArr := result.Array()

	var items []EnergyCarbonHBItem
	for _, item := range dataArr {
		items = append(items, EnergyCarbonHBItem{
			Date:   item.Get("riqi").String(),
			Price:  item.Get("cjj").Float(),
			Volume: item.Get("cjl").Float(),
			Latest: item.Get("zx").Float(),
			Change: item.Get("zd").Float(),
		})
	}

	return items, nil
}

// EnergyCarbonGZ 广州碳排放权交易中心-行情信息
// http://www.cnemission.com/article/hqxx/
func EnergyCarbonGZ() ([]EnergyCarbonGZItem, error) {
	url := "http://ets.cnemission.com/carbon/portalIndex/markethistory"
	params := map[string]string{
		"Top":       "1",
		"beginTime": "2010-01-01",
		"endTime":   "2030-09-12",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("获取广州碳交易行情失败: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(resp.Body())))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	var items []EnergyCarbonGZItem

	// 解析第二个表格
	doc.Find("table").Eq(1).Find("tr").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			return // 跳过表头
		}
		tds := s.Find("td")
		if tds.Length() >= 10 {
			item := EnergyCarbonGZItem{
				Date:    strings.TrimSpace(tds.Eq(0).Text()),
				Product: strings.TrimSpace(tds.Eq(1).Text()),
			}
			fmt.Sscanf(strings.TrimSpace(tds.Eq(2).Text()), "%f", &item.Open)
			fmt.Sscanf(strings.TrimSpace(tds.Eq(3).Text()), "%f", &item.Close)
			fmt.Sscanf(strings.TrimSpace(tds.Eq(4).Text()), "%f", &item.High)
			fmt.Sscanf(strings.TrimSpace(tds.Eq(5).Text()), "%f", &item.Low)
			fmt.Sscanf(strings.TrimSpace(tds.Eq(6).Text()), "%f", &item.Change)

			changePctStr := strings.TrimSpace(tds.Eq(7).Text())
			changePctStr = strings.TrimSuffix(changePctStr, "%")
			fmt.Sscanf(changePctStr, "%f", &item.ChangePct)

			fmt.Sscanf(strings.TrimSpace(tds.Eq(8).Text()), "%f", &item.Volume)
			fmt.Sscanf(strings.TrimSpace(tds.Eq(9).Text()), "%f", &item.Amount)

			items = append(items, item)
		}
	})

	// 按日期排序
	sort.Slice(items, func(i, j int) bool {
		return items[i].Date < items[j].Date
	})

	return items, nil
}
