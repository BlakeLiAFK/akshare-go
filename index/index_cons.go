package index

// 新浪财经相关URL
const (
	// SinaIndexStockURL 新浪指数行情接口
	SinaIndexStockURL = "http://vip.stock.finance.sina.com.cn/quotes_service/api/json_v2.php/Market_Center.getHQNodeDataSimple"
	// SinaIndexCountURL 新浪指数总数接口
	SinaIndexCountURL = "http://vip.stock.finance.sina.com.cn/quotes_service/api/json_v2.php/Market_Center.getHQNodeStockCountSimple?node=hs_s"
	// SinaIndexHistURL 新浪指数历史数据接口
	SinaIndexHistURL = "https://finance.sina.com.cn/realstock/company/%s/hisdata/klc_kl.js"
	// SinaHKIndexCountURL 新浪港股指数总数接口
	SinaHKIndexCountURL = "https://vip.stock.finance.sina.com.cn/quotes_service/api/json_v2.php/Market_Center.getNameCount?node=zs_hk"
)

// 东方财富相关URL
const (
	// EmIndexListURL 东财指数列表接口
	EmIndexListURL = "https://push2.eastmoney.com/api/qt/clist/get"
	// EmIndexKLineURL 东财指数K线接口
	EmIndexKLineURL = "https://push2his.eastmoney.com/api/qt/stock/kline/get"
	// EmIndexTrendsURL 东财指数分时接口
	EmIndexTrendsURL = "https://push2his.eastmoney.com/api/qt/stock/trends2/get"
)

// 其他数据源URL
const (
	// SinaGlobalIndexURL 新浪全球指数接口
	SinaGlobalIndexURL = "https://gi.finance.sina.com.cn/hq/daily"
	// SinaSpotGoodsURL 新浪商品现货接口
	SinaSpotGoodsURL = "https://stock.finance.sina.com.cn/futures/api/openapi.php/GoodsIndexService.get_goods_index"
	// CSIndexExportURL 中证指数导出接口
	CSIndexExportURL = "https://www.csindex.com.cn/csindex-home/exportExcel/indexAll/CH"
	// CNIIndexListURL 国证指数列表接口
	CNIIndexListURL = "https://www.cnindex.com.cn/index/indexList"
	// CNIIndexHistURL 国证指数历史数据接口
	CNIIndexHistURL = "http://hq.cnindex.com.cn/market/market/getIndexDailyDataWithDataFormat"
	// CNIIndexDetailURL 国证指数成份股接口
	CNIIndexDetailURL = "https://www.cnindex.com.cn/sample-detail/download-history"
	// CNIIndexAdjustURL 国证指数调样接口
	CNIIndexAdjustURL = "http://www.cnindex.com.cn/sample-detail/download-adjustment"
	// HKIndexSinaURL 新浪港股指数实时行情接口
	HKIndexSinaURL = "https://hq.sinajs.cn/rn=mtf2t&list=%s"
	// USIndexSinaURL 新浪美股指数接口
	USIndexSinaURL = "https://finance.sina.com.cn/staticdata/us/%s"
	// SWIndexOverviewURL 申万行业概览URL
	SWIndexOverviewURL = "https://legulegu.com/stockdata/sw-industry-overview"
	// SWIndexConsURL 申万行业成份股URL
	SWIndexConsURL = "https://legulegu.com/stockdata/index-composition?industryCode=%s"
	// DrewryWCIURL Drewry集装箱指数URL
	DrewryWCIURL = "https://infogram.com/world-container-index-1h17493095xl4zj"
	// KQIndexTableDataURL 柯桥纺织指数数据接口
	KQIndexTableDataURL = "http://www.kqindex.cn/flzs/table_data"
	// KQFashionIndexURL 柯桥时尚指数接口
	KQFashionIndexURL = "http://api.idx365.com/index/project/34/data"
	// YWIndexURL 义乌小商品指数接口
	YWIndexURL = "https://apiserver.chinagoods.com/yiwuindex/v1/active/industry/class/history/%s?gcCode="
	// SugarIndexURL 沐甜科技食糖指数接口
	SugarIndexURL = "https://www.msweet.com.cn/eportal/ui"
	// SugarInnerQuoteURL 配额内进口糖估算指数接口
	SugarInnerQuoteURL = "https://www.msweet.com.cn/datacenterapply/datacenter/json/JinKongTang.json"
	// SugarOuterQuoteURL 配额外进口糖估算指数接口
	SugarOuterQuoteURL = "https://www.msweet.com.cn/datacenterapply/datacenter/json/Jkpewlr.json"
	// ERIndexURL 浙江省排污权交易指数接口
	ERIndexURL = "https://zs.zjpwq.net/pwq-index-webapi/indexData"
	// HogIndexURL 生猪市场价格指数接口
	HogIndexURL = "https://hqb.nxin.com/pigindex/getPigIndexChart.shtml"
	// CFLPPriceIndexURL 中国公路物流运价指数接口
	CFLPPriceIndexURL = "http://index.0256.cn/expcenter_trend.action"
	// CFLPVolumeIndexURL 中国公路物流运量指数接口
	CFLPVolumeIndexURL = "http://index.0256.cn/volume_query.action"
)

