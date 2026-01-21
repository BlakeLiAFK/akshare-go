package bond

import (
	"fmt"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
	"github.com/tidwall/gjson"
)

// bondChinaCloseReturnMap 收盘收益率曲线映射数据（内部辅助函数）
//
// 从中国外汇交易中心获取收益率曲线类型映射
// 这个函数会尝试注册服务以获取访问权限
//
// 返回:
//   - gjson.Result: 包含收益率曲线类型映射的JSON数据
//   - error: 错误信息
func bondChinaCloseReturnMap() (gjson.Result, error) {
	headers := map[string]string{
		"Accept":           "application/json, text/javascript, */*; q=0.01",
		"Accept-Encoding":  "gzip, deflate, br",
		"Accept-Language":  "zh-CN,zh;q=0.9,en;q=0.8",
		"Cache-Control":    "no-cache",
		"Connection":       "keep-alive",
		"Content-Length":   "0",
		"Host":             "www.chinamoney.com.cn",
		"Origin":           "https://www.chinamoney.com.cn",
		"Pragma":           "no-cache",
		"Referer":          "https://www.chinamoney.com.cn/chinese/bkcurvclosedyhis/?bondType=CYCC000&reference=1",
		"User-Agent":       "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
		"X-Requested-With": "XMLHttpRequest",
	}

	url := "https://www.chinamoney.com.cn/ags/ms/cm-u-bk-currency/ClsYldCurvCurvGO"
	resp, err := utils.GetWithHeaders(url, nil, headers)
	if err != nil {
		return gjson.Result{}, fmt.Errorf("请求失败: %w", err)
	}

	dataJSON := gjson.Parse(resp.String())
	if !dataJSON.Get("records").Exists() {
		return gjson.Result{}, fmt.Errorf("未获取到映射数据")
	}

	return dataJSON, nil
}

// BondChinaCloseReturn 收盘收益率曲线历史数据
//
// 获取中国外汇交易中心的收盘收益率曲线历史数据
//
// 参数:
//   - symbol: 债券类型，如 "国债", "同业存单(AAA)" 等
//   - period: 期限间隔，可选值: "0.1", "0.5", "1"
//   - startDate: 开始日期，格式: "20231101"
//   - endDate: 结束日期，格式: "20231101"（结束日期和开始日期不要超过1个月）
//
// 返回:
//   - dataframe.DataFrame: 包含收益率曲线历史数据
//   - error: 错误信息
//
// 数据源: https://www.chinamoney.com.cn/chinese/bkcurvclosedyhis/?bondType=CYCC000&reference=1
func BondChinaCloseReturn(symbol, period, startDate, endDate string) (dataframe.DataFrame, error) {
	if len(startDate) != 8 || len(endDate) != 8 {
		return dataframe.DataFrame{}, fmt.Errorf("日期格式错误，应为YYYYMMDD")
	}

	// 获取映射数据
	mapData, err := bondChinaCloseReturnMap()
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("获取映射数据失败: %w", err)
	}

	// 从映射中查找symbol对应的code
	var symbolCode string
	mapData.Get("records").ForEach(func(_, item gjson.Result) bool {
		if item.Get("cnLabel").String() == symbol {
			symbolCode = item.Get("value").String()
			return false // 找到后停止遍历
		}
		return true
	})

	if symbolCode == "" {
		return dataframe.DataFrame{}, fmt.Errorf("未找到债券类型: %s", symbol)
	}

	url := "https://www.chinamoney.com.cn/ags/ms/cm-u-bk-currency/ClsYldCurvHis"
	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
	}

	params := map[string]string{
		"lang":      "CN",
		"reference": "1,2,3",
		"bondType":  symbolCode,
		"startDate": fmt.Sprintf("%s-%s-%s", startDate[:4], startDate[4:6], startDate[6:]),
		"endDate":   fmt.Sprintf("%s-%s-%s", endDate[:4], endDate[4:6], endDate[6:]),
		"termId":    period,
		"pageNum":   "1",
		"pageSize":  "50",
	}

	resp, err := utils.GetWithHeaders(url, params, headers)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	dataJSON := gjson.Parse(resp.String())
	records := dataJSON.Get("records")
	if !records.Exists() || !records.IsArray() {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	// 构建DataFrame记录
	var dfRecords []map[string]interface{}

	records.ForEach(func(_, item gjson.Result) bool {
		// 跳过newDateValue字段
		record := map[string]interface{}{
			"日期":    item.Get("showDateCN").String(),
			"期限":    utils.MustParseFloat(item.Get("term").String()),
			"到期收益率": utils.MustParseFloat(item.Get("ytmCurve").String()),
			"即期收益率": utils.MustParseFloat(item.Get("spotCurve").String()),
			"远期收益率": utils.MustParseFloat(item.Get("forwardCurve").String()),
		}
		dfRecords = append(dfRecords, record)
		return true
	})

	if len(dfRecords) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到有效数据")
	}

	df := dataframe.LoadMaps(dfRecords)
	return df, nil
}

