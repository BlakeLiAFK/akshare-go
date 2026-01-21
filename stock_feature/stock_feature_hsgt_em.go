package stock_feature

import (
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// StockHsgtFundFlowItem 沪深港通资金流向项
type StockHsgtFundFlowItem struct {
	TradeDate      string  `json:"trade_date"`      // 交易日
	Type           string  `json:"type"`            // 类型
	Board          string  `json:"board"`           // 板块
	FundsDirection string  `json:"funds_direction"` // 资金方向
	TradeStatus    string  `json:"trade_status"`    // 交易状态
	NetBuyAmt      float64 `json:"net_buy_amt"`     // 成交净买额(亿)
	NetFlowIn      float64 `json:"net_flow_in"`     // 资金净流入(亿)
	DayRemain      float64 `json:"day_remain"`      // 当日资金余额(亿)
	RiseCount      int64   `json:"rise_count"`      // 上涨数
	FlatCount      int64   `json:"flat_count"`      // 持平数
	FallCount      int64   `json:"fall_count"`      // 下跌数
	IndexName      string  `json:"index_name"`      // 相关指数
	IndexChange    float64 `json:"index_change"`    // 指数涨跌幅
}

// StockHsgtHoldItem 沪深港通持股排行项
type StockHsgtHoldItem struct {
	Index         int     `json:"index"`           // 序号
	Code          string  `json:"code"`            // 股票代码
	Name          string  `json:"name"`            // 股票简称
	LatestPrice   float64 `json:"latest_price"`    // 最新价
	ChangeRate    float64 `json:"change_rate"`     // 涨跌幅
	HoldShares    float64 `json:"hold_shares"`     // 持股数量
	HoldMarketCap float64 `json:"hold_market_cap"` // 持股市值
	HoldRatio     float64 `json:"hold_ratio"`      // 持股占流通股比
	NetBuyShares  float64 `json:"net_buy_shares"`  // 净买入股数
	Industry      string  `json:"industry"`        // 所属行业
}

// StockHsgtFundFlowSummaryEm 东方财富网-数据中心-资金流向-沪深港通资金流向
func StockHsgtFundFlowSummaryEm() ([]StockHsgtFundFlowItem, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"reportName":   "RPT_MUTUAL_QUOTA",
		"columns":      "TRADE_DATE,MUTUAL_TYPE,BOARD_TYPE,MUTUAL_TYPE_NAME,FUNDS_DIRECTION,INDEX_CODE,INDEX_NAME,BOARD_CODE",
		"quoteColumns": "status~07~BOARD_CODE,dayNetAmtIn~07~BOARD_CODE,dayAmtRemain~07~BOARD_CODE,dayAmtThreshold~07~BOARD_CODE,f104~07~BOARD_CODE,f105~07~BOARD_CODE,f106~07~BOARD_CODE,f3~03~INDEX_CODE~INDEX_f3,netBuyAmt~07~BOARD_CODE",
		"quoteType":    "0",
		"pageNumber":   "1",
		"pageSize":     "2000",
		"sortTypes":    "1",
		"sortColumns":  "MUTUAL_TYPE",
		"source":       "WEB",
		"client":       "WEB",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("获取沪深港通资金流向失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	dataArr := result.Get("result.data").Array()

	var items []StockHsgtFundFlowItem
	for _, item := range dataArr {
		items = append(items, StockHsgtFundFlowItem{
			TradeDate:      strings.Split(item.Get("TRADE_DATE").String(), " ")[0],
			Type:           item.Get("BOARD_TYPE").String(),
			Board:          item.Get("MUTUAL_TYPE_NAME").String(),
			FundsDirection: item.Get("FUNDS_DIRECTION").String(),
			TradeStatus:    item.Get("status").String(),
			NetBuyAmt:      item.Get("netBuyAmt").Float() / 10000,
			NetFlowIn:      item.Get("dayNetAmtIn").Float() / 10000,
			DayRemain:      item.Get("dayAmtRemain").Float() / 10000,
			RiseCount:      item.Get("f104").Int(),
			FlatCount:      item.Get("f106").Int(),
			FallCount:      item.Get("f105").Int(),
			IndexName:      item.Get("INDEX_NAME").String(),
			IndexChange:    item.Get("INDEX_f3").Float(),
		})
	}

	return items, nil
}

// StockHsgtHoldStockEm 东方财富-数据中心-沪深港通持股-个股排行
// market: 市场类型，可选 "北向", "沪股通", "深股通"
// indicator: 排行周期，可选 "今日排行", "3日排行", "5日排行", "10日排行", "月排行", "季排行", "年排行"
func StockHsgtHoldStockEm(market, indicator string) ([]StockHsgtHoldItem, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"

	indicatorMap := map[string]string{
		"今日排行":  "1",
		"3日排行":  "3",
		"5日排行":  "5",
		"10日排行": "10",
		"月排行":   "M",
		"季排行":   "Q",
		"年排行":   "Y",
	}

	marketMap := map[string]string{
		"北向":  "",
		"沪股通": "001",
		"深股通": "003",
	}

	indicatorType := indicatorMap[indicator]
	if indicatorType == "" {
		indicatorType = "5"
	}

	var allItems []StockHsgtHoldItem
	page := 1

	for {
		params := map[string]string{
			"sortColumns": fmt.Sprintf("ADD_SHARES_RATE%s", indicatorType),
			"sortTypes":   "-1",
			"pageSize":    "500",
			"pageNumber":  fmt.Sprintf("%d", page),
			"reportName":  "RPT_MUTUAL_STOCK_NORTHSTA",
			"columns":     "ALL",
			"source":      "WEB",
			"client":      "WEB",
		}

		if marketCode, ok := marketMap[market]; ok && marketCode != "" {
			params["filter"] = fmt.Sprintf(`(MARKET_CODE="%s")`, marketCode)
		}

		resp, err := utils.Get(url, params)
		if err != nil {
			return nil, fmt.Errorf("获取沪深港通持股排行失败: %w", err)
		}

		result := gjson.ParseBytes(resp.Body())
		totalPages := result.Get("result.pages").Int()
		dataArr := result.Get("result.data").Array()

		for _, item := range dataArr {
			allItems = append(allItems, StockHsgtHoldItem{
				Index:         len(allItems) + 1,
				Code:          item.Get("SECURITY_CODE").String(),
				Name:          item.Get("SECURITY_NAME_ABBR").String(),
				LatestPrice:   item.Get("NEW_PRICE").Float(),
				ChangeRate:    item.Get("CHANGE_RATE").Float(),
				HoldShares:    item.Get("HOLD_SHARES").Float(),
				HoldMarketCap: item.Get("HOLD_MARKET_CAP").Float(),
				HoldRatio:     item.Get("HOLD_RATIO").Float(),
				NetBuyShares:  item.Get("NET_BUY_SHARES").Float(),
				Industry:      item.Get("INDUSTRY_NAME").String(),
			})
		}

		if int64(page) >= totalPages || totalPages == 0 {
			break
		}
		page++
	}

	return allItems, nil
}

// StockHkGgtComponentsEm 东方财富网-行情中心-港股市场-港股通成份股
func StockHkGgtComponentsEm() ([]StockSpotItem, error) {
	return fetchStockSpotEm("b:DLMK0146,b:DLMK0144")
}
