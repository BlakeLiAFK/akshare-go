# stock_feature 模块计划

> 模块: 股票特色数据
> 函数数量: 197 个
> 优先级: P1
> 预计周期: 第3-5周（与 stock 模块并行）

---

## 一、模块概述

股票特色数据模块提供融资融券、龙虎榜、股东研究、机构调研、业绩预告、涨跌停池等特色数据。

---

## 二、Go 文件结构

```
stock_feature/
├── margin.go              # 融资融券
├── lhb.go                 # 龙虎榜
├── hsgt.go                # 沪深港通
├── holder.go              # 股东研究
├── institution.go         # 机构调研
├── performance.go         # 业绩数据
├── zt_pool.go             # 涨跌停池
├── board_ths.go           # 同花顺板块
├── valuation.go           # 估值数据
├── comment.go             # 股票点评
├── esg.go                 # ESG数据
├── disclosure.go          # 信息披露
├── rank.go                # 技术排名
└── types.go               # 类型定义
```

---

## 三、函数清单

### 3.1 融资融券 (margin.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `stock_margin_sse` | 上交所融资融券 | 上交所 | P0 |
| `stock_margin_szse` | 深交所融资融券 | 深交所 | P0 |
| `stock_margin_detail_sse` | 上交所融资融券明细 | 上交所 | P1 |
| `stock_margin_detail_szse` | 深交所融资融券明细 | 深交所 | P1 |
| `stock_margin_underlying_info_szse` | 融资融券标的 | 深交所 | P2 |
| `stock_margin_account_info` | 融资融券账户 | - | P2 |
| `stock_margin_ratio_pa` | 融资融券比例 | 平安证券 | P2 |

### 3.2 龙虎榜 (lhb.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `stock_lhb_detail_em` | 龙虎榜详情 | 东方财富 | P0 |
| `stock_lhb_stock_detail_em` | 个股龙虎榜详情 | 东方财富 | P0 |
| `stock_lhb_stock_detail_date_em` | 个股龙虎榜日期 | 东方财富 | P1 |
| `stock_lhb_stock_statistic_em` | 个股龙虎榜统计 | 东方财富 | P1 |
| `stock_lhb_jgmx_sina` | 机构买卖明细(新浪) | 新浪财经 | P1 |
| `stock_lhb_jgzz_sina` | 机构追踪(新浪) | 新浪财经 | P2 |
| `stock_lhb_jgstatistic_em` | 机构统计 | 东方财富 | P1 |
| `stock_lhb_jgmmtj_em` | 机构买卖统计 | 东方财富 | P1 |
| `stock_lhb_hyyyb_em` | 活跃营业部 | 东方财富 | P1 |
| `stock_lhb_yybph_em` | 营业部排行 | 东方财富 | P1 |
| `stock_lhb_yyb_detail_em` | 营业部详情 | 东方财富 | P1 |
| `stock_lhb_traderstatistic_em` | 交易商统计 | 东方财富 | P2 |
| `stock_lhb_detail_daily_sina` | 每日龙虎榜(新浪) | 新浪财经 | P1 |
| `stock_lhb_ggtj_sina` | 个股统计(新浪) | 新浪财经 | P2 |
| `stock_lhb_yytj_sina` | 营业部统计(新浪) | 新浪财经 | P2 |
| `stock_lh_yyb_most` | 最活跃营业部 | - | P2 |
| `stock_lh_yyb_capital` | 营业部资金 | - | P2 |
| `stock_lh_yyb_control` | 营业部控盘 | - | P2 |

