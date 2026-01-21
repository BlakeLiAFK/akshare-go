package economic

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
	"github.com/tidwall/gjson"
)

// fetchChinaJin10Data 金十数据通用获取函数
func fetchChinaJin10Data(attrID int, symbol string) (dataframe.DataFrame, error) {
	headers := map[string]string{
		"user-agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36",
		"x-app-id":     "rU6QIu7JHe2gOUeR",
		"x-csrf-token": "x-csrf-token",
		"x-version":    "1.0.0",
	}

	url := "https://datacenter-api.jin10.com/reports/list_v2"
	columns := []string{"商品", "日期", "今值", "预测值", "前值"}
	var allRecords [][]string
	allRecords = append(allRecords, columns)

	maxDate := ""
	for {
		params := map[string]string{
			"max_date": maxDate,
			"category": "ec",
			"attr_id":  fmt.Sprintf("%d", attrID),
			"_":        fmt.Sprintf("%d", time.Now().UnixMilli()),
		}

		resp, err := utils.GetWithHeaders(url, params, headers)
		if err != nil {
			return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
		}

		values := gjson.Get(resp.String(), "data.values").Array()
		if len(values) == 0 {
			break
		}

		for _, item := range values {
			arr := item.Array()
			if len(arr) >= 4 {
				date := arr[0].String()
				actual := arr[1].String()
				forecast := arr[2].String()
				previous := arr[3].String()
				allRecords = append(allRecords, []string{symbol, date, actual, forecast, previous})
				maxDate = date
			}
		}

		// 计算前一天日期
		t, err := time.Parse("2006-01-02", maxDate)
		if err != nil {
			break
		}
		maxDate = t.AddDate(0, 0, -1).Format("2006-01-02")
	}

	if len(allRecords) <= 1 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	// 按日期排序
	sort.Slice(allRecords[1:], func(i, j int) bool {
		return allRecords[i+1][1] < allRecords[j+1][1]
	})

	df := dataframe.LoadRecords(allRecords)
	return df, nil
}

// fetchEastmoneyEconomicData 东方财富经济数据通用获取函数
func fetchEastmoneyEconomicData(reportName, columns string, colRename map[string]string) (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"columns":     columns,
		"pageNumber":  "1",
		"pageSize":    "2000",
		"sortColumns": "REPORT_DATE",
		"sortTypes":   "-1",
		"source":      "WEB",
		"client":      "WEB",
		"reportName":  reportName,
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	data := gjson.Get(resp.String(), "result.data").Array()
	if len(data) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	// 获取列名
	var header []string
	var selectedCols []string
	for k, v := range colRename {
		selectedCols = append(selectedCols, k)
		header = append(header, v)
	}

	var records [][]string
	records = append(records, header)

	for _, item := range data {
		var record []string
		for _, col := range selectedCols {
			val := item.Get(col).String()
			record = append(record, val)
		}
		records = append(records, record)
	}

	// 按第一列排序
	if len(records) > 1 {
		sort.Slice(records[1:], func(i, j int) bool {
			return records[i+1][0] < records[j+1][0]
		})
	}

	df := dataframe.LoadRecords(records)
	return df, nil
}

