package index

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

// SWIndexHist 申万指数历史行情
type SWIndexHist struct {
	Code   string    `json:"code"`   // 代码
	Date   time.Time `json:"date"`   // 日期
	Close  float64   `json:"close"`  // 收盘
	Open   float64   `json:"open"`   // 开盘
	High   float64   `json:"high"`   // 最高
	Low    float64   `json:"low"`    // 最低
	Volume int64     `json:"volume"` // 成交量
	Amount float64   `json:"amount"` // 成交额
}

// SWIndexMin 申万指数分时数据
type SWIndexMin struct {
	Code  string    `json:"code"`  // 代码
	Name  string    `json:"name"`  // 名称
	Price float64   `json:"price"` // 价格
	Date  time.Time `json:"date"`  // 日期
	Time  string    `json:"time"`  // 时间
}

// SWIndexComponent 申万指数成份股
type SWIndexComponent struct {
	Seq          int       `json:"seq"`           // 序号
	StockCode    string    `json:"stock_code"`    // 证券代码
	StockName    string    `json:"stock_name"`    // 证券名称
	Weight       float64   `json:"weight"`        // 最新权重
	IncludedDate time.Time `json:"included_date"` // 计入日期
}

// SWIndexRealtime 申万指数实时行情
type SWIndexRealtime struct {
	IndexCode string  `json:"index_code"` // 指数代码
	IndexName string  `json:"index_name"` // 指数名称
	PreClose  float64 `json:"pre_close"`  // 昨收盘
	Open      float64 `json:"open"`       // 今开盘
	Price     float64 `json:"price"`      // 最新价
	Amount    float64 `json:"amount"`     // 成交额
	Volume    int64   `json:"volume"`     // 成交量
	High      float64 `json:"high"`       // 最高价
	Low       float64 `json:"low"`        // 最低价
}

// SWIndexAnalysis 申万指数分析数据
type SWIndexAnalysis struct {
	IndexCode    string    `json:"index_code"`    // 指数代码
	IndexName    string    `json:"index_name"`    // 指数名称
	Date         time.Time `json:"date"`          // 发布日期
	CloseIndex   float64   `json:"close_index"`   // 收盘指数
	Volume       int64     `json:"volume"`        // 成交量
	ChangePct    float64   `json:"change_pct"`    // 涨跌幅
	TurnoverRate float64   `json:"turnover_rate"` // 换手率
	PE           float64   `json:"pe"`            // 市盈率
	PB           float64   `json:"pb"`            // 市净率
	AvgPrice     float64   `json:"avg_price"`     // 均价
	AmountRatio  float64   `json:"amount_ratio"`  // 成交额占比
	FloatCap     float64   `json:"float_cap"`     // 流通市值
	AvgFloatCap  float64   `json:"avg_float_cap"` // 平均流通市值
	DividendRate float64   `json:"dividend_rate"` // 股息率
}

// SWIndexWeekMonthDate 申万指数周/月报表日期
type SWIndexWeekMonthDate struct {
	Date time.Time `json:"date"` // 日期
}

// swHistResponse 申万指数历史行情API响应
type swHistResponse struct {
	Data []struct {
		SwIndexCode   string  `json:"swindexcode"`
		BargainDate   string  `json:"bargaindate"`
		OpenIndex     float64 `json:"openindex"`
		MaxIndex      float64 `json:"maxindex"`
		MinIndex      float64 `json:"minindex"`
		CloseIndex    float64 `json:"closeindex"`
		BargainAmount int64   `json:"bargainamount"`
		BargainSum    float64 `json:"bargainsum"`
	} `json:"data"`
}

// swMinResponse 申万指数分时API响应
type swMinResponse struct {
	Data []struct {
		L1          string  `json:"l1"`
		L2          string  `json:"l2"`
		L8          float64 `json:"l8"`
		TradingDate string  `json:"trading_date"`
		TradingTime string  `json:"trading_time"`
	} `json:"data"`
}

// swComponentResponse 申万指数成份股API响应
type swComponentResponse struct {
	Data struct {
		Results []struct {
			StockCode     string  `json:"stockcode"`
			StockName     string  `json:"stockname"`
			NewWeight     float64 `json:"newweight"`
			BeginningDate string  `json:"beginningdate"`
		} `json:"results"`
	} `json:"data"`
}

