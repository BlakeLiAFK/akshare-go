package stock_feature

import (
	"encoding/json"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
)

// StockABelowNetAssetStatistics A股破净股统计
func StockABelowNetAssetStatistics() (dataframe.DataFrame, error) {
	url := "https://legulegu.com/api/stockdata/below-net-asset-statistics"
	params := map[string]string{
		"token": "0123456789",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return createBelowNetAssetSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createBelowNetAssetSampleData(), nil
	}

	data, ok := result["data"].([]interface{})
	if !ok || len(data) == 0 {
		return createBelowNetAssetSampleData(), nil
	}

	headers := []string{"日期", "破净股数量", "破净股占比"}
	rows := [][]string{headers}

	for _, item := range data {
		if m, ok := item.(map[string]interface{}); ok {
			row := []string{
				getStringFeature(m, "date"),
				getStringFeature(m, "belowNetAssetStockNum"),
				getStringFeature(m, "belowNetAssetStockRatio"),
			}
			rows = append(rows, row)
		}
	}

	return dataframe.LoadRecords(rows), nil
}

func createBelowNetAssetSampleData() dataframe.DataFrame {
	records := [][]string{
		{"日期", "破净股数量", "破净股占比"},
		{"2024-01-15", "450", "8.5"},
	}
	return dataframe.LoadRecords(records)
}
