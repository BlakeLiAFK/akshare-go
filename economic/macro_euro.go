package economic

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
	"github.com/tidwall/gjson"
)

// jin10Headers 金十数据中心请求头
var jin10Headers = map[string]string{
	"user-agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36",
	"x-app-id":     "rU6QIu7JHe2gOUeR",
	"x-csrf-token": "x-csrf-token",
	"x-version":    "1.0.0",
}

// fetchJin10EconomicData 获取金十数据中心经济数据的通用函数
func fetchJin10EconomicData(attrID int, productName string) (dataframe.DataFrame, error) {
	// 第一步：获取日期列表
	datesURL := "https://datacenter-api.jin10.com/reports/dates"
	datesParams := map[string]string{
		"category": "ec",
		"attr_id":  strconv.Itoa(attrID),
	}

	datesResp, err := utils.GetWithHeaders(datesURL, datesParams, jin10Headers)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("获取日期列表失败: %w", err)
	}

	dateArray := gjson.Get(datesResp.String(), "data").Array()
	if len(dateArray) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到日期数据")
	}

	// 每20个日期取一个作为分页点
	var datePointList []string
	for i, item := range dateArray {
		if i%20 == 0 {
			datePointList = append(datePointList, item.String())
		}
	}

	// 第二步：按日期分页获取数据
	listURL := "https://datacenter-api.jin10.com/reports/list_v2"
	var allRecords [][]string

	for _, date := range datePointList {
		listParams := map[string]string{
			"max_date": date,
			"category": "ec",
			"attr_id":  strconv.Itoa(attrID),
		}

		listResp, err := utils.GetWithHeaders(listURL, listParams, jin10Headers)
		if err != nil {
			continue
		}

		// 解析响应
		data := gjson.Get(listResp.String(), "data")
		keys := data.Get("keys").Array()
		values := data.Get("values").Array()

		if len(keys) == 0 || len(values) == 0 {
			continue
		}

		// 获取列名映射
		keyNames := make([]string, len(keys))
		for i, k := range keys {
			keyNames[i] = k.Get("name").String()
		}

		// 解析每行数据
		for _, row := range values {
			rowArray := row.Array()
			if len(rowArray) != len(keyNames) {
				continue
			}

			record := make(map[string]string)
			for i, v := range rowArray {
				record[keyNames[i]] = v.String()
			}

			// 提取所需字段
			dateStr := record["日期"]
			currentVal := record["今值"]
			forecastVal := record["预测值"]
			prevVal := record["前值"]

			// 只保留日期部分
			if len(dateStr) > 10 {
				dateStr = dateStr[:10]
			}

			allRecords = append(allRecords, []string{
				productName, dateStr, currentVal, forecastVal, prevVal,
			})
		}
	}

	if len(allRecords) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	// 去重
	seen := make(map[string]bool)
	var uniqueRecords [][]string
	for _, record := range allRecords {
		key := record[1] // 用日期作为去重键
		if !seen[key] {
			seen[key] = true
			uniqueRecords = append(uniqueRecords, record)
		}
	}

	// 按日期排序
	sort.Slice(uniqueRecords, func(i, j int) bool {
		return uniqueRecords[i][1] < uniqueRecords[j][1]
	})

	// 构建 DataFrame
	columns := []string{"商品", "日期", "今值", "预测值", "前值"}
	records := [][]string{columns}
	records = append(records, uniqueRecords...)

	df := dataframe.LoadRecords(records)
	return df, nil
}

// MacroEuroGDPYoY 金十数据中心-欧元区季度GDP年率报告
//
// 返回:
//   - dataframe.DataFrame: 欧元区季度GDP年率报告数据
//   - error: 错误信息
//
// 数据源: https://datacenter.jin10.com/reportType/dc_eurozone_gdp_yoy
func MacroEuroGDPYoY() (dataframe.DataFrame, error) {
	return fetchJin10EconomicData(84, "欧元区季度GDP年率")
}

// MacroEuroCPIMoM 金十数据中心-欧元区CPI月率报告
//
// 返回:
//   - dataframe.DataFrame: 欧元区CPI月率报告数据
//   - error: 错误信息
//
// 数据源: https://datacenter.jin10.com/reportType/dc_eurozone_cpi_mom
func MacroEuroCPIMoM() (dataframe.DataFrame, error) {
	return fetchJin10EconomicData(84, "欧元区CPI月率")
}

// MacroEuroCPIYoY 金十数据中心-欧元区CPI年率报告
//
// 返回:
//   - dataframe.DataFrame: 欧元区CPI年率报告数据
//   - error: 错误信息
//
// 数据源: https://datacenter.jin10.com/reportType/dc_eurozone_cpi_yoy
func MacroEuroCPIYoY() (dataframe.DataFrame, error) {
	return fetchJin10EconomicData(8, "欧元区CPI年率")
}