// MacroChinaSwapRate FR007利率互换曲线历史数据
//
// 获取FR007利率互换曲线历史数据（只能获取近一年的数据）
//
// 参数:
//   - startDate: 开始日期，格式: "20231101"
//   - endDate: 结束日期，格式: "20231204"（开始和结束日期不得超过一个月）
//
// 返回:
//   - dataframe.DataFrame: 包含FR007利率互换曲线历史数据
//   - error: 错误信息
//
// 数据源: https://www.chinamoney.com.cn/chinese/bkcurvfxhis/?cfgItemType=72&curveType=FR007
func MacroChinaSwapRate(startDate, endDate string) (dataframe.DataFrame, error) {
	if len(startDate) != 8 || len(endDate) != 8 {
		return dataframe.DataFrame{}, fmt.Errorf("日期格式错误，应为YYYYMMDD")
	}

	// 调用映射函数以注册服务
	_, _ = bondChinaCloseReturnMap()

	url := "https://www.chinamoney.com.cn/ags/ms/cm-u-bk-shibor/IfccHis"
	params := map[string]string{
		"cfgItemType":      "72",
		"interestRateType": "0",
		"startDate":        fmt.Sprintf("%s-%s-%s", startDate[:4], startDate[4:6], startDate[6:]),
		"endDate":          fmt.Sprintf("%s-%s-%s", endDate[:4], endDate[4:6], endDate[6:]),
		"bidAskType":       "",
		"lang":             "CN",
		"quoteTime":        "全部",
		"pageSize":         "5000",
		"pageNum":          "1",
	}

	headers := map[string]string{
		"Accept":           "application/json, text/javascript, */*; q=0.01",
		"Accept-Encoding":  "gzip, deflate, br",
		"Accept-Language":  "zh-CN,zh;q=0.9,en;q=0.8",
		"Cache-Control":    "no-cache",
		"Connection":       "keep-alive",
		"Content-Length":   "0",
		"Host":             "www.chinamoney.com.cn",
		"Origin":           "https://www.chinamoney.com.cn",
		"Pragma":           "no-cache",
		"Referer":          "https://www.chinamoney.com.cn/chinese/bkcurvfxhis/?cfgItemType=72&curveType=FR007",
		"Sec-Fetch-Dest":   "empty",
		"Sec-Fetch-Mode":   "cors",
		"Sec-Fetch-Site":   "same-origin",
		"User-Agent":       "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36",
		"X-Requested-With": "XMLHttpRequest",
	}

	resp, err := utils.PostFormWithHeaders(url, params, headers)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	dataJSON := gjson.Parse(resp.String())
	records := dataJSON.Get("records")
	if !records.Exists() || !records.IsArray() {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	// 构建DataFrame记录
	var dfRecords []map[string]interface{}

	records.ForEach(func(_, item gjson.Result) bool {
		itemArray := item.Array()
		if len(itemArray) < 17 {
			return true
		}

		// 根据Python代码的列映射
		date := itemArray[0].String()
		curveName := itemArray[11].String()
		time := itemArray[3].String()
		priceType := itemArray[9].String()

		// data字段在索引16
		dataArray := itemArray[16].Array()
		if len(dataArray) < 11 {
			return true
		}

		record := map[string]interface{}{
			"日期":   date,
			"曲线名称": curveName,
			"时刻":   time,
			"价格类型": priceType,
			"1M":   utils.MustParseFloat(dataArray[0].String()),
			"3M":   utils.MustParseFloat(dataArray[1].String()),
			"6M":   utils.MustParseFloat(dataArray[2].String()),
			"9M":   utils.MustParseFloat(dataArray[3].String()),
			"1Y":   utils.MustParseFloat(dataArray[4].String()),
			"2Y":   utils.MustParseFloat(dataArray[5].String()),
			"3Y":   utils.MustParseFloat(dataArray[6].String()),
			"4Y":   utils.MustParseFloat(dataArray[7].String()),
			"5Y":   utils.MustParseFloat(dataArray[8].String()),
			"7Y":   utils.MustParseFloat(dataArray[9].String()),
			"10Y":  utils.MustParseFloat(dataArray[10].String()),
		}
		dfRecords = append(dfRecords, record)
		return true
	})

	if len(dfRecords) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到有效数据")
	}

	df := dataframe.LoadMaps(dfRecords)

	// 按日期排序
	df = df.Arrange(dataframe.Sort("日期"))

	return df, nil
}

