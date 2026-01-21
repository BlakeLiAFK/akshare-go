# index 模块计划

> 模块: 指数数据
> 函数数量: 95 个
> 优先级: P1
> 预计周期: 第11周

---

## 一、模块概述

指数模块提供 A 股指数、全球指数、行业指数、申万指数等数据，包括实时行情、历史K线、成分股信息。

---

## 二、Go 文件结构

```
index/
├── zh_a.go              # A股指数
├── global.go            # 全球指数
├── sw.go                # 申万指数
├── cni.go               # 中证指数
├── csindex.go           # 中证官方
├── industry.go          # 行业指数
├── commodity.go         # 商品指数
├── option_qvix.go       # 期权波动率指数
├── sentiment.go         # 情绪指数
├── hk.go                # 港股指数
├── us.go                # 美股指数
└── types.go             # 类型定义
```

---

## 三、函数清单

### 3.1 A股指数 (zh_a.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `index_zh_a_hist` | A股指数历史 | 新浪财经 | P0 |
| `index_zh_a_hist_min_em` | A股指数分钟 | 东方财富 | P0 |
| `stock_zh_index_spot_em` | A股指数实时 | 东方财富 | P0 |
| `stock_zh_index_spot_sina` | A股指数实时(新浪) | 新浪财经 | P1 |
| `stock_zh_index_daily` | A股指数日K(新浪) | 新浪财经 | P1 |
| `stock_zh_index_daily_em` | A股指数日K(东财) | 东方财富 | P1 |
| `stock_zh_index_daily_tx` | A股指数日K(腾讯) | 腾讯 | P2 |
| `index_stock_info` | 指数信息 | - | P1 |
| `index_stock_cons` | 指数成分股 | - | P0 |
| `index_stock_cons_sina` | 指数成分股(新浪) | 新浪财经 | P1 |

### 3.2 全球指数 (global.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `index_global_spot_em` | 全球指数实时 | 东方财富 | P0 |
| `index_global_hist_em` | 全球指数历史 | 东方财富 | P0 |
| `index_global_hist_sina` | 全球指数历史(新浪) | 新浪财经 | P1 |
| `index_global_name_table` | 全球指数名称表 | - | P2 |

### 3.3 申万指数 (sw.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `index_hist_sw` | 申万指数历史 | 申万宏源 | P1 |
| `index_min_sw` | 申万指数分钟 | 申万宏源 | P2 |
| `index_realtime_sw` | 申万指数实时 | 申万宏源 | P1 |
| `index_realtime_fund_sw` | 申万基金实时 | 申万宏源 | P2 |
| `index_hist_fund_sw` | 申万基金历史 | 申万宏源 | P2 |
| `index_component_sw` | 申万成分股 | 申万宏源 | P1 |
| `index_analysis_daily_sw` | 申万日度分析 | 申万宏源 | P2 |
| `index_analysis_weekly_sw` | 申万周度分析 | 申万宏源 | P2 |
| `index_analysis_monthly_sw` | 申万月度分析 | 申万宏源 | P2 |
| `index_analysis_week_month_sw` | 申万周月分析 | 申万宏源 | P2 |
| `sw_index_first_info` | 申万一级行业 | 申万宏源 | P1 |
| `sw_index_second_info` | 申万二级行业 | 申万宏源 | P1 |
| `sw_index_third_info` | 申万三级行业 | 申万宏源 | P2 |
| `sw_index_third_cons` | 申万三级成分 | 申万宏源 | P2 |

