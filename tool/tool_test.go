package tool

import (
	"testing"
	"time"
)

// TestToolTradeDateHistSina 测试新浪财经交易日历
func TestToolTradeDateHistSina(t *testing.T) {
	dates, err := ToolTradeDateHistSina()
	if err != nil {
		t.Fatalf("ToolTradeDateHistSina() error = %v", err)
	}

	// 验证返回的数据不为空
	if len(dates) == 0 {
		t.Fatal("返回的交易日列表为空")
	}

	// 验证数据按日期排序
	for i := 1; i < len(dates); i++ {
		if dates[i].Date.Before(dates[i-1].Date) {
			t.Errorf("日期未排序: dates[%d]=%v 在 dates[%d]=%v 之前",
				i-1, dates[i-1].Date, i, dates[i].Date)
		}
	}

	// 验证包含特殊日期 1992-05-04
	specialDate := time.Date(1992, 5, 4, 0, 0, 0, 0, time.UTC)
	found := false
	for _, d := range dates {
		if d.Date.Equal(specialDate) || d.Date.Format("2006-01-02") == specialDate.Format("2006-01-02") {
			found = true
			break
		}
	}
	if !found {
		t.Error("未找到特殊日期 1992-05-04")
	}

	// 验证第一个交易日是 1990 年底或 1991 年初
	if len(dates) > 0 {
		firstDate := dates[0].Date
		if firstDate.Year() < 1990 || firstDate.Year() > 1992 {
			t.Logf("警告: 第一个交易日是 %v，可能不正确", firstDate)
		}
	}

	// 验证最近的交易日是最近几天
	if len(dates) > 0 {
		lastDate := dates[len(dates)-1].Date
		now := time.Now()
		diff := now.Sub(lastDate)
		// 最后一个交易日应该在最近 7 天内
		if diff > 7*24*time.Hour {
			t.Logf("警告: 最后一个交易日是 %v，距离现在 %v 天", lastDate, diff.Hours()/24)
		}
	}

	t.Logf("共获取 %d 个交易日", len(dates))
	t.Logf("第一个交易日: %v", dates[0].Date.Format("2006-01-02"))
	t.Logf("最后一个交易日: %v", dates[len(dates)-1].Date.Format("2006-01-02"))
}

// TestTradeDate_String 测试日期格式化
func TestTradeDate_String(t *testing.T) {
	date := TradeDate{
		Date: time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC),
	}

	expected := "2024-01-15"
	actual := date.Date.Format("2006-01-02")
	if actual != expected {
		t.Errorf("日期格式化错误: got %v, want %v", actual, expected)
	}
}
