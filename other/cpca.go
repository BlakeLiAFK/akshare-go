package other

import (
	"fmt"
	"strconv"

	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
)

// CarMarketTotalCPCA 乘联会-统计数据-总体市场
//
// 参数:
//   - symbol: 车型类型，"狭义乘用车" 或 "广义乘用车"
//   - indicator: 指标类型，"产量"/"批发"/"零售"/"出口"
//
// 返回:
//   - []CarMarketTotalItem: 总体市场数据
//   - error: 错误信息
func CarMarketTotalCPCA(symbol string, indicator string) ([]CarMarketTotalItem, error) {
	// 获取数据
	url := CPCAChartListURL
	resp, err := resty.New().R().
		SetQueryParam("charttype", "1").
		Get(url)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	json := resp.String()

	// 根据symbol选择数据数组
	dataIndex := 0
	if symbol == CarTypeWide {
		dataIndex = 1
	}
	dataList := gjson.Get(json, fmt.Sprintf("%d.dataList", dataIndex)).Array()

	if len(dataList) == 0 {
		return nil, fmt.Errorf("未获取到数据")
	}

	// 获取列名
	firstItem := dataList[0].Map()
	columns := make([]string, 0, len(firstItem))
	for k := range firstItem {
		columns = append(columns, k)
	}

	// 获取indicator对应的列索引
	indicatorIndex := indicatorMap[indicator]

	var result []CarMarketTotalItem

	for _, item := range dataList {
		itemMap := item.Map()

		// 获取月份
		month := itemMap["month"].String()

		// 获取当年数值
		currentYearVal := itemMap[columns[1]]
		var currentYear float64
		if currentYearVal.Exists() && currentYearVal.String() != "" {
			currentYear, _ = strconv.ParseFloat(currentYearVal.String(), 64)
		}

		// 获取去年数值
		previousYearVal := itemMap[columns[2]]
		var previousYear float64
		if previousYearVal.Exists() {
			if previousYearVal.IsArray() {
				arr := previousYearVal.Array()
				if len(arr) > indicatorIndex {
					previousYear, _ = strconv.ParseFloat(arr[indicatorIndex].String(), 64)
				}
			} else {
				previousYear, _ = strconv.ParseFloat(previousYearVal.String(), 64)
			}
		}

		result = append(result, CarMarketTotalItem{
			Month:        month,
			CurrentYear:  currentYear,
			PreviousYear: previousYear,
		})
	}

	return result, nil
}

// carMarketManRankCPCAPifa 乘联会-统计数据-厂商排名-批发
func carMarketManRankCPCAPifa(symbol string) ([]CarMarketManRankItem, error) {
	url := CPCAChartListURL
	resp, err := resty.New().R().
		SetQueryParam("charttype", "2").
		Get(url)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	json := resp.String()

	// 根据symbol选择数据数组
	dataIndex := manRankSymbolMap[symbol]
	dataList := gjson.Get(json, fmt.Sprintf("%d.dataList", dataIndex)).Array()

	if len(dataList) == 0 {
		return nil, fmt.Errorf("未获取到数据")
	}

	// 获取列名
	firstItem := dataList[0].Map()
	columns := make([]string, 0, len(firstItem))
	for k := range firstItem {
		columns = append(columns, k)
	}

	var result []CarMarketManRankItem

	for _, item := range dataList {
		itemMap := item.Map()

		// 获取厂商名称
		manufacturer := itemMap["厂商"].String()

		// 获取当年数值（第一列数据的第一个值）
		currentYearVal := itemMap[columns[1]]
		var currentYear float64
		if currentYearVal.Exists() && currentYearVal.IsArray() {
			arr := currentYearVal.Array()
			if len(arr) > 0 {
				currentYear, _ = strconv.ParseFloat(arr[0].String(), 64)
			}
		}

		// 获取去年数值（第二列数据的第一个值）
		previousYearVal := itemMap[columns[2]]
		var previousYear float64
		if previousYearVal.Exists() && previousYearVal.IsArray() {
			arr := previousYearVal.Array()
			if len(arr) > 0 {
				previousYear, _ = strconv.ParseFloat(arr[0].String(), 64)
			}
		}

		result = append(result, CarMarketManRankItem{
			Manufacturer: manufacturer,
			CurrentYear:  currentYear,
			PreviousYear: previousYear,
		})
	}

	return result, nil
}

