package article

// VolatilityData 波动率数据结构
type VolatilityData struct {
	Date  string  `json:"date"`  // 日期
	Value float64 `json:"value"` // 波动率值
}

// EPUData 经济政策不确定性指数数据
type EPUData struct {
	Year  int     `json:"year"`  // 年份
	Month int     `json:"month"` // 月份
	Index float64 `json:"index"` // EPU指数值
}

// FREDData FRED宏观经济数据
type FREDData struct {
	Date   string             `json:"date"`   // 日期
	Values map[string]float64 `json:"values"` // 指标名称和对应的值
}
