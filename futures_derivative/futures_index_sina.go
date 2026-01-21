package futures_derivative

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// ZhSubscribeExchangeSymbol 订阅指定交易所品种代码
// 参数: symbol 交易所代码，如 "dce", "czce", "shfe", "cffex", "gfex"
func ZhSubscribeExchangeSymbol(symbol string) ([]map[string]interface{}, error) {
	if symbol == "" {
		symbol = "dce"
	}

	url := "http://vip.stock.finance.sina.com.cn/quotes_service/view/js/qihuohangqing.js"

	resp, err := utils.Get(url, nil)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	text := resp.String()
	// 提取JSON部分
	startIdx := strings.Index(text, "{")
	endIdx := strings.LastIndex(text, "}")
	if startIdx == -1 || endIdx == -1 {
		return nil, fmt.Errorf("解析响应失败")
	}

	jsonText := text[startIdx : endIdx+1]
	result := gjson.Parse(jsonText)

	var items []gjson.Result
	switch symbol {
	case "czce":
		items = result.Get("czce").Array()
	case "dce":
		items = result.Get("dce").Array()
	case "shfe":
		items = result.Get("shfe").Array()
	case "cffex":
		items = result.Get("cffex").Array()
	case "gfex":
		items = result.Get("gfex").Array()
	default:
		return nil, fmt.Errorf("不支持的交易所代码: %s", symbol)
	}

	records := make([]map[string]interface{}, 0)
	for _, item := range items {
		itemStr := item.String()
		// 跳过交易所名称
		if strings.Contains(itemStr, "交易所") {
			continue
		}
		record := map[string]interface{}{
			"symbol": itemStr,
		}
		records = append(records, record)
	}

	return records, nil
}

// MatchMainContract 获取指定交易所所有可提供数据的合约
// 参数: symbol 交易所代码，如 "dce", "czce", "shfe", "cffex", "gfex"
func MatchMainContract(symbol string) ([]map[string]interface{}, error) {
	if symbol == "" {
		symbol = "shfe"
	}

	// 获取品种列表
	symbols, err := ZhSubscribeExchangeSymbol(symbol)
	if err != nil {
		return nil, fmt.Errorf("获取品种列表失败: %w", err)
	}

	url := "http://vip.stock.finance.sina.com.cn/quotes_service/api/json_v2.php/Market_Center.getHQFuturesData"
	records := make([]map[string]interface{}, 0)

	for _, item := range symbols {
		symbolCode := item["symbol"].(string)

		params := map[string]string{
			"page": "1",
			"num":  "5",
			"sort": "position",
			"asc":  "0",
			"node": symbolCode,
			"base": "futures",
		}

		resp, err := utils.Get(url, params)
		if err != nil {
			continue
		}

		var dataList []map[string]interface{}
		if err := json.Unmarshal(resp.Body(), &dataList); err != nil {
			continue
		}

		// 查找主力连续合约（包含"连续"且包含"0"）
		for _, data := range dataList {
			name, ok1 := data["name"].(string)
			symbol, ok2 := data["symbol"].(string)
			if !ok1 || !ok2 {
				continue
			}

			if strings.Contains(name, "连续") && strings.Contains(symbol, "0") {
				record := map[string]interface{}{
					"symbol": symbol,
					"name":   name,
					"trade":  data["trade"],
				}
				records = append(records, record)
				break
			}
		}
	}

	return records, nil
}

// FuturesDisplayMainSina 新浪主力连续合约品种一览表
func FuturesDisplayMainSina() ([]map[string]interface{}, error) {
	exchanges := []string{"dce", "czce", "shfe", "cffex", "gfex"}
	records := make([]map[string]interface{}, 0)

	for _, exchange := range exchanges {
		items, err := MatchMainContract(exchange)
		if err != nil {
			continue
		}
		records = append(records, items...)
	}

	return records, nil
}

// FuturesMainSina 新浪期货主力连续日数据
// 参数: symbol 品种代码，如 "V0"；startDate,endDate 日期范围，如 "20240101", "20240131"
func FuturesMainSina(symbol, startDate, endDate string) ([]map[string]interface{}, error) {
	if symbol == "" {
		symbol = "V0"
	}
	if startDate == "" {
		startDate = "19900101"
	}
	if endDate == "" {
		endDate = time.Now().Format("20060102")
	}

	tradeDate := "20210817"
	tradeDateFormatted := tradeDate[:4] + "_" + tradeDate[4:6] + "_" + tradeDate[6:]

	url := fmt.Sprintf("https://stock2.finance.sina.com.cn/futures/api/jsonp.php/var%%20_%s%s=/InnerFuturesNewService.getDailyKLine?symbol=%s&_=%s",
		symbol, tradeDateFormatted, symbol, tradeDateFormatted)

	resp, err := utils.Get(url, nil)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	text := resp.String()
	// 提取JSON部分
	startIdx := strings.Index(text, "([")
	endIdx := strings.LastIndex(text, "])")
	if startIdx == -1 || endIdx == -1 {
		return nil, fmt.Errorf("解析响应失败")
	}

	jsonText := text[startIdx+1 : endIdx+1]
	var dataList [][]interface{}
	if err := json.Unmarshal([]byte(jsonText), &dataList); err != nil {
		return nil, fmt.Errorf("解析JSON失败: %w", err)
	}

	// 解析日期范围
	startTime, _ := time.Parse("20060102", startDate)
	endTime, _ := time.Parse("20060102", endDate)

	records := make([]map[string]interface{}, 0)
	for _, item := range dataList {
		if len(item) < 8 {
			continue
		}

		// 解析日期
		dateStr, _ := item[0].(string)
		date, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			continue
		}

		// 日期过滤
		if date.Before(startTime) || date.After(endTime) {
			continue
		}

		record := map[string]interface{}{
			"日期":    dateStr,
			"开盘价":   utils.MustFloat64(fmt.Sprintf("%v", item[1])),
			"最高价":   utils.MustFloat64(fmt.Sprintf("%v", item[2])),
			"最低价":   utils.MustFloat64(fmt.Sprintf("%v", item[3])),
			"收盘价":   utils.MustFloat64(fmt.Sprintf("%v", item[4])),
			"成交量":   utils.MustFloat64(fmt.Sprintf("%v", item[5])),
			"持仓量":   utils.MustFloat64(fmt.Sprintf("%v", item[6])),
			"动态结算价": utils.MustFloat64(fmt.Sprintf("%v", item[7])),
		}
		records = append(records, record)
	}

	return records, nil
}
