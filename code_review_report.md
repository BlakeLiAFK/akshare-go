# akshare-go 代码审查报告

> 生成时间: 2026-01-21 01:05:43

## 审查范围

逐个模块对比 Python 源代码和 Go 实现的完整性：
- 函数数量对比
- 测试覆盖率
- 缺失函数清单

---

## 模块: air

### 统计

| 项目 | 数量 |
|------|------|
| Python 函数 | 9 |
| Go 函数 | 8 |
| 测试函数 | 8 |

❌ **状态: 函数缺失**（函数: 8/9, 测试: 8/��

### Python 函数清单
```
air_quality_hebei
has_month_data
air_city_table
air_quality_watch_point
air_quality_hist
air_quality_rank
sunrise_city_list
sunrise_daily
sunrise_monthly
```

### Go 函数清单
```
AirQualityHebei
AirCityTable
AirQualityWatchPoint
AirQualityHist
AirQualityRank
SunriseCityList
SunriseDaily
SunriseMonthly
```

---

## 模块: article

### 统计

| 项目 | 数量 |
|------|------|
| Python 函数 | 7 |
| Go 函数 | 7 |
| 测试函数 | 7 |

✅ **状态: 完整**（函数: 7/7, 测试: 7/��

---

## 模块: bank

### 统计

| 项目 | 数量 |
|------|------|
| Python 函数 | 4 |
| Go 函数 | 4 |
| 测试函数 | 5 |

✅ **状态: 完整**（函数: 4/4, 测试: 5/��

---

## 模块: bond

### 统计

| 项目 | 数量 |
|------|------|
| Python 函数 | 42 |
| Go 函数 | 42 |
| 测试函数 | 20 |

⚠️ **状态: 测试不足**（函数: 42/42, 测试: 20/��

---

## 模块: cal

### 统计

| 项目 | 数量 |
|------|------|
| Python 函数 | 3 |
| Go 函数 | 3 |
| 测试函数 | 5 |

✅ **状态: 完整**（函数: 3/3, 测试: 5/��

---

## 模块: crypto

### 统计

| 项目 | 数量 |
|------|------|
| Python 函数 | 2 |
| Go 函数 | 2 |
| 测试函数 | 3 |

✅ **状态: 完整**（函数: 2/2, 测试: 3/��

---

## 模块: currency

### 统计

| 项目 | 数量 |
|------|------|
| Python 函数 | 7 |
| Go 函数 | 7 |
| 测试函数 | 8 |

✅ **状态: 完整**（函数: 7/7, 测试: 8/��

---

## 模块: economic

### 统计

| 项目 | 数量 |
|------|------|
| Python 函数 | 230 |
| Go 函数 | 230 |
| 测试函数 | 111 |

⚠️ **状态: 测试不足**（函数: 230/230, 测试: 111/��

---

## 模块: energy

### 统计

| 项目 | 数量 |
|------|------|
| Python 函数 | 8 |
| Go 函数 | 8 |
| 测试函数 | 0 |

⚠️ **状态: 测试不足**（函数: 8/8, 测试: 0/��

---

## 模块: event

### 统计

| 项目 | 数量 |
|------|------|
| Python 函数 | 2 |
| Go 函数 | 6 |
| 测试函数 | 0 |

⚠️ **状态: 测试不足**（函数: 6/2, 测试: 0/��

---

## 模块: forex

### 统计

| 项目 | 数量 |
|------|------|
| Python 函数 | 2 |
| Go 函数 | 3 |
| 测试函数 | 0 |

⚠️ **状态: 测试不足**（函数: 3/2, 测试: 0/��

---

## 模块: fortune

### 统计

| 项目 | 数量 |
|------|------|
| Python 函数 | 6 |
| Go 函数 | 6 |
| 测试函数 | 0 |

⚠️ **状态: 测试不足**（函数: 6/6, 测试: 0/��

---

## 模块: fund

### 统计

| 项目 | 数量 |
|------|------|
| Python 函数 | 84 |
| Go 函数 | 86 |
| 测试函数 | 76 |

⚠️ **状态: 测试不足**（函数: 86/84, 测试: 76/��

---

## 模块: futures

### 统计

