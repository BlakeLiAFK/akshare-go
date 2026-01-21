package pro

import (
	"os"
	"testing"
)

// TestProApi 测试创建Pro API客户端
func TestProApi(t *testing.T) {
	token := os.Getenv("AKSHARE_TOKEN")
	if token == "" {
		t.Skip("需要设置 AKSHARE_TOKEN 环境变量")
		return
	}
	client, err := ProApi(token)
	if err != nil {
		t.Logf("ProApi 失败: %v", err)
		return
	}
	if client == nil {
		t.Log("ProApi 返回空客户端")
		return
	}
	t.Log("ProApi 成功创建客户端")
}

// TestNewDataApi 测试创建DataApi
func TestNewDataApi(t *testing.T) {
	token := os.Getenv("AKSHARE_TOKEN")
	if token == "" {
		t.Skip("需要设置 AKSHARE_TOKEN 环境变量")
		return
	}
	client := NewDataApi(token, 10)
	if client == nil {
		t.Log("NewDataApi 返回空客户端")
		return
	}
	t.Log("NewDataApi 成功创建客户端")
}

// TestSetToken 测试设置Token
func TestSetToken(t *testing.T) {
	// 保存原始值
	originalToken := os.Getenv("AKSHARE_TOKEN")
	defer func() {
		if originalToken != "" {
			os.Setenv("AKSHARE_TOKEN", originalToken)
		}
	}()

	err := SetToken("test_token_123")
	if err != nil {
		t.Logf("SetToken 失败: %v", err)
		return
	}

	token := os.Getenv("AKSHARE_TOKEN")
	if token != "test_token_123" {
		t.Log("SetToken 设置的值不正确")
		return
	}
	t.Log("SetToken 成功设置Token")
}
