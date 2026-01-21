package fund

import (
	"fmt"
	"strings"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// FundEtfSpotEm 获取沪深京三大市场实时 ETF 行情数据
// https://quote.eastmoney.com/center/gridlist.html#fund_etf
func FundEtfSpotEm() ([]map[string]interface{}, error) {
	url := "https://push2.eastmoney.com/api/qt/clist/get"
	params := map[string]string{
		"pn":     "1",
		"pz":     "20000",
		"fs":     "b:MK0021,b:MK0022,b:MK0023",
		"fields": "f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18,f20,f21,f23,f24,f25,f22,f11,f62,f128,f136,f115,f152",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	items := result.Get("data.diff").Array()

	records := make([]map[string]interface{}, 0, len(items))
	for _, item := range items {
		record := map[string]interface{}{
			"代码":      item.Get("f12").String(),
			"名称":      item.Get("f14").String(),
			"最新价":     utils.MustFloat64(item.Get("f2").String()),
			"涨跌幅":     utils.MustFloat64(item.Get("f3").String()),
			"涨跌额":     utils.MustFloat64(item.Get("f4").String()),
			"成交量":     utils.MustFloat64(item.Get("f5").String()),
			"成交额":     utils.MustFloat64(item.Get("f6").String()),
			"振幅":      utils.MustFloat64(item.Get("f7").String()),
			"最高":      utils.MustFloat64(item.Get("f15").String()),
			"最低":      utils.MustFloat64(item.Get("f16").String()),
			"今开":      utils.MustFloat64(item.Get("f17").String()),
			"昨收":      utils.MustFloat64(item.Get("f18").String()),
			"量比":      utils.MustFloat64(item.Get("f10").String()),
			"换手率":     utils.MustFloat64(item.Get("f8").String()),
			"市盈率动态":   utils.MustFloat64(item.Get("f9").String()),
			"市净率":     utils.MustFloat64(item.Get("f23").String()),
			"总市值":     utils.MustFloat64(item.Get("f20").String()),
			"流通市值":    utils.MustFloat64(item.Get("f21").String()),
			"涨速":      utils.MustFloat64(item.Get("f22").String()),
			"5分钟涨跌":   utils.MustFloat64(item.Get("f11").String()),
			"60日涨跌幅":  utils.MustFloat64(item.Get("f24").String()),
			"年初至今涨跌幅": utils.MustFloat64(item.Get("f25").String()),
		}
		records = append(records, record)
	}

	return records, nil
}

// getMarketId 根据代码判断市场ID
func getMarketId(symbol string) string {
	if len(symbol) == 0 {
		return "1"
	}
	firstChar := symbol[0]
	if firstChar == '5' || firstChar == '6' {
		return "1" // 上海
	}
	return "0" // 深圳
}

// FundEtfHistEm 获取 ETF 历史行情数据
// https://quote.eastmoney.com/concept/sz000062.html
// 参数:
//
//	symbol: ETF代码，如 "159707"
//	period: 周期 "daily"(日), "weekly"(周), "monthly"(月)
//	startDate: 开始日期 "19700101"
//	endDate: 结束日期 "20500101"
//	adjust: 复权类型 ""(不复权), "qfq"(前复权), "hfq"(后复权)
func FundEtfHistEm(symbol, period, startDate, endDate, adjust string) ([]map[string]interface{}, error) {
	if period == "" {
		period = "daily"
	}
	if startDate == "" {
		startDate = "19700101"
	}
	if endDate == "" {
		endDate = "20500101"
	}

	periodMap := map[string]string{
		"daily":   "101",
		"weekly":  "102",
		"monthly": "103",
	}
	adjustMap := map[string]string{
		"":    "0",
		"qfq": "1",
		"hfq": "2",
	}

	marketId := getMarketId(symbol)
	url := "https://push2his.eastmoney.com/api/qt/stock/kline/get"
	params := map[string]string{
		"fields1": "f1,f2,f3,f4,f5,f6",
		"fields2": "f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61",
		"ut":      "7eea3edcaed734bea9cbfc24409ed989",
		"klt":     periodMap[period],
		"fqt":     adjustMap[adjust],
		"secid":   marketId + "." + symbol,
		"beg":     startDate,
		"end":     endDate,
		"_":       fmt.Sprintf("%d", time.Now().UnixMilli()),
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	klines := result.Get("data.klines").Array()

	records := make([]map[string]interface{}, 0, len(klines))
	for _, kline := range klines {
		parts := strings.Split(kline.String(), ",")
		if len(parts) < 11 {
			continue
		}

		record := map[string]interface{}{
			"日期":  parts[0],
			"开盘":  utils.MustFloat64(parts[1]),
			"收盘":  utils.MustFloat64(parts[2]),
			"最高":  utils.MustFloat64(parts[3]),
			"最低":  utils.MustFloat64(parts[4]),
			"成交量": utils.MustFloat64(parts[5]),
			"成交额": utils.MustFloat64(parts[6]),
			"振幅":  utils.MustFloat64(parts[7]),
			"涨跌幅": utils.MustFloat64(parts[8]),
			"涨跌额": utils.MustFloat64(parts[9]),
			"换手率": utils.MustFloat64(parts[10]),
		}
		records = append(records, record)
	}

	return records, nil
}

// FundEtfHistMinEm 获取 ETF 分钟级别行情数据
// 参数:
//
//	symbol: ETF代码
//	startDate: 开始时间 "1979-09-01 09:32:00"
//	endDate: 结束时间 "2222-01-01 09:32:00"
//	period: 分钟周期 "5", "15", "30", "60"
//	adjust: 复权类型 ""(不复权), "qfq"(前复权), "hfq"(后复权)
func FundEtfHistMinEm(symbol, startDate, endDate, period, adjust string) ([]map[string]interface{}, error) {
	if period == "" {
		period = "5"
	}
	if startDate == "" {
		startDate = "1979-09-01 09:32:00"
	}
	if endDate == "" {
		endDate = "2222-01-01 09:32:00"
	}

	periodMap := map[string]string{
		"5":  "5",
		"15": "15",
		"30": "30",
		"60": "60",
	}
	adjustMap := map[string]string{
		"":    "0",
		"qfq": "1",
		"hfq": "2",
	}

	marketId := getMarketId(symbol)
	url := "https://push2his.eastmoney.com/api/qt/stock/kline/get"
	params := map[string]string{
		"fields1": "f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f11,f12,f13",
		"fields2": "f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61",
		"ut":      "7eea3edcaed734bea9cbfc24409ed989",
		"klt":     periodMap[period],
		"fqt":     adjustMap[adjust],
		"secid":   marketId + "." + symbol,
		"beg":     "0",
		"end":     "20500000",
		"_":       fmt.Sprintf("%d", time.Now().UnixMilli()),
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	klines := result.Get("data.klines").Array()

	records := make([]map[string]interface{}, 0, len(klines))
	for _, kline := range klines {
		parts := strings.Split(kline.String(), ",")
		if len(parts) < 11 {
			continue
		}

		record := map[string]interface{}{
			"时间":  parts[0],
			"开盘":  utils.MustFloat64(parts[1]),
			"收盘":  utils.MustFloat64(parts[2]),
			"最高":  utils.MustFloat64(parts[3]),
			"最低":  utils.MustFloat64(parts[4]),
			"成交量": utils.MustFloat64(parts[5]),
			"成交额": utils.MustFloat64(parts[6]),
			"振幅":  utils.MustFloat64(parts[7]),
			"涨跌幅": utils.MustFloat64(parts[8]),
			"涨跌额": utils.MustFloat64(parts[9]),
			"换手率": utils.MustFloat64(parts[10]),
		}
		records = append(records, record)
	}

	return records, nil
}
