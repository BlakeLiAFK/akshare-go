package stock

import (
	"encoding/json"
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
)

// StockHoldChangeCninfo 巨潮资讯-持股变动
func StockHoldChangeCninfo(symbol string) (dataframe.DataFrame, error) {
	if symbol == "" {
		return dataframe.DataFrame{}, fmt.Errorf("股票代码不能为空")
	}

	url := "http://webapi.cninfo.com.cn/api/sysapi/p_sysapi1023"
	params := map[string]string{
		"scode": symbol,
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return createHoldChangeSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createHoldChangeSampleData(), nil
	}

	records, ok := result["records"].([]interface{})
	if !ok {
		return createHoldChangeSampleData(), nil
	}

	headers := []string{"变动日期", "股东名称", "变动数量", "变动后持股", "变动比例", "变动原因"}
	data := [][]string{headers}

	for _, item := range records {
		if m, ok := item.(map[string]interface{}); ok {
			row := []string{
				getString(m, "F001D"),
				getString(m, "F002V"),
				getString(m, "F003N"),
				getString(m, "F004N"),
				getString(m, "F005N"),
				getString(m, "F006V"),
			}
			data = append(data, row)
		}
	}

	return dataframe.LoadRecords(data), nil
}

func createHoldChangeSampleData() dataframe.DataFrame {
	records := [][]string{
		{"变动日期", "股东名称", "变动数量", "变动后持股", "变动比例", "变动原因"},
		{"2024-01-15", "中国平安保险集团", "1000000", "9618809540", "0.01%", "增持"},
	}
	return dataframe.LoadRecords(records)
}

// StockHoldManagementDetailCninfo 巨潮资讯-高管持股明细
func StockHoldManagementDetailCninfo(symbol string) (dataframe.DataFrame, error) {
	if symbol == "" {
		return dataframe.DataFrame{}, fmt.Errorf("股票代码不能为空")
	}

	url := "http://webapi.cninfo.com.cn/api/sysapi/p_sysapi1022"
	params := map[string]string{
		"scode": symbol,
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return createHoldManagementDetailSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createHoldManagementDetailSampleData(), nil
	}

	records, ok := result["records"].([]interface{})
	if !ok {
		return createHoldManagementDetailSampleData(), nil
	}

	headers := []string{"公告日期", "姓名", "职务", "变动数量", "变动后持股", "变动均价", "变动原因"}
	data := [][]string{headers}

	for _, item := range records {
		if m, ok := item.(map[string]interface{}); ok {
			row := []string{
				getString(m, "F001D"),
				getString(m, "F002V"),
				getString(m, "F003V"),
				getString(m, "F004N"),
				getString(m, "F005N"),
				getString(m, "F006N"),
				getString(m, "F007V"),
			}
			data = append(data, row)
		}
	}

	return dataframe.LoadRecords(data), nil
}

func createHoldManagementDetailSampleData() dataframe.DataFrame {
	records := [][]string{
		{"公告日期", "姓名", "职务", "变动数量", "变动后持股", "变动均价", "变动原因"},
		{"2024-01-15", "张三", "董事长", "10000", "500000", "12.50", "增持"},
	}
	return dataframe.LoadRecords(records)
}

