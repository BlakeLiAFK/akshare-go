package index

import "time"

// IndexQuote 指数实时行情
type IndexQuote struct {
	Seq         int     `json:"seq"`          // 序号
	Code        string  `json:"code"`         // 代码
	Name        string  `json:"name"`         // 名称
	Price       float64 `json:"price"`        // 最新价
	ChangePct   float64 `json:"change_pct"`   // 涨跌幅(%)
	Change      float64 `json:"change"`       // 涨跌额
	Volume      int64   `json:"volume"`       // 成交量
	Amount      float64 `json:"amount"`       // 成交额
	Amplitude   float64 `json:"amplitude"`    // 振幅
	High        float64 `json:"high"`         // 最高
	Low         float64 `json:"low"`          // 最低
	Open        float64 `json:"open"`         // 今开
	PreClose    float64 `json:"pre_close"`    // 昨收
	VolumeRatio float64 `json:"volume_ratio"` // 量比
	UpdateTime  string  `json:"update_time"`  // 更新时间
}

// IndexKLine 指数K线数据
type IndexKLine struct {
	Date      time.Time `json:"date"`       // 日期
	Open      float64   `json:"open"`       // 开盘
	Close     float64   `json:"close"`      // 收盘
	High      float64   `json:"high"`       // 最高
	Low       float64   `json:"low"`        // 最低
	Volume    int64     `json:"volume"`     // 成交量
	Amount    float64   `json:"amount"`     // 成交额
	Amplitude float64   `json:"amplitude"`  // 振幅
	ChangePct float64   `json:"change_pct"` // 涨跌幅
	Change    float64   `json:"change"`     // 涨跌额
	Turnover  float64   `json:"turnover"`   // 换手率
}

// IndexIntradayQuote 指数分时数据
type IndexIntradayQuote struct {
	Time   string  `json:"time"`   // 时间
	Open   float64 `json:"open"`   // 开盘
	Close  float64 `json:"close"`  // 收盘
	High   float64 `json:"high"`   // 最高
	Low    float64 `json:"low"`    // 最低
	Volume int64   `json:"volume"` // 成交量
	Amount float64 `json:"amount"` // 成交额
	Avg    float64 `json:"avg"`    // 均价
}

// IndexCodeName 指数代码名称
type IndexCodeName struct {
	Code string `json:"code"` // 代码
	Name string `json:"name"` // 名称
}

// IndexSymbolCode 指数代码映射
type IndexSymbolCode struct {
	Code     string `json:"code"`      // 指数代码
	MarketID string `json:"market_id"` // 市场ID
}

// SpotGoodsQuote 商品现货价格指数
type SpotGoodsQuote struct {
	Date      time.Time `json:"date"`       // 日期
	Price     float64   `json:"price"`      // 指数
	Change    float64   `json:"change"`     // 涨跌额
	ChangePct float64   `json:"change_pct"` // 涨跌幅
}

// GlobalIndexKLine 全球指数K线数据
type GlobalIndexKLine struct {
	Date   time.Time `json:"date"`   // 日期
	Open   float64   `json:"open"`   // 开盘
	High   float64   `json:"high"`   // 最高
	Low    float64   `json:"low"`    // 最低
	Close  float64   `json:"close"`  // 收盘
	Volume int64     `json:"volume"` // 成交量
}

// IndexStockCons 指数成份股
type IndexStockCons struct {
	Date   string `json:"date"`   // 日期
	Code   string `json:"code"`   // 股票代码
	Name   string `json:"name"`   // 股票名称
	Sort   int    `json:"sort"`   // 排序
	Weight string `json:"weight"` // 权重
}

// IndexStockConsWeightCSIndex 中证指数成份股权重
type IndexStockConsWeightCSIndex struct {
	Date      string  `json:"date"`       // 日期
	IndexCode string  `json:"index_code"` // 指数代码
	IndexName string  `json:"index_name"` // 指数名称
	StockCode string  `json:"stock_code"` // 成分券代码
	StockName string  `json:"stock_name"` // 成分券名称
	Exchange  string  `json:"exchange"`   // 交易所
	Weight    float64 `json:"weight"`     // 权重
}

// CSIndexInfo 中证指数信息
type CSIndexInfo struct {
	IndexCode   string  `json:"index_code"`    // 指数代码
	IndexName   string  `json:"index_name"`    // 指数名称
	IndexNameEN string  `json:"index_name_en"` // 英文名称
	BaseDate    string  `json:"base_date"`     // 基日
	BasePoint   float64 `json:"base_point"`    // 基点
	PublishDate string  `json:"publish_date"`  // 发布时间
	IndexSeries string  `json:"index_series"`  // 指数系列
	Currency    string  `json:"currency"`      // 币种
	Region      string  `json:"region"`        // 区域
}

// CNIIndexInfo 国证指数信息
type CNIIndexInfo struct {
	IndexCode  string  `json:"index_code"`  // 指数代码
	IndexName  string  `json:"index_name"`  // 指数简称
	SampleNum  int     `json:"sample_num"`  // 样本数
	ClosePoint float64 `json:"close_point"` // 收盘点位
	ChangePct  float64 `json:"change_pct"`  // 涨跌幅
	PE         float64 `json:"pe"`          // PE滚动
	Volume     int64   `json:"volume"`      // 成交量
	Amount     float64 `json:"amount"`      // 成交额
	MarketCap  float64 `json:"market_cap"`  // 总市值
	FloatCap   float64 `json:"float_cap"`   // 自由流通市值
}

