package economic

import (
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
	"github.com/tidwall/gjson"
)

// CryptoJSSpot 主流加密货币的实时行情数据
//
// 数据源: https://datacenter.jin10.com/reportType/dc_bitcoin_current
//
// 返回:
//   - dataframe.DataFrame: 加密货币实时行情数据
//   - error: 错误信息
func CryptoJSSpot() (dataframe.DataFrame, error) {
	url := "https://datacenter-api.jin10.com/crypto_currency/list"

	headers := map[string]string{
		"user-agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36",
		"x-app-id":     "rU6QIu7JHe2gOUeR",
		"x-csrf-token": "x-csrf-token",
		"x-version":    "1.0.0",
	}

	resp, err := utils.GetWithHeaders(url, nil, headers)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	// 解析 JSON 响应
	dataArray := gjson.Get(resp.String(), "data").Array()
	if len(dataArray) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	// 定义列名
	columns := []string{
		"市场",
		"交易品种",
		"最近报价",
		"涨跌额",
		"涨跌幅",
		"24小时最高",
		"24小时最低",
		"24小时成交量",
		"更新时间",
	}

	// 原始字段顺序
	fields := []string{"market", "symbol", "price", "change", "percent", "high", "low", "volume", "reported_at"}

	// 构建记录
	var records [][]string
	records = append(records, columns)

	for _, item := range dataArray {
		var record []string
		for _, field := range fields {
			value := item.Get(field).String()
			record = append(record, value)
		}
		records = append(records, record)
	}

	// 创建 DataFrame
	df := dataframe.LoadRecords(records)
	return df, nil
}

// MacroFXSentiment 金十数据-外汇-投机情绪报告
//
// 外汇投机情绪报告显示当前市场多空仓位比例，数据由8家交易平台提供
//
// 参数:
//   - startDate: 开始日期，格式"20221011"
//   - endDate: 结束日期，格式"20221017"
//
// 返回:
//   - dataframe.DataFrame: 投机情绪报告数据
//   - error: 错误信息
//
// 数据源: https://datacenter.jin10.com/reportType/dc_ssi_trends
func MacroFXSentiment(startDate, endDate string) (dataframe.DataFrame, error) {
	// 格式化日期
	if len(startDate) != 8 || len(endDate) != 8 {
		return dataframe.DataFrame{}, fmt.Errorf("日期格式错误，应为YYYYMMDD格式")
	}

	formattedStartDate := fmt.Sprintf("%s-%s-%s", startDate[:4], startDate[4:6], startDate[6:])
	formattedEndDate := fmt.Sprintf("%s-%s-%s", endDate[:4], endDate[4:6], endDate[6:])

	url := "https://datacenter-api.jin10.com/sentiment/datas"

	params := map[string]string{
		"start_date":    formattedStartDate,
		"end_date":      formattedEndDate,
		"currency_pair": "",
	}

	headers := map[string]string{
		"accept":          "*/*",
		"accept-encoding": "",
		"accept-language": "zh-CN,zh;q=0.9,en;q=0.8",
		"cache-control":   "no-cache",
		"origin":          "https://datacenter.jin10.com",
		"pragma":          "no-cache",
		"referer":         "https://datacenter.jin10.com/reportType/dc_ssi_trends",
		"sec-fetch-mode":  "cors",
		"sec-fetch-site":  "same-site",
		"user-agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36",
		"x-app-id":        "rU6QIu7JHe2gOUeR",
		"x-csrf-token":    "",
		"x-version":       "1.0.0",
	}

	resp, err := utils.GetWithHeaders(url, params, headers)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	// 解析 JSON 响应
	valuesObj := gjson.Get(resp.String(), "data.values")
	if !valuesObj.Exists() {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	// values 是一个 map，key 是日期，value 是各货币对的数据
	valuesMap := valuesObj.Map()
	if len(valuesMap) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("数据为空")
	}

	// 获取所有列名（货币对）
	var allColumns []string
	for _, v := range valuesMap {
		v.ForEach(func(key, _ gjson.Result) bool {
			col := key.String()
			found := false
			for _, c := range allColumns {
				if c == col {
					found = true
					break
				}
			}
			if !found {
				allColumns = append(allColumns, col)
			}
			return true
		})
		break // 只需要第一个就够了
	}

	// 构建列名
	columns := append([]string{"date"}, allColumns...)

	// 构建记录
	var records [][]string
	records = append(records, columns)

	for date, values := range valuesMap {
		record := []string{date}
		for _, col := range allColumns {
			val := values.Get(col).String()
			if val == "" {
				val = "0"
			}
			record = append(record, val)
		}
		records = append(records, record)
	}

	// 创建 DataFrame
	df := dataframe.LoadRecords(records)
	return df, nil
}
