package economic

import (
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
	"github.com/tidwall/gjson"
)

// fetchAustraliaEconomicData 获取澳大利亚经济数据的通用函数
func fetchAustraliaEconomicData(indicatorID string) (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"

	params := map[string]string{
		"reportName":  "RPT_ECONOMICVALUE_AUSTRALIA",
		"columns":     "ALL",
		"filter":      fmt.Sprintf(`(INDICATOR_ID="%s")`, indicatorID),
		"pageNumber":  "1",
		"pageSize":    "2000",
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
		// 数据是对象格式，提取各字段
		reportDateCH := item.Get("REPORT_DATE_CH").String() // 时间
		preValue := item.Get("PRE_VALUE").String()          // 前值
		value := item.Get("VALUE").String()                 // 现值
		publishDate := item.Get("PUBLISH_DATE").String()    // 发布日期

		// 格式化发布日期 (去掉时间部分)
		if len(publishDate) > 10 {
			publishDate = publishDate[:10]
		}

		record := []string{
			reportDateCH,
			preValue,
			value,
			publishDate,
		}
		records = append(records, record)
	}

	// 创建 DataFrame
	df := dataframe.LoadRecords(records)

	// 按发布日期排序
	df = df.Arrange(dataframe.Sort("发布日期"))

	return df, nil
}

// MacroAustraliaRetailRateMonthly 东方财富-经济数据-澳大利亚-零售销售月率
//
// 返回:
//   - dataframe.DataFrame: 零售销售月率数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_5_0.html
func MacroAustraliaRetailRateMonthly() (dataframe.DataFrame, error) {
	return fetchAustraliaEconomicData("EMG00152903")
}

// MacroAustraliaTrade 东方财富-经济数据-澳大利亚-贸易帐
//
// 返回:
//   - dataframe.DataFrame: 贸易帐数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_5_1.html
func MacroAustraliaTrade() (dataframe.DataFrame, error) {
	return fetchAustraliaEconomicData("EMG00152793")
}

// MacroAustraliaUnemploymentRate 东方财富-经济数据-澳大利亚-失业率
//
// 返回:
//   - dataframe.DataFrame: 失业率数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_5_2.html
func MacroAustraliaUnemploymentRate() (dataframe.DataFrame, error) {
	return fetchAustraliaEconomicData("EMG00101141")
}

// MacroAustraliaPPIQuarterly 东方财富-经济数据-澳大利亚-生产者物价指数季率
//
// 返回:
//   - dataframe.DataFrame: 生产者物价指数季率数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_5_3.html
func MacroAustraliaPPIQuarterly() (dataframe.DataFrame, error) {
	return fetchAustraliaEconomicData("EMG00152722")
}

// MacroAustraliaCPIQuarterly 东方财富-经济数据-澳大利亚-消费者物价指数季率
//
// 返回:
//   - dataframe.DataFrame: 消费者物价指数季率数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_5_4.html
func MacroAustraliaCPIQuarterly() (dataframe.DataFrame, error) {
	return fetchAustraliaEconomicData("EMG00101104")
}

// MacroAustraliaCPIYearly 东方财富-经济数据-澳大利亚-消费者物价指数年率
//
// 返回:
//   - dataframe.DataFrame: 消费者物价指数年率数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_5_5.html
func MacroAustraliaCPIYearly() (dataframe.DataFrame, error) {
	return fetchAustraliaEconomicData("EMG00101093")
}

// MacroAustraliaBankRate 东方财富-经济数据-澳大利亚-央行公布利率决议
//
// 返回:
//   - dataframe.DataFrame: 央行公布利率决议数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_5_6.html
func MacroAustraliaBankRate() (dataframe.DataFrame, error) {
	return fetchAustraliaEconomicData("EMG00342255")
}
