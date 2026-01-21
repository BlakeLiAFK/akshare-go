// macro_china_part5.go - 中国宏观经济数据函数 (第五部分)
// 包含新浪财经数据API和部分东方财富数据

package economic

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"github.com/tidwall/gjson"
	"golang.org/x/text/encoding/simplifiedchinese"
)

// fetchSinaIndustryData 获取新浪财经宏观经济数据
// cate: 数据分类 (industry, fininfo, price等)
// event: 事件ID
// dataPath: 数据在JSON中的路径 (空字符串表示直接在data下)
func fetchSinaIndustryData(cate, event, dataPath string) (dataframe.DataFrame, []string, error) {
	baseURL := "https://quotes.sina.cn/mac/api/jsonp_v3.php/SINAREMOTECALLCALLBACK/MacPage_Service.get_pagedata"

	// 第一次请求获取总数和列配置
	params := map[string]string{
		"cate":      cate,
		"event":     event,
		"from":      "0",
		"num":       "31",
		"condition": "",
	}

	resp, err := utils.Get(baseURL, params)
	if err != nil {
		return dataframe.DataFrame{}, nil, fmt.Errorf("获取新浪财经数据失败: %w", err)
	}

	// 解析JSONP响应: SINAREMOTECALLCALLBACK({...})
	jsonStr := extractJSONFromJSONP(resp.String())
	if jsonStr == "" {
		return dataframe.DataFrame{}, nil, fmt.Errorf("解析JSONP响应失败")
	}

	// 获取总数和页数
	countStr := gjson.Get(jsonStr, "count").String()
	count, _ := strconv.Atoi(countStr)
	if count == 0 {
		return dataframe.DataFrame{}, nil, fmt.Errorf("没有数据")
	}
	pageNum := int(math.Ceil(float64(count) / 31.0))

	// 获取列配置
	configAll := gjson.Get(jsonStr, "config.all").Array()
	columns := make([]string, len(configAll))
	for i, item := range configAll {
		columns[i] = item.Get("1").String()
	}

	// 获取数据
	var allData []map[string]interface{}
	dataJSON := gjson.Get(jsonStr, "data")

	// 处理不同的数据路径
	if dataPath != "" {
		dataJSON = gjson.Get(jsonStr, "data."+dataPath)
	}

	if dataJSON.IsArray() {
		for _, item := range dataJSON.Array() {
			row := make(map[string]interface{})
			for i, col := range columns {
				row[col] = item.Get(strconv.Itoa(i)).String()
			}
			allData = append(allData, row)
		}
	}

	// 分页获取剩余数据
	for page := 1; page < pageNum; page++ {
		params["from"] = strconv.Itoa(page * 31)

		resp, err := utils.Get(baseURL, params)
		if err != nil {
			continue
		}

		jsonStr = extractJSONFromJSONP(resp.String())
		if jsonStr == "" {
			continue
		}

		dataJSON = gjson.Get(jsonStr, "data")
		if dataPath != "" {
			dataJSON = gjson.Get(jsonStr, "data."+dataPath)
		}

		if dataJSON.IsArray() {
			for _, item := range dataJSON.Array() {
				row := make(map[string]interface{})
				for i, col := range columns {
					row[col] = item.Get(strconv.Itoa(i)).String()
				}
				allData = append(allData, row)
			}
		}
	}

	if len(allData) == 0 {
		return dataframe.DataFrame{}, columns, fmt.Errorf("数据为空")
	}

	// 转换为DataFrame
	df := dataframe.LoadMaps(allData)
	return df, columns, nil
}

// extractJSONFromJSONP 从JSONP响应中提取JSON
func extractJSONFromJSONP(resp string) string {
	// 查找第一个 { 和最后一个 }
	start := strings.Index(resp, "{")
	end := strings.LastIndex(resp, "}")
	if start == -1 || end == -1 || start >= end {
		return ""
	}
	return resp[start : end+1]
}

