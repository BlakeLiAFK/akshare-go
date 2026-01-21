package stock

import (
	"fmt"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

const (
	// 东方财富个股资金流向 API
	emFundFlowURL = "https://push2.eastmoney.com/api/qt/stock/fflow/kline/get"
	// 东方财富资金流向排行 API
	emFundFlowRankURL = "https://push2.eastmoney.com/api/qt/clist/get"
	// 东方财富大盘资金流向 API
	emMarketFundFlowURL = "https://push2his.eastmoney.com/api/qt/stock/fflow/daykline/get"
)

// StockIndividualFundFlowEm 获取个股资金流向（东方财富数据源）
//
// 参数:
//   - code: 股票代码，如 "000001"
//
// 返回:
//   - []FundFlow: 资金流向列表
//   - error: 错误信息
//
// 示例:
//
//	flows, err := stock.StockIndividualFundFlowEm("000001")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, f := range flows {
//	    fmt.Printf("%s: 主力净流入 %.2f 万\n", f.Date.Format("2006-01-02"), f.MainNet/10000)
//	}
func StockIndividualFundFlowEm(code string) ([]FundFlow, error) {
	if code == "" {
		return nil, fmt.Errorf("股票代码不能为空")
	}

	secid := getSecID(code)

	params := map[string]string{
		"secid":  secid,
		"lmt":    "0",
		"klt":    "101", // 日线
		"fields": "f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61,f62,f63,f64,f65",
	}

	headers := map[string]string{
		"Referer": emQuoteReferer,
	}

	resp, err := utils.GetWithHeaders(emFundFlowURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("获取个股资金流向失败: %w", err)
	}

	// 解析 JSON 数据
	text := resp.String()
	dataPath := "data.klines"
	result := gjson.Get(text, dataPath)
	if !result.Exists() {
		return nil, fmt.Errorf("解析资金流向数据失败: 未找到 %s", dataPath)
	}

	// 获取股票名称
	stockName := gjson.Get(text, "data.name").String()

	flows := make([]FundFlow, 0)
	result.ForEach(func(_, value gjson.Result) bool {
		// 数据格式: 日期,主力净流入,小单净流入,中单净流入,大单净流入,超大单净流入
		parts := splitString(value.String(), ",")
		if len(parts) < 6 {
			return true
		}

		date, _ := time.Parse("2006-01-02", parts[0])
		mainNet := utils.MustFloat64(parts[1])
		smallNet := utils.MustFloat64(parts[2])
		midNet := utils.MustFloat64(parts[3])
		bigNet := utils.MustFloat64(parts[4])
		superNet := utils.MustFloat64(parts[5])

		flow := FundFlow{
			Date:     date,
			Code:     code,
			Name:     stockName,
			MainNet:  mainNet,
			SuperNet: superNet,
			BigNet:   bigNet,
			MidNet:   midNet,
			SmallNet: smallNet,
		}
		flows = append(flows, flow)
		return true
	})

	return flows, nil
}

// StockIndividualFundFlowRankEm 获取资金流向排行（东方财富数据源）
//
// 参数:
//   - indicator: 排行类型，"今日"/"3日"/"5日"/"10日"
//
// 返回:
//   - []FundFlow: 资金流向排行列表
//   - error: 错误信息
//
// 示例:
//
//	flows, err := stock.StockIndividualFundFlowRankEm("今日")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, f := range flows[:10] {
//	    fmt.Printf("%s(%s): 主力净流入 %.2f 万\n", f.Name, f.Code, f.MainNet/10000)
//	}
func StockIndividualFundFlowRankEm(indicator string) ([]FundFlow, error) {
	// 映射排行类型到 API 参数
	dayMap := map[string]string{
		"今日":  "f62",
		"3日":  "f267",
		"5日":  "f164",
		"10日": "f174",
	}

	sortField, ok := dayMap[indicator]
	if !ok {
		return nil, fmt.Errorf("不支持的排行类型: %s, 可选: 今日/3日/5日/10日", indicator)
	}

	params := map[string]string{
		"pn":     "1",
		"pz":     "500",
		"po":     "1",
		"np":     "1",
		"fltt":   "2",
		"invt":   "2",
		"fid":    sortField,
		"fs":     "m:0+t:6,m:0+t:80,m:1+t:2,m:1+t:23,m:0+t:81+s:2048",
		"fields": "f2,f3,f12,f14,f62,f184,f66,f69,f72,f75,f78,f81,f84,f87,f204,f205,f124",
	}

	headers := map[string]string{
		"Referer": emQuoteReferer,
	}

	resp, err := utils.GetWithHeaders(emFundFlowRankURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("获取资金流向排行失败: %w", err)
	}

	// 解析 JSON 数据
	text := resp.String()
	dataPath := "data.diff"
	result := gjson.Get(text, dataPath)
	if !result.Exists() {
		return nil, fmt.Errorf("解析资金流向排行数据失败: 未找到 %s", dataPath)
	}

	flows := make([]FundFlow, 0, result.Get("#").Int())
	now := time.Now()

	result.ForEach(func(_, value gjson.Result) bool {
		flow := FundFlow{
			Date:       now,
			Code:       value.Get("f12").String(),
			Name:       value.Get("f14").String(),
			Price:      utils.MustFloat64(value.Get("f2").String()),
			ChangePct:  utils.MustFloat64(value.Get("f3").String()),
			MainNet:    utils.MustFloat64(value.Get("f62").String()),
			MainNetPct: utils.MustFloat64(value.Get("f184").String()),
			SuperIn:    utils.MustFloat64(value.Get("f66").String()),
			SuperOut:   utils.MustFloat64(value.Get("f69").String()),
			BigIn:      utils.MustFloat64(value.Get("f72").String()),
			BigOut:     utils.MustFloat64(value.Get("f75").String()),
			MidIn:      utils.MustFloat64(value.Get("f78").String()),
			MidOut:     utils.MustFloat64(value.Get("f81").String()),
			SmallIn:    utils.MustFloat64(value.Get("f84").String()),
			SmallOut:   utils.MustFloat64(value.Get("f87").String()),
		}
		// 计算净流入
		flow.SuperNet = flow.SuperIn - flow.SuperOut
		flow.BigNet = flow.BigIn - flow.BigOut
		flow.MidNet = flow.MidIn - flow.MidOut
		flow.SmallNet = flow.SmallIn - flow.SmallOut

		flows = append(flows, flow)
		return true
	})

	return flows, nil
}

// StockMarketFundFlowEm 获取大盘资金流向（东方财富数据源）
//
// 返回:
//   - []MarketFundFlow: 大盘资金流向列表
//   - error: 错误信息
//
// 示例:
//
//	flows, err := stock.StockMarketFundFlowEm()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, f := range flows[:10] {
//	    fmt.Printf("%s: 主力净流入 %.2f 亿\n", f.Date.Format("2006-01-02"), f.MainNet/100000000)
//	}
func StockMarketFundFlowEm() ([]MarketFundFlow, error) {
	// 上证指数资金流向
	params := map[string]string{
		"secid":  "1.000001", // 上证指数
		"lmt":    "0",
		"klt":    "101",
		"fields": "f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61,f62,f63,f64,f65",
	}

	headers := map[string]string{
		"Referer": emQuoteReferer,
	}

	resp, err := utils.GetWithHeaders(emMarketFundFlowURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("获取大盘资金流向失败: %w", err)
	}

	// 解析 JSON 数据
	text := resp.String()
	dataPath := "data.klines"
	result := gjson.Get(text, dataPath)
	if !result.Exists() {
		return nil, fmt.Errorf("解析大盘资金流向数据失败: 未找到 %s", dataPath)
	}

	flows := make([]MarketFundFlow, 0)
	result.ForEach(func(_, value gjson.Result) bool {
		// 数据格式: 日期,主力净流入,小单净流入,中单净流入,大单净流入,超大单净流入
		parts := splitString(value.String(), ",")
		if len(parts) < 6 {
			return true
		}

		date, _ := time.Parse("2006-01-02", parts[0])
		mainNet := utils.MustFloat64(parts[1])
		smallNet := utils.MustFloat64(parts[2])

		flow := MarketFundFlow{
			Date:      date,
			Index:     "上证指数",
			MainNet:   mainNet,
			RetailNet: smallNet,
		}
		flows = append(flows, flow)
		return true
	})

	return flows, nil
}
