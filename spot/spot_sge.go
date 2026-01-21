package spot

import (
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
)

const (
	// 上海黄金交易所历史数据URL
	sgeHistURL = "https://www.sge.com.cn/sjzx/mrhq"

	// 上海黄金交易所实时行情URL
	sgeQuotationsURL = "https://www.sge.com.cn/graph/quotations"

	// 上海金基准价URL
	sgeGoldenBenchmarkURL = "https://www.sge.com.cn/sjzx/jzj"

	// 上海银基准价URL
	sgeSilverBenchmarkURL = "https://www.sge.com.cn/sjzx/shyjzj"
)

// SgeHistItem 上海黄金交易所历史数据结构
type SgeHistItem struct {
	Date  string  `json:"date"`  // 日期
	Open  float64 `json:"open"`  // 开盘价
	Close float64 `json:"close"` // 收盘价
	Low   float64 `json:"low"`   // 最低价
	High  float64 `json:"high"`  // 最高价
}

// SgeSymbolItem 上海黄金交易所品种数据结构
type SgeSymbolItem struct {
	Symbol string `json:"symbol"` // 品种代码
	Name   string `json:"name"`   // 品种名称
}

// SgeQuotationItem 上海黄金交易所实时行情数据结构
type SgeQuotationItem struct {
	Symbol     string  `json:"symbol"`      // 品种
	Time       string  `json:"time"`        // 时间
	Price      float64 `json:"price"`       // 现价
	UpdateTime string  `json:"update_time"` // 更新时间
}

// SgeBenchmarkItem 基准价数据结构
type SgeBenchmarkItem struct {
	Date         string  `json:"date"`          // 交易时间
	EveningPrice float64 `json:"evening_price"` // 晚盘价
	MorningPrice float64 `json:"morning_price"` // 早盘价
}

// SpotHistSge 获取上海黄金交易所-数据资讯-行情走势-历史数据
//
// 目标地址: https://www.sge.com.cn/sjzx/mrhq
//
// 参数:
//   - symbol: 品种代码，如 "Au99.99"，可通过 SpotSymbolTableSge() 获取品种表
//
// 返回:
//   - []SgeHistItem: 历史数据列表
//   - error: 错误信息
//
// 示例:
//
//	data, err := spot.SpotHistSge("Au99.99")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range data {
//	    fmt.Printf("%s 开: %.2f 收: %.2f\n", item.Date, item.Open, item.Close)
//	}
func SpotHistSge(symbol string) ([]SgeHistItem, error) {
	params := map[string]string{
		"cate": symbol,
	}

	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		"Referer":    "https://www.sge.com.cn/",
	}

	resp, err := utils.GetWithHeaders(sgeHistURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("获取上海黄金交易所历史数据失败: %w", err)
	}

	text := resp.String()

	// 尝试解析JSON
	result := gjson.Parse(text)
	if result.Get("data").Exists() {
		var items []SgeHistItem
		result.Get("data").ForEach(func(_, value gjson.Result) bool {
			item := SgeHistItem{
				Date:  value.Get("date").String(),
				Open:  value.Get("open").Float(),
				Close: value.Get("close").Float(),
				Low:   value.Get("low").Float(),
				High:  value.Get("high").Float(),
			}
			items = append(items, item)
			return true
		})
		if len(items) > 0 {
			return items, nil
		}
	}

	// 尝试解析HTML
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(text))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	var items []SgeHistItem
	doc.Find("table tbody tr").Each(func(i int, s *goquery.Selection) {
		tds := s.Find("td")
		if tds.Length() >= 5 {
			item := SgeHistItem{
				Date:  strings.TrimSpace(tds.Eq(0).Text()),
				Open:  utils.MustFloat64(strings.TrimSpace(tds.Eq(1).Text())),
				Close: utils.MustFloat64(strings.TrimSpace(tds.Eq(2).Text())),
				Low:   utils.MustFloat64(strings.TrimSpace(tds.Eq(3).Text())),
				High:  utils.MustFloat64(strings.TrimSpace(tds.Eq(4).Text())),
			}
			items = append(items, item)
		}
	})

	return items, nil
}