// 周期映射
var PeriodMap = map[string]string{
	"daily":   "101",
	"weekly":  "102",
	"monthly": "103",
	"1":       "1",
	"5":       "5",
	"15":      "15",
	"30":      "30",
	"60":      "60",
}

// 市场代码映射
var MarketMap = map[string]string{
	"sz":  "0",  // 深交所
	"sh":  "1",  // 上交所
	"csi": "2",  // 中证指数
	"bj":  "47", // 北交所
}

// GlobalSinaSymbolMap 新浪全球指数代码映射
var GlobalSinaSymbolMap = map[string]string{
	// 欧洲股市
	"英国富时100指数":     "UKX",
	"德国DAX 30种股价指数": "DAX",
	"俄罗斯MICEX指数":    "INDEXCF",
	"法CAC40指数":      "CAC",
	"瑞士股票指数":        "SWI20",
	"富时意大利MIB指数":    "FTSEMIB",
	"荷兰AEX综合指数":     "AEX",
	"西班牙IBEX指数":     "IBEX",
	"欧洲Stoxx50指数":   "SX5E",
	// 美洲股市
	"加拿大S&P/TSX综合指数": "GSPTSE",
	"墨西哥BOLSA指数":     "MXX",
	"巴西BOVESPA股票指数":  "IBOV",
	// 亚洲股市
	"中国台湾加权指数":     "TWJQ",
	"日经225指数":      "NKY",
	"首尔综合指数":       "KOSPI",
	"印度尼西亚雅加达综合指数": "JCI",
	"印度孟买SENSEX指数": "SENSEX",
	// 澳洲股市
	"澳大利亚标准普尔200指数": "AS51",
	"新西兰NZSE 50指数":  "NZ250",
	// 非洲股市
	"埃及CASE 30指数": "CASE",
}

// GlobalEMSymbolMap 东方财富全球指数代码映射
var GlobalEMSymbolMap = map[string]IndexSymbolCode{
	"波罗的海BDI指数":  {Code: "BDI", MarketID: "100"},
	"上证指数":       {Code: "000001", MarketID: "1"},
	"沪深300":      {Code: "000300", MarketID: "1"},
	"深证成指":       {Code: "399001", MarketID: "0"},
	"中小100":      {Code: "399005", MarketID: "0"},
	"创业板指":       {Code: "399006", MarketID: "0"},
	"恒生指数":       {Code: "HSI", MarketID: "100"},
	"国企指数":       {Code: "HSCEI", MarketID: "100"},
	"红筹指数":       {Code: "HSCCI", MarketID: "124"},
	"台湾加权":       {Code: "TWII", MarketID: "100"},
	"日经225":      {Code: "N225", MarketID: "100"},
	"韩国KOSPI":    {Code: "KS11", MarketID: "100"},
	"韩国KOSPI200": {Code: "KOSPI200", MarketID: "100"},
	"标普500":      {Code: "SPX", MarketID: "100"},
	"纳斯达克":       {Code: "NDX", MarketID: "100"},
	"道琼斯":        {Code: "DJIA", MarketID: "100"},
	"英国富时100":    {Code: "FTSE", MarketID: "100"},
	"德国DAX30":    {Code: "GDAXI", MarketID: "100"},
	"法国CAC40":    {Code: "FCHI", MarketID: "100"},
	"欧洲斯托克50":    {Code: "SX5E", MarketID: "100"},
	"美元指数":       {Code: "UDI", MarketID: "100"},
}

