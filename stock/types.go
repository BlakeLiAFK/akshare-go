package stock

import "time"

// StockQuote A股实时行情
type StockQuote struct {
	Code         string    `json:"code"`          // 股票代码
	Name         string    `json:"name"`          // 股票名称
	Open         float64   `json:"open"`          // 开盘价
	High         float64   `json:"high"`          // 最高价
	Low          float64   `json:"low"`           // 最低价
	Price        float64   `json:"price"`         // 最新价
	PreClose     float64   `json:"pre_close"`     // 昨收价
	Change       float64   `json:"change"`        // 涨跌额
	ChangePct    float64   `json:"change_pct"`    // 涨跌幅(%)
	Volume       int64     `json:"volume"`        // 成交量(手)
	Amount       float64   `json:"amount"`        // 成交额(元)
	TurnoverRate float64   `json:"turnover_rate"` // 换手率(%)
	PE           float64   `json:"pe"`            // 市盈率
	PB           float64   `json:"pb"`            // 市净率
	MarketCap    float64   `json:"market_cap"`    // 总市值
	CirculateCap float64   `json:"circulate_cap"` // 流通市值
	Time         time.Time `json:"time"`          // 时间
}

// StockKLine K线数据
type StockKLine struct {
	Date      time.Time `json:"date"`       // 日期
	Open      float64   `json:"open"`       // 开盘价
	High      float64   `json:"high"`       // 最高价
	Low       float64   `json:"low"`        // 最低价
	Close     float64   `json:"close"`      // 收盘价
	Volume    int64     `json:"volume"`     // 成交量
	Amount    float64   `json:"amount"`     // 成交额
	Adjust    string    `json:"adjust"`     // 复权类型: qfq/hfq/none
	Change    float64   `json:"change"`     // 涨跌额
	ChangePct float64   `json:"change_pct"` // 涨跌幅
	Turnover  float64   `json:"turnover"`   // 换手率
}

// BoardInfo 板块信息
type BoardInfo struct {
	Code         string  `json:"code"`           // 板块代码
	Name         string  `json:"name"`           // 板块名称
	Change       float64 `json:"change"`         // 涨跌额
	ChangePct    float64 `json:"change_pct"`     // 涨跌幅(%)
	Volume       int64   `json:"volume"`         // 成交量
	Amount       float64 `json:"amount"`         // 成交额
	LeadStock    string  `json:"lead_stock"`     // 领涨股
	LeadStockPct float64 `json:"lead_stock_pct"` // 领涨股涨幅
	StockCount   int     `json:"stock_count"`    // 成分股数量
	UpCount      int     `json:"up_count"`       // 上涨家数
	DownCount    int     `json:"down_count"`     // 下跌家数
}

// BoardStock 板块成分股
type BoardStock struct {
	Code      string  `json:"code"`       // 股票代码
	Name      string  `json:"name"`       // 股票名称
	Price     float64 `json:"price"`      // 最新价
	ChangePct float64 `json:"change_pct"` // 涨跌幅(%)
	Volume    int64   `json:"volume"`     // 成交量
	Amount    float64 `json:"amount"`     // 成交额
}

// FundFlow 资金流向
type FundFlow struct {
	Date       time.Time `json:"date"`         // 日期
	Code       string    `json:"code"`         // 股票代码
	Name       string    `json:"name"`         // 股票名称
	Price      float64   `json:"price"`        // 最新价
	ChangePct  float64   `json:"change_pct"`   // 涨跌幅
	MainIn     float64   `json:"main_in"`      // 主力流入
	MainOut    float64   `json:"main_out"`     // 主力流出
	MainNet    float64   `json:"main_net"`     // 主力净流入
	MainNetPct float64   `json:"main_net_pct"` // 主力净占比
	SuperIn    float64   `json:"super_in"`     // 超大单流入
	SuperOut   float64   `json:"super_out"`    // 超大单流出
	SuperNet   float64   `json:"super_net"`    // 超大单净流入
	BigIn      float64   `json:"big_in"`       // 大单流入
	BigOut     float64   `json:"big_out"`      // 大单流出
	BigNet     float64   `json:"big_net"`      // 大单净流入
	MidIn      float64   `json:"mid_in"`       // 中单流入
	MidOut     float64   `json:"mid_out"`      // 中单流出
	MidNet     float64   `json:"mid_net"`      // 中单净流入
	SmallIn    float64   `json:"small_in"`     // 小单流入
	SmallOut   float64   `json:"small_out"`    // 小单流出
	SmallNet   float64   `json:"small_net"`    // 小单净流入
}

