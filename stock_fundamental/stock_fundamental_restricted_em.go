package stock_fundamental

import (
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// symbolMap 市场代码映射
var symbolMap = map[string]string{
	"全部股票": "000300",
	"沪市A股": "000001",
	"科创板":  "000688",
	"深市A股": "399001",
	"创业板":  "399001",
	"京市A股": "999999",
}

// StockRestrictedReleaseSummaryEm 东方财富网-限售股解禁汇总
//
// 获取限售股解禁汇总数据，包含解禁时间、当日解禁股票家数、解禁数量、实际解禁市值等信息
//
// 参数:
//   - symbol: 标的市场，可选值: "全部股票", "沪市A股", "科创板", "深市A股", "创业板", "京市A股"
//   - startDate: 开始时间，格式: "20221101"
//   - endDate: 结束时间，格式: "20221209"
//
// 返回:
//   - []StockRestrictedReleaseSummaryEmItem: 限售股解禁汇总列表
//   - error: 错误信息
//
// 示例:
//
//	summary, err := stock_fundamental.StockRestrictedReleaseSummaryEm("全部股票", "20221101", "20221209")
//	if err != nil {
//	    log.Fatal(err)
//	}
func StockRestrictedReleaseSummaryEm(symbol, startDate, endDate string) ([]StockRestrictedReleaseSummaryEmItem, error) {
	// 获取市场代码
	indexCode, ok := symbolMap[symbol]
	if !ok {
		return nil, fmt.Errorf("无效的symbol参数: %s，可选值为: 全部股票, 沪市A股, 科创板, 深市A股, 创业板, 京市A股", symbol)
	}

	// 日期格式转换: 20221101 -> 2022-11-01
	startDateStr := formatDate(startDate)
	endDateStr := formatDate(endDate)

	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"

	params := map[string]string{
		"sortColumns":  "FREE_DATE",
		"sortTypes":    "1",
		"pageSize":     "500",
		"pageNumber":   "1",
		"columns":      "ALL",
		"quoteColumns": "f2~03~INDEX_CODE,f3~03~INDEX_CODE,f124~03~INDEX_CODE",
		"quoteType":    "0",
		"source":       "WEB",
		"client":       "WEB",
		"filter": fmt.Sprintf(`(INDEX_CODE="%s")(FREE_DATE>='%s')(FREE_DATE<='%s')`,
			indexCode, startDateStr, endDateStr),
		"reportName": "RPT_LIFTDAY_STA",
	}

	headers := map[string]string{
		"Referer": "https://data.eastmoney.com/dxf/marketStatistics.html",
	}

	resp, err := utils.GetWithHeaders(url, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求限售股解禁汇总失败: %w", err)
	}

	json := gjson.ParseBytes(resp.Body())
	if !json.Get("result").Exists() || !json.Get("result.data").Exists() {
		return []StockRestrictedReleaseSummaryEmItem{}, nil
	}

	dataArray := json.Get("result.data").Array()

	var items []StockRestrictedReleaseSummaryEmItem
	for i, data := range dataArray {
		liftDateStr := data.Get("FREE_DATE").String()
		liftDate, _ := utils.ParseDate(liftDateStr)

		item := StockRestrictedReleaseSummaryEmItem{
			Index:               i + 1,
			LiftDate:            liftDate,
			StockCount:          int(data.Get("COMPANY_NUM").Int()),
			LiftShares:          utils.MustFloat64(data.Get("LIFT_SHARES").String()) * 10000,
			ActualLiftShares:    utils.MustFloat64(data.Get("ABLE_FREE_SHARES").String()) * 10000,
			ActualLiftMarketCap: utils.MustFloat64(data.Get("ABLE_FREE_MARKET_CAP").String()) * 10000,
			HS300Index:          utils.MustFloat64(data.Get("INDEX_CLOSE").String()),
			HS300ChangeRatio:    utils.MustFloat64(data.Get("INDEX_CHANGE_RATE").String()),
		}
		items = append(items, item)
	}

	return items, nil
}

// StockRestrictedReleaseDetailEm 东方财富网-限售股解禁详情
//
// 获取限售股解禁详情数据，包含股票代码、股票简称、解禁时间、限售股类型、解禁数量等信息
//
// 参数:
//   - startDate: 开始时间，格式: "20221202"
//   - endDate: 结束时间，格式: "20241202"
//
// 返回:
//   - []StockRestrictedReleaseDetailEmItem: 限售股解禁详情列表
//   - error: 错误信息
//
// 示例:
//
//	detail, err := stock_fundamental.StockRestrictedReleaseDetailEm("20221202", "20221204")
//	if err != nil {
//	    log.Fatal(err)
//	}
func StockRestrictedReleaseDetailEm(startDate, endDate string) ([]StockRestrictedReleaseDetailEmItem, error) {
	// 日期格式转换: 20221202 -> 2022-12-02
	startDateStr := formatDate(startDate)
	endDateStr := formatDate(endDate)

	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"

	params := map[string]string{
		"sortColumns": "FREE_DATE,CURRENT_FREE_SHARES",
		"sortTypes":   "1,1",
		"pageSize":    "500",
		"pageNumber":  "1",
		"reportName":  "RPT_LIFT_STAGE",
		"columns":     "SECURITY_CODE,SECURITY_NAME_ABBR,FREE_DATE,CURRENT_FREE_SHARES,ABLE_FREE_SHARES,LIFT_MARKET_CAP,FREE_RATIO,NEW,B20_ADJCHRATE,A20_ADJCHRATE,FREE_SHARES_TYPE,TOTAL_RATIO,NON_FREE_SHARES,BATCH_HOLDER_NUM",
		"source":      "WEB",
		"client":      "WEB",
		"filter":      fmt.Sprintf(`(FREE_DATE>='%s')(FREE_DATE<='%s')`, startDateStr, endDateStr),
	}

	headers := map[string]string{
		"Referer": "https://data.eastmoney.com/dxf/detail.html",
	}

	// 首次请求获取总页数
	resp, err := utils.GetWithHeaders(url, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求限售股解禁详情失败: %w", err)
	}

	json := gjson.ParseBytes(resp.Body())
	pageNum := json.Get("result.pages").Int()
	if pageNum == 0 {
		pageNum = 1
	}

	var allItems []StockRestrictedReleaseDetailEmItem

	// 分页获取所有数据
	for page := 1; page <= int(pageNum); page++ {
		params["pageNumber"] = fmt.Sprintf("%d", page)
		resp, err := utils.GetWithHeaders(url, params, headers)
		if err != nil {
			return nil, fmt.Errorf("请求第%d页失败: %w", page, err)
		}

		json := gjson.ParseBytes(resp.Body())
		dataArray := json.Get("result.data").Array()

		for _, data := range dataArray {
			liftDateStr := data.Get("FREE_DATE").String()
			liftDate, _ := utils.ParseDate(liftDateStr)

			item := StockRestrictedReleaseDetailEmItem{
				Code:                data.Get("SECURITY_CODE").String(),
				Name:                data.Get("SECURITY_NAME_ABBR").String(),
				LiftDate:            liftDate,
				SharesType:          data.Get("FREE_SHARES_TYPE").String(),
				LiftShares:          utils.MustFloat64(data.Get("CURRENT_FREE_SHARES").String()) * 10000,
				ActualLiftShares:    utils.MustFloat64(data.Get("ABLE_FREE_SHARES").String()) * 10000,
				ActualLiftMarketCap: utils.MustFloat64(data.Get("LIFT_MARKET_CAP").String()) * 10000,
				CirculationRatio:    utils.MustFloat64(data.Get("FREE_RATIO").String()),
				PrevClosePrice:      utils.MustFloat64(data.Get("NEW").String()),
				Before20ChangeRatio: utils.MustFloat64(data.Get("B20_ADJCHRATE").String()),
				After20ChangeRatio:  utils.MustFloat64(data.Get("A20_ADJCHRATE").String()),
			}
			allItems = append(allItems, item)
		}
	}

	// 设置序号
	for i := range allItems {
		allItems[i].Index = i + 1
	}

	return allItems, nil
}

// StockRestrictedReleaseQueueEm 东方财富网-个股限售解禁批次
//
// 获取个股限售解禁批次数据，包含解禁时间、解禁股东数、解禁数量、实际解禁数量市值等信息
//
// 参数:
//   - symbol: 股票代码，如 "600000"
//
// 返回:
//   - []StockRestrictedReleaseQueueEmItem: 个股限售解禁批次列表
//   - error: 错误信息
//
// 示例:
//
//	queue, err := stock_fundamental.StockRestrictedReleaseQueueEm("600000")
//	if err != nil {
//	    log.Fatal(err)
//	}
func StockRestrictedReleaseQueueEm(symbol string) ([]StockRestrictedReleaseQueueEmItem, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"

	params := map[string]string{
		"sortColumns": "FREE_DATE",
		"sortTypes":   "-1",
		"pageSize":    "500",
		"pageNumber":  "1",
		"reportName":  "RPT_LIFT_STAGE",
		"filter":      fmt.Sprintf(`(SECURITY_CODE="%s")`, symbol),
		"columns":     "SECURITY_CODE,SECURITY_NAME_ABBR,FREE_DATE,CURRENT_FREE_SHARES,ABLE_FREE_SHARES,LIFT_MARKET_CAP,FREE_RATIO,NEW,B20_ADJCHRATE,A20_ADJCHRATE,FREE_SHARES_TYPE,TOTAL_RATIO,NON_FREE_SHARES,BATCH_HOLDER_NUM",
		"source":      "WEB",
		"client":      "WEB",
	}

	headers := map[string]string{
		"Referer": "https://data.eastmoney.com/dxf/q/" + symbol + ".html",
	}

	resp, err := utils.GetWithHeaders(url, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求个股限售解禁批次失败: %w", err)
	}

	json := gjson.ParseBytes(resp.Body())

	// 检查result是否存在
	if !json.Get("result").Exists() {
		return []StockRestrictedReleaseQueueEmItem{}, nil
	}

	if !json.Get("result.data").Exists() {
		return []StockRestrictedReleaseQueueEmItem{}, nil
	}

	dataArray := json.Get("result.data").Array()

	var items []StockRestrictedReleaseQueueEmItem
	for i, data := range dataArray {
		liftDateStr := data.Get("FREE_DATE").String()
		liftDate, _ := utils.ParseDate(liftDateStr)

		item := StockRestrictedReleaseQueueEmItem{
			Index:               i + 1,
			LiftDate:            liftDate,
			HolderNum:           int(data.Get("BATCH_HOLDER_NUM").Int()),
			LiftShares:          utils.MustFloat64(data.Get("CURRENT_FREE_SHARES").String()) * 10000,
			ActualLiftShares:    utils.MustFloat64(data.Get("ABLE_FREE_SHARES").String()) * 10000,
			NonLiftShares:       utils.MustFloat64(data.Get("NON_FREE_SHARES").String()) * 10000,
			ActualLiftMarketCap: utils.MustFloat64(data.Get("LIFT_MARKET_CAP").String()) * 10000,
			TotalRatio:          utils.MustFloat64(data.Get("TOTAL_RATIO").String()),
			CirculationRatio:    utils.MustFloat64(data.Get("FREE_RATIO").String()),
			PrevClosePrice:      utils.MustFloat64(data.Get("NEW").String()),
			SharesType:          data.Get("FREE_SHARES_TYPE").String(),
			Before20ChangeRatio: utils.MustFloat64(data.Get("B20_ADJCHRATE").String()),
			After20ChangeRatio:  utils.MustFloat64(data.Get("A20_ADJCHRATE").String()),
		}
		items = append(items, item)
	}

	return items, nil
}

// formatDate 将日期格式从 20221101 转换为 2022-11-01
func formatDate(dateStr string) string {
	if len(dateStr) != 8 {
		return dateStr
	}
	return dateStr[:4] + "-" + dateStr[4:6] + "-" + dateStr[6:]
}

// formatSymbol 验证symbol参数是否有效
func formatSymbol(symbol string) bool {
	_, ok := symbolMap[symbol]
	return ok
}

// GetSupportedSymbolList 获取支持的symbol列表
func GetSupportedSymbolList() []string {
	keys := make([]string, 0, len(symbolMap))
	for k := range symbolMap {
		keys = append(keys, k)
	}
	return keys
}