| 项目 | 数量 |
|------|------|
| Python 函数 | 78 |
| Go 函数 | 53 |
| 测试函数 | 20 |

❌ **状态: 函数缺失**（函数: 53/78, 测试: 20/��

### Python 函数清单
```
get_rank_sum_daily
get_rank_sum
get_shfe_rank_table
get_rank_table_czce
get_dce_rank_table
get_cffex_rank_table
futures_dce_position_rank
futures_dce_position_rank_other
futures_gfex_position_rank
futures_spot_price_daily
futures_spot_price
futures_spot_price_previous
futures_comex_inventory
futures_fees_info
futures_comm_info
futures_contract_detail
futures_contract_detail_em
get_cffex_daily
get_gfex_daily
get_ine_daily
get_czce_daily
get_shfe_daily
get_dce_daily
get_futures_daily
futures_foreign_hist
futures_foreign_detail
futures_global_spot_em
futures_global_hist_em
futures_hist_table_em
futures_hist_em
futures_foreign_commodity_subscribe_exchange_symbol
futures_hq_subscribe_exchange_symbol
futures_foreign_commodity_realtime
futures_index_ccidx
futures_inventory_99
futures_inventory_em
futures_news_shmet
get_roll_yield
get_roll_yield_bar
futures_rule
futures_rule_em
futures_trading_hours_em
futures_settlement_price_sgx
futures_spot_stock
futures_stock_shfe_js
futures_to_spot_shfe
futures_delivery_dce
futures_to_spot_dce
futures_delivery_match_dce
futures_to_spot_czce
futures_delivery_match_czce
futures_delivery_czce
futures_delivery_shfe
futures_warehouse_receipt_czce
futures_warehouse_receipt_dce
futures_shfe_warehouse_receipt
futures_gfex_warehouse_receipt
futures_symbol_mark
futures_zh_realtime
zh_subscribe_exchange_symbol
match_main_contract
futures_zh_spot
futures_zh_minute_sina
futures_zh_daily_sina
get_dce_receipt
get_shfe_receipt_1
get_shfe_receipt_2
get_czce_receipt_1
get_czce_receipt_2
get_czce_receipt_3
get_gfex_receipt
get_receipt
requests_link
pandas_read_html_link
symbol_varieties
symbol_market
find_chinese
chinese_to_english
```

### Go 函数清单
```
FuturesSHFERankTable
FuturesCZCERankTable
FuturesCFFEXRankTable
FuturesDCEPositionRank
FuturesGFEXPositionRank
FuturesRankSum
FuturesSpotPrice
FuturesSpotPriceDaily
FuturesSpotPricePrevious
FuturesComexInventoryFunc
FuturesContractDetailFunc
FuturesContractDetailEMFunc
GetCFFEXDaily
GetGFEXDaily
GetINEDaily
GetSHFEDaily
GetDCEDaily
GetCZCEDaily
GetFuturesDaily
FuturesForeignHistEM
FuturesHFSpotEMFunc
FuturesHFMinuteEMFunc
FuturesHistTableEMFunc
FuturesHistEMFunc
FuturesHQSubscribeExchangeSymbol
FuturesForeignCommoditySubscribeExchangeSymbol
FuturesForeignCommodityRealtimeFunc
FuturesIndexCCIDX
FuturesInventoryEMFunc
FuturesNewsSHMET
GetRollYield
GetRollYieldBar
FuturesRuleFunc
FuturesRuleEMFunc
FuturesToSpotSHFE
FuturesDeliveryDCE
FuturesToSpotDCE
FuturesToSpotCZCE
FuturesDeliveryCZCE
FuturesDeliverySHFE
FuturesWarehouseReceiptCZCE
FuturesWarehouseReceiptDCE
FuturesSHFEWarehouseReceipt
FuturesGFEXWarehouseReceipt
FuturesSymbolMarkFunc
FuturesZhRealtimeFunc
FuturesZhSpotFunc
FuturesZhMinuteSina
FuturesZhDailySina
SymbolVarieties
SymbolMarket
FindChinese
ChineseToEnglish
```

---

## 模块: futures_derivative

### 统计

| 项目 | 数量 |
|------|------|
| Python 函数 | 15 |
| Go 函数 | 15 |
| 测试函数 | 15 |

✅ **状态: 完整**（函数: 15/15, 测试: 15/��

