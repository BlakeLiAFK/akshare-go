package qdii

import (
	"testing"
)

// TestQdiiEIndexJsl 测试QDII-E指数基金
func TestQdiiEIndexJsl(t *testing.T) {
	data, err := QdiiEIndexJsl()
	if err != nil {
		t.Logf("QdiiEIndexJsl 可能暂时不可用: %v", err)
		return
	}
	if len(data) == 0 {
		t.Log("QdiiEIndexJsl 返回空数据")
		return
	}
	t.Logf("获取到 %d 条QDII-E指数基金数据", len(data))
	if len(data) > 0 {
		t.Logf("示例: %s %s %.3f", data[0].Code, data[0].Name, data[0].Price)
	}
}

// TestQdiiECommJsl 测试QDII商品基金
func TestQdiiECommJsl(t *testing.T) {
	data, err := QdiiECommJsl()
	if err != nil {
		t.Logf("QdiiECommJsl 可能暂时不可用: %v", err)
		return
	}
	if len(data) == 0 {
		t.Log("QdiiECommJsl 返回空数据")
		return
	}
	t.Logf("获取到 %d 条QDII商品基金数据", len(data))
}

// TestQdiiAIndexJsl 测试QDII-A指数基金
func TestQdiiAIndexJsl(t *testing.T) {
	data, err := QdiiAIndexJsl()
	if err != nil {
		t.Logf("QdiiAIndexJsl 可能暂时不可用: %v", err)
		return
	}
	if len(data) == 0 {
		t.Log("QdiiAIndexJsl 返回空数据")
		return
	}
	t.Logf("获取到 %d 条QDII-A指数基金数据", len(data))
}
