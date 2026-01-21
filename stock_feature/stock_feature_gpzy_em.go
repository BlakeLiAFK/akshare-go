package stock_feature

import (
	"fmt"
	"sort"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// StockGpzyProfileItem 股权质押市场概况项
type StockGpzyProfileItem struct {
	TradeDate       string  `json:"trade_date"`        // 交易日期
	PledgeRatio     float64 `json:"pledge_ratio"`      // A股质押总比例
	CompanyCount    int64   `json:"company_count"`     // 质押公司数量
	PledgeCount     int64   `json:"pledge_count"`      // 质押笔数
	PledgeShares    float64 `json:"pledge_shares"`     // 质押总股数
	PledgeMarketCap float64 `json:"pledge_market_cap"` // 质押总市值
	Hs300Index      float64 `json:"hs300_index"`       // 沪深300指数
	ChangeRate      float64 `json:"change_rate"`       // 涨跌幅
}

// StockGpzyPledgeRatioItem 上市公司质押比例项
type StockGpzyPledgeRatioItem struct {
	Index           int     `json:"index"`             // 序号
	Code            string  `json:"code"`              // 股票代码
	Name            string  `json:"name"`              // 股票简称
	TradeDate       string  `json:"trade_date"`        // 交易日期
	Industry        string  `json:"industry"`          // 所属行业
	PledgeRatio     float64 `json:"pledge_ratio"`      // 质押比例
	PledgeShares    float64 `json:"pledge_shares"`     // 质押股数
	PledgeCount     int64   `json:"pledge_count"`      // 质押笔数
	UnlimitedPledge float64 `json:"unlimited_pledge"`  // 无限售股质押数
	LimitedPledge   float64 `json:"limited_pledge"`    // 限售股质押数
	PledgeMarketCap float64 `json:"pledge_market_cap"` // 质押市值
	YearChangeRate  float64 `json:"year_change_rate"`  // 近一年涨跌幅
}

// StockGpzyProfileEm 东方财富网-数据中心-特色数据-股权质押-股权质押市场概况
func StockGpzyProfileEm() ([]StockGpzyProfileItem, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"

	var allItems []StockGpzyProfileItem
	page := 1

	for {
		params := map[string]string{
			"sortColumns": "TRADE_DATE",
			"sortTypes":   "-1",
			"pageSize":    "500",
			"pageNumber":  fmt.Sprintf("%d", page),
			"reportName":  "RPT_CSDC_STATISTICS",
			"columns":     "ALL",
			"source":      "WEB",
			"client":      "WEB",
		}

		resp, err := utils.Get(url, params)
		if err != nil {
			return nil, fmt.Errorf("获取股权质押市场概况失败: %w", err)
		}

		result := gjson.ParseBytes(resp.Body())
		totalPages := result.Get("result.pages").Int()
		dataArr := result.Get("result.data").Array()

		for _, item := range dataArr {
			allItems = append(allItems, StockGpzyProfileItem{
				TradeDate:       strings.Split(item.Get("TRADE_DATE").String(), " ")[0],
				PledgeRatio:     item.Get("A_PLEDGE_RATIO").Float() / 100,
				CompanyCount:    item.Get("PLEDGE_ORG_NUM").Int(),
				PledgeCount:     item.Get("PLEDGE_NUM").Int(),
				PledgeShares:    item.Get("TOTAL_PLEDGE_SHARES").Float(),
				PledgeMarketCap: item.Get("TOTAL_PLEDGE_MARKETCAP").Float(),
				Hs300Index:      item.Get("HS300_INDEX").Float(),
				ChangeRate:      item.Get("CHANGE_RATE").Float(),
			})
		}

		if int64(page) >= totalPages {
			break
		}
		page++
	}

	// 按日期排序
	sort.Slice(allItems, func(i, j int) bool {
		return allItems[i].TradeDate < allItems[j].TradeDate
	})

	return allItems, nil
}

// StockGpzyPledgeRatioEm 东方财富网-数据中心-特色数据-股权质押-上市公司质押比例
// date: 指定交易日，格式 "20240906"
func StockGpzyPledgeRatioEm(date string) ([]StockGpzyPledgeRatioItem, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	tradeDate := date[:4] + "-" + date[4:6] + "-" + date[6:]

	var allItems []StockGpzyPledgeRatioItem
	page := 1

	for {
		params := map[string]string{
			"sortColumns": "PLEDGE_RATIO",
			"sortTypes":   "-1",
			"pageSize":    "500",
			"pageNumber":  fmt.Sprintf("%d", page),
			"reportName":  "RPT_CSDC_LIST",
			"columns":     "ALL",
			"source":      "WEB",
			"client":      "WEB",
			"filter":      fmt.Sprintf(`(TRADE_DATE='%s')`, tradeDate),
		}

		resp, err := utils.Get(url, params)
		if err != nil {
			return nil, fmt.Errorf("获取上市公司质押比例失败: %w", err)
		}

		result := gjson.ParseBytes(resp.Body())
		totalPages := result.Get("result.pages").Int()
		dataArr := result.Get("result.data").Array()

		for _, item := range dataArr {
			allItems = append(allItems, StockGpzyPledgeRatioItem{
				Index:           len(allItems) + 1,
				Code:            item.Get("SECURITY_CODE").String(),
				Name:            item.Get("SECURITY_NAME_ABBR").String(),
				TradeDate:       strings.Split(item.Get("TRADE_DATE").String(), " ")[0],
				Industry:        item.Get("INDUSTRY_NAME").String(),
				PledgeRatio:     item.Get("PLEDGE_RATIO").Float(),
				PledgeShares:    item.Get("PLEDGE_SHARES").Float(),
				PledgeCount:     item.Get("PLEDGE_NUM").Int(),
				UnlimitedPledge: item.Get("UNLIMITED_PLEDGE_SHARES").Float(),
				LimitedPledge:   item.Get("LIMITED_PLEDGE_SHARES").Float(),
				PledgeMarketCap: item.Get("PLEDGE_MARKET_CAP").Float(),
				YearChangeRate:  item.Get("YEAR_CHANGE_RATE").Float(),
			})
		}

		if int64(page) >= totalPages || totalPages == 0 {
			break
		}
		page++
	}

	return allItems, nil
}
