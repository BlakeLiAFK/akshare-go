package option

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
	"github.com/xuri/excelize/v2"
)

// OptionCurrentDaySzse 深圳证券交易所-期权子网-行情数据-当日合约
//
// 返回:
//   - dataframe.DataFrame: 深圳期权当日合约数据
//   - error: 错误信息
//
// 数据源: https://www.sse.org.cn/option/quotation/contract/daycontract/index.html
func OptionCurrentDaySzse() (dataframe.DataFrame, error) {
	url := "https://www.sse.org.cn/api/report/ShowReport"
	params := map[string]string{
		"SHOWTYPE":  "xlsx",
		"CATALOGID": "option_drhy",
		"TABKEY":    "tab1",
	}

	// 请求Excel文件
	resp, err := utils.Get(url, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	// 使用excelize读取Excel文件
	f, err := excelize.OpenReader(bytes.NewReader(resp.Body()))
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("打开Excel失败: %w", err)
	}
	defer f.Close()

	// 获取第一个工作表
	sheetName := f.GetSheetName(0)
	if sheetName == "" {
		return dataframe.DataFrame{}, fmt.Errorf("未找到工作表")
	}

	// 读取所有行
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("读取工作表失败: %w", err)
	}

	if len(rows) <= 1 {
		return dataframe.DataFrame{}, fmt.Errorf("工作表数据为空")
	}

	// 处理数据
	headers := rows[0]
	var records [][]string

	for i := 1; i < len(rows); i++ {
		row := rows[i]
		if len(row) == 0 {
			continue
		}

		// 确保每行长度与表头一致
		for len(row) < len(headers) {
			row = append(row, "")
		}

		// 转换数值列
		row = convertNumericRow(row, headers)
		records = append(records, row)
	}

	if len(records) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("没有有效的数据记录")
	}

	// 按指定列顺序重新排列
	orderedHeaders := []string{
		"序号", "合约编码", "合约代码", "合约简称", "标的证券简称(代码)", "合约类型", "行权价", "合约单位",
		"最后交易日", "行权日", "到期日", "交收日", "新挂", "涨停价格", "跌停价格", "前结算价",
		"合约调整", "停牌", "合约总持仓", "挂牌原因", "原合约代码", "原合约简称", "原行权价格",
		"原合约单位", "合约到期剩余交易天数", "合约到期剩余自然天数", "下次合约调整剩余交易天数",
		"下次合约调整剩余自然天数", "交易日期",
	}

	// 创建列名到索引的映射
	colIndex := make(map[string]int)
	for i, header := range headers {
		colIndex[header] = i
	}

	// 重新排列数据
	var orderedRecords [][]string
	orderedRecords = append(orderedRecords, orderedHeaders)

	for _, record := range records {
		var orderedRow []string
		for _, header := range orderedHeaders {
			if idx, exists := colIndex[header]; exists && idx < len(record) {
				orderedRow = append(orderedRow, record[idx])
			} else {
				orderedRow = append(orderedRow, "")
			}
		}
		orderedRecords = append(orderedRecords, orderedRow)
	}

	// 构建DataFrame
	df := dataframe.LoadRecords(orderedRecords)
	return df, nil
}

// convertNumericRow 转换数值列
func convertNumericRow(row, headers []string) []string {
	numericColumns := map[string]bool{
		"序号":           true,
		"行权价":          true,
		"合约单位":         true,
		"涨停价格":         true,
		"跌停价格":         true,
		"前结算价":         true,
		"合约总持仓":        true,
		"原行权价格":        true,
		"原合约单位":        true,
		"合约到期剩余交易天数":   true,
		"合约到期剩余自然天数":   true,
		"下次合约调整剩余交易天数": true,
		"下次合约调整剩余自然天数": true,
	}

	for i, header := range headers {
		if i < len(row) && numericColumns[header] {
			row[i] = parseNumericValue(row[i])
		}
	}

	return row
}

// parseNumericValue 解析数值
func parseNumericValue(s string) string {
	s = strings.TrimSpace(s)
	if s == "" || s == "-" {
		return "0"
	}

	// 尝试直接转换为浮点数
	if _, err := strconv.ParseFloat(s, 64); err == nil {
		return s
	}

	// 移除非数字字符（除了小数点）
	result := ""
	for _, r := range s {
		if (r >= '0' && r <= '9') || r == '.' {
			result += string(r)
		}
	}

	if result == "" {
		return "0"
	}

	return result
}
