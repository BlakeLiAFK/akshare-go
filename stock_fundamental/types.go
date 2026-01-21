package stock_fundamental

import "time"

// StockProfitNoticeEm 东方财富业绩预告
type StockProfitNoticeEm struct {
	Code       string    `json:"code"`        // 股票代码
	Name       string    `json:"name"`        // 股票名称
	ReportDate time.Time `json:"report_date"` // 报告期
	ProfitType string    `json:"profit_type"` // 业绩变动类型
	ProfitMin  float64   `json:"profit_min"`  // 预告净利润下限
	ProfitMax  float64   `json:"profit_max"`  // 预告净利润上限
}

// StockHoldEm 东方财富股东持股
type StockHoldEm struct {
	Code         string    `json:"code"`          // 股票代码
	Name         string    `json:"name"`          // 股票名称
	HolderName   string    `json:"holder_name"`   // 股东名称
	Holdings     float64   `json:"holdings"`      // 持股数量
	HoldingRatio float64   `json:"holding_ratio"` // 持股比例
	Change       float64   `json:"change"`        // 变动
	EndDate      time.Time `json:"end_date"`      // 截止日期
}

// StockIpoDeclareEmItem 东方财富首发申报企业信息项
type StockIpoDeclareEmItem struct {
	Index         int       `json:"index"`          // 序号
	CompanyName   string    `json:"company_name"`   // 企业名称
	State         string    `json:"state"`          // 最新状态
	RegAddress    string    `json:"reg_address"`    // 注册地
	RecommendOrg  string    `json:"recommend_org"`  // 保荐机构
	LawFirm       string    `json:"law_firm"`       // 律师事务所
	AccountFirm   string    `json:"account_firm"`   // 会计师事务所
	PredictMarket string    `json:"predict_market"` // 拟上市地点
	EndDate       time.Time `json:"end_date"`       // 更新日期
	ProspectusURL string    `json:"prospectus_url"` // 招股说明书链接
}

// StockIpoReviewEmItem 东方财富过会企业信息项
type StockIpoReviewEmItem struct {
	Index           int       `json:"index"`            // 序号
	CompanyName     string    `json:"company_name"`     // 企业名称
	StockName       string    `json:"stock_name"`       // 股票简称
	StockCode       string    `json:"stock_code"`       // 股票代码
	TradeMarket     string    `json:"trade_market"`     // 上市板块
	ReviewDate      time.Time `json:"review_date"`      // 上会日期
	ReviewState     string    `json:"review_state"`     // 审核状态
	ReviewMember    string    `json:"review_member"`    // 发审委委员
	LeadUnderwriter string    `json:"lead_underwriter"` // 主承销商
	IssueNum        float64   `json:"issue_num"`        // 发行数量(股)
	FinanceAmt      float64   `json:"finance_amt"`      // 拟融资额(元)
	NoticeDate      time.Time `json:"notice_date"`      // 公告日期
	ListingDate     time.Time `json:"listing_date"`     // 上市日期
}

// StockIpoTutorEmItem 东方财富辅导备案信息项
type StockIpoTutorEmItem struct {
	Index             int       `json:"index"`               // 序号
	CompanyName       string    `json:"company_name"`        // 企业名称
	TutorOrg          string    `json:"tutor_org"`           // 辅导机构
	TutorProcessState string    `json:"tutor_process_state"` // 辅导状态
	ReportType        string    `json:"report_type"`         // 报告类型
	DispatchOrg       string    `json:"dispatch_org"`        // 派出机构
	ReportTitle       string    `json:"report_title"`        // 报告标题
	RecordDate        time.Time `json:"record_date"`         // 备案日期
}

// StockKcbSseItem 上交所科创板信息项
type StockKcbSseItem struct {
	Index          int       `json:"index"`            // 序号
	CompanyName    string    `json:"company_name"`     // 企业名称
	StockCode      string    `json:"stock_code"`       // 股票代码
	UpdateTime     time.Time `json:"update_time"`      // 更新日期
	AuditApplyDate time.Time `json:"audit_apply_date"` // 受理日期
	RegisterResult string    `json:"register_result"`  // 注册结果
	CommitResult   string    `json:"commit_result"`    // 审议结果
	CurrentStatus  string    `json:"current_status"`   // 当前状态
	Province       string    `json:"province"`         // 省份
	CsrcCode       string    `json:"csrc_code"`        // 证监会行业
	StockAuditNum  int       `json:"stock_audit_num"`  // 审核序号
}

