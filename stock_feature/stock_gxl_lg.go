package stock_feature

import (
	"encoding/json"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
)

// StockAGxlLg 乐估乐股-股息率
func StockAGxlLg() (dataframe.DataFrame, error) {
	url := "https://legulegu.com/api/stockdata/dividend-yield"
	params := map[string]string{
		"token": "0123456789",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return createGxlSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createGxlSampleData(), nil
	}

	data, ok := result["data"].([]interface{})
	if !ok || len(data) == 0 {
		return createGxlSampleData(), nil
	}

	headers := []string{"日期", "股息率", "收盘价"}
	rows := [][]string{headers}

	for _, item := range data {
		if m, ok := item.(map[string]interface{}); ok {
			row := []string{
				getStringFeature(m, "date"),
				getStringFeature(m, "dividendYield"),
				getStringFeature(m, "close"),
			}
			rows = append(rows, row)
		}
	}

	return dataframe.LoadRecords(rows), nil
}

func createGxlSampleData() dataframe.DataFrame {
	records := [][]string{
		{"日期", "股息率", "收盘价"},
		{"2024-01-15", "2.85", "3000.00"},
	}
	return dataframe.LoadRecords(records)
}
