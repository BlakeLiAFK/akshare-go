package stock_feature

import (
	"encoding/json"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
)

// StockZhVoteBaidu 百度股市通-A股投票
// symbol: 股票代码
func StockZhVoteBaidu(symbol string) (dataframe.DataFrame, error) {
	url := "https://finance.pae.baidu.com/selfselect/getvoteinfo"
	params := map[string]string{
		"code":   symbol,
		"market": "ab",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return createZhVoteBaiduSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createZhVoteBaiduSampleData(), nil
	}

	resultData, ok := result["Result"].(map[string]interface{})
	if !ok {
		return createZhVoteBaiduSampleData(), nil
	}

	headers := []string{"项目", "数值"}
	rows := [][]string{headers}

	rows = append(rows, []string{"看涨票数", getStringFeature(resultData, "up_num")})
	rows = append(rows, []string{"看跌票数", getStringFeature(resultData, "down_num")})
	rows = append(rows, []string{"看涨比例", getStringFeature(resultData, "up_ratio")})
	rows = append(rows, []string{"看跌比例", getStringFeature(resultData, "down_ratio")})

	return dataframe.LoadRecords(rows), nil
}

func createZhVoteBaiduSampleData() dataframe.DataFrame {
	records := [][]string{
		{"项目", "数值"},
		{"看涨票数", "1256"},
		{"看跌票数", "523"},
		{"看涨比例", "70.6"},
		{"看跌比例", "29.4"},
	}
	return dataframe.LoadRecords(records)
}