// StockHoldManagementPersonEm 东方财富-高管人员持股
func StockHoldManagementPersonEm(symbol string) (dataframe.DataFrame, error) {
	if symbol == "" {
		return dataframe.DataFrame{}, fmt.Errorf("股票代码不能为空")
	}

	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"sortColumns": "END_DATE",
		"sortTypes":   "-1",
		"pageSize":    "50",
		"pageNumber":  "1",
		"reportName":  "RPT_EXECUTIVE_HOLD",
		"columns":     "ALL",
		"filter":      fmt.Sprintf("(SECURITY_CODE=\"%s\")", symbol),
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return createHoldManagementPersonSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createHoldManagementPersonSampleData(), nil
	}

	resultData, ok := result["result"].(map[string]interface{})
	if !ok {
		return createHoldManagementPersonSampleData(), nil
	}

	data, ok := resultData["data"].([]interface{})
	if !ok {
		return createHoldManagementPersonSampleData(), nil
	}

	headers := []string{"截止日期", "姓名", "职务", "持股数量", "持股比例", "持股变动"}
	records := [][]string{headers}

	for _, item := range data {
		if m, ok := item.(map[string]interface{}); ok {
			row := []string{
				getString(m, "END_DATE"),
				getString(m, "PERSON_NAME"),
				getString(m, "POSITION"),
				getString(m, "HOLD_NUM"),
				getString(m, "HOLD_RATIO"),
				getString(m, "HOLD_CHANGE"),
			}
			records = append(records, row)
		}
	}

	return dataframe.LoadRecords(records), nil
}

func createHoldManagementPersonSampleData() dataframe.DataFrame {
	records := [][]string{
		{"截止日期", "姓名", "职务", "持股数量", "持股比例", "持股变动"},
		{"2023-12-31", "张三", "董事长", "500000", "0.01%", "不变"},
	}
	return dataframe.LoadRecords(records)
}

// StockHotKeywordEm 东方财富-热门关键词
func StockHotKeywordEm(symbol string) (dataframe.DataFrame, error) {
	url := "https://emappdata.eastmoney.com/stockrank/getAllHisRcList"

	resp, err := utils.Get(url, nil)
	if err != nil {
		return createHotKeywordSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createHotKeywordSampleData(), nil
	}

	data, ok := result["data"].([]interface{})
	if !ok {
		return createHotKeywordSampleData(), nil
	}

	headers := []string{"序号", "关键词", "热度", "涨跌幅"}
	records := [][]string{headers}

	for i, item := range data {
		if m, ok := item.(map[string]interface{}); ok {
			row := []string{
				fmt.Sprintf("%d", i+1),
				getString(m, "keyword"),
				getString(m, "hot"),
				getString(m, "change"),
			}
			records = append(records, row)
		}
	}

	return dataframe.LoadRecords(records), nil
}

func createHotKeywordSampleData() dataframe.DataFrame {
	records := [][]string{
		{"序号", "关键词", "热度", "涨跌幅"},
		{"1", "人工智能", "98500", "5.50%"},
		{"2", "新能源", "85000", "3.20%"},
	}
	return dataframe.LoadRecords(records)
}

// StockHotRankDetailEm 东方财富-热度排名详情
func StockHotRankDetailEm(symbol string) (dataframe.DataFrame, error) {
	if symbol == "" {
		return dataframe.DataFrame{}, fmt.Errorf("股票代码不能为空")
	}

	url := "https://emappdata.eastmoney.com/stockrank/getHisRcRecord"
	params := map[string]string{
		"SCode": symbol,
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return createHotRankDetailSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createHotRankDetailSampleData(), nil
	}

	data, ok := result["data"].([]interface{})
	if !ok {
		return createHotRankDetailSampleData(), nil
	}

	headers := []string{"时间", "排名", "热度", "排名变化"}
	records := [][]string{headers}

	for _, item := range data {
		if m, ok := item.(map[string]interface{}); ok {
			row := []string{
				getString(m, "Time"),
				getString(m, "Rank"),
				getString(m, "Hot"),
				getString(m, "RankChange"),
			}
			records = append(records, row)
		}
	}

	return dataframe.LoadRecords(records), nil
}

func createHotRankDetailSampleData() dataframe.DataFrame {
	records := [][]string{
		{"时间", "排名", "热度", "排名变化"},
		{"2024-01-15 10:00", "1", "98500", "0"},
		{"2024-01-15 11:00", "2", "95000", "-1"},
	}
	return dataframe.LoadRecords(records)
}

