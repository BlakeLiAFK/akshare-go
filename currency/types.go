package currency

import "time"

// CurrencyBocSinaItem 中国银行外汇牌价数据项
type CurrencyBocSinaItem struct {
	Date         time.Time // 日期
	BuyingRate   float64   // 中行汇买价
	CashBuying   float64   // 中行钞买价
	SellingRate  float64   // 中行钞卖价/汇卖价
	MiddleRate   float64   // 央行中间价
	ConvertPrice float64   // 中行折算价
}

// CurrencyBocSafeItem 人民币汇率中间价数据项
type CurrencyBocSafeItem struct {
	Date  time.Time          // 日期
	Rates map[string]float64 // 各货币汇率，key为货币代码
}

// CurrencyLatestItem 最新汇率数据项
type CurrencyLatestItem struct {
	Currency string    // 货币代码
	Date     time.Time // 日期
	Value    float64   // 汇率值
}

// CurrencyHistoryItem 历史汇率数据项
type CurrencyHistoryItem struct {
	Currency string    // 货币代码
	Date     time.Time // 日期
	Value    float64   // 汇率值
}

// CurrencyCurrenciesItem 货币列表数据项
type CurrencyCurrenciesItem struct {
	CurrencyCode string // 货币代码
	CurrencyName string // 货币名称
	Countries    string // 国家/地区
	Symbol       string // 货币符号
}

// CurrencyConvertResult 货币转换结果
type CurrencyConvertResult struct {
	From      string    // 源货币
	To        string    // 目标货币
	Amount    float64   // 源货币金额
	Value     float64   // 转换后金额
	Timestamp time.Time // 时间戳
}
