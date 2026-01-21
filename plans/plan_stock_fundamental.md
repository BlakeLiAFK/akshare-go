# stock_fundamental/ 模块移植计划

> 创建日期: 2026-01-18
> 模块: stock_fundamental (股票基本面)
> 来源: Python akshare/stock_fundamental/
> 执行顺序: Z-A (第2个)

---

## 一、模块概述

### 1.1 模块信息

```
模块名: stock_fundamental
功能: 股票基本面数据
Python源码: _akshare_source/akshare/stock_fundamental/
Go目标: stock_fundamental/
```

### 1.2 接口清单

| 序号 | 文件名 | Python函数名 | 接口数 | 数据源 |
|------|--------|--------------|--------|--------|
| 1 | stock_basic_info_xq.py | stock_basic_info_xq | 1 | 雪球 |
| 2 | stock_finance_hk_em.py | stock_finance_hk_em | 1 | 东方财富 |
| 3 | stock_finance_sina.py | stock_finance_sina | 1 | 新浪财经 |
| 4 | stock_finance_ths.py | stock_finance_ths | 1 | 同花顺 |
| 5 | stock_finance_us_em.py | stock_finance_us_em | 1 | 东方财富 |
| 6 | stock_gbjg_em.py | stock_gbjg_em | 1 | 东方财富 |
| 7 | stock_hold.py | stock_hold*, stock_individual_hold_* | ~8 | 东方财富 |
| 8 | stock_ipo_declare.py | stock_ipo_declare | 1 | 东方财富 |
| 9 | stock_ipo_review.py | stock_ipo_review | 1 | 东方财富 |
| 10 | stock_ipo_tutor.py | stock_ipo_tutor | 1 | 东方财富 |
| 11 | stock_kcb_detail_sse.py | stock_kcb_detail_sse | 1 | 上交所 |
| 12 | stock_kcb_sse.py | stock_kcb_sse | 1 | 上交所 |
| 13 | stock_notice.py | stock_notice_* | ~5 | 东方财富 |
| 14 | stock_profit_forecast_em.py | stock_profit_forecast_em | 1 | 东方财富 |
| 15 | stock_profit_forecast_hk_etnet.py | stock_profit_forecast_hk_etnet | 1 | 经济通 |
| 16 | stock_profit_forecast_ths.py | stock_profit_forecast_ths | 1 | 同花顺 |
| 17 | stock_recommend.py | stock_recommend_* | ~3 | 东方财富 |
| 18 | stock_register_em.py | stock_register_em | 1 | 东方财富 |
| 19 | stock_restricted_em.py | stock_restricted_em | 1 | 东方财富 |
| 20 | stock_zygc.py | stock_zygc | 1 | 东方财富 |
| 21 | stock_zyjs_ths.py | stock_zyjs_ths | 1 | 同花顺 |

**总计**: 约 59 个接口

---

## 二、开发计划

### 2.1 分批实施

由于接口数量较多，将分批实施：

**第一批: 财务数据 (5个接口)**
- stock_finance_hk_em
- stock_finance_sina
- stock_finance_ths
- stock_finance_us_em
- stock_basic_info_xq

**第二批: IPO 数据 (4个接口)**
- stock_ipo_declare
- stock_ipo_review
- stock_ipo_tutor
- stock_kcb_sse

**第三批: 股东数据 (约10个接口)**
- stock_hold 系列
- stock_gbjg_em
- stock_zygc

**第四批: 其他数据 (约40个接口)**
- stock_notice 系列
- stock_profit_forecast 系列
- stock_recommend 系列
- stock_register_em
- stock_restricted_em
- stock_zyjs_ths

---

## 三、技术要点

### 3.1 文件命名规范

```
stock_fundamental_{业务}_{数据源}.go
```

示例:
- `stock_fundamental_finance_hk_em.go`
- `stock_fundamental_ipo_declare.go`
- `stock_fundamental_hold_em.go`

### 3.2 函数命名规范

```python
# Python 原版
def stock_finance_hk_em() -> pd.DataFrame
```

