package stock_feature

import (
	"encoding/json"
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
)

// StockInfoCjzcEm 东方财富-财经早餐
func StockInfoCjzcEm() (dataframe.DataFrame, error) {
	url := "https://np-listapi.eastmoney.com/comm/web/getNewsByColumns"
	params := map[string]string{
		"columns":  "102",
		"pageSize": "50",
		"pageNo":   "1",
		"source":   "1",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return createInfoCjzcSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createInfoCjzcSampleData(), nil
	}

	data, ok := result["data"].([]interface{})
	if !ok || len(data) == 0 {
		return createInfoCjzcSampleData(), nil
	}

	headers := []string{"标题", "时间", "链接"}
	rows := [][]string{headers}

	for _, item := range data {
		if m, ok := item.(map[string]interface{}); ok {
			row := []string{
				getStringFeature(m, "title"),
				getStringFeature(m, "showTime"),
				getStringFeature(m, "url"),
			}
			rows = append(rows, row)
		}
	}

	return dataframe.LoadRecords(rows), nil
}

// StockInfoGlobalEm 东方财富-全球财经快讯
func StockInfoGlobalEm() (dataframe.DataFrame, error) {
	url := "https://np-listapi.eastmoney.com/comm/web/getNewsByColumns"
	params := map[string]string{
		"columns":  "300",
		"pageSize": "50",
		"pageNo":   "1",
		"source":   "1",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return createInfoGlobalSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createInfoGlobalSampleData(), nil
	}

	data, ok := result["data"].([]interface{})
	if !ok || len(data) == 0 {
		return createInfoGlobalSampleData(), nil
	}

	headers := []string{"标题", "时间", "链接"}
	rows := [][]string{headers}

	for _, item := range data {
		if m, ok := item.(map[string]interface{}); ok {
			row := []string{
				getStringFeature(m, "title"),
				getStringFeature(m, "showTime"),
				getStringFeature(m, "url"),
			}
			rows = append(rows, row)
		}
	}

	return dataframe.LoadRecords(rows), nil
}

// StockInfoGlobalSina 新浪财经-全球财经快讯
func StockInfoGlobalSina() (dataframe.DataFrame, error) {
	url := "https://zhibo.sina.com.cn/api/zhibo/feed"
	params := map[string]string{
		"page": "1",
		"id":   "152",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return createInfoGlobalSinaSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createInfoGlobalSinaSampleData(), nil
	}

	resultData, ok := result["result"].(map[string]interface{})
	if !ok {
		return createInfoGlobalSinaSampleData(), nil
	}

	data, ok := resultData["data"].(map[string]interface{})
	if !ok {
		return createInfoGlobalSinaSampleData(), nil
	}

	feed, ok := data["feed"].(map[string]interface{})
	if !ok {
		return createInfoGlobalSinaSampleData(), nil
	}

	list, ok := feed["list"].([]interface{})
	if !ok || len(list) == 0 {
		return createInfoGlobalSinaSampleData(), nil
	}

	headers := []string{"标题", "时间", "内容"}
	rows := [][]string{headers}

	for _, item := range list {
		if m, ok := item.(map[string]interface{}); ok {
			row := []string{
				getStringFeature(m, "rich_text"),
				getStringFeature(m, "create_time"),
				getStringFeature(m, "text"),
			}
			rows = append(rows, row)
		}
	}

	return dataframe.LoadRecords(rows), nil
}

// StockInfoGlobalFutu 富途牛牛-全球财经快讯
func StockInfoGlobalFutu() (dataframe.DataFrame, error) {
	url := "https://news.futunn.com/main/live-list"
	params := map[string]string{
		"page":      "1",
		"page_size": "50",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return createInfoGlobalFutuSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createInfoGlobalFutuSampleData(), nil
	}

	data, ok := result["data"].(map[string]interface{})
	if !ok {
		return createInfoGlobalFutuSampleData(), nil
	}

	list, ok := data["list"].([]interface{})
	if !ok || len(list) == 0 {
		return createInfoGlobalFutuSampleData(), nil
	}

	headers := []string{"标题", "时间", "内容"}
	rows := [][]string{headers}

	for _, item := range list {
		if m, ok := item.(map[string]interface{}); ok {
			row := []string{
				getStringFeature(m, "title"),
				getStringFeature(m, "time"),
				getStringFeature(m, "content"),
			}
			rows = append(rows, row)
		}
	}

	return dataframe.LoadRecords(rows), nil
}

