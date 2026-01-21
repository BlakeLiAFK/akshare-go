package index

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
)

// CflpIndex 中国公路物流运价/运量指数
type CflpIndex struct {
	Date       time.Time `json:"date"`        // 日期
	BaseIndex  float64   `json:"base_index"`  // 定基指数
	ChainIndex float64   `json:"chain_index"` // 环比指数
	YoyIndex   float64   `json:"yoy_index"`   // 同比指数
}

// cflpChartResponse CFLP API 响应结构
type cflpChartResponse struct {
	Chart1 struct {
		XLebal []string `json:"xLebal"`
		YLebal []string `json:"yLebal"`
	} `json:"chart1"`
	Chart2 struct {
		YLebal []string `json:"yLebal"`
	} `json:"chart2"`
	Chart3 struct {
		YLebal []string `json:"yLebal"`
	} `json:"chart3"`
}

// IndexPriceCflp 中国公路物流运价指数
//
// 数据源: http://index.0256.cn/expx.htm
//
// 参数:
//   - symbol: 指数类型，可选 "周指数", "月指数", "季度指数", "年度指数"
//
// 返回:
//   - []CflpIndex: 运价指数数据
//   - error: 错误信息
func IndexPriceCflp(symbol string) ([]CflpIndex, error) {
	symbolMap := map[string]string{
		"周指数":  "2",
		"月指数":  "3",
		"季度指数": "4",
		"年度指数": "5",
	}

	typeID, ok := symbolMap[symbol]
	if !ok {
		return nil, fmt.Errorf("无效的 symbol: %s, 可选值: 周指数, 月指数, 季度指数, 年度指数", symbol)
	}

	apiURL := "http://index.0256.cn/expcenter_trend.action"

	formData := map[string]string{
		"marketId":       "1",
		"attribute1":     "5",
		"exponentTypeId": typeID,
		"cateId":         "2",
		"attribute2":     "华北",
		"city":           "",
		"startLine":      "",
		"endLine":        "",
	}

	headers := map[string]string{
		"Origin":  "http://index.0256.cn",
		"Referer": "http://index.0256.cn/expx.htm",
	}

	respBody, err := utils.PostFormWithHeaders(apiURL, formData, headers)
	if err != nil {
		return nil, fmt.Errorf("请求运价指数失败: %w", err)
	}

	var resp cflpChartResponse
	if err := json.Unmarshal(respBody.Body(), &resp); err != nil {
		return nil, fmt.Errorf("解析运价指数响应失败: %w", err)
	}

	return parseCflpResponse(resp)
}

// IndexVolumeCflp 中国公路物流运量指数
//
// 数据源: http://index.0256.cn/expx.htm
//
// 参数:
//   - symbol: 指数类型，可选 "月指数", "季度指数", "年度指数"
//
// 返回:
//   - []CflpIndex: 运量指数数据
//   - error: 错误信息
func IndexVolumeCflp(symbol string) ([]CflpIndex, error) {
	symbolMap := map[string]string{
		"月指数":  "3",
		"季度指数": "4",
		"年度指数": "5",
	}

	typeID, ok := symbolMap[symbol]
	if !ok {
		return nil, fmt.Errorf("无效的 symbol: %s, 可选值: 月指数, 季度指数, 年度指数", symbol)
	}

	apiURL := "http://index.0256.cn/volume_query.action"

	formData := map[string]string{
		"type":       "1",
		"marketId":   "1",
		"expTypeId":  typeID,
		"startDate1": "",
		"endDate1":   "",
		"city":       "",
		"startDate3": "",
		"endDate3":   "",
	}

	headers := map[string]string{
		"Origin":  "http://index.0256.cn",
		"Referer": "http://index.0256.cn/expx.htm",
	}

	respBody, err := utils.PostFormWithHeaders(apiURL, formData, headers)
	if err != nil {
		return nil, fmt.Errorf("请求运量指数失败: %w", err)
	}

	var resp cflpChartResponse
	if err := json.Unmarshal(respBody.Body(), &resp); err != nil {
		return nil, fmt.Errorf("解析运量指数响应失败: %w", err)
	}

	return parseCflpResponse(resp)
}

// parseCflpResponse 解析 CFLP 响应数据
func parseCflpResponse(resp cflpChartResponse) ([]CflpIndex, error) {
	length := len(resp.Chart1.XLebal)
	if length == 0 {
		return []CflpIndex{}, nil
	}

	result := make([]CflpIndex, 0, length)
	for i := 0; i < length; i++ {
		date, err := time.Parse("2006-01-02", resp.Chart1.XLebal[i])
		if err != nil {
			// 尝试其他格式
			date, err = time.Parse("2006/01/02", resp.Chart1.XLebal[i])
			if err != nil {
				continue
			}
		}

		item := CflpIndex{
			Date:       date,
			BaseIndex:  utils.MustParseFloat(resp.Chart1.YLebal[i]),
			ChainIndex: utils.MustParseFloat(resp.Chart2.YLebal[i]),
			YoyIndex:   utils.MustParseFloat(resp.Chart3.YLebal[i]),
		}
		result = append(result, item)
	}

	return result, nil
}
