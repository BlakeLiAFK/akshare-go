package stock_fundamental

import (
	"fmt"
	"strings"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-gota/gota/dataframe"
	"github.com/tidwall/gjson"
)

// StockFinancialReportSina 新浪财经-财务报表-三大报表
//
// 获取指定股票的资产负债表、利润表或现金流量表数据
//
// 参数:
//   - stock: 股票代码，如 "sh600600"（上海）或 "sz000001"（深圳）
//   - symbol: 报表类型，可选值: "资产负债表", "利润表", "现金流量表"
//
// 返回:
//   - dataframe.DataFrame: 财务报表数据，每行代表一个报告期
//   - error: 错误信息
//
// 示例:
//
//	df, err := stock_fundamental.StockFinancialReportSina("sh600600", "资产负债表")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(df)
func StockFinancialReportSina(stock, symbol string) (dataframe.DataFrame, error) {
	// 报表类型映射
	symbolMap := map[string]string{
		"资产负债表": "fzb",
		"利润表":   "lrb",
		"现金流量表": "llb",
	}

	source, ok := symbolMap[symbol]
	if !ok {
		return dataframe.DataFrame{}, fmt.Errorf("不支持的报表类型: %s，可选: 资产负债表/利润表/现金流量表", symbol)
	}

	url := "https://quotes.sina.cn/cn/api/openapi.php/CompanyFinanceService.getFinanceReport2022"
	params := map[string]string{
		"paperCode": stock,
		"source":    source,
		"type":      "0",
		"page":      "1",
		"num":       "1000",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求财务报表失败: %w", err)
	}

	json := gjson.ParseBytes(resp.Body())
	resultPath := "result.data.report_date"
	reportDates := json.Get(resultPath).Array()
	if len(reportDates) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未找到报表数据")
	}

	// 提取所有报告日期
	var dates []string
	for _, item := range reportDates {
		dates = append(dates, item.Get("date_value").String())
	}

	// 构建数据
	// 第一行是表头（报告日期）
	var allRecords [][]string
	headers := append([]string{"报告日"}, dates...)
	allRecords = append(allRecords, headers)

	// 收集所有项目名称
	itemMap := make(map[string]int)
	for _, date := range dates {
		reportListPath := "result.data.report_list." + date + ".data"
		items := json.Get(reportListPath).Array()
		for _, item := range items {
			title := item.Get("item_title").String()
			if _, exists := itemMap[title]; !exists {
				itemMap[title] = len(itemMap)
			}
		}
	}

	// 按项目组织数据
	for title := range itemMap {
		row := []string{title}
		for _, date := range dates {
			reportListPath := "result.data.report_list." + date + ".data"
			items := json.Get(reportListPath).Array()
			found := false
			for _, item := range items {
				if item.Get("item_title").String() == title {
					row = append(row, item.Get("item_value").String())
					found = true
					break
				}
			}
			if !found {
				row = append(row, "")
			}
		}
		allRecords = append(allRecords, row)
	}

	// 添加元数据行 - 每个元数据项一行，包含所有日期的值
	metadataKeys := []string{"数据源", "是否审计", "公告日期", "币种", "类型", "更新日期"}
	for _, key := range metadataKeys {
		row := []string{key}
		for _, date := range dates {
			reportListPath := "result.data.report_list." + date
			var value string
			switch key {
			case "数据源":
				value = json.Get(reportListPath + ".data_source").String()
			case "是否审计":
				value = json.Get(reportListPath + ".is_audit").String()
			case "公告日期":
				value = json.Get(reportListPath + ".publish_date").String()
			case "币种":
				value = json.Get(reportListPath + ".rCurrency").String()
			case "类型":
				value = json.Get(reportListPath + ".rType").String()
			case "更新日期":
				updateTime := json.Get(reportListPath + ".update_time").Int()
				value = time.Unix(updateTime, 0).Format("2006-01-02 15:04:05")
			}
			row = append(row, value)
		}
		allRecords = append(allRecords, row)
	}

	df := dataframe.LoadRecords(allRecords)

	return df, nil
}

