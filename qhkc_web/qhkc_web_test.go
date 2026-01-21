package qhkc_web

import (
	"testing"
)

// TestQhkcWebBasis 测试基差数据
func TestQhkcWebBasis(t *testing.T) {
	data, err := QhkcWebBasis("RB")
	if err != nil {
		t.Logf("QhkcWebBasis 可能暂时不可用: %v", err)
		return
	}
	if len(data) == 0 {
		t.Log("QhkcWebBasis 返回空数据")
		return
	}
	t.Logf("获取到 %d 条基差数据", len(data))
}

// TestQhkcWebInventory 测试库存数据
func TestQhkcWebInventory(t *testing.T) {
	data, err := QhkcWebInventory("RB")
	if err != nil {
		t.Logf("QhkcWebInventory 可能暂时不可用: %v", err)
		return
	}
	if len(data) == 0 {
		t.Log("QhkcWebInventory 返回空数据")
		return
	}
	t.Logf("获取到 %d 条库存数据", len(data))
}

// TestQhkcWebProfit 测试利润数据
func TestQhkcWebProfit(t *testing.T) {
	data, err := QhkcWebProfit("RB")
	if err != nil {
		t.Logf("QhkcWebProfit 可能暂时不可用: %v", err)
		return
	}
	if len(data) == 0 {
		t.Log("QhkcWebProfit 返回空数据")
		return
	}
	t.Logf("获取到 %d 条利润数据", len(data))
}