// MacroChinaQyspjg 东方财富-经济数据一览-中国-企业商品价格指数
//
// 返回:
//   - dataframe.DataFrame: 企业商品价格指数
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/qyspjg.html
func MacroChinaQyspjg() (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"columns":     "REPORT_DATE,TIME,BASE,BASE_SAME,BASE_SEQUENTIAL,FARM_BASE,FARM_BASE_SAME,FARM_BASE_SEQUENTIAL,MINERAL_BASE,MINERAL_BASE_SAME,MINERAL_BASE_SEQUENTIAL,ENERGY_BASE,ENERGY_BASE_SAME,ENERGY_BASE_SEQUENTIAL",
		"pageNumber":  "1",
		"pageSize":    "2000",
		"sortColumns": "REPORT_DATE",
		"sortTypes":   "-1",
		"source":      "WEB",
		"client":      "WEB",
		"reportName":  "RPT_ECONOMY_GOODS_INDEX",
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
		"月份", "总指数-指数值", "总指数-同比增长", "总指数-环比增长",
		"农产品-指数值", "农产品-同比增长", "农产品-环比增长",
		"矿产品-指数值", "矿产品-同比增长", "矿产品-环比增长",
		"煤油电-指数值", "煤油电-同比增长", "煤油电-环比增长",
	}

	var records [][]string
	records = append(records, columns)

	for _, item := range data {
		record := []string{
			item.Get("TIME").String(),
			item.Get("BASE").String(),
			item.Get("BASE_SAME").String(),
			item.Get("BASE_SEQUENTIAL").String(),
			item.Get("FARM_BASE").String(),
			item.Get("FARM_BASE_SAME").String(),
			item.Get("FARM_BASE_SEQUENTIAL").String(),
			item.Get("MINERAL_BASE").String(),
			item.Get("MINERAL_BASE_SAME").String(),
			item.Get("MINERAL_BASE_SEQUENTIAL").String(),
			item.Get("ENERGY_BASE").String(),
			item.Get("ENERGY_BASE_SAME").String(),
			item.Get("ENERGY_BASE_SEQUENTIAL").String(),
		}
		records = append(records, record)
	}

	df := dataframe.LoadRecords(records)
	return df, nil
}

