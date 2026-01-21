# tool/ 模块移植计划

> 创建日期: 2026-01-18
> 模块: tool (工具函数)
> 来源: Python akshare/tool/
> 执行顺序: Z-A (第1个)

---

## 一、模块概述

### 1.1 模块信息

```
模块名: tool
功能: 工具函数
Python源码: _akshare_source/akshare/tool/
Go目标: tool/
```

### 1.2 接口清单

| 序号 | 文件名 | Python函数名 | Go函数名 | 接口数 | 数据源 |
|------|--------|--------------|----------|--------|--------|
| 1 | trade_date_hist.py | tool_trade_date_hist_sina | ToolTradeDateHistSina | 1 | 新浪财经 |

**总计**: 1 个接口

---

## 二、详细设计

### 2.1 ToolTradeDateHistSina

**Python 原版**:
```python
def tool_trade_date_hist_sina() -> pd.DataFrame
```

**Go 版本**:
```go
func ToolTradeDateHistSina() ([]TradeDate, error)
```

**功能描述**:
- 新浪财经-交易日历-历史数据
- 返回所有交易日列表

**数据源**: `https://finance.sina.com.cn/realstock/company/klc_td_sh.txt`

**返回类型**:
```go
// TradeDate 交易日
type TradeDate struct {
    Date time.Time `json:"date"` // 交易日期
}
```

**实现要点**:
1. 使用 JS 引擎 (goja) 执行解码函数
2. 解码新浪加密的交易日期数据
3. 补充缺失的日期 (1992-05-04)
4. 排序并返回

---

## 三、技术方案

### 3.1 依赖包

```go
import (
    "time"
    "github.com/dop251/goja"  // JS 引擎
    "github.com/BlakeLiAFK/akshare-go/utils"
)
```

### 3.2 JS 解码函数

需要在 Go 中嵌入 Python 版本的 `hk_js_decode` 函数。

### 3.3 文件结构

```
tool/
├── doc.go                      # 包文档 (已存在)
├── types.go                    # 类型定义 (需更新)
├── tool_trade_date_hist_sina.go # 交易日历 (新建)
└── tool_test.go                # 测试文件 (新建)
```

---

## 四、开发步骤

### Step 1: 更新 types.go

添加 `TradeDate` 类型定义。

### Step 2: 实现 tool_trade_date_hist_sina.go

1. 定义 JS 解码函数常量
2. 实现 `ToolTradeDateHistSina()` 函数
3. 处理日期排序和缺失数据

### Step 3: 编写测试

1. 测试正常获取
2. 验证日期格式
3. 验证日期排序
4. 验证特殊日期 (1992-05-04)

---

## 五、验收标准

- [ ] 编译通过: `go build ./tool/...`
- [ ] 测试通过: `go test ./tool/...`
- [ ] 测试覆盖率 > 80%
- [ ] 代码格式规范: `go fmt ./tool/...`
- [ ] 静态检查通过: `go vet ./tool/...`

---

## 六、进度跟踪

| 步骤 | 状态 | 完成时间 |
|------|------|----------|
| 创建计划文档 | ✅ 完成 | 2026-01-18 |
| 更新 types.go | ✅ 完成 | 2026-01-18 |
| 实现函数 | ✅ 完成 | 2026-01-18 |
| 编写测试 | ✅ 完成 | 2026-01-18 |
| 验收测试 | ✅ 完成 | 2026-01-18 |

---

*计划版本: v1.0*
*创建时间: 2026-01-18*
