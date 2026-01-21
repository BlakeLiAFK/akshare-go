package stock_feature

import (
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// StockTfpItem 停复牌信息项
type StockTfpItem struct {
	Index              int    `json:"index"`                // 序号
	Code               string `json:"code"`                 // 代码
	Name               string `json:"name"`                 // 名称
	SuspendTime        string `json:"suspend_time"`         // 停牌时间
	SuspendDeadline    string `json:"suspend_deadline"`     // 停牌截止时间
	SuspendPeriod      string `json:"suspend_period"`       // 停牌期限
	SuspendReason      string `json:"suspend_reason"`       // 停牌原因
	Market             string `json:"market"`               // 所属市场
	ExpectedResumeTime string `json:"expected_resume_time"` // 预计复牌时间
}

// StockTfpEm 东方财富网-数据中心-特色数据-停复牌信息
// date: 查询日期，格式 "20240426"
func StockTfpEm(date string) ([]StockTfpItem, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"

	// 格式化日期
	formattedDate := date[:4] + "-" + date[4:6] + "-" + date[6:]

	var allItems []StockTfpItem
	page := 1

	for {
		params := map[string]string{
			"sortColumns": "SUSPEND_START_DATE",
			"sortTypes":   "-1",
			"pageSize":    "500",
			"pageNumber":  fmt.Sprintf("%d", page),
			"reportName":  "RPT_CUSTOM_SUSPEND_DATA_INTERFACE",
			"columns":     "ALL",
			"source":      "WEB",
			"client":      "WEB",
			"filter":      fmt.Sprintf(`(MARKET="全部")(DATETIME='%s')`, formattedDate),
		}

		resp, err := utils.Get(url, params)
		if err != nil {
			return nil, fmt.Errorf("获取停复牌信息失败: %w", err)
		}

		result := gjson.ParseBytes(resp.Body())
		totalPages := result.Get("result.pages").Int()
		dataArr := result.Get("result.data").Array()

		for _, item := range dataArr {
			allItems = append(allItems, StockTfpItem{
				Index:              len(allItems) + 1,
				Code:               item.Get("SECURITY_CODE").String(),
				Name:               item.Get("SECURITY_NAME_ABBR").String(),
				SuspendTime:        strings.Split(item.Get("SUSPEND_START_DATE").String(), " ")[0],
				SuspendDeadline:    strings.Split(item.Get("SUSPEND_END_DATE").String(), " ")[0],
				SuspendPeriod:      item.Get("SUSPEND_PERIOD").String(),
				SuspendReason:      item.Get("SUSPEND_REASON").String(),
				Market:             item.Get("MARKET").String(),
				ExpectedResumeTime: strings.Split(item.Get("RESUMP_DATE").String(), " ")[0],
			})
		}

		if int64(page) >= totalPages || totalPages == 0 {
			break
		}
		page++
	}

	return allItems, nil
}
