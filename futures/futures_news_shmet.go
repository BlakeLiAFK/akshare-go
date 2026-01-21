package futures

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
)

// SHMETNews 上海金属网快讯
type SHMETNews struct {
	PublishTime time.Time `json:"publish_time"` // 发布时间
	Content     string    `json:"content"`      // 内容
}

// shmetNewsResponse API响应
type shmetNewsResponse struct {
	Data struct {
		DataList []shmetNewsItem `json:"dataList"`
	} `json:"data"`
}

// shmetNewsItem 快讯数据项
type shmetNewsItem struct {
	PublishTime int64  `json:"publishTime"` // 发布时间戳(毫秒)
	Content     string `json:"content"`     // 内容
}

// SHMETNewsSymbolMap 上海金属网快讯类别映射
var SHMETNewsSymbolMap = map[string]string{
	"要闻":  "0",
	"VIP": "100",
	"财经":  "999",
	"铜":   "1002",
	"铝":   "1003",
	"铅":   "1005",
	"锌":   "1004",
	"镍":   "1006",
	"锡":   "1007",
	"贵金属": "1008",
	"小金属": "1009",
}

// FuturesNewsSHMET 上海金属网-快讯
//
// 数据源: https://www.shmet.com/newsFlash/newsFlash.html
//
// 参数:
//   - symbol: 类别，可选 "全部", "要闻", "VIP", "财经", "铜", "铝", "铅", "锌", "镍", "锡", "贵金属", "小金属"
//
// 返回:
//   - []SHMETNews: 快讯数据
//   - error: 错误信息
func FuturesNewsSHMET(symbol string) ([]SHMETNews, error) {
	url := "https://www.shmet.com/api/rest/news/queryNewsflashList"

	var payload map[string]any
	if symbol == "全部" {
		payload = map[string]any{
			"currentPage": 1,
			"pageSize":    100,
		}
	} else {
		flashTag, ok := SHMETNewsSymbolMap[symbol]
		if !ok {
			return nil, fmt.Errorf("无效的类别: %s", symbol)
		}
		payload = map[string]any{
			"currentPage": 1,
			"pageSize":    2000,
			"content":     "",
			"flashTag":    flashTag,
		}
	}

	headers := map[string]string{
		"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		"Content-Type": "application/json",
	}

	resp, err := utils.PostJSONWithHeaders(url, nil, payload, headers)
	if err != nil {
		return nil, fmt.Errorf("请求上海金属网快讯失败: %w", err)
	}

	var apiResp shmetNewsResponse
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	// 加载上海时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		loc = time.FixedZone("CST", 8*3600)
	}

	var result []SHMETNews
	for _, item := range apiResp.Data.DataList {
		publishTime := time.UnixMilli(item.PublishTime).In(loc)
		result = append(result, SHMETNews{
			PublishTime: publishTime,
			Content:     item.Content,
		})
	}

	// 按时间排序
	sort.Slice(result, func(i, j int) bool {
		return result[i].PublishTime.Before(result[j].PublishTime)
	})

	return result, nil
}