### 3.3 沪深港通 (hsgt.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `stock_hsgt_hist_em` | 沪深港通历史 | 东方财富 | P0 |
| `stock_hsgt_board_rank_em` | 沪深港通板块排名 | 东方财富 | P1 |
| `stock_hsgt_fund_flow_summary_em` | 资金流向汇总 | 东方财富 | P0 |
| `stock_hsgt_fund_min_em` | 资金分钟 | 东方财富 | P1 |
| `stock_hsgt_hold_stock_em` | 持股数据 | 东方财富 | P0 |
| `stock_hsgt_stock_statistics_em` | 股票统计 | 东方财富 | P1 |
| `stock_hsgt_individual_em` | 个股数据 | 东方财富 | P1 |
| `stock_hsgt_individual_detail_em` | 个股详情 | 东方财富 | P1 |
| `stock_hsgt_institution_statistics_em` | 机构统计 | 东方财富 | P2 |
| `stock_hk_ggt_components_em` | 港股通成分 | 东方财富 | P1 |
| `stock_sgt_reference_exchange_rate_sse` | 沪市参考汇率 | 上交所 | P2 |
| `stock_sgt_reference_exchange_rate_szse` | 深市参考汇率 | 深交所 | P2 |
| `stock_sgt_settlement_exchange_rate_sse` | 沪市结算汇率 | 上交所 | P2 |
| `stock_sgt_settlement_exchange_rate_szse` | 深市结算汇率 | 深交所 | P2 |

### 3.4 股东研究 (holder.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `stock_gdfx_top_10_em` | 十大股东 | 东方财富 | P0 |
| `stock_gdfx_free_top_10_em` | 十大流通股东 | 东方财富 | P0 |
| `stock_gdfx_holding_detail_em` | 持股详情 | 东方财富 | P1 |
| `stock_gdfx_holding_change_em` | 持股变动 | 东方财富 | P1 |
| `stock_gdfx_holding_analyse_em` | 持股分析 | 东方财富 | P1 |
| `stock_gdfx_holding_statistics_em` | 持股统计 | 东方财富 | P1 |
| `stock_gdfx_holding_teamwork_em` | 抱团股 | 东方财富 | P2 |
| `stock_gdfx_free_holding_detail_em` | 流通股持股详情 | 东方财富 | P1 |
| `stock_gdfx_free_holding_change_em` | 流通股持股变动 | 东方财富 | P1 |
| `stock_gdfx_free_holding_analyse_em` | 流通股持股分析 | 东方财富 | P1 |
| `stock_gdfx_free_holding_statistics_em` | 流通股持股统计 | 东方财富 | P1 |
| `stock_gdfx_free_holding_teamwork_em` | 流通股抱团 | 东方财富 | P2 |
| `stock_zh_a_gdhs` | 股东户数 | 东方财富 | P0 |
| `stock_zh_a_gdhs_detail_em` | 股东户数详情 | 东方财富 | P1 |
| `stock_gddh_em` | 股东大会 | 东方财富 | P2 |

### 3.5 机构调研 (institution.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `stock_jgdy_tj_em` | 机构调研统计 | 东方财富 | P1 |
| `stock_jgdy_detail_em` | 机构调研详情 | 东方财富 | P1 |
| `stock_analyst_rank_em` | 分析师排名 | 东方财富 | P1 |
| `stock_analyst_detail_em` | 分析师详情 | 东方财富 | P2 |
| `stock_research_report_em` | 研究报告 | 东方财富 | P2 |

### 3.6 业绩数据 (performance.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `stock_yjyg_em` | 业绩预告 | 东方财富 | P0 |
| `stock_yjkb_em` | 业绩快报 | 东方财富 | P0 |
| `stock_yjbb_em` | 业绩报表 | 东方财富 | P0 |
| `stock_yysj_em` | 预约披露 | 东方财富 | P1 |
| `stock_dxsyl_em` | 打新收益率 | 东方财富 | P2 |
| `stock_fhps_em` | 分红配送 | 东方财富 | P0 |
| `stock_fhps_detail_em` | 分红配送详情 | 东方财富 | P1 |
| `stock_fhps_detail_ths` | 分红详情(同花顺) | 同花顺 | P2 |
| `stock_ggcg_em` | 高管持股 | 东方财富 | P1 |
| `stock_pg_em` | 配股 | 东方财富 | P2 |
| `stock_qbzf_em` | 定向增发 | 东方财富 | P2 |
| `stock_qsjy_em` | 期权激励 | 东方财富 | P2 |

