package futures

import (
	"regexp"
	"strings"
	"unicode"
)

// SymbolVarieties 从合约代码中提取品种代码
// 例如: "ru1801" -> "RU"
func SymbolVarieties(contractCode string) string {
	re := regexp.MustCompile(`\D+`)
	matches := re.FindAllString(contractCode, -1)
	result := strings.ToUpper(strings.TrimSpace(strings.Join(matches, "")))
	if result == "PTA" {
		result = "TA"
	}
	return result
}

// SymbolMarket 根据品种代码获取交易所代码
func SymbolMarket(symbolDetail string) string {
	varItem := SymbolVarieties(symbolDetail)
	for market, contracts := range MarketExchangeSymbols {
		for _, contract := range contracts {
			if varItem == contract {
				return market
			}
		}
	}
	return ""
}

// FindChinese 提取字符串中的中文字符
func FindChinese(s string) string {
	var result strings.Builder
	for _, r := range s {
		if unicode.Is(unicode.Han, r) {
			result.WriteRune(r)
		}
	}
	return result.String()
}

// ChineseToEnglish 将期货品种中文名称映射为英文缩写
func ChineseToEnglish(chineseVar string) string {
	chineseList := []string{
		"橡胶", "天然橡胶", "石油沥青", "沥青", "沥青仓库", "沥青(仓库)", "沥青厂库", "沥青(厂库)",
		"热轧卷板", "热轧卷板厂库", "热轧卷板仓库", "热轧板卷", "燃料油", "白银", "线材", "螺纹钢",
		"铅", "铜", "铝", "锌", "黄金", "钯金", "锡", "镍", "纸浆", "豆一", "大豆", "豆二",
		"胶合板", "玉米", "玉米淀粉", "聚乙烯", "LLDPE", "LDPE", "豆粕", "豆油", "大豆油",
		"棕榈油", "纤维板", "鸡蛋", "聚氯乙烯", "PVC", "聚丙烯", "PP", "焦炭", "焦煤",
		"铁矿石", "乙二醇", "强麦", "强筋小麦", " 强筋小麦", "硬冬白麦", "普麦", "硬白小麦",
		"硬白小麦（）", "皮棉", "棉花", "一号棉", "白糖", "PTA", "菜籽油", "菜油", "早籼稻",
		"早籼", "甲醇", "柴油", "玻璃", "油菜籽", "菜籽", "菜籽粕", "菜粕", "动力煤",
		"粳稻", "晚籼稻", "晚籼", "硅铁", "锰硅", "硬麦", "棉纱", "苹果", "原油",
		"中质含硫原油", "尿素", "20号胶", "苯乙烯", "不锈钢", "粳米", "20号胶20", "红枣",
		"不锈钢仓库", "不锈钢厂库", "纯碱", "液化石油气", "低硫燃料油", "纸浆仓库",
		"石油沥青厂库", "石油沥青仓库", "螺纹钢仓库", "螺纹钢厂库", "纸浆厂库",
		"低硫燃料油仓库", "低硫燃料油厂库", "短纤", "涤纶短纤", "生猪", "花生",
		"工业硅", "氧化铝", "丁二烯橡胶", "碳酸锂", "氧化铝仓库", "氧化铝厂库",
		"烧碱", "丁二烯橡胶仓库", "丁二烯橡胶厂库", "PX", "原木", "瓶片", "纯苯",
		"多晶硅", "铸造铝合金",
	}

	englishList := []string{
		"RU", "RU", "BU", "BU", "BU", "BU", "BU2", "BU2",
		"HC", "HC", "HC", "HC", "FU", "AG", "WR", "RB",
		"PB", "CU", "AL", "ZN", "AU", "AU", "SN", "NI", "SP", "A", "A", "B",
		"BB", "C", "CS", "L", "L", "L", "M", "Y", "Y",
		"P", "FB", "JD", "V", "V", "PP", "PP", "J", "JM",
		"I", "EG", "WH", "WH", "WH", "PM", "PM", "PM",
		"PM", "CF", "CF", "CF", "SR", "TA", "OI", "OI", "RI",
		"ER", "MA", "MA", "FG", "RS", "RS", "RM", "RM", "ZC",
		"JR", "LR", "LR", "SF", "SM", "WT", "CY", "AP", "SC",
		"SC", "UR", "NR", "EB", "SS", "RR", "NR", "CJ",
		"SS", "SS", "SA", "PG", "LU", "SP",
		"BU", "BU", "RB", "RB", "SP",
		"LU", "LU", "PF", "PF", "LH", "PK",
		"SI", "AO", "BR", "LC", "AO", "AO",
		"SH", "BR", "BR", "PX", "LG", "PR", "BZ",
		"PS", "AD",
	}

	for i, ch := range chineseList {
		if ch == chineseVar {
			return englishList[i]
		}
	}
	return ""
}
