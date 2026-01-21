# stock 模块计划

> 模块: 股票行情
> 函数数量: 128 个
> 优先级: P0
> 预计周期: 第3-5周

---

## 一、模块概述

股票模块是 akshare 的核心模块，提供 A 股、港股、美股的实时行情、历史 K 线、资金流向、板块数据等功能。

---

## 二、Go 文件结构

```
stock/
├── zh_a_hist.go          # A股历史行情
├── zh_a_spot.go          # A股实时行情
├── zh_a_minute.go        # A股分钟数据
├── zh_b.go               # B股数据
├── hk.go                 # 港股数据
├── us.go                 # 美股数据
├── ah.go                 # AH股比价
├── board_concept.go      # 概念板块
├── board_industry.go     # 行业板块
├── fund_flow.go          # 资金流向
├── info.go               # 股票信息
├── hot.go                # 热门股票
├── dzjy.go               # 大宗交易
├── ipo.go                # IPO信息
├── cninfo.go             # 巨潮资讯数据
├── sse.go                # 上交所数据
├── szse.go               # 深交所数据
└── types.go              # 类型定义
```

---

## 三、函数清单

### 3.1 A股历史行情 (zh_a_hist.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `stock_zh_a_hist` | A股历史日K线 | 东方财富 | P0 |
| `stock_zh_a_daily` | A股历史日K线(新浪) | 新浪财经 | P1 |
| `stock_zh_a_minute` | A股分钟K线 | 新浪财经 | P0 |
| `stock_zh_a_cdr_daily` | CDR历史行情 | 新浪财经 | P2 |
| `stock_zh_a_tick_tx_js` | A股逐笔成交 | 腾讯 | P2 |
| `stock_zh_kcb_daily` | 科创板历史行情 | 新浪财经 | P1 |

### 3.2 A股实时行情 (zh_a_spot.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `stock_zh_a_spot` | A股实时行情 | 新浪财经 | P0 |
| `stock_zh_a_spot_em` | A股实时行情(东财) | 东方财富 | P0 |
| `stock_zh_a_new` | 新股实时行情 | 新浪财经 | P1 |
| `stock_zh_a_new_em` | 新股实时行情(东财) | 东方财富 | P1 |
| `stock_zh_a_st_em` | ST股实时行情 | 东方财富 | P1 |
| `stock_zh_a_stop_em` | 停牌股票 | 东方财富 | P2 |
| `stock_zh_kcb_spot` | 科创板实时行情 | 新浪财经 | P1 |
| `stock_intraday_em` | 日内分时数据 | 东方财富 | P0 |
| `stock_intraday_sina` | 日内分时数据(新浪) | 新浪财经 | P1 |
| `stock_bid_ask_em` | 买卖五档数据 | 东方财富 | P1 |

### 3.3 B股数据 (zh_b.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `stock_zh_b_spot` | B股实时行情 | 新浪财经 | P2 |
| `stock_zh_b_daily` | B股历史行情 | 新浪财经 | P2 |
| `stock_zh_b_minute` | B股分钟数据 | 新浪财经 | P2 |

### 3.4 港股数据 (hk.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `stock_hk_spot` | 港股实时行情 | 新浪财经 | P0 |
| `stock_hk_daily` | 港股历史行情 | 新浪财经 | P0 |
| `stock_hk_spot_em` | 港股实时行情(东财) | 东方财富 | P1 |
| `stock_hk_famous_spot_em` | 港股知名股票 | 东方财富 | P2 |
| `stock_hk_company_profile_em` | 港股公司资料 | 东方财富 | P2 |
| `stock_hk_dividend_payout_em` | 港股分红派息 | 东方财富 | P2 |
| `stock_hk_financial_indicator_em` | 港股财务指标 | 东方财富 | P2 |
| `stock_hk_growth_comparison_em` | 港股成长对比 | 东方财富 | P2 |
| `stock_hk_hot_rank_em` | 港股热度排名 | 东方财富 | P2 |
| `stock_hk_hot_rank_detail_em` | 港股热度详情 | 东方财富 | P2 |
| `stock_hk_hot_rank_detail_realtime_em` | 港股热度实时 | 东方财富 | P2 |
| `stock_hk_hot_rank_latest_em` | 港股热度最新 | 东方财富 | P2 |
| `stock_hk_scale_comparison_em` | 港股规模对比 | 东方财富 | P2 |
| `stock_hk_security_profile_em` | 港股证券资料 | 东方财富 | P2 |
| `stock_hk_valuation_comparison_em` | 港股估值对比 | 东方财富 | P2 |
| `stock_hk_fhpx_detail_ths` | 港股分红详情 | 同花顺 | P2 |

