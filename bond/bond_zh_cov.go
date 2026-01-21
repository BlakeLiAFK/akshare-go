package bond

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// BondZhHsCovSpotItem 新浪财经-沪深可转债实时行情项
type BondZhHsCovSpotItem struct {
	Symbol     string  `json:"symbol"`        // 代码
	Name       string  `json:"name"`          // 名称
	Trade      float64 `json:"trade"`         // 最新价
	PriceChg   float64 `json:"pricechange"`   // 涨跌额
	ChangePct  float64 `json:"changepercent"` // 涨跌幅
	Buy        float64 `json:"buy"`           // 买入价
	Sell       float64 `json:"sell"`          // 卖出价
	Settlement float64 `json:"settlement"`    // 昨收
	Open       float64 `json:"open"`          // 开盘价
	High       float64 `json:"high"`          // 最高价
	Low        float64 `json:"low"`           // 最低价
	Volume     float64 `json:"volume"`        // 成交量
	Amount     float64 `json:"amount"`        // 成交额
}

// BondZhHsCovDailyItem 新浪财经-沪深可转债历史行情项
type BondZhHsCovDailyItem struct {
	Date   string  `json:"date"`   // 日期
	Open   float64 `json:"open"`   // 开盘价
	High   float64 `json:"high"`   // 最高价
	Low    float64 `json:"low"`    // 最低价
	Close  float64 `json:"close"`  // 收盘价
	Volume float64 `json:"volume"` // 成交量
}

// BondZhHsCovMinItem 东方财富-可转债分时行情项
type BondZhHsCovMinItem struct {
	Time      string  `json:"time"`       // 时间
	Open      float64 `json:"open"`       // 开盘
	Close     float64 `json:"close"`      // 收盘
	High      float64 `json:"high"`       // 最高
	Low       float64 `json:"low"`        // 最低
	Volume    float64 `json:"volume"`     // 成交量
	Amount    float64 `json:"amount"`     // 成交额
	Latest    float64 `json:"latest"`     // 最新价 (1分钟K线)
	Amplitude float64 `json:"amplitude"`  // 振幅 (其他K线)
	ChangePct float64 `json:"change_pct"` // 涨跌幅
	ChangeAmt float64 `json:"change_amt"` // 涨跌额
	Turnover  float64 `json:"turnover"`   // 换手率
}

// BondZhCovItem 东方财富-可转债数据项
type BondZhCovItem struct {
	BondCode           string  `json:"bond_code"`            // 债券代码
	BondName           string  `json:"bond_name"`            // 债券简称
	SubDate            string  `json:"sub_date"`             // 申购日期
	SubCode            string  `json:"sub_code"`             // 申购代码
	SubLimit           float64 `json:"sub_limit"`            // 申购上限
	StockCode          string  `json:"stock_code"`           // 正股代码
	StockName          string  `json:"stock_name"`           // 正股简称
	StockPrice         float64 `json:"stock_price"`          // 正股价
	ConvertPrice       float64 `json:"convert_price"`        // 转股价
	ConvertValue       float64 `json:"convert_value"`        // 转股价值
	BondPrice          float64 `json:"bond_price"`           // 债现价
	ConvertPremiumRate float64 `json:"convert_premium_rate"` // 转股溢价率
	EquityRegDate      string  `json:"equity_reg_date"`      // 原股东配售-股权登记日
	PerShareAllot      float64 `json:"per_share_allot"`      // 原股东配售-每股配售额
	IssueScale         float64 `json:"issue_scale"`          // 发行规模
	LotteryDate        string  `json:"lottery_date"`         // 中签号发布日
	WinRate            float64 `json:"win_rate"`             // 中签率
	ListingDate        string  `json:"listing_date"`         // 上市时间
	CreditRating       string  `json:"credit_rating"`        // 信用评级
}

