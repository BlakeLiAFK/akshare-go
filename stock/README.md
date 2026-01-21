# Stock 股票数据模块

## 功能概览

本模块提供A股市场数据获取功能，包括实时行情、历史K线、分时数据等。

## 已实现接口

### 1. StockZhASpotEm - A股实时行情

获取全部A股实时行情数据（东方财富数据源）。

```go
quotes, err := stock.StockZhASpotEm()
if err != nil {
    log.Fatal(err)
}

for _, q := range quotes[:10] {
    fmt.Printf("%s(%s): %.2f %.2f%%\n", q.Name, q.Code, q.Price, q.ChangePct)
}
```

### 2. StockZhAHist - A股历史K线

获取个股历史K线数据，支持日K、周K、月K，支持前复权、后复权。

**参数说明:**
- `code`: 股票代码，如 "000001"
- `period`: 周期，"daily"(日K)、"weekly"(周K)、"monthly"(月K)
- `startDate`: 开始日期，格式 "20240101"
- `endDate`: 结束日期，格式 "20240131"
- `adjust`: 复权类型，"qfq"(前复权)、"hfq"(后复权)、""(不复权)

**示例:**

```go
// 获取平安银行2024年1月日K线数据（前复权）
klines, err := stock.StockZhAHist("000001", "daily", "20240101", "20240131", "qfq")
if err != nil {
    log.Fatal(err)
}

for _, k := range klines {
    fmt.Printf("%s: 开盘=%.2f, 收盘=%.2f, 涨跌幅=%.2f%%\n",
        k.Date.Format("2006-01-02"), k.Open, k.Close, k.ChangePct)
}
```

**周K线示例:**

```go
// 获取贵州茅台2024年周K线（后复权）
klines, err := stock.StockZhAHist("600519", "weekly", "20240101", "20241231", "hfq")
```

**月K线示例:**

```go
// 获取宁德时代2023年月K线（不复权）
klines, err := stock.StockZhAHist("300750", "monthly", "20230101", "20231231", "")
```

### 3. StockIntradayEm - 日内分时数据

获取个股当日分时数据。

```go
quotes, err := stock.StockIntradayEm("000001")
if err != nil {
    log.Fatal(err)
}

for _, q := range quotes {
    fmt.Printf("%s: %.2f\n", q.Time, q.Price)
}
```

## 数据结构

### StockKLine - K线数据

```go
type StockKLine struct {
    Date      time.Time // 日期
    Open      float64   // 开盘价
    High      float64   // 最高价
    Low       float64   // 最低价
    Close     float64   // 收盘价
    Volume    int64     // 成交量
    Amount    float64   // 成交额
    Adjust    string    // 复权类型
    Change    float64   // 涨跌额
    ChangePct float64   // 涨跌幅
    Turnover  float64   // 换手率
}
```

### StockQuote - 实时行情

```go
type StockQuote struct {
    Code         string    // 股票代码
    Name         string    // 股票名称
    Open         float64   // 开盘价
    High         float64   // 最高价
    Low          float64   // 最低价
    Price        float64   // 最新价
    PreClose     float64   // 昨收价
    Change       float64   // 涨跌额
    ChangePct    float64   // 涨跌幅(%)
    Volume       int64     // 成交量(手)
    Amount       float64   // 成交额(元)
    TurnoverRate float64   // 换手率(%)
    PE           float64   // 市盈率
    PB           float64   // 市净率
    MarketCap    float64   // 总市值
    CirculateCap float64   // 流通市值
    Time         time.Time // 时间
}
```

## 复权说明

- **不复权 ("")**: 显示真实交易价格，不考虑分红送股影响
- **前复权 ("qfq")**: 以最新价为基准，向前调整历史价格，保持价格连续性
- **后复权 ("hfq")**: 以上市价为基准，向后调整价格，可看到真实涨幅

## 支持的股票市场

- **沪市主板**: 6开头（如：600519 贵州茅台）
- **科创板**: 688开头（如：688001）
- **深市主板**: 000开头（如：000001 平安银行）
- **创业板**: 300开头（如：300750 宁德时代）
- **北交所**: 8开头、4开头

## 性能特点

- 单次K线查询耗时约 500ms（含网络延迟）
- 内存占用约 100KB/次
- 支持HTTP重试和限流保护

## 示例程序

完整示例请参考 [examples/stock_hist_example.go](../examples/stock_hist_example.go)

## 注意事项

1. 数据来源于东方财富网，仅供学习研究使用
2. 建议添加适当的调用频率限制，避免过度请求
3. 日期格式必须为 "20240101" 格式
4. 周K和月K的日期字段为该周期的最后一个交易日

## 待实现功能

- [ ] 龙虎榜数据
- [ ] 资金流向
- [ ] 板块数据
- [ ] 财务指标
- [ ] 股东信息