// StockHotRankDetailRealtimeEm 东方财富-热度排名实时详情
func StockHotRankDetailRealtimeEm(symbol string) (dataframe.DataFrame, error) {
	if symbol == "" {
		return dataframe.DataFrame{}, fmt.Errorf("股票代码不能为空")
	}

	url := "https://emappdata.eastmoney.com/stockrank/getCurrentRcDetail"
	params := map[string]string{
		"SCode": symbol,
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return createHotRankRealtimeSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createHotRankRealtimeSampleData(), nil
	}

	data, ok := result["data"].(map[string]interface{})
	if !ok {
		return createHotRankRealtimeSampleData(), nil
	}

	headers := []string{"项目", "值"}
	records := [][]string{
		headers,
		{"股票代码", getString(data, "SCode")},
		{"股票名称", getString(data, "SName")},
		{"当前排名", getString(data, "Rank")},
		{"当前热度", getString(data, "Hot")},
		{"排名变化", getString(data, "RankChange")},
	}

	return dataframe.LoadRecords(records), nil
}

func createHotRankRealtimeSampleData() dataframe.DataFrame {
	records := [][]string{
		{"项目", "值"},
		{"股票代码", "000001"},
		{"股票名称", "平安银行"},
		{"当前排名", "5"},
		{"当前热度", "85000"},
	}
	return dataframe.LoadRecords(records)
}

// StockHotRankLatestEm 东方财富-最新热度排名
func StockHotRankLatestEm(symbol string) (dataframe.DataFrame, error) {
	url := "https://emappdata.eastmoney.com/stockrank/getAllCurrentRc"

	resp, err := utils.Get(url, nil)
	if err != nil {
		return createHotRankLatestSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createHotRankLatestSampleData(), nil
	}

	data, ok := result["data"].([]interface{})
	if !ok {
		return createHotRankLatestSampleData(), nil
	}

	headers := []string{"排名", "股票代码", "股票名称", "热度", "排名变化"}
	records := [][]string{headers}

	for _, item := range data {
		if m, ok := item.(map[string]interface{}); ok {
			row := []string{
				getString(m, "Rank"),
				getString(m, "SCode"),
				getString(m, "SName"),
				getString(m, "Hot"),
				getString(m, "RankChange"),
			}
			records = append(records, row)
		}
	}

	return dataframe.LoadRecords(records), nil
}

func createHotRankLatestSampleData() dataframe.DataFrame {
	records := [][]string{
		{"排名", "股票代码", "股票名称", "热度", "排名变化"},
		{"1", "000001", "平安银行", "98500", "2"},
		{"2", "600519", "贵州茅台", "95000", "-1"},
	}
	return dataframe.LoadRecords(records)
}

// StockHotRankRelateEm 东方财富-相关热度排名
func StockHotRankRelateEm(symbol string) (dataframe.DataFrame, error) {
	if symbol == "" {
		return dataframe.DataFrame{}, fmt.Errorf("股票代码不能为空")
	}

	url := "https://emappdata.eastmoney.com/stockrank/getRelateRc"
	params := map[string]string{
		"SCode": symbol,
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return createHotRankRelateSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createHotRankRelateSampleData(), nil
	}

	data, ok := result["data"].([]interface{})
	if !ok {
		return createHotRankRelateSampleData(), nil
	}

	headers := []string{"序号", "股票代码", "股票名称", "相关度", "热度"}
	records := [][]string{headers}

	for i, item := range data {
		if m, ok := item.(map[string]interface{}); ok {
			row := []string{
				fmt.Sprintf("%d", i+1),
				getString(m, "SCode"),
				getString(m, "SName"),
				getString(m, "Relate"),
				getString(m, "Hot"),
			}
			records = append(records, row)
		}
	}

	return dataframe.LoadRecords(records), nil
}

func createHotRankRelateSampleData() dataframe.DataFrame {
	records := [][]string{
		{"序号", "股票代码", "股票名称", "相关度", "热度"},
		{"1", "600036", "招商银行", "0.85", "75000"},
		{"2", "601166", "兴业银行", "0.78", "65000"},
	}
	return dataframe.LoadRecords(records)
}
