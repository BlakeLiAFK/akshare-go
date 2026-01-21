package stock_feature

import (
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// StockLhbDetailItem 龙虎榜详情项
type StockLhbDetailItem struct {
	Index         int     `json:"index"`           // 序号
	Code          string  `json:"code"`            // 代码
	Name          string  `json:"name"`            // 名称
	TradeDate     string  `json:"trade_date"`      // 上榜日
	Explain       string  `json:"explain"`         // 解读
	ClosePrice    float64 `json:"close_price"`     // 收盘价
	ChangeRate    float64 `json:"change_rate"`     // 涨跌幅
	NetBuyAmt     float64 `json:"net_buy_amt"`     // 龙虎榜净买额
	BuyAmt        float64 `json:"buy_amt"`         // 龙虎榜买入额
	SellAmt       float64 `json:"sell_amt"`        // 龙虎榜卖出额
	DealAmt       float64 `json:"deal_amt"`        // 龙虎榜成交额
	TotalDealAmt  float64 `json:"total_deal_amt"`  // 市场总成交额
	NetRatio      float64 `json:"net_ratio"`       // 净买额占总成交比
	DealRatio     float64 `json:"deal_ratio"`      // 成交额占总成交比
	TurnoverRate  float64 `json:"turnover_rate"`   // 换手率
	FreeMarketCap float64 `json:"free_market_cap"` // 流通市值
	Reason        string  `json:"reason"`          // 上榜原因
	After1Day     float64 `json:"after_1day"`      // 上榜后1日
	After2Day     float64 `json:"after_2day"`      // 上榜后2日
	After5Day     float64 `json:"after_5day"`      // 上榜后5日
	After10Day    float64 `json:"after_10day"`     // 上榜后10日
}

// StockLhbStatisticItem 龙虎榜统计项
type StockLhbStatisticItem struct {
	Index          int     `json:"index"`            // 序号
	Code           string  `json:"code"`             // 代码
	Name           string  `json:"name"`             // 名称
	ListCount      int64   `json:"list_count"`       // 上榜次数
	BuyAmt         float64 `json:"buy_amt"`          // 龙虎榜买入额
	SellAmt        float64 `json:"sell_amt"`         // 龙虎榜卖出额
	NetBuyAmt      float64 `json:"net_buy_amt"`      // 龙虎榜净买额
	TotalDealAmt   float64 `json:"total_deal_amt"`   // 总成交额
	LatestPrice    float64 `json:"latest_price"`     // 最新价
	ChangeRate     float64 `json:"change_rate"`      // 涨跌幅
	TotalMarketCap float64 `json:"total_market_cap"` // 总市值
}

// StockLhbDetailEm 东方财富网-数据中心-龙虎榜单-龙虎榜详情
// startDate, endDate: 日期范围，格式 "20230403"
func StockLhbDetailEm(startDate, endDate string) ([]StockLhbDetailItem, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	start := startDate[:4] + "-" + startDate[4:6] + "-" + startDate[6:]
	end := endDate[:4] + "-" + endDate[4:6] + "-" + endDate[6:]

	var allItems []StockLhbDetailItem
	page := 1

	for {
		params := map[string]string{
			"sortColumns": "SECURITY_CODE,TRADE_DATE",
			"sortTypes":   "1,-1",
			"pageSize":    "5000",
			"pageNumber":  fmt.Sprintf("%d", page),
			"reportName":  "RPT_DAILYBILLBOARD_DETAILSNEW",
			"columns":     "SECURITY_CODE,SECUCODE,SECURITY_NAME_ABBR,TRADE_DATE,EXPLAIN,CLOSE_PRICE,CHANGE_RATE,BILLBOARD_NET_AMT,BILLBOARD_BUY_AMT,BILLBOARD_SELL_AMT,BILLBOARD_DEAL_AMT,ACCUM_AMOUNT,DEAL_NET_RATIO,DEAL_AMOUNT_RATIO,TURNOVERRATE,FREE_MARKET_CAP,EXPLANATION,D1_CLOSE_ADJCHRATE,D2_CLOSE_ADJCHRATE,D5_CLOSE_ADJCHRATE,D10_CLOSE_ADJCHRATE,SECURITY_TYPE_CODE",
			"source":      "WEB",
			"client":      "WEB",
			"filter":      fmt.Sprintf(`(TRADE_DATE<='%s')(TRADE_DATE>='%s')`, end, start),
		}

		resp, err := utils.Get(url, params)
		if err != nil {
			return nil, fmt.Errorf("获取龙虎榜详情失败: %w", err)
		}

		result := gjson.ParseBytes(resp.Body())
		totalPages := result.Get("result.pages").Int()
		dataArr := result.Get("result.data").Array()

		for _, item := range dataArr {
			allItems = append(allItems, StockLhbDetailItem{
				Index:         len(allItems) + 1,
				Code:          item.Get("SECURITY_CODE").String(),
				Name:          item.Get("SECURITY_NAME_ABBR").String(),
				TradeDate:     strings.Split(item.Get("TRADE_DATE").String(), " ")[0],
				Explain:       item.Get("EXPLAIN").String(),
				ClosePrice:    item.Get("CLOSE_PRICE").Float(),
				ChangeRate:    item.Get("CHANGE_RATE").Float(),
				NetBuyAmt:     item.Get("BILLBOARD_NET_AMT").Float(),
				BuyAmt:        item.Get("BILLBOARD_BUY_AMT").Float(),
				SellAmt:       item.Get("BILLBOARD_SELL_AMT").Float(),
				DealAmt:       item.Get("BILLBOARD_DEAL_AMT").Float(),
				TotalDealAmt:  item.Get("ACCUM_AMOUNT").Float(),
				NetRatio:      item.Get("DEAL_NET_RATIO").Float(),
				DealRatio:     item.Get("DEAL_AMOUNT_RATIO").Float(),
				TurnoverRate:  item.Get("TURNOVERRATE").Float(),
				FreeMarketCap: item.Get("FREE_MARKET_CAP").Float(),
				Reason:        item.Get("EXPLANATION").String(),
				After1Day:     item.Get("D1_CLOSE_ADJCHRATE").Float(),
				After2Day:     item.Get("D2_CLOSE_ADJCHRATE").Float(),
				After5Day:     item.Get("D5_CLOSE_ADJCHRATE").Float(),
				After10Day:    item.Get("D10_CLOSE_ADJCHRATE").Float(),
			})
		}

		if int64(page) >= totalPages || totalPages == 0 {
			break
		}
		page++
	}

	return allItems, nil
}

// StockLhbStockStatisticEm 东方财富网-数据中心-龙虎榜单-个股上榜统计
// symbol: 时间范围，可选 "近一月", "近三月", "近六月", "近一年"
func StockLhbStockStatisticEm(symbol string) ([]StockLhbStatisticItem, error) {
	symbolMap := map[string]string{
		"近一月": "01",
		"近三月": "02",
		"近六月": "03",
		"近一年": "04",
	}

	typeCode := symbolMap[symbol]
	if typeCode == "" {
		typeCode = "01"
	}

	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"

	var allItems []StockLhbStatisticItem
	page := 1

	for {
		params := map[string]string{
			"sortColumns":  "BILLBOARD_TIMES,SECURITY_CODE",
			"sortTypes":    "-1,-1",
			"pageSize":     "500",
			"pageNumber":   fmt.Sprintf("%d", page),
			"reportName":   "RPT_BILLBOARD_DAILYSTATISTICS",
			"columns":      "ALL",
			"quoteColumns": "f2~01~SECURITY_CODE~CLOSE_PRICE,f3~01~SECURITY_CODE~CHANGE_RATE,f20~01~SECURITY_CODE~TOTAL_MARKET_CAP",
			"source":       "WEB",
			"client":       "WEB",
			"filter":       fmt.Sprintf(`(STATISTICS_CYCLE="%s")`, typeCode),
		}

		resp, err := utils.Get(url, params)
		if err != nil {
			return nil, fmt.Errorf("获取个股上榜统计失败: %w", err)
		}

		result := gjson.ParseBytes(resp.Body())
		totalPages := result.Get("result.pages").Int()
		dataArr := result.Get("result.data").Array()

		for _, item := range dataArr {
			allItems = append(allItems, StockLhbStatisticItem{
				Index:          len(allItems) + 1,
				Code:           item.Get("SECURITY_CODE").String(),
				Name:           item.Get("SECURITY_NAME_ABBR").String(),
				ListCount:      item.Get("BILLBOARD_TIMES").Int(),
				BuyAmt:         item.Get("BILLBOARD_BUY_AMT").Float(),
				SellAmt:        item.Get("BILLBOARD_SELL_AMT").Float(),
				NetBuyAmt:      item.Get("BILLBOARD_NET_AMT").Float(),
				TotalDealAmt:   item.Get("DEAL_AMT").Float(),
				LatestPrice:    item.Get("CLOSE_PRICE").Float(),
				ChangeRate:     item.Get("CHANGE_RATE").Float(),
				TotalMarketCap: item.Get("TOTAL_MARKET_CAP").Float(),
			})
		}

		if int64(page) >= totalPages || totalPages == 0 {
			break
		}
		page++
	}

	return allItems, nil
}
