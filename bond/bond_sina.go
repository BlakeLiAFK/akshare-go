package bond

import (
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-gota/gota/dataframe"
)

// BondCBProfileSina 新浪财经可转债详情资料
//
// 获取新浪财经网站的可转债详情资料
//
// 参数:
//   - symbol: 带市场标识的转债代码，如 "sz128039"
//
// 返回:
//   - dataframe.DataFrame: 包含可转债详情资料
//   - error: 错误信息
//
// 数据源: https://money.finance.sina.com.cn/bond/info/sz128039.html
func BondCBProfileSina(symbol string) (dataframe.DataFrame, error) {
	if symbol == "" {
		return dataframe.DataFrame{}, fmt.Errorf("转债代码不能为空")
	}

	url := fmt.Sprintf("https://money.finance.sina.com.cn/bond/info/%s.html", symbol)

	resp, err := utils.Get(url, nil)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	// 使用 goquery 解析 HTML
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("解析HTML失败: %w", err)
	}

	// 查找表格 - 新浪财经的可转债详情页使用 table 标签
	var records [][]string
	doc.Find("table").Each(func(i int, table *goquery.Selection) {
		table.Find("tr").Each(func(j int, tr *goquery.Selection) {
			var row []string
			tr.Find("td, th").Each(func(k int, cell *goquery.Selection) {
				text := strings.TrimSpace(cell.Text())
				row = append(row, text)
			})
			if len(row) > 0 {
				records = append(records, row)
			}
		})
	})

	if len(records) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未找到表格数据")
	}

	// 添加标准化列头
	headers := []string{"item", "value"}
	finalRecords := [][]string{headers}

	// 处理数据行，确保每行有2列
	for _, row := range records {
		if len(row) == 2 {
			finalRecords = append(finalRecords, row)
		} else if len(row) > 2 {
			// 如果有多列，取前两列
			finalRecords = append(finalRecords, row[:2])
		}
	}

	if len(finalRecords) <= 1 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到有效数据")
	}

	df := dataframe.LoadRecords(finalRecords)
	return df, nil
}

// BondCBSummarySina 新浪财经可转债债券概况
//
// 获取新浪财经网站的可转债债券概况
//
// 参数:
//   - symbol: 带市场标识的转债代码，如 "sh155255"
//
// 返回:
//   - dataframe.DataFrame: 包含可转债债券概况
//   - error: 错误信息
//
// 数据源: https://money.finance.sina.com.cn/bond/quotes/sh155255.html
func BondCBSummarySina(symbol string) (dataframe.DataFrame, error) {
	if symbol == "" {
		return dataframe.DataFrame{}, fmt.Errorf("转债代码不能为空")
	}

	url := fmt.Sprintf("https://money.finance.sina.com.cn/bond/quotes/%s.html", symbol)

	resp, err := utils.Get(url, nil)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	// 使用 goquery 解析 HTML
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("解析HTML失败: %w", err)
	}

	// 新浪财经的债券概况页面，第11个 table (索引10)包含概况数据
	var targetTable *goquery.Selection
	doc.Find("table").Each(func(i int, table *goquery.Selection) {
		if i == 10 {
			targetTable = table
		}
	})

	if targetTable == nil {
		return dataframe.DataFrame{}, fmt.Errorf("未找到债券概况表格")
	}

	// 提取表格数据
	var allRecords [][]string
	targetTable.Find("tr").Each(func(i int, tr *goquery.Selection) {
		var row []string
		tr.Find("td").Each(func(j int, td *goquery.Selection) {
			text := strings.TrimSpace(td.Text())
			row = append(row, text)
		})
		if len(row) > 0 {
			allRecords = append(allRecords, row)
		}
	})

	if len(allRecords) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到债券概况数据")
	}

	// 重新组织数据：将3组2列合并为单一的 item-value 对
	headers := []string{"item", "value"}
	var finalRecords [][]string
	finalRecords = append(finalRecords, headers)

	for _, row := range allRecords {
		// 每行可能有6列(3组 item-value)
		if len(row) >= 2 {
			// 第一组
			finalRecords = append(finalRecords, []string{row[0], row[1]})
		}
		if len(row) >= 4 {
			// 第二组
			finalRecords = append(finalRecords, []string{row[2], row[3]})
		}
		if len(row) >= 6 {
			// 第三组
			finalRecords = append(finalRecords, []string{row[4], row[5]})
		}
	}

	if len(finalRecords) <= 1 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到有效数据")
	}

	df := dataframe.LoadRecords(finalRecords)
	return df, nil
}
