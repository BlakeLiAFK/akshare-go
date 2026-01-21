package fortune

import (
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
)

const (
	// 胡润排行榜URL
	hurunRankURL     = "https://www.hurun.net/zh-CN/Rank/HsRankDetails?pagetype=rich"
	hurunRankListURL = "https://www.hurun.net/zh-CN/Rank/HsRankDetailsList"
)

// HurunRankItem 胡润排行榜数据结构
type HurunRankItem struct {
	Rank       string `json:"rank"`        // 排名
	Wealth     string `json:"wealth"`      // 财富
	RankChange string `json:"rank_change"` // 排名变化
	Name       string `json:"name"`        // 姓名
	Company    string `json:"company"`     // 企业
	Industry   string `json:"industry"`    // 行业
}

// HurunRank 获取胡润排行榜
//
// 目标地址: https://www.hurun.net/CN/HuList/Index
//
// 参数:
//   - indicator: 榜单类型，可选 "胡润百富榜", "胡润全球富豪榜", "胡润全球独角兽榜" 等
//   - year: 年份
//
// 返回:
//   - []HurunRankItem: 榜单数据列表
//   - error: 错误信息
//
// 示例:
//
//	data, err := fortune.HurunRank("胡润百富榜", "2023")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range data[:10] {
//	    fmt.Printf("%s %s %s %s\n", item.Rank, item.Name, item.Wealth, item.Company)
//	}
func HurunRank(indicator string, year string) ([]HurunRankItem, error) {
	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
	}

	// 获取榜单类型和URL映射
	resp, err := utils.GetWithHeaders(hurunRankURL, nil, headers)
	if err != nil {
		return nil, fmt.Errorf("获取胡润排行榜页面失败: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	// 获取榜单类型URL映射
	nameURLMap := make(map[string]string)
	doc.Find("ul.dropdown-menu a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		name := strings.TrimSpace(s.Text())
		if href != "" && name != "" {
			nameURLMap[name] = "https://www.hurun.net" + href
		}
	})

	url, ok := nameURLMap[indicator]
	if !ok {
		return nil, fmt.Errorf("未找到榜单类型: %s", indicator)
	}

	// 获取年份和代码映射
	resp, err = utils.GetWithHeaders(url, nil, headers)
	if err != nil {
		return nil, fmt.Errorf("获取胡润榜单页面失败: %w", err)
	}

	doc, err = goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	yearCodeMap := make(map[string]string)
	doc.Find("#exampleFormControlSelect1 option").Each(func(i int, s *goquery.Selection) {
		val, _ := s.Attr("value")
		yearText := strings.TrimSpace(s.Text())
		if val != "" && yearText != "" {
			// 解析value中的num参数
			if strings.Contains(val, "num=") {
				parts := strings.Split(val, "num=")
				if len(parts) >= 2 {
					code := strings.Split(parts[1], "&")[0]
					yearKey := strings.Split(yearText, " ")[0]
					yearCodeMap[yearKey] = code
				}
			}
		}
	})

	code, ok := yearCodeMap[year]
	if !ok {
		return nil, fmt.Errorf("未找到年份 %s 的数据", year)
	}

	// 获取具体数据
	params := map[string]string{
		"num":    code,
		"search": "",
		"offset": "0",
		"limit":  "20000",
	}

	resp, err = utils.GetWithHeaders(hurunRankListURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("获取胡润榜单数据失败: %w", err)
	}

	text := resp.String()
	rows := gjson.Get(text, "rows")
	if !rows.Exists() {
		return nil, fmt.Errorf("解析数据失败")
	}

	var items []HurunRankItem
	rows.ForEach(func(_, value gjson.Result) bool {
		item := HurunRankItem{}

		// 根据不同榜单类型解析不同字段
		switch indicator {
		case "胡润百富榜":
			item.Rank = value.Get("hs_Rank_Rich_Ranking").String()
			item.Wealth = value.Get("hs_Rank_Rich_Wealth").String()
			item.RankChange = value.Get("hs_Rank_Rich_Ranking_Change").String()
			item.Name = value.Get("hs_Rank_Rich_ChaName_Cn").String()
			item.Company = value.Get("hs_Rank_Rich_ComName_Cn").String()
			item.Industry = value.Get("hs_Rank_Rich_Industry_Cn").String()
		case "胡润全球富豪榜":
			item.Rank = value.Get("hs_Rank_Global_Ranking").String()
			item.Wealth = value.Get("hs_Rank_Global_Wealth").String()
			item.RankChange = value.Get("hs_Rank_Global_Ranking_Change").String()
			item.Name = value.Get("hs_Rank_Global_ChaName_Cn").String()
			item.Company = value.Get("hs_Rank_Global_ComName_Cn").String()
			item.Industry = value.Get("hs_Rank_Global_Industry_Cn").String()
		case "胡润全球独角兽榜":
			item.Rank = value.Get("hs_Rank_Unicorn_Ranking").String()
			item.Wealth = value.Get("hs_Rank_Unicorn_Wealth").String()
			item.RankChange = value.Get("hs_Rank_Unicorn_Ranking_Change").String()
			item.Name = value.Get("hs_Rank_Unicorn_ChaName_Cn").String()
			item.Company = value.Get("hs_Rank_Unicorn_ComName_Cn").String()
			item.Industry = value.Get("hs_Rank_Unicorn_Industry_Cn").String()
		case "胡润中国500强民营企业":
			item.Rank = value.Get("hs_Rank_CTop500_Ranking").String()
			item.Wealth = value.Get("hs_Rank_CTop500_Wealth").String()
			item.RankChange = value.Get("hs_Rank_CTop500_Ranking_Change").String()
			item.Name = value.Get("hs_Rank_CTop500_ChaName_Cn").String()
			item.Company = value.Get("hs_Rank_CTop500_ComName_Cn").String()
			item.Industry = value.Get("hs_Rank_CTop500_Industry_Cn").String()
		case "胡润世界500强":
			item.Rank = value.Get("hs_Rank_GTop500_Ranking").String()
			item.Wealth = value.Get("hs_Rank_GTop500_Wealth").String()
			item.RankChange = value.Get("hs_Rank_GTop500_Ranking_Change").String()
			item.Name = value.Get("hs_Rank_GTop500_ChaName_Cn").String()
			item.Company = value.Get("hs_Rank_GTop500_ComName_Cn").String()
			item.Industry = value.Get("hs_Rank_GTop500_Industry_Cn").String()
		default:
			// 尝试通用解析
			item.Rank = value.Get("hs_Rank_Rich_Ranking").String()
			item.Wealth = value.Get("hs_Rank_Rich_Wealth").String()
			item.Name = value.Get("hs_Rank_Rich_ChaName_Cn").String()
			item.Company = value.Get("hs_Rank_Rich_ComName_Cn").String()
			item.Industry = value.Get("hs_Rank_Rich_Industry_Cn").String()
		}

		items = append(items, item)
		return true
	})

	return items, nil
}
