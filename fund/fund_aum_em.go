package fund

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/PuerkitoBio/goquery"
)

// FundAumEm 东方财富-基金-基金公司排名列表
// https://fund.eastmoney.com/Company/lsgm.html
// 返回: 基金公司排名列表
func FundAumEm() ([]map[string]interface{}, error) {
	url := "https://fund.eastmoney.com/Company/home/gspmlist"
	params := map[string]string{
		"fundType": "0",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	// 解析HTML
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	records := make([]map[string]interface{}, 0)

	// 查找表格
	doc.Find("table").Each(func(i int, table *goquery.Selection) {
		table.Find("tbody tr").Each(func(rowIdx int, row *goquery.Selection) {
			cells := row.Find("td")
			if cells.Length() < 6 {
				return
			}

			// 提取单元格数据
			序号 := strings.TrimSpace(cells.Eq(0).Text())
			基金公司 := strings.TrimSpace(cells.Eq(1).Text())
			成立时间 := strings.TrimSpace(cells.Eq(2).Text())
			全部管理规模 := strings.TrimSpace(cells.Eq(3).Text())
			全部基金数 := strings.TrimSpace(cells.Eq(4).Text())
			全部经理数 := strings.TrimSpace(cells.Eq(5).Text())

			// 处理"全部管理规模"列，分离规模和更新日期
			规模部分 := strings.Split(全部管理规模, " ")
			管理规模值 := ""
			更新日期 := ""
			if len(规模部分) >= 1 {
				管理规模值 = strings.ReplaceAll(规模部分[0], ",", "")
			}
			if len(规模部分) >= 2 {
				更新日期 = 规模部分[1]
			}

			record := map[string]interface{}{
				"序号":     utils.MustInt64(序号),
				"基金公司":   基金公司,
				"成立时间":   成立时间,
				"全部管理规模": utils.MustFloat64(管理规模值),
				"全部基金数":  utils.MustInt64(全部基金数),
				"全部经理数":  utils.MustInt64(全部经理数),
				"更新日期":   更新日期,
			}
			records = append(records, record)
		})
	})

	return records, nil
}

// FundAumTrendEm 东方财富-基金-基金市场管理规模走势图
// https://fund.eastmoney.com/Company/default.html
// 返回: 基金市场管理规模走势图
func FundAumTrendEm() ([]map[string]interface{}, error) {
	url := "https://fund.eastmoney.com/Company/home/GetFundTotalScaleForChart"
	params := map[string]string{
		"fundType": "0",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	// 解析JSON
	var result struct {
		X []string  `json:"x"`
		Y []float64 `json:"y"`
	}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return nil, fmt.Errorf("解析JSON失败: %w", err)
	}

	records := make([]map[string]interface{}, 0)
	for i := 0; i < len(result.X) && i < len(result.Y); i++ {
		record := map[string]interface{}{
			"date":  result.X[i],
			"value": result.Y[i],
		}
		records = append(records, record)
	}

	return records, nil
}

// FundAumHistEm 东方财富-基金-基金公司历年管理规模排行列表
// https://fund.eastmoney.com/Company/lsgm.html
// 参数: year 年份，如 "2023"
// 返回: 基金公司历年管理规模排行列表
func FundAumHistEm(year string) ([]map[string]interface{}, error) {
	if year == "" {
		year = "2023"
	}

	url := "https://fund.eastmoney.com/Company/home/HistoryScaleTable"
	params := map[string]string{
		"year": year,
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	// 解析HTML
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	records := make([]map[string]interface{}, 0)

	// 查找表格
	doc.Find("table").Each(func(i int, table *goquery.Selection) {
		table.Find("tbody tr").Each(func(rowIdx int, row *goquery.Selection) {
			cells := row.Find("td")
			if cells.Length() < 8 {
				return
			}

			// 提取单元格数据
			序号 := strings.TrimSpace(cells.Eq(0).Text())
			基金公司 := strings.TrimSpace(cells.Eq(1).Text())
			总规模 := strings.TrimSpace(cells.Eq(2).Text())
			股票型 := strings.TrimSpace(cells.Eq(3).Text())
			混合型 := strings.TrimSpace(cells.Eq(4).Text())
			债券型 := strings.TrimSpace(cells.Eq(5).Text())
			指数型 := strings.TrimSpace(cells.Eq(6).Text())
			QDII := strings.TrimSpace(cells.Eq(7).Text())
			货币型 := strings.TrimSpace(cells.Eq(8).Text())

			record := map[string]interface{}{
				"序号":   utils.MustInt64(序号),
				"基金公司": 基金公司,
				"总规模":  utils.MustFloat64(总规模),
				"股票型":  utils.MustFloat64(股票型),
				"混合型":  utils.MustFloat64(混合型),
				"债券型":  utils.MustFloat64(债券型),
				"指数型":  utils.MustFloat64(指数型),
				"QDII": utils.MustFloat64(QDII),
				"货币型":  utils.MustFloat64(货币型),
			}
			records = append(records, record)
		})
	})

	return records, nil
}
