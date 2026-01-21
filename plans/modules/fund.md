# fund 模块计划

> 模块: 基金数据
> 函数数量: 82 个
> 优先级: P0
> 预计周期: 第6-7周

---

## 一、模块概述

基金模块提供 ETF、LOF、开放式基金、货币基金等各类基金的实时行情、历史净值、持仓信息、基金经理等数据。

---

## 二、Go 文件结构

```
fund/
├── etf.go              # ETF 基金
├── lof.go              # LOF 基金
├── open.go             # 开放式基金
├── money.go            # 货币基金
├── graded.go           # 分级基金
├── financial.go        # 理财型基金
├── hk.go               # 香港基金
├── manager.go          # 基金经理
├── portfolio.go        # 基金持仓
├── rating.go           # 基金评级
├── announcement.go     # 基金公告
├── amac.go             # 基金业协会
├── scale.go            # 基金规模
├── xueqiu.go           # 雪球数据
└── types.go            # 类型定义
```

---

## 三、函数清单

### 3.1 ETF 基金 (etf.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `fund_etf_spot_em` | ETF实时行情 | 东方财富 | P0 |
| `fund_etf_hist_em` | ETF历史行情 | 东方财富 | P0 |
| `fund_etf_hist_min_em` | ETF分钟数据 | 东方财富 | P1 |
| `fund_etf_fund_daily_em` | ETF日度数据 | 东方财富 | P1 |
| `fund_etf_fund_info_em` | ETF基金信息 | 东方财富 | P1 |
| `fund_etf_spot_ths` | ETF实时行情(同花顺) | 同花顺 | P2 |
| `fund_etf_hist_sina` | ETF历史数据(新浪) | 新浪财经 | P2 |
| `fund_etf_category_sina` | ETF分类 | 新浪财经 | P2 |
| `fund_etf_dividend_sina` | ETF分红 | 新浪财经 | P2 |

### 3.2 LOF 基金 (lof.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `fund_lof_spot_em` | LOF实时行情 | 东方财富 | P0 |
| `fund_lof_hist_em` | LOF历史行情 | 东方财富 | P0 |
| `fund_lof_hist_min_em` | LOF分钟数据 | 东方财富 | P1 |

### 3.3 开放式基金 (open.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `fund_open_fund_daily_em` | 开放式基金每日 | 东方财富 | P0 |
| `fund_open_fund_info_em` | 开放式基金信息 | 东方财富 | P0 |
| `fund_open_fund_rank_em` | 开放式基金排行 | 东方财富 | P0 |
| `fund_name_em` | 基金名称列表 | 东方财富 | P0 |
| `fund_info_index_em` | 指数基金信息 | 东方财富 | P1 |
| `fund_value_estimation_em` | 基金估值 | 东方财富 | P1 |
| `fund_purchase_em` | 基金申购状态 | 东方财富 | P2 |
| `fund_new_found_em` | 新发基金 | 东方财富 | P2 |
| `fund_cf_em` | 成立以来基金 | 东方财富 | P2 |

### 3.4 货币基金 (money.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `fund_money_fund_daily_em` | 货币基金每日 | 东方财富 | P1 |
| `fund_money_fund_info_em` | 货币基金信息 | 东方财富 | P1 |
| `fund_money_rank_em` | 货币基金排行 | 东方财富 | P1 |

### 3.5 分级基金 (graded.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `fund_graded_fund_daily_em` | 分级基金每日 | 东方财富 | P2 |
| `fund_graded_fund_info_em` | 分级基金信息 | 东方财富 | P2 |

### 3.6 理财型基金 (financial.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `fund_financial_fund_daily_em` | 理财基金每日 | 东方财富 | P2 |
| `fund_financial_fund_info_em` | 理财基金信息 | 东方财富 | P2 |

### 3.7 香港基金 (hk.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `fund_hk_fund_hist_em` | 港股基金历史 | 东方财富 | P2 |
| `fund_hk_rank_em` | 港股基金排行 | 东方财富 | P2 |

### 3.8 基金经理 (manager.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `fund_manager_em` | 基金经理信息 | 东方财富 | P1 |

