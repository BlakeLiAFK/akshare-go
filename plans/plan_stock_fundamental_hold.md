# stock_fundamental 第三批股东数据接口实现计划

> 创建日期: 2026-01-18
> 模块: stock_fundamental (股票基本面)
> 批次: 第三批 - 股东数据
> 接口数量: 4个

---

## 一、接口清单

### 1.1 待实现接口

| 序号 | 函数名 | 数据源 | 功能描述 |
|------|--------|--------|----------|
| 1 | StockInstituteHold | 新浪财经 | 机构持股一览表 |
| 2 | StockInstituteHoldDetail | 新浪财经 | 机构持股详情 |
| 3 | StockZhAGbjgEm | 东方财富 | 股本结构 |
| 4 | StockZygcEm | 东方财富 | 主营构成 |

---

## 二、实现细节

### 2.1 StockInstituteHold - 机构持股一览表

**Python源码**: `_akshare_source/akshare/stock_fundamental/stock_hold.py`

**功能**: 新浪财经-股票-机构持股一览表

**URL**: `https://vip.stock.finance.sina.com.cn/q/go.php/vComStockHold/kind/jgcg/index.phtml`

**参数**:
- `symbol`: 报告期代码，格式如 "20201"（2020年一季报）
  - 最后一位: {"一季报":1, "中报":2, "三季报":3, "年报":4}
  - 前四位: 年份

**返回字段**:
- 证券代码: string
- 证券简称: string
- 机构数: float64
- 机构数变化: float64
- 持股比例: float64
- 持股比例增幅: float64
- 占流通股比例: float64
- 占流通股比例增幅: float64

**技术要点**:
- HTML表格解析
- 证券代码需要填充为6位
- 删除"明细"列

---

### 2.2 StockInstituteHoldDetail - 机构持股详情

**Python源码**: `_akshare_source/akshare/stock_fundamental/stock_hold.py`

**功能**: 新浪财经-股票-机构持股详情

**URL**: `https://vip.stock.finance.sina.com.cn/q/api/jsonp.php/var%20details=/ComStockHoldService.getJGCGDetail`

**参数**:
- `stock`: 股票代码，如 "600433"
- `quarter`: 报告期代码，格式同上

**返回字段**:
- 持股机构类型: string (需要映射: fund->基金, socialSecurity->全国社保, qfii->QFII, insurance->保险)
- 持股机构代码: string
- 持股机构简称: string
- 持股机构全称: string
- 持股数: float64
- 最新持股数: float64
- 持股比例: float64
- 最新持股比例: float64
- 占流通股比例: float64
- 最新占流通股比例: float64
- 持股比例增幅: float64
- 占流通股比例增幅: float64

**技术要点**:
- JSONP响应解析（需要提取纯JSON）
- 数据嵌套在 data 字段中
- 持股机构类型需要翻译

---

### 2.3 StockZhAGbjgEm - 股本结构

**Python源码**: `_akshare_source/akshare/stock_fundamental/stock_gbjg_em.py`

**功能**: 东方财富-A股数据-股本结构

**URL**: `https://datacenter.eastmoney.com/securities/api/data/v1/get`

**参数**:
- `symbol`: 带市场标识的股票代码，如 "603392.SH"

**返回字段**:
- 变更日期: time.Time
- 总股本: float64
- 流通受限股份: float64
- 其他内资持股(受限): float64
- 境内法人持股(受限): float64
- 境内自然人持股(受限): float64
- 已流通股份: float64
- 已上市流通A股: float64
- 变动原因: string

**技术要点**:
- JSON API调用
- 字段映射重命名
- 日期解析

---

### 2.4 StockZygcEm - 主营构成

**Python源码**: `_akshare_source/akshare/stock_fundamental/stock_zygc.py`

**功能**: 东方财富网-个股-主营构成

**URL**: `https://emweb.securities.eastmoney.com/PC_HSF10/BusinessAnalysis/PageAjax`

**参数**:
- `symbol`: 带市场标识的股票代码，如 "SH688041"

**返回字段**:
- 股票代码: string
- 报告日期: time.Time
- 分类类型: string (映射: "2"->"按产品分类", "3"->"按地区分类")
- 主营构成: string
- 主营收入: float64
- 收入比例: float64
- 主营成本: float64
- 成本比例: float64
- 主营利润: float64
- 利润比例: float64
- 毛利率: float64

**技术要点**:
- JSON API调用
- 分类类型需要映射
- 字段过滤和重命名

---

## 三、文件结构

```
stock_fundamental/
├── types.go                           # 添加类型定义
├── stock_fundamental_hold_sina.go     # 新浪财经机构持股
├── stock_fundamental_gbjg_zygc_em.go  # 东方财富股本结构和主营构成
├── stock_fundamental_hold_test.go     # 测试用例
```

---

## 四、类型定义

### 4.1 StockInstituteHoldItem

```go
type StockInstituteHoldItem struct {
    Code                string  // 证券代码
    Name                string  // 证券简称
    InstituteCount      float64 // 机构数
    InstituteCountChange float64 // 机构数变化
    HoldingRatio        float64 // 持股比例
    HoldingRatioChange  float64 // 持股比例增幅
    CirculationRatio    float64 // 占流通股比例
    CirculationRatioChange float64 // 占流通股比例增幅
}
```

### 4.2 StockInstituteHoldDetailItem

```go
type StockInstituteHoldDetailItem struct {
    InstituteType      string  // 持股机构类型
    InstituteCode      string  // 持股机构代码
    InstituteName      string  // 持股机构简称
    InstituteFullName  string  // 持股机构全称
    Holdings           float64 // 持股数
    LatestHoldings     float64 // 最新持股数
    HoldingRatio       float64 // 持股比例
    LatestHoldingRatio float64 // 最新持股比例
    CirculationRatio   float64 // 占流通股比例
    LatestCirculationRatio float64 // 最新占流通股比例
    HoldingRatioChange float64 // 持股比例增幅
    CirculationRatioChange float64 // 占流通股比例增幅
}
```