// MacroChinaFDI 东方财富-经济数据一览-中国-外商直接投资数据
//
// 返回:
//   - dataframe.DataFrame: 外商直接投资数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/fdi.html
func MacroChinaFDI() (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"columns":     "REPORT_DATE,TIME,ACTUAL_FOREIGN,ACTUAL_FOREIGN_SAME,ACTUAL_FOREIGN_SEQUENTIAL,ACTUAL_FOREIGN_ACCUMULATE,FOREIGN_ACCUMULATE_SAME",
		"pageNumber":  "1",
		"pageSize":    "2000",
		"sortColumns": "REPORT_DATE",
		"sortTypes":   "-1",
		"source":      "WEB",
		"client":      "WEB",
		"reportName":  "RPT_ECONOMY_FDI",
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
			item.Get("ACTUAL_FOREIGN").String(),
			item.Get("ACTUAL_FOREIGN_SAME").String(),
			item.Get("ACTUAL_FOREIGN_SEQUENTIAL").String(),
			item.Get("ACTUAL_FOREIGN_ACCUMULATE").String(),
			item.Get("FOREIGN_ACCUMULATE_SAME").String(),
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

// MacroChinaShrzgm 商务数据中心-国内贸易-社会融资规模增量统计
//
// 返回:
//   - dataframe.DataFrame: 社会融资规模增量统计
//   - error: 错误信息
//
// 数据源: https://data.mofcom.gov.cn/gnmy/shrzgm.shtml
func MacroChinaShrzgm() (dataframe.DataFrame, error) {
	url := "https://data.mofcom.gov.cn/datamofcom/front/gnmy/shrzgmQuery"

	resp, err := utils.Post(url, nil)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	data := gjson.Parse(resp.String()).Array()
	if len(data) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	columns := []string{
		"月份", "社会融资规模增量", "其中-人民币贷款", "其中-委托贷款外币贷款",
		"其中-委托贷款", "其中-信托贷款", "其中-未贴现银行承兑汇票",
		"其中-企业债券", "其中-非金融企业境内股票融资",
	}

	var records [][]string
	records = append(records, columns)

	for _, item := range data {
		record := []string{
			item.Get("dataDate").String(),
			item.Get("shrzgmzl").String(),
			item.Get("rmbdk").String(),
			item.Get("wtdkwbdk").String(),
			item.Get("wtdk").String(),
			item.Get("xtdk").String(),
			item.Get("wtpyhcdp").String(),
			item.Get("qyzq").String(),
			item.Get("fjrqyjngpRz").String(),
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

// MacroChinaUrbanUnemployment 国家统计局-月度数据-城镇调查失业率
//
// 返回:
//   - dataframe.DataFrame: 城镇调查失业率
//   - error: 错误信息
//
// 数据源: https://data.stats.gov.cn/easyquery.htm
func MacroChinaUrbanUnemployment() (dataframe.DataFrame, error) {
	url := "https://data.stats.gov.cn/easyquery.htm"
	params := map[string]string{
		"m":       "QueryData",
		"dbcode":  "hgyd",
		"rowcode": "zb",
		"colcode": "sj",
		"wds":     "[]",
		"dfwds":   `[{"wdcode":"zb","valuecode":"A0E01"},{"wdcode":"sj","valuecode":"LAST72"}]`,
		"k1":      fmt.Sprintf("%d", time.Now().UnixMilli()),
		"h":       "1",
	}

	resp, err := nbsClient.R().SetQueryParams(params).Get(url)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	result := gjson.Parse(resp.String())
	datanodes := result.Get("returndata.datanodes").Array()
	if len(datanodes) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	// 获取指标名称映射
	wdnodes := result.Get("returndata.wdnodes.0.nodes").Array()
	codeNameMap := make(map[string]string)
	for _, node := range wdnodes {
		code := node.Get("code").String()
		name := node.Get("cname").String()
		codeNameMap[code] = name
	}

	columns := []string{"date", "item", "value"}
	var records [][]string
	records = append(records, columns)

	for _, node := range datanodes {
		date := node.Get("wds.1.valuecode").String()
		itemCode := node.Get("wds.0.valuecode").String()
		value := node.Get("data.data").String()

		itemName := codeNameMap[itemCode]
		if itemName == "" {
			itemName = itemCode
		}

		records = append(records, []string{date, itemName, value})
	}

	// 按日期排序
	sort.Slice(records[1:], func(i, j int) bool {
		return records[i+1][0] < records[j+1][0]
	})

	df := dataframe.LoadRecords(records)
	return df, nil
}

// MacroChinaGDPYearly 金十数据中心-中国GDP年率报告
//
// 返回:
//   - dataframe.DataFrame: 中国GDP年率报告
//   - error: 错误信息
//
// 数据源: https://datacenter.jin10.com/reportType/dc_chinese_gdp_yoy
func MacroChinaGDPYearly() (dataframe.DataFrame, error) {
	return fetchChinaJin10Data(57, "中国GDP年率报告")
}

// MacroChinaCPIYearly 金十数据中心-中国CPI年率报告
//
// 返回:
//   - dataframe.DataFrame: 中国CPI年率报告
//   - error: 错误信息
//
// 数据源: https://datacenter.jin10.com/reportType/dc_chinese_cpi_yoy
func MacroChinaCPIYearly() (dataframe.DataFrame, error) {
	return fetchChinaJin10Data(56, "中国CPI年率报告")
}

// MacroChinaCPIMonthly 金十数据中心-中国CPI月率报告
//
// 返回:
//   - dataframe.DataFrame: 中国CPI月率报告
//   - error: 错误信息
//
// 数据源: https://datacenter.jin10.com/reportType/dc_chinese_cpi_mom
func MacroChinaCPIMonthly() (dataframe.DataFrame, error) {
	return fetchChinaJin10Data(72, "中国CPI月率报告")
}

// MacroChinaPPIYearly 金十数据中心-中国PPI年率报告
//
// 返回:
//   - dataframe.DataFrame: 中国PPI年率报告
//   - error: 错误信息
//
// 数据源: https://datacenter.jin10.com/reportType/dc_chinese_ppi_yoy
func MacroChinaPPIYearly() (dataframe.DataFrame, error) {
	return fetchChinaJin10Data(60, "中国PPI年率报告")
}

// MacroChinaExportsYoY 金十数据中心-中国以美元计算出口年率报告
//
// 返回:
//   - dataframe.DataFrame: 中国以美元计算出口年率报告
//   - error: 错误信息
//
// 数据源: https://datacenter.jin10.com/reportType/dc_chinese_exports_yoy
func MacroChinaExportsYoY() (dataframe.DataFrame, error) {
	return fetchChinaJin10Data(66, "中国以美元计算出口年率报告")
}

// MacroChinaImportsYoY 金十数据中心-中国以美元计算进口年率报告
//
// 返回:
//   - dataframe.DataFrame: 中国以美元计算进口年率报告
//   - error: 错误信息
//
// 数据源: https://datacenter.jin10.com/reportType/dc_chinese_imports_yoy
func MacroChinaImportsYoY() (dataframe.DataFrame, error) {
	return fetchChinaJin10Data(77, "中国以美元计算进口年率报告")
}

// MacroChinaTradeBalance 金十数据中心-中国以美元计算贸易帐报告
//
// 返回:
//   - dataframe.DataFrame: 中国以美元计算贸易帐报告
//   - error: 错误信息
//
// 数据源: https://datacenter.jin10.com/reportType/dc_chinese_trade_balance
func MacroChinaTradeBalance() (dataframe.DataFrame, error) {
	return fetchChinaJin10Data(61, "中国以美元计算贸易帐报告")
}

// MacroChinaIndustrialProductionYoY 金十数据中心-中国规模以上工业增加值年率报告
//
// 返回:
//   - dataframe.DataFrame: 中国规模以上工业增加值年率报告
//   - error: 错误信息
//
// 数据源: https://datacenter.jin10.com/reportType/dc_chinese_industrial_production_yoy
func MacroChinaIndustrialProductionYoY() (dataframe.DataFrame, error) {
	return fetchChinaJin10Data(58, "中国规模以上工业增加值年率报告")
}

// MacroChinaPMIYearly 金十数据中心-中国官方制造业PMI报告
//
// 返回:
//   - dataframe.DataFrame: 中国官方制造业PMI报告
//   - error: 错误信息
//
// 数据源: https://datacenter.jin10.com/reportType/dc_chinese_manufacturing_pmi
func MacroChinaPMIYearly() (dataframe.DataFrame, error) {
	return fetchChinaJin10Data(65, "中国官方制造业PMI报告")
}

// MacroChinaCxPMIYearly 金十数据中心-中国财新制造业PMI终值报告
//
// 返回:
//   - dataframe.DataFrame: 中国财新制造业PMI终值报告
//   - error: 错误信息
//
// 数据源: https://datacenter.jin10.com/reportType/dc_chinese_caixin_manufacturing_pmi
func MacroChinaCxPMIYearly() (dataframe.DataFrame, error) {
	return fetchChinaJin10Data(73, "中国财新制造业PMI终值报告")
}

// MacroChinaCxServicesPMIYearly 金十数据中心-中国财新服务业PMI报告
//
// 返回:
//   - dataframe.DataFrame: 中国财新服务业PMI报告
//   - error: 错误信息
//
// 数据源: https://datacenter.jin10.com/reportType/dc_chinese_caixin_services_pmi
func MacroChinaCxServicesPMIYearly() (dataframe.DataFrame, error) {
	return fetchChinaJin10Data(74, "中国财新服务业PMI报告")
}

// MacroChinaNonManPMI 金十数据中心-中国官方非制造业PMI报告
//
// 返回:
//   - dataframe.DataFrame: 中国官方非制造业PMI报告
//   - error: 错误信息
//
// 数据源: https://datacenter.jin10.com/reportType/dc_chinese_non_manufacturing_pmi
func MacroChinaNonManPMI() (dataframe.DataFrame, error) {
	return fetchChinaJin10Data(75, "中国官方非制造业PMI报告")
}

// MacroChinaFxReservesYearly 金十数据中心-中国外汇储备报告
//
// 返回:
//   - dataframe.DataFrame: 中国外汇储备报告
//   - error: 错误信息
//
// 数据源: https://datacenter.jin10.com/reportType/dc_chinese_fx_reserves
func MacroChinaFxReservesYearly() (dataframe.DataFrame, error) {
	return fetchChinaJin10Data(64, "中国外汇储备报告")
}

// MacroChinaM2Yearly 金十数据中心-中国M2货币供应年率报告
//
// 返回:
//   - dataframe.DataFrame: 中国M2货币供应年率报告
//   - error: 错误信息
//
// 数据源: https://datacenter.jin10.com/reportType/dc_chinese_m2_money_supply_yoy
func MacroChinaM2Yearly() (dataframe.DataFrame, error) {
	return fetchChinaJin10Data(59, "中国M2货币供应年率报告")
}

// MacroChinaShiborAll 上海银行同业拆借利率(SHIBOR)
//
// 返回:
//   - dataframe.DataFrame: SHIBOR利率数据
//   - error: 错误信息
//
// 数据源: https://datacenter.jin10.com/reportType/dc_shibor
func MacroChinaShiborAll() (dataframe.DataFrame, error) {
	headers := map[string]string{
		"user-agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36",
		"x-app-id":     "rU6QIu7JHe2gOUeR",
		"x-csrf-token": "x-csrf-token",
		"x-version":    "1.0.0",
	}

	url := "https://datacenter-api.jin10.com/reports/list_v2"
	columns := []string{"日期", "O/N", "1W", "2W", "1M", "3M", "6M", "9M", "1Y"}
	var allRecords [][]string
	allRecords = append(allRecords, columns)

	maxDate := ""
	for {
		params := map[string]string{
			"max_date": maxDate,
			"category": "ec",
			"attr_id":  "24",
			"_":        fmt.Sprintf("%d", time.Now().UnixMilli()),
		}

		resp, err := utils.GetWithHeaders(url, params, headers)
		if err != nil {
			return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
		}

		values := gjson.Get(resp.String(), "data.values").Array()
		if len(values) == 0 {
			break
		}

		for _, item := range values {
			arr := item.Array()
			if len(arr) >= 9 {
				var record []string
				for i := 0; i < 9; i++ {
					record = append(record, arr[i].String())
				}
				allRecords = append(allRecords, record)
				maxDate = arr[0].String()
			}
		}

		t, err := time.Parse("2006-01-02", maxDate)
		if err != nil {
			break
		}
		maxDate = t.AddDate(0, 0, -1).Format("2006-01-02")
	}

	if len(allRecords) <= 1 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	sort.Slice(allRecords[1:], func(i, j int) bool {
		return allRecords[i+1][0] < allRecords[j+1][0]
	})

	df := dataframe.LoadRecords(allRecords)
	return df, nil
}

// MacroChinaHKMarketInfo 东方财富-沪港通资金流向
//
// 返回:
//   - dataframe.DataFrame: 沪港通资金流向数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/hsgt/index.html
func MacroChinaHKMarketInfo() (dataframe.DataFrame, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"columns":     "TRADE_DATE,MUTUAL_TYPE,BOARD_NAME,NET_INFLOW,BUY_AMT,SELL_AMT,QUOTA,QUOTA_REMAIN,NET_INFLOW_TOTAL,BUY_AMT_TOTAL,SELL_AMT_TOTAL",
		"pageNumber":  "1",
		"pageSize":    "5000",
		"sortColumns": "TRADE_DATE,MUTUAL_TYPE",
		"sortTypes":   "-1,1",
		"source":      "WEB",
		"client":      "WEB",
		"reportName":  "RPT_MUTUAL_QUOTA",
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
		"交易日期", "类型", "板块名称", "当日净流入", "买入成交金额",
		"卖出成交金额", "额度", "额度余额", "累计净流入", "累计买入成交金额", "累计卖出成交金额",
	}

	var records [][]string
	records = append(records, columns)

	for _, item := range data {
		record := []string{
			item.Get("TRADE_DATE").String()[:10],
			item.Get("MUTUAL_TYPE").String(),
			item.Get("BOARD_NAME").String(),
			item.Get("NET_INFLOW").String(),
			item.Get("BUY_AMT").String(),
			item.Get("SELL_AMT").String(),
			item.Get("QUOTA").String(),
			item.Get("QUOTA_REMAIN").String(),
			item.Get("NET_INFLOW_TOTAL").String(),
			item.Get("BUY_AMT_TOTAL").String(),
			item.Get("SELL_AMT_TOTAL").String(),
		}
		records = append(records, record)
	}

	df := dataframe.LoadRecords(records)
	return df, nil
}

// MacroChinaDailyEnergy 金十数据-日度能源数据
//
// 返回:
//   - dataframe.DataFrame: 日度能源数据
//   - error: 错误信息
//
// 数据源: https://datacenter.jin10.com/
func MacroChinaDailyEnergy() (dataframe.DataFrame, error) {
	url := "https://cdn.jin10.com/data_center/reports/fs_2.json"

	resp, err := utils.Get(url, nil)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	values := gjson.Get(resp.String(), "values").Array()
	if len(values) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	columns := []string{"日期", "焦煤", "焦炭", "动力煤", "秦皇岛港库存"}
	var records [][]string
	records = append(records, columns)

	for _, item := range values {
		arr := item.Array()
		if len(arr) >= 5 {
			record := []string{
				arr[0].String(),
				arr[1].String(),
				arr[2].String(),
				arr[3].String(),
				arr[4].String(),
			}
			records = append(records, record)
		}
	}

	sort.Slice(records[1:], func(i, j int) bool {
		return records[i+1][0] < records[j+1][0]
	})

	df := dataframe.LoadRecords(records)
	return df, nil
}

// MacroChinaRMB 人民币汇率中间价
//
// 返回:
//   - dataframe.DataFrame: 人民币汇率中间价
//   - error: 错误信息
//
// 数据源: https://datacenter.jin10.com/reportType/dc_rmb_data
func MacroChinaRMB() (dataframe.DataFrame, error) {
	headers := map[string]string{
		"user-agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36",
		"x-app-id":     "rU6QIu7JHe2gOUeR",
		"x-csrf-token": "x-csrf-token",
		"x-version":    "1.0.0",
	}

	url := "https://datacenter-api.jin10.com/reports/list_v2"
	columns := []string{
		"日期", "美元/人民币", "欧元/人民币", "100日元/人民币", "港元/人民币",
		"英镑/人民币", "澳元/人民币", "新西兰元/人民币", "新加坡元/人民币",
		"瑞郎/人民币", "加元/人民币", "林吉特/人民币", "卢布/人民币",
		"兰特/人民币", "韩元/人民币", "迪拉姆/人民币", "里亚尔/人民币",
		"福林/人民币", "兹罗提/人民币", "丹麦克朗/人民币", "瑞典克朗/人民币",
		"挪威克朗/人民币", "里拉/人民币", "比索/人民币", "泰铢/人民币",
	}
	var allRecords [][]string
	allRecords = append(allRecords, columns)

	maxDate := ""
	for {
		params := map[string]string{
			"max_date": maxDate,
			"category": "ec",
			"attr_id":  "23",
			"_":        fmt.Sprintf("%d", time.Now().UnixMilli()),
		}

		resp, err := utils.GetWithHeaders(url, params, headers)
		if err != nil {
			return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
		}

		values := gjson.Get(resp.String(), "data.values").Array()
		if len(values) == 0 {
			break
		}

		for _, item := range values {
			arr := item.Array()
			var record []string
			for i := 0; i < len(columns) && i < len(arr); i++ {
				record = append(record, arr[i].String())
			}
			// 补齐列
			for len(record) < len(columns) {
				record = append(record, "")
			}
			allRecords = append(allRecords, record)
			if len(arr) > 0 {
				maxDate = arr[0].String()
			}
		}

		t, err := time.Parse("2006-01-02", maxDate)
		if err != nil {
			break
		}
		maxDate = t.AddDate(0, 0, -1).Format("2006-01-02")
	}

	if len(allRecords) <= 1 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	sort.Slice(allRecords[1:], func(i, j int) bool {
		return allRecords[i+1][0] < allRecords[j+1][0]
	})

	df := dataframe.LoadRecords(allRecords)
	return df, nil
}

// MacroChinaMarketMarginSZ 深圳融资融券数据
//
// 返回:
//   - dataframe.DataFrame: 深圳融资融券数据
//   - error: 错误信息
//
// 数据源: https://datacenter.jin10.com/
func MacroChinaMarketMarginSZ() (dataframe.DataFrame, error) {
	headers := map[string]string{
		"user-agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36",
		"x-app-id":     "rU6QIu7JHe2gOUeR",
		"x-csrf-token": "x-csrf-token",
		"x-version":    "1.0.0",
	}

	url := "https://datacenter-api.jin10.com/reports/list_v2"
	columns := []string{"日期", "融资余额", "融券余额", "融资买入额", "融资偿还额", "融券卖出量", "融券偿还量"}
	var allRecords [][]string
	allRecords = append(allRecords, columns)

	maxDate := ""
	for {
		params := map[string]string{
			"max_date": maxDate,
			"category": "stock",
			"attr_id":  "7",
			"_":        fmt.Sprintf("%d", time.Now().UnixMilli()),
		}

		resp, err := utils.GetWithHeaders(url, params, headers)
		if err != nil {
			return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
		}

		values := gjson.Get(resp.String(), "data.values").Array()
		if len(values) == 0 {
			break
		}

		for _, item := range values {
			arr := item.Array()
			if len(arr) >= 7 {
				record := []string{
					arr[0].String(),
					arr[1].String(),
					arr[2].String(),
					arr[3].String(),
					arr[4].String(),
					arr[5].String(),
					arr[6].String(),
				}
				allRecords = append(allRecords, record)
				maxDate = arr[0].String()
			}
		}

		t, err := time.Parse("2006-01-02", maxDate)
		if err != nil {
			break
		}
		maxDate = t.AddDate(0, 0, -1).Format("2006-01-02")
	}

	if len(allRecords) <= 1 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	sort.Slice(allRecords[1:], func(i, j int) bool {
		return allRecords[i+1][0] < allRecords[j+1][0]
	})

	df := dataframe.LoadRecords(allRecords)
	return df, nil
}

// MacroChinaMarketMarginSH 上海融资融券数据
//
// 返回:
//   - dataframe.DataFrame: 上海融资融券数据
//   - error: 错误信息
//
// 数据源: https://datacenter.jin10.com/
func MacroChinaMarketMarginSH() (dataframe.DataFrame, error) {
	headers := map[string]string{
		"user-agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36",
		"x-app-id":     "rU6QIu7JHe2gOUeR",
		"x-csrf-token": "x-csrf-token",
		"x-version":    "1.0.0",
	}

	url := "https://datacenter-api.jin10.com/reports/list_v2"
	columns := []string{"日期", "融资余额", "融券余额", "融资买入额", "融资偿还额", "融券卖出量", "融券偿还量"}
	var allRecords [][]string
	allRecords = append(allRecords, columns)

	maxDate := ""
	for {
		params := map[string]string{
			"max_date": maxDate,
			"category": "stock",
			"attr_id":  "8",
			"_":        fmt.Sprintf("%d", time.Now().UnixMilli()),
		}

		resp, err := utils.GetWithHeaders(url, params, headers)
		if err != nil {
			return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
		}

		values := gjson.Get(resp.String(), "data.values").Array()
		if len(values) == 0 {
			break
		}

		for _, item := range values {
			arr := item.Array()
			if len(arr) >= 7 {
				record := []string{
					arr[0].String(),
					arr[1].String(),
					arr[2].String(),
					arr[3].String(),
					arr[4].String(),
					arr[5].String(),
					arr[6].String(),
				}
				allRecords = append(allRecords, record)
				maxDate = arr[0].String()
			}
		}

		t, err := time.Parse("2006-01-02", maxDate)
		if err != nil {
			break
		}
		maxDate = t.AddDate(0, 0, -1).Format("2006-01-02")
	}

	if len(allRecords) <= 1 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	sort.Slice(allRecords[1:], func(i, j int) bool {
		return allRecords[i+1][0] < allRecords[j+1][0]
	})

	df := dataframe.LoadRecords(allRecords)
	return df, nil
}

// 确保 strings 包被使用
var _ = strings.TrimSpace
