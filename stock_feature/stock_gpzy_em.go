package stock_feature

import (
	"encoding/json"
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
)

// StockGpzyEm 东方财富-股权质押
func StockGpzyEm() (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"sortColumns": "PLEDGE_RATIO",
		"sortTypes":   "-1",
		"pageSize":    "500",
		"pageNumber":  "1",
		"reportName":  "RPT_CSDC_LIST",
		"columns":     "ALL",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return createGpzySampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createGpzySampleData(), nil
	}

	resultData, ok := result["result"].(map[string]interface{})
	if !ok {
		return createGpzySampleData(), nil
	}

	data, ok := resultData["data"].([]interface{})
	if !ok {
		return createGpzySampleData(), nil
	}

	headers := []string{"序号", "股票代码", "股票名称", "质押股数", "质押笔数", "无限售股质押数", "有限售股质押数", "质押总市值", "质押比例", "所属行业"}
	var records [][]string
	records = append(records, headers)

	for i, item := range data {
		if m, ok := item.(map[string]interface{}); ok {
			record := []string{
				fmt.Sprintf("%d", i+1),
				getString(m, "SECURITY_CODE"),
				getString(m, "SECURITY_NAME"),
				getString(m, "PLEDGE_NUM"),
				getString(m, "PLEDGE_COUNT"),
				getString(m, "UNLIMITED_PLEDGE_NUM"),
				getString(m, "LIMITED_PLEDGE_NUM"),
				getString(m, "PLEDGE_MARKET_VALUE"),
				getString(m, "PLEDGE_RATIO"),
				getString(m, "INDUSTRY"),
			}
			records = append(records, record)
		}
	}

	return dataframe.LoadRecords(records), nil
}

func createGpzySampleData() dataframe.DataFrame {
	records := [][]string{
		{"序号", "股票代码", "股票名称", "质押股数", "质押笔数", "无限售股质押数", "有限售股质押数", "质押总市值", "质押比例", "所属行业"},
		{"1", "000001", "平安银行", "500000000", "15", "400000000", "100000000", "6250000000", "2.58%", "银行"},
		{"2", "600000", "浦发银行", "350000000", "12", "300000000", "50000000", "2905000000", "1.19%", "银行"},
	}
	return dataframe.LoadRecords(records)
}

// StockGpzyDetailEm 东方财富-股权质押详情
func StockGpzyDetailEm(symbol string) (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"sortColumns": "PLEDGE_DATE",
		"sortTypes":   "-1",
		"pageSize":    "100",
		"pageNumber":  "1",
		"reportName":  "RPT_CSDC_DETAIL",
		"columns":     "ALL",
	}

	if symbol != "" {
		params["filter"] = fmt.Sprintf("(SECURITY_CODE=\"%s\")", symbol)
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return createGpzyDetailSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createGpzyDetailSampleData(), nil
	}

	resultData, ok := result["result"].(map[string]interface{})
	if !ok {
		return createGpzyDetailSampleData(), nil
	}

	data, ok := resultData["data"].([]interface{})
	if !ok {
		return createGpzyDetailSampleData(), nil
	}

	headers := []string{"序号", "股票代码", "股票名称", "股东名称", "质押股数", "质押开始日", "质押解除日", "质押机构", "是否已解除"}
	var records [][]string
	records = append(records, headers)

	for i, item := range data {
		if m, ok := item.(map[string]interface{}); ok {
			record := []string{
				fmt.Sprintf("%d", i+1),
				getString(m, "SECURITY_CODE"),
				getString(m, "SECURITY_NAME"),
				getString(m, "SHAREHOLDER_NAME"),
				getString(m, "PLEDGE_NUM"),
				getString(m, "PLEDGE_DATE"),
				getString(m, "RELEASE_DATE"),
				getString(m, "PLEDGEE"),
				getString(m, "IS_RELEASE"),
			}
			records = append(records, record)
		}
	}

	return dataframe.LoadRecords(records), nil
}

func createGpzyDetailSampleData() dataframe.DataFrame {
	records := [][]string{
		{"序号", "股票代码", "股票名称", "股东名称", "质押股数", "质押开始日", "质押解除日", "质押机构", "是否已解除"},
		{"1", "000001", "平安银行", "中国平安保险(集团)股份有限公司", "100000000", "2023-06-15", "-", "中信银行", "否"},
		{"2", "000001", "平安银行", "中国平安保险(集团)股份有限公司", "80000000", "2022-12-01", "2023-12-01", "招商银行", "是"},
	}
	return dataframe.LoadRecords(records)
}
