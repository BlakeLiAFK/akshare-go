package economic

import (
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
	"github.com/tidwall/gjson"
)

// fetchGermanyEconomicData 获取德国经济数据的通用函数
func fetchGermanyEconomicData(indicatorID string) (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"

	params := map[string]string{
		"reportName":  "RPT_ECONOMICVALUE_GER",
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

// MacroGermanyIFO 东方财富-经济数据-德国-IFO商业景气指数
//
// 返回:
//   - dataframe.DataFrame: IFO商业景气指数数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_1_0.html
func MacroGermanyIFO() (dataframe.DataFrame, error) {
	return fetchGermanyEconomicData("EMG00179154")
}

// MacroGermanyCPIMonthly 东方财富-经济数据-德国-消费者物价指数月率终值
//
// 返回:
//   - dataframe.DataFrame: 消费者物价指数月率终值数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_1_1.html
func MacroGermanyCPIMonthly() (dataframe.DataFrame, error) {
	return fetchGermanyEconomicData("EMG00009758")
}

// MacroGermanyCPIYearly 东方财富-经济数据-德国-消费者物价指数年率终值
//
// 返回:
//   - dataframe.DataFrame: 消费者物价指数年率终值数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_1_2.html
func MacroGermanyCPIYearly() (dataframe.DataFrame, error) {
	return fetchGermanyEconomicData("EMG00009756")
}

// MacroGermanyTradeAdjusted 东方财富-经济数据-德国-贸易帐(季调后)
//
// 返回:
//   - dataframe.DataFrame: 贸易帐(季调后)数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_1_3.html
func MacroGermanyTradeAdjusted() (dataframe.DataFrame, error) {
	return fetchGermanyEconomicData("EMG00009753")
}

// MacroGermanyGDP 东方财富-经济数据-德国-GDP
//
// 返回:
//   - dataframe.DataFrame: GDP数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_1_4.html
func MacroGermanyGDP() (dataframe.DataFrame, error) {
	return fetchGermanyEconomicData("EMG00009720")
}

// MacroGermanyRetailSaleMonthly 东方财富-经济数据-德国-实际零售销售月率
//
// 返回:
//   - dataframe.DataFrame: 实际零售销售月率数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_1_5.html
func MacroGermanyRetailSaleMonthly() (dataframe.DataFrame, error) {
	return fetchGermanyEconomicData("EMG01333186")
}

// MacroGermanyRetailSaleYearly 东方财富-经济数据-德国-实际零售销售年率
//
// 返回:
//   - dataframe.DataFrame: 实际零售销售年率数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_1_6.html
func MacroGermanyRetailSaleYearly() (dataframe.DataFrame, error) {
	return fetchGermanyEconomicData("EMG01333192")
}

// MacroGermanyZEW 东方财富-经济数据-德国-ZEW经济景气指数
//
// 返回:
//   - dataframe.DataFrame: ZEW经济景气指数数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_1_7.html
func MacroGermanyZEW() (dataframe.DataFrame, error) {
	return fetchGermanyEconomicData("EMG00172577")
}

// MacroGermanyCore 东方财富-经济数据一览-德国-核心指标数据
//
// 参数:
//   - symbol: 指标代码，默认为 "EMG00179154"
//
// 返回:
//   - dataframe.DataFrame: 指定指标的经济数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_1_0.html
func MacroGermanyCore(symbol string) (dataframe.DataFrame, error) {
	if symbol == "" {
		symbol = "EMG00179154"
	}
	return fetchGermanyEconomicData(symbol)
}
