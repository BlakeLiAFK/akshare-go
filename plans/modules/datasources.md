# akshare-go 数据源分析

> 文档: 数据源详细分析
> 创建日期: 2026-01-18

---

## 一、数据源概览

| 数据源 | 接口数量 | 覆盖模块 | 难度 | 反爬措施 |
|--------|----------|----------|------|----------|
| 东方财富 | 175+ | 股票、基金、债券、指数 | 中 | 签名验证 |
| 同花顺 | 47+ | 股票、基金、指数 | 中 | Cookie验证 |
| 金十数据 | 42+ | 宏观经济、期货 | 低 | 无 |
| 新浪财经 | 41+ | 股票、期货、债券 | 中 | JS加密 |
| 乐股网 | 33+ | 股票特色数据 | 低 | Token |
| 巨潮资讯 | 26+ | 上市公司信息 | 中 | Cookie |
| 交易所官网 | 70+ | 期货、期权、债券 | 低 | 无 |
| 集思录 | 10+ | 可转债 | 中 | Cookie |
| AMAC | 14+ | 私募基金 | 低 | 无 |
| 雪球 | 10+ | 股票、基金 | 高 | Cookie+Token |

---

## 二、东方财富 (eastmoney.com)

### 2.1 接口分类

| API 域名 | 用途 | 示例 |
|----------|------|------|
| `push2.eastmoney.com` | 实时行情推送 | 股票列表、实时报价 |
| `push2his.eastmoney.com` | 历史数据 | K线、分钟数据 |
| `datacenter-web.eastmoney.com` | 数据中心 | 资金流向、龙虎榜 |
| `fund.eastmoney.com` | 基金数据 | 基金净值、排行 |
| `api.fund.eastmoney.com` | 基金API | 历史净值 |
| `fundf10.eastmoney.com` | 基金详情 | 持仓、经理 |
| `emweb.eastmoney.com` | 网页数据 | 财务报表 |

### 2.2 常用 URL 模式

```go
// 股票列表
const StockListURL = "https://push2.eastmoney.com/api/qt/clist/get"
// 参数: fs=m:0+t:6,m:0+t:80,m:1+t:2,m:1+t:23&fields=f12,f13,f14

// K线数据
const KLineURL = "https://push2his.eastmoney.com/api/qt/stock/kline/get"
// 参数: secid=1.600000&klt=101&fqt=1&beg=0&end=20500000

// 资金流向
const FundFlowURL = "https://datacenter-web.eastmoney.com/api/data/v1/get"
// 参数: sortColumns=TRADE_DATE&sortTypes=-1&reportName=RPT_MAIN_MONEYFLOW

// 基金净值
const FundNavURL = "https://api.fund.eastmoney.com/f10/lsjz"
// 参数: fundCode=000001&pageIndex=1&pageSize=20
```

### 2.3 请求特点

- **签名**: 部分接口需要动态签名
- **分页**: 支持 pageIndex/pageSize 分页
- **回调**: JSONP 回调格式，需提取数据
- **限流**: 建议 0.5 秒/请求

### 2.4 Go 实现要点

```go
// 东方财富客户端
type EastMoneyClient struct {
    *httpclient.Client
    callback string
}

// 解析 JSONP 回调
func (c *EastMoneyClient) parseCallback(response string) (string, error) {
    // jQuery112308_1642583912345({...})
    start := strings.Index(response, "(")
    end := strings.LastIndex(response, ")")
    if start == -1 || end == -1 {
        return "", errors.New("invalid jsonp")
    }
    return response[start+1 : end], nil
}
```

---

## 三、新浪财经 (sina.com.cn)

### 3.1 接口分类

| API 域名 | 用途 |
|----------|------|
| `hq.sinajs.cn` | 实时行情 |
| `vip.stock.finance.sina.com.cn` | 股票API |
| `finance.sina.com.cn` | 财经数据 |
| `money.finance.sina.com.cn` | 货币数据 |

### 3.2 常用 URL 模式

