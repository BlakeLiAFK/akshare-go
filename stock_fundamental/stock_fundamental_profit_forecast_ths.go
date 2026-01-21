package stock_fundamental

import (
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/PuerkitoBio/goquery"
)

// StockProfitForecastThsEpsItem 同花顺盈利预测-预测年报每股收益项
type StockProfitForecastThsEpsItem struct {
	Year   string  `json:"year"`    // 年度
	OrgNum int     `json:"org_num"` // 预测机构数
	Avg    float64 `json:"avg"`     // 平均值
	Max    float64 `json:"max"`     // 最大值
	Min    float64 `json:"min"`     // 最小值
}

// StockProfitForecastThsProfitItem 同花顺盈利预测-预测年报净利润项
type StockProfitForecastThsProfitItem struct {
	Year   string  `json:"year"`    // 年度
	OrgNum int     `json:"org_num"` // 预测机构数
	Avg    float64 `json:"avg"`     // 平均值(亿)
	Max    float64 `json:"max"`     // 最大值(亿)
	Min    float64 `json:"min"`     // 最小值(亿)
}

// StockProfitForecastThsOrgItem 同花顺盈利预测-业绩预测详表-机构项
type StockProfitForecastThsOrgItem struct {
	OrgName    string  `json:"org_name"`    // 机构名称
	ReportDate string  `json:"report_date"` // 报告日期
	EpsAvg     float64 `json:"eps_avg"`     // 预测年报每股收益-平均
	EpsMax     float64 `json:"eps_max"`     // 预测年报每股收益-最大
	EpsMin     float64 `json:"eps_min"`     // 预测年报每股收益-最小
	ProfitAvg  float64 `json:"profit_avg"`  // 预测年报净利润-平均(亿)
	ProfitMax  float64 `json:"profit_max"`  // 预测年报净利润-最大(亿)
	ProfitMin  float64 `json:"profit_min"`  // 预测年报净利润-最小(亿)
}

// StockProfitForecastThsDetailItem 同花顺盈利预测-业绩预测详表-详细指标预测项
type StockProfitForecastThsDetailItem struct {
	Indicator string            `json:"indicator"` // 指标名称
	Data      map[string]string `json:"data"`      // 各年度数据
}

// StockProfitForecastThsDetail 同花顺-盈利预测（详细版）
// symbol: 股票代码
// indicator: 指标类型，可选 "预测年报每股收益", "预测年报净利润", "业绩预测详表-机构", "业绩预测详表-详细指标预测"
func StockProfitForecastThsDetail(symbol, indicator string) (interface{}, error) {
	url := fmt.Sprintf("https://basic.10jqka.com.cn/new/%s/worth.html", symbol)
	headers := map[string]string{
		"User-Agent": utils.DefaultUserAgent,
	}

	resp, err := utils.GetWithHeaders(url, nil, headers)
	if err != nil {
		return nil, fmt.Errorf("获取盈利预测数据失败: %w", err)
	}

	html := resp.String()
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	// 检查是否无机构预测
	noForecast := strings.Contains(html, "本年度暂无机构做出业绩预测")

	tables := doc.Find("table")
	tableCount := tables.Length()

	switch indicator {
	case "预测年报每股收益":
		if noForecast {
			return []StockProfitForecastThsEpsItem{}, nil
		}
		return parseEpsTable(tables.Eq(0))

	case "预测年报净利润":
		if noForecast {
			return []StockProfitForecastThsProfitItem{}, nil
		}
		if tableCount > 1 {
			return parseProfitTable(tables.Eq(1))
		}
		return []StockProfitForecastThsProfitItem{}, nil

	case "业绩预测详表-机构":
		var tableIdx int
		if noForecast {
			tableIdx = 0
		} else {
			tableIdx = 2
		}
		if tableCount > tableIdx {
			return parseOrgTable(tables.Eq(tableIdx))
		}
		return []StockProfitForecastThsOrgItem{}, nil

	case "业绩预测详表-详细指标预测":
		var tableIdx int
		if noForecast {
			tableIdx = 1
		} else {
			tableIdx = 3
		}
		if tableCount > tableIdx {
			return parseDetailTable(tables.Eq(tableIdx))
		}
		return []StockProfitForecastThsDetailItem{}, nil

	default:
		return nil, fmt.Errorf("不支持的指标类型: %s", indicator)
	}
}

