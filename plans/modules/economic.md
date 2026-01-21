# economic 模块计划

> 模块: 宏观经济数据
> 函数数量: 215 个
> 优先级: P1
> 预计周期: 第13-14周

---

## 一、模块概述

宏观经济模块提供中国、美国、欧洲、日本等主要经济体的宏观经济指标，包括 GDP、CPI、PPI、PMI、利率、就业、贸易等数据。

---

## 二、Go 文件结构

```
economic/
├── china.go             # 中国宏观数据
├── china_money.go       # 中国货币金融
├── china_real.go        # 中国实体经济
├── usa.go               # 美国宏观数据
├── euro.go              # 欧洲宏观数据
├── uk.go                # 英国宏观数据
├── germany.go           # 德国宏观数据
├── japan.go             # 日本宏观数据
├── australia.go         # 澳大利亚数据
├── canada.go            # 加拿大数据
├── swiss.go             # 瑞士数据
├── hk.go                # 香港宏观数据
├── global.go            # 全球综合数据
├── central_bank.go      # 央行利率
├── commodity.go         # 商品数据
├── shipping.go          # 航运指数
└── types.go             # 类型定义
```

---

## 三、函数清单

### 3.1 中国宏观数据 (china.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `macro_china_gdp` | 中国GDP | 金十数据 | P0 |
| `macro_china_gdp_yearly` | 中国GDP年度 | 金十数据 | P0 |
| `macro_china_cpi` | 中国CPI | 金十数据 | P0 |
| `macro_china_cpi_monthly` | 中国CPI月度 | 金十数据 | P0 |
| `macro_china_cpi_yearly` | 中国CPI年度 | 金十数据 | P0 |
| `macro_china_ppi` | 中国PPI | 金十数据 | P0 |
| `macro_china_ppi_yearly` | 中国PPI年度 | 金十数据 | P0 |
| `macro_china_pmi` | 中国PMI | 金十数据 | P0 |
| `macro_china_pmi_yearly` | 中国PMI年度 | 金十数据 | P0 |
| `macro_china_non_man_pmi` | 中国非制造业PMI | 金十数据 | P1 |
| `macro_china_cx_pmi_yearly` | 财新PMI | 金十数据 | P1 |
| `macro_china_cx_services_pmi_yearly` | 财新服务业PMI | 金十数据 | P1 |
| `macro_china_exports_yoy` | 出口同比 | 金十数据 | P0 |
| `macro_china_imports_yoy` | 进口同比 | 金十数据 | P0 |
| `macro_china_trade_balance` | 贸易差额 | 金十数据 | P0 |
| `macro_china_industrial_production_yoy` | 工业增加值 | 金十数据 | P0 |
| `macro_china_urban_unemployment` | 城镇失业率 | 金十数据 | P1 |
| `macro_china_consumer_goods_retail` | 社会消费品零售 | 金十数据 | P0 |
| `macro_china_fdi` | 外商直接投资 | 金十数据 | P1 |

### 3.2 中国货币金融 (china_money.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `macro_china_m2_yearly` | M2货币供应 | 金十数据 | P0 |
| `macro_china_money_supply` | 货币供应量 | 金十数据 | P0 |
| `macro_china_supply_of_money` | 货币供应详情 | 金十数据 | P1 |
| `macro_china_lpr` | LPR利率 | 金十数据 | P0 |
| `macro_china_shibor_all` | Shibor利率 | 金十数据 | P0 |
| `macro_china_new_financial_credit` | 新增人民币贷款 | 金十数据 | P0 |
| `macro_china_shrzgm` | 社会融资规模 | 金十数据 | P0 |
| `macro_china_fx_gold` | 外汇和黄金储备 | 金十数据 | P1 |
| `macro_china_fx_reserves_yearly` | 外汇储备年度 | 金十数据 | P1 |
| `macro_china_foreign_exchange_gold` | 外汇黄金详情 | 金十数据 | P1 |
| `macro_china_central_bank_balance` | 央行资产负债表 | 金十数据 | P2 |
| `macro_china_reserve_requirement_ratio` | 存款准备金率 | 金十数据 | P1 |
| `macro_china_market_margin_sh` | 上海融资融券 | 金十数据 | P2 |
| `macro_china_market_margin_sz` | 深圳融资融券 | 金十数据 | P2 |
| `macro_china_stock_market_cap` | 股市市值 | 金十数据 | P2 |
| `macro_china_rmb` | 人民币汇率 | 金十数据 | P1 |
| `macro_china_bank_financing` | 银行理财 | 金十数据 | P2 |
| `macro_china_insurance` | 保险数据 | 金十数据 | P2 |
| `macro_china_insurance_income` | 保险收入 | 金十数据 | P2 |
| `macro_rmb_deposit` | 人民币存款 | 金十数据 | P2 |
| `macro_rmb_loan` | 人民币贷款 | 金十数据 | P2 |

