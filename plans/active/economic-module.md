# economic 模块实现计划

> 创建: 2026-01-18
> 状态: 进行中

## 目标

实现 akshare economic/ 模块的宏观经济数据接口（236个接口）。

## 模块分析

### 文件统计

| 文件名 | 函数数 | 主要数据源 |
|--------|--------|-----------|
| macro_china.py | 78 | 金十数据中心、东方财富 |
| macro_usa.py | 49 | 金十数据中心 |
| macro_euro.py | 16 | 金十数据中心 |
| macro_uk.py | 16 | 金十数据中心 |
| macro_bank.py | 12 | 金十数据中心-央行利率 |
| macro_canada.py | 10 | 金十数据中心 |
| macro_china_hk.py | 10 | 金十数据中心 |
| macro_germany.py | 9 | 金十数据中心 |
| macro_australia.py | 7 | 金十数据中心 |
| macro_swiss.py | 7 | 金十数据中心 |
| macro_japan.py | 6 | 金十数据中心 |
| macro_china_nbs.py | 5 | 国家统计局 |
| macro_constitute.py | 3 | 金十数据中心 |
| macro_finance_ths.py | 3 | 同花顺 |
| macro_info_ws.py | 3 | 金十数据中心 |
| macro_other.py | 2 | 金十数据中心 |
| **总计** | **236** | - |

### 数据源分类

1. **金十数据中心** (jin10.com) - 约220个接口
   - 统一的 API 接口模式
   - 需要特殊 headers: x-app-id, x-csrf-token, x-version
   - 数据格式统一: 日期、今值、预测值、前值
   - 分页获取数据

2. **国家统计局** (data.stats.gov.cn) - 5个接口
   - 需要处理树形结构数据
   - 使用 @lru_cache 缓存
   - HTTPS 证书验证问题
   - 复杂度较高

3. **东方财富/同花顺** - 约11个接口
   - 各自独立的 API 接口
   - 需要单独实现

### 复杂度评估

- **简单** (12个): macro_bank.py - 央行利率接口，只需修改参数
- **中等** (约210个): 金十数据中心接口 - 统一模式，易于批量实现
- **复杂** (5个): macro_china_nbs.py - 需要处理树形结构和缓存
- **中等** (9个): 其他数据源 - 需要单独分析实现

## 分批实施计划

### 第一批: 央行利率（12个接口）✅ 推荐先实现

**文件**: `economic/macro_bank_em.go`

**原因**:
- 最简单，统一使用 `__get_interest_rate_data` 基础函数
- 只需修改 `attr_id` 参数
- 可快速完成，建立信心

**接口清单**:
1. MacroBankUsaInterestRate - 美联储利率决议
2. MacroBankEuroInterestRate - 欧洲央行利率决议
3. MacroBankNewzealandInterestRate - 新西兰联储利率决议
4. MacroBankChinaInterestRate - 中国央行利率决议
5. MacroBankSwitzerlandInterestRate - 瑞士央行利率决议
6. MacroBankEnglishInterestRate - 英国央行利率决议
7. MacroBankAustraliaInterestRate - 澳洲联储利率决议
8. MacroBankJapanInterestRate - 日本央行利率决议
9. MacroBankRussiaInterestRate - 俄罗斯央行利率决议
10. MacroBankIndiaInterestRate - 印度央行利率决议
11. MacroBankBrazilInterestRate - 巴西央行利率决议
12. MacroBankSouthKoreaInterestRate - 韩国央行利率决议（如有）

**预计工时**: 2-3小时

---

### 第二批: 美国宏观数据（49个接口）

**文件**: `economic/macro_usa_em.go`

**接口特点**:
- 统一使用 `__macro_usa_base_func` 基础函数
- 数据格式一致
- 需要分页处理

**预计工时**: 8-10小时

---

### 第三批: 其他国家宏观数据（约60个接口）

**文件**:
- `economic/macro_euro_em.go` - 欧元区（16个）
- `economic/macro_uk_em.go` - 英国（16个）
- `economic/macro_canada_em.go` - 加拿大（10个）
- `economic/macro_australia_em.go` - 澳大利亚（7个）
- `economic/macro_japan_em.go` - 日本（6个）
- `economic/macro_germany_em.go` - 德国（9个）
- `economic/macro_swiss_em.go` - 瑞士（7个）

**预计工时**: 10-12小时

---

### 第四批: 中国宏观数据（78个接口）

**文件**: `economic/macro_china_em.go`

**接口特点**:
- 数量最多
- 数据源多样：金十数据中心、东方财富、统计局等
- 统一使用 `__macro_china_base_func` 基础函数
- 需要特殊处理的接口较多

**预计工时**: 12-15小时

---

### 第五批: 国家统计局数据（5个接口）

**文件**: `economic/macro_china_nbs.go`

**接口特点**:
- 复杂度最高
- 需要处理树形结构数据
- 使用 LRU 缓存
- HTTPS 证书验证问题
- 需要递归解析数据

**预计工时**: 4-5小时

---

### 第六批: 其他数据源（约32个接口）

**文件**:
- `economic/macro_china_hk_em.go` - 中国香港（10个）
- `economic/macro_constitute_em.go` - 成分股（3个）
- `economic/macro_finance_ths.go` - 同花顺财经（3个）
- `economic/macro_info_ws.go` - 信息（3个）
- `economic/macro_other_em.go` - 其他（2个）
- 其他杂项接口

**预计工时**: 6-8小时

---

## 技术要点