// carMarketManRankCPCALingshou 乘联会-统计数据-厂商排名-零售
func carMarketManRankCPCALingshou(symbol string) ([]CarMarketManRankItem, error) {
	url := CPCAChartList2URL
	resp, err := resty.New().R().
		SetQueryParam("charttype", "2").
		Get(url)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	json := resp.String()

	// 根据symbol选择数据数组
	dataIndex := manRankSymbolMap[symbol]
	dataList := gjson.Get(json, fmt.Sprintf("%d.dataList", dataIndex)).Array()

	if len(dataList) == 0 {
		return nil, fmt.Errorf("未获取到数据")
	}

	// 获取列名
	firstItem := dataList[0].Map()
	columns := make([]string, 0, len(firstItem))
	for k := range firstItem {
		columns = append(columns, k)
	}

	var result []CarMarketManRankItem

	for _, item := range dataList {
		itemMap := item.Map()

		// 获取厂商名称
		manufacturer := itemMap["厂商"].String()

		// 获取当年数值（第一列数据的第二个值）
		currentYearVal := itemMap[columns[1]]
		var currentYear float64
		if currentYearVal.Exists() && currentYearVal.IsArray() {
			arr := currentYearVal.Array()
			if len(arr) > 1 {
				currentYear, _ = strconv.ParseFloat(arr[1].String(), 64)
			}
		}

		// 获取去年数值（第二列数据的第二个值）
		previousYearVal := itemMap[columns[2]]
		var previousYear float64
		if previousYearVal.Exists() && previousYearVal.IsArray() {
			arr := previousYearVal.Array()
			if len(arr) > 1 {
				previousYear, _ = strconv.ParseFloat(arr[1].String(), 64)
			}
		}

		result = append(result, CarMarketManRankItem{
			Manufacturer: manufacturer,
			CurrentYear:  currentYear,
			PreviousYear: previousYear,
		})
	}

	return result, nil
}

// CarMarketManRankCPCA 乘联会-统计数据-厂商排名
//
// 参数:
//   - symbol: 统计类型，如 "狭义乘用车-单月"/"狭义乘用车-累计"/"广义乘用车-单月"/"广义乘用车-累计"
//   - indicator: 指标类型，"批发" 或 "零售"
//
// 返回:
//   - []CarMarketManRankItem: 厂商排名数据
//   - error: 错误信息
func CarMarketManRankCPCA(symbol string, indicator string) ([]CarMarketManRankItem, error) {
	if indicator == IndicatorWholesale {
		return carMarketManRankCPCAPifa(symbol)
	}
	return carMarketManRankCPCALingshou(symbol)
}

// carMarketCateCPCAPifa 乘联会-统计数据-车型大类-批发
func carMarketCateCPCAPifa(symbol string) ([]CarMarketCateItem, error) {
	url := CPCAChartListURL
	resp, err := resty.New().R().
		SetQueryParam("charttype", "3").
		Get(url)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	json := resp.String()

	// 根据symbol选择数据数组
	dataIndex := categorySymbolMap[symbol]
	dataList := gjson.Get(json, fmt.Sprintf("%d.dataList", dataIndex)).Array()

	if len(dataList) == 0 {
		return nil, fmt.Errorf("未获取到数据")
	}

	// 获取列名
	firstItem := dataList[0].Map()
	columns := make([]string, 0, len(firstItem))
	for k := range firstItem {
		columns = append(columns, k)
	}

	var result []CarMarketCateItem

	for _, item := range dataList {
		itemMap := item.Map()

		// 获取月份
		month := itemMap["month"].String()
		if month == "" {
			month = itemMap["月份"].String()
		}

		// 获取当年数值（columns[2] 的索引1）
		currentYearVal := itemMap[columns[2]]
		var currentYear float64
		if currentYearVal.Exists() && currentYearVal.IsArray() {
			arr := currentYearVal.Array()
			if len(arr) > 1 {
				currentYear, _ = strconv.ParseFloat(arr[1].String(), 64)
			}
		}

		// 获取去年数值（columns[1] 的索引1）
		previousYearVal := itemMap[columns[1]]
		var previousYear float64
		if previousYearVal.Exists() && previousYearVal.IsArray() {
			arr := previousYearVal.Array()
			if len(arr) > 1 {
				previousYear, _ = strconv.ParseFloat(arr[1].String(), 64)
			}
		}

		result = append(result, CarMarketCateItem{
			Month:        month,
			CurrentYear:  currentYear,
			PreviousYear: previousYear,
		})
	}

	return result, nil
}

