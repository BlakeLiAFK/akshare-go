# AKShare-Go 接口文档

> 本文档包含所有模块的接口说明和使用示例

## 📚 模块索引

### 股票相关

| 模块 | 说明 | 接口数量 | 文档链接 |
|-----|------|---------|---------|
| stock | 股票数据 | 167 | [查看文档](stock.md) |
| stock_feature | 股票特色数据 | 148 | [查看文档](stock_feature.md) |
| stock_fundamental | 股票基本面 | 55 | [查看文档](stock_fundamental.md) |
| stock_a | A股数据 | 3 | [查看文档](stock_a.md) |

### 基金相关

| 模块 | 说明 | 接口数量 | 文档链接 |
|-----|------|---------|---------|
| fund | 基金数据 | 86 | [查看文档](fund.md) |
| qdii | QDII基金 | 3 | [查看文档](qdii.md) |

### 期货相关

| 模块 | 说明 | 接口数量 | 文档链接 |
|-----|------|---------|---------|
| futures | 期货数据 | 53 | [查看文档](futures.md) |
| futures_derivative | 期货衍生 | 15 | [查看文档](futures_derivative.md) |
| qhkc | 期货会员持仓 | 5 | [查看文档](qhkc.md) |
| qhkc_web | 期货会员持仓(网页版) | 3 | [查看文档](qhkc_web.md) |

### 债券相关

| 模块 | 说明 | 接口数量 | 文档链接 |
|-----|------|---------|---------|
| bond | 债券数据 | 42 | [查看文档](bond.md) |
| reits | REITs数据 | 3 | [查看文档](reits.md) |

### 指数&期权

| 模块 | 说明 | 接口数量 | 文档链接 |
|-----|------|---------|---------|
| index | 指数数据 | 105 | [查看文档](index.md) |
| option | 期权数据 | 39 | [查看文档](option.md) |

### 宏观经济

| 模块 | 说明 | 接口数量 | 文档链接 |
|-----|------|---------|---------|
| economic | 宏观经济 | 230 | [查看文档](economic.md) |
| pro | 专业数据 | 14 | [查看文档](pro.md) |
| interest_rate | 利率数据 | 1 | [查看文档](interest_rate.md) |

### 外汇&加密货币

| 模块 | 说明 | 接口数量 | 文档链接 |
|-----|------|---------|---------|
| forex | 外汇数据 | 3 | [查看文档](forex.md) |
| fx | 外汇行情 | 6 | [查看文档](fx.md) |
| currency | 货币汇率 | 7 | [查看文档](currency.md) |
| crypto | 加密货币 | 2 | [查看文档](crypto.md) |

### 现货&能源

| 模块 | 说明 | 接口数量 | 文档链接 |
|-----|------|---------|---------|
| spot | 现货数据 | 7 | [查看文档](spot.md) |
| energy | 能源数据 | 8 | [查看文档](energy.md) |

### 金融机构

| 模块 | 说明 | 接口数量 | 文档链接 |
|-----|------|---------|---------|
| bank | 银行数据 | 4 | [查看文档](bank.md) |
| hf | 对冲基金 | 1 | [查看文档](hf.md) |

### 其他数据

| 模块 | 说明 | 接口数量 | 文档链接 |
|-----|------|---------|---------|
| air | 空气质量 | 8 | [查看文档](air.md) |
| article | 文章资讯 | 7 | [查看文档](article.md) |
| cal | 日历数据 | 3 | [查看文档](cal.md) |
| event | 事件数据 | 6 | [查看文档](event.md) |
| fortune | 财富榜单 | 6 | [查看文档](fortune.md) |
| movie | 电影票房 | 12 | [查看文档](movie.md) |
| news | 新闻资讯 | 6 | [查看文档](news.md) |
| nlp | NLP工具 | 6 | [查看文档](nlp.md) |
| rate | 利率数据 | 4 | [查看文档](rate.md) |
| tool | 工具函数 | 1 | [查看文档](tool.md) |
| other | 其他数据 | 7 | [查看文档](other.md) |

## 📊 统计信息

- **模块总数**: 36 个
- **接口总数**: 1065+ 个
- **覆盖范围**: 股票、基金、期货、债券、指数、期权、外汇、宏观经济等

## 🚀 快速开始

### 安装

```bash
go get github.com/BlakeLiAFK/akshare
```

### 示例

```go
package main

import (
    "fmt"
    "github.com/BlakeLiAFK/akshare/stock"
)

func main() {
    // 获取股票实时行情
    data, err := stock.StockZhASpotEm()
    if err != nil {
        panic(err)
    }

    fmt.Println(data)
}
```

## 📖 文档说明

每个模块的文档包含：

- **接口名称**: Go函数名
- **描述**: 接口功能说明
- **数据源**: 原始数据网站地址
- **输入参数**: 参数名称、类型和说明
- **输出参数**: 返回数据的字段说明
- **代码示例**: 完整的Go代码示例

## 🔗 相关链接

- [GitHub 仓库](https://github.com/BlakeLiAFK/akshare)
- [Python 原版 AKShare](https://github.com/akfamily/akshare)
- [AKShare 官方文档](https://akshare.akfamily.xyz/)

## 📝 贡献

欢迎提交 Issue 和 Pull Request 改进文档！

## ⚠️ 声明

所有数据仅供学术研究使用，不构成任何投资建议。