// MarketFundFlow 大盘资金流向
type MarketFundFlow struct {
	Date       time.Time `json:"date"`        // 日期
	Index      string    `json:"index"`       // 指数名称
	ClosePrice float64   `json:"close_price"` // 收盘价
	ChangePct  float64   `json:"change_pct"`  // 涨跌幅
	MainIn     float64   `json:"main_in"`     // 主力流入
	MainOut    float64   `json:"main_out"`    // 主力流出
	MainNet    float64   `json:"main_net"`    // 主力净流入
	RetailIn   float64   `json:"retail_in"`   // 散户流入
	RetailOut  float64   `json:"retail_out"`  // 散户流出
	RetailNet  float64   `json:"retail_net"`  // 散户净流入
}

// StockInfo 股票基本信息
type StockInfo struct {
	Code         string  `json:"code"`          // 股票代码
	Name         string  `json:"name"`          // 股票名称
	Industry     string  `json:"industry"`      // 行业
	ListDate     string  `json:"list_date"`     // 上市日期
	TotalShares  float64 `json:"total_shares"`  // 总股本(股)
	FloatShares  float64 `json:"float_shares"`  // 流通股本(股)
	MarketCap    float64 `json:"market_cap"`    // 总市值
	CirculateCap float64 `json:"circulate_cap"` // 流通市值
	PE           float64 `json:"pe"`            // 市盈率(动态)
	PEStatic     float64 `json:"pe_static"`     // 市盈率(静态)
	PB           float64 `json:"pb"`            // 市净率
	ROE          float64 `json:"roe"`           // ROE
	EPS          float64 `json:"eps"`           // 每股收益
	BPS          float64 `json:"bps"`           // 每股净资产
}

// HKStockQuote 港股实时行情
type HKStockQuote struct {
	Code      string    `json:"code"`       // 股票代码
	Name      string    `json:"name"`       // 股票名称
	NameEn    string    `json:"name_en"`    // 英文名称
	Open      float64   `json:"open"`       // 开盘价
	High      float64   `json:"high"`       // 最高价
	Low       float64   `json:"low"`        // 最低价
	Price     float64   `json:"price"`      // 最新价
	PreClose  float64   `json:"pre_close"`  // 昨收价
	Change    float64   `json:"change"`     // 涨跌额
	ChangePct float64   `json:"change_pct"` // 涨跌幅(%)
	Volume    int64     `json:"volume"`     // 成交量
	Amount    float64   `json:"amount"`     // 成交额
	Time      time.Time `json:"time"`       // 时间
}

// USStockQuote 美股实时行情
type USStockQuote struct {
	Code      string    `json:"code"`       // 股票代码
	Name      string    `json:"name"`       // 股票名称
	NameCN    string    `json:"name_cn"`    // 中文名称
	Open      float64   `json:"open"`       // 开盘价
	High      float64   `json:"high"`       // 最高价
	Low       float64   `json:"low"`        // 最低价
	Price     float64   `json:"price"`      // 最新价
	PreClose  float64   `json:"pre_close"`  // 昨收价
	Change    float64   `json:"change"`     // 涨跌额
	ChangePct float64   `json:"change_pct"` // 涨跌幅(%)
	Volume    int64     `json:"volume"`     // 成交量
	Amount    float64   `json:"amount"`     // 成交额
	MarketCap float64   `json:"market_cap"` // 市值
	PE        float64   `json:"pe"`         // 市盈率
	Time      time.Time `json:"time"`       // 时间
}

// StockCodeName 股票代码名称
type StockCodeName struct {
	Code string `json:"code"` // 股票代码
	Name string `json:"name"` // 股票名称
}

// IntradayQuote 日内分时数据
type IntradayQuote struct {
	Time     string  `json:"time"`      // 时间 HH:MM
	Price    float64 `json:"price"`     // 价格
	Volume   int64   `json:"volume"`    // 成交量
	Amount   float64 `json:"amount"`    // 成交额
	AvgPrice float64 `json:"avg_price"` // 均价
}
