package futures

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/PuerkitoBio/goquery"
	"github.com/xuri/excelize/v2"
)

// ToSpotSHFE 上期所期转现
type ToSpotSHFE struct {
	Date     string  `json:"date"`     // 日期
	Contract string  `json:"contract"` // 合约
	Delivery float64 `json:"delivery"` // 交割量
	EFP      float64 `json:"efp"`      // 期转现量
}

// DeliveryDCE 大商所交割统计
type DeliveryDCE struct {
	Date     string  `json:"date"`     // 交割日期
	Variety  string  `json:"variety"`  // 品种
	Delivery float64 `json:"delivery"` // 交割量
	Amount   float64 `json:"amount"`   // 交割金额
}

// ToSpotDCE 大商所期转现
type ToSpotDCE struct {
	Date     string  `json:"date"`     // 期转现发生日期
	Contract string  `json:"contract"` // 合约代码
	Quantity float64 `json:"quantity"` // 期转现数量
}

// DeliveryMatchDCE 大商所交割配对
type DeliveryMatchDCE struct {
	Date        string  `json:"date"`         // 配对日期
	Contract    string  `json:"contract"`     // 合约代码
	MatchQty    float64 `json:"match_qty"`    // 配对手数
	SettlePrice float64 `json:"settle_price"` // 交割结算价
}

// ToSpotCZCE 郑商所期转现
type ToSpotCZCE struct {
	Contract string  `json:"contract"` // 合约代码
	Quantity float64 `json:"quantity"` // 合约数量
}

// DeliveryCZCE 郑商所月度交割
type DeliveryCZCE struct {
	Variety  string  `json:"variety"`  // 品种
	Quantity float64 `json:"quantity"` // 交割数量
	Amount   float64 `json:"amount"`   // 交割额
}

// DeliverySHFE 上期所交割情况
type DeliverySHFE struct {
	Variety       string  `json:"variety"`        // 品种
	MonthDelivery float64 `json:"month_delivery"` // 交割量-本月
	Ratio         float64 `json:"ratio"`          // 交割量-比重
	YearDelivery  float64 `json:"year_delivery"`  // 交割量-本年累计
	YoYChange     float64 `json:"yoy_change"`     // 交割量-累计同比
}

// shfeToSpotResponse SHFE期转现API响应
type shfeToSpotResponse struct {
	ExchangeDelivery []struct {
		Col1     string `json:"WBIESSION_ROWNUM"`
		Date     string `json:"DELIVERYDATE"`
		Delivery any    `json:"DELIVERY_QTY"`
		Col4     string `json:"VARCODE"`
		EFP      any    `json:"EFP_QTY"`
		Contract string `json:"DELIVERYID"`
		Col7     string `json:"VARNAME"`
		Col8     string `json:"EXCHANGENAMESC"`
	} `json:"ExchangeDelivery"`
}

// FuturesToSpotSHFE 上海期货交易所-期转现
//
// 数据源: https://tsite.shfe.com.cn/statements/dataview.html?paramid=kx
//
// 参数:
//   - date: 年月，格式 "202312"
//
// 返回:
//   - []ToSpotSHFE: 期转现数据
//   - error: 错误信息
func FuturesToSpotSHFE(date string) ([]ToSpotSHFE, error) {
	url := fmt.Sprintf("https://tsite.shfe.com.cn/data/instrument/ExchangeDelivery%s.dat", date)
	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
	}

	resp, err := utils.GetWithHeaders(url, nil, headers)
	if err != nil {
		return nil, fmt.Errorf("请求上期所期转现数据失败: %w", err)
	}

	var apiResp shfeToSpotResponse
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	var result []ToSpotSHFE
	for _, item := range apiResp.ExchangeDelivery {
		delivery, _ := utils.ToFloat64(item.Delivery)
		efp, _ := utils.ToFloat64(item.EFP)

		result = append(result, ToSpotSHFE{
			Date:     item.Date,
			Contract: item.Contract,
			Delivery: delivery,
			EFP:      efp,
		})
	}

	return result, nil
}

