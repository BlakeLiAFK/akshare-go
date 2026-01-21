package option

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"time"
)

// 期权相关URL常量
const (
	// 中国金融期货交易所
	CFFEXOptionURL300 = "http://www.cffex.com.cn/quote_IO.txt"

	// 深圳证券交易所
	SZOptionURL300 = "http://www.szse.cn/api/report/ShowReport?SHOWTYPE=xlsx&CATALOGID=ysplbrb&TABKEY=tab1&random=0.10432465776720479"

	// 上海证券交易所 - 50ETF期权
	SHOptionURL50     = "http://yunhq.sse.com.cn:32041/v1/sh1/list/self/510050"
	SHOptionURLKing50 = "http://yunhq.sse.com.cn:32041/v1/sho/list/tstyle/510050_%s"

	// 上海证券交易所 - 300ETF期权
	SHOptionURL300     = "http://yunhq.sse.com.cn:32041/v1/sh1/list/self/510300"
	SHOptionURLKing300 = "http://yunhq.sse.com.cn:32041/v1/sho/list/tstyle/510300_%s"

	// 上海证券交易所 - 500ETF期权
	SHOptionURL500     = "http://yunhq.sse.com.cn:32041/v1/sh1/list/self/510500"
	SHOptionURLKing500 = "http://yunhq.sse.com.cn:32041/v1/sho/list/tstyle/510500_%s"

	// 上海证券交易所 - 科创50期权
	SHOptionURLKC50     = "http://yunhq.sse.com.cn:32041/v1/sh1/list/self/588000"
	SHOptionURLKingKC50 = "http://yunhq.sse.com.cn:32041/v1/sho/list/tstyle/588000_%s"

	// 上海证券交易所 - 科创50期权(期权)
	SHOptionURLKC50YFD   = "http://yunhq.sse.com.cn:32041/v1/sh1/list/self/588080"
	SHOptionURLKing50YFD = "http://yunhq.sse.com.cn:32041/v1/sho/list/tstyle/588080_%s"

	// 大连商品交易所
	DCEOptionURL      = "http://portal.dce.com.cn/publicweb/quotesdata/dayQuotesCh.html"
	DCEDailyOptionURL = "http://portal.dce.com.cn/publicweb/quotesdata/exportDayQuotesChData.html"

	// 上海期货交易所
	SHFEOptionURL = "https://tsite.shfe.com.cn/data/dailydata/option/kx/kx%s.dat"

	// 郑州商品交易所
	CZCEDailyOptionURL3 = "http://www.czce.com.cn/cn/DFSStaticFiles/Option/%s/%s/OptionDataDaily.txt"
)

// 上海期权请求参数
var SHOptionPayload = map[string]string{
	"select": "select: code,name,last,change,chg_rate,amp_rate,volume,amount,prev_close",
}

var SHOptionPayloadOther = map[string]string{
	"select": "contractid,last,chg_rate,presetpx,exepx",
}

// 上期所请求头
var SHFEHeaders = map[string]string{
	"User-Agent": "Mozilla/4.0 (compatible; MSIE 5.5; Windows NT)",
}

// 日期正则表达式
var datePattern = regexp.MustCompile(`^([0-9]{4})[-/]?([0-9]{2})[-/]?([0-9]{2})`)

// ConvertDate 转换日期字符串为time.Time对象
//
// 参数:
//   - dateStr: 日期字符串，格式如 "2016-01-01", "20160101", "2016/01/01"
//
// 返回:
//   - time.Time: 解析后的日期对象
//   - error: 错误信息
func ConvertDate(dateStr string) (time.Time, error) {
	if dateStr == "" {
		return time.Time{}, fmt.Errorf("日期字符串为空")
	}

	matches := datePattern.FindStringSubmatch(dateStr)
	if len(matches) != 4 {
		return time.Time{}, fmt.Errorf("日期格式错误: %s", dateStr)
	}

	year, err1 := strconv.Atoi(matches[1])
	month, err2 := strconv.Atoi(matches[2])
	day, err3 := strconv.Atoi(matches[3])

	if err1 != nil || err2 != nil || err3 != nil {
		return time.Time{}, fmt.Errorf("日期解析失败: %s", dateStr)
	}

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC), nil
}

// GetJSONPath 获取JSON配置文件的路径
//
// 参数:
//   - name: 文件名
//   - moduleFile: 模块文件路径
//
// 返回:
//   - string: JSON文件路径
func GetJSONPath(name, moduleFile string) string {
	moduleFolder := filepath.Dir(filepath.Dir(moduleFile))
	return filepath.Join(moduleFolder, "file_fold", name)
}

// GetCalendar 获取交易日历
//
// 返回:
//   - []string: 交易日历列表
//   - error: 错误信息
func GetCalendar() ([]string, error) {
	settingFileName := "calendar.json"

	// 获取当前文件路径
	wd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("获取工作目录失败: %w", err)
	}

	settingFilePath := filepath.Join(wd, "file_fold", settingFileName)

	data, err := os.ReadFile(settingFilePath)
	if err != nil {
		return nil, fmt.Errorf("读取日历文件失败: %w", err)
	}

	var calendar []string
	if err := json.Unmarshal(data, &calendar); err != nil {
		return nil, fmt.Errorf("解析日历文件失败: %w", err)
	}

	return calendar, nil
}

// LastTradingDay 获取前一个交易日
//
// 参数:
//   - day: 日期字符串或time.Time对象
//
// 返回:
//   - string: 前一个交易日字符串
//   - error: 错误信息
func LastTradingDay(day interface{}) (string, error) {
	calendar, err := GetCalendar()
	if err != nil {
		return "", fmt.Errorf("获取交易日历失败: %w", err)
	}

	var dayStr string
	switch v := day.(type) {
	case string:
		dayStr = v
	case time.Time:
		dayStr = v.Format("20060102")
	default:
		return "", fmt.Errorf("不支持的日期类型")
	}

	// 查找当前交易日在日历中的位置
	pos := -1
	for i, calDay := range calendar {
		if calDay == dayStr {
			pos = i
			break
		}
	}

	if pos == -1 {
		return "", fmt.Errorf("今天不是交易日: %s", dayStr)
	}

	if pos == 0 {
		return "", fmt.Errorf("没有找到前一个交易日")
	}

	return calendar[pos-1], nil
}

// GetLatestDataDate 获取最新的有数据的交易日
//
// 参数:
//   - day: 当前时间
//
// 返回:
//   - string: 最新交易日字符串
//   - error: 错误信息
func GetLatestDataDate(day time.Time) (string, error) {
	calendar, err := GetCalendar()
	if err != nil {
		return "", fmt.Errorf("获取交易日历失败: %w", err)
	}

	dayStr := day.Format("20060102")

	// 检查今天是否是交易日
	for _, calDay := range calendar {
		if calDay == dayStr {
			// 如果是交易日且时间超过17点，返回今天
			if day.Hour() >= 17 {
				return dayStr, nil
			} else {
				// 否则返回前一个交易日
				return LastTradingDay(dayStr)
			}
		}
	}

	// 如果今天不是交易日，向前查找最近的交易日
	for i := len(calendar) - 1; i >= 0; i-- {
		calDay, _ := time.Parse("20060102", calendar[i])
		if calDay.Before(day) || calDay.Equal(day) {
			return calendar[i], nil
		}
	}

	return "", fmt.Errorf("未找到有效的交易日")
}
