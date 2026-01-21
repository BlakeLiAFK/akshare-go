package index

import (
	"fmt"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// IndexUSSpotSina 新浪财经-美股指数实时行情
//
// 返回:
//   - []IndexQuote: 美股指数实时行情列表
//   - error: 错误信息
func IndexUSSpotSina() ([]IndexQuote, error) {
	// 新浪美股指数代码
	symbols := []string{".IXIC", ".DJI", ".INX", ".NDX"}

	quotes := make([]IndexQuote, 0)
	for _, symbol := range symbols {
		klines, err := IndexUSStockSina(symbol)
		if err != nil {
			continue
		}
		if len(klines) > 0 {
			latest := klines[0]
			quotes = append(quotes, IndexQuote{
				Code:       symbol,
				Name:       USIndexSymbolMap[symbol],
				Price:      latest.Close,
				UpdateTime: latest.Date.Format("2006-01-02"),
			})
		}
	}

	return quotes, nil
}

// IndexUSStockSina 新浪财经-美股指数历史行情
//
// 参数:
//   - symbol: 美股指数代码，如 ".INX"(标普500), ".DJI"(道琼斯), ".IXIC"(纳斯达克), ".NDX"(纳斯达克100)
//
// 返回:
//   - []GlobalIndexKLine: 美股指数K线数据
//   - error: 错误信息
//
// 示例:
//
//	// 获取标普500历史行情
//	klines, err := index.IndexUSStockSina(".INX")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, k := range klines {
//	    fmt.Printf("%s: %.2f\n", k.Date.Format("2006-01-02"), k.Close)
//	}
func IndexUSStockSina(symbol string) ([]GlobalIndexKLine, error) {
	url := fmt.Sprintf(USIndexSinaURL, symbol)

	headers := map[string]string{
		"Referer": "https://stock.finance.sina.com.cn/usstock/quotes/" + symbol + ".html",
	}

	resp, err := utils.GetWithHeaders(url, nil, headers)
	if err != nil {
		return nil, fmt.Errorf("请求美股指数数据失败: %w", err)
	}

	// 新浪美股数据格式: index_data = "..." (需要JS解密)
	// 这里简化处理，直接解析
	text := resp.String()

	// 提取数据部分
	dataStart := -1
	for i := 0; i < len(text)-10; i++ {
		if text[i:i+10] == "\"data\":[" {
			dataStart = i + 8
			break
		}
	}

	if dataStart == -1 {
		return nil, fmt.Errorf("未找到美股指数数据")
	}

	// 查找数据结束位置
	dataEnd := dataStart
	braceCount := 0
	inArray := false
	for i := dataStart; i < len(text); i++ {
		if text[i] == '[' {
			inArray = true
			braceCount++
		} else if text[i] == ']' {
			braceCount--
			if inArray && braceCount == 0 {
				dataEnd = i + 1
				break
			}
		}
	}

	if dataEnd <= dataStart {
		return nil, fmt.Errorf("解析美股指数数据失败")
	}

	jsonData := text[dataStart:dataEnd]
	result := gjson.Parse(jsonData)

	klines := make([]GlobalIndexKLine, 0)
	for _, item := range result.Array() {
		dateStr := item.Get("date").String()
		date, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			continue
		}

		kline := GlobalIndexKLine{
			Date:   date,
			Open:   item.Get("open").Float(),
			High:   item.Get("high").Float(),
			Low:    item.Get("low").Float(),
			Close:  item.Get("close").Float(),
			Volume: int64(item.Get("volume").Float()),
		}
		klines = append(klines, kline)
	}

	return klines, nil
}
