package stock

import (
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// StockBoardIndustryNameEm 获取行业板块名称列表（东方财富数据源）
//
// 返回:
//   - []BoardInfo: 行业板块列表
//   - error: 错误信息
//
// 示例:
//
//	boards, err := stock.StockBoardIndustryNameEm()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, b := range boards[:10] {
//	    fmt.Printf("%s(%s): %.2f%%\n", b.Name, b.Code, b.ChangePct)
//	}
func StockBoardIndustryNameEm() ([]BoardInfo, error) {
	params := map[string]string{
		"pn":     "1",
		"pz":     "500",
		"po":     "1",
		"np":     "1",
		"fltt":   "2",
		"invt":   "2",
		"fid":    "f3",
		"fs":     "m:90+t:2+f:!50", // 行业板块筛选条件 (t:2 表示行业)
		"fields": conceptBoardFields,
	}

	headers := map[string]string{
		"Referer": emQuoteReferer,
	}

	resp, err := utils.GetWithHeaders(emBoardURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("获取行业板块失败: %w", err)
	}

	// 解析 JSON 数据
	text := resp.String()
	dataPath := "data.diff"
	result := gjson.Get(text, dataPath)
	if !result.Exists() {
		return nil, fmt.Errorf("解析行业板块数据失败: 未找到 %s", dataPath)
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

// StockBoardIndustrySpotEm 获取行业板块实时行情（东方财富数据源）
//
// 返回:
//   - []BoardInfo: 行业板块实时行情列表
//   - error: 错误信息
//
// 说明: 与 StockBoardIndustryNameEm 返回相同数据
func StockBoardIndustrySpotEm() ([]BoardInfo, error) {
	return StockBoardIndustryNameEm()
}

// StockBoardIndustryConsEm 获取行业板块成分股（东方财富数据源）
//
// 参数:
//   - boardCode: 板块代码，如 "BK0475"
//
// 返回:
//   - []BoardStock: 成分股列表
//   - error: 错误信息
//
// 示例:
//
//	stocks, err := stock.StockBoardIndustryConsEm("BK0475")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, s := range stocks {
//	    fmt.Printf("%s(%s): %.2f %.2f%%\n", s.Name, s.Code, s.Price, s.ChangePct)
//	}
func StockBoardIndustryConsEm(boardCode string) ([]BoardStock, error) {
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
		return nil, fmt.Errorf("获取行业板块成分股失败: %w", err)
	}

	// 解析 JSON 数据
	text := resp.String()
	dataPath := "data.diff"
	result := gjson.Get(text, dataPath)
	if !result.Exists() {
		return nil, fmt.Errorf("解析行业板块成分股数据失败: 未找到 %s", dataPath)
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
