package forex

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// parseFloat 安全解析浮点数
func parseFloat(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

// ForexSpotItem 外汇实时行情项
type ForexSpotItem struct {
	Index        int     `json:"index"`         // 序号
	Code         string  `json:"code"`          // 代码
	Name         string  `json:"name"`          // 名称
	LatestPrice  float64 `json:"latest_price"`  // 最新价
	ChangeAmount float64 `json:"change_amount"` // 涨跌额
	ChangeRate   float64 `json:"change_rate"`   // 涨跌幅
	Open         float64 `json:"open"`          // 今开
	High         float64 `json:"high"`          // 最高
	Low          float64 `json:"low"`           // 最低
	PrevClose    float64 `json:"prev_close"`    // 昨收
}

// ForexHistItem 外汇历史行情项
type ForexHistItem struct {
	Date      string  `json:"date"`      // 日期
	Open      float64 `json:"open"`      // 今开
	Close     float64 `json:"close"`     // 最新价
	High      float64 `json:"high"`      // 最高
	Low       float64 `json:"low"`       // 最低
	Amplitude float64 `json:"amplitude"` // 振幅
	Code      string  `json:"code"`      // 代码
	Name      string  `json:"name"`      // 名称
}

// 货币市场代码映射
var symbolMarketMap = map[string]string{
	"USDCNH": "133", "EURCNH": "133", "GBPCNH": "133", "JPYCNH": "133",
	"HKDCNH": "133", "AUDCNH": "133", "CADCNH": "133", "CHFCNH": "133",
	"EURUSD": "119", "GBPUSD": "119", "USDJPY": "119", "USDCHF": "119",
	"AUDUSD": "119", "USDCAD": "119", "NZDUSD": "119", "EURGBP": "119",
}

// ForexSpotEm 东方财富网-行情中心-外汇市场-所有汇率-实时行情数据
func ForexSpotEm() ([]ForexSpotItem, error) {
	url := "https://push2.eastmoney.com/api/qt/clist/get"

	var allItems []ForexSpotItem
	page := 1

	for {
		params := map[string]string{
			"np":     "1",
			"fltt":   "2",
			"invt":   "2",
			"fs":     "m:119,m:120,m:133",
			"fields": "f12,f13,f14,f1,f2,f4,f3,f152,f17,f18,f15,f16",
			"fid":    "f3",
			"pn":     fmt.Sprintf("%d", page),
			"pz":     "100",
			"po":     "1",
			"dect":   "1",
		}

		resp, err := utils.Get(url, params)
		if err != nil {
			return nil, fmt.Errorf("获取外汇实时行情失败: %w", err)
		}

		result := gjson.ParseBytes(resp.Body())
		total := result.Get("data.total").Int()
		dataArr := result.Get("data.diff").Array()

		if len(dataArr) == 0 {
			break
		}

		for _, item := range dataArr {
			allItems = append(allItems, ForexSpotItem{
				Index:        len(allItems) + 1,
				Code:         item.Get("f12").String(),
				Name:         item.Get("f14").String(),
				LatestPrice:  item.Get("f2").Float(),
				ChangeAmount: item.Get("f4").Float(),
				ChangeRate:   item.Get("f3").Float(),
				Open:         item.Get("f17").Float(),
				High:         item.Get("f15").Float(),
				Low:          item.Get("f16").Float(),
				PrevClose:    item.Get("f18").Float(),
			})
		}

		if int64(len(allItems)) >= total {
			break
		}
		page++
	}

	return allItems, nil
}

// ForexHistEm 东方财富网-行情中心-外汇市场-所有汇率-历史行情数据
// symbol: 品种代码，如 "USDCNH"
func ForexHistEm(symbol string) ([]ForexHistItem, error) {
	marketCode, ok := symbolMarketMap[symbol]
	if !ok {
		marketCode = "133"
	}

	url := "https://push2his.eastmoney.com/api/qt/stock/kline/get"
	params := map[string]string{
		"secid":   fmt.Sprintf("%s.%s", marketCode, symbol),
		"klt":     "101",
		"fqt":     "1",
		"lmt":     "50000",
		"end":     "20500000",
		"iscca":   "1",
		"fields1": "f1,f2,f3,f4,f5,f6,f7,f8",
		"fields2": "f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61,f62,f63,f64",
		"ut":      "f057cbcbce2a86e2866ab8877db1d059",
		"forcect": "1",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("获取外汇历史行情失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	code := result.Get("data.code").String()
	name := result.Get("data.name").String()
	klinesArr := result.Get("data.klines").Array()

	var items []ForexHistItem
	for _, kline := range klinesArr {
		parts := strings.Split(kline.String(), ",")
		if len(parts) < 8 {
			continue
		}
		items = append(items, ForexHistItem{
			Date:      parts[0],
			Open:      parseFloat(parts[1]),
			Close:     parseFloat(parts[2]),
			High:      parseFloat(parts[3]),
			Low:       parseFloat(parts[4]),
			Amplitude: parseFloat(parts[7]),
			Code:      code,
			Name:      name,
		})
	}

	return items, nil
}