// StockNoticeEm 东方财富公告
type StockNoticeEm struct {
	Code        string    `json:"code"`         // 股票代码
	Title       string    `json:"title"`        // 公告标题
	PublishDate time.Time `json:"publish_date"` // 发布日期
	Type        string    `json:"type"`         // 公告类型
}

// StockRecommendEm 东方财富机构推荐
type StockRecommendEm struct {
	Code          string    `json:"code"`           // 股票代码
	Name          string    `json:"name"`           // 股票名称
	InstituteName string    `json:"institute_name"` // 机构名称
	Rating        string    `json:"rating"`         // 评级
	TargetPrice   float64   `json:"target_price"`   // 目标价
	ReportDate    time.Time `json:"report_date"`    // 报告日期
}

// StockShareStructureEm 东方财富股本结构
type StockShareStructureEm struct {
	Code           string    `json:"code"`            // 股票代码
	Name           string    `json:"name"`            // 股票名称
	TotalShare     float64   `json:"total_share"`     // 总股本
	CirculateShare float64   `json:"circulate_share"` // 流通股本
	EndDate        time.Time `json:"end_date"`        // 截止日期
}

// StockRestrictedStockEm 东方财富限售股
type StockRestrictedStockEm struct {
	Code            string    `json:"code"`             // 股票代码
	Name            string    `json:"name"`             // 股票名称
	HolderName      string    `json:"holder_name"`      // 股东名称
	RestrictedShare float64   `json:"restricted_share"` // 限售股数
	LiftDate        time.Time `json:"lift_date"`        // 解禁日期
}

// StockZyjsEm 东方财富重大股改（重命名以避免与函数名冲突）
type StockZyjsEm struct {
	Code         string    `json:"code"`          // 股票代码
	Name         string    `json:"name"`          // 股票名称
	ProcessType  string    `json:"process_type"`  // 流程类型
	AnnounceDate time.Time `json:"announce_date"` // 公告日期
}

// StockGbjgEm 东方财富股本结构
type StockGbjgEm struct {
	Code       string    `json:"code"`        // 股票代码
	Name       string    `json:"name"`        // 股票名称
	ChangeDate time.Time `json:"change_date"` // 变更日期
	ChangeType string    `json:"change_type"` // 变更类型
	ChangeDesc string    `json:"change_desc"` // 变动描述
}

// StockFinancialAnalysisIndicatorEmItem 东方财富财务分析指标项
type StockFinancialAnalysisIndicatorEmItem struct {
	ReportDate        string  `json:"report_date"`         // 报告期
	ROE               float64 `json:"roe"`                 // 净资产收益率
	ROA               float64 `json:"roa"`                 // 总资产净利率
	GrossProfitMargin float64 `json:"gross_profit_margin"` // 销售毛利率
	NetProfitMargin   float64 `json:"net_profit_margin"`   // 销售净利率
	// 更多字段根据实际响应添加
}

// StockHistoryDividendItem 历史分红数据项
type StockHistoryDividendItem struct {
	Code            string    `json:"code"`             // 股票代码
	Name            string    `json:"name"`             // 股票名称
	ListDate        time.Time `json:"list_date"`        // 上市日期
	TotalDividend   float64   `json:"total_dividend"`   // 累计股息
	AvgDividend     float64   `json:"avg_dividend"`     // 年均股息
	DividendCount   int       `json:"dividend_count"`   // 分红次数
	FinancingAmount float64   `json:"financing_amount"` // 融资总额
	FinancingCount  int       `json:"financing_count"`  // 融资次数
}

// StockHistoryDividendDetailItem 分红详情项
type StockHistoryDividendDetailItem struct {
	AnnounceDate   time.Time `json:"announce_date"`    // 公告日期
	BonusShare     float64   `json:"bonus_share"`      // 送股
	TransferShare  float64   `json:"transfer_share"`   // 转增
	Dividend       float64   `json:"dividend"`         // 派息
	Status         string    `json:"status"`           // 进度
	ExDividendDate time.Time `json:"ex_dividend_date"` // 除权除息日
	RecordDate     time.Time `json:"record_date"`      // 股权登记日
	ListingDate    time.Time `json:"listing_date"`     // 红股上市日
}

