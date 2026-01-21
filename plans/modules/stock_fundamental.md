# stock_fundamental 模块计划

> 模块: 股票基本面数据
> 函数数量: 57 个
> 优先级: P1
> 预计周期: 第4-5周

---

## 一、模块概述

股票基本面模块提供财务报表、财务分析指标、分红历史、股本结构、机构持股等基本面数据。

---

## 二、Go 文件结构

```
stock_fundamental/
├── financial_report.go    # 财务报表
├── financial_analysis.go  # 财务分析
├── dividend.go            # 分红历史
├── share_structure.go     # 股本结构
├── institutional.go       # 机构持股
├── ipo.go                 # IPO信息
├── forecast.go            # 业绩预测
├── restricted.go          # 限售解禁
├── register.go            # 注册制
└── types.go               # 类型定义
```

---

## 三、函数清单

### 3.1 财务报表 (financial_report.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `stock_financial_report_sina` | 财务报表(新浪) | 新浪财经 | P0 |
| `stock_balance_sheet_by_report_em` | 资产负债表 | 东方财富 | P0 |
| `stock_balance_sheet_by_yearly_em` | 资产负债表年度 | 东方财富 | P0 |
| `stock_balance_sheet_by_report_delisted_em` | 退市公司资产负债表 | 东方财富 | P2 |
| `stock_profit_sheet_by_report_em` | 利润表 | 东方财富 | P0 |
| `stock_profit_sheet_by_quarterly_em` | 利润表季度 | 东方财富 | P0 |
| `stock_profit_sheet_by_yearly_em` | 利润表年度 | 东方财富 | P0 |
| `stock_profit_sheet_by_report_delisted_em` | 退市公司利润表 | 东方财富 | P2 |
| `stock_cash_flow_sheet_by_report_em` | 现金流量表 | 东方财富 | P0 |
| `stock_cash_flow_sheet_by_quarterly_em` | 现金流量表季度 | 东方财富 | P0 |
| `stock_cash_flow_sheet_by_yearly_em` | 现金流量表年度 | 东方财富 | P0 |
| `stock_cash_flow_sheet_by_report_delisted_em` | 退市公司现金流量表 | 东方财富 | P2 |
| `stock_lrb_em` | 利润表(东财) | 东方财富 | P1 |
| `stock_zcfz_em` | 资产负债表(东财) | 东方财富 | P1 |
| `stock_zcfz_bj_em` | 北交所资产负债表 | 东方财富 | P2 |
| `stock_xjll_em` | 现金流量表(东财) | 东方财富 | P1 |

### 3.2 财务分析 (financial_analysis.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `stock_financial_abstract` | 财务摘要 | 新浪财经 | P0 |
| `stock_financial_abstract_ths` | 财务摘要(同花顺) | 同花顺 | P1 |
| `stock_financial_abstract_new_ths` | 财务摘要新版(同花顺) | 同花顺 | P1 |
| `stock_financial_analysis_indicator` | 财务分析指标 | 新浪财经 | P0 |
| `stock_financial_analysis_indicator_em` | 财务分析指标(东财) | 东方财富 | P0 |
| `stock_financial_hk_analysis_indicator_em` | 港股财务指标 | 东方财富 | P1 |
| `stock_financial_hk_report_em` | 港股财报 | 东方财富 | P1 |
| `stock_financial_us_analysis_indicator_em` | 美股财务指标 | 东方财富 | P1 |
| `stock_financial_us_report_em` | 美股财报 | 东方财富 | P1 |
| `stock_financial_benefit_ths` | 盈利能力(同花顺) | 同花顺 | P1 |
| `stock_financial_benefit_new_ths` | 盈利能力新版 | 同花顺 | P1 |
| `stock_financial_cash_ths` | 现金流(同花顺) | 同花顺 | P1 |
| `stock_financial_cash_new_ths` | 现金流新版 | 同花顺 | P1 |
| `stock_financial_debt_ths` | 偿债能力(同花顺) | 同花顺 | P1 |
| `stock_financial_debt_new_ths` | 偿债能力新版 | 同花顺 | P1 |
| `stock_zygc_em` | 主营构成 | 东方财富 | P1 |
| `stock_zyjs_ths` | 主营解释(同花顺) | 同花顺 | P2 |