// StockFinancialAbstract 新浪财经-财务报表-关键指标
//
// 获取指定股票的关键财务指标摘要
//
// 参数:
//   - symbol: 股票代码，如 "600004"（不带市场前缀）
//
// 返回:
//   - dataframe.DataFrame: 关键财务指标，包含常用指标、每股指标、盈利能力等分类
//   - error: 错误信息
//
// 示例:
//
//	df, err := stock_fundamental.StockFinancialAbstract("600004")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(df)
func StockFinancialAbstract(symbol string) (dataframe.DataFrame, error) {
	if symbol == "" {
		return dataframe.DataFrame{}, fmt.Errorf("股票代码不能为空")
	}

	url := "https://quotes.sina.cn/cn/api/openapi.php/CompanyFinanceService.getFinanceReport2022"
	params := map[string]string{
		"paperCode": "sh" + symbol,
		"source":    "gjzb",
		"type":      "0",
		"page":      "1",
		"num":       "1000",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求财务摘要失败: %w", err)
	}

	json := gjson.ParseBytes(resp.Body())
	reportListPath := "result.data.report_list"
	reportList := json.Get(reportListPath).Map()

	if len(reportList) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未找到财务摘要数据")
	}

	// 获取所有报告期
	var dates []string
	for date := range reportList {
		dates = append(dates, date)
	}

	// 构建指标映射
	indicatorMap := make(map[string][]string) // indicator name -> values per period

	// 第一期数据建立指标列表
	for _, date := range dates {
		dataPath := "result.data.report_list." + date + ".data"
		items := json.Get(dataPath).Array()
		for _, item := range items {
			title := item.Get("item_title").String()
			value := item.Get("item_value").String()
			indicatorMap[title] = append(indicatorMap[title], value)
		}
	}

	// 构建数据
	// 按照Python代码的分类方式组织数据
	categories := []struct {
		name     string
		startKey string
		endKey   string
	}{
		{"常用指标", "常用指标", "每股指标"},
		{"每股指标", "每股指标", "盈利能力"},
		{"盈利能力", "盈利能力", "成长能力"},
		{"成长能力", "成长能力", "收益质量"},
		{"收益质量", "收益质量", "财务风险"},
		{"财务风险", "财务风险", "营运能力"},
		{"营运能力", "营运能力", ""},
	}

	var allRecords [][]string
	headers := append([]string{"选项", "指标"}, dates...)
	allRecords = append(allRecords, headers)

	// 需要找到每个分类的起止指标
	for _, cat := range categories {
		// 获取该分类下的所有指标
		var indicators []string
		inCategory := false

		// 从原始数据中提取分类
		for title := range indicatorMap {
			if title == cat.startKey {
				inCategory = true
			}
			if inCategory {
				if cat.endKey != "" && title == cat.endKey {
					break
				}
				if title != cat.startKey {
					indicators = append(indicators, title)
				}
			}
		}

		// 为每个指标创建一行
		for _, indicator := range indicators {
			row := []string{cat.name, indicator}
			row = append(row, indicatorMap[indicator]...)
			allRecords = append(allRecords, row)
		}
	}

	df := dataframe.LoadRecords(allRecords)

	return df, nil
}

