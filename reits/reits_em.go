package reits

import (
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

const (
	// 东方财富REITs实时行情URL
	emReitsRealtimeURL = "https://push2.eastmoney.com/api/qt/clist/get"

	// 东方财富REITs历史行情URL
	emReitsHistURL = "https://push2his.eastmoney.com/api/qt/stock/kline/get"
)

// ReitsRealtimeItem REITs实时行情数据结构
type ReitsRealtimeItem struct {
	Seq       int     `json:"seq"`        // 序号
	Code      string  `json:"code"`       // 代码
	Name      string  `json:"name"`       // 名称
	Price     float64 `json:"price"`      // 最新价
	ChangePct float64 `json:"change_pct"` // 涨跌幅
	Change    float64 `json:"change"`     // 涨跌额
	Volume    int64   `json:"volume"`     // 成交量
	Amount    float64 `json:"amount"`     // 成交额
	Open      float64 `json:"open"`       // 开盘价
	High      float64 `json:"high"`       // 最高价
	Low       float64 `json:"low"`        // 最低价
	PreClose  float64 `json:"pre_close"`  // 昨收
}

// ReitsHistItem REITs历史行情数据结构
type ReitsHistItem struct {
	Date      string  `json:"date"`      // 日期
	Open      float64 `json:"open"`      // 今开
	High      float64 `json:"high"`      // 最高
	Low       float64 `json:"low"`       // 最低
	Close     float64 `json:"close"`     // 最新价
	Volume    int64   `json:"volume"`    // 成交量
	Amount    float64 `json:"amount"`    // 成交额
	Amplitude float64 `json:"amplitude"` // 振幅
	Turnover  float64 `json:"turnover"`  // 换手
}

// ReitsRealtimeEm 获取东方财富网-REITs-沪深REITs-实时行情
//
// 目标地址: http://quote.eastmoney.com/center/gridlist.html#fund_reits_all
//
// 返回:
//   - []ReitsRealtimeItem: REITs实时行情数据列表
//   - error: 错误信息
//
// 示例:
//
//	data, err := reits.ReitsRealtimeEm()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range data {
//	    fmt.Printf("%s %s %.3f %.2f%%\n", item.Code, item.Name, item.Price, item.ChangePct)
//	}
func ReitsRealtimeEm() ([]ReitsRealtimeItem, error) {
	params := map[string]string{
		"pn":     "1",
		"pz":     "200",
		"po":     "1",
		"np":     "1",
		"fltt":   "2",
		"invt":   "2",
		"fid":    "f3",
		"fs":     "b:MK0802",
		"fields": "f1,f2,f3,f4,f5,f6,f7,f12,f13,f14,f15,f16,f17,f18",
	}

	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		"Referer":    "http://quote.eastmoney.com/",
	}

	resp, err := utils.GetWithHeaders(emReitsRealtimeURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("获取REITs实时行情失败: %w", err)
	}

	text := resp.String()
	dataPath := "data.diff"
	result := gjson.Get(text, dataPath)
	if !result.Exists() {
		return nil, fmt.Errorf("解析REITs数据失败: 未找到 %s", dataPath)
	}

	var items []ReitsRealtimeItem
	seq := 1
	result.ForEach(func(_, value gjson.Result) bool {
		item := ReitsRealtimeItem{
			Seq:       seq,
			Code:      value.Get("f12").String(),
			Name:      value.Get("f14").String(),
			Price:     value.Get("f2").Float(),
			ChangePct: value.Get("f3").Float(),
			Change:    value.Get("f4").Float(),
			Volume:    value.Get("f5").Int(),
			Amount:    value.Get("f6").Float(),
			Open:      value.Get("f17").Float(),
			High:      value.Get("f15").Float(),
			Low:       value.Get("f16").Float(),
			PreClose:  value.Get("f18").Float(),
		}
		items = append(items, item)
		seq++
		return true
	})

	return items, nil
}

