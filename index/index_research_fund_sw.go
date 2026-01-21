package index

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
)

// SWFundIndexRealtime 申万基金指数实时行情
type SWFundIndexRealtime struct {
	IndexCode      string  `json:"index_code"`       // 指数代码
	IndexName      string  `json:"index_name"`       // 指数名称
	LastClose      float64 `json:"last_close"`       // 昨收盘
	DailyChangePct float64 `json:"daily_change_pct"` // 日涨跌幅
	YearChangePct  float64 `json:"year_change_pct"`  // 年涨跌幅
}

// SWFundIndexHist 申万基金指数历史行情
type SWFundIndexHist struct {
	Date      time.Time `json:"date"`       // 日期
	Close     float64   `json:"close"`      // 收盘指数
	Open      float64   `json:"open"`       // 开盘指数
	High      float64   `json:"high"`       // 最高指数
	Low       float64   `json:"low"`        // 最低指数
	ChangePct float64   `json:"change_pct"` // 涨跌幅
}

// swFundRealtimeResponse 申万基金实时行情API响应
type swFundRealtimeResponse struct {
	Data struct {
		List []struct {
			SwIndexCode    string  `json:"swIndexCode"`
			SwIndexName    string  `json:"swIndexName"`
			LastCloseIndex float64 `json:"lastCloseIndex"`
			LastMarkup     float64 `json:"lastMarkup"`
			YearMarkup     float64 `json:"yearMarkup"`
		} `json:"list"`
	} `json:"data"`
}

// swFundHistResponse 申万基金历史行情API响应
type swFundHistResponse struct {
	Data []struct {
		BargainDate string  `json:"bargaindate"`
		CloseIndex  float64 `json:"closeindex"`
		OpenIndex   float64 `json:"openindex"`
		MaxIndex    float64 `json:"maxindex"`
		MinIndex    float64 `json:"minindex"`
		Markup      float64 `json:"markup"`
	} `json:"data"`
}

// IndexRealtimeFundSW 申万基金指数实时行情
//
// 数据源: https://www.swsresearch.com/institute_sw/allIndex/releasedIndex
//
// 参数:
//   - symbol: 指数类型，可选 "基础一级", "基础二级", "基础三级", "特色指数"
//
// 返回:
//   - []SWFundIndexRealtime: 基金指数实时行情数据
//   - error: 错误信息
func IndexRealtimeFundSW(symbol string) ([]SWFundIndexRealtime, error) {
	apiURL := "https://www.swsresearch.com/insWechatSw/fundIndex/pageList"

	payload := map[string]interface{}{
		"pageNo":        1,
		"pageSize":      50,
		"indexTypeName": symbol,
		"sortField":     "",
		"rule":          "",
		"indexType":     1,
	}

	headers := map[string]string{
		"Content-Type": "application/json",
		"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36",
	}

	resp, err := utils.PostJSONWithHeaders(apiURL, nil, payload, headers)
	if err != nil {
		return nil, fmt.Errorf("请求申万基金指数实时行情失败: %w", err)
	}

	var apiResp swFundRealtimeResponse
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, fmt.Errorf("解析申万基金指数实时行情响应失败: %w", err)
	}

	result := make([]SWFundIndexRealtime, 0, len(apiResp.Data.List))
	for _, item := range apiResp.Data.List {
		result = append(result, SWFundIndexRealtime{
			IndexCode:      item.SwIndexCode,
			IndexName:      item.SwIndexName,
			LastClose:      item.LastCloseIndex,
			DailyChangePct: item.LastMarkup,
			YearChangePct:  item.YearMarkup,
		})
	}

	return result, nil
}

// IndexHistFundSW 申万基金指数历史行情
//
// 数据源: https://www.swsresearch.com/institute_sw/allIndex/releasedIndex/fundDetail?code=807100
//
// 参数:
//   - symbol: 基金指数代码，如 "807200"
//   - period: 周期，可选 "day", "week", "month"
//
// 返回:
//   - []SWFundIndexHist: 基金指数历史行情数据
//   - error: 错误信息
func IndexHistFundSW(symbol, period string) ([]SWFundIndexHist, error) {
	periodMap := map[string]string{
		"day":   "DAY",
		"week":  "WEEK",
		"month": "MONTH",
	}

	periodValue, ok := periodMap[period]
	if !ok {
		return nil, fmt.Errorf("无效的 period: %s，可选 day, week, month", period)
	}

	apiURL := "https://www.swsresearch.com/insWechatSw/fundIndex/getFundKChartData"

	payload := map[string]interface{}{
		"swIndexCode": symbol,
		"type":        periodValue,
	}

	headers := map[string]string{
		"Content-Type": "application/json",
		"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36",
	}

	resp, err := utils.PostJSONWithHeaders(apiURL, nil, payload, headers)
	if err != nil {
		return nil, fmt.Errorf("请求申万基金指数历史行情失败: %w", err)
	}

	var apiResp swFundHistResponse
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, fmt.Errorf("解析申万基金指数历史行情响应失败: %w", err)
	}

	result := make([]SWFundIndexHist, 0, len(apiResp.Data))
	for _, item := range apiResp.Data {
		date, err := time.Parse("2006-01-02", item.BargainDate[:10])
		if err != nil {
			continue
		}

		result = append(result, SWFundIndexHist{
			Date:      date,
			Close:     item.CloseIndex,
			Open:      item.OpenIndex,
			High:      item.MaxIndex,
			Low:       item.MinIndex,
			ChangePct: item.Markup,
		})
	}

	return result, nil
}
