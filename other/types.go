package other

// CarMarketTotalItem 乘联会-总体市场数据项
type CarMarketTotalItem struct {
	Month        string  `json:"month"`         // 月份
	CurrentYear  float64 `json:"current_year"`  // 当年数值
	PreviousYear float64 `json:"previous_year"` // 去年数值
}

// CarMarketManRankItem 乘联会-厂商排名数据项
type CarMarketManRankItem struct {
	Manufacturer string  `json:"manufacturer"`  // 厂商
	CurrentYear  float64 `json:"current_year"`  // 当年数值
	PreviousYear float64 `json:"previous_year"` // 去年数值
}

// CarMarketCateItem 乘联会-车型大类数据项
type CarMarketCateItem struct {
	Month        string  `json:"month"`         // 月份
	CurrentYear  float64 `json:"current_year"`  // 当年数值
	PreviousYear float64 `json:"previous_year"` // 去年数值
}

// CarMarketCountryRow 乘联会-国别细分市场数据行
type CarMarketCountryRow struct {
	Country string  `json:"country"` // 国别
	Value1  float64 `json:"value1"`  // 数值1
	Value2  float64 `json:"value2"`  // 数值2
	Value3  float64 `json:"value3"`  // 数值3
	Value4  float64 `json:"value4"`  // 数值4
}

// CarMarketCountryItem 乘联会-国别细分市场数据项
type CarMarketCountryItem struct {
	Month  string                `json:"month"`  // 月份
	Values []CarMarketCountryRow `json:"values"` // 各国别数据
}

// CarMarketSegmentRow 乘联会-级别细分市场数据行
type CarMarketSegmentRow struct {
	Segment string  `json:"segment"` // 级别
	Value1  float64 `json:"value1"`  // 数值1
	Value2  float64 `json:"value2"`  // 数值2
	Value3  float64 `json:"value3"`  // 数值3
}

// CarMarketSegmentItem 乘联会-级别细分市场数据项
type CarMarketSegmentItem struct {
	Month  string                `json:"month"`  // 月份
	Values []CarMarketSegmentRow `json:"values"` // 各级别数据
}

// CarMarketFuelItem 乘联会-新能源细分市场数据项
type CarMarketFuelItem struct {
	Month        string  `json:"month"`         // 月份
	CurrentYear  float64 `json:"current_year"`  // 当年数值
	PreviousYear float64 `json:"previous_year"` // 去年数值
}

// CarSaleRankGasgooItem 盖世汽车-销量排名数据项
type CarSaleRankGasgooItem struct {
	Rank         int     `json:"rank"`           // 排名
	Name         string  `json:"name"`           // 名称
	Sales        float64 `json:"sales"`          // 销量
	YearOnYear   float64 `json:"year_on_year"`   // 同比
	MonthOnMonth float64 `json:"month_on_month"` // 环比
}