// StockFinancialAnalysisIndicatorEm 东方财富-A股-财务分析-主要指标
//
// 获取指定股票的财务分析指标
//
// 参数:
//   - symbol: 股票代码（带市场标识），如 "301389.SZ" 或 "600000.SH"
//   - indicator: 数据类型，"按报告期" 或 "按单季度"
//
// 返回:
//   - []StockFinancialAnalysisIndicatorEm: 财务分析指标列表
//   - error: 错误信息
//
// 示例:
//
//	indicators, err := stock_fundamental.StockFinancialAnalysisIndicatorEm("301389.SZ", "按报告期")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range indicators {
//	    fmt.Printf("%s: ROE=%.2f%%\n", item.ReportDate, item.ROE)
//	}
func StockFinancialAnalysisIndicatorEm(symbol, indicator string) ([]StockFinancialAnalysisIndicatorEmItem, error) {
	if symbol == "" {
		return nil, fmt.Errorf("股票代码不能为空")
	}

	var url string
	var params map[string]string

	if indicator == "按报告期" {
		url = "https://datacenter.eastmoney.com/securities/api/data/get"
		params = map[string]string{
			"type":         "RPT_F10_FINANCE_MAINFINADATA",
			"sty":          "APP_F10_MAINFINADATA",
			"quoteColumns": "",
			"filter":       fmt.Sprintf(`(SECUCODE="%s")`, symbol),
			"p":            "1",
			"ps":           "200",
			"sr":           "-1",
			"st":           "REPORT_DATE",
			"source":       "HSF10",
			"client":       "PC",
		}
	} else {
		url = "https://datacenter.eastmoney.com/securities/api/data/v1/get"
		params = map[string]string{
			"reportName":   "RPT_F10_QTR_MAINFINADATA",
			"columns":      "ALL",
			"quoteColumns": "",
			"filter":       fmt.Sprintf(`(SECUCODE="%s")`, symbol),
			"pageNumber":   "1",
			"pageSize":     "200",
			"sortTypes":    "-1",
			"sortColumns":  "REPORT_DATE",
			"source":       "HSF10",
			"client":       "PC",
		}
	}

	headers := map[string]string{
		"Referer": "https://emweb.securities.eastmoney.com/",
	}

	resp, err := utils.GetWithHeaders(url, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求财务分析指标失败: %w", err)
	}

	json := gjson.ParseBytes(resp.Body())
	dataPath := "result.data"
	result := json.Get(dataPath)

	if !result.Exists() || len(result.Array()) == 0 {
		return nil, fmt.Errorf("未找到财务分析指标数据")
	}

	var indicators []StockFinancialAnalysisIndicatorEmItem
	result.ForEach(func(_, value gjson.Result) bool {
		item := StockFinancialAnalysisIndicatorEmItem{
			ReportDate:        value.Get("REPORT_DATE").String(),
			ROE:               utils.MustFloat64(value.Get("ROE").String()),
			ROA:               utils.MustFloat64(value.Get("ROA").String()),
			GrossProfitMargin: utils.MustFloat64(value.Get("GrossProfitMargin").String()),
			NetProfitMargin:   utils.MustFloat64(value.Get("NetProfitMargin").String()),
		}
		indicators = append(indicators, item)
		return true
	})

	return indicators, nil
}

// StockHistoryDividend 新浪财经-历史分红
//
// 获取所有股票的历史分红数据
//
// 返回:
//   - []StockHistoryDividend: 历史分红数据列表
//   - error: 错误信息
//
// 示例:
//
//	dividends, err := stock_fundamental.StockHistoryDividend()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, d := range dividends {
//	    fmt.Printf("%s(%s): 累计股息=%.2f, 分红次数=%d\n", d.Name, d.Code, d.TotalDividend, d.DividendCount)
//	}
func StockHistoryDividend() ([]StockHistoryDividendItem, error) {
	url := "https://vip.stock.finance.sina.com.cn/q/go.php/vInvestConsult/kind/lsfh/index.phtml"
	params := map[string]string{
		"p":   "1",
		"num": "50000",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求历史分红失败: %w", err)
	}

	// 解析HTML表格
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	var dividends []StockHistoryDividendItem

	// 查找数据表格
	doc.Find("table").Each(func(i int, table *goquery.Selection) {
		// 通常第一个表格是数据表
		if i == 0 {
			table.Find("tr").Each(func(rowIdx int, row *goquery.Selection) {
				// 跳过表头
				if rowIdx == 0 {
					return
				}

				var cells []string
				row.Find("td").Each(func(j int, cell *goquery.Selection) {
					text := strings.TrimSpace(cell.Text())
					cells = append(cells, text)
				})

				if len(cells) >= 8 {
					// 代码填充为6位
					code := cells[0]
					for len(code) < 6 {
						code = "0" + code
					}

					listDate, _ := utils.ParseDate(cells[2])

					dividend := StockHistoryDividendItem{
						Code:            code,
						Name:            cells[1],
						ListDate:        listDate,
						TotalDividend:   utils.MustFloat64(cells[3]),
						AvgDividend:     utils.MustFloat64(cells[4]),
						DividendCount:   utils.MustInt(cells[5]),
						FinancingAmount: utils.MustFloat64(cells[6]),
						FinancingCount:  utils.MustInt(cells[7]),
					}
					dividends = append(dividends, dividend)
				}
			})
		}
	})

	return dividends, nil
}

