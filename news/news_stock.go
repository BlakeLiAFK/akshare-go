package news

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// StockNewsEm 东方财富-个股新闻-最近 100 条新闻
// https://so.eastmoney.com/news/s?keyword=603777
// symbol: 股票代码
func StockNewsEm(symbol string) ([]StockNews, error) {
	url := "https://search-api-web.eastmoney.com/search/jsonp"

	// 构建内部参数
	innerParam := map[string]interface{}{
		"uid":           "",
		"keyword":       symbol,
		"type":          []string{"cmsArticleWebOld"},
		"client":        "web",
		"clientType":    "web",
		"clientVersion": "curr",
		"param": map[string]interface{}{
			"cmsArticleWebOld": map[string]interface{}{
				"searchScope": "default",
				"sort":        "default",
				"pageIndex":   1,
				"pageSize":    100,
				"preTag":      "<em>",
				"postTag":     "</em>",
			},
		},
	}

	innerParamBytes, err := json.Marshal(innerParam)
	if err != nil {
		return nil, fmt.Errorf("序列化参数失败: %w", err)
	}

	params := map[string]string{
		"cb":    "jQuery35101792940631092459_1764599530165",
		"param": string(innerParamBytes),
		"_":     "1764599530176",
	}

	headers := map[string]string{
		"Accept":          "*/*",
		"Accept-Encoding": "gzip, deflate, br, zstd",
		"Accept-Language": "en,zh-CN;q=0.9,zh;q=0.8",
		"Cache-Control":   "no-cache",
		"Connection":      "keep-alive",
		"Host":            "search-api-web.eastmoney.com",
		"Pragma":          "no-cache",
		"Referer":         fmt.Sprintf("https://so.eastmoney.com/news/s?keyword=%s", symbol),
		"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36",
	}

	resp, err := utils.GetWithHeaders(url, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	// 解析 JSONP 响应
	text := resp.String()
	// 移除 JSONP 包装: jQuery35101792940631092459_1764599530165(...);
	re := regexp.MustCompile(`^jQuery\d+_\d+\((.*)\);?$`)
	matches := re.FindStringSubmatch(text)
	if len(matches) < 2 {
		// 尝试简单的方式解析
		start := strings.Index(text, "(")
		end := strings.LastIndex(text, ")")
		if start == -1 || end == -1 || start >= end {
			return nil, fmt.Errorf("解析JSONP响应失败")
		}
		text = text[start+1 : end]
	} else {
		text = matches[1]
	}

	// 解析 JSON
	result := gjson.Get(text, "result.cmsArticleWebOld")
	if !result.Exists() {
		return nil, fmt.Errorf("未找到新闻数据")
	}

	var newsList []StockNews
	// 用于清理 HTML 标签的正则
	emTagRe := regexp.MustCompile(`</?em>`)
	emParenRe := regexp.MustCompile(`\(<em>|</em>\)`)
	spaceRe := regexp.MustCompile(`\x{3000}`)

	result.ForEach(func(key, value gjson.Result) bool {
		code := value.Get("code").String()
		title := value.Get("title").String()
		content := value.Get("content").String()
		date := value.Get("date").String()
		mediaName := value.Get("mediaName").String()

		// 构建新闻链接
		newsUrl := fmt.Sprintf("http://finance.eastmoney.com/a/%s.html", code)

		// 清理标题中的 HTML 标签
		title = emParenRe.ReplaceAllString(title, "")
		title = emTagRe.ReplaceAllString(title, "")

		// 清理内容中的 HTML 标签和特殊字符
		content = emParenRe.ReplaceAllString(content, "")
		content = emTagRe.ReplaceAllString(content, "")
		content = spaceRe.ReplaceAllString(content, "")
		content = strings.ReplaceAll(content, "\r\n", " ")

		newsList = append(newsList, StockNews{
			Keyword:     symbol,
			Title:       title,
			Content:     content,
			PublishTime: date,
			Source:      mediaName,
			Url:         newsUrl,
		})
		return true
	})

	return newsList, nil
}