// CNIIndexDetail 国证指数成份股详情
type CNIIndexDetail struct {
	Date      string  `json:"date"`       // 日期
	StockCode string  `json:"stock_code"` // 样本代码
	StockName string  `json:"stock_name"` // 样本简称
	Industry  string  `json:"industry"`   // 所属行业
	MarketCap float64 `json:"market_cap"` // 总市值
	Weight    float64 `json:"weight"`     // 权重
}

// SWIndexInfo 申万行业指数信息
type SWIndexInfo struct {
	IndexCode     string  `json:"index_code"`     // 行业代码
	IndexName     string  `json:"index_name"`     // 行业名称
	StockCount    int     `json:"stock_count"`    // 成份个数
	PEStatic      float64 `json:"pe_static"`      // 静态市盈率
	PE            float64 `json:"pe"`             // TTM(滚动)市盈率
	PB            float64 `json:"pb"`             // 市净率
	DividendYield float64 `json:"dividend_yield"` // 静态股息率
}

// SWIndexCons 申万行业指数成份股
type SWIndexCons struct {
	Seq               int     `json:"seq"`                  // 序号
	StockCode         string  `json:"stock_code"`           // 股票代码
	StockName         string  `json:"stock_name"`           // 股票简称
	IncludedDate      string  `json:"included_date"`        // 纳入时间
	SWLevel1          string  `json:"sw_level1"`            // 申万1级
	SWLevel2          string  `json:"sw_level2"`            // 申万2级
	SWLevel3          string  `json:"sw_level3"`            // 申万3级
	Price             float64 `json:"price"`                // 价格
	PE                float64 `json:"pe"`                   // 市盈率
	PETTM             float64 `json:"pe_ttm"`               // 市盈率ttm
	PB                float64 `json:"pb"`                   // 市净率
	DividendYield     float64 `json:"dividend_yield"`       // 股息率
	MarketCap         float64 `json:"market_cap"`           // 市值
	NetProfitGrowth09 float64 `json:"net_profit_growth_09"` // 归母净利润同比增长(09-30)
	NetProfitGrowth06 float64 `json:"net_profit_growth_06"` // 归母净利润同比增长(06-30)
	RevenueGrowth09   float64 `json:"revenue_growth_09"`    // 营业收入同比增长(09-30)
	RevenueGrowth06   float64 `json:"revenue_growth_06"`    // 营业收入同比增长(06-30)
}

// HKIndexQuote 港股指数行情
type HKIndexQuote struct {
	Seq       int     `json:"seq"`        // 序号
	MarketID  string  `json:"market_id"`  // 内部编号
	Code      string  `json:"code"`       // 代码
	Name      string  `json:"name"`       // 名称
	Price     float64 `json:"price"`      // 最新价
	Change    float64 `json:"change"`     // 涨跌额
	ChangePct float64 `json:"change_pct"` // 涨跌幅
	Open      float64 `json:"open"`       // 今开
	High      float64 `json:"high"`       // 最高
	Low       float64 `json:"low"`        // 最低
	PreClose  float64 `json:"pre_close"`  // 昨收
	Volume    int64   `json:"volume"`     // 成交量
	Amount    float64 `json:"amount"`     // 成交额
}

// YWIndexQuote 义乌小商品指数
type YWIndexQuote struct {
	PeriodNo       string  `json:"period_no"`       // 期数
	PriceIndex     float64 `json:"price_index"`     // 价格指数
	FieldPrice     float64 `json:"field_price"`     // 场内价格指数
	NetPrice       float64 `json:"net_price"`       // 网上价格指数
	OrderPrice     float64 `json:"order_price"`     // 订单价格指数
	ExportPrice    float64 `json:"export_price"`    // 出口价格指数
	BoomIndex      float64 `json:"boom_index"`      // 景气指数
	ScopeIndex     float64 `json:"scope_index"`     // 规模指数
	ProfitIndex    float64 `json:"profit_index"`    // 效益指数
	ConfidentIndex float64 `json:"confident_index"` // 市场信心指数
}

// SugarIndex 中国食糖指数
type SugarIndex struct {
	Date           string  `json:"date"`            // 日期
	CompositePrice float64 `json:"composite_price"` // 综合价格
	RawSugarPrice  float64 `json:"raw_sugar_price"` // 原糖价格
	SpotPrice      float64 `json:"spot_price"`      // 现货价格
}

// DrewryContainerIndex Drewry集装箱指数
type DrewryContainerIndex struct {
	Date string  `json:"date"` // 日期
	WCI  float64 `json:"wci"`  // 集装箱指数
}

// KQIndexQuote 柯桥纺织指数
type KQIndexQuote struct {
	Period      string  `json:"period"`      // 期次
	Index       float64 `json:"index"`       // 指数
	ChangePct   float64 `json:"change_pct"`  // 涨跌幅
	TotalIndex  float64 `json:"total_index"` // 总景气指数
	Circulation float64 `json:"circulation"` // 流通景气指数
	Production  float64 `json:"production"`  // 生产景气指数
	PriceIndex  float64 `json:"price_index"` // 价格指数
	BoomIndex   float64 `json:"boom_index"`  // 景气指数
}
