package futures

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
)

// FuturesSymbolMark 期货品种代码映射
type FuturesSymbolMark struct {
	Exchange string `json:"exchange"` // 交易所
	Symbol   string `json:"symbol"`   // 品种名称
	Mark     string `json:"mark"`     // 品种代码
}

// FuturesZhRealtime 期货实时行情
type FuturesZhRealtime struct {
	Symbol        string  `json:"symbol"`         // 合约代码
	Name          string  `json:"name"`           // 合约名称
	Trade         float64 `json:"trade"`          // 最新价
	Settlement    float64 `json:"settlement"`     // 结算价
	PreSettlement float64 `json:"pre_settlement"` // 昨结算
	Open          float64 `json:"open"`           // 开盘价
	High          float64 `json:"high"`           // 最高价
	Low           float64 `json:"low"`            // 最低价
	Close         float64 `json:"close"`          // 收盘价
	BidPrice1     float64 `json:"bid_price1"`     // 买一价
	AskPrice1     float64 `json:"ask_price1"`     // 卖一价
	BidVol1       int64   `json:"bid_vol1"`       // 买一量
	AskVol1       int64   `json:"ask_vol1"`       // 卖一量
	Volume        int64   `json:"volume"`         // 成交量
	Position      int64   `json:"position"`       // 持仓量
	PreClose      float64 `json:"pre_close"`      // 昨收盘
	ChangePct     float64 `json:"change_pct"`     // 涨跌幅
}

// FuturesZhSpot 期货实时行情(spot)
type FuturesZhSpot struct {
	Symbol          string  `json:"symbol"`            // 品种
	Time            string  `json:"time"`              // 时间
	Open            float64 `json:"open"`              // 开盘
	High            float64 `json:"high"`              // 最高
	Low             float64 `json:"low"`               // 最低
	CurrentPrice    float64 `json:"current_price"`     // 最新价
	BidPrice        float64 `json:"bid_price"`         // 买价
	AskPrice        float64 `json:"ask_price"`         // 卖价
	BuyVol          int64   `json:"buy_vol"`           // 买量
	SellVol         int64   `json:"sell_vol"`          // 卖量
	Hold            int64   `json:"hold"`              // 持仓
	Volume          int64   `json:"volume"`            // 成交量
	AvgPrice        float64 `json:"avg_price"`         // 均价
	LastClose       float64 `json:"last_close"`        // 昨收
	LastSettlePrice float64 `json:"last_settle_price"` // 昨结算
}

// FuturesZhMinute 期货分钟数据
type FuturesZhMinute struct {
	Datetime string  `json:"datetime"` // 日期时间
	Open     float64 `json:"open"`     // 开盘
	High     float64 `json:"high"`     // 最高
	Low      float64 `json:"low"`      // 最低
	Close    float64 `json:"close"`    // 收盘
	Volume   int64   `json:"volume"`   // 成交量
	Hold     int64   `json:"hold"`     // 持仓量
}

// FuturesZhDaily 期货日线数据
type FuturesZhDaily struct {
	Date   time.Time `json:"date"`   // 日期
	Open   float64   `json:"open"`   // 开盘
	High   float64   `json:"high"`   // 最高
	Low    float64   `json:"low"`    // 最低
	Close  float64   `json:"close"`  // 收盘
	Volume int64     `json:"volume"` // 成交量
	Hold   int64     `json:"hold"`   // 持仓量
	Settle float64   `json:"settle"` // 结算价
}

