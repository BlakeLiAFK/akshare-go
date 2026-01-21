package economic

import (
	"testing"
	"time"
)

func TestMacroInfoWS(t *testing.T) {
	// 使用昨天的日期，确保有数据
	yesterday := time.Now().AddDate(0, 0, -1).Format("20060102")
	df, err := MacroInfoWS(yesterday)
	if err != nil {
		t.Fatalf("MacroInfoWS 失败: %v", err)
	}
	t.Logf("宏观日历数据 (%s) 行数: %d, 列数: %d", yesterday, df.Nrow(), df.Ncol())
}

func TestMacroInfoWSToday(t *testing.T) {
	// 测试今天的数据
	today := time.Now().Format("20060102")
	df, err := MacroInfoWS(today)
	if err != nil {
		t.Fatalf("MacroInfoWS 今天数据失败: %v", err)
	}
	t.Logf("宏观日历数据 (%s) 行数: %d, 列数: %d", today, df.Nrow(), df.Ncol())
}

func TestMacroInfoWSInvalidDate(t *testing.T) {
	// 测试无效日期格式
	_, err := MacroInfoWS("invalid")
	if err == nil {
		t.Error("应该返回错误，但没有")
	}
}