// StockHistoryDividendDetail 新浪财经-分红配股详情
//
// 获取指定股票的分红或配股详情
//
// 参数:
//   - symbol: 股票代码，如 "000002"
//   - indicator: 类型，"分红" 或 "配股"
//   - date: 具体日期（可选），如 "1994-12-24"，为空则返回所有记录
//
// 返回:
//   - []StockHistoryDividendDetail: 分红配股详情列表
//   - error: 错误信息
//
// 示例:
//
//	details, err := stock_fundamental.StockHistoryDividendDetail("000002", "分红", "")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, d := range details {
//	    fmt.Printf("%s: 派息=%.2f, 送股=%.2f\n", d.AnnounceDate, d.Dividend, d.BonusShare)
//	}
func StockHistoryDividendDetail(symbol, indicator, date string) ([]StockHistoryDividendDetailItem, error) {
	if symbol == "" {
		return nil, fmt.Errorf("股票代码不能为空")
	}

	url := fmt.Sprintf("https://vip.stock.finance.sina.com.cn/corp/go.php/vISSUE_ShareBonus/stockid/%s.phtml", symbol)

	resp, err := utils.Get(url, nil)
	if err != nil {
		return nil, fmt.Errorf("请求分红详情失败: %w", err)
	}

	// 解析HTML
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	var details []StockHistoryDividendDetailItem
	tableIndex := 12 // 分红表格在索引12
	if indicator == "配股" {
		tableIndex = 13 // 配股表格在索引13
	}

	// 查找表格
	currentTable := 0
	doc.Find("table").Each(func(i int, table *goquery.Selection) {
		if currentTable == tableIndex {
			// 检查是否有数据
			firstRowText := table.Find("tr").First().Text()
			if strings.Contains(firstRowText, "暂时没有数据") {
				return
			}

			table.Find("tr").Each(func(rowIdx int, row *goquery.Selection) {
				// 跳过表头
				if rowIdx == 0 {
					return
				}

				var cells []string
				row.Find("td").Each(func(j int, cell *goquery.Selection) {
					text := utils.TrimSpace(cell.Text())
					cells = append(cells, text)
				})

				if len(cells) < 8 {
					return
				}

				// 解析日期
				announceDate, _ := utils.ParseDate(cells[0])
				exDate, _ := utils.ParseDate(cells[5])
				recordDate, _ := utils.ParseDate(cells[6])
				listingDate, _ := utils.ParseDate(cells[7])

				detail := StockHistoryDividendDetailItem{
					AnnounceDate:   announceDate,
					BonusShare:     utils.MustFloat64(cells[1]),
					TransferShare:  utils.MustFloat64(cells[2]),
					Dividend:       utils.MustFloat64(cells[3]),
					Status:         cells[4],
					ExDividendDate: exDate,
					RecordDate:     recordDate,
					ListingDate:    listingDate,
				}

				// 如果指定了日期，只返回匹配的记录
				if date == "" || cells[0] == date {
					details = append(details, detail)
				}
			})
		}
		currentTable++
	})

	// 如果指定了日期且需要详细信息
	if date != "" && len(details) > 0 {
		// 可以添加获取详细信息的逻辑
		// 这里简化处理，直接返回基本信息
	}

	return details, nil
}

