package utils

import (
	"bytes"
	"fmt"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-resty/resty/v2"
	"github.com/xuri/excelize/v2"
)

// ReadExcelFromBytes 从字节数据中读取 Excel 文件并转换为 DataFrame
//
// 参数:
//   - data: Excel 文件的字节数据
//
// 返回:
//   - dataframe.DataFrame: 解析后的数据框
//   - error: 错误信息
func ReadExcelFromBytes(data []byte) (dataframe.DataFrame, error) {
	// 打开 Excel 文件
	f, err := excelize.OpenReader(bytes.NewReader(data))
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("打开Excel文件失败: %w", err)
	}
	defer f.Close()

	// 获取第一个工作表
	sheets := f.GetSheetList()
	if len(sheets) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("Excel文件无工作表")
	}

	// 读取第一个工作表的所有行
	rows, err := f.GetRows(sheets[0])
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("读取工作表失败: %w", err)
	}

	if len(rows) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("工作表无数据")
	}

	// 将行数据转换为 map 列表
	headers := rows[0]
	var records []map[string]interface{}

	for i := 1; i < len(rows); i++ {
		record := make(map[string]interface{})
		for j, header := range headers {
			if j < len(rows[i]) {
				record[header] = rows[i][j]
			} else {
				record[header] = ""
			}
		}
		records = append(records, record)
	}

	// 转换为 DataFrame
	df := dataframe.LoadMaps(records)
	return df, nil
}

// GetWithHeadersAndParams 发起带自定义请求头和查询参数的 GET 请求
//
// 参数:
//   - url: 请求URL
//   - params: 查询参数
//   - headers: 请求头
//
// 返回:
//   - *resty.Response: 响应对象
//   - error: 错误信息
//
// 注意: 这个函数是 GetWithHeaders 的别名，为了兼容性而保留
func GetWithHeadersAndParams(url string, params, headers map[string]string) (*resty.Response, error) {
	return GetWithHeaders(url, params, headers)
}
