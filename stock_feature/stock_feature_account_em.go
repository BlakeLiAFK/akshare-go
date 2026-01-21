package stock_feature

import (
	"fmt"
	"sort"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// StockAccountStatisticsEm 东方财富网-数据中心-特色数据-股票账户统计
func StockAccountStatisticsEm() ([]StockAccountStatisticsItem, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"reportName":  "RPT_STOCK_OPEN_DATA",
		"columns":     "ALL",
		"pageSize":    "500",
		"sortColumns": "STATISTICS_DATE",
		"sortTypes":   "-1",
		"source":      "WEB",
		"client":      "WEB",
		"p":           "1",
		"pageNo":      "1",
		"pageNum":     "1",
		"pageNumber":  "1",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("获取股票账户统计失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	dataArr := result.Get("result.data").Array()

	var items []StockAccountStatisticsItem
	for _, item := range dataArr {
		items = append(items, StockAccountStatisticsItem{
			DataDate:           item.Get("STATISTICS_DATE").String(),
			NewInvestorCount:   item.Get("ADD_INVESTOR").Float(),
			NewInvestorMoM:     item.Get("ADD_INVESTOR_QOQ").Float(),
			NewInvestorYoY:     item.Get("ADD_INVESTOR_YOY").Float(),
			TotalInvestor:      item.Get("END_INVESTOR").Float(),
			AShareAccount:      item.Get("END_A_INVESTOR").Float(),
			BShareAccount:      item.Get("END_B_INVESTOR").Float(),
			TotalMarketCap:     item.Get("TOTAL_MARKET_CAP").Float(),
			AvgMarketCap:       item.Get("AVG_MARKET_CAP").Float(),
			ShangHaiIndexClose: item.Get("SH_INDEX").Float(),
			ShangHaiIndexPct:   item.Get("SH_INDEX_CHG").Float(),
		})
	}

	// 按日期排序
	sort.Slice(items, func(i, j int) bool {
		return items[i].DataDate < items[j].DataDate
	})

	return items, nil
}