### 3.9 基金持仓 (portfolio.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `fund_portfolio_hold_em` | 基金持仓 | 东方财富 | P0 |
| `fund_portfolio_bond_hold_em` | 基金债券持仓 | 东方财富 | P1 |
| `fund_portfolio_change_em` | 基金持仓变动 | 东方财富 | P1 |
| `fund_portfolio_industry_allocation_em` | 行业配置 | 东方财富 | P1 |
| `fund_hold_structure_em` | 持有人结构 | 东方财富 | P2 |

### 3.10 基金评级 (rating.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `fund_rating_all` | 综合评级 | 天天基金 | P2 |
| `fund_rating_sh` | 上海证券评级 | 天天基金 | P2 |
| `fund_rating_zs` | 招商证券评级 | 天天基金 | P2 |
| `fund_rating_ja` | 济安金信评级 | 天天基金 | P2 |

### 3.11 基金公告 (announcement.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `fund_announcement_dividend_em` | 分红公告 | 东方财富 | P2 |
| `fund_announcement_personnel_em` | 人事公告 | 东方财富 | P2 |
| `fund_announcement_report_em` | 定期报告 | 东方财富 | P2 |

### 3.12 基金业协会 (amac.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `amac_fund_info` | 私募基金信息 | AMAC | P2 |
| `amac_fund_sub_info` | 私募子基金信息 | AMAC | P2 |
| `amac_fund_abs` | 资产支持专项 | AMAC | P2 |
| `amac_fund_account_info` | 基金账户信息 | AMAC | P2 |
| `amac_manager_info` | 管理人信息 | AMAC | P2 |
| `amac_manager_classify_info` | 管理人分类 | AMAC | P2 |
| `amac_manager_cancelled_info` | 注销管理人 | AMAC | P2 |
| `amac_member_info` | 会员信息 | AMAC | P2 |
| `amac_member_sub_info` | 会员子信息 | AMAC | P2 |
| `amac_aoin_info` | 境外投资顾问 | AMAC | P3 |
| `amac_futures_info` | 期货公司信息 | AMAC | P2 |
| `amac_securities_info` | 证券公司信息 | AMAC | P2 |
| `amac_person_bond_org_list` | 债券组织 | AMAC | P3 |
| `amac_person_fund_org_list` | 基金组织 | AMAC | P3 |

### 3.13 基金规模 (scale.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `fund_scale_open_sina` | 开放式规模 | 新浪财经 | P2 |
| `fund_scale_close_sina` | 封闭式规模 | 新浪财经 | P2 |
| `fund_scale_structured_sina` | 分级基金规模 | 新浪财经 | P2 |
| `fund_scale_change_em` | 规模变动 | 东方财富 | P2 |
| `fund_aum_em` | 资产规模 | 东方财富 | P1 |
| `fund_aum_hist_em` | 资产规模历史 | 东方财富 | P2 |
| `fund_aum_trend_em` | 资产规模趋势 | 东方财富 | P2 |
| `fund_overview_em` | 基金概览 | 东方财富 | P1 |

### 3.14 雪球数据 (xueqiu.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `fund_individual_basic_info_xq` | 基金基本信息 | 雪球 | P2 |
| `fund_individual_detail_info_xq` | 基金详细信息 | 雪球 | P2 |
| `fund_individual_detail_hold_xq` | 基金持仓详情 | 雪球 | P2 |
| `fund_individual_achievement_xq` | 基金业绩 | 雪球 | P2 |
| `fund_individual_analysis_xq` | 基金分析 | 雪球 | P2 |
| `fund_individual_profit_probability_xq` | 盈利概率 | 雪球 | P2 |

### 3.15 其他基金功能

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `fund_fee_em` | 基金费率 | 东方财富 | P2 |
| `fund_fh_em` | 基金分红 | 东方财富 | P1 |
| `fund_fh_rank_em` | 基金分红排行 | 东方财富 | P2 |
| `fund_exchange_rank_em` | 交易所基金排行 | 东方财富 | P2 |
| `fund_lcx_rank_em` | 理财通排行 | 东方财富 | P3 |
| `fund_balance_position_lg` | 持仓(乐股) | 乐股网 | P2 |
| `fund_stock_position_lg` | 股票持仓(乐股) | 乐股网 | P2 |
| `fund_linghuo_position_lg` | 灵活持仓(乐股) | 乐股网 | P2 |
| `fund_report_asset_allocation_cninfo` | 资产配置(巨潮) | 巨潮资讯 | P2 |
| `fund_report_industry_allocation_cninfo` | 行业配置(巨潮) | 巨潮资讯 | P2 |
| `fund_report_stock_cninfo` | 持股(巨潮) | 巨潮资讯 | P2 |

