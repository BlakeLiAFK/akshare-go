# Stock 模块对齐计划

> 创建时间: 2026-01-19 23:50

## 1. stock_a 模块分析

### Python 函数 (3个)

| 函数名                          | 参数                   | 返回值    |
| ------------------------------- | ---------------------- | --------- |
| stock_board_concept_name_em     | 无                     | DataFrame |
| stock_individual_fund_flow_rank | indicator: str = "5日" | DataFrame |
| stock_zh_a_spot_em              | 无                     | DataFrame |

### Go 实现 (3个) ✅ 完成

| 函数名                      | 状态 |
| --------------------------- | ---- |
| StockBoardConceptNameEm     | ✅   |
| StockIndividualFundFlowRank | ✅   |
| StockZhASpotEm              | ✅   |

**stock_a 模块: 100% 完成**

---

## 2. stock 模块分析

### Python 公开函数 (约140个)

需要实现的主要函数:

| Python函数                                              | Go函数                    | 状态          |
| ------------------------------------------------------- | ------------------------- | ------------- |
| stock_allotment_cninfo                                  | StockAllotmentCninfo      | ✅            |
| stock_bid_ask_em                                        | StockAskBidEm             | ✅            |
| stock_board_concept_cons_em                             | StockBoardConceptConsEm   | ✅            |
| stock_board_concept_name_em                             | StockBoardConceptNameEm   | ✅            |
| stock_board_concept_spot_em                             | StockBoardConceptSpotEm   | ✅            |
| stock_board_industry_cons_em                            | StockBoardIndustryConsEm  | ✅            |
| stock_board_industry_name_em                            | StockBoardIndustryNameEm  | ✅            |
| stock_board_industry_spot_em                            | StockBoardIndustrySpotEm  | ✅            |
| stock_cg_equity_mortgage_cninfo                         | StockCgEquityMortgageEm   | ✅            |
| stock_cg_guarantee_cninfo                               | StockCgGuaranteeEm        | ✅            |
| stock_cg_lawsuit_cninfo                                 | StockCgLawsuitEm          | ✅            |
| stock_dividend_cninfo                                   | StockDividendCninfo       | ✅            |
| stock_dzjy_mrmx                                         | StockDzjyGgMx             | ⚠️ 名称不一致 |
| stock_dzjy_mrtj                                         | StockDzjyMrTj             | ✅            |
| stock_dzjy_yybph                                        | StockDzjyYybPm            | ⚠️ 名称不一致 |
| stock_gsrl_gsdt_em                                      | StockGsrlEm               | ✅            |
| stock_hk_daily                                          | StockHkDailyEm            | ✅            |
| stock_hk_famous_spot_em                                 | StockHkFamousEm           | ✅            |
| stock_hk_fhpx_detail_ths                                | StockHkFhpxThs            | ✅            |
| stock_hk_hot_rank_em                                    | StockHkHotRankEm          | ✅            |
| stock_hk_spot                                           | StockHkSpotSina           | ✅            |
| stock_hold_control_cninfo                               | StockHoldControlCninfo    | ✅            |
| stock_hold_control_em → stock_hold_management_detail_em | StockHoldControlEm        | ⚠️            |
| stock_hold_num_cninfo                                   | StockHoldNumCninfo        | ✅            |
| stock_hot_rank_em                                       | StockHotRankEm            | ✅            |
| stock_hot_search_baidu                                  | StockHotSearchBaidu       | ✅            |
| stock_hot_up_em                                         | StockHotUpEm              | ✅            |
| stock_individual_fund_flow                              | StockIndividualFundFlowEm | ✅            |
| stock_individual_info_em                                | StockIndividualInfoEm     | ✅            |
| stock_industry_category_cninfo                          | StockIndustryCategory     | ✅            |
| stock_industry_pe_ratio_cninfo                          | StockIndustryPeCninfo     | ✅            |
| stock_info_a_code_name                                  | StockInfoACodeNameEm      | ✅            |
| stock_intraday_em                                       | StockIntradayEm           | ✅            |
| stock_intraday_sina                                     | StockIntradaySina         | ✅            |
| stock_ipo_summary_cninfo                                | StockIpoSummaryCninfo     | ✅            |
| stock_new_ipo_cninfo                                    | StockNewCninfo            | ✅            |
| stock_news_main_cx                                      | StockNewsEm               | ⚠️ 名称不一致 |
| stock_profile_cninfo                                    | StockProfileCninfo        | ✅            |
| stock_rank_forecast_cninfo                              | StockRankForecast         | ✅            |
| stock_repurchase_em                                     | StockRepurchaseEm         | ✅            |
| stock_share_change_cninfo                               | StockShareChangesEm       | ✅            |
| stock_sse_summary                                       | StockSseSummary           | ✅            |
| stock_szse_summary                                      | StockSzseSummary          | ✅            |
| stock_us_daily                                          | StockUsDailyEm            | ✅            |
| stock_us_famous_spot_em                                 | StockUsFamousEm           | ✅            |
| stock_us_spot                                           | StockUsSpotSina           | ✅            |
| stock_weibo_nlp → stock_js_weibo_nlp_time               | StockWeiboNlp             | ✅            |
| stock_zh_a_daily                                        | StockZhAHistSina          | ✅            |
| stock_zh_a_minute                                       | StockZhAMinuteSina        | ✅            |
| stock_zh_a_spot                                         | StockZhASpotSina          | ✅            |
| stock_zh_a_tick_tx_js                                   | StockZhATickTx            | ✅            |
| stock_zh_ah_spot                                        | StockZhAhTx               | ✅            |
| stock_zh_b_spot                                         | StockZhBSpotSina          | ✅            |
| stock_zh_kcb_report_em                                  | StockKcbReport            | ✅            |