### 1. Jin10 数据中心接口

**通用 Headers**:
```go
headers := map[string]string{
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
    "x-app-id": "rU6QIu7JHe2gOUeR",
    "x-csrf-token": "x-csrf-token",
    "x-version": "1.0.0",
}
```

**分页逻辑**:
- 首次请求不带 `max_date` 参数
- 后续请求使用上一页最后一条数据的日期减1天作为 `max_date`
- 直到返回数据为空

**数据格式**:
```go
type MacroEconomicItem struct {
    Commodity string    // 商品/指标名称
    Date      time.Time // 日期
    Current   float64   // 今值
    Forecast  float64   // 预测值
    Previous  float64   // 前值
}
```

### 2. 国家统计局接口

**技术难点**:
- 树形结构数据递归解析
- LRU 缓存实现（可用 `groupcache/lru` 或自实现）
- HTTPS 证书验证问题（使用 `InsecureSkipVerify`）
- 复杂的参数构造

### 3. 文件命名规范

```
economic/
├── doc.go                      # 包文档
├── types.go                    # 公共数据类型
├── macro_bank_em.go            # 央行利率（第一批）
├── macro_usa_em.go             # 美国宏观数据（第二批）
├── macro_euro_em.go            # 欧元区（第三批）
├── macro_uk_em.go              # 英国（第三批）
├── macro_canada_em.go          # 加拿大（第三批）
├── macro_australia_em.go       # 澳大利亚（第三批）
├── macro_japan_em.go           # 日本（第三批）
├── macro_germany_em.go         # 德国（第三批）
├── macro_swiss_em.go           # 瑞士（第三批）
├── macro_china_em.go           # 中国宏观（第四批）
├── macro_china_nbs.go          # 国家统计局（第五批）
├── macro_china_hk_em.go        # 中国香港（第六批）
├── macro_other_em.go           # 其他（第六批）
└── economic_test.go            # 测试文件
```

## 实施步骤

### Step 1: 创建基础结构
- [x] 分析模块结构
- [x] 制定分批计划
- [x] 创建 economic/ 目录
- [x] 创建 doc.go
- [x] 创建 types.go

### Step 2: 第一批 - 央行利率（11个）
- [x] 实现基础函数 `getInterestRateData`
- [x] 实现11个央行利率接口
- [x] 编写测试用例
- [x] 测试通过

### Step 3: 第二批 - 美国宏观数据（49个）
- [x] 实现基础函数 `macroUsaBaseFunc`
- [x] 实现49个接口
- [x] 编写测试用例
- [x] 测试通过

### Step 4: 第三批 - 其他国家（约60个）
- [ ] 实现各国基础函数
- [ ] 实现各国接口
- [ ] 编写测试用例
- [ ] 测试通过

### Step 5: 第四批 - 中国宏观数据（78个）
- [ ] 实现基础函数 `macroChinaBaseFunc`
- [ ] 实现78个接口
- [ ] 编写测试用例
- [ ] 测试通过

### Step 6: 第五批 - 国家统计局（5个）
- [ ] 实现树形结构解析
- [ ] 实现缓存机制
- [ ] 实现5个接口
- [ ] 编写测试用例
- [ ] 测试通过

### Step 7: 第六批 - 其他数据源（约32个）
- [ ] 分析各数据源特点
- [ ] 实现各接口
- [ ] 编写测试用例
- [ ] 测试通过

### Step 8: 整体测试
- [ ] 运行全部测试
- [ ] 确保测试覆盖率 > 80%
- [ ] 代码格式检查
- [ ] 静态分析检查

## 验收标准

- [ ] 编译通过: `go build ./economic/...`
- [ ] 测试通过: `go test ./economic/...`
- [ ] 测试覆盖率 > 80%
- [ ] 代码格式规范: `go fmt ./economic/...`
- [ ] 静态检查通过: `go vet ./economic/...`
- [ ] 实现全部236个接口

## 进度跟踪

| 批次 | 接口数 | 状态 | 完成时间 |
|------|--------|------|----------|
| 创建计划文档 | - | ✅ 完成 | 2026-01-18 |
| Step 1: 基础结构 | - | ✅ 完成 | 2026-01-18 |
| 第一批: 央行利率 | 11 | ✅ 完成 | 2026-01-18 |
| 第二批: 美国宏观 | 49 | ✅ 完成 | 2026-01-18 |
| 第三批: 其他国家 | ~60 | 🔜 待开始 | - |
| 第四批: 中国宏观 | 78 | 🔜 待开始 | - |
| 第五批: 国家统计局 | 5 | 🔜 待开始 | - |
| 第六批: 其他数据源 | ~32 | 🔜 待开始 | - |
| 整体测试验收 | - | 🔜 待开始 | - |

**总进度**: 60/236 (25.4%)

## 风险和挑战

1. **接口数量大**: 236个接口需要较长时间完成，建议严格按批次实施
2. **数据源多样**: 不同数据源可能有不同的反爬机制
3. **国家统计局接口**: 复杂度高，可能需要额外时间调试
4. **测试覆盖**: 接口多，确保每个接口都有测试用例工作量大

## 备注

- 优先实现简单的央行利率接口，建立统一的实现模式
- 金十数据中心接口占比最大，实现统一的基础函数可大幅减少工作量
- 国家统计局接口最复杂，建议放在后期实现
- 每个批次完成后立即测试验收，避免问题累积

---

*计划版本: v1.0*
*创建时间: 2026-01-18*
*预计总工时: 42-53小时*