// FuturesDeliveryDCE 大连商品交易所-交割统计
//
// 数据源: http://www.dce.com.cn/dalianshangpin/xqsj/tjsj26/jgtj/jgsj/index.html
//
// 参数:
//   - date: 年月，格式 "202312"
//
// 返回:
//   - []DeliveryDCE: 交割统计数据
//   - error: 错误信息
func FuturesDeliveryDCE(date string) ([]DeliveryDCE, error) {
	url := "http://www.dce.com.cn/publicweb/quotesdata/delivery.html"

	endMonth := fmt.Sprintf("%d", mustAtoi(date)+1)

	params := map[string]string{
		"deliveryQuotes.variety":     "all",
		"year":                       "",
		"month":                      "",
		"deliveryQuotes.begin_month": date,
		"deliveryQuotes.end_month":   endMonth,
	}

	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
	}

	resp, err := utils.PostFormWithHeaders(url+"?"+buildQueryString(params), nil, headers)
	if err != nil {
		return nil, fmt.Errorf("请求大商所交割统计失败: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(resp.Body())))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	var result []DeliveryDCE
	doc.Find("table").First().Find("tr").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			return // 跳过表头
		}
		var row []string
		s.Find("td").Each(func(_ int, td *goquery.Selection) {
			row = append(row, strings.TrimSpace(td.Text()))
		})

		if len(row) >= 4 {
			variety := row[1]
			// 跳过小计和总计行
			if strings.Contains(variety, "小计") || strings.Contains(variety, "总计") {
				return
			}

			result = append(result, DeliveryDCE{
				Date:     strings.Split(row[0], ".")[0],
				Variety:  variety,
				Delivery: utils.MustParseFloat(strings.ReplaceAll(row[2], ",", "")),
				Amount:   utils.MustParseFloat(strings.ReplaceAll(row[3], ",", "")),
			})
		}
	})

	return result, nil
}

// FuturesToSpotDCE 大连商品交易所-期转现
//
// 数据源: http://www.dce.com.cn/dalianshangpin/xqsj/tjsj26/jgtj/qzxcx/index.html
//
// 参数:
//   - date: 年月，格式 "202312"
//
// 返回:
//   - []ToSpotDCE: 期转现数据
//   - error: 错误信息
func FuturesToSpotDCE(date string) ([]ToSpotDCE, error) {
	url := "http://www.dce.com.cn/publicweb/quotesdata/ftsDeal.html"

	params := map[string]string{
		"ftsDealQuotes.variety":     "all",
		"year":                      "",
		"month":                     "",
		"ftsDealQuotes.begin_month": date,
		"ftsDealQuotes.end_month":   date,
	}

	resp, err := utils.PostFormWithHeaders(url+"?"+buildQueryString(params), nil, nil)
	if err != nil {
		return nil, fmt.Errorf("请求大商所期转现数据失败: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(resp.Body())))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	var result []ToSpotDCE
	doc.Find("table").First().Find("tr").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			return // 跳过表头
		}
		var row []string
		s.Find("td").Each(func(_ int, td *goquery.Selection) {
			row = append(row, strings.TrimSpace(td.Text()))
		})

		if len(row) >= 3 {
			contract := row[1]
			// 跳过小计和总计行
			if strings.Contains(contract, "小计") || strings.Contains(contract, "总计") {
				return
			}

			result = append(result, ToSpotDCE{
				Date:     strings.Split(row[0], ".")[0],
				Contract: contract,
				Quantity: utils.MustParseFloat(strings.ReplaceAll(row[2], ",", "")),
			})
		}
	})

	return result, nil
}

// FuturesToSpotCZCE 郑州商品交易所-期转现统计
//
// 数据源: http://www.czce.com.cn/cn/jysj/qzxtj/H770311index_1.htm
//
// 参数:
//   - date: 年月日，格式 "20231228"
//
// 返回:
//   - []ToSpotCZCE: 期转现统计数据
//   - error: 错误信息
func FuturesToSpotCZCE(date string) ([]ToSpotCZCE, error) {
	url := fmt.Sprintf("http://www.czce.com.cn/cn/DFSStaticFiles/Future/%s/%s/FutureDataTrdtrades.xls", date[:4], date)
	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
	}

	resp, err := utils.GetWithHeaders(url, nil, headers)
	if err != nil {
		return nil, fmt.Errorf("请求郑商所期转现数据失败: %w", err)
	}

	f, err := excelize.OpenReader(strings.NewReader(string(resp.Body())))
	if err != nil {
		return nil, fmt.Errorf("解析Excel失败: %w", err)
	}
	defer f.Close()

	sheets := f.GetSheetList()
	if len(sheets) == 0 {
		return nil, fmt.Errorf("Excel文件没有工作表")
	}

	rows, err := f.GetRows(sheets[0])
	if err != nil {
		return nil, fmt.Errorf("读取工作表失败: %w", err)
	}

	var result []ToSpotCZCE
	for i, row := range rows {
		if i <= 1 || len(row) < 2 { // 跳过表头行
			continue
		}

		contract := row[0]
		// 跳过小计和合计行
		if strings.Contains(contract, "小计") || strings.Contains(contract, "合计") {
			continue
		}

		result = append(result, ToSpotCZCE{
			Contract: contract,
			Quantity: utils.MustParseFloat(strings.ReplaceAll(row[1], ",", "")),
		})
	}

	return result, nil
}

