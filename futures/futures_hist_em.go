package futures

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
)

// FuturesHistEM 东方财富期货历史数据
type FuturesHistEM struct {
	Date      time.Time `json:"date"`       // 日期
	Open      float64   `json:"open"`       // 开盘
	High      float64   `json:"high"`       // 最高
	Low       float64   `json:"low"`        // 最低
	Close     float64   `json:"close"`      // 收盘
	Change    float64   `json:"change"`     // 涨跌
	ChangePct float64   `json:"change_pct"` // 涨跌幅
	Volume    int64     `json:"volume"`     // 成交量
	Amount    float64   `json:"amount"`     // 成交额
	Hold      int64     `json:"hold"`       // 持仓量
}

// FuturesHistTableEM 期货合约对照表
type FuturesHistTableEM struct {
	MarketName   string `json:"market_name"`   // 市场简称
	ContractName string `json:"contract_name"` // 合约中文代码
	ContractCode string `json:"contract_code"` // 合约代码
}

// 缓存交易所品种映射
var (
	exchangeSymbolCache     map[string]interface{}
	exchangeSymbolCacheLock sync.RWMutex
)

// emHistResponse 东方财富历史数据API响应
type emHistResponse struct {
	Data struct {
		Code   string   `json:"code"`
		Klines []string `json:"klines"`
	} `json:"data"`
}

// emRedisResponse 东方财富品种映射API响应
type emRedisItem struct {
	MktID   int    `json:"mktid"`
	MktName string `json:"mktname"`
	Name    string `json:"name"`
	Code    string `json:"code"`
	VCode   string `json:"vcode"`
	VName   string `json:"vname"`
}

// fetchExchangeSymbolRawEM 获取东方财富交易所品种对照表原始数据
func fetchExchangeSymbolRawEM() ([]emRedisItem, error) {
	exchangeSymbolCacheLock.RLock()
	if exchangeSymbolCache != nil {
		if items, ok := exchangeSymbolCache["items"].([]emRedisItem); ok {
			exchangeSymbolCacheLock.RUnlock()
			return items, nil
		}
	}
	exchangeSymbolCacheLock.RUnlock()

	url := "https://futsse-static.eastmoney.com/redis"

	// 获取市场列表
	resp, err := utils.Get(url, map[string]string{"msgid": "gnweb"})
	if err != nil {
		return nil, fmt.Errorf("获取市场列表失败: %w", err)
	}

	var markets []struct {
		MktID int `json:"mktid"`
	}
	if err := json.Unmarshal(resp.Body(), &markets); err != nil {
		return nil, fmt.Errorf("解析市场列表失败: %w", err)
	}

	var allItems []emRedisItem

	for _, market := range markets {
		// 获取每个市场的品种
		for num := 1; num <= 10; num++ {
			resp, err := utils.Get(url, map[string]string{
				"msgid": fmt.Sprintf("%d_%d", market.MktID, num),
			})
			if err != nil {
				break
			}

			var items []emRedisItem
			if err := json.Unmarshal(resp.Body(), &items); err != nil {
				break
			}
			if len(items) == 0 {
				break
			}
			allItems = append(allItems, items...)
		}
	}

	// 缓存结果
	exchangeSymbolCacheLock.Lock()
	exchangeSymbolCache = map[string]interface{}{
		"items": allItems,
	}
	exchangeSymbolCacheLock.Unlock()

	return allItems, nil
}

// FuturesHistTableEM 东方财富-期货行情-交易所品种对照表
//
// 数据源: https://quote.eastmoney.com/qihuo/al2505.html
//
// 返回:
//   - []FuturesHistTableEM: 品种对照表
//   - error: 错误信息
func FuturesHistTableEMFunc() ([]FuturesHistTableEM, error) {
	items, err := fetchExchangeSymbolRawEM()
	if err != nil {
		return nil, err
	}

	result := make([]FuturesHistTableEM, 0, len(items))
	for _, item := range items {
		result = append(result, FuturesHistTableEM{
			MarketName:   item.MktName,
			ContractName: item.Name,
			ContractCode: item.Code,
		})
	}
	return result, nil
}

// separateCharAndNumbers 分离字符和数字
func separateCharAndNumbers(symbol string) (string, string) {
	charRe := regexp.MustCompile(`[\x{4e00}-\x{9fa5}a-zA-Z]+`)
	numRe := regexp.MustCompile(`\d+`)

	chars := charRe.FindString(symbol)
	nums := numRe.FindString(symbol)

	return chars, nums
}

