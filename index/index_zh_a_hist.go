package index

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

var (
	codeIDMap sync.Map // 指数代码ID缓存
)

// IndexCodeIDMapEM 获取东方财富指数代码ID映射
// 使用缓存提高性能
func IndexCodeIDMapEM() (map[string]string, error) {
	// 尝试从缓存获取
	if m, ok := codeIDMap.Load("map"); ok {
		return m.(map[string]string), nil
	}

	url := EmIndexListURL
	params := map[string]string{
		"pn":     "1",
		"pz":     "100",
		"po":     "1",
		"np":     "1",
		"ut":     "bd1d9ddb04089700cf9c27f6f7426281",
		"fltt":   "2",
		"invt":   "2",
		"fid":    "f3",
		"fs":     "b:MK0010,m:1+t:1,m:0 t:5,m:1+s:3,m:0+t:5,m:2",
		"fields": "f3,f12,f13",
	}

	headers := map[string]string{
		"Referer": "https://finance.eastmoney.com/",
	}

	resp, err := utils.GetWithHeaders(url, params, headers)
	if err != nil {
		return nil, fmt.Errorf("获取指数代码映射失败: %w", err)
	}

	// 解析所有分页数据
	result := gjson.Get(resp.String(), "data.diff")
	codeMap := make(map[string]string)

	for _, item := range result.Array() {
		code := item.Get("f12").String()
		id := item.Get("f13").String()
		if code != "" && id != "" {
			codeMap[code] = id
		}
	}

	// 存入缓存
	codeIDMap.Store("map", codeMap)
	return codeMap, nil
}

// IndexZHAHist 获取中国股票指数历史行情数据
//
// 参数:
//   - symbol: 指数代码，如 "000001"(上证指数), "399001"(深证成指)
//   - period: 周期，"daily"(日K), "weekly"(周K), "monthly"(月K)
//   - startDate: 开始日期，格式 "20240101"
//   - endDate: 结束日期，格式 "20240131"
//
// 返回:
//   - []IndexKLine: K线数据列表
//   - error: 错误信息
//
// 示例:
//
//	// 获取上证指数2024年1月日K线数据
//	klines, err := index.IndexZHAHist("000001", "daily", "20240101", "20240131")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, kline := range klines {
//	    fmt.Printf("%s: 开盘=%.2f, 收盘=%.2f, 涨跌幅=%.2f%%\n",
//	        kline.Date.Format("2006-01-02"), kline.Open, kline.Close, kline.ChangePct)
//	}
func IndexZHAHist(symbol, period, startDate, endDate string) ([]IndexKLine, error) {
	// 获取指数代码ID映射
	codeMap, err := IndexCodeIDMapEM()
	if err != nil {
		return nil, err
	}

	// 周期映射
	klt, ok := PeriodMap[period]
	if !ok {
		klt = "101" // 默认日K
	}

	// 尝试从映射表获取市场ID
	var secid string
	if id, ok := codeMap[symbol]; ok {
		secid = fmt.Sprintf("%s.%s", id, symbol)
	} else {
		// 尝试常见市场
		markets := []string{"1", "0", "2", "47"}
		for _, market := range markets {
			secid = fmt.Sprintf("%s.%s", market, symbol)
			if klines, err := fetchKLines(secid, klt); err == nil && len(klines) > 0 {
				return filterByDate(klines, startDate, endDate), nil
			}
		}
		return nil, fmt.Errorf("未找到指数: %s", symbol)
	}

	// 获取K线数据
	klines, err := fetchKLines(secid, klt)
	if err != nil {
		return nil, err
	}

	return filterByDate(klines, startDate, endDate), nil
}

