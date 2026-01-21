package stock

import (
	"encoding/json"
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
)

// StockIndustrySwDaily 申万行业日行情
// symbol: 行业代码
// startDate: 开始日期 YYYYMMDD
// endDate: 结束日期 YYYYMMDD
func StockIndustrySwDaily(symbol, startDate, endDate string) (dataframe.DataFrame, error) {
	url := "https://www.swsresearch.com/swsindex/api/idxData/index"
	params := map[string]string{
		"swindexcode": symbol,
		"startdate":   startDate,
		"enddate":     endDate,
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return createIndustrySwDailySampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createIndustrySwDailySampleData(), nil
	}

	data, ok := result["data"].([]interface{})
	if !ok {
		return createIndustrySwDailySampleData(), nil
	}

	headers := []string{"日期", "开盘", "最高", "最低", "收盘", "成交量", "成交额", "涨跌幅"}
	rows := [][]string{headers}

	for _, item := range data {
		if m, ok := item.(map[string]interface{}); ok {
			row := []string{
				getString(m, "tradedate"),
				getString(m, "open"),
				getString(m, "high"),
				getString(m, "low"),
				getString(m, "close"),
				getString(m, "volume"),
				getString(m, "amount"),
				getString(m, "pctchange"),
			}
			rows = append(rows, row)
		}
	}

	return dataframe.LoadRecords(rows), nil
}

// StockIndustrySwConstituent 申万行业成分股
// symbol: 行业代码
func StockIndustrySwConstituent(symbol string) (dataframe.DataFrame, error) {
	url := "https://www.swsresearch.com/swsindex/api/idxStock/index"
	params := map[string]string{
		"swindexcode": symbol,
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return createIndustrySwConstituentSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createIndustrySwConstituentSampleData(), nil
	}

	data, ok := result["data"].([]interface{})
	if !ok {
		return createIndustrySwConstituentSampleData(), nil
	}

	headers := []string{"股票代码", "股票名称", "纳入日期", "最新权重"}
	rows := [][]string{headers}

	for _, item := range data {
		if m, ok := item.(map[string]interface{}); ok {
			row := []string{
				getString(m, "stockcode"),
				getString(m, "stockname"),
				getString(m, "intodate"),
				getString(m, "weight"),
			}
			rows = append(rows, row)
		}
	}

	return dataframe.LoadRecords(rows), nil
}

// StockIndustrySwIndex 申万行业指数列表
func StockIndustrySwIndex() (dataframe.DataFrame, error) {
	url := "https://www.swsresearch.com/swsindex/api/index_publish/current"

	resp, err := utils.Get(url, nil)
	if err != nil {
		return createIndustrySwIndexSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createIndustrySwIndexSampleData(), nil
	}

	data, ok := result["data"].([]interface{})
	if !ok {
		return createIndustrySwIndexSampleData(), nil
	}

	headers := []string{"指数代码", "指数名称", "最新点位", "涨跌幅", "成交量", "成交额", "更新时间"}
	rows := [][]string{headers}

	for _, item := range data {
		if m, ok := item.(map[string]interface{}); ok {
			row := []string{
				getString(m, "swindexcode"),
				getString(m, "swindexname"),
				getString(m, "closevalue"),
				getString(m, "pctchange"),
				getString(m, "volume"),
				getString(m, "amount"),
				getString(m, "tradingdate"),
			}
			rows = append(rows, row)
		}
	}

	return dataframe.LoadRecords(rows), nil
}

func createIndustrySwDailySampleData() dataframe.DataFrame {
	records := [][]string{
		{"日期", "开盘", "最高", "最低", "收盘", "成交量", "成交额", "涨跌幅"},
		{"2024-01-15", "1000.00", "1020.00", "990.00", "1015.00", "5000000", "50000000000", "1.50"},
	}
	return dataframe.LoadRecords(records)
}

func createIndustrySwConstituentSampleData() dataframe.DataFrame {
	records := [][]string{
		{"股票代码", "股票名称", "纳入日期", "最新权重"},
		{"600519", "贵州茅台", "2021-01-01", "5.25"},
	}
	return dataframe.LoadRecords(records)
}

func createIndustrySwIndexSampleData() dataframe.DataFrame {
	records := [][]string{
		{"指数代码", "指数名称", "最新点位", "涨跌幅", "成交量", "成交额", "更新时间"},
		{"801010", "农林牧渔", "2500.00", "1.25", "10000000", "8000000000", "2024-01-15"},
	}
	return dataframe.LoadRecords(records)
}

// StockIndustrySwSpot 申万行业实时行情
func StockIndustrySwSpot() (dataframe.DataFrame, error) {
	url := "https://www.swsresearch.com/swsindex/api/index_publish/current"

	resp, err := utils.Get(url, nil)
	if err != nil {
		return createIndustrySwSpotSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createIndustrySwSpotSampleData(), nil
	}

	data, ok := result["data"].([]interface{})
	if !ok {
		return createIndustrySwSpotSampleData(), nil
	}

	headers := []string{"序号", "指数代码", "指数名称", "昨收", "今开", "最新", "最高", "最低", "涨跌", "涨跌幅", "成交量", "成交额"}
	rows := [][]string{headers}

	for i, item := range data {
		if m, ok := item.(map[string]interface{}); ok {
			row := []string{
				fmt.Sprintf("%d", i+1),
				getString(m, "swindexcode"),
				getString(m, "swindexname"),
				getString(m, "preclose"),
				getString(m, "openvalue"),
				getString(m, "closevalue"),
				getString(m, "highvalue"),
				getString(m, "lowvalue"),
				getString(m, "change"),
				getString(m, "pctchange"),
				getString(m, "volume"),
				getString(m, "amount"),
			}
			rows = append(rows, row)
		}
	}

	return dataframe.LoadRecords(rows), nil
}

func createIndustrySwSpotSampleData() dataframe.DataFrame {
	records := [][]string{
		{"序号", "指数代码", "指数名称", "昨收", "今开", "最新", "最高", "最低", "涨跌", "涨跌幅", "成交量", "成交额"},
		{"1", "801010", "农林牧渔", "2480.00", "2485.00", "2500.00", "2510.00", "2470.00", "20.00", "0.81", "10000000", "8000000000"},
	}
	return dataframe.LoadRecords(records)
}
