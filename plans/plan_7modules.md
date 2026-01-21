# 7个模块实现计划

> 创建时间: 2026-01-20
> 完成时间: 2026-01-20
> 状态: ✅ 已完成

## 模块清单

| 模块     | 文件数 | 需实现函数                                                                                                 | 状态      |
| -------- | ------ | ---------------------------------------------------------------------------------------------------------- | --------- |
| pro      | 3      | pro_api, DataApi类, cons常量                                                                               | ✅ 已完成 |
| qdii     | 1      | qdii_e_index_jsl, qdii_e_comm_jsl, qdii_a_index_jsl                                                        | ✅ 已完成 |
| qhkc     | 1      | qhkc_index (通过pro_api访问)                                                                               | ✅ 已完成 |
| qhkc_web | 3      | 网页版数据接口                                                                                             | ✅ 已完成 |
| rate     | 1      | 银行间利率数据                                                                                             | ✅ 已完成 |
| reits    | 1      | reits_realtime_em, reits_hist_em                                                                           | ✅ 已完成 |
| spot     | 3      | spot_price_qh, spot_hist_sge, spot_symbol_table_sge, spot_quotations_sge, spot_golden/silver_benchmark_sge | ✅ 已完成 |

## 详细实现清单

### 1. pro 模块 (3文件)

| 文件        | 函数/结构     | 描述              |
| ----------- | ------------- | ----------------- |
| cons.go     | 常量定义      | TOKEN相关常量     |
| client.go   | DataApi结构体 | 奇货可查API客户端 |
| data_pro.go | ProApi()      | 初始化pro API     |

### 2. qdii 模块 (1文件 -> 3函数)

| 文件        | 函数            | 描述                |
| ----------- | --------------- | ------------------- |
| qdii_jsl.go | QdiiEIndexJsl() | 集思录-欧美指数QDII |
| qdii_jsl.go | QdiiECommJsl()  | 集思录-欧美商品QDII |
| qdii_jsl.go | QdiiAIndexJsl() | 集思录-亚洲指数QDII |

### 3. qhkc 模块 (1文件)

| 文件          | 函数            | 描述                      |
| ------------- | --------------- | ------------------------- |
| qhkc_index.go | 通过pro模块访问 | 奇货可查指数数据(需token) |

### 4. qhkc_web 模块 (3文件)

| 文件        | 函数           | 描述                |
| ----------- | -------------- | ------------------- |
| qhkc_web.go | 网页版数据接口 | 无需token的公开数据 |

### 5. rate 模块 (1文件)

| 文件              | 函数               | 描述           |
| ----------------- | ------------------ | -------------- |
| interbank_rate.go | RateInterbankLpr() | 银行间利率数据 |

### 6. reits 模块 (1文件 -> 2函数)

| 文件        | 函数              | 描述                  |
| ----------- | ----------------- | --------------------- |
| reits_em.go | ReitsRealtimeEm() | 东方财富REITs实时行情 |
| reits_em.go | ReitsHistEm()     | 东方财富REITs历史行情 |

### 7. spot 模块 (3文件 -> 6函数)

| 文件        | 函数                     | 描述                   |
| ----------- | ------------------------ | ---------------------- |
| spot_qh.go  | SpotPriceQh()            | 99期货现货走势         |
| spot_qh.go  | SpotPriceTableQh()       | 99期货品种表           |
| spot_sge.go | SpotHistSge()            | 上海黄金交易所历史数据 |
| spot_sge.go | SpotSymbolTableSge()     | 上海黄金交易所品种表   |
| spot_sge.go | SpotQuotationsSge()      | 上海黄金交易所实时行情 |
| spot_sge.go | SpotGoldenBenchmarkSge() | 上海金基准价           |
| spot_sge.go | SpotSilverBenchmarkSge() | 上海银基准价           |

## 执行顺序

1. ✅ 分析Python源码和文档
2. ✅ 实现 pro 模块
3. ✅ 实现 qdii 模块
4. ✅ 实现 qhkc 模块
5. ✅ 实现 qhkc_web 模块
6. ✅ 实现 rate 模块
7. ✅ 实现 reits 模块
8. ✅ 实现 spot 模块
9. ✅ 编译验证
10. ✅ 更新 STATUS.md