---

## 模块: fx

### 统计

| 项目 | 数量 |
|------|------|
| Python 函数 | 6 |
| Go 函数 | 6 |
| 测试函数 | 0 |

⚠️ **状态: 测试不足**（函数: 6/6, 测试: 0/��

---

## 模块: hf

### 统计

| 项目 | 数量 |
|------|------|
| Python 函数 | 1 |
| Go 函数 | 1 |
| 测试函数 | 0 |

⚠️ **状态: 测试不足**（函数: 1/1, 测试: 0/��

---

## 模块: index

### 统计

| 项目 | 数量 |
|------|------|
| Python 函数 | 98 |
| Go 函数 | 105 |
| 测试函数 | 75 |

⚠️ **状态: 测试不足**（函数: 105/98, 测试: 75/��

---

## 模块: interest_rate

### 统计

| 项目 | 数量 |
|------|------|
| Python 函数 | 1 |
| Go 函数 | 1 |
| 测试函数 | 0 |

⚠️ **状态: 测试不足**（函数: 1/1, 测试: 0/��

---

## 模块: movie

### 统计

| 项目 | 数量 |
|------|------|
| Python 函数 | 16 |
| Go 函数 | 12 |
| 测试函数 | 1 |

❌ **状态: 函数缺失**（函数: 12/16, 测试: 1/��

### Python 函数清单
```
decrypt
business_value_artist
online_value_artist
get_current_week
decrypt
movie_boxoffice_realtime
movie_boxoffice_daily
movie_boxoffice_weekly
movie_boxoffice_monthly
movie_boxoffice_yearly
movie_boxoffice_yearly_first_week
movie_boxoffice_cinema_daily
movie_boxoffice_cinema_weekly
decrypt
video_tv
video_variety_show
```

### Go 函数清单
```
BusinessValueArtist
OnlineValueArtist
MovieBoxofficeRealtime
MovieBoxofficeDaily
MovieBoxofficeWeekly
MovieBoxofficeMonthly
MovieBoxofficeYearly
MovieBoxofficeYearlyFirstWeek
MovieBoxofficeCinemaDaily
MovieBoxofficeCinemaWeekly
VideoTv
VideoVarietyShow
```

---

## 模块: news

### 统计

| 项目 | 数量 |
|------|------|
| Python 函数 | 6 |
| Go 函数 | 6 |
| 测试函数 | 8 |

✅ **状态: 完整**（函数: 6/6, 测试: 8/��

---

## 模块: nlp

### 统计

| 项目 | 数量 |
|------|------|
| Python 函数 | 2 |
| Go 函数 | 6 |
| 测试函数 | 8 |

✅ **状态: 完整**（函数: 6/2, 测试: 8/��

---

## 模块: option

### 统计

| 项目 | 数量 |
|------|------|
| Python 函数 | 47 |
| Go 函数 | 39 |
| 测试函数 | 10 |

❌ **状态: 函数缺失**（函数: 39/47, 测试: 10/��

### Python 函数清单
```
option_comm_symbol
option_comm_info
option_hist_dce
option_hist_czce
option_hist_shfe
option_vol_shfe
option_hist_gfex
option_vol_gfex
option_commodity_contract_sina
option_commodity_contract_table_sina
option_commodity_hist_sina
option_contract_info_ctp
option_current_day_sse
option_current_day_szse
option_hist_yearly_czce
option_daily_stats_sse
option_daily_stats_szse
option_current_em
option_current_cffex_em
option_finance_sse_underlying
option_finance_board
option_cffex_sz50_list_sina
option_cffex_hs300_list_sina
option_cffex_zz1000_list_sina
option_cffex_sz50_spot_sina
option_cffex_hs300_spot_sina
option_cffex_zz1000_spot_sina
option_cffex_sz50_daily_sina
option_cffex_hs300_daily_sina
option_cffex_zz1000_daily_sina
option_sse_list_sina
option_sse_expire_day_sina
option_sse_codes_sina
option_sse_spot_price_sina
option_sse_underlying_spot_price_sina
option_sse_greeks_sina
option_sse_minute_sina
option_sse_daily_sina
option_finance_minute_sina
option_minute_em
option_lhb_em
option_margin_symbol
option_margin
option_premium_analysis_em
option_risk_analysis_em
option_risk_indicator_sse
option_value_analysis_em
```

