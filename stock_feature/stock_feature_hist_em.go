package stock_feature

import (
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// StockSpotItem 股票实时行情项
type StockSpotItem struct {
	Index           int     `json:"index"`             // 序号
	Code            string  `json:"code"`              // 代码
	Name            string  `json:"name"`              // 名称
	LatestPrice     float64 `json:"latest_price"`      // 最新价
	ChangeRate      float64 `json:"change_rate"`       // 涨跌幅
	ChangeAmount    float64 `json:"change_amount"`     // 涨跌额
	Volume          float64 `json:"volume"`            // 成交量
	Turnover        float64 `json:"turnover"`          // 成交额
	Amplitude       float64 `json:"amplitude"`         // 振幅
	High            float64 `json:"high"`              // 最高
	Low             float64 `json:"low"`               // 最低
	Open            float64 `json:"open"`              // 今开
	PrevClose       float64 `json:"prev_close"`        // 昨收
	VolumeRatio     float64 `json:"volume_ratio"`      // 量比
	TurnoverRate    float64 `json:"turnover_rate"`     // 换手率
	PeDynamic       float64 `json:"pe_dynamic"`        // 市盈率-动态
	Pb              float64 `json:"pb"`                // 市净率
	TotalMarketCap  float64 `json:"total_market_cap"`  // 总市值
	CirculateMktCap float64 `json:"circulate_mkt_cap"` // 流通市值
	ChangeSpeed     float64 `json:"change_speed"`      // 涨速
	Change5Min      float64 `json:"change_5min"`       // 5分钟涨跌
	Change60Day     float64 `json:"change_60day"`      // 60日涨跌幅
	ChangeYTD       float64 `json:"change_ytd"`        // 年初至今涨跌幅
}

// fetchStockSpotEm 通用获取实时行情函数
func fetchStockSpotEm(fs string) ([]StockSpotItem, error) {
	url := "https://82.push2.eastmoney.com/api/qt/clist/get"

	var allItems []StockSpotItem
	page := 1

	for {
		params := map[string]string{
			"pn":     fmt.Sprintf("%d", page),
			"pz":     "1000",
			"po":     "1",
			"np":     "1",
			"ut":     "bd1d9ddb04089700cf9c27f6f7426281",
			"fltt":   "2",
			"invt":   "2",
			"fid":    "f12",
			"fs":     fs,
			"fields": "f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18,f20,f21,f23,f24,f25,f22,f11,f62,f128,f136,f115,f152",
		}

		resp, err := utils.Get(url, params)
		if err != nil {
			return nil, fmt.Errorf("获取实时行情失败: %w", err)
		}

		result := gjson.ParseBytes(resp.Body())
		total := result.Get("data.total").Int()
		dataArr := result.Get("data.diff").Array()

		if len(dataArr) == 0 {
			break
		}

		for _, item := range dataArr {
			allItems = append(allItems, StockSpotItem{
				Index:           len(allItems) + 1,
				Code:            item.Get("f12").String(),
				Name:            item.Get("f14").String(),
				LatestPrice:     item.Get("f2").Float(),
				ChangeRate:      item.Get("f3").Float(),
				ChangeAmount:    item.Get("f4").Float(),
				Volume:          item.Get("f5").Float(),
				Turnover:        item.Get("f6").Float(),
				Amplitude:       item.Get("f7").Float(),
				High:            item.Get("f15").Float(),
				Low:             item.Get("f16").Float(),
				Open:            item.Get("f17").Float(),
				PrevClose:       item.Get("f18").Float(),
				VolumeRatio:     item.Get("f10").Float(),
				TurnoverRate:    item.Get("f8").Float(),
				PeDynamic:       item.Get("f9").Float(),
				Pb:              item.Get("f23").Float(),
				TotalMarketCap:  item.Get("f20").Float(),
				CirculateMktCap: item.Get("f21").Float(),
				ChangeSpeed:     item.Get("f22").Float(),
				Change5Min:      item.Get("f11").Float(),
				Change60Day:     item.Get("f24").Float(),
				ChangeYTD:       item.Get("f25").Float(),
			})
		}

		if int64(len(allItems)) >= total {
			break
		}
		page++
	}

	return allItems, nil
}

// StockZhASpotEm 东方财富网-沪深京 A 股-实时行情
func StockZhASpotEm() ([]StockSpotItem, error) {
	return fetchStockSpotEm("m:0 t:6,m:0 t:80,m:1 t:2,m:1 t:23,m:0 t:81 s:2048")
}

// StockShASpotEm 东方财富网-沪 A 股-实时行情
func StockShASpotEm() ([]StockSpotItem, error) {
	return fetchStockSpotEm("m:1 t:2,m:1 t:23")
}

// StockSzASpotEm 东方财富网-深 A 股-实时行情
func StockSzASpotEm() ([]StockSpotItem, error) {
	return fetchStockSpotEm("m:0 t:6,m:0 t:80")
}

// StockBjASpotEm 东方财富网-北 A 股-实时行情
func StockBjASpotEm() ([]StockSpotItem, error) {
	return fetchStockSpotEm("m:0 t:81 s:2048")
}

// StockCybASpotEm 东方财富网-创业板-实时行情
func StockCybASpotEm() ([]StockSpotItem, error) {
	return fetchStockSpotEm("m:0 t:80")
}

// StockKcbASpotEm 东方财富网-科创板-实时行情
func StockKcbASpotEm() ([]StockSpotItem, error) {
	return fetchStockSpotEm("m:1 t:23")
}

// StockNewSpotEm 东方财富网-新股-实时行情
func StockNewSpotEm() ([]StockSpotItem, error) {
	return fetchStockSpotEm("m:0 f:8,m:1 f:8")
}

// StockHsgtSpotEm 东方财富网-沪深港通持股-实时行情
func StockHsgtSpotEm() ([]StockSpotItem, error) {
	return fetchStockSpotEm("b:BK0707")
}

// StockZhBSpotEm 东方财富网-沪深 B 股-实时行情
func StockZhBSpotEm() ([]StockSpotItem, error) {
	return fetchStockSpotEm("m:0 t:7,m:1 t:3")
}
