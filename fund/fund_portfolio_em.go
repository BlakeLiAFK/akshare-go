package fund

import (
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
)

// FundPortfolioHoldEm 获取基金股票持仓
func FundPortfolioHoldEm(symbol, date string) ([]map[string]interface{}, error) {
	if symbol == "" {
		symbol = "000001"
	}
	if date == "" {
		date = "2024"
	}

	url := "https://api.fund.eastmoney.com/f10/FundArchivesDatas"
	params := map[string]string{
		"type": "jjcc",
		"code": symbol,
		"year": date,
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	items := result.Get("Data.fundStocks").Array()

	records := make([]map[string]interface{}, 0, len(items))
	for _, item := range items {
		record := map[string]interface{}{
			"股票代码": item.Get("GPDM").String(),
			"股票名称": item.Get("GPJC").String(),
			"持仓市值": utils.MustFloat64(item.Get("SZ").String()),
		}
		records = append(records, record)
	}

	return records, nil
}

// FundPortfolioBondHoldEm 获取基金债券持仓
// https://fundf10.eastmoney.com/ccmx1_000001.html
func FundPortfolioBondHoldEm(symbol, date string) ([]map[string]interface{}, error) {
	if symbol == "" {
		symbol = "000001"
	}
	if date == "" {
		date = "2023"
	}

	url := "https://fundf10.eastmoney.com/FundArchivesDatas.aspx"
	params := map[string]string{
		"type": "zqcc",
		"code": symbol,
		"year": date,
		"rt":   "0.913877030254846",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	// 提取JSON部分（从第一个{到最后一个}）
	text := resp.String()
	startIdx := strings.Index(text, "{")
	if startIdx == -1 {
		return nil, fmt.Errorf("解析响应失败: 未找到JSON数据")
	}
	jsonText := text[startIdx:]

	// 使用gjson解析
	result := gjson.Parse(jsonText)
	content := result.Get("content").String()
	if content == "" {
		return []map[string]interface{}{}, nil
	}

	// 解析HTML内容
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	// 提取季度标签
	var quarters []string
	doc.Find("h4.t").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		// 分割文本，提取季度信息（格式: "xxx  第x季度"）
		parts := strings.Split(text, "\u00a0\u00a0") // \u00a0 是 &nbsp;
		if len(parts) >= 2 {
			quarters = append(quarters, strings.TrimSpace(parts[1]))
		}
	})

	records := make([]map[string]interface{}, 0)
	tableIdx := 0

	// 查找所有表格
	doc.Find("table").Each(func(i int, table *goquery.Selection) {
		if tableIdx >= len(quarters) {
			return
		}
		quarter := quarters[tableIdx]
		tableIdx++

		// 处理表格数据
		table.Find("tbody tr").Each(func(rowIdx int, row *goquery.Selection) {
			cells := row.Find("td")
			if cells.Length() < 5 {
				return
			}

			// 提取单元格数据
			序号 := strings.TrimSpace(cells.Eq(0).Text())
			债券代码 := strings.TrimSpace(cells.Eq(1).Text())
			债券名称 := strings.TrimSpace(cells.Eq(2).Text())
			占净值比例 := strings.TrimSpace(cells.Eq(3).Text())
			持仓市值 := strings.TrimSpace(cells.Eq(4).Text())

			// 处理占净值比例（去掉%）
			占净值比例 = strings.TrimSuffix(占净值比例, "%")

			record := map[string]interface{}{
				"序号":    utils.MustInt64(序号),
				"债券代码":  债券代码,
				"债券名称":  债券名称,
				"占净值比例": utils.MustFloat64(占净值比例),
				"持仓市值":  utils.MustFloat64(持仓市值),
				"季度":    quarter,
			}
			records = append(records, record)
		})
	})

	// 重新编号
	for i := range records {
		records[i]["序号"] = i + 1
	}

	return records, nil
}

