package futures

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
)

// FuturesDailyBar 期货日线行情
type FuturesDailyBar struct {
	Symbol       string  `json:"symbol"`        // 合约代码
	Date         string  `json:"date"`          // 日期
	Open         float64 `json:"open"`          // 开盘价
	High         float64 `json:"high"`          // 最高价
	Low          float64 `json:"low"`           // 最低价
	Close        float64 `json:"close"`         // 收盘价
	Volume       float64 `json:"volume"`        // 成交量
	OpenInterest float64 `json:"open_interest"` // 持仓量
	Turnover     float64 `json:"turnover"`      // 成交额
	Settle       float64 `json:"settle"`        // 结算价
	PreSettle    float64 `json:"pre_settle"`    // 前结算价
	Variety      string  `json:"variety"`       // 品种
}

// GetCFFEXDaily 中国金融期货交易所-日频率交易数据
//
// 数据源: http://www.cffex.com.cn/rtj/
//
// 参数:
//   - date: 交易日，格式 "20100416"，数据开始时间为 20100416
//
// 返回:
//   - []FuturesDailyBar: 日频率交易数据
//   - error: 错误信息
func GetCFFEXDaily(date string) ([]FuturesDailyBar, error) {
	if len(date) < 8 {
		return nil, fmt.Errorf("日期格式错误: %s", date)
	}

	url := fmt.Sprintf("http://www.cffex.com.cn/sj/historysj/%s/zip/%s.zip", date[:6], date[:6])
	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
	}

	resp, err := utils.GetWithHeaders(url, nil, headers)
	if err != nil {
		return nil, fmt.Errorf("请求中金所数据失败: %w", err)
	}

	zipReader, err := zip.NewReader(bytes.NewReader(resp.Body()), int64(len(resp.Body())))
	if err != nil {
		return nil, fmt.Errorf("解压数据失败: %w", err)
	}

	csvFileName := fmt.Sprintf("%s_1.csv", date)
	var csvData []byte
	for _, f := range zipReader.File {
		if f.Name == csvFileName {
			rc, err := f.Open()
			if err != nil {
				return nil, fmt.Errorf("打开CSV文件失败: %w", err)
			}
			csvData, err = io.ReadAll(rc)
			rc.Close()
			if err != nil {
				return nil, fmt.Errorf("读取CSV文件失败: %w", err)
			}
			break
		}
	}

	if csvData == nil {
		return []FuturesDailyBar{}, nil
	}

	lines := strings.Split(string(csvData), "\n")
	if len(lines) < 2 {
		return []FuturesDailyBar{}, nil
	}

	varietyRe := regexp.MustCompile(`[a-zA-Z_]+`)
	var result []FuturesDailyBar

	for i, line := range lines[1:] {
		fields := strings.Split(line, ",")
		if len(fields) < 11 {
			continue
		}

		symbol := strings.TrimSpace(fields[0])
		if symbol == "小计" || symbol == "合计" || strings.Contains(symbol, "IO") ||
			strings.Contains(symbol, "MO") || strings.Contains(symbol, "HO") {
			continue
		}

		variety := ""
		if matches := varietyRe.FindString(symbol); matches != "" {
			variety = matches
		}

		bar := FuturesDailyBar{
			Symbol:       symbol,
			Date:         date,
			Open:         utils.MustParseFloat(fields[1]),
			High:         utils.MustParseFloat(fields[2]),
			Low:          utils.MustParseFloat(fields[3]),
			Volume:       utils.MustParseFloat(fields[4]),
			Turnover:     utils.MustParseFloat(fields[5]),
			OpenInterest: utils.MustParseFloat(fields[6]),
			Variety:      variety,
		}

		// 根据字段数量设置收盘价等
		if len(fields) >= 15 {
			bar.Close = utils.MustParseFloat(fields[8])
			bar.Settle = utils.MustParseFloat(fields[9])
			bar.PreSettle = utils.MustParseFloat(fields[10])
		} else if len(fields) >= 14 {
			bar.Close = utils.MustParseFloat(fields[8])
			bar.Settle = utils.MustParseFloat(fields[9])
			bar.PreSettle = utils.MustParseFloat(fields[10])
		}

		result = append(result, bar)
		_ = i // 防止未使用警告
	}

	return result, nil
}