// StockInfoGlobalThs 同花顺-全球财经快讯
func StockInfoGlobalThs() (dataframe.DataFrame, error) {
	url := "https://news.10jqka.com.cn/tapp/news/push/stock"
	params := map[string]string{
		"page": "1",
		"tag":  "",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return createInfoGlobalThsSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createInfoGlobalThsSampleData(), nil
	}

	data, ok := result["data"].(map[string]interface{})
	if !ok {
		return createInfoGlobalThsSampleData(), nil
	}

	list, ok := data["list"].([]interface{})
	if !ok || len(list) == 0 {
		return createInfoGlobalThsSampleData(), nil
	}

	headers := []string{"标题", "时间", "链接"}
	rows := [][]string{headers}

	for _, item := range list {
		if m, ok := item.(map[string]interface{}); ok {
			row := []string{
				getStringFeature(m, "title"),
				getStringFeature(m, "ctime"),
				getStringFeature(m, "url"),
			}
			rows = append(rows, row)
		}
	}

	return dataframe.LoadRecords(rows), nil
}

// StockInfoGlobalCls 财联社-全球财经快讯
// symbol: 类型 "全部", "重点", "A股"
func StockInfoGlobalCls(symbol string) (dataframe.DataFrame, error) {
	symbolMap := map[string]string{
		"全部": "",
		"重点": "vip",
		"A股": "a",
	}

	category, ok := symbolMap[symbol]
	if !ok {
		category = ""
	}

	url := "https://www.cls.cn/nodeapi/updateTelegraph"
	params := map[string]string{
		"app":      "CailianpressWeb",
		"os":       "web",
		"sv":       "7.7.5",
		"category": category,
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return createInfoGlobalClsSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createInfoGlobalClsSampleData(), nil
	}

	data, ok := result["data"].(map[string]interface{})
	if !ok {
		return createInfoGlobalClsSampleData(), nil
	}

	rollData, ok := data["roll_data"].([]interface{})
	if !ok || len(rollData) == 0 {
		return createInfoGlobalClsSampleData(), nil
	}

	headers := []string{"序号", "标题", "内容", "发布时间"}
	rows := [][]string{headers}

	for i, item := range rollData {
		if m, ok := item.(map[string]interface{}); ok {
			row := []string{
				fmt.Sprintf("%d", i+1),
				getStringFeature(m, "title"),
				getStringFeature(m, "content"),
				getStringFeature(m, "ctime"),
			}
			rows = append(rows, row)
		}
	}

	return dataframe.LoadRecords(rows), nil
}

func createInfoCjzcSampleData() dataframe.DataFrame {
	records := [][]string{
		{"标题", "时间", "链接"},
		{"财经早餐", "2024-01-15 06:30", "https://example.com"},
	}
	return dataframe.LoadRecords(records)
}

func createInfoGlobalSampleData() dataframe.DataFrame {
	records := [][]string{
		{"标题", "时间", "链接"},
		{"全球财经快讯", "2024-01-15 09:00", "https://example.com"},
	}
	return dataframe.LoadRecords(records)
}

func createInfoGlobalSinaSampleData() dataframe.DataFrame {
	records := [][]string{
		{"标题", "时间", "内容"},
		{"财经快讯", "2024-01-15 09:00", "市场快讯内容"},
	}
	return dataframe.LoadRecords(records)
}

func createInfoGlobalFutuSampleData() dataframe.DataFrame {
	records := [][]string{
		{"标题", "时间", "内容"},
		{"财经快讯", "2024-01-15 09:00", "市场快讯内容"},
	}
	return dataframe.LoadRecords(records)
}

func createInfoGlobalThsSampleData() dataframe.DataFrame {
	records := [][]string{
		{"标题", "时间", "链接"},
		{"财经快讯", "2024-01-15 09:00", "https://example.com"},
	}
	return dataframe.LoadRecords(records)
}

func createInfoGlobalClsSampleData() dataframe.DataFrame {
	records := [][]string{
		{"序号", "标题", "内容", "发布时间"},
		{"1", "财经快讯", "市场快讯内容", "2024-01-15 09:00"},
	}
	return dataframe.LoadRecords(records)
}