// carMarketCateCPCALingshou 乘联会-统计数据-车型大类-零售
func carMarketCateCPCALingshou(symbol string) ([]CarMarketCateItem, error) {
	url := CPCAChartListURL
	resp, err := resty.New().R().
		SetQueryParam("charttype", "3").
		Get(url)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	json := resp.String()

	// 根据symbol选择数据数组
	dataIndex := categorySymbolMap[symbol]
	dataList := gjson.Get(json, fmt.Sprintf("%d.dataList", dataIndex)).Array()

	if len(dataList) == 0 {
		return nil, fmt.Errorf("未获取到数据")
	}

	// 获取列名
	firstItem := dataList[0].Map()
	columns := make([]string, 0, len(firstItem))
	for k := range firstItem {
		columns = append(columns, k)
	}

	var result []CarMarketCateItem

	for _, item := range dataList {
		itemMap := item.Map()

		// 获取月份
		month := itemMap["month"].String()
		if month == "" {
			month = itemMap["月份"].String()
		}

		// 获取当年数值（columns[2] 的索引2）
		currentYearVal := itemMap[columns[2]]
		var currentYear float64
		if currentYearVal.Exists() && currentYearVal.IsArray() {
			arr := currentYearVal.Array()
			if len(arr) > 2 {
				currentYear, _ = strconv.ParseFloat(arr[2].String(), 64)
			}
		}

		// 获取去年数值（columns[1] 的索引2）
		previousYearVal := itemMap[columns[1]]
		var previousYear float64
		if previousYearVal.Exists() && previousYearVal.IsArray() {
			arr := previousYearVal.Array()
			if len(arr) > 2 {
				previousYear, _ = strconv.ParseFloat(arr[2].String(), 64)
			}
		}

		result = append(result, CarMarketCateItem{
			Month:        month,
			CurrentYear:  currentYear,
			PreviousYear: previousYear,
		})
	}

	return result, nil
}

// CarMarketCateCPCA 乘联会-统计数据-车型大类
//
// 参数:
//   - symbol: 车型大类，"轿车"/"MPV"/"SUV"/"占比"
//   - indicator: 指标类型，"批发" 或 "零售"
//
// 返回:
//   - []CarMarketCateItem: 车型大类数据
//   - error: 错误信息
func CarMarketCateCPCA(symbol string, indicator string) ([]CarMarketCateItem, error) {
	if indicator == IndicatorWholesale {
		return carMarketCateCPCAPifa(symbol)
	}
	return carMarketCateCPCALingshou(symbol)
}

// CarMarketCountryCPCA 乘联会-统计数据-国别细分市场
//
// 返回:
//   - []CarMarketCountryItem: 国别细分市场数据
//   - error: 错误信息
func CarMarketCountryCPCA() ([]CarMarketCountryItem, error) {
	url := CPCAChartListURL
	resp, err := resty.New().R().
		SetQueryParam("charttype", "4").
		Get(url)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	json := resp.String()

	dataList := gjson.Get(json, "0.dataList").Array()

	if len(dataList) == 0 {
		return nil, fmt.Errorf("未获取到数据")
	}

	var result []CarMarketCountryItem

	for _, item := range dataList {
		itemMap := item.Map()

		// 第一列是月份
		month := ""
		for k, v := range itemMap {
			if k == "月份" || k == "month" {
				month = v.String()
				break
			}
		}

		var values []CarMarketCountryRow
		for k, v := range itemMap {
			if k == "月份" || k == "month" {
				continue
			}
			if v.IsArray() {
				arr := v.Array()
				if len(arr) >= 3 {
					val1, _ := strconv.ParseFloat(arr[0].String(), 64)
					val2, _ := strconv.ParseFloat(arr[1].String(), 64)
					val3, _ := strconv.ParseFloat(arr[2].String(), 64)
					val4 := 0.0
					if len(arr) > 3 {
						val4, _ = strconv.ParseFloat(arr[3].String(), 64)
					}
					values = append(values, CarMarketCountryRow{
						Country: k,
						Value1:  val1,
						Value2:  val2,
						Value3:  val3,
						Value4:  val4,
					})
				}
			}
		}

		result = append(result, CarMarketCountryItem{
			Month:  month,
			Values: values,
		})
	}

	return result, nil
}

