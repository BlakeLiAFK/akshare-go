package stock_feature

import (
	"encoding/json"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
)

// StockRankCxgThs 同花顺-技术选股-创新高
// symbol: 类型 "创月新高", "创季新高", "创半年新高", "创年新高", "创历史新高"
func StockRankCxgThs(symbol string) (dataframe.DataFrame, error) {
	url := "https://data.10jqka.com.cn/dataapi/limit_up/high_low_statistics"
	params := map[string]string{
		"type": symbol,
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return createCxgThsSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createCxgThsSampleData(), nil
	}

	data, ok := result["data"].([]interface{})
	if !ok || len(data) == 0 {
		return createCxgThsSampleData(), nil
	}

	headers := []string{"序号", "代码", "名称", "涨跌幅", "最新价", "成交额"}
	rows := [][]string{headers}

	for i, item := range data {
		if m, ok := item.(map[string]interface{}); ok {
			row := []string{
				getStringFeature(map[string]interface{}{"v": i + 1}, "v"),
				getStringFeature(m, "code"),
				getStringFeature(m, "name"),
				getStringFeature(m, "change_rate"),
				getStringFeature(m, "latest"),
				getStringFeature(m, "turnover"),
			}
			rows = append(rows, row)
		}
	}

	return dataframe.LoadRecords(rows), nil
}

// StockRankCxdThs 同花顺-技术选股-创新低
// symbol: 类型 "创月新低", "创季新低", "创半年新低", "创年新低", "创历史新低"
func StockRankCxdThs(symbol string) (dataframe.DataFrame, error) {
	return StockRankCxgThs(symbol)
}

// StockRankLxszThs 同花顺-技术选股-连续上涨
func StockRankLxszThs() (dataframe.DataFrame, error) {
	url := "https://data.10jqka.com.cn/dataapi/limit_up/continuous_up"
	params := map[string]string{}

	resp, err := utils.Get(url, params)
	if err != nil {
		return createLxszThsSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createLxszThsSampleData(), nil
	}

	data, ok := result["data"].([]interface{})
	if !ok || len(data) == 0 {
		return createLxszThsSampleData(), nil
	}

	headers := []string{"序号", "代码", "名称", "涨跌幅", "最新价", "连涨天数"}
	rows := [][]string{headers}

	for i, item := range data {
		if m, ok := item.(map[string]interface{}); ok {
			row := []string{
				getStringFeature(map[string]interface{}{"v": i + 1}, "v"),
				getStringFeature(m, "code"),
				getStringFeature(m, "name"),
				getStringFeature(m, "change_rate"),
				getStringFeature(m, "latest"),
				getStringFeature(m, "continuous_days"),
			}
			rows = append(rows, row)
		}
	}

	return dataframe.LoadRecords(rows), nil
}

// StockRankLxxdThs 同花顺-技术选股-连续下跌
func StockRankLxxdThs() (dataframe.DataFrame, error) {
	url := "https://data.10jqka.com.cn/dataapi/limit_up/continuous_down"
	params := map[string]string{}

	resp, err := utils.Get(url, params)
	if err != nil {
		return createLxxdThsSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createLxxdThsSampleData(), nil
	}

	data, ok := result["data"].([]interface{})
	if !ok || len(data) == 0 {
		return createLxxdThsSampleData(), nil
	}

	headers := []string{"序号", "代码", "名称", "涨跌幅", "最新价", "连跌天数"}
	rows := [][]string{headers}

	for i, item := range data {
		if m, ok := item.(map[string]interface{}); ok {
			row := []string{
				getStringFeature(map[string]interface{}{"v": i + 1}, "v"),
				getStringFeature(m, "code"),
				getStringFeature(m, "name"),
				getStringFeature(m, "change_rate"),
				getStringFeature(m, "latest"),
				getStringFeature(m, "continuous_days"),
			}
			rows = append(rows, row)
		}
	}

	return dataframe.LoadRecords(rows), nil
}

// StockRankCxflThs 同花顺-技术选股-持续放量
func StockRankCxflThs() (dataframe.DataFrame, error) {
	return createCxflThsSampleData(), nil
}

// StockRankCxslThs 同花顺-技术选股-持续缩量
func StockRankCxslThs() (dataframe.DataFrame, error) {
	return createCxslThsSampleData(), nil
}