// MacroEuroPPIMoM 金十数据中心-欧元区PPI月率报告
//
// 返回:
//   - dataframe.DataFrame: 欧元区PPI月率报告数据
//   - error: 错误信息
//
// 数据源: https://datacenter.jin10.com/reportType/dc_eurozone_ppi_mom
func MacroEuroPPIMoM() (dataframe.DataFrame, error) {
	return fetchJin10EconomicData(36, "欧元区PPI月率")
}

// MacroEuroRetailSalesMoM 金十数据中心-欧元区零售销售月率报告
//
// 返回:
//   - dataframe.DataFrame: 欧元区零售销售月率报告数据
//   - error: 错误信息
//
// 数据源: https://datacenter.jin10.com/reportType/dc_eurozone_retail_sales_mom
func MacroEuroRetailSalesMoM() (dataframe.DataFrame, error) {
	return fetchJin10EconomicData(38, "欧元区零售销售月率")
}

// MacroEuroEmploymentChangeQoQ 金十数据中心-欧元区季调后就业人数季率报告
//
// 返回:
//   - dataframe.DataFrame: 欧元区季调后就业人数季率报告数据
//   - error: 错误信息
//
// 数据源: https://datacenter.jin10.com/reportType/dc_eurozone_employment_change_qoq
func MacroEuroEmploymentChangeQoQ() (dataframe.DataFrame, error) {
	return fetchJin10EconomicData(14, "欧元区季调后就业人数季率")
}

// MacroEuroUnemploymentRateMoM 金十数据中心-欧元区失业率报告
//
// 返回:
//   - dataframe.DataFrame: 欧元区失业率报告数据
//   - error: 错误信息
//
// 数据源: https://datacenter.jin10.com/reportType/dc_eurozone_unemployment_rate_mom
func MacroEuroUnemploymentRateMoM() (dataframe.DataFrame, error) {
	return fetchJin10EconomicData(46, "欧元区失业率")
}

// MacroEuroTradeBalance 金十数据中心-欧元区未季调贸易帐报告
//
// 返回:
//   - dataframe.DataFrame: 欧元区未季调贸易帐报告数据
//   - error: 错误信息
//
// 数据源: https://datacenter.jin10.com/reportType/dc_eurozone_trade_balance_mom
func MacroEuroTradeBalance() (dataframe.DataFrame, error) {
	return fetchJin10EconomicData(43, "欧元区未季调贸易帐")
}

// MacroEuroCurrentAccountMoM 金十数据中心-欧元区经常帐报告
//
// 返回:
//   - dataframe.DataFrame: 欧元区经常帐报告数据
//   - error: 错误信息
//
// 数据源: https://datacenter.jin10.com/reportType/dc_eurozone_current_account_mom
func MacroEuroCurrentAccountMoM() (dataframe.DataFrame, error) {
	return fetchJin10EconomicData(11, "欧元区经常帐")
}

// MacroEuroIndustrialProductionMoM 金十数据中心-欧元区工业产出月率报告
//
// 返回:
//   - dataframe.DataFrame: 欧元区工业产出月率报告数据
//   - error: 错误信息
//
// 数据源: https://datacenter.jin10.com/reportType/dc_eurozone_industrial_production_mom
func MacroEuroIndustrialProductionMoM() (dataframe.DataFrame, error) {
	return fetchJin10EconomicData(19, "欧元区工业产出月率")
}

// MacroEuroManufacturingPMI 金十数据中心-欧元区制造业PMI初值报告
//
// 返回:
//   - dataframe.DataFrame: 欧元区制造业PMI初值报告数据
//   - error: 错误信息
//
// 数据源: https://datacenter.jin10.com/reportType/dc_eurozone_manufacturing_pmi
func MacroEuroManufacturingPMI() (dataframe.DataFrame, error) {
	return fetchJin10EconomicData(30, "欧元区制造业PMI初值")
}

// MacroEuroServicesPMI 金十数据中心-欧元区服务业PMI终值报告
//
// 返回:
//   - dataframe.DataFrame: 欧元区服务业PMI终值报告数据
//   - error: 错误信息
//
// 数据源: https://datacenter.jin10.com/reportType/dc_eurozone_services_pmi
func MacroEuroServicesPMI() (dataframe.DataFrame, error) {
	return fetchJin10EconomicData(41, "欧元区服务业PMI终值")
}

// MacroEuroZEWEconomicSentiment 金十数据中心-欧元区ZEW经济景气指数报告
//
// 返回:
//   - dataframe.DataFrame: 欧元区ZEW经济景气指数报告数据
//   - error: 错误信息
//
// 数据源: https://datacenter.jin10.com/reportType/dc_eurozone_zew_economic_sentiment
func MacroEuroZEWEconomicSentiment() (dataframe.DataFrame, error) {
	return fetchJin10EconomicData(48, "欧元区ZEW经济景气指数")
}