### 3.4 中证指数 (cni.go / csindex.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `index_all_cni` | 中证指数列表 | 中证 | P1 |
| `index_hist_cni` | 中证指数历史 | 中证 | P1 |
| `index_detail_cni` | 中证指数详情 | 中证 | P2 |
| `index_detail_hist_cni` | 中证详情历史 | 中证 | P2 |
| `index_detail_hist_adjust_cni` | 中证调整历史 | 中证 | P2 |
| `index_csindex_all` | 中证官方列表 | 中证官网 | P1 |
| `stock_zh_index_hist_csindex` | 中证指数历史 | 中证官网 | P1 |
| `stock_zh_index_value_csindex` | 中证估值 | 中证官网 | P2 |
| `index_stock_cons_csindex` | 中证成分股 | 中证官网 | P1 |
| `index_stock_cons_weight_csindex` | 中证成分权重 | 中证官网 | P2 |
| `index_code_id_map_em` | 指数代码映射 | 东方财富 | P2 |

### 3.5 商品/特色指数 (commodity.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `index_sugar_msweet` | 白糖指数 | 甜蜜网 | P2 |
| `index_inner_quote_sugar_msweet` | 糖内盘报价 | 甜蜜网 | P3 |
| `index_outer_quote_sugar_msweet` | 糖外盘报价 | 甜蜜网 | P3 |
| `index_hog_spot_price` | 生猪现货价格 | - | P2 |
| `index_price_cflp` | 物流价格指数 | 中物联 | P2 |
| `index_volume_cflp` | 物流量指数 | 中物联 | P2 |
| `index_kq_fz` | 科琪纺织指数 | 科琪 | P3 |
| `index_kq_fashion` | 科琪时尚指数 | 科琪 | P3 |
| `drewry_wci_index` | 集装箱运价指数 | Drewry | P2 |
| `spot_goods` | 现货商品 | - | P2 |

### 3.6 期权波动率指数 (option_qvix.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `index_option_50etf_qvix` | 50ETF波动率 | - | P2 |
| `index_option_50etf_min_qvix` | 50ETF波动率分钟 | - | P2 |
| `index_option_300etf_qvix` | 300ETF波动率 | - | P2 |
| `index_option_300etf_min_qvix` | 300ETF波动率分钟 | - | P2 |
| `index_option_500etf_qvix` | 500ETF波动率 | - | P2 |
| `index_option_500etf_min_qvix` | 500ETF波动率分钟 | - | P2 |
| `index_option_100etf_qvix` | 100ETF波动率 | - | P2 |
| `index_option_100etf_min_qvix` | 100ETF波动率分钟 | - | P2 |
| `index_option_300index_qvix` | 300指数波动率 | - | P2 |
| `index_option_300index_min_qvix` | 300指数波动率分钟 | - | P2 |
| `index_option_1000index_qvix` | 1000指数波动率 | - | P2 |
| `index_option_1000index_min_qvix` | 1000指数波动率分钟 | - | P2 |
| `index_option_50index_qvix` | 50指数波动率 | - | P2 |
| `index_option_50index_min_qvix` | 50指数波动率分钟 | - | P2 |
| `index_option_cyb_qvix` | 创业板波动率 | - | P2 |
| `index_option_cyb_min_qvix` | 创业板波动率分钟 | - | P2 |
| `index_option_kcb_qvix` | 科创板波动率 | - | P2 |
| `index_option_kcb_min_qvix` | 科创板波动率分钟 | - | P2 |

### 3.7 情绪/特色指数 (sentiment.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `index_news_sentiment_scope` | 新闻情绪指数 | - | P2 |
| `index_eri` | 经济恢复指数 | - | P2 |
| `index_yw` | 义乌小商品指数 | - | P3 |
| `index_ai_cx` | AI财新指数 | 财新 | P2 |
| `index_awpr_cx` | 气候指数 | 财新 | P3 |
| `index_bei_cx` | 营商环境指数 | 财新 | P3 |
| `index_bi_cx` | 破产指数 | 财新 | P3 |
| `index_cci_cx` | 消费者信心 | 财新 | P2 |
| `index_ci_cx` | 建筑指数 | 财新 | P3 |
| `index_dei_cx` | 数字经济指数 | 财新 | P2 |
| `index_fi_cx` | 金融指数 | 财新 | P2 |
| `index_ii_cx` | 产业指数 | 财新 | P3 |
| `index_li_cx` | 劳动力指数 | 财新 | P3 |
| `index_neaw_cx` | 新经济活动指数 | 财新 | P2 |
| `index_neei_cx` | 新经济效率指数 | 财新 | P2 |
| `index_nei_cx` | 新经济指数 | 财新 | P2 |
| `index_pmi_com_cx` | 综合PMI | 财新 | P1 |
| `index_pmi_man_cx` | 制造业PMI | 财新 | P1 |
| `index_pmi_ser_cx` | 服务业PMI | 财新 | P1 |
| `index_qli_cx` | 生活质量指数 | 财新 | P3 |
| `index_si_cx` | 服务指数 | 财新 | P3 |
| `index_ti_cx` | 技术指数 | 财新 | P3 |

