# 测试失败分析报告

生成时间: 2026-01-21

## 测试结果概览

### 通过的模块 ✅
- air (cached)
- article (cached)
- bank (cached)
- cal (cached)
- crypto (cached)
- currency (cached)
- movie (2.392s)
- news (cached)
- nlp (cached)
- other (cached)
- rate (9.940s)
- reits (32.941s)
- stock (cached)
- stock_a (114.231s)
- tool (cached)
- utils (cached)
- futures_derivative (155.558s)

### 失败的模块 ❌

#### 1. stock_feature [编译失败]
**严重程度**: 🔴 高 - 编译错误导致无法运行测试

**错误类型**: 编译错误

**具体错误**:
```
stock_feature/stock_feature_test.go:11:17: undefined: StockAHighLowStatistics
stock_feature/stock_feature_test.go:48:43: invalid argument: result (variable of struct type dataframe.DataFrame) for built-in len
stock_feature/stock_feature_test.go:53:17: undefined: StockFhpsDetailEm
stock_feature/stock_feature_test.go:79:43: invalid argument: result (variable of struct type dataframe.DataFrame) for built-in len
stock_feature/stock_feature_test.go:93:49: invalid argument: result (variable of struct type dataframe.DataFrame) for built-in len
stock_feature/stock_feature_test.go:119:17: not enough arguments in call to StockDxsylEm
	have ()
	want (string)
stock_feature/stock_feature_test.go:121:46: invalid argument: result (variable of struct type dataframe.DataFrame) for built-in len
stock_feature/stock_feature_test.go:128:43: invalid argument: result (variable of struct type dataframe.DataFrame) for built-in len
```

**问题分析**:
1. 缺少函数定义: `StockAHighLowStatistics`, `StockFhpsDetailEm`
2. DataFrame API 使用错误: 应使用 `.Nrow()` 而非 `len()`
3. 函数参数不匹配: `StockDxsylEm` 需要一个 string 参数

**修复优先级**: P0 - 必须立即修复

---

#### 2. bond (51.994s)
**严重程度**: 🟡 中 - 网络相关失败

**失败测试**:
1. `TestBondSHBuyBackEM` - 上证质押式回购
2. `TestBondSZBuyBackEM` - 深证质押式回购
3. `TestBondSpotQuote` - 现券市场做市报价
4. `TestBondSpotDeal` - 现券市场成交行情
5. `TestBondChinaCloseReturn` - 收盘收益率曲线

**错误原因**:
- 前两个: 网络 EOF 错误（东方财富API连接失败）
- 后三个: 未获取到有效数据

**修复优先级**: P2 - 网络问题，可能需要调试 API 或标记为 Skip

---

#### 3. stock_fundamental (36.064s)
**严重程度**: 🔴 高 - 包含 panic

**失败测试**:
1. `TestStockIndividualBasicInfoXq` - 雪球A股公司简介
2. `TestStockIndividualBasicInfoUsXq` - 雪球美股公司简介
3. `TestStockIndividualBasicInfoHkXq` - 雪球港股公司简介
4. `TestStockFinancialReportSina` - 新浪财务报表（**PANIC**）

**错误原因**:
- 前三个: 雪球API返回 `code=400016: 遇到错误，请刷新页面或者重新登录帐号后再试`（需要登录）
- 第四个: `panic: runtime error: index out of range [2] with length 2`
  - 位置: `stock_fundamental/stock_fundamental_finance_sina.go:133`
  - 原因: DataFrame 列索引越界

**修复优先级**:
- 雪球API: P2 - 标记为 Skip（需要登录）
- Panic: P0 - 必须修复，影响程序稳定性

---

#### 4. economic (600.898s)
**严重程度**: 🟡 中 - 测试超时

**问题**: 测试运行时间过长（10分钟），可能某些测试挂起或网络超时

**修复优先级**: P1 - 需要定位超时原因

---

#### 5. fund (601.753s)
**严重程度**: 🟡 中 - 测试超时

**问题**: 与 economic 类似，测试运行超时

**修复优先级**: P1 - 需要定位超时原因

---

#### 6. futures (357.379s)
**严重程度**: 🟡 中 - 测试超时或部分失败

**问题**: 测试运行时间较长

**修复优先级**: P1 - 需要查看具体失败原因

---

#### 7. index (601.506s)
**严重程度**: 🟡 中 - 测试超时

**问题**: 测试运行超时

**修复优先级**: P1 - 需要定位超时原因

---

#### 8. option (73.627s)
**严重程度**: 🟡 中 - 测试失败

**问题**: 需要查看具体失败原因

**修复优先级**: P1 - 需要定位失败测试

---

## 修复计划

### 立即修复 (P0)
1. **stock_feature 编译错误**
   - 实现缺失的函数
   - 修复 DataFrame API 使用错误
   - 修正函数参数

2. **stock_fundamental panic**
   - 修复 `stock_fundamental_finance_sina.go:133` 的索引越界问题
   - 添加边界检查

### 高优先级 (P1)
1. 定位 economic, fund, index 超时原因
2. 查看 futures, option 具体失败测试
3. 对超时测试添加合理的 timeout 或标记为 Skip

### 中优先级 (P2)
1. 修复 bond 模块的网络相关测试
2. 将雪球 API 相关测试标记为 Skip（需要登录）

## 后续行动

1. 先修复 P0 级别的编译错误和 panic
2. 运行单个模块测试以定位 P1 问题
3. 根据实际情况调整测试策略（Skip vs 修复）
