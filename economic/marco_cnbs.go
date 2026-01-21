package economic

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"github.com/xuri/excelize/v2"
)

// MacroCNBS 国家金融与发展实验室-中国宏观杠杆率数据
//
// 数据源: http://114.115.232.154:8080/
//
// 返回:
//   - dataframe.DataFrame: 中国宏观杠杆率数据
//   - error: 错误信息
func MacroCNBS() (dataframe.DataFrame, error) {
	url := "http://114.115.232.154:8080/handler/download.ashx"

	// 下载 Excel 文件
	resp, err := http.Get(url)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("下载Excel文件失败: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("读取响应失败: %w", err)
	}

	// 使用 excelize 解析 Excel
	f, err := excelize.OpenReader(strings.NewReader(string(body)))
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("解析Excel文件失败: %w", err)
	}
	defer f.Close()

	// 读取 Data sheet
	rows, err := f.GetRows("Data")
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("读取Data表失败: %w", err)
	}

	if len(rows) < 3 {
		return dataframe.DataFrame{}, fmt.Errorf("Excel数据行数不足")
	}

	// 跳过第一行（标题说明），第二行是列名
	header := rows[1]

	// 列名映射
	columnMapping := map[string]string{
		"Period":                           "年份",
		"Household":                        "居民部门",
		"Non-financial corporations":       "非金融企业部门",
		"Central government ":              "中央政府",
		"Central government":               "中央政府",
		"Local government":                 "地方政府",
		"General government":               "政府部门",
		"Non financial sector":             "实体经济部门",
		"Financial sector(asset side)":     "金融部门资产方",
		"Financial sector(liability side)": "金融部门负债方",
	}

	// 构建列名列表
	var columns []string
	for _, col := range header {
		col = strings.TrimSpace(col)
		if mappedName, ok := columnMapping[col]; ok {
			columns = append(columns, mappedName)
		} else {
			columns = append(columns, col)
		}
	}

	// 构建记录
	var records [][]string
	records = append(records, columns)

	for i := 2; i < len(rows); i++ {
		row := rows[i]
		if len(row) == 0 {
			continue
		}

		var record []string
		for j, cell := range row {
			if j >= len(columns) {
				break
			}

			// 处理日期列
			if j == 0 && columns[j] == "年份" {
				// 尝试解析日期
				if t, err := time.Parse("2006-01-02", cell); err == nil {
					cell = t.Format("2006-01")
				} else if t, err := time.Parse("01-02-06", cell); err == nil {
					cell = t.Format("2006-01")
				} else if t, err := time.Parse("1/2/2006", cell); err == nil {
					cell = t.Format("2006-01")
				} else {
					// 尝试解析 Excel 序列号
					if serial, err := strconv.ParseFloat(cell, 64); err == nil && serial > 0 {
						// Excel 日期序列号转换
						t := time.Date(1899, 12, 30, 0, 0, 0, 0, time.UTC).Add(time.Duration(serial*24) * time.Hour)
						cell = t.Format("2006-01")
					}
				}
			}
			record = append(record, cell)
		}

		// 补齐缺失列
		for len(record) < len(columns) {
			record = append(record, "")
		}
		records = append(records, record)
	}

	// 创建 DataFrame
	df := dataframe.LoadRecords(records)

	// 按指定顺序重排列
	columnOrder := []string{
		"年份",
		"居民部门",
		"非金融企业部门",
		"政府部门",
		"中央政府",
		"地方政府",
		"实体经济部门",
		"金融部门资产方",
		"金融部门负债方",
	}

	// 检查并选择存在的列
	var selectCols []string
	for _, col := range columnOrder {
		for _, dfCol := range df.Names() {
			if dfCol == col {
				selectCols = append(selectCols, col)
				break
			}
		}
	}

	if len(selectCols) > 0 {
		df = df.Select(selectCols)
	}

	// 转换数值列
	numericCols := []string{"居民部门", "非金融企业部门", "政府部门", "中央政府", "地方政府", "实体经济部门", "金融部门资产方", "金融部门负债方"}
	for _, col := range numericCols {
		if containsColumn(df.Names(), col) {
			df = convertToNumeric(df, col)
		}
	}

	return df, nil
}

// containsColumn 检查列名是否存在
func containsColumn(columns []string, col string) bool {
	for _, c := range columns {
		if c == col {
			return true
		}
	}
	return false
}

// convertToNumeric 将列转换为数值类型
func convertToNumeric(df dataframe.DataFrame, colName string) dataframe.DataFrame {
	col := df.Col(colName)
	var floatValues []float64
	for i := 0; i < col.Len(); i++ {
		val := col.Elem(i).String()
		if f, err := strconv.ParseFloat(val, 64); err == nil {
			floatValues = append(floatValues, f)
		} else {
			floatValues = append(floatValues, 0)
		}
	}
	newCol := series.Floats(floatValues)
	newCol.Name = colName
	return df.Mutate(newCol)
}