// StockIpoInfo 新浪财经-新股发行
//
// 获取指定股票的新股发行信息
//
// 参数:
//   - stock: 股票代码，如 "600004"
//
// 返回:
//   - []StockIpoInfo: 新股发行信息列表
//   - error: 错误信息
//
// 示例:
//
//	info, err := stock_fundamental.StockIpoInfo("600004")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range info {
//	    fmt.Printf("%s: %s\n", item.Item, item.Value)
//	}
func StockIpoInfo(stock string) ([]StockIpoInfoItem, error) {
	if stock == "" {
		return nil, fmt.Errorf("股票代码不能为空")
	}

	url := fmt.Sprintf("https://vip.stock.finance.sina.com.cn/corp/go.php/vISSUE_NewStock/stockid/%s.phtml", stock)

	resp, err := utils.Get(url, nil)
	if err != nil {
		return nil, fmt.Errorf("请求新股发行信息失败: %w", err)
	}

	// 解析HTML
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	var infoList []StockIpoInfoItem

	// 查找表格（通常在第12个表格）
	currentTable := 0
	doc.Find("table").Each(func(i int, table *goquery.Selection) {
		if currentTable == 12 {
			table.Find("tr").Each(func(rowIdx int, row *goquery.Selection) {
				var cells []string
				row.Find("td").Each(func(j int, cell *goquery.Selection) {
					text := utils.TrimSpace(cell.Text())
					cells = append(cells, text)
				})

				if len(cells) == 2 {
					info := StockIpoInfoItem{
						Item:  cells[0],
						Value: cells[1],
					}
					infoList = append(infoList, info)
				}
			})
		}
		currentTable++
	})

	return infoList, nil
}

// StockMainStockHolder 新浪财经-股本股东-主要股东
//
// 获取指定股票的主要股东信息
//
// 参数:
//   - stock: 股票代码，如 "600004"
//
// 返回:
//   - []StockMainStockHolder: 主要股东信息列表
//   - error: 错误信息
//
// 示例:
//
//	holders, err := stock_fundamental.StockMainStockHolder("600004")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, h := range holders {
//	    fmt.Printf("%s: 持股=%.2f万股, 比例=%.2f%%\n", h.HolderName, h.Holdings, h.HoldingRatio)
//	}
func StockMainStockHolder(stock string) ([]StockMainStockHolderItem, error) {
	if stock == "" {
		return nil, fmt.Errorf("股票代码不能为空")
	}

	url := fmt.Sprintf("https://vip.stock.finance.sina.com.cn/corp/go.php/vCI_StockHolder/stockid/%s.phtml", stock)

	resp, err := utils.Get(url, nil)
	if err != nil {
		return nil, fmt.Errorf("请求主要股东信息失败: %w", err)
	}

	// 解析HTML
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	var holders []StockMainStockHolderItem

	// 查找表格（通常在第13个表格）
	currentTable := 0
	doc.Find("table").Each(func(i int, table *goquery.Selection) {
		if currentTable == 13 {
			var endDate string

			table.Find("tr").Each(func(rowIdx int, row *goquery.Selection) {
				cellsText := row.Find("td").First().Text()
				cellsText = utils.TrimSpace(cellsText)

				// 检查是否是截止日期行
				if strings.HasPrefix(cellsText, "截至日期") || strings.HasPrefix(cellsText, "截止日期") {
					// 提取日期
					var cells []string
					row.Find("td").Each(func(j int, cell *goquery.Selection) {
						text := utils.TrimSpace(cell.Text())
						cells = append(cells, text)
					})
					if len(cells) >= 2 {
						endDate = cells[1]
					}
					return
				}

				// 解析股东数据行（通常有5列以上）
				var cells []string
				row.Find("td").Each(func(j int, cell *goquery.Selection) {
					text := utils.TrimSpace(cell.Text())
					cells = append(cells, text)
				})

				// 股东数据行通常有编号、股东名称、持股数量、持股比例等列
				if len(cells) >= 4 {
					// 第一列通常是编号
					if cells[0] != "1" && cells[0] != "2" && cells[0] != "3" &&
						cells[0] != "4" && cells[0] != "5" && cells[0] != "6" &&
						cells[0] != "7" && cells[0] != "8" && cells[0] != "9" &&
						cells[0] != "10" {
						return
					}

					// 清理持股比例（去除↓符号和多余字符）
					ratioStr := cells[3]
					ratioStr = strings.TrimSpace(strings.TrimSuffix(strings.TrimSpace(ratioStr), "↓"))
					ratioStr = strings.TrimSpace(strings.TrimSuffix(strings.TrimSpace(ratioStr), "↑"))

					endDateParsed, _ := utils.ParseDate(endDate)

					holder := StockMainStockHolderItem{
						Code:         stock,
						Name:         "", // 需要从其他接口获取
						HolderName:   cells[1],
						Holdings:     utils.MustFloat64(cells[2]),
						HoldingRatio: utils.MustFloat64(ratioStr),
						EndDate:      endDateParsed,
					}
					holders = append(holders, holder)
				}
			})
		}
		currentTable++
	})

	return holders, nil
}

