package index

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// IndexZHSpotSymbol 新浪指数系列类型
type IndexZHSpotSymbol string

const (
	// IndexZHMain 沪深重要指数
	IndexZHMain IndexZHSpotSymbol = "沪深重要指数"
	// IndexZHSHSeries 上证系列指数
	IndexZHSHSeries IndexZHSpotSymbol = "上证系列指数"
	// IndexZHSZSeries 深证系列指数
	IndexZHSZSeries IndexZHSpotSymbol = "深证系列指数"
	// IndexZHCons 指数成份
	IndexZHCons IndexZHSpotSymbol = "指数成份"
	// IndexZHCSISeries 中证系列指数
	IndexZHCSISeries IndexZHSpotSymbol = "中证系列指数"
)

// StockZHIndexSpotEM 东方财富网-中国指数实时行情
//
// 参数:
//   - symbol: 指数类型，可选值见 IndexZHSpotSymbol 常量
//
// 返回:
//   - []IndexQuote: 指数实时行情列表
//   - error: 错误信息
//
// 示例:
//
//	// 获取沪深重要指数实时行情
//	quotes, err := index.StockZHIndexSpotEM(index.IndexZHMain)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, q := range quotes {
//	    fmt.Printf("%s(%s): %.2f, %.2f%%\n", q.Name, q.Code, q.Price, q.ChangePct)
//	}
func StockZHIndexSpotEM(symbol string) ([]IndexQuote, error) {
	// 沪深重要指数使用特殊接口
	if symbol == string(IndexZHMain) {
		return fetchMainIndexSpot()
	}

	// 其他指数使用通用接口
	symbolMap := map[string]string{
		string(IndexZHSHSeries):  "m:1+t:1",
		string(IndexZHSZSeries):  "m:0 t:5",
		string(IndexZHCons):      "m:1+s:3,m:0+t:5",
		string(IndexZHCSISeries): "m:2",
	}

	fs, ok := symbolMap[symbol]
	if !ok {
		return nil, fmt.Errorf("不支持的指数类型: %s", symbol)
	}

	return fetchAllIndexSpot(fs)
}

// fetchMainIndexSpot 获取沪深重要指数实时行情
func fetchMainIndexSpot() ([]IndexQuote, error) {
	params := map[string]string{
		"pn":    "1",
		"pz":    "100",
		"po":    "1",
		"np":    "1",
		"ut":    "bd1d9ddb04089700cf9c27f6f7426281",
		"fltt":  "2",
		"invt":  "2",
		"wbp2u": "|0|0|0|web",
		"fid":   "",
		"fs":    "b:MK0010",
		"fields": "f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18,f20,f21," +
			"f23,f24,f25,f26,f22,f11,f62,f128,f136,f115,f152",
	}

	headers := map[string]string{
		"Referer": "https://quote.eastmoney.com/center/hszs.html",
	}

	resp, err := utils.GetWithHeaders(EmIndexListURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求指数行情失败: %w", err)
	}

	return parseIndexSpotResponse(resp.String())
}

// fetchAllIndexSpot 获取全部指数实时行情
func fetchAllIndexSpot(fs string) ([]IndexQuote, error) {
	params := map[string]string{
		"pn":    "1",
		"pz":    "100",
		"po":    "1",
		"np":    "1",
		"ut":    "bd1d9ddb04089700cf9c27f6f7426281",
		"fltt":  "2",
		"invt":  "2",
		"wbp2u": "|0|0|0|web",
		"fid":   "f12",
		"fs":    fs,
		"fields": "f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18,f20,f21," +
			"f23,f24,f25,f26,f22,f33,f11,f62,f128,f136,f115,f152",
	}

	headers := map[string]string{
		"Referer": "https://quote.eastmoney.com/center/gridlist.html#index_sz",
	}

	resp, err := utils.GetWithHeaders(EmIndexListURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求指数行情失败: %w", err)
	}

	return parseIndexSpotResponse(resp.String())
}

