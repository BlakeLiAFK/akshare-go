package index

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
)

// StockZhIndexSpot 股票指数实时行情
type StockZhIndexSpot struct {
	Seq       int     `json:"seq"`        // 序号
	Code      string  `json:"code"`       // 代码
	Name      string  `json:"name"`       // 名称
	Price     float64 `json:"price"`      // 最新价
	ChangePct float64 `json:"change_pct"` // 涨跌幅
	Change    float64 `json:"change"`     // 涨跌额
	Volume    int64   `json:"volume"`     // 成交量
	Amount    float64 `json:"amount"`     // 成交额
	Amplitude float64 `json:"amplitude"`  // 振幅
	High      float64 `json:"high"`       // 最高
	Low       float64 `json:"low"`        // 最低
	Open      float64 `json:"open"`       // 今开
	PreClose  float64 `json:"pre_close"`  // 昨收
	VolumeRat float64 `json:"volume_rat"` // 量比
}

// StockZhIndexDaily 股票指数日线数据
type StockZhIndexDaily struct {
	Date   time.Time `json:"date"`   // 日期
	Open   float64   `json:"open"`   // 开盘
	Close  float64   `json:"close"`  // 收盘
	High   float64   `json:"high"`   // 最高
	Low    float64   `json:"low"`    // 最低
	Volume int64     `json:"volume"` // 成交量
	Amount float64   `json:"amount"` // 成交额
}

// emIndexSpotResponse 东方财富指数实时行情API响应
type emIndexSpotResponse struct {
	Data struct {
		Total int `json:"total"`
		Diff  []struct {
			F1  interface{} `json:"f1"`
			F2  interface{} `json:"f2"`  // 最新价
			F3  interface{} `json:"f3"`  // 涨跌幅
			F4  interface{} `json:"f4"`  // 涨跌额
			F5  interface{} `json:"f5"`  // 成交量
			F6  interface{} `json:"f6"`  // 成交额
			F7  interface{} `json:"f7"`  // 振幅
			F10 interface{} `json:"f10"` // 量比
			F12 string      `json:"f12"` // 代码
			F14 string      `json:"f14"` // 名称
			F15 interface{} `json:"f15"` // 最高
			F16 interface{} `json:"f16"` // 最低
			F17 interface{} `json:"f17"` // 今开
			F18 interface{} `json:"f18"` // 昨收
		} `json:"diff"`
	} `json:"data"`
}

// emIndexDailyResponse 东方财富指数日线API响应
type emIndexDailyResponse struct {
	Data struct {
		Code   string   `json:"code"`
		Klines []string `json:"klines"`
	} `json:"data"`
}

// StockZhIndexSpotEM 东方财富-沪深京指数实时行情
//
// 数据源: https://quote.eastmoney.com/center/hszs.html
//
// 参数:
//   - symbol: 指数类型，可选:
//     "沪深重要指数", "上证系列指数", "深证系列指数", "指数成份", "中证系列指数"
//
// 返回:
//   - []StockZhIndexSpot: 指数实时行情数据
//   - error: 错误信息
func StockZhIndexSpotEM(symbol string) ([]StockZhIndexSpot, error) {
	if symbol == "沪深重要指数" {
		return stockZhMainSpotEM()
	}

	symbolMap := map[string]string{
		"上证系列指数": "m:1+t:1",
		"深证系列指数": "m:0 t:5",
		"指数成份":   "m:1+s:3,m:0+t:5",
		"中证系列指数": "m:2",
	}

	fs, ok := symbolMap[symbol]
	if !ok {
		return nil, fmt.Errorf("无效的 symbol: %s", symbol)
	}

	apiURL := "https://48.push2.eastmoney.com/api/qt/clist/get"

	result := make([]StockZhIndexSpot, 0)
	page := 1

	for {
		params := map[string]string{
			"pn":     fmt.Sprintf("%d", page),
			"pz":     "100",
			"po":     "1",
			"np":     "1",
			"ut":     "bd1d9ddb04089700cf9c27f6f7426281",
			"fltt":   "2",
			"invt":   "2",
			"wbp2u":  "|0|0|0|web",
			"fid":    "f12",
			"fs":     fs,
			"fields": "f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18,f20,f21,f23,f24,f25,f26,f22,f33,f11,f62,f128,f136,f115,f152",
		}

		resp, err := utils.Get(apiURL, params)
		if err != nil {
			return nil, fmt.Errorf("请求东方财富指数实时行情失败: %w", err)
		}

		var apiResp emIndexSpotResponse
		if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
			return nil, fmt.Errorf("解析东方财富指数实时行情响应失败: %w", err)
		}

		if len(apiResp.Data.Diff) == 0 {
			break
		}

		for i, item := range apiResp.Data.Diff {
			result = append(result, StockZhIndexSpot{
				Seq:       (page-1)*100 + i + 1,
				Code:      item.F12,
				Name:      item.F14,
				Price:     toFloat64(item.F2),
				ChangePct: toFloat64(item.F3),
				Change:    toFloat64(item.F4),
				Volume:    int64(toFloat64(item.F5)),
				Amount:    toFloat64(item.F6),
				Amplitude: toFloat64(item.F7),
				High:      toFloat64(item.F15),
				Low:       toFloat64(item.F16),
				Open:      toFloat64(item.F17),
				PreClose:  toFloat64(item.F18),
				VolumeRat: toFloat64(item.F10),
			})
		}

		if len(apiResp.Data.Diff) < 100 {
			break
		}
		page++
	}

	return result, nil
}

