package stock

import (
	"fmt"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

const (
	// 东方财富 A 股列表 API
	emStockListURL = "https://push2.eastmoney.com/api/qt/clist/get"
	// 东方财富 Referer
	emQuoteReferer = "https://quote.eastmoney.com/"
)

// StockZhASpotEm 获取 A 股实时行情（东方财富数据源）
//
// 返回:
//   - []StockQuote: 股票行情列表
//   - error: 错误信息
//
// 示例:
//
//	quotes, err := stock.StockZhASpotEm()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, q := range quotes[:10] {
//	    fmt.Printf("%s(%s): %.2f %.2f%%\n", q.Name, q.Code, q.Price, q.ChangePct)
//	}
func StockZhASpotEm() ([]StockQuote, error) {
	params := map[string]string{
		"pn":     "1",
		"pz":     "10000",
		"po":     "1",
		"np":     "1",
		"fltt":   "2",
		"invt":   "2",
		"fid":    "f3",
		"fs":     "m:0+t:6,m:0+t:80,m:1+t:2,m:1+t:23,m:0+t:81+s:2048",
		"fields": "f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18,f20,f21,f23,f24,f25,f22,f11,f62,f128,f136,f115,f152",
	}

	headers := map[string]string{
		"Referer": emQuoteReferer,
	}

	resp, err := utils.GetWithHeaders(emStockListURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("获取A股行情失败: %w", err)
	}

	// 解析 JSON 数据
	text := resp.String()
	dataPath := "data.diff"
	result := gjson.Get(text, dataPath)
	if !result.Exists() {
		return nil, fmt.Errorf("解析A股行情数据失败: 未找到 %s", dataPath)
	}

	quotes := make([]StockQuote, 0, result.Get("#").Int())
	now := time.Now()

	result.ForEach(func(_, value gjson.Result) bool {
		quote := StockQuote{
			Code:         value.Get("f12").String(),
			Name:         value.Get("f14").String(),
			Price:        utils.MustFloat64(value.Get("f2").String()),
			ChangePct:    utils.MustFloat64(value.Get("f3").String()),
			Change:       utils.MustFloat64(value.Get("f4").String()),
			Volume:       utils.MustInt64(value.Get("f5").String()),
			Amount:       utils.MustFloat64(value.Get("f6").String()),
			High:         utils.MustFloat64(value.Get("f15").String()),
			Low:          utils.MustFloat64(value.Get("f16").String()),
			Open:         utils.MustFloat64(value.Get("f17").String()),
			PreClose:     utils.MustFloat64(value.Get("f18").String()),
			TurnoverRate: utils.MustFloat64(value.Get("f8").String()),
			PE:           utils.MustFloat64(value.Get("f9").String()),
			PB:           utils.MustFloat64(value.Get("f23").String()),
			MarketCap:    utils.MustFloat64(value.Get("f20").String()),
			CirculateCap: utils.MustFloat64(value.Get("f21").String()),
			Time:         now,
		}
		quotes = append(quotes, quote)
		return true
	})

	return quotes, nil
}

// StockIntradayEm 获取个股日内分时数据（东方财富数据源）
//
// 参数:
//   - code: 股票代码，如 "000001"
//
// 返回:
//   - []IntradayQuote: 分时数据列表
//   - error: 错误信息
func StockIntradayEm(code string) ([]IntradayQuote, error) {
	secid := getSecID(code)

	params := map[string]string{
		"secid":  secid,
		"fields": "f51,f52,f53,f54,f55,f56,f57,f58",
		"ndays":  "1",
	}

	headers := map[string]string{
		"Referer": emQuoteReferer,
	}

	resp, err := utils.GetWithHeaders("https://push2.eastmoney.com/api/qt/stock/trends2/get", params, headers)
	if err != nil {
		return nil, fmt.Errorf("获取分时数据失败: %w", err)
	}

	text := resp.String()
	result := gjson.Get(text, "data.trends")
	if !result.Exists() {
		return nil, fmt.Errorf("解析分时数据失败")
	}

	quotes := make([]IntradayQuote, 0)
	result.ForEach(func(_, value gjson.Result) bool {
		parts := splitString(value.String(), ",")
		if len(parts) >= 6 {
			quote := IntradayQuote{
				Time:     extractTime(parts[0]),
				Price:    utils.MustFloat64(parts[2]),
				Volume:   utils.MustInt64(parts[5]),
				AvgPrice: utils.MustFloat64(parts[7]),
			}
			quotes = append(quotes, quote)
		}
		return true
	})

	return quotes, nil
}

// getSecID 转换股票代码为东方财富格式
func getSecID(code string) string {
	if len(code) < 1 {
		return "0." + code
	}
	// 沪市: 6开头, 5开头(基金), 9开头(B股)
	// 深市: 0开头, 3开头, 2开头(B股)
	// 北交所: 8开头, 4开头
	first := code[0]
	switch first {
	case '6', '5', '9':
		return "1." + code // 沪市
	case '0', '3', '2':
		return "0." + code // 深市
	case '8', '4':
		return "0." + code // 北交所
	default:
		return "0." + code
	}
}

// splitString 分割字符串
func splitString(s, sep string) []string {
	if s == "" {
		return nil
	}
	var result []string
	start := 0
	for i := 0; i < len(s); i++ {
		if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
			result = append(result, s[start:i])
			start = i + len(sep)
			i += len(sep) - 1
		}
	}
	result = append(result, s[start:])
	return result
}

// extractTime 从时间字符串提取时间部分 "2024-01-15 09:30" -> "09:30"
func extractTime(s string) string {
	if len(s) >= 16 {
		return s[11:16]
	}
	return s
}
