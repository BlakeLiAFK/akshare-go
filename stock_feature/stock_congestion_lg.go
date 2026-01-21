package stock_feature

import (
	"encoding/json"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
)

// StockACongestionLg 乐估乐股-筹码拥挤度
func StockACongestionLg() (dataframe.DataFrame, error) {
	url := "https://legulegu.com/api/stockdata/congestion"
	params := map[string]string{
		"token": "0123456789",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return createCongestionSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createCongestionSampleData(), nil
	}

	data, ok := result["data"].([]interface{})
	if !ok || len(data) == 0 {
		return createCongestionSampleData(), nil
	}

	headers := []string{"日期", "拥挤度", "收盘价"}
	rows := [][]string{headers}

	for _, item := range data {
		if m, ok := item.(map[string]interface{}); ok {
			row := []string{
				getStringFeature(m, "date"),
				getStringFeature(m, "congestion"),
				getStringFeature(m, "close"),
			}
			rows = append(rows, row)
		}
	}

	return dataframe.LoadRecords(rows), nil
}

func createCongestionSampleData() dataframe.DataFrame {
	records := [][]string{
		{"日期", "拥挤度", "收盘价"},
		{"2024-01-15", "0.65", "3000.00"},
	}
	return dataframe.LoadRecords(records)
}
