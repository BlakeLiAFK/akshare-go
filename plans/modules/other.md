# other 模块计划

> 模块: 其他模块集合
> 函数数量: ~100 个
> 优先级: P2-P3
> 预计周期: 第15-16周

---

## 一、模块概述

此文档包含多个小型模块的计划：外汇、加密货币、能源、财富排行、新闻电影、空气质量、学术指数、现货、REITs、利率等。

---

## 二、currency 外汇模块 (9 函数)

### Go 文件结构
```
currency/
├── boc.go        # 中国银行
├── convert.go    # 汇率转换
├── history.go    # 历史汇率
└── types.go
```

### 函数清单

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `currency_boc_safe` | 中行外汇牌价(安全) | 中国银行 | P1 |
| `currency_boc_sina` | 中行外汇牌价(新浪) | 新浪财经 | P1 |
| `currency_convert` | 汇率转换 | - | P1 |
| `currency_currencies` | 货币列表 | - | P2 |
| `currency_history` | 历史汇率 | - | P1 |
| `currency_latest` | 最新汇率 | - | P1 |
| `currency_time_series` | 汇率时间序列 | - | P2 |
| `forex_hist_em` | 外汇历史(东财) | 东方财富 | P1 |
| `forex_spot_em` | 外汇实时(东财) | 东方财富 | P1 |

---

## 三、crypto 加密货币模块 (2 函数)

### Go 文件结构
```
crypto/
├── bitcoin.go
└── types.go
```

### 函数清单

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `crypto_bitcoin_cme` | CME比特币 | CME | P2 |
| `crypto_bitcoin_hold_report` | 比特币持仓报告 | - | P3 |

---

## 四、energy 能源模块 (8 函数)

### Go 文件结构
```
energy/
├── carbon.go     # 碳排放
├── oil.go        # 原油
└── types.go
```

### 函数清单

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `energy_carbon_bj` | 北京碳市场 | - | P2 |
| `energy_carbon_domestic` | 国内碳市场 | - | P2 |
| `energy_carbon_eu` | 欧洲碳市场 | - | P2 |
| `energy_carbon_gz` | 广州碳市场 | - | P2 |
| `energy_carbon_hb` | 湖北碳市场 | - | P2 |
| `energy_carbon_sz` | 深圳碳市场 | - | P2 |
| `energy_oil_detail` | 原油详情 | - | P2 |
| `energy_oil_hist` | 原油历史 | - | P2 |

---

## 五、fortune 财富排行模块 (6 函数)

### Go 文件结构
```
fortune/
├── forbes.go
├── hurun.go
├── bloomberg.go
└── types.go
```

### 函数清单

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `forbes_rank` | 福布斯排行 | 福布斯 | P3 |
| `fortune_rank` | 财富排行 | 财富杂志 | P3 |
| `hurun_rank` | 胡润排行 | 胡润 | P3 |
| `xincaifu_rank` | 新财富排行 | 新财富 | P3 |
| `index_bloomberg_billionaires` | 彭博亿万富翁 | 彭博 | P3 |
| `index_bloomberg_billionaires_hist` | 彭博亿万富翁历史 | 彭博 | P3 |

---

## 六、news/movie 新闻电影模块 (19 函数)

### Go 文件结构
```
news/
├── cctv.go       # 央视新闻
├── baidu.go      # 百度新闻
├── stock.go      # 股票新闻
└── types.go

movie/
├── boxoffice.go  # 票房
├── tv.go         # 电视
└── types.go
```

### 函数清单

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `news_cctv` | 央视新闻 | 央视 | P3 |
| `news_economic_baidu` | 百度财经新闻 | 百度 | P3 |
| `news_report_time_baidu` | 百度财报时间 | 百度 | P3 |
| `news_trade_notify_dividend_baidu` | 百度分红通知 | 百度 | P3 |
| `news_trade_notify_suspend_baidu` | 百度停牌通知 | 百度 | P3 |
| `stock_news_em` | 股票新闻(东财) | 东方财富 | P2 |
| `movie_boxoffice_daily` | 单日票房 | 猫眼 | P3 |
| `movie_boxoffice_weekly` | 周票房 | 猫眼 | P3 |
| `movie_boxoffice_monthly` | 月票房 | 猫眼 | P3 |
| `movie_boxoffice_yearly` | 年票房 | 猫眼 | P3 |
| `movie_boxoffice_yearly_first_week` | 年度首周票房 | 猫眼 | P3 |
| `movie_boxoffice_realtime` | 实时票房 | 猫眼 | P3 |
| `movie_boxoffice_cinema_daily` | 影院日票房 | 猫眼 | P3 |
| `movie_boxoffice_cinema_weekly` | 影院周票房 | 猫眼 | P3 |
| `video_tv` | 电视剧 | - | P3 |
| `video_variety_show` | 综艺节目 | - | P3 |
| `business_value_artist` | 艺人商业价值 | - | P3 |
| `online_value_artist` | 艺人网络价值 | - | P3 |

