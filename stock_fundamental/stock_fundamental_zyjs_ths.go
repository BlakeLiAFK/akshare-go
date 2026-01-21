package stock_fundamental

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// StockZyjsThs 同花顺-主营介绍
//
// 获取同花顺网站的主营介绍数据，包含产品名称、业务类型等信息
//
// 参数:
//   - symbol: 股票代码，如 "000066"
//
// 返回:
//   - map[string]string: 主营介绍数据，键值对形式，如 "产品名称": "业务类型"
//   - error: 错误信息
//
// 示例:
//
//	zyjs, err := stock_fundamental.StockZyjsThs("000066")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for key, value := range zyjs {
//	    fmt.Printf("%s: %s\n", key, value)
//	}
func StockZyjsThs(symbol string) (map[string]string, error) {
	if symbol == "" {
		return nil, fmt.Errorf("股票代码不能为空")
	}

	url := fmt.Sprintf("https://basic.10jqka.com.cn/new/%s/operate.html", symbol)

	resp, err := utils.Get(url, nil)
	if err != nil {
		return nil, fmt.Errorf("请求同花顺主营介绍失败: %w", err)
	}

	// GB2312 解码
	reader := transform.NewReader(strings.NewReader(resp.String()), simplifiedchinese.GBK.NewDecoder())
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	result := make(map[string]string)

	// 查找 ul.main_intro_list
	doc.Find("ul.main_intro_list").Each(func(i int, s *goquery.Selection) {
		// 提取每个 li 元素
		s.Find("li").Each(func(j int, li *goquery.Selection) {
			text := strings.TrimSpace(li.Text())
			if text == "" {
				return
			}

			// 按 "：" 分割
			parts := strings.SplitN(text, "：", 2)
			if len(parts) == 2 {
				key := strings.TrimSpace(parts[0])
				value := cleanValue(parts[1])
				if key != "" {
					result[key] = value
				}
			}
		})
	})

	if len(result) == 0 {
		return nil, fmt.Errorf("未找到主营介绍数据")
	}

	// 添加股票代码
	result["股票代码"] = symbol

	return result, nil
}

// cleanValue 清理值字符串，去除制表符、换行符、多余空格
func cleanValue(s string) string {
	// 移除制表符和换行符
	s = strings.ReplaceAll(s, "\t", "")
	s = strings.ReplaceAll(s, "\n", "")
	// 移除所有空格
	s = utils.RemoveAllSpaces(s)
	return strings.TrimSpace(s)
}

// StockProfitForecastThs 同花顺-盈利预测
//
// 获取同花顺网站的盈利预测数据，支持多种指标类型
//
// 参数:
//   - symbol: 股票代码，如 "600519"
//   - indicator: 指标类型，可选值:
//   - "预测年报每股收益"
//   - "预测年报净利润"
//   - "业绩预测详表-机构"
//   - "业绩预测详表-详细指标预测"
//
// 返回:
//   - []map[string]interface{}: 盈利预测数据（动态列）
//   - error: 错误信息
//
// 示例:
//
//	// 获取预测年报每股收益
//	forecast, err := stock_fundamental.StockProfitForecastThs("600519", "预测年报每股收益")
//	if err != nil {
//	    log.Fatal(err)
//	}
func StockProfitForecastThs(symbol, indicator string) ([]map[string]interface{}, error) {
	if symbol == "" {
		return nil, fmt.Errorf("股票代码不能为空")
	}

	validIndicators := map[string]bool{
		"预测年报每股收益":      true,
		"预测年报净利润":       true,
		"业绩预测详表-机构":     true,
		"业绩预测详表-详细指标预测": true,
	}

	if !validIndicators[indicator] {
		return nil, fmt.Errorf("无效的indicator参数，可选值: %v", getMapKeys(validIndicators))
	}

	url := fmt.Sprintf("https://basic.10jqka.com.cn/new/%s/worth.html", symbol)

	resp, err := utils.Get(url, nil)
	if err != nil {
		return nil, fmt.Errorf("请求同花顺盈利预测失败: %w", err)
	}

	// GBK 解码
	reader := transform.NewReader(strings.NewReader(resp.String()), simplifiedchinese.GBK.NewDecoder())
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	// 检查是否包含 "本年度暂无机构做出业绩预测"
	htmlText := resp.String()
	hasNoForecast := strings.Contains(htmlText, "本年度暂无机构做出业绩预测")

	var results []map[string]interface{}

	if hasNoForecast {
		// 处理无预测数据的情况
		if indicator == "预测年报每股收益" || indicator == "预测年报净利润" {
			return results, nil // 返回空数组
		} else if indicator == "业绩预测详表-机构" {
			results, err = parseProfitForecastOrgTable(doc, 0)
		} else if indicator == "业绩预测详表-详细指标预测" {
			results, err = parseProfitForecastDetailTable(doc, 1)
		}
	} else {
		// 有预测数据的情况
		if indicator == "预测年报每股收益" {
			results, err = parseProfitForecastEPSTable(doc, 0)
		} else if indicator == "预测年报净利润" {
			results, err = parseProfitForecastProfitTable(doc, 1)
		} else if indicator == "业绩预测详表-机构" {
			results, err = parseProfitForecastOrgTable(doc, 2)
		} else if indicator == "业绩预测详表-详细指标预测" {
			results, err = parseProfitForecastDetailTable(doc, 3)
		}
	}

	if err != nil {
		return nil, err
	}

	return results, nil
}

