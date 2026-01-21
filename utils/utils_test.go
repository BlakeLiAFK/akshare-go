package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestConfig 测试配置管理
func TestConfig(t *testing.T) {
	config := GetConfig()
	assert.NotNil(t, config, "配置实例不应为空")

	// 测试代理设置
	SetGlobalProxy("http://127.0.0.1:7890")
	assert.Equal(t, "http://127.0.0.1:7890", GetGlobalProxy())

	// 清除代理
	SetGlobalProxy("")
	assert.Equal(t, "", GetGlobalProxy())
}

// TestProxyScope 测试代理作用域
func TestProxyScope(t *testing.T) {
	// 设置初始代理
	SetGlobalProxy("http://initial:8080")
	assert.Equal(t, "http://initial:8080", GetGlobalProxy())

	// 创建临时作用域
	scope := WithProxy("http://temp:9090")
	assert.Equal(t, "http://temp:9090", GetGlobalProxy())

	// 恢复
	scope.Restore()
	assert.Equal(t, "http://initial:8080", GetGlobalProxy())

	// 清理
	SetGlobalProxy("")
}

// TestParseFloat 测试浮点数解析
func TestParseFloat(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"123.45", 123.45},
		{"1,234.56", 1234.56},
		{"-123.45", -123.45},
		{"50%", 0.5},
		{"1.5万", 15000},
		{"2亿", 200000000},
		{"-", 0},
		{"--", 0},
		{"", 0},
	}

	for _, tt := range tests {
		result, err := ParseFloat(tt.input)
		assert.NoError(t, err)
		assert.Equal(t, tt.expected, result, "输入: %s", tt.input)
	}
}

// TestParseInt 测试整数解析
func TestParseInt(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"123", 123},
		{"1,234", 1234},
		{"-123", -123},
		{"-", 0},
		{"", 0},
	}

	for _, tt := range tests {
		result, err := ParseInt(tt.input)
		assert.NoError(t, err)
		assert.Equal(t, tt.expected, result, "输入: %s", tt.input)
	}
}

// TestParseDate 测试日期解析
func TestParseDate(t *testing.T) {
	tests := []string{
		"2024-01-15",
		"20240115",
		"2024/01/15",
		"2024.01.15",
	}

	for _, input := range tests {
		result, err := ParseDate(input)
		assert.NoError(t, err)
		assert.Equal(t, 2024, result.Year(), "输入: %s", input)
		assert.Equal(t, 1, int(result.Month()), "输入: %s", input)
		assert.Equal(t, 15, result.Day(), "输入: %s", input)
	}
}

// TestChunk 测试切片分块
func TestChunk(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	chunks := Chunk(slice, 3)

	assert.Len(t, chunks, 3)
	assert.Equal(t, []int{1, 2, 3}, chunks[0])
	assert.Equal(t, []int{4, 5, 6}, chunks[1])
	assert.Equal(t, []int{7}, chunks[2])
}

// TestContains 测试包含检查
func TestContains(t *testing.T) {
	slice := []string{"a", "b", "c"}
	assert.True(t, Contains(slice, "b"))
	assert.False(t, Contains(slice, "d"))
}

// TestUnique 测试去重
func TestUnique(t *testing.T) {
	slice := []int{1, 2, 2, 3, 3, 3, 4}
	result := Unique(slice)
	assert.Equal(t, []int{1, 2, 3, 4}, result)
}

// TestMap 测试 Map 函数
func TestMap(t *testing.T) {
	slice := []int{1, 2, 3}
	result := Map(slice, func(i int) int { return i * 2 })
	assert.Equal(t, []int{2, 4, 6}, result)
}

// TestFilter 测试 Filter 函数
func TestFilter(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	result := Filter(slice, func(i int) bool { return i%2 == 0 })
	assert.Equal(t, []int{2, 4}, result)
}

// TestToken 测试 Token 管理
func TestToken(t *testing.T) {
	// 先删除可能存在的 token
	_ = DeleteToken()

	// 检查不存在
	assert.False(t, HasToken())

	// 设置 token
	err := SetToken("test_token_12345")
	assert.NoError(t, err)

	// 验证存在
	assert.True(t, HasToken())

	// 读取 token
	token, err := GetToken()
	assert.NoError(t, err)
	assert.Equal(t, "test_token_12345", token)

	// 清理
	err = DeleteToken()
	assert.NoError(t, err)
	assert.False(t, HasToken())
}

// TestToString 测试字符串转换
func TestToString(t *testing.T) {
	assert.Equal(t, "123", ToString(123))
	assert.Equal(t, "123.45", ToString(123.45))
	assert.Equal(t, "true", ToString(true))
	assert.Equal(t, "hello", ToString("hello"))
	assert.Equal(t, "", ToString(nil))
}

// TestToBool 测试布尔转换
func TestToBool(t *testing.T) {
	assert.True(t, ToBool(true))
	assert.True(t, ToBool(1))
	assert.True(t, ToBool("true"))
	assert.False(t, ToBool(false))
	assert.False(t, ToBool(0))
	assert.False(t, ToBool(nil))
}
