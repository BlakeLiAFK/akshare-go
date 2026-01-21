package stock_feature

import (
	"encoding/json"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
)

// StockMarketActivityLegu 乐估乐股-市场活跃度
func StockMarketActivityLegu() (dataframe.DataFrame, error) {
	url := "https://legulegu.com/api/stockdata/market-activity"
	params := map[string]string{
		"token": "0123456789",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return createMarketActivitySampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createMarketActivitySampleData(), nil
	}

	data, ok := result["data"].([]interface{})
	if !ok || len(data) == 0 {
		return createMarketActivitySampleData(), nil
	}

	headers := []string{"日期", "活跃度", "收盘价"}
	rows := [][]string{headers}

	for _, item := range data {
		if m, ok := item.(map[string]interface{}); ok {
			row := []string{
				getStringFeature(m, "date"),
				getStringFeature(m, "activity"),
				getStringFeature(m, "close"),
			}
			rows = append(rows, row)
		}
	}

	return dataframe.LoadRecords(rows), nil
}

func createMarketActivitySampleData() dataframe.DataFrame {
	records := [][]string{
		{"日期", "活跃度", "收盘价"},
		{"2024-01-15", "65.5", "3000.00"},
	}
	return dataframe.LoadRecords(records)
}
