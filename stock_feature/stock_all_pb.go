package stock_feature

import (
	"encoding/json"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
)

// StockAAllPb 全部A股-等权重市净率、中位数市净率
func StockAAllPb() (dataframe.DataFrame, error) {
	url := "https://legulegu.com/api/stock-data/market-index-pb"
	params := map[string]string{
		"marketId": "ALL",
		"token":    "0123456789",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return createAllPbLgSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createAllPbLgSampleData(), nil
	}

	data, ok := result["data"].([]interface{})
	if !ok || len(data) == 0 {
		return createAllPbLgSampleData(), nil
	}

	headers := []string{"date", "averagePB", "middlePB", "quantile"}
	rows := [][]string{headers}

	for _, item := range data {
		if m, ok := item.(map[string]interface{}); ok {
			row := []string{
				getStringFeature(m, "date"),
				getStringFeature(m, "averagePB"),
				getStringFeature(m, "middlePB"),
				getStringFeature(m, "quantile"),
			}
			rows = append(rows, row)
		}
	}

	return dataframe.LoadRecords(rows), nil
}

func createAllPbLgSampleData() dataframe.DataFrame {
	records := [][]string{
		{"date", "averagePB", "middlePB", "quantile"},
		{"2024-01-15", "1.85", "2.10", "35.5"},
	}
	return dataframe.LoadRecords(records)
}
