package utils

import (
	"strings"
	"time"
)

// 中国时区
var ChinaLoc = time.FixedZone("Asia/Shanghai", 8*3600)

// 支持的日期格式
var dateFormats = []string{
	"2006-01-02",
	"20060102",
	"2006/01/02",
	"2006.01.02",
	"2006年01月02日",
	"2006-1-2",
	"2006/1/2",
}

// ParseDate 解析日期字符串为 time.Time
func ParseDate(s string) (time.Time, error) {
	if s == "" {
		return time.Time{}, nil
	}

	s = strings.TrimSpace(s)

	// 尝试各种格式
	for _, format := range dateFormats {
		if t, err := time.ParseInLocation(format, s, ChinaLoc); err == nil {
			return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, ChinaLoc), nil
		}
	}

	return time.Time{}, nil
}

// FormatDate 格式化日期为指定格式字符串
func FormatDate(t time.Time, format string) string {
	if format == "" {
		format = "2006-01-02"
	}
	return t.In(ChinaLoc).Format(format)
}

// Today 获取今天的日期（零点时间）
func Today() time.Time {
	now := time.Now().In(ChinaLoc)
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, ChinaLoc)
}
