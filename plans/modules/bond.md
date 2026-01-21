# bond 模块计划

> 模块: 债券数据
> 函数数量: 40 个
> 优先级: P1
> 预计周期: 第10周

---

## 一、模块概述

债券模块提供可转债、国债、企业债的行情数据、收益率曲线、发行信息等。

---

## 二、Go 文件结构

```
bond/
├── convertible.go       # 可转债
├── treasury.go          # 国债
├── corporate.go         # 企业债
├── yield.go             # 收益率曲线
├── issue.go             # 债券发行
├── repo.go              # 回购
├── index.go             # 债券指数
├── jisilu.go            # 集思录数据
├── chinamoney.go        # 中国货币网
└── types.go             # 类型定义
```

---

## 三、函数清单

### 3.1 可转债 (convertible.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `bond_zh_cov` | 可转债列表 | 东方财富 | P0 |
| `bond_zh_cov_info` | 可转债信息 | 东方财富 | P0 |
| `bond_zh_cov_info_ths` | 可转债信息(同花顺) | 同花顺 | P1 |
| `bond_zh_cov_value_analysis` | 可转债价值分析 | 东方财富 | P1 |
| `bond_zh_hs_cov_spot` | 可转债实时行情 | 东方财富 | P0 |
| `bond_zh_hs_cov_daily` | 可转债日K | 东方财富 | P0 |
| `bond_zh_hs_cov_min` | 可转债分钟K | 东方财富 | P1 |
| `bond_zh_hs_cov_pre_min` | 可转债盘前分钟 | 东方财富 | P2 |
| `bond_cov_comparison` | 可转债对比 | 东方财富 | P2 |
| `bond_cov_issue_cninfo` | 可转债发行(巨潮) | 巨潮资讯 | P2 |
| `bond_cov_stock_issue_cninfo` | 可转债对应股票 | 巨潮资讯 | P2 |

### 3.2 集思录数据 (jisilu.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `bond_cb_jsl` | 集思录可转债 | 集思录 | P0 |
| `bond_cb_index_jsl` | 可转债等权指数 | 集思录 | P1 |
| `bond_cb_redeem_jsl` | 可转债强赎 | 集思录 | P1 |
| `bond_cb_adj_logs_jsl` | 可转债调整记录 | 集思录 | P2 |
| `bond_cb_profile_sina` | 可转债资料(新浪) | 新浪财经 | P2 |
| `bond_cb_summary_sina` | 可转债汇总(新浪) | 新浪财经 | P2 |

### 3.3 国债/收益率 (treasury.go / yield.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `bond_zh_hs_spot` | 债券实时行情 | 东方财富 | P1 |
| `bond_zh_hs_daily` | 债券日K | 东方财富 | P1 |
| `bond_china_yield` | 中国国债收益率 | 中国货币网 | P0 |
| `bond_china_close_return` | 收益率曲线 | 中国货币网 | P1 |
| `bond_china_close_return_map` | 收益率曲线图 | 中国货币网 | P2 |
| `bond_zh_us_rate` | 中美国债收益率 | - | P1 |

### 3.4 债券指数 (index.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `bond_composite_index_cbond` | 中债综合指数 | 中债 | P2 |
| `bond_new_composite_index_cbond` | 新综合指数 | 中债 | P2 |

### 3.5 债券发行 (issue.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `bond_treasure_issue_cninfo` | 国债发行(巨潮) | 巨潮资讯 | P2 |
| `bond_local_government_issue_cninfo` | 地方债发行 | 巨潮资讯 | P2 |
| `bond_corporate_issue_cninfo` | 企业债发行 | 巨潮资讯 | P2 |
| `bond_debt_nafmii` | 债务融资工具 | NAFMII | P2 |

### 3.6 债券回购 (repo.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `bond_buy_back_hist_em` | 债券回购历史 | 东方财富 | P2 |
| `bond_sh_buy_back_em` | 上交所回购 | 东方财富 | P2 |
| `bond_sz_buy_back_em` | 深交所回购 | 东方财富 | P2 |