// StockFundStockHolder 新浪财经-股本股东-基金持股
//
// 获取指定股票的基金持股信息
//
// 参数:
//   - symbol: 股票代码，如 "600004"
//
// 返回:
//   - []StockFundStockHolder: 基金持股信息列表
//   - error: 错误信息
//
// 示例:
//
//	funds, err := stock_fundamental.StockFundStockHolder("601318")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, f := range funds {
//	    fmt.Printf("%s(%s): 持仓=%.2f万股, 比例=%.2f%%\n", f.FundName, f.FundCode, f.Holdings, f.HoldingRatio)
//	}
func StockFundStockHolder(symbol string) ([]StockFundStockHolderItem, error) {
	if symbol == "" {
		return nil, fmt.Errorf("股票代码不能为空")
	}

	url := fmt.Sprintf("https://vip.stock.finance.sina.com.cn/corp/go.php/vCI_FundStockHolder/stockid/%s.phtml", symbol)

	resp, err := utils.Get(url, nil)
	if err != nil {
		return nil, fmt.Errorf("请求基金持股信息失败: %w", err)
	}

	// 解析HTML
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	var funds []StockFundStockHolderItem

	// 查找表格（通常在第13个表格）
	currentTable := 0
	doc.Find("table").Each(func(i int, table *goquery.Selection) {
		if currentTable == 13 {
			var endDate string

			table.Find("tr").Each(func(rowIdx int, row *goquery.Selection) {
				cellsText := row.Find("td").First().Text()
				cellsText = utils.TrimSpace(cellsText)

				// 检查是否是截止日期行
				if strings.HasPrefix(cellsText, "截止日期") {
					// 提取日期
					var cells []string
					row.Find("td").Each(func(j int, cell *goquery.Selection) {
						text := utils.TrimSpace(cell.Text())
						cells = append(cells, text)
					})
					if len(cells) >= 2 {
						endDate = cells[1]
					}
					return
				}

				// 解析基金数据行（通常有6列以上）
				var cells []string
				row.Find("td").Each(func(j int, cell *goquery.Selection) {
					text := utils.TrimSpace(cell.Text())
					cells = append(cells, text)
				})

				// 基金数据行通常有基金名称、基金代码、持仓数量等列
				if len(cells) >= 6 {
					// 跳过非数据行
					if cells[0] == "" || cells[1] == "" {
						return
					}

					endDateParsed, _ := utils.ParseDate(endDate)

					fund := StockFundStockHolderItem{
						Code:          symbol,
						Name:          "", // 需要从其他接口获取
						FundName:      cells[0],
						FundCode:      cells[1],
						Holdings:      utils.MustFloat64(cells[2]),
						HoldingRatio:  utils.MustFloat64(cells[3]),
						MarketValue:   utils.MustFloat64(cells[4]),
						NetValueRatio: utils.MustFloat64(cells[5]),
						EndDate:       endDateParsed,
					}
					funds = append(funds, fund)
				}
			})
		}
		currentTable++
	})

	return funds, nil
}