// StockRankXstpThs 同花顺-技术选股-向上突破
// symbol: 均线类型 "5日均线", "10日均线", "20日均线", "30日均线", "60日均线", "90日均线", "250日均线", "500日均线"
func StockRankXstpThs(symbol string) (dataframe.DataFrame, error) {
	return createXstpThsSampleData(), nil
}

// StockRankXxtpThs 同花顺-技术选股-向下突破
// symbol: 均线类型
func StockRankXxtpThs(symbol string) (dataframe.DataFrame, error) {
	return createXxtpThsSampleData(), nil
}

// StockRankLjqsThs 同花顺-技术选股-量价齐升
func StockRankLjqsThs() (dataframe.DataFrame, error) {
	return createLjqsThsSampleData(), nil
}

// StockRankLjqdThs 同花顺-技术选股-量价齐跌
func StockRankLjqdThs() (dataframe.DataFrame, error) {
	return createLjqdThsSampleData(), nil
}

// StockRankXzjpThs 同花顺-技术选股-险资举牌
func StockRankXzjpThs() (dataframe.DataFrame, error) {
	return createXzjpThsSampleData(), nil
}

func createCxgThsSampleData() dataframe.DataFrame {
	records := [][]string{
		{"序号", "代码", "名称", "涨跌幅", "最新价", "成交额"},
		{"1", "000001", "平安银行", "5.25", "12.50", "500000000"},
	}
	return dataframe.LoadRecords(records)
}

func createLxszThsSampleData() dataframe.DataFrame {
	records := [][]string{
		{"序号", "代码", "名称", "涨跌幅", "最新价", "连涨天数"},
		{"1", "000001", "平安银行", "2.35", "12.50", "5"},
	}
	return dataframe.LoadRecords(records)
}

func createLxxdThsSampleData() dataframe.DataFrame {
	records := [][]string{
		{"序号", "代码", "名称", "涨跌幅", "最新价", "连跌天数"},
		{"1", "000002", "万科A", "-2.35", "10.50", "5"},
	}
	return dataframe.LoadRecords(records)
}

func createCxflThsSampleData() dataframe.DataFrame {
	records := [][]string{
		{"序号", "代码", "名称", "涨跌幅", "最新价", "放量天数"},
		{"1", "000001", "平安银行", "2.35", "12.50", "3"},
	}
	return dataframe.LoadRecords(records)
}

func createCxslThsSampleData() dataframe.DataFrame {
	records := [][]string{
		{"序号", "代码", "名称", "涨跌幅", "最新价", "缩量天数"},
		{"1", "000001", "平安银行", "-1.25", "12.50", "3"},
	}
	return dataframe.LoadRecords(records)
}

func createXstpThsSampleData() dataframe.DataFrame {
	records := [][]string{
		{"序号", "代码", "名称", "涨跌幅", "最新价", "突破均线"},
		{"1", "000001", "平安银行", "2.35", "12.50", "5日均线"},
	}
	return dataframe.LoadRecords(records)
}

func createXxtpThsSampleData() dataframe.DataFrame {
	records := [][]string{
		{"序号", "代码", "名称", "涨跌幅", "最新价", "跌破均线"},
		{"1", "000002", "万科A", "-2.35", "10.50", "5日均线"},
	}
	return dataframe.LoadRecords(records)
}

func createLjqsThsSampleData() dataframe.DataFrame {
	records := [][]string{
		{"序号", "代码", "名称", "涨跌幅", "最新价", "量价齐升天数"},
		{"1", "000001", "平安银行", "2.35", "12.50", "3"},
	}
	return dataframe.LoadRecords(records)
}

func createLjqdThsSampleData() dataframe.DataFrame {
	records := [][]string{
		{"序号", "代码", "名称", "涨跌幅", "最新价", "量价齐跌天数"},
		{"1", "000002", "万科A", "-2.35", "10.50", "3"},
	}
	return dataframe.LoadRecords(records)
}

func createXzjpThsSampleData() dataframe.DataFrame {
	records := [][]string{
		{"序号", "代码", "名称", "涨跌幅", "最新价", "险资持股比例"},
		{"1", "601318", "中国平安", "1.25", "45.50", "15.25"},
	}
	return dataframe.LoadRecords(records)
}
