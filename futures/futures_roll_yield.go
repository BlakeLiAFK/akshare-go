package futures

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

// FuturesRollYield 展期收益率
type FuturesRollYield struct {
	Variety   string  `json:"variety"`    // 品种
	RollYield float64 `json:"roll_yield"` // 展期收益率
	NearBy    string  `json:"near_by"`    // 近月合约
	Deferred  string  `json:"deferred"`   // 远月合约
	Date      string  `json:"date"`       // 日期
}

// GetRollYield 指定交易日指定品种（主力和次主力）或任意两个合约的展期收益率
//
// 参数:
//   - date: 日期，格式 "YYYYMMDD"
//   - variety: 合约品种，如 "RB", "AL"
//   - symbol1: 合约1（可选），如 "rb1810"
//   - symbol2: 合约2（可选），如 "rb1812"
//   - data: 日线数据（可选）
//
// 返回:
//   - *FuturesRollYield: 展期收益率数据
//   - error: 错误信息
func GetRollYield(date, variety, symbol1, symbol2 string, data []FuturesDailyBar) (*FuturesRollYield, error) {
	if symbol1 != "" {
		variety = SymbolVarieties(symbol1)
	}

	// 如果没有传入数据，则获取数据
	if len(data) == 0 {
		market := SymbolMarket(variety)
		var err error
		data, err = GetFuturesDaily(date, date, market)
		if err != nil {
			return nil, fmt.Errorf("获取日线数据失败: %w", err)
		}
	}

	// 过滤品种数据
	var varietyData []FuturesDailyBar
	for _, bar := range data {
		if !strings.Contains(bar.Symbol, "efp") && bar.Variety == variety {
			varietyData = append(varietyData, bar)
		}
	}

	// 按持仓量排序
	sort.Slice(varietyData, func(i, j int) bool {
		return varietyData[i].OpenInterest > varietyData[j].OpenInterest
	})

	if len(varietyData) < 2 {
		return nil, fmt.Errorf("品种 %s 合约数量不足", variety)
	}

	if symbol1 == "" {
		symbol1 = varietyData[0].Symbol
	}
	if symbol2 == "" {
		symbol2 = varietyData[1].Symbol
	}

	// 获取收盘价
	var close1, close2 float64
	for _, bar := range varietyData {
		if bar.Symbol == symbol1 {
			close1 = bar.Close
		}
		if bar.Symbol == symbol2 {
			close2 = bar.Close
		}
	}

	if close1 == 0 || close2 == 0 {
		return nil, fmt.Errorf("无法获取合约收盘价")
	}

	// 计算月份差
	re := regexp.MustCompile(`\d+`)
	a := re.FindString(symbol1)
	b := re.FindString(symbol2)

	if len(a) < 4 || len(b) < 4 {
		return nil, fmt.Errorf("合约代码格式错误")
	}

	a1, _ := strconv.Atoi(a[:len(a)-2])
	a2, _ := strconv.Atoi(a[len(a)-2:])
	b1, _ := strconv.Atoi(b[:len(b)-2])
	b2, _ := strconv.Atoi(b[len(b)-2:])

	monthDiff := (a1-b1)*12 + (a2 - b2)

	if monthDiff == 0 {
		return nil, fmt.Errorf("合约月份相同")
	}

	rollYield := math.Log(close2/close1) / float64(monthDiff) * 12

	result := &FuturesRollYield{
		Variety:   variety,
		RollYield: rollYield,
		Date:      date,
	}

	if monthDiff > 0 {
		result.NearBy = symbol2
		result.Deferred = symbol1
	} else {
		result.NearBy = symbol1
		result.Deferred = symbol2
	}

	return result, nil
}

// GetRollYieldBar 展期收益率
//
// 参数:
//   - typeMethod: 类型，可选:
//     "symbol": 获取指定交易日指定品种所有交割月合约的收盘价
//     "var": 获取指定交易日所有品种两个主力合约的展期收益率(展期收益率横截面)
//     "date": 获取指定品种每天的两个主力合约的展期收益率(展期收益率时间序列)
//   - variety: 合约品种，如 "RB"
//   - date: 指定交易日，格式 "YYYYMMDD"
//   - startDay: 开始日期，格式 "YYYYMMDD"
//   - endDay: 结束日期，格式 "YYYYMMDD"
//
// 返回:
//   - []FuturesRollYield: 展期收益率数据
//   - error: 错误信息
func GetRollYieldBar(typeMethod, variety, date, startDay, endDay string) ([]FuturesRollYield, error) {
	switch typeMethod {
	case "symbol":
		// 获取指定品种所有合约的日线数据
		market := SymbolMarket(variety)
		data, err := GetFuturesDaily(date, date, market)
		if err != nil {
			return nil, err
		}

		var result []FuturesRollYield
		for _, bar := range data {
			if bar.Variety == variety {
				result = append(result, FuturesRollYield{
					Variety:   variety,
					NearBy:    bar.Symbol,
					RollYield: bar.Close,
					Date:      date,
				})
			}
		}
		return result, nil

	case "var":
		// 获取所有品种的展期收益率
		var allData []FuturesDailyBar
		for _, market := range []string{"dce", "cffex", "shfe", "czce", "gfex"} {
			data, err := GetFuturesDaily(date, date, market)
			if err != nil {
				continue
			}
			allData = append(allData, data...)
		}

		// 获取所有品种
		varietySet := make(map[string]bool)
		for _, bar := range allData {
			if bar.Variety != "" && bar.Variety != "IO" && bar.Variety != "MO" && bar.Variety != "HO" {
				varietySet[bar.Variety] = true
			}
		}

		var result []FuturesRollYield
		for v := range varietySet {
			ry, err := GetRollYield(date, v, "", "", allData)
			if err == nil && ry != nil {
				result = append(result, *ry)
			}
		}

		// 按展期收益率排序
		sort.Slice(result, func(i, j int) bool {
			return result[i].RollYield < result[j].RollYield
		})

		return result, nil

	case "date":
		// 获取指定品种的时间序列展期收益率
		start, err := time.Parse("20060102", startDay)
		if err != nil {
			return nil, fmt.Errorf("开始日期格式错误: %w", err)
		}

		end, err := time.Parse("20060102", endDay)
		if err != nil {
			return nil, fmt.Errorf("结束日期格式错误: %w", err)
		}

		var result []FuturesRollYield
		for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
			dateStr := d.Format("20060102")
			ry, err := GetRollYield(dateStr, variety, "", "", nil)
			if err == nil && ry != nil {
				result = append(result, *ry)
			}
		}

		return result, nil

	default:
		return nil, fmt.Errorf("无效的类型: %s", typeMethod)
	}
}
