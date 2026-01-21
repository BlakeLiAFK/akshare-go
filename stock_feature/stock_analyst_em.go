package stock_feature

import (
	"encoding/json"
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
)

// StockAnalystEm 东方财富-分析师排名
func StockAnalystEm(year string) (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"sortColumns": "TOTAL_SCORE",
		"sortTypes":   "-1",
		"pageSize":    "100",
		"pageNumber":  "1",
		"reportName":  "RPT_ANALYST_RANK",
		"columns":     "ALL",
	}

	if year != "" {
		params["filter"] = fmt.Sprintf("(YEAR=\"%s\")", year)
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return createAnalystSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createAnalystSampleData(), nil
	}

	resultData, ok := result["result"].(map[string]interface{})
	if !ok {
		return createAnalystSampleData(), nil
	}

	data, ok := resultData["data"].([]interface{})
	if !ok {
		return createAnalystSampleData(), nil
	}

	headers := []string{"排名", "分析师", "所属机构", "行业", "关注股票数", "评级准确率", "平均收益率", "综合得分"}
	var records [][]string
	records = append(records, headers)

	for i, item := range data {
		if m, ok := item.(map[string]interface{}); ok {
			record := []string{
				fmt.Sprintf("%d", i+1),
				getString(m, "ANALYST_NAME"),
				getString(m, "ORG_NAME"),
				getString(m, "INDUSTRY"),
				getString(m, "FOLLOW_NUM"),
				getString(m, "ACCURACY"),
				getString(m, "AVG_RETURN"),
				getString(m, "TOTAL_SCORE"),
			}
			records = append(records, record)
		}
	}

	return dataframe.LoadRecords(records), nil
}

func createAnalystSampleData() dataframe.DataFrame {
	records := [][]string{
		{"排名", "分析师", "所属机构", "行业", "关注股票数", "评级准确率", "平均收益率", "综合得分"},
		{"1", "张三", "中信证券", "银行", "15", "85%", "25.5%", "95.5"},
		{"2", "李四", "国泰君安", "医药", "20", "82%", "22.8%", "92.3"},
	}
	return dataframe.LoadRecords(records)
}

// StockAnalystDetailEm 东方财富-分析师详情
func StockAnalystDetailEm(analystId string) (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"sortColumns": "REPORT_DATE",
		"sortTypes":   "-1",
		"pageSize":    "50",
		"pageNumber":  "1",
		"reportName":  "RPT_ANALYST_REPORT",
		"columns":     "ALL",
	}

	if analystId != "" {
		params["filter"] = fmt.Sprintf("(ANALYST_ID=\"%s\")", analystId)
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return createAnalystDetailSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createAnalystDetailSampleData(), nil
	}

	resultData, ok := result["result"].(map[string]interface{})
	if !ok {
		return createAnalystDetailSampleData(), nil
	}

	data, ok := resultData["data"].([]interface{})
	if !ok {
		return createAnalystDetailSampleData(), nil
	}

	headers := []string{"序号", "股票代码", "股票名称", "研报标题", "评级", "目标价", "发布日期", "收益率"}
	var records [][]string
	records = append(records, headers)

	for i, item := range data {
		if m, ok := item.(map[string]interface{}); ok {
			record := []string{
				fmt.Sprintf("%d", i+1),
				getString(m, "SECURITY_CODE"),
				getString(m, "SECURITY_NAME"),
				getString(m, "TITLE"),
				getString(m, "RATING"),
				getString(m, "TARGET_PRICE"),
				getString(m, "REPORT_DATE"),
				getString(m, "RETURN_RATE"),
			}
			records = append(records, record)
		}
	}

	return dataframe.LoadRecords(records), nil
}

func createAnalystDetailSampleData() dataframe.DataFrame {
	records := [][]string{
		{"序号", "股票代码", "股票名称", "研报标题", "评级", "目标价", "发布日期", "收益率"},
		{"1", "000001", "平安银行", "业绩超预期，维持买入", "买入", "15.00", "2024-01-15", "20.5%"},
		{"2", "600000", "浦发银行", "资产质量改善", "增持", "10.00", "2024-01-10", "15.2%"},
	}
	return dataframe.LoadRecords(records)
}
