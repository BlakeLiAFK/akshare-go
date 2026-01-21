# Bank & Bond 模块完成任务

> 创建: 2026-01-19
> 状态: 进行中

## 目标

完成 bank 和 bond 两个模块的所有剩余函数实现，确保功能完整性和测试覆盖。

## 当前状态

### Bank 模块 (✅ 已完成)
- Python文件数: 2 (bank_cbirc_2020.py, cons.py)
- Go实现数: 4个函数
- 完成度: 100%

已实现函数:
1. ✅ BankFjcfTotalNum - 获取总记录数
2. ✅ BankFjcfTotalPage - 计算总页数
3. ✅ BankFjcfPageUrl - 获取分页数据列表
4. ✅ BankFjcfTableDetail - 获取详细表格数据

### Bond 模块 (🔶 33% -> 100%)
- Python文件数: 15
- Go实现数: 5/41个函数
- 待实现: 36个函数

已实现函数 (5个):
1. ✅ BondCBJSL - 集思录可转债列表
2. ✅ BondCBRedeemJSL - 集思录可转债强赎数据
3. ✅ BondCBAdjLogsJSL - 集思录转股价调整记录
4. ✅ BondCBIndexJSL - 集思录可转债等权指数
5. ✅ BondZHUSRate - 中美国债收益率

## 实施步骤

### 第一批：质押式回购 (bond_buy_back_em.py - 3个函数)
- [ ] bond_sh_buy_back_em - 上证质押式回购
- [ ] bond_sz_buy_back_em - 深证质押式回购
- [ ] bond_buy_back_hist_em - 质押式回购历史数据

### 第二批：新浪财经可转债 (bond_cb_sina.py - 2个函数)
- [ ] bond_cb_profile_sina - 可转债详情资料
- [ ] bond_cb_summary_sina - 可转债债券概况

### 第三批：同花顺可转债 (bond_cb_ths.py - 1个函数)
- [ ] bond_zh_cov_info_ths - 同花顺可转债数据

### 第四批：中国债券信息网 (bond_cbond.py - 2个函数)
- [ ] bond_new_composite_index_cbond - 中债新综合指数
- [ ] bond_composite_index_cbond - 中债综合指数

### 第五批：外汇交易中心 (bond_china.py - 3个函数)
- [ ] bond_spot_quote - 现券市场做市报价
- [ ] bond_spot_deal - 现券市场成交行情
- [ ] bond_china_yield - 中国债券收益率曲线

### 第六批：收盘收益率 (bond_china_money.py - 4个函数)
- [ ] bond_china_close_return_map - 收盘收益率曲线映射
- [ ] bond_china_close_return - 收盘收益率曲线历史数据
- [ ] macro_china_swap_rate - 利率互换曲线历史数据
- [ ] macro_china_bond_public - 公开市场业务交易公告

### 第七批：债券信息 (bond_info_cm.py - 3个函数)
- [ ] bond_info_cm_query - 债券信息查询字段
- [ ] bond_info_cm - 债券信息查询
- [ ] bond_info_detail_cm - 债券详细信息

### 第八批：债券发行 (bond_issue_cninfo.py - 5个函数)
- [ ] bond_treasure_issue_cninfo - 国债发行
- [ ] bond_local_government_issue_cninfo - 地方政府债发行
- [ ] bond_corporate_issue_cninfo - 企业债发行
- [ ] bond_cov_issue_cninfo - 可转债发行
- [ ] bond_cov_stock_issue_cninfo - 可转债转股

### 第九批：银行间市场 (bond_nafmii.py - 1个函数)
- [ ] bond_debt_nafmii - 银行间市场债务融资工具

### 第十批：上交所汇总 (bond_summary.py - 2个函数)
- [ ] bond_cash_summary_sse - 上交所债券现券汇总
- [ ] bond_deal_summary_sse - 上交所债券成交汇总

### 第十一批：沪深可转债 (bond_zh_cov.py - 8个函数)
- [ ] bond_zh_hs_cov_spot - 沪深可转债实时行情
- [ ] bond_zh_hs_cov_daily - 沪深可转债日线行情
- [ ] bond_zh_hs_cov_min - 沪深可转债分钟行情
- [ ] bond_zh_hs_cov_pre_min - 沪深可转债盘前分钟行情
- [ ] bond_zh_cov - 可转债行情
- [ ] bond_cov_comparison - 可转债比价
- [ ] bond_zh_cov_info - 可转债基本信息
- [ ] bond_zh_cov_value_analysis - 可转债价值分析

### 第十二批：新浪沪深债券 (bond_zh_sina.py - 2个函数)
- [ ] bond_zh_hs_spot - 沪深债券实时行情
- [ ] bond_zh_hs_daily - 沪深债券日线行情

## 实施策略

### 文件组织
根据 Python 源文件对应创建 Go 文件：
- bond_buy_back_em.go - 质押式回购
- bond_sina.go - 新浪财经债券
- bond_ths.go - 同花顺债券
- bond_cbond.go - 中国债券信息网
- bond_china.go - 外汇交易中心
- bond_china_money.go - 收盘收益率
- bond_info.go - 债券信息查询
- bond_issue.go - 债券发行
- bond_nafmii.go - 银行间市场
- bond_summary.go - 上交所汇总
- bond_zh_cov.go - 沪深可转债
- bond_zh_sina.go - 新浪沪深债券

### 测试要求
每个函数必须有对应的测试用例：
- 在 bond_test.go 或单独的测试文件中添加测试
- 至少覆盖正常情况和边界情况
- 验证返回的 DataFrame 结构和数据类型

### 代码规范
- 所有注释使用中文
- 函数命名遵循 Go 命名规范（PascalCase）
- 保持与 Python 版本的参数和返回值一致性
- 合理使用 utils 包中的工具函数

## 完成标准

- [x] Bank 模块所有函数实现并通过测试
- [ ] Bond 模块所有36个待实现函数完成
- [ ] 每个函数都有对应的单元测试
- [ ] 所有测试通过
- [ ] 代码格式化 (go fmt)
- [ ] 静态检查通过 (go vet)
- [ ] 更新 plans/STATUS.md

## 预估工作量

- Bank 模块: ✅ 已完成
- Bond 模块:
  - 每个函数平均 30-60 分钟（含测试）
  - 36个函数 × 45分钟 = 约 27 小时
  - 按批次实施，每批 2-4 小时

## 备注

- Bond 模块函数较多，需要分批实施
- 某些函数可能依赖特定的网站权限或 Cookie
- 部分数据接口可能需要处理 JavaScript 加密
- 优先实现常用和重要的函数
