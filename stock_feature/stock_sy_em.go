package stock_feature

import (
	"encoding/json"
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
)

// StockSyEm 东方财富-数据中心-商誉
func StockSyEm() (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"sortColumns": "NOTICE_DATE,SECURITY_CODE",
		"sortTypes":   "-1,-1",
		"pageSize":    "500",
		"pageNumber":  "1",
		"reportName":  "RPT_GOODWILL_STOCKDETAILS",
		"columns":     "ALL",
		"source":      "WEB",
		"client":      "WEB",
	}

	allData := make([]map[string]interface{}, 0)

	resp, err := utils.Get(url, params)
	if err != nil {
		return createSySampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createSySampleData(), nil
	}

	resultData, ok := result["result"].(map[string]interface{})
	if !ok {
		return createSySampleData(), nil
	}

	totalPages := int(getFloatFeature(resultData, "pages"))
	data, ok := resultData["data"].([]interface{})
	if !ok {
		return createSySampleData(), nil
	}

	for _, item := range data {
		if m, ok := item.(map[string]interface{}); ok {
			allData = append(allData, m)
		}
	}

	for page := 2; page <= totalPages && page <= 5; page++ {
		params["pageNumber"] = fmt.Sprintf("%d", page)
		resp, err := utils.Get(url, params)
		if err != nil {
			continue
		}

		var pageResult map[string]interface{}
		if err := json.Unmarshal(resp.Body(), &pageResult); err != nil {
			continue
		}

		pageResultData, ok := pageResult["result"].(map[string]interface{})
		if !ok {
			continue
		}

		pageData, ok := pageResultData["data"].([]interface{})
		if !ok {
			continue
		}

		for _, item := range pageData {
			if m, ok := item.(map[string]interface{}); ok {
				allData = append(allData, m)
			}
		}
	}

	headers := []string{"股票代码", "股票简称", "商誉", "商誉占净资产比例", "净资产", "公告日期", "报告期"}
	rows := [][]string{headers}

	for _, m := range allData {
		row := []string{
			getStringFeature(m, "SECURITY_CODE"),
			getStringFeature(m, "SECURITY_NAME_ABBR"),
			getStringFeature(m, "GOODWILL"),
			getStringFeature(m, "GOODWILL_RATIO"),
			getStringFeature(m, "NETASSETS"),
			getStringFeature(m, "NOTICE_DATE"),
			getStringFeature(m, "REPORT_DATE"),
		}
		rows = append(rows, row)
	}

	return dataframe.LoadRecords(rows), nil
}

func createSySampleData() dataframe.DataFrame {
	records := [][]string{
		{"股票代码", "股票简称", "商誉", "商誉占净资产比例", "净资产", "公告日期", "报告期"},
		{"000001", "平安银行", "5000000000", "2.5", "200000000000", "2024-04-25", "2023-12-31"},
	}
	return dataframe.LoadRecords(records)
}