// GetGFEXDaily 广州期货交易所-日频率-量价数据
//
// 数据源: http://www.gfex.com.cn/gfex/rihq/hqsj_tjsj.shtml
//
// 参数:
//   - date: 日期，格式 "20221223"
//
// 返回:
//   - []FuturesDailyBar: 广州期货交易所日频率量价数据
//   - error: 错误信息
func GetGFEXDaily(date string) ([]FuturesDailyBar, error) {
	url := "http://www.gfex.com.cn/u/interfacesWebTiDayQuotes/loadList"

	payload := map[string]string{
		"trade_date": date,
		"trade_type": "0",
	}

	headers := map[string]string{
		"Accept":           "application/json, text/javascript, */*; q=0.01",
		"Content-Type":     "application/x-www-form-urlencoded; charset=UTF-8",
		"User-Agent":       "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		"X-Requested-With": "XMLHttpRequest",
	}

	resp, err := utils.PostFormWithHeaders(url, payload, headers)
	if err != nil {
		return nil, fmt.Errorf("请求广期所数据失败: %w", err)
	}

	var dataJSON struct {
		Data []struct {
			Variety      string `json:"variety"`
			VarietyOrder string `json:"varietyOrder"`
			DelivMonth   string `json:"delivMonth"`
			Open         string `json:"open"`
			High         string `json:"high"`
			Low          string `json:"low"`
			Close        string `json:"close"`
			Volumn       string `json:"volumn"`
			OpenInterest string `json:"openInterest"`
			Turnover     string `json:"turnover"`
			ClearPrice   string `json:"clearPrice"`
			LastClear    string `json:"lastClear"`
		} `json:"data"`
	}

	if err := json.Unmarshal(resp.Body(), &dataJSON); err != nil {
		return nil, fmt.Errorf("解析广期所数据失败: %w", err)
	}

	var result []FuturesDailyBar
	for _, item := range dataJSON.Data {
		if strings.Contains(item.Variety, "小计") || strings.Contains(item.Variety, "总计") {
			continue
		}

		symbol := strings.ToUpper(item.VarietyOrder) + item.DelivMonth
		variety := strings.ToUpper(item.VarietyOrder)

		result = append(result, FuturesDailyBar{
			Symbol:       symbol,
			Date:         date,
			Open:         utils.MustParseFloat(item.Open),
			High:         utils.MustParseFloat(item.High),
			Low:          utils.MustParseFloat(item.Low),
			Close:        utils.MustParseFloat(item.Close),
			Volume:       utils.MustParseFloat(item.Volumn),
			OpenInterest: utils.MustParseFloat(item.OpenInterest),
			Turnover:     utils.MustParseFloat(item.Turnover),
			Settle:       utils.MustParseFloat(item.ClearPrice),
			PreSettle:    utils.MustParseFloat(item.LastClear),
			Variety:      variety,
		})
	}

	return result, nil
}

