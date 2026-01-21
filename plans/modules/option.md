# option 模块计划

> 模块: 期权数据
> 函数数量: 50 个
> 优先级: P2
> 预计周期: 第12周

---

## 一、模块概述

期权模块提供商品期权（各大商品交易所）和金融期权（上交所、深交所、中金所）的行情、持仓、希腊值等数据。

---

## 二、Go 文件结构

```
option/
├── commodity.go         # 商品期权
├── financial.go         # 金融期权
├── sse.go               # 上交所期权
├── szse.go              # 深交所期权
├── cffex.go             # 中金所期权
├── greeks.go            # 希腊值
├── volatility.go        # 波动率
├── analysis.go          # 期权分析
├── lhb.go               # 龙虎榜
└── types.go             # 类型定义
```

---

## 三、函数清单

### 3.1 商品期权 (commodity.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `option_commodity_contract_sina` | 商品期权合约(新浪) | 新浪财经 | P0 |
| `option_commodity_contract_table_sina` | 商品期权合约表 | 新浪财经 | P1 |
| `option_commodity_hist_sina` | 商品期权历史 | 新浪财经 | P0 |
| `option_comm_info` | 商品期权信息 | - | P1 |
| `option_comm_symbol` | 商品期权品种 | - | P1 |
| `option_hist_czce` | 郑商所期权历史 | 郑商所 | P1 |
| `option_hist_dce` | 大商所期权历史 | 大商所 | P1 |
| `option_hist_shfe` | 上期所期权历史 | 上期所 | P1 |
| `option_hist_gfex` | 广期所期权历史 | 广期所 | P1 |
| `option_hist_yearly_czce` | 郑商所年度历史 | 郑商所 | P2 |
| `option_vol_shfe` | 上期所成交量 | 上期所 | P2 |
| `option_vol_gfex` | 广期所成交量 | 广期所 | P2 |

### 3.2 金融期权 (financial.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `option_finance_board` | 金融期权行情 | - | P0 |
| `option_finance_minute_sina` | 金融期权分钟 | 新浪财经 | P1 |
| `option_finance_sse_underlying` | 上交所标的 | 上交所 | P1 |
| `option_current_em` | 期权实时(东财) | 东方财富 | P0 |
| `option_minute_em` | 期权分钟(东财) | 东方财富 | P1 |
| `option_current_cffex_em` | 中金所期权实时 | 东方财富 | P1 |

### 3.3 上交所期权 (sse.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `option_sse_list_sina` | 上交所期权列表 | 新浪财经 | P0 |
| `option_sse_expire_day_sina` | 上交所到期日 | 新浪财经 | P1 |
| `option_sse_codes_sina` | 上交所期权代码 | 新浪财经 | P1 |
| `option_sse_spot_price_sina` | 上交所期权实时 | 新浪财经 | P0 |
| `option_sse_underlying_spot_price_sina` | 上交所标的实时 | 新浪财经 | P1 |
| `option_sse_greeks_sina` | 上交所希腊值 | 新浪财经 | P1 |
| `option_sse_minute_sina` | 上交所分钟 | 新浪财经 | P1 |
| `option_sse_daily_sina` | 上交所日K | 新浪财经 | P1 |
| `option_current_day_sse` | 上交所当日期权 | 上交所 | P1 |
| `option_daily_stats_sse` | 上交所每日统计 | 上交所 | P2 |
| `option_risk_indicator_sse` | 上交所风险指标 | 上交所 | P2 |

### 3.4 深交所期权 (szse.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `option_current_day_szse` | 深交所当日期权 | 深交所 | P1 |
| `option_daily_stats_szse` | 深交所每日统计 | 深交所 | P2 |

### 3.5 中金所期权 (cffex.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `option_cffex_hs300_list_sina` | 沪深300期权列表 | 新浪财经 | P1 |
| `option_cffex_hs300_spot_sina` | 沪深300期权实时 | 新浪财经 | P1 |
| `option_cffex_hs300_daily_sina` | 沪深300期权日K | 新浪财经 | P1 |
| `option_cffex_sz50_list_sina` | 上证50期权列表 | 新浪财经 | P1 |
| `option_cffex_sz50_spot_sina` | 上证50期权实时 | 新浪财经 | P1 |
| `option_cffex_sz50_daily_sina` | 上证50期权日K | 新浪财经 | P1 |
| `option_cffex_zz1000_list_sina` | 中证1000期权列表 | 新浪财经 | P1 |
| `option_cffex_zz1000_spot_sina` | 中证1000期权实时 | 新浪财经 | P1 |
| `option_cffex_zz1000_daily_sina` | 中证1000期权日K | 新浪财经 | P1 |

