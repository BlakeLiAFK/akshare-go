package stock

import (
	"fmt"
	"strings"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

const (
	// 东方财富港股列表 API
	emHKStockListURL = "https://push2.eastmoney.com/api/qt/clist/get"
	// 东方财富港股K线 API
	emHKKlineURL = "https://push2his.eastmoney.com/api/qt/stock/kline/get"
)

// StockHkSpotEm 获取港股实时行情（东方财富数据源）
//
// 返回:
//   - []HKStockQuote: 港股行情列表
//   - error: 错误信息
//
// 示例:
//
//	quotes, err := stock.StockHkSpotEm()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, q := range quotes[:10] {
//	    fmt.Printf("%s(%s): %.3f %.2f%%\n", q.Name, q.Code, q.Price, q.ChangePct)
//	}
func StockHkSpotEm() ([]HKStockQuote, error) {
	params := map[string]string{
		"pn":     "1",
		"pz":     "5000",
		"po":     "1",
		"np":     "1",
		"fltt":   "2",
		"invt":   "2",
		"fid":    "f3",
		"fs":     "m:116+t:1,m:116+t:2,m:116+t:3,m:116+t:4", // 港股主板
		"fields": "f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18,f20,f21,f23,f24,f25,f152",
	}

	headers := map[string]string{
		"Referer": emQuoteReferer,
	}

	resp, err := utils.GetWithHeaders(emHKStockListURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("获取港股行情失败: %w", err)
	}

	// 解析 JSON 数据
	text := resp.String()
	dataPath := "data.diff"
	result := gjson.Get(text, dataPath)
	if !result.Exists() {
		return nil, fmt.Errorf("解析港股行情数据失败: 未找到 %s", dataPath)
	}

	quotes := make([]HKStockQuote, 0, result.Get("#").Int())
	now := time.Now()

	result.ForEach(func(_, value gjson.Result) bool {
		quote := HKStockQuote{
			Code:      value.Get("f12").String(),
			Name:      value.Get("f14").String(),
			Price:     utils.MustFloat64(value.Get("f2").String()),
			ChangePct: utils.MustFloat64(value.Get("f3").String()),
			Change:    utils.MustFloat64(value.Get("f4").String()),
			Volume:    utils.MustInt64(value.Get("f5").String()),
			Amount:    utils.MustFloat64(value.Get("f6").String()),
			High:      utils.MustFloat64(value.Get("f15").String()),
			Low:       utils.MustFloat64(value.Get("f16").String()),
			Open:      utils.MustFloat64(value.Get("f17").String()),
			PreClose:  utils.MustFloat64(value.Get("f18").String()),
			Time:      now,
		}
		quotes = append(quotes, quote)
		return true
	})

	return quotes, nil
}

// StockHkDailyEm 获取港股历史K线（东方财富数据源）
//
// 参数:
//   - code: 港股代码，如 "00700"
//   - startDate: 开始日期，格式 "20240101"
//   - endDate: 结束日期，格式 "20240131"
//   - adjust: 复权类型，"qfq"(前复权), "hfq"(后复权), ""(不复权)
//
// 返回:
//   - []StockKLine: K线数据列表
//   - error: 错误信息
//
// 示例:
//
//	klines, err := stock.StockHkDailyEm("00700", "20240101", "20240131", "qfq")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, k := range klines {
//	    fmt.Printf("%s: %.3f\n", k.Date.Format("2006-01-02"), k.Close)
//	}
func StockHkDailyEm(code, startDate, endDate, adjust string) ([]StockKLine, error) {
	if code == "" {
		return nil, fmt.Errorf("股票代码不能为空")
	}

	// 港股代码转换
	secid := "116." + code

	// 复权类型映射
	fqt := "0" // 不复权
	switch adjust {
	case "qfq":
		fqt = "1" // 前复权
	case "hfq":
		fqt = "2" // 后复权
	}

	params := map[string]string{
		"secid":   secid,
		"klt":     "101", // 日K
		"fqt":     fqt,
		"beg":     startDate,
		"end":     endDate,
		"fields1": "f1,f2,f3,f4,f5,f6",
		"fields2": "f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61",
	}

	headers := map[string]string{
		"Referer": emQuoteReferer,
	}

	resp, err := utils.GetWithHeaders(emHKKlineURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("获取港股K线失败: %w", err)
	}

	// 解析 JSON 数据
	text := resp.String()
	dataPath := "data.klines"
	result := gjson.Get(text, dataPath)
	if !result.Exists() {
		return nil, fmt.Errorf("解析港股K线数据失败: 未找到 %s", dataPath)
	}

	klines := make([]StockKLine, 0)
	for _, item := range result.Array() {
		// 数据格式: "日期,开,收,高,低,成交量,成交额,振幅,涨跌幅,涨跌额,换手率"
		parts := strings.Split(item.String(), ",")
		if len(parts) < 11 {
			continue
		}

		date, _ := time.Parse("2006-01-02", parts[0])

		kline := StockKLine{
			Date:      date,
			Open:      utils.MustFloat64(parts[1]),
			Close:     utils.MustFloat64(parts[2]),
			High:      utils.MustFloat64(parts[3]),
			Low:       utils.MustFloat64(parts[4]),
			Volume:    utils.MustInt64(parts[5]),
			Amount:    utils.MustFloat64(parts[6]),
			ChangePct: utils.MustFloat64(parts[8]),
			Change:    utils.MustFloat64(parts[9]),
			Turnover:  utils.MustFloat64(parts[10]),
			Adjust:    adjust,
		}
		klines = append(klines, kline)
	}

	return klines, nil
}
