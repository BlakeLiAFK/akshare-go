package stock_fundamental

import (
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// StockZhAGbjgEm 东方财富-A股数据-股本结构
//
// 获取指定股票的股本结构数据
//
// 参数:
//   - symbol: 带市场标识的股票代码，如 "603392.SH"
//
// 返回:
//   - []StockZhAGbjgEmItem: 股本结构数据
//   - error: 错误信息
//
// 示例:
//
//	gbjg, err := stock_fundamental.StockZhAGbjgEm("603392.SH")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range gbjg {
//	    fmt.Printf("%s: 总股本=%.2f万股, 流通A股=%.2f万股\n", item.ChangeDate, item.TotalShares, item.ListedAShares)
//	}
func StockZhAGbjgEm(symbol string) ([]StockZhAGbjgEmItem, error) {
	if symbol == "" {
		return nil, fmt.Errorf("股票代码不能为空")
	}

	url := "https://datacenter.eastmoney.com/securities/api/data/v1/get"
	params := map[string]string{
		"reportName": "RPT_F10_EH_EQUITY",
		"columns": "SECUCODE,SECURITY_CODE,END_DATE,TOTAL_SHARES,LIMITED_SHARES,LIMITED_OTHARS," +
			"LIMITED_DOMESTIC_NATURAL,LIMITED_STATE_LEGAL,LIMITED_OVERSEAS_NOSTATE,LIMITED_OVERSEAS_NATURAL," +
			"UNLIMITED_SHARES,LISTED_A_SHARES,B_FREE_SHARE,H_FREE_SHARE,FREE_SHARES,LIMITED_A_SHARES," +
			"NON_FREE_SHARES,LIMITED_B_SHARES,OTHER_FREE_SHARES,LIMITED_STATE_SHARES," +
			"LIMITED_DOMESTIC_NOSTATE,LOCK_SHARES,LIMITED_FOREIGN_SHARES,LIMITED_H_SHARES," +
			"SPONSOR_SHARES,STATE_SPONSOR_SHARES,SPONSOR_SOCIAL_SHARES,RAISE_SHARES," +
			"RAISE_STATE_SHARES,RAISE_DOMESTIC_SHARES,RAISE_OVERSEAS_SHARES,CHANGE_REASON",
		"quoteColumns": "",
		"filter":       fmt.Sprintf(`(SECUCODE="%s")`, symbol),
		"pageNumber":   "1",
		"pageSize":     "20",
		"sortTypes":    "-1",
		"sortColumns":  "END_DATE",
		"source":       "HSF10",
		"client":       "PC",
		"v":            "047483522105257925",
	}

	headers := map[string]string{
		"Referer": "https://emweb.securities.eastmoney.com/",
	}

	resp, err := utils.GetWithHeaders(url, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求股本结构失败: %w", err)
	}

	json := gjson.ParseBytes(resp.Body())
	dataPath := "result.data"
	result := json.Get(dataPath)

	if !result.Exists() || len(result.Array()) == 0 {
		return nil, fmt.Errorf("未找到股本结构数据")
	}

	var gbjgList []StockZhAGbjgEmItem
	result.ForEach(func(_, value gjson.Result) bool {
		changeDate, _ := utils.ParseDate(value.Get("END_DATE").String())

		gbjg := StockZhAGbjgEmItem{
			ChangeDate:      changeDate,
			TotalShares:     utils.MustFloat64(value.Get("TOTAL_SHARES").String()),
			LimitedShares:   utils.MustFloat64(value.Get("LIMITED_A_SHARES").String()),
			LimitedOthARS:   utils.MustFloat64(value.Get("LIMITED_OTHARS").String()),
			LimitedDomestic: utils.MustFloat64(value.Get("LIMITED_DOMESTIC_NOSTATE").String()),
			LimitedNatural:  utils.MustFloat64(value.Get("LIMITED_DOMESTIC_NATURAL").String()),
			UnlimitedShares: utils.MustFloat64(value.Get("FREE_SHARES").String()),
			ListedAShares:   utils.MustFloat64(value.Get("LISTED_A_SHARES").String()),
			ChangeReason:    value.Get("CHANGE_REASON").String(),
		}
		gbjgList = append(gbjgList, gbjg)
		return true
	})

	return gbjgList, nil
}

// StockZygcEm 东方财富网-个股-主营构成
//
// 获取指定股票的主营构成数据
//
// 参数:
//   - symbol: 带市场标识的股票代码，如 "SH688041"
//
// 返回:
//   - []StockZygcEmItem: 主营构成数据
//   - error: 错误信息
//
// 示例:
//
//	zygc, err := stock_fundamental.StockZygcEm("SH688041")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range zygc {
//	    fmt.Printf("%s %s: %s 收入=%.2f万元, 比例=%.2f%%\n",
//	        item.ReportDate, item.CategoryType, item.MainComposition, item.MainIncome, item.IncomeRatio)
//	}
func StockZygcEm(symbol string) ([]StockZygcEmItem, error) {
	if symbol == "" {
		return nil, fmt.Errorf("股票代码不能为空")
	}

	url := "https://emweb.securities.eastmoney.com/PC_HSF10/BusinessAnalysis/PageAjax"
	params := map[string]string{
		"code": symbol,
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求主营构成失败: %w", err)
	}

	json := gjson.ParseBytes(resp.Body())
	dataPath := "zygcfx"
	result := json.Get(dataPath)

	if !result.Exists() || len(result.Array()) == 0 {
		return nil, fmt.Errorf("未找到主营构成数据")
	}

	var zygcList []StockZygcEmItem
	result.ForEach(func(_, value gjson.Result) bool {
		reportDate, _ := utils.ParseDate(value.Get("REPORT_DATE").String())

		// 分类类型映射
		categoryType := value.Get("MAINOP_TYPE").String()
		categoryType = translateCategoryType(categoryType)

		zygc := StockZygcEmItem{
			Code:             value.Get("SECURITY_CODE").String(),
			ReportDate:       reportDate,
			CategoryType:     categoryType,
			MainComposition:  value.Get("ITEM_NAME").String(),
			MainIncome:       utils.MustFloat64(value.Get("MAIN_BUSINESS_INCOME").String()),
			IncomeRatio:      utils.MustFloat64(value.Get("MBI_RATIO").String()),
			MainCost:         utils.MustFloat64(value.Get("MAIN_BUSINESS_COST").String()),
			CostRatio:        utils.MustFloat64(value.Get("MBC_RATIO").String()),
			MainProfit:       utils.MustFloat64(value.Get("MAIN_BUSINESS_RPOFIT").String()),
			ProfitRatio:      utils.MustFloat64(value.Get("MBR_RATIO").String()),
			GrossProfitRatio: utils.MustFloat64(value.Get("GROSS_RPOFIT_RATIO").String()),
		}
		zygcList = append(zygcList, zygc)
		return true
	})

	return zygcList, nil
}

// translateCategoryType 翻译分类类型
func translateCategoryType(categoryType string) string {
	switch categoryType {
	case "2":
		return "按产品分类"
	case "3":
		return "按地区分类"
	default:
		return categoryType
	}
}