// parseEpsTable 解析预测年报每股收益表格
func parseEpsTable(table *goquery.Selection) ([]StockProfitForecastThsEpsItem, error) {
	var items []StockProfitForecastThsEpsItem

	table.Find("tbody tr").Each(func(i int, tr *goquery.Selection) {
		tds := tr.Find("td")
		if tds.Length() >= 5 {
			item := StockProfitForecastThsEpsItem{
				Year:   strings.TrimSpace(tds.Eq(0).Text()),
				OrgNum: int(utils.MustParseInt(tds.Eq(1).Text())),
				Avg:    utils.MustParseFloat(tds.Eq(2).Text()),
				Max:    utils.MustParseFloat(tds.Eq(3).Text()),
				Min:    utils.MustParseFloat(tds.Eq(4).Text()),
			}
			items = append(items, item)
		}
	})

	return items, nil
}

// parseProfitTable 解析预测年报净利润表格
func parseProfitTable(table *goquery.Selection) ([]StockProfitForecastThsProfitItem, error) {
	var items []StockProfitForecastThsProfitItem

	table.Find("tbody tr").Each(func(i int, tr *goquery.Selection) {
		tds := tr.Find("td")
		if tds.Length() >= 5 {
			item := StockProfitForecastThsProfitItem{
				Year:   strings.TrimSpace(tds.Eq(0).Text()),
				OrgNum: int(utils.MustParseInt(tds.Eq(1).Text())),
				Avg:    utils.MustParseFloat(tds.Eq(2).Text()),
				Max:    utils.MustParseFloat(tds.Eq(3).Text()),
				Min:    utils.MustParseFloat(tds.Eq(4).Text()),
			}
			items = append(items, item)
		}
	})

	return items, nil
}

// parseOrgTable 解析业绩预测详表-机构表格
func parseOrgTable(table *goquery.Selection) ([]StockProfitForecastThsOrgItem, error) {
	var items []StockProfitForecastThsOrgItem

	table.Find("tbody tr").Each(func(i int, tr *goquery.Selection) {
		tds := tr.Find("td")
		if tds.Length() >= 8 {
			item := StockProfitForecastThsOrgItem{
				OrgName:    strings.TrimSpace(tds.Eq(0).Text()),
				ReportDate: strings.TrimSpace(tds.Eq(1).Text()),
				EpsAvg:     utils.MustParseFloat(tds.Eq(2).Text()),
				EpsMax:     utils.MustParseFloat(tds.Eq(3).Text()),
				EpsMin:     utils.MustParseFloat(tds.Eq(4).Text()),
				ProfitAvg:  utils.MustParseFloat(tds.Eq(5).Text()),
				ProfitMax:  utils.MustParseFloat(tds.Eq(6).Text()),
				ProfitMin:  utils.MustParseFloat(tds.Eq(7).Text()),
			}
			items = append(items, item)
		}
	})

	return items, nil
}

// parseDetailTable 解析业绩预测详表-详细指标预测表格
func parseDetailTable(table *goquery.Selection) ([]StockProfitForecastThsDetailItem, error) {
	var items []StockProfitForecastThsDetailItem

	// 获取表头（年度列）
	var headers []string
	table.Find("thead tr th").Each(func(i int, th *goquery.Selection) {
		text := strings.TrimSpace(th.Text())
		text = strings.ReplaceAll(text, "（", "-")
		text = strings.ReplaceAll(text, "）", "")
		headers = append(headers, text)
	})

	// 解析数据行
	table.Find("tbody tr").Each(func(i int, tr *goquery.Selection) {
		tds := tr.Find("td")
		if tds.Length() > 0 {
			item := StockProfitForecastThsDetailItem{
				Indicator: strings.TrimSpace(tds.Eq(0).Text()),
				Data:      make(map[string]string),
			}

			tds.Each(func(j int, td *goquery.Selection) {
				if j > 0 && j < len(headers) {
					item.Data[headers[j]] = strings.TrimSpace(td.Text())
				}
			})

			items = append(items, item)
		}
	})

	return items, nil
}
