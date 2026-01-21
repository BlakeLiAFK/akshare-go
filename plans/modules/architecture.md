# akshare-go 技术架构设计

> 模块: 基础设施
> 优先级: P0
> 预计周期: 第1-2周

---

## 一、目录结构详细设计

```
akshare-go/
├── go.mod
├── go.sum
├── README.md
├── LICENSE                     # Apache 2.0
├── Makefile                    # 构建脚本
├── .golangci.yml               # 代码检查配置
│
├── akshare.go                  # 主入口文件，导出所有公开函数
│
├── internal/                   # 内部包（不对外暴露）
│   │
│   ├── httpclient/             # HTTP 客户端封装
│   │   ├── client.go           # 基础客户端结构
│   │   ├── client_test.go
│   │   ├── retry.go            # 指数退避重试
│   │   ├── retry_test.go
│   │   ├── proxy.go            # 代理管理
│   │   ├── proxy_test.go
│   │   ├── ratelimit.go        # 请求限流
│   │   ├── ratelimit_test.go
│   │   ├── middleware.go       # 中间件（日志、监控）
│   │   └── options.go          # 配置选项
│   │
│   ├── parser/                 # 数据解析器
│   │   ├── html.go             # HTML/表格解析 (goquery)
│   │   ├── html_test.go
│   │   ├── json.go             # JSON 解析
│   │   ├── json_test.go
│   │   ├── table.go            # 通用表格处理
│   │   ├── table_test.go
│   │   ├── js.go               # JavaScript 解密 (goja)
│   │   ├── js_test.go
│   │   ├── csv.go              # CSV 解析
│   │   └── excel.go            # Excel 解析
│   │
│   ├── utils/                  # 工具函数
│   │   ├── date.go             # 日期时间处理
│   │   ├── date_test.go
│   │   ├── number.go           # 数字格式转换
│   │   ├── number_test.go
│   │   ├── encoding.go         # 编码处理（GBK等）
│   │   ├── encoding_test.go
│   │   ├── validate.go         # 参数验证
│   │   ├── validate_test.go
│   │   ├── cache.go            # 内存缓存
│   │   └── trading_day.go      # 交易日历
│   │
│   └── types/                  # 内部类型定义
│       ├── errors.go           # 错误类型
│       ├── dataframe.go        # DataFrame 兼容
│       ├── config.go           # 配置结构
│       └── constants.go        # 常量定义
│
├── pkg/                        # 公开包
│   │
│   ├── types/                  # 公开类型定义
│   │   ├── stock.go            # 股票相关类型
│   │   ├── fund.go             # 基金相关类型
│   │   ├── futures.go          # 期货相关类型
│   │   ├── bond.go             # 债券相关类型
│   │   ├── index.go            # 指数相关类型
│   │   ├── option.go           # 期权相关类型
│   │   ├── economic.go         # 宏观经济类型
│   │   ├── common.go           # 通用类型
│   │   └── errors.go           # 公开错误类型
│   │
│   └── config/                 # 配置管理
│       ├── config.go           # 全局配置
│       ├── config_test.go
│       └── options.go          # 配置选项
│
├── stock/                      # 股票模块
├── stock_feature/              # 股票特色数据
├── stock_fundamental/          # 股票基本面
├── fund/                       # 基金模块
├── futures/                    # 期货模块
├── futures_derivative/         # 期货衍生数据
├── bond/                       # 债券模块
├── index/                      # 指数模块
├── option/                     # 期权模块
├── economic/                   # 宏观经济模块
├── currency/                   # 外汇模块
├── crypto/                     # 加密货币模块
├── energy/                     # 能源模块
├── fortune/                    # 财富排行模块
├── news/                       # 新闻模块
├── movie/                      # 电影票房模块
├── air/                        # 空气质量模块
├── article/                    # 学术指数模块
├── spot/                       # 现货模块
├── reits/                      # REITs 模块
├── rate/                       # 利率模块
├── fx/                         # 外汇交易模块
├── bank/                       # 银行模块
│
├── examples/                   # 示例代码
│   ├── stock_example.go
│   ├── fund_example.go
│   └── ...
│
├── testdata/                   # 测试数据
│   ├── stock/
│   ├── fund/
│   └── ...
│
└── scripts/                    # 脚本
    ├── generate.go             # 代码生成
    └── sync_akshare.go         # 同步 akshare 更新
```

---

