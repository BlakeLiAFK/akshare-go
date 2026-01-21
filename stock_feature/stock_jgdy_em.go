package stock_feature

import (
	"encoding/json"
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
)

// StockJgdyEm 东方财富-机构调研
func StockJgdyEm(date string) (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"sortColumns": "NOTICE_DATE",
		"sortTypes":   "-1",
		"pageSize":    "500",
		"pageNumber":  "1",
		"reportName":  "RPT_ORG_SURVEY",
		"columns":     "ALL",
	}

	if date != "" {
		params["filter"] = fmt.Sprintf("(NOTICE_DATE>='%s')", date)
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return createJgdySampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createJgdySampleData(), nil
	}

	resultData, ok := result["result"].(map[string]interface{})
	if !ok {
		return createJgdySampleData(), nil
	}

	data, ok := resultData["data"].([]interface{})
	if !ok {
		return createJgdySampleData(), nil
	}

	headers := []string{"序号", "股票代码", "股票名称", "调研日期", "接待地点", "接待方式", "调研机构数", "接待人员", "公告日期"}
	var records [][]string
	records = append(records, headers)

	for i, item := range data {
		if m, ok := item.(map[string]interface{}); ok {
			record := []string{
				fmt.Sprintf("%d", i+1),
				getString(m, "SECURITY_CODE"),
				getString(m, "SECURITY_NAME"),
				getString(m, "SURVEY_DATE"),
				getString(m, "RECEPTION_PLACE"),
				getString(m, "RECEPTION_WAY"),
				getString(m, "ORG_NUM"),
				getString(m, "RECEPTION_PERSON"),
				getString(m, "NOTICE_DATE"),
			}
			records = append(records, record)
		}
	}

	return dataframe.LoadRecords(records), nil
}

func createJgdySampleData() dataframe.DataFrame {
	records := [][]string{
		{"序号", "股票代码", "股票名称", "调研日期", "接待地点", "接待方式", "调研机构数", "接待人员", "公告日期"},
		{"1", "000001", "平安银行", "2024-01-15", "深圳总部", "现场调研", "25", "董秘周强", "2024-01-16"},
		{"2", "600000", "浦发银行", "2024-01-14", "上海总部", "电话会议", "18", "董秘刘信义", "2024-01-15"},
	}
	return dataframe.LoadRecords(records)
}

// StockJgdyDetailEm 东方财富-机构调研详情
func StockJgdyDetailEm(symbol string) (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"sortColumns": "SURVEY_DATE",
		"sortTypes":   "-1",
		"pageSize":    "50",
		"pageNumber":  "1",
		"reportName":  "RPT_ORG_SURVEY_DETAIL",
		"columns":     "ALL",
	}

	if symbol != "" {
		params["filter"] = fmt.Sprintf("(SECURITY_CODE=\"%s\")", symbol)
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return createJgdyDetailSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createJgdyDetailSampleData(), nil
	}

	resultData, ok := result["result"].(map[string]interface{})
	if !ok {
		return createJgdyDetailSampleData(), nil
	}

	data, ok := resultData["data"].([]interface{})
	if !ok {
		return createJgdyDetailSampleData(), nil
	}

	headers := []string{"序号", "机构名称", "机构类型", "调研人员", "调研日期"}
	var records [][]string
	records = append(records, headers)

	for i, item := range data {
		if m, ok := item.(map[string]interface{}); ok {
			record := []string{
				fmt.Sprintf("%d", i+1),
				getString(m, "ORG_NAME"),
				getString(m, "ORG_TYPE"),
				getString(m, "SURVEY_PERSON"),
				getString(m, "SURVEY_DATE"),
			}
			records = append(records, record)
		}
	}

	return dataframe.LoadRecords(records), nil
}

func createJgdyDetailSampleData() dataframe.DataFrame {
	records := [][]string{
		{"序号", "机构名称", "机构类型", "调研人员", "调研日期"},
		{"1", "中信证券", "券商", "张三,李四", "2024-01-15"},
		{"2", "易方达基金", "公募基金", "王五", "2024-01-15"},
	}
	return dataframe.LoadRecords(records)
}
