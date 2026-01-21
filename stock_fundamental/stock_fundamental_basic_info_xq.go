package stock_fundamental

import (
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// StockIndividualBasicInfoXq 雪球-A股公司简介
//
// 获取雪球网站的A股公司简介数据
//
// 参数:
//   - symbol: 股票代码，如 "SH601127"
//   - token: 雪球 token，可选参数。如果为空，使用默认值
//
// 返回:
//   - []StockIpoInfoItem: 公司简介数据，每个item包含字段名和值
//   - error: 错误信息
//
// 示例:
//
//	info, err := stock_fundamental.StockIndividualBasicInfoXq("SH601127", "")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range info {
//	    fmt.Printf("%s: %s\n", item.Item, item.Value)
//	}
func StockIndividualBasicInfoXq(symbol, token string) ([]StockIpoInfoItem, error) {
	if symbol == "" {
		return nil, fmt.Errorf("股票代码不能为空")
	}

	// 如果token为空，使用默认token
	if token == "" {
		token = "e09cdd7f42e4b7f7e1b088e29df8ee6df995ce4e" // 默认token
	}

	url := "https://stock.xueqiu.com/v5/stock/f10/cn/company.json"
	params := map[string]string{
		"symbol": symbol,
	}

	headers := map[string]string{
		"cookie": fmt.Sprintf("xq_a_token=%s;", token),
	}

	resp, err := utils.GetWithHeaders(url, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求雪球A股公司简介失败: %w", err)
	}

	json := gjson.ParseBytes(resp.Body())

	// 检查错误码
	errorCode := json.Get("error_code").Int()
	if errorCode != 0 {
		errorMsg := json.Get("error_description").String()
		return nil, fmt.Errorf("API返回错误(code=%d): %s", errorCode, errorMsg)
	}

	data := json.Get("data")

	if !data.Exists() || data.Type.String() == "null" {
		return nil, fmt.Errorf("未找到公司简介数据，响应: %s", truncateString(resp.String(), 200))
	}

	var items []StockIpoInfoItem

	// data是对象类型
	if data.Type.String() == "JSON" {
		// 遍历data对象，提取所有字段
		data.ForEach(func(key, value gjson.Result) bool {
			item := StockIpoInfoItem{
				Item:  key.String(),
				Value: value.String(),
			}
			items = append(items, item)
			return true
		})
	}

	if len(items) == 0 {
		return nil, fmt.Errorf("未找到公司简介数据，响应: %s", truncateString(resp.String(), 200))
	}

	return items, nil
}

// min 返回两个整数中的较小值
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// StockIndividualBasicInfoUsXq 雪球-美股公司简介
//
// 获取雪球网站的美股公司简介数据
//
// 参数:
//   - symbol: 美股代码，如 "NVDA"
//   - token: 雪球 token，可选参数。如果为空，使用默认值
//
// 返回:
//   - []StockIpoInfoItem: 公司简介数据，每个item包含字段名和值
//   - error: 错误信息
//
// 示例:
//
//	info, err := stock_fundamental.StockIndividualBasicInfoUsXq("NVDA", "")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range info {
//	    fmt.Printf("%s: %s\n", item.Item, item.Value)
//	}
func StockIndividualBasicInfoUsXq(symbol, token string) ([]StockIpoInfoItem, error) {
	if symbol == "" {
		return nil, fmt.Errorf("股票代码不能为空")
	}

	// 如果token为空，使用默认token
	if token == "" {
		token = "e09cdd7f42e4b7f7e1b088e29df8ee6df995ce4e" // 默认token
	}

	url := "https://stock.xueqiu.com/v5/stock/f10/us/company.json"
	params := map[string]string{
		"symbol": symbol,
	}

	headers := map[string]string{
		"cookie": fmt.Sprintf("xq_a_token=%s;", token),
	}

	resp, err := utils.GetWithHeaders(url, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求雪球美股公司简介失败: %w", err)
	}

	json := gjson.ParseBytes(resp.Body())

	// 检查错误码
	errorCode := json.Get("error_code").Int()
	if errorCode != 0 {
		errorMsg := json.Get("error_description").String()
		return nil, fmt.Errorf("API返回错误(code=%d): %s", errorCode, errorMsg)
	}

	data := json.Get("data")

	if !data.Exists() || data.Type.String() == "null" {
		return nil, fmt.Errorf("未找到公司简介数据，响应: %s", truncateString(resp.String(), 200))
	}

	var items []StockIpoInfoItem

	// data是对象类型
	if data.Type.String() == "JSON" {
		// 遍历data对象，提取所有字段
		data.ForEach(func(key, value gjson.Result) bool {
			item := StockIpoInfoItem{
				Item:  key.String(),
				Value: value.String(),
			}
			items = append(items, item)
			return true
		})
	}

	if len(items) == 0 {
		return nil, fmt.Errorf("未找到公司简介数据，响应: %s", truncateString(resp.String(), 200))
	}

	return items, nil
}

// StockIndividualBasicInfoHkXq 雪球-港股公司简介
//
// 获取雪球网站的港股公司简介数据
//
// 参数:
//   - symbol: 港股代码，如 "02097"
//   - token: 雪球 token，可选参数。如果为空，使用默认值
//
// 返回:
//   - []StockIpoInfoItem: 公司简介数据，每个item包含字段名和值
//   - error: 错误信息
//
// 示例:
//
//	info, err := stock_fundamental.StockIndividualBasicInfoHkXq("02097", "")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range info {
//	    fmt.Printf("%s: %s\n", item.Item, item.Value)
//	}
func StockIndividualBasicInfoHkXq(symbol, token string) ([]StockIpoInfoItem, error) {
	if symbol == "" {
		return nil, fmt.Errorf("股票代码不能为空")
	}

	// 如果token为空，使用默认token
	if token == "" {
		token = "e09cdd7f42e4b7f7e1b088e29df8ee6df995ce4e" // 默认token
	}

	url := "https://stock.xueqiu.com/v5/stock/f10/hk/company.json"
	params := map[string]string{
		"symbol": symbol,
	}

	headers := map[string]string{
		"cookie": fmt.Sprintf("xq_a_token=%s;", token),
	}

	resp, err := utils.GetWithHeaders(url, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求雪球港股公司简介失败: %w", err)
	}

	json := gjson.ParseBytes(resp.Body())

	// 检查错误码
	errorCode := json.Get("error_code").Int()
	if errorCode != 0 {
		errorMsg := json.Get("error_description").String()
		return nil, fmt.Errorf("API返回错误(code=%d): %s", errorCode, errorMsg)
	}

	data := json.Get("data")

	if !data.Exists() || data.Type.String() == "null" {
		return nil, fmt.Errorf("未找到公司简介数据，响应: %s", truncateString(resp.String(), 200))
	}

	var items []StockIpoInfoItem

	// data是对象类型
	if data.Type.String() == "JSON" {
		// 遍历data对象，提取所有字段
		data.ForEach(func(key, value gjson.Result) bool {
			item := StockIpoInfoItem{
				Item:  key.String(),
				Value: value.String(),
			}
			items = append(items, item)
			return true
		})
	}

	if len(items) == 0 {
		return nil, fmt.Errorf("未找到公司简介数据，响应: %s", truncateString(resp.String(), 200))
	}

	return items, nil
}

// truncateString 截断字符串到指定长度
func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen]
}