// MacroChinaBondPublic 中国-债券信息披露-债券发行
//
// 获取中国外汇交易中心的债券发行信息
//
// 参数:
//   - 无
//
// 返回:
//   - dataframe.DataFrame: 包含债券发行信息
//   - error: 错误信息
//
// 数据源: https://www.chinamoney.com.cn/chinese/xzjfx/
func MacroChinaBondPublic() (dataframe.DataFrame, error) {
	// 调用映射函数以注册服务
	_, _ = bondChinaCloseReturnMap()

	url := "https://www.chinamoney.com.cn/ags/ms/cm-u-bond-an/bnBondEmit"
	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36",
	}

	// 先获取总页数
	payload := map[string]string{
		"enty":            "",
		"bondType":        "",
		"bondNameCode":    "",
		"leadUnderwriter": "",
		"pageNo":          "1",
		"pageSize":        "10",
		"limit":           "1",
	}

	resp, err := utils.PostFormWithHeaders(url, payload, headers)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	dataJSON := gjson.Parse(resp.String())
	totalPages := int(dataJSON.Get("data.pageTotalSize").Int()) + 1

	// 获取所有页面的数据
	var allRecords []map[string]interface{}

	for page := 1; page < totalPages; page++ {
		payload["pageNo"] = fmt.Sprintf("%d", page)

		resp, err := utils.PostFormWithHeaders(url, payload, headers)
		if err != nil {
			continue
		}

		dataJSON := gjson.Parse(resp.String())
		records := dataJSON.Get("records")
		if !records.Exists() || !records.IsArray() {
			continue
		}

		records.ForEach(func(_, item gjson.Result) bool {
			itemArray := item.Array()
			if len(itemArray) < 13 {
				return true
			}

			// 根据Python代码的列映射:
			// 债券全称:0, 债券类型:1, 发行日期:3, 计息方式:5, 债券期限:7, 债券评级:9, 价格:11, 计划发行量:12
			record := map[string]interface{}{
				"债券全称":  itemArray[0].String(),
				"债券类型":  itemArray[1].String(),
				"发行日期":  itemArray[3].String(),
				"计息方式":  itemArray[5].String(),
				"价格":    utils.MustParseFloat(itemArray[11].String()),
				"债券期限":  itemArray[7].String(),
				"计划发行量": utils.MustParseFloat(itemArray[12].String()),
				"债券评级":  itemArray[9].String(),
			}
			allRecords = append(allRecords, record)
			return true
		})

		// 添加小延迟避免请求过快
		time.Sleep(100 * time.Millisecond)
	}

	if len(allRecords) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到有效数据")
	}

	df := dataframe.LoadMaps(allRecords)
	return df, nil
}
