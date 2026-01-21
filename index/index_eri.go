package index

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
)

// EriIndex 浙江省排污权交易指数
type EriIndex struct {
	Date       time.Time `json:"date"`        // 日期
	TradeIndex float64   `json:"trade_index"` // 交易指数
	Volume     float64   `json:"volume"`      // 成交量
	Amount     float64   `json:"amount"`      // 成交额
}

// eriIndexDataResponse 排污权指数数据响应
type eriIndexDataResponse struct {
	Data []struct {
		IndexValue float64 `json:"indexValue"`
		Stage      struct {
			PublishTime string `json:"publishTime"`
		} `json:"stage"`
	} `json:"data"`
}

// eriStatisticsResponse 排污权统计数据响应
type eriStatisticsResponse struct {
	Data []struct {
		TotalQuantity float64 `json:"totalQuantity"`
		TotalCost     float64 `json:"totalCost"`
	} `json:"data"`
}

// IndexEri 浙江省排污权交易指数
//
// 数据源: https://zs.zjpwq.net
//
// 参数:
//   - symbol: 周期类型，可选 "月度", "季度"
//
// 返回:
//   - []EriIndex: 排污权交易指数数据
//   - error: 错误信息
func IndexEri(symbol string) ([]EriIndex, error) {
	symbolMap := map[string]string{
		"月度": "MONTH",
		"季度": "QUARTER",
	}

	cycle, ok := symbolMap[symbol]
	if !ok {
		return nil, fmt.Errorf("无效的 symbol: %s, 可选值: 月度, 季度", symbol)
	}

	params := map[string]string{
		"cycle":    cycle,
		"regionId": "1",
		"structId": "1",
		"pageSize": "5000",
		"indexId":  "1",
		"orderBy":  "stage.publishTime",
	}

	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36",
	}

	// 获取指数数据
	indexURL := "https://zs.zjpwq.net/pwq-index-webapi/indexData"
	indexResp, err := utils.GetWithHeaders(indexURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求排污权交易指数失败: %w", err)
	}

	var indexData eriIndexDataResponse
	if err := json.Unmarshal(indexResp.Body(), &indexData); err != nil {
		return nil, fmt.Errorf("解析排污权交易指数响应失败: %w", err)
	}

	// 获取统计数据
	statsURL := "https://zs.zjpwq.net/pwq-index-webapi/dataStatistics"
	statsResp, err := utils.GetWithHeaders(statsURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求排污权统计数据失败: %w", err)
	}

	var statsData eriStatisticsResponse
	if err := json.Unmarshal(statsResp.Body(), &statsData); err != nil {
		return nil, fmt.Errorf("解析排污权统计数据响应失败: %w", err)
	}

	// 合并数据
	length := len(indexData.Data)
	if len(statsData.Data) < length {
		length = len(statsData.Data)
	}

	result := make([]EriIndex, 0, length)
	for i := 0; i < length; i++ {
		date, err := time.Parse("2006-01-02", indexData.Data[i].Stage.PublishTime[:10])
		if err != nil {
			continue
		}

		item := EriIndex{
			Date:       date,
			TradeIndex: indexData.Data[i].IndexValue,
			Volume:     statsData.Data[i].TotalQuantity,
			Amount:     statsData.Data[i].TotalCost,
		}
		result = append(result, item)
	}

	return result, nil
}
