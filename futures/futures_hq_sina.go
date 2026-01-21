package futures

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
)

// FuturesHQSubscribeSymbol 外盘期货品种对照表
type FuturesHQSubscribeSymbol struct {
	Symbol string `json:"symbol"` // 品种名称
	Code   string `json:"code"`   // 代码
}

// FuturesForeignCommodityRealtime 外盘期货实时行情
type FuturesForeignCommodityRealtime struct {
	Name            string  `json:"name"`              // 名称
	CurrentPrice    float64 `json:"current_price"`     // 最新价
	CurrentPriceRMB float64 `json:"current_price_rmb"` // 人民币报价
	Change          float64 `json:"change"`            // 涨跌额
	ChangePct       float64 `json:"change_pct"`        // 涨跌幅
	Open            float64 `json:"open"`              // 开盘价
	High            float64 `json:"high"`              // 最高价
	Low             float64 `json:"low"`               // 最低价
	LastSettle      float64 `json:"last_settle"`       // 昨日结算价
	Hold            int64   `json:"hold"`              // 持仓量
	Bid             float64 `json:"bid"`               // 买价
	Ask             float64 `json:"ask"`               // 卖价
	Time            string  `json:"time"`              // 行情时间
	Date            string  `json:"date"`              // 日期
}

// 外盘期货品种映射
var foreignCommoditySymbolMap = map[string]string{
	"新加坡铁矿石":   "FEF",
	"马棕油":      "FCPO",
	"日橡胶":      "RSS3",
	"美国原糖":     "RS",
	"CME比特币期货": "BTC",
	"NYBOT-棉花": "CT",
	"LME镍3个月":  "NID",
	"LME铅3个月":  "PBD",
	"LME锡3个月":  "SND",
	"LME锌3个月":  "ZSD",
	"LME铝3个月":  "AHD",
	"LME铜3个月":  "CAD",
	"CBOT-黄豆":  "S",
	"CBOT-小麦":  "W",
	"CBOT-玉米":  "C",
	"CBOT-黄豆油": "BO",
	"CBOT-黄豆粉": "SM",
	"日本橡胶":     "TRB",
	"COMEX铜":   "HG",
	"NYMEX天然气": "NG",
	"NYMEX原油":  "CL",
	"COMEX白银":  "SI",
	"COMEX黄金":  "GC",
	"CME-瘦肉猪":  "LHC",
	"布伦特原油":    "OIL",
	"伦敦金":      "XAU",
	"伦敦银":      "XAG",
	"伦敦铂金":     "XPT",
	"伦敦钯金":     "XPD",
	"欧洲碳排放":    "EUA",
}

// 反向映射
var foreignCommodityCodeMap map[string]string

func init() {
	foreignCommodityCodeMap = make(map[string]string)
	for name, code := range foreignCommoditySymbolMap {
		foreignCommodityCodeMap[code] = name
	}
}

// FuturesHQSubscribeExchangeSymbol 获取外盘期货品种对照表
//
// 数据源: https://finance.sina.com.cn/money/future/hf.html
//
// 返回:
//   - []FuturesHQSubscribeSymbol: 品种对照表
//   - error: 错误信息
func FuturesHQSubscribeExchangeSymbol() ([]FuturesHQSubscribeSymbol, error) {
	result := make([]FuturesHQSubscribeSymbol, 0, len(foreignCommoditySymbolMap))
	for name, code := range foreignCommoditySymbolMap {
		result = append(result, FuturesHQSubscribeSymbol{
			Symbol: name,
			Code:   code,
		})
	}
	return result, nil
}

