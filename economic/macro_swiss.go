package economic

import (
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
	"github.com/tidwall/gjson"
)

// fetchSwissEconomicData 获取瑞士经济数据的通用函数
func fetchSwissEconomicData(indicatorID string) (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"

	params := map[string]string{
		"reportName":  "RPT_ECONOMICVALUE_CH",
		"columns":     "ALL",
		"filter":      fmt.Sprintf(`(INDICATOR_ID="%s")`, indicatorID),
		"pageNumber":  "1",
		"pageSize":    "5000",
		"sortColumns": "REPORT_DATE",
		"sortTypes":   "-1",
		"source":      "WEB",
		"client":      "WEB",
		"p":           "1",
		"pageNo":      "1",
		"pageNum":     "1",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	// 解析 JSON 响应
	dataArray := gjson.Get(resp.String(), "result.data").Array()
	if len(dataArray) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	// 定义列名
	columns := []string{"时间", "前值", "现值", "发布日期"}

	// 构建记录
	var records [][]string
	records = append(records, columns)

	for _, item := range dataArray {
		reportDateCH := item.Get("REPORT_DATE_CH").String()
		preValue := item.Get("PRE_VALUE").String()
		value := item.Get("VALUE").String()
		publishDate := item.Get("PUBLISH_DATE").String()

		if len(publishDate) > 10 {
			publishDate = publishDate[:10]
		}

		record := []string{reportDateCH, preValue, value, publishDate}
		records = append(records, record)
	}

	df := dataframe.LoadRecords(records)
	df = df.Arrange(dataframe.Sort("发布日期"))

	return df, nil
}

// MacroSwissSVME 东方财富-经济数据-瑞士-SVME采购经理人指数
//
// 返回:
//   - dataframe.DataFrame: SVME采购经理人指数数据
//   - error: 错误信息
//
// 数据源: http://data.eastmoney.com/cjsj/foreign_2_0.html
func MacroSwissSVME() (dataframe.DataFrame, error) {
	return fetchSwissEconomicData("EMG00341602")
}

// MacroSwissTrade 东方财富-经济数据-瑞士-贸易帐
//
// 返回:
//   - dataframe.DataFrame: 贸易帐数据
//   - error: 错误信息
//
// 数据源: http://data.eastmoney.com/cjsj/foreign_2_1.html
func MacroSwissTrade() (dataframe.DataFrame, error) {
	return fetchSwissEconomicData("EMG00341603")
}

// MacroSwissCPIYearly 东方财富-经济数据-瑞士-消费者物价指数年率
//
// 返回:
//   - dataframe.DataFrame: 消费者物价指数年率数据
//   - error: 错误信息
//
// 数据源: http://data.eastmoney.com/cjsj/foreign_2_2.html
func MacroSwissCPIYearly() (dataframe.DataFrame, error) {
	return fetchSwissEconomicData("EMG00341604")
}

// MacroSwissGDPQuarterly 东方财富-经济数据-瑞士-GDP季率
//
// 返回:
//   - dataframe.DataFrame: GDP季率数据
//   - error: 错误信息
//
// 数据源: http://data.eastmoney.com/cjsj/foreign_2_3.html
func MacroSwissGDPQuarterly() (dataframe.DataFrame, error) {
	return fetchSwissEconomicData("EMG00341600")
}

// MacroSwissGDPYearly 东方财富-经济数据-瑞士-GDP年率
//
// 返回:
//   - dataframe.DataFrame: GDP年率数据
//   - error: 错误信息
//
// 数据源: http://data.eastmoney.com/cjsj/foreign_2_4.html
func MacroSwissGDPYearly() (dataframe.DataFrame, error) {
	return fetchSwissEconomicData("EMG00341601")
}

// MacroSwissBankRate 东方财富-经济数据-瑞士-央行公布利率决议
//
// 返回:
//   - dataframe.DataFrame: 央行公布利率决议数据
//   - error: 错误信息
//
// 数据源: http://data.eastmoney.com/cjsj/foreign_2_5.html
func MacroSwissBankRate() (dataframe.DataFrame, error) {
	return fetchSwissEconomicData("EMG00341606")
}

// MacroSwissCore 东方财富-经济数据一览-瑞士-核心指标数据
//
// 参数:
//   - symbol: 指标代码，默认为 "EMG00341602"
//
// 返回:
//   - dataframe.DataFrame: 指定指标的经济数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_1_0.html
func MacroSwissCore(symbol string) (dataframe.DataFrame, error) {
	if symbol == "" {
		symbol = "EMG00341602"
	}
	return fetchSwissEconomicData(symbol)
}
