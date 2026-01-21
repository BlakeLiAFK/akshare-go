package option

import (
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

// OptionCommoditySina 新浪商品期权-实时行情
//
// 参数:
//   - symbol: 期权品种代码，如 "CU0" (铜期权)
//
// 返回:
//   - dataframe.DataFrame: 实时行情数据
//   - error: 错误信息
//
// 数据源: https://finance.sina.com.cn/futures/quotes.shtml
func OptionCommoditySina(symbol string) (dataframe.DataFrame, error) {
	if symbol == "" {
		return dataframe.DataFrame{}, fmt.Errorf("品种代码不能为空")
	}

	// 构建URL - 新浪期货期权页面
	url := "https://finance.sina.com.cn/futures/quotes.shtml"

	// 请求数据
	resp, err := utils.Get(url, nil)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	// 解析HTML
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("解析HTML失败: %w", err)
	}

	// 查找期权数据表格
	var records [][]string
	var headers []string

	// 查找包含期权数据的表格
	doc.Find("table").Each(func(i int, table *goquery.Selection) {
		table.Find("tr").Each(func(j int, row *goquery.Selection) {
			var rowData []string
			row.Find("td, th").Each(func(k int, cell *goquery.Selection) {
				text := strings.TrimSpace(cell.Text())
				rowData = append(rowData, text)
			})

			if len(rowData) > 0 {
				if j == 0 {
					// 表头
					headers = rowData
				} else {
					// 数据行
					// 检查是否包含指定的期权品种
					if len(rowData) > 0 && strings.Contains(rowData[0], symbol) {
						// 转换数值列
						rowData = convertSinaNumericRow(rowData)
						records = append(records, rowData)
					}
				}
			}
		})
	})

	if len(headers) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未找到数据表格")
	}

	if len(records) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未找到品种 %s 的数据", symbol)
	}

	// 构建DataFrame
	allRecords := append([][]string{headers}, records...)
	df := dataframe.LoadRecords(allRecords)

	return df, nil
}

// convertSinaNumericRow 转换新浪数据行的数值列
func convertSinaNumericRow(row []string) []string {
	// 假设数值列的位置（需要根据实际表格结构调整）
	numericCols := []int{1, 2, 3, 4, 5, 6, 7, 8} // 价格、成交量等列

	for _, col := range numericCols {
		if col < len(row) {
			row[col] = parseNumericValue(row[col])
		}
	}

	return row
}

// OptionCommoditySinaAll 新浪商品期权-所有品种实时行情
//
// 返回:
//   - dataframe.DataFrame: 所有期权品种的实时行情数据
//   - error: 错误信息
//
// 数据源: https://finance.sina.com.cn/futures/quotes.shtml
func OptionCommoditySinaAll() (dataframe.DataFrame, error) {
	// 构建URL
	url := "https://finance.sina.com.cn/futures/quotes.shtml"

	// 请求数据
	resp, err := utils.Get(url, nil)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	// 解析HTML
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("解析HTML失败: %w", err)
	}

	// 查找期权数据表格
	var records [][]string
	var headers []string

	// 查找包含期权数据的表格
	doc.Find("table").Each(func(i int, table *goquery.Selection) {
		table.Find("tr").Each(func(j int, row *goquery.Selection) {
			var rowData []string
			row.Find("td, th").Each(func(k int, cell *goquery.Selection) {
				text := strings.TrimSpace(cell.Text())
				rowData = append(rowData, text)
			})

			if len(rowData) > 0 {
				if j == 0 {
					// 表头
					headers = rowData
				} else {
					// 数据行 - 筛选期权相关数据
					if len(rowData) > 0 && isOptionRow(rowData) {
						// 转换数值列
						rowData = convertSinaNumericRow(rowData)
						records = append(records, rowData)
					}
				}
			}
		})
	})

	if len(headers) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未找到数据表格")
	}

	if len(records) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未找到期权数据")
	}

	// 构建DataFrame
	allRecords := append([][]string{headers}, records...)
	df := dataframe.LoadRecords(allRecords)

	return df, nil
}

