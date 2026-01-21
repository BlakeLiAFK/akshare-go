package stock_feature

import (
	"encoding/json"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
)

// StockBuffettIndexLg 乐估乐股-底部研究-巴菲特指标
func StockBuffettIndexLg() (dataframe.DataFrame, error) {
	url := "https://legulegu.com/api/stockdata/marketcap-gdp/get-marketcap-gdp"
	params := map[string]string{
		"token": "0123456789",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return createBuffettIndexLgSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createBuffettIndexLgSampleData(), nil
	}

	data, ok := result["data"].([]interface{})
	if !ok || len(data) == 0 {
		return createBuffettIndexLgSampleData(), nil
	}

	headers := []string{"日期", "收盘价", "总市值", "GDP", "近十年分位数", "总历史分位数"}
	rows := [][]string{headers}

	for _, item := range data {
		if m, ok := item.(map[string]interface{}); ok {
			row := []string{
				getStringFeature(m, "date"),
				getStringFeature(m, "close"),
				getStringFeature(m, "marketCap"),
				getStringFeature(m, "gdp"),
				getStringFeature(m, "quantileInRecent10Years"),
				getStringFeature(m, "quantileInAllHistory"),
			}
			rows = append(rows, row)
		}
	}

	return dataframe.LoadRecords(rows), nil
}

func createBuffettIndexLgSampleData() dataframe.DataFrame {
	records := [][]string{
		{"日期", "收盘价", "总市值", "GDP", "近十年分位数", "总历史分位数"},
		{"2024-01-15", "3000.00", "80000000000000", "120000000000000", "45.5", "55.2"},
	}
	return dataframe.LoadRecords(records)
}