### 3.6 期权分析 (analysis.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `option_risk_analysis_em` | 期权风险分析 | 东方财富 | P2 |
| `option_value_analysis_em` | 期权价值分析 | 东方财富 | P2 |
| `option_premium_analysis_em` | 期权溢价分析 | 东方财富 | P2 |
| `option_margin` | 期权保证金 | - | P2 |
| `option_margin_symbol` | 期权保证金品种 | - | P2 |
| `option_contract_info_ctp` | CTP合约信息 | CTP | P2 |

### 3.7 龙虎榜 (lhb.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `option_lhb_em` | 期权龙虎榜 | 东方财富 | P2 |

### 3.8 工具函数

| 函数名 | 功能 | 优先级 |
|--------|------|--------|
| `convert_date` | 日期转换 | P2 |
| `get_calendar` | 交易日历 | P2 |
| `get_json_path` | JSON路径 | P3 |
| `get_latest_data_date` | 最新数据日期 | P2 |
| `last_trading_day` | 上一交易日 | P2 |

---

## 四、类型定义示例

```go
// types.go
package option

import "time"

// OptionQuote 期权实时行情
type OptionQuote struct {
    Code         string    `json:"code"`           // 期权代码
    Name         string    `json:"name"`           // 期权名称
    Underlying   string    `json:"underlying"`     // 标的代码
    Strike       float64   `json:"strike"`         // 行权价
    Type         string    `json:"type"`           // 类型(C/P)
    Expiry       time.Time `json:"expiry"`         // 到期日
    Price        float64   `json:"price"`          // 现价
    Change       float64   `json:"change"`         // 涨跌
    ChangePct    float64   `json:"change_pct"`     // 涨跌幅
    Volume       int64     `json:"volume"`         // 成交量
    OpenInt      int64     `json:"open_int"`       // 持仓量
    IV           float64   `json:"iv"`             // 隐含波动率
    Time         time.Time `json:"time"`
}

// OptionKLine 期权K线
type OptionKLine struct {
    Date    time.Time `json:"date"`
    Open    float64   `json:"open"`
    High    float64   `json:"high"`
    Low     float64   `json:"low"`
    Close   float64   `json:"close"`
    Settle  float64   `json:"settle"`
    Volume  int64     `json:"volume"`
    OpenInt int64     `json:"open_int"`
}

// OptionGreeks 希腊值
type OptionGreeks struct {
    Code  string  `json:"code"`
    Delta float64 `json:"delta"`
    Gamma float64 `json:"gamma"`
    Theta float64 `json:"theta"`
    Vega  float64 `json:"vega"`
    Rho   float64 `json:"rho"`
    IV    float64 `json:"iv"` // 隐含波动率
}

// OptionChain 期权链
type OptionChain struct {
    Underlying string        `json:"underlying"`
    Expiry     time.Time     `json:"expiry"`
    Calls      []OptionQuote `json:"calls"`
    Puts       []OptionQuote `json:"puts"`
}

// OptionContract 期权合约信息
type OptionContract struct {
    Code           string    `json:"code"`
    Name           string    `json:"name"`
    Underlying     string    `json:"underlying"`
    UnderlyingName string    `json:"underlying_name"`
    Strike         float64   `json:"strike"`
    Type           string    `json:"type"`
    ExerciseType   string    `json:"exercise_type"` // 欧式/美式
    Expiry         time.Time `json:"expiry"`
    Multiplier     int       `json:"multiplier"`    // 合约乘数
    MinChange      float64   `json:"min_change"`    // 最小变动
}
```

---

## 五、实现优先级

### P0 - 第一批实现

1. `option_current_em` - 期权实时(东财)
2. `option_sse_spot_price_sina` - 上交所期权实时
3. `option_commodity_hist_sina` - 商品期权历史
4. `option_finance_board` - 金融期权行情

### P1 - 第二批实现

1. 各交易所期权历史
2. 中金所期权系列
3. 希腊值数据
4. 期权分钟数据

### P2 - 第三批实现

1. 期权分析
2. 龙虎榜
3. 风险指标

---

*文档版本: v1.0*
*最后更新: 2026-01-18*
