package fortune

import (
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/PuerkitoBio/goquery"
)

const (
	// 财富500强URL
	fortune500URL = "https://www.fortunechina.com/fortune500/index.htm"
)

// Fortune500Item 财富500强数据结构
type Fortune500Item struct {
	Rank    string `json:"rank"`    // 排名
	Company string `json:"company"` // 公司名称
	Revenue string `json:"revenue"` // 营收
	Profit  string `json:"profit"`  // 利润
	Country string `json:"country"` // 国家
}

// FortuneRank 获取历年世界500强榜单数据
//
// 目标地址: https://www.fortunechina.com/fortune500/index.htm
//
// 参数:
//   - year: 年份，从1996年开始
//
// 返回:
//   - []Fortune500Item: 财富500强数据列表
//   - error: 错误信息
//
// 示例:
//
//	data, err := fortune.FortuneRank("2023")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range data[:10] {
//	    fmt.Printf("%s %s %s\n", item.Rank, item.Company, item.Revenue)
//	}
func FortuneRank(year string) ([]Fortune500Item, error) {
	// 首先获取年份和URL的映射
	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
	}

	resp, err := utils.GetWithHeaders(fortune500URL, nil, headers)
	if err != nil {
		return nil, fmt.Errorf("获取财富500强页面失败: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	// 获取年份URL映射
	yearURLMap := make(map[string]string)
	doc.Find("div.swiper-slide a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		yearText := strings.TrimSpace(s.Text())
		if href != "" && yearText != "" {
			yearURLMap[yearText] = href
		}
	})

	// 2023年特殊处理
	yearURLMap["2023"] = "https://www.fortunechina.com/fortune500/c/2023-08/02/content_436874.htm"

	url, ok := yearURLMap[year]
	if !ok {
		return nil, fmt.Errorf("未找到年份 %s 的数据", year)
	}

	// 获取具体年份的数据
	resp, err = utils.GetWithHeaders(url, nil, headers)
	if err != nil {
		return nil, fmt.Errorf("获取 %s 年数据失败: %w", year, err)
	}

	doc, err = goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	var items []Fortune500Item
	doc.Find("table tbody tr").Each(func(i int, s *goquery.Selection) {
		tds := s.Find("td")
		if tds.Length() >= 3 {
			item := Fortune500Item{
				Rank:    strings.TrimSpace(tds.Eq(0).Text()),
				Company: strings.TrimSpace(tds.Eq(1).Text()),
				Revenue: strings.TrimSpace(tds.Eq(2).Text()),
			}
			if tds.Length() >= 4 {
				item.Profit = strings.TrimSpace(tds.Eq(3).Text())
			}
			if tds.Length() >= 5 {
				item.Country = strings.TrimSpace(tds.Eq(4).Text())
			}
			items = append(items, item)
		}
	})

	return items, nil
}
