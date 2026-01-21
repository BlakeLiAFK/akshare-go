package fund

import (
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/PuerkitoBio/goquery"
)

// FundOverviewEm 获取天天基金-基金档案-基本概况
// https://fundf10.eastmoney.com/jbgk_015641.html
// 参数:
//
//	symbol: 基金代码，例如 "015641"
//
// 返回:
//
//	基金基本概况数据
func FundOverviewEm(symbol string) (map[string]interface{}, error) {
	if symbol == "" {
		symbol = "015641"
	}

	url := fmt.Sprintf("https://fundf10.eastmoney.com/jbgk_%s.html", symbol)

	// 发送HTTP请求
	resp, err := utils.Get(url, nil)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	// 解析HTML
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	result := make(map[string]interface{})

	// 查找所有表格
	tables := doc.Find("table")
	if tables.Length() == 0 {
		return result, nil
	}

	// 获取最后一个表格（基本概况数据在最后一个表格）
	lastTable := tables.Last()

	// 遍历表格行，提取key-value数据
	lastTable.Find("tr").Each(func(i int, row *goquery.Selection) {
		tds := row.Find("td")
		if tds.Length() >= 4 {
			// 每行有4列：key1, value1, key2, value2
			key1 := strings.TrimSpace(tds.Eq(0).Text())
			val1 := strings.TrimSpace(tds.Eq(1).Text())
			key2 := strings.TrimSpace(tds.Eq(2).Text())
			val3 := strings.TrimSpace(tds.Eq(3).Text())

			if key1 != "" {
				result[key1] = val1
			}
			if key2 != "" {
				result[key2] = val3
			}
		} else if tds.Length() >= 2 {
			// 有些行只有2列
			key := strings.TrimSpace(tds.Eq(0).Text())
			val := strings.TrimSpace(tds.Eq(1).Text())
			if key != "" {
				result[key] = val
			}
		}
	})

	return result, nil
}
