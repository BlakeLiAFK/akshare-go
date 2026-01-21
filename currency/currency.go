package currency

import (
	"fmt"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// CurrencyLatest 获取最新汇率数据
//
// 从 currencyscoop.com 获取最新汇率数据（需要API Key）
//
// 参数:
//   - base: 基准货币代码，如 "USD"
//   - symbols: 目标货币代码列表，多个用逗号分隔，如 "CNY,EUR,GBP"，空字符串表示所有货币
//   - apiKey: CurrencyScoop API Key（在 https://currencyscoop.com/ 注册获取）
//
// 返回:
//   - []CurrencyLatestItem: 最新汇率数据列表
//   - error: 错误信息
//
// 示例:
//
//	rates, err := currency.CurrencyLatest("USD", "CNY,EUR", "your-api-key")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range rates {
//	    fmt.Printf("%s: %.4f (日期=%s)\n", item.Currency, item.Value, item.Date)
//	}
func CurrencyLatest(base, symbols, apiKey string) ([]CurrencyLatestItem, error) {
	url := "https://api.currencyscoop.com/v1/latest"

	params := map[string]string{
		"base":    base,
		"symbols": symbols,
		"api_key": apiKey,
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求最新汇率数据失败: %w", err)
	}

	json := gjson.ParseBytes(resp.Body())

	// 检查API错误
	if !json.Get("meta.code").Exists() || json.Get("meta.code").Int() != 200 {
		errMsg := json.Get("meta.error_detail").String()
		if errMsg == "" {
			errMsg = "未知错误"
		}
		return nil, fmt.Errorf("API返回错误: %s", errMsg)
	}

	response := json.Get("response")
	dateStr := response.Get("date").String()
	date, _ := time.Parse("2006-01-02", dateStr)

	var items []CurrencyLatestItem
	response.Get("rates").ForEach(func(key, value gjson.Result) bool {
		items = append(items, CurrencyLatestItem{
			Currency: key.String(),
			Date:     date,
			Value:    value.Float(),
		})
		return true
	})

	return items, nil
}

// CurrencyHistory 获取历史汇率数据
//
// 从 currencyscoop.com 获取指定日期的历史汇率数据（需要API Key）
//
// 参数:
//   - base: 基准货币代码，如 "USD"
//   - date: 日期，格式 "2023-02-03"
//   - symbols: 目标货币代码列表，多个用逗号分隔，空字符串表示所有货币
//   - apiKey: CurrencyScoop API Key
//
// 返回:
//   - []CurrencyHistoryItem: 历史汇率数据列表
//   - error: 错误信息
//
// 示例:
//
//	rates, err := currency.CurrencyHistory("USD", "2023-02-03", "CNY,EUR", "your-api-key")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range rates {
//	    fmt.Printf("%s: %.4f (日期=%s)\n", item.Currency, item.Value, item.Date)
//	}
func CurrencyHistory(base, date, symbols, apiKey string) ([]CurrencyHistoryItem, error) {
	url := "https://api.currencyscoop.com/v1/historical"

	params := map[string]string{
		"base":    base,
		"date":    date,
		"symbols": symbols,
		"api_key": apiKey,
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求历史汇率数据失败: %w", err)
	}

	json := gjson.ParseBytes(resp.Body())

	// 检查API错误
	if !json.Get("meta.code").Exists() || json.Get("meta.code").Int() != 200 {
		errMsg := json.Get("meta.error_detail").String()
		if errMsg == "" {
			errMsg = "未知错误"
		}
		return nil, fmt.Errorf("API返回错误: %s", errMsg)
	}

	response := json.Get("response")
	dateStr := response.Get("date").String()
	queryDate, _ := time.Parse("2006-01-02", dateStr)

	var items []CurrencyHistoryItem
	response.Get("rates").ForEach(func(key, value gjson.Result) bool {
		items = append(items, CurrencyHistoryItem{
			Currency: key.String(),
			Date:     queryDate,
			Value:    value.Float(),
		})
		return true
	})

	return items, nil
}

// CurrencyTimeSeries 获取汇率时间序列数据
//
// 从 currencyscoop.com 获取时间序列汇率数据（需要API Key和特殊权限）
//
// 参数:
//   - base: 基准货币代码，如 "USD"
//   - startDate: 开始日期，格式 "2023-02-03"
//   - endDate: 结束日期，格式 "2023-03-04"
//   - symbols: 目标货币代码列表，多个用逗号分隔，空字符串表示所有货币
//   - apiKey: CurrencyScoop API Key
//
// 返回:
//   - map[string][]CurrencyHistoryItem: 时间序列数据，key为日期，value为该日期的汇率列表
//   - error: 错误信息
//
// 示例:
//
//	timeSeries, err := currency.CurrencyTimeSeries("USD", "2023-02-03", "2023-03-04", "CNY", "your-api-key")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for date, rates := range timeSeries {
//	    fmt.Printf("日期 %s:\n", date)
//	    for _, rate := range rates {
//	        fmt.Printf("  %s: %.4f\n", rate.Currency, rate.Value)
//	    }
//	}
func CurrencyTimeSeries(base, startDate, endDate, symbols, apiKey string) (map[string][]CurrencyHistoryItem, error) {
	url := "https://api.currencyscoop.com/v1/timeseries"

	params := map[string]string{
		"base":       base,
		"start_date": startDate,
		"end_date":   endDate,
		"symbols":    symbols,
		"api_key":    apiKey,
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求时间序列数据失败: %w", err)
	}

	json := gjson.ParseBytes(resp.Body())

	// 检查API错误
	if !json.Get("meta.code").Exists() || json.Get("meta.code").Int() != 200 {
		errMsg := json.Get("meta.error_detail").String()
		if errMsg == "" {
			errMsg = "未知错误"
		}
		return nil, fmt.Errorf("API返回错误: %s", errMsg)
	}

	response := json.Get("response")
	result := make(map[string][]CurrencyHistoryItem)

	response.ForEach(func(dateKey, ratesObj gjson.Result) bool {
		dateStr := dateKey.String()
		date, _ := time.Parse("2006-01-02", dateStr)

		var items []CurrencyHistoryItem
		ratesObj.ForEach(func(currKey, value gjson.Result) bool {
			items = append(items, CurrencyHistoryItem{
				Currency: currKey.String(),
				Date:     date,
				Value:    value.Float(),
			})
			return true
		})

		result[dateStr] = items
		return true
	})

	return result, nil
}

// CurrencyCurrencies 获取货币列表
//
// 从 currencyscoop.com 获取支持的货币列表（需要API Key）
//
// 参数:
//   - cType: 货币类型，目前只支持 "fiat"（法定货币）
//   - apiKey: CurrencyScoop API Key
//
// 返回:
//   - []CurrencyCurrenciesItem: 货币列表
//   - error: 错误信息
//
// 示例:
//
//	currencies, err := currency.CurrencyCurrencies("fiat", "your-api-key")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, curr := range currencies {
//	    fmt.Printf("%s (%s): %s - %s\n", curr.CurrencyCode, curr.Symbol, curr.CurrencyName, curr.Countries)
//	}
func CurrencyCurrencies(cType, apiKey string) ([]CurrencyCurrenciesItem, error) {
	url := "https://api.currencyscoop.com/v1/currencies"

	params := map[string]string{
		"type":    cType,
		"api_key": apiKey,
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求货币列表失败: %w", err)
	}

	json := gjson.ParseBytes(resp.Body())

	// 检查API错误
	if !json.Get("meta.code").Exists() || json.Get("meta.code").Int() != 200 {
		errMsg := json.Get("meta.error_detail").String()
		if errMsg == "" {
			errMsg = "未知错误"
		}
		return nil, fmt.Errorf("API返回错误: %s", errMsg)
	}

	response := json.Get("response.fiats")
	var items []CurrencyCurrenciesItem

	response.ForEach(func(key, value gjson.Result) bool {
		items = append(items, CurrencyCurrenciesItem{
			CurrencyCode: value.Get("currency_code").String(),
			CurrencyName: value.Get("currency_name").String(),
			Countries:    value.Get("countries").String(),
			Symbol:       value.Get("symbol").String(),
		})
		return true
	})

	return items, nil
}

// CurrencyConvert 货币转换
//
// 从 currencyscoop.com 进行货币转换计算（需要API Key）
//
// 参数:
//   - from: 源货币代码，如 "USD"
//   - to: 目标货币代码，如 "CNY"
//   - amount: 源货币金额
//   - apiKey: CurrencyScoop API Key
//
// 返回:
//   - *CurrencyConvertResult: 转换结果
//   - error: 错误信息
//
// 示例:
//
//	result, err := currency.CurrencyConvert("USD", "CNY", 10000, "your-api-key")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("%.2f %s = %.2f %s (时间=%s)\n",
//	    result.Amount, result.From, result.Value, result.To, result.Timestamp)
func CurrencyConvert(from, to string, amount float64, apiKey string) (*CurrencyConvertResult, error) {
	url := "https://api.currencyscoop.com/v1/convert"

	params := map[string]string{
		"from":    from,
		"to":      to,
		"amount":  fmt.Sprintf("%.2f", amount),
		"api_key": apiKey,
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求货币转换失败: %w", err)
	}

	json := gjson.ParseBytes(resp.Body())

	// 检查API错误
	if !json.Get("meta.code").Exists() || json.Get("meta.code").Int() != 200 {
		errMsg := json.Get("meta.error_detail").String()
		if errMsg == "" {
			errMsg = "未知错误"
		}
		return nil, fmt.Errorf("API返回错误: %s", errMsg)
	}

	response := json.Get("response")
	timestamp := time.Unix(response.Get("timestamp").Int(), 0)

	result := &CurrencyConvertResult{
		From:      response.Get("from").String(),
		To:        response.Get("to").String(),
		Amount:    response.Get("amount").Float(),
		Value:     response.Get("value").Float(),
		Timestamp: timestamp,
	}

	return result, nil
}