// BondCovComparisonItem 东方财富-可转债比价表项
type BondCovComparisonItem struct {
	Seq                 int     `json:"seq"`                   // 序号
	BondCode            string  `json:"bond_code"`             // 转债代码
	BondName            string  `json:"bond_name"`             // 转债名称
	BondPrice           float64 `json:"bond_price"`            // 转债最新价
	BondChangePct       float64 `json:"bond_change_pct"`       // 转债涨跌幅
	StockCode           string  `json:"stock_code"`            // 正股代码
	StockName           string  `json:"stock_name"`            // 正股名称
	StockPrice          float64 `json:"stock_price"`           // 正股最新价
	StockChangePct      float64 `json:"stock_change_pct"`      // 正股涨跌幅
	ConvertPrice        float64 `json:"convert_price"`         // 转股价
	ConvertValue        float64 `json:"convert_value"`         // 转股价值
	ConvertPremium      float64 `json:"convert_premium"`       // 转股溢价率
	PureBondPremium     float64 `json:"pure_bond_premium"`     // 纯债溢价率
	ResaleTrigPrice     float64 `json:"resale_trig_price"`     // 回售触发价
	RedeemTrigPrice     float64 `json:"redeem_trig_price"`     // 强赎触发价
	MaturityRedeemPrice float64 `json:"maturity_redeem_price"` // 到期赎回价
	PureBondValue       float64 `json:"pure_bond_value"`       // 纯债价值
	ConvertStartDate    string  `json:"convert_start_date"`    // 开始转股日
	ListingDate         string  `json:"listing_date"`          // 上市日期
	SubDate             string  `json:"sub_date"`              // 申购日期
}

// BondZhCovValueAnalysisItem 东方财富-可转债价值分析项
type BondZhCovValueAnalysisItem struct {
	Date            string  `json:"date"`              // 日期
	ClosePrice      float64 `json:"close_price"`       // 收盘价
	PureBondValue   float64 `json:"pure_bond_value"`   // 纯债价值
	ConvertValue    float64 `json:"convert_value"`     // 转股价值
	PureBondPremium float64 `json:"pure_bond_premium"` // 纯债溢价率
	ConvertPremium  float64 `json:"convert_premium"`   // 转股溢价率
}

// BondZhHsCovSpot 新浪财经-债券-沪深可转债的实时行情数据
// https://vip.stock.finance.sina.com.cn/mkt/#hskzz_z
func BondZhHsCovSpot() ([]BondZhHsCovSpotItem, error) {
	countURL := "https://vip.stock.finance.sina.com.cn/quotes_service/api/json_v2.php/Market_Center.getHQNodeStockCount"
	params := map[string]string{"node": "hskzz_z"}

	resp, err := utils.Get(countURL, params)
	if err != nil {
		return nil, fmt.Errorf("获取总数失败: %w", err)
	}

	totalCount := gjson.ParseBytes(resp.Body()).Int()
	pageCount := int(math.Ceil(float64(totalCount) / 80))

	var items []BondZhHsCovSpotItem

	dataURL := "https://vip.stock.finance.sina.com.cn/quotes_service/api/json_v2.php/Market_Center.getHQNodeData"
	for page := 1; page <= pageCount; page++ {
		params := map[string]string{
			"page":   fmt.Sprintf("%d", page),
			"num":    "80",
			"sort":   "symbol",
			"asc":    "1",
			"node":   "hskzz_z",
			"symbol": "",
			"_s_r_a": "page",
		}

		resp, err := utils.Get(dataURL, params)
		if err != nil {
			continue
		}

		dataArr := gjson.ParseBytes(resp.Body()).Array()
		for _, item := range dataArr {
			items = append(items, BondZhHsCovSpotItem{
				Symbol:     item.Get("symbol").String(),
				Name:       item.Get("name").String(),
				Trade:      item.Get("trade").Float(),
				PriceChg:   item.Get("pricechange").Float(),
				ChangePct:  item.Get("changepercent").Float(),
				Buy:        item.Get("buy").Float(),
				Sell:       item.Get("sell").Float(),
				Settlement: item.Get("settlement").Float(),
				Open:       item.Get("open").Float(),
				High:       item.Get("high").Float(),
				Low:        item.Get("low").Float(),
				Volume:     item.Get("volume").Float(),
				Amount:     item.Get("amount").Float(),
			})
		}
	}

	return items, nil
}

