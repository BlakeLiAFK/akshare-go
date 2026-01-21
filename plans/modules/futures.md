# futures 模块计划

> 模块: 期货数据
> 函数数量: 85 个
> 优先级: P0
> 预计周期: 第8-9周

---

## 一、模块概述

期货模块提供国内四大期货交易所（上期所、大商所、郑商所、中金所）及广期所的期货行情、持仓排名、仓单数据、现货基差等信息。

---

## 二、Go 文件结构

```
futures/
├── hist.go               # 期货历史行情
├── spot.go               # 期货实时行情
├── position.go           # 持仓排名
├── warehouse.go          # 仓单数据
├── delivery.go           # 交割数据
├── contract.go           # 合约信息
├── basis.go              # 基差数据
├── inventory.go          # 库存数据
├── foreign.go            # 外盘期货
├── global.go             # 全球期货
├── calendar.go           # 交易日历
├── utils.go              # 工具函数
└── types.go              # 类型定义
```

---

## 三、函数清单

### 3.1 期货历史行情 (hist.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `futures_zh_daily_sina` | 期货日K(新浪) | 新浪财经 | P0 |
| `futures_zh_minute_sina` | 期货分钟(新浪) | 新浪财经 | P0 |
| `futures_hist_em` | 期货历史(东财) | 东方财富 | P0 |
| `futures_hist_table_em` | 期货历史表(东财) | 东方财富 | P1 |
| `futures_global_hist_em` | 全球期货历史 | 东方财富 | P1 |
| `get_futures_daily` | 期货日数据 | 交易所 | P0 |
| `get_shfe_daily` | 上期所日数据 | 上期所 | P1 |
| `get_dce_daily` | 大商所日数据 | 大商所 | P1 |
| `get_czce_daily` | 郑商所日数据 | 郑商所 | P1 |
| `get_cffex_daily` | 中金所日数据 | 中金所 | P1 |
| `get_ine_daily` | 能源中心日数据 | 上期能源 | P1 |
| `get_gfex_daily` | 广期所日数据 | 广期所 | P1 |

### 3.2 期货实时行情 (spot.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `futures_zh_spot` | 期货实时行情 | 新浪财经 | P0 |
| `futures_zh_realtime` | 期货实时(详细) | 新浪财经 | P0 |
| `futures_global_spot_em` | 全球期货实时 | 东方财富 | P1 |
| `futures_main_sina` | 主力合约 | 新浪财经 | P0 |
| `futures_display_main_sina` | 主力合约展示 | 新浪财经 | P1 |
| `match_main_contract` | 匹配主力合约 | - | P1 |

### 3.3 持仓排名 (position.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `futures_dce_position_rank` | 大商所持仓排名 | 大商所 | P0 |
| `futures_dce_position_rank_other` | 大商所其他持仓 | 大商所 | P2 |
| `futures_gfex_position_rank` | 广期所持仓排名 | 广期所 | P1 |
| `futures_hold_pos_sina` | 持仓(新浪) | 新浪财经 | P1 |
| `get_cffex_rank_table` | 中金所持仓排名 | 中金所 | P1 |
| `get_shfe_rank_table` | 上期所持仓排名 | 上期所 | P1 |
| `get_dce_rank_table` | 大商所持仓排名 | 大商所 | P1 |
| `get_rank_table_czce` | 郑商所持仓排名 | 郑商所 | P1 |
| `get_rank_sum` | 持仓排名汇总 | - | P2 |
| `get_rank_sum_daily` | 每日持仓排名汇总 | - | P2 |

### 3.4 仓单数据 (warehouse.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `futures_shfe_warehouse_receipt` | 上期所仓单 | 上期所 | P1 |
| `futures_warehouse_receipt_czce` | 郑商所仓单 | 郑商所 | P1 |
| `futures_warehouse_receipt_dce` | 大商所仓单 | 大商所 | P1 |
| `futures_gfex_warehouse_receipt` | 广期所仓单 | 广期所 | P1 |
| `get_shfe_receipt_1` | 上期所仓单1 | 上期所 | P2 |
| `get_shfe_receipt_2` | 上期所仓单2 | 上期所 | P2 |
| `get_dce_receipt` | 大商所仓单 | 大商所 | P2 |
| `get_czce_receipt_1` | 郑商所仓单1 | 郑商所 | P2 |
| `get_czce_receipt_2` | 郑商所仓单2 | 郑商所 | P2 |
| `get_czce_receipt_3` | 郑商所仓单3 | 郑商所 | P2 |
| `get_gfex_receipt` | 广期所仓单 | 广期所 | P2 |
| `get_receipt` | 通用仓单获取 | - | P2 |