// MacroChinaSocietyElectricity 获取全社会用电分类情况表
// 数据源: 新浪财经
// https://finance.sina.com.cn/mac/#industry-6-0-31-1
func MacroChinaSocietyElectricity() (dataframe.DataFrame, error) {
	baseURL := "https://quotes.sina.cn/mac/api/jsonp_v3.php/SINAREMOTECALLCALLBACK/MacPage_Service.get_pagedata"

	params := map[string]string{
		"cate":      "industry",
		"event":     "6",
		"from":      "0",
		"num":       "31",
		"condition": "",
	}

	// 第一次请求
	resp, err := utils.Get(baseURL, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("获取全社会用电数据失败: %w", err)
	}

	jsonStr := extractJSONFromJSONP(resp.String())
	if jsonStr == "" {
		return dataframe.DataFrame{}, fmt.Errorf("解析JSONP响应失败")
	}

	// 获取总页数
	countStr := gjson.Get(jsonStr, "count").String()
	count, _ := strconv.Atoi(countStr)
	pageNum := int(math.Ceil(float64(count) / 31.0))

	// 定义列名
	columns := []string{
		"统计时间", "全社会用电量", "全社会用电量同比",
		"各行业用电量合计", "各行业用电量合计同比",
		"第一产业用电量", "第一产业用电量同比",
		"第二产业用电量", "第二产业用电量同比",
		"第三产业用电量", "第三产业用电量同比",
		"城乡居民生活用电量合计", "城乡居民生活用电量合计同比",
		"城镇居民用电量", "城镇居民用电量同比",
		"乡村居民用电量", "乡村居民用电量同比",
	}

	var allRows [][]string

	// 解析第一页数据
	dataJSON := gjson.Get(jsonStr, "data")
	for _, item := range dataJSON.Array() {
		row := make([]string, len(columns))
		for i := range columns {
			row[i] = item.Get(strconv.Itoa(i)).String()
		}
		allRows = append(allRows, row)
	}

	// 分页获取剩余数据
	for page := 1; page < pageNum; page++ {
		params["from"] = strconv.Itoa(page * 31)
		resp, err := utils.Get(baseURL, params)
		if err != nil {
			continue
		}

		jsonStr = extractJSONFromJSONP(resp.String())
		if jsonStr == "" {
			continue
		}

		dataJSON = gjson.Get(jsonStr, "data")
		for _, item := range dataJSON.Array() {
			row := make([]string, len(columns))
			for i := range columns {
				row[i] = item.Get(strconv.Itoa(i)).String()
			}
			allRows = append(allRows, row)
		}
	}

	if len(allRows) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("数据为空")
	}

	// 构建DataFrame
	seriesList := make([]series.Series, len(columns))
	for i, col := range columns {
		vals := make([]string, len(allRows))
		for j, row := range allRows {
			vals[j] = row[i]
		}
		if i == 0 {
			seriesList[i] = series.New(vals, series.String, col)
		} else {
			seriesList[i] = series.New(vals, series.Float, col)
		}
	}

	df := dataframe.New(seriesList...)
	df = df.Arrange(dataframe.Sort("统计时间"))
	return df, nil
}