### 3.5 美股数据 (us.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `stock_us_spot` | 美股实时行情 | 新浪财经 | P0 |
| `stock_us_daily` | 美股历史行情 | 新浪财经 | P0 |
| `stock_us_famous_spot_em` | 美股知名股票 | 东方财富 | P2 |
| `stock_us_pink_spot_em` | 美股粉单市场 | 东方财富 | P3 |
| `get_us_stock_name` | 获取美股名称 | 新浪财经 | P2 |

### 3.6 AH股比价 (ah.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `stock_zh_ah_spot` | AH股实时行情 | 新浪财经 | P1 |
| `stock_zh_ah_spot_em` | AH股实时行情(东财) | 东方财富 | P1 |
| `stock_zh_ah_daily` | AH股历史行情 | 新浪财经 | P2 |
| `stock_zh_ah_name` | AH股列表 | 新浪财经 | P2 |

### 3.7 概念板块 (board_concept.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `stock_board_concept_name_em` | 概念板块名称 | 东方财富 | P0 |
| `stock_board_concept_spot_em` | 概念板块实时 | 东方财富 | P0 |
| `stock_board_concept_cons_em` | 概念板块成分股 | 东方财富 | P0 |
| `stock_board_concept_hist_em` | 概念板块历史 | 东方财富 | P1 |
| `stock_board_concept_hist_min_em` | 概念板块分钟 | 东方财富 | P2 |
| `stock_concept_fund_flow_hist` | 概念板块资金流 | 东方财富 | P1 |

### 3.8 行业板块 (board_industry.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `stock_board_industry_name_em` | 行业板块名称 | 东方财富 | P0 |
| `stock_board_industry_spot_em` | 行业板块实时 | 东方财富 | P0 |
| `stock_board_industry_cons_em` | 行业板块成分股 | 东方财富 | P0 |
| `stock_board_industry_hist_em` | 行业板块历史 | 东方财富 | P1 |
| `stock_board_industry_hist_min_em` | 行业板块分钟 | 东方财富 | P2 |
| `stock_industry_category_cninfo` | 行业分类(巨潮) | 巨潮资讯 | P2 |
| `stock_industry_change_cninfo` | 行业变更(巨潮) | 巨潮资讯 | P2 |
| `stock_industry_clf_hist_sw` | 申万行业历史 | 申万宏源 | P2 |
| `stock_industry_pe_ratio_cninfo` | 行业市盈率 | 巨潮资讯 | P2 |

### 3.9 资金流向 (fund_flow.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `stock_individual_fund_flow` | 个股资金流向 | 东方财富 | P0 |
| `stock_individual_fund_flow_rank` | 个股资金排名 | 东方财富 | P0 |
| `stock_market_fund_flow` | 大盘资金流向 | 东方财富 | P0 |
| `stock_main_fund_flow` | 主力资金流向 | 东方财富 | P0 |
| `stock_sector_fund_flow_rank` | 板块资金排名 | 东方财富 | P1 |
| `stock_sector_fund_flow_hist` | 板块资金历史 | 东方财富 | P1 |
| `stock_sector_fund_flow_summary` | 板块资金汇总 | 东方财富 | P1 |
| `stock_sector_spot` | 板块实时行情 | 东方财富 | P1 |
| `stock_sector_detail` | 板块详情 | 东方财富 | P1 |

### 3.10 股票信息 (info.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `stock_info_a_code_name` | A股代码名称 | 新浪财经 | P0 |
| `stock_info_sh_name_code` | 沪市代码名称 | 上交所 | P1 |
| `stock_info_sz_name_code` | 深市代码名称 | 深交所 | P1 |
| `stock_info_bj_name_code` | 北交所代码名称 | 北交所 | P1 |
| `stock_info_sh_delist` | 沪市退市股票 | 上交所 | P2 |
| `stock_info_sz_delist` | 深市退市股票 | 深交所 | P2 |
| `stock_info_change_name` | 股票改名记录 | 新浪财经 | P2 |
| `stock_info_sz_change_name` | 深市改名记录 | 深交所 | P2 |
| `stock_individual_info_em` | 个股信息 | 东方财富 | P0 |
| `stock_individual_spot_xq` | 个股实时(雪球) | 雪球 | P2 |
| `stock_profile_cninfo` | 公司简介 | 巨潮资讯 | P2 |

### 3.11 热门股票 (hot.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `stock_hot_rank_em` | 热门排名 | 东方财富 | P1 |
| `stock_hot_rank_latest_em` | 热门排名最新 | 东方财富 | P1 |
| `stock_hot_rank_detail_em` | 热门排名详情 | 东方财富 | P2 |
| `stock_hot_rank_detail_realtime_em` | 热门排名实时 | 东方财富 | P2 |
| `stock_hot_rank_relate_em` | 热门相关股票 | 东方财富 | P2 |
| `stock_hot_up_em` | 人气飙升榜 | 东方财富 | P2 |
| `stock_hot_keyword_em` | 热门关键词 | 东方财富 | P2 |
| `stock_hot_search_baidu` | 百度热搜 | 百度 | P3 |

