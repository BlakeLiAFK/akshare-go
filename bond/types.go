package bond

// ConvertibleBondJSL 集思录可转债数据结构
type ConvertibleBondJSL struct {
	BondID       string  `json:"bond_id"`       // 代码
	BondNM       string  `json:"bond_nm"`       // 转债名称
	Price        float64 `json:"price"`         // 现价
	IncreaseRT   string  `json:"increase_rt"`   // 涨跌幅
	StockID      string  `json:"stock_id"`      // 正股代码
	StockNM      string  `json:"stock_nm"`      // 正股名称
	SPrice       float64 `json:"sprice"`        // 正股价
	SIncreaseRT  string  `json:"sincrease_rt"`  // 正股涨跌
	PB           float64 `json:"pb"`            // 正股PB
	ConvertPrice float64 `json:"convert_price"` // 转股价
}

// BondIndexJSL 集思录可转债指数数据结构
type BondIndexJSL struct {
	Date  string  `json:"date"`  // 日期
	Value float64 `json:"value"` // 指数值
}

// BondSpotInfo 债券现货行情数据结构
type BondSpotInfo struct {
	Symbol        string  `json:"symbol"`        // 代码
	Name          string  `json:"name"`          // 名称
	Trade         float64 `json:"trade"`         // 现价
	PriceChange   float64 `json:"pricechange"`   // 涨跌额
	ChangePercent float64 `json:"changepercent"` // 涨跌幅
	Buy           float64 `json:"buy"`           // 买入
	Sell          float64 `json:"sell"`          // 卖出
	Settlement    float64 `json:"settlement"`    // 昨收
	Open          float64 `json:"open"`          // 今开
	High          float64 `json:"high"`          // 最高
	Low           float64 `json:"low"`           // 最低
	Volume        int64   `json:"volume"`        // 成交量
	Amount        float64 `json:"amount"`        // 成交额
	Time          string  `json:"time"`          // 时间
}

// BondDailyInfo 债券日线数据结构
type BondDailyInfo struct {
	Date   string  `json:"date"`   // 日期
	Open   float64 `json:"open"`   // 开盘
	High   float64 `json:"high"`   // 最高
	Low    float64 `json:"low"`    // 最低
	Close  float64 `json:"close"`  // 收盘
	Volume int64   `json:"volume"` // 成交量
}

// BondCovInfo 可转债信息数据结构
type BondCovInfo struct {
	Code         string  `json:"code"`          // 代码
	Name         string  `json:"name"`          // 名称
	StockCode    string  `json:"stock_code"`    // 正股代码
	StockName    string  `json:"stock_name"`    // 正股名称
	ConvertPrice float64 `json:"convert_price"` // 转股价
	ConvertValue float64 `json:"convert_value"` // 转股价值
	Premium      float64 `json:"premium"`       // 溢价率
	IssueDate    string  `json:"issue_date"`    // 发行日期
	MaturityDate string  `json:"maturity_date"` // 到期日期
}

// BondIssueInfo 债券发行信息数据结构
type BondIssueInfo struct {
	Code         string  `json:"code"`          // 债券代码
	Name         string  `json:"name"`          // 债券名称
	IssueDate    string  `json:"issue_date"`    // 发行日期
	IssuePrice   float64 `json:"issue_price"`   // 发行价格
	IssueAmount  float64 `json:"issue_amount"`  // 发行总额
	MaturityDate string  `json:"maturity_date"` // 到期日期
	CouponRate   float64 `json:"coupon_rate"`   // 票面利率
	Rating       string  `json:"rating"`        // 债券评级
}

// BondYield 债券收益率数据结构
type BondYield struct {
	Date        string  `json:"date"`        // 日期
	Period      string  `json:"period"`      // 期限
	Yield       float64 `json:"yield"`       // 收益率
	BondType    string  `json:"bond_type"`   // 债券类型
	Description string  `json:"description"` // 说明
}

// BondBuyBackInfo 债券回购信息数据结构
type BondBuyBackInfo struct {
	Code   string  `json:"code"`   // 代码
	Name   string  `json:"name"`   // 名称
	Date   string  `json:"date"`   // 日期
	Amount float64 `json:"amount"` // 成交额
	Volume int64   `json:"volume"` // 成交量
	Price  float64 `json:"price"`  // 价格
	Change float64 `json:"change"` // 涨跌幅
}

// BondCompositeIndex 债券综合指数数据结构
type BondCompositeIndex struct {
	Date       string  `json:"date"`        // 日期
	IndexValue float64 `json:"index_value"` // 指数值
	Change     float64 `json:"change"`      // 涨跌幅
	IndexName  string  `json:"index_name"`  // 指数名称
}
