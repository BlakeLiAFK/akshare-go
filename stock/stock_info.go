package stock

import (
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

const (
	// 东方财富股票信息 API
	emStockInfoURL = "https://push2.eastmoney.com/api/qt/stock/get"
)

// StockInfoACodeNameEm 获取 A 股代码名称列表（东方财富数据源）
//
// 返回:
//   - []StockCodeName: 股票代码名称列表
//   - error: 错误信息
//
// 示例:
//
//	stocks, err := stock.StockInfoACodeNameEm()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, s := range stocks[:10] {
//	    fmt.Printf("%s: %s\n", s.Code, s.Name)
//	}
func StockInfoACodeNameEm() ([]StockCodeName, error) {
	params := map[string]string{
		"pn":     "1",
		"pz":     "10000",
		"po":     "1",
		"np":     "1",
		"fltt":   "2",
		"invt":   "2",
		"fid":    "f12",
		"fs":     "m:0+t:6,m:0+t:80,m:1+t:2,m:1+t:23,m:0+t:81+s:2048",
		"fields": "f12,f14",
	}

	headers := map[string]string{
		"Referer": emQuoteReferer,
	}

	resp, err := utils.GetWithHeaders(emStockListURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("获取股票列表失败: %w", err)
	}

	// 解析 JSON 数据
	text := resp.String()
	dataPath := "data.diff"
	result := gjson.Get(text, dataPath)
	if !result.Exists() {
		return nil, fmt.Errorf("解析股票列表数据失败: 未找到 %s", dataPath)
	}

	stocks := make([]StockCodeName, 0, result.Get("#").Int())
	result.ForEach(func(_, value gjson.Result) bool {
		stock := StockCodeName{
			Code: value.Get("f12").String(),
			Name: value.Get("f14").String(),
		}
		stocks = append(stocks, stock)
		return true
	})

	return stocks, nil
}

// StockIndividualInfoEm 获取个股详细信息（东方财富数据源）
//
// 参数:
//   - code: 股票代码，如 "000001"
//
// 返回:
//   - *StockInfo: 股票详细信息
//   - error: 错误信息
//
// 示例:
//
//	info, err := stock.StockIndividualInfoEm("000001")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("%s(%s): 行业=%s, 市盈率=%.2f\n", info.Name, info.Code, info.Industry, info.PE)
func StockIndividualInfoEm(code string) (*StockInfo, error) {
	if code == "" {
		return nil, fmt.Errorf("股票代码不能为空")
	}

	secid := getSecID(code)

	params := map[string]string{
		"secid":  secid,
		"fields": "f57,f58,f84,f85,f116,f117,f162,f163,f167,f169,f170,f171,f173,f183,f184,f185,f186,f187,f188,f189,f190,f191,f192",
	}

	headers := map[string]string{
		"Referer": emQuoteReferer,
	}

	resp, err := utils.GetWithHeaders(emStockInfoURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("获取股票信息失败: %w", err)
	}

	// 解析 JSON 数据
	text := resp.String()
	dataPath := "data"
	result := gjson.Get(text, dataPath)
	if !result.Exists() {
		return nil, fmt.Errorf("解析股票信息数据失败: 未找到 %s", dataPath)
	}

	info := &StockInfo{
		Code:         result.Get("f57").String(),
		Name:         result.Get("f58").String(),
		TotalShares:  utils.MustFloat64(result.Get("f84").String()),
		FloatShares:  utils.MustFloat64(result.Get("f85").String()),
		MarketCap:    utils.MustFloat64(result.Get("f116").String()),
		CirculateCap: utils.MustFloat64(result.Get("f117").String()),
		PE:           utils.MustFloat64(result.Get("f162").String()),
		PEStatic:     utils.MustFloat64(result.Get("f163").String()),
		PB:           utils.MustFloat64(result.Get("f167").String()),
		ROE:          utils.MustFloat64(result.Get("f173").String()),
		EPS:          utils.MustFloat64(result.Get("f183").String()),
		BPS:          utils.MustFloat64(result.Get("f184").String()),
		Industry:     result.Get("f127").String(),
	}

	return info, nil
}

// StockSectorCodeNameEm 获取板块代码名称列表（东方财富数据源）
//
// 参数:
//   - sectorType: 板块类型，"concept"(概念板块) 或 "industry"(行业板块)
//
// 返回:
//   - []StockCodeName: 板块代码名称列表
//   - error: 错误信息
//
// 示例:
//
//	sectors, err := stock.StockSectorCodeNameEm("concept")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, s := range sectors[:10] {
//	    fmt.Printf("%s: %s\n", s.Code, s.Name)
//	}
func StockSectorCodeNameEm(sectorType string) ([]StockCodeName, error) {
	// 根据板块类型设置筛选条件
	var fs string
	switch sectorType {
	case "concept":
		fs = "m:90+t:3+f:!50" // 概念板块
	case "industry":
		fs = "m:90+t:2+f:!50" // 行业板块
	default:
		return nil, fmt.Errorf("不支持的板块类型: %s, 可选: concept/industry", sectorType)
	}

	params := map[string]string{
		"pn":     "1",
		"pz":     "500",
		"po":     "1",
		"np":     "1",
		"fltt":   "2",
		"invt":   "2",
		"fid":    "f12",
		"fs":     fs,
		"fields": "f12,f14",
	}

	headers := map[string]string{
		"Referer": emQuoteReferer,
	}

	resp, err := utils.GetWithHeaders(emBoardURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("获取板块列表失败: %w", err)
	}

	// 解析 JSON 数据
	text := resp.String()
	dataPath := "data.diff"
	result := gjson.Get(text, dataPath)
	if !result.Exists() {
		return nil, fmt.Errorf("解析板块列表数据失败: 未找到 %s", dataPath)
	}

	sectors := make([]StockCodeName, 0, result.Get("#").Int())
	result.ForEach(func(_, value gjson.Result) bool {
		sector := StockCodeName{
			Code: value.Get("f12").String(),
			Name: value.Get("f14").String(),
		}
		sectors = append(sectors, sector)
		return true
	})

	return sectors, nil
}