// SpotSymbolTableSge 获取上海黄金交易所品种表
//
// 目标地址: https://www.sge.com.cn/sjzx/mrhq
//
// 返回:
//   - []SgeSymbolItem: 品种列表
//   - error: 错误信息
//
// 示例:
//
//	data, err := spot.SpotSymbolTableSge()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range data {
//	    fmt.Printf("%s\n", item.Symbol)
//	}
func SpotSymbolTableSge() ([]SgeSymbolItem, error) {
	// 返回常用品种列表
	items := []SgeSymbolItem{
		{Symbol: "Au99.99", Name: "黄金99.99"},
		{Symbol: "Au99.95", Name: "黄金99.95"},
		{Symbol: "Au100g", Name: "黄金100克"},
		{Symbol: "Au(T+D)", Name: "黄金T+D"},
		{Symbol: "Au(T+N1)", Name: "黄金T+N1"},
		{Symbol: "Au(T+N2)", Name: "黄金T+N2"},
		{Symbol: "mAu(T+D)", Name: "迷你黄金T+D"},
		{Symbol: "Ag99.99", Name: "白银99.99"},
		{Symbol: "Ag(T+D)", Name: "白银T+D"},
		{Symbol: "iAu99.99", Name: "国际板黄金"},
		{Symbol: "iAu99.5", Name: "国际板黄金99.5"},
		{Symbol: "iAu100g", Name: "国际板黄金100克"},
		{Symbol: "PGC30g", Name: "铂金30克"},
	}

	return items, nil
}

// SpotQuotationsSge 获取上海黄金交易所-数据资讯-行情走势-实时数据
//
// 目标地址: https://www.sge.com.cn/
//
// 参数:
//   - symbol: 品种代码，如 "Au99.99"，可通过 SpotSymbolTableSge() 获取品种表
//
// 返回:
//   - []SgeQuotationItem: 实时行情数据列表
//   - error: 错误信息
//
// 示例:
//
//	data, err := spot.SpotQuotationsSge("Au99.99")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range data {
//	    fmt.Printf("%s %s %.2f\n", item.Symbol, item.Time, item.Price)
//	}
func SpotQuotationsSge(symbol string) ([]SgeQuotationItem, error) {
	params := map[string]string{
		"cate": symbol,
	}

	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		"Referer":    "https://www.sge.com.cn/",
	}

	resp, err := utils.GetWithHeaders(sgeQuotationsURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("获取上海黄金交易所实时行情失败: %w", err)
	}

	text := resp.String()
	result := gjson.Parse(text)

	var items []SgeQuotationItem
	if result.IsArray() {
		result.ForEach(func(_, value gjson.Result) bool {
			item := SgeQuotationItem{
				Symbol:     value.Get("symbol").String(),
				Time:       value.Get("time").String(),
				Price:      value.Get("price").Float(),
				UpdateTime: value.Get("updateTime").String(),
			}
			if item.Symbol == "" {
				item.Symbol = symbol
			}
			items = append(items, item)
			return true
		})
	} else if result.Get("data").Exists() {
		result.Get("data").ForEach(func(_, value gjson.Result) bool {
			item := SgeQuotationItem{
				Symbol:     symbol,
				Time:       value.Get("time").String(),
				Price:      value.Get("price").Float(),
				UpdateTime: value.Get("updateTime").String(),
			}
			items = append(items, item)
			return true
		})
	}

	return items, nil
}

