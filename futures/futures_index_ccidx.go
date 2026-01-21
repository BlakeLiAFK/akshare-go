package futures

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/BlakeLiAFK/akshare/utils"
)

// CCIIndex 中证商品指数
type CCIIndex struct {
	Date        string  `json:"date"`         // 日期
	IndexID     string  `json:"index_id"`     // 指数代码
	ClosePrice  float64 `json:"close_price"`  // 收盘点位
	SettlePrice float64 `json:"settle_price"` // 结算点位
	Change      float64 `json:"change"`       // 涨跌
	ChangeRatio float64 `json:"change_ratio"` // 涨跌幅
}

// cciIndexResponse API响应
type cciIndexResponse struct {
	Data struct {
		DateLineJSON []string `json:"dateLineJson"`
	} `json:"data"`
}

// cciIndexItem 单条数据
type cciIndexItem struct {
	TradeDate                          string  `json:"tradeDate"`
	IndexID                            string  `json:"indexId"`
	ClosingPrice                       float64 `json:"closingPrice"`
	SettlePrice                        float64 `json:"settlePrice"`
	DailyIncreaseAndDecrease           float64 `json:"dailyIncreaseAndDecrease"`
	DailyIncreaseAndDecreasePercentage float64 `json:"dailyIncreaseAndDecreasePercentage"`
}

// CCIIndexMap 中证商品指数代码映射
var CCIIndexMap = map[string]string{
	"中证商品期货指数":   "100001.CCI",
	"中证商品期货价格指数": "000001.CCI",
}

// FuturesIndexCCIDX 中证商品指数-商品指数-日频率
//
// 数据源: http://www.ccidx.com/index.html
//
// 参数:
//   - symbol: 指数名称，可选 "中证商品期货指数", "中证商品期货价格指数"
//
// 返回:
//   - []CCIIndex: 商品指数数据
//   - error: 错误信息
func FuturesIndexCCIDX(symbol string) ([]CCIIndex, error) {
	indexID, ok := CCIIndexMap[symbol]
	if !ok {
		return nil, fmt.Errorf("无效的指数名称: %s", symbol)
	}

	url := "http://www.ccidx.com/CCI-ZZZS/index/getDateLine"
	params := map[string]string{
		"indexId": indexID,
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求中证商品指数失败: %w", err)
	}

	var apiResp cciIndexResponse
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	var result []CCIIndex
	for _, itemStr := range apiResp.Data.DateLineJSON {
		var item cciIndexItem
		if err := json.Unmarshal([]byte(itemStr), &item); err != nil {
			continue
		}

		result = append(result, CCIIndex{
			Date:        item.TradeDate,
			IndexID:     item.IndexID,
			ClosePrice:  item.ClosingPrice,
			SettlePrice: item.SettlePrice,
			Change:      item.DailyIncreaseAndDecrease,
			ChangeRatio: item.DailyIncreaseAndDecreasePercentage,
		})
	}

	// 按日期排序
	sort.Slice(result, func(i, j int) bool {
		return result[i].Date < result[j].Date
	})

	return result, nil
}
