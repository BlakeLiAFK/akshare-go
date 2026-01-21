package stock

import (
	"encoding/json"
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
)

// StockZtPoolEm 东方财富-涨停板池
func StockZtPoolEm(date string) (dataframe.DataFrame, error) {
	url := "http://push2ex.eastmoney.com/getTopicZTPool"
	params := map[string]string{
		"ut":        "7eea3edcaed734bea9cbfc24409ed989",
		"dpt":       "wz.ztzt",
		"Pageindex": "0",
		"pagesize":  "500",
		"sort":      "fbt:asc",
		"date":      date,
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("解析JSON失败: %w", err)
	}

	data, ok := result["data"].(map[string]interface{})
	if !ok {
		return dataframe.DataFrame{}, fmt.Errorf("数据格式错误")
	}

	pool, ok := data["pool"].([]interface{})
	if !ok {
		return dataframe.DataFrame{}, fmt.Errorf("数据格式错误")
	}

	headers := []string{"代码", "名称", "涨跌幅", "最新价", "成交额", "流通市值", "总市值", "换手率", "封板资金", "首次封板时间", "最后封板时间", "炸板次数", "涨停统计", "连板数", "所属行业"}
	var records [][]string
	records = append(records, headers)

	for _, item := range pool {
		if m, ok := item.(map[string]interface{}); ok {
			record := []string{
				getString(m, "c"),
				getString(m, "n"),
				getString(m, "zdp"),
				getString(m, "p"),
				getString(m, "amount"),
				getString(m, "ltsz"),
				getString(m, "tshare"),
				getString(m, "hs"),
				getString(m, "fund"),
				getString(m, "fbt"),
				getString(m, "lbt"),
				getString(m, "zbc"),
				getString(m, "zttj"),
				getString(m, "lbc"),
				getString(m, "hybk"),
			}
			records = append(records, record)
		}
	}

	return dataframe.LoadRecords(records), nil
}

// StockDtPoolEm 东方财富-跌停板池
func StockDtPoolEm(date string) (dataframe.DataFrame, error) {
	url := "http://push2ex.eastmoney.com/getTopicDTPool"
	params := map[string]string{
		"ut":        "7eea3edcaed734bea9cbfc24409ed989",
		"dpt":       "wz.ztzt",
		"Pageindex": "0",
		"pagesize":  "500",
		"sort":      "fund:asc",
		"date":      date,
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("解析JSON失败: %w", err)
	}

	data, ok := result["data"].(map[string]interface{})
	if !ok {
		return dataframe.DataFrame{}, fmt.Errorf("数据格式错误")
	}

	pool, ok := data["pool"].([]interface{})
	if !ok {
		return dataframe.DataFrame{}, fmt.Errorf("数据格式错误")
	}

	headers := []string{"代码", "名称", "涨跌幅", "最新价", "成交额", "流通市值", "总市值", "换手率", "动态市盈率", "封单资金", "最后封板时间", "板上成交额", "连续跌停", "开板次数", "所属行业"}
	var records [][]string
	records = append(records, headers)

	for _, item := range pool {
		if m, ok := item.(map[string]interface{}); ok {
			record := []string{
				getString(m, "c"),
				getString(m, "n"),
				getString(m, "zdp"),
				getString(m, "p"),
				getString(m, "amount"),
				getString(m, "ltsz"),
				getString(m, "tshare"),
				getString(m, "hs"),
				getString(m, "pe"),
				getString(m, "fund"),
				getString(m, "lbt"),
				getString(m, "bamount"),
				getString(m, "days"),
				getString(m, "oc"),
				getString(m, "hybk"),
			}
			records = append(records, record)
		}
	}

	return dataframe.LoadRecords(records), nil
}

// StockZbPoolEm 东方财富-炸板池
func StockZbPoolEm(date string) (dataframe.DataFrame, error) {
	url := "http://push2ex.eastmoney.com/getTopicZBPool"
	params := map[string]string{
		"ut":        "7eea3edcaed734bea9cbfc24409ed989",
		"dpt":       "wz.ztzt",
		"Pageindex": "0",
		"pagesize":  "500",
		"sort":      "fbt:asc",
		"date":      date,
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("解析JSON失败: %w", err)
	}

	data, ok := result["data"].(map[string]interface{})
	if !ok {
		return dataframe.DataFrame{}, fmt.Errorf("数据格式错误")
	}

	pool, ok := data["pool"].([]interface{})
	if !ok {
		return dataframe.DataFrame{}, fmt.Errorf("数据格式错误")
	}

	headers := []string{"代码", "名称", "涨跌幅", "最新价", "成交额", "流通市值", "总市值", "换手率", "动态市盈率", "首次封板时间", "最后封板时间", "炸板次数", "涨停统计", "振幅", "所属行业"}
	var records [][]string
	records = append(records, headers)

	for _, item := range pool {
		if m, ok := item.(map[string]interface{}); ok {
			record := []string{
				getString(m, "c"),
				getString(m, "n"),
				getString(m, "zdp"),
				getString(m, "p"),
				getString(m, "amount"),
				getString(m, "ltsz"),
				getString(m, "tshare"),
				getString(m, "hs"),
				getString(m, "pe"),
				getString(m, "fbt"),
				getString(m, "lbt"),
				getString(m, "zbc"),
				getString(m, "zttj"),
				getString(m, "zf"),
				getString(m, "hybk"),
			}
			records = append(records, record)
		}
	}

	return dataframe.LoadRecords(records), nil
}

// StockQsJyEm 东方财富-强势股票
func StockQsJyEm() (dataframe.DataFrame, error) {
	url := "http://push2ex.eastmoney.com/getTopicQSJY"
	params := map[string]string{
		"ut":        "7eea3edcaed734bea9cbfc24409ed989",
		"dpt":       "wz.ztzt",
		"Pageindex": "0",
		"pagesize":  "500",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("解析JSON失败: %w", err)
	}

	data, ok := result["data"].(map[string]interface{})
	if !ok {
		return dataframe.DataFrame{}, fmt.Errorf("数据格式错误")
	}

	pool, ok := data["pool"].([]interface{})
	if !ok {
		return dataframe.DataFrame{}, fmt.Errorf("数据格式错误")
	}

	headers := []string{"代码", "名称", "涨跌幅", "最新价", "成交额", "流通市值", "总市值", "换手率", "涨停价", "是否新股", "入选理由"}
	var records [][]string
	records = append(records, headers)

	for _, item := range pool {
		if m, ok := item.(map[string]interface{}); ok {
			record := []string{
				getString(m, "c"),
				getString(m, "n"),
				getString(m, "zdp"),
				getString(m, "p"),
				getString(m, "amount"),
				getString(m, "ltsz"),
				getString(m, "tshare"),
				getString(m, "hs"),
				getString(m, "ztp"),
				getString(m, "isnew"),
				getString(m, "reason"),
			}
			records = append(records, record)
		}
	}

	return dataframe.LoadRecords(records), nil
}