// SpotGoldenBenchmarkSge 获取上海黄金交易所-数据资讯-上海金基准价-历史数据
//
// 目标地址: https://www.sge.com.cn/sjzx/jzj
//
// 返回:
//   - []SgeBenchmarkItem: 上海金基准价数据列表
//   - error: 错误信息
//
// 示例:
//
//	data, err := spot.SpotGoldenBenchmarkSge()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range data {
//	    fmt.Printf("%s 晚盘: %.2f 早盘: %.2f\n", item.Date, item.EveningPrice, item.MorningPrice)
//	}
func SpotGoldenBenchmarkSge() ([]SgeBenchmarkItem, error) {
	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		"Referer":    "https://www.sge.com.cn/",
	}

	resp, err := utils.GetWithHeaders(sgeGoldenBenchmarkURL, nil, headers)
	if err != nil {
		return nil, fmt.Errorf("获取上海金基准价失败: %w", err)
	}

	text := resp.String()

	// 尝试解析JSON
	result := gjson.Parse(text)
	if result.Get("data").Exists() {
		var items []SgeBenchmarkItem
		result.Get("data").ForEach(func(_, value gjson.Result) bool {
			item := SgeBenchmarkItem{
				Date:         value.Get("date").String(),
				EveningPrice: value.Get("eveningPrice").Float(),
				MorningPrice: value.Get("morningPrice").Float(),
			}
			items = append(items, item)
			return true
		})
		if len(items) > 0 {
			return items, nil
		}
	}

	// 尝试解析HTML
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(text))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	var items []SgeBenchmarkItem
	doc.Find("table tbody tr").Each(func(i int, s *goquery.Selection) {
		tds := s.Find("td")
		if tds.Length() >= 3 {
			item := SgeBenchmarkItem{
				Date:         strings.TrimSpace(tds.Eq(0).Text()),
				EveningPrice: utils.MustFloat64(strings.TrimSpace(tds.Eq(1).Text())),
				MorningPrice: utils.MustFloat64(strings.TrimSpace(tds.Eq(2).Text())),
			}
			items = append(items, item)
		}
	})

	return items, nil
}

// SpotSilverBenchmarkSge 获取上海黄金交易所-数据资讯-上海银基准价-历史数据
//
// 目标地址: https://www.sge.com.cn/sjzx/shyjzj
//
// 返回:
//   - []SgeBenchmarkItem: 上海银基准价数据列表
//   - error: 错误信息
//
// 示例:
//
//	data, err := spot.SpotSilverBenchmarkSge()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range data {
//	    fmt.Printf("%s 晚盘: %.2f 早盘: %.2f\n", item.Date, item.EveningPrice, item.MorningPrice)
//	}
func SpotSilverBenchmarkSge() ([]SgeBenchmarkItem, error) {
	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		"Referer":    "https://www.sge.com.cn/",
	}

	resp, err := utils.GetWithHeaders(sgeSilverBenchmarkURL, nil, headers)
	if err != nil {
		return nil, fmt.Errorf("获取上海银基准价失败: %w", err)
	}

	text := resp.String()

	// 尝试解析JSON
	result := gjson.Parse(text)
	if result.Get("data").Exists() {
		var items []SgeBenchmarkItem
		result.Get("data").ForEach(func(_, value gjson.Result) bool {
			item := SgeBenchmarkItem{
				Date:         value.Get("date").String(),
				EveningPrice: value.Get("eveningPrice").Float(),
				MorningPrice: value.Get("morningPrice").Float(),
			}
			items = append(items, item)
			return true
		})
		if len(items) > 0 {
			return items, nil
		}
	}

	// 尝试解析HTML
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(text))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	var items []SgeBenchmarkItem
	doc.Find("table tbody tr").Each(func(i int, s *goquery.Selection) {
		tds := s.Find("td")
		if tds.Length() >= 3 {
			item := SgeBenchmarkItem{
				Date:         strings.TrimSpace(tds.Eq(0).Text()),
				EveningPrice: utils.MustFloat64(strings.TrimSpace(tds.Eq(1).Text())),
				MorningPrice: utils.MustFloat64(strings.TrimSpace(tds.Eq(2).Text())),
			}
			items = append(items, item)
		}
	})

	return items, nil
}
