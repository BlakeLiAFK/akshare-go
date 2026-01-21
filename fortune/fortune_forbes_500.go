package fortune

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	// 福布斯榜单URL
	forbesListURL = "https://www.forbeschina.com/lists"
)

// ForbesRankItem 福布斯榜单数据结构
type ForbesRankItem struct {
	Rank   string `json:"rank"`   // 排名
	Name   string `json:"name"`   // 姓名
	Wealth string `json:"wealth"` // 财富
	Source string `json:"source"` // 来源
	Age    string `json:"age"`    // 年龄
}

// ForbesRank 获取福布斯中国榜单
//
// 目标地址: https://www.forbeschina.com/lists
//
// 参数:
//   - symbol: 榜单名称，如 "2021福布斯中国创投人100" 等
//
// 返回:
//   - []ForbesRankItem: 榜单数据列表
//   - error: 错误信息
//
// 示例:
//
//	data, err := fortune.ForbesRank("2021福布斯中国香港富豪榜")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range data[:10] {
//	    fmt.Printf("%s %s %s\n", item.Rank, item.Name, item.Wealth)
//	}
func ForbesRank(symbol string) ([]ForbesRankItem, error) {
	// 跳过TLS验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// 首先获取榜单列表
	req, err := http.NewRequest("GET", forbesListURL, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("获取福布斯榜单列表失败: %w", err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	// 获取所有榜单链接
	nameURLMap := make(map[string]string)
	doc.Find("div.col-sm-4 a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		name := strings.TrimSpace(s.Text())
		if href != "" && name != "" {
			nameURLMap[name] = "https://www.forbeschina.com" + href
		}
	})

	url, ok := nameURLMap[symbol]
	if !ok {
		return nil, fmt.Errorf("未找到榜单: %s", symbol)
	}

	// 获取具体榜单数据
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")

	resp, err = client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("获取福布斯榜单数据失败: %w", err)
	}
	defer resp.Body.Close()

	doc, err = goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	var items []ForbesRankItem
	doc.Find("table tbody tr").Each(func(i int, s *goquery.Selection) {
		tds := s.Find("td")
		if tds.Length() >= 2 {
			item := ForbesRankItem{
				Rank: strings.TrimSpace(tds.Eq(0).Text()),
				Name: strings.TrimSpace(tds.Eq(1).Text()),
			}
			if tds.Length() >= 3 {
				item.Wealth = strings.TrimSpace(tds.Eq(2).Text())
			}
			if tds.Length() >= 4 {
				item.Source = strings.TrimSpace(tds.Eq(3).Text())
			}
			if tds.Length() >= 5 {
				item.Age = strings.TrimSpace(tds.Eq(4).Text())
			}
			items = append(items, item)
		}
	})

	return items, nil
}
