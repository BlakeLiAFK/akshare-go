package futures

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
)

// FuturesHFSpotEM 高频期货实时行情
type FuturesHFSpotEM struct {
	Symbol     string    `json:"symbol"`      // 合约代码
	Name       string    `json:"name"`        // 合约名称
	Price      float64   `json:"price"`       // 最新价
	ChangePct  float64   `json:"change_pct"`  // 涨跌幅
	Change     float64   `json:"change"`      // 涨跌额
	Volume     int64     `json:"volume"`      // 成交量
	Amount     float64   `json:"amount"`      // 成交额
	Hold       int64     `json:"hold"`        // 持仓量
	High       float64   `json:"high"`        // 最高
	Low        float64   `json:"low"`         // 最低
	Open       float64   `json:"open"`        // 开盘
	PreClose   float64   `json:"pre_close"`   // 昨收
	PreSettle  float64   `json:"pre_settle"`  // 昨结算
	UpdateTime time.Time `json:"update_time"` // 更新时间
}

// hfSpotResponse 高频行情API响应
type hfSpotResponse struct {
	Data struct {
		Total int `json:"total"`
		Diff  []struct {
			F12 string      `json:"f12"` // 代码
			F14 string      `json:"f14"` // 名称
			F2  interface{} `json:"f2"`  // 最新价
			F3  interface{} `json:"f3"`  // 涨跌幅
			F4  interface{} `json:"f4"`  // 涨跌额
			F5  interface{} `json:"f5"`  // 成交量
			F6  interface{} `json:"f6"`  // 成交额
			F7  interface{} `json:"f7"`  // 振幅
			F15 interface{} `json:"f15"` // 最高
			F16 interface{} `json:"f16"` // 最低
			F17 interface{} `json:"f17"` // 今开
			F18 interface{} `json:"f18"` // 昨收
			F57 interface{} `json:"f57"` // 昨结算
			F8  interface{} `json:"f8"`  // 持仓量
		} `json:"diff"`
	} `json:"data"`
}

// FuturesHFSpotEMFunc 东方财富-高频期货实时行情
//
// 数据源: https://quote.eastmoney.com/center/gridlist.html#futures_qhjy
//
// 参数:
//   - symbol: 市场类型，可选:
//     "全部合约", "上期所", "大商所", "郑商所", "中金所", "上期能源", "广期所"
//
// 返回:
//   - []FuturesHFSpotEM: 高频期货实时行情
//   - error: 错误信息
func FuturesHFSpotEMFunc(symbol string) ([]FuturesHFSpotEM, error) {
	symbolMap := map[string]string{
		"全部合约": "m:113,m:114,m:115,m:140,m:141,m:142",
		"上期所":  "m:113",
		"大商所":  "m:114",
		"郑商所":  "m:115",
		"中金所":  "m:8",
		"上期能源": "m:142",
		"广期所":  "m:225",
	}

	fs, ok := symbolMap[symbol]
	if !ok {
		return nil, fmt.Errorf("无效的 symbol: %s", symbol)
	}

	url := "https://80.push2.eastmoney.com/api/qt/clist/get"

	var result []FuturesHFSpotEM
	page := 1

	for {
		params := map[string]string{
			"pn":     fmt.Sprintf("%d", page),
			"pz":     "500",
			"po":     "1",
			"np":     "1",
			"ut":     "bd1d9ddb04089700cf9c27f6f7426281",
			"fltt":   "2",
			"invt":   "2",
			"fid":    "f3",
			"fs":     fs,
			"fields": "f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18,f20,f21,f23,f24,f25,f26,f22,f57",
		}

		resp, err := utils.Get(url, params)
		if err != nil {
			return nil, fmt.Errorf("请求高频期货行情失败: %w", err)
		}

		var apiResp hfSpotResponse
		if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
			return nil, fmt.Errorf("解析高频期货行情响应失败: %w", err)
		}

		if len(apiResp.Data.Diff) == 0 {
			break
		}

		for _, item := range apiResp.Data.Diff {
			result = append(result, FuturesHFSpotEM{
				Symbol:    item.F12,
				Name:      item.F14,
				Price:     toFloat64(item.F2),
				ChangePct: toFloat64(item.F3),
				Change:    toFloat64(item.F4),
				Volume:    int64(toFloat64(item.F5)),
				Amount:    toFloat64(item.F6),
				Hold:      int64(toFloat64(item.F8)),
				High:      toFloat64(item.F15),
				Low:       toFloat64(item.F16),
				Open:      toFloat64(item.F17),
				PreClose:  toFloat64(item.F18),
				PreSettle: toFloat64(item.F57),
			})
		}

		if len(apiResp.Data.Diff) < 500 {
			break
		}
		page++
	}

	return result, nil
}

// FuturesHFMinuteEM 高频期货分钟数据
type FuturesHFMinuteEM struct {
	Time   string  `json:"time"`   // 时间
	Price  float64 `json:"price"`  // 价格
	Volume int64   `json:"volume"` // 成交量
	Hold   int64   `json:"hold"`   // 持仓量
	AvgP   float64 `json:"avg_p"`  // 均价
}

// hfMinuteResponse 高频分钟数据API响应
type hfMinuteResponse struct {
	Data struct {
		Code   string   `json:"code"`
		Trends []string `json:"trends"`
	} `json:"data"`
}

// FuturesHFMinuteEMFunc 东方财富-高频期货分钟数据
//
// 数据源: https://quote.eastmoney.com/qihuo/sc2501.html
//
// 参数:
//   - symbol: 合约代码，如 "sc2501"
//
// 返回:
//   - []FuturesHFMinuteEM: 高频期货分钟数据
//   - error: 错误信息
func FuturesHFMinuteEMFunc(symbol string) ([]FuturesHFMinuteEM, error) {
	// 确定市场代码
	marketMap := map[string]string{
		"shfe":  "113", // 上期所
		"dce":   "114", // 大商所
		"czce":  "115", // 郑商所
		"cffex": "8",   // 中金所
		"ine":   "142", // 上期能源
		"gfex":  "225", // 广期所
	}

	var secid string
	symbolUpper := strings.ToUpper(symbol)
	market := SymbolMarket(symbolUpper)
	if mktCode, ok := marketMap[market]; ok {
		secid = fmt.Sprintf("%s.%s", mktCode, symbol)
	} else {
		// 默认使用上期所
		secid = fmt.Sprintf("113.%s", symbol)
	}

	url := "https://push2.eastmoney.com/api/qt/stock/trends2/get"
	params := map[string]string{
		"secid":   secid,
		"fields1": "f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f11,f12,f13",
		"fields2": "f51,f52,f53,f54,f55,f56,f57,f58",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求高频期货分钟数据失败: %w", err)
	}

	var apiResp hfMinuteResponse
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, fmt.Errorf("解析高频期货分钟数据响应失败: %w", err)
	}

	result := make([]FuturesHFMinuteEM, 0, len(apiResp.Data.Trends))
	for _, line := range apiResp.Data.Trends {
		parts := strings.Split(line, ",")
		if len(parts) < 5 {
			continue
		}

		result = append(result, FuturesHFMinuteEM{
			Time:   parts[0],
			Price:  utils.MustParseFloat(parts[1]),
			Volume: int64(utils.MustParseFloat(parts[2])),
			Hold:   int64(utils.MustParseFloat(parts[3])),
			AvgP:   utils.MustParseFloat(parts[4]),
		})
	}

	return result, nil
}