// parseProfitForecastEPSTable 解析预测年报每股收益表格
func parseProfitForecastEPSTable(doc *goquery.Document, tableIndex int) ([]map[string]interface{}, error) {
	var results []map[string]interface{}

	tables := doc.Find("table")
	if tables.Length() <= tableIndex {
		return results, nil
	}

	tables.Eq(tableIndex).Find("tr").Each(func(i int, row *goquery.Selection) {
		if i == 0 {
			// 跳过表头
			return
		}

		var cells []string
		row.Find("td").Each(func(j int, cell *goquery.Selection) {
			cells = append(cells, strings.TrimSpace(cell.Text()))
		})

		if len(cells) < 2 {
			return
		}

		rowData := map[string]interface{}{
			"年度":   cells[0],
			"机构数":  parseInt(cells[1]),
			"平均值":  parseFloat(cells[2]),
			"最大值":  parseFloat(cells[3]),
			"最小值":  parseFloat(cells[4]),
			"指标类型": "预测年报每股收益",
		}

		results = append(results, rowData)
	})

	return results, nil
}

// parseProfitForecastProfitTable 解析预测年报净利润表格
func parseProfitForecastProfitTable(doc *goquery.Document, tableIndex int) ([]map[string]interface{}, error) {
	var results []map[string]interface{}

	tables := doc.Find("table")
	if tables.Length() <= tableIndex {
		return results, nil
	}

	tables.Eq(tableIndex).Find("tr").Each(func(i int, row *goquery.Selection) {
		if i == 0 {
			return
		}

		var cells []string
		row.Find("td").Each(func(j int, cell *goquery.Selection) {
			cells = append(cells, strings.TrimSpace(cell.Text()))
		})

		if len(cells) < 2 {
			return
		}

		rowData := map[string]interface{}{
			"年度":   cells[0],
			"机构数":  parseInt(cells[1]),
			"平均值":  parseFloat(cells[2]),
			"最大值":  parseFloat(cells[3]),
			"最小值":  parseFloat(cells[4]),
			"指标类型": "预测年报净利润",
		}

		results = append(results, rowData)
	})

	return results, nil
}

// parseProfitForecastOrgTable 解析业绩预测详表-机构表格
func parseProfitForecastOrgTable(doc *goquery.Document, tableIndex int) ([]map[string]interface{}, error) {
	var results []map[string]interface{}

	tables := doc.Find("table")
	if tables.Length() <= tableIndex {
		return results, nil
	}

	// 获取表头（处理多级表头）
	var columns []string
	tables.Eq(tableIndex).Find("tr").First().Find("td").Each(func(j int, cell *goquery.Selection) {
		text := strings.TrimSpace(cell.Text())
		if j >= 2 && j <= 4 {
			columns = append(columns, "预测年报每股收益"+text)
		} else if j >= 5 && j <= 7 {
			columns = append(columns, "预测年报净利润"+text)
		} else {
			columns = append(columns, text)
		}
	})

	tables.Eq(tableIndex).Find("tr").Each(func(i int, row *goquery.Selection) {
		if i == 0 || i == 1 {
			// 跳过表头
			return
		}

		var cells []string
		row.Find("td").Each(func(j int, cell *goquery.Selection) {
			cells = append(cells, strings.TrimSpace(cell.Text()))
		})

		if len(cells) < 8 {
			return
		}

		rowData := map[string]interface{}{
			"机构名称":        cells[0],
			"报告日期":        parseDate(cells[1]),
			"预测年报每股收益-平均": parseFloat(cells[2]),
			"预测年报每股收益-最大": parseFloat(cells[3]),
			"预测年报每股收益-最小": parseFloat(cells[4]),
			"预测年报净利润-平均":  parseFloat(cells[5]),
			"预测年报净利润-最大":  parseFloat(cells[6]),
			"预测年报净利润-最小":  parseFloat(cells[7]),
			"指标类型":        "业绩预测详表-机构",
		}

		results = append(results, rowData)
	})

	return results, nil
}

// parseProfitForecastDetailTable 解析业绩预测详表-详细指标预测表格
func parseProfitForecastDetailTable(doc *goquery.Document, tableIndex int) ([]map[string]interface{}, error) {
	var results []map[string]interface{}

	tables := doc.Find("table")
	if tables.Length() <= tableIndex {
		return results, nil
	}

	tables.Eq(tableIndex).Find("tr").Each(func(i int, row *goquery.Selection) {
		if i == 0 {
			return
		}

		var cells []string
		row.Find("td").Each(func(j int, cell *goquery.Selection) {
			text := strings.TrimSpace(cell.Text())
			// 替换括号
			text = regexp.MustCompile(`（`).ReplaceAllString(text, "-")
			text = regexp.MustCompile(`）`).ReplaceAllString(text, "")
			cells = append(cells, text)
		})

		if len(cells) < 2 {
			return
		}

		rowData := make(map[string]interface{})
		rowData["指标类型"] = "业绩预测详表-详细指标预测"
		for j, cell := range cells {
			if j == 0 {
				rowData["指标"] = cell
			} else {
				rowData[fmt.Sprintf("值%d", j)] = cell
			}
		}

		results = append(results, rowData)
	})

	return results, nil
}

// parseInt 解析整数
func parseInt(s string) int {
	s = strings.TrimSpace(s)
	s = utils.RemoveAllSpaces(s)
	if s == "" || s == "-" {
		return 0
	}
	val := utils.MustInt(s)
	return int(val)
}

// parseFloat 解析浮点数
func parseFloat(s string) float64 {
	s = strings.TrimSpace(s)
	s = utils.RemoveAllSpaces(s)
	if s == "" || s == "-" {
		return 0
	}
	return utils.MustParseFloat(s)
}

// parseDate 解析日期
func parseDate(s string) string {
	s = strings.TrimSpace(s)
	s = utils.RemoveAllSpaces(s)
	return s
}

// getMapKeys 获取map的键列表
func getMapKeys(m map[string]bool) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
