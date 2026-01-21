package stock_fundamental

import (
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// StockIpoDeclareEm 东方财富网-数据中心-新股申购-首发申报企业信息
//
// 获取首发申报企业信息，包含企业名称、最新状态、注册地、保荐机构等信息
//
// 返回:
//   - []StockIpoDeclareEmItem: 首发申报企业信息列表
//   - error: 错误信息
//
// 示例:
//
//	declare, err := stock_fundamental.StockIpoDeclareEm()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range declare {
//	    fmt.Printf("%s: %s, 保荐机构=%s\n", item.CompanyName, item.State, item.RecommendOrg)
//	}
func StockIpoDeclareEm() ([]StockIpoDeclareEmItem, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"

	params := map[string]string{
		"sortColumns": "END_DATE,SECURITY_CODE",
		"sortTypes":   "-1,-1",
		"pageSize":    "500",
		"pageNumber":  "1",
		"reportName":  "RPT_IPO_DECORGNEWEST",
		"columns":     "DECLARE_ORG,STATE,REG_ADDRESS,RECOMMEND_ORG,LAW_FIRM,ACCOUNT_FIRM,IS_SUBMIT,PREDICT_LISTING_MARKET,END_DATE,INFO_CODE,SECURITY_CODE,ORG_CODE,IS_REGISTER,STATE_CODE,DERIVE_SECURITY_CODE,ORG_CODE_OLD,IS_STATE",
		"source":      "WEB",
		"client":      "WEB",
	}

	headers := map[string]string{
		"Referer": "https://data.eastmoney.com/xg/xg/sbqy.html",
	}

	// 首次请求获取总页数
	resp, err := utils.GetWithHeaders(url, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求首发申报企业信息失败: %w", err)
	}

	json := gjson.ParseBytes(resp.Body())
	pageNum := json.Get("result.pages").Int()
	if pageNum == 0 {
		pageNum = 1
	}

	var allItems []StockIpoDeclareEmItem

	// 分页获取所有数据
	for page := 1; page <= int(pageNum); page++ {
		params["pageNumber"] = fmt.Sprintf("%d", page)
		resp, err := utils.GetWithHeaders(url, params, headers)
		if err != nil {
			return nil, fmt.Errorf("请求第%d页失败: %w", page, err)
		}

		json := gjson.ParseBytes(resp.Body())
		dataArray := json.Get("result.data").Array()

		for _, data := range dataArray {
			endDateStr := data.Get("END_DATE").String()
			endDate, _ := utils.ParseDate(endDateStr)

			// 构建招股说明书链接
			infoCode := data.Get("INFO_CODE").String()
			prospectusURL := ""
			if infoCode != "" {
				prospectusURL = fmt.Sprintf("https://pdf.dfcfw.com/pdf/H2_%s_1.pdf", infoCode)
			}

			item := StockIpoDeclareEmItem{
				CompanyName:   data.Get("DECLARE_ORG").String(),
				State:         data.Get("STATE").String(),
				RegAddress:    data.Get("REG_ADDRESS").String(),
				RecommendOrg:  data.Get("RECOMMEND_ORG").String(),
				LawFirm:       data.Get("LAW_FIRM").String(),
				AccountFirm:   data.Get("ACCOUNT_FIRM").String(),
				PredictMarket: data.Get("PREDICT_LISTING_MARKET").String(),
				EndDate:       endDate,
				ProspectusURL: prospectusURL,
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

// parseJSONP 解析 JSONP 响应，提取纯 JSON 部分
func parseJSONP(text string) (string, error) {
	// 找到第一个 { 和最后一个 } 之间的内容
	start := strings.Index(text, "{")
	end := strings.LastIndex(text, "}")

	if start == -1 || end == -1 || end <= start {
		return "", fmt.Errorf("无法找到 JSON 对象")
	}

	return text[start : end+1], nil
}

// StockIpoReviewEm 东方财富网-数据中心-新股申购-过会企业信息
//
// 获取新股上会信息，包含企业名称、股票简称、上市板块、审核状态等信息
//
// 返回:
//   - []StockIpoReviewEmItem: 过会企业信息列表
//   - error: 错误信息
//
// 示例:
//
//	review, err := stock_fundamental.StockIpoReviewEm()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range review {
//	    fmt.Printf("%s: %s, 审核状态=%s\n", item.CompanyName, item.StockName, item.ReviewState)
//	}
func StockIpoReviewEm() ([]StockIpoReviewEmItem, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"

	params := map[string]string{
		"sortColumns": "REVIEW_DATE,ORG_CODE",
		"sortTypes":   "-1,-1",
		"pageSize":    "500",
		"pageNumber":  "1",
		"reportName":  "RPT_IPO_REVIEW",
		"columns":     "ALL",
		"source":      "WEB",
		"client":      "WEB",
	}

	headers := map[string]string{
		"Referer": "https://data.eastmoney.com/xg/gh/default.html",
	}

	// 首次请求获取总页数
	resp, err := utils.GetWithHeaders(url, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求过会企业信息失败: %w", err)
	}

	// 解析 JSONP 响应
	jsonStr, err := parseJSONP(resp.String())
	if err != nil {
		return nil, fmt.Errorf("解析 JSONP 响应失败: %w", err)
	}

	json := gjson.Parse(jsonStr)
	pageNum := json.Get("result.pages").Int()
	if pageNum == 0 {
		pageNum = 1
	}

	var allItems []StockIpoReviewEmItem

	// 分页获取所有数据
	for page := 1; page <= int(pageNum); page++ {
		params["pageNumber"] = fmt.Sprintf("%d", page)
		resp, err := utils.GetWithHeaders(url, params, headers)
		if err != nil {
			return nil, fmt.Errorf("请求第%d页失败: %w", page, err)
		}

		// 解析 JSONP 响应
		jsonStr, err := parseJSONP(resp.String())
		if err != nil {
			// 跳过解析失败的页面
			continue
		}

		json := gjson.Parse(jsonStr)
		dataArray := json.Get("result.data").Array()

		for _, data := range dataArray {
			reviewDateStr := data.Get("REVIEW_DATE").String()
			reviewDate, _ := utils.ParseDate(reviewDateStr)

			noticeDateStr := data.Get("NOTICE_DATE").String()
			noticeDate, _ := utils.ParseDate(noticeDateStr)

			listingDateStr := data.Get("LISTING_DATE").String()
			listingDate, _ := utils.ParseDate(listingDateStr)

			item := StockIpoReviewEmItem{
				CompanyName:     data.Get("ORG_NAME").String(),
				StockName:       data.Get("SECURITY_NAME_ABBR").String(),
				StockCode:       data.Get("SECURITY_CODE").String(),
				TradeMarket:     data.Get("TRADE_MARKET").String(),
				ReviewDate:      reviewDate,
				ReviewState:     data.Get("REVIEW_STATE").String(),
				ReviewMember:    data.Get("REVIEW_MEMBER").String(),
				LeadUnderwriter: data.Get("LEAD_UNDERWRITER").String(),
				IssueNum:        utils.MustFloat64(data.Get("ISSUE_NUM").String()),
				FinanceAmt:      utils.MustFloat64(data.Get("FINANCE_AMT_UPPER").String()),
				NoticeDate:      noticeDate,
				ListingDate:     listingDate,
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

// StockIpoTutorEm 东方财富网-数据中心-新股申购-辅导备案信息
//
// 获取IPO辅导信息，包含企业名称、辅导机构、辅导状态、报告类型等信息
//
// 返回:
//   - []StockIpoTutorEmItem: 辅导备案信息列表
//   - error: 错误信息
//
// 示例:
//
//	tutor, err := stock_fundamental.StockIpoTutorEm()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range tutor {
//	    fmt.Printf("%s: %s, 辅导状态=%s\n", item.CompanyName, item.TutorOrg, item.TutorProcessState)
//	}
func StockIpoTutorEm() ([]StockIpoTutorEmItem, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"

	params := map[string]string{
		"sortColumns": "RECORD_DATE,TUTOR_OBJECT",
		"sortTypes":   "-1,-1",
		"pageSize":    "500",
		"pageNumber":  "1",
		"reportName":  "RPT_IPO_TUTRECORD",
		"columns":     "TUTOR_OBJECT,ORG_CODE,TUTOR_ORG_CODE,TUTOR_ORG,TUTOR_PROCESS_STATE,REPORT_TYPE,DISPATCH_ORG,REPORT_TITLE,RECORD_DATE",
		"source":      "WEB",
		"client":      "WEB",
	}

	headers := map[string]string{
		"Referer": "https://data.eastmoney.com/xg/ipo/fd.html",
	}

	// 首次请求获取总页数
	resp, err := utils.GetWithHeaders(url, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求辅导备案信息失败: %w", err)
	}

	// 解析 JSONP 响应
	jsonStr, err := parseJSONP(resp.String())
	if err != nil {
		return nil, fmt.Errorf("解析 JSONP 响应失败: %w", err)
	}

	json := gjson.Parse(jsonStr)
	pageNum := json.Get("result.pages").Int()
	if pageNum == 0 {
		pageNum = 1
	}

	var allItems []StockIpoTutorEmItem

	// 分页获取所有数据
	for page := 1; page <= int(pageNum); page++ {
		params["pageNumber"] = fmt.Sprintf("%d", page)
		resp, err := utils.GetWithHeaders(url, params, headers)
		if err != nil {
			return nil, fmt.Errorf("请求第%d页失败: %w", page, err)
		}

		// 解析 JSONP 响应
		jsonStr, err := parseJSONP(resp.String())
		if err != nil {
			// 跳过解析失败的页面
			continue
		}

		json := gjson.Parse(jsonStr)
		dataArray := json.Get("result.data").Array()

		for _, data := range dataArray {
			recordDateStr := data.Get("RECORD_DATE").String()
			recordDate, _ := utils.ParseDate(recordDateStr)

			item := StockIpoTutorEmItem{
				CompanyName:       data.Get("TUTOR_OBJECT").String(),
				TutorOrg:          data.Get("TUTOR_ORG").String(),
				TutorProcessState: data.Get("TUTOR_PROCESS_STATE").String(),
				ReportType:        data.Get("REPORT_TYPE").String(),
				DispatchOrg:       data.Get("DISPATCH_ORG").String(),
				ReportTitle:       data.Get("REPORT_TITLE").String(),
				RecordDate:        recordDate,
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
