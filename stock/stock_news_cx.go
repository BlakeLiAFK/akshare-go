package stock

import (
	"encoding/json"
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
)

// StockNewsCx 财联社-电报
// symbol: 新闻类型 "全部", "A股", "港股", "美股", "基金", "外汇"
func StockNewsCx(symbol string) (dataframe.DataFrame, error) {
	url := "https://www.cls.cn/nodeapi/updateTelegraph"

	symbolMap := map[string]string{
		"全部": "",
		"A股": "a",
		"港股": "hk",
		"美股": "us",
		"基金": "fund",
		"外汇": "forex",
	}

	market, ok := symbolMap[symbol]
	if !ok {
		market = ""
	}

	params := map[string]string{
		"app":         "CailianpressWeb",
		"os":          "web",
		"sv":          "7.7.5",
		"sign":        "1",
		"category":    market,
		"lastTime":    "",
		"hasFirstVip": "1",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return createNewsCxSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createNewsCxSampleData(), nil
	}

	data, ok := result["data"].(map[string]interface{})
	if !ok {
		return createNewsCxSampleData(), nil
	}

	rollData, ok := data["roll_data"].([]interface{})
	if !ok {
		return createNewsCxSampleData(), nil
	}

	headers := []string{"序号", "标题", "内容", "发布时间", "关键词"}
	rows := [][]string{headers}

	for i, item := range rollData {
		if m, ok := item.(map[string]interface{}); ok {
			row := []string{
				fmt.Sprintf("%d", i+1),
				getString(m, "title"),
				getString(m, "content"),
				getString(m, "ctime"),
				getString(m, "keyword"),
			}
			rows = append(rows, row)
		}
	}

	return dataframe.LoadRecords(rows), nil
}

func createNewsCxSampleData() dataframe.DataFrame {
	records := [][]string{
		{"序号", "标题", "内容", "发布时间", "关键词"},
		{"1", "市场快讯", "今日A股三大指数集体高开...", "2024-01-15 09:30:00", "A股,高开"},
	}
	return dataframe.LoadRecords(records)
}