---

## 七、air/article 空气质量/学术指数模块 (16 函数)

### Go 文件结构
```
air/
├── quality.go
├── city.go
└── types.go

article/
├── index.go
├── fred.go
└── types.go
```

### 函数清单

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `air_quality_hist` | 空气质量历史 | - | P3 |
| `air_quality_rank` | 空气质量排名 | - | P3 |
| `air_quality_watch_point` | 空气监测点 | - | P3 |
| `air_quality_hebei` | 河北空气质量 | - | P3 |
| `air_city_table` | 城市列表 | - | P3 |
| `sunrise_city_list` | 日出城市列表 | - | P3 |
| `sunrise_daily` | 每日日出 | - | P3 |
| `sunrise_monthly` | 月度日出 | - | P3 |
| `article_epu_index` | 经济政策不确定性指数 | - | P2 |
| `article_ff_crr` | FF因子 | - | P2 |
| `article_oman_rv` | Oman波动率 | - | P3 |
| `article_oman_rv_short` | Oman短期波动率 | - | P3 |
| `article_rlab_rv` | RLAB波动率 | - | P3 |
| `fred_md` | FRED月度数据 | FRED | P2 |
| `fred_qd` | FRED季度数据 | FRED | P2 |
| `has_month_data` | 月度数据检查 | - | P3 |

---

## 八、spot/reits/rate 现货/REITs/利率模块 (21 函数)

### Go 文件结构
```
spot/
├── sge.go        # 上海黄金交易所
├── commodity.go  # 商品现货
└── types.go

reits/
├── reits.go
└── types.go

rate/
├── interbank.go  # 银行间利率
├── repo.go       # 回购利率
└── types.go
```

### 函数清单

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `spot_golden_benchmark_sge` | 黄金基准价 | 上海黄金交易所 | P2 |
| `spot_silver_benchmark_sge` | 白银基准价 | 上海黄金交易所 | P2 |
| `spot_hist_sge` | 贵金属历史 | 上海黄金交易所 | P2 |
| `spot_quotations_sge` | 贵金属报价 | 上海黄金交易所 | P2 |
| `spot_symbol_table_sge` | 贵金属品种 | 上海黄金交易所 | P2 |
| `spot_corn_price_soozhu` | 玉米价格 | 搜猪网 | P3 |
| `spot_soybean_price_soozhu` | 大豆价格 | 搜猪网 | P3 |
| `spot_hog_soozhu` | 生猪价格 | 搜猪网 | P2 |
| `spot_hog_lean_price_soozhu` | 瘦肉猪价格 | 搜猪网 | P3 |
| `spot_hog_crossbred_soozhu` | 杂交猪价格 | 搜猪网 | P3 |
| `spot_hog_three_way_soozhu` | 三元猪价格 | 搜猪网 | P3 |
| `spot_hog_year_trend_soozhu` | 生猪年度趋势 | 搜猪网 | P3 |
| `spot_mixed_feed_soozhu` | 配合饲料 | 搜猪网 | P3 |
| `spot_price_qh` | 现货价格 | 期货网 | P2 |
| `spot_price_table_qh` | 现货价格表 | 期货网 | P2 |
| `reits_realtime_em` | REITs实时 | 东方财富 | P2 |
| `reits_hist_em` | REITs历史 | 东方财富 | P2 |
| `reits_hist_min_em` | REITs分钟 | 东方财富 | P2 |
| `rate_interbank` | 银行间利率 | - | P1 |
| `repo_rate_hist` | 回购利率历史 | - | P2 |
| `repo_rate_query` | 回购利率查询 | - | P2 |

---

## 九、fx/bank/other 外汇交易/银行/其他模块 (17 函数)

