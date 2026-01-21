package stock_fundamental

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/PuerkitoBio/goquery"
)

// StockHkProfitForecastEt 经济通-港股盈利预测
//
// 获取经济通网站的港股盈利预测数据，支持多种指标类型
//
// 参数:
//   - symbol: 港股代码，如 "09999"
//   - indicator: 指标类型，可选值:
//   - "评级总览"
//   - "去年度业绩表现"
//   - "综合盈利预测"
//   - "盈利预测概览"
//
// 返回:
//   - []map[string]interface{}: 盈利预测数据（动态列）
//   - error: 错误信息
//
// 示例:
//
//	// 获取评级总览
//	forecast, err := stock_fundamental.StockHkProfitForecastEt("09999", "评级总览")
//	if err != nil {
//	    log.Fatal(err)
//	}
func StockHkProfitForecastEt(symbol, indicator string) ([]map[string]interface{}, error) {
	if symbol == "" {
		return nil, fmt.Errorf("股票代码不能为空")
	}

	validIndicators := map[string]bool{
		"评级总览":    true,
		"去年度业绩表现": true,
		"综合盈利预测":  true,
		"盈利预测概览":  true,
	}

	if !validIndicators[indicator] {
		return nil, fmt.Errorf("无效的indicator参数，可选值: %v", getMapKeys(validIndicators))
	}

	url := "https://www.etnet.com.hk/www/sc/stocks/realtime/quote_profit.php"

	// 移除股票代码前导零
	codeInt := parseIntString(symbol)
	params := map[string]string{
		"code": strconv.Itoa(codeInt),
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求经济通盈利预测失败: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	var results []map[string]interface{}

	switch indicator {
	case "评级总览":
		results, err = parseRatingOverviewTable(doc)
	case "去年度业绩表现":
		results, err = parsePastPerformanceTable(doc)
	case "综合盈利预测":
		results, err = parseComprehensiveForecastTable(doc)
	case "盈利预测概览":
		results, err = parseForecastOverviewTable(doc)
	}

	if err != nil {
		return nil, err
	}

	return results, nil
}

// parseIntString 解析整数字符串，移除前导零
func parseIntString(s string) int {
	s = strings.TrimSpace(s)
	val, _ := strconv.ParseInt(s, 10, 64)
	return int(val)
}

// parseRatingOverviewTable 解析评级总览表格（索引0）
func parseRatingOverviewTable(doc *goquery.Document) ([]map[string]interface{}, error) {
	var results []map[string]interface{}

	tables := doc.Find("table")
	if tables.Length() < 1 {
		return results, nil
	}

	// 获取第一个表格的第一行第一列
	firstCell := tables.Eq(0).Find("tr").First().Find("td").First()
	text := strings.TrimSpace(firstCell.Text())

	if text == "" {
		return results, nil
	}

	// 按空格分割
	parts := strings.Fields(text)
	var cleanParts []string
	for _, part := range parts {
		if part != "" && part != "平均评级" {
			cleanParts = append(cleanParts, part)
		}
	}

	if len(cleanParts) >= 3 {
		rowData := map[string]interface{}{
			"方向":   cleanParts[0],
			"评级数量": parseInt(cleanParts[1]),
			"平均评级": cleanParts[2],
			"指标类型": "评级总览",
		}
		results = append(results, rowData)
	}

	return results, nil
}

// parsePastPerformanceTable 解析去年度业绩表现表格（索引2）
func parsePastPerformanceTable(doc *goquery.Document) ([]map[string]interface{}, error) {
	var results []map[string]interface{}

	tables := doc.Find("table")
	if tables.Length() < 3 {
		return results, nil
	}

	table := tables.Eq(2)
	table.Find("tr").Each(func(i int, row *goquery.Selection) {
		var cells []string
		row.Find("td").Each(func(j int, cell *goquery.Selection) {
			text := strings.TrimSpace(cell.Text())
			cells = append(cells, text)
		})

		// 处理前两列（左半部分）
		if len(cells) >= 2 && cells[0] != "" && cells[1] != "" {
			rowData := map[string]interface{}{
				"item":  cells[0],
				"value": cells[1],
				"指标类型":  "去年度业绩表现",
			}
			results = append(results, rowData)
		}

		// 处理后两列（右半部分，从索引3开始）
		if len(cells) >= 4 && cells[2] != "" && cells[3] != "" {
			rowData := map[string]interface{}{
				"item":  cells[2],
				"value": cells[3],
				"指标类型":  "去年度业绩表现",
			}
			results = append(results, rowData)
		}
	})

	return results, nil
}

// parseComprehensiveForecastTable 解析综合盈利预测表格（索引3）
func parseComprehensiveForecastTable(doc *goquery.Document) ([]map[string]interface{}, error) {
	var results []map[string]interface{}

	tables := doc.Find("table")
	if tables.Length() < 4 {
		return results, nil
	}

	table := tables.Eq(3)

	// 获取表头
	var headers []string
	table.Find("tr").First().Find("th").Each(func(j int, th *goquery.Selection) {
		header := strings.TrimSpace(th.Text())
		header = normalizeColumnName(header)
		headers = append(headers, header)
	})

	if len(headers) == 0 {
		// 尝试使用 td 作为表头
		table.Find("tr").First().Find("td").Each(func(j int, td *goquery.Selection) {
			header := strings.TrimSpace(td.Text())
			header = normalizeColumnName(header)
			headers = append(headers, header)
		})
	}

	table.Find("tr").Each(func(i int, row *goquery.Selection) {
		if i == 0 {
			return
		}

		var cells []string
		row.Find("td").Each(func(j int, cell *goquery.Selection) {
			text := strings.TrimSpace(cell.Text())
			cells = append(cells, text)
		})

		if len(cells) < 2 {
			return
		}

		rowData := map[string]interface{}{
			"指标类型": "综合盈利预测",
		}

		for j, cell := range cells {
			if j < len(headers) {
				key := headers[j]
				if key == "" {
					key = fmt.Sprintf("列%d", j)
				}
				// 数值列转换为浮点数
				if j > 0 {
					rowData[key] = parseFloat(cell)
				} else {
					rowData[key] = cell
				}
			}
		}

		results = append(results, rowData)
	})

	return results, nil
}

// parseForecastOverviewTable 解析盈利预测概览表格（索引4）
func parseForecastOverviewTable(doc *goquery.Document) ([]map[string]interface{}, error) {
	var results []map[string]interface{}

	tables := doc.Find("table")
	if tables.Length() < 5 {
		return results, nil
	}

	table := tables.Eq(4)

	// 获取表头
	var headers []string
	table.Find("tr").First().Find("th").Each(func(j int, th *goquery.Selection) {
		header := strings.TrimSpace(th.Text())
		header = normalizeColumnName(header)
		// 跳过重复的列
		if header != "目标价* (港元).1" {
			headers = append(headers, header)
		}
	})

	if len(headers) == 0 {
		table.Find("tr").First().Find("td").Each(func(j int, td *goquery.Selection) {
			header := strings.TrimSpace(td.Text())
			header = normalizeColumnName(header)
			if header != "目标价* (港元).1" {
				headers = append(headers, header)
			}
		})
	}

	table.Find("tr").Each(func(i int, row *goquery.Selection) {
		if i == 0 {
			return
		}

		var cells []string
		row.Find("td").Each(func(j int, cell *goquery.Selection) {
			text := strings.TrimSpace(cell.Text())
			cells = append(cells, text)
		})

		if len(cells) < 2 {
			return
		}

		// 跳过空行
		hasData := false
		for _, cell := range cells {
			if cell != "" {
				hasData = true
				break
			}
		}
		if !hasData {
			return
		}

		rowData := map[string]interface{}{
			"指标类型": "盈利预测概览",
		}

		colIndex := 0
		for j, cell := range cells {
			// 跳过重复列（索引8）
			if j == 8 {
				continue
			}

			if colIndex < len(headers) {
				key := headers[colIndex]
				if key == "" {
					key = fmt.Sprintf("列%d", colIndex)
				}

				// 根据列类型处理值
				switch key {
				case "财政年度":
					// 转换为字符串
					rowData[key] = cell
				case "更新日期":
					// 尝试解析日期
					rowData[key] = parseDateEt(cell)
				default:
					// 数值列
					rowData[key] = parseFloat(cell)
				}
			}
			colIndex++
		}

		results = append(results, rowData)
	})

	return results, nil
}

// normalizeColumnName 标准化列名
func normalizeColumnName(name string) string {
	// 替换各种列名为标准名称
	replacements := map[string]string{
		"纯利/(亏损)  (百万元人民币)": "纯利/亏损",
		"纯利/(亏损)  (百万港元)":   "纯利/亏损",
		"每股盈利/  (亏损)(分)":    "每股盈利/每股亏损",
		"每股盈利/  (亏损)(港仙)":   "每股盈利/每股亏损",
		"每股派息  (分)":         "每股派息",
		"每股派息  (港仙)":        "每股派息",
		"每股资产净值  (人民币元)":    "每股资产净值",
		"每股资产净值  (港元)":      "每股资产净值",
		"最高  (百万元人民币)":      "最高",
		"最高  (百万港元)":        "最高",
		"最低  (百万元人民币)":      "最低",
		"最低  (百万港元)":        "最低",
		"每股盈利*/ (亏损)  (港仙)": "每股盈利",
		"每股盈利*/ (亏损)  (分)":  "每股盈利",
		"每股派息*  (分)":        "每股派息",
		"每股派息*  (港仙)":       "每股派息",
		"目标价* (港元)":         "目标价",
		"更新日期":              "更新日期",
		"财政年度":              "财政年度",
	}

	// 精确匹配
	if replacement, ok := replacements[name]; ok {
		return replacement
	}

	return name
}

// parseDateEt 解析经济通日期格式
func parseDateEt(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return ""
	}

	// 尝试解析多种日期格式
	layouts := []string{
		"2006-01-02",
		"2006/01/02",
		"02/01/2006",
		"01/02/2006",
	}

	for _, layout := range layouts {
		if t, err := time.Parse(layout, s); err == nil {
			return t.Format("2006-01-02")
		}
	}

	return s
}

// cleanCellValueEt 清理经济通单元格值
func cleanCellValueEt(s string) string {
	s = regexp.MustCompile(`\s+`).ReplaceAllString(s, "")
	return strings.TrimSpace(s)
}
