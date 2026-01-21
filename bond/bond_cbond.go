package bond

import (
	"fmt"
	"strconv"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
	"github.com/tidwall/gjson"
)

var (
	indicatorMap = map[string]string{
		"全价":          "QJZS",
		"净价":          "JJZS",
		"财富":          "CFZS",
		"平均市值法久期":     "PJSZFJQ",
		"平均现金流法久期":    "PJXJLFJQ",
		"平均市值法凸性":     "PJSZFTX",
		"平均现金流法凸性":    "PJXJLFTX",
		"平均现金流法到期收益率": "PJDQSYL",
		"平均市值法到期收益率":  "PJSZFDQSYL",
		"平均基点价值":      "PJJDJZ",
		"平均待偿期":       "PJDCQ",
		"平均派息率":       "PJPXL",
		"指数上日总市值":     "ZSZSZ",
		"财富指数涨跌幅":     "CFZSZDF",
		"全价指数涨跌幅":     "QJZSZDF",
		"净价指数涨跌幅":     "JJZSZDF",
		"现券结算量":       "XQJSL",
	}

	periodMap = map[string]string{
		"总值":     "00",
		"1年以下":   "01",
		"1-3年":   "02",
		"3-5年":   "03",
		"5-7年":   "04",
		"7-10年":  "05",
		"10年以上":  "06",
		"0-3个月":  "07",
		"3-6个月":  "08",
		"6-9个月":  "09",
		"9-12个月": "10",
		"0-6个月":  "11",
		"6-12个月": "12",
	}
)

// BondNewCompositeIndexCBond 中债新综合指数
//
// 获取中国债券信息网的中债新综合指数数据
//
// 参数:
//   - indicator: 指标类型，可选值见 indicatorMap
//   - period: 期限类型，可选值见 periodMap
//
// 返回:
//   - dataframe.DataFrame: 包含指数历史数据
//   - error: 错误信息
//
// 数据源: https://yield.chinabond.com.cn/cbweb-mn/indices/single_index_query
func BondNewCompositeIndexCBond(indicator, period string) (dataframe.DataFrame, error) {
	indicatorCode, ok1 := indicatorMap[indicator]
	if !ok1 {
		return dataframe.DataFrame{}, fmt.Errorf("无效的指标类型: %s", indicator)
	}

	periodCode, ok2 := periodMap[period]
	if !ok2 {
		return dataframe.DataFrame{}, fmt.Errorf("无效的期限类型: %s", period)
	}

	url := "https://yield.chinabond.com.cn/cbweb-mn/indices/singleIndexQuery"
	params := map[string]string{
		"indexid": "8a8b2ca0332abed20134ea76d8885831",
		"qxlxt":   periodCode,
		"ltcslx":  "",
		"zslxt":   indicatorCode,
		"lx":      "1",
		"locale":  "",
	}

	resp, err := utils.PostForm(url, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	// 解析JSON响应 - 数据是对象格式，key为时间戳，value为指数值
	dataKey := fmt.Sprintf("%s_%s", indicatorCode, periodCode)
	dataObj := gjson.Get(resp.String(), dataKey)
	if !dataObj.Exists() {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	// 构建DataFrame记录
	var records []map[string]interface{}

	dataObj.ForEach(func(key, value gjson.Result) bool {
		// key 是毫秒时间戳
		timestamp, _ := strconv.ParseInt(key.String(), 10, 64)
		date := time.UnixMilli(timestamp).Format("2006-01-02")

		record := map[string]interface{}{
			"date":  date,
			"value": utils.MustParseFloat(value.String()),
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

// BondCompositeIndexCBond 中债综合指数
//
// 获取中国债券信息网的中债综合指数数据
//
// 参数:
//   - indicator: 指标类型，可选值见 indicatorMap
//   - period: 期限类型，可选值见 periodMap
//
// 返回:
//   - dataframe.DataFrame: 包含指数历史数据
//   - error: 错误信息
//
// 数据源: https://yield.chinabond.com.cn/cbweb-mn/indices/single_index_query
func BondCompositeIndexCBond(indicator, period string) (dataframe.DataFrame, error) {
	indicatorCode, ok1 := indicatorMap[indicator]
	if !ok1 {
		return dataframe.DataFrame{}, fmt.Errorf("无效的指标类型: %s", indicator)
	}

	periodCode, ok2 := periodMap[period]
	if !ok2 {
		return dataframe.DataFrame{}, fmt.Errorf("无效的期限类型: %s", period)
	}

	url := "https://yield.chinabond.com.cn/cbweb-mn/indices/singleIndexQuery"
	params := map[string]string{
		"indexid": "2c90818811afed8d0111c0c672b31578",
		"qxlxt":   periodCode,
		"zslxt":   indicatorCode,
		"lx":      "1",
		"locale":  "",
	}

	resp, err := utils.PostForm(url, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	// 解析JSON响应
	dataKey := fmt.Sprintf("%s_%s", indicatorCode, periodCode)
	dataObj := gjson.Get(resp.String(), dataKey)
	if !dataObj.Exists() {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	// 构建DataFrame记录
	var records []map[string]interface{}

	dataObj.ForEach(func(key, value gjson.Result) bool {
		// key 是毫秒时间戳
		timestamp, _ := strconv.ParseInt(key.String(), 10, 64)
		date := time.UnixMilli(timestamp).Format("2006-01-02")

		record := map[string]interface{}{
			"date":  date,
			"value": utils.MustParseFloat(value.String()),
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
