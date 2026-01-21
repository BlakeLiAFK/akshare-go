package stock

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// 雪球Token
const XQAToken = "ac2a78f80e88c34c3386da6345e0f7562548dd15"

// 新浪科创板URL配置
const (
	ZhSinaKcbStockURL       = "http://vip.stock.finance.sina.com.cn/quotes_service/api/json_v2.php/Market_Center.getHQNodeData"
	ZhSinaKcbStockCountURL  = "http://vip.stock.finance.sina.com.cn/quotes_service/api/json_v2.php/Market_Center.getHQNodeStockCount?node=kcb"
	ZhSinaKcbStockHistURL   = "https://quotes.sina.cn/cn/api/jsonp.php/var_%s%s=/KC_MarketDataService.getKLineData?symbol=%s"
	ZhSinaKcbStockAmountURL = "https://stock.finance.sina.com.cn/stock/api/jsonp.php/var%20KKE_ShareAmount_%s=/StockService.getAmountBySymbol?_=20&symbol=%s"
	ZhSinaKcbStockHfqURL    = "https://finance.sina.com.cn/realstock/company/%s/hfq.js"
	ZhSinaKcbStockQfqURL    = "https://finance.sina.com.cn/realstock/company/%s/qfq.js"
)

// 新浪A股URL配置
const (
	ZhSinaAStockURL           = "http://vip.stock.finance.sina.com.cn/quotes_service/api/json_v2.php/Market_Center.getHQNodeData"
	ZhSinaAStockCountURL      = "http://vip.stock.finance.sina.com.cn/quotes_service/api/json_v2.php/Market_Center.getHQNodeStockCount?node=hs_a"
	ZhSinaAStockHistURL       = "https://finance.sina.com.cn/realstock/company/%s/hisdata_klc2/klc_kl.js"
	ZhSinaAStockAmountURL     = "https://stock.finance.sina.com.cn/stock/api/jsonp.php/var%20KKE_ShareAmount_%s=/StockService.getAmountBySymbol?_=20&symbol=%s"
	ZhSinaAStockHfqURL        = "https://finance.sina.com.cn/realstock/company/%s/hfq.js"
	ZhSinaAStockQfqURL        = "https://finance.sina.com.cn/realstock/company/%s/qfq.js"
	ZhSinaAStockAmountPageURL = "https://money.finance.sina.com.cn/corp/go.php/vCI_StockStructureHistory/stockid/%s/stocktype/TotalStock.phtml"
)

// 新浪美股URL配置
const (
	USSinaStockHistQfqURL = "https://finance.sina.com.cn/us_stock/company/reinstatement/%s_qfq.js"
	USSinaStockHistURL    = "https://finance.sina.com.cn/us_stock/company/hisdata/klc_kl_%s.js"
	USSinaStockListURL    = "http://stock.finance.sina.com.cn/usstock/api/jsonp.php/IO.XSRV2.CallbackList[%s]/US_CategoryService.getList"
)

// 新浪B股URL配置
const (
	ZhSinaBStockURL       = "http://vip.stock.finance.sina.com.cn/quotes_service/api/json_v2.php/Market_Center.getHQNodeData"
	ZhSinaBStockCountURL  = "http://vip.stock.finance.sina.com.cn/quotes_service/api/json_v2.php/Market_Center.getHQNodeStockCount?node=hs_b"
	ZhSinaBStockHistURL   = "https://finance.sina.com.cn/realstock/company/%s/hisdata_klc2/klc_kl.js"
	ZhSinaBStockAmountURL = "https://stock.finance.sina.com.cn/stock/api/jsonp.php/var%20KKE_ShareAmount_%s=/StockService.getAmountBySymbol?_=20&symbol=%s"
)

// 东方财富URL配置
const (
	EmStockBoardConceptURL  = "http://push2.eastmoney.com/api/qt/clist/get"
	EmStockBoardIndustryURL = "http://push2.eastmoney.com/api/qt/clist/get"
	EmStockFundFlowURL      = "http://push2.eastmoney.com/api/qt/stock/fflow/kline/get"
	EmStockInfoURL          = "http://push2.eastmoney.com/api/qt/stock/get"
	EmStockProfileURL       = "http://push2.eastmoney.com/api/qt/stock/get"
	EmStockZtListURL        = "http://push2.eastmoney.com/api/qt/clist/get"
)

