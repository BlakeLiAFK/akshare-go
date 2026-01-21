package economic

import (
	"fmt"
	"sort"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-gota/gota/dataframe"
)

// thsHeaders 同花顺请求头
var thsHeaders = map[string]string{
	"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36",
}

// MacroStockFinance 同花顺-数据中心-宏观数据-股票筹资
//
// 返回:
//   - dataframe.DataFrame: 股票筹资数据
//   - error: 错误信息
//
// 数据源: https://data.10jqka.com.cn/macro/finance/
func MacroStockFinance() (dataframe.DataFrame, error) {
	url := "https://data.10jqka.com.cn/macro/finance/"

	resp, err := utils.GetWithHeaders(url, nil, thsHeaders)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	// 解析 HTML
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("HTML解析失败: %w", err)
	}

	// 定义列名
	columns := []string{"月份", "募集资金", "首发募集资金", "增发募集资金", "配股募集资金"}

	// 构建记录
	var records [][]string
	records = append(records, columns)

	// 查找表格
	doc.Find("table tbody tr").Each(func(i int, s *goquery.Selection) {
		var record []string
		s.Find("td").Each(func(j int, td *goquery.Selection) {
			text := strings.TrimSpace(td.Text())
			record = append(record, text)
		})
		if len(record) >= 5 {
			records = append(records, record[:5])
		}
	})

	if len(records) <= 1 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	// 按月份排序
	sort.Slice(records[1:], func(i, j int) bool {
		return records[i+1][0] < records[j+1][0]
	})

	df := dataframe.LoadRecords(records)
	return df, nil
}

// MacroRMBLoan 同花顺-数据中心-宏观数据-新增人民币贷款
//
// 返回:
//   - dataframe.DataFrame: 新增人民币贷款数据
//   - error: 错误信息
//
// 数据源: https://data.10jqka.com.cn/macro/loan/
func MacroRMBLoan() (dataframe.DataFrame, error) {
	url := "https://data.10jqka.com.cn/macro/loan/"

	resp, err := utils.GetWithHeaders(url, nil, thsHeaders)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	// 解析 HTML
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("HTML解析失败: %w", err)
	}

	// 定义列名
	columns := []string{
		"月份",
		"新增人民币贷款-总额",
		"新增人民币贷款-同比",
		"新增人民币贷款-环比",
		"累计人民币贷款-总额",
		"累计人民币贷款-同比",
	}

	// 构建记录
	var records [][]string
	records = append(records, columns)

	// 查找表格
	doc.Find("table tbody tr").Each(func(i int, s *goquery.Selection) {
		var record []string
		s.Find("td").Each(func(j int, td *goquery.Selection) {
			text := strings.TrimSpace(td.Text())
			record = append(record, text)
		})
		if len(record) >= 6 {
			records = append(records, record[:6])
		}
	})

	if len(records) <= 1 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	// 按月份排序
	sort.Slice(records[1:], func(i, j int) bool {
		return records[i+1][0] < records[j+1][0]
	})

	df := dataframe.LoadRecords(records)
	return df, nil
}

// MacroRMBDeposit 同花顺-数据中心-宏观数据-人民币存款余额
//
// 返回:
//   - dataframe.DataFrame: 人民币存款余额数据
//   - error: 错误信息
//
// 数据源: https://data.10jqka.com.cn/macro/rmb/
func MacroRMBDeposit() (dataframe.DataFrame, error) {
	url := "https://data.10jqka.com.cn/macro/rmb/"

	resp, err := utils.GetWithHeaders(url, nil, thsHeaders)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	// 解析 HTML
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("HTML解析失败: %w", err)
	}

	// 定义列名
	columns := []string{
		"月份",
		"新增存款-数量",
		"新增存款-同比",
		"新增存款-环比",
		"新增企业存款-数量",
		"新增企业存款-同比",
		"新增企业存款-环比",
		"新增储蓄存款-数量",
		"新增储蓄存款-同比",
		"新增储蓄存款-环比",
		"新增其他存款-数量",
		"新增其他存款-同比",
		"新增其他存款-环比",
	}

	// 构建记录
	var records [][]string
	records = append(records, columns)

	// 查找表格
	doc.Find("table tbody tr").Each(func(i int, s *goquery.Selection) {
		var record []string
		s.Find("td").Each(func(j int, td *goquery.Selection) {
			text := strings.TrimSpace(td.Text())
			record = append(record, text)
		})
		if len(record) >= 13 {
			records = append(records, record[:13])
		}
	})

	if len(records) <= 1 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	// 按月份排序
	sort.Slice(records[1:], func(i, j int) bool {
		return records[i+1][0] < records[j+1][0]
	})

	df := dataframe.LoadRecords(records)
	return df, nil
}