// FuturesSymbolMarkFunc 获取期货品种代码映射
//
// 数据源: https://vip.stock.finance.sina.com.cn/quotes_service/view/js/qihuohangqing.js
//
// 返回:
//   - []FuturesSymbolMark: 期货品种代码映射
//   - error: 错误信息
func FuturesSymbolMarkFunc() ([]FuturesSymbolMark, error) {
	url := "https://vip.stock.finance.sina.com.cn/quotes_service/view/js/qihuohangqing.js"

	resp, err := utils.Get(url, nil)
	if err != nil {
		return nil, fmt.Errorf("请求期货品种映射失败: %w", err)
	}

	text := string(resp.Body())

	// 提取JSON部分
	startIdx := strings.Index(text, "{")
	endIdx := strings.Index(text, "}") + 1
	if startIdx == -1 || endIdx <= startIdx {
		return nil, fmt.Errorf("解析期货品种映射失败: 无法找到JSON数据")
	}

	jsonStr := text[startIdx:endIdx]

	// 使用更宽松的解析方式
	var dataMap map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &dataMap); err != nil {
		// 尝试修复JavaScript对象格式
		jsonStr = fixJSObject(jsonStr)
		if err := json.Unmarshal([]byte(jsonStr), &dataMap); err != nil {
			return nil, fmt.Errorf("解析期货品种映射JSON失败: %w", err)
		}
	}

	var result []FuturesSymbolMark

	exchanges := []string{"czce", "dce", "shfe", "cffex", "gfex"}
	exchangeNames := map[string]string{
		"czce":  "郑州商品交易所",
		"dce":   "大连商品交易所",
		"shfe":  "上海期货交易所",
		"cffex": "中国金融期货交易所",
		"gfex":  "广州期货交易所",
	}

	for _, exchange := range exchanges {
		if data, ok := dataMap[exchange]; ok {
			if items, ok := data.([]interface{}); ok {
				exchangeName := exchangeNames[exchange]
				for i, item := range items {
					if i == 0 {
						continue // 跳过第一个元素（交易所名称）
					}
					if arr, ok := item.([]interface{}); ok && len(arr) >= 2 {
						symbol := fmt.Sprintf("%v", arr[0])
						mark := fmt.Sprintf("%v", arr[1])
						result = append(result, FuturesSymbolMark{
							Exchange: exchangeName,
							Symbol:   symbol,
							Mark:     mark,
						})
					}
				}
			}
		}
	}

	return result, nil
}

// fixJSObject 修复JavaScript对象格式为JSON
func fixJSObject(js string) string {
	// 简单的修复：给键名加引号
	re := regexp.MustCompile(`(\w+):`)
	return re.ReplaceAllString(js, `"$1":`)
}

// FuturesZhRealtimeFunc 获取期货品种当前时刻所有可交易的合约实时数据
//
// 数据源: https://vip.stock.finance.sina.com.cn/quotes_service/view/qihuohangqing.html
//
// 参数:
//   - symbol: 品种名称，如 "PTA", "螺纹钢" 等
//
// 返回:
//   - []FuturesZhRealtime: 期货实时行情
//   - error: 错误信息
func FuturesZhRealtimeFunc(symbol string) ([]FuturesZhRealtime, error) {
	marks, err := FuturesSymbolMarkFunc()
	if err != nil {
		return nil, err
	}

	// 构建映射
	symbolMarkMap := make(map[string]string)
	for _, m := range marks {
		symbolMarkMap[m.Symbol] = m.Mark
	}

	mark, ok := symbolMarkMap[symbol]
	if !ok {
		return nil, fmt.Errorf("无效的品种: %s", symbol)
	}

	url := "https://vip.stock.finance.sina.com.cn/quotes_service/api/json_v2.php/Market_Center.getHQFuturesData"
	params := map[string]string{
		"page": "1",
		"sort": "position",
		"asc":  "0",
		"node": mark,
		"base": "futures",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求期货实时行情失败: %w", err)
	}

	var data []map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("解析期货实时行情失败: %w", err)
	}

	result := make([]FuturesZhRealtime, 0, len(data))
	for _, item := range data {
		result = append(result, FuturesZhRealtime{
			Symbol:        getString(item, "symbol"),
			Name:          getString(item, "name"),
			Trade:         getFloat(item, "trade"),
			Settlement:    getFloat(item, "settlement"),
			PreSettlement: getFloat(item, "presettlement"),
			Open:          getFloat(item, "open"),
			High:          getFloat(item, "high"),
			Low:           getFloat(item, "low"),
			Close:         getFloat(item, "close"),
			BidPrice1:     getFloat(item, "bidprice1"),
			AskPrice1:     getFloat(item, "askprice1"),
			BidVol1:       int64(getFloat(item, "bidvol1")),
			AskVol1:       int64(getFloat(item, "askvol1")),
			Volume:        int64(getFloat(item, "volume")),
			Position:      int64(getFloat(item, "position")),
			PreClose:      getFloat(item, "preclose"),
			ChangePct:     getFloat(item, "changepercent"),
		})
	}

	return result, nil
}

