package fund

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/PuerkitoBio/goquery"
)

// FundFeeEm 获取天天基金-基金档案-购买信息
// https://fundf10.eastmoney.com/jjfl_015641.html
// indicator 可选值：
// "交易状态", "申购与赎回金额", "交易确认日", "运作费用",
// "认购费率（前端）", "认购费率（后端）", "申购费率（前端）", "赎回费率"
func FundFeeEm(symbol, indicator string) ([][]string, error) {
	if symbol == "" {
		symbol = "015641"
	}
	if indicator == "" {
		indicator = "认购费率"
	}

	url := fmt.Sprintf("https://fundf10.eastmoney.com/jjfl_%s.html", symbol)
	resp, err := utils.Get(url, nil)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	// 存储所有表格
	tablesDict := make(map[string][][]string)

	// 查找所有 h4.t 标题元素
	doc.Find("h4.t").Each(func(i int, title *goquery.Selection) {
		titleText := title.Text()
		// 清理空白字符
		re := regexp.MustCompile(`\s+`)
		titleText = strings.TrimSpace(re.ReplaceAllString(titleText, " "))

		var tableData [][]string

		if titleText == "申购与赎回金额" {
			// 特殊处理：需要合并两个表格
			tables := title.NextAllFiltered("table").First().Parent().Find("table")
			if tables.Length() >= 2 {
				// 解析第一个表格
				table1Data := parseTable(tables.Eq(0))
				// 解析第二个表格
				table2Data := parseTable(tables.Eq(1))
				// 合并数据
				tableData = append(table1Data, table2Data...)
			}
		} else {
			// 查找标题后的第一个 table
			table := title.NextFiltered("table")
			if table.Length() == 0 {
				// 如果直接后面没有table，尝试查找所有后续元素中的第一个table
				table = title.NextAll().Find("table").First()
			}
			if table.Length() > 0 {
				tableData = parseTable(table)
			}
		}

		if len(tableData) > 0 {
			tablesDict[titleText] = tableData
		}
	})

	// 根据 indicator 返回对应的表格数据
	tempData, ok := tablesDict[indicator]
	if !ok {
		return [][]string{}, nil
	}

	// 对某些 indicator 需要进行特殊处理
	if indicator == "认购费率（前端）" {
		return processFrontendSubscriptionFee(tempData), nil
	} else if indicator == "申购费率（前端）" {
		return processFrontendPurchaseFee(tempData), nil
	}

	return tempData, nil
}

// parseTable 解析HTML表格为二维字符串数组
func parseTable(table *goquery.Selection) [][]string {
	var result [][]string

	table.Find("tr").Each(func(i int, row *goquery.Selection) {
		var rowData []string
		row.Find("td, th").Each(func(j int, cell *goquery.Selection) {
			cellText := strings.TrimSpace(cell.Text())
			rowData = append(rowData, cellText)
		})
		if len(rowData) > 0 {
			result = append(result, rowData)
		}
	})

	return result
}

// processFrontendSubscriptionFee 处理认购费率（前端）的特殊格式
func processFrontendSubscriptionFee(data [][]string) [][]string {
	if len(data) == 0 {
		return data
	}

	// 检查是否有"原费率|天天基金优惠费率"列
	headers := data[0]
	targetColIdx := -1
	for i, header := range headers {
		if strings.Contains(header, "原费率") && strings.Contains(header, "天天基金优惠费率") {
			targetColIdx = i
			break
		}
	}

	if targetColIdx == -1 {
		return data
	}

	// 创建新的表格，添加拆分后的列
	result := make([][]string, 0, len(data))

	// 处理表头
	newHeaders := make([]string, 0, len(headers)+1)
	for i, header := range headers {
		if i == targetColIdx {
			newHeaders = append(newHeaders, "原费率", "天天基金优惠费率")
		} else {
			newHeaders = append(newHeaders, header)
		}
	}
	result = append(result, newHeaders)

	// 处理数据行
	for rowIdx := 1; rowIdx < len(data); rowIdx++ {
		row := data[rowIdx]
		newRow := make([]string, 0, len(newHeaders))

		for colIdx, cell := range row {
			if colIdx == targetColIdx {
				parts := strings.Split(cell, "|")
				if len(parts) >= 2 {
					newRow = append(newRow, strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]))
				} else {
					newRow = append(newRow, strings.TrimSpace(cell), strings.TrimSpace(cell))
				}
			} else {
				newRow = append(newRow, cell)
			}
		}
		result = append(result, newRow)
	}

	// 特殊处理：第4行（索引3）的天天基金优惠费率应该等于原费率
	if len(result) > 4 {
		originalFeeIdx := -1
		discountFeeIdx := -1
		for i, header := range result[0] {
			if header == "原费率" {
				originalFeeIdx = i
			} else if header == "天天基金优惠费率" {
				discountFeeIdx = i
			}
		}
		if originalFeeIdx != -1 && discountFeeIdx != -1 && len(result[4]) > discountFeeIdx {
			result[4][discountFeeIdx] = result[4][originalFeeIdx]
		}
	}

	return result
}

