package bond

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
	"github.com/tidwall/gjson"
)

// BondSHBuyBackEM 上证质押式回购
//
// 获取东方财富网上证质押式回购实时行情数据
//
// 参数:
//   - 无
//
// 返回:
//   - dataframe.DataFrame: 包含上证质押式回购行情数据
//   - error: 错误信息
//
// 数据源: https://quote.eastmoney.com/center/gridlist.html#bond_sh_buyback
func BondSHBuyBackEM() (dataframe.DataFrame, error) {
	url := "https://push2.eastmoney.com/api/qt/clist/get"
	params := map[string]string{
		"np":     "1",
		"fltt":   "1",
		"invt":   "2",
		"fs":     "m:1+b:MK0356",
		"fields": "f12,f13,f14,f1,f2,f4,f3,f152,f17,f18,f15,f16,f5,f6",
		"fid":    "f6",
		"pn":     "1",
		"pz":     "20",
		"po":     "1",
		"dect":   "1",
		"wbp2u":  "|0|0|0|web",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	// 解析JSON响应
	diff := gjson.Get(resp.String(), "data.diff")
	if !diff.Exists() || !diff.IsArray() {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	// 构建DataFrame记录
	var records []map[string]interface{}

	diff.ForEach(func(idx, item gjson.Result) bool {
		record := map[string]interface{}{
			"序号":  int(idx.Int()) + 1,
			"代码":  item.Get("f12").String(),
			"名称":  item.Get("f14").String(),
			"最新价": parseAndDivide(item.Get("f2").String(), 1000),
			"涨跌额": parseAndDivide(item.Get("f4").String(), 1000),
			"涨跌幅": parseAndDivide(item.Get("f3").String(), 100),
			"今开":  parseAndDivide(item.Get("f17").String(), 1000),
			"最高":  parseAndDivide(item.Get("f15").String(), 1000),
			"最低":  parseAndDivide(item.Get("f16").String(), 1000),
			"昨收":  parseAndDivide(item.Get("f18").String(), 1000),
			"成交量": utils.MustParseFloat(item.Get("f5").String()),
			"成交额": utils.MustParseFloat(item.Get("f6").String()),
		}
		records = append(records, record)
		return true
	})

	if len(records) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到有效数据")
	}

	df := dataframe.LoadMaps(records)
	return df, nil
}

// BondSZBuyBackEM 深证质押式回购
//
// 获取东方财富网深证质押式回购实时行情数据
//
// 参数:
//   - 无
//
// 返回:
//   - dataframe.DataFrame: 包含深证质押式回购行情数据
//   - error: 错误信息
//
// 数据源: https://quote.eastmoney.com/center/gridlist.html#bond_sz_buyback
func BondSZBuyBackEM() (dataframe.DataFrame, error) {
	url := "https://push2.eastmoney.com/api/qt/clist/get"
	params := map[string]string{
		"np":     "1",
		"fltt":   "1",
		"invt":   "2",
		"fs":     "m:0+b:MK0356",
		"fields": "f12,f13,f14,f1,f2,f4,f3,f152,f17,f18,f15,f16,f5,f6",
		"fid":    "f6",
		"pn":     "1",
		"pz":     "20",
		"po":     "1",
		"dect":   "1",
		"wbp2u":  "|0|0|0|web",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	// 解析JSON响应
	diff := gjson.Get(resp.String(), "data.diff")
	if !diff.Exists() || !diff.IsArray() {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	// 构建DataFrame记录
	var records []map[string]interface{}

	diff.ForEach(func(idx, item gjson.Result) bool {
		record := map[string]interface{}{
			"序号":  int(idx.Int()) + 1,
			"代码":  item.Get("f12").String(),
			"名称":  item.Get("f14").String(),
			"最新价": parseAndDivide(item.Get("f2").String(), 1000),
			"涨跌额": parseAndDivide(item.Get("f4").String(), 1000),
			"涨跌幅": parseAndDivide(item.Get("f3").String(), 100),
			"今开":  parseAndDivide(item.Get("f17").String(), 1000),
			"最高":  parseAndDivide(item.Get("f15").String(), 1000),
			"最低":  parseAndDivide(item.Get("f16").String(), 1000),
			"昨收":  parseAndDivide(item.Get("f18").String(), 1000),
			"成交量": utils.MustParseFloat(item.Get("f5").String()),
			"成交额": utils.MustParseFloat(item.Get("f6").String()),
		}
		records = append(records, record)
		return true
	})

	if len(records) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到有效数据")
	}

	df := dataframe.LoadMaps(records)
	return df, nil
}

// BondBuyBackHistEM 质押式回购历史数据
//
// 获取东方财富网质押式回购历史K线数据
//
// 参数:
//   - symbol: 质押式回购代码，如 "204001"(上证) 或 "131810"(深证)
//
// 返回:
//   - dataframe.DataFrame: 包含历史K线数据
//   - error: 错误信息
//
// 数据源: https://quote.eastmoney.com/center/gridlist.html#bond_sh_buyback
func BondBuyBackHistEM(symbol string) (dataframe.DataFrame, error) {
	if symbol == "" {
		return dataframe.DataFrame{}, fmt.Errorf("代码不能为空")
	}

	// 根据代码首位判断市场
	marketID := "1"
	if strings.HasPrefix(symbol, "1") {
		marketID = "0"
	}

	url := "https://push2his.eastmoney.com/api/qt/stock/kline/get"
	params := map[string]string{
		"secid":   fmt.Sprintf("%s.%s", marketID, symbol),
		"klt":     "101",
		"fqt":     "1",
		"lmt":     "10000",
		"end":     "20500000",
		"iscca":   "1",
		"fields1": "f1,f2,f3,f4,f5,f6,f7,f8",
		"fields2": "f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61,f62,f63,f64",
		"forcect": "1",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	// 解析JSON响应
	klines := gjson.Get(resp.String(), "data.klines")
	if !klines.Exists() || !klines.IsArray() {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到K线数据")
	}

	// 构建DataFrame记录
	var records []map[string]interface{}

	klines.ForEach(func(_, item gjson.Result) bool {
		parts := strings.Split(item.String(), ",")
		if len(parts) < 7 {
			return true
		}

		record := map[string]interface{}{
			"日期":  parts[0],
			"开盘":  utils.MustParseFloat(parts[1]),
			"收盘":  utils.MustParseFloat(parts[2]),
			"最高":  utils.MustParseFloat(parts[3]),
			"最低":  utils.MustParseFloat(parts[4]),
			"成交量": utils.MustParseFloat(parts[5]),
			"成交额": utils.MustParseFloat(parts[6]),
		}
		records = append(records, record)
		return true
	})

	if len(records) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到有效数据")
	}

	df := dataframe.LoadMaps(records)
	return df, nil
}

// parseAndDivide 解析字符串并除以指定的除数
func parseAndDivide(s string, divisor float64) float64 {
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return val / divisor
}