// CarMarketSegmentCPCA 乘联会-统计数据-级别细分市场
//
// 参数:
//   - symbol: 车型大类，"轿车"/"MPV"/"SUV"
//
// 返回:
//   - []CarMarketSegmentItem: 级别细分市场数据
//   - error: 错误信息
func CarMarketSegmentCPCA(symbol string) ([]CarMarketSegmentItem, error) {
	url := CPCAChartListURL
	resp, err := resty.New().R().
		SetQueryParam("charttype", "5").
		Get(url)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	json := resp.String()

	// 根据symbol选择数据数组
	dataIndex := segmentSymbolMap[symbol]
	dataList := gjson.Get(json, fmt.Sprintf("%d.dataList", dataIndex)).Array()

	if len(dataList) == 0 {
		return nil, fmt.Errorf("未获取到数据")
	}

	var result []CarMarketSegmentItem

	for _, item := range dataList {
		itemMap := item.Map()

		// 第一列是月份
		month := ""
		for k, v := range itemMap {
			if k == "月份" || k == "month" {
				month = v.String()
				break
			}
		}

		var values []CarMarketSegmentRow
		for k, v := range itemMap {
			if k == "月份" || k == "month" {
				continue
			}
			if v.IsArray() {
				arr := v.Array()
				if len(arr) >= 2 {
					val1, _ := strconv.ParseFloat(arr[0].String(), 64)
					val2, _ := strconv.ParseFloat(arr[1].String(), 64)
					val3 := 0.0
					if len(arr) > 2 {
						val3, _ = strconv.ParseFloat(arr[2].String(), 64)
					}
					values = append(values, CarMarketSegmentRow{
						Segment: k,
						Value1:  val1,
						Value2:  val2,
						Value3:  val3,
					})
				}
			}
		}

		result = append(result, CarMarketSegmentItem{
			Month:  month,
			Values: values,
		})
	}

	return result, nil
}

// CarMarketFuelCPCA 乘联会-统计数据-新能源细分市场
//
// 参数:
//   - symbol: 燃料类型，"整体市场"/"销量占比-PHEV-BEV"/"销量占比-ICE-NEV"
//
// 返回:
//   - []CarMarketFuelItem: 新能源细分市场数据
//   - error: 错误信息
func CarMarketFuelCPCA(symbol string) ([]CarMarketFuelItem, error) {
	url := CPCAChartListURL
	resp, err := resty.New().R().
		SetQueryParam("charttype", "6").
		Get(url)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	json := resp.String()

	// 根据symbol选择数据数组
	dataIndex := fuelSymbolMap[symbol]
	dataList := gjson.Get(json, fmt.Sprintf("%d.dataList", dataIndex)).Array()

	if len(dataList) == 0 {
		return nil, fmt.Errorf("未获取到数据")
	}

	// 获取列名
	firstItem := dataList[0].Map()
	columns := make([]string, 0, len(firstItem))
	for k := range firstItem {
		columns = append(columns, k)
	}

	var result []CarMarketFuelItem

	for _, item := range dataList {
		itemMap := item.Map()

		// 获取月份
		month := itemMap["month"].String()
		if month == "" {
			month = itemMap["月份"].String()
		}

		// 获取当年数值（索引2）
		currentYearVal := itemMap[columns[1]]
		var currentYear float64
		if currentYearVal.Exists() && currentYearVal.IsArray() {
			arr := currentYearVal.Array()
			if len(arr) > 2 {
				currentYear, _ = strconv.ParseFloat(arr[2].String(), 64)
			}
		}

		// 获取去年数值（索引2）
		previousYearVal := itemMap[columns[2]]
		var previousYear float64
		if previousYearVal.Exists() && previousYearVal.IsArray() {
			arr := previousYearVal.Array()
			if len(arr) > 2 {
				previousYear, _ = strconv.ParseFloat(arr[2].String(), 64)
			}
		}

		result = append(result, CarMarketFuelItem{
			Month:        month,
			CurrentYear:  currentYear,
			PreviousYear: previousYear,
		})
	}

	return result, nil
}
