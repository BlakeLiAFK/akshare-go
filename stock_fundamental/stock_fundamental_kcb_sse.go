package stock_fundamental

import (
	"fmt"
	"strconv"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// StockKcbSse 上交所科创板-股票发行
//
// 获取上交所科创板股票发行信息
//
// 返回:
//   - []StockKcbSseItem: 科创板信息列表
//   - error: 错误信息
//
// 示例:
//
//	kcb, err := stock_fundamental.StockKcbSse()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range kcb {
//	    fmt.Printf("%s: %s, 状态=%s\n", item.CompanyName, item.StockCode, item.CurrentStatus)
//	}
func StockKcbSse() ([]StockKcbSseItem, error) {
	url := "http://query.sse.com.cn/statusAction.do"

	headers := map[string]string{
		"Host":    "query.sse.com.cn",
		"Referer": "http://kcb.sse.com.cn/",
	}

	// 首次请求获取总页数
	params := map[string]string{
		"isPagination":        "true",
		"sqlId":               "SH_XM_LB",
		"pageHelp.pageSize":   "20",
		"offerType":           "",
		"commitiResult":       "",
		"registeResult":       "",
		"province":            "",
		"csrcCode":            "",
		"currStatus":          "",
		"order":               "updateDate|desc,stockAuditNum|desc",
		"keyword":             "",
		"auditApplyDateBegin": "",
		"auditApplyDateEnd":   "",
		"pageHelp.pageNo":     "1",
		"pageHelp.beginPage":  "1",
		"pageHelp.endPage":    "1",
		"_":                   strconv.FormatInt(time.Now().UnixMilli(), 10),
	}

	resp, err := utils.GetWithHeaders(url, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求科创板信息失败: %w", err)
	}

	json := gjson.ParseBytes(resp.Body())

	// 解析总页数
	totalPages := json.Get("pageHelp.totalPages").Int()
	if totalPages == 0 {
		totalPages = 1
	}

	_ = json.Get("pageHelp.totalCount").Int() // 总数，供参考

	var allItems []StockKcbSseItem

	// 分页获取所有数据
	for page := 1; page <= int(totalPages); page++ {
		params["pageHelp.pageNo"] = strconv.Itoa(page)
		params["pageHelp.beginPage"] = strconv.Itoa(page)
		params["pageHelp.endPage"] = strconv.Itoa(page)

		resp, err := utils.GetWithHeaders(url, params, headers)
		if err != nil {
			return nil, fmt.Errorf("请求第%d页失败: %w", page, err)
		}

		json := gjson.ParseBytes(resp.Body())
		resultArray := json.Get("result").Array()

		for _, data := range resultArray {
			// 解析更新日期 (格式: 20260116171217)
			updateTimeStr := data.Get("updateDate").String()
			var updateTime time.Time
			if len(updateTimeStr) >= 8 {
				updateTime, _ = time.ParseInLocation("20060102", updateTimeStr[:8], time.Local)
			}

			// 解析受理日期
			auditApplyDateStr := data.Get("auditApplyDate").String()
			var auditApplyDate time.Time
			if len(auditApplyDateStr) >= 8 {
				auditApplyDate, _ = time.ParseInLocation("20060102", auditApplyDateStr[:8], time.Local)
			}

			// 当前状态转为字符串
			currStatus := data.Get("currStatus").Int()
			currStatusStr := ""
			switch currStatus {
			case 1:
				currStatusStr = "已受理"
			case 2:
				currStatusStr = "已问询"
			case 3:
				currStatusStr = "上市委会议通过"
			case 4:
				currStatusStr = "提交注册"
			case 5:
				currStatusStr = "注册结果"
			case 6:
				currStatusStr = "中止"
			case 7:
				currStatusStr = "终止"
			default:
				currStatusStr = fmt.Sprintf("%d", currStatus)
			}

			item := StockKcbSseItem{
				CompanyName:    data.Get("stockAuditName").String(),
				StockCode:      data.Get("stockCode").String(),
				UpdateTime:     updateTime,
				AuditApplyDate: auditApplyDate,
				RegisterResult: data.Get("registeResult").String(),
				CommitResult:   data.Get("commitiResult").String(),
				CurrentStatus:  currStatusStr,
				Province:       data.Get("province").String(),
				CsrcCode:       data.Get("csrcCode").String(),
				StockAuditNum:  utils.MustInt(data.Get("stockAuditNum").String()),
			}
			allItems = append(allItems, item)
		}
	}

	// 设置序号
	for i := range allItems {
		allItems[i].Index = i + 1
	}

	return allItems, nil
}