### 缺失函数 (需要添加)

| Python函数                          | 参数                                         | 说明                    |
| ----------------------------------- | -------------------------------------------- | ----------------------- |
| stock_board_concept_hist_em         | symbol, period, start_date, end_date, adjust | 概念板块历史行情        |
| stock_board_concept_hist_min_em     | symbol, period                               | 概念板块分钟行情        |
| stock_board_industry_hist_em        | symbol, period, start_date, end_date, adjust | 行业板块历史行情        |
| stock_board_industry_hist_min_em    | symbol, period                               | 行业板块分钟行情        |
| stock_dzjy_hygtj                    | 无                                           | 大宗交易-活跃股统计     |
| stock_dzjy_hyyybtj                  | 无                                           | 大宗交易-活跃营业部统计 |
| stock_dzjy_sctj                     | 无                                           | 大宗交易-市场统计       |
| stock_hk_company_profile_em         | symbol                                       | 港股公司概况            |
| stock_hk_dividend_payout_em         | symbol                                       | 港股派息                |
| stock_hk_financial_indicator_em     | symbol                                       | 港股财务指标            |
| stock_hk_growth_comparison_em       | symbol                                       | 港股成长对比            |
| stock_hk_scale_comparison_em        | symbol                                       | 港股规模对比            |
| stock_hk_security_profile_em        | symbol                                       | 港股证券资料            |
| stock_hk_valuation_comparison_em    | symbol                                       | 港股估值对比            |
| stock_hold_change_cninfo            | symbol                                       | 持股变动                |
| stock_hold_management_detail_cninfo | symbol                                       | 高管持股明细            |
| stock_hold_management_person_em     | symbol                                       | 高管人员持股            |
| stock_hot_keyword_em                | symbol                                       | 热门关键词              |
| stock_hot_rank_detail_em            | symbol                                       | 热度排名详情            |
| stock_hot_rank_detail_realtime_em   | symbol                                       | 热度排名实时详情        |
| stock_hot_rank_latest_em            | symbol                                       | 最新热度排名            |
| stock_hot_rank_relate_em            | symbol                                       | 相关热度排名            |
| stock_hsgt_sh_hk_spot_em            | 无                                           | 沪港通实时行情          |
| stock_industry_change_cninfo        | symbol                                       | 行业变动                |
| stock_industry_clf_hist_sw          | symbol                                       | 申万行业历史            |
| stock_info_bj_name_code             | 无                                           | 北交所股票代码          |
| stock_info_change_name              | symbol                                       | 股票更名历史            |
| stock_info_sh_delist                | 无                                           | 上交所退市股票          |
| stock_info_sh_name_code             | symbol                                       | 上交所股票代码          |
| stock_info_sz_change_name           | symbol                                       | 深交所更名历史          |
| stock_info_sz_delist                | 无                                           | 深交所退市股票          |
| stock_info_sz_name_code             | symbol                                       | 深交所股票代码          |
| stock_main_fund_flow                | 无                                           | 主力资金流向            |
| stock_market_fund_flow              | 无                                           | 市场资金流向            |
| stock_new_gh_cninfo                 | 无                                           | 新股过会                |
| stock_price_js                      | 无                                           | 股票价格JS              |
| stock_report_fund_hold              | symbol, date                                 | 基金持仓                |
| stock_report_fund_hold_detail       | symbol, date                                 | 基金持仓详情            |
| stock_sector_detail                 | sector                                       | 板块详情                |
| stock_sector_fund_flow_hist         | symbol                                       | 板块资金历史            |
| stock_sector_fund_flow_rank         | indicator                                    | 板块资金排名            |
| stock_sector_fund_flow_summary      | indicator                                    | 板块资金汇总            |
| stock_sector_spot                   | indicator                                    | 板块实时行情            |
| stock_share_hold_change_bse         | 无                                           | 北交所股份变动          |
| stock_share_hold_change_sse         | 无                                           | 上交所股份变动          |
| stock_share_hold_change_szse        | 无                                           | 深交所股份变动          |
| stock_sse_deal_daily                | date                                         | 上交所每日成交          |
| stock_staq_net_stop                 | 无                                           | STAQ/NET停牌            |
| stock_szse_area_summary             | date                                         | 深交所地区汇总          |
| stock_szse_sector_summary           | date                                         | 深交所板块汇总          |
| stock_us_pink_spot_em               | 无                                           | 美股粉单实时            |
| stock_xq → stock_individual_spot_xq | symbol                                       | 雪球个股行情            |
| stock_zh_a_cdr_daily                | symbol                                       | CDR日线                 |
| stock_zh_a_new                      | 无                                           | 新股                    |
| stock_zh_a_new_em                   | 无                                           | 新股EM                  |
| stock_zh_a_st_em                    | 无                                           | ST股票                  |
| stock_zh_a_stop_em                  | 无                                           | 停牌股票                |
| stock_zh_ah_daily                   | symbol                                       | AH股日线                |
| stock_zh_ah_name                    | 无                                           | AH股名称                |
| stock_zh_ah_spot_em                 | 无                                           | AH股实时EM              |
| stock_zh_b_daily                    | symbol                                       | B股日线                 |
| stock_zh_b_minute                   | symbol                                       | B股分钟                 |
| stock_zh_dupont_comparison_em       | symbol                                       | 杜邦分析对比            |
| stock_zh_growth_comparison_em       | symbol                                       | 成长对比                |
| stock_zh_kcb_daily                  | symbol                                       | 科创板日线              |
| stock_zh_kcb_spot                   | 无                                           | 科创板实时              |
| stock_zh_scale_comparison_em        | symbol                                       | 规模对比                |
| stock_zh_valuation_comparison_em    | symbol                                       | 估值对比                |