// swRealtimeResponse 申万指数实时行情API响应
type swRealtimeResponse struct {
	Data struct {
		Count   int `json:"count"`
		Results []struct {
			SwIndexCode   string  `json:"swindexcode"`
			SwIndexName   string  `json:"swindexname"`
			LastClose     float64 `json:"lastclose"`
			Open          float64 `json:"open"`
			BargainSum    float64 `json:"bargainsum"`
			MaxIndex      float64 `json:"maxindex"`
			MinIndex      float64 `json:"minindex"`
			CloseIndex    float64 `json:"closeindex"`
			BargainAmount int64   `json:"bargainamount"`
		} `json:"results"`
	} `json:"data"`
}

// swAnalysisResponse 申万指数分析API响应
type swAnalysisResponse struct {
	Data struct {
		Count   int `json:"count"`
		Results []struct {
			SwIndexCode         string  `json:"swindexcode"`
			SwIndexName         string  `json:"swindexname"`
			BargainDate         string  `json:"bargaindate"`
			CloseIndex          float64 `json:"closeindex"`
			BargainAmount       int64   `json:"bargainamount"`
			Markup              float64 `json:"markup"`
			TurnoverRate        float64 `json:"turnoverrate"`
			PE                  float64 `json:"pe"`
			PB                  float64 `json:"pb"`
			MeanPrice           float64 `json:"meanprice"`
			BargainSumRate      float64 `json:"bargainsumrate"`
			NegotiableShareSum1 float64 `json:"negotiablessharesum1"`
			NegotiableShareSum2 float64 `json:"negotiablessharesum2"`
			DP                  float64 `json:"dp"`
		} `json:"results"`
	} `json:"data"`
}

// swWeekMonthResponse 申万周/月报表日期API响应
type swWeekMonthResponse struct {
	Data []struct {
		BargainDate string `json:"bargaindate"`
	} `json:"data"`
}

// createSWClient 创建申万研究客户端（跳过SSL验证）
func createSWClient() *resty.Client {
	return resty.New().
		SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetTransport(&http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}).
		SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")
}

// IndexHistSW 申万指数历史行情
//
// 数据源: https://www.swsresearch.com/institute_sw/allIndex/releasedIndex
//
// 参数:
//   - symbol: 指数代码，如 "801030"
//   - period: 周期，可选 "day", "week", "month"
//
// 返回:
//   - []SWIndexHist: 指数历史行情数据
//   - error: 错误信息
func IndexHistSW(symbol, period string) ([]SWIndexHist, error) {
	periodMap := map[string]string{
		"day":   "DAY",
		"week":  "WEEK",
		"month": "MONTH",
	}

	periodValue, ok := periodMap[period]
	if !ok {
		return nil, fmt.Errorf("无效的 period: %s，可选 day, week, month", period)
	}

	apiURL := "https://www.swsresearch.com/institute-sw/api/index_publish/trend/"
	client := createSWClient()

	resp, err := client.R().
		SetQueryParams(map[string]string{
			"swindexcode": symbol,
			"period":      periodValue,
		}).
		Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("请求申万指数历史行情失败: %w", err)
	}

	var apiResp swHistResponse
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, fmt.Errorf("解析申万指数历史行情响应失败: %w", err)
	}

	result := make([]SWIndexHist, 0, len(apiResp.Data))
	for _, item := range apiResp.Data {
		date, err := time.Parse("2006-01-02", item.BargainDate[:10])
		if err != nil {
			continue
		}

		result = append(result, SWIndexHist{
			Code:   item.SwIndexCode,
			Date:   date,
			Close:  item.CloseIndex,
			Open:   item.OpenIndex,
			High:   item.MaxIndex,
			Low:    item.MinIndex,
			Volume: item.BargainAmount,
			Amount: item.BargainSum,
		})
	}

	return result, nil
}

