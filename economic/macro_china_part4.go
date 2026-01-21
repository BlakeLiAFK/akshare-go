package economic

import (
	"fmt"
	"sort"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
	"github.com/tidwall/gjson"
)

// MacroChinaGdzctz 东方财富-中国城镇固定资产投资
//
// 返回:
//   - dataframe.DataFrame: 中国城镇固定资产投资
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/gdzctz.html
func MacroChinaGdzctz() (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"columns":     "REPORT_DATE,TIME,BASE,BASE_SAME,BASE_SEQUENTIAL,BASE_ACCUMULATE",
		"pageNumber":  "1",
		"pageSize":    "2000",
		"sortColumns": "REPORT_DATE",
		"sortTypes":   "-1",
		"source":      "WEB",
		"client":      "WEB",
		"reportName":  "RPT_ECONOMY_ASSET_INVEST",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	data := gjson.Get(resp.String(), "result.data").Array()
	if len(data) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	columns := []string{"月份", "当月", "同比增长", "环比增长", "自年初累计"}
	var records [][]string
	records = append(records, columns)

	for _, item := range data {
		record := []string{
			item.Get("TIME").String(),
			item.Get("BASE").String(),
			item.Get("BASE_SAME").String(),
			item.Get("BASE_SEQUENTIAL").String(),
			item.Get("BASE_ACCUMULATE").String(),
		}
		records = append(records, record)
	}

	df := dataframe.LoadRecords(records)
	return df, nil
}

// MacroChinaHgjck 东方财富-海关进出口增减情况一览表
//
// 返回:
//   - dataframe.DataFrame: 海关进出口增减情况一览表
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/hgjck.html
func MacroChinaHgjck() (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"columns":     "REPORT_DATE,TIME,EXIT_BASE,IMPORT_BASE,EXIT_BASE_SAME,IMPORT_BASE_SAME,EXIT_BASE_SEQUENTIAL,IMPORT_BASE_SEQUENTIAL,EXIT_ACCUMULATE,IMPORT_ACCUMULATE,EXIT_ACCUMULATE_SAME,IMPORT_ACCUMULATE_SAME",
		"pageNumber":  "1",
		"pageSize":    "2000",
		"sortColumns": "REPORT_DATE",
		"sortTypes":   "-1",
		"source":      "WEB",
		"client":      "WEB",
		"reportName":  "RPT_ECONOMY_CUSTOMS",
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
		"当月出口额-金额", "当月出口额-同比增长", "当月出口额-环比增长",
		"当月进口额-金额", "当月进口额-同比增长", "当月进口额-环比增长",
		"累计出口额-金额", "累计出口额-同比增长",
		"累计进口额-金额", "累计进口额-同比增长",
	}
	var records [][]string
	records = append(records, columns)

	for _, item := range data {
		record := []string{
			item.Get("TIME").String(),
			item.Get("EXIT_BASE").String(),
			item.Get("EXIT_BASE_SAME").String(),
			item.Get("EXIT_BASE_SEQUENTIAL").String(),
			item.Get("IMPORT_BASE").String(),
			item.Get("IMPORT_BASE_SAME").String(),
			item.Get("IMPORT_BASE_SEQUENTIAL").String(),
			item.Get("EXIT_ACCUMULATE").String(),
			item.Get("EXIT_ACCUMULATE_SAME").String(),
			item.Get("IMPORT_ACCUMULATE").String(),
			item.Get("IMPORT_ACCUMULATE_SAME").String(),
		}
		records = append(records, record)
	}

	df := dataframe.LoadRecords(records)
	return df, nil
}

