# stock_fundamental 模块第四批核心接口实现计划

## 任务概述

实现 stock_fundamental 模块的第四批核心接口，共7个函数。

## 实现状态

已完成所有7个核心接口的实现，代码已编译通过。

## 实现列表

### 1. StockProfitForecastEm - 东方财富盈利预测

**功能**: 东方财富网-数据中心-研究报告-盈利预测
**Python源码**: `_akshare_source/akshare/stock_fundamental/stock_profit_forecast_em.py`
**文件**: `stock_fundamental_profit_forecast_em.go`

**参数**:
- `symbol string`: 行业板块，默认为空获取全部

**返回类型**: `[]StockProfitForecastEmItem`

**技术要点**:
- URL: `https://datacenter-web.eastmoney.com/api/data/v1/get`
- 动态获取年份列 (YEAR1-YEAR4)
- 分页获取所有数据
- 按研报数排序

---

### 2. StockRegisterAllEm - IPO审核信息-全部

**功能**: 东方财富网-IPO审核信息-全部
**Python源码**: `_akshare_source/akshare/stock_fundamental/stock_register_em.py`
**文件**: `stock_fundamental_register_em.go`

**返回类型**: `[]StockRegisterEmItem`

**技术要点**:
- URL: `https://datacenter-web.eastmoney.com/api/data/v1/get`
- reportName: `RPT_IPO_INFOALLNEW`
- 分页获取
- 生成招股说明书链接

---

### 3. StockRegisterKcb - IPO审核信息-科创板

**功能**: 东方财富网-IPO审核信息-科创板
**Python源码**: `_akshare_source/akshare/stock_fundamental/stock_register_em.py`
**文件**: `stock_fundamental_register_em.go`

**返回类型**: `[]StockRegisterEmItem`

**技术要点**:
- filter: `(PREDICT_LISTING_MARKET="科创板")`
- 其他同 StockRegisterAllEm

---

### 4. StockRegisterCyb - IPO审核信息-创业板

**功能**: 东方财富网-IPO审核信息-创业板
**Python源码**: `_akshare_source/akshare/stock_fundamental/stock_register_em.py`
**文件**: `stock_fundamental_register_em.go`

**返回类型**: `[]StockRegisterEmItem`

**技术要点**:
- filter: `(PREDICT_LISTING_MARKET="创业板")`

---

### 5. StockRestrictedReleaseSummaryEm - 限售股解禁汇总

**功能**: 东方财富网-限售股解禁汇总
**Python源码**: `_akshare_source/akshare/stock_fundamental/stock_restricted_em.py`
**文件**: `stock_fundamental_restricted_em.go`

**参数**:
- `symbol string`: 市场代码
- `startDate string`: 开始日期 (格式: 20221101)
- `endDate string`: 结束日期 (格式: 20221209)

**返回类型**: `[]StockRestrictedReleaseSummaryEmItem`

**技术要点**:
- symbol_map 映射
- 日期格式转换: 20221101 -> 2022-11-01
- reportName: `RPT_LIFTDAY_STA`

---

### 6. StockRestrictedReleaseDetailEm - 限售股解禁详情

**功能**: 东方财富网-限售股解禁详情
**Python源码**: `_akshare_source/akshare/stock_fundamental/stock_restricted_em.py`
**文件**: `stock_fundamental_restricted_em.go`

**参数**:
- `startDate string`: 开始日期
- `endDate string`: 结束日期

**返回类型**: `[]StockRestrictedReleaseDetailEmItem`

**技术要点**:
- 日期格式转换
- 分页获取
- reportName: `RPT_LIFT_STAGE`

---

### 7. StockRestrictedReleaseQueueEm - 个股限售解禁批次

**功能**: 东方财富网-个股限售解禁批次
**Python源码**: `_akshare_source/akshare/stock_fundamental/stock_restricted_em.py`
**文件**: `stock_fundamental_restricted_em.go`

**参数**:
- `symbol string`: 股票代码

**返回类型**: `[]StockRestrictedReleaseQueueEmItem`

**技术要点**:
- filter: `(SECURITY_CODE="{symbol}")`
- 检查 result 是否存在

---

## 已实现的文件

1. **types.go** - 更新类型定义
   - StockProfitForecastEmItem
   - StockRegisterEmItem
   - StockRestrictedReleaseSummaryEmItem
   - StockRestrictedReleaseDetailEmItem
   - StockRestrictedReleaseQueueEmItem

2. **stock_fundamental_profit_forecast_em.go** - 盈利预测接口
   - StockProfitForecastEm()
   - getMostCommonYear()

3. **stock_fundamental_register_em.go** - IPO审核信息接口
   - StockRegisterAllEm()
   - StockRegisterKcb()
   - StockRegisterCyb()
   - stockRegisterEm() (内部实现)

4. **stock_fundamental_restricted_em.go** - 限售股解禁接口
   - StockRestrictedReleaseSummaryEm()
   - StockRestrictedReleaseDetailEm()
   - StockRestrictedReleaseQueueEm()
   - formatDate() (内部辅助函数)
   - GetSupportedSymbolList() (辅助函数)

5. **stock_fundamental_other_test.go** - 测试文件