```go
// 实时行情
const RealtimeURL = "https://hq.sinajs.cn/list=sh600000,sz000001"
// 返回: var hq_str_sh600000="...";

// 股票历史
const HistoryURL = "https://finance.sina.com.cn/realstock/company/%s/hisdata/day/qianfuquan.js"

// 期货行情
const FuturesURL = "https://hq.sinajs.cn/list=nf_RB0"
```

### 3.3 数据加密

部分接口（如港股历史）返回加密数据，需要 JavaScript 解密：

```go
// JS 解密示例
func (c *SinaClient) decryptHKData(encrypted string) (string, error) {
    vm := goja.New()

    // 加载解密脚本
    decryptScript := `
    function decrypt(data) {
        // 新浪解密逻辑
        return decrypted;
    }
    `
    vm.RunString(decryptScript)

    decrypt, _ := goja.AssertFunction(vm.Get("decrypt"))
    result, _ := decrypt(goja.Undefined(), vm.ToValue(encrypted))
    return result.String(), nil
}
```

### 3.4 Go 实现要点

```go
// 解析新浪行情数据
func parseSinaQuote(data string) (*Quote, error) {
    // var hq_str_sh600000="...";
    parts := strings.Split(data, "=")
    if len(parts) != 2 {
        return nil, errors.New("invalid format")
    }

    // 提取引号内内容
    content := strings.Trim(parts[1], "\";\n")
    fields := strings.Split(content, ",")

    return &Quote{
        Name:  fields[0],
        Open:  parseFloat(fields[1]),
        Close: parseFloat(fields[3]),
        // ...
    }, nil
}
```

---

## 四、同花顺 (10jqka.com.cn)

### 4.1 接口分类

| API 域名 | 用途 |
|----------|------|
| `q.10jqka.com.cn` | 行情数据 |
| `data.10jqka.com.cn` | 数据中心 |
| `basic.10jqka.com.cn` | 基础数据 |

### 4.2 请求特点

- **Cookie**: 需要携带有效 Cookie
- **Hexin-V**: 部分接口需要 Hexin-V 头
- **反爬**: 有 WAF 防护

### 4.3 Go 实现要点

```go
// 同花顺客户端
type THSClient struct {
    *httpclient.Client
    hexinV string
}

// 获取 Hexin-V
func (c *THSClient) getHexinV() (string, error) {
    // 访问首页获取 Cookie 和 Hexin-V
    resp, err := c.Get(ctx, "https://q.10jqka.com.cn/", nil)
    if err != nil {
        return "", err
    }

    // 从响应中提取
    hexinV := resp.Header.Get("Hexin-V")
    return hexinV, nil
}
```

---

## 五、金十数据 (jin10.com)

### 5.1 接口特点

- **开放性**: 大部分接口无需认证
- **数据格式**: JSON 格式，结构清晰
- **更新频率**: 实时更新

### 5.2 常用 URL 模式

```go
// 宏观数据日历
const CalendarURL = "https://cdn-rili.jin10.com/web_data/%s/%s.json"

// 经济指标
const IndicatorURL = "https://datacenter-api.jin10.com/reports/list"
```

### 5.3 Go 实现要点

```go
// 金十数据客户端
type Jin10Client struct {
    *httpclient.Client
}

// 获取经济日历
func (c *Jin10Client) GetEconomicCalendar(date time.Time) ([]Event, error) {
    url := fmt.Sprintf(CalendarURL,
        date.Format("200601"),
        date.Format("20060102"))

    var events []Event
    err := c.GetJSON(ctx, url, nil, &events)
    return events, err
}
```

---

## 六、巨潮资讯 (cninfo.com.cn)

### 6.1 接口特点

- **官方数据**: 证监会指定信披平台
- **Cookie**: 需要获取 Cookie
- **限流**: 有请求频率限制

### 6.2 常用 URL 模式

```go
// 股票数据
const StockURL = "https://webapi.cninfo.com.cn/api/stock/"

// 公告信息
const DisclosureURL = "https://webapi.cninfo.com.cn/api/disclosure/"
```

---

## 七、交易所官网

### 7.1 期货交易所