// IndexMinSW 申万指数分时数据
//
// 数据源: https://www.swsresearch.com/institute_sw/allIndex/releasedIndex
//
// 参数:
//   - symbol: 指数代码，如 "801001"
//
// 返回:
//   - []SWIndexMin: 指数分时数据
//   - error: 错误信息
func IndexMinSW(symbol string) ([]SWIndexMin, error) {
	apiURL := "https://www.swsresearch.com/institute-sw/api/index_publish/details/timelines/"
	client := createSWClient()

	resp, err := client.R().
		SetQueryParams(map[string]string{
			"swindexcode": symbol,
		}).
		Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("请求申万指数分时数据失败: %w", err)
	}

	var apiResp swMinResponse
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, fmt.Errorf("解析申万指数分时数据响应失败: %w", err)
	}

	result := make([]SWIndexMin, 0, len(apiResp.Data))
	for _, item := range apiResp.Data {
		date, err := time.Parse("2006-01-02", item.TradingDate[:10])
		if err != nil {
			continue
		}

		result = append(result, SWIndexMin{
			Code:  item.L1,
			Name:  item.L2,
			Price: item.L8,
			Date:  date,
			Time:  item.TradingTime,
		})
	}

	return result, nil
}

// IndexComponentSW 申万指数成份股
//
// 数据源: https://www.swsresearch.com/institute_sw/allIndex/releasedIndex
//
// 参数:
//   - symbol: 指数代码，如 "801001"
//
// 返回:
//   - []SWIndexComponent: 指数成份股数据
//   - error: 错误信息
func IndexComponentSW(symbol string) ([]SWIndexComponent, error) {
	apiURL := "https://www.swsresearch.com/institute-sw/api/index_publish/details/component_stocks/"
	client := createSWClient()

	resp, err := client.R().
		SetQueryParams(map[string]string{
			"swindexcode": symbol,
			"page":        "1",
			"page_size":   "10000",
		}).
		Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("请求申万指数成份股失败: %w", err)
	}

	var apiResp swComponentResponse
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, fmt.Errorf("解析申万指数成份股响应失败: %w", err)
	}

	result := make([]SWIndexComponent, 0, len(apiResp.Data.Results))
	for i, item := range apiResp.Data.Results {
		date, _ := time.Parse("2006-01-02", item.BeginningDate[:10])

		result = append(result, SWIndexComponent{
			Seq:          i + 1,
			StockCode:    item.StockCode,
			StockName:    item.StockName,
			Weight:       item.NewWeight,
			IncludedDate: date,
		})
	}

	return result, nil
}

// IndexRealtimeSW 申万指数实时行情
//
// 数据源: https://www.swsresearch.com/institute_sw/allIndex/releasedIndex
//
// 参数:
//   - symbol: 指数类型，可选 "市场表征", "一级行业", "二级行业", "风格指数", "大类风格指数", "金创指数"
//
// 返回:
//   - []SWIndexRealtime: 指数实时行情数据
//   - error: 错误信息
func IndexRealtimeSW(symbol string) ([]SWIndexRealtime, error) {
	// 大类风格指数和金创指数使用不同的API
	if symbol == "大类风格指数" || symbol == "金创指数" {
		return indexRealtimeSWSpecial(symbol)
	}

	apiURL := "https://www.swsresearch.com/institute-sw/api/index_publish/current/"
	client := createSWClient()

	// 获取第一页以获取总数
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"page":      "1",
			"page_size": "50",
			"indextype": symbol,
		}).
		Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("请求申万指数实时行情失败: %w", err)
	}

	var apiResp swRealtimeResponse
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, fmt.Errorf("解析申万指数实时行情响应失败: %w", err)
	}

	totalPages := (apiResp.Data.Count + 49) / 50
	result := make([]SWIndexRealtime, 0, apiResp.Data.Count)

	// 获取所有页
	for page := 1; page <= totalPages; page++ {
		if page > 1 {
			resp, err = client.R().
				SetQueryParams(map[string]string{
					"page":      fmt.Sprintf("%d", page),
					"page_size": "50",
					"indextype": symbol,
				}).
				Get(apiURL)
			if err != nil {
				continue
			}

			if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
				continue
			}
		}

		for _, item := range apiResp.Data.Results {
			result = append(result, SWIndexRealtime{
				IndexCode: item.SwIndexCode,
				IndexName: item.SwIndexName,
				PreClose:  item.LastClose,
				Open:      item.Open,
				Price:     item.CloseIndex,
				Amount:    item.BargainSum,
				Volume:    item.BargainAmount,
				High:      item.MaxIndex,
				Low:       item.MinIndex,
			})
		}
	}

	return result, nil
}