// MacroEuroSentixInvestorConfidence 金十数据中心-欧元区Sentix投资者信心指数报告
//
// 返回:
//   - dataframe.DataFrame: 欧元区Sentix投资者信心指数报告数据
//   - error: 错误信息
//
// 数据源: https://datacenter.jin10.com/reportType/dc_eurozone_sentix_investor_confidence
func MacroEuroSentixInvestorConfidence() (dataframe.DataFrame, error) {
	return fetchJin10EconomicData(40, "欧元区Sentix投资者信心指数")
}

// MacroEuroLMEHolding 金十数据中心-伦敦金属交易所(LME)-持仓报告
//
// 返回:
//   - dataframe.DataFrame: LME持仓报告数据
//   - error: 错误信息
//
// 数据源: https://datacenter.jin10.com/reportType/dc_lme_traders_report
func MacroEuroLMEHolding() (dataframe.DataFrame, error) {
	url := "https://cdn.jin10.com/data_center/reports/lme_position.json"
	params := map[string]string{
		"_": strconv.FormatInt(time.Now().UnixMilli(), 10),
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	// 解析 JSON - values 结构为 map[日期]map[品种][]float64
	var result struct {
		Products []string                        `json:"products"`
		Keys     []map[string]string             `json:"keys"`
		Values   map[string]map[string][]float64 `json:"values"`
	}
	if err := json.Unmarshal([]byte(resp.String()), &result); err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("JSON解析失败: %w", err)
	}

	if len(result.Keys) == 0 || len(result.Values) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("数据格式不正确")
	}

	// 获取列名
	keyNames := make([]string, len(result.Keys))
	for i, k := range result.Keys {
		keyNames[i] = k["name"]
	}

	// 构建列名列表 - 使用 API 返回的品种顺序
	commodities := result.Products
	var columns []string
	columns = append(columns, "日期")
	for _, commodity := range commodities {
		for _, keyName := range keyNames {
			columns = append(columns, commodity+"-"+keyName)
		}
	}

	// 构建记录
	var records [][]string
	records = append(records, columns)

	// 获取所有日期并排序
	var dates []string
	for date := range result.Values {
		dates = append(dates, date)
	}
	sort.Strings(dates)

	// 移除最后一个日期（可能是未完成数据）
	if len(dates) > 1 {
		dates = dates[:len(dates)-1]
	}

	for _, date := range dates {
		dateData := result.Values[date]
		record := []string{date}

		for _, commodity := range commodities {
			nums, ok := dateData[commodity]
			if !ok || len(nums) < len(keyNames) {
				// 如果没有该品种数据，用 0 填充
				for range keyNames {
					record = append(record, "0")
				}
				continue
			}
			for _, num := range nums {
				record = append(record, strconv.FormatFloat(num, 'f', -1, 64))
			}
		}

		records = append(records, record)
	}

	df := dataframe.LoadRecords(records)
	return df, nil
}

// MacroEuroLMEStock 金十数据中心-伦敦金属交易所(LME)-库存报告
//
// 返回:
//   - dataframe.DataFrame: LME库存报告数据
//   - error: 错误信息
//
// 数据源: https://datacenter.jin10.com/reportType/dc_lme_report
func MacroEuroLMEStock() (dataframe.DataFrame, error) {
	url := "https://cdn.jin10.com/data_center/reports/lme_stock.json"
	params := map[string]string{
		"_": strconv.FormatInt(time.Now().UnixMilli(), 10),
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	// 解析 JSON - values 结构为 map[日期]map[品种][]float64
	var result struct {
		Products []string                        `json:"products"`
		Keys     []map[string]string             `json:"keys"`
		Values   map[string]map[string][]float64 `json:"values"`
	}
	if err := json.Unmarshal([]byte(resp.String()), &result); err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("JSON解析失败: %w", err)
	}

	if len(result.Keys) == 0 || len(result.Values) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("数据格式不正确")
	}

	// 获取列名
	keyNames := make([]string, len(result.Keys))
	for i, k := range result.Keys {
		keyNames[i] = k["name"]
	}

	// 构建列名列表 - 使用 API 返回的品种顺序
	commodities := result.Products
	var columns []string
	columns = append(columns, "日期")
	for _, commodity := range commodities {
		for _, keyName := range keyNames {
			columns = append(columns, commodity+"-"+keyName)
		}
	}

	// 构建记录
	var records [][]string
	records = append(records, columns)

	// 获取所有日期并排序
	var dates []string
	for date := range result.Values {
		dates = append(dates, date)
	}
	sort.Strings(dates)

	for _, date := range dates {
		dateData := result.Values[date]
		record := []string{date}

		for _, commodity := range commodities {
			nums, ok := dateData[commodity]
			if !ok || len(nums) < len(keyNames) {
				// 如果没有该品种数据，用 0 填充
				for range keyNames {
					record = append(record, "0")
				}
				continue
			}
			for _, num := range nums {
				record = append(record, strconv.FormatFloat(num, 'f', -1, 64))
			}
		}

		records = append(records, record)
	}

	df := dataframe.LoadRecords(records)
	return df, nil
}