### 3.7 涨跌停池 (zt_pool.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `stock_zt_pool_em` | 涨停池 | 东方财富 | P0 |
| `stock_zt_pool_previous_em` | 昨日涨停 | 东方财富 | P0 |
| `stock_zt_pool_strong_em` | 强势股池 | 东方财富 | P1 |
| `stock_zt_pool_sub_new_em` | 次新股池 | 东方财富 | P1 |
| `stock_zt_pool_dtgc_em` | 跌停股池 | 东方财富 | P0 |
| `stock_zt_pool_zbgc_em` | 炸板股池 | 东方财富 | P1 |

### 3.8 同花顺板块 (board_ths.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `stock_board_concept_name_ths` | 概念板块名称 | 同花顺 | P1 |
| `stock_board_concept_index_ths` | 概念板块指数 | 同花顺 | P1 |
| `stock_board_concept_info_ths` | 概念板块信息 | 同花顺 | P1 |
| `stock_board_concept_summary_ths` | 概念板块汇总 | 同花顺 | P1 |
| `stock_board_industry_name_ths` | 行业板块名称 | 同花顺 | P1 |
| `stock_board_industry_index_ths` | 行业板块指数 | 同花顺 | P1 |
| `stock_board_industry_info_ths` | 行业板块信息 | 同花顺 | P1 |
| `stock_board_industry_summary_ths` | 行业板块汇总 | 同花顺 | P1 |
| `stock_board_change_em` | 板块变动 | 东方财富 | P2 |

### 3.9 估值数据 (valuation.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `stock_a_ttm_lyr` | A股TTM市盈率 | - | P1 |
| `stock_a_all_pb` | A股市净率 | - | P1 |
| `stock_a_below_net_asset_statistics` | 破净股统计 | - | P2 |
| `stock_a_high_low_statistics` | 创新高低统计 | - | P2 |
| `stock_a_congestion_lg` | 筹码分布 | 乐股网 | P2 |
| `stock_a_gxl_lg` | 股息率 | 乐股网 | P1 |
| `stock_hk_gxl_lg` | 港股股息率 | 乐股网 | P2 |
| `stock_buffett_index_lg` | 巴菲特指标 | 乐股网 | P2 |
| `stock_ebs_lg` | EBS估值 | 乐股网 | P2 |
| `stock_market_pe_lg` | 市场市盈率 | 乐股网 | P1 |
| `stock_market_pb_lg` | 市场市净率 | 乐股网 | P1 |
| `stock_index_pe_lg` | 指数市盈率 | 乐股网 | P1 |
| `stock_index_pb_lg` | 指数市净率 | 乐股网 | P1 |
| `stock_market_activity_legu` | 市场活跃度 | 乐股网 | P2 |
| `stock_zh_valuation_baidu` | A股估值(百度) | 百度 | P2 |
| `stock_hk_valuation_baidu` | 港股估值(百度) | 百度 | P2 |
| `stock_us_valuation_baidu` | 美股估值(百度) | 百度 | P2 |
| `stock_zh_vote_baidu` | 股票投票(百度) | 百度 | P3 |
| `stock_value_em` | 估值(东财) | 东方财富 | P1 |
| `stock_cyq_em` | 筹码(东财) | 东方财富 | P2 |

### 3.10 股票点评 (comment.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `stock_comment_em` | 股票点评 | 东方财富 | P2 |
| `stock_comment_detail_scrd_desire_em` | 市场愿望 | 东方财富 | P2 |
| `stock_comment_detail_scrd_focus_em` | 关注焦点 | 东方财富 | P2 |
| `stock_comment_detail_zhpj_lspf_em` | 历史评分 | 东方财富 | P2 |
| `stock_comment_detail_zlkp_jgcyd_em` | 机构参与度 | 东方财富 | P2 |