### 3.7 中国货币网 (chinamoney.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `bond_info_cm` | 债券信息 | 中国货币网 | P1 |
| `bond_info_cm_query` | 债券查询 | 中国货币网 | P1 |
| `bond_info_detail_cm` | 债券详情 | 中国货币网 | P2 |
| `bond_spot_deal` | 现券成交 | 中国货币网 | P2 |
| `bond_spot_quote` | 现券报价 | 中国货币网 | P2 |
| `bond_deal_summary_sse` | 上交所成交汇总 | 上交所 | P2 |
| `bond_cash_summary_sse` | 上交所现金汇总 | 上交所 | P2 |
| `macro_china_bond_public` | 债券公开发行 | - | P2 |
| `macro_china_swap_rate` | 利率互换 | - | P2 |
| `get_zh_bond_hs_page_count` | 债券页数 | - | P3 |

---

## 四、数据源 URL 模式

### 4.1 集思录

| URL 模式 | 用途 |
|----------|------|
| `www.jisilu.cn/data/cbnew/` | 可转债数据 |

### 4.2 中国货币网

| URL 模式 | 用途 |
|----------|------|
| `www.chinamoney.com.cn/chinese/` | 债券信息 |
| `yield.chinabond.com.cn/` | 收益率曲线 |

---

## 五、类型定义示例

```go
// types.go
package bond

import "time"

// ConvertibleBond 可转债
type ConvertibleBond struct {
    Code         string    `json:"code"`          // 转债代码
    Name         string    `json:"name"`          // 转债名称
    StockCode    string    `json:"stock_code"`    // 正股代码
    StockName    string    `json:"stock_name"`    // 正股名称
    Price        float64   `json:"price"`         // 现价
    Change       float64   `json:"change"`        // 涨跌幅
    StockPrice   float64   `json:"stock_price"`   // 正股价
    ConvertPrice float64   `json:"convert_price"` // 转股价
    ConvertValue float64   `json:"convert_value"` // 转股价值
    Premium      float64   `json:"premium"`       // 溢价率
    PB           float64   `json:"pb"`            // 纯债价值
    YTM          float64   `json:"ytm"`           // 到期收益率
    Volume       int64     `json:"volume"`        // 成交量
    MaturityDate time.Time `json:"maturity_date"` // 到期日
}

// BondQuote 债券实时行情
type BondQuote struct {
    Code      string    `json:"code"`
    Name      string    `json:"name"`
    Price     float64   `json:"price"`
    Change    float64   `json:"change"`
    ChangePct float64   `json:"change_pct"`
    Volume    int64     `json:"volume"`
    Amount    float64   `json:"amount"`
    YTM       float64   `json:"ytm"`
    Time      time.Time `json:"time"`
}

// TreasuryYield 国债收益率
type TreasuryYield struct {
    Date   time.Time `json:"date"`
    M1     float64   `json:"1m"`    // 1个月
    M3     float64   `json:"3m"`    // 3个月
    M6     float64   `json:"6m"`    // 6个月
    Y1     float64   `json:"1y"`    // 1年
    Y2     float64   `json:"2y"`    // 2年
    Y3     float64   `json:"3y"`    // 3年
    Y5     float64   `json:"5y"`    // 5年
    Y7     float64   `json:"7y"`    // 7年
    Y10    float64   `json:"10y"`   // 10年
    Y20    float64   `json:"20y"`   // 20年
    Y30    float64   `json:"30y"`   // 30年
}

// BondIssue 债券发行
type BondIssue struct {
    Code        string    `json:"code"`
    Name        string    `json:"name"`
    Issuer      string    `json:"issuer"`       // 发行人
    IssueDate   time.Time `json:"issue_date"`   // 发行日期
    IssueSize   float64   `json:"issue_size"`   // 发行规模
    Coupon      float64   `json:"coupon"`       // 票面利率
    Term        string    `json:"term"`         // 期限
    Rating      string    `json:"rating"`       // 评级
}
```

---

## 六、实现优先级

### P0 - 第一批实现

1. `bond_cb_jsl` - 集思录可转债
2. `bond_zh_cov` - 可转债列表
3. `bond_zh_hs_cov_spot` - 可转债实时
4. `bond_zh_hs_cov_daily` - 可转债日K
5. `bond_china_yield` - 国债收益率

### P1 - 第二批实现

1. 可转债指数
2. 债券信息查询
3. 中美国债对比

### P2 - 第三批实现

1. 债券发行
2. 债券回购
3. 债券指数

---

*文档版本: v1.0*
*最后更新: 2026-01-18*