// getExchangeSymbolMap 获取交易所品种映射
func getExchangeSymbolMap() (cContractMkt, cContractToEContract, eSymbolMkt, cSymbolMkt map[string]int, err error) {
	items, err := fetchExchangeSymbolRawEM()
	if err != nil {
		return nil, nil, nil, nil, err
	}

	cContractMkt = make(map[string]int)
	cContractToEContract = make(map[string]int)
	eSymbolMkt = make(map[string]int)
	cSymbolMkt = make(map[string]int)

	for _, item := range items {
		cContractMkt[item.Name] = item.MktID
		eSymbolMkt[item.VCode] = item.MktID
		cSymbolMkt[item.VName] = item.MktID
	}

	return cContractMkt, cContractToEContract, eSymbolMkt, cSymbolMkt, nil
}

// FuturesHistEM 东方财富-期货行情-历史数据
//
// 数据源: https://qhweb.eastmoney.com/quote
//
// 参数:
//   - symbol: 期货代码，如 "热卷主连", "RB2505"
//   - period: 周期，可选 "daily", "weekly", "monthly"
//   - startDate: 开始日期，格式 "19900101"
//   - endDate: 结束日期，格式 "20500101"
//
// 返回:
//   - []FuturesHistEM: 历史行情数据
//   - error: 错误信息
func FuturesHistEMFunc(symbol, period, startDate, endDate string) ([]FuturesHistEM, error) {
	periodMap := map[string]string{
		"daily":   "101",
		"weekly":  "102",
		"monthly": "103",
	}

	klt, ok := periodMap[period]
	if !ok {
		return nil, fmt.Errorf("无效的 period: %s, 可选 daily, weekly, monthly", period)
	}

	items, err := fetchExchangeSymbolRawEM()
	if err != nil {
		return nil, err
	}

	// 构建映射
	cContractMkt := make(map[string]int)
	cContractToEContract := make(map[string]string)
	eSymbolMkt := make(map[string]int)
	cSymbolMkt := make(map[string]int)

	for _, item := range items {
		cContractMkt[item.Name] = item.MktID
		cContractToEContract[item.Name] = item.Code
		eSymbolMkt[item.VCode] = item.MktID
		cSymbolMkt[item.VName] = item.MktID
	}

	// 确定 secid
	var secid string
	if mktID, ok := cContractMkt[symbol]; ok {
		secid = fmt.Sprintf("%d.%s", mktID, cContractToEContract[symbol])
	} else {
		symbolChar, _ := separateCharAndNumbers(symbol)
		isChinese := regexp.MustCompile(`^[\x{4e00}-\x{9fa5}]+$`).MatchString(symbolChar)
		if isChinese {
			if mktID, ok := cSymbolMkt[symbolChar]; ok {
				secid = fmt.Sprintf("%d.%s", mktID, symbol)
			}
		} else {
			if mktID, ok := eSymbolMkt[symbolChar]; ok {
				secid = fmt.Sprintf("%d.%s", mktID, symbol)
			}
		}
	}

	if secid == "" {
		return nil, fmt.Errorf("无法找到品种映射: %s", symbol)
	}

	url := "https://push2his.eastmoney.com/api/qt/stock/kline/get"
	params := map[string]string{
		"secid":   secid,
		"klt":     klt,
		"fqt":     "1",
		"lmt":     "10000",
		"end":     "20500000",
		"iscca":   "1",
		"fields1": "f1,f2,f3,f4,f5,f6,f7,f8",
		"fields2": "f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61,f62,f63,f64",
		"ut":      "7eea3edcaed734bea9cbfc24409ed989",
		"forcect": "1",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求东方财富期货历史数据失败: %w", err)
	}

	var apiResp emHistResponse
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, fmt.Errorf("解析东方财富期货历史数据响应失败: %w", err)
	}

	if len(apiResp.Data.Klines) == 0 {
		return []FuturesHistEM{}, nil
	}

	// 解析日期范围
	var start, end time.Time
	if startDate != "" {
		start, _ = time.Parse("20060102", startDate)
	}
	if endDate != "" {
		end, _ = time.Parse("20060102", endDate)
	}

	result := make([]FuturesHistEM, 0, len(apiResp.Data.Klines))
	for _, line := range apiResp.Data.Klines {
		parts := strings.Split(line, ",")
		if len(parts) < 13 {
			continue
		}

		date, err := time.Parse("2006-01-02", parts[0])
		if err != nil {
			continue
		}

		// 日期范围过滤
		if !start.IsZero() && date.Before(start) {
			continue
		}
		if !end.IsZero() && date.After(end) {
			continue
		}

		result = append(result, FuturesHistEM{
			Date:      date,
			Open:      utils.MustParseFloat(parts[1]),
			Close:     utils.MustParseFloat(parts[2]),
			High:      utils.MustParseFloat(parts[3]),
			Low:       utils.MustParseFloat(parts[4]),
			Volume:    int64(utils.MustParseFloat(parts[5])),
			Amount:    utils.MustParseFloat(parts[6]),
			ChangePct: utils.MustParseFloat(parts[8]),
			Change:    utils.MustParseFloat(parts[9]),
			Hold:      int64(utils.MustParseFloat(parts[12])),
		})
	}

	return result, nil
}
