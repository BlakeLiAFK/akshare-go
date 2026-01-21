package stock

import (
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

const (
	// 东方财富板块 API
	emBoardURL = "https://push2.eastmoney.com/api/qt/clist/get"

	// 概念板块字段列表
	conceptBoardFields = "f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18,f20,f21,f24,f25,f26,f22,f33,f11,f62,f128,f136,f115,f152,f124,f107,f104,f105,f140,f141,f207,f208,f209,f222"

	// 板块成分股字段列表
	boardStockFields = "f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18"
)

// StockBoardConceptNameEm 获取概念板块名称列表（东方财富数据源）
//
// 返回:
//   - []BoardInfo: 概念板块列表
//   - error: 错误信息
//
// 示例:
//
//	boards, err := stock.StockBoardConceptNameEm()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, b := range boards[:10] {
//	    fmt.Printf("%s(%s): %.2f%%\n", b.Name, b.Code, b.ChangePct)
//	}
func StockBoardConceptNameEm() ([]BoardInfo, error) {
	params := map[string]string{
		"pn":     "1",
		"pz":     "500",
		"po":     "1",
		"np":     "1",
		"fltt":   "2",
		"invt":   "2",
		"fid":    "f3",
		"fs":     "m:90+t:3+f:!50", // 概念板块筛选条件
		"fields": conceptBoardFields,
	}

	headers := map[string]string{
		"Referer": emQuoteReferer,
	}

	resp, err := utils.GetWithHeaders(emBoardURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("获取概念板块失败: %w", err)
	}

	// 解析 JSON 数据
	text := resp.String()
	dataPath := "data.diff"
	result := gjson.Get(text, dataPath)
	if !result.Exists() {
		return nil, fmt.Errorf("解析概念板块数据失败: 未找到 %s", dataPath)
	}

	boards := make([]BoardInfo, 0, result.Get("#").Int())
	result.ForEach(func(_, value gjson.Result) bool {
		upCount := utils.MustInt(value.Get("f104").String())
		downCount := utils.MustInt(value.Get("f105").String())

		board := BoardInfo{
			Code:         value.Get("f12").String(),
			Name:         value.Get("f14").String(),
			Change:       utils.MustFloat64(value.Get("f4").String()),
			ChangePct:    utils.MustFloat64(value.Get("f3").String()),
			Volume:       utils.MustInt64(value.Get("f5").String()),
			Amount:       utils.MustFloat64(value.Get("f6").String()),
			LeadStock:    value.Get("f128").String(),
			LeadStockPct: utils.MustFloat64(value.Get("f140").String()),
			UpCount:      upCount,
			DownCount:    downCount,
			StockCount:   upCount + downCount,
		}
		boards = append(boards, board)
		return true
	})

	return boards, nil
}

// StockBoardConceptSpotEm 获取概念板块实时行情（东方财富数据源）
//
// 返回:
//   - []BoardInfo: 概念板块实时行情列表
//   - error: 错误信息
//
// 说明: 与 StockBoardConceptNameEm 返回相同数据
func StockBoardConceptSpotEm() ([]BoardInfo, error) {
	return StockBoardConceptNameEm()
}

// StockBoardConceptConsEm 获取概念板块成分股（东方财富数据源）
//
// 参数:
//   - boardCode: 板块代码，如 "BK0493"
//
// 返回:
//   - []BoardStock: 成分股列表
//   - error: 错误信息
//
// 示例:
//
//	stocks, err := stock.StockBoardConceptConsEm("BK0493")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, s := range stocks {
//	    fmt.Printf("%s(%s): %.2f %.2f%%\n", s.Name, s.Code, s.Price, s.ChangePct)
//	}
func StockBoardConceptConsEm(boardCode string) ([]BoardStock, error) {
	if boardCode == "" {
		return nil, fmt.Errorf("板块代码不能为空")
	}

	params := map[string]string{
		"pn":     "1",
		"pz":     "2000",
		"po":     "1",
		"np":     "1",
		"fltt":   "2",
		"invt":   "2",
		"fid":    "f3",
		"fs":     fmt.Sprintf("b:%s+f:!50", boardCode), // 板块成分股筛选
		"fields": boardStockFields,
	}

	headers := map[string]string{
		"Referer": emQuoteReferer,
	}

	resp, err := utils.GetWithHeaders(emBoardURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("获取板块成分股失败: %w", err)
	}

	// 解析 JSON 数据
	text := resp.String()
	dataPath := "data.diff"
	result := gjson.Get(text, dataPath)
	if !result.Exists() {
		return nil, fmt.Errorf("解析板块成分股数据失败: 未找到 %s", dataPath)
	}

	stocks := make([]BoardStock, 0, result.Get("#").Int())
	result.ForEach(func(_, value gjson.Result) bool {
		stock := BoardStock{
			Code:      value.Get("f12").String(),
			Name:      value.Get("f14").String(),
			Price:     utils.MustFloat64(value.Get("f2").String()),
			ChangePct: utils.MustFloat64(value.Get("f3").String()),
			Volume:    utils.MustInt64(value.Get("f5").String()),
			Amount:    utils.MustFloat64(value.Get("f6").String()),
		}
		stocks = append(stocks, stock)
		return true
	})

	return stocks, nil
}