### 3.3 分红历史 (dividend.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `stock_history_dividend` | 历史分红 | 新浪财经 | P0 |
| `stock_history_dividend_detail` | 历史分红详情 | 新浪财经 | P0 |

### 3.4 股本结构 (share_structure.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `stock_zh_a_gbjg_em` | 股本结构 | 东方财富 | P0 |
| `stock_add_stock` | 增发信息 | - | P2 |

### 3.5 机构持股 (institutional.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `stock_institute_hold` | 机构持股 | 新浪财经 | P0 |
| `stock_institute_hold_detail` | 机构持股详情 | 新浪财经 | P1 |
| `stock_institute_recommend` | 机构推荐 | 新浪财经 | P1 |
| `stock_institute_recommend_detail` | 机构推荐详情 | 新浪财经 | P1 |
| `stock_fund_stock_holder` | 基金持股 | 新浪财经 | P0 |
| `stock_main_stock_holder` | 主要股东 | 新浪财经 | P0 |
| `stock_circulate_stock_holder` | 流通股东 | 新浪财经 | P0 |

### 3.6 IPO 信息 (ipo.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `stock_ipo_info` | IPO信息 | 新浪财经 | P1 |
| `stock_ipo_declare_em` | IPO申报 | 东方财富 | P1 |
| `stock_ipo_review_em` | IPO审核 | 东方财富 | P1 |
| `stock_ipo_tutor_em` | IPO辅导 | 东方财富 | P2 |
| `stock_ipo_benefit_ths` | IPO受益股 | 同花顺 | P2 |

### 3.7 业绩预测 (forecast.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `stock_profit_forecast_em` | 业绩预测(东财) | 东方财富 | P0 |
| `stock_profit_forecast_ths` | 业绩预测(同花顺) | 同花顺 | P1 |
| `stock_hk_profit_forecast_et` | 港股业绩预测 | - | P2 |

### 3.8 限售解禁 (restricted.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `stock_restricted_release_queue_em` | 解禁计划 | 东方财富 | P0 |
| `stock_restricted_release_summary_em` | 解禁汇总 | 东方财富 | P1 |
| `stock_restricted_release_detail_em` | 解禁详情 | 东方财富 | P1 |
| `stock_restricted_release_stockholder_em` | 解禁股东 | 东方财富 | P2 |
| `stock_restricted_release_queue_sina` | 解禁计划(新浪) | 新浪财经 | P2 |

### 3.9 注册制 (register.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `stock_register_all_em` | 全部注册制 | 东方财富 | P1 |
| `stock_register_kcb` | 科创板注册 | 东方财富 | P1 |
| `stock_register_cyb` | 创业板注册 | 东方财富 | P1 |
| `stock_register_bj` | 北交所注册 | 东方财富 | P1 |
| `stock_register_db` | 主板注册 | 东方财富 | P1 |
| `stock_register_sh` | 沪市注册 | 东方财富 | P2 |
| `stock_register_sz` | 深市注册 | 东方财富 | P2 |
| `stock_kcb_renewal` | 科创板续审 | 东方财富 | P2 |
| `stock_kcb_detail_renewal` | 科创板续审详情 | 东方财富 | P2 |

### 3.10 其他功能

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `stock_notice_report` | 公告报告 | 新浪财经 | P2 |
| `stock_shareholder_change_ths` | 股东变动 | 同花顺 | P1 |
| `stock_management_change_ths` | 管理层变动 | 同花顺 | P2 |
| `stock_individual_basic_info_xq` | 个股基本信息(雪球) | 雪球 | P2 |
| `stock_individual_basic_info_hk_xq` | 港股基本信息(雪球) | 雪球 | P2 |
| `stock_individual_basic_info_us_xq` | 美股基本信息(雪球) | 雪球 | P2 |

---

## 四、类型定义示例

