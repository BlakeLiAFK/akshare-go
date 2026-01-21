package forex

import (
	"testing"
)

// TestForexSpotEm 测试东方财富外汇实时行情
func TestForexSpotEm(t *testing.T) {
	data, err := ForexSpotEm()
	if err != nil {
		t.Logf("ForexSpotEm 可能暂时不可用: %v", err)
		return
	}
	if len(data) == 0 {
		t.Log("ForexSpotEm 返回空数据")
		return
	}
	t.Logf("获取到 %d 条外汇实时数据", len(data))
	if len(data) > 0 {
		t.Logf("示例: %s %.4f %.4f%%", data[0].Name, data[0].LatestPrice, data[0].ChangeRate)
	}
}

// TestForexHistEm 测试东方财富外汇历史行情
func TestForexHistEm(t *testing.T) {
	data, err := ForexHistEm("USDCNH")
	if err != nil {
		t.Logf("ForexHistEm 可能暂时不可用: %v", err)
		return
	}
	if len(data) == 0 {
		t.Log("ForexHistEm 返回空数据")
		return
	}
	t.Logf("获取到 %d 条外汇历史数据", len(data))
	if len(data) > 0 {
		t.Logf("示例: %s %.4f %.4f", data[0].Date, data[0].Open, data[0].Close)
	}
}

// TestGetMarketCode 测试获取市场代码
func TestGetMarketCode(t *testing.T) {
	code := GetMarketCode("USDCNH")
	if code == 0 {
		t.Log("GetMarketCode 返回默认值")
		return
	}
	t.Logf("USDCNH 市场代码: %d", code)
}
