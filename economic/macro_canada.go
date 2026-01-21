package economic

import (
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
	"github.com/tidwall/gjson"
)

// fetchCanadaEconomicData 获取加拿大经济数据的通用函数
func fetchCanadaEconomicData(indicatorID string) (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"

	params := map[string]string{
		"reportName":  "RPT_ECONOMICVALUE_CA",
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

// MacroCanadaNewHouseRate 东方财富-经济数据-加拿大-新屋开工
//
// 返回:
//   - dataframe.DataFrame: 新屋开工数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_7_0.html
func MacroCanadaNewHouseRate() (dataframe.DataFrame, error) {
	return fetchCanadaEconomicData("EMG00342247")
}

// MacroCanadaUnemploymentRate 东方财富-经济数据-加拿大-失业率
//
// 返回:
//   - dataframe.DataFrame: 失业率数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_7_1.html
func MacroCanadaUnemploymentRate() (dataframe.DataFrame, error) {
	return fetchCanadaEconomicData("EMG00157746")
}

// MacroCanadaTrade 东方财富-经济数据-加拿大-贸易帐
//
// 返回:
//   - dataframe.DataFrame: 贸易帐数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_7_2.html
func MacroCanadaTrade() (dataframe.DataFrame, error) {
	return fetchCanadaEconomicData("EMG00102022")
}

// MacroCanadaRetailRateMonthly 东方财富-经济数据-加拿大-零售销售月率
//
// 返回:
//   - dataframe.DataFrame: 零售销售月率数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_7_3.html
func MacroCanadaRetailRateMonthly() (dataframe.DataFrame, error) {
	return fetchCanadaEconomicData("EMG01337094")
}

// MacroCanadaBankRate 东方财富-经济数据-加拿大-央行公布利率决议
//
// 返回:
//   - dataframe.DataFrame: 央行公布利率决议数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_7_4.html
func MacroCanadaBankRate() (dataframe.DataFrame, error) {
	return fetchCanadaEconomicData("EMG00342248")
}

// MacroCanadaCoreCPIYearly 东方财富-经济数据-加拿大-核心消费者物价指数年率
//
// 返回:
//   - dataframe.DataFrame: 核心消费者物价指数年率数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_7_5.html
func MacroCanadaCoreCPIYearly() (dataframe.DataFrame, error) {
	return fetchCanadaEconomicData("EMG00102030")
}

// MacroCanadaCoreCPIMonthly 东方财富-经济数据-加拿大-核心消费者物价指数月率
//
// 返回:
//   - dataframe.DataFrame: 核心消费者物价指数月率数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_7_6.html
func MacroCanadaCoreCPIMonthly() (dataframe.DataFrame, error) {
	return fetchCanadaEconomicData("EMG00102044")
}

// MacroCanadaCPIYearly 东方财富-经济数据-加拿大-消费者物价指数年率
//
// 返回:
//   - dataframe.DataFrame: 消费者物价指数年率数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_7_7.html
func MacroCanadaCPIYearly() (dataframe.DataFrame, error) {
	return fetchCanadaEconomicData("EMG00102029")
}

// MacroCanadaCPIMonthly 东方财富-经济数据-加拿大-消费者物价指数月率
//
// 返回:
//   - dataframe.DataFrame: 消费者物价指数月率数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_7_8.html
func MacroCanadaCPIMonthly() (dataframe.DataFrame, error) {
	return fetchCanadaEconomicData("EMG00158719")
}

// MacroCanadaGDPMonthly 东方财富-经济数据-加拿大-GDP月率
//
// 返回:
//   - dataframe.DataFrame: GDP月率数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_7_9.html
func MacroCanadaGDPMonthly() (dataframe.DataFrame, error) {
	return fetchCanadaEconomicData("EMG00159259")
}