// parseIndexSpotResponse 解析指数行情响应
func parseIndexSpotResponse(text string) ([]IndexQuote, error) {
	result := gjson.Get(text, "data.diff")
	if !result.Exists() {
		return nil, fmt.Errorf("未找到指数数据")
	}

	quotes := make([]IndexQuote, 0)
	for _, item := range result.Array() {
		quote := IndexQuote{
			Seq:         int(item.Get("f0").Int()) + 1,
			Code:        item.Get("f12").String(),
			Name:        item.Get("f14").String(),
			Price:       item.Get("f2").Float() / 100,
			ChangePct:   item.Get("f3").Float() / 100,
			Change:      item.Get("f4").Float() / 100,
			Volume:      item.Get("f5").Int(),
			Amount:      item.Get("f6").Float() / 10000,
			Amplitude:   item.Get("f7").Float() / 100,
			High:        item.Get("f15").Float() / 100,
			Low:         item.Get("f16").Float() / 100,
			Open:        item.Get("f17").Float() / 100,
			PreClose:    item.Get("f18").Float() / 100,
			VolumeRatio: item.Get("f10").Float(),
		}

		// 过滤无效数据
		if quote.Code != "" {
			quotes = append(quotes, quote)
		}
	}

	return quotes, nil
}

// GetZHIndexPageCount 获取新浪指数总页数
//
// 返回:
//   - int: 总页数
//   - error: 错误信息
func GetZHIndexPageCount() (int, error) {
	headers := map[string]string{
		"Referer": "https://vip.stock.finance.sina.com.cn/",
	}

	resp, err := utils.GetWithHeaders(SinaIndexCountURL, nil, headers)
	if err != nil {
		return 0, fmt.Errorf("请求指数页数失败: %w", err)
	}

	// 解析页数: 从响应中提取数字
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(resp.String(), -1)
	if len(matches) == 0 {
		return 0, fmt.Errorf("未找到页数")
	}

	count, _ := strconv.Atoi(matches[0])
	pageCount := count / 80
	if count%80 != 0 {
		pageCount++
	}

	return pageCount, nil
}

// StockZHIndexSpotSina 新浪财经-中国指数实时行情
//
// 返回:
//   - []IndexQuote: 指数实时行情列表
//   - error: 错误信息
//
// 注意: 大量采集可能被目标网站服务器封禁IP
func StockZHIndexSpotSina() ([]IndexQuote, error) {
	pageCount, err := GetZHIndexPageCount()
	if err != nil {
		return nil, err
	}

	allQuotes := make([]IndexQuote, 0)

	headers := map[string]string{
		"Referer": "https://vip.stock.finance.sina.com.cn/mkt/#hs_s",
	}

	for page := 1; page <= pageCount; page++ {
		params := map[string]string{
			"page":   strconv.Itoa(page),
			"num":    "80",
			"sort":   "symbol",
			"asc":    "1",
			"node":   "hs_s",
			"_s_r_a": "page",
		}

		resp, err := utils.GetWithHeaders(SinaIndexStockURL, params, headers)
		if err != nil {
			continue
		}

		// 新浪返回的是非标准JSON，需要解析
		quotes := parseSinaIndexSpot(resp.String())
		allQuotes = append(allQuotes, quotes...)
	}

	return allQuotes, nil
}