// FuturesForeignCommoditySubscribeExchangeSymbol 获取需要订阅的外盘期货代码列表
//
// 数据源: https://finance.sina.com.cn/money/future/hf.html
//
// 返回:
//   - []string: 代码列表
//   - error: 错误信息
func FuturesForeignCommoditySubscribeExchangeSymbol() ([]string, error) {
	url := "https://finance.sina.com.cn/money/future/hf.html"

	resp, err := utils.Get(url, nil)
	if err != nil {
		return nil, fmt.Errorf("请求外盘期货代码列表失败: %w", err)
	}

	text := string(resp.Body())
	// 查找 oHF_1 变量
	startIdx := strings.Index(text, "var oHF_1 = ")
	if startIdx == -1 {
		// 如果无法从网页获取，返回内置的代码列表
		codes := make([]string, 0, len(foreignCommoditySymbolMap))
		for _, code := range foreignCommoditySymbolMap {
			codes = append(codes, code)
		}
		return codes, nil
	}

	endIdx := strings.Index(text[startIdx:], "var oHF_2")
	if endIdx == -1 {
		endIdx = len(text) - startIdx
	}

	jsonStr := strings.TrimSpace(text[startIdx+12 : startIdx+endIdx-2])
	jsonStr = strings.ReplaceAll(jsonStr, "\n\t", "")

	var dataMap map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &dataMap); err != nil {
		// 返回内置的代码列表
		codes := make([]string, 0, len(foreignCommoditySymbolMap))
		for _, code := range foreignCommoditySymbolMap {
			codes = append(codes, code)
		}
		return codes, nil
	}

	codes := make([]string, 0, len(dataMap))
	for code := range dataMap {
		codes = append(codes, code)
	}
	return codes, nil
}

// FuturesForeignCommodityRealtime 获取外盘期货实时行情
//
// 数据源: https://finance.sina.com.cn/money/future/hf.html
//
// 参数:
//   - symbol: 品种代码，可以是单个代码或用逗号分隔的多个代码，如 "CT,NID" 或 ["XAU"]
//
// 返回:
//   - []FuturesForeignCommodityRealtime: 外盘期货实时行情
//   - error: 错误信息
func FuturesForeignCommodityRealtimeFunc(symbol string) ([]FuturesForeignCommodityRealtime, error) {
	symbols := strings.Split(symbol, ",")
	payload := "?list=" + strings.Join(func() []string {
		result := make([]string, len(symbols))
		for i, s := range symbols {
			result[i] = "hf_" + strings.TrimSpace(s)
		}
		return result
	}(), ",")

	url := "https://hq.sinajs.cn/" + payload

	headers := map[string]string{
		"Accept":          "*/*",
		"Accept-Encoding": "gzip, deflate, br",
		"Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8",
		"Cache-Control":   "no-cache",
		"Connection":      "keep-alive",
		"Host":            "hq.sinajs.cn",
		"Pragma":          "no-cache",
		"Referer":         "https://finance.sina.com.cn/",
		"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36",
	}

	resp, err := utils.GetWithHeaders(url, nil, headers)
	if err != nil {
		return nil, fmt.Errorf("请求外盘期货实时行情失败: %w", err)
	}

	text := string(resp.Body())
	lines := strings.Split(text, ";")

	var result []FuturesForeignCommodityRealtime
	symbolIdx := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.Split(line, "=")
		if len(parts) != 2 {
			continue
		}

		data := strings.Trim(parts[1], "\"")
		fields := strings.Split(data, ",")
		if len(fields) < 14 {
			continue
		}

		currentPrice := utils.MustParseFloat(strings.Trim(fields[0], "\""))
		lastSettle := utils.MustParseFloat(fields[7])

		symbolCode := ""
		if symbolIdx < len(symbols) {
			symbolCode = strings.TrimSpace(symbols[symbolIdx])
			symbolIdx++
		}

		name := foreignCommodityCodeMap[symbolCode]
		if name == "" {
			name = symbolCode
		}

		change := currentPrice - lastSettle
		changePct := 0.0
		if lastSettle != 0 {
			changePct = (currentPrice - lastSettle) / lastSettle * 100
		}

		result = append(result, FuturesForeignCommodityRealtime{
			Name:            name,
			CurrentPrice:    currentPrice,
			CurrentPriceRMB: utils.MustParseFloat(strings.Trim(fields[len(fields)-1], "\"")),
			Change:          change,
			ChangePct:       changePct,
			Open:            utils.MustParseFloat(fields[8]),
			High:            utils.MustParseFloat(fields[4]),
			Low:             utils.MustParseFloat(fields[5]),
			LastSettle:      lastSettle,
			Hold:            int64(utils.MustParseFloat(fields[9])),
			Bid:             utils.MustParseFloat(fields[2]),
			Ask:             utils.MustParseFloat(fields[3]),
			Time:            fields[6],
			Date:            fields[12],
		})
	}

	return result, nil
}
