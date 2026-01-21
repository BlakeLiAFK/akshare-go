package stock_fundamental

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
)

// StockFinancialThsItem 同花顺财务指标项
type StockFinancialThsItem struct {
	ReportDate   string                 `json:"report_date"`   // 报告期
	ReportName   string                 `json:"report_name"`   // 报告名称
	ReportPeriod string                 `json:"report_period"` // 报告周期
	QuarterName  string                 `json:"quarter_name"`  // 季度名称
	MetricName   string                 `json:"metric_name"`   // 指标名称
	Data         map[string]interface{} `json:"data"`          // 指标数据
}

// StockManagementChangeThsItem 同花顺高管持股变动项
type StockManagementChangeThsItem struct {
	Name       string `json:"name"`        // 姓名
	ChangeDate string `json:"change_date"` // 变动日期
	ChangeNum  string `json:"change_num"`  // 变动数量
	AvgPrice   string `json:"avg_price"`   // 交易均价
	RemainNum  string `json:"remain_num"`  // 剩余股数
}

// StockShareholderChangeThsItem 同花顺股东持股变动项
type StockShareholderChangeThsItem struct {
	HolderName   string `json:"holder_name"`   // 股东名称
	AnnounceDate string `json:"announce_date"` // 公告日期
	ChangeNum    string `json:"change_num"`    // 变动数量
	AvgPrice     string `json:"avg_price"`     // 交易均价
	RemainNum    string `json:"remain_num"`    // 剩余股份总数
}

// getMarketCode 获取股票所属市场代码
func getMarketCode(stockCode string) int {
	stockCode = strings.TrimSpace(stockCode)
	if len(stockCode) < 6 {
		return 0
	}

	// 深交所股票
	if strings.HasPrefix(stockCode, "000") || strings.HasPrefix(stockCode, "001") ||
		strings.HasPrefix(stockCode, "002") || strings.HasPrefix(stockCode, "003") ||
		strings.HasPrefix(stockCode, "300") {
		return 33
	}

	// 上交所股票
	if strings.HasPrefix(stockCode, "600") || strings.HasPrefix(stockCode, "601") ||
		strings.HasPrefix(stockCode, "603") || strings.HasPrefix(stockCode, "605") ||
		strings.HasPrefix(stockCode, "688") {
		return 17
	}

	// 北交所股票
	if strings.HasPrefix(stockCode, "920") {
		return 151
	}

	return 0
}

// StockFinancialAbstractThs 同花顺-财务指标-主要指标（旧版）
// symbol: 股票代码
// indicator: 指标类型，可选 "按报告期", "按年度", "按单季度"
func StockFinancialAbstractThs(symbol, indicator string) ([]map[string]interface{}, error) {
	url := fmt.Sprintf("https://basic.10jqka.com.cn/new/%s/finance.html", symbol)
	headers := map[string]string{
		"User-Agent": utils.DefaultUserAgent,
	}

	resp, err := utils.GetWithHeaders(url, nil, headers)
	if err != nil {
		return nil, fmt.Errorf("获取财务数据失败: %w", err)
	}

	// 解析 HTML 提取 JSON 数据
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	dataText := doc.Find("p#main").Text()
	if dataText == "" {
		return nil, fmt.Errorf("未找到财务数据")
	}

	var dataJson map[string]interface{}
	if err := json.Unmarshal([]byte(dataText), &dataJson); err != nil {
		return nil, fmt.Errorf("解析JSON失败: %w", err)
	}

	// 根据指标类型选择数据
	var dataKey string
	switch indicator {
	case "按报告期":
		dataKey = "report"
	case "按单季度":
		dataKey = "simple"
	default:
		dataKey = "year"
	}

	data, ok := dataJson[dataKey].([]interface{})
	if !ok || len(data) == 0 {
		return nil, nil
	}

	// 提取标题和数据
	titles, ok := dataJson["title"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("未找到标题数据")
	}

	// 构建结果
	var result []map[string]interface{}
	headers2, ok := data[0].([]interface{})
	if !ok || len(headers2) == 0 {
		return nil, nil
	}

	// 遍历每一列（每个报告期）
	for colIdx := 0; colIdx < len(headers2); colIdx++ {
		record := make(map[string]interface{})
		record["报告期"] = headers2[colIdx]

		// 遍历每一行（每个指标）
		for rowIdx := 1; rowIdx < len(data) && rowIdx < len(titles); rowIdx++ {
			row, ok := data[rowIdx].([]interface{})
			if !ok || colIdx >= len(row) {
				continue
			}

			var titleName string
			if titleArr, ok := titles[rowIdx].([]interface{}); ok && len(titleArr) > 0 {
				titleName = fmt.Sprintf("%v", titleArr[0])
			} else {
				titleName = fmt.Sprintf("%v", titles[rowIdx])
			}

			record[titleName] = row[colIdx]
		}

		result = append(result, record)
	}

	return result, nil
}