### 3.3 中国实体经济 (china_real.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `macro_china_gdzctz` | 固定资产投资 | 金十数据 | P0 |
| `macro_china_gyzjz` | 工业增加值 | 金十数据 | P0 |
| `macro_china_hgjck` | 海关进出口 | 金十数据 | P1 |
| `macro_china_czsr` | 财政收入 | 金十数据 | P1 |
| `macro_china_whxd` | 外汇信贷 | 金十数据 | P2 |
| `macro_china_wbck` | 外币存款 | 金十数据 | P2 |
| `macro_china_xfzxx` | 消费者信心 | 金十数据 | P1 |
| `macro_china_enterprise_boom_index` | 企业景气指数 | 金十数据 | P2 |
| `macro_china_real_estate` | 房地产数据 | 金十数据 | P1 |
| `macro_china_new_house_price` | 新房价格 | 金十数据 | P1 |
| `macro_china_society_electricity` | 社会用电量 | 金十数据 | P2 |
| `macro_china_society_traffic_volume` | 社会货运量 | 金十数据 | P2 |
| `macro_china_passenger_load_factor` | 客运量 | 金十数据 | P2 |
| `macro_china_postal_telecommunicational` | 邮电业务 | 金十数据 | P3 |
| `macro_china_international_tourism_fx` | 国际旅游外汇 | 金十数据 | P3 |
| `macro_china_mobile_number` | 移动用户数 | 金十数据 | P3 |
| `macro_china_national_tax_receipts` | 国家税收 | 金十数据 | P2 |
| `macro_china_vegetable_basket` | 菜篮子价格 | 金十数据 | P2 |
| `macro_china_agricultural_product` | 农产品价格 | 金十数据 | P2 |
| `macro_china_agricultural_index` | 农产品指数 | 金十数据 | P2 |
| `macro_china_energy_index` | 能源价格指数 | 金十数据 | P2 |
| `macro_china_commodity_price_index` | 商品价格指数 | 金十数据 | P2 |
| `macro_china_construction_index` | 建材价格指数 | 金十数据 | P2 |
| `macro_china_construction_price_index` | 建筑价格指数 | 金十数据 | P2 |
| `macro_china_lpi_index` | 物流景气指数 | 金十数据 | P2 |
| `macro_china_bdti_index` | 油轮运价指数 | 金十数据 | P2 |
| `macro_china_bsi_index` | 超灵便型指数 | 金十数据 | P2 |
| `macro_china_qyspjg` | 企业商品价格 | 金十数据 | P2 |
| `macro_china_retail_price_index` | 零售价格指数 | 金十数据 | P2 |
| `macro_china_daily_energy` | 日度能源 | 金十数据 | P2 |
| `macro_china_freight_index` | 运价指数 | 金十数据 | P2 |
| `macro_china_au_report` | 黄金报告 | 金十数据 | P2 |
| `macro_china_yw_electronic_index` | 义乌电子指数 | 金十数据 | P3 |
| `macro_china_nbs_nation` | 统计局全国 | 统计局 | P1 |
| `macro_china_nbs_region` | 统计局地区 | 统计局 | P1 |
| `macro_cnbs` | 中国银保监 | 银保监 | P2 |

