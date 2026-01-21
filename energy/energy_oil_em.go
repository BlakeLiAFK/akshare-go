package energy

import (
	"fmt"
	"sort"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// EnergyOilHistItem 汽柴油历史调价信息项
type EnergyOilHistItem struct {
	Date           string  `json:"date"`            // 调整日期
	GasolinePrice  float64 `json:"gasoline_price"`  // 汽油价格
	DieselPrice    float64 `json:"diesel_price"`    // 柴油价格
	GasolineChange float64 `json:"gasoline_change"` // 汽油涨跌
	DieselChange   float64 `json:"diesel_change"`   // 柴油涨跌
}

// EnergyOilDetailItem 地区油价详情项
type EnergyOilDetailItem struct {
	Date   string  `json:"date"`   // 日期
	Region string  `json:"region"` // 地区
	V0     float64 `json:"v_0"`    // 0号柴油价格
	V92    float64 `json:"v_92"`   // 92号汽油价格
	V95    float64 `json:"v_95"`   // 95号汽油价格
	V89    float64 `json:"v_89"`   // 89号汽油价格
}

// EnergyOilHist 汽柴油历史调价信息
func EnergyOilHist() ([]EnergyOilHistItem, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"reportName":  "RPTA_WEB_YJ_BD",
		"columns":     "ALL",
		"sortColumns": "dim_date",
		"sortTypes":   "-1",
		"token":       "894050c76af8597a853f5b408b759f5d",
		"pageNumber":  "1",
		"pageSize":    "1000",
		"source":      "WEB",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("获取汽柴油历史调价信息失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	dataArr := result.Get("result.data").Array()

	var items []EnergyOilHistItem
	for _, item := range dataArr {
		items = append(items, EnergyOilHistItem{
			Date:           strings.Split(item.Get("dim_date").String(), " ")[0],
			GasolinePrice:  item.Get("V_92").Float(),
			DieselPrice:    item.Get("V_0").Float(),
			GasolineChange: item.Get("ZDE_92").Float(),
			DieselChange:   item.Get("ZDE_0").Float(),
		})
	}

	// 按日期排序
	sort.Slice(items, func(i, j int) bool {
		return items[i].Date < items[j].Date
	})

	return items, nil
}

// EnergyOilDetail 全国各地区的汽油和柴油油价
// date: 日期，格式 "20220517"
func EnergyOilDetail(date string) ([]EnergyOilDetailItem, error) {
	formattedDate := date[:4] + "-" + date[4:6] + "-" + date[6:]
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"reportName":  "RPTA_WEB_YJ_JH",
		"columns":     "ALL",
		"filter":      fmt.Sprintf(`(dim_date='%s')`, formattedDate),
		"sortColumns": "cityname",
		"sortTypes":   "1",
		"token":       "894050c76af8597a853f5b408b759f5d",
		"pageNumber":  "1",
		"pageSize":    "1000",
		"source":      "WEB",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("获取地区油价详情失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	dataArr := result.Get("result.data").Array()

	var items []EnergyOilDetailItem
	for _, item := range dataArr {
		items = append(items, EnergyOilDetailItem{
			Date:   strings.Split(item.Get("dim_date").String(), " ")[0],
			Region: item.Get("cityname").String(),
			V0:     item.Get("V_0").Float(),
			V92:    item.Get("V_92").Float(),
			V95:    item.Get("V_95").Float(),
			V89:    item.Get("V_89").Float(),
		})
	}

	return items, nil
}
