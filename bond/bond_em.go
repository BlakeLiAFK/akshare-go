package bond

import (
	"fmt"
	"sort"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// BondZhUsRateItem 中美国债收益率项
type BondZhUsRateItem struct {
	Date         string  `json:"date"`            // 日期
	CnYield2Y    float64 `json:"cn_yield_2y"`     // 中国国债收益率2年
	CnYield5Y    float64 `json:"cn_yield_5y"`     // 中国国债收益率5年
	CnYield10Y   float64 `json:"cn_yield_10y"`    // 中国国债收益率10年
	CnYield30Y   float64 `json:"cn_yield_30y"`    // 中国国债收益率30年
	CnYield10Y2Y float64 `json:"cn_yield_10y_2y"` // 中国国债收益率10年-2年
	CnGdpYoy     float64 `json:"cn_gdp_yoy"`      // 中国GDP年增率
	UsYield2Y    float64 `json:"us_yield_2y"`     // 美国国债收益率2年
	UsYield5Y    float64 `json:"us_yield_5y"`     // 美国国债收益率5年
	UsYield10Y   float64 `json:"us_yield_10y"`    // 美国国债收益率10年
	UsYield30Y   float64 `json:"us_yield_30y"`    // 美国国债收益率30年
	UsYield10Y2Y float64 `json:"us_yield_10y_2y"` // 美国国债收益率10年-2年
	UsGdpYoy     float64 `json:"us_gdp_yoy"`      // 美国GDP年增率
}

// BondZhUsRate 东方财富网-数据中心-经济数据-中美国债收益率
// startDate: 开始统计时间，格式 "19901219"
func BondZhUsRate(startDate string) ([]BondZhUsRateItem, error) {
	url := "https://datacenter.eastmoney.com/api/data/get"

	var allItems []BondZhUsRateItem
	page := 1

	for {
		params := map[string]string{
			"type":    "RPTA_WEB_TREASURYYIELD",
			"sty":     "ALL",
			"st":      "SOLAR_DATE",
			"sr":      "-1",
			"token":   "894050c76af8597a853f5b408b759f5d",
			"p":       fmt.Sprintf("%d", page),
			"ps":      "500",
			"pageNo":  fmt.Sprintf("%d", page),
			"pageNum": fmt.Sprintf("%d", page),
		}

		resp, err := utils.Get(url, params)
		if err != nil {
			return nil, fmt.Errorf("获取中美国债收益率失败: %w", err)
		}

		result := gjson.ParseBytes(resp.Body())
		totalPages := result.Get("result.pages").Int()
		dataArr := result.Get("result.data").Array()

		for _, item := range dataArr {
			dateStr := strings.Split(item.Get("SOLAR_DATE").String(), " ")[0]
			allItems = append(allItems, BondZhUsRateItem{
				Date:         dateStr,
				CnYield2Y:    item.Get("EMM00588704").Float(),
				CnYield5Y:    item.Get("EMM00166462").Float(),
				CnYield10Y:   item.Get("EMM00166466").Float(),
				CnYield30Y:   item.Get("EMM00166469").Float(),
				CnYield10Y2Y: item.Get("EMM01276014").Float(),
				CnGdpYoy:     item.Get("EMM00000024").Float(),
				UsYield2Y:    item.Get("EMG00001306").Float(),
				UsYield5Y:    item.Get("EMG00001308").Float(),
				UsYield10Y:   item.Get("EMG00001310").Float(),
				UsYield30Y:   item.Get("EMG00001312").Float(),
				UsYield10Y2Y: item.Get("EMG01339436").Float(),
				UsGdpYoy:     item.Get("EMG00159635").Float(),
			})
		}

		if int64(page) >= totalPages || totalPages == 0 {
			break
		}
		page++
	}

	// 按日期排序
	sort.Slice(allItems, func(i, j int) bool {
		return allItems[i].Date < allItems[j].Date
	})

	// 过滤起始日期之后的数据
	if startDate != "" {
		formattedStart := startDate[:4] + "-" + startDate[4:6] + "-" + startDate[6:]
		var filtered []BondZhUsRateItem
		for _, item := range allItems {
			if item.Date >= formattedStart {
				filtered = append(filtered, item)
			}
		}
		allItems = filtered
	}

	return allItems, nil
}
