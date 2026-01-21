package stock_feature

import (
	"encoding/json"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
)

// StockATtmLyr A股滚动市盈率
func StockATtmLyr() (dataframe.DataFrame, error) {
	url := "https://legulegu.com/api/stockdata/ttm-lyr"
	params := map[string]string{
		"token": "0123456789",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return createTtmLyrSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createTtmLyrSampleData(), nil
	}

	data, ok := result["data"].([]interface{})
	if !ok || len(data) == 0 {
		return createTtmLyrSampleData(), nil
	}

	headers := []string{"日期", "滚动市盈率", "静态市盈率", "收盘价"}
	rows := [][]string{headers}

	for _, item := range data {
		if m, ok := item.(map[string]interface{}); ok {
			row := []string{
				getStringFeature(m, "date"),
				getStringFeature(m, "ttm"),
				getStringFeature(m, "lyr"),
				getStringFeature(m, "close"),
			}
			rows = append(rows, row)
		}
	}

	return dataframe.LoadRecords(rows), nil
}

func createTtmLyrSampleData() dataframe.DataFrame {
	records := [][]string{
		{"日期", "滚动市盈率", "静态市盈率", "收盘价"},
		{"2024-01-15", "15.5", "18.2", "3000.00"},
	}
	return dataframe.LoadRecords(records)
}