---

## 3. stock_feature 模块分析

### Python 公开函数 (约160个)

Go实现状态对比...

### 已实现 (主要函数)

- StockAIndicator, StockMarginEm, StockHsgtEmFlow, StockHistEm
- StockAccountEm, StockReportEm, StockJgdyEm, StockFundFlow
- StockAnalystEm, StockValueEm, StockZtPoolEm 等

### 缺失函数 (部分列表)

| Python函数                 | 说明               |
| -------------------------- | ------------------ |
| stock_a_pe_and_pb          | A股市盈率市净率    |
| stock_all_pb               | 全市场市净率       |
| stock_board_concept_ths    | 同花顺概念板块     |
| stock_board_industry_ths   | 同花顺行业板块     |
| stock_comment_em           | 东方财富股票评论   |
| stock_cyq_em               | 东方财富筹码分布   |
| stock_dxsyl_em             | 东方财富打新收益率 |
| stock*esg*\*               | ESG相关函数        |
| stock_fhps_detail_ths      | 同花顺分红派送详情 |
| stock_fund_flow_big_deal   | 大单资金流         |
| stock_fund_flow_concept    | 概念资金流         |
| stock_fund_flow_individual | 个股资金流         |
| stock_fund_flow_industry   | 行业资金流         |
| stock_gddh_em              | 股东大会           |
| stock*gdfx*\*              | 股东分析系列       |
| stock_ggcg_em              | 高管持股           |
| stock*gpzy*\*              | 股票质押系列       |
| stock*lhb*\*               | 龙虎榜系列         |
| stock*margin*\*            | 融资融券系列       |
| stock*rank*\*              | 排名系列           |
| stock_yjbb_em              | 业绩报表           |
| stock_yjkb_em              | 业绩快报           |
| stock_yjyg_em              | 业绩预告           |
| stock*zt_pool*\*           | 涨停池系列         |

---

## 4. stock_fundamental 模块分析

### Python 公开函数 (约55个)

Go实现状态对比...

### 已实现函数

- StockFinancialAbstract, StockFinancialReportSina
- StockIpoDeclareEm, StockIpoReviewEm, StockIpoTutorEm
- StockKcbSse, StockProfitForecastEm, StockProfitForecastThs
- StockRegisterAllEm, StockRestrictedRelease\* 等

### 缺失函数

| Python函数                              | 说明           |
| --------------------------------------- | -------------- |
| stock_add_stock                         | 增发股票       |
| stock_circulate_stock_holder            | 流通股股东     |
| stock_kcb_detail_renewal                | 科创板续贷详情 |
| stock_kcb_renewal                       | 科创板续贷     |
| stock_register_bj                       | 北交所注册     |
| stock_register_db                       | 主板注册       |
| stock_register_sh                       | 上交所注册     |
| stock_register_sz                       | 深交所注册     |
| stock_restricted_release_stockholder_em | 限售股解禁股东 |

---

## 5. 执行计划

### 阶段1: 补全 stock 模块缺失函数 (优先级高)

1. 板块历史行情函数 (4个)
2. 大宗交易函数 (3个)
3. 港股函数 (7个)
4. 持股变动函数 (4个)
5. 热度排名函数 (5个)

### 阶段2: 补全 stock_feature 模块缺失函数

1. ESG相关函数 (4个)
2. 股东分析函数 (10个)
3. 龙虎榜函数 (15个)
4. 业绩相关函数 (5个)

### 阶段3: 补全 stock_fundamental 模块缺失函数

1. 注册函数 (4个)
2. 科创板函数 (2个)
3. 限售股函数 (1个)

### 阶段4: 验证和测试

1. 编译验证所有模块
2. 运行现有测试
3. 更新STATUS.md
