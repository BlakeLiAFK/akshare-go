package option

// OptionCommSymbolItem 商品期权品种
// 对应 Python: option_comm_symbol
type OptionCommSymbolItem struct {
	Name string `json:"name"` // 品种名称
	Code string `json:"code"` // 品种代码
}

// OptionCommInfoItem 商品期权手续费信息
// 对应 Python: option_comm_info
type OptionCommInfoItem struct {
	Contract        string  `json:"contract"`          // 合约
	CurrentPrice    float64 `json:"current_price"`     // 现价
	Volume          int64   `json:"volume"`            // 成交量
	TickProfit      float64 `json:"tick_profit"`       // 每跳毛利/元
	TickNetProfit   float64 `json:"tick_net_profit"`   // 每跳净利/元
	Exchange        string  `json:"exchange"`          // 交易所
	CommUpdateTime  string  `json:"comm_update_time"`  // 手续费更新时间
	PriceUpdateTime string  `json:"price_update_time"` // 价格更新时间
}

// OptionCurrentEmItem 东方财富期权行情
// 对应 Python: option_current_em
type OptionCurrentEmItem struct {
	Index        int     `json:"index"`         // 序号
	Code         string  `json:"code"`          // 代码
	Name         string  `json:"name"`          // 名称
	LatestPrice  float64 `json:"latest_price"`  // 最新价
	Change       float64 `json:"change"`        // 涨跌额
	ChangeRate   float64 `json:"change_rate"`   // 涨跌幅
	Volume       int64   `json:"volume"`        // 成交量
	Amount       float64 `json:"amount"`        // 成交额
	Position     int64   `json:"position"`      // 持仓量
	StrikePrice  float64 `json:"strike_price"`  // 行权价
	RemainingDay int     `json:"remaining_day"` // 剩余日
	DayIncrease  int64   `json:"day_increase"`  // 日增
	PreSettle    float64 `json:"pre_settle"`    // 昨结
	Open         float64 `json:"open"`          // 今开
	MarketId     string  `json:"market_id"`     // 市场标识
}

// OptionCurrentDaySseItem 上交所期权当日合约
// 对应 Python: option_current_day_sse
type OptionCurrentDaySseItem struct {
	SecurityId     string `json:"security_id"`     // 合约编码
	ContractId     string `json:"contract_id"`     // 合约交易代码
	ContractSymbol string `json:"contract_symbol"` // 合约简称
	SecurityName   string `json:"security_name"`   // 标的券名称及代码
	CallOrPut      string `json:"call_or_put"`     // 类型
	ExercisePrice  string `json:"exercise_price"`  // 行权价
	ContractUnit   string `json:"contract_unit"`   // 合约单位
	EndDate        string `json:"end_date"`        // 期权行权日
	DeliveryDate   string `json:"delivery_date"`   // 行权交收日
	ExpireDate     string `json:"expire_date"`     // 到期日
	StartDate      string `json:"start_date"`      // 开始日期
}