// BondZhHsCovDaily 新浪财经-债券-沪深可转债的历史行情数据
// symbol: 沪深可转债代码，如 "sh010107"
// https://vip.stock.finance.sina.com.cn/mkt/#hskzz_z
func BondZhHsCovDaily(symbol string) ([]BondZhHsCovDailyItem, error) {
	now := time.Now().Format("2006_01_02")
	url := fmt.Sprintf("https://finance.sina.com.cn/realstock/company/%s/hisdata/klc_kl.js?d=%s", symbol, now)

	resp, err := utils.GetWithHeaders(url, nil, map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
	})
	if err != nil {
		return nil, fmt.Errorf("获取历史行情失败: %w", err)
	}

	// 解析JS响应
	bodyStr := string(resp.Body())
	// 提取数据部分
	startIdx := strings.Index(bodyStr, "=")
	if startIdx == -1 {
		return nil, fmt.Errorf("解析响应失败")
	}
	dataStr := strings.TrimSpace(bodyStr[startIdx+1:])
	dataStr = strings.TrimSuffix(dataStr, ";")
	dataStr = strings.ReplaceAll(dataStr, `"`, `"`)

	result := gjson.Parse(dataStr)
	dataArr := result.Array()

	var items []BondZhHsCovDailyItem
	for _, item := range dataArr {
		items = append(items, BondZhHsCovDailyItem{
			Date:   item.Get("date").String(),
			Open:   item.Get("open").Float(),
			High:   item.Get("high").Float(),
			Low:    item.Get("low").Float(),
			Close:  item.Get("close").Float(),
			Volume: item.Get("volume").Float(),
		})
	}

	return items, nil
}

// BondZhHsCovMin 东方财富网-可转债-分时行情
// symbol: 转债代码，如 "sz128039"
// period: K线周期，可选 {"1", "5", "15", "30", "60"}
// adjust: 复权类型，可选 {"", "qfq", "hfq"}
// startDate: 开始日期，如 "1979-09-01 09:32:00"
// endDate: 结束日期，如 "2222-01-01 09:32:00"
// https://quote.eastmoney.com/concept/sz128039.html
func BondZhHsCovMin(symbol, period, adjust, startDate, endDate string) ([]BondZhHsCovMinItem, error) {
	marketType := map[string]string{"sh": "1", "sz": "0"}

	if period == "1" {
		url := "https://push2.eastmoney.com/api/qt/stock/trends2/get"
		params := map[string]string{
			"secid":   fmt.Sprintf("%s.%s", marketType[symbol[:2]], symbol[2:]),
			"fields1": "f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f11,f12,f13",
			"fields2": "f51,f52,f53,f54,f55,f56,f57,f58",
			"iscr":    "0",
			"iscca":   "0",
			"ut":      "f057cbcbce2a86e2866ab8877db1d059",
			"ndays":   "1",
		}

		resp, err := utils.Get(url, params)
		if err != nil {
			return nil, fmt.Errorf("获取分时行情失败: %w", err)
		}

		trends := gjson.Get(string(resp.Body()), "data.trends").Array()
		var items []BondZhHsCovMinItem

		for _, trend := range trends {
			parts := strings.Split(trend.String(), ",")
			if len(parts) >= 8 {
				item := BondZhHsCovMinItem{
					Time: parts[0],
				}
				fmt.Sscanf(parts[1], "%f", &item.Open)
				fmt.Sscanf(parts[2], "%f", &item.Close)
				fmt.Sscanf(parts[3], "%f", &item.High)
				fmt.Sscanf(parts[4], "%f", &item.Low)
				fmt.Sscanf(parts[5], "%f", &item.Volume)
				fmt.Sscanf(parts[6], "%f", &item.Amount)
				fmt.Sscanf(parts[7], "%f", &item.Latest)
				items = append(items, item)
			}
		}

		return filterByTime(items, startDate, endDate), nil
	}

	// 其他周期
	adjustMap := map[string]string{"": "0", "qfq": "1", "hfq": "2"}
	url := "https://push2his.eastmoney.com/api/qt/stock/kline/get"
	params := map[string]string{
		"secid":   fmt.Sprintf("%s.%s", marketType[symbol[:2]], symbol[2:]),
		"klt":     period,
		"fqt":     adjustMap[adjust],
		"lmt":     "66",
		"end":     "20500000",
		"iscca":   "1",
		"fields1": "f1,f2,f3,f4,f5",
		"fields2": "f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61",
		"ut":      "7eea3edcaed734bea9cbfc24409ed989",
		"forcect": "1",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("获取K线数据失败: %w", err)
	}

	klines := gjson.Get(string(resp.Body()), "data.klines").Array()
	var items []BondZhHsCovMinItem

	for _, kline := range klines {
		parts := strings.Split(kline.String(), ",")
		if len(parts) >= 11 {
			item := BondZhHsCovMinItem{
				Time: parts[0],
			}
			fmt.Sscanf(parts[1], "%f", &item.Open)
			fmt.Sscanf(parts[2], "%f", &item.Close)
			fmt.Sscanf(parts[3], "%f", &item.High)
			fmt.Sscanf(parts[4], "%f", &item.Low)
			fmt.Sscanf(parts[5], "%f", &item.Volume)
			fmt.Sscanf(parts[6], "%f", &item.Amount)
			fmt.Sscanf(parts[7], "%f", &item.Amplitude)
			fmt.Sscanf(parts[8], "%f", &item.ChangePct)
			fmt.Sscanf(parts[9], "%f", &item.ChangeAmt)
			fmt.Sscanf(parts[10], "%f", &item.Turnover)
			items = append(items, item)
		}
	}

	return filterByTime(items, startDate, endDate), nil
}