// StockIpoInfoItem 新股发行信息项
type StockIpoInfoItem struct {
	Item  string `json:"item"`  // 项目
	Value string `json:"value"` // 数值
}

// StockMainStockHolderItem 主要股东信息项
type StockMainStockHolderItem struct {
	Code         string    `json:"code"`          // 股票代码
	Name         string    `json:"name"`          // 股票名称
	HolderName   string    `json:"holder_name"`   // 股东名称
	Holdings     float64   `json:"holdings"`      // 持股数量
	HoldingRatio float64   `json:"holding_ratio"` // 持股比例
	EndDate      time.Time `json:"end_date"`      // 截止日期
}

// StockFundStockHolderItem 基金持股信息项
type StockFundStockHolderItem struct {
	Code          string    `json:"code"`            // 股票代码
	Name          string    `json:"name"`            // 股票名称
	FundName      string    `json:"fund_name"`       // 基金名称
	FundCode      string    `json:"fund_code"`       // 基金代码
	Holdings      float64   `json:"holdings"`        // 持仓数量
	HoldingRatio  float64   `json:"holding_ratio"`   // 占流通股比例
	MarketValue   float64   `json:"market_value"`    // 持股市值
	NetValueRatio float64   `json:"net_value_ratio"` // 占净值比例
	EndDate       time.Time `json:"end_date"`        // 截止日期
}

// StockInstituteHoldItem 新浪财经-机构持股一览表项
type StockInstituteHoldItem struct {
	Code                   string  `json:"code"`                     // 证券代码
	Name                   string  `json:"name"`                     // 证券简称
	InstituteCount         float64 `json:"institute_count"`          // 机构数
	InstituteCountChange   float64 `json:"institute_count_change"`   // 机构数变化
	HoldingRatio           float64 `json:"holding_ratio"`            // 持股比例
	HoldingRatioChange     float64 `json:"holding_ratio_change"`     // 持股比例增幅
	CirculationRatio       float64 `json:"circulation_ratio"`        // 占流通股比例
	CirculationRatioChange float64 `json:"circulation_ratio_change"` // 占流通股比例增幅
}

// StockInstituteHoldDetailItem 新浪财经-机构持股详情项
type StockInstituteHoldDetailItem struct {
	InstituteType          string  `json:"institute_type"`           // 持股机构类型
	InstituteCode          string  `json:"institute_code"`           // 持股机构代码
	InstituteName          string  `json:"institute_name"`           // 持股机构简称
	InstituteFullName      string  `json:"institute_full_name"`      // 持股机构全称
	Holdings               float64 `json:"holdings"`                 // 持股数
	LatestHoldings         float64 `json:"latest_holdings"`          // 最新持股数
	HoldingRatio           float64 `json:"holding_ratio"`            // 持股比例
	LatestHoldingRatio     float64 `json:"latest_holding_ratio"`     // 最新持股比例
	CirculationRatio       float64 `json:"circulation_ratio"`        // 占流通股比例
	LatestCirculationRatio float64 `json:"latest_circulation_ratio"` // 最新占流通股比例
	HoldingRatioChange     float64 `json:"holding_ratio_change"`     // 持股比例增幅
	CirculationRatioChange float64 `json:"circulation_ratio_change"` // 占流通股比例增幅
}

// StockZhAGbjgEmItem 东方财富-股本结构项
type StockZhAGbjgEmItem struct {
	ChangeDate      time.Time `json:"change_date"`      // 变更日期
	TotalShares     float64   `json:"total_shares"`     // 总股本
	LimitedShares   float64   `json:"limited_shares"`   // 流通受限股份
	LimitedOthARS   float64   `json:"limited_oth_ars"`  // 其他内资持股(受限)
	LimitedDomestic float64   `json:"limited_domestic"` // 境内法人持股(受限)
	LimitedNatural  float64   `json:"limited_natural"`  // 境内自然人持股(受限)
	UnlimitedShares float64   `json:"unlimited_shares"` // 已流通股份
	ListedAShares   float64   `json:"listed_a_shares"`  // 已上市流通A股
	ChangeReason    string    `json:"change_reason"`    // 变动原因
}

