package index

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
)

// KqFzPriceIndex 柯桥纺织价格指数
type KqFzPriceIndex struct {
	Period     time.Time `json:"period"`      // 期次
	Index      float64   `json:"index"`       // 指数
	ChangeRate float64   `json:"change_rate"` // 涨跌幅
}

// KqFzProsperityIndex 柯桥纺织景气指数
type KqFzProsperityIndex struct {
	Period               time.Time `json:"period"`               // 期次
	TotalProsperityIndex float64   `json:"total_prosperity_idx"` // 总景气指数
	ChangeRate           float64   `json:"change_rate"`          // 涨跌幅
	CirculationIndex     float64   `json:"circulation_idx"`      // 流通景气指数
	ProductionIndex      float64   `json:"production_idx"`       // 生产景气指数
}

// KqFzForeignTradeIndex 柯桥纺织外贸指数
type KqFzForeignTradeIndex struct {
	Period               time.Time `json:"period"`                 // 期次
	PriceIndex           float64   `json:"price_index"`            // 价格指数
	PriceChangeRate      float64   `json:"price_change_rate"`      // 价格指数涨跌幅
	ProsperityIndex      float64   `json:"prosperity_index"`       // 景气指数
	ProsperityChangeRate float64   `json:"prosperity_change_rate"` // 景气指数涨跌幅
}

// kqFzResponse API 响应结构
type kqFzResponse struct {
	Page   int               `json:"page"`
	Result []json.RawMessage `json:"result"`
}

// IndexKqFzPrice 柯桥纺织价格指数
//
// 数据源: http://www.kqindex.cn/flzs/jiage
//
// 返回:
//   - []KqFzPriceIndex: 柯桥纺织价格指数数据
//   - error: 错误信息
func IndexKqFzPrice() ([]KqFzPriceIndex, error) {
	return fetchKqFzPrice("1_1")
}

// IndexKqFzProsperity 柯桥纺织景气指数
//
// 数据源: http://www.kqindex.cn/flzs/jingqi
//
// 返回:
//   - []KqFzProsperityIndex: 柯桥纺织景气指数数据
//   - error: 错误信息
func IndexKqFzProsperity() ([]KqFzProsperityIndex, error) {
	return fetchKqFzProsperity("1_2")
}

// IndexKqFzForeignTrade 柯桥纺织外贸指数
//
// 数据源: http://www.kqindex.cn/flzs/waimao
//
// 返回:
//   - []KqFzForeignTradeIndex: 柯桥纺织外贸指数数据
//   - error: 错误信息
func IndexKqFzForeignTrade() ([]KqFzForeignTradeIndex, error) {
	return fetchKqFzForeignTrade("2")
}

// fetchKqFzPrice 获取柯桥纺织价格指数
func fetchKqFzPrice(indexType string) ([]KqFzPriceIndex, error) {
	apiURL := "http://www.kqindex.cn/flzs/table_data"

	// 先获取总页数
	params := map[string]string{
		"category":  "0",
		"start":     "",
		"end":       "",
		"indexType": indexType,
		"pageindex": "1",
	}

	resp, err := utils.Get(apiURL, params)
	if err != nil {
		return nil, fmt.Errorf("请求柯桥纺织价格指数失败: %w", err)
	}

	var apiResp kqFzResponse
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, fmt.Errorf("解析柯桥纺织价格指数响应失败: %w", err)
	}

	result := make([]KqFzPriceIndex, 0)

	// 遍历所有页
	for page := 1; page <= apiResp.Page; page++ {
		params["pageindex"] = fmt.Sprintf("%d", page)
		resp, err := utils.Get(apiURL, params)
		if err != nil {
			continue
		}

		var pageResp kqFzResponse
		if err := json.Unmarshal(resp.Body(), &pageResp); err != nil {
			continue
		}

		for _, item := range pageResp.Result {
			var row []interface{}
			if err := json.Unmarshal(item, &row); err != nil {
				continue
			}
			if len(row) < 3 {
				continue
			}

			period, _ := time.Parse("2006-01-02", toString(row[0])[:10])
			result = append(result, KqFzPriceIndex{
				Period:     period,
				Index:      toFloat64(row[1]),
				ChangeRate: toFloat64(row[2]),
			})
		}
	}

	return result, nil
}