// FuturesDeliveryCZCE 郑州商品交易所-月度交割查询
//
// 数据源: http://www.czce.com.cn/cn/jysj/ydjgcx/H770316index_1.htm
//
// 参数:
//   - date: 年月日，格式 "20210112"
//
// 返回:
//   - []DeliveryCZCE: 月度交割数据
//   - error: 错误信息
func FuturesDeliveryCZCE(date string) ([]DeliveryCZCE, error) {
	url := fmt.Sprintf("http://www.czce.com.cn/cn/DFSStaticFiles/Future/%s/%s/FutureDataSettlematched.xls", date[:4], date)

	resp, err := utils.Get(url, nil)
	if err != nil {
		return nil, fmt.Errorf("请求郑商所月度交割数据失败: %w", err)
	}

	f, err := excelize.OpenReader(strings.NewReader(string(resp.Body())))
	if err != nil {
		return nil, fmt.Errorf("解析Excel失败: %w", err)
	}
	defer f.Close()

	sheets := f.GetSheetList()
	if len(sheets) == 0 {
		return nil, fmt.Errorf("Excel文件没有工作表")
	}

	rows, err := f.GetRows(sheets[0])
	if err != nil {
		return nil, fmt.Errorf("读取工作表失败: %w", err)
	}

	var result []DeliveryCZCE
	for i, row := range rows {
		if i <= 1 || len(row) < 3 { // 跳过表头行
			continue
		}

		result = append(result, DeliveryCZCE{
			Variety:  row[0],
			Quantity: utils.MustParseFloat(strings.ReplaceAll(row[1], ",", "")),
			Amount:   utils.MustParseFloat(strings.ReplaceAll(row[2], ",", "")),
		})
	}

	return result, nil
}

// shfeDeliveryResponse SHFE交割情况API响应
type shfeDeliveryResponse struct {
	OCurDelivery []struct {
		Variety       string `json:"PRODUCTNAME"`
		Code          string `json:"PRODUCTID"`
		Col3          string `json:"EXCHANGENAMESC"`
		MonthDelivery any    `json:"DELIVERYAMOUNT"`
		Ratio         any    `json:"DELIVERYAMOUNTRATIO"`
		YearDelivery  any    `json:"YEARDELIVERYAMOUNT"`
		YoYChange     any    `json:"YONLASTYEAR"`
	} `json:"o_curdelivery"`
}

// FuturesDeliverySHFE 上海期货交易所-交割情况表
//
// 数据源: https://tsite.shfe.com.cn/statements/dataview.html?paramid=kx
//
// 参数:
//   - date: 年月，格式 "202312"
//
// 返回:
//   - []DeliverySHFE: 交割情况数据
//   - error: 错误信息
func FuturesDeliverySHFE(date string) ([]DeliverySHFE, error) {
	url := fmt.Sprintf("https://tsite.shfe.com.cn/data/dailydata/%smonthvarietystatistics.dat", date)
	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
	}

	resp, err := utils.GetWithHeaders(url, nil, headers)
	if err != nil {
		return nil, fmt.Errorf("请求上期所交割情况失败: %w", err)
	}

	var apiResp shfeDeliveryResponse
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	var result []DeliverySHFE
	for _, item := range apiResp.OCurDelivery {
		monthDelivery, _ := utils.ToFloat64(item.MonthDelivery)
		ratio, _ := utils.ToFloat64(item.Ratio)
		yearDelivery, _ := utils.ToFloat64(item.YearDelivery)
		yoyChange, _ := utils.ToFloat64(item.YoYChange)

		result = append(result, DeliverySHFE{
			Variety:       item.Variety,
			MonthDelivery: monthDelivery,
			Ratio:         ratio,
			YearDelivery:  yearDelivery,
			YoYChange:     yoyChange,
		})
	}

	return result, nil
}

// mustAtoi 转换字符串为整数，失败返回0
func mustAtoi(s string) int {
	var n int
	fmt.Sscanf(s, "%d", &n)
	return n
}

// buildQueryString 构建查询字符串
func buildQueryString(params map[string]string) string {
	var parts []string
	for k, v := range params {
		parts = append(parts, fmt.Sprintf("%s=%s", k, v))
	}
	return strings.Join(parts, "&")
}
