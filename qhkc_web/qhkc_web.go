package qhkc_web

import (
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

const (
	// 奇货可查网页版API地址
	qhkcWebURL = "https://qhkch.com/ajax"
)

// FuturesData 期货数据结构
type FuturesData struct {
	Symbol     string  `json:"symbol"`      // 品种代码
	Name       string  `json:"name"`        // 品种名称
	Price      float64 `json:"price"`       // 价格
	Change     float64 `json:"change"`      // 涨跌
	ChangePct  float64 `json:"change_pct"`  // 涨跌幅
	Volume     int64   `json:"volume"`      // 成交量
	OpenInt    int64   `json:"open_int"`    // 持仓量
	UpdateTime string  `json:"update_time"` // 更新时间
}

// BasisData 基差数据结构
type BasisData struct {
	Symbol    string  `json:"symbol"`     // 品种代码
	Name      string  `json:"name"`       // 品种名称
	SpotPrice float64 `json:"spot_price"` // 现货价格
	FutPrice  float64 `json:"fut_price"`  // 期货价格
	Basis     float64 `json:"basis"`      // 基差
	BasisPct  float64 `json:"basis_pct"`  // 基差率
	Date      string  `json:"date"`       // 日期
}

// InventoryData 库存数据结构
type InventoryData struct {
	Symbol    string  `json:"symbol"`     // 品种代码
	Name      string  `json:"name"`       // 品种名称
	Inventory float64 `json:"inventory"`  // 库存量
	Change    float64 `json:"change"`     // 变化量
	ChangePct float64 `json:"change_pct"` // 变化率
	Date      string  `json:"date"`       // 日期
}

// QhkcWebBasis 获取奇货可查网页版基差数据
//
// 目标地址: https://qhkch.com/
//
// 参数:
//   - symbol: 品种代码，如 "RB"
//
// 返回:
//   - []BasisData: 基差数据列表
//   - error: 错误信息
//
// 示例:
//
//	data, err := qhkc_web.QhkcWebBasis("RB")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range data {
//	    fmt.Printf("%s 基差: %.2f 基差率: %.2f%%\n", item.Symbol, item.Basis, item.BasisPct)
//	}
func QhkcWebBasis(symbol string) ([]BasisData, error) {
	url := fmt.Sprintf("%s/basis/%s", qhkcWebURL, strings.ToUpper(symbol))

	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		"Referer":    "https://qhkch.com/",
	}

	resp, err := utils.GetWithHeaders(url, nil, headers)
	if err != nil {
		return nil, fmt.Errorf("获取基差数据失败: %w", err)
	}

	text := resp.String()
	result := gjson.Parse(text)

	var items []BasisData
	if result.IsArray() {
		result.ForEach(func(_, value gjson.Result) bool {
			item := BasisData{
				Symbol:    value.Get("symbol").String(),
				Name:      value.Get("name").String(),
				SpotPrice: value.Get("spot_price").Float(),
				FutPrice:  value.Get("fut_price").Float(),
				Basis:     value.Get("basis").Float(),
				BasisPct:  value.Get("basis_pct").Float(),
				Date:      value.Get("date").String(),
			}
			items = append(items, item)
			return true
		})
	}

	return items, nil
}

// QhkcWebInventory 获取奇货可查网页版库存数据
//
// 目标地址: https://qhkch.com/
//
// 参数:
//   - symbol: 品种代码，如 "RB"
//
// 返回:
//   - []InventoryData: 库存数据列表
//   - error: 错误信息
//
// 示例:
//
//	data, err := qhkc_web.QhkcWebInventory("RB")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range data {
//	    fmt.Printf("%s 库存: %.2f 变化: %.2f\n", item.Symbol, item.Inventory, item.Change)
//	}
func QhkcWebInventory(symbol string) ([]InventoryData, error) {
	url := fmt.Sprintf("%s/inventory/%s", qhkcWebURL, strings.ToUpper(symbol))

	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		"Referer":    "https://qhkch.com/",
	}

	resp, err := utils.GetWithHeaders(url, nil, headers)
	if err != nil {
		return nil, fmt.Errorf("获取库存数据失败: %w", err)
	}

	text := resp.String()
	result := gjson.Parse(text)

	var items []InventoryData
	if result.IsArray() {
		result.ForEach(func(_, value gjson.Result) bool {
			item := InventoryData{
				Symbol:    value.Get("symbol").String(),
				Name:      value.Get("name").String(),
				Inventory: value.Get("inventory").Float(),
				Change:    value.Get("change").Float(),
				ChangePct: value.Get("change_pct").Float(),
				Date:      value.Get("date").String(),
			}
			items = append(items, item)
			return true
		})
	}

	return items, nil
}

// QhkcWebProfit 获取奇货可查网页版利润数据
//
// 目标地址: https://qhkch.com/
//
// 参数:
//   - symbol: 品种代码，如 "RB"
//
// 返回:
//   - []map[string]interface{}: 利润数据列表
//   - error: 错误信息
//
// 示例:
//
//	data, err := qhkc_web.QhkcWebProfit("RB")
//	if err != nil {
//	    log.Fatal(err)
//	}
func QhkcWebProfit(symbol string) ([]map[string]interface{}, error) {
	url := fmt.Sprintf("%s/profit/%s", qhkcWebURL, strings.ToUpper(symbol))

	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		"Referer":    "https://qhkch.com/",
	}

	resp, err := utils.GetWithHeaders(url, nil, headers)
	if err != nil {
		return nil, fmt.Errorf("获取利润数据失败: %w", err)
	}

	text := resp.String()
	result := gjson.Parse(text)

	var items []map[string]interface{}
	if result.IsArray() {
		result.ForEach(func(_, value gjson.Result) bool {
			item := make(map[string]interface{})
			value.ForEach(func(key, val gjson.Result) bool {
				item[key.String()] = val.Value()
				return true
			})
			items = append(items, item)
			return true
		})
	}

	return items, nil
}