// GetINEDaily 上海国际能源交易中心-日频率-量价数据
//
// 数据源: https://www.ine.cn/statements/daily/?paramid=kx
//
// 参数:
//   - date: 日期，格式 "20241129"
//
// 返回:
//   - []FuturesDailyBar: 上海国际能源交易中心日频率量价数据
//   - error: 错误信息
func GetINEDaily(date string) ([]FuturesDailyBar, error) {
	url := fmt.Sprintf("https://www.ine.cn/data/tradedata/future/dailydata/kx%s.dat", date)
	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
	}

	resp, err := utils.GetWithHeaders(url, nil, headers)
	if err != nil {
		return nil, fmt.Errorf("请求上期能源数据失败: %w", err)
	}

	var dataJSON struct {
		OCurinstrument []struct {
			ProductGroupID     string `json:"PRODUCTGROUPID"`
			ProductID          string `json:"PRODUCTID"`
			DeliveryMonth      string `json:"DELIVERYMONTH"`
			ProductName        string `json:"PRODUCTNAME"`
			OpenPrice          any    `json:"OPENPRICE"`
			HighestPrice       any    `json:"HIGHESTPRICE"`
			LowestPrice        any    `json:"LOWESTPRICE"`
			ClosePrice         any    `json:"CLOSEPRICE"`
			Volume             any    `json:"VOLUME"`
			OpenInterest       any    `json:"OPENINTEREST"`
			Turnover           any    `json:"TURNOVER"`
			SettlementPrice    any    `json:"SETTLEMENTPRICE"`
			PreSettlementPrice any    `json:"PRESETTLEMENTPRICE"`
		} `json:"o_curinstrument"`
	}

	if err := json.Unmarshal(resp.Body(), &dataJSON); err != nil {
		return nil, fmt.Errorf("解析上期能源数据失败: %w", err)
	}

	var result []FuturesDailyBar
	for _, item := range dataJSON.OCurinstrument {
		if item.DeliveryMonth == "小计" || strings.Contains(item.ProductName, "总计") {
			continue
		}

		variety := strings.ToUpper(strings.TrimSpace(item.ProductGroupID))
		if variety == "" {
			parts := strings.Split(item.ProductID, "_")
			if len(parts) > 0 {
				variety = strings.ToUpper(strings.TrimSpace(parts[0]))
			}
		}

		symbol := variety + item.DeliveryMonth
		if strings.Contains(symbol, "efp") || symbol == "总计" {
			continue
		}

		result = append(result, FuturesDailyBar{
			Symbol:       symbol,
			Date:         date,
			Open:         toFloat64(item.OpenPrice),
			High:         toFloat64(item.HighestPrice),
			Low:          toFloat64(item.LowestPrice),
			Close:        toFloat64(item.ClosePrice),
			Volume:       toFloat64(item.Volume),
			OpenInterest: toFloat64(item.OpenInterest),
			Turnover:     toFloat64(item.Turnover),
			Settle:       toFloat64(item.SettlementPrice),
			PreSettle:    toFloat64(item.PreSettlementPrice),
			Variety:      variety,
		})
	}

	return result, nil
}

// GetSHFEDaily 上海期货交易所-日频率-量价数据
//
// 数据源: https://tsite.shfe.com.cn/statements/dataview.html?paramid=kx
//
// 参数:
//   - date: 日期，格式 "20220415"
//
// 返回:
//   - []FuturesDailyBar: 上海期货交易所日频率量价数据
//   - error: 错误信息
func GetSHFEDaily(date string) ([]FuturesDailyBar, error) {
	url := fmt.Sprintf("https://www.shfe.com.cn/data/dailydata/%sdailyTimePrice.dat", date)
	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
	}

	resp, err := utils.GetWithHeaders(url, nil, headers)
	if err != nil {
		return nil, fmt.Errorf("请求上期所数据失败: %w", err)
	}

	var dataJSON struct {
		OCurinstrument []struct {
			ProductGroupID     string `json:"PRODUCTGROUPID"`
			ProductID          string `json:"PRODUCTID"`
			DeliveryMonth      string `json:"DELIVERYMONTH"`
			OpenPrice          any    `json:"OPENPRICE"`
			HighestPrice       any    `json:"HIGHESTPRICE"`
			LowestPrice        any    `json:"LOWESTPRICE"`
			ClosePrice         any    `json:"CLOSEPRICE"`
			Volume             any    `json:"VOLUME"`
			OpenInterest       any    `json:"OPENINTEREST"`
			Turnover           any    `json:"TURNOVER"`
			SettlementPrice    any    `json:"SETTLEMENTPRICE"`
			PreSettlementPrice any    `json:"PRESETTLEMENTPRICE"`
		} `json:"o_curinstrument"`
	}

	if err := json.Unmarshal(resp.Body(), &dataJSON); err != nil {
		return nil, fmt.Errorf("解析上期所数据失败: %w", err)
	}

	var result []FuturesDailyBar
	for _, item := range dataJSON.OCurinstrument {
		if item.DeliveryMonth == "小计" || item.DeliveryMonth == "合计" || item.DeliveryMonth == "" {
			continue
		}

		variety := strings.ToUpper(strings.TrimSpace(item.ProductGroupID))
		if variety == "" {
			parts := strings.Split(item.ProductID, "_")
			if len(parts) > 0 {
				variety = strings.ToUpper(strings.TrimSpace(parts[0]))
			}
		}

		symbol := variety + item.DeliveryMonth
		if strings.Contains(symbol, "efp") {
			continue
		}

		result = append(result, FuturesDailyBar{
			Symbol:       symbol,
			Date:         date,
			Open:         toFloat64(item.OpenPrice),
			High:         toFloat64(item.HighestPrice),
			Low:          toFloat64(item.LowestPrice),
			Close:        toFloat64(item.ClosePrice),
			Volume:       toFloat64(item.Volume),
			OpenInterest: toFloat64(item.OpenInterest),
			Turnover:     toFloat64(item.Turnover),
			Settle:       toFloat64(item.SettlementPrice),
			PreSettle:    toFloat64(item.PreSettlementPrice),
			Variety:      variety,
		})
	}

	return result, nil
}