## 二、核心组件设计

### 2.1 HTTP 客户端 (internal/httpclient)

```go
// client.go
package httpclient

import (
    "context"
    "net/http"
    "time"
)

// Client HTTP 客户端封装
type Client struct {
    httpClient *http.Client
    baseURL    string
    headers    map[string]string
    proxy      *ProxyManager
    limiter    *RateLimiter
    retry      *RetryConfig
}

// Options 客户端配置选项
type Options struct {
    Timeout       time.Duration     // 请求超时
    Proxy         string            // 代理地址
    MaxRetries    int               // 最大重试次数
    RetryDelay    time.Duration     // 重试延迟
    RateLimit     float64           // 每秒请求数
    Headers       map[string]string // 默认请求头
    EnableCookies bool              // 启用 Cookie
}

// NewClient 创建新客户端
func NewClient(opts ...Option) *Client

// Get 发送 GET 请求
func (c *Client) Get(ctx context.Context, url string, params map[string]string) (*Response, error)

// Post 发送 POST 请求
func (c *Client) Post(ctx context.Context, url string, body interface{}) (*Response, error)

// GetJSON 获取 JSON 响应并解析
func (c *Client) GetJSON(ctx context.Context, url string, params map[string]string, result interface{}) error

// GetHTML 获取 HTML 响应
func (c *Client) GetHTML(ctx context.Context, url string, params map[string]string) (*goquery.Document, error)
```

```go
// retry.go
package httpclient

// RetryConfig 重试配置
type RetryConfig struct {
    MaxRetries    int           // 最大重试次数
    InitialDelay  time.Duration // 初始延迟
    MaxDelay      time.Duration // 最大延迟
    Multiplier    float64       // 延迟倍数
    RetryableFunc func(error) bool // 判断是否可重试
}

// WithRetry 添加重试中间件
func (c *Client) WithRetry(config *RetryConfig) *Client

// DefaultRetryConfig 默认重试配置
func DefaultRetryConfig() *RetryConfig
```

```go
// proxy.go
package httpclient

// ProxyManager 代理管理器
type ProxyManager struct {
    proxies    []string
    current    int
    mu         sync.RWMutex
    rotateMode RotateMode
}

// RotateMode 代理轮换模式
type RotateMode int

const (
    RotateRoundRobin RotateMode = iota // 轮询
    RotateRandom                        // 随机
    RotateFailover                      // 故障转移
)

// NewProxyManager 创建代理管理器
func NewProxyManager(proxies []string, mode RotateMode) *ProxyManager

// GetProxy 获取代理
func (pm *ProxyManager) GetProxy() string

// MarkFailed 标记代理失败
func (pm *ProxyManager) MarkFailed(proxy string)
```

```go
// ratelimit.go
package httpclient

import "golang.org/x/time/rate"

// RateLimiter 请求限流器
type RateLimiter struct {
    limiter *rate.Limiter
    domain  map[string]*rate.Limiter // 按域名限流
    mu      sync.RWMutex
}

// NewRateLimiter 创建限流器
func NewRateLimiter(rps float64) *RateLimiter

// Wait 等待限流
func (rl *RateLimiter) Wait(ctx context.Context) error

// WaitDomain 按域名限流
func (rl *RateLimiter) WaitDomain(ctx context.Context, domain string) error
```

---

### 2.2 数据解析器 (internal/parser)

```go
// html.go
package parser

import "github.com/PuerkitoBio/goquery"

// HTMLParser HTML 解析器
type HTMLParser struct {
    doc *goquery.Document
}

// NewHTMLParser 从 HTML 字符串创建解析器
func NewHTMLParser(html string) (*HTMLParser, error)

// NewHTMLParserFromReader 从 Reader 创建解析器
func NewHTMLParserFromReader(r io.Reader) (*HTMLParser, error)

// Find 查找元素
func (p *HTMLParser) Find(selector string) *goquery.Selection

// ParseTable 解析 HTML 表格
func (p *HTMLParser) ParseTable(selector string) ([][]string, error)

// ParseTableToMap 解析表格为 map 切片
func (p *HTMLParser) ParseTableToMap(selector string) ([]map[string]string, error)

// ExtractLinks 提取所有链接
func (p *HTMLParser) ExtractLinks(selector string) []string
```

