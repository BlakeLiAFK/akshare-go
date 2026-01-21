package air

import (
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-gota/gota/dataframe"
)

// SunriseCityList 获取支持查询日出日落数据的城市列表
//
// 数据源: Time and Date
// URL: https://www.timeanddate.com/astronomy/china
//
// 返回:
//   - []string: 城市名称列表（小写英文）
//   - error: 错误信息
//
// 示例:
//
//	cities, err := air.SunriseCityList()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(cities)
func SunriseCityList() ([]string, error) {
	url := "https://www.timeanddate.com/astronomy/china"

	// 发起HTTP请求
	resp, err := utils.Get(url, nil)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	// 解析HTML
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	var cities []string
	citySet := make(map[string]bool)

	// 查找所有表格中的城市
	doc.Find("table").Each(func(i int, table *goquery.Selection) {
		table.Find("td a").Each(func(j int, link *goquery.Selection) {
			city := strings.ToLower(strings.TrimSpace(link.Text()))
			if city != "" && !citySet[city] {
				cities = append(cities, city)
				citySet[city] = true
			}
		})
	})

	return cities, nil
}

// SunriseDaily 获取指定城市指定日期的日出日落数据
//
// 数据源: Time and Date
// URL: https://www.timeanddate.com/sun/china/{city}
//
// 参数:
//   - date: 日期，格式 "20240428"
//   - city: 城市名称（小写英文），如 "beijing"
//
// 返回:
//   - dataframe.DataFrame: 包含日期、日出、日落、日照时长等数据
//   - error: 错误信息
//
// 示例:
//
//	df, err := air.SunriseDaily("20240428", "beijing")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(df)
func SunriseDaily(date, city string) (dataframe.DataFrame, error) {
	// 验证城市
	cities, err := SunriseCityList()
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("获取城市列表失败: %w", err)
	}

	cityLower := strings.ToLower(city)
	found := false
	for _, c := range cities {
		if c == cityLower {
			found = true
			break
		}
	}
	if !found {
		return dataframe.DataFrame{}, fmt.Errorf("不支持的城市: %s", city)
	}

	// 解析日期
	if len(date) != 8 {
		return dataframe.DataFrame{}, fmt.Errorf("日期格式错误，应为 YYYYMMDD")
	}
	year := date[:4]
	month := date[4:6]
	day := date[6:]

	// 构建URL
	url := fmt.Sprintf("https://www.timeanddate.com/sun/china/%s?month=%s&year=%s",
		cityLower, month, year)

	// 发起HTTP请求（禁用SSL验证）
	resp, err := utils.Get(url, nil)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	// 解析HTML表格
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("解析HTML失败: %w", err)
	}

	// 查找数据表
	var targetTable *goquery.Selection
	doc.Find("table").Each(func(i int, table *goquery.Selection) {
		// 查找包含日出日落数据的表格
		if table.Find("th:contains('Sunrise')").Length() > 0 {
			targetTable = table
		}
	})

	if targetTable == nil {
		return dataframe.DataFrame{}, fmt.Errorf("未找到日出日落数据表格")
	}

	// 解析表格找到指定日期的行
	var dayRow *goquery.Selection
	targetTable.Find("tr").Each(func(i int, row *goquery.Selection) {
		firstCell := row.Find("td").First().Text()
		firstCell = strings.TrimSpace(firstCell)
		if firstCell == day || firstCell == strings.TrimLeft(day, "0") {
			dayRow = row
		}
	})

	if dayRow == nil {
		return dataframe.DataFrame{}, fmt.Errorf("未找到日期 %s 的数据", date)
	}

	// 提取数据
	var cells []string
	dayRow.Find("td").Each(func(i int, cell *goquery.Selection) {
		cells = append(cells, strings.TrimSpace(cell.Text()))
	})

	if len(cells) < 4 {
		return dataframe.DataFrame{}, fmt.Errorf("数据格式错误")
	}

	// 构建记录
	record := map[string]interface{}{
		"date":      date,
		"day":       cells[0],
		"sunrise":   cells[1],
		"sunset":    cells[2],
		"daylength": cells[3],
	}

	// 创建DataFrame
	df := dataframe.LoadMaps([]map[string]interface{}{record})
	return df, nil
}

// SunriseMonthly 获取指定月份的每日日出日落数据
//
// 数据源: Time and Date
//
// 参数:
//   - date: 日期（用于指定月份），格式 "20240428"
//   - city: 城市名称（小写英文），如 "beijing"
//
// 返回:
//   - dataframe.DataFrame: 包含整月的日出日落数据
//   - error: 错误信息
//
// 示例:
//
//	df, err := air.SunriseMonthly("20240428", "beijing")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(df)
func SunriseMonthly(date, city string) (dataframe.DataFrame, error) {
	// 验证城市
	cities, err := SunriseCityList()
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("获取城市列表失败: %w", err)
	}

	cityLower := strings.ToLower(city)
	found := false
	for _, c := range cities {
		if c == cityLower {
			found = true
			break
		}
	}
	if !found {
		return dataframe.DataFrame{}, fmt.Errorf("不支持的城市: %s", city)
	}

	// 解析日期
	if len(date) != 8 {
		return dataframe.DataFrame{}, fmt.Errorf("日期格式错误，应为 YYYYMMDD")
	}
	year := date[:4]
	month := date[4:6]

	// 构建URL
	url := fmt.Sprintf("https://www.timeanddate.com/sun/china/%s?month=%s&year=%s",
		cityLower, month, year)

	// 发起HTTP请求
	resp, err := utils.Get(url, nil)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	// 解析HTML表格
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("解析HTML失败: %w", err)
	}

	// 查找数据表
	var targetTable *goquery.Selection
	doc.Find("table").Each(func(i int, table *goquery.Selection) {
		if table.Find("th:contains('Sunrise')").Length() > 0 {
			targetTable = table
		}
	})

	if targetTable == nil {
		return dataframe.DataFrame{}, fmt.Errorf("未找到日出日落数据表格")
	}

	// 解析所有数据行
	var records []map[string]interface{}
	targetTable.Find("tr").Each(func(i int, row *goquery.Selection) {
		if i == 0 {
			// 跳过表头
			return
		}

		var cells []string
		row.Find("td").Each(func(j int, cell *goquery.Selection) {
			cells = append(cells, strings.TrimSpace(cell.Text()))
		})

		if len(cells) >= 4 && cells[0] != "" {
			dayNum := strings.TrimSpace(cells[0])
			if len(dayNum) == 1 {
				dayNum = "0" + dayNum
			}
			dateStr := year + month + dayNum

			record := map[string]interface{}{
				"date":      dateStr,
				"day":       cells[0],
				"sunrise":   cells[1],
				"sunset":    cells[2],
				"daylength": cells[3],
			}
			records = append(records, record)
		}
	})

	if len(records) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	// 创建DataFrame
	df := dataframe.LoadMaps(records)
	return df, nil
}