// MacroChinaSocietyTrafficVolume 获取全社会客货运输量
// 数据源: 新浪财经
// https://finance.sina.com.cn/mac/#industry-10-0-31-1
func MacroChinaSocietyTrafficVolume() (dataframe.DataFrame, error) {
	baseURL := "https://quotes.sina.cn/mac/api/jsonp_v3.php/SINAREMOTECALLCALLBACK/MacPage_Service.get_pagedata"

	params := map[string]string{
		"cate":      "industry",
		"event":     "10",
		"from":      "0",
		"num":       "31",
		"condition": "",
	}

	// 第一次请求
	resp, err := utils.Get(baseURL, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("获取全社会客货运输量数据失败: %w", err)
	}

	jsonStr := extractJSONFromJSONP(resp.String())
	if jsonStr == "" {
		return dataframe.DataFrame{}, fmt.Errorf("解析JSONP响应失败")
	}

	// 获取总页数
	countStr := gjson.Get(jsonStr, "count").String()
	count, _ := strconv.Atoi(countStr)
	pageNum := int(math.Ceil(float64(count) / 31.0))

	// 从config中获取列名
	configAll := gjson.Get(jsonStr, "config.all").Array()
	columns := make([]string, len(configAll))
	for i, item := range configAll {
		columns[i] = item.Get("1").String()
	}

	var allRows [][]string

	// 解析第一页数据 - 使用"非累计"
	dataJSON := gjson.Get(jsonStr, "data.非累计")
	for _, item := range dataJSON.Array() {
		row := make([]string, len(columns))
		for i := range columns {
			row[i] = item.Get(strconv.Itoa(i)).String()
		}
		allRows = append(allRows, row)
	}

	// 分页获取剩余数据
	for page := 1; page < pageNum; page++ {
		params["from"] = strconv.Itoa(page * 31)
		resp, err := utils.Get(baseURL, params)
		if err != nil {
			continue
		}

		jsonStr = extractJSONFromJSONP(resp.String())
		if jsonStr == "" {
			continue
		}

		dataJSON = gjson.Get(jsonStr, "data.非累计")
		for _, item := range dataJSON.Array() {
			row := make([]string, len(columns))
			for i := range columns {
				row[i] = item.Get(strconv.Itoa(i)).String()
			}
			allRows = append(allRows, row)
		}
	}

	if len(allRows) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("数据为空")
	}

	// 构建DataFrame
	seriesList := make([]series.Series, len(columns))
	for i, col := range columns {
		vals := make([]string, len(allRows))
		for j, row := range allRows {
			vals[j] = row[i]
		}
		if i == 0 {
			seriesList[i] = series.New(vals, series.String, col)
		} else {
			seriesList[i] = series.New(vals, series.Float, col)
		}
	}

	return dataframe.New(seriesList...), nil
}

// MacroChinaPostalTelecommunicational 获取邮电业务基本情况
// 数据源: 新浪财经
// https://finance.sina.com.cn/mac/#industry-11-0-31-1
func MacroChinaPostalTelecommunicational() (dataframe.DataFrame, error) {
	baseURL := "https://quotes.sina.cn/mac/api/jsonp_v3.php/SINAREMOTECALLCALLBACK/MacPage_Service.get_pagedata"

	params := map[string]string{
		"cate":      "industry",
		"event":     "11",
		"from":      "0",
		"num":       "31",
		"condition": "",
	}

	resp, err := utils.Get(baseURL, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("获取邮电业务数据失败: %w", err)
	}

	jsonStr := extractJSONFromJSONP(resp.String())
	if jsonStr == "" {
		return dataframe.DataFrame{}, fmt.Errorf("解析JSONP响应失败")
	}

	countStr := gjson.Get(jsonStr, "count").String()
	count, _ := strconv.Atoi(countStr)
	pageNum := int(math.Ceil(float64(count) / 31.0))

	// 获取列名
	configAll := gjson.Get(jsonStr, "config.all").Array()
	columns := make([]string, len(configAll))
	for i, item := range configAll {
		columns[i] = item.Get("1").String()
	}

	var allRows [][]string

	dataJSON := gjson.Get(jsonStr, "data.非累计")
	for _, item := range dataJSON.Array() {
		row := make([]string, len(columns))
		for i := range columns {
			row[i] = item.Get(strconv.Itoa(i)).String()
		}
		allRows = append(allRows, row)
	}

	for page := 1; page < pageNum; page++ {
		params["from"] = strconv.Itoa(page * 31)
		resp, err := utils.Get(baseURL, params)
		if err != nil {
			continue
		}

		jsonStr = extractJSONFromJSONP(resp.String())
		if jsonStr == "" {
			continue
		}

		dataJSON = gjson.Get(jsonStr, "data.非累计")
		for _, item := range dataJSON.Array() {
			row := make([]string, len(columns))
			for i := range columns {
				row[i] = item.Get(strconv.Itoa(i)).String()
			}
			allRows = append(allRows, row)
		}
	}

	if len(allRows) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("数据为空")
	}

	seriesList := make([]series.Series, len(columns))
	for i, col := range columns {
		vals := make([]string, len(allRows))
		for j, row := range allRows {
			vals[j] = row[i]
		}
		if i == 0 {
			seriesList[i] = series.New(vals, series.String, col)
		} else {
			seriesList[i] = series.New(vals, series.Float, col)
		}
	}

	return dataframe.New(seriesList...), nil
}

