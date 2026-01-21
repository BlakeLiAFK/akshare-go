package stock_feature

import "time"

// StockHighLowStatisticsItem 创新高新低统计项
type StockHighLowStatisticsItem struct {
	Date    time.Time `json:"date"`    // 日期
	Close   float64   `json:"close"`   // 收盘价
	High20  int       `json:"high20"`  // 20日新高数量
	Low20   int       `json:"low20"`   // 20日新低数量
	High60  int       `json:"high60"`  // 60日新高数量
	Low60   int       `json:"low60"`   // 60日新低数量
	High120 int       `json:"high120"` // 120日新高数量
	Low120  int       `json:"low120"`  // 120日新低数量
}

// StockAccountStatisticsItem 股票账户统计项
type StockAccountStatisticsItem struct {
	DataDate           string  `json:"data_date"`            // 数据日期
	NewInvestorCount   float64 `json:"new_investor_count"`   // 新增投资者数量
	NewInvestorMoM     float64 `json:"new_investor_mom"`     // 新增投资者环比
	NewInvestorYoY     float64 `json:"new_investor_yoy"`     // 新增投资者同比
	TotalInvestor      float64 `json:"total_investor"`       // 期末投资者总量
	AShareAccount      float64 `json:"a_share_account"`      // A股账户
	BShareAccount      float64 `json:"b_share_account"`      // B股账户
	TotalMarketCap     float64 `json:"total_market_cap"`     // 沪深总市值
	AvgMarketCap       float64 `json:"avg_market_cap"`       // 沪深户均市值
	ShangHaiIndexClose float64 `json:"shanghai_index_close"` // 上证指数收盘
	ShangHaiIndexPct   float64 `json:"shanghai_index_pct"`   // 上证指数涨跌幅
}

// StockAllPbItem 全部A股市净率项
type StockAllPbItem struct {
	Date          time.Time `json:"date"`            // 日期
	EqualWeightPB float64   `json:"equal_weight_pb"` // 等权重市净率
	MedianPB      float64   `json:"median_pb"`       // 中位数市净率
	Close         float64   `json:"close"`           // 收盘价
}

// StockBuffettIndexItem 巴菲特指标项
type StockBuffettIndexItem struct {
	Date               time.Time `json:"date"`                 // 日期
	Close              float64   `json:"close"`                // 收盘价
	MarketCap          float64   `json:"market_cap"`           // 总市值
	GDP                float64   `json:"gdp"`                  // GDP
	Quantile10Years    float64   `json:"quantile_10_years"`    // 近十年分位数
	QuantileAllHistory float64   `json:"quantile_all_history"` // 总历史分位数
}

// StockCongestionItem 大盘拥挤度项
type StockCongestionItem struct {
	Date       time.Time `json:"date"`       // 日期
	Close      float64   `json:"close"`      // 收盘价
	Congestion float64   `json:"congestion"` // 拥挤度
}

// StockEbsItem 股债利差项
type StockEbsItem struct {
	Date            time.Time `json:"date"`              // 日期
	HS300Index      float64   `json:"hs300_index"`       // 沪深300指数
	PeSpread        float64   `json:"pe_spread"`         // 股债利差
	PeSpreadAverage float64   `json:"pe_spread_average"` // 股债利差均线
}

// StockGxlItem 股息率项
type StockGxlItem struct {
	Date         time.Time `json:"date"`          // 日期
	DividendRate float64   `json:"dividend_rate"` // 股息率
}

// StockTtmLyrItem 市盈率项
type StockTtmLyrItem struct {
	Date          time.Time `json:"date"`            // 日期
	EqualWeightPE float64   `json:"equal_weight_pe"` // 等权重市盈率
	MedianPE      float64   `json:"median_pe"`       // 中位数市盈率
	Close         float64   `json:"close"`           // 收盘价
}

// StockHkIndicatorItem 港股指标项
type StockHkIndicatorItem struct {
	StockID   string  `json:"stock_id"`   // 股票代码
	StockName string  `json:"stock_name"` // 股票名称
	Date      string  `json:"date"`       // 日期
	Value     float64 `json:"value"`      // 指标值
}

// StockMarketActivityItem 赚钱效应分析项
type StockMarketActivityItem struct {
	Item  string `json:"item"`  // 项目
	Value string `json:"value"` // 值
}

// StockBelowNetAssetItem 破净股统计项
type StockBelowNetAssetItem struct {
	Date          time.Time `json:"date"`            // 日期
	Close         float64   `json:"close"`           // 收盘价
	BelowNetCount int       `json:"below_net_count"` // 破净股数量
	TotalCount    int       `json:"total_count"`     // 总股票数量
	BelowNetRatio float64   `json:"below_net_ratio"` // 破净股比例
}

// StockPeAndPbItem 市盈率市净率项
type StockPeAndPbItem struct {
	Date  time.Time `json:"date"`  // 日期
	PE    float64   `json:"pe"`    // 市盈率
	PB    float64   `json:"pb"`    // 市净率
	Close float64   `json:"close"` // 收盘价
}
