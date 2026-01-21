# Bond 模块实现计划

> 创建: 2026-01-18
> 状态: 进行中

## 目标

实现 akshare bond 模块的所有 42 个债券数据接口，为 akshare-go 提供完整的债券数据支持。

## 模块概况

- **总接口数**: 42个
- **已实现**: 5个 (2个历史 + 3个新增)
- **待实现**: 37个
- **数据源**: 集思录、新浪财经、巨潮资讯、中国货币网、东方财富
- **当前进度**: 11.9% (5/42)

## 已实现接口 (5个)

| 接口名 | 文件位置 | 数据源 | 实现日期 |
|-------|---------|--------|----------|
| BondCBIndexJSL | bond/convertible.go | 集思录 | 历史 |
| BondZHUSRate | bond/other.go | 东方财富 | 历史 |
| BondCBJSL | bond/bond_jsl.go | 集思录 | 2026-01-18 |
| BondCBRedeemJSL | bond/bond_jsl.go | 集思录 | 2026-01-18 |
| BondCBAdjLogsJSL | bond/bond_jsl.go | 集思录 | 2026-01-18 |

## 实施批次

### 第一批: 集思录可转债数据 (3个) - P0 ✅ 已完成

**目标**: 集思录(Jisilu)可转债核心数据

| 函数 | 描述 | 优先级 | 状态 |
|------|------|-------|------|
| BondCBJSL | 集思录可转债列表 | P0 | ✅ |
| BondCBRedeemJSL | 集思录可转债强赎数据 | P0 | ✅ |
| BondCBAdjLogsJSL | 集思录转股价调整记录 | P0 | ✅ |

**数据源**: https://www.jisilu.cn/data/cbnew/
**实现文件**: bond/bond_jsl.go
**测试文件**: bond/bond_jsl_test.go

**步骤**:
- [x] 创建 bond/bond_jsl.go
- [x] 实现 BondCBJSL (集思录可转债列表)
- [x] 实现 BondCBRedeemJSL (强赎数据)
- [x] 实现 BondCBAdjLogsJSL (调整记录)
- [x] 编写测试用例
- [x] 运行测试验证

**完成时间**: 2026-01-18
**测试结果**: 所有测试通过
- BondCBJSL: 获取 30 行数据, 23 列
- BondCBRedeemJSL: 获取 385 行数据, 18 列
- BondCBAdjLogsJSL: 正常处理有/无调整记录的情况

### 第二批: 东方财富可转债数据 (8个) - P0

**目标**: 东方财富和新浪财经可转债行情与分析

| 函数 | 描述 | 优先级 |
|------|------|-------|
| BondZHHSCovSpot | 新浪可转债实时行情 | P0 |
| BondZHHSCovDaily | 新浪可转债日K线 | P0 |
| BondZHHSCovMin | 东财可转债分时行情 | P0 |
| BondZHHSCovPreMin | 东财可转债盘前分时 | P0 |
| BondZHCov | 东财可转债数据 | P0 |
| BondCovComparison | 东财可转债比价表 | P0 |
| BondZHCovInfo | 东财可转债详情 | P0 |
| BondZHCovValueAnalysis | 东财可转债价值分析 | P0 |

**数据源**:
- https://vip.stock.finance.sina.com.cn/mkt/#hskzz_z
- https://data.eastmoney.com/kzz/default.html

**实现文件**: bond/bond_zh_cov.go

**步骤**:
- [ ] 创建 bond/bond_zh_cov.go
- [ ] 实现新浪财经接口（2个）
- [ ] 实现东方财富接口（6个）
- [ ] 编写测试用例
- [ ] 运行测试验证

### 第三批: 巨潮资讯债券发行 (5个) - P1

**目标**: 巨潮资讯(CNINFO)债券发行数据

| 函数 | 描述 | 优先级 |
|------|------|-------|
| BondTreasureIssueCNINFO | 国债发行 | P1 |
| BondLocalGovernmentIssueCNINFO | 地方债发行 | P1 |
| BondCorporateIssueCNINFO | 企业债发行 | P1 |
| BondCovIssueCNINFO | 可转债发行 | P1 |
| BondCovStockIssueCNINFO | 可转债转股 | P1 |

**数据源**: http://webapi.cninfo.com.cn/#/thematicStatistics
**实现文件**: bond/bond_issue_cninfo.go
**特殊要求**: 需要 JS 解密（py_mini_racer）

**步骤**:
- [ ] 创建 bond/bond_issue_cninfo.go
- [ ] 实现 JS 解密逻辑
- [ ] 实现所有 5 个接口
- [ ] 编写测试用例
- [ ] 运行测试验证

### 第四批: 中国货币网债券数据 (7个) - P1

**目标**: 中国货币网(chinamoney.com.cn)债券市场数据

