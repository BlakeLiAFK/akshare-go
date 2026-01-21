package futures

// 期货配置常量

// FuturesInventoryEMSymbolDict 期货库存品种代码映射 (东方财富)
var FuturesInventoryEMSymbolDict = map[string]string{
	"a":  "A",  // 豆一
	"ag": "AG", // 沪银
	"al": "AL", // 沪铝
	"ao": "AO", // 氧化铝
	"AP": "AP", // 苹果
	"au": "AU", // 沪金
	"b":  "B",  // 豆二
	"br": "BR", // BR橡胶
	"bu": "BU", // 沥青
	"c":  "C",  // 玉米
	"CF": "CF", // 棉花/郑棉
	"CJ": "CJ", // 红枣
	"cs": "CS", // 淀粉/玉米淀粉
	"cu": "CU", // 沪铜
	"CY": "CY", // 棉纱
	"eb": "EB", // 苯乙烯
	"ec": "ec", // 集运欧线
	"eg": "EG", // 乙二醇
	"FG": "FG", // 玻璃
	"fu": "FU", // 燃料油
	"hc": "HC", // 热卷
	"i":  "I",  // 铁矿石
	"IC": "IC", // 中证500
	"IF": "IF", // 沪深300
	"IH": "IH", // 上证50
	"IM": "IM", // 中证1000
	"j":  "J",  // 焦炭
	"jd": "JD", // 鸡蛋
	"jm": "JM", // 焦煤
	"l":  "L",  // 塑料
	"lc": "lc", // 碳酸锂
	"lh": "LH", // 生猪
	"lu": "lu", // 低硫燃料油
	"m":  "M",  // 豆粕
	"MA": "MA", // 甲醇
	"ni": "NI", // 沪镍
	"nr": "nr", // 20号胶
	"OI": "OI", // 菜籽油
	"p":  "P",  // 棕榈油
	"pb": "PB", // 沪铅
	"PF": "PF", // 短纤
	"pg": "PG", // 液化气
	"PK": "PK", // 花生
	"pp": "PP", // 聚丙烯
	"PX": "PX", // 对二甲苯
	"rb": "RB", // 螺纹钢
	"RM": "RM", // 菜籽粕
	"RS": "RS", // 油菜籽
	"ru": "RU", // 橡胶
	"SA": "SA", // 纯碱
	"SF": "SF", // 硅铁
	"SH": "SH", // 烧碱
	"si": "si", // 工业硅
	"SM": "SM", // 锰硅
	"sn": "SN", // 沪锡
	"sp": "SP", // 纸浆
	"SR": "SR", // 白糖
	"ss": "SS", // 不锈钢
	"T":  "T",  // 十年国债
	"TA": "TA", // PTA
	"TF": "TF", // 五年国债
	"TL": "TL", // 三十年国债
	"TS": "TS", // 二年国债
	"UR": "UR", // 尿素
	"v":  "V",  // PVC
	"y":  "Y",  // 豆油
	"zn": "ZN", // 沪锌
}

// MarketExchangeSymbols 交易所品种映射
var MarketExchangeSymbols = map[string][]string{
	"cffex": {"IF", "IC", "IM", "IH", "T", "TF", "TS", "TL"},
	"dce": {
		"C", "CS", "A", "B", "M", "Y", "P", "FB", "BB", "JD",
		"L", "V", "PP", "J", "JM", "I", "EG", "RR", "EB", "PG", "LH", "LG", "BZ",
	},
	"czce": {
		"WH", "PM", "CF", "SR", "TA", "OI", "RI", "MA", "ME", "FG",
		"RS", "RM", "ZC", "JR", "LR", "SF", "SM", "WT", "TC", "GN",
		"RO", "ER", "SRX", "SRY", "WSX", "WSY", "CY", "AP", "UR",
		"CJ", "SA", "PK", "PF", "PX", "SH", "PR",
	},
	"shfe": {
		"CU", "AL", "ZN", "PB", "NI", "SN", "AU", "AG", "RB", "WR",
		"HC", "FU", "BU", "RU", "SC", "NR", "SP", "SS", "LU", "BC", "AO", "BR", "EC", "AD",
	},
	"gfex": {"SI", "LC", "PS"},
}

// ContractSymbols 所有合约品种列表
var ContractSymbols []string

func init() {
	for _, symbols := range MarketExchangeSymbols {
		ContractSymbols = append(ContractSymbols, symbols...)
	}
}

// DCEMap 大连商品交易所品种映射
var DCEMap = map[string]string{
	"大豆":      "A",
	"豆一":      "A",
	"豆二":      "B",
	"豆粕":      "M",
	"豆油":      "Y",
	"棕榈油":     "P",
	"玉米":      "C",
	"玉米淀粉":    "CS",
	"鸡蛋":      "JD",
	"纤维板":     "FB",
	"胶合板":     "BB",
	"聚乙烯":     "L",
	"聚氯乙烯":    "V",
	"聚丙烯":     "PP",
	"焦炭":      "J",
	"焦煤":      "JM",
	"铁矿石":     "I",
	"乙二醇":     "EG",
	"粳米":      "RR",
	"苯乙烯":     "EB",
	"液化石油气":   "PG",
	"生猪":      "LH",
	"原木":      "LG",
	"纯苯":      "BZ",
	"聚氯乙烯月均价": "VF",
	"聚丙烯月均价":  "PPF",
	"聚乙烯月均价":  "LF",
}

// HQSinaSpotHeaders 新浪期货行情请求头
var HQSinaSpotHeaders = map[string]string{
	"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8",
	"Accept-Encoding": "gzip, deflate",
	"Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8",
	"Cache-Control":   "no-cache",
	"Connection":      "keep-alive",
	"Host":            "finance.sina.com.cn",
	"Pragma":          "no-cache",
	"Referer":         "https://finance.sina.com.cn/futuremarket/",
	"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.97 Safari/537.36",
}

// ZhSinaSpotHeaders 新浪国内期货行情请求头
var ZhSinaSpotHeaders = map[string]string{
	"Accept":          "*/*",
	"Accept-Encoding": "gzip, deflate, br",
	"Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8",
	"Cache-Control":   "no-cache",
	"Connection":      "keep-alive",
	"Host":            "hq.sinajs.cn",
	"Pragma":          "no-cache",
	"Referer":         "https://finance.sina.com.cn/futuremarket/",
	"Sec-Fetch-Mode":  "no-cors",
	"Sec-Fetch-Site":  "cross-site",
	"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.97 Safari/537.36",
}

// URL常量
const (
	ZhSubscribeExchangeSymbolURL = "http://vip.stock.finance.sina.com.cn/quotes_service/view/js/qihuohangqing.js"
	ZhMatchMainContractURL       = "http://vip.stock.finance.sina.com.cn/quotes_service/api/json_v2.php/Market_Center.getHQFuturesData"
)