// OptionCurrentDaySzseItem 深交所期权当日合约
// 对应 Python: option_current_day_szse
type OptionCurrentDaySzseItem struct {
	Index                 int     `json:"index"`                    // 序号
	ContractCode          string  `json:"contract_code"`            // 合约编码
	ContractId            string  `json:"contract_id"`              // 合约代码
	ContractName          string  `json:"contract_name"`            // 合约简称
	UnderlyingName        string  `json:"underlying_name"`          // 标的证券简称(代码)
	ContractType          string  `json:"contract_type"`            // 合约类型
	StrikePrice           float64 `json:"strike_price"`             // 行权价
	ContractUnit          int     `json:"contract_unit"`            // 合约单位
	LastTradingDay        string  `json:"last_trading_day"`         // 最后交易日
	ExerciseDay           string  `json:"exercise_day"`             // 行权日
	ExpireDay             string  `json:"expire_day"`               // 到期日
	SettlementDay         string  `json:"settlement_day"`           // 交收日
	IsNew                 string  `json:"is_new"`                   // 新挂
	UpperLimit            float64 `json:"upper_limit"`              // 涨停价格
	LowerLimit            float64 `json:"lower_limit"`              // 跌停价格
	PreSettlement         float64 `json:"pre_settlement"`           // 前结算价
	ContractAdjust        string  `json:"contract_adjust"`          // 合约调整
	Suspended             string  `json:"suspended"`                // 停牌
	TotalPosition         int64   `json:"total_position"`           // 合约总持仓
	ListReason            string  `json:"list_reason"`              // 挂牌原因
	OrigContractId        string  `json:"orig_contract_id"`         // 原合约代码
	OrigContractName      string  `json:"orig_contract_name"`       // 原合约简称
	OrigStrikePrice       float64 `json:"orig_strike_price"`        // 原行权价格
	OrigContractUnit      int     `json:"orig_contract_unit"`       // 原合约单位
	RemainingTradingDays  int     `json:"remaining_trading_days"`   // 合约到期剩余交易天数
	RemainingNaturalDays  int     `json:"remaining_natural_days"`   // 合约到期剩余自然天数
	NextAdjustTradingDays int     `json:"next_adjust_trading_days"` // 下次合约调整剩余交易天数
	NextAdjustNaturalDays int     `json:"next_adjust_natural_days"` // 下次合约调整剩余自然天数
	TradeDate             string  `json:"trade_date"`               // 交易日期
}

// OptionDailyStatsSseItem 上交所期权每日统计
// 对应 Python: option_daily_stats_sse
type OptionDailyStatsSseItem struct {
	SecurityCode     string  `json:"security_code"`      // 合约标的代码
	SecurityName     string  `json:"security_name"`      // 合约标的名称
	ContractCount    int     `json:"contract_count"`     // 合约数量
	TotalAmount      float64 `json:"total_amount"`       // 总成交额
	TotalVolume      int64   `json:"total_volume"`       // 总成交量
	CallVolume       int64   `json:"call_volume"`        // 认购成交量
	PutVolume        int64   `json:"put_volume"`         // 认沽成交量
	PutCallRatio     float64 `json:"put_call_ratio"`     // 认沽/认购
	OpenInterest     int64   `json:"open_interest"`      // 未平仓合约总数
	OpenCallInterest int64   `json:"open_call_interest"` // 未平仓认购合约数
	OpenPutInterest  int64   `json:"open_put_interest"`  // 未平仓认沽合约数
	TradeDate        string  `json:"trade_date"`         // 交易日
}

// OptionDailyStatsSzseItem 深交所期权每日统计
// 对应 Python: option_daily_stats_szse
type OptionDailyStatsSzseItem struct {
	SecurityCode     string  `json:"security_code"`      // 合约标的代码
	SecurityName     string  `json:"security_name"`      // 合约标的名称
	Volume           int64   `json:"volume"`             // 成交量
	CallVolume       int64   `json:"call_volume"`        // 认购成交量
	PutVolume        int64   `json:"put_volume"`         // 认沽成交量
	PutCallRatio     float64 `json:"put_call_ratio"`     // 认沽/认购持仓比
	OpenInterest     int64   `json:"open_interest"`      // 未平仓合约总数
	OpenCallInterest int64   `json:"open_call_interest"` // 未平仓认购合约数
	OpenPutInterest  int64   `json:"open_put_interest"`  // 未平仓认沽合约数
	TradeDate        string  `json:"trade_date"`         // 交易日
}

// OptionRiskIndicatorSseItem 上交所期权风险指标
// 对应 Python: option_risk_indicator_sse
type OptionRiskIndicatorSseItem struct {
	TradeDate      string  `json:"trade_date"`      // 交易日
	SecurityId     string  `json:"security_id"`     // 合约编码
	ContractId     string  `json:"contract_id"`     // 合约代码
	ContractSymbol string  `json:"contract_symbol"` // 合约简称
	Delta          float64 `json:"delta"`           // Delta
	Theta          float64 `json:"theta"`           // Theta
	Gamma          float64 `json:"gamma"`           // Gamma
	Vega           float64 `json:"vega"`            // Vega
	Rho            float64 `json:"rho"`             // Rho
	ImpliedVol     float64 `json:"implied_vol"`     // 隐含波动率
}

