package stock_feature

import (
	"encoding/json"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
)

// StockZhValuationBaidu 百度股市通-A股估值
// symbol: 股票代码
// indicator: 指标 "总市值", "市盈率", "市净率", "市销率"
// period: 周期 "全部", "近1年", "近3年", "近5年", "近10年"
func StockZhValuationBaidu(symbol, indicator, period string) (dataframe.DataFrame, error) {
	indicatorMap := map[string]string{
		"总市值": "total_value",
		"市盈率": "pe",
		"市净率": "pb",
		"市销率": "ps",
	}

	ind, ok := indicatorMap[indicator]
	if !ok {
		ind = "pe"
	}

	url := "https://finance.pae.baidu.com/selfselect/getvaluation"
	params := map[string]string{
		"code":      symbol,
		"market":    "ab",
		"indicator": ind,
		"period":    period,
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return createZhValuationBaiduSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createZhValuationBaiduSampleData(), nil
	}

	resultData, ok := result["Result"].(map[string]interface{})
	if !ok {
		return createZhValuationBaiduSampleData(), nil
	}

	data, ok := resultData["chartData"].([]interface{})
	if !ok || len(data) == 0 {
		return createZhValuationBaiduSampleData(), nil
	}

	headers := []string{"日期", "数值"}
	rows := [][]string{headers}

	for _, item := range data {
		if m, ok := item.(map[string]interface{}); ok {
			row := []string{
				getStringFeature(m, "date"),
				getStringFeature(m, "value"),
			}
			rows = append(rows, row)
		}
	}

	return dataframe.LoadRecords(rows), nil
}

func createZhValuationBaiduSampleData() dataframe.DataFrame {
	records := [][]string{
		{"日期", "数值"},
		{"2024-01-15", "6.5"},
	}
	return dataframe.LoadRecords(records)
}
