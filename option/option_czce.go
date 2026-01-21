package option

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

// 郑商所期权品种上市时间映射
var czceOptionListingYear = map[string]string{
	"SR": "2017", // 白糖
	"CF": "2019", // 棉花
	"TA": "2019", // PTA
	"MA": "2019", // 甲醇
	"RM": "2020", // 菜籽粕
	"ZC": "2020", // 动力煤
	"OI": "2022", // 菜籽油
	"PK": "2022", // 花生
	"PX": "2023", // 对二甲苯
	"SH": "2023", // 烧碱
	"SA": "2023", // 纯碱
	"PF": "2023", // 短纤
	"SM": "2023", // 锰硅
	"SF": "2023", // 硅铁
	"UR": "2023", // 尿素
	"AP": "2023", // 苹果
	"CJ": "2024", // 红枣
	"FG": "2024", // 玻璃
	"PR": "2024", // 瓶片
}

// 郑商所期权品种名称映射
var czceOptionNames = map[string]string{
	"SR": "白糖",
	"CF": "棉花",
	"TA": "PTA",
	"MA": "甲醇",
	"RM": "菜籽粕",
	"ZC": "动力煤",
	"OI": "菜籽油",
	"PK": "花生",
	"PX": "对二甲苯",
	"SH": "烧碱",
	"SA": "纯碱",
	"PF": "短纤",
	"SM": "锰硅",
	"SF": "硅铁",
	"UR": "尿素",
	"AP": "苹果",
	"CJ": "红枣",
	"FG": "玻璃",
	"PR": "瓶片",
}

// OptionHistYearlyCzce 郑州商品交易所-交易数据-历史行情下载-期权历史行情下载
//
// 参数:
//   - symbol: 期权品种代码，如 "SR", "CF", "TA" 等
//   - year: 需要获取数据的年份，注意品种的上市时间
//
// 返回:
//   - dataframe.DataFrame: 指定年份的日频期权数据
//   - error: 错误信息
//
// 数据源: http://www.czce.com.cn/cn/jysj/lshqxz/H770319index_1.htm
func OptionHistYearlyCzce(symbol, year string) (dataframe.DataFrame, error) {
	if symbol == "" {
		return dataframe.DataFrame{}, fmt.Errorf("品种代码不能为空")
	}
	if year == "" {
		return dataframe.DataFrame{}, fmt.Errorf("年份不能为空")
	}

	// 检查品种是否有效
	listingYear, exists := czceOptionListingYear[symbol]
	if !exists {
		return dataframe.DataFrame{}, fmt.Errorf("不支持的品种代码: %s", symbol)
	}

	// 检查年份是否在上市时间之前
	listingYearInt, err1 := strconv.Atoi(listingYear)
	yearInt, err2 := strconv.Atoi(year)
	if err1 != nil || err2 != nil {
		return dataframe.DataFrame{}, fmt.Errorf("年份格式错误")
	}

	if yearInt < listingYearInt {
		name, _ := czceOptionNames[symbol]
		return dataframe.DataFrame{}, fmt.Errorf("%s年，品种%s(%s)尚未上市", year, name, symbol)
	}

	// 构建URL
	url := fmt.Sprintf("http://www.czce.com.cn/cn/DFSStaticFiles/Option/%s/OptionDataAllHistory/%sOPTIONS%s.txt", year, symbol, year)

	// 请求数据
	resp, err := utils.Get(url, nil)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	// 解析文本数据
	text := resp.String()
	if text == "" {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	lines := strings.Split(text, "\n")
	if len(lines) <= 1 {
		return dataframe.DataFrame{}, fmt.Errorf("数据格式错误，行数不足")
	}

	// 跳过第一行（标题行），从第二行开始解析数据
	var records [][]string
	for i := 1; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}

		// 按管道符分割字段
		fields := strings.Split(line, "|")

		// 清理字段
		var cleanFields []string
		for _, field := range fields {
			cleanField := strings.TrimSpace(field)
			cleanFields = append(cleanFields, cleanField)
		}

		if len(cleanFields) > 0 {
			records = append(records, cleanFields)
		}
	}

	if len(records) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("没有有效的数据记录")
	}

	// 获取列名（从第一行解析）
	headers := strings.Split(strings.TrimSpace(lines[0]), "|")
	for i, header := range headers {
		headers[i] = strings.TrimSpace(header)
	}

	// 确保所有记录的列数与表头一致
	maxCols := len(headers)
	for i := range records {
		for len(records[i]) < maxCols {
			records[i] = append(records[i], "")
		}
	}

	// 转换数值列
	records = convertNumericRecords(records, headers)

	// 构建DataFrame
	allRecords := append([][]string{headers}, records...)
	df := dataframe.LoadRecords(allRecords)

	return df, nil
}

// convertNumericRecords 转换数值记录
func convertNumericRecords(records [][]string, headers []string) [][]string {
	// 假设数值列（根据实际数据结构调整）
	// 这里需要根据实际的郑商所期权数据格式来确定哪些列是数值列
	numericCols := make(map[int]bool)

	// 根据列名判断是否为数值列
	for i, header := range headers {
		if strings.Contains(header, "价") || strings.Contains(header, "量") ||
			strings.Contains(header, "额") || strings.Contains(header, "率") ||
			strings.Contains(header, "数") || strings.Contains(header, "日") {
			numericCols[i] = true
		}
	}

	// 转换数值
	for i := range records {
		for colIdx := range numericCols {
			if colIdx < len(records[i]) {
				records[i][colIdx] = parseNumericValue(records[i][colIdx])
			}
		}
	}

	return records
}

// GetCzceOptionSymbols 获取郑商所期权品种列表
//
// 返回:
//   - dataframe.DataFrame: 包含品种代码和名称的数据
func GetCzceOptionSymbols() (dataframe.DataFrame, error) {
	var symbols []string
	var names []string
	var listingYears []string

	for symbol, year := range czceOptionListingYear {
		symbols = append(symbols, symbol)
		name, exists := czceOptionNames[symbol]
		if !exists {
			name = symbol
		}
		names = append(names, name)
		listingYears = append(listingYears, year)
	}

	df := dataframe.New(
		series.New(symbols, series.String, "symbol"),
		series.New(names, series.String, "name"),
		series.New(listingYears, series.String, "listing_year"),
	)

	return df, nil
}
