package fortune

import (
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/PuerkitoBio/goquery"
)

// BloombergBillionaireItem 彭博亿万富豪数据结构
type BloombergBillionaireItem struct {
	Rank          string `json:"rank"`            // 排名
	Name          string `json:"name"`            // 姓名
	TotalNetWorth string `json:"total_net_worth"` // 总净资产
	LastChange    string `json:"last_change"`     // 最新变化
	YTDChange     string `json:"ytd_change"`      // 年初至今变化
	Country       string `json:"country"`         // 国家
	Industry      string `json:"industry"`        // 行业
}

// IndexBloombergBillionaires 获取彭博亿万富豪指数
//
// 目标地址: https://www.bloomberg.com/billionaires/
//
// 返回:
//   - []BloombergBillionaireItem: 富豪数据列表
//   - error: 错误信息
//
// 示例:
//
//	data, err := fortune.IndexBloombergBillionaires()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range data[:10] {
//	    fmt.Printf("%s %s %s\n", item.Rank, item.Name, item.TotalNetWorth)
//	}
func IndexBloombergBillionaires() ([]BloombergBillionaireItem, error) {
	url := "https://www.bloomberg.com/billionaires"

	headers := map[string]string{
		"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8",
		"accept-language":           "zh-CN,zh;q=0.9,en;q=0.8",
		"cache-control":             "no-cache",
		"pragma":                    "no-cache",
		"sec-fetch-dest":            "document",
		"sec-fetch-mode":            "navigate",
		"sec-fetch-site":            "same-origin",
		"upgrade-insecure-requests": "1",
		"referer":                   "https://www.bloomberg.com/",
		"user-agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36",
	}

	resp, err := utils.GetWithHeaders(url, nil, headers)
	if err != nil {
		return nil, fmt.Errorf("获取彭博富豪数据失败: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	var items []BloombergBillionaireItem
	doc.Find(".table-chart .table-row").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		text = strings.ReplaceAll(text, "\n", "")
		parts := strings.Fields(text)
		if len(parts) >= 7 {
			item := BloombergBillionaireItem{
				Rank:          parts[0],
				Name:          strings.Join(parts[1:len(parts)-5], " "),
				TotalNetWorth: parts[len(parts)-5],
				LastChange:    parts[len(parts)-4],
				YTDChange:     parts[len(parts)-3],
				Country:       parts[len(parts)-2],
				Industry:      parts[len(parts)-1],
			}
			items = append(items, item)
		}
	})

	return items, nil
}

// IndexBloombergBillionairesHist 获取彭博亿万富豪指数历史数据
//
// 目标地址: https://stats.areppim.com/stats/links_billionairexlists.htm
//
// 参数:
//   - year: 年份，如 "2021", "2019", "2018" 等
//
// 返回:
//   - []BloombergBillionaireItem: 富豪数据列表
//   - error: 错误信息
//
// 示例:
//
//	data, err := fortune.IndexBloombergBillionairesHist("2021")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range data[:10] {
//	    fmt.Printf("%s %s %s\n", item.Rank, item.Name, item.TotalNetWorth)
//	}
func IndexBloombergBillionairesHist(year string) ([]BloombergBillionaireItem, error) {
	// 获取年份后两位
	yearSuffix := year
	if len(year) >= 2 {
		yearSuffix = year[len(year)-2:]
	}

	url := fmt.Sprintf("https://stats.areppim.com/listes/list_billionairesx%sxwor.htm", yearSuffix)

	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
	}

	resp, err := utils.GetWithHeaders(url, nil, headers)
	if err != nil {
		return nil, fmt.Errorf("获取彭博富豪历史数据失败: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	var items []BloombergBillionaireItem
	doc.Find("table tr").Each(func(i int, s *goquery.Selection) {
		tds := s.Find("td")
		if tds.Length() >= 4 {
			rank := strings.TrimSpace(tds.Eq(0).Text())
			// 跳过非数字排名行
			if rank == "" || !isDigit(rank[0]) {
				return
			}

			item := BloombergBillionaireItem{
				Rank: rank,
				Name: strings.TrimSpace(tds.Eq(1).Text()),
			}
			if tds.Length() >= 3 {
				item.TotalNetWorth = strings.TrimSpace(tds.Eq(2).Text())
			}
			if tds.Length() >= 4 {
				item.Country = strings.TrimSpace(tds.Eq(3).Text())
			}
			if tds.Length() >= 5 {
				item.Industry = strings.TrimSpace(tds.Eq(4).Text())
			}
			items = append(items, item)
		}
	})

	return items, nil
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}
