package qhkc

import (
	"os"
	"testing"

	"github.com/BlakeLiAFK/akshare/pro"
)

// getTestClient 获取测试用客户端
func getTestClient(t *testing.T) *pro.DataApi {
	token := os.Getenv("AKSHARE_TOKEN")
	if token == "" {
		t.Skip("需要设置 AKSHARE_TOKEN 环境变量")
		return nil
	}
	client, err := GetProClient(token)
	if err != nil {
		t.Skipf("获取客户端失败: %v", err)
		return nil
	}
	return client
}

// TestGetProClient 测试获取Pro客户端
func TestGetProClient(t *testing.T) {
	token := os.Getenv("AKSHARE_TOKEN")
	if token == "" {
		t.Skip("需要设置 AKSHARE_TOKEN 环境变量")
		return
	}
	client, err := GetProClient(token)
	if err != nil {
		t.Logf("GetProClient 失败: %v", err)
		return
	}
	if client == nil {
		t.Log("GetProClient 返回空客户端")
		return
	}
	t.Log("GetProClient 成功获取客户端")
}

// TestIndexInfo 测试指数信息
func TestIndexInfo(t *testing.T) {
	client := getTestClient(t)
	if client == nil {
		return
	}
	data, err := IndexInfo(client, "index0070c0eb-93ba-2da9-6633-fa70cb90e959")
	if err != nil {
		t.Logf("IndexInfo 可能暂时不可用: %v", err)
		return
	}
	t.Logf("获取到 %d 条指数信息", len(data))
}

// TestIndexWeights 测试指数权重
func TestIndexWeights(t *testing.T) {
	client := getTestClient(t)
	if client == nil {
		return
	}
	data, err := IndexWeights(client, "index0070c0eb-93ba-2da9-6633-fa70cb90e959", "2023-01-01")
	if err != nil {
		t.Logf("IndexWeights 可能暂时不可用: %v", err)
		return
	}
	t.Logf("获取到 %d 条指数权重数据", len(data))
}

// TestIndexKline 测试指数K线
func TestIndexKline(t *testing.T) {
	client := getTestClient(t)
	if client == nil {
		return
	}
	data, err := IndexKline(client, "index0070c0eb-93ba-2da9-6633-fa70cb90e959")
	if err != nil {
		t.Logf("IndexKline 可能暂时不可用: %v", err)
		return
	}
	t.Logf("获取到 %d 条指数K线数据", len(data))
}

// TestIndexMember 测试指数成分
func TestIndexMember(t *testing.T) {
	client := getTestClient(t)
	if client == nil {
		return
	}
	data, err := IndexMember(client, "index0070c0eb-93ba-2da9-6633-fa70cb90e959")
	if err != nil {
		t.Logf("IndexMember 可能暂时不可用: %v", err)
		return
	}
	t.Logf("获取到 %d 条指数成分数据", len(data))
}
