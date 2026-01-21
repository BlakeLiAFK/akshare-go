package economic

import (
	"fmt"
	"sort"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
	"github.com/tidwall/gjson"
)

// MacroChinaAuReport 上海黄金交易所报告
//
// 返回:
//   - dataframe.DataFrame: 上海黄金交易所报告
//   - error: 错误信息
//
// 数据源: https://datacenter.jin10.com/reportType/dc_sge_report
func MacroChinaAuReport() (dataframe.DataFrame, error) {
	url := "https://cdn.jin10.com/data_center/reports/sge.json"

	resp, err := utils.Get(url, nil)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	values := gjson.Get(resp.String(), "values")
	if !values.Exists() {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	columns := []string{
		"日期", "商品", "开盘价", "最高价", "最低价", "收盘价",
		"涨跌", "涨跌幅", "加权平均价", "成交量", "成交金额",
		"持仓量", "交收方向", "交收量",
	}

	var records [][]string
	records = append(records, columns)

	values.ForEach(func(date, items gjson.Result) bool {
		for _, item := range items.Array() {
			arr := item.Array()
			if len(arr) >= 13 {
				record := []string{
					date.String(),
					arr[0].String(),
					arr[1].String(),
					arr[2].String(),
					arr[3].String(),
					arr[4].String(),
					arr[5].String(),
					arr[6].String(),
					arr[7].String(),
					arr[8].String(),
					arr[9].String(),
					arr[10].String(),
					arr[11].String(),
					arr[12].String(),
				}
				records = append(records, record)
			}
		}
		return true
	})

	if len(records) <= 1 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	// 按日期排序
	sort.Slice(records[1:], func(i, j int) bool {
		return records[i+1][0] < records[j+1][0]
	})

	df := dataframe.LoadRecords(records)
	return df, nil
}

// MacroChinaLPR LPR品种详细数据
//
// 返回:
//   - dataframe.DataFrame: LPR品种详细数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/globalRateLPR.html
func MacroChinaLPR() (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"

	columns := []string{"TRADE_DATE", "LPR1Y", "LPR5Y", "RATE_1", "RATE_2"}
	var allRecords [][]string
	allRecords = append(allRecords, columns)

	for page := 1; page <= 100; page++ {
		params := map[string]string{
			"reportName":  "RPTA_WEB_RATE",
			"columns":     "ALL",
			"sortColumns": "TRADE_DATE",
			"sortTypes":   "-1",
			"token":       "894050c76af8597a853f5b408b759f5d",
			"pageNumber":  fmt.Sprintf("%d", page),
			"pageSize":    "500",
		}

		resp, err := utils.Get(url, params)
		if err != nil {
			return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
		}

		data := gjson.Get(resp.String(), "result.data").Array()
		if len(data) == 0 {
			break
		}

		for _, item := range data {
			record := []string{
				item.Get("TRADE_DATE").String()[:10],
				item.Get("LPR1Y").String(),
				item.Get("LPR5Y").String(),
				item.Get("RATE_1").String(),
				item.Get("RATE_2").String(),
			}
			allRecords = append(allRecords, record)
		}

		pages := gjson.Get(resp.String(), "result.pages").Int()
		if int64(page) >= pages {
			break
		}
	}

	if len(allRecords) <= 1 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	// 按日期排序
	sort.Slice(allRecords[1:], func(i, j int) bool {
		return allRecords[i+1][0] < allRecords[j+1][0]
	})

	df := dataframe.LoadRecords(allRecords)
	return df, nil
}

// MacroChinaNewHousePrice 中国-新房价指数
//
// 参数:
//   - cityFirst: 第一个城市
//   - citySecond: 第二个城市
//
// 返回:
//   - dataframe.DataFrame: 新房价指数
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/newhouse.html
func MacroChinaNewHousePrice(cityFirst, citySecond string) (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"reportName":  "RPT_ECONOMY_HOUSE_PRICE",
		"columns":     "REPORT_DATE,CITY,FIRST_COMHOUSE_SAME,FIRST_COMHOUSE_SEQUENTIAL,FIRST_COMHOUSE_BASE,SECOND_HOUSE_SAME,SECOND_HOUSE_SEQUENTIAL,SECOND_HOUSE_BASE,REPORT_DAY",
		"filter":      fmt.Sprintf(`(CITY in ("%s","%s"))`, cityFirst, citySecond),
		"pageNumber":  "1",
		"pageSize":    "500",
		"sortColumns": "REPORT_DATE,CITY",
		"sortTypes":   "-1,-1",
		"source":      "WEB",
		"client":      "WEB",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	data := gjson.Get(resp.String(), "result.data").Array()
	if len(data) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	columns := []string{
		"日期", "城市",
		"新建商品住宅价格指数-同比", "新建商品住宅价格指数-环比", "新建商品住宅价格指数-定基",
		"二手住宅价格指数-同比", "二手住宅价格指数-环比", "二手住宅价格指数-定基",
	}

	var records [][]string
	records = append(records, columns)

	for _, item := range data {
		record := []string{
			item.Get("REPORT_DATE").String()[:10],
			item.Get("CITY").String(),
			item.Get("FIRST_COMHOUSE_SAME").String(),
			item.Get("FIRST_COMHOUSE_SEQUENTIAL").String(),
			item.Get("FIRST_COMHOUSE_BASE").String(),
			item.Get("SECOND_HOUSE_SAME").String(),
			item.Get("SECOND_HOUSE_SEQUENTIAL").String(),
			item.Get("SECOND_HOUSE_BASE").String(),
		}
		records = append(records, record)
	}

	// 按日期排序
	sort.Slice(records[1:], func(i, j int) bool {
		return records[i+1][0] < records[j+1][0]
	})

	df := dataframe.LoadRecords(records)
	return df, nil
}

// MacroChinaEnterpriseBoomIndex 中国-企业景气及企业家信心指数
//
// 返回:
//   - dataframe.DataFrame: 企业景气及企业家信心指数
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/qyjqzs.html
func MacroChinaEnterpriseBoomIndex() (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"columns":     "REPORT_DATE,TIME,BOOM_INDEX,FAITH_INDEX,BOOM_INDEX_SAME,BOOM_INDEX_SEQUENTIAL,FAITH_INDEX_SAME,FAITH_INDEX_SEQUENTIAL",
		"pageNumber":  "1",
		"pageSize":    "500",
		"sortColumns": "REPORT_DATE",
		"sortTypes":   "-1",
		"source":      "WEB",
		"client":      "WEB",
		"reportName":  "RPT_ECONOMY_BOOM_INDEX",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	data := gjson.Get(resp.String(), "result.data").Array()
	if len(data) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	columns := []string{
		"季度", "企业景气指数-指数", "企业景气指数-同比", "企业景气指数-环比",
		"企业家信心指数-指数", "企业家信心指数-同比", "企业家信心指数-环比",
	}

	var records [][]string
	records = append(records, columns)

	for _, item := range data {
		record := []string{
			item.Get("TIME").String(),
			item.Get("BOOM_INDEX").String(),
			item.Get("BOOM_INDEX_SAME").String(),
			item.Get("BOOM_INDEX_SEQUENTIAL").String(),
			item.Get("FAITH_INDEX").String(),
			item.Get("FAITH_INDEX_SAME").String(),
			item.Get("FAITH_INDEX_SEQUENTIAL").String(),
		}
		records = append(records, record)
	}

	df := dataframe.LoadRecords(records)
	return df, nil
}

// MacroChinaNationalTaxReceipts 中国-全国税收收入
//
// 返回:
//   - dataframe.DataFrame: 全国税收收入
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/qgsssr.html
func MacroChinaNationalTaxReceipts() (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"columns":     "REPORT_DATE,TIME,TAX_INCOME,TAX_INCOME_SAME,TAX_INCOME_SEQUENTIAL",
		"pageNumber":  "1",
		"pageSize":    "500",
		"sortColumns": "REPORT_DATE",
		"sortTypes":   "-1",
		"source":      "WEB",
		"client":      "WEB",
		"reportName":  "RPT_ECONOMY_TAX",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	data := gjson.Get(resp.String(), "result.data").Array()
	if len(data) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	columns := []string{"季度", "税收收入合计", "较上年同期", "季度环比"}

	var records [][]string
	records = append(records, columns)

	for _, item := range data {
		record := []string{
			item.Get("TIME").String(),
			item.Get("TAX_INCOME").String(),
			item.Get("TAX_INCOME_SAME").String(),
			item.Get("TAX_INCOME_SEQUENTIAL").String(),
		}
		records = append(records, record)
	}

	df := dataframe.LoadRecords(records)
	return df, nil
}

// MacroChinaBankFinancing 银行理财产品发行数量
//
// 返回:
//   - dataframe.DataFrame: 银行理财产品发行数量
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/hyzs_list_EMI01516267.html
func MacroChinaBankFinancing() (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"sortColumns": "REPORT_DATE",
		"sortTypes":   "-1",
		"pageSize":    "1000",
		"pageNumber":  "1",
		"reportName":  "RPT_INDUSTRY_INDEX",
		"columns":     "REPORT_DATE,INDICATOR_VALUE,CHANGE_RATE,CHANGERATE_3M,CHANGERATE_6M,CHANGERATE_1Y,CHANGERATE_2Y,CHANGERATE_3Y",
		"filter":      `(INDICATOR_ID="EMI01516267")`,
		"source":      "WEB",
		"client":      "WEB",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	data := gjson.Get(resp.String(), "result.data").Array()
	if len(data) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	columns := []string{"日期", "最新值", "涨跌幅", "近3月涨跌幅", "近6月涨跌幅", "近1年涨跌幅", "近2年涨跌幅", "近3年涨跌幅"}

	var records [][]string
	records = append(records, columns)

	for _, item := range data {
		record := []string{
			item.Get("REPORT_DATE").String()[:10],
			item.Get("INDICATOR_VALUE").String(),
			item.Get("CHANGE_RATE").String(),
			item.Get("CHANGERATE_3M").String(),
			item.Get("CHANGERATE_6M").String(),
			item.Get("CHANGERATE_1Y").String(),
			item.Get("CHANGERATE_2Y").String(),
			item.Get("CHANGERATE_3Y").String(),
		}
		records = append(records, record)
	}

	// 按日期排序
	sort.Slice(records[1:], func(i, j int) bool {
		return records[i+1][0] < records[j+1][0]
	})

	df := dataframe.LoadRecords(records)
	return df, nil
}

// MacroChinaInsuranceIncome 原保险保费收入
//
// 返回:
//   - dataframe.DataFrame: 原保险保费收入
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/hyzs_list_EMM00088870.html
func MacroChinaInsuranceIncome() (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"sortColumns": "REPORT_DATE",
		"sortTypes":   "-1",
		"pageSize":    "1000",
		"pageNumber":  "1",
		"reportName":  "RPT_INDUSTRY_INDEX",
		"columns":     "REPORT_DATE,INDICATOR_VALUE,CHANGE_RATE,CHANGERATE_3M,CHANGERATE_6M,CHANGERATE_1Y,CHANGERATE_2Y,CHANGERATE_3Y",
		"filter":      `(INDICATOR_ID="EMM00088870")`,
		"source":      "WEB",
		"client":      "WEB",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	data := gjson.Get(resp.String(), "result.data").Array()
	if len(data) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	columns := []string{"日期", "最新值", "涨跌幅", "近3月涨跌幅", "近6月涨跌幅", "近1年涨跌幅", "近2年涨跌幅", "近3年涨跌幅"}

	var records [][]string
	records = append(records, columns)

	for _, item := range data {
		record := []string{
			item.Get("REPORT_DATE").String()[:10],
			item.Get("INDICATOR_VALUE").String(),
			item.Get("CHANGE_RATE").String(),
			item.Get("CHANGERATE_3M").String(),
			item.Get("CHANGERATE_6M").String(),
			item.Get("CHANGERATE_1Y").String(),
			item.Get("CHANGERATE_2Y").String(),
			item.Get("CHANGERATE_3Y").String(),
		}
		records = append(records, record)
	}

	// 按日期排序
	sort.Slice(records[1:], func(i, j int) bool {
		return records[i+1][0] < records[j+1][0]
	})

	df := dataframe.LoadRecords(records)
	return df, nil
}

// MacroChinaMobileNumber 移动电话用户数
//
// 返回:
//   - dataframe.DataFrame: 移动电话用户数
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/hyzs_list_EMI00009870.html
func MacroChinaMobileNumber() (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"sortColumns": "REPORT_DATE",
		"sortTypes":   "-1",
		"pageSize":    "1000",
		"pageNumber":  "1",
		"reportName":  "RPT_INDUSTRY_INDEX",
		"columns":     "REPORT_DATE,INDICATOR_VALUE,CHANGE_RATE,CHANGERATE_3M,CHANGERATE_6M,CHANGERATE_1Y,CHANGERATE_2Y,CHANGERATE_3Y",
		"filter":      `(INDICATOR_ID="EMI00009870")`,
		"source":      "WEB",
		"client":      "WEB",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	data := gjson.Get(resp.String(), "result.data").Array()
	if len(data) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	columns := []string{"日期", "最新值", "涨跌幅", "近3月涨跌幅", "近6月涨跌幅", "近1年涨跌幅", "近2年涨跌幅", "近3年涨跌幅"}

	var records [][]string
	records = append(records, columns)

	for _, item := range data {
		record := []string{
			item.Get("REPORT_DATE").String()[:10],
			item.Get("INDICATOR_VALUE").String(),
			item.Get("CHANGE_RATE").String(),
			item.Get("CHANGERATE_3M").String(),
			item.Get("CHANGERATE_6M").String(),
			item.Get("CHANGERATE_1Y").String(),
			item.Get("CHANGERATE_2Y").String(),
			item.Get("CHANGERATE_3Y").String(),
		}
		records = append(records, record)
	}

	// 按日期排序
	sort.Slice(records[1:], func(i, j int) bool {
		return records[i+1][0] < records[j+1][0]
	})

	df := dataframe.LoadRecords(records)
	return df, nil
}

// MacroChinaVegetableBasket 菜篮子产品批发价格指数
//
// 返回:
//   - dataframe.DataFrame: 菜篮子产品批发价格指数
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/hyzs_list_EMI00009855.html
func MacroChinaVegetableBasket() (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"sortColumns": "REPORT_DATE",
		"sortTypes":   "-1",
		"pageSize":    "1000",
		"pageNumber":  "1",
		"reportName":  "RPT_INDUSTRY_INDEX",
		"columns":     "REPORT_DATE,INDICATOR_VALUE,CHANGE_RATE,CHANGERATE_3M,CHANGERATE_6M,CHANGERATE_1Y,CHANGERATE_2Y,CHANGERATE_3Y",
		"filter":      `(INDICATOR_ID="EMI00009855")`,
		"source":      "WEB",
		"client":      "WEB",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	data := gjson.Get(resp.String(), "result.data").Array()
	if len(data) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	columns := []string{"日期", "最新值", "涨跌幅", "近3月涨跌幅", "近6月涨跌幅", "近1年涨跌幅", "近2年涨跌幅", "近3年涨跌幅"}

	var records [][]string
	records = append(records, columns)

	for _, item := range data {
		record := []string{
			item.Get("REPORT_DATE").String()[:10],
			item.Get("INDICATOR_VALUE").String(),
			item.Get("CHANGE_RATE").String(),
			item.Get("CHANGERATE_3M").String(),
			item.Get("CHANGERATE_6M").String(),
			item.Get("CHANGERATE_1Y").String(),
			item.Get("CHANGERATE_2Y").String(),
			item.Get("CHANGERATE_3Y").String(),
		}
		records = append(records, record)
	}

	// 按日期排序
	sort.Slice(records[1:], func(i, j int) bool {
		return records[i+1][0] < records[j+1][0]
	})

	df := dataframe.LoadRecords(records)
	return df, nil
}

// MacroChinaAgriculturalProduct 农产品批发价格200指数
//
// 返回:
//   - dataframe.DataFrame: 农产品批发价格200指数
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/hyzs_list_EMI00009856.html
func MacroChinaAgriculturalProduct() (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"sortColumns": "REPORT_DATE",
		"sortTypes":   "-1",
		"pageSize":    "1000",
		"pageNumber":  "1",
		"reportName":  "RPT_INDUSTRY_INDEX",
		"columns":     "REPORT_DATE,INDICATOR_VALUE,CHANGE_RATE,CHANGERATE_3M,CHANGERATE_6M,CHANGERATE_1Y,CHANGERATE_2Y,CHANGERATE_3Y",
		"filter":      `(INDICATOR_ID="EMI00009856")`,
		"source":      "WEB",
		"client":      "WEB",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	data := gjson.Get(resp.String(), "result.data").Array()
	if len(data) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	columns := []string{"日期", "最新值", "涨跌幅", "近3月涨跌幅", "近6月涨跌幅", "近1年涨跌幅", "近2年涨跌幅", "近3年涨跌幅"}

	var records [][]string
	records = append(records, columns)

	for _, item := range data {
		record := []string{
			item.Get("REPORT_DATE").String()[:10],
			item.Get("INDICATOR_VALUE").String(),
			item.Get("CHANGE_RATE").String(),
			item.Get("CHANGERATE_3M").String(),
			item.Get("CHANGERATE_6M").String(),
			item.Get("CHANGERATE_1Y").String(),
			item.Get("CHANGERATE_2Y").String(),
			item.Get("CHANGERATE_3Y").String(),
		}
		records = append(records, record)
	}

	// 按日期排序
	sort.Slice(records[1:], func(i, j int) bool {
		return records[i+1][0] < records[j+1][0]
	})

	df := dataframe.LoadRecords(records)
	return df, nil
}

// MacroChinaAgriculturalIndex 农副指数
//
// 返回:
//   - dataframe.DataFrame: 农副指数
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/hyzs_list_EMI00662535.html
func MacroChinaAgriculturalIndex() (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"sortColumns": "REPORT_DATE",
		"sortTypes":   "-1",
		"pageSize":    "1000",
		"pageNumber":  "1",
		"reportName":  "RPT_INDUSTRY_INDEX",
		"columns":     "REPORT_DATE,INDICATOR_VALUE,CHANGE_RATE,CHANGERATE_3M,CHANGERATE_6M,CHANGERATE_1Y,CHANGERATE_2Y,CHANGERATE_3Y",
		"filter":      `(INDICATOR_ID="EMI00662535")`,
		"source":      "WEB",
		"client":      "WEB",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	data := gjson.Get(resp.String(), "result.data").Array()
	if len(data) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	columns := []string{"日期", "最新值", "涨跌幅", "近3月涨跌幅", "近6月涨跌幅", "近1年涨跌幅", "近2年涨跌幅", "近3年涨跌幅"}

	var records [][]string
	records = append(records, columns)

	for _, item := range data {
		record := []string{
			item.Get("REPORT_DATE").String()[:10],
			item.Get("INDICATOR_VALUE").String(),
			item.Get("CHANGE_RATE").String(),
			item.Get("CHANGERATE_3M").String(),
			item.Get("CHANGERATE_6M").String(),
			item.Get("CHANGERATE_1Y").String(),
			item.Get("CHANGERATE_2Y").String(),
			item.Get("CHANGERATE_3Y").String(),
		}
		records = append(records, record)
	}

	// 按日期排序
	sort.Slice(records[1:], func(i, j int) bool {
		return records[i+1][0] < records[j+1][0]
	})

	df := dataframe.LoadRecords(records)
	return df, nil
}

// MacroChinaEnergyIndex 能源指数
//
// 返回:
//   - dataframe.DataFrame: 能源指数
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/hyzs_list_EMI00662534.html
func MacroChinaEnergyIndex() (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"sortColumns": "REPORT_DATE",
		"sortTypes":   "-1",
		"pageSize":    "1000",
		"pageNumber":  "1",
		"reportName":  "RPT_INDUSTRY_INDEX",
		"columns":     "REPORT_DATE,INDICATOR_VALUE,CHANGE_RATE,CHANGERATE_3M,CHANGERATE_6M,CHANGERATE_1Y,CHANGERATE_2Y,CHANGERATE_3Y",
		"filter":      `(INDICATOR_ID="EMI00662534")`,
		"source":      "WEB",
		"client":      "WEB",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	data := gjson.Get(resp.String(), "result.data").Array()
	if len(data) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	columns := []string{"日期", "最新值", "涨跌幅", "近3月涨跌幅", "近6月涨跌幅", "近1年涨跌幅", "近2年涨跌幅", "近3年涨跌幅"}

	var records [][]string
	records = append(records, columns)

	for _, item := range data {
		record := []string{
			item.Get("REPORT_DATE").String()[:10],
			item.Get("INDICATOR_VALUE").String(),
			item.Get("CHANGE_RATE").String(),
			item.Get("CHANGERATE_3M").String(),
			item.Get("CHANGERATE_6M").String(),
			item.Get("CHANGERATE_1Y").String(),
			item.Get("CHANGERATE_2Y").String(),
			item.Get("CHANGERATE_3Y").String(),
		}
		records = append(records, record)
	}

	// 按日期排序
	sort.Slice(records[1:], func(i, j int) bool {
		return records[i+1][0] < records[j+1][0]
	})

	df := dataframe.LoadRecords(records)
	return df, nil
}

// MacroChinaCommodityPriceIndex 大宗商品价格指数
//
// 返回:
//   - dataframe.DataFrame: 大宗商品价格指数
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/hyzs_list_EMI00662542.html
func MacroChinaCommodityPriceIndex() (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"sortColumns": "REPORT_DATE",
		"sortTypes":   "-1",
		"pageSize":    "1000",
		"pageNumber":  "1",
		"reportName":  "RPT_INDUSTRY_INDEX",
		"columns":     "REPORT_DATE,INDICATOR_VALUE,CHANGE_RATE,CHANGERATE_3M,CHANGERATE_6M,CHANGERATE_1Y,CHANGERATE_2Y,CHANGERATE_3Y",
		"filter":      `(INDICATOR_ID="EMI00662542")`,
		"source":      "WEB",
		"client":      "WEB",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	data := gjson.Get(resp.String(), "result.data").Array()
	if len(data) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	columns := []string{"日期", "最新值", "涨跌幅", "近3月涨跌幅", "近6月涨跌幅", "近1年涨跌幅", "近2年涨跌幅", "近3年涨跌幅"}

	var records [][]string
	records = append(records, columns)

	for _, item := range data {
		record := []string{
			item.Get("REPORT_DATE").String()[:10],
			item.Get("INDICATOR_VALUE").String(),
			item.Get("CHANGE_RATE").String(),
			item.Get("CHANGERATE_3M").String(),
			item.Get("CHANGERATE_6M").String(),
			item.Get("CHANGERATE_1Y").String(),
			item.Get("CHANGERATE_2Y").String(),
			item.Get("CHANGERATE_3Y").String(),
		}
		records = append(records, record)
	}

	// 按日期排序
	sort.Slice(records[1:], func(i, j int) bool {
		return records[i+1][0] < records[j+1][0]
	})

	df := dataframe.LoadRecords(records)
	return df, nil
}
