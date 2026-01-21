package futures

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
)

// FuturesForeign 外盘期货数据
type FuturesForeign struct {
	Date   time.Time `json:"date"`   // 日期
	Open   float64   `json:"open"`   // 开盘
	High   float64   `json:"high"`   // 最高
	Low    float64   `json:"low"`    // 最低
	Close  float64   `json:"close"`  // 收盘
	Volume int64     `json:"volume"` // 成交量
	Hold   int64     `json:"hold"`   // 持仓量
}

// foreignDailyResponse 外盘期货日线API响应
type foreignDailyResponse struct {
	Data struct {
		Code   string   `json:"code"`
		Klines []string `json:"klines"`
	} `json:"data"`
}

// FuturesForeignHistEM 东方财富-外盘期货历史数据
//
// 数据源: https://quote.eastmoney.com/qihuo/CL.html
//
// 参数:
//   - symbol: 期货代码，如 "CL" (原油), "GC" (黄金), "SI" (白银)
//   - period: 周期，可选 "daily", "weekly", "monthly"
//   - startDate: 开始日期，格式 "19900101"
//   - endDate: 结束日期，格式 "20500101"
//
// 返回:
//   - []FuturesForeign: 外盘期货历史数据
//   - error: 错误信息
func FuturesForeignHistEM(symbol, period, startDate, endDate string) ([]FuturesForeign, error) {
	periodMap := map[string]string{
		"daily":   "101",
		"weekly":  "102",
		"monthly": "103",
	}

	klt, ok := periodMap[period]
	if !ok {
		return nil, fmt.Errorf("无效的 period: %s, 可选 daily, weekly, monthly", period)
	}

	// 外盘期货使用特定的市场代码
	// 常见外盘期货代码映射
	marketMap := map[string]string{
		"CL":  "101", // NYMEX原油
		"GC":  "101", // COMEX黄金
		"SI":  "101", // COMEX白银
		"HG":  "101", // COMEX铜
		"NG":  "101", // NYMEX天然气
		"OIL": "101", // 布伦特原油
	}

	marketCode := marketMap[strings.ToUpper(symbol)]
	if marketCode == "" {
		marketCode = "101" // 默认市场代码
	}

	secid := fmt.Sprintf("%s.%s", marketCode, symbol)

	url := "https://push2his.eastmoney.com/api/qt/stock/kline/get"
	params := map[string]string{
		"secid":   secid,
		"klt":     klt,
		"fqt":     "1",
		"lmt":     "10000",
		"end":     "20500000",
		"iscca":   "1",
		"fields1": "f1,f2,f3,f4,f5,f6,f7,f8",
		"fields2": "f51,f52,f53,f54,f55,f56,f57,f58",
		"ut":      "7eea3edcaed734bea9cbfc24409ed989",
		"forcect": "1",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求外盘期货历史数据失败: %w", err)
	}

	var apiResp foreignDailyResponse
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, fmt.Errorf("解析外盘期货历史数据响应失败: %w", err)
	}

	if len(apiResp.Data.Klines) == 0 {
		return []FuturesForeign{}, nil
	}

	// 解析日期范围
	var start, end time.Time
	if startDate != "" {
		start, _ = time.Parse("20060102", startDate)
	}
	if endDate != "" {
		end, _ = time.Parse("20060102", endDate)
	}

	result := make([]FuturesForeign, 0, len(apiResp.Data.Klines))
	for _, line := range apiResp.Data.Klines {
		parts := strings.Split(line, ",")
		if len(parts) < 7 {
			continue
		}

		date, err := time.Parse("2006-01-02", parts[0])
		if err != nil {
			continue
		}

		// 日期范围过滤
		if !start.IsZero() && date.Before(start) {
			continue
		}
		if !end.IsZero() && date.After(end) {
			continue
		}

		result = append(result, FuturesForeign{
			Date:   date,
			Open:   utils.MustParseFloat(parts[1]),
			Close:  utils.MustParseFloat(parts[2]),
			High:   utils.MustParseFloat(parts[3]),
			Low:    utils.MustParseFloat(parts[4]),
			Volume: int64(utils.MustParseFloat(parts[5])),
			Hold:   int64(utils.MustParseFloat(parts[6])),
		})
	}

	return result, nil
}