// parseSinaIndexSpot 解析新浪指数行情响应
func parseSinaIndexSpot(text string) []IndexQuote {
	// 新浪返回格式类似: [{code:"000001", name:"上证指数", ...}]
	// 需要特殊处理
	quotes := make([]IndexQuote, 0)

	// 简单解析：提取数据项
	items := strings.Split(text, "},{")
	for _, item := range items {
		// 提取code
		codeStart := strings.Index(item, `"code":"`)
		if codeStart == -1 {
			continue
		}
		codeStart += 8
		codeEnd := strings.Index(item[codeStart:], `"`)
		if codeEnd == -1 {
			continue
		}
		code := item[codeStart : codeStart+codeEnd]

		// 提取name
		nameStart := strings.Index(item, `"name":"`)
		if nameStart == -1 {
			continue
		}
		nameStart += 8
		nameEnd := strings.Index(item[nameStart:], `"`)
		if nameEnd == -1 {
			continue
		}
		name := item[nameStart : nameStart+nameEnd]

		// 提取trade (最新价)
		tradeStart := strings.Index(item, `"trade":"`)
		var price float64
		if tradeStart != -1 {
			tradeStart += 9
			tradeEnd := strings.Index(item[tradeStart:], `"`)
			if tradeEnd != -1 {
				priceStr := strings.Replace(item[tradeStart:tradeStart+tradeEnd], ",", "", -1)
				price = utils.MustFloat64(priceStr)
			}
		}

		// 提取changeprice (涨跌额)
		changeStart := strings.Index(item, `"changeprice":"`)
		var change float64
		if changeStart != -1 {
			changeStart += 14
			changeEnd := strings.Index(item[changeStart:], `"`)
			if changeEnd != -1 {
				changeStr := strings.Replace(item[changeStart:changeStart+changeEnd], ",", "", -1)
				change = utils.MustFloat64(changeStr)
			}
		}

		// 提取changepercent (涨跌幅)
		changePctStart := strings.Index(item, `"changepercent":"`)
		var changePct float64
		if changePctStart != -1 {
			changePctStart += 17
			changePctEnd := strings.Index(item[changePctStart:], `"`)
			if changePctEnd != -1 {
				changePctStr := strings.Replace(item[changePctStart:changePctStart+changePctEnd], ",", "", -1)
				changePct = utils.MustFloat64(changePctStr)
			}
		}

		if code != "" {
			quotes = append(quotes, IndexQuote{
				Code:      code,
				Name:      name,
				Price:     price,
				Change:    change,
				ChangePct: changePct,
			})
		}
	}

	return quotes
}

// StockZHIndexDaily 东方财富网-指数日K线数据
//
// 参数:
//   - symbol: 指数代码，如 "000001", "399001", "csi931151"
//   - startDate: 开始日期，格式 "19900101"
//   - endDate: 结束日期，格式 "20500101"
//
// 返回:
//   - []IndexKLine: K线数据列表
//   - error: 错误信息
func StockZHIndexDaily(symbol, startDate, endDate string) ([]IndexKLine, error) {
	// 确定市场代码
	marketCode, code := parseSymbolCode(symbol)

	secid := fmt.Sprintf("%s.%s", marketCode, code)

	params := map[string]string{
		"secid":   secid,
		"fields1": "f1,f2,f3,f4,f5",
		"fields2": "f51,f52,f53,f54,f55,f56,f57,f58",
		"klt":     "101", // 日频率
		"fqt":     "0",
		"beg":     startDate,
		"end":     endDate,
	}

	headers := map[string]string{
		"Referer": "https://quote.eastmoney.com/center/hszs.html",
	}

	resp, err := utils.GetWithHeaders(EmIndexKLineURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求指数K线失败: %w", err)
	}

	result := gjson.Get(resp.String(), "data.klines")
	if !result.Exists() {
		return nil, fmt.Errorf("未找到K线数据")
	}

	klines := make([]IndexKLine, 0)
	for _, item := range result.Array() {
		// 数据格式: "日期,开,收,高,低,成交量,成交额,_"
		parts := strings.Split(item.String(), ",")
		if len(parts) < 7 {
			continue
		}

		date, _ := utils.ParseDate(parts[0])
		kline := IndexKLine{
			Date:   date,
			Open:   utils.MustFloat64(parts[1]),
			Close:  utils.MustFloat64(parts[2]),
			High:   utils.MustFloat64(parts[3]),
			Low:    utils.MustFloat64(parts[4]),
			Volume: utils.MustInt64(parts[5]),
			Amount: utils.MustFloat64(parts[6]),
		}
		klines = append(klines, kline)
	}

	return klines, nil
}

// parseSymbolCode 解析指数代码，返回市场代码和纯代码
func parseSymbolCode(symbol string) (string, string) {
	if strings.HasPrefix(symbol, "sz") {
		return "0", strings.TrimPrefix(symbol, "sz")
	}
	if strings.HasPrefix(symbol, "sh") {
		return "1", strings.TrimPrefix(symbol, "sh")
	}
	if strings.HasPrefix(symbol, "csi") {
		return "2", strings.TrimPrefix(symbol, "csi")
	}
	if strings.HasPrefix(symbol, "bj") {
		return "47", strings.TrimPrefix(symbol, "bj")
	}
	// 默认尝试深交所
	return "0", symbol
}
