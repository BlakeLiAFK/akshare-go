package stock_feature

import (
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// StockYjbbItem 业绩报表项
type StockYjbbItem struct {
	Index               int     `json:"index"`                  // 序号
	Code                string  `json:"code"`                   // 股票代码
	Name                string  `json:"name"`                   // 股票简称
	EPS                 float64 `json:"eps"`                    // 每股收益
	TotalRevenue        float64 `json:"total_revenue"`          // 营业总收入
	TotalRevenueYoY     float64 `json:"total_revenue_yoy"`      // 营业总收入同比增长
	TotalRevenueQoQ     float64 `json:"total_revenue_qoq"`      // 营业总收入季度环比增长
	NetProfit           float64 `json:"net_profit"`             // 净利润
	NetProfitYoY        float64 `json:"net_profit_yoy"`         // 净利润同比增长
	NetProfitQoQ        float64 `json:"net_profit_qoq"`         // 净利润季度环比增长
	BPS                 float64 `json:"bps"`                    // 每股净资产
	ROE                 float64 `json:"roe"`                    // 净资产收益率
	OperatingCashFlowPS float64 `json:"operating_cash_flow_ps"` // 每股经营现金流量
	GrossProfitMargin   float64 `json:"gross_profit_margin"`    // 销售毛利率
	Industry            string  `json:"industry"`               // 所处行业
	LatestNoticeDate    string  `json:"latest_notice_date"`     // 最新公告日期
}

// StockYjbbEm 东方财富-数据中心-年报季报-业绩快报-业绩报表
// date: 报告期，如 "20200331", "20200630", "20200930", "20201231"
func StockYjbbEm(date string) ([]StockYjbbItem, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"

	// 格式化日期
	formattedDate := date[:4] + "-" + date[4:6] + "-" + date[6:]

	var allItems []StockYjbbItem
	page := 1

	for {
		params := map[string]string{
			"sortColumns": "UPDATE_DATE,SECURITY_CODE",
			"sortTypes":   "-1,-1",
			"pageSize":    "500",
			"pageNumber":  fmt.Sprintf("%d", page),
			"reportName":  "RPT_LICO_FN_CPD",
			"columns":     "ALL",
			"filter":      fmt.Sprintf(`(REPORTDATE='%s')`, formattedDate),
		}

		resp, err := utils.Get(url, params)
		if err != nil {
			return nil, fmt.Errorf("获取业绩报表失败: %w", err)
		}

		result := gjson.ParseBytes(resp.Body())
		totalPages := result.Get("result.pages").Int()
		dataArr := result.Get("result.data").Array()

		for _, item := range dataArr {
			allItems = append(allItems, StockYjbbItem{
				Index:               len(allItems) + 1,
				Code:                item.Get("SECURITY_CODE").String(),
				Name:                item.Get("SECURITY_NAME_ABBR").String(),
				EPS:                 item.Get("BASIC_EPS").Float(),
				TotalRevenue:        item.Get("TOTAL_OPERATE_INCOME").Float(),
				TotalRevenueYoY:     item.Get("YSTZ").Float(),
				TotalRevenueQoQ:     item.Get("YSHZ").Float(),
				NetProfit:           item.Get("PARENT_NETPROFIT").Float(),
				NetProfitYoY:        item.Get("SJLTZ").Float(),
				NetProfitQoQ:        item.Get("SJLHZ").Float(),
				BPS:                 item.Get("BPS").Float(),
				ROE:                 item.Get("WEIGHTAVG_ROE").Float(),
				OperatingCashFlowPS: item.Get("MGJYXJJE").Float(),
				GrossProfitMargin:   item.Get("XSMLL").Float(),
				Industry:            item.Get("INDUSTRY").String(),
				LatestNoticeDate:    strings.Split(item.Get("UPDATE_DATE").String(), " ")[0],
			})
		}

		if int64(page) >= totalPages {
			break
		}
		page++
	}

	return allItems, nil
}
