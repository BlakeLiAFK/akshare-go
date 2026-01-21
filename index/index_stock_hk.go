package index

import (
	"fmt"
	"strings"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// StockHKIndexSpotSina 新浪财经-港股指数实时行情
//
// 返回:
//   - []HKIndexQuote: 港股指数实时行情列表
//   - error: 错误信息
//
// 注意: 大量采集可能被目标网站服务器封禁IP
func StockHKIndexSpotSina() ([]HKIndexQuote, error) {
	// 构建指数代码列表
	codeList := strings.Join(HKIndexList, ",")
	url := fmt.Sprintf(HKIndexSinaURL, codeList)

	headers := map[string]string{
		"Referer": "https://vip.stock.finance.sina.com.cn/",
	}

	resp, err := utils.GetWithHeaders(url, nil, headers)
	if err != nil {
		return nil, fmt.Errorf("请求港股指数行情失败: %w", err)
	}

	return parseHKIndexSinaResponse(resp.String())
}

// parseHKIndexSinaResponse 解析新浪港股指数响应
func parseHKIndexSinaResponse(text string) ([]HKIndexQuote, error) {
	// 新浪港股数据格式: hq_str_hkCES100="CES100,恒生香港中资企业指数,..."
	lines := strings.Split(text, "\n")
	quotes := make([]HKIndexQuote, 0)

	for _, line := range lines {
		if !strings.Contains(line, "=") {
			continue
		}

		// 提取数据部分
		parts := strings.SplitN(line, "=", 2)
		if len(parts) < 2 {
			continue
		}

		// 移除引号
		dataStr := strings.Trim(parts[1], `"`)
		if dataStr == "" {
			continue
		}

		// 解析数据: 代码,名称,今开,昨收,最高,最低,最新价,涨跌额,涨跌幅,...
		fields := strings.Split(dataStr, ",")
		if len(fields) < 9 {
			continue
		}

		quote := HKIndexQuote{
			Code:      fields[0],
			Name:      fields[1],
			Open:      utils.MustFloat64(fields[2]),
			PreClose:  utils.MustFloat64(fields[3]),
			High:      utils.MustFloat64(fields[4]),
			Low:       utils.MustFloat64(fields[5]),
			Price:     utils.MustFloat64(fields[6]),
			Change:    utils.MustFloat64(fields[7]),
			ChangePct: utils.MustFloat64(fields[8]),
		}
		quotes = append(quotes, quote)
	}

	return quotes, nil
}

// StockHKIndexSpotEM 东方财富网-港股指数实时行情
//
// 返回:
//   - []HKIndexQuote: 港股指数实时行情列表
//   - error: 错误信息
func StockHKIndexSpotEM() ([]HKIndexQuote, error) {
	params := map[string]string{
		"pn":    "1",
		"pz":    "100",
		"po":    "1",
		"np":    "1",
		"ut":    "bd1d9ddb04089700cf9c27f6f7426281",
		"fltt":  "2",
		"invt":  "2",
		"wbp2u": "|0|0|0|web",
		"fid":   "f3",
		"fs":    "m:124,m:125,m:305",
		"fields": "f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18,f20,f21," +
			"f23,f24,f25,f26,f22,f33,f11,f62,f128,f136,f115,f152",
	}

	headers := map[string]string{
		"Referer": "https://quote.eastmoney.com/center/gridlist.html#hk_index",
	}

	resp, err := utils.GetWithHeaders(EmIndexListURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求港股指数行情失败: %w", err)
	}

	return parseHKIndexEMResponse(resp.String())
}

// parseHKIndexEMResponse 解析东财港股指数响应
func parseHKIndexEMResponse(text string) ([]HKIndexQuote, error) {
	result := gjson.Get(text, "data.diff")
	if !result.Exists() {
		return nil, fmt.Errorf("未找到港股指数数据")
	}

	quotes := make([]HKIndexQuote, 0)
	for _, item := range result.Array() {
		quote := HKIndexQuote{
			Seq:       int(item.Get("f0").Int()) + 1,
			MarketID:  item.Get("f13").String(),
			Code:      item.Get("f12").String(),
			Name:      item.Get("f14").String(),
			Price:     item.Get("f2").Float() / 100,
			ChangePct: item.Get("f3").Float() / 100,
			Change:    item.Get("f4").Float() / 100,
			Open:      item.Get("f17").Float() / 100,
			High:      item.Get("f15").Float() / 100,
			Low:       item.Get("f16").Float() / 100,
			PreClose:  item.Get("f18").Float() / 100,
			Volume:    item.Get("f5").Int(),
			Amount:    item.Get("f6").Float() / 10000,
		}
		quotes = append(quotes, quote)
	}

	return quotes, nil
}

// StockHKIndexDailyEM 东方财富网-港股指数历史行情
//
// 参数:
//   - symbol: 港股指数代码，如 "HSTECH", "HSTECF2L"
//   - 可以通过 StockHKIndexSpotEM() 获取所有可用代码
//
// 返回:
//   - []GlobalIndexKLine: 港股指数K线数据
//   - error: 错误信息
//
// 示例:
//
//	// 获取恒生科技指数历史行情
//	klines, err := index.StockHKIndexDailyEM("HSTECH")
//	if err != nil {
//	    log.Fatal(err)
//	}
func StockHKIndexDailyEM(symbol string) ([]GlobalIndexKLine, error) {
	// 获取代码映射
	codeMap, err := getHKSymbolCodeMap()
	if err != nil {
		return nil, err
	}

	marketID, ok := codeMap[symbol]
	if !ok {
		// 默认使用市场100
		marketID = "100"
	}

	secid := fmt.Sprintf("%s.%s", marketID, symbol)

	params := map[string]string{
		"secid":   secid,
		"klt":     "101",
		"fqt":     "1",
		"lmt":     "10000",
		"end":     "20500000",
		"iscca":   "1",
		"fields1": "f1,f2,f3,f4,f5,f6,f7,f8",
		"fields2": "f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61,f62,f63,f64",
		"ut":      "f057cbcbce2a86e2866ab8877db1d059",
		"forcect": "1",
	}

	headers := map[string]string{
		"Referer": "https://quote.eastmoney.com/gb/zsHSTECF2L.html",
	}

	resp, err := utils.GetWithHeaders(EmIndexKLineURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求港股指数历史数据失败: %w", err)
	}

	klinesData := gjson.Get(resp.String(), "data.klines")
	if !klinesData.Exists() {
		return nil, fmt.Errorf("未找到历史数据")
	}

	klines := make([]GlobalIndexKLine, 0)
	for _, item := range klinesData.Array() {
		// 数据格式: "日期,开,最新价,高,低,-,-,-,-,-,-,-,-"
		parts := strings.Split(item.String(), ",")
		if len(parts) < 5 {
			continue
		}

		date, _ := time.Parse("2006-01-02", parts[0])

		kline := GlobalIndexKLine{
			Date:  date,
			Open:  utils.MustFloat64(parts[1]),
			Close: utils.MustFloat64(parts[2]),
			High:  utils.MustFloat64(parts[3]),
			Low:   utils.MustFloat64(parts[4]),
		}
		klines = append(klines, kline)
	}

	return klines, nil
}

// getHKSymbolCodeMap 获取港股指数代码映射表
func getHKSymbolCodeMap() (map[string]string, error) {
	// 先获取实时行情来构建映射表
	quotes, err := StockHKIndexSpotEM()
	if err != nil {
		return nil, err
	}

	codeMap := make(map[string]string)
	codeMap["HSAHP"] = "100" // 特殊处理

	for _, quote := range quotes {
		if quote.MarketID != "" {
			codeMap[quote.Code] = quote.MarketID
		}
	}

	return codeMap, nil
}

// StockHKIndexDailySina 新浪财经-港股指数历史行情
//
// 参数:
//   - symbol: 港股指数代码，如 "CES100", "HSCI"
//
// 返回:
//   - []GlobalIndexKLine: 港股指数K线数据
//   - error: 错误信息
func StockHKIndexDailySina(symbol string) ([]GlobalIndexKLine, error) {
	url := fmt.Sprintf("https://finance.sina.com.cn/stock/hkstock/%s/klc_kl.js", symbol)

	params := map[string]string{
		"d": "2023_5_01",
	}

	headers := map[string]string{
		"Referer": "https://stock.finance.sina.com.cn/hkstock/quotes/" + symbol + ".html",
	}

	_, err := utils.GetWithHeaders(url, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求港股指数历史数据失败: %w", err)
	}

	// 新浪港股数据需要JS解密，这里暂未实现
	// 返回空数据
	return []GlobalIndexKLine{}, fmt.Errorf("新浪港股指数历史数据需要JS解密，暂未实现")
}