| 交易所 | 域名 | 数据类型 |
|--------|------|----------|
| 上期所 | www.shfe.com.cn | 日数据、持仓、仓单 |
| 大商所 | www.dce.com.cn | 日数据、持仓、仓单 |
| 郑商所 | www.czce.com.cn | 日数据、持仓、仓单 |
| 中金所 | www.cffex.com.cn | 日数据、持仓 |
| 广期所 | www.gfex.com.cn | 日数据、持仓、仓单 |
| 上期能源 | www.ine.cn | 原油、LPG |

### 7.2 证券交易所

| 交易所 | 域名 | 数据类型 |
|--------|------|----------|
| 上交所 | www.sse.com.cn | 股票、债券、期权 |
| 深交所 | www.szse.cn | 股票、债券、期权 |
| 北交所 | www.bse.cn | 股票 |

### 7.3 Go 实现要点

```go
// 交易所数据通常是静态文件
const SHFEDailyURL = "https://www.shfe.com.cn/data/dailydata/%s/%s_daily.dat"

// 下载并解析
func (c *ExchangeClient) GetSHFEDaily(date time.Time) ([]FuturesDaily, error) {
    url := fmt.Sprintf(SHFEDailyURL,
        date.Format("200601"),
        date.Format("20060102"))

    resp, err := c.Get(ctx, url, nil)
    // 解析 DAT 文件格式
}
```

---

## 八、其他数据源

### 8.1 集思录 (jisilu.cn)

- **用途**: 可转债数据
- **认证**: Cookie 登录
- **限制**: 部分数据需要会员

### 8.2 AMAC (amac.org.cn)

- **用途**: 基金业协会私募数据
- **特点**: 官方数据，更新较慢
- **认证**: 无需认证

### 8.3 雪球 (xueqiu.com)

- **用途**: 股票、基金社区数据
- **认证**: 需要 Cookie 和 Token
- **反爬**: 较严格

---

## 九、通用实现建议

### 9.1 HTTP 客户端配置

```go
// 针对不同数据源的配置
var DataSourceConfigs = map[string]ClientConfig{
    "eastmoney": {
        RateLimit:   2.0,    // 每秒2个请求
        Timeout:     10 * time.Second,
        MaxRetries:  3,
        UserAgent:   "Mozilla/5.0...",
    },
    "sina": {
        RateLimit:   5.0,
        Timeout:     5 * time.Second,
        MaxRetries:  2,
    },
    "ths": {
        RateLimit:   1.0,
        Timeout:     15 * time.Second,
        MaxRetries:  3,
        NeedCookie:  true,
    },
}
```

### 9.2 错误处理

```go
// 数据源特定错误
var (
    ErrDataSourceUnavailable = errors.New("data source unavailable")
    ErrRateLimited           = errors.New("rate limited")
    ErrAuthRequired          = errors.New("authentication required")
    ErrDataExpired           = errors.New("data expired")
)

// 自动重试策略
func shouldRetry(err error) bool {
    switch {
    case errors.Is(err, ErrRateLimited):
        return true
    case errors.Is(err, ErrDataSourceUnavailable):
        return true
    default:
        return false
    }
}
```

### 9.3 缓存策略

```go
// 数据缓存配置
var CacheConfigs = map[string]time.Duration{
    "realtime":   0,               // 不缓存
    "daily":      1 * time.Hour,   // 日数据缓存1小时
    "historical": 24 * time.Hour,  // 历史数据缓存1天
    "static":     7 * 24 * time.Hour, // 静态数据缓存1周
}
```

---

## 十、监控与报警

### 10.1 接口健康检查

```go
// 定期检查数据源可用性
func (m *Monitor) CheckDataSources() map[string]HealthStatus {
    results := make(map[string]HealthStatus)

    for name, checker := range m.checkers {
        status := checker.Check()
        results[name] = status

        if status.Status != "healthy" {
            m.alerter.Send(fmt.Sprintf("数据源 %s 异常: %s", name, status.Message))
        }
    }

    return results
}
```

### 10.2 数据质量检查

```go
// 检查数据完整性
func (v *Validator) ValidateData(data interface{}) error {
    // 检查空值
    // 检查数据范围
    // 检查时间连续性
    return nil
}
```

---

*文档版本: v1.0*
*最后更新: 2026-01-18*
