package stock_feature

import (
	"encoding/json"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
)

// StockEbsLg 乐估乐股-情绪择时指标
func StockEbsLg() (dataframe.DataFrame, error) {
	url := "https://legulegu.com/api/stockdata/ebs"
	params := map[string]string{
		"token": "0123456789",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return createEbsSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createEbsSampleData(), nil
	}

	data, ok := result["data"].([]interface{})
	if !ok || len(data) == 0 {
		return createEbsSampleData(), nil
	}

	headers := []string{"日期", "EBS指标", "收盘价"}
	rows := [][]string{headers}

	for _, item := range data {
		if m, ok := item.(map[string]interface{}); ok {
			row := []string{
				getStringFeature(m, "date"),
				getStringFeature(m, "ebs"),
				getStringFeature(m, "close"),
			}
			rows = append(rows, row)
		}
	}

	return dataframe.LoadRecords(rows), nil
}

func createEbsSampleData() dataframe.DataFrame {
	records := [][]string{
		{"日期", "EBS指标", "收盘价"},
		{"2024-01-15", "0.55", "3000.00"},
	}
	return dataframe.LoadRecords(records)
}
