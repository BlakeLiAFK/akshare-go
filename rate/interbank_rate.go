package rate

import (
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
)

const (
	// Shibor官网URL
	shiborURL = "https://www.shibor.org/shibor/web/ShiborJPG.jsp"

	// LPR数据URL
	lprURL = "https://www.chinamoney.com.cn/chinese/bklpr/"
)

// ShiborRate Shibor利率数据结构
type ShiborRate struct {
	Date      string  `json:"date"`      // 日期
	Overnight float64 `json:"overnight"` // 隔夜
	Week1     float64 `json:"week1"`     // 1周
	Week2     float64 `json:"week2"`     // 2周
	Month1    float64 `json:"month1"`    // 1月
	Month3    float64 `json:"month3"`    // 3月
	Month6    float64 `json:"month6"`    // 6月
	Month9    float64 `json:"month9"`    // 9月
	Year1     float64 `json:"year1"`     // 1年
}

// LPRRate LPR利率数据结构
type LPRRate struct {
	Date  string  `json:"date"`   // 日期
	LPR1Y float64 `json:"lpr_1y"` // 1年期LPR
	LPR5Y float64 `json:"lpr_5y"` // 5年期以上LPR
}

// RateShibor 获取Shibor利率数据
//
// 目标地址: https://www.shibor.org/
//
// 返回:
//   - []ShiborRate: Shibor利率数据列表
//   - error: 错误信息
//
// 示例:
//
//	data, err := rate.RateShibor()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range data {
//	    fmt.Printf("%s 隔夜: %.4f%% 1年: %.4f%%\n", item.Date, item.Overnight, item.Year1)
//	}
func RateShibor() ([]ShiborRate, error) {
	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		"Referer":    "https://www.shibor.org/",
	}

	resp, err := utils.GetWithHeaders(shiborURL, nil, headers)
	if err != nil {
		return nil, fmt.Errorf("获取Shibor数据失败: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	var items []ShiborRate

	doc.Find("table tbody tr").Each(func(i int, s *goquery.Selection) {
		tds := s.Find("td")
		if tds.Length() < 9 {
			return
		}

		item := ShiborRate{
			Date:      strings.TrimSpace(tds.Eq(0).Text()),
			Overnight: utils.MustFloat64(strings.TrimSpace(tds.Eq(1).Text())),
			Week1:     utils.MustFloat64(strings.TrimSpace(tds.Eq(2).Text())),
			Week2:     utils.MustFloat64(strings.TrimSpace(tds.Eq(3).Text())),
			Month1:    utils.MustFloat64(strings.TrimSpace(tds.Eq(4).Text())),
			Month3:    utils.MustFloat64(strings.TrimSpace(tds.Eq(5).Text())),
			Month6:    utils.MustFloat64(strings.TrimSpace(tds.Eq(6).Text())),
			Month9:    utils.MustFloat64(strings.TrimSpace(tds.Eq(7).Text())),
			Year1:     utils.MustFloat64(strings.TrimSpace(tds.Eq(8).Text())),
		}
		items = append(items, item)
	})

	return items, nil
}

// RateLpr 获取LPR利率数据
//
// 目标地址: https://www.chinamoney.com.cn/chinese/bklpr/
//
// 返回:
//   - []LPRRate: LPR利率数据列表
//   - error: 错误信息
//
// 示例:
//
//	data, err := rate.RateLpr()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range data {
//	    fmt.Printf("%s 1年期LPR: %.2f%% 5年期LPR: %.2f%%\n", item.Date, item.LPR1Y, item.LPR5Y)
//	}
func RateLpr() ([]LPRRate, error) {
	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		"Referer":    "https://www.chinamoney.com.cn/",
	}

	resp, err := utils.GetWithHeaders(lprURL, nil, headers)
	if err != nil {
		return nil, fmt.Errorf("获取LPR数据失败: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	var items []LPRRate

	doc.Find("table tbody tr").Each(func(i int, s *goquery.Selection) {
		tds := s.Find("td")
		if tds.Length() < 3 {
			return
		}

		item := LPRRate{
			Date:  strings.TrimSpace(tds.Eq(0).Text()),
			LPR1Y: utils.MustFloat64(strings.TrimSpace(tds.Eq(1).Text())),
			LPR5Y: utils.MustFloat64(strings.TrimSpace(tds.Eq(2).Text())),
		}
		items = append(items, item)
	})

	return items, nil
}

// RepoRateItem 回购定盘利率数据结构
type RepoRateItem struct {
	Date   string  `json:"date"`   // 日期
	FR001  float64 `json:"fr001"`  // FR001
	FR007  float64 `json:"fr007"`  // FR007
	FR014  float64 `json:"fr014"`  // FR014
	FDR001 float64 `json:"fdr001"` // FDR001
	FDR007 float64 `json:"fdr007"` // FDR007
	FDR014 float64 `json:"fdr014"` // FDR014
}

// RepoRateQuery 获取回购定盘利率查询数据
//
// 目标地址: https://www.chinamoney.com.cn/chinese/bkfrr/
//
// 参数:
//   - symbol: "回购定盘利率" 或 "银银间回购定盘利率"
//
// 返回:
//   - []RepoRateItem: 回购定盘利率数据列表
//   - error: 错误信息
func RepoRateQuery(symbol string) ([]RepoRateItem, error) {
	var url string
	if symbol == "回购定盘利率" {
		url = "https://www.chinamoney.com.cn/r/cms/www/chinamoney/data/currency/frr-chrt.csv"
	} else {
		url = "https://www.chinamoney.com.cn/r/cms/www/chinamoney/data/currency/fdr-chrt.csv"
	}

	resp, err := utils.Get(url, nil)
	if err != nil {
		return nil, fmt.Errorf("获取回购定盘利率数据失败: %w", err)
	}

	lines := strings.Split(resp.String(), "\n")
	var items []RepoRateItem

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		if len(parts) < 4 {
			continue
		}

		item := RepoRateItem{
			Date: parts[0],
		}
		if symbol == "回购定盘利率" {
			item.FR001 = utils.MustFloat64(parts[1])
			item.FR007 = utils.MustFloat64(parts[2])
			item.FR014 = utils.MustFloat64(parts[3])
		} else {
			item.FDR001 = utils.MustFloat64(parts[1])
			item.FDR007 = utils.MustFloat64(parts[2])
			item.FDR014 = utils.MustFloat64(parts[3])
		}
		items = append(items, item)
	}

	return items, nil
}

// RepoRateHist 获取回购定盘利率历史数据
//
// 目标地址: https://www.chinamoney.com.cn/chinese/bkfrr/
//
// 参数:
//   - startDate: 开始日期，格式 "20200930"
//   - endDate: 结束日期，格式 "20201029"
//
// 返回:
//   - []RepoRateItem: 回购定盘利率历史数据列表
//   - error: 错误信息
func RepoRateHist(startDate, endDate string) ([]RepoRateItem, error) {
	// 格式化日期
	formattedStart := fmt.Sprintf("%s-%s-%s", startDate[:4], startDate[4:6], startDate[6:])
	formattedEnd := fmt.Sprintf("%s-%s-%s", endDate[:4], endDate[4:6], endDate[6:])

	url := "https://www.chinamoney.com.cn/ags/ms/cm-u-bk-currency/FrrHis"
	params := map[string]string{
		"lang":      "CN",
		"startDate": formattedStart,
		"endDate":   formattedEnd,
	}

	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
	}

	resp, err := utils.PostForm(url, params)
	if err != nil {
		// 尝试使用带headers的POST
		resp, err = utils.PostFormWithHeaders(url, params, headers)
		if err != nil {
			return nil, fmt.Errorf("获取回购定盘利率历史数据失败: %w", err)
		}
	}

	text := resp.String()
	records := gjson.Get(text, "records")
	if !records.Exists() {
		return nil, fmt.Errorf("解析回购定盘利率数据失败")
	}

	var items []RepoRateItem
	records.ForEach(func(_, value gjson.Result) bool {
		frValueMap := value.Get("frValueMap")
		if !frValueMap.Exists() {
			return true
		}
		item := RepoRateItem{
			Date:   frValueMap.Get("date").String(),
			FR001:  frValueMap.Get("FR001").Float(),
			FR007:  frValueMap.Get("FR007").Float(),
			FR014:  frValueMap.Get("FR014").Float(),
			FDR001: frValueMap.Get("FDR001").Float(),
			FDR007: frValueMap.Get("FDR007").Float(),
			FDR014: frValueMap.Get("FDR014").Float(),
		}
		items = append(items, item)
		return true
	})

	return items, nil
}
