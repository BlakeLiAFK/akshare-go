package economic

import (
	"fmt"
	"sort"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
	"github.com/tidwall/gjson"
)

// fetchEastmoneyIndustryIndex 东方财富行业指数通用获取函数
// indicatorID: 指标ID
func fetchEastmoneyIndustryIndex(indicatorID string) (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"

	columns := []string{"日期", "最新值", "涨跌幅", "近3月涨跌幅", "近6月涨跌幅", "近1年涨跌幅", "近2年涨跌幅", "近3年涨跌幅"}
	var allRecords [][]string
	allRecords = append(allRecords, columns)

	for page := 1; page <= 100; page++ {
		params := map[string]string{
			"sortColumns": "REPORT_DATE",
			"sortTypes":   "-1",
			"pageSize":    "500",
			"pageNumber":  fmt.Sprintf("%d", page),
			"reportName":  "RPT_INDUSTRY_INDEX",
			"columns":     "REPORT_DATE,INDICATOR_VALUE,CHANGE_RATE,CHANGERATE_3M,CHANGERATE_6M,CHANGERATE_1Y,CHANGERATE_2Y,CHANGERATE_3Y",
			"filter":      fmt.Sprintf(`(INDICATOR_ID="%s")`, indicatorID),
			"source":      "WEB",
			"client":      "WEB",
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
			date := item.Get("REPORT_DATE").String()
			if len(date) >= 10 {
				date = date[:10]
			}
			record := []string{
				date,
				item.Get("INDICATOR_VALUE").String(),
				item.Get("CHANGE_RATE").String(),
				item.Get("CHANGERATE_3M").String(),
				item.Get("CHANGERATE_6M").String(),
				item.Get("CHANGERATE_1Y").String(),
				item.Get("CHANGERATE_2Y").String(),
				item.Get("CHANGERATE_3Y").String(),
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

// MacroChinaYwElectronicIndex 义乌小商品指数-电子元器件
//
// 返回:
//   - dataframe.DataFrame: 义乌小商品指数-电子元器件
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/hyzs_list_EMI00055551.html
func MacroChinaYwElectronicIndex() (dataframe.DataFrame, error) {
	return fetchEastmoneyIndustryIndex("EMI00055551")
}

// MacroChinaConstructionIndex 建材指数
//
// 返回:
//   - dataframe.DataFrame: 建材指数
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/hyzs_list_EMI00662541.html
func MacroChinaConstructionIndex() (dataframe.DataFrame, error) {
	return fetchEastmoneyIndustryIndex("EMI00662541")
}

// MacroChinaConstructionPriceIndex 建材价格指数
//
// 返回:
//   - dataframe.DataFrame: 建材价格指数
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/hyzs_list_EMI00237146.html
func MacroChinaConstructionPriceIndex() (dataframe.DataFrame, error) {
	return fetchEastmoneyIndustryIndex("EMI00237146")
}

// MacroChinaLPIIndex 物流景气指数
//
// 返回:
//   - dataframe.DataFrame: 物流景气指数
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/hyzs_list_EMI00352262.html
func MacroChinaLPIIndex() (dataframe.DataFrame, error) {
	return fetchEastmoneyIndustryIndex("EMI00352262")
}

// MacroChinaBDTIIndex 原油运输指数
//
// 返回:
//   - dataframe.DataFrame: 原油运输指数
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/hyzs_list_EMI00107668.html
func MacroChinaBDTIIndex() (dataframe.DataFrame, error) {
	return fetchEastmoneyIndustryIndex("EMI00107668")
}

// MacroChinaBSIIndex 超灵便型船运价指数
//
// 返回:
//   - dataframe.DataFrame: 超灵便型船运价指数
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/hyzs_list_EMI00107667.html
func MacroChinaBSIIndex() (dataframe.DataFrame, error) {
	return fetchEastmoneyIndustryIndex("EMI00107667")
}

// MacroShippingBCI 海岬型运费指数(BCI)
//
// 返回:
//   - dataframe.DataFrame: 海岬型运费指数
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/hyzs_list_EMI00107666.html
func MacroShippingBCI() (dataframe.DataFrame, error) {
	return fetchEastmoneyIndustryIndex("EMI00107666")
}

// MacroShippingBDI 波罗的海干散货指数(BDI)
//
// 返回:
//   - dataframe.DataFrame: 波罗的海干散货指数
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/hyzs_list_EMI00107664.html
func MacroShippingBDI() (dataframe.DataFrame, error) {
	return fetchEastmoneyIndustryIndex("EMI00107664")
}

// MacroShippingBPI 巴拿马型运费指数(BPI)
//
// 返回:
//   - dataframe.DataFrame: 巴拿马型运费指数
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/hyzs_list_EMI00107665.html
func MacroShippingBPI() (dataframe.DataFrame, error) {
	return fetchEastmoneyIndustryIndex("EMI00107665")
}

// MacroShippingBCTI 成品油运输指数(BCTI)
//
// 返回:
//   - dataframe.DataFrame: 成品油运输指数
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/hyzs_list_EMI00107669.html
func MacroShippingBCTI() (dataframe.DataFrame, error) {
	return fetchEastmoneyIndustryIndex("EMI00107669")
}

// MacroChinaNewFinancialCredit 中国-新增信贷数据
//
// 返回:
//   - dataframe.DataFrame: 新增信贷数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/xzxd.html
func MacroChinaNewFinancialCredit() (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"columns":     "REPORT_DATE,TIME,RMB_LOAN,RMB_LOAN_SAME,RMB_LOAN_SEQUENTIAL,RMB_LOAN_ACCUMULATE,LOAN_ACCUMULATE_SAME",
		"pageNumber":  "1",
		"pageSize":    "2000",
		"sortColumns": "REPORT_DATE",
		"sortTypes":   "-1",
		"source":      "WEB",
		"client":      "WEB",
		"reportName":  "RPT_ECONOMY_RMB_LOAN",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	data := gjson.Get(resp.String(), "result.data").Array()
	if len(data) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	columns := []string{"月份", "当月", "当月-同比增长", "当月-环比增长", "累计", "累计-同比增长"}
	var records [][]string
	records = append(records, columns)

	for _, item := range data {
		record := []string{
			item.Get("TIME").String(),
			item.Get("RMB_LOAN").String(),
			item.Get("RMB_LOAN_SAME").String(),
			item.Get("RMB_LOAN_SEQUENTIAL").String(),
			item.Get("RMB_LOAN_ACCUMULATE").String(),
			item.Get("LOAN_ACCUMULATE_SAME").String(),
		}
		records = append(records, record)
	}

	df := dataframe.LoadRecords(records)
	return df, nil
}

// MacroChinaFxGold 东方财富-外汇和黄金储备
//
// 返回:
//   - dataframe.DataFrame: 外汇和黄金储备
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/hjwh.html
func MacroChinaFxGold() (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"columns":     "REPORT_DATE,TIME,GOLD_RESERVES,GOLD_RESERVES_SAME,GOLD_RESERVES_SEQUENTIAL,FOREX,FOREX_SAME,FOREX_SEQUENTIAL",
		"pageNumber":  "1",
		"pageSize":    "1000",
		"sortColumns": "REPORT_DATE",
		"sortTypes":   "-1",
		"source":      "WEB",
		"client":      "WEB",
		"reportName":  "RPT_ECONOMY_GOLD_CURRENCY",
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
		"月份", "黄金储备-数值", "黄金储备-同比", "黄金储备-环比",
		"国家外汇储备-数值", "国家外汇储备-同比", "国家外汇储备-环比",
	}
	var records [][]string
	records = append(records, columns)

	for _, item := range data {
		record := []string{
			item.Get("TIME").String(),
			item.Get("GOLD_RESERVES").String(),
			item.Get("GOLD_RESERVES_SAME").String(),
			item.Get("GOLD_RESERVES_SEQUENTIAL").String(),
			item.Get("FOREX").String(),
			item.Get("FOREX_SAME").String(),
			item.Get("FOREX_SEQUENTIAL").String(),
		}
		records = append(records, record)
	}

	// 按月份排序
	sort.Slice(records[1:], func(i, j int) bool {
		return records[i+1][0] < records[j+1][0]
	})

	df := dataframe.LoadRecords(records)
	return df, nil
}

// MacroChinaStockMarketCap 东方财富-全国股票交易统计表
//
// 返回:
//   - dataframe.DataFrame: 全国股票交易统计表
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/gpjytj.html
func MacroChinaStockMarketCap() (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"reportName":  "RPT_ECONOMY_STOCK_STATISTICS",
		"columns":     "REPORT_DATE,TIME,TOTAL_SHARES_SH,TOTAL_MARKE_SH,DEAL_AMOUNT_SH,VOLUME_SH,HIGH_INDEX_SH,LOW_INDEX_SH,TOTAL_SZARES_SZ,TOTAL_MARKE_SZ,DEAL_AMOUNT_SZ,VOLUME_SZ,HIGH_INDEX_SZ,LOW_INDEX_SZ",
		"sortColumns": "REPORT_DATE",
		"sortTypes":   "-1",
		"pageNumber":  "1",
		"pageSize":    "1000",
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
		"数据日期",
		"发行总股本-上海", "发行总股本-深圳",
		"市价总值-上海", "市价总值-深圳",
		"成交金额-上海", "成交金额-深圳",
		"成交量-上海", "成交量-深圳",
		"A股最高综合股价指数-上海", "A股最高综合股价指数-深圳",
		"A股最低综合股价指数-上海", "A股最低综合股价指数-深圳",
	}
	var records [][]string
	records = append(records, columns)

	for _, item := range data {
		record := []string{
			item.Get("TIME").String(),
			item.Get("TOTAL_SHARES_SH").String(),
			item.Get("TOTAL_SZARES_SZ").String(),
			item.Get("TOTAL_MARKE_SH").String(),
			item.Get("TOTAL_MARKE_SZ").String(),
			item.Get("DEAL_AMOUNT_SH").String(),
			item.Get("DEAL_AMOUNT_SZ").String(),
			item.Get("VOLUME_SH").String(),
			item.Get("VOLUME_SZ").String(),
			item.Get("HIGH_INDEX_SH").String(),
			item.Get("HIGH_INDEX_SZ").String(),
			item.Get("LOW_INDEX_SH").String(),
			item.Get("LOW_INDEX_SZ").String(),
		}
		records = append(records, record)
	}

	df := dataframe.LoadRecords(records)
	return df, nil
}

// MacroChinaMoneySupply 东方财富-货币供应量
//
// 返回:
//   - dataframe.DataFrame: 货币供应量
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/hbgyl.html
func MacroChinaMoneySupply() (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"columns":     "REPORT_DATE,TIME,BASIC_CURRENCY,BASIC_CURRENCY_SAME,BASIC_CURRENCY_SEQUENTIAL,CURRENCY,CURRENCY_SAME,CURRENCY_SEQUENTIAL,FREE_CASH,FREE_CASH_SAME,FREE_CASH_SEQUENTIAL",
		"pageNumber":  "1",
		"pageSize":    "2000",
		"sortColumns": "REPORT_DATE",
		"sortTypes":   "-1",
		"source":      "WEB",
		"client":      "WEB",
		"reportName":  "RPT_ECONOMY_CURRENCY_SUPPLY",
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
		"月份",
		"货币和准货币(M2)-数量(亿元)", "货币和准货币(M2)-同比增长", "货币和准货币(M2)-环比增长",
		"货币(M1)-数量(亿元)", "货币(M1)-同比增长", "货币(M1)-环比增长",
		"流通中的现金(M0)-数量(亿元)", "流通中的现金(M0)-同比增长", "流通中的现金(M0)-环比增长",
	}
	var records [][]string
	records = append(records, columns)

	for _, item := range data {
		record := []string{
			item.Get("TIME").String(),
			item.Get("BASIC_CURRENCY").String(),
			item.Get("BASIC_CURRENCY_SAME").String(),
			item.Get("BASIC_CURRENCY_SEQUENTIAL").String(),
			item.Get("CURRENCY").String(),
			item.Get("CURRENCY_SAME").String(),
			item.Get("CURRENCY_SEQUENTIAL").String(),
			item.Get("FREE_CASH").String(),
			item.Get("FREE_CASH_SAME").String(),
			item.Get("FREE_CASH_SEQUENTIAL").String(),
		}
		records = append(records, record)
	}

	df := dataframe.LoadRecords(records)
	return df, nil
}

// MacroChinaCPI 东方财富-中国居民消费价格指数
//
// 返回:
//   - dataframe.DataFrame: 中国居民消费价格指数
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/cpi.html
func MacroChinaCPI() (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"columns":     "REPORT_DATE,TIME,NATIONAL_SAME,NATIONAL_BASE,NATIONAL_SEQUENTIAL,NATIONAL_ACCUMULATE,CITY_SAME,CITY_BASE,CITY_SEQUENTIAL,CITY_ACCUMULATE,RURAL_SAME,RURAL_BASE,RURAL_SEQUENTIAL,RURAL_ACCUMULATE",
		"pageNumber":  "1",
		"pageSize":    "2000",
		"sortColumns": "REPORT_DATE",
		"sortTypes":   "-1",
		"source":      "WEB",
		"client":      "WEB",
		"reportName":  "RPT_ECONOMY_CPI",
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
		"月份",
		"全国-当月", "全国-同比增长", "全国-环比增长", "全国-累计",
		"城市-当月", "城市-同比增长", "城市-环比增长", "城市-累计",
		"农村-当月", "农村-同比增长", "农村-环比增长", "农村-累计",
	}
	var records [][]string
	records = append(records, columns)

	for _, item := range data {
		record := []string{
			item.Get("TIME").String(),
			item.Get("NATIONAL_BASE").String(),
			item.Get("NATIONAL_SAME").String(),
			item.Get("NATIONAL_SEQUENTIAL").String(),
			item.Get("NATIONAL_ACCUMULATE").String(),
			item.Get("CITY_BASE").String(),
			item.Get("CITY_SAME").String(),
			item.Get("CITY_SEQUENTIAL").String(),
			item.Get("CITY_ACCUMULATE").String(),
			item.Get("RURAL_BASE").String(),
			item.Get("RURAL_SAME").String(),
			item.Get("RURAL_SEQUENTIAL").String(),
			item.Get("RURAL_ACCUMULATE").String(),
		}
		records = append(records, record)
	}

	df := dataframe.LoadRecords(records)
	return df, nil
}