### 3.4 香港宏观 (hk.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `macro_china_hk_gbp` | 香港GDP | 金十数据 | P1 |
| `macro_china_hk_gbp_ratio` | 香港GDP增速 | 金十数据 | P1 |
| `macro_china_hk_cpi` | 香港CPI | 金十数据 | P1 |
| `macro_china_hk_cpi_ratio` | 香港CPI增速 | 金十数据 | P1 |
| `macro_china_hk_core` | 香港核心通胀 | 金十数据 | P2 |
| `macro_china_hk_ppi` | 香港PPI | 金十数据 | P2 |
| `macro_china_hk_rate_of_unemployment` | 香港失业率 | 金十数据 | P1 |
| `macro_china_hk_trade_diff_ratio` | 香港贸易差额 | 金十数据 | P2 |
| `macro_china_hk_building_volume` | 香港楼宇成交量 | 金十数据 | P2 |
| `macro_china_hk_building_amount` | 香港楼宇成交额 | 金十数据 | P2 |
| `macro_china_hk_market_info` | 香港市场信息 | 金十数据 | P2 |

### 3.5 美国宏观数据 (usa.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `macro_usa_gdp_monthly` | 美国GDP | 金十数据 | P0 |
| `macro_usa_cpi_monthly` | 美国CPI月度 | 金十数据 | P0 |
| `macro_usa_cpi_yoy` | 美国CPI同比 | 金十数据 | P0 |
| `macro_usa_core_cpi_monthly` | 美国核心CPI | 金十数据 | P0 |
| `macro_usa_ppi` | 美国PPI | 金十数据 | P0 |
| `macro_usa_core_ppi` | 美国核心PPI | 金十数据 | P1 |
| `macro_usa_pmi` | 美国PMI | 金十数据 | P0 |
| `macro_usa_ism_pmi` | ISM制造业PMI | 金十数据 | P0 |
| `macro_usa_ism_non_pmi` | ISM非制造业PMI | 金十数据 | P0 |
| `macro_usa_services_pmi` | 服务业PMI | 金十数据 | P1 |
| `macro_usa_non_farm` | 非农就业 | 金十数据 | P0 |
| `macro_usa_unemployment_rate` | 美国失业率 | 金十数据 | P0 |
| `macro_usa_adp_employment` | ADP就业 | 金十数据 | P0 |
| `macro_usa_initial_jobless` | 初请失业金 | 金十数据 | P0 |
| `macro_usa_job_cuts` | 挑战者裁员 | 金十数据 | P2 |
| `macro_usa_trade_balance` | 美国贸易差额 | 金十数据 | P0 |
| `macro_usa_current_account` | 美国经常账户 | 金十数据 | P1 |
| `macro_usa_retail_sales` | 美国零售销售 | 金十数据 | P0 |
| `macro_usa_personal_spending` | 个人消费支出 | 金十数据 | P1 |
| `macro_usa_real_consumer_spending` | 实际消费支出 | 金十数据 | P2 |
| `macro_usa_core_pce_price` | 核心PCE物价 | 金十数据 | P0 |
| `macro_usa_industrial_production` | 工业产出 | 金十数据 | P1 |
| `macro_usa_durable_goods_orders` | 耐用品订单 | 金十数据 | P1 |
| `macro_usa_factory_orders` | 工厂订单 | 金十数据 | P2 |
| `macro_usa_business_inventories` | 商业库存 | 金十数据 | P2 |
| `macro_usa_cb_consumer_confidence` | 消费者信心 | 金十数据 | P0 |
| `macro_usa_michigan_consumer_sentiment` | 密歇根消费者信心 | 金十数据 | P1 |
| `macro_usa_nfib_small_business` | 小企业乐观指数 | 金十数据 | P2 |
| `macro_usa_lmci` | 劳动力市场指数 | 金十数据 | P2 |
| `macro_usa_building_permits` | 建筑许可 | 金十数据 | P1 |
| `macro_usa_house_starts` | 新屋开工 | 金十数据 | P1 |
| `macro_usa_new_home_sales` | 新屋销售 | 金十数据 | P1 |
| `macro_usa_exist_home_sales` | 成屋销售 | 金十数据 | P1 |
| `macro_usa_pending_home_sales` | 成屋签约销售 | 金十数据 | P2 |
| `macro_usa_phs` | PHS指数 | 金十数据 | P2 |
| `macro_usa_nahb_house_market_index` | NAHB房屋市场指数 | 金十数据 | P2 |
| `macro_usa_house_price_index` | 房价指数 | 金十数据 | P2 |
| `macro_usa_spcs20` | 标普/CS房价指数 | 金十数据 | P2 |
| `macro_usa_import_price` | 进口物价 | 金十数据 | P2 |
| `macro_usa_export_price` | 出口物价 | 金十数据 | P2 |
| `macro_usa_eia_crude_rate` | EIA原油库存 | 金十数据 | P1 |
| `macro_usa_api_crude_stock` | API原油库存 | 金十数据 | P1 |
| `macro_usa_crude_inner` | 原油产量 | 金十数据 | P2 |
| `macro_usa_rig_count` | 钻井数 | 金十数据 | P2 |
| `macro_usa_cftc_nc_holding` | CFTC非商业持仓 | 金十数据 | P1 |
| `macro_usa_cftc_c_holding` | CFTC商业持仓 | 金十数据 | P1 |
| `macro_usa_cftc_merchant_currency_holding` | CFTC货币持仓 | 金十数据 | P2 |
| `macro_usa_cftc_merchant_goods_holding` | CFTC商品持仓 | 金十数据 | P2 |
| `macro_usa_cme_merchant_goods_holding` | CME持仓 | 金十数据 | P2 |

