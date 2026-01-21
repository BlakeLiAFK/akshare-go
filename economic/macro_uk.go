package economic

import (
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
	"github.com/tidwall/gjson"
)

// fetchUKEconomicData 获取英国经济数据的通用函数
func fetchUKEconomicData(indicatorID string) (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"

	params := map[string]string{
		"reportName":  "RPT_ECONOMICVALUE_BRITAIN",
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

// MacroUKHalifaxMonthly 东方财富-经济数据-英国-Halifax房价指数月率
//
// 返回:
//   - dataframe.DataFrame: Halifax房价指数月率数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_4_0.html
func MacroUKHalifaxMonthly() (dataframe.DataFrame, error) {
	return fetchUKEconomicData("EMG00342256")
}

// MacroUKHalifaxYearly 东方财富-经济数据-英国-Halifax房价指数年率
//
// 返回:
//   - dataframe.DataFrame: Halifax房价指数年率数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_4_1.html
func MacroUKHalifaxYearly() (dataframe.DataFrame, error) {
	return fetchUKEconomicData("EMG00010370")
}

// MacroUKTrade 东方财富-经济数据-英国-贸易帐
//
// 返回:
//   - dataframe.DataFrame: 贸易帐数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_4_2.html
func MacroUKTrade() (dataframe.DataFrame, error) {
	return fetchUKEconomicData("EMG00158309")
}

// MacroUKBankRate 东方财富-经济数据-英国-央行公布利率决议
//
// 返回:
//   - dataframe.DataFrame: 央行公布利率决议数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_4_3.html
func MacroUKBankRate() (dataframe.DataFrame, error) {
	return fetchUKEconomicData("EMG00342253")
}

// MacroUKCoreCPIYearly 东方财富-经济数据-英国-核心消费者物价指数年率
//
// 返回:
//   - dataframe.DataFrame: 核心消费者物价指数年率数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_4_4.html
func MacroUKCoreCPIYearly() (dataframe.DataFrame, error) {
	return fetchUKEconomicData("EMG00010279")
}

// MacroUKCoreCPIMonthly 东方财富-经济数据-英国-核心消费者物价指数月率
//
// 返回:
//   - dataframe.DataFrame: 核心消费者物价指数月率数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_4_5.html
func MacroUKCoreCPIMonthly() (dataframe.DataFrame, error) {
	return fetchUKEconomicData("EMG00010291")
}

// MacroUKCPIYearly 东方财富-经济数据-英国-消费者物价指数年率
//
// 返回:
//   - dataframe.DataFrame: 消费者物价指数年率数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_4_6.html
func MacroUKCPIYearly() (dataframe.DataFrame, error) {
	return fetchUKEconomicData("EMG00010267")
}

// MacroUKCPIMonthly 东方财富-经济数据-英国-消费者物价指数月率
//
// 返回:
//   - dataframe.DataFrame: 消费者物价指数月率数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_4_7.html
func MacroUKCPIMonthly() (dataframe.DataFrame, error) {
	return fetchUKEconomicData("EMG00010291")
}

// MacroUKRetailMonthly 东方财富-经济数据-英国-零售销售月率
//
// 返回:
//   - dataframe.DataFrame: 零售销售月率数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_4_8.html
func MacroUKRetailMonthly() (dataframe.DataFrame, error) {
	return fetchUKEconomicData("EMG00158298")
}

// MacroUKRetailYearly 东方财富-经济数据-英国-零售销售年率
//
// 返回:
//   - dataframe.DataFrame: 零售销售年率数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_4_9.html
func MacroUKRetailYearly() (dataframe.DataFrame, error) {
	return fetchUKEconomicData("EMG00158297")
}

// MacroUKRightmoveYearly 东方财富-经济数据-英国-Rightmove房价指数年率
//
// 返回:
//   - dataframe.DataFrame: Rightmove房价指数年率数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_4_10.html
func MacroUKRightmoveYearly() (dataframe.DataFrame, error) {
	return fetchUKEconomicData("EMG00341608")
}

// MacroUKRightmoveMonthly 东方财富-经济数据-英国-Rightmove房价指数月率
//
// 返回:
//   - dataframe.DataFrame: Rightmove房价指数月率数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_4_11.html
func MacroUKRightmoveMonthly() (dataframe.DataFrame, error) {
	return fetchUKEconomicData("EMG00341607")
}

// MacroUKGDPQuarterly 东方财富-经济数据-英国-GDP季率初值
//
// 返回:
//   - dataframe.DataFrame: GDP季率初值数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_4_12.html
func MacroUKGDPQuarterly() (dataframe.DataFrame, error) {
	return fetchUKEconomicData("EMG00158277")
}

// MacroUKGDPYearly 东方财富-经济数据-英国-GDP年率初值
//
// 返回:
//   - dataframe.DataFrame: GDP年率初值数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_4_13.html
func MacroUKGDPYearly() (dataframe.DataFrame, error) {
	return fetchUKEconomicData("EMG00158276")
}

// MacroUKUnemploymentRate 东方财富-经济数据-英国-失业率
//
// 返回:
//   - dataframe.DataFrame: 失业率数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_4_14.html
func MacroUKUnemploymentRate() (dataframe.DataFrame, error) {
	return fetchUKEconomicData("EMG00010348")
}

// MacroUKCore 东方财富-经济数据一览-英国-核心指标数据
//
// 参数:
//   - symbol: 指标代码，默认为 "EMG00010348"
//
// 返回:
//   - dataframe.DataFrame: 指定指标的经济数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_4_0.html
func MacroUKCore(symbol string) (dataframe.DataFrame, error) {
	if symbol == "" {
		symbol = "EMG00010348"
	}
	return fetchUKEconomicData(symbol)
}