### Go 函数清单
```
ConvertDate
GetJSONPath
GetCalendar
LastTradingDay
GetLatestDataDate
OptionPremiumAnalysisEm
OptionRiskAnalysisEm
OptionValueAnalysisEm
OptionCommSymbol
OptionCommInfo
OptionHistDce
GetDceOptionSymbols
OptionHistShfe
OptionCommoditySina
OptionCommoditySinaAll
OptionFinanceSina
GetSinaOptionSymbols
OptionContractInfoCtp
OptionCurrentDaySse
OptionCurrentDaySzse
OptionHistYearlyCzce
GetCzceOptionSymbols
OptionDailyStatsSse
OptionDailyStatsSzse
OptionCurrentEm
OptionCurrentCffexEm
OptionFinance
OptionFinanceSSE
OptionFinanceSZSE
OptionFinanceDaily
GetFinanceOptionSymbols
OptionFinanceSinaDetail
OptionFinanceSinaChain
OptionFinanceSinaVolatility
GetSinaFinanceOptionSymbols
OptionLhbEm
OptionMarginSymbol
OptionMargin
OptionRiskIndicatorSse
```

---

## 模块: other

### 统计

| 项目 | 数量 |
|------|------|
| Python 函数 | 7 |
| Go 函数 | 7 |
| 测试函数 | 7 |

✅ **状态: 完整**（函数: 7/7, 测试: 7/��

---

## 模块: pro

### 统计

| 项目 | 数量 |
|------|------|
| Python 函数 | 1 |
| Go 函数 | 3 |
| 测试函数 | 0 |

⚠️ **状态: 测试不足**（函数: 3/1, 测试: 0/��

---

## 模块: qdii

### 统计

| 项目 | 数量 |
|------|------|
| Python 函数 | 3 |
| Go 函数 | 3 |
| 测试函数 | 0 |

⚠️ **状态: 测试不足**（函数: 3/3, 测试: 0/��

---

## 模块: qhkc

### 统计

| 项目 | 数量 |
|------|------|
| Python 函数 | 0 |
| Go 函数 | 5 |
| 测试函数 | 0 |

⚠️ **状态: 测试不足**（函数: 5/0, 测试: 0/��

---

## 模块: qhkc_web

### 统计

| 项目 | 数量 |
|------|------|
| Python 函数 | 10 |
| Go 函数 | 3 |
| 测试函数 | 0 |

❌ **状态: 函数缺失**（函数: 3/10, 测试: 0/��

### Python 函数清单
```
get_qhkc_fund_bs
get_qhkc_fund_position
get_qhkc_fund_position_change
get_qhkc_fund_money_change
get_qhkc_index
get_qhkc_index_trend
get_qhkc_index_profit_loss
qhkc_tool_foreign
qhkc_tool_nebula
qhkc_tool_gdp
```

### Go 函数清单
```
QhkcWebBasis
QhkcWebInventory
QhkcWebProfit
```

---

## 模块: rate

### 统计

| 项目 | 数量 |
|------|------|
| Python 函数 | 2 |
| Go 函数 | 4 |
| 测试函数 | 5 |

✅ **状态: 完整**（函数: 4/2, 测试: 5/��

---

## 模块: reits

### 统计

| 项目 | 数量 |
|------|------|
| Python 函数 | 3 |
| Go 函数 | 3 |
| 测试函数 | 3 |

✅ **状态: 完整**（函数: 3/3, 测试: 3/��

---

## 模块: spot

### 统计

| 项目 | 数量 |
|------|------|
| Python 函数 | 15 |
| Go 函数 | 7 |
| 测试函数 | 0 |

❌ **状态: 函数缺失**（函数: 7/15, 测试: 0/��

### Python 函数清单
```
spot_hog_soozhu
spot_hog_year_trend_soozhu
spot_hog_lean_price_soozhu
spot_hog_three_way_soozhu
spot_hog_crossbred_soozhu
spot_corn_price_soozhu
spot_soybean_price_soozhu
spot_mixed_feed_soozhu
spot_price_table_qh
spot_price_qh
spot_symbol_table_sge
spot_quotations_sge
spot_hist_sge
spot_golden_benchmark_sge
spot_silver_benchmark_sge
```

