package stock_feature

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
)

// GetTokenLg 生成乐咕的 token
func GetTokenLg() string {
	currentDate := time.Now().Format("2006-01-02")
	hash := md5.Sum([]byte(currentDate))
	return hex.EncodeToString(hash[:])
}

// GetCookieCsrf 获取乐咕的 Cookie 和 CSRF Token
func GetCookieCsrf(pageURL string) (map[string]string, error) {
	resp, err := utils.Get(pageURL, nil)
	if err != nil {
		return nil, fmt.Errorf("获取页面失败: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	csrfToken := ""
	doc.Find("meta[name='_csrf']").Each(func(i int, s *goquery.Selection) {
		if content, exists := s.Attr("content"); exists {
			csrfToken = content
		}
	})

	return map[string]string{
		"X-CSRF-Token": csrfToken,
	}, nil
}

// StockHkIndicatorEniu 亿牛网-港股指标
// symbol: 港股代码，如 "hk01093"
// indicator: 指标类型，可选 "港股", "市盈率", "市净率", "股息率", "ROE", "市值"
func StockHkIndicatorEniu(symbol, indicator string) ([]map[string]interface{}, error) {
	headers := map[string]string{
		"User-Agent": utils.DefaultUserAgent,
	}

	if indicator == "港股" {
		url := "https://eniu.com/static/data/stock_list.json"
		resp, err := utils.GetWithHeaders(url, nil, headers)
		if err != nil {
			return nil, fmt.Errorf("获取港股列表失败: %w", err)
		}

		var items []map[string]interface{}
		result := gjson.ParseBytes(resp.Body())
		result.ForEach(func(key, value gjson.Result) bool {
			stockID := value.Get("stock_id").String()
			if strings.HasPrefix(stockID, "hk") {
				items = append(items, value.Value().(map[string]interface{}))
			}
			return true
		})
		return items, nil
	}

	var url string
	switch indicator {
	case "市盈率":
		url = fmt.Sprintf("https://eniu.com/chart/peh/%s", symbol)
	case "市净率":
		url = fmt.Sprintf("https://eniu.com/chart/pbh/%s", symbol)
	case "股息率":
		url = fmt.Sprintf("https://eniu.com/chart/dvh/%s", symbol)
	case "ROE":
		url = fmt.Sprintf("https://eniu.com/chart/roeh/%s", symbol)
	default:
		url = fmt.Sprintf("https://eniu.com/chart/marketvalueh/%s", symbol)
	}

	resp, err := utils.GetWithHeaders(url, nil, headers)
	if err != nil {
		return nil, fmt.Errorf("获取港股指标失败: %w", err)
	}

	var items []map[string]interface{}
	result := gjson.ParseBytes(resp.Body())
	result.ForEach(func(key, value gjson.Result) bool {
		if value.IsObject() {
			if m, ok := value.Value().(map[string]interface{}); ok {
				items = append(items, m)
			}
		} else if value.IsArray() {
			// 处理数组元素
			value.ForEach(func(_, v gjson.Result) bool {
				if m, ok := v.Value().(map[string]interface{}); ok {
					items = append(items, m)
				}
				return true
			})
		}
		return true
	})

	return items, nil
}