// 腾讯URL配置
const (
	TxStockAhURL   = "http://qt.gtimg.cn/q=sh%s,sz%s,hk%s"
	TxStockTickURL = "http://web.ifzq.gtimg.cn/appstock/app/HSFQV2/GetHSFQData.cgi"
)

// 巨潮资讯URL配置
const (
	CninfoStockIpoURL      = "http://static.cninfo.com.cn/static/common/cninfo_new/data/ipo/ipo_list.json"
	CninfoStockNewURL      = "http://static.cninfo.com.cn/static/common/cninfo_new/data/new_stock/new_stock_list.json"
	CninfoStockDividendURL = "http://static.cninfo.com.cn/static/common/cninfo_new/data/dividend/dividend_list.json"
)

// 默认请求参数
var (
	ZhSinaKcbStockPayload = map[string]string{
		"page":   "1",
		"num":    "80",
		"sort":   "symbol",
		"asc":    "1",
		"node":   "kcb",
		"symbol": "",
		"_s_r_a": "auto",
	}

	ZhSinaAStockPayload = map[string]string{
		"page":   "1",
		"num":    "80",
		"sort":   "symbol",
		"asc":    "1",
		"node":   "hs_a",
		"symbol": "",
		"_s_r_a": "page",
	}

	ZhSinaBStockPayload = map[string]string{
		"page":   "1",
		"num":    "80",
		"sort":   "symbol",
		"asc":    "1",
		"node":   "hs_b",
		"symbol": "",
		"_s_r_a": "page",
	}

	USSinaStockDictPayload = map[string]string{
		"page":   "2",
		"num":    "20",
		"sort":   "",
		"asc":    "0",
		"market": "",
		"id":     "",
	}
)

// 东方财富请求参数
var (
	EmStockBoardPayload = map[string]string{
		"fid":    "f62",
		"po":     "1",
		"pz":     "500",
		"pn":     "1",
		"np":     "1",
		"fltt":   "2",
		"invt":   "2",
		"fs":     "m:0+t:6,m:0+t:80,m:1+t:2,m:1+t:23",
		"fields": "f12,f14,f2,f3,f62,f184,f66,f69,f72,f75,f78,f81,f84,f87,f204,f205,f206,f207,f208,f209,f210,f211,f212,f213,f214,f215",
	}

	EmStockFundFlowPayload = map[string]string{
		"fields1": "f1,f2,f3,f7",
		"fields2": "f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61,f62,f63",
		"ut":      "b2884a393a59ad64002292a3e90d46a5",
	}

	EmStockInfoPayload = map[string]string{
		"fields": "f57,f58,f59,f60,f61,f62,f63,f64,f65,f67,f70,f71,f74,f75,f76,f77,f78,f79,f80,f81,f82,f84,f85,f86,f87,f88,f89,f90,f91,f92,f93,f95,f96,f97,f98,f99,f100,f101,f102,f103,f104,f105,f106,f107,f108,f109,f110,f111,f112,f113,f114,f115,f116,f117,f118,f119,f120,f121,f122,f123,f124,f125,f126,f127,f128,f129",
		"ut":     "b2884a393a59ad64002292a3e90d46a5",
	}
)

// 市场代码映射
var MarketCodeMap = map[string]string{
	"sh": "1",   // 上海证券交易所
	"sz": "0",   // 深圳证券交易所
	"hk": "116", // 香港交易所
	"us": "105", // 美国市场
}

// 股票类型映射
var StockTypeMap = map[string]string{
	"hs_a": "A股",
	"hs_b": "B股",
	"kcb":  "科创板",
	"cyb":  "创业板",
	"zxb":  "中小板",
}

// 交易日相关函数