// filterByTime 按时间过滤数据
func filterByTime(items []BondZhHsCovMinItem, startDate, endDate string) []BondZhHsCovMinItem {
	if startDate == "" && endDate == "" {
		return items
	}

	var filtered []BondZhHsCovMinItem
	for _, item := range items {
		if (startDate == "" || item.Time >= startDate) && (endDate == "" || item.Time <= endDate) {
			filtered = append(filtered, item)
		}
	}
	return filtered
}

// BondZhHsCovPreMin 东方财富网-可转债-分时行情-盘前
// symbol: 转债代码，如 "sh113570"
// https://quote.eastmoney.com/concept/sz128039.html
func BondZhHsCovPreMin(symbol string) ([]BondZhHsCovMinItem, error) {
	marketType := map[string]string{"sh": "1", "sz": "0"}
	url := "https://push2.eastmoney.com/api/qt/stock/trends2/get"
	params := map[string]string{
		"fields1": "f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f11,f12,f13",
		"fields2": "f51,f52,f53,f54,f55,f56,f57,f58",
		"ndays":   "1",
		"iscr":    "1",
		"iscca":   "0",
		"secid":   fmt.Sprintf("%s.%s", marketType[symbol[:2]], symbol[2:]),
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("获取盘前分时行情失败: %w", err)
	}

	trends := gjson.Get(string(resp.Body()), "data.trends").Array()
	var items []BondZhHsCovMinItem

	for _, trend := range trends {
		parts := strings.Split(trend.String(), ",")
		if len(parts) >= 8 {
			item := BondZhHsCovMinItem{
				Time: parts[0],
			}
			fmt.Sscanf(parts[1], "%f", &item.Open)
			fmt.Sscanf(parts[2], "%f", &item.Close)
			fmt.Sscanf(parts[3], "%f", &item.High)
			fmt.Sscanf(parts[4], "%f", &item.Low)
			fmt.Sscanf(parts[5], "%f", &item.Volume)
			fmt.Sscanf(parts[6], "%f", &item.Amount)
			fmt.Sscanf(parts[7], "%f", &item.Latest)
			items = append(items, item)
		}
	}

	return items, nil
}