// StockZygcEmItem 东方财富-主营构成项
type StockZygcEmItem struct {
	Code             string    `json:"code"`               // 股票代码
	ReportDate       time.Time `json:"report_date"`        // 报告日期
	CategoryType     string    `json:"category_type"`      // 分类类型
	MainComposition  string    `json:"main_composition"`   // 主营构成
	MainIncome       float64   `json:"main_income"`        // 主营收入
	IncomeRatio      float64   `json:"income_ratio"`       // 收入比例
	MainCost         float64   `json:"main_cost"`          // 主营成本
	CostRatio        float64   `json:"cost_ratio"`         // 成本比例
	MainProfit       float64   `json:"main_profit"`        // 主营利润
	ProfitRatio      float64   `json:"profit_ratio"`       // 利润比例
	GrossProfitRatio float64   `json:"gross_profit_ratio"` // 毛利率
}

// StockProfitForecastEmItem 东方财富盈利预测项
type StockProfitForecastEmItem struct {
	Index         int     `json:"index"`          // 序号
	Code          string  `json:"code"`           // 代码
	Name          string  `json:"name"`           // 名称
	ReportCount   int     `json:"report_count"`   // 研报数
	RatingBuy     int     `json:"rating_buy"`     // 机构投资评级(近六个月)-买入
	RatingAdd     int     `json:"rating_add"`     // 机构投资评级(近六个月)-增持
	RatingNeutral int     `json:"rating_neutral"` // 机构投资评级(近六个月)-中性
	RatingReduce  int     `json:"rating_reduce"`  // 机构投资评级(近六个月)-减持
	RatingSell    int     `json:"rating_sell"`    // 机构投资评级(近六个月)-卖出
	Year1Forecast float64 `json:"year1_forecast"` // 第一年预测每股收益
	Year2Forecast float64 `json:"year2_forecast"` // 第二年预测每股收益
	Year3Forecast float64 `json:"year3_forecast"` // 第三年预测每股收益
	Year4Forecast float64 `json:"year4_forecast"` // 第四年预测每股收益
}

// StockRegisterEmItem IPO审核信息项
type StockRegisterEmItem struct {
	Index         int       `json:"index"`          // 序号
	CompanyName   string    `json:"company_name"`   // 企业名称
	State         string    `json:"state"`          // 最新状态
	RegAddress    string    `json:"reg_address"`    // 注册地
	Industry      string    `json:"industry"`       // 行业
	RecommendOrg  string    `json:"recommend_org"`  // 保荐机构
	LawFirm       string    `json:"law_firm"`       // 律师事务所
	AccountFirm   string    `json:"account_firm"`   // 会计师事务所
	UpdateDate    time.Time `json:"update_date"`    // 更新日期
	AcceptDate    time.Time `json:"accept_date"`    // 受理日期
	PredictMarket string    `json:"predict_market"` // 拟上市地点
	ProspectusURL string    `json:"prospectus_url"` // 招股说明书
}

// StockRestrictedReleaseSummaryEmItem 限售股解禁汇总项
type StockRestrictedReleaseSummaryEmItem struct {
	Index               int       `json:"index"`              // 序号
	LiftDate            time.Time `json:"lift_date"`          // 解禁时间
	StockCount          int       `json:"stock_count"`        // 当日解禁股票家数
	LiftShares          float64   `json:"lift_shares"`        // 解禁数量
	ActualLiftShares    float64   `json:"actual_lift_shares"` // 实际解禁数量
	ActualLiftMarketCap float64   `json:"actual_lift_cap"`    // 实际解禁市值
	HS300Index          float64   `json:"hs300_index"`        // 沪深300指数
	HS300ChangeRatio    float64   `json:"hs300_change_ratio"` // 沪深300指数涨跌幅
}