```go
// types.go
package stock_fundamental

import "time"

// BalanceSheet 资产负债表
type BalanceSheet struct {
    ReportDate        time.Time `json:"report_date"`
    TotalAssets       float64   `json:"total_assets"`        // 总资产
    TotalLiabilities  float64   `json:"total_liabilities"`   // 总负债
    TotalEquity       float64   `json:"total_equity"`        // 股东权益
    Cash              float64   `json:"cash"`                // 货币资金
    Receivables       float64   `json:"receivables"`         // 应收账款
    Inventory         float64   `json:"inventory"`           // 存货
    FixedAssets       float64   `json:"fixed_assets"`        // 固定资产
    ShortTermDebt     float64   `json:"short_term_debt"`     // 短期借款
    LongTermDebt      float64   `json:"long_term_debt"`      // 长期借款
}

// IncomeStatement 利润表
type IncomeStatement struct {
    ReportDate      time.Time `json:"report_date"`
    Revenue         float64   `json:"revenue"`          // 营业收入
    OperatingCost   float64   `json:"operating_cost"`   // 营业成本
    GrossProfit     float64   `json:"gross_profit"`     // 毛利润
    OperatingProfit float64   `json:"operating_profit"` // 营业利润
    NetProfit       float64   `json:"net_profit"`       // 净利润
    EPS             float64   `json:"eps"`              // 每股收益
}

// CashFlowStatement 现金流量表
type CashFlowStatement struct {
    ReportDate       time.Time `json:"report_date"`
    OperatingCF      float64   `json:"operating_cf"`      // 经营活动现金流
    InvestingCF      float64   `json:"investing_cf"`      // 投资活动现金流
    FinancingCF      float64   `json:"financing_cf"`      // 筹资活动现金流
    NetCashFlow      float64   `json:"net_cash_flow"`     // 净现金流
    EndingCash       float64   `json:"ending_cash"`       // 期末现金
}

// FinancialIndicator 财务指标
type FinancialIndicator struct {
    ReportDate      time.Time `json:"report_date"`
    ROE             float64   `json:"roe"`               // 净资产收益率
    ROA             float64   `json:"roa"`               // 总资产收益率
    GrossProfitRate float64   `json:"gross_profit_rate"` // 毛利率
    NetProfitRate   float64   `json:"net_profit_rate"`   // 净利率
    DebtRatio       float64   `json:"debt_ratio"`        // 资产负债率
    CurrentRatio    float64   `json:"current_ratio"`     // 流动比率
    QuickRatio      float64   `json:"quick_ratio"`       // 速动比率
}

// DividendHistory 分红历史
type DividendHistory struct {
    Year          int     `json:"year"`
    ReportDate    string  `json:"report_date"`
    DividendRatio float64 `json:"dividend_ratio"` // 每股派息
    BonusRatio    float64 `json:"bonus_ratio"`    // 送股比例
    TransferRatio float64 `json:"transfer_ratio"` // 转增比例
    ExDivDate     string  `json:"ex_div_date"`    // 除权日
    PayDate       string  `json:"pay_date"`       // 派息日
}

// RestrictedRelease 限售解禁
type RestrictedRelease struct {
    Date          time.Time `json:"date"`
    Code          string    `json:"code"`
    Name          string    `json:"name"`
    ShareCount    int64     `json:"share_count"`    // 解禁股数
    ShareRatio    float64   `json:"share_ratio"`    // 占总股本比
    LockupType    string    `json:"lockup_type"`    // 限售类型
    Shareholders  []string  `json:"shareholders"`   // 解禁股东
}
```

---

## 五、实现优先级

### P0 - 第一批实现

1. 三大财务报表
2. 财务分析指标
3. 分红历史
4. 股本结构
5. 机构持股

### P1 - 第二批实现

1. 同花顺财务数据
2. 业绩预测
3. 限售解禁
4. 注册制信息

### P2 - 第三批实现

1. 退市公司财报
2. IPO 辅导
3. 雪球数据

---

*文档版本: v1.0*
*最后更新: 2026-01-18*