// ReitsHistEm 获取东方财富网-REITs-沪深REITs-历史行情
//
// 目标地址: https://quote.eastmoney.com/sh508097.html
//
// 参数:
//   - symbol: REITs代码，如 "508097"
//
// 返回:
//   - []ReitsHistItem: REITs历史行情数据列表
//   - error: 错误信息
//
// 示例:
//
//	data, err := reits.ReitsHistEm("508097")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range data {
//	    fmt.Printf("%s %.3f %.3f %.3f %.3f\n", item.Date, item.Open, item.High, item.Low, item.Close)
//	}
func ReitsHistEm(symbol string) ([]ReitsHistItem, error) {
	// 确定市场代码
	secid := "1." + symbol // 默认上海
	if len(symbol) >= 1 && (symbol[0] == '1' || symbol[0] == '0') {
		secid = "0." + symbol // 深圳
	}

	params := map[string]string{
		"secid":   secid,
		"fields1": "f1,f2,f3,f4,f5,f6",
		"fields2": "f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61",
		"klt":     "101",
		"fqt":     "1",
		"beg":     "0",
		"end":     "20500101",
	}

	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		"Referer":    "http://quote.eastmoney.com/",
	}

	resp, err := utils.GetWithHeaders(emReitsHistURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("获取REITs历史行情失败: %w", err)
	}

	text := resp.String()
	klines := gjson.Get(text, "data.klines")
	if !klines.Exists() {
		return nil, fmt.Errorf("解析REITs历史数据失败")
	}

	var items []ReitsHistItem
	klines.ForEach(func(_, value gjson.Result) bool {
		parts := splitKline(value.String())
		if len(parts) >= 9 {
			item := ReitsHistItem{
				Date:      parts[0],
				Open:      utils.MustFloat64(parts[1]),
				Close:     utils.MustFloat64(parts[2]),
				High:      utils.MustFloat64(parts[3]),
				Low:       utils.MustFloat64(parts[4]),
				Volume:    utils.MustInt64(parts[5]),
				Amount:    utils.MustFloat64(parts[6]),
				Amplitude: utils.MustFloat64(parts[7]),
				Turnover:  utils.MustFloat64(parts[8]),
			}
			items = append(items, item)
		}
		return true
	})

	return items, nil
}

// splitKline 分割K线数据字符串
func splitKline(s string) []string {
	var result []string
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ',' {
			result = append(result, s[start:i])
			start = i + 1
		}
	}
	if start < len(s) {
		result = append(result, s[start:])
	}
	return result
}

// ReitsHistMinItem REITs分钟行情数据结构
type ReitsHistMinItem struct {
	Time     string  `json:"time"`      // 时间
	Price    float64 `json:"price"`     // 最新价
	High     float64 `json:"high"`      // 最高
	Low      float64 `json:"low"`       // 最低
	Volume   int64   `json:"volume"`    // 成交量
	Amount   float64 `json:"amount"`    // 成交额
	PreClose float64 `json:"pre_close"` // 昨收
}

// ReitsHistMinEm 获取东方财富网-REITs-沪深REITs-分钟行情
//
// 目标地址: https://quote.eastmoney.com/sh508097.html
//
// 参数:
//   - symbol: REITs代码，如 "508097"
//
// 返回:
//   - []ReitsHistMinItem: REITs分钟行情数据列表
//   - error: 错误信息
//
// 示例:
//
//	data, err := reits.ReitsHistMinEm("508097")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range data {
//	    fmt.Printf("%s %.3f %.3f %.3f\n", item.Time, item.Price, item.High, item.Low)
//	}
func ReitsHistMinEm(symbol string) ([]ReitsHistMinItem, error) {
	// 确定市场代码
	secid := "1." + symbol // 默认上海
	if len(symbol) >= 1 && (symbol[0] == '1' || symbol[0] == '0') {
		secid = "0." + symbol // 深圳
	}

	url := "https://push2.eastmoney.com/api/qt/stock/trends2/get"
	params := map[string]string{
		"secid":   secid,
		"fields1": "f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f11,f12,f13,f14,f17",
		"fields2": "f51,f53,f54,f55,f56,f57,f58",
		"iscr":    "0",
		"iscca":   "0",
		"ndays":   "5",
	}

	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		"Referer":    "http://quote.eastmoney.com/",
	}

	resp, err := utils.GetWithHeaders(url, params, headers)
	if err != nil {
		return nil, fmt.Errorf("获取REITs分钟行情失败: %w", err)
	}

	text := resp.String()
	trends := gjson.Get(text, "data.trends")
	if !trends.Exists() {
		return nil, fmt.Errorf("解析REITs分钟数据失败")
	}

	var items []ReitsHistMinItem
	trends.ForEach(func(_, value gjson.Result) bool {
		parts := splitKline(value.String())
		if len(parts) >= 7 {
			item := ReitsHistMinItem{
				Time:     parts[0],
				Price:    utils.MustFloat64(parts[1]),
				High:     utils.MustFloat64(parts[2]),
				Low:      utils.MustFloat64(parts[3]),
				Volume:   utils.MustInt64(parts[4]),
				Amount:   utils.MustFloat64(parts[5]),
				PreClose: utils.MustFloat64(parts[6]),
			}
			items = append(items, item)
		}
		return true
	})

	return items, nil
}
