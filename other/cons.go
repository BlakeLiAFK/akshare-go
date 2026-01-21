package other

// 乘联会相关URL
const (
	CPCAChartListURL  = "http://data.cpcadata.com/api/chartlist"
	CPCAChartList2URL = "http://data.cpcadata.com/api/chartlist_2"
)

// 盖世汽车相关URL
const (
	GasgooSalesRankURL = "https://i.gasgoo.com/data/sales/AutoModelSalesRank.aspx/GetSalesRank"
)

// 乘联会参数常量
const (
	// 车型类型
	CarTypeNarrow = "狭义乘用车"
	CarTypeWide   = "广义乘用车"

	// 指标类型
	IndicatorProduction = "产量"
	IndicatorWholesale  = "批发"
	IndicatorRetail     = "零售"
	IndicatorExport     = "出口"

	// 统计类型
	StatTypeMonthly    = "单月"
	StatTypeCumulative = "累计"

	// 车型大类
	CategorySedan = "轿车"
	CategoryMPV   = "MPV"
	CategorySUV   = "SUV"
	CategoryRatio = "占比"

	// 燃料类型
	FuelOverallMarket = "整体市场"
	FuelRatioPHEVBEV  = "销量占比-PHEV-BEV"
	FuelRatioICENEV   = "销量占比-ICE-NEV"
)

// 盖世汽车参数常量
const (
	// 排行榜类型
	RankTypeModel = "车型榜"
	RankTypeFirm  = "车企榜"
	RankTypeBrand = "品牌榜"
)

// symbol映射表 - 车型类型
var symbolMap = map[string]int{
	CarTypeNarrow: 0,
	CarTypeWide:   1,
}

// indicator映射表 - 指标类型
var indicatorMap = map[string]int{
	IndicatorProduction: 0,
	IndicatorWholesale:  1,
	IndicatorRetail:     2,
	IndicatorExport:     3,
}

// 厂商排名symbol映射表
var manRankSymbolMap = map[string]int{
	CarTypeNarrow + "-" + StatTypeCumulative: 0,
	CarTypeNarrow + "-" + StatTypeMonthly:    1,
	CarTypeWide + "-" + StatTypeCumulative:   2,
	CarTypeWide + "-" + StatTypeMonthly:      3,
}

// 车型大类symbol映射表
var categorySymbolMap = map[string]int{
	CategoryMPV:   0,
	CategorySUV:   1,
	CategorySedan: 2,
	CategoryRatio: 3,
}

// 级别细分symbol映射表
var segmentSymbolMap = map[string]int{
	CategoryMPV:   0,
	CategorySUV:   1,
	CategorySedan: 2,
}

// 燃料类型symbol映射表
var fuelSymbolMap = map[string]int{
	FuelOverallMarket: 0,
	FuelRatioPHEVBEV:  1,
	FuelRatioICENEV:   2,
}

// 盖世汽车排行榜类型映射表
var rankTypeMap = map[string]string{
	RankTypeModel: "M",
	RankTypeFirm:  "F",
	RankTypeBrand: "B",
}