### Go 文件结构
```
fx/
├── quote.go
├── swap.go
└── types.go

bank/
├── penalty.go    # 银行罚单
└── types.go

other/
├── car.go        # 汽车销量
└── types.go
```

### 函数清单

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `fx_pair_quote` | 外汇对报价 | - | P2 |
| `fx_spot_quote` | 外汇现货报价 | - | P2 |
| `fx_swap_quote` | 外汇掉期报价 | - | P2 |
| `fx_c_swap_cm` | 人民币掉期 | 中国货币网 | P2 |
| `fx_quote_baidu` | 外汇报价(百度) | 百度 | P2 |
| `currency_pair_map` | 货币对映射 | - | P3 |
| `bank_fjcf_table_detail` | 银行罚单详情 | - | P3 |
| `bank_fjcf_total_num` | 银行罚单数量 | - | P3 |
| `bank_fjcf_total_page` | 银行罚单页数 | - | P3 |
| `bank_fjcf_page_url` | 银行罚单URL | - | P3 |
| `car_market_total_cpca` | 汽车销量总量 | 乘联会 | P3 |
| `car_market_man_rank_cpca` | 厂商排名 | 乘联会 | P3 |
| `car_market_segment_cpca` | 细分市场 | 乘联会 | P3 |
| `car_market_fuel_cpca` | 燃料类型 | 乘联会 | P3 |
| `car_market_country_cpca` | 国别销量 | 乘联会 | P3 |
| `car_market_cate_cpca` | 车型分类 | 乘联会 | P3 |
| `car_sale_rank_gasgoo` | 汽车销量排名 | 盖世汽车 | P3 |

---

## 十、futures_derivative 期货衍生模块 (15 函数)

### Go 文件结构
```
futures_derivative/
├── contract_info.go  # 合约信息
├── hog.go            # 生猪期货
├── main.go           # 主力合约
└── types.go
```

### 函数清单

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `futures_contract_info_shfe` | 上期所合约 | 上期所 | P1 |
| `futures_contract_info_dce` | 大商所合约 | 大商所 | P1 |
| `futures_contract_info_czce` | 郑商所合约 | 郑商所 | P1 |
| `futures_contract_info_cffex` | 中金所合约 | 中金所 | P1 |
| `futures_contract_info_ine` | 能源中心合约 | 上期能源 | P2 |
| `futures_contract_info_gfex` | 广期所合约 | 广期所 | P2 |
| `futures_display_main_sina` | 主力合约展示 | 新浪财经 | P1 |
| `futures_main_sina` | 主力合约 | 新浪财经 | P1 |
| `futures_hold_pos_sina` | 持仓(新浪) | 新浪财经 | P2 |
| `futures_hog_core` | 生猪核心 | - | P2 |
| `futures_hog_cost` | 生猪成本 | - | P2 |
| `futures_hog_supply` | 生猪供应 | - | P2 |
| `futures_spot_sys` | 现货系统 | - | P2 |
| `match_main_contract` | 匹配主力合约 | - | P2 |
| `zh_subscribe_exchange_symbol` | 订阅交易所品种 | - | P3 |

---

## 十一、类型定义示例

```go
// 外汇报价
type ForexQuote struct {
    Pair       string    `json:"pair"`        // 货币对
    Bid        float64   `json:"bid"`         // 买入价
    Ask        float64   `json:"ask"`         // 卖出价
    Mid        float64   `json:"mid"`         // 中间价
    Change     float64   `json:"change"`      // 涨跌
    ChangePct  float64   `json:"change_pct"`  // 涨跌幅
    Time       time.Time `json:"time"`
}

// 碳排放
type CarbonPrice struct {
    Date       time.Time `json:"date"`
    Price      float64   `json:"price"`       // 价格
    Volume     float64   `json:"volume"`      // 成交量
    Amount     float64   `json:"amount"`      // 成交额
    Market     string    `json:"market"`      // 市场
}

// REITs
type REITsQuote struct {
    Code       string    `json:"code"`
    Name       string    `json:"name"`
    Price      float64   `json:"price"`
    Change     float64   `json:"change"`
    ChangePct  float64   `json:"change_pct"`
    Volume     int64     `json:"volume"`
    Yield      float64   `json:"yield"`       // 分派率
    Time       time.Time `json:"time"`
}
```

---

*文档版本: v1.0*
*最后更新: 2026-01-18*
