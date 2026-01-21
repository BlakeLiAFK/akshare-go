package economic

import "time"

// MacroEconomicItem 宏观经济指标数据项（通用）
// 用于金十数据中心的大部分宏观经济指标
type MacroEconomicItem struct {
	Commodity string    `json:"商品"`  // 商品/指标名称
	Date      time.Time `json:"日期"`  // 日期
	Current   float64   `json:"今值"`  // 今值（当前值）
	Forecast  float64   `json:"预测值"` // 预测值
	Previous  float64   `json:"前值"`  // 前值（上期值）
}

// InterestRateItem 央行利率决议数据项
// 用于各国央行利率决议报告
type InterestRateItem struct {
	Date     time.Time `json:"日期"`  // 日期
	Current  float64   `json:"今值"`  // 今值（当前利率）
	Forecast float64   `json:"预测值"` // 预测值
	Previous float64   `json:"前值"`  // 前值（上期利率）
}

// NBSDataItem 国家统计局数据项
// 用于中国国家统计局宏观数据
type NBSDataItem struct {
	Code  string  `json:"code"`  // 指标代码
	Name  string  `json:"name"`  // 指标名称
	Value float64 `json:"value"` // 数据值
	Unit  string  `json:"unit"`  // 单位
	Date  string  `json:"date"`  // 日期（格式根据数据源而定）
}

// RigCountItem 贝克休斯钻井平台数据项
// 用于美国贝克休斯钻井报告
type RigCountItem struct {
	Date                   string  `json:"日期"`          // 日期
	TotalRigCount          float64 `json:"钻井总数_钻井数"`    // 钻井总数-钻井数
	TotalRigChange         float64 `json:"钻井总数_变化"`     // 钻井总数-变化
	USAOilRigCount         float64 `json:"美国石油钻井_钻井数"`  // 美国石油钻井-钻井数
	USAOilRigChange        float64 `json:"美国石油钻井_变化"`   // 美国石油钻井-变化
	MixedRigCount          float64 `json:"混合钻井_钻井数"`    // 混合钻井-钻井数
	MixedRigChange         float64 `json:"混合钻井_变化"`     // 混合钻井-变化
	USANaturalGasRigCount  float64 `json:"美国天然气钻井_钻井数"` // 美国天然气钻井-钻井数
	USANaturalGasRigChange float64 `json:"美国天然气钻井_变化"`  // 美国天然气钻井-变化
}

// EastMoneyEconomicItem 东方财富宏观经济数据项
// 用于东方财富经济数据接口
type EastMoneyEconomicItem struct {
	Time        string  `json:"时间"`   // 时间
	Previous    float64 `json:"前值"`   // 前值
	Current     float64 `json:"现值"`   // 现值
	PublishDate string  `json:"发布日期"` // 发布日期
}

// CrudeOilProductionItem 美国原油产量数据项
// 用于美国原油产量报告
type CrudeOilProductionItem struct {
	Date                 string  `json:"日期"`             // 日期
	TotalProduction      float64 `json:"美国国内原油总量_产量"`    // 美国国内原油总量-产量
	TotalChange          float64 `json:"美国国内原油总量_变化"`    // 美国国内原油总量-变化
	Mainland48Production float64 `json:"美国本土48州原油产量_产量"` // 美国本土48州原油产量-产量
	Mainland48Change     float64 `json:"美国本土48州原油产量_变化"` // 美国本土48州原油产量-变化
	AlaskaProduction     float64 `json:"美国阿拉斯加州原油产量_产量"` // 美国阿拉斯加州原油产量-产量
	AlaskaChange         float64 `json:"美国阿拉斯加州原油产量_变化"` // 美国阿拉斯加州原油产量-变化
}

// CFTCItem CFTC持仓数据项（通用）
// 用于美国商品期货交易委员会CFTC各类持仓报告
// 由于列数动态变化，使用map存储
type CFTCItem struct {
	Date string             `json:"日期"` // 日期
	Data map[string]float64 `json:"数据"` // 动态数据，键为"品种-指标"
}

// CMEItem CME贵金属数据项
// 用于CME贵金属报告
type CMEItem struct {
	Date    string  `json:"日期"`  // 日期
	Variety string  `json:"品种"`  // 品种
	Volume  float64 `json:"成交量"` // 成交量
}

// jin10CommonHeaders 金十数据中心通用请求头
// 用于所有金十数据中心接口
var jin10CommonHeaders = map[string]string{
	"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36",
	"x-app-id":     "rU6QIu7JHe2gOUeR",
	"x-csrf-token": "x-csrf-token",
	"x-version":    "1.0.0",
}