// MacroChinaGDP 东方财富-中国国内生产总值
//
// 返回:
//   - dataframe.DataFrame: 中国国内生产总值
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/gdp.html
func MacroChinaGDP() (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"columns":     "REPORT_DATE,TIME,DOMESTICL_PRODUCT_BASE,FIRST_PRODUCT_BASE,SECOND_PRODUCT_BASE,THIRD_PRODUCT_BASE,SUM_SAME,FIRST_SAME,SECOND_SAME,THIRD_SAME",
		"pageNumber":  "1",
		"pageSize":    "2000",
		"sortColumns": "REPORT_DATE",
		"sortTypes":   "-1",
		"source":      "WEB",
		"client":      "WEB",
		"reportName":  "RPT_ECONOMY_GDP",
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
		"季度",
		"国内生产总值-绝对值", "国内生产总值-同比增长",
		"第一产业-绝对值", "第一产业-同比增长",
		"第二产业-绝对值", "第二产业-同比增长",
		"第三产业-绝对值", "第三产业-同比增长",
	}
	var records [][]string
	records = append(records, columns)

	for _, item := range data {
		record := []string{
			item.Get("TIME").String(),
			item.Get("DOMESTICL_PRODUCT_BASE").String(),
			item.Get("SUM_SAME").String(),
			item.Get("FIRST_PRODUCT_BASE").String(),
			item.Get("FIRST_SAME").String(),
			item.Get("SECOND_PRODUCT_BASE").String(),
			item.Get("SECOND_SAME").String(),
			item.Get("THIRD_PRODUCT_BASE").String(),
			item.Get("THIRD_SAME").String(),
		}
		records = append(records, record)
	}

	df := dataframe.LoadRecords(records)
	return df, nil
}

