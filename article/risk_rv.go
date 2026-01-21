package article

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

// ArticleOmanRV 获取 Oxford-Man 已实现波动率数据
//
// 参数:
//   - symbol: 指数代码，如 "FTSE"
//   - index: 指标代码，如 "rk_th2"
//
// 返回:
//   - dataframe.DataFrame: 包含日期和波动率数据的 DataFrame
//   - error: 错误信息
//
// 数据源: https://realized.oxford-man.ox.ac.uk
func ArticleOmanRV(symbol, index string) (dataframe.DataFrame, error) {
	url := "https://realized.oxford-man.ox.ac.uk/theme/js/visualization-data.js?20191111113154"

	resp, err := utils.Get(url, nil)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	// 解析HTML提取JSON数据
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("解析HTML失败: %w", err)
	}

	// 提取<p>标签内容
	pText := doc.Find("p").Text()

	// 提取JSON数据（从第一个{到最后一个}）
	startIdx := strings.Index(pText, "{")
	endIdx := strings.LastIndex(pText, "};")
	if startIdx == -1 || endIdx == -1 {
		return dataframe.DataFrame{}, fmt.Errorf("未找到有效的JSON数据")
	}
	jsonStr := pText[startIdx : endIdx+1]

	// 解析JSON
	var dataMap map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &dataMap); err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("解析JSON失败: %w", err)
	}

	// 访问嵌套数据: data[".{symbol}"][index]
	symbolKey := "." + symbol
	symbolData, ok := dataMap[symbolKey].(map[string]interface{})
	if !ok {
		return dataframe.DataFrame{}, fmt.Errorf("未找到指数 %s 的数据", symbol)
	}

	indexData, ok := symbolData[index].(map[string]interface{})
	if !ok {
		return dataframe.DataFrame{}, fmt.Errorf("未找到指标 %s 的数据", index)
	}

	// 提取时间戳和数值数组
	timestamps, ok := indexData["t"].([]interface{})
	if !ok {
		return dataframe.DataFrame{}, fmt.Errorf("时间戳数据格式错误")
	}

	values, ok := indexData["v"].([]interface{})
	if !ok {
		return dataframe.DataFrame{}, fmt.Errorf("数值数据格式错误")
	}

	if len(timestamps) != len(values) {
		return dataframe.DataFrame{}, fmt.Errorf("时间戳和数值长度不匹配")
	}

	// 构建DataFrame
	var dates []string
	var vals []float64

	for i := range timestamps {
		// 时间戳转换为日期（毫秒转秒）
		ts, ok := timestamps[i].(float64)
		if !ok {
			continue
		}
		t := time.Unix(int64(ts)/1000, 0)
		dates = append(dates, t.Format("2006-01-02"))

		// 提取数值
		val, ok := values[i].(float64)
		if !ok {
			vals = append(vals, 0)
		} else {
			vals = append(vals, val)
		}
	}

	df := dataframe.New(
		series.New(dates, series.String, "date"),
		series.New(vals, series.Float, "value"),
	)

	return df, nil
}

// ArticleRlabRV 获取 Risk-Lab 已实现波动率数据
//
// 参数:
//   - symbol: 指数代码，如 "39693"
//
// 返回:
//   - dataframe.DataFrame: 包含日期和波动率数据的 DataFrame
//   - error: 错误信息
//
// 数据源: https://dachxiu.chicagobooth.edu/
func ArticleRlabRV(symbol string) (dataframe.DataFrame, error) {
	url := "https://dachxiu.chicagobooth.edu/data.php"

	params := map[string]string{
		"ticker": symbol,
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	// 解析HTML提取数据
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("解析HTML失败: %w", err)
	}

	// 提取<p>标签内容
	pText := doc.Find("p").Text()

	// 按符号分割数据
	parts := strings.Split(pText, symbol)
	if len(parts) < 2 {
		return dataframe.DataFrame{}, fmt.Errorf("未找到有效数据")
	}

	// 获取数据部分（跳过前两行标题）
	dataText := parts[1]
	lines := strings.Split(dataText, "\n")

	var dates []string
	var vals []float64

	// 从第3行开始处理数据
	for i := 2; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}

		// 第一列是日期（YYYYMMDD格式）
		dateStr := fields[0]
		if len(dateStr) == 8 {
			// 格式化为 YYYY-MM-DD
			date := dateStr[0:4] + "-" + dateStr[4:6] + "-" + dateStr[6:8]
			dates = append(dates, date)

			// 第二列及之后是数值，这里取第二列
			if len(fields) > 1 {
				var val float64
				fmt.Sscanf(fields[1], "%f", &val)
				vals = append(vals, val)
			}
		}
	}

	if len(dates) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未能解析出有效数据")
	}

	df := dataframe.New(
		series.New(dates, series.String, "date"),
		series.New(vals, series.Float, "value"),
	)

	return df, nil
}

