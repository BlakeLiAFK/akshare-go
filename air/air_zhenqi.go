package air

import (
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-gota/gota/dataframe"
)

// AirCityTable 获取真气网所有城市列表及其空气质量
//
// 数据源: 真气网
// URL: https://www.zq12369.com/environment.php
//
// 返回:
//   - dataframe.DataFrame: 包含序号、省份、城市、AQI、空气质量、PM2.5浓度、首要污染物
//   - error: 错误信息
//
// 示例:
//
//	df, err := air.AirCityTable()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(df)
func AirCityTable() (dataframe.DataFrame, error) {
	url := "https://www.zq12369.com/environment.php"
	params := map[string]string{
		"date":  "2020-05-01",
		"tab":   "rank",
		"order": "DESC",
		"type":  "DAY",
	}
	
	// 发起HTTP请求
	resp, err := utils.Get(url, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}
	
	// 解析HTML
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("解析HTML失败: %w", err)
	}
	
	// 查找表格
	var df dataframe.DataFrame
	doc.Find("table").Each(func(i int, table *goquery.Selection) {
		if i == 1 { // 第二个表格是数据表
			df = parseAirCityTable(table)
		}
	})
	
	if df.Nrow() == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未找到数据表格")
	}
	
	return df, nil
}

// parseAirCityTable 解析城市空气质量表格
func parseAirCityTable(table *goquery.Selection) dataframe.DataFrame {
	var records [][]string
	
	// 读取表格数据
	table.Find("tr").Each(func(i int, row *goquery.Selection) {
		if i == 0 {
			// 跳过表头
			return
		}
		var rowData []string
		row.Find("td").Each(func(j int, cell *goquery.Selection) {
			// 跳过"降序"列（假设是最后一列或特定列）
			text := strings.TrimSpace(cell.Text())
			if text != "降序" {
				rowData = append(rowData, text)
			}
		})
		if len(rowData) > 0 {
			records = append(records, rowData)
		}
	})
	
	if len(records) == 0 {
		return dataframe.DataFrame{}
	}

	// 构建DataFrame，添加列名作为第一行
	columns := []string{"序号", "省份", "城市", "AQI", "空气质量", "PM2.5浓度", "首要污染物"}
	allRecords := append([][]string{columns}, records...)
	df := dataframe.LoadRecords(allRecords)

	return df
}

// AirQualityWatchPoint 获取指定城市监测点的空气质量数据
//
// 注意: 此接口需要JS加密，暂未完全实现
//
// 参数:
//   - city: 城市名称（如"杭州"），可通过 AirCityTable() 获取
//   - startDate: 开始日期，格式 "20220408"
//   - endDate: 结束日期，格式 "20220409"
//
// 返回:
//   - dataframe.DataFrame: 监测点空气质量数据
//   - error: 错误信息
func AirQualityWatchPoint(city, startDate, endDate string) (dataframe.DataFrame, error) {
	return dataframe.DataFrame{}, fmt.Errorf("此接口需要JS加密功能，暂未实现")
}

// AirQualityHist 获取指定城市的历史空气质量数据
//
// 注意: 此接口需要JS加密，暂未完全实现
//
// 参数:
//   - city: 城市名称（如"杭州"）
//   - period: 数据周期，"hour"(小时)/"day"(天)/"month"(月)
//   - startDate: 开始日期，格式 "20190327"
//   - endDate: 结束日期，格式 "20200427"
//
// 返回:
//   - dataframe.DataFrame: 历史空气质量数据
//   - error: 错误信息
func AirQualityHist(city, period, startDate, endDate string) (dataframe.DataFrame, error) {
	return dataframe.DataFrame{}, fmt.Errorf("此接口需要JS加密功能，暂未实现")
}

// AirQualityRank 获取空气质量排名
//
// 数据源: 真气网
//
// 返回:
//   - dataframe.DataFrame: 城市空气质量排名
//   - error: 错误信息
func AirQualityRank() (dataframe.DataFrame, error) {
	// 实际上AirCityTable就是排名数据
	return AirCityTable()
}
