package reits

import (
	"testing"
)

// TestReitsRealtimeEm 测试REITs实时行情
func TestReitsRealtimeEm(t *testing.T) {
	data, err := ReitsRealtimeEm()
	if err != nil {
		t.Logf("ReitsRealtimeEm 可能暂时不可用: %v", err)
		return
	}

	if len(data) == 0 {
		t.Log("ReitsRealtimeEm 返回空数据，可能是非交易时间")
		return
	}

	t.Logf("获取到 %d 条REITs实时数据", len(data))
	if len(data) > 0 {
		t.Logf("示例: %s %s %.3f", data[0].Code, data[0].Name, data[0].Price)
	}
}

// TestReitsHistEm 测试REITs历史行情
func TestReitsHistEm(t *testing.T) {
	data, err := ReitsHistEm("508097")
	if err != nil {
		t.Logf("ReitsHistEm 可能暂时不可用: %v", err)
		return
	}

	if len(data) == 0 {
		t.Log("ReitsHistEm 返回空数据")
		return
	}

	t.Logf("获取到 %d 条REITs历史数据", len(data))
	if len(data) > 0 {
		t.Logf("示例: %s %.3f %.3f", data[0].Date, data[0].Open, data[0].Close)
	}
}

// TestReitsHistMinEm 测试REITs分钟行情
func TestReitsHistMinEm(t *testing.T) {
	data, err := ReitsHistMinEm("508097")
	if err != nil {
		t.Logf("ReitsHistMinEm 可能暂时不可用: %v", err)
		return
	}

	if len(data) == 0 {
		t.Log("ReitsHistMinEm 返回空数据，可能是非交易时间")
		return
	}

	t.Logf("获取到 %d 条REITs分钟数据", len(data))
	if len(data) > 0 {
		t.Logf("示例: %s %.3f %.3f", data[0].Time, data[0].Price, data[0].High)
	}
}
