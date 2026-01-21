# 模块完成计划

> 最后更新: 2026-01-20

## 概述

本计划目标是将所有标记为"部分完成"的模块补充完整，确保文件、函数、函数参数与Python源代码严格一致。

## 待实现清单

### 1. Bond模块 (87% → 100%)

#### 缺失文件：bond_cb_sina.py

| Python函数           | 参数        | Go函数            | 状态   |
| -------------------- | ----------- | ----------------- | ------ |
| bond_cb_profile_sina | symbol: str | BondCBProfileSina | 待实现 |
| bond_cb_summary_sina | symbol: str | BondCBSummarySina | 待实现 |

#### 缺失文件：bond_cb_ths.py

| Python函数           | 参数 | Go函数           | 状态   |
| -------------------- | ---- | ---------------- | ------ |
| bond_zh_cov_info_ths | 无   | BondZhCovInfoThs | 待实现 |

#### 缺失文件：bond_info_cm.py

| Python函数          | 参数                                                                                     | Go函数           | 状态   |
| ------------------- | ---------------------------------------------------------------------------------------- | ---------------- | ------ |
| bond_info_cm_query  | symbol: str                                                                              | BondInfoCMQuery  | 待实现 |
| bond_info_cm        | bond_name, bond_code, bond_issue, bond_type, coupon_type, issue_year, underwriter, grade | BondInfoCM       | 待实现 |
| bond_info_detail_cm | symbol: str                                                                              | BondInfoDetailCM | 待实现 |

#### 缺失文件：bond_issue_cninfo.py

| Python函数                         | 参数                 | Go函数                         | 状态   |
| ---------------------------------- | -------------------- | ------------------------------ | ------ |
| bond_treasure_issue_cninfo         | start_date, end_date | BondTreasureIssueCninfo        | 待实现 |
| bond_local_government_issue_cninfo | start_date, end_date | BondLocalGovernmentIssueCninfo | 待实现 |
| bond_corporate_issue_cninfo        | start_date, end_date | BondCorporateIssueCninfo       | 待实现 |
| bond_cov_issue_cninfo              | start_date, end_date | BondCovIssueCninfo             | 待实现 |
| bond_cov_stock_issue_cninfo        | 无                   | BondCovStockIssueCninfo        | 待实现 |

#### 缺失文件：bond_zh_cov.py

| Python函数                 | 参数                                         | Go函数                 | 状态   |
| -------------------------- | -------------------------------------------- | ---------------------- | ------ |
| bond_zh_hs_cov_spot        | 无                                           | BondZhHsCovSpot        | 待实现 |
| bond_zh_hs_cov_daily       | symbol: str                                  | BondZhHsCovDaily       | 待实现 |
| bond_zh_hs_cov_min         | symbol, period, adjust, start_date, end_date | BondZhHsCovMin         | 待实现 |
| bond_zh_hs_cov_pre_min     | symbol: str                                  | BondZhHsCovPreMin      | 待实现 |
| bond_zh_cov                | 无                                           | BondZhCov              | 待实现 |
| bond_cov_comparison        | 无                                           | BondCovComparison      | 待实现 |
| bond_zh_cov_info           | symbol, indicator                            | BondZhCovInfo          | 待实现 |
| bond_zh_cov_value_analysis | symbol: str                                  | BondZhCovValueAnalysis | 待实现 |

### 2. Energy模块 (50% → 100%)

#### 缺失文件：energy_carbon.py

| Python函数             | 参数        | Go函数               | 状态   |
| ---------------------- | ----------- | -------------------- | ------ |
| energy_carbon_domestic | symbol: str | EnergyCarbonDomestic | 待实现 |
| energy_carbon_bj       | 无          | EnergyCarbonBJ       | 待实现 |
| energy_carbon_sz       | 无          | EnergyCarbonSZ       | 待实现 |
| energy_carbon_eu       | 无          | EnergyCarbonEU       | 待实现 |
| energy_carbon_hb       | 无          | EnergyCarbonHB       | 待实现 |
| energy_carbon_gz       | 无          | EnergyCarbonGZ       | 待实现 |

### 3. Event模块 (50% → 100%)

#### 缺失文件：cons.py (常量)

| 常量          | Go常量       | 状态   |
| ------------- | ------------ | ------ |
| province_dict | ProvinceDict | 待实现 |
| city_dict     | CityDict     | 待实现 |

### 4. Forex模块 (50% → 100%)

#### 缺失文件：cons.py (常量)

| 常量              | Go常量          | 状态   |
| ----------------- | --------------- | ------ |
| symbol_market_map | SymbolMarketMap | 待实现 |

### 5. Utils模块 (88% → 100%)

#### func.py中需要的函数

| Python函数           | 参数                      | Go函数             | 状态   |
| -------------------- | ------------------------- | ------------------ | ------ |
| fetch_paginated_data | url, base_params, timeout | FetchPaginatedData | 待实现 |
| set_df_columns       | df, cols                  | SetDFColumns       | 待实现 |

注: tqdm.py 和 multi_decrypt.py 在Go中不需要实现

## 执行顺序

1. ✅ 创建计划文件
2. ✅ Energy模块 - energy_carbon.go (6个函数)
3. ✅ Event模块 - cons.go (已存在完整)
4. ✅ Forex模块 - cons.go (已存在完整)
5. ✅ Utils模块 - func.go (已存在完整)
6. ✅ Bond模块 - bond_cb_sina.go (已在bond_sina.go中实现)
7. ✅ Bond模块 - bond_cb_ths.go (1个函数)
8. ✅ Bond模块 - bond_info_cm.go (已存在完整)
9. ✅ Bond模块 - bond_issue_cninfo.go (已存在完整)
10. ✅ Bond模块 - bond_zh_cov.go (8个函数)
11. ✅ 编译验证通过

## 完成标准

- 所有函数名与Python严格对应
- 所有参数名与Python严格对应
- 返回类型为[]map[string]interface{}或类似结构
