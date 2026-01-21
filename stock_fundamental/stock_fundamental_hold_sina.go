package stock_fundamental

import (
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
)

// StockInstituteHold 新浪财经-股票-机构持股一览表
//
// 获取指定报告期的机构持股一览表数据
//
// 参数:
//   - symbol: 报告期代码，格式如 "20201"（2020年一季报）
//   - 最后一位: {"一季报":1, "中报":2, "三季报":3, "年报":4}
//   - 前四位: 年份
//
// 返回:
//   - []StockInstituteHoldItem: 机构持股一览表数据
//   - error: 错误信息
//
// 示例:
//
//	holds, err := stock_fundamental.StockInstituteHold("20201")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range holds {
//	    fmt.Printf("%s(%s): 机构数=%.0f, 持股比例=%.2f%%\n", item.Name, item.Code, item.InstituteCount, item.HoldingRatio)
//	}
func StockInstituteHold(symbol string) ([]StockInstituteHoldItem, error) {
	if symbol == "" || len(symbol) < 5 {
		return nil, fmt.Errorf("报告期代码格式错误，应为5位数字，如 \"20201\"")
	}

	// 解析年份和季度
	reportDate := symbol[:len(symbol)-1]
	quarter := symbol[len(symbol)-1:]

	url := "https://vip.stock.finance.sina.com.cn/q/go.php/vComStockHold/kind/jgcg/index.phtml"
	params := map[string]string{
		"p":          "1",
		"num":        "10000",
		"reportdate": reportDate,
		"quarter":    quarter,
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求机构持股一览表失败: %w", err)
	}

	// 解析HTML表格
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	var holds []StockInstituteHoldItem

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

				// 需要至少9列（包含明细列）
				if len(cells) < 9 {
					return
				}

				// 证券代码填充为6位
				code := cells[0]
				for len(code) < 6 {
					code = "0" + code
				}

				hold := StockInstituteHoldItem{
					Code:                   code,
					Name:                   cells[1],
					InstituteCount:         utils.MustFloat64(cells[2]),
					InstituteCountChange:   utils.MustFloat64(cells[3]),
					HoldingRatio:           utils.MustFloat64(cells[4]),
					HoldingRatioChange:     utils.MustFloat64(cells[5]),
					CirculationRatio:       utils.MustFloat64(cells[6]),
					CirculationRatioChange: utils.MustFloat64(cells[7]),
				}
				holds = append(holds, hold)
			})
		}
	})

	if len(holds) == 0 {
		return nil, fmt.Errorf("未找到机构持股数据")
	}

	return holds, nil
}

// StockInstituteHoldDetail 新浪财经-股票-机构持股详情
//
// 获取指定股票和报告期的机构持股详情数据
//
// 参数:
//   - stock: 股票代码，如 "600433"
//   - quarter: 报告期代码，格式如 "20201"（2020年一季报）
//
// 返回:
//   - []StockInstituteHoldDetailItem: 机构持股详情数据
//   - error: 错误信息
//
// 示例:
//
//	details, err := stock_fundamental.StockInstituteHoldDetail("600433", "20201")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range details {
//	    fmt.Printf("%s: 持股=%.2f万股, 比例=%.2f%%\n", item.InstituteName, item.Holdings, item.HoldingRatio)
//	}
func StockInstituteHoldDetail(stock, quarter string) ([]StockInstituteHoldDetailItem, error) {
	if stock == "" {
		return nil, fmt.Errorf("股票代码不能为空")
	}
	if quarter == "" || len(quarter) < 5 {
		return nil, fmt.Errorf("报告期代码格式错误，应为5位数字，如 \"20201\"")
	}

	url := "https://vip.stock.finance.sina.com.cn/q/api/jsonp.php/var%20details=/ComStockHoldService.getJGCGDetail"
	params := map[string]string{
		"symbol":  stock,
		"quarter": quarter,
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求机构持股详情失败: %w", err)
	}

	// 解析JSONP响应
	jsonStr, err := parseJSONP(resp.String())
	if err != nil {
		return nil, fmt.Errorf("解析JSONP失败: %w", err)
	}

	json := gjson.Parse(jsonStr)
	dataPath := "data"
	result := json.Get(dataPath)

	if !result.Exists() {
		return nil, fmt.Errorf("未找到机构持股详情数据")
	}

	var details []StockInstituteHoldDetailItem

	// 遍历所有机构类型的数据 (fund, socialSecurity, insurance, qfii)
	result.ForEach(func(instituteTypeKey, instituteTypeValue gjson.Result) bool {
		// 跳过 stock 和 total 字段
		instituteType := instituteTypeKey.String()
		if instituteType == "stock" || instituteType == "total" {
			return true
		}

		// 遍历该类型下的所有机构
		instituteTypeValue.ForEach(func(orgKey, orgValue gjson.Result) bool {
			// 跳过 total 字段
			if orgKey.String() == "total" {
				return true
			}

			// orgValue 是一个对象，包含各个字段
			detail := StockInstituteHoldDetailItem{
				InstituteType:          translateInstituteType(instituteType),
				InstituteCode:          orgValue.Get("orgCode").String(),
				InstituteName:          orgValue.Get("orgName").String(),
				InstituteFullName:      orgValue.Get("orgFullName").String(),
				Holdings:               utils.MustFloat64(orgValue.Get("stockAmount").String()),
				LatestHoldings:         utils.MustFloat64(orgValue.Get("stockAmountLast").String()),
				HoldingRatio:           utils.MustFloat64(orgValue.Get("stockPercent").String()),
				LatestHoldingRatio:     utils.MustFloat64(orgValue.Get("stockPercentLast").String()),
				CirculationRatio:       utils.MustFloat64(orgValue.Get("currentPercent").String()),
				LatestCirculationRatio: utils.MustFloat64(orgValue.Get("lastCurrentPercent").String()),
				HoldingRatioChange:     utils.MustFloat64(orgValue.Get("stockPercentBalance").String()),
				CirculationRatioChange: utils.MustFloat64(orgValue.Get("currentPercentBalance").String()),
			}
			details = append(details, detail)
			return true
		})
		return true
	})

	if len(details) == 0 {
		return nil, fmt.Errorf("未找到机构持股详情数据")
	}

	return details, nil
}

// translateInstituteType 翻译机构类型
func translateInstituteType(instituteType string) string {
	switch instituteType {
	case "fund":
		return "基金"
	case "socialSecurity":
		return "全国社保"
	case "qfii":
		return "QFII"
	case "insurance":
		return "保险"
	default:
		return instituteType
	}
}