### 3.6 欧洲宏观数据 (euro.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `macro_euro_gdp_yoy` | 欧元区GDP | 金十数据 | P0 |
| `macro_euro_cpi_mom` | 欧元区CPI月度 | 金十数据 | P0 |
| `macro_euro_cpi_yoy` | 欧元区CPI同比 | 金十数据 | P0 |
| `macro_euro_ppi_mom` | 欧元区PPI | 金十数据 | P1 |
| `macro_euro_manufacturing_pmi` | 制造业PMI | 金十数据 | P0 |
| `macro_euro_services_pmi` | 服务业PMI | 金十数据 | P0 |
| `macro_euro_unemployment_rate_mom` | 失业率 | 金十数据 | P0 |
| `macro_euro_employment_change_qoq` | 就业人数变化 | 金十数据 | P1 |
| `macro_euro_trade_balance` | 贸易差额 | 金十数据 | P1 |
| `macro_euro_current_account_mom` | 经常账户 | 金十数据 | P2 |
| `macro_euro_retail_sales_mom` | 零售销售 | 金十数据 | P1 |
| `macro_euro_industrial_production_mom` | 工业产出 | 金十数据 | P1 |
| `macro_euro_sentix_investor_confidence` | 投资者信心 | 金十数据 | P2 |
| `macro_euro_zew_economic_sentiment` | ZEW经济景气 | 金十数据 | P1 |
| `macro_euro_lme_holding` | LME持仓 | 金十数据 | P2 |
| `macro_euro_lme_stock` | LME库存 | 金十数据 | P2 |

### 3.7 英国宏观数据 (uk.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `macro_uk_gdp_quarterly` | 英国GDP季度 | 金十数据 | P1 |
| `macro_uk_gdp_yearly` | 英国GDP年度 | 金十数据 | P1 |
| `macro_uk_cpi_monthly` | 英国CPI月度 | 金十数据 | P1 |
| `macro_uk_cpi_yearly` | 英国CPI年度 | 金十数据 | P1 |
| `macro_uk_core_cpi_monthly` | 核心CPI月度 | 金十数据 | P1 |
| `macro_uk_core_cpi_yearly` | 核心CPI年度 | 金十数据 | P1 |
| `macro_uk_core` | 核心通胀 | 金十数据 | P2 |
| `macro_uk_unemployment_rate` | 失业率 | 金十数据 | P1 |
| `macro_uk_trade` | 贸易差额 | 金十数据 | P2 |
| `macro_uk_retail_monthly` | 零售月度 | 金十数据 | P1 |
| `macro_uk_retail_yearly` | 零售年度 | 金十数据 | P1 |
| `macro_uk_halifax_monthly` | 房价月度 | 金十数据 | P2 |
| `macro_uk_halifax_yearly` | 房价年度 | 金十数据 | P2 |
| `macro_uk_rightmove_monthly` | Rightmove房价月度 | 金十数据 | P2 |
| `macro_uk_rightmove_yearly` | Rightmove房价年度 | 金十数据 | P2 |
| `macro_uk_bank_rate` | 英国央行利率 | 金十数据 | P0 |

