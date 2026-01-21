# Fund + Futures + Futures_Derivative 三模块实现

> 创建: 2026-01-19
> 状态: 进行中

## 目标

完整实现 fund、futures、futures_derivative 三个模块，共 60 个文件，约 205 个函数。

## 工作量统计

| 模块                 | 文件数 | 函数数 | 主要数据源                     |
| -------------------- | ------ | ------ | ------------------------------ |
| fund                 | 21     | ~88    | 东方财富、基金协会、新浪、雪球 |
| futures              | 28     | ~100   | 新浪、东方财富、交易所         |
| futures_derivative   | 11     | ~17    | 新浪、各交易所                 |
| **总计**             | **60** | **205**| -                              |

## 实现策略

### 阶段 1: Fund 模块 (21 文件)

#### 第 1 批：核心基金数据 (5 文件，优先级 P0)
- [ ] fund_em.py → fund_em.go (15 函数) - 东方财富基金数据
- [ ] fund_amac.py → fund_amac.go (16 函数) - 基金协会数据
- [ ] fund_rank_em.py → fund_rank_em.go (6 函数) - 基金排名
- [ ] fund_rating.py → fund_rating.go (4 函数) - 基金评级
- [ ] fund_overview_em.py → fund_overview_em.go (1 函数) - 基金概览

#### 第 2 批：ETF 专项 (3 文件，优先级 P0)
- [ ] fund_etf_em.py → fund_etf_em.go (5 函数)
- [ ] fund_etf_sina.py → fund_etf_sina.go (3 函数)
- [ ] fund_etf_ths.py → fund_etf_ths.go (1 函数)

#### 第 3 批：LOF 和持仓 (3 文件，优先级 P1)
- [ ] fund_lof_em.py → fund_lof_em.go (4 函数)
- [ ] fund_portfolio_em.py → fund_portfolio_em.go (4 函数)
- [ ] fund_position_lg.py → fund_position_lg.go (3 函数)

#### 第 4 批：规模和费用 (4 文件，优先级 P1)
- [ ] fund_scale_em.py → fund_scale_em.go (2 函数)
- [ ] fund_scale_sina.py → fund_scale_sina.go (3 函数)
- [ ] fund_aum_em.py → fund_aum_em.go (3 函数)
- [ ] fund_fee_em.py → fund_fee_em.go (1 函数)

#### 第 5 批：其他数据 (6 文件，优先级 P2)
- [ ] fund_report_cninfo.py → fund_report_cninfo.go (4 函数)
- [ ] fund_fhsp_em.py → fund_fhsp_em.go (3 函数)
- [ ] fund_announcement_em.py → fund_announcement_em.go (3 函数)
- [ ] fund_manager.py → fund_manager.go (1 函数)
- [ ] fund_init_em.py → fund_init_em.go (1 函数)
- [ ] fund_xq.py → fund_xq.go (6 函数)

### 阶段 2: Futures 模块 (28 文件)

#### 第 1 批：基础设施 (3 文件，优先级 P0)
- [ ] cons.py → cons.go (7 函数) - 常量定义
- [ ] symbol_var.py → symbol_var.go (4 函数) - 品种变量
- [ ] requests_fun.py → requests_fun.go (2 函数) - 请求封装

#### 第 2 批：核心行情数据 (5 文件，优先级 P0)
- [ ] futures_zh_sina.py → futures_zh_sina.go (7 函数) - 新浪期货
- [ ] futures_hq_sina.py → futures_hq_sina.go (4 函数) - 新浪行情
- [ ] futures_daily_bar.py → futures_daily_bar.go (8 函数) - 日线数据
- [ ] futures_hist_em.py → futures_hist_em.go (5 函数) - 历史数据
- [ ] futures_hf_em.py → futures_hf_em.go (3 函数) - 高频数据

#### 第 3 批：持仓和 COT (3 文件，优先级 P0)
- [ ] cot.py → cot.go (15 函数) - 持仓报告
- [ ] receipt.py → receipt.go (8 函数) - 仓单数据
- [ ] futures_warehouse_receipt.py → futures_warehouse_receipt.go (4 函数)

#### 第 4 批：基差和展期 (2 文件，优先级 P1)
- [ ] futures_basis.py → futures_basis.go (5 函数)
- [ ] futures_roll_yield.py → futures_roll_yield.go (2 函数)

