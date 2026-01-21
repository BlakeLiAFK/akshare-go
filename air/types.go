package air

// HebeiAirQuality 河北省空气质量数据
type HebeiAirQuality struct {
	City      string  `json:"城市"`        // 城市
	Region    string  `json:"区域"`        // 区域
	Station   string  `json:"监测点"`       // 监测点
	DateTime  string  `json:"时间"`        // 时间
	AQI       float64 `json:"AQI"`       // 空气质量指数
	Level     string  `json:"空气质量等级"`    // 空气质量等级
	MaxPoll   string  `json:"首要污染物"`     // 首要污染物
	Longitude float64 `json:"经度"`        // 经度
	Latitude  float64 `json:"纬度"`        // 纬度
	
	// 污染物浓度
	SO2Value   float64 `json:"二氧化硫_浓度"`    // 二氧化硫浓度
	SO2IAQI    float64 `json:"二氧化硫_IAQI"`   // 二氧化硫IAQI
	COValue    float64 `json:"一氧化碳_浓度"`    // 一氧化碳浓度
	COIAQI     float64 `json:"一氧化碳_IAQI"`   // 一氧化碳IAQI
	NO2Value   float64 `json:"二氧化氮_浓度"`    // 二氧化氮浓度
	NO2IAQI    float64 `json:"二氧化氮_IAQI"`   // 二氧化氮IAQI
	O31HValue  float64 `json:"臭氧1小时_浓度"`   // 臭氧1小时浓度
	O31HIAQI   float64 `json:"臭氧1小时_IAQI"`  // 臭氧1小时IAQI
	O38HValue  float64 `json:"臭氧8小时_浓度"`   // 臭氧8小时浓度
	O38HIAQI   float64 `json:"臭氧8小时_IAQI"`  // 臭氧8小时IAQI
	PM25Value  float64 `json:"PM2.5_浓度"`    // PM2.5浓度
	PM25IAQI   float64 `json:"PM2.5_IAQI"`   // PM2.5 IAQI
	PM10Value  float64 `json:"PM10_浓度"`     // PM10浓度
	PM10IAQI   float64 `json:"PM10_IAQI"`    // PM10 IAQI
}

// CityAir 城市空气质量数据
type CityAir struct {
	No         int     `json:"序号"`        // 序号
	Province   string  `json:"省份"`        // 省份
	City       string  `json:"城市"`        // 城市
	AQI        float64 `json:"AQI"`       // 空气质量指数
	Quality    string  `json:"空气质量"`      // 空气质量等级
	PM25       float64 `json:"PM2.5浓度"`   // PM2.5浓度
	MainPoll   string  `json:"首要污染物"`     // 首要污染物
}

// WatchPointAir 监测点空气质量数据
type WatchPointAir struct {
	StationName string  `json:"stationName"` // 监测点名称
	Time        string  `json:"time"`        // 时间
	AQI         float64 `json:"AQI"`         // 空气质量指数
	PM25        float64 `json:"PM2.5"`       // PM2.5浓度
	PM10        float64 `json:"PM10"`        // PM10浓度
	SO2         float64 `json:"SO2"`         // 二氧化硫
	NO2         float64 `json:"NO2"`         // 二氧化氮
	CO          float64 `json:"CO"`          // 一氧化碳
	O3          float64 `json:"O3"`          // 臭氧
}

// AirHist 历史空气质量数据
type AirHist struct {
	Time     string  `json:"time"`     // 时间
	AQI      float64 `json:"AQI"`      // 空气质量指数
	Quality  string  `json:"quality"`  // 空气质量等级
	PM25     float64 `json:"PM2.5"`    // PM2.5浓度
	PM10     float64 `json:"PM10"`     // PM10浓度
	SO2      float64 `json:"SO2"`      // 二氧化硫
	NO2      float64 `json:"NO2"`      // 二氧化氮
	CO       float64 `json:"CO"`       // 一氧化碳
	O3       float64 `json:"O3"`       // 臭氧
}

// SunriseData 日出日落数据
type SunriseData struct {
	Date      string `json:"date"`      // 日期
	Sunrise   string `json:"sunrise"`   // 日出时间
	Sunset    string `json:"sunset"`    // 日落时间
	Daylength string `json:"daylength"` // 日照时长
}