### 3.8 德国宏观数据 (germany.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `macro_germany_gdp` | 德国GDP | 金十数据 | P1 |
| `macro_germany_cpi_monthly` | 德国CPI月度 | 金十数据 | P1 |
| `macro_germany_cpi_yearly` | 德国CPI年度 | 金十数据 | P1 |
| `macro_germany_core` | 核心通胀 | 金十数据 | P2 |
| `macro_germany_retail_sale_monthly` | 零售销售月度 | 金十数据 | P2 |
| `macro_germany_retail_sale_yearly` | 零售销售年度 | 金十数据 | P2 |
| `macro_germany_trade_adjusted` | 贸易差额 | 金十数据 | P2 |
| `macro_germany_ifo` | IFO商业景气 | 金十数据 | P1 |
| `macro_germany_zew` | ZEW经济景气 | 金十数据 | P1 |

### 3.9 日本宏观数据 (japan.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `macro_japan_cpi_yearly` | 日本CPI | 金十数据 | P1 |
| `macro_japan_core_cpi_yearly` | 日本核心CPI | 金十数据 | P1 |
| `macro_japan_core` | 核心通胀 | 金十数据 | P2 |
| `macro_japan_unemployment_rate` | 失业率 | 金十数据 | P1 |
| `macro_japan_head_indicator` | 领先指标 | 金十数据 | P2 |
| `macro_japan_bank_rate` | 日本央行利率 | 金十数据 | P0 |

### 3.10 其他国家 (australia.go / canada.go / swiss.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `macro_australia_cpi_quarterly` | 澳大利亚CPI季度 | 金十数据 | P2 |
| `macro_australia_cpi_yearly` | 澳大利亚CPI年度 | 金十数据 | P2 |
| `macro_australia_ppi_quarterly` | 澳大利亚PPI | 金十数据 | P2 |
| `macro_australia_unemployment_rate` | 澳大利亚失业率 | 金十数据 | P2 |
| `macro_australia_retail_rate_monthly` | 澳大利亚零售 | 金十数据 | P2 |
| `macro_australia_trade` | 澳大利亚贸易 | 金十数据 | P2 |
| `macro_australia_bank_rate` | 澳大利亚利率 | 金十数据 | P1 |
| `macro_canada_gdp_monthly` | 加拿大GDP | 金十数据 | P2 |
| `macro_canada_cpi_monthly` | 加拿大CPI月度 | 金十数据 | P2 |
| `macro_canada_cpi_yearly` | 加拿大CPI年度 | 金十数据 | P2 |
| `macro_canada_core_cpi_monthly` | 加拿大核心CPI月度 | 金十数据 | P2 |
| `macro_canada_core_cpi_yearly` | 加拿大核心CPI年度 | 金十数据 | P2 |
| `macro_canada_unemployment_rate` | 加拿大失业率 | 金十数据 | P2 |
| `macro_canada_retail_rate_monthly` | 加拿大零售 | 金十数据 | P2 |
| `macro_canada_trade` | 加拿大贸易 | 金十数据 | P2 |
| `macro_canada_new_house_rate` | 加拿大新屋价格 | 金十数据 | P2 |
| `macro_canada_bank_rate` | 加拿大利率 | 金十数据 | P1 |
| `macro_swiss_gdp_quarterly` | 瑞士GDP | 金十数据 | P2 |
| `macro_swiss_gdp_yearly` | 瑞士GDP年度 | 金十数据 | P2 |
| `macro_swiss_cpi_yearly` | 瑞士CPI | 金十数据 | P2 |
| `macro_swiss_core` | 瑞士核心通胀 | 金十数据 | P2 |
| `macro_swiss_trade` | 瑞士贸易 | 金十数据 | P2 |
| `macro_swiss_svme` | 瑞士PMI | 金十数据 | P2 |
| `macro_swiss_gbd_bank_rate` | 瑞士央行利率 | 金十数据 | P1 |
| `macro_swiss_gbd_yearly` | 瑞士央行年度 | 金十数据 | P2 |

### 3.11 央行利率 (central_bank.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `macro_bank_usa_interest_rate` | 美联储利率 | 金十数据 | P0 |
| `macro_bank_euro_interest_rate` | 欧央行利率 | 金十数据 | P0 |
| `macro_bank_english_interest_rate` | 英央行利率 | 金十数据 | P0 |
| `macro_bank_japan_interest_rate` | 日央行利率 | 金十数据 | P0 |
| `macro_bank_china_interest_rate` | 中国央行利率 | 金十数据 | P0 |
| `macro_bank_australia_interest_rate` | 澳央行利率 | 金十数据 | P1 |
| `macro_bank_newzealand_interest_rate` | 新西兰利率 | 金十数据 | P2 |
| `macro_bank_switzerland_interest_rate` | 瑞士央行利率 | 金十数据 | P1 |
| `macro_bank_russia_interest_rate` | 俄罗斯利率 | 金十数据 | P2 |
| `macro_bank_india_interest_rate` | 印度央行利率 | 金十数据 | P2 |
| `macro_bank_brazil_interest_rate` | 巴西央行利率 | 金十数据 | P2 |