// ArticleOmanRVShort 获取 Oxford-Man 简化版已实现波动率数据
//
// 参数:
//   - symbol: 指数代码，如 "FTSE", "GDAXI", "RUT", "SPX", "STOXX50E", "SSEC", "N225"
//
// 返回:
//   - dataframe.DataFrame: 包含日期和波动率数据的 DataFrame
//   - error: 错误信息
//
// 数据源: https://realized.oxford-man.ox.ac.uk
func ArticleOmanRVShort(symbol string) (dataframe.DataFrame, error) {
	url := "https://realized.oxford-man.ox.ac.uk/theme/js/front-page-chart.js"

	headers := map[string]string{
		"Accept":          "*/*",
		"Accept-Encoding": "gzip, deflate, br",
		"Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8",
		"Cache-Control":   "no-cache",
		"Connection":      "keep-alive",
		"Host":            "realized.oxford-man.ox.ac.uk",
		"Pragma":          "no-cache",
		"Referer":         "https://realized.oxford-man.ox.ac.uk/?from=groupmessage&isappinstalled=0",
		"Sec-Fetch-Mode":  "no-cors",
		"Sec-Fetch-Site":  "same-origin",
		"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.97 Safari/537.36",
	}

	resp, err := utils.Get(url, headers)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	// 解析HTML提取JSON数据
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("解析HTML失败: %w", err)
	}

	// 提取<p>标签内容
	pText := doc.Find("p").Text()

	// 提取JSON数据（从第一个{到最后一个}）
	startIdx := strings.Index(pText, "{")
	endIdx := strings.LastIndex(pText, "}")
	if startIdx == -1 || endIdx == -1 {
		return dataframe.DataFrame{}, fmt.Errorf("未找到有效的JSON数据")
	}
	jsonStr := pText[startIdx : endIdx+1]

	// 解析JSON
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("解析JSON失败: %w", err)
	}

	// 获取指定symbol的数据
	symbolKey := "." + symbol
	symbolData, ok := data[symbolKey].(map[string]interface{})
	if !ok {
		return dataframe.DataFrame{}, fmt.Errorf("未找到symbol %s 的数据", symbol)
	}

	// 获取data数组
	dataArray, ok := symbolData["data"].([]interface{})
	if !ok {
		return dataframe.DataFrame{}, fmt.Errorf("数据格式错误")
	}

	var dates []time.Time
	var values []float64

	// 解析数据
	for _, item := range dataArray {
		if arr, ok := item.([]interface{}); ok && len(arr) >= 2 {
			// 第一列是时间戳（毫秒）
			if timestamp, ok := arr[0].(float64); ok {
				date := time.Unix(0, int64(timestamp)*int64(time.Millisecond))
				dates = append(dates, date)

				// 第二列是数值
				if value, ok := arr[1].(float64); ok {
					values = append(values, value)
				} else {
					values = append(values, 0)
				}
			}
		}
	}

	if len(dates) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未能解析出有效数据")
	}

	// 转换日期为字符串
	var dateStrs []string
	for _, date := range dates {
		dateStrs = append(dateStrs, date.Format("2006-01-02"))
	}

	df := dataframe.New(
		series.New(dateStrs, series.String, "date"),
		series.New(values, series.Float, symbol),
	)

	return df, nil
}
