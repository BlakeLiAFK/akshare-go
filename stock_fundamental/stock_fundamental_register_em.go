package stock_fundamental

import (
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// StockRegisterAllEm 东方财富网-IPO审核信息-全部
//
// 获取全部IPO审核信息，包含企业名称、最新状态、注册地、行业、保荐机构等信息
//
// 返回:
//   - []StockRegisterEmItem: IPO审核信息列表
//   - error: 错误信息
//
// 示例:
//
//	register, err := stock_fundamental.StockRegisterAllEm()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range register {
//	    fmt.Printf("%s: %s, 保荐机构=%s\n", item.CompanyName, item.State, item.RecommendOrg)
//	}
func StockRegisterAllEm() ([]StockRegisterEmItem, error) {
	return stockRegisterEm("")
}

// StockRegisterKcb 东方财富网-IPO审核信息-科创板
//
// 获取科创板IPO审核信息
//
// 返回:
//   - []StockRegisterEmItem: 科创板IPO审核信息列表
//   - error: 错误信息
//
// 示例:
//
//	register, err := stock_fundamental.StockRegisterKcb()
//	if err != nil {
//	    log.Fatal(err)
//	}
func StockRegisterKcb() ([]StockRegisterEmItem, error) {
	return stockRegisterEm("(PREDICT_LISTING_MARKET=\"科创板\")")
}

// StockRegisterCyb 东方财富网-IPO审核信息-创业板
//
// 获取创业板IPO审核信息
//
// 返回:
//   - []StockRegisterEmItem: 创业板IPO审核信息列表
//   - error: 错误信息
//
// 示例:
//
//	register, err := stock_fundamental.StockRegisterCyb()
//	if err != nil {
//	    log.Fatal(err)
//	}
func StockRegisterCyb() ([]StockRegisterEmItem, error) {
	return stockRegisterEm("(PREDICT_LISTING_MARKET=\"创业板\")")
}

// stockRegisterEm IPO审核信息通用实现
func stockRegisterEm(filter string) ([]StockRegisterEmItem, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"

	params := map[string]string{
		"sortColumns": "UPDATE_DATE,ORG_CODE",
		"sortTypes":   "-1,-1",
		"pageSize":    "500",
		"pageNumber":  "1",
		"reportName":  "RPT_IPO_INFOALLNEW",
		"columns":     "SECURITY_CODE,STATE,REG_ADDRESS,INFO_CODE,CSRC_INDUSTRY,ACCEPT_DATE,DECLARE_ORG,PREDICT_LISTING_MARKET,LAW_FIRM,ACCOUNT_FIRM,ORG_CODE,UPDATE_DATE,RECOMMEND_ORG,IS_REGISTRATION",
		"source":      "WEB",
		"client":      "WEB",
	}

	if filter != "" {
		params["filter"] = filter
	}

	headers := map[string]string{
		"Referer": "https://data.eastmoney.com/xg/ipo/",
	}

	// 首次请求获取总页数
	resp, err := utils.GetWithHeaders(url, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求IPO审核信息失败: %w", err)
	}

	json := gjson.ParseBytes(resp.Body())
	pageNum := json.Get("result.pages").Int()
	if pageNum == 0 {
		pageNum = 1
	}

	var allItems []StockRegisterEmItem

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
			updateDateStr := data.Get("UPDATE_DATE").String()
			updateDate, _ := utils.ParseDate(updateDateStr)

			acceptDateStr := data.Get("ACCEPT_DATE").String()
			acceptDate, _ := utils.ParseDate(acceptDateStr)

			// 构建招股说明书链接
			infoCode := data.Get("INFO_CODE").String()
			prospectusURL := ""
			if infoCode != "" {
				prospectusURL = fmt.Sprintf("https://pdf.dfcfw.com/pdf/H2_%s_1.pdf", infoCode)
			}

			item := StockRegisterEmItem{
				CompanyName:   data.Get("DECLARE_ORG").String(),
				State:         data.Get("STATE").String(),
				RegAddress:    data.Get("REG_ADDRESS").String(),
				Industry:      data.Get("CSRC_INDUSTRY").String(),
				RecommendOrg:  data.Get("RECOMMEND_ORG").String(),
				LawFirm:       data.Get("LAW_FIRM").String(),
				AccountFirm:   data.Get("ACCOUNT_FIRM").String(),
				UpdateDate:    updateDate,
				AcceptDate:    acceptDate,
				PredictMarket: data.Get("PREDICT_LISTING_MARKET").String(),
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