// fetchKqFzProsperity 获取柯桥纺织景气指数
func fetchKqFzProsperity(indexType string) ([]KqFzProsperityIndex, error) {
	apiURL := "http://www.kqindex.cn/flzs/table_data"

	params := map[string]string{
		"category":  "0",
		"start":     "",
		"end":       "",
		"indexType": indexType,
		"pageindex": "1",
	}

	resp, err := utils.Get(apiURL, params)
	if err != nil {
		return nil, fmt.Errorf("请求柯桥纺织景气指数失败: %w", err)
	}

	var apiResp kqFzResponse
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, fmt.Errorf("解析柯桥纺织景气指数响应失败: %w", err)
	}

	result := make([]KqFzProsperityIndex, 0)

	for page := 1; page <= apiResp.Page; page++ {
		params["pageindex"] = fmt.Sprintf("%d", page)
		resp, err := utils.Get(apiURL, params)
		if err != nil {
			continue
		}

		var pageResp kqFzResponse
		if err := json.Unmarshal(resp.Body(), &pageResp); err != nil {
			continue
		}

		for _, item := range pageResp.Result {
			var row []interface{}
			if err := json.Unmarshal(item, &row); err != nil {
				continue
			}
			if len(row) < 5 {
				continue
			}

			period, _ := time.Parse("2006-01-02", toString(row[0])[:10])
			result = append(result, KqFzProsperityIndex{
				Period:               period,
				TotalProsperityIndex: toFloat64(row[1]),
				ChangeRate:           toFloat64(row[2]),
				CirculationIndex:     toFloat64(row[3]),
				ProductionIndex:      toFloat64(row[4]),
			})
		}
	}

	return result, nil
}

// fetchKqFzForeignTrade 获取柯桥纺织外贸指数
func fetchKqFzForeignTrade(indexType string) ([]KqFzForeignTradeIndex, error) {
	apiURL := "http://www.kqindex.cn/flzs/table_data"

	params := map[string]string{
		"category":  "0",
		"start":     "",
		"end":       "",
		"indexType": indexType,
		"pageindex": "1",
	}

	resp, err := utils.Get(apiURL, params)
	if err != nil {
		return nil, fmt.Errorf("请求柯桥纺织外贸指数失败: %w", err)
	}

	var apiResp kqFzResponse
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, fmt.Errorf("解析柯桥纺织外贸指数响应失败: %w", err)
	}

	result := make([]KqFzForeignTradeIndex, 0)

	for page := 1; page <= apiResp.Page; page++ {
		params["pageindex"] = fmt.Sprintf("%d", page)
		resp, err := utils.Get(apiURL, params)
		if err != nil {
			continue
		}

		var pageResp kqFzResponse
		if err := json.Unmarshal(resp.Body(), &pageResp); err != nil {
			continue
		}

		for _, item := range pageResp.Result {
			var row []interface{}
			if err := json.Unmarshal(item, &row); err != nil {
				continue
			}
			if len(row) < 5 {
				continue
			}

			period, _ := time.Parse("2006-01-02", toString(row[0])[:10])
			result = append(result, KqFzForeignTradeIndex{
				Period:               period,
				PriceIndex:           toFloat64(row[1]),
				PriceChangeRate:      toFloat64(row[2]),
				ProsperityIndex:      toFloat64(row[3]),
				ProsperityChangeRate: toFloat64(row[4]),
			})
		}
	}

	return result, nil
}

// toString 安全转换为 string
func toString(v interface{}) string {
	switch val := v.(type) {
	case string:
		return val
	case float64:
		return fmt.Sprintf("%v", val)
	default:
		return ""
	}
}

// toFloat64 安全转换为 float64
func toFloat64(v interface{}) float64 {
	switch val := v.(type) {
	case float64:
		return val
	case int:
		return float64(val)
	case int64:
		return float64(val)
	case string:
		return utils.MustParseFloat(val)
	default:
		return 0
	}
}
