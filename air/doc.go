// Package air 提供空气质量和日出日落时间数据接口
//
// 主要功能：
//   - 河北省空气质量数据
//   - 真气网城市空气质量数据
//   - 日出日落时间查询
//
// 数据源：
//   - 河北省环境监测: http://110.249.223.67/
//   - 真气网: https://www.zq12369.com/
//   - Time and Date: https://www.timeanddate.com/
//
// 示例：
//
//	// 获取河北省空气质量
//	df, err := air.AirQualityHebei()
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	// 获取城市空气质量列表
//	df, err = air.AirCityTable()
//
//	// 获取日出日落时间
//	df, err = air.SunriseDaily("Beijing", "2024-01-01")
package air