### 3.12 大宗交易 (dzjy.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `stock_dzjy_mrmx` | 大宗交易每日明细 | 东方财富 | P1 |
| `stock_dzjy_mrtj` | 大宗交易每日统计 | 东方财富 | P1 |
| `stock_dzjy_sctj` | 大宗交易市场统计 | 东方财富 | P2 |
| `stock_dzjy_hygtj` | 活跃股统计 | 东方财富 | P2 |
| `stock_dzjy_hyyybtj` | 活跃营业部统计 | 东方财富 | P2 |
| `stock_dzjy_yybph` | 营业部排行 | 东方财富 | P2 |

### 3.13 IPO 信息 (ipo.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `stock_new_gh_cninfo` | 新股公告 | 巨潮资讯 | P2 |
| `stock_new_ipo_cninfo` | 新股IPO信息 | 巨潮资讯 | P2 |
| `stock_ipo_summary_cninfo` | IPO汇总 | 巨潮资讯 | P2 |

### 3.14 巨潮资讯数据 (cninfo.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `stock_dividend_cninfo` | 分红数据 | 巨潮资讯 | P2 |
| `stock_hold_num_cninfo` | 股东人数 | 巨潮资讯 | P2 |
| `stock_hold_change_cninfo` | 持股变动 | 巨潮资讯 | P2 |
| `stock_hold_control_cninfo` | 实控人信息 | 巨潮资讯 | P2 |
| `stock_hold_management_detail_cninfo` | 管理层持股详情 | 巨潮资讯 | P2 |
| `stock_allotment_cninfo` | 配股信息 | 巨潮资讯 | P2 |
| `stock_cg_equity_mortgage_cninfo` | 股权质押 | 巨潮资讯 | P2 |
| `stock_cg_guarantee_cninfo` | 担保信息 | 巨潮资讯 | P2 |
| `stock_cg_lawsuit_cninfo` | 诉讼信息 | 巨潮资讯 | P2 |
| `stock_share_change_cninfo` | 股本变动 | 巨潮资讯 | P2 |
| `stock_rank_forecast_cninfo` | 业绩预测排名 | 巨潮资讯 | P2 |

### 3.15 交易所数据 (sse.go / szse.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `stock_sse_summary` | 上交所市场概况 | 上交所 | P1 |
| `stock_sse_deal_daily` | 上交所每日成交 | 上交所 | P1 |
| `stock_szse_summary` | 深交所市场概况 | 深交所 | P1 |
| `stock_szse_area_summary` | 深市地区分布 | 深交所 | P2 |
| `stock_szse_sector_summary` | 深市行业分布 | 深交所 | P2 |
| `stock_share_hold_change_sse` | 沪市股权变动 | 上交所 | P2 |
| `stock_share_hold_change_szse` | 深市股权变动 | 深交所 | P2 |
| `stock_share_hold_change_bse` | 北交所股权变动 | 北交所 | P2 |
| `stock_staq_net_stop` | 老三板停牌 | - | P3 |

### 3.16 其他股票功能

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `stock_repurchase_em` | 股票回购 | 东方财富 | P2 |
| `stock_gsrl_gsdt_em` | 公司日历 | 东方财富 | P2 |
| `stock_report_fund_hold` | 基金持股 | 东方财富 | P2 |
| `stock_report_fund_hold_detail` | 基金持股详情 | 东方财富 | P2 |
| `stock_news_main_cx` | 股票新闻 | 财新 | P3 |
| `stock_js_weibo_report` | 微博舆情 | 微博 | P3 |
| `stock_js_weibo_nlp_time` | 微博NLP分析 | 微博 | P3 |
| `stock_price_js` | 股价JavaScript解析 | - | P3 |
| `stock_hold_management_person_em` | 高管持股 | 东方财富 | P2 |
| `stock_hold_management_detail_em` | 高管持股详情 | 东方财富 | P2 |
| `stock_zh_dupont_comparison_em` | 杜邦分析对比 | 东方财富 | P2 |
| `stock_zh_growth_comparison_em` | 成长能力对比 | 东方财富 | P2 |
| `stock_zh_scale_comparison_em` | 规模对比 | 东方财富 | P2 |
| `stock_zh_valuation_comparison_em` | 估值对比 | 东方财富 | P2 |
| `stock_zh_kcb_report_em` | 科创板报告 | 东方财富 | P2 |
| `stock_hsgt_sh_hk_spot_em` | 沪深港通实时 | 东方财富 | P1 |
| `get_zh_kcb_page_count` | 科创板页数 | - | P3 |