// MacroChinaInternationalTourismFx 获取国际旅游外汇收入构成
// 数据源: 新浪财经
// https://finance.sina.com.cn/mac/#industry-15-0-31-3
func MacroChinaInternationalTourismFx() (dataframe.DataFrame, error) {
	baseURL := "https://quotes.sina.cn/mac/api/jsonp_v3.php/SINAREMOTECALLCALLBACK/MacPage_Service.get_pagedata"

	params := map[string]string{
		"cate":      "industry",
		"event":     "15",
		"from":      "0",
		"num":       "31",
		"condition": "",
	}

	resp, err := utils.Get(baseURL, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("获取国际旅游外汇收入数据失败: %w", err)
	}

	jsonStr := extractJSONFromJSONP(resp.String())
	if jsonStr == "" {
		return dataframe.DataFrame{}, fmt.Errorf("解析JSONP响应失败")
	}

	countStr := gjson.Get(jsonStr, "count").String()
	count, _ := strconv.Atoi(countStr)
	pageNum := int(math.Ceil(float64(count) / 31.0))

	configAll := gjson.Get(jsonStr, "config.all").Array()
	columns := make([]string, len(configAll))
	for i, item := range configAll {
		columns[i] = item.Get("1").String()
	}

	var allRows [][]string

	dataJSON := gjson.Get(jsonStr, "data")
	for _, item := range dataJSON.Array() {
		row := make([]string, len(columns))
		for i := range columns {
			row[i] = item.Get(strconv.Itoa(i)).String()
		}
		allRows = append(allRows, row)
	}

	for page := 1; page < pageNum; page++ {
		params["from"] = strconv.Itoa(page * 31)
		resp, err := utils.Get(baseURL, params)
		if err != nil {
			continue
		}

		jsonStr = extractJSONFromJSONP(resp.String())
		if jsonStr == "" {
			continue
		}

		dataJSON = gjson.Get(jsonStr, "data")
		for _, item := range dataJSON.Array() {
			row := make([]string, len(columns))
			for i := range columns {
				row[i] = item.Get(strconv.Itoa(i)).String()
			}
			allRows = append(allRows, row)
		}
	}

	if len(allRows) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("数据为空")
	}

	seriesList := make([]series.Series, len(columns))
	for i, col := range columns {
		vals := make([]string, len(allRows))
		for j, row := range allRows {
			vals[j] = row[i]
		}
		// 数量和比重为数值类型
		if col == "数量" || col == "比重" {
			seriesList[i] = series.New(vals, series.Float, col)
		} else {
			seriesList[i] = series.New(vals, series.String, col)
		}
	}

	return dataframe.New(seriesList...), nil
}

// MacroChinaPassengerLoadFactor 获取民航客座率及载运率
// 数据源: 新浪财经
// https://finance.sina.com.cn/mac/#industry-20-0-31-1
func MacroChinaPassengerLoadFactor() (dataframe.DataFrame, error) {
	baseURL := "https://quotes.sina.cn/mac/api/jsonp_v3.php/SINAREMOTECALLCALLBACK/MacPage_Service.get_pagedata"

	params := map[string]string{
		"cate":      "industry",
		"event":     "20",
		"from":      "0",
		"num":       "31",
		"condition": "",
	}

	resp, err := utils.Get(baseURL, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("获取民航客座率数据失败: %w", err)
	}

	jsonStr := extractJSONFromJSONP(resp.String())
	if jsonStr == "" {
		return dataframe.DataFrame{}, fmt.Errorf("解析JSONP响应失败")
	}

	countStr := gjson.Get(jsonStr, "count").String()
	count, _ := strconv.Atoi(countStr)
	pageNum := int(math.Ceil(float64(count) / 31.0))

	configAll := gjson.Get(jsonStr, "config.all").Array()
	columns := make([]string, len(configAll))
	for i, item := range configAll {
		columns[i] = item.Get("1").String()
	}

	var allRows [][]string

	dataJSON := gjson.Get(jsonStr, "data")
	for _, item := range dataJSON.Array() {
		row := make([]string, len(columns))
		for i := range columns {
			row[i] = item.Get(strconv.Itoa(i)).String()
		}
		allRows = append(allRows, row)
	}

	for page := 1; page < pageNum; page++ {
		params["from"] = strconv.Itoa(page * 31)
		resp, err := utils.Get(baseURL, params)
		if err != nil {
			continue
		}

		jsonStr = extractJSONFromJSONP(resp.String())
		if jsonStr == "" {
			continue
		}

		dataJSON = gjson.Get(jsonStr, "data")
		for _, item := range dataJSON.Array() {
			row := make([]string, len(columns))
			for i := range columns {
				row[i] = item.Get(strconv.Itoa(i)).String()
			}
			allRows = append(allRows, row)
		}
	}

	if len(allRows) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("数据为空")
	}

	seriesList := make([]series.Series, len(columns))
	for i, col := range columns {
		vals := make([]string, len(allRows))
		for j, row := range allRows {
			vals[j] = row[i]
		}
		if col == "客座率" || col == "载运率" {
			seriesList[i] = series.New(vals, series.Float, col)
		} else {
			seriesList[i] = series.New(vals, series.String, col)
		}
	}

	return dataframe.New(seriesList...), nil
}

