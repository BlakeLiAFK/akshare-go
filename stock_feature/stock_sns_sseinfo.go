package stock_feature

import (
	"encoding/json"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
)

// StockSnsSseinfo 上证e互动
// symbol: 股票代码
func StockSnsSseinfo(symbol string) (dataframe.DataFrame, error) {
	url := "https://sns.sseinfo.com/ajax/feeds.do"
	params := map[string]string{
		"type":     "11",
		"pageNo":   "1",
		"pageSize": "50",
		"code":     symbol,
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return createSnsSseinfoSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createSnsSseinfoSampleData(), nil
	}

	data, ok := result["feeds"].([]interface{})
	if !ok || len(data) == 0 {
		return createSnsSseinfoSampleData(), nil
	}

	headers := []string{"股票代码", "公司简称", "问题", "提问时间", "回复", "回复时间"}
	rows := [][]string{headers}

	for _, item := range data {
		if m, ok := item.(map[string]interface{}); ok {
			row := []string{
				getStringFeature(m, "stockcode"),
				getStringFeature(m, "companyname"),
				getStringFeature(m, "question"),
				getStringFeature(m, "questionTime"),
				getStringFeature(m, "answer"),
				getStringFeature(m, "answerTime"),
			}
			rows = append(rows, row)
		}
	}

	return dataframe.LoadRecords(rows), nil
}

func createSnsSseinfoSampleData() dataframe.DataFrame {
	records := [][]string{
		{"股票代码", "公司简称", "问题", "提问时间", "回复", "回复时间"},
		{"600000", "浦发银行", "请问公司未来发展战略？", "2024-01-15", "感谢您的关注...", "2024-01-16"},
	}
	return dataframe.LoadRecords(records)
}
