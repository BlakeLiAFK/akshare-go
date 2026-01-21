package index

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
)

// NewsSentimentIndex A股新闻情绪指数
type NewsSentimentIndex struct {
	Date           time.Time `json:"date"`            // 日期
	SentimentIndex float64   `json:"sentiment_index"` // 市场情绪指数
	HS300Index     float64   `json:"hs300_index"`     // 沪深300指数
}

// scopeSentimentResponse 数库情绪指数API响应
type scopeSentimentResponse []struct {
	TradeDate   string  `json:"tradeDate"`
	MaIndex1    float64 `json:"maIndex1"`
	MarketClose float64 `json:"marketClose"`
}

// IndexNewsSentimentScope 数库-A股新闻情绪指数
//
// 数据源: https://www.chinascope.com/reasearch.html
//
// 返回:
//   - []NewsSentimentIndex: A股新闻情绪指数数据
//   - error: 错误信息
func IndexNewsSentimentScope() ([]NewsSentimentIndex, error) {
	apiURL := "https://www.chinascope.com/inews/senti/index"

	params := map[string]string{
		"period": "YEAR",
	}

	resp, err := utils.Get(apiURL, params)
	if err != nil {
		return nil, fmt.Errorf("请求A股新闻情绪指数失败: %w", err)
	}

	var apiResp scopeSentimentResponse
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, fmt.Errorf("解析A股新闻情绪指数响应失败: %w", err)
	}

	result := make([]NewsSentimentIndex, 0, len(apiResp))
	for _, item := range apiResp {
		date, err := time.Parse("2006-01-02", item.TradeDate[:10])
		if err != nil {
			continue
		}

		result = append(result, NewsSentimentIndex{
			Date:           date,
			SentimentIndex: item.MaIndex1,
			HS300Index:     item.MarketClose,
		})
	}

	return result, nil
}
