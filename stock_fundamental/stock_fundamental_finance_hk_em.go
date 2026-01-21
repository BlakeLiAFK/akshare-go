package stock_fundamental

import (
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// StockFinancialHkReportEmItem 港股财务报表项
type StockFinancialHkReportEmItem struct {
	SecuCode     string  `json:"secu_code"`     // 证券代码
	SecurityCode string  `json:"security_code"` // 股票代码
	SecurityName string  `json:"security_name"` // 股票简称
	OrgCode      string  `json:"org_code"`      // 机构代码
	ReportDate   string  `json:"report_date"`   // 报告日期
	DateTypeCode string  `json:"date_type"`     // 日期类型代码
	FiscalYear   string  `json:"fiscal_year"`   // 财政年度
	StdItemCode  string  `json:"std_item_code"` // 标准项目代码
	StdItemName  string  `json:"std_item_name"` // 标准项目名称
	Amount       float64 `json:"amount"`        // 金额
}

// StockFinancialHkIndicatorEmItem 港股财务分析主要指标项
type StockFinancialHkIndicatorEmItem struct {
	SecuCode             string  `json:"secu_code"`               // 证券代码
	SecurityCode         string  `json:"security_code"`           // 股票代码
	SecurityName         string  `json:"security_name"`           // 股票简称
	ReportDate           string  `json:"report_date"`             // 报告日期
	DateTypeCode         string  `json:"date_type"`               // 日期类型代码
	ROE                  float64 `json:"roe"`                     // 净资产收益率
	ROA                  float64 `json:"roa"`                     // 总资产净利率
	GrossProfitMargin    float64 `json:"gross_profit_margin"`     // 销售毛利率
	NetProfitMargin      float64 `json:"net_profit_margin"`       // 销售净利率
	EPS                  float64 `json:"eps"`                     // 每股收益
	BPS                  float64 `json:"bps"`                     // 每股净资产
	TotalRevenue         float64 `json:"total_revenue"`           // 营业总收入
	TotalRevenueYOY      float64 `json:"total_revenue_yoy"`       // 营业总收入同比
	NetProfit            float64 `json:"net_profit"`              // 净利润
	NetProfitYOY         float64 `json:"net_profit_yoy"`          // 净利润同比
	DeductedNetProfit    float64 `json:"deducted_net_profit"`     // 扣非净利润
	DeductedNetProfitYOY float64 `json:"deducted_net_profit_yoy"` // 扣非净利润同比
	OperatingCashFlow    float64 `json:"operating_cash_flow"`     // 经营活动产生的现金流量净额
	OperatingCashFlowYOY float64 `json:"operating_cash_flow_yoy"` // 经营活动产生的现金流量净额同比
	CurrentRatio         float64 `json:"current_ratio"`           // 流动比率
	QuickRatio           float64 `json:"quick_ratio"`             // 速动比率
	DebtToAssetRatio     float64 `json:"debt_to_asset_ratio"`     // 资产负债率
	EquityMultiplier     float64 `json:"equity_multiplier"`       // 权益乘数
	TotalAssetsTurnover  float64 `json:"total_assets_turnover"`   // 总资产周转率
	InventoryTurnover    float64 `json:"inventory_turnover"`      // 存货周转率
	ReceivablesTurnover  float64 `json:"receivables_turnover"`    // 应收账款周转率
	OperatingCycle       float64 `json:"operating_cycle"`         // 营业周期
	CashConversionCycle  float64 `json:"cash_conversion_cycle"`   // 现金转换周期
}

// StockFinancialHkReportEm 东方财富-港股-财务报表-三大报表
// stock: 股票代码，如 "00700"
// symbol: 报表类型，可选 "资产负债表", "利润表", "现金流量表"
// indicator: 指标类型，可选 "年度", "报告期"
func StockFinancialHkReportEm(stock, symbol, indicator string) ([]StockFinancialHkReportEmItem, error) {
	baseURL := "https://datacenter.eastmoney.com/securities/api/data/v1/get"

	// 先获取报告期列表
	params := map[string]string{
		"reportName":   "RPT_CUSTOM_HKSK_APPFN_CASHFLOW_SUMMARY",
		"columns":      "SECUCODE,SECURITY_CODE,SECURITY_NAME_ABBR,START_DATE,REPORT_DATE,FISCAL_YEAR,CURRENCY,ACCOUNT_STANDARD,REPORT_TYPE",
		"quoteColumns": "",
		"filter":       fmt.Sprintf(`(SECUCODE="%s.HK")`, stock),
		"source":       "F10",
		"client":       "PC",
	}

	resp, err := utils.Get(baseURL, params)
	if err != nil {
		return nil, fmt.Errorf("获取报告期列表失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	reportList := result.Get("result.data.0.REPORT_LIST").Array()
	if len(reportList) == 0 {
		return nil, nil
	}

	// 筛选年报或所有报告期
	var yearList []string
	for _, item := range reportList {
		reportType := item.Get("REPORT_TYPE").String()
		reportDate := item.Get("REPORT_DATE").String()
		if indicator == "年度" {
			if reportType == "年报" {
				parts := strings.Split(reportDate, " ")
				if len(parts) > 0 {
					yearList = append(yearList, parts[0])
				}
			}
		} else {
			parts := strings.Split(reportDate, " ")
			if len(parts) > 0 {
				yearList = append(yearList, parts[0])
			}
		}
	}

	if len(yearList) == 0 {
		return nil, nil
	}

	// 构建查询过滤条件
	yearFilter := "'" + strings.Join(yearList, "','") + "'"

	// 根据报表类型选择报告名称
	var reportName, columns string
	switch symbol {
	case "资产负债表":
		reportName = "RPT_HKF10_FN_BALANCE_PC"
		columns = "SECUCODE,SECURITY_CODE,SECURITY_NAME_ABBR,ORG_CODE,REPORT_DATE,DATE_TYPE_CODE,FISCAL_YEAR,STD_ITEM_CODE,STD_ITEM_NAME,AMOUNT,STD_REPORT_DATE"
	case "利润表":
		reportName = "RPT_HKF10_FN_INCOME_PC"
		columns = "SECUCODE,SECURITY_CODE,SECURITY_NAME_ABBR,ORG_CODE,REPORT_DATE,DATE_TYPE_CODE,FISCAL_YEAR,START_DATE,STD_ITEM_CODE,STD_ITEM_NAME,AMOUNT"
	case "现金流量表":
		reportName = "RPT_HKF10_FN_CASHFLOW_PC"
		columns = "SECUCODE,SECURITY_CODE,SECURITY_NAME_ABBR,ORG_CODE,REPORT_DATE,DATE_TYPE_CODE,FISCAL_YEAR,START_DATE,STD_ITEM_CODE,STD_ITEM_NAME,AMOUNT"
	default:
		return nil, fmt.Errorf("不支持的报表类型: %s", symbol)
	}

	params = map[string]string{
		"reportName":   reportName,
		"columns":      columns,
		"quoteColumns": "",
		"filter":       fmt.Sprintf(`(SECUCODE="%s.HK")(REPORT_DATE in (%s))`, stock, yearFilter),
		"pageNumber":   "1",
		"pageSize":     "",
		"sortTypes":    "-1,1",
		"sortColumns":  "REPORT_DATE,STD_ITEM_CODE",
		"source":       "F10",
		"client":       "PC",
	}

	resp, err = utils.Get(baseURL, params)
	if err != nil {
		return nil, fmt.Errorf("获取报表数据失败: %w", err)
	}

	result = gjson.ParseBytes(resp.Body())
	dataArr := result.Get("result.data").Array()

	var items []StockFinancialHkReportEmItem
	for _, item := range dataArr {
		items = append(items, StockFinancialHkReportEmItem{
			SecuCode:     item.Get("SECUCODE").String(),
			SecurityCode: item.Get("SECURITY_CODE").String(),
			SecurityName: item.Get("SECURITY_NAME_ABBR").String(),
			OrgCode:      item.Get("ORG_CODE").String(),
			ReportDate:   item.Get("REPORT_DATE").String(),
			DateTypeCode: item.Get("DATE_TYPE_CODE").String(),
			FiscalYear:   item.Get("FISCAL_YEAR").String(),
			StdItemCode:  item.Get("STD_ITEM_CODE").String(),
			StdItemName:  item.Get("STD_ITEM_NAME").String(),
			Amount:       item.Get("AMOUNT").Float(),
		})
	}

	return items, nil
}

// StockFinancialHkAnalysisIndicatorEm 东方财富-港股-财务分析-主要指标
// symbol: 股票代码，如 "00700"
// indicator: 指标类型，可选 "年度", "报告期"
func StockFinancialHkAnalysisIndicatorEm(symbol, indicator string) ([]map[string]interface{}, error) {
	baseURL := "https://datacenter.eastmoney.com/securities/api/data/v1/get"

	params := map[string]string{
		"reportName":   "RPT_HKF10_FN_MAININDICATOR",
		"columns":      "HKF10_FN_MAININDICATOR",
		"quoteColumns": "",
		"pageNumber":   "1",
		"pageSize":     "9",
		"sortTypes":    "-1",
		"sortColumns":  "STD_REPORT_DATE",
		"source":       "F10",
		"client":       "PC",
	}

	if indicator == "年度" {
		params["filter"] = fmt.Sprintf(`(SECUCODE="%s.HK")(DATE_TYPE_CODE="001")`, symbol)
	} else {
		params["filter"] = fmt.Sprintf(`(SECUCODE="%s.HK")`, symbol)
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
