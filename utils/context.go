package utils

import (
	"sync"
)

// Config 全局配置（单例模式）
type Config struct {
	proxy string
	mu    sync.RWMutex
}

var (
	globalConfig *Config
	once         sync.Once
)

// GetConfig 获取全局配置实例
func GetConfig() *Config {
	once.Do(func() {
		globalConfig = &Config{}
	})
	return globalConfig
}

// SetProxy 设置代理
func (c *Config) SetProxy(proxy string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.proxy = proxy

	// 同步更新 HTTP 客户端代理
	if proxy == "" {
		HTTPClient.RemoveProxy()
	} else {
		HTTPClient.SetProxy(proxy)
	}
}

// GetProxy 获取代理
func (c *Config) GetProxy() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.proxy
}

// SetGlobalProxy 设置全局代理（快捷函数）
func SetGlobalProxy(proxy string) {
	GetConfig().SetProxy(proxy)
}

// GetGlobalProxy 获取全局代理（快捷函数）
func GetGlobalProxy() string {
	return GetConfig().GetProxy()
}

// ProxyScope 代理作用域，用于临时切换代理
type ProxyScope struct {
	oldProxy string
}

// WithProxy 创建临时代理作用域
// 用法:
//
//	scope := utils.WithProxy("http://127.0.0.1:7890")
//	defer scope.Restore()
//	// 在此作用域内使用临时代理
func WithProxy(proxy string) *ProxyScope {
	scope := &ProxyScope{
		oldProxy: GetGlobalProxy(),
	}
	SetGlobalProxy(proxy)
	return scope
}

// Restore 恢复原代理设置
func (s *ProxyScope) Restore() {
	SetGlobalProxy(s.oldProxy)
}
