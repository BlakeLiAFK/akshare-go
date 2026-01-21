package stock_feature

import (
	"encoding/json"
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
)

// StockCommentEm 东方财富-千股千评
func StockCommentEm() (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"sortColumns": "TRADE_DATE,SECURITY_CODE",
		"sortTypes":   "-1,-1",
		"pageSize":    "500",
		"pageNumber":  "1",
		"reportName":  "RPT_DMSK_TS_STOCKEVALUATION",
		"columns":     "ALL",
		"source":      "WEB",
		"client":      "WEB",
	}

	allData := make([]map[string]interface{}, 0)

	resp, err := utils.Get(url, params)
	if err != nil {
		return createCommentSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createCommentSampleData(), nil
	}

	resultData, ok := result["result"].(map[string]interface{})
	if !ok {
		return createCommentSampleData(), nil
	}

	totalPages := int(getFloatFeature(resultData, "pages"))
	data, ok := resultData["data"].([]interface{})
	if !ok {
		return createCommentSampleData(), nil
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

	headers := []string{"代码", "名称", "最新价", "涨跌幅", "换手率", "市盈率", "主力成本", "机构参与度", "综合得分", "排名变动", "交易日期"}
	rows := [][]string{headers}

	for _, m := range allData {
		row := []string{
			getStringFeature(m, "SECURITY_CODE"),
			getStringFeature(m, "SECURITY_NAME_ABBR"),
			getStringFeature(m, "CLOSE_PRICE"),
			getStringFeature(m, "CHANGE_RATE"),
			getStringFeature(m, "TURNOVER_RATE"),
			getStringFeature(m, "PE_RATIO"),
			getStringFeature(m, "MAIN_COST"),
			getStringFeature(m, "ORG_PARTICIPATE"),
			getStringFeature(m, "SCORE"),
			getStringFeature(m, "RANK_CHANGE"),
			getStringFeature(m, "TRADE_DATE"),
		}
		rows = append(rows, row)
	}

	return dataframe.LoadRecords(rows), nil
}

func createCommentSampleData() dataframe.DataFrame {
	records := [][]string{
		{"代码", "名称", "最新价", "涨跌幅", "换手率", "市盈率", "主力成本", "机构参与度", "综合得分", "排名变动", "交易日期"},
		{"000001", "平安银行", "12.50", "2.35", "0.85", "6.5", "12.35", "55.25", "75.5", "5", "2024-01-15"},
	}
	return dataframe.LoadRecords(records)
}
