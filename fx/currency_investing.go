package fx

import (
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/PuerkitoBio/goquery"
)

const (
	// 英为财情货币URL
	investingCurrencyURL = "https://cn.investing.com/currencies/Service/currency"
	investingRegionURL   = "https://cn.investing.com/currencies/Service/region"
)

// CurrencyPairItem 货币对数据结构
type CurrencyPairItem struct {
	Name string `json:"name"` // 货币对名称
	Code string `json:"code"` // 货币对代码
}

// CurrencyPairMap 获取指定货币的所有可获取货币对的数据
//
// 目标地址: https://cn.investing.com/currencies/cny-jmd
//
// 参数:
//   - symbol: 指定货币，如 "美元", "人民币" 等
//
// 返回:
//   - []CurrencyPairItem: 货币对数据列表
//   - error: 错误信息
//
// 示例:
//
//	data, err := fx.CurrencyPairMap("人民币")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range data {
//	    fmt.Printf("%s: %s\n", item.Name, item.Code)
//	}
func CurrencyPairMap(symbol string) ([]CurrencyPairItem, error) {
	headers := map[string]string{
		"Accept":           "application/json, text/javascript, */*; q=0.01",
		"Accept-Language":  "zh-CN,zh;q=0.9,en;q=0.8",
		"Cache-Control":    "no-cache",
		"Connection":       "keep-alive",
		"Host":             "cn.investing.com",
		"Pragma":           "no-cache",
		"Referer":          "https://cn.investing.com/currencies/single-currency-crosses",
		"Sec-Fetch-Mode":   "cors",
		"Sec-Fetch-Site":   "same-origin",
		"User-Agent":       "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36",
		"X-Requested-With": "XMLHttpRequest",
	}

	// 获取区域和货币ID映射
	regionCode := make([]string, 0)
	regionName := make([]string, 0)

	regionIDs := []string{"4", "1", "8", "7", "6"}
	for _, regionID := range regionIDs {
		params := map[string]string{
			"region_ID":   regionID,
			"currency_ID": "false",
		}

		resp, err := utils.GetWithHeaders(investingRegionURL, params, headers)
		if err != nil {
			continue
		}

		doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
		if err != nil {
			continue
		}

		doc.Find("[data-sml-id]").Each(func(i int, s *goquery.Selection) {
			if _, hasTitle := s.Attr("title"); !hasTitle {
				if continentID, ok := s.Attr("continentid"); ok {
					regionCode = append(regionCode, continentID+"-"+regionID)
					if iElem := s.Find("i"); iElem.Length() > 0 {
						regionName = append(regionName, strings.TrimSpace(iElem.Text()))
					}
				}
			}
		})
	}

	// 构建名称到ID的映射
	nameIDMap := make(map[string]string)
	for i, name := range regionName {
		if i < len(regionCode) {
			nameIDMap[name] = regionCode[i]
		}
	}

	// 获取指定货币的货币对
	code, ok := nameIDMap[symbol]
	if !ok {
		return nil, fmt.Errorf("未找到货币: %s", symbol)
	}

	parts := strings.Split(code, "-")
	if len(parts) != 2 {
		return nil, fmt.Errorf("无效的货币代码: %s", code)
	}

	params := map[string]string{
		"region_ID":   parts[1],
		"currency_ID": parts[0],
	}

	resp, err := utils.GetWithHeaders(investingCurrencyURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("获取货币对失败: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	var items []CurrencyPairItem
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		title, _ := s.Attr("title")
		if href != "" && title != "" {
			parts := strings.Split(href, "/")
			codeVal := ""
			if len(parts) > 0 {
				codeVal = parts[len(parts)-1]
			}
			items = append(items, CurrencyPairItem{
				Name: strings.ReplaceAll(title, " ", "-"),
				Code: codeVal,
			})
		}
	})

	return items, nil
}
