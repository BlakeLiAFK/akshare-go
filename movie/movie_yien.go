package movie

import (
	"fmt"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

const (
	// 艺恩API URL
	yienAPIURL = "https://www.endata.com.cn/API/GetData.ashx"
)

// MovieBoxofficeRealtimeItem 实时票房数据结构
type MovieBoxofficeRealtimeItem struct {
	Rank        int    `json:"rank"`         // 排序
	MovieName   string `json:"movie_name"`   // 影片名称
	RealtimeBox string `json:"realtime_box"` // 实时票房
	BoxPct      string `json:"box_pct"`      // 票房占比
	ShowDays    int    `json:"show_days"`    // 上映天数
	TotalBox    string `json:"total_box"`    // 累计票房
}

// MovieBoxofficeDailyItem 单日票房数据结构
type MovieBoxofficeDailyItem struct {
	Rank       int     `json:"rank"`        // 排序
	MovieName  string  `json:"movie_name"`  // 影片名称
	DailyBox   string  `json:"daily_box"`   // 单日票房
	DayOnDay   string  `json:"day_on_day"`  // 环比变化
	TotalBox   string  `json:"total_box"`   // 累计票房
	AvgPrice   float64 `json:"avg_price"`   // 平均票价
	AvgPerson  float64 `json:"avg_person"`  // 场均人次
	ScoreIndex float64 `json:"score_index"` // 口碑指数
	ShowDays   int     `json:"show_days"`   // 上映天数
}

// MovieBoxofficeWeeklyItem 单周票房数据结构
type MovieBoxofficeWeeklyItem struct {
	Rank       int     `json:"rank"`         // 排序
	MovieName  string  `json:"movie_name"`   // 影片名称
	RankChange string  `json:"rank_change"`  // 排名变化
	WeeklyBox  float64 `json:"weekly_box"`   // 单周票房
	WeekOnWeek float64 `json:"week_on_week"` // 环比变化
	TotalBox   float64 `json:"total_box"`    // 累计票房
	AvgPrice   float64 `json:"avg_price"`    // 平均票价
	AvgPerson  float64 `json:"avg_person"`   // 场均人次
	ScoreIndex float64 `json:"score_index"`  // 口碑指数
	ShowDays   int     `json:"show_days"`    // 上映天数
}

// MovieBoxofficeMonthlyItem 单月票房数据结构
type MovieBoxofficeMonthlyItem struct {
	Rank        int     `json:"rank"`         // 排序
	MovieName   string  `json:"movie_name"`   // 影片名称
	MonthlyBox  string  `json:"monthly_box"`  // 单月票房
	MonthPct    string  `json:"month_pct"`    // 月度占比
	AvgPrice    float64 `json:"avg_price"`    // 平均票价
	AvgPerson   float64 `json:"avg_person"`   // 场均人次
	ReleaseDate string  `json:"release_date"` // 上映日期
	ScoreIndex  float64 `json:"score_index"`  // 口碑指数
	MonthDays   int     `json:"month_days"`   // 月内天数
}

// MovieBoxofficeYearlyItem 年度票房数据结构
type MovieBoxofficeYearlyItem struct {
	Rank        int     `json:"rank"`         // 排序
	MovieName   string  `json:"movie_name"`   // 影片名称
	Genre       string  `json:"genre"`        // 类型
	TotalBox    string  `json:"total_box"`    // 总票房
	AvgPrice    float64 `json:"avg_price"`    // 平均票价
	AvgPerson   float64 `json:"avg_person"`   // 场均人次
	Country     string  `json:"country"`      // 国家及地区
	ReleaseDate string  `json:"release_date"` // 上映日期
}

// MovieBoxofficeYearlyFirstWeekItem 年度首周票房数据结构
type MovieBoxofficeYearlyFirstWeekItem struct {
	Rank          int     `json:"rank"`            // 排序
	MovieName     string  `json:"movie_name"`      // 影片名称
	Genre         string  `json:"genre"`           // 类型
	FirstWeekBox  string  `json:"first_week_box"`  // 首周票房
	BoxPct        string  `json:"box_pct"`         // 占总票房比重
	AvgPerson     float64 `json:"avg_person"`      // 场均人次
	Country       string  `json:"country"`         // 国家及地区
	ReleaseDate   string  `json:"release_date"`    // 上映日期
	FirstWeekDays int     `json:"first_week_days"` // 首周天数
}

// MovieBoxofficeCinemaDailyItem 影院日票房数据结构
type MovieBoxofficeCinemaDailyItem struct {
	Rank       int     `json:"rank"`        // 排序
	CinemaName string  `json:"cinema_name"` // 影院名称
	DailyBox   string  `json:"daily_box"`   // 单日票房
	DailySess  int     `json:"daily_sess"`  // 单日场次
	AvgPerson  float64 `json:"avg_person"`  // 场均人次
	AvgPrice   float64 `json:"avg_price"`   // 场均票价
	Occupancy  string  `json:"occupancy"`   // 上座率
}

// MovieBoxofficeCinemaWeeklyItem 影院周票房数据结构
type MovieBoxofficeCinemaWeeklyItem struct {
	Rank             int     `json:"rank"`                // 排序
	CinemaName       string  `json:"cinema_name"`         // 影院名称
	WeeklyBox        string  `json:"weekly_box"`          // 当周票房
	PerScreenBox     string  `json:"per_screen_box"`      // 单银幕票房
	AvgPerson        float64 `json:"avg_person"`          // 场均人次
	DailyPerHallBox  string  `json:"daily_per_hall_box"`  // 单日单厅票房
	DailyPerHallSess string  `json:"daily_per_hall_sess"` // 单日单厅场次
}

// MovieBoxofficeRealtime 获取电影票房-实时票房
//
// 目标地址: https://ys.endata.cn/BoxOffice/Movie
//
// 返回:
//   - []MovieBoxofficeRealtimeItem: 实时票房数据列表
//   - error: 错误信息
//
// 注意: 此接口需要解密，当前返回示例数据
func MovieBoxofficeRealtime() ([]MovieBoxofficeRealtimeItem, error) {
	today := time.Now().Format("2006-01-02")

	params := map[string]string{
		"showDate":   "",
		"tdate":      today,
		"MethodName": "BoxOffice_GetHourBoxOffice",
	}

	headers := map[string]string{
		"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := utils.PostFormWithHeaders(yienAPIURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("获取实时票房失败: %w", err)
	}

	// 艺恩数据需要JS解密，这里尝试直接解析
	text := resp.String()
	result := gjson.Parse(text)
	data := result.Get("Data.Table1")

	var items []MovieBoxofficeRealtimeItem
	if data.Exists() && data.IsArray() {
		data.ForEach(func(_, value gjson.Result) bool {
			item := MovieBoxofficeRealtimeItem{
				Rank:        int(value.Get("Irank").Int()),
				MovieName:   value.Get("MovieName").String(),
				RealtimeBox: value.Get("BoxOffice").String(),
				TotalBox:    value.Get("SumBoxOffice").String(),
				ShowDays:    int(value.Get("Days").Int()),
				BoxPct:      value.Get("BoxPer").String(),
			}
			items = append(items, item)
			return true
		})
	}

	// 如果解析失败，返回提示
	if len(items) == 0 {
		return nil, fmt.Errorf("数据需要解密，请参考Python版本使用jm.js解密")
	}

	return items, nil
}

// MovieBoxofficeDaily 获取电影票房-单日票房
//
// 目标地址: https://www.endata.com.cn/BoxOffice/BO/Day/index.html
//
// 参数:
//   - date: 日期，格式 "20240219"
//
// 返回:
//   - []MovieBoxofficeDailyItem: 单日票房数据列表
//   - error: 错误信息
func MovieBoxofficeDaily(date string) ([]MovieBoxofficeDailyItem, error) {
	// 格式化日期
	sdate := fmt.Sprintf("%s-%s-%s", date[:4], date[4:6], date[6:])

	params := map[string]string{
		"sdate":      sdate,
		"edate":      sdate,
		"MethodName": "BoxOffice_GetDayBoxOffice",
	}

	headers := map[string]string{
		"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := utils.PostFormWithHeaders(yienAPIURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("获取单日票房失败: %w", err)
	}

	text := resp.String()
	result := gjson.Parse(text)
	data := result.Get("Data.Table")

	var items []MovieBoxofficeDailyItem
	if data.Exists() && data.IsArray() {
		data.ForEach(func(_, value gjson.Result) bool {
			item := MovieBoxofficeDailyItem{
				Rank:       int(value.Get("Irank").Int()),
				MovieName:  value.Get("MovieName").String(),
				DailyBox:   value.Get("BoxOffice").String(),
				DayOnDay:   value.Get("BoxOffice_Up").String(),
				TotalBox:   value.Get("SumBoxOffice").String(),
				AvgPrice:   value.Get("AvgPrice").Float(),
				AvgPerson:  value.Get("AvgPeoPle").Float(),
				ScoreIndex: value.Get("WomIndex").Float(),
				ShowDays:   int(value.Get("Days").Int()),
			}
			items = append(items, item)
			return true
		})
	}

	if len(items) == 0 {
		return nil, fmt.Errorf("数据需要解密，请参考Python版本使用jm.js解密")
	}

	return items, nil
}

// MovieBoxofficeWeekly 获取电影票房-单周票房
//
// 目标地址: https://www.endata.com.cn/BoxOffice/BO/Week/oneWeek.html
//
// 参数:
//   - date: 日期，格式 "20240218"
//
// 返回:
//   - []MovieBoxofficeWeeklyItem: 单周票房数据列表
//   - error: 错误信息
func MovieBoxofficeWeekly(date string) ([]MovieBoxofficeWeeklyItem, error) {
	// 获取当周周一
	t, _ := time.Parse("20060102", date)
	for t.Weekday() != time.Monday {
		t = t.AddDate(0, 0, -1)
	}
	monday := t.Format("2006-01-02")

	params := map[string]string{
		"sdate":      monday,
		"MethodName": "BoxOffice_GetWeekInfoData",
	}

	headers := map[string]string{
		"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := utils.PostFormWithHeaders(yienAPIURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("获取单周票房失败: %w", err)
	}

	text := resp.String()
	result := gjson.Parse(text)
	data := result.Get("Data.Table")

	var items []MovieBoxofficeWeeklyItem
	if data.Exists() && data.IsArray() {
		data.ForEach(func(_, value gjson.Result) bool {
			item := MovieBoxofficeWeeklyItem{
				Rank:       int(value.Get("Irank").Int()),
				MovieName:  value.Get("MovieName").String(),
				RankChange: value.Get("RankChange").String(),
				WeeklyBox:  value.Get("BoxOffice").Float(),
				WeekOnWeek: value.Get("BoxOffice_Up").Float(),
				TotalBox:   value.Get("SumBoxOffice").Float(),
				AvgPrice:   value.Get("AvgPrice").Float(),
				AvgPerson:  value.Get("AvgPeoPle").Float(),
				ScoreIndex: value.Get("WomIndex").Float(),
				ShowDays:   int(value.Get("Days").Int()),
			}
			items = append(items, item)
			return true
		})
	}

	if len(items) == 0 {
		return nil, fmt.Errorf("数据需要解密，请参考Python版本使用jm.js解密")
	}

	return items, nil
}

// MovieBoxofficeMonthly 获取电影票房-单月票房
//
// 目标地址: https://www.endata.com.cn/BoxOffice/BO/Month/oneMonth.html
//
// 参数:
//   - date: 日期，格式 "20240218"
//
// 返回:
//   - []MovieBoxofficeMonthlyItem: 单月票房数据列表
//   - error: 错误信息
func MovieBoxofficeMonthly(date string) ([]MovieBoxofficeMonthlyItem, error) {
	startTime := fmt.Sprintf("%s-%s-01", date[:4], date[4:6])

	params := map[string]string{
		"startTime":  startTime,
		"MethodName": "BoxOffice_GetMonthBox",
	}

	headers := map[string]string{
		"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := utils.PostFormWithHeaders(yienAPIURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("获取单月票房失败: %w", err)
	}

	text := resp.String()
	result := gjson.Parse(text)
	data := result.Get("Data.Table")

	var items []MovieBoxofficeMonthlyItem
	if data.Exists() && data.IsArray() {
		data.ForEach(func(_, value gjson.Result) bool {
			item := MovieBoxofficeMonthlyItem{
				Rank:        int(value.Get("Irank").Int()),
				MovieName:   value.Get("MovieName").String(),
				MonthlyBox:  value.Get("BoxOffice").String(),
				MonthPct:    value.Get("BoxPer").String(),
				AvgPrice:    value.Get("AvgPrice").Float(),
				AvgPerson:   value.Get("AvgPeoPle").Float(),
				ReleaseDate: value.Get("ReleaseTime").String(),
				ScoreIndex:  value.Get("WomIndex").Float(),
				MonthDays:   int(value.Get("Days").Int()),
			}
			items = append(items, item)
			return true
		})
	}

	if len(items) == 0 {
		return nil, fmt.Errorf("数据需要解密，请参考Python版本使用jm.js解密")
	}

	return items, nil
}

// MovieBoxofficeYearly 获取电影票房-年度票房
//
// 目标地址: https://www.endata.com.cn/BoxOffice/BO/Year/index.html
//
// 参数:
//   - date: 日期，格式 "20240218"
//
// 返回:
//   - []MovieBoxofficeYearlyItem: 年度票房数据列表
//   - error: 错误信息
func MovieBoxofficeYearly(date string) ([]MovieBoxofficeYearlyItem, error) {
	year := date[:4]

	params := map[string]string{
		"year":       year,
		"MethodName": "BoxOffice_GetYearInfoData",
	}

	headers := map[string]string{
		"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := utils.PostFormWithHeaders(yienAPIURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("获取年度票房失败: %w", err)
	}

	text := resp.String()
	result := gjson.Parse(text)
	data := result.Get("Data.Table")

	var items []MovieBoxofficeYearlyItem
	if data.Exists() && data.IsArray() {
		rank := 1
		data.ForEach(func(_, value gjson.Result) bool {
			item := MovieBoxofficeYearlyItem{
				Rank:        rank,
				MovieName:   value.Get("MovieName").String(),
				Genre:       value.Get("MovieType").String(),
				TotalBox:    value.Get("BoxOffice").String(),
				AvgPrice:    value.Get("AvgPrice").Float(),
				AvgPerson:   value.Get("AvgPeoPle").Float(),
				Country:     value.Get("Area").String(),
				ReleaseDate: value.Get("ReleaseTime").String(),
			}
			items = append(items, item)
			rank++
			return true
		})
	}

	if len(items) == 0 {
		return nil, fmt.Errorf("数据需要解密，请参考Python版本使用jm.js解密")
	}

	return items, nil
}

// MovieBoxofficeYearlyFirstWeek 获取电影票房-年度票房-年度首周票房
//
// 目标地址: https://www.endata.com.cn/BoxOffice/BO/Year/firstWeek.html
//
// 参数:
//   - date: 日期，格式 "20201018"，获取该日期所在年度的年度首周票房数据
//
// 返回:
//   - []MovieBoxofficeYearlyFirstWeekItem: 年度首周票房数据列表
//   - error: 错误信息
func MovieBoxofficeYearlyFirstWeek(date string) ([]MovieBoxofficeYearlyFirstWeekItem, error) {
	year := date[:4]

	params := map[string]string{
		"year":       year,
		"MethodName": "BoxOffice_getYearInfo_fData",
	}

	headers := map[string]string{
		"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := utils.PostFormWithHeaders(yienAPIURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("获取年度首周票房失败: %w", err)
	}

	text := resp.String()
	result := gjson.Parse(text)
	data := result.Get("Data.Table")

	var items []MovieBoxofficeYearlyFirstWeekItem
	if data.Exists() && data.IsArray() {
		rank := 1
		data.ForEach(func(_, value gjson.Result) bool {
			item := MovieBoxofficeYearlyFirstWeekItem{
				Rank:          rank,
				MovieName:     value.Get("MovieName").String(),
				Genre:         value.Get("Type").String(),
				FirstWeekBox:  value.Get("BoxOffice").String(),
				BoxPct:        value.Get("BoxRate").String(),
				AvgPerson:     value.Get("AvgPeoPle").Float(),
				Country:       value.Get("Area").String(),
				ReleaseDate:   value.Get("ReleaseTime").String(),
				FirstWeekDays: int(value.Get("FirstWeekDays").Int()),
			}
			items = append(items, item)
			rank++
			return true
		})
	}

	if len(items) == 0 {
		return nil, fmt.Errorf("数据需要解密，请参考Python版本使用jm.js解密")
	}

	return items, nil
}

// MovieBoxofficeCinemaDaily 获取电影票房-影院票房-日票房排行
//
// 目标地址: https://www.endata.com.cn/BoxOffice/BO/Cinema/day.html
//
// 参数:
//   - date: 日期，格式 "20240219"
//
// 返回:
//   - []MovieBoxofficeCinemaDailyItem: 影院日票房数据列表
//   - error: 错误信息
func MovieBoxofficeCinemaDaily(date string) ([]MovieBoxofficeCinemaDailyItem, error) {
	params := map[string]string{
		"rowNum1":    "1",
		"rowNum2":    "100",
		"date":       date,
		"MethodName": "BoxOffice_GetCinemaDayBoxOffice",
	}

	headers := map[string]string{
		"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := utils.PostFormWithHeaders(yienAPIURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("获取影院日票房失败: %w", err)
	}

	text := resp.String()
	result := gjson.Parse(text)
	data := result.Get("Data.Table")

	var items []MovieBoxofficeCinemaDailyItem
	if data.Exists() && data.IsArray() {
		data.ForEach(func(_, value gjson.Result) bool {
			item := MovieBoxofficeCinemaDailyItem{
				Rank:       int(value.Get("Irank").Int()),
				CinemaName: value.Get("CinemaName").String(),
				DailyBox:   value.Get("BoxOffice").String(),
				DailySess:  int(value.Get("SessionCnt").Int()),
				AvgPerson:  value.Get("AvgPeoPle").Float(),
				AvgPrice:   value.Get("AvgPrice").Float(),
				Occupancy:  value.Get("Attendance").String(),
			}
			items = append(items, item)
			return true
		})
	}

	if len(items) == 0 {
		return nil, fmt.Errorf("数据需要解密，请参考Python版本使用jm.js解密")
	}

	return items, nil
}

// MovieBoxofficeCinemaWeekly 获取电影票房-影院票房-周票房排行
//
// 目标地址: https://www.endata.com.cn/BoxOffice/BO/Cinema/week.html
//
// 参数:
//   - date: 日期，格式 "20240219"
//
// 返回:
//   - []MovieBoxofficeCinemaWeeklyItem: 影院周票房数据列表
//   - error: 错误信息
func MovieBoxofficeCinemaWeekly(date string) ([]MovieBoxofficeCinemaWeeklyItem, error) {
	// 计算周数ID
	t, _ := time.Parse("20060102", date)
	_, week := t.ISOWeek()
	dateID := fmt.Sprintf("%d", week-1-41+1128)

	params := map[string]string{
		"dateID":     dateID,
		"rowNum1":    "1",
		"rowNum2":    "100",
		"MethodName": "BoxOffice_GetCinemaWeekBoxOffice",
	}

	headers := map[string]string{
		"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := utils.PostFormWithHeaders(yienAPIURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("获取影院周票房失败: %w", err)
	}

	text := resp.String()
	result := gjson.Parse(text)
	data := result.Get("Data.Table")

	var items []MovieBoxofficeCinemaWeeklyItem
	if data.Exists() && data.IsArray() {
		data.ForEach(func(_, value gjson.Result) bool {
			item := MovieBoxofficeCinemaWeeklyItem{
				Rank:             int(value.Get("Irank").Int()),
				CinemaName:       value.Get("CinemaName").String(),
				WeeklyBox:        value.Get("BoxOffice").String(),
				PerScreenBox:     value.Get("PerScreenBoxOffice").String(),
				AvgPerson:        value.Get("AvgPeoPle").Float(),
				DailyPerHallBox:  value.Get("DailyPerHallBox").String(),
				DailyPerHallSess: value.Get("DailyPerHallSession").String(),
			}
			items = append(items, item)
			return true
		})
	}

	if len(items) == 0 {
		return nil, fmt.Errorf("数据需要解密，请参考Python版本使用jm.js解密")
	}

	return items, nil
}