// OptionPremiumAnalysisEmItem 东方财富期权折溢价分析
// 对应 Python: option_premium_analysis_em
type OptionPremiumAnalysisEmItem struct {
	OptionCode       string  `json:"option_code"`       // 期权代码
	OptionName       string  `json:"option_name"`       // 期权名称
	LatestPrice      float64 `json:"latest_price"`      // 最新价
	ChangeRate       float64 `json:"change_rate"`       // 涨跌幅
	StrikePrice      float64 `json:"strike_price"`      // 行权价
	PremiumRate      float64 `json:"premium_rate"`      // 折溢价率
	UnderlyingName   string  `json:"underlying_name"`   // 标的名称
	UnderlyingPrice  float64 `json:"underlying_price"`  // 标的最新价
	UnderlyingChange float64 `json:"underlying_change"` // 标的涨跌幅
	BreakEvenPrice   float64 `json:"break_even_price"`  // 盈亏平衡价
	ExpireDate       string  `json:"expire_date"`       // 到期日
}

// OptionRiskAnalysisEmItem 东方财富期权风险分析
// 对应 Python: option_risk_analysis_em
type OptionRiskAnalysisEmItem struct {
	OptionCode   string  `json:"option_code"`   // 期权代码
	OptionName   string  `json:"option_name"`   // 期权名称
	LatestPrice  float64 `json:"latest_price"`  // 最新价
	ChangeRate   float64 `json:"change_rate"`   // 涨跌幅
	Leverage     float64 `json:"leverage"`      // 杠杆比率
	RealLeverage float64 `json:"real_leverage"` // 实际杠杆比率
	Delta        float64 `json:"delta"`         // Delta
	Gamma        float64 `json:"gamma"`         // Gamma
	Vega         float64 `json:"vega"`          // Vega
	Rho          float64 `json:"rho"`           // Rho
	Theta        float64 `json:"theta"`         // Theta
	ExpireDate   string  `json:"expire_date"`   // 到期日
}

// OptionValueAnalysisEmItem 东方财富期权价值分析
// 对应 Python: option_value_analysis_em
type OptionValueAnalysisEmItem struct {
	OptionCode        string  `json:"option_code"`         // 期权代码
	OptionName        string  `json:"option_name"`         // 期权名称
	LatestPrice       float64 `json:"latest_price"`        // 最新价
	TimeValue         float64 `json:"time_value"`          // 时间价值
	IntrinsicValue    float64 `json:"intrinsic_value"`     // 内在价值
	ImpliedVol        float64 `json:"implied_vol"`         // 隐含波动率
	TheoreticalPrice  float64 `json:"theoretical_price"`   // 理论价格
	UnderlyingName    string  `json:"underlying_name"`     // 标的名称
	UnderlyingPrice   float64 `json:"underlying_price"`    // 标的最新价
	UnderlyingYearVol float64 `json:"underlying_year_vol"` // 标的近一年波动率
	ExpireDate        string  `json:"expire_date"`         // 到期日
}

// OptionLhbEmItem 东方财富期权龙虎榜
// 对应 Python: option_lhb_em
type OptionLhbEmItem struct {
	TradeType    string  `json:"trade_type"`    // 交易类型
	TradeDate    string  `json:"trade_date"`    // 交易日期
	SecurityCode string  `json:"security_code"` // 证券代码
	TargetName   string  `json:"target_name"`   // 标的名称
	Rank         int     `json:"rank"`          // 名次
	Member       string  `json:"member"`        // 机构
	Volume       int64   `json:"volume"`        // 交易量/持仓量
	Change       int64   `json:"change"`        // 增减
	NetVolume    int64   `json:"net_volume"`    // 净认沽量/净持仓量/净交易量
	Ratio        float64 `json:"ratio"`         // 占总交易量比例
}
