package stock_fundamental

import (
	"fmt"
	"math"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// StockNoticeReportItem 股票公告项
type StockNoticeReportItem struct {
	Code       string `json:"code"`        // 股票代码
	Name       string `json:"name"`        // 股票名称
	Title      string `json:"title"`       // 公告标题
	NoticeType string `json:"notice_type"` // 公告类型
	NoticeDate string `json:"notice_date"` // 公告日期
	URL        string `json:"url"`         // 公告网址
}

// StockNoticeReport 东方财富网-数据中心-公告大全-沪深京 A 股公告
// symbol: 报告类型，可选 "全部", "重大事项", "财务报告", "融资公告", "风险提示", "资产重组", "信息变更", "持股变动"
// date: 指定日期，格式 "20220511"
func StockNoticeReport(symbol, date string) ([]StockNoticeReportItem, error) {
	baseURL := "https://np-anotice-stock.eastmoney.com/api/security/ann"

	reportMap := map[string]string{
		"全部":   "0",
		"财务报告": "1",
		"融资公告": "2",
		"风险提示": "3",
		"信息变更": "4",
		"重大事项": "5",
		"资产重组": "6",
		"持股变动": "7",
	}

	nodeCode, ok := reportMap[symbol]
	if !ok {
		nodeCode = "0"
	}

	// 格式化日期
	formattedDate := date
	if len(date) == 8 {
		formattedDate = date[:4] + "-" + date[4:6] + "-" + date[6:]
	}

	params := map[string]string{
		"sr":            "-1",
		"page_size":     "100",
		"page_index":    "1",
		"ann_type":      "A",
		"client_source": "web",
		"f_node":        nodeCode,
		"s_node":        "0",
		"begin_time":    formattedDate,
		"end_time":      formattedDate,
	}

	resp, err := utils.Get(baseURL, params)
	if err != nil {
		return nil, fmt.Errorf("获取公告数据失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	totalHits := result.Get("data.total_hits").Int()
	if totalHits == 0 {
		return nil, nil
	}

	totalPage := int(math.Ceil(float64(totalHits) / 100))

	var items []StockNoticeReportItem

	for page := 1; page <= totalPage; page++ {
		params["page_index"] = fmt.Sprintf("%d", page)

		resp, err := utils.Get(baseURL, params)
		if err != nil {
			continue
		}

		result := gjson.ParseBytes(resp.Body())
		list := result.Get("data.list").Array()

		for _, item := range list {
			artCode := item.Get("art_code").String()
			title := item.Get("title").String()
			noticeDate := item.Get("notice_date").String()

			// 获取公告类型
			columnName := ""
			columns := item.Get("columns").Array()
			if len(columns) > 0 {
				columnName = columns[0].Get("column_name").String()
			}

			// 获取股票代码和名称
			var stockCode, shortName string
			codes := item.Get("codes").Array()
			if len(codes) == 1 {
				stockCode = codes[0].Get("stock_code").String()
				shortName = codes[0].Get("short_name").String()
			} else {
				for _, code := range codes {
					annType := code.Get("ann_type").String()
					if strings.HasPrefix(annType, "A") {
						stockCode = code.Get("stock_code").String()
						shortName = code.Get("short_name").String()
						break
					}
				}
			}

			// 构建公告URL
			noticeURL := fmt.Sprintf("https://data.eastmoney.com/notices/detail/%s/%s.html", stockCode, artCode)

			// 格式化日期
			if len(noticeDate) >= 10 {
				noticeDate = noticeDate[:10]
			}

			items = append(items, StockNoticeReportItem{
				Code:       stockCode,
				Name:       shortName,
				Title:      title,
				NoticeType: columnName,
				NoticeDate: noticeDate,
				URL:        noticeURL,
			})
		}
	}

	return items, nil
}