### 3.5 交割数据 (delivery.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `futures_delivery_shfe` | 上期所交割 | 上期所 | P2 |
| `futures_delivery_dce` | 大商所交割 | 大商所 | P2 |
| `futures_delivery_czce` | 郑商所交割 | 郑商所 | P2 |
| `futures_delivery_match_czce` | 郑商所交割配对 | 郑商所 | P2 |
| `futures_delivery_match_dce` | 大商所交割配对 | 大商所 | P2 |

### 3.6 合约信息 (contract.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `futures_comm_info` | 商品期货信息 | - | P1 |
| `futures_contract_detail` | 合约详情 | - | P1 |
| `futures_contract_detail_em` | 合约详情(东财) | 东方财富 | P1 |
| `futures_fees_info` | 期货手续费 | - | P2 |
| `futures_rule` | 交易规则 | - | P2 |
| `futures_rule_em` | 交易规则(东财) | 东方财富 | P2 |
| `futures_trading_hours_em` | 交易时间 | 东方财富 | P2 |
| `futures_symbol_mark` | 合约标识 | - | P2 |
| `symbol_varieties` | 品种列表 | - | P1 |
| `symbol_market` | 品种市场 | - | P2 |

### 3.7 基差数据 (basis.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `futures_to_spot_shfe` | 上期所期现价差 | 上期所 | P1 |
| `futures_to_spot_dce` | 大商所期现价差 | 大商所 | P1 |
| `futures_to_spot_czce` | 郑商所期现价差 | 郑商所 | P1 |
| `futures_spot_price` | 现货价格 | - | P1 |
| `futures_spot_price_daily` | 现货日价格 | - | P1 |
| `futures_spot_price_previous` | 前期现货价格 | - | P2 |
| `futures_spot_stock` | 现货库存 | - | P2 |
| `futures_spot_sys` | 现货系统 | - | P2 |
| `get_roll_yield` | 展期收益率 | - | P2 |
| `get_roll_yield_bar` | 展期收益率图 | - | P2 |

### 3.8 库存数据 (inventory.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `futures_inventory_99` | 99期货库存 | 99期货 | P2 |
| `futures_inventory_em` | 库存(东财) | 东方财富 | P1 |
| `futures_comex_inventory` | COMEX库存 | - | P2 |
| `futures_stock_shfe_js` | 上期所库存JS | 上期所 | P2 |

### 3.9 外盘期货 (foreign.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `futures_foreign_hist` | 外盘期货历史 | - | P1 |
| `futures_foreign_detail` | 外盘期货详情 | - | P1 |
| `futures_foreign_commodity_realtime` | 外盘商品实时 | - | P1 |
| `futures_foreign_commodity_subscribe_exchange_symbol` | 外盘订阅 | - | P2 |
| `futures_settlement_price_sgx` | 新加坡交易所结算价 | SGX | P2 |
| `futures_news_shmet` | 上海有色网新闻 | 上海有色 | P3 |

### 3.10 全球期货 (global.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `futures_index_ccidx` | 商品指数 | 中期协 | P2 |

### 3.11 交易日历 (calendar.go)

| 函数名 | 功能 | 数据源 | 优先级 |
|--------|------|--------|--------|
| `get_calendar` | 获取交易日历 | - | P1 |
| `last_trading_day` | 上一交易日 | - | P1 |
| `get_latest_data_date` | 最新数据日期 | - | P2 |

### 3.12 工具函数 (utils.go)

| 函数名 | 功能 | 优先级 |
|--------|------|--------|
| `chinese_to_english` | 中英文转换 | P2 |
| `convert_date` | 日期转换 | P2 |
| `find_chinese` | 查找中文 | P3 |
| `get_json_path` | 获取JSON路径 | P3 |
| `get_pk_data` | 获取PK数据 | P3 |
| `get_pk_path` | 获取PK路径 | P3 |
| `futures_hq_subscribe_exchange_symbol` | 行情订阅 | P2 |
| `zh_subscribe_exchange_symbol` | 中国交易所订阅 | P2 |
| `pandas_read_html_link` | 读取HTML表格 | P3 |
| `requests_link` | 请求链接 | P3 |

---

## 四、数据源 URL 模式

### 4.1 交易所官网

