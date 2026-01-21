package news

// StockNews 个股新闻数据结构
// 对应 Python: stock_news_em 返回的 DataFrame
type StockNews struct {
	Keyword     string `json:"keyword"`      // 关键词
	Title       string `json:"title"`        // 新闻标题
	Content     string `json:"content"`      // 新闻内容
	PublishTime string `json:"publish_time"` // 发布时间
	Source      string `json:"source"`       // 文章来源
	Url         string `json:"url"`          // 新闻链接
}

// CctvNews 新闻联播文字稿数据结构
// 对应 Python: news_cctv 返回的 DataFrame
type CctvNews struct {
	Date    string `json:"date"`    // 日期
	Title   string `json:"title"`   // 标题
	Content string `json:"content"` // 内容
}

// EconomicData 百度股市通-经济数据
// 对应 Python: news_economic_baidu 返回的 DataFrame
type EconomicData struct {
	Date       string  `json:"date"`       // 日期
	Time       string  `json:"time"`       // 时间
	Country    string  `json:"country"`    // 国家
	Region     string  `json:"region"`     // 地区
	Event      string  `json:"event"`      // 事件
	Period     string  `json:"period"`     // 统计周期
	Actual     float64 `json:"actual"`     // 公布
	Forecast   float64 `json:"forecast"`   // 预期
	Previous   float64 `json:"previous"`   // 前值
	Importance int     `json:"importance"` // 重要性
}

// SuspendNotify 百度股市通-停复牌提醒
// 对应 Python: news_trade_notify_suspend_baidu 返回的 DataFrame
type SuspendNotify struct {
	Code         string `json:"code"`          // 股票代码
	Name         string `json:"name"`          // 股票简称
	Exchange     string `json:"exchange"`      // 交易所代码
	SuspendTime  string `json:"suspend_time"`  // 停牌时间
	ResumeTime   string `json:"resume_time"`   // 复牌时间
	Reason       string `json:"reason"`        // 停牌事项说明
	MarketValue  string `json:"market_value"`  // 市值
	AnnounceDate string `json:"announce_date"` // 公告日期
	AnnounceTime string `json:"announce_time"` // 公告时间
	SecType      string `json:"sec_type"`      // 证券类型
	MarketType   string `json:"market_type"`   // 市场类型
	IsSkip       string `json:"is_skip"`       // 是否跳过
}

// DividendNotify 百度股市通-分红派息提醒
// 对应 Python: news_trade_notify_dividend_baidu 返回的 DataFrame
type DividendNotify struct {
	Code       string `json:"code"`        // 股票代码
	ExDate     string `json:"ex_date"`     // 除权日
	Dividend   string `json:"dividend"`    // 分红
	Bonus      string `json:"bonus"`       // 送股
	Transfer   string `json:"transfer"`    // 转增
	Physical   string `json:"physical"`    // 实物
	Exchange   string `json:"exchange"`    // 交易所
	Name       string `json:"name"`        // 股票简称
	ReportDate string `json:"report_date"` // 报告期
}

// ReportTime 百度股市通-财报发行
// 对应 Python: news_report_time_baidu 返回的 DataFrame
type ReportTime struct {
	Code        string  `json:"code"`         // 股票代码
	Name        string  `json:"name"`         // 股票简称
	Exchange    string  `json:"exchange"`     // 交易所
	ReportType  string  `json:"report_type"`  // 财报类型
	PublishTime string  `json:"publish_time"` // 发布时间
	MarketValue float64 `json:"market_value"` // 市值
	PublishDate string  `json:"publish_date"` // 发布日期
}