### 3.8 港股/美股指数 (hk.go / us.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `stock_hk_index_spot_em` | 港股指数实时 | 东方财富 | P1 |
| `stock_hk_index_spot_sina` | 港股指数实时(新浪) | 新浪财经 | P2 |
| `stock_hk_index_daily_em` | 港股指数日K | 东方财富 | P1 |
| `stock_hk_index_daily_sina` | 港股指数日K(新浪) | 新浪财经 | P2 |
| `index_us_stock_sina` | 美股指数 | 新浪财经 | P1 |

### 3.9 工具函数

| 函数名 | 功能 | 优先级 |
|--------|------|--------|
| `get_zh_index_page_count` | A股指数页数 | P3 |
| `get_hk_index_page_count` | 港股指数页数 | P3 |
| `get_tx_start_year` | 腾讯起始年份 | P3 |
| `stock_a_code_to_symbol` | 代码转symbol | P2 |

---

## 四、类型定义示例

```go
// types.go
package index

import "time"

// IndexQuote 指数实时行情
type IndexQuote struct {
    Code      string    `json:"code"`
    Name      string    `json:"name"`
    Open      float64   `json:"open"`
    High      float64   `json:"high"`
    Low       float64   `json:"low"`
    Close     float64   `json:"close"`
    PreClose  float64   `json:"pre_close"`
    Change    float64   `json:"change"`
    ChangePct float64   `json:"change_pct"`
    Volume    int64     `json:"volume"`
    Amount    float64   `json:"amount"`
    Time      time.Time `json:"time"`
}

// IndexKLine 指数K线
type IndexKLine struct {
    Date   time.Time `json:"date"`
    Open   float64   `json:"open"`
    High   float64   `json:"high"`
    Low    float64   `json:"low"`
    Close  float64   `json:"close"`
    Volume int64     `json:"volume"`
    Amount float64   `json:"amount"`
}

// IndexConstituent 指数成分股
type IndexConstituent struct {
    Code   string  `json:"code"`
    Name   string  `json:"name"`
    Weight float64 `json:"weight"` // 权重
}

// SWIndex 申万行业指数
type SWIndex struct {
    Code      string  `json:"code"`
    Name      string  `json:"name"`
    Level     int     `json:"level"`      // 级别(1/2/3)
    ParentCode string `json:"parent_code"` // 上级代码
    Change    float64 `json:"change"`
    ChangePct float64 `json:"change_pct"`
    PE        float64 `json:"pe"`
    PB        float64 `json:"pb"`
}
```

---

## 五、实现优先级

### P0 - 第一批实现

1. `stock_zh_index_spot_em` - A股指数实时
2. `index_zh_a_hist` - A股指数历史
3. `index_global_spot_em` - 全球指数实时
4. `index_stock_cons` - 指数成分股

### P1 - 第二批实现

1. 申万指数系列
2. 中证指数系列
3. 港股/美股指数
4. PMI指数

### P2 - 第三批实现

1. 期权波动率
2. 情绪指数
3. 商品指数

---

*文档版本: v1.0*
*最后更新: 2026-01-18*