---

## 四、数据源 URL 模式

### 4.1 东方财富

| URL 模式 | 用途 |
|----------|------|
| `fund.eastmoney.com/Data/Fund_JJJZ_Data.aspx` | 基金净值 |
| `api.fund.eastmoney.com/f10/lsjz` | 历史净值 |
| `fundf10.eastmoney.com/F10DataApi.aspx` | 基金详情 |
| `push2.eastmoney.com/api/qt/clist/get` | ETF列表 |

### 4.2 AMAC

| URL 模式 | 用途 |
|----------|------|
| `gs.amac.org.cn/amac-infodisc/api/pof/` | 私募基金 |
| `gs.amac.org.cn/amac-infodisc/api/aoin/` | 境外顾问 |

---

## 五、类型定义示例

```go
// types.go
package fund

import "time"

// FundInfo 基金基本信息
type FundInfo struct {
    Code       string    `json:"code"`        // 基金代码
    Name       string    `json:"name"`        // 基金名称
    Type       string    `json:"type"`        // 基金类型
    Manager    string    `json:"manager"`     // 基金经理
    Company    string    `json:"company"`     // 基金公司
    SetupDate  time.Time `json:"setup_date"`  // 成立日期
    Scale      float64   `json:"scale"`       // 基金规模（亿）
    FeeRate    float64   `json:"fee_rate"`    // 管理费率
}

// FundNav 基金净值
type FundNav struct {
    Date         time.Time `json:"date"`          // 日期
    Nav          float64   `json:"nav"`           // 单位净值
    AccNav       float64   `json:"acc_nav"`       // 累计净值
    DailyReturn  float64   `json:"daily_return"`  // 日收益率
    Status       string    `json:"status"`        // 申购状态
}

// ETFQuote ETF实时行情
type ETFQuote struct {
    Code      string    `json:"code"`
    Name      string    `json:"name"`
    Price     float64   `json:"price"`
    Change    float64   `json:"change"`
    ChangePct float64   `json:"change_pct"`
    Volume    int64     `json:"volume"`
    Amount    float64   `json:"amount"`
    Nav       float64   `json:"nav"`       // 净值
    Premium   float64   `json:"premium"`   // 溢价率
    Time      time.Time `json:"time"`
}

// FundHolding 基金持仓
type FundHolding struct {
    ReportDate time.Time `json:"report_date"` // 报告日期
    StockCode  string    `json:"stock_code"`  // 股票代码
    StockName  string    `json:"stock_name"`  // 股票名称
    Shares     int64     `json:"shares"`      // 持股数量
    Value      float64   `json:"value"`       // 持仓市值
    Ratio      float64   `json:"ratio"`       // 占净值比
}

// FundManager 基金经理
type FundManager struct {
    ID          string    `json:"id"`
    Name        string    `json:"name"`
    Company     string    `json:"company"`
    StartDate   time.Time `json:"start_date"`   // 从业日期
    FundCount   int       `json:"fund_count"`   // 管理基金数
    TotalScale  float64   `json:"total_scale"`  // 管理规模
    BestReturn  float64   `json:"best_return"`  // 最佳回报
}
```

---

## 六、实现优先级

### P0 - 第一批实现

1. `fund_etf_spot_em` - ETF实时行情
2. `fund_etf_hist_em` - ETF历史行情
3. `fund_lof_spot_em` - LOF实时行情
4. `fund_open_fund_daily_em` - 开放式基金
5. `fund_name_em` - 基金名称列表
6. `fund_portfolio_hold_em` - 基金持仓

### P1 - 第二批实现

1. 货币基金数据
2. 基金经理信息
3. 基金分红
4. 基金规模

### P2 - 第三批实现

1. AMAC 协会数据
2. 基金评级
3. 基金公告
4. 雪球数据

---

*文档版本: v1.0*
*最后更新: 2026-01-18*
