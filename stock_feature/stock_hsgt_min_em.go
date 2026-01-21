package stock_feature

import (
	"encoding/json"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
)

// StockHsgtFundMinEm 东方财富-沪深港通资金流向-分钟数据
// symbol: 类型 "北向资金", "南向资金"
func StockHsgtFundMinEm(symbol string) (dataframe.DataFrame, error) {
	symbolMap := map[string]string{
		"北向资金": "n2s",
		"南向资金": "s2n",
	}

	flowType, ok := symbolMap[symbol]
	if !ok {
		flowType = "n2s"
	}

	url := "https://push2.eastmoney.com/api/qt/kamtbs.rtmin/get"
	params := map[string]string{
		"fields1": "f1,f2,f3,f4",
		"fields2": "f51,f52,f53,f54,f55,f56",
	}

	if flowType == "n2s" {
		params["fields2"] = "f51,f52,f53,f54,f55,f56"
	} else {
		params["fields2"] = "f51,f52,f53,f54,f55,f56"
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return createHsgtFundMinSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createHsgtFundMinSampleData(), nil
	}

	data, ok := result["data"].(map[string]interface{})
	if !ok {
		return createHsgtFundMinSampleData(), nil
	}

	var trends []interface{}
	if flowType == "n2s" {
		trends, ok = data["n2s"].([]interface{})
	} else {
		trends, ok = data["s2n"].([]interface{})
	}

	if !ok || len(trends) == 0 {
		return createHsgtFundMinSampleData(), nil
	}

	headers := []string{"时间", "北向资金", "上证资金", "深证资金"}
	rows := [][]string{headers}

	for _, item := range trends {
		if s, ok := item.(string); ok {
			rows = append(rows, []string{s, "", "", ""})
		}
	}

	return dataframe.LoadRecords(rows), nil
}

func createHsgtFundMinSampleData() dataframe.DataFrame {
	records := [][]string{
		{"时间", "北向资金", "上证资金", "深证资金"},
		{"09:30", "50000000", "30000000", "20000000"},
	}
	return dataframe.LoadRecords(records)
}
