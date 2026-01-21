package other

import (
	"fmt"
	"strconv"

	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
)

// CarSaleRankGasgoo 盖世汽车-汽车行业制造企业数据库-销量数据
//
// 参数:
//   - symbol: 排行榜类型，"车企榜"/"品牌榜"/"车型榜"
//   - date: 查询的年份和月份，格式 "202309"
//
// 返回:
//   - []CarSaleRankGasgooItem: 销量排名数据
//   - error: 错误信息
func CarSaleRankGasgoo(symbol string, date string) ([]CarSaleRankGasgooItem, error) {
	// 解析日期
	if len(date) != 6 {
		return nil, fmt.Errorf("日期格式错误，应为6位数字如202309")
	}
	year := date[:4]
	month := date[4:6]
	monthInt, err := strconv.Atoi(month)
	if err != nil {
		return nil, fmt.Errorf("月份解析失败: %w", err)
	}

	// 获取排行榜类型代码
	rankType, ok := rankTypeMap[symbol]
	if !ok {
		return nil, fmt.Errorf("不支持的排行榜类型: %s", symbol)
	}

	// 构建请求体
	payload := map[string]interface{}{
		"countryID":    "",
		"endM":         monthInt,
		"endY":         year,
		"energy":       "",
		"modelGradeID": "",
		"modelTypeID":  "",
		"orderBy":      fmt.Sprintf("%s-%d", year, monthInt),
		"queryDate":    fmt.Sprintf("%s-%d", year, monthInt),
		"rankType":     rankType,
		"startY":       year,
		"startM":       monthInt,
	}

	// 发送请求
	resp, err := resty.New().R().
		SetHeader("Content-Type", "application/json; charset=UTF-8").
		SetHeader("Accept", "application/json, text/javascript, */*; q=0.01").
		SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36").
		SetBody(payload).
		Post(GasgooSalesRankURL)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	json := resp.String()

	// 获取数据（在"d"字段中）
	dataStr := gjson.Get(json, "d").String()
	if dataStr == "" {
		return nil, fmt.Errorf("未获取到数据")
	}

	// 解码JSON字符串（gjson会自动处理）
	dataResult := gjson.Parse(dataStr)
	dataList := dataResult.Array()

	if len(dataList) == 0 {
		return nil, fmt.Errorf("数据列表为空")
	}

	var result []CarSaleRankGasgooItem

	for _, item := range dataList {
		itemMap := item.Map()

		rank := 0
		if r := itemMap["rank"]; r.Exists() {
			rank = int(r.Int())
		}

		name := ""
		if n := itemMap["name"]; n.Exists() {
			name = n.String()
		} else if n := itemMap["厂商"]; n.Exists() {
			name = n.String()
		} else if n := itemMap["品牌"]; n.Exists() {
			name = n.String()
		} else if n := itemMap["车型"]; n.Exists() {
			name = n.String()
		}

		sales := 0.0
		if s := itemMap["sales"]; s.Exists() {
			sales, _ = strconv.ParseFloat(s.String(), 64)
		} else if s := itemMap["销量"]; s.Exists() {
			sales, _ = strconv.ParseFloat(s.String(), 64)
		}

		yearOnYear := 0.0
		if y := itemMap["year_on_year"]; y.Exists() {
			yearOnYear, _ = strconv.ParseFloat(y.String(), 64)
		} else if y := itemMap["同比"]; y.Exists() {
			yearOnYear, _ = strconv.ParseFloat(y.String(), 64)
		}

		monthOnMonth := 0.0
		if m := itemMap["month_on_month"]; m.Exists() {
			monthOnMonth, _ = strconv.ParseFloat(m.String(), 64)
		} else if m := itemMap["环比"]; m.Exists() {
			monthOnMonth, _ = strconv.ParseFloat(m.String(), 64)
		}

		result = append(result, CarSaleRankGasgooItem{
			Rank:         rank,
			Name:         name,
			Sales:        sales,
			YearOnYear:   yearOnYear,
			MonthOnMonth: monthOnMonth,
		})
	}

	return result, nil
}