// MacroChinaFreightIndex 获取航贸运价指数
// 数据源: 新浪财经
// https://finance.sina.com.cn/mac/#industry-22-0-31-2
func MacroChinaFreightIndex() (dataframe.DataFrame, error) {
	// 使用Excel导出接口获取完整数据
	url := "http://quotes.sina.cn/mac/view/vMacExcle.php"

	params := map[string]string{
		"cate":      "industry",
		"event":     "22",
		"from":      "0",
		"num":       "5000",
		"condition": "",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("获取航贸运价指数数据失败: %w", err)
	}

	// 解码GBK响应
	decoder := simplifiedchinese.GBK.NewDecoder()
	content, err := decoder.String(resp.String())
	if err != nil {
		// 尝试直接使用
		content = resp.String()
	}

	// 按行分割
	lines := strings.Split(content, "\n")
	if len(lines) < 4 {
		return dataframe.DataFrame{}, fmt.Errorf("数据格式错误")
	}

	// 获取列名 (第3行)
	columnsLine := strings.TrimSpace(lines[2])
	columns := strings.Split(columnsLine, ", ")
	for i := range columns {
		columns[i] = strings.TrimSpace(columns[i])
	}

	// 解析数据行 (从第4行开始)
	var allRows [][]string
	for i := 3; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}
		fields := strings.Split(line, ", ")
		if len(fields) >= len(columns)-1 {
			row := make([]string, len(columns))
			for j := 0; j < len(columns) && j < len(fields); j++ {
				row[j] = strings.TrimSpace(fields[j])
			}
			allRows = append(allRows, row)
		}
	}

	if len(allRows) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("数据为空")
	}

	// 构建DataFrame
	seriesList := make([]series.Series, len(columns))
	for i, col := range columns {
		vals := make([]string, len(allRows))
		for j, row := range allRows {
			if i < len(row) {
				vals[j] = row[i]
			}
		}
		if i == 0 {
			seriesList[i] = series.New(vals, series.String, col)
		} else {
			seriesList[i] = series.New(vals, series.Float, col)
		}
	}

	df := dataframe.New(seriesList...)
	// 移除全空列
	df = df.Drop(df.Ncol() - 1)
	return df, nil
}