// MacroChinaPPI 东方财富-中国工业品出厂价格指数
//
// 返回:
//   - dataframe.DataFrame: 中国工业品出厂价格指数
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/ppi.html
func MacroChinaPPI() (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"columns":     "REPORT_DATE,TIME,BASE,BASE_SAME,BASE_ACCUMULATE",
		"pageNumber":  "1",
		"pageSize":    "2000",
		"sortColumns": "REPORT_DATE",
		"sortTypes":   "-1",
		"source":      "WEB",
		"client":      "WEB",
		"reportName":  "RPT_ECONOMY_PPI",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	data := gjson.Get(resp.String(), "result.data").Array()
	if len(data) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	columns := []string{"月份", "当月", "当月同比增长", "累计"}
	var records [][]string
	records = append(records, columns)

	for _, item := range data {
		record := []string{
			item.Get("TIME").String(),
			item.Get("BASE").String(),
			item.Get("BASE_SAME").String(),
			item.Get("BASE_ACCUMULATE").String(),
		}
		records = append(records, record)
	}

	df := dataframe.LoadRecords(records)
	return df, nil
}

// MacroChinaPMI 东方财富-中国采购经理人指数
//
// 返回:
//   - dataframe.DataFrame: 中国采购经理人指数
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/pmi.html
func MacroChinaPMI() (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"columns":     "REPORT_DATE,TIME,MAKE_INDEX,MAKE_SAME,NMAKE_INDEX,NMAKE_SAME",
		"pageNumber":  "1",
		"pageSize":    "2000",
		"sortColumns": "REPORT_DATE",
		"sortTypes":   "-1",
		"source":      "WEB",
		"client":      "WEB",
		"reportName":  "RPT_ECONOMY_PMI",
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
		"月份",
		"制造业-指数", "制造业-同比增长",
		"非制造业-指数", "非制造业-同比增长",
	}
	var records [][]string
	records = append(records, columns)

	for _, item := range data {
		record := []string{
			item.Get("TIME").String(),
			item.Get("MAKE_INDEX").String(),
			item.Get("MAKE_SAME").String(),
			item.Get("NMAKE_INDEX").String(),
			item.Get("NMAKE_SAME").String(),
		}
		records = append(records, record)
	}

	df := dataframe.LoadRecords(records)
	return df, nil
}
