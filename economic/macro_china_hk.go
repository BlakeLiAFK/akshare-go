package economic

import (
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
	"github.com/tidwall/gjson"
)

// fetchChinaHKEconomicData 获取中国香港经济数据的通用函数
func fetchChinaHKEconomicData(indicatorID string) (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"

	params := map[string]string{
		"reportName":  "RPT_ECONOMICVALUE_HK",
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

// MacroChinaHKCPI 东方财富-经济数据一览-中国香港-消费者物价指数
//
// 返回:
//   - dataframe.DataFrame: 消费者物价指数数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_8_0.html
func MacroChinaHKCPI() (dataframe.DataFrame, error) {
	return fetchChinaHKEconomicData("EMG01336996")
}

// MacroChinaHKCPIRatio 东方财富-经济数据一览-中国香港-消费者物价指数年率
//
// 返回:
//   - dataframe.DataFrame: 消费者物价指数年率数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_8_1.html
func MacroChinaHKCPIRatio() (dataframe.DataFrame, error) {
	return fetchChinaHKEconomicData("EMG00059282")
}

// MacroChinaHKRateOfUnemployment 东方财富-经济数据一览-中国香港-失业率
//
// 返回:
//   - dataframe.DataFrame: 失业率数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_8_2.html
func MacroChinaHKRateOfUnemployment() (dataframe.DataFrame, error) {
	return fetchChinaHKEconomicData("EMG00059647")
}

// MacroChinaHKGBP 东方财富-经济数据一览-中国香港-香港GDP
//
// 返回:
//   - dataframe.DataFrame: 香港GDP数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_8_3.html
func MacroChinaHKGBP() (dataframe.DataFrame, error) {
	return fetchChinaHKEconomicData("EMG01337008")
}

// MacroChinaHKGBPRatio 东方财富-经济数据一览-中国香港-香港GDP同比
//
// 返回:
//   - dataframe.DataFrame: 香港GDP同比数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_8_4.html
func MacroChinaHKGBPRatio() (dataframe.DataFrame, error) {
	return fetchChinaHKEconomicData("EMG01337009")
}

// MacroChinaHKBuildingVolume 东方财富-经济数据一览-中国香港-香港楼宇买卖合约数量
//
// 返回:
//   - dataframe.DataFrame: 香港楼宇买卖合约数量数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_8_5.html
func MacroChinaHKBuildingVolume() (dataframe.DataFrame, error) {
	return fetchChinaHKEconomicData("EMG00158055")
}

// MacroChinaHKBuildingAmount 东方财富-经济数据一览-中国香港-香港楼宇买卖合约成交金额
//
// 返回:
//   - dataframe.DataFrame: 香港楼宇买卖合约成交金额数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_8_6.html
func MacroChinaHKBuildingAmount() (dataframe.DataFrame, error) {
	return fetchChinaHKEconomicData("EMG00158066")
}

// MacroChinaHKTradeDiffRatio 东方财富-经济数据一览-中国香港-香港商品贸易差额年率
//
// 返回:
//   - dataframe.DataFrame: 香港商品贸易差额年率数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_8_7.html
func MacroChinaHKTradeDiffRatio() (dataframe.DataFrame, error) {
	return fetchChinaHKEconomicData("EMG00157898")
}

// MacroChinaHKPPI 东方财富-经济数据一览-中国香港-香港制造业PPI年率
//
// 返回:
//   - dataframe.DataFrame: 香港制造业PPI年率数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_8_8.html
func MacroChinaHKPPI() (dataframe.DataFrame, error) {
	return fetchChinaHKEconomicData("EMG00157818")
}

// MacroChinaHKCore 东方财富-经济数据一览-中国香港-核心指标数据
//
// 参数:
//   - symbol: 指标代码，默认为 "EMG00341602"
//
// 返回:
//   - dataframe.DataFrame: 指定指标的经济数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_8_0.html
func MacroChinaHKCore(symbol string) (dataframe.DataFrame, error) {
	if symbol == "" {
		symbol = "EMG00341602"
	}
	return fetchChinaHKEconomicData(symbol)
}
