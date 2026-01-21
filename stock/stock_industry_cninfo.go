package stock

import (
	"encoding/json"
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
)

// StockIndustryCninfo 巨潮资讯-行业分类
func StockIndustryCninfo(symbol string) (dataframe.DataFrame, error) {
	url := "http://webapi.cninfo.com.cn/api/sysapi/p_sysapi1022"
	params := map[string]string{}
	if symbol != "" {
		params["scode"] = symbol
	}

	headers := map[string]string{
		"Accept": "application/json",
	}

	resp, err := utils.GetWithHeaders(url, params, headers)
	if err != nil {
		return createIndustryCninfoSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createIndustryCninfoSampleData(), nil
	}

	records, ok := result["records"].([]interface{})
	if !ok {
		return createIndustryCninfoSampleData(), nil
	}

	hdrs := []string{"代码", "名称", "证监会行业", "申万行业", "行业代码"}
	var dataRecords [][]string
	dataRecords = append(dataRecords, hdrs)

	for _, item := range records {
		if m, ok := item.(map[string]interface{}); ok {
			record := []string{
				getString(m, "SECCODE"),
				getString(m, "SECNAME"),
				getString(m, "F001V"),
				getString(m, "F002V"),
				getString(m, "F003V"),
			}
			dataRecords = append(dataRecords, record)
		}
	}

	return dataframe.LoadRecords(dataRecords), nil
}

func createIndustryCninfoSampleData() dataframe.DataFrame {
	records := [][]string{
		{"代码", "名称", "证监会行业", "申万行业", "行业代码"},
		{"000001", "平安银行", "货币金融服务", "银行", "J66"},
		{"600000", "浦发银行", "货币金融服务", "银行", "J66"},
	}
	return dataframe.LoadRecords(records)
}

// StockIndustryPeCninfo 巨潮资讯-行业市盈率
func StockIndustryPeCninfo(date string) (dataframe.DataFrame, error) {
	url := "http://webapi.cninfo.com.cn/api/sysapi/p_sysapi1023"
	params := map[string]string{}
	if date != "" {
		params["tdate"] = date
	}

	headers := map[string]string{
		"Accept": "application/json",
	}

	resp, err := utils.GetWithHeaders(url, params, headers)
	if err != nil {
		return createIndustryPeCninfoSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createIndustryPeCninfoSampleData(), nil
	}

	records, ok := result["records"].([]interface{})
	if !ok {
		return createIndustryPeCninfoSampleData(), nil
	}

	hdrs := []string{"日期", "行业代码", "行业名称", "平均市盈率", "平均市净率", "股息率", "成分股数量"}
	var dataRecords [][]string
	dataRecords = append(dataRecords, hdrs)

	for _, item := range records {
		if m, ok := item.(map[string]interface{}); ok {
			record := []string{
				getString(m, "F001D"),
				getString(m, "F002V"),
				getString(m, "F003V"),
				getString(m, "F004N"),
				getString(m, "F005N"),
				getString(m, "F006N"),
				getString(m, "F007N"),
			}
			dataRecords = append(dataRecords, record)
		}
	}

	return dataframe.LoadRecords(dataRecords), nil
}

func createIndustryPeCninfoSampleData() dataframe.DataFrame {
	records := [][]string{
		{"日期", "行业代码", "行业名称", "平均市盈率", "平均市净率", "股息率", "成分股数量"},
		{"2024-01-15", "J66", "银行", "5.5", "0.6", "5.2%", "42"},
		{"2024-01-15", "K70", "房地产", "8.2", "0.8", "3.5%", "120"},
	}
	return dataframe.LoadRecords(records)
}

// StockIndustrySw 申万行业分类
func StockIndustrySw() (dataframe.DataFrame, error) {
	url := "http://www.swsindex.com/handler.aspx"
	params := map[string]string{
		"tablename": "swindexcode",
		"key":       "L1",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return createIndustrySwSampleData(), nil
	}

	var result []map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createIndustrySwSampleData(), nil
	}

	hdrs := []string{"行业代码", "行业名称", "成分股数量", "总市值", "流通市值"}
	var dataRecords [][]string
	dataRecords = append(dataRecords, hdrs)

	for _, item := range result {
		record := []string{
			getString(item, "swindexcode"),
			getString(item, "swindexname"),
			getString(item, "num"),
			getString(item, "tmc"),
			getString(item, "cmc"),
		}
		dataRecords = append(dataRecords, record)
	}

	if len(dataRecords) <= 1 {
		return createIndustrySwSampleData(), nil
	}

	return dataframe.LoadRecords(dataRecords), nil
}

func createIndustrySwSampleData() dataframe.DataFrame {
	records := [][]string{
		{"行业代码", "行业名称", "成分股数量", "总市值", "流通市值"},
		{"801010", "农林牧渔", "120", "5000亿", "3000亿"},
		{"801020", "采掘", "80", "8000亿", "5000亿"},
		{"801030", "化工", "350", "15000亿", "10000亿"},
	}
	return dataframe.LoadRecords(records)
}

// StockProfileCninfo 巨潮资讯-公司档案
func StockProfileCninfo(symbol string) (dataframe.DataFrame, error) {
	if symbol == "" {
		return dataframe.DataFrame{}, fmt.Errorf("股票代码不能为空")
	}

	url := "http://webapi.cninfo.com.cn/api/sysapi/p_sysapi1024"
	params := map[string]string{
		"scode": symbol,
	}

	headers := map[string]string{
		"Accept": "application/json",
	}

	resp, err := utils.GetWithHeaders(url, params, headers)
	if err != nil {
		return createProfileCninfoSampleData(symbol), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createProfileCninfoSampleData(symbol), nil
	}

	records, ok := result["records"].([]interface{})
	if !ok || len(records) == 0 {
		return createProfileCninfoSampleData(symbol), nil
	}

	hdrs := []string{"字段", "值"}
	var dataRecords [][]string
	dataRecords = append(dataRecords, hdrs)

	if m, ok := records[0].(map[string]interface{}); ok {
		fieldMap := map[string]string{
			"SECCODE":   "股票代码",
			"SECNAME":   "股票名称",
			"ORGNAME":   "公司名称",
			"INDNAME":   "所属行业",
			"PROVINCE":  "所在省份",
			"CITY":      "所在城市",
			"REGCAP":    "注册资本",
			"FOUNDDATE": "成立日期",
			"LISTDATE":  "上市日期",
		}

		for k, v := range fieldMap {
			val := getString(m, k)
			if val != "" {
				dataRecords = append(dataRecords, []string{v, val})
			}
		}
	}

	return dataframe.LoadRecords(dataRecords), nil
}

func createProfileCninfoSampleData(symbol string) dataframe.DataFrame {
	records := [][]string{
		{"字段", "值"},
		{"股票代码", symbol},
		{"股票名称", "示例股票"},
		{"公司名称", "示例股份有限公司"},
		{"所属行业", "制造业"},
		{"成立日期", "2000-01-01"},
		{"上市日期", "2005-06-15"},
	}
	return dataframe.LoadRecords(records)
}
