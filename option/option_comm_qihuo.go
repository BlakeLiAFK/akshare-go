package option

import (
	"crypto/tls"
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
)

// OptionCommSymbol 获取商品期权品种列表
// https://www.9qihuo.com/qiquanshouxufei
func OptionCommSymbol() ([]OptionCommSymbolItem, error) {
	url := "https://www.9qihuo.com/qiquanshouxufei"

	client := resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	resp, err := client.R().Get(url)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	var result []OptionCommSymbolItem
	doc.Find("div#inst_list a").Each(func(i int, s *goquery.Selection) {
		name := strings.TrimSpace(s.Text())
		href, exists := s.Attr("href")
		if exists && strings.Contains(href, "?") {
			parts := strings.Split(href, "?")
			if len(parts) > 1 {
				params := strings.Split(parts[1], "=")
				if len(params) > 1 {
					result = append(result, OptionCommSymbolItem{
						Name: name,
						Code: params[1],
					})
				}
			}
		}
	})

	return result, nil
}

// OptionCommInfo 获取商品期权手续费信息
// https://www.9qihuo.com/qiquanshouxufei
// symbol: 品种名称，如 "工业硅期权"
func OptionCommInfo(symbol string) ([]OptionCommInfoItem, error) {
	// 获取品种代码
	symbols, err := OptionCommSymbol()
	if err != nil {
		return nil, err
	}

	var symbolCode string
	for _, s := range symbols {
		if strings.Contains(s.Name, symbol) {
			symbolCode = s.Code
			break
		}
	}

	if symbolCode == "" {
		return nil, fmt.Errorf("未找到品种: %s", symbol)
	}

	url := fmt.Sprintf("https://www.9qihuo.com/qiquanshouxufei?heyue=%s", symbolCode)

	client := resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	resp, err := client.R().Get(url)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	// 获取交易所名称
	var exchange string
	doc.Find("table").First().Find("tr").First().Find("td").First().Each(func(i int, s *goquery.Selection) {
		exchange = strings.TrimSpace(s.Text())
	})

	// 获取更新时间
	var commUpdateTime, priceUpdateTime string
	doc.Find("a#dlink").Each(func(i int, s *goquery.Selection) {
		prevText := s.Parent().Text()
		if strings.Contains(prevText, "手续费更新时间：") {
			parts := strings.Split(prevText, "，")
			if len(parts) >= 2 {
				commUpdateTime = strings.TrimPrefix(parts[0], "（手续费更新时间：")
				priceUpdateTime = strings.TrimSuffix(strings.TrimPrefix(parts[1], "价格更新时间："), "。）")
			}
		}
	})

	var result []OptionCommInfoItem
	doc.Find("table").First().Find("tr").Each(func(i int, s *goquery.Selection) {
		if i < 3 {
			return // 跳过表头
		}

		tds := s.Find("td")
		if tds.Length() < 5 {
			return
		}

		contract := strings.TrimSpace(tds.Eq(0).Text())
		if contract == "" {
			return
		}

		result = append(result, OptionCommInfoItem{
			Contract:        contract,
			CurrentPrice:    utils.MustParseFloat(strings.TrimSpace(tds.Eq(1).Text())),
			Volume:          utils.MustParseInt(strings.TrimSpace(tds.Eq(2).Text())),
			TickProfit:      utils.MustParseFloat(strings.TrimSpace(tds.Eq(3).Text())),
			TickNetProfit:   utils.MustParseFloat(strings.TrimSpace(tds.Eq(4).Text())),
			Exchange:        exchange,
			CommUpdateTime:  commUpdateTime,
			PriceUpdateTime: priceUpdateTime,
		})
	})

	return result, nil
}
