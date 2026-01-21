package stock_feature

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

// StockAHighLowStatistics 乐咕乐股-创新高、新低的股票数量
// https://www.legulegu.com/stockdata/high-low-statistics
// symbol: choice of {"all", "sz50", "hs300", "zz500"}
func StockAHighLowStatistics(symbol string) (dataframe.DataFrame, error) {
	url := fmt.Sprintf("https://www.legulegu.com/stockdata/member-ship/get-high-low-statistics/%s", symbol)

	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
	}

	resp, err := utils.GetWithHeaders(url, nil, headers)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("获取创新高新低统计失败: %w", err)
	}

	var dataJSON []map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &dataJSON); err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("解析JSON失败: %w", err)
	}

	if len(dataJSON) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	// 准备 DataFrame 的列
	dates := make([]string, 0, len(dataJSON))
	closes := make([]float64, 0, len(dataJSON))
	high20s := make([]float64, 0, len(dataJSON))
	low20s := make([]float64, 0, len(dataJSON))
	high60s := make([]float64, 0, len(dataJSON))
	low60s := make([]float64, 0, len(dataJSON))
	high120s := make([]float64, 0, len(dataJSON))
	low120s := make([]float64, 0, len(dataJSON))

	for _, item := range dataJSON {
		// 将毫秒时间戳转换为时间字符串
		if dateVal, ok := item["date"].(float64); ok {
			t := time.UnixMilli(int64(dateVal))
			dates = append(dates, t.Format("2006-01-02"))
		}

		closes = append(closes, getFloat64Feature(item, "close"))
		high20s = append(high20s, getFloat64Feature(item, "high20"))
		low20s = append(low20s, getFloat64Feature(item, "low20"))
		high60s = append(high60s, getFloat64Feature(item, "high60"))
		low60s = append(low60s, getFloat64Feature(item, "low60"))
		high120s = append(high120s, getFloat64Feature(item, "high120"))
		low120s = append(low120s, getFloat64Feature(item, "low120"))
	}

	df := dataframe.New(
		series.New(dates, series.String, "date"),
		series.New(closes, series.Float, "close"),
		series.New(high20s, series.Float, "high20"),
		series.New(low20s, series.Float, "low20"),
		series.New(high60s, series.Float, "high60"),
		series.New(low60s, series.Float, "low60"),
		series.New(high120s, series.Float, "high120"),
		series.New(low120s, series.Float, "low120"),
	)

	return df, nil
}

// getFloat64Feature 从 map 中获取 float64 值
func getFloat64Feature(m map[string]interface{}, key string) float64 {
	if v, ok := m[key]; ok {
		if f, ok := v.(float64); ok {
			return f
		}
	}
	return 0
}