## 编译和测试

```bash
# 编译
go build ./stock_fundamental/...
# ✅ 编译成功

# 测试
go test -v ./stock_fundamental/ -run "TestStockRestrictedReleaseQueueEm"
go test -v ./stock_fundamental/ -run "TestGetSupportedSymbolList"
go test -v ./stock_fundamental/ -run "TestStockRestrictedReleaseSummaryEmInvalidSymbol"
# ✅ 测试通过
```

## 类型定义 (types.go)

已添加以下结构体类型:

```go
// StockProfitForecastEmItem 东方财富盈利预测项
type StockProfitForecastEmItem struct {
    Index           int     `json:"index"`
    Code            string  `json:"code"`
    Name            string  `json:"name"`
    ReportCount     int     `json:"report_count"`
    RatingBuy       int     `json:"rating_buy"`
    RatingAdd       int     `json:"rating_add"`
    RatingNeutral   int     `json:"rating_neutral"`
    RatingReduce    int     `json:"rating_reduce"`
    RatingSell      int     `json:"rating_sell"`
    Year1Forecast   float64 `json:"year1_forecast"`
    Year2Forecast   float64 `json:"year2_forecast"`
    Year3Forecast   float64 `json:"year3_forecast"`
    Year4Forecast   float64 `json:"year4_forecast"`
}

// StockRegisterEmItem IPO审核信息项
type StockRegisterEmItem struct {
    Index            int       `json:"index"`
    CompanyName      string    `json:"company_name"`
    State            string    `json:"state"`
    RegAddress       string    `json:"reg_address"`
    Industry         string    `json:"industry"`
    RecommendOrg     string    `json:"recommend_org"`
    LawFirm          string    `json:"law_firm"`
    AccountFirm      string    `json:"account_firm"`
    UpdateDate       time.Time `json:"update_date"`
    AcceptDate       time.Time `json:"accept_date"`
    PredictMarket    string    `json:"predict_market"`
    ProspectusURL    string    `json:"prospectus_url"`
}

// StockRestrictedReleaseSummaryEmItem 限售股解禁汇总项
type StockRestrictedReleaseSummaryEmItem struct {
    Index                    int       `json:"index"`
    LiftDate                 time.Time `json:"lift_date"`
    StockCount               int       `json:"stock_count"`
    LiftShares               float64   `json:"lift_shares"`
    ActualLiftShares         float64   `json:"actual_lift_shares"`
    ActualLiftMarketCap      float64   `json:"actual_lift_market_cap"`
    HS300Index               float64   `json:"hs300_index"`
    HS300ChangeRatio         float64   `json:"hs300_change_ratio"`
}

// StockRestrictedReleaseDetailEmItem 限售股解禁详情项
type StockRestrictedReleaseDetailEmItem struct {
    Index                  int       `json:"index"`
    Code                   string    `json:"code"`
    Name                   string    `json:"name"`
    LiftDate               time.Time `json:"lift_date"`
    SharesType             string    `json:"shares_type"`
    LiftShares             float64   `json:"lift_shares"`
    ActualLiftShares       float64   `json:"actual_lift_shares"`
    ActualLiftMarketCap    float64   `json:"actual_lift_market_cap"`
    CirculationRatio       float64   `json:"circulation_ratio"`
    PrevClosePrice         float64   `json:"prev_close_price"`
    Before20ChangeRatio    float64   `json:"before_20_change_ratio"`
    After20ChangeRatio     float64   `json:"after_20_change_ratio"`
}

// StockRestrictedReleaseQueueEmItem 个股限售解禁批次项
type StockRestrictedReleaseQueueEmItem struct {
    Index                int       `json:"index"`
    LiftDate             time.Time `json:"lift_date"`
    HolderNum            int       `json:"holder_num"`
    LiftShares           float64   `json:"lift_shares"`
    ActualLiftShares     float64   `json:"actual_lift_shares"`
    NonLiftShares        float64   `json:"non_lift_shares"`
    ActualLiftMarketCap  float64   `json:"actual_lift_market_cap"`
    TotalRatio           float64   `json:"total_ratio"`
    CirculationRatio     float64   `json:"circulation_ratio"`
    PrevClosePrice       float64   `json:"prev_close_price"`
    SharesType           string    `json:"shares_type"`
    Before20ChangeRatio  float64   `json:"before_20_change_ratio"`
    After20ChangeRatio   float64   `json:"after_20_change_ratio"`
}
```

---

## 文件结构

```
stock_fundamental/
├── types.go                                    # 更新类型定义
├── stock_fundamental_profit_forecast_em.go     # 新增
├── stock_fundamental_register_em.go            # 新增
├── stock_fundamental_restricted_em.go          # 新增
└── stock_fundamental_other_test.go             # 新增测试
```

---

## 编码规范

1. 使用中文注释
2. 导出函数使用 PascalCase
3. 错误使用 `fmt.Errorf` 包装
4. 日期使用 `utils.ParseDate`
5. 数值转换使用 `utils.MustFloat64`
6. 复用 `parseJSONP` 函数

---

## 测试计划

每个接口编写基础测试用例，验证:
- 函数能正常调用
- 返回数据结构正确
- 日期解析正确