```go
// Go 版本
func StockFinanceHkEm() ([]StockFinanceHK, error)
```

---

## 四、开发步骤

### Step 1: 创建基础结构
- [ ] 更新 doc.go
- [ ] 创建 types.go

### Step 2: 第一批实现
- [ ] 财务数据接口 (5个)

### Step 3: 第二批实现
- [ ] IPO 数据接口 (4个)

### Step 4: 第三批实现
- [ ] 股东数据接口 (10个)

### Step 5: 第四批实现
- [ ] 其他数据接口 (40个)

### Step 6: 测试
- [ ] 编写测试用例
- [ ] 确保测试覆盖率 > 80%

---

## 五、验收标准

- [ ] 编译通过: `go build ./stock_fundamental/...`
- [ ] 测试通过: `go test ./stock_fundamental/...`
- [ ] 测试覆盖率 > 80%
- [ ] 代码格式规范: `go fmt ./stock_fundamental/...`
- [ ] 静态检查通过: `go vet ./stock_fundamental/...`

---

## 六、进度跟踪

| 批次 | 状态 | 完成时间 |
|------|------|----------|
| 创建计划文档 | ✅ 完成 | 2026-01-18 |
| 第一批: 财务数据 (8个) | ✅ 完成 | 2026-01-18 |
| 第二批: IPO 数据 (4个) | ✅ 完成 | 2026-01-18 |
| 第三批: 股东数据 (4个) | ✅ 完成 | 2026-01-18 |
| 第四批: 其他数据 (7个) | ✅ 完成 | 2026-01-18 |
| 测试验收 | 🔜 待开始 | - |

### 已实现接口清单 (23/59)

| 序号 | 函数名 | 功能 | 状态 |
|------|--------|------|------|
| 1 | StockFinancialReportSina | 新浪财经-财务报表 | ✅ |
| 2 | StockFinancialAbstract | 新浪财经-财务摘要 | ✅ |
| 3 | StockFinancialAnalysisIndicatorEm | 东方财富-财务分析指标 | ✅ |
| 4 | StockHistoryDividend | 新浪财经-历史分红 | ✅ |
| 5 | StockHistoryDividendDetail | 新浪财经-分红详情 | ✅ |
| 6 | StockIpoInfo | 新浪财经-IPO信息 | ✅ |
| 7 | StockMainStockHolder | 新浪财经-主要股东 | ✅ |
| 8 | StockFundStockHolder | 新浪财经-基金持股 | ✅ |
| 9 | StockIpoDeclareEm | 东方财富-IPO申报企业信息 | ✅ |
| 10 | StockIpoReviewEm | 东方财富-IPO过会企业信息 | ✅ |
| 11 | StockIpoTutorEm | 东方财富-IPO辅导备案信息 | ✅ |
| 12 | StockKcbSse | 上交所-科创板 | ✅ |
| 13 | StockInstituteHold | 新浪财经-机构持股一览表 | ✅ |
| 14 | StockInstituteHoldDetail | 新浪财经-机构持股详情 | ✅ |
| 15 | StockZhAGbjgEm | 东方财富-股本结构 | ✅ |
| 16 | StockZygcEm | 东方财富-主营构成 | ✅ |
| 17 | StockProfitForecastEm | 东方财富-盈利预测 | ✅ |
| 18 | StockRegisterAllEm | 东方财富-IPO审核(全部) | ✅ |
| 19 | StockRegisterKcb | 东方财富-IPO审核(科创板) | ✅ |
| 20 | StockRegisterCyb | 东方财富-IPO审核(创业板) | ✅ |
| 21 | StockRestrictedReleaseSummaryEm | 东方财富-限售股解禁汇总 | ✅ |
| 22 | StockRestrictedReleaseDetailEm | 东方财富-限售股解禁详情 | ✅ |
| 23 | StockRestrictedReleaseQueueEm | 东方财富-个股限售解禁批次 | ✅ |

**进度**: 23/59 (39.0%)

---

*计划版本: v1.1*
*创建时间: 2026-01-18*
*更新时间: 2026-01-18*
