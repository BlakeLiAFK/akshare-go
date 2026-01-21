package stock

import (
	"fmt"
	"strings"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// StockZhAHist 获取 A 股历史 K 线数据
//
// 参数:
//   - code: 股票代码，如 "000001"
//   - period: 周期，"daily"(日K), "weekly"(周K), "monthly"(月K)
//   - startDate: 开始日期，格式 "20240101"
//   - endDate: 结束日期，格式 "20240115"
//   - adjust: 复权类型，"qfq"(前复权), "hfq"(后复权), ""(不复权)
//
// 返回:
//   - []StockKLine: K线数据列表
//   - error: 错误信息
//
// 示例:
//
//	// 获取平安银行2024年1月日K线数据（前复权）
//	klines, err := StockZhAHist("000001", "daily", "20240101", "20240131", "qfq")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, kline := range klines {
//	    fmt.Printf("%s: 开盘=%.2f, 收盘=%.2f, 涨跌幅=%.2f%%\n",
//	        kline.Date.Format("2006-01-02"), kline.Open, kline.Close, kline.ChangePct)
//	}
func StockZhAHist(code, period, startDate, endDate, adjust string) ([]StockKLine, error) {
	// 市场代码转换（东方财富格式）
	secid := getSecID(code)

	// 周期映射
	klt := "101" // 日K
	switch period {
	case "weekly":
		klt = "102" // 周K
	case "monthly":
		klt = "103" // 月K
	}

	// 复权类型映射
	fqt := "0" // 不复权
	switch adjust {
	case "qfq":
		fqt = "1" // 前复权
	case "hfq":
		fqt = "2" // 后复权
	}

	// 构建请求参数
	params := map[string]string{
		"secid":   secid,
		"klt":     klt,
		"fqt":     fqt,
		"beg":     startDate,
		"end":     endDate,
		"fields1": "f1,f2,f3,f4,f5,f6",
		"fields2": "f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61",
	}

	// 发送HTTP请求
	headers := map[string]string{
		"Referer": "https://finance.eastmoney.com/",
	}

	resp, err := utils.GetWithHeaders("https://push2his.eastmoney.com/api/qt/stock/kline/get", params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求K线数据失败: %w", err)
	}

	// 解析JSON响应
	result := gjson.Get(resp.String(), "data.klines")
	if !result.Exists() {
		return nil, fmt.Errorf("未找到K线数据")
	}

	// 将JSON数据转换为结构体切片
	klines := make([]StockKLine, 0)
	for _, item := range result.Array() {
		// 数据格式: "日期,开,收,高,低,成交量,成交额,振幅,涨跌幅,涨跌额,换手率"
		parts := strings.Split(item.String(), ",")
		if len(parts) < 11 {
			continue
		}

		// 解析日期
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