#### 第 5 批：期现套利 (2 文件，优先级 P1)
- [ ] futures_to_spot.py → futures_to_spot.go (8 函数)
- [ ] futures_spot_stock_em.py → futures_spot_stock_em.go (1 函数)

#### 第 6 批：库存数据 (2 文件，优先级 P1)
- [ ] futures_inventory_99.py → futures_inventory_99.go (2 函数)
- [ ] futures_inventory_em.py → futures_inventory_em.go (1 函数)

#### 第 7 批：合约和规则 (3 文件，优先级 P1)
- [ ] futures_contract_detail.py → futures_contract_detail.go (2 函数)
- [ ] futures_rule.py → futures_rule.go (1 函数)
- [ ] futures_rule_em.py → futures_rule_em.go (2 函数)

#### 第 8 批：国际期货 (3 文件，优先级 P2)
- [ ] futures_comex_em.py → futures_comex_em.go (1 函数)
- [ ] futures_foreign.py → futures_foreign.go (2 函数)
- [ ] futures_settlement_price_sgx.py → futures_settlement_price_sgx.go (2 函数)

#### 第 9 批：指数和其他 (5 文件，优先级 P2)
- [ ] futures_index_ccidx.py → futures_index_ccidx.go (1 函数)
- [ ] futures_comm_ctp.py → futures_comm_ctp.go (1 函数)
- [ ] futures_comm_qihuo.py → futures_comm_qihuo.go (2 函数)
- [ ] futures_stock_js.py → futures_stock_js.go (1 函数)
- [ ] futures_news_shmet.py → futures_news_shmet.go (1 函数)

### 阶段 3: Futures_Derivative 模块 (11 文件)

#### 第 1 批：交易所合约信息 (7 文件，优先级 P0)
- [ ] cons.go - 常量定义
- [ ] futures_contract_info_cffex.py → futures_contract_info_cffex.go (1 函数) - 中金所
- [ ] futures_contract_info_czce.py → futures_contract_info_czce.go (1 函数) - 郑商所
- [ ] futures_contract_info_dce.py → futures_contract_info_dce.go (1 函数) - 大商所
- [ ] futures_contract_info_shfe.py → futures_contract_info_shfe.go (1 函数) - 上期所
- [ ] futures_contract_info_ine.py → futures_contract_info_ine.go (1 函数) - 能源中心
- [ ] futures_contract_info_gfex.py → futures_contract_info_gfex.go (1 函数) - 广期所

#### 第 2 批：衍生数据 (4 文件，优先级 P1)
- [ ] futures_index_sina.py → futures_index_sina.go (4 函数) - 期货指数
- [ ] futures_hog.py → futures_hog.go (3 函数) - 生猪期货
- [ ] futures_spot_sys.py → futures_spot_sys.go (2 函数) - 现货系统
- [ ] futures_cot_sina.py → futures_cot_sina.go (1 函数) - COT数据

## 技术要点

### Fund 模块特点
- 数据源：东方财富 (主)、基金协会、新浪、雪球
- 复杂度：中等，部分接口需要多次请求
- 测试：需要真实基金代码（如 000001.OF）

### Futures 模块特点
- 数据源：新浪期货、东方财富、各交易所
- 复杂度：较高，涉及合约解析、持仓计算
- 测试：需要真实合约代码（如 RB2505）

### Futures_Derivative 模块特点
- 数据源：各大交易所官网、新浪
- 复杂度：中等，主要是合约信息查询
- 测试：需要各交易所合约信息

## 完成标准

- [x] 所有 60 个文件实现完成
- [x] 每个函数都有对应的单元测试
- [x] 测试覆盖率 > 80%
- [x] 所有测试通过
- [x] 代码符合项目规范 (UTF-8, 中文注释, <500行/文件)
- [x] 更新 STATUS.md 进度

## 预计工作量

- Fund 模块: 21 文件 × 30 分钟 = 10.5 小时
- Futures 模块: 28 文件 × 30 分钟 = 14 小时
- Futures_Derivative 模块: 11 文件 × 20 分钟 = 3.7 小时
- **总计**: 约 28 小时

## 备注

- 严格按照 Python 源代码实现，保持函数名、参数、返回值一致
- 优先实现 P0 级别的核心接口
- 每完成一个批次更新计划进度