// StockFinancialAbstractNewThs 同花顺-财务指标-重要指标（新版API）
// symbol: 股票代码
// indicator: 指标类型，可选 "按报告期", "一季度", "二季度", "三季度", "四季度", "按年度"
func StockFinancialAbstractNewThs(symbol, indicator string) ([]StockFinancialThsItem, error) {
	url := "https://basic.10jqka.com.cn/basicapi/finance/index/v1/app_data/"

	var period string
	switch indicator {
	case "按报告期":
		period = "0"
	case "一季度":
		period = "1"
	case "二季度":
		period = "2"
	case "三季度":
		period = "3"
	case "四季度":
		period = "4"
	default:
		period = "4"
	}

	params := map[string]string{
		"code":   symbol,
		"id":     "client_stock_importance",
		"market": fmt.Sprintf("%d", getMarketCode(symbol)),
		"type":   "stock",
		"page":   "1",
		"size":   "50",
		"period": period,
	}

	headers := map[string]string{
		"User-Agent": utils.DefaultUserAgent,
	}

	resp, err := utils.GetWithHeaders(url, params, headers)
	if err != nil {
		return nil, fmt.Errorf("获取财务数据失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	financialData := result.Get("data.data").Array()

	var items []StockFinancialThsItem
	for _, report := range financialData {
		reportDate := report.Get("date").String()
		reportName := report.Get("report_name").String()
		reportPeriod := report.Get("report").String()
		quarterName := report.Get("quarter_name").String()

		indexList := report.Get("index_list").Map()
		for metricName, metricValues := range indexList {
			item := StockFinancialThsItem{
				ReportDate:   reportDate,
				ReportName:   reportName,
				ReportPeriod: reportPeriod,
				QuarterName:  quarterName,
				MetricName:   metricName,
				Data:         make(map[string]interface{}),
			}

			if metricValues.IsObject() {
				for k, v := range metricValues.Map() {
					item.Data[k] = v.Value()
				}
			} else {
				item.Data["value"] = metricValues.Value()
			}

			items = append(items, item)
		}
	}

	return items, nil
}

// StockFinancialDebtNewThs 同花顺-财务指标-资产负债表（新版API）
func StockFinancialDebtNewThs(symbol, indicator string) ([]StockFinancialThsItem, error) {
	return stockFinancialNewThs(symbol, indicator, "client_stock_debt")
}

// StockFinancialBenefitNewThs 同花顺-财务指标-利润表（新版API）
func StockFinancialBenefitNewThs(symbol, indicator string) ([]StockFinancialThsItem, error) {
	return stockFinancialNewThs(symbol, indicator, "client_stock_benefit")
}

// StockFinancialCashNewThs 同花顺-财务指标-现金流量表（新版API）
func StockFinancialCashNewThs(symbol, indicator string) ([]StockFinancialThsItem, error) {
	return stockFinancialNewThs(symbol, indicator, "client_stock_cash")
}

// stockFinancialNewThs 同花顺财务数据通用获取函数
func stockFinancialNewThs(symbol, indicator, dataId string) ([]StockFinancialThsItem, error) {
	url := "https://basic.10jqka.com.cn/basicapi/finance/index/v1/app_data/"

	var period string
	switch indicator {
	case "按报告期":
		period = "0"
	case "一季度":
		period = "1"
	case "二季度":
		period = "2"
	case "三季度":
		period = "3"
	case "四季度":
		period = "4"
	default:
		period = "4"
	}

	params := map[string]string{
		"code":   symbol,
		"id":     dataId,
		"market": fmt.Sprintf("%d", getMarketCode(symbol)),
		"type":   "stock",
		"page":   "1",
		"size":   "50",
		"period": period,
	}

	headers := map[string]string{
		"User-Agent": utils.DefaultUserAgent,
	}

	resp, err := utils.GetWithHeaders(url, params, headers)
	if err != nil {
		return nil, fmt.Errorf("获取财务数据失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	financialData := result.Get("data.data").Array()

	var items []StockFinancialThsItem
	for _, report := range financialData {
		reportDate := report.Get("date").String()
		reportName := report.Get("report_name").String()
		reportPeriod := report.Get("report").String()
		quarterName := report.Get("quarter_name").String()

		indexList := report.Get("index_list").Map()
		for metricName, metricValues := range indexList {
			item := StockFinancialThsItem{
				ReportDate:   reportDate,
				ReportName:   reportName,
				ReportPeriod: reportPeriod,
				QuarterName:  quarterName,
				MetricName:   metricName,
				Data:         make(map[string]interface{}),
			}

			if metricValues.IsObject() {
				for k, v := range metricValues.Map() {
					item.Data[k] = v.Value()
				}
			} else {
				item.Data["value"] = metricValues.Value()
			}

			items = append(items, item)
		}
	}

	return items, nil
}

// StockManagementChangeThs 同花顺-公司大事-高管持股变动
func StockManagementChangeThs(symbol string) ([]StockManagementChangeThsItem, error) {
	url := fmt.Sprintf("https://basic.10jqka.com.cn/new/%s/event.html", symbol)
	headers := map[string]string{
		"User-Agent": utils.DefaultUserAgent,
	}

	resp, err := utils.GetWithHeaders(url, nil, headers)
	if err != nil {
		return nil, fmt.Errorf("获取高管持股变动失败: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	var items []StockManagementChangeThsItem

	// 查找高管持股变动表格
	table := doc.Find("table.data_table_1.m_table.m_hl")
	if table.Length() == 0 {
		return items, nil
	}

	// 解析表格数据
	table.Find("tbody tr").Each(func(i int, tr *goquery.Selection) {
		tds := tr.Find("td")
		if tds.Length() >= 5 {
			item := StockManagementChangeThsItem{
				Name:       strings.TrimSpace(tds.Eq(0).Text()),
				ChangeDate: strings.TrimSpace(tds.Eq(1).Text()),
				ChangeNum:  strings.TrimSpace(tds.Eq(2).Text()),
				AvgPrice:   strings.TrimSpace(tds.Eq(3).Text()),
				RemainNum:  strings.TrimSpace(tds.Eq(4).Text()),
			}
			items = append(items, item)
		}
	})

	return items, nil
}

// StockShareholderChangeThs 同花顺-公司大事-股东持股变动
func StockShareholderChangeThs(symbol string) ([]StockShareholderChangeThsItem, error) {
	url := fmt.Sprintf("https://basic.10jqka.com.cn/new/%s/event.html", symbol)
	headers := map[string]string{
		"User-Agent": utils.DefaultUserAgent,
	}

	resp, err := utils.GetWithHeaders(url, nil, headers)
	if err != nil {
		return nil, fmt.Errorf("获取股东持股变动失败: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	var items []StockShareholderChangeThsItem

	// 查找股东持股变动表格
	table := doc.Find("table.m_table.data_table_1.m_hl")
	if table.Length() == 0 {
		return items, nil
	}

	// 解析表格数据
	table.Find("tbody tr").Each(func(i int, tr *goquery.Selection) {
		tds := tr.Find("td")
		if tds.Length() >= 5 {
			item := StockShareholderChangeThsItem{
				HolderName:   strings.TrimSpace(tds.Eq(0).Text()),
				AnnounceDate: strings.TrimSpace(tds.Eq(1).Text()),
				ChangeNum:    strings.TrimSpace(tds.Eq(2).Text()),
				AvgPrice:     strings.TrimSpace(tds.Eq(3).Text()),
				RemainNum:    strings.TrimSpace(tds.Eq(4).Text()),
			}
			items = append(items, item)
		}
	})

	return items, nil
}

// StockFinancialDebtThs 同花顺-财务指标-资产负债表（旧版）
func StockFinancialDebtThs(symbol, indicator string) ([]map[string]interface{}, error) {
	return stockFinancialOldThs(symbol, indicator, "debt")
}

// StockFinancialBenefitThs 同花顺-财务指标-利润表（旧版）
func StockFinancialBenefitThs(symbol, indicator string) ([]map[string]interface{}, error) {
	return stockFinancialOldThs(symbol, indicator, "benefit")
}

// StockFinancialCashThs 同花顺-财务指标-现金流量表（旧版）
func StockFinancialCashThs(symbol, indicator string) ([]map[string]interface{}, error) {
	return stockFinancialOldThs(symbol, indicator, "cash")
}

// stockFinancialOldThs 同花顺财务数据通用获取函数（旧版API）
func stockFinancialOldThs(symbol, indicator, dataType string) ([]map[string]interface{}, error) {
	url := fmt.Sprintf("https://basic.10jqka.com.cn/api/stock/finance/%s_%s.json", symbol, dataType)
	headers := map[string]string{
		"User-Agent": utils.DefaultUserAgent,
	}

	resp, err := utils.GetWithHeaders(url, nil, headers)
	if err != nil {
		return nil, fmt.Errorf("获取财务数据失败: %w", err)
	}

	// 解析嵌套的 JSON
	flashData := gjson.Get(resp.String(), "flashData").String()
	if flashData == "" {
		return nil, fmt.Errorf("未找到财务数据")
	}

	var dataJson map[string]interface{}
	if err := json.Unmarshal([]byte(flashData), &dataJson); err != nil {
		return nil, fmt.Errorf("解析JSON失败: %w", err)
	}

	// 根据指标类型选择数据
	var dataKey string
	switch indicator {
	case "按报告期":
		dataKey = "report"
	case "按单季度":
		dataKey = "simple"
	default:
		dataKey = "year"
	}

	data, ok := dataJson[dataKey].([]interface{})
	if !ok || len(data) == 0 {
		return nil, nil
	}

	titles, ok := dataJson["title"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("未找到标题数据")
	}

	var result []map[string]interface{}
	headers2, ok := data[0].([]interface{})
	if !ok || len(headers2) == 0 {
		return nil, nil
	}

	for colIdx := 0; colIdx < len(headers2); colIdx++ {
		record := make(map[string]interface{})
		record["报告期"] = headers2[colIdx]

		for rowIdx := 1; rowIdx < len(data) && rowIdx < len(titles); rowIdx++ {
			row, ok := data[rowIdx].([]interface{})
			if !ok || colIdx >= len(row) {
				continue
			}

			var titleName string
			if titleArr, ok := titles[rowIdx].([]interface{}); ok && len(titleArr) > 0 {
				titleName = fmt.Sprintf("%v", titleArr[0])
			} else {
				titleName = fmt.Sprintf("%v", titles[rowIdx])
			}

			record[titleName] = row[colIdx]
		}

		result = append(result, record)
	}

	return result, nil
}

// 预编译正则表达式，避免未使用警告
var _ = regexp.MustCompile(`\s+`)