// isOptionRow 判断是否为期权数据行
func isOptionRow(row []string) bool {
	if len(row) == 0 {
		return false
	}

	// 根据期权代码特征判断
	code := strings.ToUpper(row[0])

	// 期权代码通常包含数字和字母的组合
	if len(code) >= 3 {
		// 检查是否包含期权相关的标识
		optionIndicators := []string{"C", "P", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
		hasOptionIndicator := false

		for _, indicator := range optionIndicators {
			if strings.Contains(code, indicator) {
				hasOptionIndicator = true
				break
			}
		}

		return hasOptionIndicator
	}

	return false
}

// OptionFinanceSina 新浪金融期权-实时行情
//
// 参数:
//   - symbol: 期权品种代码，如 "510050" (50ETF期权)
//
// 返回:
//   - dataframe.DataFrame: 实时行情数据
//   - error: 错误信息
//
// 数据源: https://finance.sina.com.cn/
func OptionFinanceSina(symbol string) (dataframe.DataFrame, error) {
	if symbol == "" {
		return dataframe.DataFrame{}, fmt.Errorf("品种代码不能为空")
	}

	// 构建URL - 新浪财经期权页面
	url := fmt.Sprintf("https://hq.sinajs.cn/list=%s", symbol)

	// 请求数据
	resp, err := utils.Get(url, nil)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	// 解析响应数据
	text := strings.TrimSpace(resp.String())
	if text == "" {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	// 新浪返回的数据格式: var hq_str_510050="期权名称,最新价,涨跌,涨跌幅,成交量,持仓量,..."
	lines := strings.Split(text, "\n")
	if len(lines) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("数据格式错误")
	}

	// 提取数据部分
	dataLine := lines[0]
	if !strings.Contains(dataLine, "=") {
		return dataframe.DataFrame{}, fmt.Errorf("数据格式错误")
	}

	dataPart := strings.Split(dataLine, "=")[1]
	dataPart = strings.Trim(dataPart, `";`)

	if dataPart == "" {
		return dataframe.DataFrame{}, fmt.Errorf("数据为空")
	}

	// 分割字段
	fields := strings.Split(dataPart, ",")
	if len(fields) < 10 {
		return dataframe.DataFrame{}, fmt.Errorf("数据字段不足")
	}

	// 定义列名
	headers := []string{
		"期权名称", "最新价", "涨跌", "涨跌幅", "成交量", "持仓量",
		"开盘价", "昨收", "最高", "最低", "时间",
	}

	// 确保字段数量与列名一致
	for len(fields) < len(headers) {
		fields = append(fields, "")
	}

	// 转换数值列
	for i := 1; i < len(fields) && i < len(headers); i++ {
		if headers[i] == "最新价" || headers[i] == "涨跌" || headers[i] == "涨跌幅" ||
			headers[i] == "成交量" || headers[i] == "持仓量" || headers[i] == "开盘价" ||
			headers[i] == "昨收" || headers[i] == "最高" || headers[i] == "最低" {
			fields[i] = parseNumericValue(fields[i])
		}
	}

	// 构建DataFrame
	allRecords := append([][]string{headers}, fields)
	df := dataframe.LoadRecords(allRecords)

	return df, nil
}

// GetSinaOptionSymbols 获取新浪期权品种列表
//
// 返回:
//   - dataframe.DataFrame: 包含期权品种代码和名称的数据
func GetSinaOptionSymbols() (dataframe.DataFrame, error) {
	// 常见的期权品种代码
	symbols := map[string]string{
		"510050": "50ETF期权",
		"510300": "300ETF期权",
		"510500": "500ETF期权",
		"588000": "科创50ETF期权",
		"588080": "科创50ETF期权(期权)",
	}

	var codes []string
	var names []string

	for code, name := range symbols {
		codes = append(codes, code)
		names = append(names, name)
	}

	df := dataframe.New(
		series.New(codes, series.String, "code"),
		series.New(names, series.String, "name"),
	)

	return df, nil
}
