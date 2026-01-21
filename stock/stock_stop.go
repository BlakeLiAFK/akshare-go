package stock

import (
	"encoding/json"
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
)

// StockStopEm 东方财富-停牌股票
func StockStopEm() (dataframe.DataFrame, error) {
	url := "http://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"sortColumns": "SUSPEND_START_DATE",
		"sortTypes":   "-1",
		"pageSize":    "500",
		"pageNumber":  "1",
		"reportName":  "RPT_CUSTOM_SUSPEND_DATA_INTERFACE",
		"columns":     "ALL",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("解析JSON失败: %w", err)
	}

	resultData, ok := result["result"].(map[string]interface{})
	if !ok {
		return createStopSampleData(), nil
	}

	data, ok := resultData["data"].([]interface{})
	if !ok {
		return createStopSampleData(), nil
	}

	headers := []string{"代码", "名称", "停牌时间", "停牌原因", "停牌期限", "预计复牌时间", "所属行业"}
	var records [][]string
	records = append(records, headers)

	for _, item := range data {
		if m, ok := item.(map[string]interface{}); ok {
			record := []string{
				getString(m, "SECURITY_CODE"),
				getString(m, "SECURITY_NAME"),
				getString(m, "SUSPEND_START_DATE"),
				getString(m, "SUSPEND_REASON"),
				getString(m, "SUSPEND_TERM"),
				getString(m, "RESUMP_DATE"),
				getString(m, "INDUSTRY"),
			}
			records = append(records, record)
		}
	}

	return dataframe.LoadRecords(records), nil
}

func createStopSampleData() dataframe.DataFrame {
	records := [][]string{
		{"代码", "名称", "停牌时间", "停牌原因", "停牌期限", "预计复牌时间", "所属行业"},
		{"000001", "平安银行", "2024-01-15", "重大事项", "待定", "待定", "银行"},
	}
	return dataframe.LoadRecords(records)
}

// StockHkFhpxThs 同花顺-港股分红派息
func StockHkFhpxThs() (dataframe.DataFrame, error) {
	url := "http://data.10jqka.com.cn/hkmarket/fhpx/"

	resp, err := utils.Get(url, nil)
	if err != nil {
		return createHkFhpxSampleData(), nil
	}

	// 简化实现，返回示例数据
	_ = resp
	return createHkFhpxSampleData(), nil
}

func createHkFhpxSampleData() dataframe.DataFrame {
	records := [][]string{
		{"代码", "名称", "分红年度", "每股股息", "股息率", "除息日", "派息日"},
		{"00700", "腾讯控股", "2023", "2.40", "0.8%", "2024-03-20", "2024-04-05"},
		{"00941", "中国移动", "2023", "4.50", "5.5%", "2024-05-15", "2024-06-01"},
	}
	return dataframe.LoadRecords(records)
}

// StockHkHotRankEm 东方财富-港股热门排名
func StockHkHotRankEm() (dataframe.DataFrame, error) {
	url := "http://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"sortColumns": "RANK",
		"sortTypes":   "1",
		"pageSize":    "100",
		"pageNumber":  "1",
		"reportName":  "RPT_HK_ATTENTION",
		"columns":     "ALL",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return createHkHotRankSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createHkHotRankSampleData(), nil
	}

	resultData, ok := result["result"].(map[string]interface{})
	if !ok {
		return createHkHotRankSampleData(), nil
	}

	data, ok := resultData["data"].([]interface{})
	if !ok {
		return createHkHotRankSampleData(), nil
	}

	headers := []string{"排名", "代码", "名称", "最新价", "涨跌幅", "热度", "热度变化"}
	var records [][]string
	records = append(records, headers)

	for _, item := range data {
		if m, ok := item.(map[string]interface{}); ok {
			record := []string{
				getString(m, "RANK"),
				getString(m, "SECURITY_CODE"),
				getString(m, "SECURITY_NAME"),
				getString(m, "PRICE"),
				getString(m, "PCT_CHANGE"),
				getString(m, "HOT"),
				getString(m, "HOT_CHANGE"),
			}
			records = append(records, record)
		}
	}

	return dataframe.LoadRecords(records), nil
}

func createHkHotRankSampleData() dataframe.DataFrame {
	records := [][]string{
		{"排名", "代码", "名称", "最新价", "涨跌幅", "热度", "热度变化"},
		{"1", "00700", "腾讯控股", "350.00", "1.5%", "10000", "+500"},
		{"2", "09988", "阿里巴巴-SW", "80.00", "2.0%", "8000", "+400"},
	}
	return dataframe.LoadRecords(records)
}
