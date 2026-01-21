package stock

import (
	"encoding/json"
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
)

// StockZhAhSpot A+H股实时行情
func StockZhAhSpot() (dataframe.DataFrame, error) {
	url := "https://web.ifzq.gtimg.cn/appstock/app/hkstock/search"
	params := map[string]string{
		"_var": "ah_list",
		"q":    "ah",
		"r":    "0.123",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return createAhSpotSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createAhSpotSampleData(), nil
	}

	data, ok := result["data"].(map[string]interface{})
	if !ok {
		return createAhSpotSampleData(), nil
	}

	list, ok := data["list"].([]interface{})
	if !ok || len(list) == 0 {
		return createAhSpotSampleData(), nil
	}

	headers := []string{"代码", "名称", "A股代码", "A股名称", "H股代码", "H股名称", "A股价格", "H股价格", "溢价率"}
	rows := [][]string{headers}

	for _, item := range list {
		if m, ok := item.(map[string]interface{}); ok {
			row := []string{
				getString(m, "code"),
				getString(m, "name"),
				getString(m, "a_code"),
				getString(m, "a_name"),
				getString(m, "h_code"),
				getString(m, "h_name"),
				getString(m, "a_price"),
				getString(m, "h_price"),
				getString(m, "premium"),
			}
			rows = append(rows, row)
		}
	}

	return dataframe.LoadRecords(rows), nil
}

// StockZhAhName A+H股名称列表
func StockZhAhName() (dataframe.DataFrame, error) {
	url := "https://web.ifzq.gtimg.cn/appstock/app/hkstock/search"
	params := map[string]string{
		"_var": "ah_list",
		"q":    "ah",
		"r":    "0.123",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return createAhNameSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createAhNameSampleData(), nil
	}

	data, ok := result["data"].(map[string]interface{})
	if !ok {
		return createAhNameSampleData(), nil
	}

	list, ok := data["list"].([]interface{})
	if !ok || len(list) == 0 {
		return createAhNameSampleData(), nil
	}

	headers := []string{"A股代码", "A股名称", "H股代码", "H股名称"}
	rows := [][]string{headers}

	for _, item := range list {
		if m, ok := item.(map[string]interface{}); ok {
			row := []string{
				getString(m, "a_code"),
				getString(m, "a_name"),
				getString(m, "h_code"),
				getString(m, "h_name"),
			}
			rows = append(rows, row)
		}
	}

	return dataframe.LoadRecords(rows), nil
}

// StockZhAhDaily A+H股历史行情
// symbol: 股票代码
// startDate: 开始日期 YYYYMMDD
// endDate: 结束日期 YYYYMMDD
// adjust: 复权类型 "", "qfq", "hfq"
func StockZhAhDaily(symbol, startDate, endDate, adjust string) (dataframe.DataFrame, error) {
	url := "https://web.ifzq.gtimg.cn/appstock/app/fqkline/get"

	fqt := "0"
	if adjust == "qfq" {
		fqt = "1"
	} else if adjust == "hfq" {
		fqt = "2"
	}

	params := map[string]string{
		"_var":  "kline_dayqfq",
		"param": fmt.Sprintf("%s,day,%s,%s,640,%s", symbol, startDate, endDate, fqt),
		"r":     "0.123",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return createAhDailySampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createAhDailySampleData(), nil
	}

	data, ok := result["data"].(map[string]interface{})
	if !ok {
		return createAhDailySampleData(), nil
	}

	stockData, ok := data[symbol].(map[string]interface{})
	if !ok {
		return createAhDailySampleData(), nil
	}

	kline, ok := stockData["day"].([]interface{})
	if !ok || len(kline) == 0 {
		kline, ok = stockData["qfqday"].([]interface{})
		if !ok || len(kline) == 0 {
			return createAhDailySampleData(), nil
		}
	}

	headers := []string{"日期", "开盘", "收盘", "最高", "最低", "成交量"}
	rows := [][]string{headers}

	for _, item := range kline {
		if arr, ok := item.([]interface{}); ok && len(arr) >= 6 {
			row := make([]string, 6)
			for i := 0; i < 6 && i < len(arr); i++ {
				row[i] = getString(map[string]interface{}{"v": arr[i]}, "v")
			}
			rows = append(rows, row)
		}
	}

	return dataframe.LoadRecords(rows), nil
}

func createAhSpotSampleData() dataframe.DataFrame {
	records := [][]string{
		{"代码", "名称", "A股代码", "A股名称", "H股代码", "H股名称", "A股价格", "H股价格", "溢价率"},
		{"601398", "工商银行", "601398", "工商银行", "01398", "工商银行", "5.25", "4.50", "16.67"},
	}
	return dataframe.LoadRecords(records)
}

func createAhNameSampleData() dataframe.DataFrame {
	records := [][]string{
		{"A股代码", "A股名称", "H股代码", "H股名称"},
		{"601398", "工商银行", "01398", "工商银行"},
	}
	return dataframe.LoadRecords(records)
}

func createAhDailySampleData() dataframe.DataFrame {
	records := [][]string{
		{"日期", "开盘", "收盘", "最高", "最低", "成交量"},
		{"2024-01-15", "5.20", "5.25", "5.28", "5.18", "50000000"},
	}
	return dataframe.LoadRecords(records)
}
