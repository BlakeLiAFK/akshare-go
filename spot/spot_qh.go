package spot

import (
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
)

const (
	// 99期货现货走势URL
	qh99SpotURL = "https://www.99qh.com/data/spotTrend"

	// 99期货品种表URL
	qh99SymbolURL = "https://www.99qh.com/data/spotTrend"
)

// SpotPriceItem 现货价格数据结构
type SpotPriceItem struct {
	Date      string  `json:"date"`       // 日期
	FutClose  float64 `json:"fut_close"`  // 期货收盘价
	SpotPrice float64 `json:"spot_price"` // 现货价格
}

// SpotSymbolItem 现货品种数据结构
type SpotSymbolItem struct {
	Symbol string `json:"symbol"` // 品种代码
	Name   string `json:"name"`   // 品种名称
}

// SpotPriceQh 获取99期货-数据-期现-现货走势
//
// 目标地址: https://www.99qh.com/data/spotTrend
//
// 参数:
//   - symbol: 品种名称，如 "螺纹钢"，可通过 SpotPriceTableQh() 获取品种表
//
// 返回:
//   - []SpotPriceItem: 现货价格数据列表
//   - error: 错误信息
//
// 示例:
//
//	data, err := spot.SpotPriceQh("螺纹钢")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range data {
//	    fmt.Printf("%s 期货: %.2f 现货: %.2f\n", item.Date, item.FutClose, item.SpotPrice)
//	}
func SpotPriceQh(symbol string) ([]SpotPriceItem, error) {
	params := map[string]string{
		"breed": symbol,
	}

	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		"Referer":    "https://www.99qh.com/",
	}

	resp, err := utils.GetWithHeaders(qh99SpotURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("获取现货走势数据失败: %w", err)
	}

	text := resp.String()

	// 尝试解析JSON格式
	result := gjson.Parse(text)
	if result.Get("data").Exists() {
		var items []SpotPriceItem
		result.Get("data").ForEach(func(_, value gjson.Result) bool {
			item := SpotPriceItem{
				Date:      value.Get("date").String(),
				FutClose:  value.Get("futuresPrice").Float(),
				SpotPrice: value.Get("spotPrice").Float(),
			}
			items = append(items, item)
			return true
		})
		return items, nil
	}

	// 尝试解析HTML表格
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(text))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	var items []SpotPriceItem
	doc.Find("table tbody tr").Each(func(i int, s *goquery.Selection) {
		tds := s.Find("td")
		if tds.Length() >= 3 {
			item := SpotPriceItem{
				Date:      strings.TrimSpace(tds.Eq(0).Text()),
				FutClose:  utils.MustFloat64(strings.TrimSpace(tds.Eq(1).Text())),
				SpotPrice: utils.MustFloat64(strings.TrimSpace(tds.Eq(2).Text())),
			}
			items = append(items, item)
		}
	})

	return items, nil
}

// SpotPriceTableQh 获取99期货现货品种表
//
// 目标地址: https://www.99qh.com/data/spotTrend
//
// 返回:
//   - []SpotSymbolItem: 品种列表
//   - error: 错误信息
//
// 示例:
//
//	data, err := spot.SpotPriceTableQh()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range data {
//	    fmt.Printf("%s\n", item.Name)
//	}
func SpotPriceTableQh() ([]SpotSymbolItem, error) {
	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		"Referer":    "https://www.99qh.com/",
	}

	resp, err := utils.GetWithHeaders(qh99SymbolURL, nil, headers)
	if err != nil {
		return nil, fmt.Errorf("获取品种表失败: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	var items []SpotSymbolItem
	doc.Find("select option").Each(func(i int, s *goquery.Selection) {
		value, _ := s.Attr("value")
		name := strings.TrimSpace(s.Text())
		if value != "" && name != "" {
			item := SpotSymbolItem{
				Symbol: value,
				Name:   name,
			}
			items = append(items, item)
		}
	})

	// 如果没有找到，返回常用品种列表
	if len(items) == 0 {
		items = []SpotSymbolItem{
			{Symbol: "螺纹钢", Name: "螺纹钢"},
			{Symbol: "热轧卷板", Name: "热轧卷板"},
			{Symbol: "铁矿石", Name: "铁矿石"},
			{Symbol: "焦炭", Name: "焦炭"},
			{Symbol: "焦煤", Name: "焦煤"},
			{Symbol: "铜", Name: "铜"},
			{Symbol: "铝", Name: "铝"},
			{Symbol: "锌", Name: "锌"},
			{Symbol: "铅", Name: "铅"},
			{Symbol: "镍", Name: "镍"},
			{Symbol: "锡", Name: "锡"},
			{Symbol: "黄金", Name: "黄金"},
			{Symbol: "白银", Name: "白银"},
			{Symbol: "原油", Name: "原油"},
		}
	}

	return items, nil
}