// indexRealtimeSWSpecial 申万特殊指数实时行情（大类风格指数、金创指数）
func indexRealtimeSWSpecial(symbol string) ([]SWIndexRealtime, error) {
	apiURL := "https://www.swsresearch.com/insWechatSw/dflgOrJcIndex/pageList"
	client := createSWClient()

	payload := map[string]interface{}{
		"pageNo":        1,
		"pageSize":      10,
		"indexTypeName": symbol,
		"sortField":     "",
		"rule":          "",
		"indexType":     1,
	}

	payloadBytes, _ := json.Marshal(payload)
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payloadBytes).
		Post(apiURL)
	if err != nil {
		return nil, fmt.Errorf("请求申万特殊指数实时行情失败: %w", err)
	}

	var apiResp struct {
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
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, fmt.Errorf("解析申万特殊指数实时行情响应失败: %w", err)
	}

	result := make([]SWIndexRealtime, 0, len(apiResp.Data.List))
	for _, item := range apiResp.Data.List {
		result = append(result, SWIndexRealtime{
			IndexCode: item.SwIndexCode,
			IndexName: item.SwIndexName,
			PreClose:  item.LastCloseIndex,
		})
	}

	return result, nil
}

// IndexAnalysisDailySW 申万指数分析-日报告
//
// 数据源: https://www.swsresearch.com/institute_sw/allIndex/analysisIndex
//
// 参数:
//   - symbol: 指数类型，可选 "市场表征", "一级行业", "二级行业", "风格指数"
//   - startDate: 开始日期，格式 "20221103"
//   - endDate: 结束日期，格式 "20221103"
//
// 返回:
//   - []SWIndexAnalysis: 指数分析数据
//   - error: 错误信息
func IndexAnalysisDailySW(symbol, startDate, endDate string) ([]SWIndexAnalysis, error) {
	apiURL := "https://www.swsresearch.com/institute-sw/api/index_analysis/index_analysis_report/"
	client := createSWClient()

	startFormatted := fmt.Sprintf("%s-%s-%s", startDate[:4], startDate[4:6], startDate[6:])
	endFormatted := fmt.Sprintf("%s-%s-%s", endDate[:4], endDate[4:6], endDate[6:])

	// 获取第一页以获取总数
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"page":        "1",
			"page_size":   "50",
			"index_type":  symbol,
			"start_date":  startFormatted,
			"end_date":    endFormatted,
			"type":        "DAY",
			"swindexcode": "all",
		}).
		Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("请求申万指数分析日报告失败: %w", err)
	}

	var apiResp swAnalysisResponse
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, fmt.Errorf("解析申万指数分析日报告响应失败: %w", err)
	}

	totalPages := (apiResp.Data.Count + 49) / 50
	result := make([]SWIndexAnalysis, 0, apiResp.Data.Count)

	// 获取所有页
	for page := 1; page <= totalPages; page++ {
		if page > 1 {
			resp, err = client.R().
				SetQueryParams(map[string]string{
					"page":        fmt.Sprintf("%d", page),
					"page_size":   "50",
					"index_type":  symbol,
					"start_date":  startFormatted,
					"end_date":    endFormatted,
					"type":        "DAY",
					"swindexcode": "all",
				}).
				Get(apiURL)
			if err != nil {
				continue
			}

			if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
				continue
			}
		}

		for _, item := range apiResp.Data.Results {
			date, _ := time.Parse("2006-01-02", item.BargainDate[:10])
			result = append(result, SWIndexAnalysis{
				IndexCode:    item.SwIndexCode,
				IndexName:    item.SwIndexName,
				Date:         date,
				CloseIndex:   item.CloseIndex,
				Volume:       item.BargainAmount,
				ChangePct:    item.Markup,
				TurnoverRate: item.TurnoverRate,
				PE:           item.PE,
				PB:           item.PB,
				AvgPrice:     item.MeanPrice,
				AmountRatio:  item.BargainSumRate,
				FloatCap:     item.NegotiableShareSum1,
				AvgFloatCap:  item.NegotiableShareSum2,
				DividendRate: item.DP,
			})
		}
	}

	return result, nil
}

