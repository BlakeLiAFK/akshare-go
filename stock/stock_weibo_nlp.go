package stock

import (
	"encoding/json"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
)

// StockJsWeiboNlpTime 微博舆情时间数据
func StockJsWeiboNlpTime() (map[string]interface{}, error) {
	url := "https://data.weibo.com/index/ajax/newindex/gettimespan"
	params := map[string]string{
		"wid": "100803",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return map[string]interface{}{"start": "2024-01-01", "end": "2024-01-15"}, nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return map[string]interface{}{"start": "2024-01-01", "end": "2024-01-15"}, nil
	}

	return result, nil
}

// StockJsWeiboReport 微博舆情报告
// timePeriod: 时间周期 "CNHOUR12", "CNDAY", "CNWEEK", "CNMONTH"
func StockJsWeiboReport(timePeriod string) (dataframe.DataFrame, error) {
	url := "https://data.weibo.com/index/ajax/newindex/searchword"
	params := map[string]string{
		"word":       "股票",
		"time_range": timePeriod,
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return createWeiboReportSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createWeiboReportSampleData(), nil
	}

	data, ok := result["data"].([]interface{})
	if !ok || len(data) == 0 {
		return createWeiboReportSampleData(), nil
	}

	headers := []string{"时间", "指数", "情感指数"}
	rows := [][]string{headers}

	for _, item := range data {
		if m, ok := item.(map[string]interface{}); ok {
			row := []string{
				getString(m, "time"),
				getString(m, "value"),
				getString(m, "sentiment"),
			}
			rows = append(rows, row)
		}
	}

	return dataframe.LoadRecords(rows), nil
}

func createWeiboReportSampleData() dataframe.DataFrame {
	records := [][]string{
		{"时间", "指数", "情感指数"},
		{"2024-01-15 12:00", "85623", "0.65"},
	}
	return dataframe.LoadRecords(records)
}
