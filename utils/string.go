package utils

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// TrimSpace 去除字符串首尾空白（包括全角空格）
func TrimSpace(s string) string {
	return strings.TrimFunc(s, func(r rune) bool {
		return unicode.IsSpace(r) || r == '\u3000'
	})
}

// CleanString 清理字符串（去除所有空白字符和特殊字符）
func CleanString(s string) string {
	re := regexp.MustCompile(`[\s\u00A0\u3000]+`)
	return strings.TrimSpace(re.ReplaceAllString(s, " "))
}

// RemoveAllSpaces 移除所有空格
func RemoveAllSpaces(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, s)
}

// ParseFloat 解析浮点数，支持带逗号的数字
func ParseFloat(s string) (float64, error) {
	// 移除逗号和空格
	s = strings.ReplaceAll(s, ",", "")
	s = strings.ReplaceAll(s, " ", "")
	s = TrimSpace(s)

	// 处理百分号
	isPercent := strings.HasSuffix(s, "%")
	if isPercent {
		s = strings.TrimSuffix(s, "%")
	}

	// 处理中文单位
	multiplier := 1.0
	if strings.HasSuffix(s, "万") {
		s = strings.TrimSuffix(s, "万")
		multiplier = 10000
	} else if strings.HasSuffix(s, "亿") {
		s = strings.TrimSuffix(s, "亿")
		multiplier = 100000000
	}

	// 处理负号
	if s == "-" || s == "--" || s == "" {
		return 0, nil
	}

	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}

	val *= multiplier
	if isPercent {
		val /= 100
	}

	return val, nil
}

// ParseInt 解析整数，支持带逗号的数字
func ParseInt(s string) (int64, error) {
	s = strings.ReplaceAll(s, ",", "")
	s = TrimSpace(s)

	if s == "-" || s == "--" || s == "" {
		return 0, nil
	}

	return strconv.ParseInt(s, 10, 64)
}

// MustParseFloat 解析浮点数，失败返回0
func MustParseFloat(s string) float64 {
	v, _ := ParseFloat(s)
	return v
}

// MustParseInt 解析整数，失败返回0
func MustParseInt(s string) int64 {
	v, _ := ParseInt(s)
	return v
}