// FuturesZhSpotFunc 获取期货实时行情
//
// 数据源: https://vip.stock.finance.sina.com.cn/quotes_service/view/qihuohangqing.html
//
// 参数:
//   - symbol: 合约代码，多个用逗号分隔，如 "V2309,V2401"
//   - market: 市场类型，"CF" 为商品期货
//
// 返回:
//   - []FuturesZhSpot: 期货实时行情
//   - error: 错误信息
func FuturesZhSpotFunc(symbol, market string) ([]FuturesZhSpot, error) {
	rnCode := fmt.Sprintf("%x", rand.Int63())
	symbols := strings.Split(symbol, ",")
	subscribeList := make([]string, 0, len(symbols))
	for _, s := range symbols {
		subscribeList = append(subscribeList, "nf_"+strings.TrimSpace(s))
	}

	url := fmt.Sprintf("https://hq.sinajs.cn/rn=%s&list=%s", rnCode, strings.Join(subscribeList, ","))

	headers := map[string]string{
		"Accept":           "*/*",
		"Accept-Encoding":  "gzip, deflate",
		"Accept-Language":  "zh-CN,zh;q=0.9,en;q=0.8",
		"Cache-Control":    "no-cache",
		"Host":             "hq.sinajs.cn",
		"Pragma":           "no-cache",
		"Proxy-Connection": "keep-alive",
		"Referer":          "https://vip.stock.finance.sina.com.cn/",
		"User-Agent":       "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36",
	}

	resp, err := utils.GetWithHeaders(url, nil, headers)
	if err != nil {
		return nil, fmt.Errorf("请求期货实时行情失败: %w", err)
	}

	text := string(resp.Body())
	lines := strings.Split(text, ";")

	var result []FuturesZhSpot
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
		if len(fields) < 15 {
			continue
		}

		if market == "CF" {
			result = append(result, FuturesZhSpot{
				Symbol:          strings.Trim(fields[0], "\""),
				Time:            fields[1],
				Open:            utils.MustParseFloat(fields[2]),
				High:            utils.MustParseFloat(fields[3]),
				Low:             utils.MustParseFloat(fields[4]),
				LastClose:       utils.MustParseFloat(fields[5]),
				BidPrice:        utils.MustParseFloat(fields[6]),
				AskPrice:        utils.MustParseFloat(fields[7]),
				CurrentPrice:    utils.MustParseFloat(fields[8]),
				AvgPrice:        utils.MustParseFloat(fields[9]),
				LastSettlePrice: utils.MustParseFloat(fields[10]),
				BuyVol:          int64(utils.MustParseFloat(fields[11])),
				SellVol:         int64(utils.MustParseFloat(fields[12])),
				Hold:            int64(utils.MustParseFloat(fields[13])),
				Volume:          int64(utils.MustParseFloat(fields[14])),
			})
		}
	}

	return result, nil
}

