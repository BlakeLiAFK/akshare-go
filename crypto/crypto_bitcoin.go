package crypto

import (
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
	"github.com/tidwall/gjson"
)

// CryptoBitcoinCME 芝加哥商业交易所-比特币成交量报告
//
// 参数:
//
//	date: 指定日期，格式"20230830"
//
// 返回:
//   - dataframe.DataFrame: 比特币成交量报告数据
//   - error: 错误信息
//
// 数据源: https://datacenter.jin10.com/reportType/dc_cme_btc_report
func CryptoBitcoinCME(date string) (dataframe.DataFrame, error) {
	// 将日期从 "20230830" 格式转换为 "2023-08-30"
	if len(date) != 8 {
		return dataframe.DataFrame{}, fmt.Errorf("日期格式错误，应为YYYYMMDD格式，如：20230830")
	}
	formattedDate := fmt.Sprintf("%s-%s-%s", date[:4], date[4:6], date[6:])

	url := "https://datacenter-api.jin10.com/reports/list"
	params := map[string]string{
		"category": "cme",
		"date":     formattedDate,
		"attr_id":  "4",
	}

	headers := map[string]string{
		"accept":          "*/*",
		"accept-encoding": "gzip, deflate, br",
		"accept-language": "zh-CN,zh;q=0.9,en;q=0.8",
		"cache-control":   "no-cache",
		"origin":          "https://datacenter.jin10.com",
		"pragma":          "no-cache",
		"referer":         "https://datacenter.jin10.com/",
		"user-agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.106 Safari/537.36",
		"x-app-id":        "rU6QIu7JHe2gOUeR",
		"x-csrf-token":    "",
		"x-version":       "1.0.0",
	}

	resp, err := utils.GetWithHeaders(url, params, headers)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	// 解析 JSON 响应
	dataObj := gjson.Get(resp.String(), "data")
	if !dataObj.Exists() {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	// 获取列名
	keysArray := gjson.Get(resp.String(), "data.keys").Array()
	if len(keysArray) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到列名")
	}

	var columns []string
	for _, key := range keysArray {
		name := key.Get("name").String()
		columns = append(columns, name)
	}

	// 获取数据
	valuesArray := gjson.Get(resp.String(), "data.values").Array()
	if len(valuesArray) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据值")
	}

	// 构建记录
	var records [][]string
	records = append(records, columns) // 添加列头

	for _, row := range valuesArray {
		var record []string
		for i := 0; i < len(columns); i++ {
			value := row.Array()[i].String()
			record = append(record, value)
		}
		records = append(records, record)
	}

	// 创建 DataFrame
	df := dataframe.LoadRecords(records)
	return df, nil
}

// CryptoBitcoinHoldReport 金十数据-比特币持仓报告
//
// 参数:
//
//	无
//
// 返回:
//   - dataframe.DataFrame: 比特币持仓报告数据
//   - error: 错误信息
//
// 数据源: https://datacenter.jin10.com/dc_report?name=bitcoint
func CryptoBitcoinHoldReport() (dataframe.DataFrame, error) {
	url := "https://datacenter-api.jin10.com/bitcoin_treasuries/list"

	headers := map[string]string{
		"X-App-Id":  "lnFP5lxse24wPgtY",
		"X-Version": "1.0.0",
	}

	resp, err := utils.GetWithHeaders(url, nil, headers)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	// 解析 JSON 响应
	dataObj := gjson.Get(resp.String(), "data.values")
	if !dataObj.Exists() {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	valuesArray := dataObj.Array()
	if len(valuesArray) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据值")
	}

	// 定义列名（与 Python 版本一致）
	columns := []string{
		"代码",
		"公司名称-英文",
		"公司名称-中文",
		"国家/地区",
		"市值",
		"比特币占市值比重",
		"持仓成本",
		"持仓占比",
		"持仓量",
		"当日持仓市值",
		"查询日期",
		"公告链接",
		"分类",
		"倍数",
	}

	// 原始数据的列索引映射（Python中的顺序）
	// 0: 代码, 1: 公司名称-英文, 2: 国家/地区, 3: 市值, 4: 比特币占市值比重,
	// 5: 持仓成本, 6: 持仓占比, 7: 持仓量, 8: 当日持仓市值, 9: 查询日期,
	// 10: 公告链接, 11: _, 12: 分类, 13: 倍数, 14: _, 15: 公司名称-中文
	columnIndexMap := []int{0, 1, 15, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 13}

	// 构建记录
	var records [][]string
	records = append(records, columns) // 添加列头

	for _, row := range valuesArray {
		rowArray := row.Array()
		if len(rowArray) < 16 {
			continue // 跳过不完整的行
		}

		var record []string
		for _, idx := range columnIndexMap {
			value := rowArray[idx].String()
			record = append(record, value)
		}
		records = append(records, record)
	}

	// 创建 DataFrame
	df := dataframe.LoadRecords(records)
	return df, nil
}