// MacroChinaCzsr 东方财富-财政收入
//
// 返回:
//   - dataframe.DataFrame: 财政收入
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/czsr.html
func MacroChinaCzsr() (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"columns":     "REPORT_DATE,TIME,BASE,BASE_SAME,BASE_SEQUENTIAL,BASE_ACCUMULATE,ACCUMULATE_SAME",
		"pageNumber":  "1",
		"pageSize":    "2000",
		"sortColumns": "REPORT_DATE",
		"sortTypes":   "-1",
		"source":      "WEB",
		"client":      "WEB",
		"reportName":  "RPT_ECONOMY_INCOME",
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
			item.Get("BASE").String(),
			item.Get("BASE_SAME").String(),
			item.Get("BASE_SEQUENTIAL").String(),
			item.Get("BASE_ACCUMULATE").String(),
			item.Get("ACCUMULATE_SAME").String(),
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

// MacroChinaWhxd 东方财富-外汇贷款数据
//
// 返回:
//   - dataframe.DataFrame: 外汇贷款数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/whxd.html
func MacroChinaWhxd() (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"columns":     "REPORT_DATE,TIME,BASE,BASE_SAME,BASE_SEQUENTIAL,BASE_ACCUMULATE",
		"pageNumber":  "1",
		"pageSize":    "2000",
		"sortColumns": "REPORT_DATE",
		"sortTypes":   "-1",
		"source":      "WEB",
		"client":      "WEB",
		"reportName":  "RPT_ECONOMY_FOREX_LOAN",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	data := gjson.Get(resp.String(), "result.data").Array()
	if len(data) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	columns := []string{"月份", "当月", "同比增长", "环比增长", "累计"}
	var records [][]string
	records = append(records, columns)

	for _, item := range data {
		record := []string{
			item.Get("TIME").String(),
			item.Get("BASE").String(),
			item.Get("BASE_SAME").String(),
			item.Get("BASE_SEQUENTIAL").String(),
			item.Get("BASE_ACCUMULATE").String(),
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

// MacroChinaWbck 东方财富-本外币存款
//
// 返回:
//   - dataframe.DataFrame: 本外币存款
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/wbck.html
func MacroChinaWbck() (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"columns":     "REPORT_DATE,TIME,BASE,BASE_SAME,BASE_SEQUENTIAL,BASE_ACCUMULATE",
		"pageNumber":  "1",
		"pageSize":    "2000",
		"sortColumns": "REPORT_DATE",
		"sortTypes":   "-1",
		"source":      "WEB",
		"client":      "WEB",
		"reportName":  "RPT_ECONOMY_FOREX_DEPOSIT",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	data := gjson.Get(resp.String(), "result.data").Array()
	if len(data) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	columns := []string{"月份", "当月", "同比增长", "环比增长", "累计"}
	var records [][]string
	records = append(records, columns)

	for _, item := range data {
		record := []string{
			item.Get("TIME").String(),
			item.Get("BASE").String(),
			item.Get("BASE_SAME").String(),
			item.Get("BASE_SEQUENTIAL").String(),
			item.Get("BASE_ACCUMULATE").String(),
		}
		records = append(records, record)
	}

	df := dataframe.LoadRecords(records)
	return df, nil
}

// MacroChinaXfzxx 东方财富-消费者信心指数
//
// 返回:
//   - dataframe.DataFrame: 消费者信心指数
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/xfzxx.html
func MacroChinaXfzxx() (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"columns":     "REPORT_DATE,TIME,CONSUMERS_FAITH_INDEX,FAITH_INDEX_SAME,FAITH_INDEX_SEQUENTIAL,CONSUMERS_ASTIS_INDEX,ASTIS_INDEX_SAME,ASTIS_INDEX_SEQUENTIAL,CONSUMERS_EXPECT_INDEX,EXPECT_INDEX_SAME,EXPECT_INDEX_SEQUENTIAL",
		"pageNumber":  "1",
		"pageSize":    "2000",
		"sortColumns": "REPORT_DATE",
		"sortTypes":   "-1",
		"source":      "WEB",
		"client":      "WEB",
		"reportName":  "RPT_ECONOMY_FAITH_INDEX",
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
		"消费者信心指数-指数值", "消费者信心指数-同比增长", "消费者信心指数-环比增长",
		"消费者满意指数-指数值", "消费者满意指数-同比增长", "消费者满意指数-环比增长",
		"消费者预期指数-指数值", "消费者预期指数-同比增长", "消费者预期指数-环比增长",
	}
	var records [][]string
	records = append(records, columns)

	for _, item := range data {
		record := []string{
			item.Get("TIME").String(),
			item.Get("CONSUMERS_FAITH_INDEX").String(),
			item.Get("FAITH_INDEX_SAME").String(),
			item.Get("FAITH_INDEX_SEQUENTIAL").String(),
			item.Get("CONSUMERS_ASTIS_INDEX").String(),
			item.Get("ASTIS_INDEX_SAME").String(),
			item.Get("ASTIS_INDEX_SEQUENTIAL").String(),
			item.Get("CONSUMERS_EXPECT_INDEX").String(),
			item.Get("EXPECT_INDEX_SAME").String(),
			item.Get("EXPECT_INDEX_SEQUENTIAL").String(),
		}
		records = append(records, record)
	}

	df := dataframe.LoadRecords(records)
	return df, nil
}

// MacroChinaGyzjz 东方财富-工业增加值增长
//
// 返回:
//   - dataframe.DataFrame: 工业增加值增长
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/gyzjz.html
func MacroChinaGyzjz() (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"columns":     "REPORT_DATE,TIME,BASE_SAME,BASE_ACCUMULATE",
		"pageNumber":  "1",
		"pageSize":    "2000",
		"sortColumns": "REPORT_DATE",
		"sortTypes":   "-1",
		"source":      "WEB",
		"client":      "WEB",
		"reportName":  "RPT_ECONOMY_INDUS_GROW",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	data := gjson.Get(resp.String(), "result.data").Array()
	if len(data) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	columns := []string{"月份", "同比增长", "累计增长", "发布时间"}
	var records [][]string
	records = append(records, columns)

	for _, item := range data {
		date := item.Get("REPORT_DATE").String()
		if len(date) >= 10 {
			date = date[:10]
		}
		record := []string{
			item.Get("TIME").String(),
			item.Get("BASE_SAME").String(),
			item.Get("BASE_ACCUMULATE").String(),
			date,
		}
		records = append(records, record)
	}

	// 按发布时间排序
	sort.Slice(records[1:], func(i, j int) bool {
		return records[i+1][3] < records[j+1][3]
	})

	df := dataframe.LoadRecords(records)
	return df, nil
}

// MacroChinaReserveRequirementRatio 存款准备金率
//
// 返回:
//   - dataframe.DataFrame: 存款准备金率
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/ckzbj.html
func MacroChinaReserveRequirementRatio() (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"columns":     "REPORT_DATE,PUBLISH_DATE,TRADE_DATE,INTEREST_RATE_BB,INTEREST_RATE_BA,CHANGE_RATE_B,INTEREST_RATE_SB,INTEREST_RATE_SA,CHANGE_RATE_S,NEXT_SH_RATE,NEXT_SZ_RATE,REMARK",
		"pageNumber":  "1",
		"pageSize":    "2000",
		"sortColumns": "PUBLISH_DATE,TRADE_DATE",
		"sortTypes":   "-1,-1",
		"source":      "WEB",
		"client":      "WEB",
		"reportName":  "RPT_ECONOMY_DEPOSIT_RESERVE",
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
		"公布时间", "生效时间",
		"大型金融机构-调整前", "大型金融机构-调整后", "大型金融机构-调整幅度",
		"中小金融机构-调整前", "中小金融机构-调整后", "中小金融机构-调整幅度",
		"消息公布次日指数涨跌-上证", "消息公布次日指数涨跌-深证",
		"备注",
	}
	var records [][]string
	records = append(records, columns)

	for _, item := range data {
		publishDate := item.Get("PUBLISH_DATE").String()
		if len(publishDate) >= 10 {
			publishDate = publishDate[:10]
		}
		tradeDate := item.Get("TRADE_DATE").String()
		if len(tradeDate) >= 10 {
			tradeDate = tradeDate[:10]
		}
		record := []string{
			publishDate,
			tradeDate,
			item.Get("INTEREST_RATE_BB").String(),
			item.Get("INTEREST_RATE_BA").String(),
			item.Get("CHANGE_RATE_B").String(),
			item.Get("INTEREST_RATE_SB").String(),
			item.Get("INTEREST_RATE_SA").String(),
			item.Get("CHANGE_RATE_S").String(),
			item.Get("NEXT_SH_RATE").String(),
			item.Get("NEXT_SZ_RATE").String(),
			item.Get("REMARK").String(),
		}
		records = append(records, record)
	}

	df := dataframe.LoadRecords(records)
	return df, nil
}

// MacroChinaConsumerGoodsRetail 东方财富-社会消费品零售总额
//
// 返回:
//   - dataframe.DataFrame: 社会消费品零售总额
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/xfp.html
func MacroChinaConsumerGoodsRetail() (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"columns":     "REPORT_DATE,TIME,RETAIL_TOTAL,RETAIL_TOTAL_SAME,RETAIL_TOTAL_SEQUENTIAL,RETAIL_TOTAL_ACCUMULATE,RETAIL_ACCUMULATE_SAME",
		"pageNumber":  "1",
		"pageSize":    "1000",
		"sortColumns": "REPORT_DATE",
		"sortTypes":   "-1",
		"source":      "WEB",
		"client":      "WEB",
		"reportName":  "RPT_ECONOMY_TOTAL_RETAIL",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	data := gjson.Get(resp.String(), "result.data").Array()
	if len(data) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	columns := []string{"月份", "当月", "同比增长", "环比增长", "累计", "累计-同比增长"}
	var records [][]string
	records = append(records, columns)

	for _, item := range data {
		record := []string{
			item.Get("TIME").String(),
			item.Get("RETAIL_TOTAL").String(),
			item.Get("RETAIL_TOTAL_SAME").String(),
			item.Get("RETAIL_TOTAL_SEQUENTIAL").String(),
			item.Get("RETAIL_TOTAL_ACCUMULATE").String(),
			item.Get("RETAIL_ACCUMULATE_SAME").String(),
		}
		records = append(records, record)
	}

	df := dataframe.LoadRecords(records)
	return df, nil
}