// MacroChinaCentralBankBalance 获取央行货币当局资产负债
// 数据源: 新浪财经
// https://finance.sina.com.cn/mac/#fininfo-8-0-31-2
func MacroChinaCentralBankBalance() (dataframe.DataFrame, error) {
	baseURL := "https://quotes.sina.cn/mac/api/jsonp_v3.php/SINAREMOTECALLCALLBACK/MacPage_Service.get_pagedata"

	params := map[string]string{
		"cate":      "fininfo",
		"event":     "8",
		"from":      "0",
		"num":       "31",
		"condition": "",
	}

	resp, err := utils.Get(baseURL, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("获取央行资产负债数据失败: %w", err)
	}

	jsonStr := extractJSONFromJSONP(resp.String())
	if jsonStr == "" {
		return dataframe.DataFrame{}, fmt.Errorf("解析JSONP响应失败")
	}

	countStr := gjson.Get(jsonStr, "count").String()
	count, _ := strconv.Atoi(countStr)
	pageNum := int(math.Ceil(float64(count) / 31.0))

	configAll := gjson.Get(jsonStr, "config.all").Array()
	columns := make([]string, len(configAll))
	for i, item := range configAll {
		columns[i] = item.Get("1").String()
	}

	var allRows [][]string

	dataJSON := gjson.Get(jsonStr, "data")
	for _, item := range dataJSON.Array() {
		row := make([]string, len(columns))
		for i := range columns {
			row[i] = item.Get(strconv.Itoa(i)).String()
		}
		allRows = append(allRows, row)
	}

	for page := 1; page < pageNum; page++ {
		params["from"] = strconv.Itoa(page * 31)
		resp, err := utils.Get(baseURL, params)
		if err != nil {
			continue
		}

		jsonStr = extractJSONFromJSONP(resp.String())
		if jsonStr == "" {
			continue
		}

		dataJSON = gjson.Get(jsonStr, "data")
		for _, item := range dataJSON.Array() {
			row := make([]string, len(columns))
			for i := range columns {
				row[i] = item.Get(strconv.Itoa(i)).String()
			}
			allRows = append(allRows, row)
		}
	}

	if len(allRows) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("数据为空")
	}

	seriesList := make([]series.Series, len(columns))
	for i, col := range columns {
		vals := make([]string, len(allRows))
		for j, row := range allRows {
			vals[j] = row[i]
		}
		if i == 0 {
			seriesList[i] = series.New(vals, series.String, col)
		} else {
			seriesList[i] = series.New(vals, series.Float, col)
		}
	}

	return dataframe.New(seriesList...), nil
}

// MacroChinaInsurance 获取保险业经营情况
// 数据源: 新浪财经
// https://finance.sina.com.cn/mac/#fininfo-19-0-31-3
func MacroChinaInsurance() (dataframe.DataFrame, error) {
	baseURL := "https://quotes.sina.cn/mac/api/jsonp_v3.php/SINAREMOTECALLCALLBACK/MacPage_Service.get_pagedata"

	params := map[string]string{
		"cate":      "fininfo",
		"event":     "19",
		"from":      "0",
		"num":       "31",
		"condition": "",
	}

	resp, err := utils.Get(baseURL, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("获取保险业经营数据失败: %w", err)
	}

	jsonStr := extractJSONFromJSONP(resp.String())
	if jsonStr == "" {
		return dataframe.DataFrame{}, fmt.Errorf("解析JSONP响应失败")
	}

	countStr := gjson.Get(jsonStr, "count").String()
	count, _ := strconv.Atoi(countStr)
	pageNum := int(math.Ceil(float64(count) / 31.0))

	configAll := gjson.Get(jsonStr, "config.all").Array()
	columns := make([]string, len(configAll))
	for i, item := range configAll {
		columns[i] = item.Get("1").String()
	}

	var allRows [][]string

	dataJSON := gjson.Get(jsonStr, "data")
	for _, item := range dataJSON.Array() {
		row := make([]string, len(columns))
		for i := range columns {
			row[i] = item.Get(strconv.Itoa(i)).String()
		}
		allRows = append(allRows, row)
	}

	for page := 1; page < pageNum; page++ {
		params["from"] = strconv.Itoa(page * 31)
		resp, err := utils.Get(baseURL, params)
		if err != nil {
			continue
		}

		jsonStr = extractJSONFromJSONP(resp.String())
		if jsonStr == "" {
			continue
		}

		dataJSON = gjson.Get(jsonStr, "data")
		for _, item := range dataJSON.Array() {
			row := make([]string, len(columns))
			for i := range columns {
				row[i] = item.Get(strconv.Itoa(i)).String()
			}
			allRows = append(allRows, row)
		}
	}

	if len(allRows) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("数据为空")
	}

	seriesList := make([]series.Series, len(columns))
	for i, col := range columns {
		vals := make([]string, len(allRows))
		for j, row := range allRows {
			vals[j] = row[i]
		}
		// 前两列为字符串类型
		if i < 2 {
			seriesList[i] = series.New(vals, series.String, col)
		} else {
			seriesList[i] = series.New(vals, series.Float, col)
		}
	}

	return dataframe.New(seriesList...), nil
}

