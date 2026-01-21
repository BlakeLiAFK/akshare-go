# cal 模块实现计划

> 创建: 2026-01-18
> 状态: 部分完成（1/3 接口）

## 目标

实现 akshare cal/ 模块的已实现波动率（Realized Volatility）计算接口。

## 接口清单

### 已实现（1个）

1. ✅ **VolatilityYZRV** - Yang-Zhang 已实现波动率计算
   - 函数：`VolatilityYZRV(data dataframe.DataFrame) (dataframe.DataFrame, error)`
   - 说明：纯数学计算函数，不依赖外部API
   - 状态：已完成并测试通过

### 待实现（2个）- 依赖其他模块

2. ⏳ **RVFromStockZhAHistMinEM** - 从股票分钟数据计算波动率
   - 函数：`RVFromStockZhAHistMinEM(symbol, startDate, endDate, period, adjust string) (dataframe.DataFrame, error)`
   - 依赖：`stock_feature.StockZhAHistMinEM`
   - 状态：占位符已创建，等待 stock_feature 模块实现
   - Python源码：`akshare/cal/rv.py:13`

3. ⏳ **RVFromFuturesZhMinuteSina** - 从期货分钟数据计算波动率
   - 函数：`RVFromFuturesZhMinuteSina(symbol, period string) (dataframe.DataFrame, error)`
   - 依赖：`futures.FuturesZhMinuteSina`
   - 状态：占位符已创建，等待 futures 模块实现
   - Python源码：`akshare/cal/rv.py:61`

## 实现策略

由于 cal 模块依赖于尚未实现的 futures 和 stock_feature 模块，采用以下策略：

1. **优先实现独立函数**：VolatilityYZRV 不依赖其他模块，已完整实现
2. **创建占位符**：为依赖其他模块的函数创建占位符，返回明确的依赖错误
3. **延迟完善**：等待依赖模块实现后再补充完整功能

## 理论基础

Yang-Zhang 已实现波动率公式：

```
RV^2 = Vo + k*Vc + (1-k)*Vrs
```

其中：
- **Vo**: 隔夜波动率 = 1/(n-1) * Σ(Oi-Ōi)²
- **Vc**: 收盘波动率 = 1/(n-1) * Σ(ci-c̄i)²
- **k**: 权重系数 = 0.34 / (1.34 + (n+1)/(n-1))
- **Vrs**: Rogers-Satchell 波动率 = ui(ui-ci) + di(di-ci)
  - ui = ln(Hi/Oi)
  - ci = ln(Ci/Oi)
  - di = ln(Li/Oi)
  - oi = ln(Oi/Ci-1)

## 文件结构

```
cal/
├── doc.go          # 包文档
├── rv.go           # 已实现波动率接口实现
└── rv_test.go      # 测试文件
```

## 测试覆盖

- ✅ 依赖错误测试（2个）
- ✅ 正常计算测试（1个）
- ✅ 空数据测试（1个）
- ✅ 缺少列测试（1个）

总计：5个测试用例，全部通过

## 完成标准

- [x] 实现 VolatilityYZRV 函数
- [x] 为依赖函数创建占位符
- [x] 编写测试用例
- [x] 所有测试通过
- [x] 代码编译通过
- [ ] 补充剩余2个函数（等待依赖模块）

## 备注

**依赖关系**：
- `stock_feature` 模块：预计在字母表中较后实现
- `futures` 模块：预计在字母表中较后实现

**后续行动**：
1. 继续按字母顺序实现 crypto/ 模块
2. 当 futures/ 和 stock_feature/ 实现后，回来补充 cal/ 的剩余函数
