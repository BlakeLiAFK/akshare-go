package news

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/PuerkitoBio/goquery"
)

// NewsCctv 新闻联播文字稿
// https://tv.cctv.com/lm/xwlb
// date: 需要获取数据的日期，格式 YYYYMMDD，目前支持 20160203 年后
func NewsCctv(date string) ([]CctvNews, error) {
	dateInt := 0
	fmt.Sscanf(date, "%d", &dateInt)

	var pageUrls []string
	var err error

	// 根据日期选择不同的URL模式
	if dateInt <= 20130708 {
		pageUrls, err = getCctvUrlsBefore2013(date)
	} else if dateInt < 20160203 {
		pageUrls, err = getCctvUrls2013To2016(date)
	} else {
		pageUrls, err = getCctvUrlsAfter2016(date)
	}

	if err != nil {
		return nil, err
	}

	// 获取每个页面的内容
	var newsList []CctvNews
	headers := getCctvHeaders()

	for _, pageUrl := range pageUrls {
		title, content, err := fetchCctvPage(pageUrl, headers)
		if err != nil {
			// 跳过错误的页面，与 Python 版本行为一致
			continue
		}
		newsList = append(newsList, CctvNews{
			Date:    date,
			Title:   title,
			Content: content,
		})
	}

	return newsList, nil
}

// getCctvHeaders 获取 CCTV 请求头
func getCctvHeaders() map[string]string {
	return map[string]string{
		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
		"Accept-Encoding":           "gzip, deflate",
		"Accept-Language":           "zh-CN,zh;q=0.9,en;q=0.8",
		"Cache-Control":             "no-cache",
		"Cookie":                    "cna=DLYSGBDthG4CAbRVCNxSxGT6",
		"Pragma":                    "no-cache",
		"Upgrade-Insecure-Requests": "1",
		"User-Agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36",
	}
}

// getCctvUrlsBefore2013 获取 2013 年之前的新闻链接
func getCctvUrlsBefore2013(date string) ([]string, error) {
	url := fmt.Sprintf("https://cctv.cntv.cn/lm/xinwenlianbo/%s.shtml", date)

	resp, err := utils.Get(url, nil)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	text := resp.String()
	// 提取页面链接
	re := regexp.MustCompile(`title_array_01\((.*)\)`)
	matches := re.FindAllStringSubmatch(text, -1)

	var urls []string
	urlRe := regexp.MustCompile(`(http[^']+)`)
	for i, match := range matches {
		if i == 0 {
			continue // 跳过第一个
		}
		if len(match) > 1 {
			urlMatch := urlRe.FindStringSubmatch(match[1])
			if len(urlMatch) > 1 {
				urls = append(urls, urlMatch[1])
			}
		}
	}

	return urls, nil
}

// getCctvUrls2013To2016 获取 2013-2016 年的新闻链接
func getCctvUrls2013To2016(date string) ([]string, error) {
	url := fmt.Sprintf("https://cctv.cntv.cn/lm/xinwenlianbo/%s.shtml", date)

	resp, err := utils.Get(url, nil)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	var urls []string
	doc.Find("#contentELMT1368521805488378 li").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			return // 跳过第一个
		}
		if href, exists := s.Find("a").Attr("href"); exists {
			urls = append(urls, href)
		}
	})

	return urls, nil
}

// getCctvUrlsAfter2016 获取 2016 年之后的新闻链接
func getCctvUrlsAfter2016(date string) ([]string, error) {
	url := fmt.Sprintf("https://tv.cctv.com/lm/xwlb/day/%s.shtml", date)

	resp, err := utils.Get(url, nil)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	var urls []string
	doc.Find("li").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			return // 跳过第一个
		}
		if href, exists := s.Find("a").Attr("href"); exists {
			urls = append(urls, href)
		}
	})

	return urls, nil
}

// fetchCctvPage 获取单个新闻页面的标题和内容
func fetchCctvPage(pageUrl string, headers map[string]string) (string, string, error) {
	resp, err := utils.GetWithHeaders(pageUrl, nil, headers)
	if err != nil {
		return "", "", err
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return "", "", err
	}

	// 获取标题
	var title string
	if h3 := doc.Find("h3").First(); h3.Length() > 0 {
		title = h3.Text()
	} else if tit := doc.Find("div.tit").First(); tit.Length() > 0 {
		title = tit.Text()
	}

	// 获取内容
	var content string
	if cntBd := doc.Find("div.cnt_bd").First(); cntBd.Length() > 0 {
		content = cntBd.Text()
	} else if contentArea := doc.Find("div.content_area").First(); contentArea.Length() > 0 {
		content = contentArea.Text()
	}

	// 清理标题
	title = strings.TrimPrefix(title, "[视频]")
	title = strings.TrimSpace(title)
	title = strings.ReplaceAll(title, "\n", " ")

	// 清理内容
	content = strings.TrimSpace(content)
	content = strings.TrimPrefix(content, "央视网消息(新闻联播)：")
	content = strings.TrimPrefix(content, "央视网消息（新闻联播）：")
	content = strings.TrimPrefix(content, "(新闻联播)：")
	content = strings.TrimSpace(content)
	content = strings.ReplaceAll(content, "\n", " ")

	return title, content, nil
}
