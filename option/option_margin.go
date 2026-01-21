package option

import (
	"fmt"
	"strings"
	"sync"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

// symbolCache 缓存期权品种信息
var (
	symbolCache     []OptionSymbol
	symbolCacheOnce sync.Once
)

// OptionSymbol 期权品种信息
type OptionSymbol struct {
	Symbol string `json:"symbol"` // 品种名称
	URL    string `json:"url"`    // 品种URL
}

// OptionMarginSymbol 获取商品期权品种代码和名称
//
// 返回:
//   - dataframe.DataFrame: 包含品种代码和URL的数据
//   - error: 错误信息
//
// 数据源: https://www.iweiai.com/qiquan/yuanyou
func OptionMarginSymbol() (dataframe.DataFrame, error) {
	symbolCacheOnce.Do(func() {
		symbols, err := fetchOptionSymbols()
		if err == nil {
			symbolCache = symbols
		}
	})

	if len(symbolCache) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未能获取期权品种信息")
	}

	// 转换为DataFrame
	var symbols []string
	var urls []string

	for _, symbol := range symbolCache {
		symbols = append(symbols, symbol.Symbol)
		urls = append(urls, symbol.URL)
	}

	df := dataframe.New(
		series.New(symbols, series.String, "symbol"),
		series.New(urls, series.String, "url"),
	)

	return df, nil
}

// fetchOptionSymbols 获取期权品种信息
func fetchOptionSymbols() ([]OptionSymbol, error) {
	url := "https://www.iweiai.com/qiquan/yuanyou"

	resp, err := utils.Get(url, nil)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	// 解析HTML
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	var symbols []OptionSymbol

	// 查找包含"qiquan"的链接
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists && strings.Contains(href, "qiquan") {
			text := strings.TrimSpace(s.Text())
			if text != "" {
				symbols = append(symbols, OptionSymbol{
					Symbol: text,
					URL:    href,
				})
			}
		}
	})

	if len(symbols) == 0 {
		return nil, fmt.Errorf("未找到期权品种信息")
	}

	return symbols, nil
}

// OptionMargin 获取商品期权保证金
//
// 参数:
//   - symbol: 商品期权品种名称，如 "原油期权"
//
// 返回:
//   - dataframe.DataFrame: 商品期权保证金数据
//   - error: 错误信息
//
// 数据源: https://www.iweiai.com/qihuo/yuanyou
func OptionMargin(symbol string) (dataframe.DataFrame, error) {
	if symbol == "" {
		return dataframe.DataFrame{}, fmt.Errorf("品种名称不能为空")
	}

	// 获取品种信息
	symbols, err := fetchOptionSymbols()
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("获取品种信息失败: %w", err)
	}

	// 查找指定品种的URL
	var targetURL string
	for _, s := range symbols {
		if s.Symbol == symbol {
			targetURL = s.URL
			break
		}
	}

	if targetURL == "" {
		return dataframe.DataFrame{}, fmt.Errorf("未找到品种: %s", symbol)
	}

	// 如果URL是相对路径，补充完整域名
	if strings.HasPrefix(targetURL, "/") {
		targetURL = "https://www.iweiai.com" + targetURL
	}

	// 请求期权保证金数据
	resp, err := utils.Get(targetURL, nil)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	// 解析HTML获取更新时间
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("解析HTML失败: %w", err)
	}

	var updateTime string
	doc.Find("small").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		if strings.Contains(text, "最近更新：") {
			updateTime = strings.TrimPrefix(text, "最近更新：")
		}
	})

	// 解析表格数据
	var records [][]string
	doc.Find("table").Each(func(i int, table *goquery.Selection) {
		table.Find("tr").Each(func(j int, row *goquery.Selection) {
			var rowData []string
			row.Find("td, th").Each(func(k int, cell *goquery.Selection) {
				text := strings.TrimSpace(cell.Text())
				rowData = append(rowData, text)
			})

			if len(rowData) > 0 {
				records = append(records, rowData)
			}
		})
	})

	if len(records) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未找到数据表格")
	}

	// 添加更新时间列
	for i := range records {
		if i == 0 {
			// 表头添加更新时间列
			records[i] = append(records[i], "更新时间")
		} else {
			// 数据行添加更新时间
			records[i] = append(records[i], updateTime)
		}
	}

	// 转换数值列
	if len(records) > 1 {
		// 假设数值列的位置（需要根据实际表格结构调整）
		numericCols := []int{1, 2, 3, 4, 5, 6, 7, 8} // 结算价、交易乘数等列
		for i := 1; i < len(records); i++ {
			for _, col := range numericCols {
				if col < len(records[i]) {
					records[i][col] = parseNumeric(records[i][col])
				}
			}
		}
	}

	// 构建DataFrame
	df := dataframe.LoadRecords(records)
	return df, nil
}

// parseNumeric 解析数值字符串
func parseNumeric(s string) string {
	s = strings.TrimSpace(s)
	if s == "" || s == "-" {
		return "0"
	}

	// 移除非数字字符（除了小数点）
	result := ""
	for _, r := range s {
		if (r >= '0' && r <= '9') || r == '.' {
			result += string(r)
		}
	}

	if result == "" {
		return "0"
	}

	return result
}