// BondZhCov 东方财富网-数据中心-新股数据-可转债数据
// https://data.eastmoney.com/kzz/default.html
func BondZhCov() ([]BondZhCovItem, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"sortColumns":  "PUBLIC_START_DATE",
		"sortTypes":    "-1",
		"pageSize":     "500",
		"pageNumber":   "1",
		"reportName":   "RPT_BOND_CB_LIST",
		"columns":      "ALL",
		"quoteColumns": "f2~01~CONVERT_STOCK_CODE~CONVERT_STOCK_PRICE,f235~10~SECURITY_CODE~TRANSFER_PRICE,f236~10~SECURITY_CODE~TRANSFER_VALUE,f2~10~SECURITY_CODE~CURRENT_BOND_PRICE,f237~10~SECURITY_CODE~TRANSFER_PREMIUM_RATIO,f239~10~SECURITY_CODE~RESALE_TRIG_PRICE,f240~10~SECURITY_CODE~REDEEM_TRIG_PRICE,f23~01~CONVERT_STOCK_CODE~PBV_RATIO",
		"source":       "WEB",
		"client":       "WEB",
	}

	var items []BondZhCovItem

	// 获取第一页确定总页数
	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("获取可转债数据失败: %w", err)
	}

	totalPage := gjson.Get(string(resp.Body()), "result.pages").Int()

	for page := int64(1); page <= totalPage; page++ {
		params["pageNumber"] = fmt.Sprintf("%d", page)
		resp, err := utils.Get(url, params)
		if err != nil {
			continue
		}

		dataArr := gjson.Get(string(resp.Body()), "result.data").Array()
		for _, item := range dataArr {
			bondPrice := item.Get("CURRENT_BOND_PRICE").Float()
			if bondPrice == 0 {
				bondPrice = 100
			}

			items = append(items, BondZhCovItem{
				BondCode:           item.Get("SECURITY_CODE").String(),
				BondName:           item.Get("SECURITY_NAME_ABBR").String(),
				SubDate:            item.Get("PUBLIC_START_DATE").String()[:10],
				SubCode:            item.Get("CORRECODE").String(),
				SubLimit:           item.Get("BONDSTART").Float(),
				StockCode:          item.Get("CONVERT_STOCK_CODE").String(),
				StockName:          item.Get("SECURITY_SHORT_NAME").String(),
				StockPrice:         item.Get("CONVERT_STOCK_PRICE").Float(),
				ConvertPrice:       item.Get("TRANSFER_PRICE").Float(),
				ConvertValue:       item.Get("TRANSFER_VALUE").Float(),
				BondPrice:          bondPrice,
				ConvertPremiumRate: item.Get("TRANSFER_PREMIUM_RATIO").Float(),
				EquityRegDate:      strings.Split(item.Get("SECURITY_START_DATE").String(), " ")[0],
				PerShareAllot:      item.Get("INITIAL_TRANSFER_PRICE").Float(),
				IssueScale:         item.Get("ACTUAL_ISSUE_SCALE").Float(),
				LotteryDate:        strings.Split(item.Get("BOND_START_DATE").String(), " ")[0],
				WinRate:            item.Get("ONLINE_GENERAL_LWR").Float(),
				ListingDate:        strings.Split(item.Get("LISTING_DATE").String(), " ")[0],
				CreditRating:       item.Get("RATING").String(),
			})
		}
	}

	return items, nil
}

// BondCovComparison 东方财富网-行情中心-债券市场-可转债比价表
// https://quote.eastmoney.com/center/fullscreenlist.html#convertible_comparison
func BondCovComparison() ([]BondCovComparisonItem, error) {
	url := "https://16.push2.eastmoney.com/api/qt/clist/get"
	params := map[string]string{
		"pn":     "1",
		"pz":     "100",
		"po":     "1",
		"np":     "1",
		"ut":     "bd1d9ddb04089700cf9c27f6f7426281",
		"fltt":   "2",
		"invt":   "2",
		"fid":    "f243",
		"fs":     "b:MK0354",
		"fields": "f1,f152,f2,f3,f12,f13,f14,f227,f228,f229,f230,f231,f232,f233,f234,f235,f236,f237,f238,f239,f240,f241,f242,f26,f243",
	}

	var items []BondCovComparisonItem
	page := 1

	for {
		params["pn"] = fmt.Sprintf("%d", page)
		resp, err := utils.Get(url, params)
		if err != nil {
			break
		}

		dataArr := gjson.Get(string(resp.Body()), "data.diff").Array()
		if len(dataArr) == 0 {
			break
		}

		for _, item := range dataArr {
			items = append(items, BondCovComparisonItem{
				Seq:                 int(item.Get("f1").Int()),
				BondCode:            item.Get("f12").String(),
				BondName:            item.Get("f14").String(),
				BondPrice:           item.Get("f2").Float(),
				BondChangePct:       item.Get("f3").Float(),
				StockCode:           item.Get("f234").String(),
				StockName:           item.Get("f233").String(),
				StockPrice:          item.Get("f231").Float(),
				StockChangePct:      item.Get("f232").Float(),
				ConvertPrice:        item.Get("f235").Float(),
				ConvertValue:        item.Get("f236").Float(),
				ConvertPremium:      item.Get("f237").Float(),
				PureBondPremium:     item.Get("f238").Float(),
				ResaleTrigPrice:     item.Get("f239").Float(),
				RedeemTrigPrice:     item.Get("f240").Float(),
				MaturityRedeemPrice: item.Get("f241").Float(),
				PureBondValue:       item.Get("f229").Float(),
				ConvertStartDate:    item.Get("f242").String(),
				ListingDate:         item.Get("f227").String(),
				SubDate:             item.Get("f243").String(),
			})
		}

		total := gjson.Get(string(resp.Body()), "data.total").Int()
		if int64(page*100) >= total {
			break
		}
		page++
	}

	return items, nil
}