// GetDCEDaily 大连商品交易所日交易数据
//
// 数据源: http://www.dce.com.cn/dalianshangpin/xqsj/tjsj26/rtj/rxq/index.html
//
// 参数:
//   - date: 交易日，格式 "20251027"
//
// 返回:
//   - []FuturesDailyBar: 大连商品交易所日交易数据
//   - error: 错误信息
func GetDCEDaily(date string) ([]FuturesDailyBar, error) {
	url := "http://www.dce.com.cn/dcereport/publicweb/dailystat/dayQuotes"
	payload := map[string]any{
		"contractId":     "",
		"lang":           "zh",
		"optionSeries":   "",
		"statisticsType": "0",
		"tradeDate":      date,
		"tradeType":      "1",
		"varietyId":      "all",
	}

	headers := map[string]string{
		"Content-Type": "application/json",
		"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
	}

	resp, err := utils.PostJSONWithHeaders(url, nil, payload, headers)
	if err != nil {
		return nil, fmt.Errorf("请求大商所数据失败: %w", err)
	}

	var dataJSON struct {
		Data []struct {
			Variety      string `json:"variety"`
			ContractID   string `json:"contractId"`
			Open         any    `json:"open"`
			High         any    `json:"high"`
			Low          any    `json:"low"`
			Close        any    `json:"close"`
			LastClear    any    `json:"lastClear"`
			ClearPrice   any    `json:"clearPrice"`
			Volumn       any    `json:"volumn"`
			OpenInterest any    `json:"openInterest"`
			Turnover     any    `json:"turnover"`
		} `json:"data"`
	}

	if err := json.Unmarshal(resp.Body(), &dataJSON); err != nil {
		return nil, fmt.Errorf("解析大商所数据失败: %w", err)
	}

	// 大商所品种中文名到代码的映射
	dceMap := map[string]string{
		"豆一":    "A",
		"豆二":    "B",
		"豆粕":    "M",
		"豆油":    "Y",
		"棕榈油":   "P",
		"玉米":    "C",
		"玉米淀粉":  "CS",
		"鸡蛋":    "JD",
		"生猪":    "LH",
		"纤维板":   "FB",
		"胶合板":   "BB",
		"聚乙烯":   "L",
		"聚氯乙烯":  "V",
		"聚丙烯":   "PP",
		"乙二醇":   "EG",
		"苯乙烯":   "EB",
		"焦炭":    "J",
		"焦煤":    "JM",
		"铁矿石":   "I",
		"液化石油气": "PG",
		"粳米":    "RR",
		"粳稻":    "JR",
	}

	var result []FuturesDailyBar
	for _, item := range dataJSON.Data {
		if strings.Contains(item.Variety, "小计") || strings.Contains(item.Variety, "总计") {
			continue
		}

		variety := dceMap[item.Variety]
		if variety == "" {
			variety = item.Variety
		}

		result = append(result, FuturesDailyBar{
			Symbol:       item.ContractID,
			Date:         date,
			Open:         toFloat64(item.Open),
			High:         toFloat64(item.High),
			Low:          toFloat64(item.Low),
			Close:        toFloat64(item.Close),
			Volume:       toFloat64(item.Volumn),
			OpenInterest: toFloat64(item.OpenInterest),
			Turnover:     toFloat64(item.Turnover),
			Settle:       toFloat64(item.ClearPrice),
			PreSettle:    toFloat64(item.LastClear),
			Variety:      variety,
		})
	}

	return result, nil
}