### Go 函数清单
```
SpotPriceQh
SpotPriceTableQh
SpotHistSge
SpotSymbolTableSge
SpotQuotationsSge
SpotGoldenBenchmarkSge
SpotSilverBenchmarkSge
```

---

## 模块: stock

### 统计

| 项目 | 数量 |
|------|------|
| Python 函数 | 130 |
| Go 函数 | 167 |
| 测试函数 | 3 |

⚠️ **状态: 测试不足**（函数: 167/130, 测试: 3/��

---

## 模块: stock_a

### 统计

| 项目 | 数量 |
|------|------|
| Python 函数 | 6 |
| Go 函数 | 3 |
| 测试函数 | 4 |

❌ **状态: 函数缺失**（函数: 3/6, 测试: 4/��

### Python 函数清单
```
process_concept_board_data
stock_board_concept_name_em
process_fund_flow_data
stock_individual_fund_flow_rank
process_data
stock_zh_a_spot_em
```

### Go 函数清单
```
StockBoardConceptNameEm
StockIndividualFundFlowRank
StockZhASpotEm
```

---

## 模块: stock_feature

### 统计

| 项目 | 数量 |
|------|------|
| Python 函数 | 207 |
| Go 函数 | 146 |
| 测试函数 | 27 |

❌ **状态: 函数缺失**（函数: 146/207, 测试: 27/��

### Python 函数清单
```
stock_a_below_net_asset_statistics
stock_a_high_low_statistics
get_cookie_csrf
get_token_lg
stock_hk_indicator_eniu
stock_market_pe_lg
stock_index_pe_lg
stock_market_pb_lg
stock_index_pb_lg
stock_account_statistics_em
stock_a_all_pb
stock_analyst_rank_em
stock_analyst_detail_em
stock_board_concept_name_ths
stock_board_concept_info_ths
stock_board_concept_index_ths
stock_board_concept_summary_ths
stock_board_industry_name_ths
stock_board_industry_info_ths
stock_board_industry_index_ths
stock_xgsr_ths
stock_ipo_benefit_ths
stock_board_industry_summary_ths
stock_buffett_index_lg
stock_classify_board
stock_classify_sina
stock_comment_em
stock_comment_detail_zlkp_jgcyd_em
stock_comment_detail_zhpj_lspf_em
stock_comment_detail_scrd_focus_em
stock_comment_detail_scrd_desire_em
stock_concept_cons_futu
stock_a_congestion_lg
stock_cyq_em
stock_zh_a_disclosure_report_cninfo
stock_zh_a_disclosure_relation_cninfo
stock_dxsyl_em
stock_xgsglb_em
stock_ebs_lg
stock_esg_msci_sina
stock_esg_rft_sina
stock_esg_rate_sina
stock_esg_zd_sina
stock_esg_hz_sina
stock_fhps_em
stock_fhps_detail_em
stock_fhps_detail_ths
stock_fund_flow_individual
stock_fund_flow_concept
stock_fund_flow_industry
stock_fund_flow_big_deal
stock_gddh_em
stock_gdfx_free_holding_statistics_em
stock_gdfx_holding_statistics_em
stock_gdfx_free_holding_change_em
stock_gdfx_holding_change_em
stock_gdfx_free_top_10_em
stock_gdfx_top_10_em
stock_gdfx_free_holding_detail_em
stock_gdfx_holding_detail_em
stock_gdfx_free_holding_analyse_em
stock_gdfx_holding_analyse_em
stock_gdfx_free_holding_teamwork_em
stock_gdfx_holding_teamwork_em
stock_zh_a_gdhs
stock_zh_a_gdhs_detail_em
stock_ggcg_em
stock_gpzy_profile_em
stock_gpzy_pledge_ratio_em
stock_gpzy_pledge_ratio_detail_em
stock_gpzy_distribute_statistics_company_em
stock_gpzy_distribute_statistics_bank_em
stock_gpzy_industry_data_em
stock_a_gxl_lg
stock_hk_gxl_lg
stock_zh_a_spot_em
stock_sh_a_spot_em
stock_sz_a_spot_em
stock_bj_a_spot_em
stock_new_a_spot_em
stock_cy_a_spot_em
stock_kc_a_spot_em
stock_zh_ab_comparison_em
stock_zh_b_spot_em
stock_zh_a_hist
stock_zh_a_hist_min_em
stock_zh_a_hist_pre_min_em
stock_hk_spot_em
stock_hk_main_board_spot_em
stock_hk_hist
stock_hk_hist_min_em
stock_us_spot_em
stock_us_hist
stock_us_hist_min_em
stock_zh_a_hist_tx
stock_hk_valuation_baidu
stock_hot_follow_xq
stock_hot_tweet_xq
stock_hot_deal_xq
stock_hsgt_fund_flow_summary_em
stock_hk_ggt_components_em
stock_hsgt_hold_stock_em
stock_hsgt_stock_statistics_em
stock_hsgt_institution_statistics_em
stock_hsgt_hist_em
stock_hsgt_board_rank_em
stock_hsgt_individual_em
stock_hsgt_individual_detail_em
stock_sgt_settlement_exchange_rate_szse
stock_sgt_reference_exchange_rate_szse
stock_sgt_reference_exchange_rate_sse
stock_sgt_settlement_exchange_rate_sse
stock_hsgt_fund_min_em
stock_info_cjzc_em
stock_info_global_em
stock_info_global_sina
stock_info_global_futu
stock_info_global_ths
stock_info_global_cls
stock_inner_trade_xq
stock_irm_cninfo
stock_irm_ans_cninfo
stock_jgdy_tj_em
stock_jgdy_detail_em
stock_lh_yyb_most
stock_lh_yyb_capital
stock_lh_yyb_control
stock_lhb_detail_em
stock_lhb_stock_statistic_em
stock_lhb_jgmmtj_em
stock_lhb_jgstatistic_em
stock_lhb_hyyyb_em
stock_lhb_yybph_em
stock_lhb_traderstatistic_em
stock_lhb_stock_detail_date_em
stock_lhb_stock_detail_em
stock_lhb_yyb_detail_em
stock_lhb_detail_daily_sina
stock_lhb_ggtj_sina
stock_lhb_yytj_sina
stock_lhb_jgzz_sina
stock_lhb_jgmx_sina
stock_margin_account_info
stock_margin_ratio_pa
stock_margin_sse
stock_margin_detail_sse
stock_margin_underlying_info_szse
stock_margin_szse
stock_margin_detail_szse
stock_market_activity_legu
stock_changes_em
stock_board_change_em
stock_qsjy_em
stock_zcfz_em
stock_zcfz_bj_em
stock_lrb_em
stock_xjll_em
stock_research_report_em
stock_sns_sseinfo
stock_sy_profile_em
stock_sy_yq_em
stock_sy_jz_em
stock_sy_em
stock_sy_hy_em
stock_rank_cxg_ths
stock_rank_cxd_ths
stock_rank_lxsz_ths
stock_rank_lxxd_ths
stock_rank_cxfl_ths
stock_rank_cxsl_ths
stock_rank_xstp_ths
stock_rank_xxtp_ths
stock_rank_ljqs_ths
stock_rank_ljqd_ths
stock_rank_xzjp_ths
stock_tfp_em
stock_balance_sheet_by_report_em
stock_balance_sheet_by_yearly_em
stock_profit_sheet_by_report_em
stock_profit_sheet_by_yearly_em
stock_profit_sheet_by_quarterly_em
stock_cash_flow_sheet_by_report_em
stock_cash_flow_sheet_by_yearly_em
stock_cash_flow_sheet_by_quarterly_em
stock_balance_sheet_by_report_delisted_em
stock_profit_sheet_by_report_delisted_em
stock_cash_flow_sheet_by_report_delisted_em
stock_a_ttm_lyr
stock_us_valuation_baidu
stock_value_em
stock_yjbb_em
stock_report_disclosure
stock_yjkb_em
stock_yjyg_em
stock_yysj_em
stock_yzxdr_em
stock_zdhtmx_em
stock_qbzf_em
stock_pg_em
stock_zh_valuation_baidu
stock_zh_vote_baidu
stock_zt_pool_em
stock_zt_pool_previous_em
stock_zt_pool_strong_em
stock_zt_pool_sub_new_em
stock_zt_pool_zbgc_em
stock_zt_pool_dtgc_em
```