```go
// json.go
package parser

import "github.com/tidwall/gjson"

// JSONParser JSON 解析器
type JSONParser struct {
    data string
}

// NewJSONParser 创建 JSON 解析器
func NewJSONParser(data string) *JSONParser

// Get 获取 JSON 路径值
func (p *JSONParser) Get(path string) gjson.Result

// GetArray 获取数组
func (p *JSONParser) GetArray(path string) []gjson.Result

// GetMap 获取 map
func (p *JSONParser) GetMap(path string) map[string]gjson.Result

// Unmarshal 反序列化到结构体
func (p *JSONParser) Unmarshal(v interface{}) error

// ExtractDataFromCallback 提取 JSONP 回调中的数据
func ExtractDataFromCallback(response string) (string, error)
```

```go
// js.go
package parser

import "github.com/dop251/goja"

// JSExecutor JavaScript 执行器
type JSExecutor struct {
    vm *goja.Runtime
}

// NewJSExecutor 创建 JS 执行器
func NewJSExecutor() *JSExecutor

// Execute 执行 JavaScript 代码
func (e *JSExecutor) Execute(script string) (goja.Value, error)

// Call 调用 JavaScript 函数
func (e *JSExecutor) Call(funcName string, args ...interface{}) (goja.Value, error)

// DecryptSinaData 解密新浪数据
func (e *JSExecutor) DecryptSinaData(encryptedData string) (string, error)

// SetGlobal 设置全局变量
func (e *JSExecutor) SetGlobal(name string, value interface{}) error
```

---

### 2.3 工具函数 (internal/utils)

```go
// date.go
package utils

import "time"

// ParseDate 解析多种格式的日期字符串
func ParseDate(s string) (time.Time, error)

// FormatDate 格式化日期
func FormatDate(t time.Time, format string) string

// TradingDates 获取交易日列表
func TradingDates(start, end time.Time) []time.Time

// IsTradingDay 判断是否为交易日
func IsTradingDay(t time.Time) bool

// LastTradingDay 获取最近的交易日
func LastTradingDay() time.Time

// NextTradingDay 获取下一个交易日
func NextTradingDay(t time.Time) time.Time
```

```go
// number.go
package utils

// ParseFloat 解析浮点数（支持百分比、千分位等）
func ParseFloat(s string) (float64, error)

// ParseInt 解析整数
func ParseInt(s string) (int64, error)

// FormatFloat 格式化浮点数
func FormatFloat(f float64, precision int) string

// ParsePercent 解析百分比
func ParsePercent(s string) (float64, error)

// ParseChinese 解析中文数字（亿、万等）
func ParseChinese(s string) (float64, error)
```

```go
// encoding.go
package utils

// DecodeGBK 解码 GBK 字符串
func DecodeGBK(data []byte) (string, error)

// EncodeGBK 编码为 GBK
func EncodeGBK(s string) ([]byte, error)

// DetectEncoding 检测字符编码
func DetectEncoding(data []byte) string
```

---

### 2.4 类型定义 (internal/types & pkg/types)

```go
// internal/types/errors.go
package types

import "errors"

var (
    ErrNetwork        = errors.New("network error")
    ErrRateLimit      = errors.New("rate limit exceeded")
    ErrInvalidParam   = errors.New("invalid parameter")
    ErrDataNotFound   = errors.New("data not found")
    ErrParseError     = errors.New("parse error")
    ErrAPIError       = errors.New("api error")
)

// APIError API 错误
type APIError struct {
    Code    int
    Message string
    URL     string
}

func (e *APIError) Error() string
func (e *APIError) Unwrap() error
```

```go
// pkg/types/common.go
package types

import "time"

// DataFrame 通用数据表结构
type DataFrame struct {
    Columns []string                 // 列名
    Data    []map[string]interface{} // 数据行
}

// ToJSON 转换为 JSON
func (df *DataFrame) ToJSON() ([]byte, error)

// ToCSV 转换为 CSV
func (df *DataFrame) ToCSV() ([]byte, error)

// Filter 过滤数据
func (df *DataFrame) Filter(fn func(row map[string]interface{}) bool) *DataFrame

// Sort 排序
func (df *DataFrame) Sort(column string, ascending bool) *DataFrame

// OHLC K 线数据
type OHLC struct {
    Date   time.Time `json:"date"`
    Open   float64   `json:"open"`
    High   float64   `json:"high"`
    Low    float64   `json:"low"`
    Close  float64   `json:"close"`
    Volume int64     `json:"volume"`
    Amount float64   `json:"amount"`
}

// Quote 实时行情
type Quote struct {
    Code      string    `json:"code"`
    Name      string    `json:"name"`
    Price     float64   `json:"price"`
    Change    float64   `json:"change"`
    ChangePct float64   `json:"change_pct"`
    Volume    int64     `json:"volume"`
    Amount    float64   `json:"amount"`
    Time      time.Time `json:"time"`
}
```