### 3.12 商品数据 (commodity.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `macro_cons_gold` | 黄金消费 | 金十数据 | P2 |
| `macro_cons_silver` | 白银消费 | 金十数据 | P2 |
| `macro_cons_opec_month` | OPEC月报 | 金十数据 | P2 |

### 3.13 航运指数 (shipping.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `macro_shipping_bdi` | 波罗的海干散货指数 | 金十数据 | P1 |
| `macro_shipping_bci` | 海岬型指数 | 金十数据 | P2 |
| `macro_shipping_bpi` | 巴拿马型指数 | 金十数据 | P2 |
| `macro_shipping_bcti` | 原油运输指数 | 金十数据 | P2 |
| `macro_global_sox_index` | 费城半导体指数 | 金十数据 | P1 |

### 3.14 其他功能

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `crypto_js_spot` | 加密货币实时 | 金十数据 | P2 |
| `macro_fx_sentiment` | 外汇情绪 | 金十数据 | P2 |
| `macro_info_ws` | 财经资讯 | 金十数据 | P3 |
| `macro_stock_finance` | 股票财经 | 金十数据 | P3 |

---

## 四、类型定义示例

```go
// types.go
package economic

import "time"

// EconomicIndicator 经济指标
type EconomicIndicator struct {
    Date     time.Time `json:"date"`
    Actual   float64   `json:"actual"`   // 实际值
    Previous float64   `json:"previous"` // 前值
    Forecast float64   `json:"forecast"` // 预期值
    Country  string    `json:"country"`  // 国家
    Name     string    `json:"name"`     // 指标名称
}

// GDP 国内生产总值
type GDP struct {
    Date     time.Time `json:"date"`
    Value    float64   `json:"value"`    // GDP值
    YoY      float64   `json:"yoy"`      // 同比
    QoQ      float64   `json:"qoq"`      // 环比
    Country  string    `json:"country"`
}

// CPI 消费者物价指数
type CPI struct {
    Date     time.Time `json:"date"`
    Value    float64   `json:"value"`
    YoY      float64   `json:"yoy"`
    MoM      float64   `json:"mom"`
    Core     float64   `json:"core"`     // 核心CPI
    Country  string    `json:"country"`
}

// PMI 采购经理指数
type PMI struct {
    Date         time.Time `json:"date"`
    Manufacturing float64  `json:"manufacturing"` // 制造业
    NonMfg       float64   `json:"non_mfg"`       // 非制造业
    Composite    float64   `json:"composite"`     // 综合
    Country      string    `json:"country"`
}

// InterestRate 利率
type InterestRate struct {
    Date       time.Time `json:"date"`
    Rate       float64   `json:"rate"`        // 利率
    Change     float64   `json:"change"`      // 变动
    Country    string    `json:"country"`
    CentralBank string   `json:"central_bank"` // 央行
}

// Employment 就业数据
type Employment struct {
    Date            time.Time `json:"date"`
    NonFarm         int64     `json:"non_farm"`         // 非农就业
    UnemploymentRate float64  `json:"unemployment_rate"` // 失业率
    ADPEmployment   int64     `json:"adp_employment"`   // ADP就业
    Country         string    `json:"country"`
}
```

---

## 五、实现优先级

### P0 - 第一批实现 (核心指标)

1. 中国GDP、CPI、PPI、PMI
2. 美国GDP、CPI、非农、失业率
3. 主要央行利率
4. LPR、M2、社融

### P1 - 第二批实现 (常用指标)

1. 贸易数据
2. 工业产出
3. 消费者信心
4. 房地产数据
5. 航运指数

### P2 - 第三批实现 (其他指标)

1. 各类价格指数
2. 小国经济数据
3. CFTC持仓
4. LME数据

---

*文档版本: v1.0*
*最后更新: 2026-01-18*