// fetchKLines 获取K线数据
func fetchKLines(secid, klt string) ([]IndexKLine, error) {
	params := map[string]string{
		"secid":   secid,
		"ut":      "7eea3edcaed734bea9cbfc24409ed989",
		"fields1": "f1,f2,f3,f4,f5,f6",
		"fields2": "f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61",
		"klt":     klt,
		"fqt":     "0",
		"beg":     "0",
		"end":     "20500000",
	}

	headers := map[string]string{
		"Referer": "https://finance.eastmoney.com/",
	}

	resp, err := utils.GetWithHeaders(EmIndexKLineURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求K线数据失败: %w", err)
	}

	// 解析JSON响应
	result := gjson.Get(resp.String(), "data.klines")
	if !result.Exists() {
		return nil, fmt.Errorf("未找到K线数据")
	}

	klines := make([]IndexKLine, 0)
	for _, item := range result.Array() {
		// 数据格式: "日期,开,收,高,低,成交量,成交额,振幅,涨跌幅,涨跌额,换手率"
		parts := strings.Split(item.String(), ",")
		if len(parts) < 11 {
			continue
		}

		date, _ := time.Parse("2006-01-02", parts[0])
		kline := IndexKLine{
			Date:      date,
			Open:      utils.MustFloat64(parts[1]),
			Close:     utils.MustFloat64(parts[2]),
			High:      utils.MustFloat64(parts[3]),
			Low:       utils.MustFloat64(parts[4]),
			Volume:    utils.MustInt64(parts[5]),
			Amount:    utils.MustFloat64(parts[6]),
			Amplitude: utils.MustFloat64(parts[7]),
			ChangePct: utils.MustFloat64(parts[8]),
			Change:    utils.MustFloat64(parts[9]),
			Turnover:  utils.MustFloat64(parts[10]),
		}
		klines = append(klines, kline)
	}

	return klines, nil
}

// filterByDate 按日期范围过滤K线数据
func filterByDate(klines []IndexKLine, startDate, endDate string) []IndexKLine {
	if startDate == "" && endDate == "" {
		return klines
	}

	var start, end time.Time
	var startErr, endErr error

	if startDate != "" {
		start, startErr = time.Parse("20060102", startDate)
	}
	if endDate != "" {
		end, endErr = time.Parse("20060102", endDate)
	}

	if startErr != nil && endErr != nil {
		return klines
	}

	filtered := make([]IndexKLine, 0)
	for _, kline := range klines {
		include := true
		if !start.IsZero() && kline.Date.Before(start) {
			include = false
		}
		if !end.IsZero() && kline.Date.After(end) {
			include = false
		}
		if include {
			filtered = append(filtered, kline)
		}
	}

	return filtered
}

// IndexZHAHistMinEM 获取中国指数分时行情数据
//
// 参数:
//   - symbol: 指数代码，如 "000001", "399006"
//   - period: 周期，"1"(1分钟), "5"(5分钟), "15"(15分钟), "30"(30分钟), "60"(60分钟)
//   - startDate: 开始日期时间，格式 "2024-01-01 09:30:00"
//   - endDate: 结束日期时间，格式 "2024-01-01 15:00:00"
//
// 返回:
//   - []IndexIntradayQuote: 分时数据列表
//   - error: 错误信息
func IndexZHAHistMinEM(symbol, period, startDate, endDate string) ([]IndexIntradayQuote, error) {
	// 获取指数代码ID映射
	codeMap, err := IndexCodeIDMapEM()
	if err != nil {
		return nil, err
	}

	// 1分钟分时使用特殊接口
	if period == "1" {
		return fetchIntradayData(symbol, codeMap, startDate, endDate)
	}

	// 其他周期使用K线接口
	return fetchMinKLineData(symbol, codeMap, period, startDate, endDate)
}

// fetchIntradayData 获取1分钟分时数据
func fetchIntradayData(symbol string, codeMap map[string]string, startDate, endDate string) ([]IndexIntradayQuote, error) {
	var secid string
	if id, ok := codeMap[symbol]; ok {
		secid = fmt.Sprintf("%s.%s", id, symbol)
	} else {
		// 尝试常见市场
		markets := []string{"1", "0", "2", "47"}
		for _, market := range markets {
			secid = fmt.Sprintf("%s.%s", market, symbol)
			if quotes, err := fetchTrends(secid); err == nil && len(quotes) > 0 {
				return filterIntradayByDate(quotes, startDate, endDate), nil
			}
		}
		return nil, fmt.Errorf("未找到指数: %s", symbol)
	}

	quotes, err := fetchTrends(secid)
	if err != nil {
		return nil, err
	}

	return filterIntradayByDate(quotes, startDate, endDate), nil
}

