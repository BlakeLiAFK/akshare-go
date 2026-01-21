package economic

import (
	"fmt"
	"sort"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
	"github.com/tidwall/gjson"
)

// MacroConsGold 全球最大黄金ETF-SPDR Gold Trust持仓报告
//
// 数据区间从 20041118-至今
//
// 返回:
//   - dataframe.DataFrame: 持仓报告数据
//   - error: 错误信息
//
// 数据源: https://datacenter.jin10.com/reportType/dc_etf_gold
func MacroConsGold() (dataframe.DataFrame, error) {
	return fetchETFData("1", "黄金")
}

// MacroConsSilver 全球最大白银ETF-iShares Silver Trust持仓报告
//
// 数据区间从 20041118-至今
//
// 返回:
//   - dataframe.DataFrame: 持仓报告数据
//   - error: 错误信息
//
// 数据源: https://datacenter.jin10.com/reportType/dc_etf_sliver
func MacroConsSilver() (dataframe.DataFrame, error) {
	return fetchETFData("2", "白银")
}

// fetchETFData 获取ETF持仓数据的通用函数
func fetchETFData(attrID, commodity string) (dataframe.DataFrame, error) {
	t := time.Now().UnixMilli()
	url := "https://datacenter-api.jin10.com/reports/list_v2"

	headers := map[string]string{
		"user-agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36",
		"x-app-id":     "rU6QIu7JHe2gOUeR",
		"x-csrf-token": "x-csrf-token",
		"x-version":    "1.0.0",
	}

	// 定义列名
	columns := []string{"商品", "日期", "总库存", "增持/减持", "总价值"}

	var allRecords [][]string
	allRecords = append(allRecords, columns)

	maxDate := ""
	for {
		params := map[string]string{
			"category": "etf",
			"attr_id":  attrID,
			"max_date": maxDate,
			"_":        fmt.Sprintf("%d", t),
		}

		resp, err := utils.GetWithHeaders(url, params, headers)
		if err != nil {
			return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
		}

		valuesArray := gjson.Get(resp.String(), "data.values").Array()
		if len(valuesArray) == 0 {
			break
		}

		for _, row := range valuesArray {
			rowArray := row.Array()
			if len(rowArray) < 4 {
				continue
			}

			record := []string{
				commodity,
				rowArray[0].String(), // 日期
				rowArray[1].String(), // 总库存
				rowArray[2].String(), // 增持/减持
				rowArray[3].String(), // 总价值
			}
			allRecords = append(allRecords, record)
		}

		// 获取最后一条记录的日期，减一天作为下次请求的 max_date
		lastDate := valuesArray[len(valuesArray)-1].Array()[0].String()
		parsedDate, err := time.Parse("2006-01-02", lastDate)
		if err != nil {
			break
		}
		maxDate = parsedDate.AddDate(0, 0, -1).Format("2006-01-02")
	}

	// 创建 DataFrame
	df := dataframe.LoadRecords(allRecords)

	// 按日期排序
	df = df.Arrange(dataframe.Sort("日期"))

	return df, nil
}

// MacroConsOpecMonth 欧佩克报告-月度
//
// 数据区间从 20170118-至今
// 这里返回的具体索引日期的数据为上一个月的数据
//
// 返回:
//   - dataframe.DataFrame: 欧佩克报告数据
//   - error: 错误信息
//
// 数据源: https://datacenter.jin10.com/reportType/dc_opec_report
func MacroConsOpecMonth() (dataframe.DataFrame, error) {
	t := time.Now().UnixMilli()

	headers := map[string]string{
		"accept":          "*/*",
		"accept-encoding": "gzip, deflate, br",
		"accept-language": "zh-CN,zh;q=0.9,en;q=0.8",
		"cache-control":   "no-cache",
		"origin":          "https://datacenter.jin10.com",
		"pragma":          "no-cache",
		"referer":         "https://datacenter.jin10.com/reportType/dc_opec_report",
		"sec-fetch-mode":  "cors",
		"sec-fetch-site":  "same-site",
		"user-agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.117 Safari/537.36",
		"x-app-id":        "rU6QIu7JHe2gOUeR",
		"x-csrf-token":    "",
		"x-version":       "1.0.0",
	}

	// 获取日期列表
	datesURL := fmt.Sprintf("https://datacenter-api.jin10.com/reports/dates?category=opec&_=%d", t)
	datesResp, err := utils.GetWithHeaders(datesURL, nil, headers)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("获取日期列表失败: %w", err)
	}

	dateList := gjson.Get(datesResp.String(), "data").Array()
	if len(dateList) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到日期数据")
	}

	// 目标国家列表
	targetCountries := []string{
		"阿尔及利亚", "安哥拉", "加蓬", "伊朗", "伊拉克",
		"科威特", "利比亚", "尼日利亚", "沙特", "阿联酋",
		"委内瑞拉", "欧佩克产量",
	}

	// 定义列名
	columns := append([]string{"日期"}, targetCountries...)

	var allRecords [][]string
	allRecords = append(allRecords, columns)

	// 逆序遍历日期
	for i := len(dateList) - 1; i >= 0; i-- {
		dateStr := dateList[i].String()

		dataURL := fmt.Sprintf("https://datacenter-api.jin10.com/reports/list?category=opec&date=%s&_=%d", dateStr, t)
		dataResp, err := utils.GetWithHeaders(dataURL, nil, headers)
		if err != nil {
			continue
		}

		// 解析数据
		keysArray := gjson.Get(dataResp.String(), "data.keys").Array()
		valuesArray := gjson.Get(dataResp.String(), "data.values").Array()

		if len(keysArray) == 0 || len(valuesArray) == 0 {
			continue
		}

		// 构建列名映射
		colNames := make([]string, len(keysArray))
		for j, key := range keysArray {
			colNames[j] = key.Get("name").String()
		}

		// 构建数据 map (转置)
		dataMap := make(map[string][]string)
		for _, row := range valuesArray {
			rowArray := row.Array()
			for j, cell := range rowArray {
				if j < len(colNames) {
					dataMap[colNames[j]] = append(dataMap[colNames[j]], cell.String())
				}
			}
		}

		// 提取目标国家数据（取倒数第二行，如果没有则取最后一行）
		record := []string{dateStr}
		for _, country := range targetCountries {
			if values, ok := dataMap[country]; ok && len(values) > 0 {
				// 取倒数第二个值，如果没有则取最后一个
				idx := len(values) - 2
				if idx < 0 {
					idx = len(values) - 1
				}
				record = append(record, values[idx])
			} else {
				record = append(record, "")
			}
		}
		allRecords = append(allRecords, record)
	}

	// 创建 DataFrame
	df := dataframe.LoadRecords(allRecords)

	// 按日期排序
	df = df.Arrange(dataframe.Sort("日期"))

	return df, nil
}

