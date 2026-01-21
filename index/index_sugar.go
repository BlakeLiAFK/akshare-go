package index

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
)

// SugarIndex 已在 types.go 中定义

// InnerQuoteSugar 配额内进口糖估算指数
type InnerQuoteSugar struct {
	Date               time.Time `json:"date"`           // 日期
	ProfitMargin       float64   `json:"profit_margin"`  // 利润空间
	ThailandSugar      float64   `json:"thailand_sugar"` // 泰国糖
	ThailandMA5        float64   `json:"thailand_ma5"`   // 泰国MA5
	BrazilMA5          float64   `json:"brazil_ma5"`     // 巴西MA5
	ProfitMA5          float64   `json:"profit_ma5"`     // 利润MA5
	BrazilMA10         float64   `json:"brazil_ma10"`    // 巴西MA10
	BrazilSugar        float64   `json:"brazil_sugar"`   // 巴西糖
	LiuzhouSpotPrice   float64   `json:"liuzhou_spot"`   // 柳州现货价
	GuangzhouSpotPrice float64   `json:"guangzhou_spot"` // 广州现货价
	ThailandMA10       float64   `json:"thailand_ma10"`  // 泰国MA10
	ProfitMA30         float64   `json:"profit_ma30"`    // 利润MA30
	ProfitMA10         float64   `json:"profit_ma10"`    // 利润MA10
}

// OuterQuoteSugar 配额外进口糖估算指数
type OuterQuoteSugar struct {
	Date                 time.Time `json:"date"`                   // 日期
	BrazilImportCost     float64   `json:"brazil_import_cost"`     // 巴西糖进口成本
	ThailandImportProfit float64   `json:"thailand_import_profit"` // 泰国糖进口利润空间
	BrazilImportProfit   float64   `json:"brazil_import_profit"`   // 巴西糖进口利润空间
	ThailandImportCost   float64   `json:"thailand_import_cost"`   // 泰国糖进口成本
	RizhaoSpotPrice      float64   `json:"rizhao_spot"`            // 日照现货价
}

// sugarIndexResponse 食糖指数API响应
type sugarIndexResponse struct {
	Category []string        `json:"category"`
	Data     [][]interface{} `json:"data"`
}

// IndexSugarMsweet 中国食糖指数
//
// 数据源: https://www.msweet.com.cn/mtkj/sjzx13/index.html
//
// 返回:
//   - []SugarIndex: 中国食糖指数数据
//   - error: 错误信息
func IndexSugarMsweet() ([]SugarIndex, error) {
	apiURL := "https://www.msweet.com.cn/eportal/ui"

	params := map[string]string{
		"struts.portlet.action": "/portlet/price!getSTZSJson.action",
		"moduleId":              "cb752447cfe24b44b18c7a7e9abab048",
	}

	resp, err := utils.Get(apiURL, params)
	if err != nil {
		return nil, fmt.Errorf("请求中国食糖指数失败: %w", err)
	}

	var apiResp sugarIndexResponse
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, fmt.Errorf("解析中国食糖指数响应失败: %w", err)
	}

	length := len(apiResp.Category)
	if length == 0 {
		return []SugarIndex{}, nil
	}

	result := make([]SugarIndex, 0, length)
	for i := 0; i < length; i++ {
		dateStr := apiResp.Category[i]
		if len(dateStr) >= 10 {
			dateStr = dateStr[:10]
		}

		var composite, rawSugar, spot float64
		if len(apiResp.Data) > 0 && i < len(apiResp.Data[0]) {
			composite = toFloat64(apiResp.Data[0][i])
		}
		if len(apiResp.Data) > 1 && i < len(apiResp.Data[1]) {
			rawSugar = toFloat64(apiResp.Data[1][i])
		}
		if len(apiResp.Data) > 2 && i < len(apiResp.Data[2]) {
			spot = toFloat64(apiResp.Data[2][i])
		}

		result = append(result, SugarIndex{
			Date:           dateStr,
			CompositePrice: composite,
			RawSugarPrice:  rawSugar,
			SpotPrice:      spot,
		})
	}

	return result, nil
}

// innerQuoteSugarResponse 配额内进口糖API响应
type innerQuoteSugarResponse struct {
	Category []string        `json:"category"`
	Data     [][]interface{} `json:"data"`
}

