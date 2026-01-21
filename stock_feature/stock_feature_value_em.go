package stock_feature

import (
	"fmt"
	"sort"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// StockValueItem 估值分析项
type StockValueItem struct {
	TradeDate          string  `json:"trade_date"`           // 数据日期
	ClosePrice         float64 `json:"close_price"`          // 当日收盘价
	ChangeRate         float64 `json:"change_rate"`          // 当日涨跌幅
	TotalMarketCap     float64 `json:"total_market_cap"`     // 总市值
	CirculateMarketCap float64 `json:"circulate_market_cap"` // 流通市值
	TotalShares        float64 `json:"total_shares"`         // 总股本
	FreeShares         float64 `json:"free_shares"`          // 流通股本
	PeTTM              float64 `json:"pe_ttm"`               // PE(TTM)
	PeLAR              float64 `json:"pe_lar"`               // PE(静)
	PbMRQ              float64 `json:"pb_mrq"`               // 市净率
	PegCAR             float64 `json:"peg_car"`              // PEG值
	PcfOcfTTM          float64 `json:"pcf_ocf_ttm"`          // 市现率
	PsTTM              float64 `json:"ps_ttm"`               // 市销率
}

// StockValueEm 东方财富网-数据中心-估值分析
// symbol: 股票代码
func StockValueEm(symbol string) ([]StockValueItem, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"sortColumns":  "TRADE_DATE",
		"sortTypes":    "-1",
		"pageSize":     "5000",
		"pageNumber":   "1",
		"reportName":   "RPT_VALUEANALYSIS_DET",
		"columns":      "ALL",
		"quoteColumns": "",
		"source":       "WEB",
		"client":       "WEB",
		"filter":       fmt.Sprintf(`(SECURITY_CODE="%s")`, symbol),
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("获取估值分析失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	dataArr := result.Get("result.data").Array()

	var items []StockValueItem
	for _, item := range dataArr {
		items = append(items, StockValueItem{
			TradeDate:          strings.Split(item.Get("TRADE_DATE").String(), " ")[0],
			ClosePrice:         item.Get("CLOSE_PRICE").Float(),
			ChangeRate:         item.Get("CHANGE_RATE").Float(),
			TotalMarketCap:     item.Get("TOTAL_MARKET_CAP").Float(),
			CirculateMarketCap: item.Get("NOTLIMITED_MARKETCAP_A").Float(),
			TotalShares:        item.Get("TOTAL_SHARES").Float(),
			FreeShares:         item.Get("FREE_SHARES_A").Float(),
			PeTTM:              item.Get("PE_TTM").Float(),
			PeLAR:              item.Get("PE_LAR").Float(),
			PbMRQ:              item.Get("PB_MRQ").Float(),
			PegCAR:             item.Get("PEG_CAR").Float(),
			PcfOcfTTM:          item.Get("PCF_OCF_TTM").Float(),
			PsTTM:              item.Get("PS_TTM").Float(),
		})
	}

	// 按日期排序
	sort.Slice(items, func(i, j int) bool {
		return items[i].TradeDate < items[j].TradeDate
	})

	return items, nil
}