### 3.11 ESG数据 (esg.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `stock_esg_rate_sina` | ESG评级 | 新浪财经 | P2 |
| `stock_esg_hz_sina` | ESG评级汇总 | 新浪财经 | P2 |
| `stock_esg_zd_sina` | ESG评级综合 | 新浪财经 | P2 |
| `stock_esg_msci_sina` | MSCI ESG | 新浪财经 | P2 |
| `stock_esg_rft_sina` | ESG研报 | 新浪财经 | P2 |

### 3.12 信息披露 (disclosure.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `stock_zh_a_disclosure_report_cninfo` | 定期报告 | 巨潮资讯 | P1 |
| `stock_zh_a_disclosure_relation_cninfo` | 关联交易 | 巨潮资讯 | P2 |
| `stock_report_disclosure` | 报告披露 | 东方财富 | P1 |

### 3.13 技术排名 (rank.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `stock_rank_cxg_ths` | 创新高 | 同花顺 | P1 |
| `stock_rank_cxd_ths` | 创新低 | 同花顺 | P1 |
| `stock_rank_lxsz_ths` | 连续上涨 | 同花顺 | P1 |
| `stock_rank_lxxd_ths` | 连续下跌 | 同花顺 | P1 |
| `stock_rank_cxfl_ths` | 持续放量 | 同花顺 | P2 |
| `stock_rank_cxsl_ths` | 持续缩量 | 同花顺 | P2 |
| `stock_rank_ljqd_ths` | 量价齐跌 | 同花顺 | P2 |
| `stock_rank_ljqs_ths` | 量价齐升 | 同花顺 | P2 |
| `stock_rank_xstp_ths` | 向上突破 | 同花顺 | P2 |
| `stock_rank_xxtp_ths` | 向下突破 | 同花顺 | P2 |
| `stock_rank_xzjp_ths` | 险资举牌 | 同花顺 | P2 |

### 3.14 其他功能

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `stock_changes_em` | 股票异动 | 东方财富 | P1 |
| `stock_fund_flow_individual` | 个股资金流 | 东方财富 | P0 |
| `stock_fund_flow_concept` | 概念资金流 | 东方财富 | P1 |
| `stock_fund_flow_industry` | 行业资金流 | 东方财富 | P1 |
| `stock_fund_flow_big_deal` | 大单资金流 | 东方财富 | P1 |
| `stock_gpzy_pledge_ratio_em` | 股权质押比例 | 东方财富 | P1 |
| `stock_gpzy_pledge_ratio_detail_em` | 股权质押详情 | 东方财富 | P1 |
| `stock_gpzy_profile_em` | 股权质押概况 | 东方财富 | P2 |
| `stock_gpzy_industry_data_em` | 股权质押行业 | 东方财富 | P2 |
| `stock_gpzy_distribute_statistics_bank_em` | 质押银行统计 | 东方财富 | P2 |
| `stock_gpzy_distribute_statistics_company_em` | 质押公司统计 | 东方财富 | P2 |
| `stock_sy_em` | 商誉 | 东方财富 | P2 |
| `stock_sy_hy_em` | 商誉行业 | 东方财富 | P2 |
| `stock_sy_jz_em` | 商誉减值 | 东方财富 | P2 |
| `stock_sy_profile_em` | 商誉概况 | 东方财富 | P2 |
| `stock_sy_yq_em` | 商誉预期 | 东方财富 | P2 |
| `stock_tfp_em` | 停复牌 | 东方财富 | P1 |
| `stock_xgsglb_em` | 新股申购 | 东方财富 | P1 |
| `stock_xgsr_ths` | 新股上市 | 同花顺 | P1 |
| `stock_yzxdr_em` | 一字涨跌 | 东方财富 | P2 |
| `stock_zdhtmx_em` | 重大合同 | 东方财富 | P2 |

---

