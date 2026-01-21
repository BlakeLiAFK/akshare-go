package economic

import (
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
	"github.com/tidwall/gjson"
)

// fetchJapanEconomicData 获取日本经济数据的通用函数
func fetchJapanEconomicData(indicatorID string) (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"

	params := map[string]string{
		"reportName":  "RPT_ECONOMICVALUE_JPAN",
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

// MacroJapanBankRate 东方财富-经济数据-日本-央行公布利率决议
//
// 返回:
//   - dataframe.DataFrame: 央行公布利率决议数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_3_0.html
func MacroJapanBankRate() (dataframe.DataFrame, error) {
	return fetchJapanEconomicData("EMG00342252")
}

// MacroJapanCPIYearly 东方财富-经济数据-日本-全国消费者物价指数年率
//
// 返回:
//   - dataframe.DataFrame: 全国消费者物价指数年率数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_3_1.html
func MacroJapanCPIYearly() (dataframe.DataFrame, error) {
	return fetchJapanEconomicData("EMG00005004")
}

// MacroJapanCoreCPIYearly 东方财富-经济数据-日本-全国核心消费者物价指数年率
//
// 返回:
//   - dataframe.DataFrame: 全国核心消费者物价指数年率数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_2_2.html
func MacroJapanCoreCPIYearly() (dataframe.DataFrame, error) {
	return fetchJapanEconomicData("EMG00158099")
}

// MacroJapanUnemploymentRate 东方财富-经济数据-日本-失业率
//
// 返回:
//   - dataframe.DataFrame: 失业率数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_2_3.html
func MacroJapanUnemploymentRate() (dataframe.DataFrame, error) {
	return fetchJapanEconomicData("EMG00005047")
}

// MacroJapanHeadIndicator 东方财富-经济数据-日本-领先指标终值
//
// 返回:
//   - dataframe.DataFrame: 领先指标终值数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_3_4.html
func MacroJapanHeadIndicator() (dataframe.DataFrame, error) {
	return fetchJapanEconomicData("EMG00005117")
}

// MacroJapanCore 东方财富-经济数据一览-日本-核心指标数据
//
// 参数:
//   - symbol: 指标代码，默认为 "EMG00341602"
//
// 返回:
//   - dataframe.DataFrame: 指定指标的经济数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/foreign_1_0.html
func MacroJapanCore(symbol string) (dataframe.DataFrame, error) {
	if symbol == "" {
		symbol = "EMG00341602"
	}
	return fetchJapanEconomicData(symbol)
}