// StockRestrictedReleaseDetailEmItem 限售股解禁详情项
type StockRestrictedReleaseDetailEmItem struct {
	Index               int       `json:"index"`              // 序号
	Code                string    `json:"code"`               // 股票代码
	Name                string    `json:"name"`               // 股票简称
	LiftDate            time.Time `json:"lift_date"`          // 解禁时间
	SharesType          string    `json:"shares_type"`        // 限售股类型
	LiftShares          float64   `json:"lift_shares"`        // 解禁数量
	ActualLiftShares    float64   `json:"actual_lift_shares"` // 实际解禁数量
	ActualLiftMarketCap float64   `json:"actual_lift_cap"`    // 实际解禁市值
	CirculationRatio    float64   `json:"circulation_ratio"`  // 占解禁前流通市值比例
	PrevClosePrice      float64   `json:"prev_close_price"`   // 解禁前一交易日收盘价
	Before20ChangeRatio float64   `json:"before_20_change"`   // 解禁前20日涨跌幅
	After20ChangeRatio  float64   `json:"after_20_change"`    // 解禁后20日涨跌幅
}

// StockRestrictedReleaseQueueEmItem 个股限售解禁批次项
type StockRestrictedReleaseQueueEmItem struct {
	Index               int       `json:"index"`              // 序号
	LiftDate            time.Time `json:"lift_date"`          // 解禁时间
	HolderNum           int       `json:"holder_num"`         // 解禁股东数
	LiftShares          float64   `json:"lift_shares"`        // 解禁数量
	ActualLiftShares    float64   `json:"actual_lift_shares"` // 实际解禁数量
	NonLiftShares       float64   `json:"non_lift_shares"`    // 未解禁数量
	ActualLiftMarketCap float64   `json:"actual_lift_cap"`    // 实际解禁数量市值
	TotalRatio          float64   `json:"total_ratio"`        // 占总市值比例
	CirculationRatio    float64   `json:"circulation_ratio"`  // 占流通市值比例
	PrevClosePrice      float64   `json:"prev_close_price"`   // 解禁前一交易日收盘价
	SharesType          string    `json:"shares_type"`        // 限售股类型
	Before20ChangeRatio float64   `json:"before_20_change"`   // 解禁前20日涨跌幅
	After20ChangeRatio  float64   `json:"after_20_change"`    // 解禁后20日涨跌幅
}

// StockProfitForecastThsItem 同花顺盈利预测项
type StockProfitForecastThsItem struct {
	Year      string  `json:"year"`      // 年度
	OrgNum    int     `json:"org_num"`   // 机构数
	Avg       float64 `json:"avg"`       // 平均值
	Max       float64 `json:"max"`       // 最大值
	Min       float64 `json:"min"`       // 最小值
	Indicator string  `json:"indicator"` // 指标类型
}

// StockProfitForecastOrgThsItem 同花顺业绩预测详表-机构项
type StockProfitForecastOrgThsItem struct {
	OrgName    string    `json:"org_name"`    // 机构名称
	ReportDate time.Time `json:"report_date"` // 报告日期
	EPS1Avg    float64   `json:"eps1_avg"`    // 预测年报每股收益-平均
	EPS1Max    float64   `json:"eps1_max"`    // 预测年报每股收益-最大
	EPS1Min    float64   `json:"eps1_min"`    // 预测年报每股收益-最小
	ProfitAvg  float64   `json:"profit_avg"`  // 预测年报净利润-平均
	ProfitMax  float64   `json:"profit_max"`  // 预测年报净利润-最大
	ProfitMin  float64   `json:"profit_min"`  // 预测年报净利润-最小
}

// StockHkProfitForecastEtItem 经济通港股盈利预测项
type StockHkProfitForecastEtItem struct {
	FiscalYear  string    `json:"fiscal_year"`  // 财政年度
	PureProfit  float64   `json:"pure_profit"`  // 纯利/亏损
	EPS         float64   `json:"eps"`          // 每股盈利/每股亏损
	Dividend    float64   `json:"dividend"`     // 每股派息
	NAV         float64   `json:"nav"`          // 每股资产净值
	High        float64   `json:"high"`         // 最高
	Low         float64   `json:"low"`          // 最低
	TargetPrice float64   `json:"target_price"` // 目标价
	UpdateDate  time.Time `json:"update_date"`  // 更新日期
	Indicator   string    `json:"indicator"`    // 指标类型
}