// IndexInnerQuoteSugarMsweet 配额内进口糖估算指数
//
// 数据源: https://www.msweet.com.cn/mtkj/sjzx13/index.html
//
// 返回:
//   - []InnerQuoteSugar: 配额内进口糖估算指数数据
//   - error: 错误信息
func IndexInnerQuoteSugarMsweet() ([]InnerQuoteSugar, error) {
	apiURL := "https://www.msweet.com.cn/datacenterapply/datacenter/json/JinKongTang.json"

	resp, err := utils.Get(apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("请求配额内进口糖估算指数失败: %w", err)
	}

	var apiResp innerQuoteSugarResponse
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, fmt.Errorf("解析配额内进口糖估算指数响应失败: %w", err)
	}

	length := len(apiResp.Category)
	if length == 0 {
		return []InnerQuoteSugar{}, nil
	}

	result := make([]InnerQuoteSugar, 0, length)
	for i := 0; i < length; i++ {
		dateStr := strings.ReplaceAll(apiResp.Category[i], "/", "-")
		date, err := time.Parse("2006-01-02", dateStr[:10])
		if err != nil {
			continue
		}

		item := InnerQuoteSugar{
			Date: date,
		}

		if len(apiResp.Data) > 0 && i < len(apiResp.Data[0]) {
			item.ProfitMargin = toFloat64(apiResp.Data[0][i])
		}
		if len(apiResp.Data) > 1 && i < len(apiResp.Data[1]) {
			item.ThailandSugar = toFloat64(apiResp.Data[1][i])
		}
		if len(apiResp.Data) > 2 && i < len(apiResp.Data[2]) {
			item.ThailandMA5 = toFloat64(apiResp.Data[2][i])
		}
		if len(apiResp.Data) > 3 && i < len(apiResp.Data[3]) {
			item.BrazilMA5 = toFloat64(apiResp.Data[3][i])
		}
		if len(apiResp.Data) > 4 && i < len(apiResp.Data[4]) {
			item.ProfitMA5 = toFloat64(apiResp.Data[4][i])
		}
		if len(apiResp.Data) > 5 && i < len(apiResp.Data[5]) {
			item.BrazilMA10 = toFloat64(apiResp.Data[5][i])
		}
		if len(apiResp.Data) > 6 && i < len(apiResp.Data[6]) {
			item.BrazilSugar = toFloat64(apiResp.Data[6][i])
		}
		if len(apiResp.Data) > 7 && i < len(apiResp.Data[7]) {
			item.LiuzhouSpotPrice = toFloat64(apiResp.Data[7][i])
		}
		if len(apiResp.Data) > 8 && i < len(apiResp.Data[8]) {
			item.GuangzhouSpotPrice = toFloat64(apiResp.Data[8][i])
		}
		if len(apiResp.Data) > 9 && i < len(apiResp.Data[9]) {
			item.ThailandMA10 = toFloat64(apiResp.Data[9][i])
		}
		if len(apiResp.Data) > 10 && i < len(apiResp.Data[10]) {
			item.ProfitMA30 = toFloat64(apiResp.Data[10][i])
		}
		if len(apiResp.Data) > 11 && i < len(apiResp.Data[11]) {
			item.ProfitMA10 = toFloat64(apiResp.Data[11][i])
		}

		result = append(result, item)
	}

	return result, nil
}

// IndexOuterQuoteSugarMsweet 配额外进口糖估算指数
//
// 数据源: https://www.msweet.com.cn/mtkj/sjzx13/index.html
//
// 返回:
//   - []OuterQuoteSugar: 配额外进口糖估算指数数据
//   - error: 错误信息
func IndexOuterQuoteSugarMsweet() ([]OuterQuoteSugar, error) {
	apiURL := "https://www.msweet.com.cn/datacenterapply/datacenter/json/Jkpewlr.json"

	resp, err := utils.Get(apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("请求配额外进口糖估算指数失败: %w", err)
	}

	var apiResp innerQuoteSugarResponse
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, fmt.Errorf("解析配额外进口糖估算指数响应失败: %w", err)
	}

	length := len(apiResp.Category)
	if length == 0 {
		return []OuterQuoteSugar{}, nil
	}

	result := make([]OuterQuoteSugar, 0, length)
	for i := 0; i < length; i++ {
		dateStr := strings.ReplaceAll(apiResp.Category[i], "/", "-")
		date, err := time.Parse("2006-01-02", dateStr[:10])
		if err != nil {
			continue
		}

		item := OuterQuoteSugar{
			Date: date,
		}

		if len(apiResp.Data) > 0 && i < len(apiResp.Data[0]) {
			item.BrazilImportCost = toFloat64(apiResp.Data[0][i])
		}
		if len(apiResp.Data) > 1 && i < len(apiResp.Data[1]) {
			item.ThailandImportProfit = toFloat64(apiResp.Data[1][i])
		}
		if len(apiResp.Data) > 2 && i < len(apiResp.Data[2]) {
			item.BrazilImportProfit = toFloat64(apiResp.Data[2][i])
		}
		if len(apiResp.Data) > 3 && i < len(apiResp.Data[3]) {
			item.ThailandImportCost = toFloat64(apiResp.Data[3][i])
		}
		if len(apiResp.Data) > 4 && i < len(apiResp.Data[4]) {
			item.RizhaoSpotPrice = toFloat64(apiResp.Data[4][i])
		}

		result = append(result, item)
	}

	return result, nil
}