---

### 2.5 配置管理 (pkg/config)

```go
// config.go
package config

import (
    "sync"
    "time"
)

// Config 全局配置
type Config struct {
    // HTTP 配置
    Timeout    time.Duration
    Proxy      string
    MaxRetries int
    RateLimit  float64

    // 日志配置
    LogLevel string
    LogFile  string

    // 缓存配置
    EnableCache bool
    CacheTTL    time.Duration
}

var (
    globalConfig *Config
    once         sync.Once
)

// Get 获取全局配置
func Get() *Config

// Set 设置全局配置
func Set(cfg *Config)

// SetProxy 设置代理
func SetProxy(proxy string)

// SetTimeout 设置超时
func SetTimeout(timeout time.Duration)

// SetRateLimit 设置限流
func SetRateLimit(rps float64)

// Load 从文件加载配置
func Load(path string) (*Config, error)

// LoadFromEnv 从环境变量加载
func LoadFromEnv() *Config
```

---

## 三、依赖列表

```go
// go.mod
module github.com/BlakeLiAFK/akshare-go

go 1.21

require (
    // HTTP 客户端
    github.com/go-resty/resty/v2 v2.11.0

    // HTML 解析
    github.com/PuerkitoBio/goquery v1.8.1

    // JSON 解析
    github.com/tidwall/gjson v1.17.0
    github.com/json-iterator/go v1.1.12

    // JavaScript 执行
    github.com/dop251/goja v0.0.0-20231027120936-b396bb4c349d

    // DataFrame
    github.com/go-gota/gota v0.12.0

    // 并发控制
    golang.org/x/sync v0.6.0
    golang.org/x/time v0.5.0

    // 编码处理
    golang.org/x/text v0.14.0

    // 测试
    github.com/stretchr/testify v1.8.4

    // Excel 处理
    github.com/xuri/excelize/v2 v2.8.0
)
```

---

## 四、实现任务清单

### 4.1 HTTP 客户端

- [ ] `internal/httpclient/client.go` - 基础客户端
- [ ] `internal/httpclient/options.go` - 配置选项
- [ ] `internal/httpclient/retry.go` - 重试机制
- [ ] `internal/httpclient/proxy.go` - 代理管理
- [ ] `internal/httpclient/ratelimit.go` - 请求限流
- [ ] `internal/httpclient/middleware.go` - 中间件
- [ ] `internal/httpclient/client_test.go` - 单元测试

### 4.2 数据解析器

- [ ] `internal/parser/html.go` - HTML 解析
- [ ] `internal/parser/json.go` - JSON 解析
- [ ] `internal/parser/table.go` - 表格解析
- [ ] `internal/parser/js.go` - JavaScript 执行
- [ ] `internal/parser/csv.go` - CSV 解析
- [ ] `internal/parser/excel.go` - Excel 解析

### 4.3 工具函数

- [ ] `internal/utils/date.go` - 日期处理
- [ ] `internal/utils/number.go` - 数字转换
- [ ] `internal/utils/encoding.go` - 编码处理
- [ ] `internal/utils/validate.go` - 参数验证
- [ ] `internal/utils/cache.go` - 缓存
- [ ] `internal/utils/trading_day.go` - 交易日历

### 4.4 类型定义

- [ ] `internal/types/errors.go` - 错误类型
- [ ] `internal/types/config.go` - 配置结构
- [ ] `internal/types/constants.go` - 常量
- [ ] `pkg/types/common.go` - 通用类型
- [ ] `pkg/types/stock.go` - 股票类型
- [ ] `pkg/types/fund.go` - 基金类型
- [ ] `pkg/types/futures.go` - 期货类型
- [ ] `pkg/types/bond.go` - 债券类型

### 4.5 配置管理

- [ ] `pkg/config/config.go` - 全局配置
- [ ] `pkg/config/options.go` - 配置选项

---

*文档版本: v1.0*
*最后更新: 2026-01-18*
