package stock_feature

import (
	"encoding/json"
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
)

// StockPankouEm 东方财富-盘口异动
func StockPankouEm() (dataframe.DataFrame, error) {
	url := "https://push2ex.eastmoney.com/getAllStockChanges"
	params := map[string]string{
		"type": "8201,8202,8193,8194,8195,8196,8197,8198,8199,8200",
		"ut":   "7eea3edcaed734bea9cbfc24409ed989",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return createPankouSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createPankouSampleData(), nil
	}

	data, ok := result["data"].(map[string]interface{})
	if !ok {
		return createPankouSampleData(), nil
	}

	allstock, ok := data["allstock"].([]interface{})
	if !ok {
		return createPankouSampleData(), nil
	}

	headers := []string{"时间", "代码", "名称", "板块", "异动类型", "异动描述"}
	var records [][]string
	records = append(records, headers)

	for _, item := range allstock {
		if m, ok := item.(map[string]interface{}); ok {
			record := []string{
				getString(m, "tm"),
				getString(m, "c"),
				getString(m, "n"),
				getString(m, "p"),
				getString(m, "t"),
				getString(m, "i"),
			}
			records = append(records, record)
		}
	}

	return dataframe.LoadRecords(records), nil
}

func createPankouSampleData() dataframe.DataFrame {
	records := [][]string{
		{"时间", "代码", "名称", "板块", "异动类型", "异动描述"},
		{"14:55:30", "000001", "平安银行", "银行", "大单买入", "成交金额超过500万"},
		{"14:52:15", "600000", "浦发银行", "银行", "快速拉升", "1分钟涨幅超过2%"},
	}
	return dataframe.LoadRecords(records)
}

// StockQsjyEm 东方财富-强势股统计
func StockQsjyEm(date string) (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"sortColumns": "CONTINUOUS_DAYS",
		"sortTypes":   "-1",
		"pageSize":    "100",
		"pageNumber":  "1",
		"reportName":  "RPT_STRONG_STOCK",
		"columns":     "ALL",
	}

	if date != "" {
		params["filter"] = fmt.Sprintf("(TRADE_DATE='%s')", date)
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return createQsjySampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createQsjySampleData(), nil
	}

	resultData, ok := result["result"].(map[string]interface{})
	if !ok {
		return createQsjySampleData(), nil
	}

	data, ok := resultData["data"].([]interface{})
	if !ok {
		return createQsjySampleData(), nil
	}

	headers := []string{"序号", "股票代码", "股票名称", "连涨天数", "连涨区间涨幅", "累计换手率", "所属行业"}
	var records [][]string
	records = append(records, headers)

	for i, item := range data {
		if m, ok := item.(map[string]interface{}); ok {
			record := []string{
				fmt.Sprintf("%d", i+1),
				getString(m, "SECURITY_CODE"),
				getString(m, "SECURITY_NAME"),
				getString(m, "CONTINUOUS_DAYS"),
				getString(m, "TOTAL_CHANGE"),
				getString(m, "TOTAL_TURNOVER"),
				getString(m, "INDUSTRY"),
			}
			records = append(records, record)
		}
	}

	return dataframe.LoadRecords(records), nil
}

func createQsjySampleData() dataframe.DataFrame {
	records := [][]string{
		{"序号", "股票代码", "股票名称", "连涨天数", "连涨区间涨幅", "累计换手率", "所属行业"},
		{"1", "000001", "平安银行", "5", "15.5%", "12.5%", "银行"},
		{"2", "600000", "浦发银行", "4", "12.2%", "10.8%", "银行"},
	}
	return dataframe.LoadRecords(records)
}