## 四、类型定义示例

```go
// types.go
package stock_feature

import "time"

// MarginData 融资融券数据
type MarginData struct {
    Date          time.Time `json:"date"`
    RzBalance     float64   `json:"rz_balance"`     // 融资余额
    RzBuy         float64   `json:"rz_buy"`         // 融资买入
    RzRepay       float64   `json:"rz_repay"`       // 融资偿还
    RqBalance     float64   `json:"rq_balance"`     // 融券余额
    RqSell        float64   `json:"rq_sell"`        // 融券卖出
    RqRepay       float64   `json:"rq_repay"`       // 融券偿还
    RzRqBalance   float64   `json:"rzrq_balance"`   // 融资融券余额
}

// LHBDetail 龙虎榜详情
type LHBDetail struct {
    Date       time.Time `json:"date"`
    Code       string    `json:"code"`
    Name       string    `json:"name"`
    Close      float64   `json:"close"`
    Change     float64   `json:"change"`
    Reason     string    `json:"reason"`      // 上榜原因
    BuyAmount  float64   `json:"buy_amount"`  // 买入金额
    SellAmount float64   `json:"sell_amount"` // 卖出金额
    NetAmount  float64   `json:"net_amount"`  // 净买入
}

// HSGTFlow 沪深港通资金流向
type HSGTFlow struct {
    Date     time.Time `json:"date"`
    HGTIn    float64   `json:"hgt_in"`     // 沪股通流入
    HGTOut   float64   `json:"hgt_out"`    // 沪股通流出
    HGTNet   float64   `json:"hgt_net"`    // 沪股通净流入
    SGTIn    float64   `json:"sgt_in"`     // 深股通流入
    SGTOut   float64   `json:"sgt_out"`    // 深股通流出
    SGTNet   float64   `json:"sgt_net"`    // 深股通净流入
    TotalNet float64   `json:"total_net"`  // 北向总净流入
}

// ZTPool 涨停池
type ZTPool struct {
    Code       string    `json:"code"`
    Name       string    `json:"name"`
    Price      float64   `json:"price"`
    ChangePct  float64   `json:"change_pct"`
    FirstTime  string    `json:"first_time"`  // 首次涨停
    LastTime   string    `json:"last_time"`   // 最后涨停
    OpenTimes  int       `json:"open_times"`  // 打开次数
    ZTReason   string    `json:"zt_reason"`   // 涨停原因
    ZTDays     int       `json:"zt_days"`     // 连板天数
    Date       time.Time `json:"date"`
}

// PerformanceForecast 业绩预告
type PerformanceForecast struct {
    Code       string  `json:"code"`
    Name       string  `json:"name"`
    Type       string  `json:"type"`        // 预增/预减/预盈/预亏
    Summary    string  `json:"summary"`     // 业绩摘要
    NetProfitL float64 `json:"net_profit_l"` // 净利润下限
    NetProfitH float64 `json:"net_profit_h"` // 净利润上限
    ChangeL    float64 `json:"change_l"`    // 变动幅度下限
    ChangeH    float64 `json:"change_h"`    // 变动幅度上限
    ReportDate string  `json:"report_date"` // 报告期
}
```

---

## 五、实现优先级

### P0 - 第一批实现

1. 融资融券 (`stock_margin_sse/szse`)
2. 龙虎榜 (`stock_lhb_detail_em`)
3. 沪深港通 (`stock_hsgt_hist_em`, `stock_hsgt_fund_flow_summary_em`)
4. 股东户数 (`stock_zh_a_gdhs`)
5. 业绩预告 (`stock_yjyg_em`)
6. 涨停池 (`stock_zt_pool_em`)

### P1 - 第二批实现

1. 机构调研
2. 估值数据
3. 技术排名
4. 板块数据

### P2 - 第三批实现

1. ESG 数据
2. 股票点评
3. 商誉数据

---

*文档版本: v1.0*
*最后更新: 2026-01-18*
