package futures

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
)

// FuturesComexInventory COMEX库存数据
type FuturesComexInventory struct {
	Seq        int       `json:"seq"`         // 序号
	Date       time.Time `json:"date"`        // 日期
	StorageTon float64   `json:"storage_ton"` // 库存量-吨
	StorageOz  float64   `json:"storage_oz"`  // 库存量-盎司
}

// comexInventoryResponse COMEX库存API响应
type comexInventoryResponse struct {
	Result struct {
		Pages int `json:"pages"`
		Data  []struct {
			ReportDate string  `json:"REPORT_DATE"`
			StorageTon float64 `json:"STORAGE_TON"`
			StorageOz  float64 `json:"STORAGE_OUNCE"`
		} `json:"data"`
	} `json:"result"`
}

// FuturesComexInventoryFunc 东方财富-COMEX库存数据
//
// 数据源: https://data.eastmoney.com/pmetal/comex/by.html
//
// 参数:
//   - symbol: 品种，可选 "黄金", "白银"
//
// 返回:
//   - []FuturesComexInventory: COMEX库存数据
//   - error: 错误信息
func FuturesComexInventoryFunc(symbol string) ([]FuturesComexInventory, error) {
	symbolMap := map[string]string{
		"黄金": "EMI00069026",
		"白银": "EMI00069027",
	}

	indicatorID, ok := symbolMap[symbol]
	if !ok {
		return nil, fmt.Errorf("无效的 symbol: %s, 可选 黄金, 白银", symbol)
	}

	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"

	var allData []FuturesComexInventory
	page := 1

	for {
		params := map[string]string{
			"sortColumns":  "REPORT_DATE",
			"sortTypes":    "-1",
			"pageSize":     "500",
			"pageNumber":   fmt.Sprintf("%d", page),
			"reportName":   "RPT_FUTUOPT_GOLDSIL",
			"columns":      "ALL",
			"quoteColumns": "",
			"source":       "WEB",
			"client":       "WEB",
			"filter":       fmt.Sprintf(`(INDICATOR_ID1="%s")(@STORAGE_TON<>"NULL")`, indicatorID),
		}

		resp, err := utils.Get(url, params)
		if err != nil {
			return nil, fmt.Errorf("请求COMEX库存数据失败: %w", err)
		}

		var apiResp comexInventoryResponse
		if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
			return nil, fmt.Errorf("解析COMEX库存数据响应失败: %w", err)
		}

		if len(apiResp.Result.Data) == 0 {
			break
		}

		for _, item := range apiResp.Result.Data {
			date, err := time.Parse("2006-01-02 15:04:05", item.ReportDate)
			if err != nil {
				date, _ = time.Parse("2006-01-02", item.ReportDate[:10])
			}

			allData = append(allData, FuturesComexInventory{
				Date:       date,
				StorageTon: item.StorageTon,
				StorageOz:  item.StorageOz,
			})
		}

		if page >= apiResp.Result.Pages {
			break
		}
		page++
	}

	// 按日期排序
	sort.Slice(allData, func(i, j int) bool {
		return allData[i].Date.Before(allData[j].Date)
	})

	// 添加序号
	for i := range allData {
		allData[i].Seq = i + 1
	}

	return allData, nil
}