// MacroConsOpecMonthWithLimit 欧佩克报告-月度（限制获取数量）
//
// 由于数据量较大，提供限制获取数量的版本
//
// 参数:
//   - limit: 最多获取的月份数量，0表示全部
//
// 返回:
//   - dataframe.DataFrame: 欧佩克报告数据
//   - error: 错误信息
func MacroConsOpecMonthWithLimit(limit int) (dataframe.DataFrame, error) {
	t := time.Now().UnixMilli()

	headers := map[string]string{
		"accept":          "*/*",
		"accept-encoding": "gzip, deflate, br",
		"accept-language": "zh-CN,zh;q=0.9,en;q=0.8",
		"cache-control":   "no-cache",
		"origin":          "https://datacenter.jin10.com",
		"pragma":          "no-cache",
		"referer":         "https://datacenter.jin10.com/reportType/dc_opec_report",
		"user-agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.117 Safari/537.36",
		"x-app-id":        "rU6QIu7JHe2gOUeR",
		"x-csrf-token":    "",
		"x-version":       "1.0.0",
	}

	// 获取日期列表
	datesURL := fmt.Sprintf("https://datacenter-api.jin10.com/reports/dates?category=opec&_=%d", t)
	datesResp, err := utils.GetWithHeaders(datesURL, nil, headers)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("获取日期列表失败: %w", err)
	}

	dateList := gjson.Get(datesResp.String(), "data").Array()
	if len(dateList) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到日期数据")
	}

	// 限制数量
	if limit > 0 && limit < len(dateList) {
		dateList = dateList[len(dateList)-limit:]
	}

	// 目标国家列表
	targetCountries := []string{
		"阿尔及利亚", "安哥拉", "加蓬", "伊朗", "伊拉克",
		"科威特", "利比亚", "尼日利亚", "沙特", "阿联酋",
		"委内瑞拉", "欧佩克产量",
	}

	// 定义列名
	columns := append([]string{"日期"}, targetCountries...)

	var allRecords [][]string
	allRecords = append(allRecords, columns)

	// 遍历日期
	for _, dateItem := range dateList {
		dateStr := dateItem.String()

		dataURL := fmt.Sprintf("https://datacenter-api.jin10.com/reports/list?category=opec&date=%s&_=%d", dateStr, t)
		dataResp, err := utils.GetWithHeaders(dataURL, nil, headers)
		if err != nil {
			continue
		}

		// 解析数据
		keysArray := gjson.Get(dataResp.String(), "data.keys").Array()
		valuesArray := gjson.Get(dataResp.String(), "data.values").Array()

		if len(keysArray) == 0 || len(valuesArray) == 0 {
			continue
		}

		// 构建列名映射
		colNames := make([]string, len(keysArray))
		for j, key := range keysArray {
			colNames[j] = key.Get("name").String()
		}

		// 构建数据 map (转置)
		dataMap := make(map[string][]string)
		for _, row := range valuesArray {
			rowArray := row.Array()
			for j, cell := range rowArray {
				if j < len(colNames) {
					dataMap[colNames[j]] = append(dataMap[colNames[j]], cell.String())
				}
			}
		}

		// 提取目标国家数据
		record := []string{dateStr}
		for _, country := range targetCountries {
			if values, ok := dataMap[country]; ok && len(values) > 0 {
				idx := len(values) - 2
				if idx < 0 {
					idx = len(values) - 1
				}
				record = append(record, values[idx])
			} else {
				record = append(record, "")
			}
		}
		allRecords = append(allRecords, record)
	}

	// 创建 DataFrame
	df := dataframe.LoadRecords(allRecords)

	// 按日期排序
	dates := make([]string, df.Nrow())
	for i := 0; i < df.Nrow(); i++ {
		dates[i] = df.Elem(i, 0).String()
	}
	sort.Strings(dates)

	return df, nil
}