// BondZhCovInfo 东方财富网-数据中心-新股数据-可转债详情
// symbol: 可转债代码，如 "123121"
// indicator: 指标类型，可选 {"基本信息", "中签号", "筹资用途", "重要日期"}
// https://data.eastmoney.com/kzz/detail/123121.html
func BondZhCovInfo(symbol, indicator string) ([]map[string]interface{}, error) {
	indicatorMap := map[string]string{
		"基本信息": "RPT_BOND_CB_LIST",
		"中签号":  "RPT_CB_BALLOTNUM",
		"筹资用途": "RPT_BOND_BS_OPRFINVESTITEM",
		"重要日期": "RPT_CB_IMPORTANTDATE",
	}

	url := "https://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"reportName":   indicatorMap[indicator],
		"columns":      "ALL",
		"quoteColumns": "",
		"quoteType":    "0",
		"source":       "WEB",
		"client":       "WEB",
		"filter":       fmt.Sprintf(`(SECURITY_CODE="%s")`, symbol),
	}

	if indicator == "基本信息" {
		params["quoteColumns"] = "f2~01~CONVERT_STOCK_CODE~CONVERT_STOCK_PRICE,f235~10~SECURITY_CODE~TRANSFER_PRICE,f236~10~SECURITY_CODE~TRANSFER_VALUE,f2~10~SECURITY_CODE~CURRENT_BOND_PRICE,f237~10~SECURITY_CODE~TRANSFER_PREMIUM_RATIO,f239~10~SECURITY_CODE~RESALE_TRIG_PRICE,f240~10~SECURITY_CODE~REDEEM_TRIG_PRICE,f23~01~CONVERT_STOCK_CODE~PBV_RATIO"
	} else if indicator == "筹资用途" {
		params["sortColumns"] = "SORT"
		params["sortTypes"] = "1"
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("获取可转债详情失败: %w", err)
	}

	dataArr := gjson.Get(string(resp.Body()), "result.data").Array()
	var items []map[string]interface{}

	for _, item := range dataArr {
		record := make(map[string]interface{})
		item.ForEach(func(key, value gjson.Result) bool {
			record[key.String()] = value.Value()
			return true
		})
		items = append(items, record)
	}

	return items, nil
}

// BondZhCovValueAnalysis 东方财富网-数据中心-新股数据-可转债数据-价值分析-溢价率分析
// symbol: 可转债代码，如 "113527"
// https://data.eastmoney.com/kzz/detail/113527.html
func BondZhCovValueAnalysis(symbol string) ([]BondZhCovValueAnalysisItem, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/get"
	params := map[string]string{
		"sty":    "ALL",
		"token":  "894050c76af8597a853f5b408b759f5d",
		"st":     "date",
		"sr":     "1",
		"source": "WEB",
		"type":   "RPTA_WEB_KZZ_LS",
		"filter": fmt.Sprintf(`(zcode="%s")`, symbol),
		"p":      "1",
		"ps":     "8000",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("获取可转债价值分析失败: %w", err)
	}

	dataArr := gjson.Get(string(resp.Body()), "result.data").Array()
	var items []BondZhCovValueAnalysisItem

	for _, item := range dataArr {
		items = append(items, BondZhCovValueAnalysisItem{
			Date:            item.Get("date").String()[:10],
			ClosePrice:      item.Get("spj").Float(),
			PureBondValue:   item.Get("czjz").Float(),
			ConvertValue:    item.Get("zgjz").Float(),
			PureBondPremium: item.Get("czyl").Float(),
			ConvertPremium:  item.Get("zgyl").Float(),
		})
	}

	return items, nil
}