// IsTradingDay 判断是否为交易日
func IsTradingDay(date string) bool {
	// 简化实现，实际应该查询交易日历
	if len(date) != 8 {
		return false
	}

	// 检查是否为周末（简化判断）
	year, _ := strconv.Atoi(date[:4])
	month, _ := strconv.Atoi(date[4:6])
	day, _ := strconv.Atoi(date[6:8])

	t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	weekday := t.Weekday()

	return weekday != time.Saturday && weekday != time.Sunday
}

// GetPreviousTradingDay 获取前一个交易日
func GetPreviousTradingDay(date string) string {
	if len(date) != 8 {
		return ""
	}

	year, _ := strconv.Atoi(date[:4])
	month, _ := strconv.Atoi(date[4:6])
	day, _ := strconv.Atoi(date[6:8])

	t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)

	// 向前查找交易日
	for i := 1; i <= 7; i++ {
		prev := t.AddDate(0, 0, -i)
		prevStr := prev.Format("20060102")
		if IsTradingDay(prevStr) {
			return prevStr
		}
	}

	return ""
}

// FormatStockCode 格式化股票代码
func FormatStockCode(symbol, market string) string {
	symbol = strings.ToUpper(strings.TrimSpace(symbol))
	market = strings.ToLower(strings.TrimSpace(market))

	switch market {
	case "sh", "sz":
		return market + symbol
	case "hk":
		return symbol + ".HK"
	case "us":
		return symbol
	default:
		return symbol
	}
}

// ParseStockCode 解析股票代码
func ParseStockCode(code string) (symbol, market string) {
	code = strings.ToUpper(strings.TrimSpace(code))

	if strings.HasSuffix(code, ".HK") {
		return strings.TrimSuffix(code, ".HK"), "hk"
	} else if strings.HasPrefix(code, "SH") {
		return strings.TrimPrefix(code, "SH"), "sh"
	} else if strings.HasPrefix(code, "SZ") {
		return strings.TrimPrefix(code, "SZ"), "sz"
	} else if len(code) == 6 {
		// 6位数字代码，需要判断市场
		if strings.HasPrefix(code, "6") {
			return code, "sh"
		} else if strings.HasPrefix(code, "0") || strings.HasPrefix(code, "3") {
			return code, "sz"
		}
	}

	return code, ""
}

// GetStockName 获取股票名称（简化实现）
func GetStockName(symbol, market string) string {
	// 简化实现，实际应该调用API查询
	code := FormatStockCode(symbol, market)

	// 一些常见的股票名称映射
	stockNames := map[string]string{
		"sh600000": "浦发银行",
		"sh600036": "招商银行",
		"sz000001": "平安银行",
		"sz000002": "万科A",
		"hk00700":  "腾讯控股",
		"hk00941":  "中国移动",
		"usAAPL":   "苹果",
		"usGOOGL":  "谷歌",
	}

	if name, exists := stockNames[code]; exists {
		return name
	}

	return code
}

// ValidateStockCode 验证股票代码格式
func ValidateStockCode(code string) bool {
	if code == "" {
		return false
	}

	code = strings.ToUpper(strings.TrimSpace(code))

	// A股代码：6位数字
	if matched := len(code) == 6 && isAllDigits(code); matched {
		return true
	}

	// 港股代码：5位数字 + .HK
	if strings.HasSuffix(code, ".HK") {
		prefix := strings.TrimSuffix(code, ".HK")
		return len(prefix) == 5 && isAllDigits(prefix)
	}

	// 美股代码：字母
	return isAllLetters(code)
}

// isAllDigits 检查字符串是否全为数字
func isAllDigits(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

// isAllLetters 检查字符串是否全为字母
func isAllLetters(s string) bool {
	for _, r := range s {
		if !((r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')) {
			return false
		}
	}
	return true
}

// BuildURL 构建完整URL
func BuildURL(baseURL string, params map[string]string) string {
	if len(params) == 0 {
		return baseURL
	}

	var paramPairs []string
	for key, value := range params {
		if value != "" {
			paramPairs = append(paramPairs, fmt.Sprintf("%s=%s", key, value))
		}
	}

	return baseURL + "?" + strings.Join(paramPairs, "&")
}
