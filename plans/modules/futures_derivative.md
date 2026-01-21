# futures_derivative 模块计划

> 模块: 期货衍生数据
> 函数数量: 15 个
> 优先级: P2
> 预计周期: 第9周(与期货模块一起)

---

## 一、模块概述

期货衍生数据模块提供各交易所期货合约信息、主力合约识别、持仓分析、生猪期货相关数据等衍生分析功能。

---

## 二、Go 文件结构

```
futures_derivative/
├── contract_info.go    # 合约信息
├── main_contract.go    # 主力合约
├── hog.go              # 生猪数据
├── spot.go             # 现货数据
└── types.go            # 类型定义
```

---

## 三、函数清单

### 3.1 合约信息 (contract_info.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `futures_contract_info_shfe` | 上期所合约信息 | 上期所 | P0 |
| `futures_contract_info_dce` | 大商所合约信息 | 大商所 | P0 |
| `futures_contract_info_czce` | 郑商所合约信息 | 郑商所 | P0 |
| `futures_contract_info_cffex` | 中金所合约信息 | 中金所 | P0 |
| `futures_contract_info_gfex` | 广期所合约信息 | 广期所 | P1 |
| `futures_contract_info_ine` | 上期能源合约信息 | 上期能源 | P1 |

### 3.2 主力合约 (main_contract.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `zh_subscribe_exchange_symbol` | 交易所品种列表 | 新浪财经 | P0 |
| `match_main_contract` | 匹配主力合约 | 新浪财经 | P0 |
| `futures_display_main_sina` | 显示主力合约 | 新浪财经 | P0 |
| `futures_main_sina` | 主力合约数据 | 新浪财经 | P0 |
| `futures_hold_pos_sina` | 持仓排名数据 | 新浪财经 | P1 |

### 3.3 生猪数据 (hog.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `futures_hog_core` | 生猪核心数据 | 搜猪网 | P2 |
| `futures_hog_cost` | 生猪养殖成本 | 搜猪网 | P2 |
| `futures_hog_supply` | 生猪供应数据 | 搜猪网 | P2 |

### 3.4 现货数据 (spot.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `futures_spot_sys` | 期现价格系统 | 生意社 | P1 |

---

## 四、类型定义

```go
// types.go
package futures_derivative

import "time"

// ContractInfo 合约信息
type ContractInfo struct {
    Symbol          string    `json:"symbol"`           // 合约代码
    Name            string    `json:"name"`             // 合约名称
    Exchange        string    `json:"exchange"`         // 交易所
    ProductClass    string    `json:"product_class"`    // 品种分类
    Unit            int       `json:"unit"`             // 交易单位
    PriceTick       float64   `json:"price_tick"`       // 最小变动价位
    MarginRatio     float64   `json:"margin_ratio"`     // 保证金率
    LastTradeDate   time.Time `json:"last_trade_date"`  // 最后交易日
    DeliveryMonth   string    `json:"delivery_month"`   // 交割月份
}

// MainContract 主力合约
type MainContract struct {
    Symbol       string  `json:"symbol"`        // 品种代码
    MainSymbol   string  `json:"main_symbol"`   // 主力合约
    Price        float64 `json:"price"`         // 当前价格
    Change       float64 `json:"change"`        // 涨跌
    ChangeRate   float64 `json:"change_rate"`   // 涨跌幅
    Volume       int64   `json:"volume"`        // 成交量
    OpenInterest int64   `json:"open_interest"` // 持仓量
}

// HogData 生猪数据
type HogData struct {
    Date     time.Time `json:"date"`
    Province string    `json:"province"`
    Price    float64   `json:"price"`
    Change   float64   `json:"change"`
    Type     string    `json:"type"` // 外三元、内三元等
}

// SpotFuturesData 期现数据
type SpotFuturesData struct {
    Date         time.Time `json:"date"`
    Product      string    `json:"product"`
    SpotPrice    float64   `json:"spot_price"`
    FuturesPrice float64   `json:"futures_price"`
    Basis        float64   `json:"basis"`       // 基差
    BasisRate    float64   `json:"basis_rate"`  // 基差率
}
```

---

## 五、数据源

| 数据源 | URL模式 | 认证 |
|--------|---------|------|
| 上期所 | www.shfe.com.cn | 无 |
| 大商所 | www.dce.com.cn | 无 |
| 郑商所 | www.czce.com.cn | 无 |
| 中金所 | www.cffex.com.cn | 无 |
| 广期所 | www.gfex.com.cn | 无 |
| 新浪财经 | finance.sina.com.cn | 无 |
| 搜猪网 | www.soozhu.com | 无 |
| 生意社 | www.100ppi.com | 无 |

---

## 六、实现优先级

### P0 - 第一批实现
1. 各交易所合约信息
2. 主力合约识别

### P1 - 第二批实现
1. 持仓排名
2. 期现价格

### P2 - 第三批实现
1. 生猪相关数据

---

*文档版本: v1.0*
*最后更新: 2026-01-18*