---

## 四、数据源 URL 模式

### 4.1 东方财富

| URL 模式 | 用途 |
|----------|------|
| `push2.eastmoney.com/api/qt/clist/get` | 股票列表 |
| `push2his.eastmoney.com/api/qt/stock/kline/get` | K线数据 |
| `push2.eastmoney.com/api/qt/stock/get` | 实时行情 |
| `datacenter-web.eastmoney.com/api/data/v1/get` | 数据中心 |

### 4.2 新浪财经

| URL 模式 | 用途 |
|----------|------|
| `hq.sinajs.cn/list=` | 实时行情 |
| `vip.stock.finance.sina.com.cn/quotes_service/api/` | 行情API |
| `finance.sina.com.cn/realstock/company/` | 历史数据 |

### 4.3 巨潮资讯

| URL 模式 | 用途 |
|----------|------|
| `webapi.cninfo.com.cn/api/stock/` | 股票数据 |
| `webapi.cninfo.com.cn/api/disclosure/` | 公告信息 |

---

## 五、类型定义示例

```go
// types.go
package stock

import "time"

// StockQuote A股实时行情
type StockQuote struct {
    Code      string    `json:"code"`       // 股票代码
    Name      string    `json:"name"`       // 股票名称
    Open      float64   `json:"open"`       // 开盘价
    High      float64   `json:"high"`       // 最高价
    Low       float64   `json:"low"`        // 最低价
    Close     float64   `json:"close"`      // 收盘价/最新价
    PreClose  float64   `json:"pre_close"`  // 昨收价
    Change    float64   `json:"change"`     // 涨跌额
    ChangePct float64   `json:"change_pct"` // 涨跌幅
    Volume    int64     `json:"volume"`     // 成交量（手）
    Amount    float64   `json:"amount"`     // 成交额（元）
    Time      time.Time `json:"time"`       // 时间
}

// StockKLine K线数据
type StockKLine struct {
    Date   time.Time `json:"date"`
    Open   float64   `json:"open"`
    High   float64   `json:"high"`
    Low    float64   `json:"low"`
    Close  float64   `json:"close"`
    Volume int64     `json:"volume"`
    Amount float64   `json:"amount"`
    Adjust string    `json:"adjust"` // qfq/hfq/none
}

// BoardInfo 板块信息
type BoardInfo struct {
    Code      string  `json:"code"`
    Name      string  `json:"name"`
    Change    float64 `json:"change"`
    ChangePct float64 `json:"change_pct"`
    Volume    int64   `json:"volume"`
    Amount    float64 `json:"amount"`
    LeadStock string  `json:"lead_stock"` // 领涨股
}

// FundFlow 资金流向
type FundFlow struct {
    Date       time.Time `json:"date"`
    MainIn     float64   `json:"main_in"`     // 主力流入
    MainOut    float64   `json:"main_out"`    // 主力流出
    MainNet    float64   `json:"main_net"`    // 主力净流入
    RetailIn   float64   `json:"retail_in"`   // 散户流入
    RetailOut  float64   `json:"retail_out"`  // 散户流出
    RetailNet  float64   `json:"retail_net"`  // 散户净流入
}

// StockInfo 股票基本信息
type StockInfo struct {
    Code         string  `json:"code"`
    Name         string  `json:"name"`
    Industry     string  `json:"industry"`
    MarketCap    float64 `json:"market_cap"`    // 总市值
    CirculateCap float64 `json:"circulate_cap"` // 流通市值
    PE           float64 `json:"pe"`            // 市盈率
    PB           float64 `json:"pb"`            // 市净率
    TotalShares  int64   `json:"total_shares"`  // 总股本
    FloatShares  int64   `json:"float_shares"`  // 流通股本
}
```

---

## 六、实现优先级

### P0 - 第一批实现 (高频核心)

1. `stock_zh_a_spot_em` - A股实时行情
2. `stock_zh_a_hist` - A股历史K线
3. `stock_individual_fund_flow` - 个股资金流向
4. `stock_board_concept_name_em` - 概念板块
5. `stock_board_industry_name_em` - 行业板块
6. `stock_hk_spot` - 港股实时
7. `stock_us_spot` - 美股实时
8. `stock_info_a_code_name` - 股票代码名称

### P1 - 第二批实现 (常用功能)

1. 港股/美股历史数据
2. 资金流向排名
3. 板块成分股
4. 热门排名
5. 沪深港通

### P2 - 第三批实现 (进阶功能)

1. 大宗交易
2. IPO 信息
3. 巨潮资讯数据
4. 公司信息

### P3 - 第四批实现 (低频功能)

1. 微博舆情
2. 百度热搜
3. 老三板数据

---

*文档版本: v1.0*
*最后更新: 2026-01-18*