// HKIndexList 新浪港股指数代码列表
var HKIndexList = []string{
	"CES100", "CES120", "CES280", "CES300", "CESA80", "CESG10",
	"CESHKM", "CSCMC", "CSHK100", "CSHKDIV", "CSHKLC", "CSHKLRE", "CSHKMCS", "CSHKME",
	"CSHKPE", "CSHKSE", "CSI300", "CSRHK50", "GEM", "HKL", "HSCCI", "HSCEI", "HSI",
	"HSMBI", "HSMOGI", "HSMPI", "HSTECH", "SSE180", "SSE180GV", "SSE380", "SSE50",
	"SSECEQT", "SSECOMP", "SSEDIV", "SSEITOP", "SSEMCAP", "SSEMEGA", "VHSI",
}

// SpotGoodsSymbolMap 商品现货指数代码映射
var SpotGoodsSymbolMap = map[string]string{
	"波罗的海干散货指数": "BDI",
	"钢坯价格指数":    "GP",
	"澳大利亚粉矿价格":  "PB",
}

// SWIndexTypeMap 申万行业指数类型映射
var SWIndexTypeMap = map[string]string{
	"价格指数": "1_1",
	"景气指数": "1_2",
	"外贸指数": "2",
}

// SWFashionIndexMap 柯桥时尚指数类型映射
var SWFashionIndexMap = map[string]string{
	"柯桥时尚指数":   "root",
	"时尚创意指数":   "01",
	"时尚设计人才数":  "0101",
	"新花型推出数":   "0102",
	"创意产品成交数":  "0103",
	"创意企业数量":   "0104",
	"时尚活跃度指数":  "02",
	"电商运行数":    "0201",
	"时尚平台拓展数":  "0201",
	"新产品销售额占比": "0201",
	"企业合作占比":   "0201",
	"品牌传播费用":   "0201",
	"时尚推广度指数":  "03",
	"国际交流合作次数": "0301",
	"企业参展次数":   "0302",
	"外商驻点数量变化": "0302",
	"时尚评价指数":   "04",
}

// YWIndexTypeMap 义乌小商品指数类型映射
var YWIndexTypeMap = map[string]string{
	"周价格指数": "piweek",
	"月价格指数": "month",
}

// ERIndexTypeMap 排污权交易指数类型映射
var ERIndexTypeMap = map[string]string{
	"月度": "MONTH",
	"季度": "QUARTER",
}

// CFLPPriceIndexTypeMap 公路物流运价指数类型映射
var CFLPPriceIndexTypeMap = map[string]string{
	"周指数":  "2",
	"月指数":  "3",
	"季度指数": "4",
	"年度指数": "5",
}

// CFLPVolumeIndexTypeMap 公路物流运量指数类型映射
var CFLPVolumeIndexTypeMap = map[string]string{
	"月指数":  "3",
	"季度指数": "4",
	"年度指数": "5",
}

// DrewryWCISymbolMap Drewry集装箱航线映射
var DrewryWCISymbolMap = map[string]int{
	"composite":            0,
	"shanghai-rotterdam":   1,
	"rotterdam-shanghai":   2,
	"shanghai-los angeles": 3,
	"los angeles-shanghai": 4,
	"shanghai-genoa":       5,
	"new york-rotterdam":   6,
	"rotterdam-new york":   7,
}

// USIndexSymbolMap 美股指数代码映射
var USIndexSymbolMap = map[string]string{
	".IXIC": "纳斯达克",
	".DJI":  "道琼斯",
	".INX":  "标普500",
	".NDX":  "纳斯达克100",
}
