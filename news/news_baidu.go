package news

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// getBaiduCookie 获取百度股市通所需的 Cookie
func getBaiduCookie(headers map[string]string) (string, error) {
	// 第一步：获取基础Cookie
	resp1, err := utils.GetWithHeaders("https://gushitong.baidu.com/calendar", nil, headers)
	if err != nil {
		return "", fmt.Errorf("获取百度Cookie失败: %w", err)
	}

	// 从响应中提取 hm.js URL
	text := resp1.String()
	hmRe := regexp.MustCompile(`https://hm\.baidu\.com/hm\.js\?\w+`)
	hmMatch := hmRe.FindString(text)
	if hmMatch == "" {
		// 尝试备用模式
		hmRe2 := regexp.MustCompile(`//hm\.baidu\.com/hm\.js\?\w+`)
		hmMatch = hmRe2.FindString(text)
		if hmMatch != "" {
			hmMatch = "https:" + hmMatch
		}
	}

	if hmMatch == "" {
		return "", fmt.Errorf("无法提取hm.js URL")
	}

	// 第二步：请求 hm.js
	_, err = utils.GetWithHeaders(hmMatch, nil, headers)
	if err != nil {
		return "", fmt.Errorf("请求hm.js失败: %w", err)
	}

	// 注意：Go 的 resty 库不会自动管理 Cookie
	// 这里简化处理，返回空字符串，让调用方自行传入 cookie
	return "", nil
}

// baiduFinanceCalendar 百度股市通日历数据基础函数
func baiduFinanceCalendar(date, cate, cookie string) (gjson.Result, error) {
	// 日期格式转换: YYYYMMDD -> YYYY-MM-DD
	if len(date) != 8 {
		return gjson.Result{}, fmt.Errorf("日期格式错误，应为 YYYYMMDD")
	}
	formattedDate := date[:4] + "-" + date[4:6] + "-" + date[6:]

	headers := map[string]string{
		"Accept":          "application/vnd.finance-web.v1+json",
		"Accept-Encoding": "gzip, deflate, br, zstd",
		"Accept-Language": "en,zh-CN;q=0.9,zh;q=0.8",
		"Cache-Control":   "no-cache",
		"Origin":          "https://gushitong.baidu.com",
		"Pragma":          "no-cache",
		"Referer":         "https://gushitong.baidu.com/",
		"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36",
	}

	if cookie != "" {
		headers["Cookie"] = cookie
	}

	url := "https://finance.pae.baidu.com/sapi/v1/financecalendar"

	// 第一次请求获取总记录数
	params := map[string]string{
		"start_date":    formattedDate,
		"end_date":      formattedDate,
		"pn":            "0",
		"rn":            "100",
		"cate":          cate,
		"finClientType": "pc",
	}

	resp, err := utils.GetWithHeaders(url, params, headers)
	if err != nil {
		return gjson.Result{}, fmt.Errorf("请求失败: %w", err)
	}

	dataJson := gjson.Parse(resp.String())

	// 获取指定日期的总记录数
	var totalRecords int64
	calendarInfo := dataJson.Get("Result.calendarInfo")
	calendarInfo.ForEach(func(key, value gjson.Result) bool {
		if value.Get("date").String() == formattedDate {
			totalRecords = value.Get("total").Int()
			return false
		}
		return true
	})

	// 计算总页数
	totalPages := int(math.Ceil(float64(totalRecords) / 100))
	if totalPages == 0 {
		totalPages = 1
	}

	// 收集所有数据
	var allItems []gjson.Result

	// 处理第一页数据
	calendarInfo.ForEach(func(key, value gjson.Result) bool {
		if value.Get("date").String() == formattedDate {
			value.Get("list").ForEach(func(k, v gjson.Result) bool {
				allItems = append(allItems, v)
				return true
			})
			return false
		}
		return true
	})

	// 获取剩余页
	for page := 1; page < totalPages; page++ {
		params["pn"] = strconv.Itoa(page)
		resp, err := utils.GetWithHeaders(url, params, headers)
		if err != nil {
			continue
		}

		dataJson := gjson.Parse(resp.String())
		dataJson.Get("Result.calendarInfo").ForEach(func(key, value gjson.Result) bool {
			if value.Get("date").String() == formattedDate {
				value.Get("list").ForEach(func(k, v gjson.Result) bool {
					allItems = append(allItems, v)
					return true
				})
				return false
			}
			return true
		})
	}

	// 将所有项目合并为一个 JSON 数组
	return gjson.Parse(itemsToJsonArray(allItems)), nil
}

// itemsToJsonArray 将 gjson.Result 数组转为 JSON 数组字符串
func itemsToJsonArray(items []gjson.Result) string {
	if len(items) == 0 {
		return "[]"
	}
	result := "["
	for i, item := range items {
		if i > 0 {
			result += ","
		}
		result += item.Raw
	}
	result += "]"
	return result
}