### Go 函数清单
```
StockABelowNetAssetStatistics
StockAIndicator
StockAPe
StockAPb
StockAPeAndPb
StockAHighLow
StockAccountEm
StockInfoChange
StockAAllPb
StockAnalystEm
StockAnalystDetailEm
StockBoardConceptThs
StockBoardConceptConstThs
StockBoardIndustryThs
StockBoardIndustryConstThs
StockBuffettIndexLg
StockClassifySina
StockClassifyConstSina
StockCommentEm
StockConceptFutu
StockConceptConstFutu
StockACongestionLg
StockCyqEm
StockCyqDetailEm
StockDisclosureCninfo
StockIrmCninfo
StockDxsylEm
StockEbsLg
StockEsgMsciSina
StockEsgRateSina
StockEsgRftSina
StockEsgZdSina
StockAccountStatisticsEm
StockGdfxFreeHoldingStatisticsEm
StockGdfxHoldingStatisticsEm
StockGpzyProfileEm
StockGpzyPledgeRatioEm
StockZhASpotEm
StockShASpotEm
StockSzASpotEm
StockBjASpotEm
StockCybASpotEm
StockKcbASpotEm
StockNewSpotEm
StockHsgtSpotEm
StockZhBSpotEm
StockHsgtFundFlowSummaryEm
StockHsgtHoldStockEm
StockHkGgtComponentsEm
GetTokenLg
GetCookieCsrf
StockHkIndicatorEniu
StockLhbDetailEm
StockLhbStockStatisticEm
StockMarginAccountInfo
StockChangesEm
StockBoardChangeEm
StockTfpEm
StockValueEm
StockYjbbEm
StockZtPoolEm
StockZtPoolPreviousEm
StockDtPoolEm
StockZbPoolEm
StockFhpsEm
StockFhpsDetailThs
StockFundFlow
StockFundFlowIndividual
StockGddhEm
StockGdfxHoldingDetailEm
StockGdfxHoldingAnalyseEm
StockGdfxFreeHoldingDetailEm
StockGdfxFreeHoldingAnalyseEm
StockZhAGdhs
StockZhAGdhsDetailEm
StockGgcgEm
StockGpzyEm
StockGpzyDetailEm
StockAGxlLg
StockHistEm
StockZhAHistTx
StockHkValuationBaidu
StockHotFollowXq
StockHotTweetXq
StockHotDealXq
StockHsgtEmFlow
StockHsgtMinEm
StockHsgtExchangeRate
StockSgtSettlementExchangeRateSzse
StockSgtReferenceExchangeRateSzse
StockSgtReferenceExchangeRateSse
StockSgtSettlementExchangeRateSse
StockHsgtFundMinEm
StockInfoCjzcEm
StockInfoGlobalEm
StockInfoGlobalSina
StockInfoGlobalFutu
StockInfoGlobalThs
StockInfoGlobalCls
StockInnerTradeXq
StockHotXq
StockJgdyEm
StockJgdyDetailEm
StockLhYybMost
StockLhYybCapital
StockLhYybControl
StockLhbGgtjSina
StockLhbJgmmtjEm
StockLhbJgstatisticEm
StockLhbStockDetailEm
StockLhbTraderstatisticEm
StockLhbYybDetailEm
StockLhbYybphEm
StockMarginEm
StockMarginSse
StockMarginSzse
StockMarketActivityLegu
StockPankouEm
StockQsjyEm
StockReportEm
StockResearchReportEm
StockSnsSseinfo
StockSyEm
StockRankCxgThs
StockRankCxdThs
StockRankLxszThs
StockRankLxxdThs
StockRankCxflThs
StockRankCxslThs
StockRankXstpThs
StockRankXxtpThs
StockRankLjqsThs
StockRankLjqdThs
StockRankXzjpThs
StockThreeReportEm
StockATtmLyr
StockUsValuationBaidu
StockReportDisclosure
StockYjygEm
StockYjkbEm
StockYzxdrEm
StockZdhtmxEm
StockQbzfEm
StockPgEm
StockZhValuationBaidu
StockZhVoteBaidu
```

---

## 模块: stock_fundamental

