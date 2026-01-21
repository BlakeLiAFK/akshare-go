package stock_feature

import (
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// StockGdfxHoldingItem 股东持股统计项
type StockGdfxHoldingItem struct {
	Index           int     `json:"index"`            // 序号
	HolderName      string  `json:"holder_name"`      // 股东名称
	HolderType      string  `json:"holder_type"`      // 股东类型
	StatisticsCount int64   `json:"statistics_count"` // 统计次数
	Avg10DayChange  float64 `json:"avg_10day_change"` // 10日平均涨幅
	Max10DayChange  float64 `json:"max_10day_change"` // 10日最大涨幅
	Min10DayChange  float64 `json:"min_10day_change"` // 10日最小涨幅
	Avg30DayChange  float64 `json:"avg_30day_change"` // 30日平均涨幅
	Max30DayChange  float64 `json:"max_30day_change"` // 30日最大涨幅
	Min30DayChange  float64 `json:"min_30day_change"` // 30日最小涨幅
	Avg60DayChange  float64 `json:"avg_60day_change"` // 60日平均涨幅
	Max60DayChange  float64 `json:"max_60day_change"` // 60日最大涨幅
	Min60DayChange  float64 `json:"min_60day_change"` // 60日最小涨幅
	HoldingStocks   string  `json:"holding_stocks"`   // 持有个股
}

// StockGdfxFreeHoldingStatisticsEm 东方财富网-数据中心-股东分析-股东持股统计-十大流通股东
// date: 报告期，格式 "20210630"
func StockGdfxFreeHoldingStatisticsEm(date string) ([]StockGdfxHoldingItem, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	formattedDate := date[:4] + "-" + date[4:6] + "-" + date[6:]

	var allItems []StockGdfxHoldingItem
	page := 1

	for {
		params := map[string]string{
			"sortColumns": "STATISTICS_TIMES,COOPERATION_HOLDER_MARK",
			"sortTypes":   "-1,-1",
			"pageSize":    "500",
			"pageNumber":  fmt.Sprintf("%d", page),
			"reportName":  "RPT_COOPFREEHOLDERS_ANALYSIS",
			"columns":     "ALL",
			"source":      "WEB",
			"client":      "WEB",
			"filter":      fmt.Sprintf(`(HOLDNUM_CHANGE_TYPE="001")(END_DATE='%s')`, formattedDate),
		}

		resp, err := utils.Get(url, params)
		if err != nil {
			return nil, fmt.Errorf("获取十大流通股东统计失败: %w", err)
		}

		result := gjson.ParseBytes(resp.Body())
		totalPages := result.Get("result.pages").Int()
		dataArr := result.Get("result.data").Array()

		for _, item := range dataArr {
			allItems = append(allItems, StockGdfxHoldingItem{
				Index:           len(allItems) + 1,
				HolderName:      item.Get("COOPERATION_HOLDER_NAME").String(),
				HolderType:      item.Get("HOLDER_TYPE").String(),
				StatisticsCount: item.Get("STATISTICS_TIMES").Int(),
				Avg10DayChange:  item.Get("AVERAGE_10_ADJCHRATE").Float(),
				Max10DayChange:  item.Get("MAX_10_ADJCHRATE").Float(),
				Min10DayChange:  item.Get("MIN_10_ADJCHRATE").Float(),
				Avg30DayChange:  item.Get("AVERAGE_30_ADJCHRATE").Float(),
				Max30DayChange:  item.Get("MAX_30_ADJCHRATE").Float(),
				Min30DayChange:  item.Get("MIN_30_ADJCHRATE").Float(),
				Avg60DayChange:  item.Get("AVERAGE_60_ADJCHRATE").Float(),
				Max60DayChange:  item.Get("MAX_60_ADJCHRATE").Float(),
				Min60DayChange:  item.Get("MIN_60_ADJCHRATE").Float(),
				HoldingStocks:   item.Get("HOLD_SECURITY_NAMES").String(),
			})
		}

		if int64(page) >= totalPages || totalPages == 0 {
			break
		}
		page++
	}

	return allItems, nil
}

// StockGdfxHoldingStatisticsEm 东方财富网-数据中心-股东分析-股东持股统计-十大股东
// date: 报告期，格式 "20210930"
func StockGdfxHoldingStatisticsEm(date string) ([]StockGdfxHoldingItem, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	formattedDate := date[:4] + "-" + date[4:6] + "-" + date[6:]

	var allItems []StockGdfxHoldingItem
	page := 1

	for {
		params := map[string]string{
			"sortColumns": "STATISTICS_TIMES,COOPERATION_HOLDER_MARK",
			"sortTypes":   "-1,-1",
			"pageSize":    "500",
			"pageNumber":  fmt.Sprintf("%d", page),
			"reportName":  "RPT_COOPHOLDERS_ANALYSIS",
			"columns":     "ALL",
			"source":      "WEB",
			"client":      "WEB",
			"filter":      fmt.Sprintf(`(HOLDNUM_CHANGE_TYPE="001")(END_DATE='%s')`, formattedDate),
		}

		resp, err := utils.Get(url, params)
		if err != nil {
			return nil, fmt.Errorf("获取十大股东统计失败: %w", err)
		}

		result := gjson.ParseBytes(resp.Body())
		totalPages := result.Get("result.pages").Int()
		dataArr := result.Get("result.data").Array()

		for _, item := range dataArr {
			allItems = append(allItems, StockGdfxHoldingItem{
				Index:           len(allItems) + 1,
				HolderName:      item.Get("COOPERATION_HOLDER_NAME").String(),
				HolderType:      item.Get("HOLDER_TYPE").String(),
				StatisticsCount: item.Get("STATISTICS_TIMES").Int(),
				Avg10DayChange:  item.Get("AVERAGE_10_ADJCHRATE").Float(),
				Max10DayChange:  item.Get("MAX_10_ADJCHRATE").Float(),
				Min10DayChange:  item.Get("MIN_10_ADJCHRATE").Float(),
				Avg30DayChange:  item.Get("AVERAGE_30_ADJCHRATE").Float(),
				Max30DayChange:  item.Get("MAX_30_ADJCHRATE").Float(),
				Min30DayChange:  item.Get("MIN_30_ADJCHRATE").Float(),
				Avg60DayChange:  item.Get("AVERAGE_60_ADJCHRATE").Float(),
				Max60DayChange:  item.Get("MAX_60_ADJCHRATE").Float(),
				Min60DayChange:  item.Get("MIN_60_ADJCHRATE").Float(),
				HoldingStocks:   item.Get("HOLD_SECURITY_NAMES").String(),
			})
		}

		if int64(page) >= totalPages || totalPages == 0 {
			break
		}
		page++
	}

	return allItems, nil
}
