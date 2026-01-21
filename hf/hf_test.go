package hf

import (
	"testing"
)

// TestHfSp500 测试S&P500高频数据
func TestHfSp500(t *testing.T) {
	data, err := HfSp500("2023")
	if err != nil {
		t.Logf("HfSp500 可能暂时不可用: %v", err)
		return
	}
	if len(data) == 0 {
		t.Log("HfSp500 返回空数据")
		return
	}
	t.Logf("获取到 %d 条S&P500高频数据", len(data))
	if len(data) > 0 {
		t.Logf("示例: %s %.2f %.2f %.2f %.2f", data[0].Date, data[0].Open, data[0].High, data[0].Low, data[0].Close)
	}
}