// MacroChinaSupplyOfMoney 获取货币供应量
// 数据源: 新浪财经
// https://finance.sina.com.cn/mac/#fininfo-1-0-31-1
func MacroChinaSupplyOfMoney() (dataframe.DataFrame, error) {
	baseURL := "https://quotes.sina.cn/mac/api/jsonp_v3.php/SINAREMOTECALLCALLBACK/MacPage_Service.get_pagedata"

	params := map[string]string{
		"cate":      "fininfo",
		"event":     "1",
		"from":      "0",
		"num":       "31",
		"condition": "",
	}

	resp, err := utils.Get(baseURL, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("获取货币供应量数据失败: %w", err)
	}

	jsonStr := extractJSONFromJSONP(resp.String())
	if jsonStr == "" {
		return dataframe.DataFrame{}, fmt.Errorf("解析JSONP响应失败")
	}

	countStr := gjson.Get(jsonStr, "count").String()
	count, _ := strconv.Atoi(countStr)
	pageNum := int(math.Ceil(float64(count) / 31.0))

	configAll := gjson.Get(jsonStr, "config.all").Array()
	columns := make([]string, len(configAll))
	for i, item := range configAll {
		columns[i] = item.Get("1").String()
	}

	var allRows [][]string

	dataJSON := gjson.Get(jsonStr, "data")
	for _, item := range dataJSON.Array() {
		row := make([]string, len(columns))
		for i := range columns {
			row[i] = item.Get(strconv.Itoa(i)).String()
		}
		allRows = append(allRows, row)
	}

	for page := 1; page < pageNum; page++ {
		params["from"] = strconv.Itoa(page * 31)
		resp, err := utils.Get(baseURL, params)
		if err != nil {
			continue
		}

		jsonStr = extractJSONFromJSONP(resp.String())
		if jsonStr == "" {
			continue
		}

		dataJSON = gjson.Get(jsonStr, "data")
		for _, item := range dataJSON.Array() {
			row := make([]string, len(columns))
			for i := range columns {
				row[i] = item.Get(strconv.Itoa(i)).String()
			}
			allRows = append(allRows, row)
		}
	}

	if len(allRows) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("数据为空")
	}

	seriesList := make([]series.Series, len(columns))
	for i, col := range columns {
		vals := make([]string, len(allRows))
		for j, row := range allRows {
			vals[j] = row[i]
		}
		if i == 0 {
			seriesList[i] = series.New(vals, series.String, col)
		} else {
			seriesList[i] = series.New(vals, series.Float, col)
		}
	}

	return dataframe.New(seriesList...), nil
}

// MacroChinaForeignExchangeGold 获取央行黄金和外汇储备
// 数据源: 新浪财经
// https://finance.sina.com.cn/mac/#fininfo-5-0-31-2
func MacroChinaForeignExchangeGold() (dataframe.DataFrame, error) {
	baseURL := "https://quotes.sina.cn/mac/api/jsonp_v3.php/SINAREMOTECALLCALLBACK/MacPage_Service.get_pagedata"

	params := map[string]string{
		"cate":      "fininfo",
		"event":     "5",
		"from":      "0",
		"num":       "31",
		"condition": "",
	}

	resp, err := utils.Get(baseURL, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("获取央行黄金和外汇储备数据失败: %w", err)
	}

	jsonStr := extractJSONFromJSONP(resp.String())
	if jsonStr == "" {
		return dataframe.DataFrame{}, fmt.Errorf("解析JSONP响应失败")
	}

	countStr := gjson.Get(jsonStr, "count").String()
	count, _ := strconv.Atoi(countStr)
	pageNum := int(math.Ceil(float64(count) / 31.0))

	configAll := gjson.Get(jsonStr, "config.all").Array()
	columns := make([]string, len(configAll))
	for i, item := range configAll {
		columns[i] = item.Get("1").String()
	}

	var allRows [][]string

	dataJSON := gjson.Get(jsonStr, "data")
	for _, item := range dataJSON.Array() {
		row := make([]string, len(columns))
		for i := range columns {
			row[i] = item.Get(strconv.Itoa(i)).String()
		}
		allRows = append(allRows, row)
	}

	for page := 1; page < pageNum; page++ {
		params["from"] = strconv.Itoa(page * 31)
		resp, err := utils.Get(baseURL, params)
		if err != nil {
			continue
		}

		jsonStr = extractJSONFromJSONP(resp.String())
		if jsonStr == "" {
			continue
		}

		dataJSON = gjson.Get(jsonStr, "data")
		for _, item := range dataJSON.Array() {
			row := make([]string, len(columns))
			for i := range columns {
				row[i] = item.Get(strconv.Itoa(i)).String()
			}
			allRows = append(allRows, row)
		}
	}

	if len(allRows) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("数据为空")
	}

	seriesList := make([]series.Series, len(columns))
	for i, col := range columns {
		vals := make([]string, len(allRows))
		for j, row := range allRows {
			vals[j] = row[i]
		}
		if i == 0 {
			seriesList[i] = series.New(vals, series.String, col)
		} else {
			seriesList[i] = series.New(vals, series.Float, col)
		}
	}

	df := dataframe.New(seriesList...)
	df = df.Arrange(dataframe.Sort("统计时间"))
	return df, nil
}