// FuturesZhMinuteSina 获取期货分钟数据
//
// 数据源: https://vip.stock.finance.sina.com.cn/quotes_service/view/qihuohangqing.html
//
// 参数:
//   - symbol: 合约代码，如 "IF2008", "RB0"
//   - period: 周期，可选 "1", "5", "15", "30", "60"
//
// 返回:
//   - []FuturesZhMinute: 期货分钟数据
//   - error: 错误信息
func FuturesZhMinuteSina(symbol, period string) ([]FuturesZhMinute, error) {
	url := "https://stock2.finance.sina.com.cn/futures/api/jsonp.php/=/InnerFuturesNewService.getFewMinLine"
	params := map[string]string{
		"symbol": symbol,
		"type":   period,
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求期货分钟数据失败: %w", err)
	}

	text := string(resp.Body())
	// 提取JSON部分
	startIdx := strings.Index(text, "([")
	endIdx := strings.LastIndex(text, "])")
	if startIdx == -1 || endIdx == -1 {
		return nil, fmt.Errorf("解析期货分钟数据失败: 无法找到数据")
	}

	jsonStr := text[startIdx+1 : endIdx+1]

	var data [][]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
		return nil, fmt.Errorf("解析期货分钟数据JSON失败: %w", err)
	}

	result := make([]FuturesZhMinute, 0, len(data))
	for _, item := range data {
		if len(item) < 7 {
			continue
		}
		result = append(result, FuturesZhMinute{
			Datetime: fmt.Sprintf("%v", item[0]),
			Open:     toFloat64(item[1]),
			High:     toFloat64(item[2]),
			Low:      toFloat64(item[3]),
			Close:    toFloat64(item[4]),
			Volume:   int64(toFloat64(item[5])),
			Hold:     int64(toFloat64(item[6])),
		})
	}

	return result, nil
}

// FuturesZhDailySina 获取期货日线数据
//
// 数据源: https://finance.sina.com.cn/futures/quotes/V2105.shtml
//
// 参数:
//   - symbol: 合约代码，如 "RB0", "RB2410"
//
// 返回:
//   - []FuturesZhDaily: 期货日线数据
//   - error: 错误信息
func FuturesZhDailySina(symbol string) ([]FuturesZhDaily, error) {
	date := "20210412"
	url := "https://stock2.finance.sina.com.cn/futures/api/jsonp.php/var%20_V21052021_4_12=/InnerFuturesNewService.getDailyKLine"
	params := map[string]string{
		"symbol": symbol,
		"type":   fmt.Sprintf("%s_%s_%s", date[:4], date[4:6], date[6:]),
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求期货日线数据失败: %w", err)
	}

	text := string(resp.Body())
	// 提取JSON部分
	startIdx := strings.Index(text, "([")
	endIdx := strings.LastIndex(text, "])")
	if startIdx == -1 || endIdx == -1 {
		return nil, fmt.Errorf("解析期货日线数据失败: 无法找到数据")
	}

	jsonStr := text[startIdx+1 : endIdx+1]

	var data [][]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
		return nil, fmt.Errorf("解析期货日线数据JSON失败: %w", err)
	}

	result := make([]FuturesZhDaily, 0, len(data))
	for _, item := range data {
		if len(item) < 8 {
			continue
		}
		dateStr := fmt.Sprintf("%v", item[0])
		date, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			continue
		}
		result = append(result, FuturesZhDaily{
			Date:   date,
			Open:   toFloat64(item[1]),
			High:   toFloat64(item[2]),
			Low:    toFloat64(item[3]),
			Close:  toFloat64(item[4]),
			Volume: int64(toFloat64(item[5])),
			Hold:   int64(toFloat64(item[6])),
			Settle: toFloat64(item[7]),
		})
	}

	return result, nil
}

// 辅助函数
func getString(m map[string]interface{}, key string) string {
	if v, ok := m[key]; ok {
		return fmt.Sprintf("%v", v)
	}
	return ""
}

func getFloat(m map[string]interface{}, key string) float64 {
	if v, ok := m[key]; ok {
		return toFloat64(v)
	}
	return 0
}

func toFloat64(v interface{}) float64 {
	switch val := v.(type) {
	case float64:
		return val
	case float32:
		return float64(val)
	case int:
		return float64(val)
	case int64:
		return float64(val)
	case string:
		return utils.MustParseFloat(val)
	default:
		return 0
	}
}