### 统计

| 项目 | 数量 |
|------|------|
| Python 函数 | 56 |
| Go 函数 | 55 |
| 测试函数 | 60 |

❌ **状态: 函数缺失**（函数: 55/56, 测试: 60/��

### Python 函数清单
```
stock_individual_basic_info_xq
stock_individual_basic_info_us_xq
stock_individual_basic_info_hk_xq
stock_financial_hk_report_em
stock_financial_hk_analysis_indicator_em
stock_financial_report_sina
stock_financial_abstract
stock_financial_analysis_indicator_em
stock_financial_analysis_indicator
stock_history_dividend
stock_history_dividend_detail
stock_ipo_info
stock_add_stock
stock_restricted_release_queue_sina
stock_circulate_stock_holder
stock_fund_stock_holder
stock_main_stock_holder
stock_financial_abstract_ths
stock_financial_debt_ths
stock_financial_benefit_ths
stock_financial_cash_ths
stock_financial_abstract_new_ths
stock_financial_debt_new_ths
stock_financial_benefit_new_ths
stock_financial_cash_new_ths
stock_management_change_ths
stock_shareholder_change_ths
stock_financial_us_report_em
stock_financial_us_analysis_indicator_em
stock_zh_a_gbjg_em
stock_institute_hold
stock_institute_hold_detail
stock_ipo_declare_em
stock_ipo_review_em
stock_ipo_tutor_em
stock_kcb_detail_renewal
stock_kcb_renewal
stock_notice_report
stock_profit_forecast_em
stock_hk_profit_forecast_et
stock_profit_forecast_ths
stock_institute_recommend
stock_institute_recommend_detail
stock_register_all_em
stock_register_kcb
stock_register_cyb
stock_register_bj
stock_register_sh
stock_register_sz
stock_register_db
stock_restricted_release_summary_em
stock_restricted_release_detail_em
stock_restricted_release_queue_em
stock_restricted_release_stockholder_em
stock_zygc_em
stock_zyjs_ths
```

### Go 函数清单
```
StockIndividualBasicInfoXq
StockIndividualBasicInfoUsXq
StockIndividualBasicInfoHkXq
StockFinancialHkReportEm
StockFinancialHkAnalysisIndicatorEm
StockFinancialReportSina
StockFinancialAbstract
StockFinancialAnalysisIndicatorEm
StockHistoryDividend
StockHistoryDividendDetail
StockIpoInfo
StockMainStockHolder
StockFundStockHolder
StockFinancialAbstractThs
StockFinancialAbstractNewThs
StockFinancialDebtNewThs
StockFinancialBenefitNewThs
StockFinancialCashNewThs
StockManagementChangeThs
StockShareholderChangeThs
StockFinancialDebtThs
StockFinancialBenefitThs
StockFinancialCashThs
StockFinancialUsReportEm
StockFinancialUsAnalysisIndicatorEm
StockZhAGbjgEm
StockZygcEm
StockInstituteHold
StockInstituteHoldDetail
StockIpoDeclareEm
StockIpoReviewEm
StockIpoTutorEm
StockKcbSse
StockNoticeReport
StockProfitForecastEm
StockHkProfitForecastEt
StockProfitForecastThsDetail
StockInstituteRecommend
StockInstituteRecommendDetail
StockRegisterAllEm
StockRegisterKcb
StockRegisterCyb
StockRestrictedReleaseSummaryEm
StockRestrictedReleaseDetailEm
StockRestrictedReleaseQueueEm
GetSupportedSymbolList
StockZyjsThs
StockProfitForecastThs
StockRegisterBj
StockRegisterDb
StockRegisterSh
StockRegisterSz
StockAddStock
StockCirculateStockHolder
StockRestrictedReleaseStockholderEm
```

---

## 模块: tool

### 统计

| 项目 | 数量 |
|------|------|
| Python 函数 | 1 |
| Go 函数 | 1 |
| 测试函数 | 2 |

✅ **状态: 完整**（函数: 1/1, 测试: 2/��

---

## 汇总统计

| 项目 | 数量 | 百分比 |
|------|------|--------|
| 总模块数 | 36 | 100% |
| 完整模块 | 12 | 33% |
| 不完整模块 | 24 | 66% |