// FundPortfolioIndustryAllocationEm 获取基金行业配置
func FundPortfolioIndustryAllocationEm(symbol, date string) ([]map[string]interface{}, error) {
	if symbol == "" {
		symbol = "000001"
	}
	if date == "" {
		date = "2024"
	}

	url := "https://api.fund.eastmoney.com/f10/HYPZ/"
	params := map[string]string{
		"fundCode": symbol,
		"year":     date,
	}

	headers := map[string]string{
		"Referer": "https://fundf10.eastmoney.com/",
	}

	resp, err := utils.GetWithHeaders(url, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	// 移除JSONP包装
	text := resp.String()
	startIdx := strings.Index(text, "(")
	endIdx := strings.LastIndex(text, ")")
	if startIdx == -1 || endIdx == -1 {
		return nil, fmt.Errorf("解析响应失败")
	}
	jsonText := text[startIdx+1 : endIdx]

	result := gjson.Parse(jsonText)
	if result.Get("ErrCode").Int() != 0 {
		return nil, fmt.Errorf("API返回错误")
	}

	records := make([]map[string]interface{}, 0)

	// 解析各个季度的数据
	quarters := []string{"Data_quarterAllocationValue", "Data_thirdQuarterAllocationValue",
		"Data_secondQuarterAllocationValue", "Data_firstQuarterAllocationValue"}
	quarterNames := []string{"第四季度", "第三季度", "第二季度", "第一季度"}

	for i, quarter := range quarters {
		items := result.Get(quarter).Array()
		for _, item := range items {
			record := map[string]interface{}{
				"季度":   quarterNames[i],
				"行业名称": item.Get("HYMC").String(),
				"占净值比": utils.MustFloat64(item.Get("ZJZBL").String()),
				"市值":   utils.MustFloat64(item.Get("SZ").String()),
			}
			records = append(records, record)
		}
	}

	return records, nil
}

// FundPortfolioChangeEm 获取基金持仓变动
// https://fundf10.eastmoney.com/ccbd_000001.html
// indicator: choice of {"累计买入", "累计卖出"}
func FundPortfolioChangeEm(symbol, indicator, date string) ([]map[string]interface{}, error) {
	if symbol == "" {
		symbol = "003567"
	}
	if indicator == "" {
		indicator = "累计买入"
	}
	if date == "" {
		date = "2023"
	}

	indicatorMap := map[string]string{
		"累计买入": "1",
		"累计卖出": "2",
	}

	zdbd, ok := indicatorMap[indicator]
	if !ok {
		return nil, fmt.Errorf("请输入正确的indicator参数: 累计买入, 累计卖出")
	}

	url := "https://fundf10.eastmoney.com/FundArchivesDatas.aspx"
	params := map[string]string{
		"type": "zdbd",
		"code": symbol,
		"zdbd": zdbd,
		"year": date,
		"rt":   "0.913877030254846",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	// 提取JSON部分
	text := resp.String()
	startIdx := strings.Index(text, "{")
	if startIdx == -1 {
		return nil, fmt.Errorf("解析响应失败: 未找到JSON数据")
	}
	jsonText := text[startIdx:]

	// 使用gjson解析
	result := gjson.Parse(jsonText)
	content := result.Get("content").String()
	if content == "" {
		return []map[string]interface{}{}, nil
	}

	// 解析HTML内容
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	// 提取季度标签
	var quarters []string
	doc.Find("h4.t").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		parts := strings.Split(text, "\u00a0\u00a0")
		if len(parts) >= 2 {
			quarters = append(quarters, strings.TrimSpace(parts[1]))
		}
	})

	records := make([]map[string]interface{}, 0)
	tableIdx := 0

	// 查找所有表格
	doc.Find("table").Each(func(i int, table *goquery.Selection) {
		if tableIdx >= len(quarters) {
			return
		}
		quarter := quarters[tableIdx]
		tableIdx++

		// 处理表格数据
		table.Find("tbody tr").Each(func(rowIdx int, row *goquery.Selection) {
			cells := row.Find("td")
			if cells.Length() < 5 {
				return
			}

			// 提取单元格数据（跳过"相关资讯"列）
			股票代码 := strings.TrimSpace(cells.Eq(1).Text())
			股票名称 := strings.TrimSpace(cells.Eq(2).Text())
			累计金额 := strings.TrimSpace(cells.Eq(3).Text())
			占净值比例 := strings.TrimSpace(cells.Eq(4).Text())

			// 处理占净值比例（去掉%）
			占净值比例 = strings.TrimSuffix(占净值比例, "%")

			record := map[string]interface{}{
				"股票代码":        股票代码,
				"股票名称":        股票名称,
				"本期累计买入金额":    utils.MustFloat64(累计金额),
				"占期初基金资产净值比例": utils.MustFloat64(占净值比例),
				"季度":          quarter,
			}
			records = append(records, record)
		})
	})

	// 重新编号
	for i := range records {
		records[i]["序号"] = i + 1
	}

	return records, nil
}
