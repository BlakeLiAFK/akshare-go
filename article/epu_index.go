package article

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
	"github.com/xuri/excelize/v2"
)

// ArticleEPUIndex 获取主要国家和地区的经济政策不确定性(EPU)指数
//
// 参数:
//   - symbol: 国家或地区名称，如 "China", "USA", "Europe" 等
//
// 返回:
//   - dataframe.DataFrame: EPU指数数据
//   - error: 错误信息
//
// 数据源: https://www.policyuncertainty.com/
func ArticleEPUIndex(symbol string) (dataframe.DataFrame, error) {
	// 符号映射
	symbolMap := map[string]string{
		"China":       "SCMP_China",
		"China New":   "SCMP_China",
		"USA":         "US",
		"Hong Kong":   "HongKong",
		"France":      "France",
		"Germany":     "Germany",
		"Greece":      "FKT_Greece",
		"Italy":       "Italy",
		"Netherlands": "Netherlands",
		"Spain":       "Spain",
		"UK":          "UK",
		"Russia":      "Russia",
		"Korea":       "Korea",
		"Japan":       "Japan",
		"India":       "India",
		"Singapore":   "Singapore",
		"Brazil":      "Brazil",
		"Chile":       "Chile",
		"Colombia":    "Colombia",
		"Mexico":      "Mexico",
		"Europe":      "Europe",
		"Australia":   "Australia",
		"Canada":      "Canada",
		"Ireland":     "Ireland",
		"Sweden":      "Sweden",
	}

	mappedSymbol, ok := symbolMap[symbol]
	if !ok {
		mappedSymbol = symbol // 使用原始输入
	}

	// 构建URL
	var url string
	isExcel := false

	// Excel文件的国家/地区
	excelSymbols := map[string]bool{
		"HongKong":    true,
		"France":      true,
		"Germany":     true,
		"Italy":       true,
		"Netherlands": true,
		"Spain":       true,
		"UK":          true,
		"FKT_Greece":  true,
		"Europe":      true,
		"Australia":   true,
		"Canada":      true,
		"Ireland":     true,
		"Korea":       true,
		"Russia":      true,
		"Singapore":   true,
		"Sweden":      true,
	}

	if excelSymbols[mappedSymbol] {
		isExcel = true
		if mappedSymbol == "FKT_Greece" {
			url = fmt.Sprintf("http://www.policyuncertainty.com/media/%s_EPU_Data_Annotated.xlsx", mappedSymbol)
		} else if mappedSymbol == "Europe" {
			url = "http://www.policyuncertainty.com/media/Europe_Policy_Uncertainty_Data.xlsx"
		} else {
			url = fmt.Sprintf("http://www.policyuncertainty.com/media/%s_Policy_Uncertainty_Data.xlsx", mappedSymbol)
		}
	} else {
		url = fmt.Sprintf("http://www.policyuncertainty.com/%s_Policy_Uncertainty_Data.csv", mappedSymbol)
	}

	// 下载数据
	resp, err := utils.Get(url, nil)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	// 根据文件类型解析
	if isExcel {
		return parseExcelData(resp.Body())
	}
	return parseCSVData(resp.String())
}

// parseExcelData 解析Excel数据
func parseExcelData(data []byte) (dataframe.DataFrame, error) {
	// 创建临时文件读取Excel
	f, err := excelize.OpenReader(bytes.NewReader(data))
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("打开Excel失败: %w", err)
	}
	defer f.Close()

	// 获取第一个工作表
	sheetName := f.GetSheetName(0)
	if sheetName == "" {
		return dataframe.DataFrame{}, fmt.Errorf("未找到工作表")
	}

	// 读取数据
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("读取工作表失败: %w", err)
	}

	if len(rows) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("工作表为空")
	}

	// 构建DataFrame
	df := dataframe.LoadRecords(rows)
	return df, nil
}

// parseCSVData 解析CSV数据
func parseCSVData(data string) (dataframe.DataFrame, error) {
	reader := csv.NewReader(strings.NewReader(data))
	records, err := reader.ReadAll()
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("解析CSV失败: %w", err)
	}

	if len(records) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("CSV数据为空")
	}

	df := dataframe.LoadRecords(records)
	return df, nil
}