// IndexAnalysisWeekMonthSW 申万周/月报表日期序列
//
// 数据源: https://www.swsresearch.com/institute_sw/allIndex/analysisIndex
//
// 参数:
//   - symbol: 报表类型，可选 "week", "month"
//
// 返回:
//   - []SWIndexWeekMonthDate: 日期序列
//   - error: 错误信息
func IndexAnalysisWeekMonthSW(symbol string) ([]SWIndexWeekMonthDate, error) {
	apiURL := "https://www.swsresearch.com/institute-sw/api/index_analysis/week_month_datetime/"
	client := createSWClient()

	typeValue := "WEEK"
	if symbol == "month" {
		typeValue = "MONTH"
	}

	resp, err := client.R().
		SetQueryParams(map[string]string{
			"type": typeValue,
		}).
		Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("请求申万周/月报表日期序列失败: %w", err)
	}

	var apiResp swWeekMonthResponse
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, fmt.Errorf("解析申万周/月报表日期序列响应失败: %w", err)
	}

	result := make([]SWIndexWeekMonthDate, 0, len(apiResp.Data))
	for _, item := range apiResp.Data {
		date, err := time.Parse("2006-01-02", item.BargainDate[:10])
		if err != nil {
			continue
		}
		result = append(result, SWIndexWeekMonthDate{Date: date})
	}

	return result, nil
}

// IndexAnalysisWeeklySW 申万指数分析-周报告
//
// 数据源: https://www.swsresearch.com/institute_sw/allIndex/analysisIndex
//
// 参数:
//   - symbol: 指数类型，可选 "市场表征", "一级行业", "二级行业", "风格指数"
//   - date: 查询日期，格式 "20221104"
//
// 返回:
//   - []SWIndexAnalysis: 指数分析数据
//   - error: 错误信息
func IndexAnalysisWeeklySW(symbol, date string) ([]SWIndexAnalysis, error) {
	return indexAnalysisPeriodSW(symbol, date, "WEEK")
}

// IndexAnalysisMonthlySW 申万指数分析-月报告
//
// 数据源: https://www.swsresearch.com/institute_sw/allIndex/analysisIndex
//
// 参数:
//   - symbol: 指数类型，可选 "市场表征", "一级行业", "二级行业", "风格指数"
//   - date: 查询日期，格式 "20221031"
//
// 返回:
//   - []SWIndexAnalysis: 指数分析数据
//   - error: 错误信息
func IndexAnalysisMonthlySW(symbol, date string) ([]SWIndexAnalysis, error) {
	return indexAnalysisPeriodSW(symbol, date, "MONTH")
}

// indexAnalysisPeriodSW 申万指数分析-周/月报告
func indexAnalysisPeriodSW(symbol, date, periodType string) ([]SWIndexAnalysis, error) {
	apiURL := "https://www.swsresearch.com/institute-sw/api/index_analysis/index_analysis_reports/"
	client := createSWClient()

	dateFormatted := fmt.Sprintf("%s-%s-%s", date[:4], date[4:6], date[6:])

	// 获取第一页以获取总数
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"page":        "1",
			"page_size":   "50",
			"index_type":  symbol,
			"bargaindate": dateFormatted,
			"type":        periodType,
			"swindexcode": "all",
		}).
		Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("请求申万指数分析周/月报告失败: %w", err)
	}

	var apiResp swAnalysisResponse
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, fmt.Errorf("解析申万指数分析周/月报告响应失败: %w", err)
	}

	totalPages := (apiResp.Data.Count + 49) / 50
	result := make([]SWIndexAnalysis, 0, apiResp.Data.Count)

	// 获取所有页
	for page := 1; page <= totalPages; page++ {
		if page > 1 {
			resp, err = client.R().
				SetQueryParams(map[string]string{
					"page":        fmt.Sprintf("%d", page),
					"page_size":   "50",
					"index_type":  symbol,
					"bargaindate": dateFormatted,
					"type":        periodType,
					"swindexcode": "all",
				}).
				Get(apiURL)
			if err != nil {
				continue
			}

			if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
				continue
			}
		}

		for _, item := range apiResp.Data.Results {
			dateVal, _ := time.Parse("2006-01-02", item.BargainDate[:10])
			result = append(result, SWIndexAnalysis{
				IndexCode:    item.SwIndexCode,
				IndexName:    item.SwIndexName,
				Date:         dateVal,
				CloseIndex:   item.CloseIndex,
				Volume:       item.BargainAmount,
				ChangePct:    item.Markup,
				TurnoverRate: item.TurnoverRate,
				PE:           item.PE,
				PB:           item.PB,
				AvgPrice:     item.MeanPrice,
				AmountRatio:  item.BargainSumRate,
				FloatCap:     item.NegotiableShareSum1,
				AvgFloatCap:  item.NegotiableShareSum2,
				DividendRate: item.DP,
			})
		}
	}

	return result, nil
}