// GetCZCEDaily 郑州商品交易所-日频率-量价数据
//
// 数据源: http://www.czce.com.cn/cn/jysj/mrhq/H770301index_1.htm
//
// 参数:
//   - date: 日期，格式 "20050525"，日期需要大于 20100824
//
// 返回:
//   - []FuturesDailyBar: 郑州商品交易所日频率量价数据
//   - error: 错误信息
func GetCZCEDaily(date string) ([]FuturesDailyBar, error) {
	dateTime, _ := time.Parse("20060102", date)
	cutoffDate, _ := time.Parse("20060102", "20151111")

	var url string
	if dateTime.After(cutoffDate) {
		url = fmt.Sprintf("http://www.czce.com.cn/cn/DFSStaticFiles/Future/%s/%s/FutureDataDaily.htm",
			date[:4], date)
	} else {
		url = fmt.Sprintf("http://www.czce.com.cn/cn/exchange/jyxx/hq/hq%s.html", date)
	}

	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
	}

	resp, err := utils.GetWithHeaders(url, nil, headers)
	if err != nil {
		return nil, fmt.Errorf("请求郑商所数据失败: %w", err)
	}

	text := string(resp.Body())
	if strings.Contains(text, "您的访问出错了") || strings.Contains(text, "无期权每日行情交易记录") {
		return []FuturesDailyBar{}, nil
	}

	// 解析数据
	lines := strings.Split(text, "\n")
	varietyRe := regexp.MustCompile(`^([A-Za-z]+)`)

	var result []FuturesDailyBar
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "小") {
			continue
		}

		fields := strings.Split(strings.ReplaceAll(line, " ", ""), "|")
		if len(fields) < 10 {
			continue
		}

		symbol := strings.TrimSpace(fields[0])
		if !varietyRe.MatchString(symbol) {
			continue
		}

		matches := varietyRe.FindStringSubmatch(symbol)
		variety := ""
		if len(matches) > 1 {
			variety = matches[1]
		}

		result = append(result, FuturesDailyBar{
			Symbol:       symbol,
			Date:         date,
			Open:         parseFloatWithComma(fields[2]),
			High:         parseFloatWithComma(fields[3]),
			Low:          parseFloatWithComma(fields[4]),
			Close:        parseFloatWithComma(fields[5]),
			Settle:       parseFloatWithComma(fields[6]),
			Volume:       parseFloatWithComma(fields[9]),
			OpenInterest: parseFloatWithComma(fields[10]),
			PreSettle:    parseFloatWithComma(fields[1]),
			Variety:      variety,
		})
	}

	return result, nil
}

// GetFuturesDaily 交易所日交易数据
//
// 参数:
//   - startDate: 开始日期，格式 "20220208"
//   - endDate: 结束日期，格式 "20220208"
//   - market: 交易所，可选 "CFFEX" 中金所, "CZCE" 郑商所, "SHFE" 上期所, "DCE" 大商所, "INE" 上期能源, "GFEX" 广期所
//
// 返回:
//   - []FuturesDailyBar: 交易所日交易数据
//   - error: 错误信息
func GetFuturesDaily(startDate, endDate, market string) ([]FuturesDailyBar, error) {
	var getFunc func(string) ([]FuturesDailyBar, error)

	switch strings.ToUpper(market) {
	case "CFFEX":
		getFunc = GetCFFEXDaily
	case "CZCE":
		getFunc = GetCZCEDaily
	case "SHFE":
		getFunc = GetSHFEDaily
	case "DCE":
		getFunc = GetDCEDaily
	case "INE":
		getFunc = GetINEDaily
	case "GFEX":
		getFunc = GetGFEXDaily
	default:
		return nil, fmt.Errorf("无效的交易所代码: %s", market)
	}

	start, err := time.Parse("20060102", startDate)
	if err != nil {
		return nil, fmt.Errorf("开始日期格式错误: %w", err)
	}

	end, err := time.Parse("20060102", endDate)
	if err != nil {
		return nil, fmt.Errorf("结束日期格式错误: %w", err)
	}

	var result []FuturesDailyBar
	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		dateStr := d.Format("20060102")
		data, err := getFunc(dateStr)
		if err != nil {
			continue // 跳过错误日期
		}
		result = append(result, data...)
	}

	// 过滤掉包含 efp 的合约
	var filtered []FuturesDailyBar
	for _, bar := range result {
		if !strings.Contains(bar.Symbol, "efp") {
			filtered = append(filtered, bar)
		}
	}

	return filtered, nil
}

// parseFloatWithComma 解析带逗号的浮点数
func parseFloatWithComma(s string) float64 {
	s = strings.ReplaceAll(s, ",", "")
	s = strings.TrimSpace(s)
	if s == "" || s == "-" {
		return 0
	}
	v, _ := strconv.ParseFloat(s, 64)
	return v
}