| 交易所 | URL 模式 | 用途 |
|--------|----------|------|
| 上期所 | `www.shfe.com.cn/data/` | 日数据、持仓、仓单 |
| 大商所 | `www.dce.com.cn/publicweb/` | 日数据、持仓、仓单 |
| 郑商所 | `www.czce.com.cn/cn/` | 日数据、持仓、仓单 |
| 中金所 | `www.cffex.com.cn/` | 日数据、持仓 |
| 广期所 | `www.gfex.com.cn/` | 日数据、持仓、仓单 |
| 能源中心 | `www.ine.cn/` | 原油期货 |

### 4.2 新浪财经

| URL 模式 | 用途 |
|----------|------|
| `hq.sinajs.cn/list=` | 实时行情 |
| `vip.stock.finance.sina.com.cn/` | 历史数据 |

---

## 五、类型定义示例

```go
// types.go
package futures

import "time"

// FuturesQuote 期货实时行情
type FuturesQuote struct {
    Symbol    string    `json:"symbol"`     // 合约代码
    Name      string    `json:"name"`       // 合约名称
    Exchange  string    `json:"exchange"`   // 交易所
    Open      float64   `json:"open"`       // 开盘价
    High      float64   `json:"high"`       // 最高价
    Low       float64   `json:"low"`        // 最低价
    Close     float64   `json:"close"`      // 收盘价
    PreSettle float64   `json:"pre_settle"` // 昨结算
    Settle    float64   `json:"settle"`     // 今结算
    Change    float64   `json:"change"`     // 涨跌
    ChangePct float64   `json:"change_pct"` // 涨跌幅
    Volume    int64     `json:"volume"`     // 成交量
    Amount    float64   `json:"amount"`     // 成交额
    OpenInt   int64     `json:"open_int"`   // 持仓量
    Time      time.Time `json:"time"`
}

// FuturesKLine 期货K线
type FuturesKLine struct {
    Date      time.Time `json:"date"`
    Open      float64   `json:"open"`
    High      float64   `json:"high"`
    Low       float64   `json:"low"`
    Close     float64   `json:"close"`
    Volume    int64     `json:"volume"`
    OpenInt   int64     `json:"open_int"`
    Settle    float64   `json:"settle"`
}

// PositionRank 持仓排名
type PositionRank struct {
    Date       time.Time `json:"date"`
    Symbol     string    `json:"symbol"`
    Rank       int       `json:"rank"`
    Member     string    `json:"member"`      // 会员名称
    Volume     int64     `json:"volume"`      // 成交量
    VolumeChg  int64     `json:"volume_chg"`  // 成交量变化
    LongPos    int64     `json:"long_pos"`    // 多头持仓
    LongChg    int64     `json:"long_chg"`    // 多头变化
    ShortPos   int64     `json:"short_pos"`   // 空头持仓
    ShortChg   int64     `json:"short_chg"`   // 空头变化
}

// WarehouseReceipt 仓单数据
type WarehouseReceipt struct {
    Date       time.Time `json:"date"`
    Symbol     string    `json:"symbol"`
    Warehouse  string    `json:"warehouse"`  // 仓库
    Quantity   int64     `json:"quantity"`   // 数量
    QuantityChg int64    `json:"quantity_chg"` // 变化
    Unit       string    `json:"unit"`       // 单位
}

// ContractInfo 合约信息
type ContractInfo struct {
    Symbol       string    `json:"symbol"`
    Name         string    `json:"name"`
    Exchange     string    `json:"exchange"`
    Unit         float64   `json:"unit"`         // 交易单位
    PriceUnit    string    `json:"price_unit"`   // 报价单位
    MinChange    float64   `json:"min_change"`   // 最小变动
    DailyLimit   float64   `json:"daily_limit"`  // 涨跌停板
    LastTrading  string    `json:"last_trading"` // 最后交易日
    Delivery     string    `json:"delivery"`     // 交割日
    Margin       float64   `json:"margin"`       // 保证金
    Commission   float64   `json:"commission"`   // 手续费
}
```

---

## 六、实现优先级

### P0 - 第一批实现

1. `futures_zh_spot` - 期货实时行情
2. `futures_zh_daily_sina` - 期货日K
3. `futures_main_sina` - 主力合约
4. `futures_dce_position_rank` - 持仓排名
5. `get_futures_daily` - 期货日数据

### P1 - 第二批实现

1. 各交易所持仓排名
2. 仓单数据
3. 合约信息
4. 基差数据
5. 外盘期货

### P2 - 第三批实现

1. 交割数据
2. 库存数据
3. 工具函数

---

*文档版本: v1.0*
*最后更新: 2026-01-18*
