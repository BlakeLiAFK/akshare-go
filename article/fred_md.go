package article

import (
	"encoding/csv"
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
)

// FredMD 获取FRED-MD月度宏观经济数据集
//
// 参数:
//   - date: 日期字符串，格式如 "2020-01" 或 "2020-03"
//
// 返回:
//   - dataframe.DataFrame: FRED-MD月度数据
//   - error: 错误信息
//
// 数据源: https://research.stlouisfed.org/
func FredMD(date string) (dataframe.DataFrame, error) {
	// 构建URL
	// 从AWS S3存储桶下载CSV文件
	url := fmt.Sprintf("https://s3.amazonaws.com/files.research.stlouisfed.org/fred-md/monthly/%s.csv", date)

	resp, err := utils.Get(url, nil)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	// 解析CSV
	reader := csv.NewReader(strings.NewReader(resp.String()))
	reader.LazyQuotes = true // 允许裸引号
	records, err := reader.ReadAll()
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("解析CSV失败: %w", err)
	}

	if len(records) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("数据为空")
	}

	// 构建DataFrame
	df := dataframe.LoadRecords(records)
	return df, nil
}

// FredQD 获取FRED-QD季度宏观经济数据集
//
// 参数:
//   - date: 日期字符串，格式如 "2020-01" 或 "2020-03"
//
// 返回:
//   - dataframe.DataFrame: FRED-QD季度数据
//   - error: 错误信息
//
// 数据源: https://research.stlouisfed.org/
func FredQD(date string) (dataframe.DataFrame, error) {
	// 构建URL
	// 从AWS S3存储桶下载CSV文件
	url := fmt.Sprintf("https://s3.amazonaws.com/files.research.stlouisfed.org/fred-qd/quarterly/%s.csv", date)

	resp, err := utils.Get(url, nil)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	// 解析CSV
	reader := csv.NewReader(strings.NewReader(resp.String()))
	reader.LazyQuotes = true // 允许裸引号
	records, err := reader.ReadAll()
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("解析CSV失败: %w", err)
	}

	if len(records) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("数据为空")
	}

	// 构建DataFrame
	df := dataframe.LoadRecords(records)
	return df, nil
}
