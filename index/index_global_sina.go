package index

import (
	"fmt"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// IndexGlobalNameTable 获取新浪全球指数名称代码映射表
//
// 返回:
//   - []IndexCodeName: 指数代码名称列表
//   - error: 错误信息
//
// 示例:
//
//	names, err := index.IndexGlobalNameTable()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, n := range names {
//	    fmt.Printf("%s: %s\n", n.Name, n.Code)
//	}
func IndexGlobalNameTable() ([]IndexCodeName, error) {
	result := make([]IndexCodeName, 0, len(GlobalSinaSymbolMap))
	for name, code := range GlobalSinaSymbolMap {
		result = append(result, IndexCodeName{
			Name: name,
			Code: code,
		})
	}
	return result, nil
}

// IndexGlobalHistSina 新浪财经-全球指数历史行情
//
// 参数:
//   - symbol: 指数名称，可以通过 IndexGlobalNameTable() 获取
//   - 如 "英国富时100指数"、"德国DAX 30种股价指数"等
//
// 返回:
//   - []GlobalIndexKLine: 全球指数K线数据
//   - error: 错误信息
//
// 示例:
//
//	// 获取英国富时100指数历史行情
//	klines, err := index.IndexGlobalHistSina("英国富时100指数")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, k := range klines {
//	    fmt.Printf("%s: %.2f\n", k.Date.Format("2006-01-02"), k.Close)
//	}
func IndexGlobalHistSina(symbol string) ([]GlobalIndexKLine, error) {
	code, ok := GlobalSinaSymbolMap[symbol]
	if !ok {
		return nil, fmt.Errorf("未找到指数: %s", symbol)
	}

	params := map[string]string{
		"symbol": code,
		"num":    "10000",
	}

	headers := map[string]string{
		"Referer": "https://finance.sina.com.cn/stock/globalindex/quotes/",
	}

	resp, err := utils.GetWithHeaders(SinaGlobalIndexURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求全球指数数据失败: %w", err)
	}

	result := gjson.Get(resp.String(), "result.data")
	if !result.Exists() {
		return nil, fmt.Errorf("未找到全球指数数据")
	}

	klines := make([]GlobalIndexKLine, 0)
	for _, item := range result.Array() {
		dateStr := item.Get("d").String()
		date, _ := time.Parse("2006-01-02", dateStr)

		kline := GlobalIndexKLine{
			Date:   date,
			Open:   item.Get("o").Float(),
			High:   item.Get("h").Float(),
			Low:    item.Get("l").Float(),
			Close:  item.Get("c").Float(),
			Volume: int64(item.Get("v").Float()),
		}
		klines = append(klines, kline)
	}

	return klines, nil
}

// MustParseDate 解析日期字符串，支持多种格式
func MustParseDate(s string) time.Time {
	// 尝试多种日期格式
	formats := []string{
		"2006-01-02",
		"2006/01/02",
		"20060102",
		time.RFC3339,
	}

	for _, format := range formats {
		if t, err := time.Parse(format, s); err == nil {
			return t
		}
	}

	return time.Time{}
}