// MacroChinaRetailPriceIndex 获取商品零售价格指数
// 数据源: 新浪财经
// https://finance.sina.com.cn/mac/#price-12-0-31-1
func MacroChinaRetailPriceIndex() (dataframe.DataFrame, error) {
	baseURL := "https://quotes.sina.cn/mac/api/jsonp_v3.php/SINAREMOTECALLCALLBACK/MacPage_Service.get_pagedata"

	params := map[string]string{
		"cate":      "price",
		"event":     "12",
		"from":      "0",
		"num":       "31",
		"condition": "",
	}

	resp, err := utils.Get(baseURL, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("获取商品零售价格指数数据失败: %w", err)
	}

	jsonStr := extractJSONFromJSONP(resp.String())
	if jsonStr == "" {
		return dataframe.DataFrame{}, fmt.Errorf("解析JSONP响应失败")
	}

	countStr := gjson.Get(jsonStr, "count").String()
	count, _ := strconv.Atoi(countStr)
	pageNum := int(math.Ceil(float64(count) / 31.0))

	configAll := gjson.Get(jsonStr, "config.all").Array()
	columns := make([]string, len(configAll))
	for i, item := range configAll {
		columns[i] = item.Get("1").String()
	}

	var allRows [][]string

	dataJSON := gjson.Get(jsonStr, "data")
	for _, item := range dataJSON.Array() {
		row := make([]string, len(columns))
		for i := range columns {
			row[i] = item.Get(strconv.Itoa(i)).String()
		}
		allRows = append(allRows, row)
	}

	for page := 1; page < pageNum; page++ {
		params["from"] = strconv.Itoa(page * 31)
		resp, err := utils.Get(baseURL, params)
		if err != nil {
			continue
		}

		jsonStr = extractJSONFromJSONP(resp.String())
		if jsonStr == "" {
			continue
		}

		dataJSON = gjson.Get(jsonStr, "data")
		for _, item := range dataJSON.Array() {
			row := make([]string, len(columns))
			for i := range columns {
				row[i] = item.Get(strconv.Itoa(i)).String()
			}
			allRows = append(allRows, row)
		}
	}

	if len(allRows) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("数据为空")
	}

	seriesList := make([]series.Series, len(columns))
	for i, col := range columns {
		vals := make([]string, len(allRows))
		for j, row := range allRows {
			vals[j] = row[i]
		}
		if i == 0 {
			seriesList[i] = series.New(vals, series.String, col)
		} else {
			seriesList[i] = series.New(vals, series.Float, col)
		}
	}

	df := dataframe.New(seriesList...)
	// 按统计月份排序
	if df.Ncol() > 0 {
		sortCol := columns[0]
		df = df.Arrange(dataframe.Sort(sortCol))
	}
	return df, nil
}

// MacroChinaRealEstate 获取国房景气指数
// 数据源: 东方财富
// https://data.eastmoney.com/cjsj/hyzs_list_EMM00121987.html
func MacroChinaRealEstate() (dataframe.DataFrame, error) {
	return fetchEastmoneyIndustryIndex("EMM00121987")
}
