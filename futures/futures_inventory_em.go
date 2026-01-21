package futures

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
)

// FuturesInventoryEM 期货库存数据
type FuturesInventoryEM struct {
	Date      time.Time `json:"date"`      // 日期
	Inventory float64   `json:"inventory"` // 库存
	Change    float64   `json:"change"`    // 增减
}

// inventoryCodeResponse 期货品种代码API响应
type inventoryCodeResponse struct {
	Result struct {
		Data []struct {
			TradeCode string `json:"TRADE_CODE"`
			TradeType string `json:"TRADE_TYPE"`
		} `json:"data"`
	} `json:"result"`
}

// inventoryDataResponse 期货库存数据API响应
type inventoryDataResponse struct {
	Result struct {
		Data []struct {
			TradeDate    string  `json:"TRADE_DATE"`
			OnWarrantNum float64 `json:"ON_WARRANT_NUM"`
			AddChange    float64 `json:"ADDCHANGE"`
		} `json:"data"`
	} `json:"result"`
}

// FuturesInventoryEMFunc 东方财富-期货库存数据
//
// 数据源: https://data.eastmoney.com/ifdata/kcsj.html
//
// 参数:
//   - symbol: 品种代码或中文名称，如 "a", "豆一"
//
// 返回:
//   - []FuturesInventoryEM: 期货库存数据
//   - error: 错误信息
func FuturesInventoryEMFunc(symbol string) ([]FuturesInventoryEM, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"

	// 首先获取品种代码映射
	codeParams := map[string]string{
		"reportName": "RPT_FUTU_POSITIONCODE",
		"columns":    "TRADE_MARKET_CODE,TRADE_CODE,TRADE_TYPE",
		"filter":     `(IS_MAINCODE="1")`,
		"pageNumber": "1",
		"pageSize":   "500",
		"source":     "WEB",
		"client":     "WEB",
	}

	resp, err := utils.Get(url, codeParams)
	if err != nil {
		return nil, fmt.Errorf("请求期货品种代码失败: %w", err)
	}

	var codeResp inventoryCodeResponse
	if err := json.Unmarshal(resp.Body(), &codeResp); err != nil {
		return nil, fmt.Errorf("解析期货品种代码响应失败: %w", err)
	}

	// 构建品种名称到代码的映射
	symbolDict := make(map[string]string)
	for _, item := range codeResp.Result.Data {
		symbolDict[item.TradeType] = item.TradeCode
	}

	// 确定产品ID
	var productID string
	if code, ok := symbolDict[symbol]; ok {
		productID = code
	} else if code, ok := FuturesInventoryEMSymbolDict[symbol]; ok {
		productID = code
	} else {
		return nil, fmt.Errorf("无效的 symbol: %s", symbol)
	}

	// 获取库存数据
	dataParams := map[string]string{
		"reportName":  "RPT_FUTU_STOCKDATA",
		"columns":     "SECURITY_CODE,TRADE_DATE,ON_WARRANT_NUM,ADDCHANGE",
		"filter":      fmt.Sprintf(`(SECURITY_CODE="%s")(TRADE_DATE>='2020-10-28')`, productID),
		"pageNumber":  "1",
		"pageSize":    "500",
		"sortTypes":   "-1",
		"sortColumns": "TRADE_DATE",
		"source":      "WEB",
		"client":      "WEB",
	}

	resp, err = utils.Get(url, dataParams)
	if err != nil {
		return nil, fmt.Errorf("请求期货库存数据失败: %w", err)
	}

	var dataResp inventoryDataResponse
	if err := json.Unmarshal(resp.Body(), &dataResp); err != nil {
		return nil, fmt.Errorf("解析期货库存数据响应失败: %w", err)
	}

	result := make([]FuturesInventoryEM, 0, len(dataResp.Result.Data))
	for _, item := range dataResp.Result.Data {
		date, err := time.Parse("2006-01-02 15:04:05", item.TradeDate)
		if err != nil {
			date, _ = time.Parse("2006-01-02", item.TradeDate[:10])
		}

		result = append(result, FuturesInventoryEM{
			Date:      date,
			Inventory: item.OnWarrantNum,
			Change:    item.AddChange,
		})
	}

	// 按日期排序
	sort.Slice(result, func(i, j int) bool {
		return result[i].Date.Before(result[j].Date)
	})

	return result, nil
}
