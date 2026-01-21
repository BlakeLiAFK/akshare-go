package stock_fundamental

import (
	"fmt"
	"sort"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// StockFinancialUsReportEmItem 美股财务报表项
type StockFinancialUsReportEmItem struct {
	SecuCode     string  `json:"secu_code"`     // 证券代码
	SecurityCode string  `json:"security_code"` // 股票代码
	SecurityName string  `json:"security_name"` // 股票简称
	ReportDate   string  `json:"report_date"`   // 报告日期
	ReportType   string  `json:"report_type"`   // 报告类型
	Report       string  `json:"report"`        // 报告期
	StdItemCode  string  `json:"std_item_code"` // 标准项目代码
	ItemName     string  `json:"item_name"`     // 项目名称
	Amount       float64 `json:"amount"`        // 金额
}

// StockFinancialUsIndicatorEmItem 美股财务分析主要指标项
type StockFinancialUsIndicatorEmItem struct {
	SecuCode        string  `json:"secu_code"`         // 证券代码
	SecurityCode    string  `json:"security_code"`     // 股票代码
	SecurityName    string  `json:"security_name"`     // 股票简称
	ReportDate      string  `json:"report_date"`       // 报告日期
	DateTypeCode    string  `json:"date_type_code"`    // 日期类型代码
	TotalRevenue    float64 `json:"total_revenue"`     // 营业总收入
	TotalRevenueYOY float64 `json:"total_revenue_yoy"` // 营业总收入同比
	NetProfit       float64 `json:"net_profit"`        // 归母净利润
	NetProfitYOY    float64 `json:"net_profit_yoy"`    // 归母净利润同比
	BasicEPS        float64 `json:"basic_eps"`         // 基本每股收益
	BasicEPSYOY     float64 `json:"basic_eps_yoy"`     // 基本每股收益同比
	DilutedEPS      float64 `json:"diluted_eps"`       // 稀释每股收益
	PayoutRatio     float64 `json:"payout_ratio"`      // 派息比率
	CapitalRatio    float64 `json:"capital_ratio"`     // 资本比率
	ROE             float64 `json:"roe"`               // 净资产收益率
	ROEYOY          float64 `json:"roe_yoy"`           // 净资产收益率同比
	ROA             float64 `json:"roa"`               // 总资产净利率
	ROAYOY          float64 `json:"roa_yoy"`           // 总资产净利率同比
	DebtRatio       float64 `json:"debt_ratio"`        // 资产负债率
	DebtRatioYOY    float64 `json:"debt_ratio_yoy"`    // 资产负债率同比
	EquityRatio     float64 `json:"equity_ratio"`      // 权益比率
}

// stockFinancialUsReportQueryMarketEm 查询美股市场代码
func stockFinancialUsReportQueryMarketEm(symbol string) (string, error) {
	baseURL := "https://datacenter.eastmoney.com/securities/api/data/v1/get"

	params := map[string]string{
		"reportName":   "RPT_USF10_INFO_ORGPROFILE",
		"columns":      "SECUCODE,SECURITY_CODE,ORG_CODE,SECURITY_INNER_CODE,ORG_NAME,ORG_EN_ABBR,BELONG_INDUSTRY,FOUND_DATE,CHAIRMAN,REG_PLACE,ADDRESS,EMP_NUM,ORG_TEL,ORG_FAX,ORG_EMAIL,ORG_WEB,ORG_PROFILE",
		"quoteColumns": "",
		"filter":       fmt.Sprintf(`(SECURITY_CODE="%s")`, symbol),
		"pageNumber":   "1",
		"pageSize":     "200",
		"sortTypes":    "",
		"sortColumns":  "",
		"source":       "SECURITIES",
		"client":       "PC",
	}

	resp, err := utils.Get(baseURL, params)
	if err != nil {
		return "", fmt.Errorf("查询市场代码失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	stockCode := result.Get("result.data.0.SECUCODE").String()
	if stockCode == "" {
		return "", fmt.Errorf("未找到股票代码: %s", symbol)
	}

	return stockCode, nil
}

// stockFinancialUsReportEm 获取美股报告期列表
func stockFinancialUsReportEm(stock, symbol, indicator string) (string, error) {
	baseURL := "https://datacenter.eastmoney.com/securities/api/data/v1/get"

	stockCode, err := stockFinancialUsReportQueryMarketEm(stock)
	if err != nil {
		return "", err
	}

	var reportName string
	switch symbol {
	case "资产负债表":
		reportName = "RPT_USF10_FN_BALANCE"
	case "综合损益表":
		reportName = "RPT_USF10_FN_INCOME"
	case "现金流量表":
		reportName = "RPT_USSK_FN_CASHFLOW"
	default:
		return "", fmt.Errorf("不支持的报表类型: %s", symbol)
	}

	params := map[string]string{
		"reportName":   reportName,
		"columns":      "SECUCODE,SECURITY_CODE,SECURITY_NAME_ABBR,REPORT,REPORT_DATE,FISCAL_YEAR,CURRENCY,ACCOUNT_STANDARD,REPORT_TYPE,DATE_TYPE_CODE",
		"quoteColumns": "",
		"filter":       fmt.Sprintf(`(SECUCODE="%s")`, stockCode),
		"pageNumber":   "",
		"pageSize":     "",
		"sortTypes":    "-1",
		"sortColumns":  "REPORT_DATE",
		"source":       "SECURITIES",
		"client":       "PC",
	}

	resp, err := utils.Get(baseURL, params)
	if err != nil {
		return "", fmt.Errorf("获取报告期列表失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	dataArr := result.Get("result.data").Array()

	// 收集所有报告期
	reportSet := make(map[string]bool)
	for _, item := range dataArr {
		report := strings.TrimSpace(item.Get("REPORT").String())
		reportSet[report] = true
	}

	// 根据指标类型筛选
	var filteredReports []string
	for report := range reportSet {
		switch indicator {
		case "年报":
			if strings.Contains(report, "FY") {
				filteredReports = append(filteredReports, report)
			}
		case "单季报":
			if strings.Contains(report, "Q1") || strings.Contains(report, "Q2") ||
				strings.Contains(report, "Q3") || strings.Contains(report, "Q4") {
				filteredReports = append(filteredReports, report)
			}
		case "累计季报":
			if strings.Contains(report, "Q6") || strings.Contains(report, "Q9") {
				filteredReports = append(filteredReports, report)
			}
		}
	}

	// 排序
	sort.Slice(filteredReports, func(i, j int) bool {
		partsI := strings.Split(filteredReports[i], "/")
		partsJ := strings.Split(filteredReports[j], "/")
		if len(partsI) > 0 && len(partsJ) > 0 {
			return partsI[0] > partsJ[0]
		}
		return filteredReports[i] > filteredReports[j]
	})

	// 构建过滤字符串
	if len(filteredReports) == 0 {
		return "", nil
	}

	quotedReports := make([]string, len(filteredReports))
	for i, r := range filteredReports {
		quotedReports[i] = `"` + r + `"`
	}
	return "(" + strings.Join(quotedReports, ",") + ")", nil
}

// StockFinancialUsReportEm 东方财富-美股-财务分析-三大报表
// stock: 股票代码，如 "TSLA"
// symbol: 报表类型，可选 "资产负债表", "综合损益表", "现金流量表"
// indicator: 指标类型，可选 "年报", "单季报", "累计季报"
func StockFinancialUsReportEm(stock, symbol, indicator string) ([]StockFinancialUsReportEmItem, error) {
	baseURL := "https://datacenter.eastmoney.com/securities/api/data/v1/get"

	dateStr, err := stockFinancialUsReportEm(stock, symbol, indicator)
	if err != nil {
		return nil, err
	}
	if dateStr == "" {
		return nil, nil
	}

	stockCode, err := stockFinancialUsReportQueryMarketEm(stock)
	if err != nil {
		return nil, err
	}

	var reportName string
	switch symbol {
	case "资产负债表":
		reportName = "RPT_USF10_FN_BALANCE"
	case "综合损益表":
		reportName = "RPT_USF10_FN_INCOME"
	case "现金流量表":
		reportName = "RPT_USSK_FN_CASHFLOW"
	default:
		return nil, fmt.Errorf("不支持的报表类型: %s", symbol)
	}

	params := map[string]string{
		"reportName":   reportName,
		"columns":      "SECUCODE,SECURITY_CODE,SECURITY_NAME_ABBR,REPORT_DATE,REPORT_TYPE,REPORT,STD_ITEM_CODE,AMOUNT,ITEM_NAME",
		"quoteColumns": "",
		"filter":       fmt.Sprintf(`(SECUCODE="%s")(REPORT in %s)`, stockCode, dateStr),
		"pageNumber":   "",
		"pageSize":     "",
		"sortTypes":    "1,-1",
		"sortColumns":  "STD_ITEM_CODE,REPORT_DATE",
		"source":       "SECURITIES",
		"client":       "PC",
	}

	resp, err := utils.Get(baseURL, params)
	if err != nil {
		return nil, fmt.Errorf("获取报表数据失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	dataArr := result.Get("result.data").Array()

	var items []StockFinancialUsReportEmItem
	for _, item := range dataArr {
		items = append(items, StockFinancialUsReportEmItem{
			SecuCode:     item.Get("SECUCODE").String(),
			SecurityCode: item.Get("SECURITY_CODE").String(),
			SecurityName: item.Get("SECURITY_NAME_ABBR").String(),
			ReportDate:   item.Get("REPORT_DATE").String(),
			ReportType:   item.Get("REPORT_TYPE").String(),
			Report:       item.Get("REPORT").String(),
			StdItemCode:  item.Get("STD_ITEM_CODE").String(),
			ItemName:     item.Get("ITEM_NAME").String(),
			Amount:       item.Get("AMOUNT").Float(),
		})
	}

	return items, nil
}

// StockFinancialUsAnalysisIndicatorEm 东方财富-美股-财务分析-主要指标
// symbol: 股票代码，如 "TSLA"
// indicator: 指标类型，可选 "年报", "单季报", "累计季报"
func StockFinancialUsAnalysisIndicatorEm(symbol, indicator string) ([]map[string]interface{}, error) {
	baseURL := "https://datacenter.eastmoney.com/securities/api/data/v1/get"

	stockCode, err := stockFinancialUsReportQueryMarketEm(symbol)
	if err != nil {
		return nil, err
	}

	params := map[string]string{
		"reportName":   "RPT_USF10_FN_GMAININDICATOR",
		"columns":      "USF10_FN_GMAININDICATOR",
		"quoteColumns": "",
		"pageNumber":   "",
		"pageSize":     "",
		"sortTypes":    "-1",
		"sortColumns":  "REPORT_DATE",
		"source":       "SECURITIES",
		"client":       "PC",
	}

	// 特殊处理带下划线的代码（如 BRK_A）
	if strings.Contains(stockCode, "_") {
		params["reportName"] = "RPT_USF10_FN_IMAININDICATOR"
		params["columns"] = "ORG_CODE,SECURITY_CODE,SECUCODE,SECURITY_NAME_ABBR,SECURITY_INNER_CODE,STD_REPORT_DATE,REPORT_DATE,DATE_TYPE,DATE_TYPE_CODE,REPORT_TYPE,REPORT_DATA_TYPE,FISCAL_YEAR,START_DATE,NOTICE_DATE,ACCOUNT_STANDARD,ACCOUNT_STANDARD_NAME,CURRENCY,CURRENCY_NAME,ORGTYPE,TOTAL_INCOME,TOTAL_INCOME_YOY,PREMIUM_INCOME,PREMIUM_INCOME_YOY,PARENT_HOLDER_NETPROFIT,PARENT_HOLDER_NETPROFIT_YOY,BASIC_EPS_CS,BASIC_EPS_CS_YOY,DILUTED_EPS_CS,PAYOUT_RATIO,CAPITIAL_RATIO,ROE,ROE_YOY,ROA,ROA_YOY,DEBT_RATIO,DEBT_RATIO_YOY,EQUITY_RATIO"
	}

	switch indicator {
	case "年报":
		params["filter"] = fmt.Sprintf(`(SECUCODE="%s")(DATE_TYPE_CODE="001")`, stockCode)
	case "单季报":
		params["filter"] = fmt.Sprintf(`(SECUCODE="%s")(DATE_TYPE_CODE in ("003","006","007","008"))`, stockCode)
	case "累计季报":
		params["filter"] = fmt.Sprintf(`(SECUCODE="%s")(DATE_TYPE_CODE in ("002","004"))`, stockCode)
	default:
		return nil, fmt.Errorf("不支持的指标类型: %s", indicator)
	}

	resp, err := utils.Get(baseURL, params)
	if err != nil {
		return nil, fmt.Errorf("获取主要指标失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	dataArr := result.Get("result.data").Array()

	var items []map[string]interface{}
	for _, item := range dataArr {
		m := item.Value().(map[string]interface{})
		items = append(items, m)
	}

	return items, nil
}
