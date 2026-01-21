package stock_fundamental

import (
	"fmt"
	"sort"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// StockProfitForecastEm 东方财富网-数据中心-研究报告-盈利预测
//
// 获取东方财富网的盈利预测数据，包含代码、名称、研报数、机构投资评级、预测每股收益等信息
//
// 参数:
//   - symbol: 行业板块，默认为空获取全部。可以通过 stock_board_industry_name_em() 获取行业板块列表
//
// 返回:
//   - []StockProfitForecastEmItem: 盈利预测数据列表
//   - error: 错误信息
//
// 示例:
//
//	// 获取全部盈利预测
//	forecast, err := stock_fundamental.StockProfitForecastEm("")
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	// 获取航运港口板块的盈利预测
//	forecast, err := stock_fundamental.StockProfitForecastEm("航运港口")
func StockProfitForecastEm(symbol string) ([]StockProfitForecastEmItem, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"

	params := map[string]string{
		"reportName":  "RPT_WEB_RESPREDICT",
		"columns":     "WEB_RESPREDICT",
		"pageNumber":  "1",
		"pageSize":    "500",
		"sortTypes":   "-1",
		"sortColumns": "RATING_ORG_NUM",
		"p":           "1",
		"pageNo":      "1",
		"pageNum":     "1",
	}

	if symbol != "" {
		params["filter"] = fmt.Sprintf(`(INDUSTRY_BOARD="%s")`, symbol)
	}

	headers := map[string]string{
		"Referer": "https://data.eastmoney.com/report/profitforecast.jshtml",
	}

	// 首次请求获取总页数
	resp, err := utils.GetWithHeaders(url, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求盈利预测失败: %w", err)
	}

	json := gjson.ParseBytes(resp.Body())
	pageNum := json.Get("result.pages").Int()
	if pageNum == 0 {
		pageNum = 1
	}

	var allItems []StockProfitForecastEmItem
	year1Count := make(map[string]int)
	year2Count := make(map[string]int)
	year3Count := make(map[string]int)
	year4Count := make(map[string]int)

	// 分页获取所有数据
	for page := 1; page <= int(pageNum); page++ {
		params["pageNumber"] = fmt.Sprintf("%d", page)
		params["p"] = fmt.Sprintf("%d", page)
		params["pageNo"] = fmt.Sprintf("%d", page)
		params["pageNum"] = fmt.Sprintf("%d", page)

		resp, err := utils.GetWithHeaders(url, params, headers)
		if err != nil {
			return nil, fmt.Errorf("请求第%d页失败: %w", page, err)
		}

		json := gjson.ParseBytes(resp.Body())
		dataArray := json.Get("result.data").Array()

		for _, data := range dataArray {
			// 统计年份
			year1 := data.Get("YEAR1").String()
			year2 := data.Get("YEAR2").String()
			year3 := data.Get("YEAR3").String()
			year4 := data.Get("YEAR4").String()

			if year1 != "" {
				year1Count[year1]++
			}
			if year2 != "" {
				year2Count[year2]++
			}
			if year3 != "" {
				year3Count[year3]++
			}
			if year4 != "" {
				year4Count[year4]++
			}

			item := StockProfitForecastEmItem{
				Code:          data.Get("SECURITY_CODE").String(),
				Name:          data.Get("SECURITY_NAME_ABBR").String(),
				ReportCount:   int(data.Get("RATING_ORG_NUM").Int()),
				RatingBuy:     int(data.Get("BUY_NUM").Int()),
				RatingAdd:     int(data.Get("ADD_NUM").Int()),
				RatingNeutral: int(data.Get("NEUTRAL_NUM").Int()),
				RatingReduce:  int(data.Get("REDUCE_NUM").Int()),
				RatingSell:    int(data.Get("SELL_NUM").Int()),
				Year1Forecast: utils.MustFloat64(data.Get("AVG_FPS1").String()),
				Year2Forecast: utils.MustFloat64(data.Get("AVG_FPS2").String()),
				Year3Forecast: utils.MustFloat64(data.Get("AVG_FPS3").String()),
				Year4Forecast: utils.MustFloat64(data.Get("AVG_FPS4").String()),
			}
			allItems = append(allItems, item)
		}
	}

	// 获取最常见的年份
	year1 := getMostCommonYear(year1Count)
	year2 := getMostCommonYear(year2Count)
	year3 := getMostCommonYear(year3Count)
	year4 := getMostCommonYear(year4Count)

	// 按研报数排序
	sort.Slice(allItems, func(i, j int) bool {
		return allItems[i].ReportCount > allItems[j].ReportCount
	})

	// 设置序号
	for i := range allItems {
		allItems[i].Index = i + 1
	}

	// 打印年份信息（用于调试）
	if year1 != "" {
		fmt.Printf("年份列: %s, %s, %s, %s\n", year1, year2, year3, year4)
	}

	return allItems, nil
}

// getMostCommonYear 获取出现次数最多的年份
func getMostCommonYear(yearCount map[string]int) string {
	if len(yearCount) == 0 {
		return ""
	}

	maxCount := 0
	var year string
	for y, count := range yearCount {
		if count > maxCount {
			maxCount = count
			year = y
		}
	}
	return year
}