### 4.3 StockZhAGbjgEmItem

```go
type StockZhAGbjgEmItem struct {
    ChangeDate         time.Time // 变更日期
    TotalShares        float64   // 总股本
    LimitedShares      float64   // 流通受限股份
    LimitedOthARS      float64   // 其他内资持股(受限)
    LimitedDomestic    float64   // 境内法人持股(受限)
    LimitedNatural     float64   // 境内自然人持股(受限)
    UnlimitedShares    float64   // 已流通股份
    ListedAShares      float64   // 已上市流通A股
    ChangeReason       string    // 变动原因
}
```

### 4.4 StockZygcEmItem

```go
type StockZygcEmItem struct {
    Code              string    // 股票代码
    ReportDate        time.Time // 报告日期
    CategoryType      string    // 分类类型
    MainComposition   string    // 主营构成
    MainIncome        float64   // 主营收入
    IncomeRatio       float64   // 收入比例
    MainCost          float64   // 主营成本
    CostRatio         float64   // 成本比例
    MainProfit        float64   // 主营利润
    ProfitRatio       float64   // 利润比例
    GrossProfitRatio  float64   // 毛利率
}
```

---

## 五、开发步骤

### Step 1: 更新 types.go
- [ ] 添加 StockInstituteHoldItem
- [ ] 添加 StockInstituteHoldDetailItem
- [ ] 添加 StockZhAGbjgEmItem
- [ ] 添加 StockZygcEmItem

### Step 2: 实现 stock_fundamental_hold_sina.go
- [ ] 实现 StockInstituteHold
- [ ] 实现 StockInstituteHoldDetail
- [ ] 添加 parseJSONP 辅助函数

### Step 3: 实现 stock_fundamental_gbjg_zygc_em.go
- [ ] 实现 StockZhAGbjgEm
- [ ] 实现 StockZygcEm

### Step 4: 编写测试用例
- [ ] 测试 StockInstituteHold
- [ ] 测试 StockInstituteHoldDetail
- [ ] 测试 StockZhAGbjgEm
- [ ] 测试 StockZygcEm
- [ ] 测试无效参数

### Step 5: 验收
- [ ] 编译通过: `go build ./stock_fundamental/...`
- [ ] 测试通过: `go test ./stock_fundamental/...`
- [ ] 代码格式规范: `go fmt ./stock_fundamental/...`

---

## 六、进度跟踪

| 任务 | 状态 | 完成时间 |
|------|------|----------|
| 创建计划文档 | ✅ 完成 | 2026-01-18 |
| 更新 types.go | ✅ 完成 | 2026-01-18 |
| 实现 hold_sina.go | ✅ 完成 | 2026-01-18 |
| 实现 gbjg_zygc_em.go | ✅ 完成 | 2026-01-18 |
| 编写测试用例 | ✅ 完成 | 2026-01-18 |
| 验收测试 | ✅ 完成 | 2026-01-18 |

---

## 七、实现总结

### 7.1 已完成文件

1. **types.go** - 添加了4个新类型：
   - `StockInstituteHoldItem` - 机构持股一览表项
   - `StockInstituteHoldDetailItem` - 机构持股详情项
   - `StockZhAGbjgEmItem` - 股本结构项
   - `StockZygcEmItem` - 主营构成项

2. **stock_fundamental_hold_sina.go** - 新浪财经机构持股接口：
   - `StockInstituteHold(symbol string)` - 机构持股一览表
   - `StockInstituteHoldDetail(stock, quarter string)` - 机构持股详情
   - `translateInstituteType(instituteType string)` - 机构类型翻译
   - 复用 `parseJSONP` 函数（来自 stock_fundamental_ipo_em.go）

3. **stock_fundamental_gbjg_zygc_em.go** - 东方财富股本结构和主营构成接口：
   - `StockZhAGbjgEm(symbol string)` - 股本结构
   - `StockZygcEm(symbol string)` - 主营构成
   - `translateCategoryType(categoryType string)` - 分类类型翻译
   - `cleanNumber(s string)` - 数字清理

4. **stock_fundamental_hold_test.go** - 测试用例：
   - 8个功能测试
   - 4个无效参数测试
   - 3个辅助函数测试

### 7.2 测试结果

所有测试通过：
- `TestStockInstituteHold` ✅ - 获取1795条记录
- `TestStockInstituteHoldDetail` ✅ - 获取16条记录
- `TestStockZhAGbjgEm` ✅ - 获取13条记录
- `TestStockZygcEm` ✅ - 获取73条记录
- 参数验证测试 ✅
- 辅助函数测试 ✅

### 7.3 技术要点

1. **HTML表格解析** - 使用 goquery 解析新浪财经的HTML表格
2. **JSONP解析** - 复用现有的 parseJSONP 函数处理新浪财经的JSONP响应
3. **JSON API** - 使用 gjson 解析东方财富的JSON API响应
4. **类型映射** - 实现机构类型和分类类型的中文翻译
5. **错误处理** - 完善的参数验证和错误处理

### 7.4 注意事项

- StockZygcEm 类型重命名为 StockZyjsEm 以避免与函数名冲突
- parseJSONP 函数在 stock_fundamental_ipo_em.go 中已存在，避免重复定义
- 某些股票可能没有机构持股数据，测试中已处理这种情况

---

*计划版本: v1.1*
*创建时间: 2026-01-18*
*完成时间: 2026-01-18*