// stockZhMainSpotEM 东方财富-沪深重要指数实时行情
func stockZhMainSpotEM() ([]StockZhIndexSpot, error) {
	apiURL := "https://33.push2.eastmoney.com/api/qt/clist/get"

	params := map[string]string{
		"pn":     "1",
		"pz":     "100",
		"po":     "1",
		"np":     "1",
		"ut":     "bd1d9ddb04089700cf9c27f6f7426281",
		"fltt":   "2",
		"invt":   "2",
		"dect":   "1",
		"wbp2u":  "|0|0|0|web",
		"fid":    "",
		"fs":     "b:MK0010",
		"fields": "f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18,f20,f21,f23,f24,f25,f26,f22,f11,f62,f128,f136,f115,f152",
	}

	resp, err := utils.Get(apiURL, params)
	if err != nil {
		return nil, fmt.Errorf("请求东方财富沪深重要指数实时行情失败: %w", err)
	}

	var apiResp emIndexSpotResponse
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, fmt.Errorf("解析东方财富沪深重要指数实时行情响应失败: %w", err)
	}

	result := make([]StockZhIndexSpot, 0, len(apiResp.Data.Diff))
	for i, item := range apiResp.Data.Diff {
		result = append(result, StockZhIndexSpot{
			Seq:       i + 1,
			Code:      item.F12,
			Name:      item.F14,
			Price:     toFloat64(item.F2),
			ChangePct: toFloat64(item.F3),
			Change:    toFloat64(item.F4),
			Volume:    int64(toFloat64(item.F5)),
			Amount:    toFloat64(item.F6),
			Amplitude: toFloat64(item.F7),
			High:      toFloat64(item.F15),
			Low:       toFloat64(item.F16),
			Open:      toFloat64(item.F17),
			PreClose:  toFloat64(item.F18),
			VolumeRat: toFloat64(item.F10),
		})
	}

	return result, nil
}

// StockZhIndexDailyEM 东方财富-股票指数日线数据
//
// 数据源: https://quote.eastmoney.com/center/hszs.html
//
// 参数:
//   - symbol: 带市场标识的指数代码
//     sz: 深交所, sh: 上交所, csi: 中证指数, bj: 北交所
//     例如: "sz399001", "sh000001", "csi931151"
//   - startDate: 开始日期，格式 "19900101"
//   - endDate: 结束日期，格式 "20500101"
//
// 返回:
//   - []StockZhIndexDaily: 指数日线数据
//   - error: 错误信息
func StockZhIndexDailyEM(symbol, startDate, endDate string) ([]StockZhIndexDaily, error) {
	marketMap := map[string]string{
		"sz":  "0",
		"sh":  "1",
		"csi": "2",
		"bj":  "0",
	}

	var secid string
	for prefix, marketID := range marketMap {
		if strings.HasPrefix(symbol, prefix) {
			code := strings.TrimPrefix(symbol, prefix)
			secid = fmt.Sprintf("%s.%s", marketID, code)
			break
		}
	}

	if secid == "" {
		return nil, fmt.Errorf("无效的 symbol: %s，需要带市场前缀（sz/sh/csi/bj）", symbol)
	}

	apiURL := "https://push2his.eastmoney.com/api/qt/stock/kline/get"

	params := map[string]string{
		"secid":   secid,
		"fields1": "f1,f2,f3,f4,f5",
		"fields2": "f51,f52,f53,f54,f55,f56,f57,f58",
		"klt":     "101", // 日频
		"fqt":     "0",
		"beg":     startDate,
		"end":     endDate,
	}

	resp, err := utils.Get(apiURL, params)
	if err != nil {
		return nil, fmt.Errorf("请求东方财富指数日线数据失败: %w", err)
	}

	var apiResp emIndexDailyResponse
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, fmt.Errorf("解析东方财富指数日线数据响应失败: %w", err)
	}

	if len(apiResp.Data.Klines) == 0 {
		return []StockZhIndexDaily{}, nil
	}

	result := make([]StockZhIndexDaily, 0, len(apiResp.Data.Klines))
	for _, line := range apiResp.Data.Klines {
		parts := strings.Split(line, ",")
		if len(parts) < 7 {
			continue
		}

		date, err := time.Parse("2006-01-02", parts[0])
		if err != nil {
			continue
		}

		result = append(result, StockZhIndexDaily{
			Date:   date,
			Open:   utils.MustParseFloat(parts[1]),
			Close:  utils.MustParseFloat(parts[2]),
			High:   utils.MustParseFloat(parts[3]),
			Low:    utils.MustParseFloat(parts[4]),
			Volume: int64(utils.MustParseFloat(parts[5])),
			Amount: utils.MustParseFloat(parts[6]),
		})
	}

	return result, nil
}