| 函数 | 描述 | 优先级 |
|------|------|-------|
| BondChinaCloseReturnMap | 收益率曲线映射表 | P1 |
| BondChinaCloseReturn | 收盘收益率曲线历史 | P1 |
| MacroChinaSwapRate | FR007利率互换曲线 | P1 |
| MacroChinaBondPublic | 债券信息披露-发行 | P1 |
| BondInfoCMQuery | 中国货币网查询参数 | P1 |
| BondInfoCM | 中国货币网债券信息查询 | P1 |
| BondInfoDetailCM | 中国货币网债券详情 | P1 |

**数据源**:
- https://www.chinamoney.com.cn/chinese/bkcurvclosedyhis/
- https://www.chinamoney.com.cn/chinese/scsjzqxx/

**实现文件**: bond/bond_china_money.go, bond/bond_info_cm.go
**特殊要求**: 需要处理 session 注册和 cookie

**步骤**:
- [ ] 创建 bond/bond_china_money.go
- [ ] 创建 bond/bond_info_cm.go
- [ ] 实现 session 注册逻辑
- [ ] 实现所有 7 个接口
- [ ] 编写测试用例
- [ ] 运行测试验证

### 第五批: 新浪财经沪深债券 (3个) - P1

**目标**: 新浪财经沪深债券行情

| 函数 | 描述 | 优先级 |
|------|------|-------|
| BondZHHSSpot | 新浪沪深债券实时行情 | P1 |
| BondZHHSDaily | 新浪沪深债券日K线 | P1 |
| GetZHBondHSPageCount | 辅助-获取总页数 | P1 |

**数据源**: https://vip.stock.finance.sina.com.cn/mkt/#hs_z
**实现文件**: bond/bond_zh_sina.go
**特殊要求**: 需要 demjson 解码和 JS 解密

**步骤**:
- [ ] 创建 bond/bond_zh_sina.go
- [ ] 实现 demjson 解码
- [ ] 实现 JS 解密
- [ ] 实现所有 3 个接口
- [ ] 编写测试用例
- [ ] 运行测试验证

### 第六批: 中国债券市场行情 (3个) - P2

**目标**: 中国货币网现券市场行情

| 函数 | 描述 | 优先级 |
|------|------|-------|
| BondSpotQuote | 现券市场做市报价 | P2 |
| BondSpotDeal | 现券市场成交行情 | P2 |
| BondChinaYield | 国债收益率曲线 | P2 |

**数据源**:
- https://www.chinamoney.com.cn/chinese/mkdatabond/
- https://yield.chinabond.com.cn/

**实现文件**: bond/bond_china.go

**步骤**:
- [ ] 创建 bond/bond_china.go
- [ ] 实现所有 3 个接口
- [ ] 编写测试用例
- [ ] 运行测试验证

### 第七批: 东方财富质押式回购 (3个) - P2

**目标**: 东方财富质押式回购数据

| 函数 | 描述 | 优先级 |
|------|------|-------|
| BondSHBuyBackEM | 上证质押式回购 | P2 |
| BondSZBuyBackEM | 深证质押式回购 | P2 |
| BondBuyBackHistEM | 质押式回购历史数据 | P2 |

**数据源**: https://quote.eastmoney.com/center/gridlist.html#bond_sz_buyback
**实现文件**: bond/bond_buy_back_em.go

**步骤**:
- [ ] 创建 bond/bond_buy_back_em.go
- [ ] 实现所有 3 个接口
- [ ] 编写测试用例
- [ ] 运行测试验证

## 完成标准

- [x] 分析所有 Python 源码，理解接口逻辑
- [ ] 实现所有 40 个待开发接口
- [ ] 每个接口都有对应的测试用例
- [ ] 所有测试通过
- [ ] 代码通过 go vet 和 go fmt 检查
- [ ] 更新模块文档

## 技术挑战

1. **JS 解密**: 新浪财经使用 JS 加密，需要用 goja 或其他方式处理
2. **Session 管理**: 中国货币网需要注册 session
3. **demjson 解码**: 新浪财经返回特殊格式 JSON
4. **分页处理**: 多个接口需要处理分页逻辑
5. **日期格式**: 不同数据源使用不同的日期格式

## 参考资源

- Python 源码: `_akshare_source/akshare/bond/`
- Go 实现示例: `bond/convertible.go`, `bond/other.go`
- 工具函数: `utils/` 目录

## 数据源总结

| 数据源 | 接口数 | 说明 |
|-------|-------|------|
| 集思录 (jisilu.cn) | 3 | 可转债核心数据 |
| 新浪财经 (sina.com.cn) | 5 | 可转债和债券行情 |
| 东方财富 (eastmoney.com) | 11 | 可转债和回购数据 |
| 巨潮资讯 (cninfo.com.cn) | 5 | 债券发行数据 |
| 中国货币网 (chinamoney.com.cn) | 10 | 债券信息和收益率 |
| 中债网 (chinabond.com.cn) | 1 | 国债收益率 |

## 备注

- 优先实现 P0 接口（可转债核心数据）
- P1 接口为重要功能，第二批实现
- P2 接口为补充功能，最后实现
- 每批实现后立即编写测试并验证