// fetchTrends 获取分时走势数据
func fetchTrends(secid string) ([]IndexIntradayQuote, error) {
	params := map[string]string{
		"fields1": "f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f11,f12,f13",
		"fields2": "f51,f52,f53,f54,f55,f56,f57,f58",
		"iscr":    "0",
		"ndays":   "5",
		"secid":   secid,
	}

	headers := map[string]string{
		"Referer": "https://quote.eastmoney.com/",
	}

	resp, err := utils.GetWithHeaders(EmIndexTrendsURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求分时数据失败: %w", err)
	}

	result := gjson.Get(resp.String(), "data.trends")
	if !result.Exists() {
		return nil, fmt.Errorf("未找到分时数据")
	}

	quotes := make([]IndexIntradayQuote, 0)
	for _, item := range result.Array() {
		// 数据格式: "时间,开,收,高,低,成交量,成交额,均价"
		parts := strings.Split(item.String(), ",")
		if len(parts) < 8 {
			continue
		}

		quote := IndexIntradayQuote{
			Time:   parts[0],
			Open:   utils.MustFloat64(parts[1]),
			Close:  utils.MustFloat64(parts[2]),
			High:   utils.MustFloat64(parts[3]),
			Low:    utils.MustFloat64(parts[4]),
			Volume: utils.MustInt64(parts[5]),
			Amount: utils.MustFloat64(parts[6]),
			Avg:    utils.MustFloat64(parts[7]),
		}
		quotes = append(quotes, quote)
	}

	return quotes, nil
}

// fetchMinKLineData 获取分钟级K线数据
func fetchMinKLineData(symbol string, codeMap map[string]string, period, startDate, endDate string) ([]IndexIntradayQuote, error) {
	var secid string
	if id, ok := codeMap[symbol]; ok {
		secid = fmt.Sprintf("%s.%s", id, symbol)
	} else {
		// 尝试深交所
		secid = fmt.Sprintf("0.%s", symbol)
	}

	params := map[string]string{
		"secid":   secid,
		"ut":      "7eea3edcaed734bea9cbfc24409ed989",
		"fields1": "f1,f2,f3,f4,f5,f6",
		"fields2": "f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61",
		"klt":     period,
		"fqt":     "1",
		"beg":     "0",
		"end":     "20500000",
	}

	headers := map[string]string{
		"Referer": "https://quote.eastmoney.com/",
	}

	resp, err := utils.GetWithHeaders(EmIndexKLineURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求分钟K线数据失败: %w", err)
	}

	result := gjson.Get(resp.String(), "data.klines")
	if !result.Exists() {
		return nil, fmt.Errorf("未找到分钟K线数据")
	}

	quotes := make([]IndexIntradayQuote, 0)
	for _, item := range result.Array() {
		// 数据格式: "时间,开,收,高,低,成交量,成交额,振幅,涨跌幅,涨跌额,换手率"
		parts := strings.Split(item.String(), ",")
		if len(parts) < 11 {
			continue
		}

		quote := IndexIntradayQuote{
			Time:   parts[0],
			Open:   utils.MustFloat64(parts[1]),
			Close:  utils.MustFloat64(parts[2]),
			High:   utils.MustFloat64(parts[3]),
			Low:    utils.MustFloat64(parts[4]),
			Volume: utils.MustInt64(parts[5]),
			Amount: utils.MustFloat64(parts[6]),
		}
		quotes = append(quotes, quote)
	}

	return filterIntradayByDate(quotes, startDate, endDate), nil
}

// filterIntradayByDate 按日期时间范围过滤分时数据
func filterIntradayByDate(quotes []IndexIntradayQuote, startDate, endDate string) []IndexIntradayQuote {
	if startDate == "" && endDate == "" {
		return quotes
	}

	var start, end time.Time
	var startErr, endErr error

	layout := "2006-01-02 15:04:05"
	if startDate != "" {
		start, startErr = time.Parse(layout, startDate)
	}
	if endDate != "" {
		end, endErr = time.Parse(layout, endDate)
	}

	if startErr != nil && endErr != nil {
		return quotes
	}

	filtered := make([]IndexIntradayQuote, 0)
	for _, quote := range quotes {
		// 尝试解析时间
		t, err := time.Parse(layout, quote.Time)
		if err != nil {
			continue
		}

		include := true
		if !start.IsZero() && t.Before(start) {
			include = false
		}
		if !end.IsZero() && t.After(end) {
			include = false
		}
		if include {
			filtered = append(filtered, quote)
		}
	}

	return filtered
}