// processFrontendPurchaseFee 处理申购费率（前端）的特殊格式
func processFrontendPurchaseFee(data [][]string) [][]string {
	if len(data) == 0 {
		return data
	}

	// 检查是否有"原费率|天天基金优惠费率 银行卡购买|活期宝购买"列
	headers := data[0]
	targetColIdx := -1
	for i, header := range headers {
		if strings.Contains(header, "原费率") && strings.Contains(header, "银行卡购买") {
			targetColIdx = i
			break
		}
	}

	if targetColIdx == -1 {
		return data
	}

	// 创建新的表格，添加拆分后的列
	result := make([][]string, 0, len(data))

	// 处理表头
	newHeaders := make([]string, 0, len(headers)+2)
	for i, header := range headers {
		if i == targetColIdx {
			newHeaders = append(newHeaders, "原费率", "天天基金优惠费率-银行卡购买", "天天基金优惠费率-活期宝购买")
		} else {
			newHeaders = append(newHeaders, header)
		}
	}
	result = append(result, newHeaders)

	// 处理数据行
	for rowIdx := 1; rowIdx < len(data); rowIdx++ {
		row := data[rowIdx]
		newRow := make([]string, 0, len(newHeaders))

		for colIdx, cell := range row {
			if colIdx == targetColIdx {
				parts := strings.Split(cell, "|")
				if len(parts) == 1 {
					// 只有一个值，三列都用同一个值
					val := strings.TrimSpace(parts[0])
					newRow = append(newRow, val, val, val)
				} else if len(parts) == 3 {
					// 有三个值，分别赋值
					newRow = append(newRow,
						strings.TrimSpace(parts[0]),
						strings.TrimSpace(parts[1]),
						strings.TrimSpace(parts[2]))
				} else if len(parts) >= 2 {
					// 有两个或更多值
					newRow = append(newRow,
						strings.TrimSpace(parts[0]),
						strings.TrimSpace(parts[1]),
						strings.TrimSpace(parts[len(parts)-1]))
				} else {
					newRow = append(newRow, cell, cell, cell)
				}
			} else {
				newRow = append(newRow, cell)
			}
		}
		result = append(result, newRow)
	}

	// 特殊处理：某些行的优惠费率应该等于原费率
	originalFeeIdx := -1
	bankDiscountIdx := -1
	hqbDiscountIdx := -1
	for i, header := range result[0] {
		if header == "原费率" {
			originalFeeIdx = i
		} else if header == "天天基金优惠费率-银行卡购买" {
			bankDiscountIdx = i
		} else if header == "天天基金优惠费率-活期宝购买" {
			hqbDiscountIdx = i
		}
	}

	if originalFeeIdx != -1 && bankDiscountIdx != -1 && hqbDiscountIdx != -1 {
		// 对第3行（索引3）和第4行（索引4）进行特殊处理
		for _, rowIdx := range []int{3, 4} {
			if len(result) > rowIdx && len(result[rowIdx]) > hqbDiscountIdx {
				result[rowIdx][bankDiscountIdx] = result[rowIdx][originalFeeIdx]
				result[rowIdx][hqbDiscountIdx] = result[rowIdx][originalFeeIdx]
			}
		}
	}

	return result
}