// NewsEconomicBaidu 百度股市通-经济数据
// https://gushitong.baidu.com/calendar
// date: 查询日期，格式 YYYYMMDD
// cookie: 可选的 cookie 字符串
func NewsEconomicBaidu(date string, cookie string) ([]EconomicData, error) {
	result, err := baiduFinanceCalendar(date, "economic_data", cookie)
	if err != nil {
		return nil, err
	}

	var dataList []EconomicData
	result.ForEach(func(key, value gjson.Result) bool {
		dataList = append(dataList, EconomicData{
			Date:       value.Get("date").String(),
			Time:       value.Get("time").String(),
			Country:    value.Get("country").String(),
			Region:     value.Get("region").String(),
			Event:      value.Get("title").String(),
			Period:     value.Get("timePeriod").String(),
			Actual:     value.Get("pubVal").Float(),
			Forecast:   value.Get("indicateVal").Float(),
			Previous:   value.Get("formerVal").Float(),
			Importance: int(value.Get("star").Int()),
		})
		return true
	})

	return dataList, nil
}

// NewsTradeNotifySuspendBaidu 百度股市通-交易提醒-停复牌
// https://gushitong.baidu.com/calendar
// date: 查询日期，格式 YYYYMMDD
// cookie: 可选的 cookie 字符串
func NewsTradeNotifySuspendBaidu(date string, cookie string) ([]SuspendNotify, error) {
	result, err := baiduFinanceCalendar(date, "notify_suspend", cookie)
	if err != nil {
		return nil, err
	}

	var dataList []SuspendNotify
	result.ForEach(func(key, value gjson.Result) bool {
		resumeTime := value.Get("end").String()
		if resumeTime == "" {
			resumeTime = "-"
		}

		dataList = append(dataList, SuspendNotify{
			Code:         value.Get("code").String(),
			Name:         value.Get("name").String(),
			Exchange:     value.Get("exchange").String(),
			SuspendTime:  value.Get("start").String(),
			ResumeTime:   resumeTime,
			Reason:       value.Get("reason").String(),
			MarketValue:  value.Get("marketValue").String(),
			AnnounceDate: value.Get("date").String(),
			AnnounceTime: value.Get("time").String(),
			SecType:      value.Get("type").String(),
			MarketType:   value.Get("market").String(),
			IsSkip:       value.Get("isSkip").String(),
		})
		return true
	})

	return dataList, nil
}

// NewsTradeNotifyDividendBaidu 百度股市通-交易提醒-分红派息
// https://gushitong.baidu.com/calendar
// date: 查询日期，格式 YYYYMMDD
// cookie: 可选的 cookie 字符串
func NewsTradeNotifyDividendBaidu(date string, cookie string) ([]DividendNotify, error) {
	result, err := baiduFinanceCalendar(date, "notify_divide", cookie)
	if err != nil {
		return nil, err
	}

	var dataList []DividendNotify
	result.ForEach(func(key, value gjson.Result) bool {
		dividend := value.Get("diviCash").String()
		if dividend == "" {
			dividend = "-"
		}
		bonus := value.Get("shareDivide").String()
		if bonus == "" {
			bonus = "-"
		}
		transfer := value.Get("transfer").String()
		if transfer == "" {
			transfer = "-"
		}
		physical := value.Get("physical").String()
		if physical == "" {
			physical = "-"
		}

		dataList = append(dataList, DividendNotify{
			Code:       value.Get("code").String(),
			ExDate:     value.Get("diviDate").String(),
			Dividend:   dividend,
			Bonus:      bonus,
			Transfer:   transfer,
			Physical:   physical,
			Exchange:   value.Get("exchange").String(),
			Name:       value.Get("name").String(),
			ReportDate: value.Get("date").String(),
		})
		return true
	})

	return dataList, nil
}

// NewsReportTimeBaidu 百度股市通-财报发行
// https://gushitong.baidu.com/calendar
// date: 查询日期，格式 YYYYMMDD
// cookie: 可选的 cookie 字符串
func NewsReportTimeBaidu(date string, cookie string) ([]ReportTime, error) {
	result, err := baiduFinanceCalendar(date, "report_time", cookie)
	if err != nil {
		return nil, err
	}

	var dataList []ReportTime
	result.ForEach(func(key, value gjson.Result) bool {
		marketValue := value.Get("marketValue").Float()
		if marketValue == 0 {
			marketValue = value.Get("capitalization").Float()
		}

		dataList = append(dataList, ReportTime{
			Code:        value.Get("code").String(),
			Name:        value.Get("name").String(),
			Exchange:    value.Get("exchange").String(),
			ReportType:  value.Get("reportType").String(),
			PublishTime: value.Get("time").String(),
			MarketValue: marketValue,
			PublishDate: value.Get("date").String(),
		})
		return true
	})

	return dataList, nil
}
