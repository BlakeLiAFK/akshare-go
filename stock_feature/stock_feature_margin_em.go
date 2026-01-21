package stock_feature

import (
	"fmt"
	"sort"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// StockMarginAccountInfoItem 融资融券账户统计项
type StockMarginAccountInfoItem struct {
	Date                  string  `json:"date"`                     // 日期
	FinBalance            float64 `json:"fin_balance"`              // 融资余额
	LoanBalance           float64 `json:"loan_balance"`             // 融券余额
	FinBuyAmt             float64 `json:"fin_buy_amt"`              // 融资买入额
	LoanSellAmt           float64 `json:"loan_sell_amt"`            // 融券卖出额
	SecurityOrgNum        int     `json:"security_org_num"`         // 证券公司数量
	OperateDeptNum        int     `json:"operate_dept_num"`         // 营业部数量
	PersonalInvestorNum   int     `json:"personal_investor_num"`    // 个人投资者数量
	OrgInvestorNum        int     `json:"org_investor_num"`         // 机构投资者数量
	InvestorNum           int     `json:"investor_num"`             // 参与交易的投资者数量
	MarginLiabInvestorNum int     `json:"margin_liab_investor_num"` // 有融资融券负债的投资者数量
	TotalGuarantee        float64 `json:"total_guarantee"`          // 担保物总价值
	AvgGuaranteeRatio     float64 `json:"avg_guarantee_ratio"`      // 平均维持担保比例
}

// StockMarginAccountInfo 东方财富网-数据中心-融资融券-融资融券账户统计-两融账户信息
func StockMarginAccountInfo() ([]StockMarginAccountInfoItem, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"

	var allItems []StockMarginAccountInfoItem
	page := 1

	for {
		params := map[string]string{
			"reportName":  "RPTA_WEB_MARGIN_DAILYTRADE",
			"columns":     "ALL",
			"pageNumber":  fmt.Sprintf("%d", page),
			"pageSize":    "500",
			"sortColumns": "STATISTICS_DATE",
			"sortTypes":   "-1",
		}

		resp, err := utils.Get(url, params)
		if err != nil {
			return nil, fmt.Errorf("获取融资融券账户统计失败: %w", err)
		}

		result := gjson.ParseBytes(resp.Body())
		totalPages := result.Get("result.pages").Int()
		dataArr := result.Get("result.data").Array()

		for _, item := range dataArr {
			allItems = append(allItems, StockMarginAccountInfoItem{
				Date:                  strings.Split(item.Get("STATISTICS_DATE").String(), " ")[0],
				FinBalance:            item.Get("FIN_BALANCE").Float(),
				LoanBalance:           item.Get("LOAN_BALANCE").Float(),
				FinBuyAmt:             item.Get("FIN_BUY_AMT").Float(),
				LoanSellAmt:           item.Get("LOAN_SELL_AMT").Float(),
				SecurityOrgNum:        int(item.Get("SECURITY_ORG_NUM").Int()),
				OperateDeptNum:        int(item.Get("OPERATEDEPT_NUM").Int()),
				PersonalInvestorNum:   int(item.Get("PERSONAL_INVESTOR_NUM").Int()),
				OrgInvestorNum:        int(item.Get("ORG_INVESTOR_NUM").Int()),
				InvestorNum:           int(item.Get("INVESTOR_NUM").Int()),
				MarginLiabInvestorNum: int(item.Get("MARGINLIAB_INVESTOR_NUM").Int()),
				TotalGuarantee:        item.Get("TOTAL_GUARANTEE").Float(),
				AvgGuaranteeRatio:     item.Get("AVG_GUARANTEE_RATIO").Float(),
			})
		}

		if int64(page) >= totalPages {
			break
		}
		page++
	}

	// 按日期排序
	sort.Slice(allItems, func(i, j int) bool {
		return allItems[i].Date < allItems[j].Date
	})

	return allItems, nil
}
