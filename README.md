# AKShare - Go

[![Go Version](https://img.shields.io/badge/Go-1.21%2B-blue)](https://go.dev/)
[![Go Reference](https://pkg.go.dev/badge/github.com/BlakeLiAFK/akshare.svg)](https://pkg.go.dev/github.com/BlakeLiAFK/akshare)
[![MIT License](https://img.shields.io/badge/license-MIT-blue)](https://github.com/BlakeLiAFK/akshare/blob/master/LICENSE)
[![akshare](https://img.shields.io/badge/Data%20Science-AKShare-green)](https://github.com/akfamily/akshare)

## 概述

[AKShare](https://github.com/akfamily/akshare) 的 Go 语言实现版本，提供简单易用的金融数据接口。

**Write less, get more!**

- 原始 Python 版本: [AKShare](https://github.com/akfamily/akshare)
- 官方文档: [中文文档](https://akshare.akfamily.xyz/)

## 安装

### 要求

- Go 1.21 或更高版本

### 使用 go get 安装

```bash
go get github.com/BlakeLiAFK/akshare
```

### 使用 go mod

在项目中添加依赖：

```bash
go mod init your-project
go get github.com/BlakeLiAFK/akshare
```

## 快速开始

### 获取股票历史数据

```go
package main

import (
    "fmt"
    "github.com/BlakeLiAFK/akshare/stock"
)

func main() {
    // 获取平安银行(000001)的历史行情数据
    df, err := stock.StockZhAHist(
        "000001",           // 股票代码
        "daily",            // 周期: daily, weekly, monthly
        "20230101",         // 开始日期
        "20231231",         // 结束日期
        "qfq",              // 复权类型: qfq(前复权), hfq(后复权), ""(不复权)
    )
    if err != nil {
        panic(err)
    }

    fmt.Println(df)
}
```

### 获取实时行情

```go
package main

import (
    "fmt"
    "github.com/BlakeLiAFK/akshare/stock"
)

func main() {
    // 获取沪深A股实时行情
    df, err := stock.StockZhASpotEm()
    if err != nil {
        panic(err)
    }

    fmt.Println(df)
}
```

## 主要模块

本项目包含以下数据模块：

| 模块 | 说明 | 接口数量 |
|-----|------|---------|
| `stock` | 股票数据 | 128 |
| `stock_feature` | 股票特色数据 | 207 |
| `stock_fundamental` | 股票基本面 | 56 |
| `fund` | 基金数据 | 84 |
| `futures` | 期货数据 | 85 |
| `futures_derivative` | 期货衍生 | 15 |
| `bond` | 债券数据 | 42 |
| `index` | 指数数据 | 95 |
| `option` | 期权数据 | 52 |
| `economic` | 宏观经济 | 230 |
| `forex` | 外汇数据 | 2 |
| `crypto` | 加密货币 | 2 |
| `air` | 空气质量 | 9 |
| 其他模块 | ... | ... |

**总计: 1154+ 数据接口**

## 特性

- **易于使用**: 一行代码获取数据
- **类型安全**: 利用 Go 的静态类型系统
- **高性能**: Go 语言的高并发特性
- **完整覆盖**: 对标 Python 版本的所有接口

## 数据源

主要数据来源：

- 东方财富 (eastmoney.com)
- 新浪财经 (sina.com.cn)
- 同花顺 (10jqka.com.cn)
- 证券交易所官网
- 期货交易所官网
- 其他权威金融数据网站

## 文档

详细的接口文档请参考：

- [Go Package 文档](https://pkg.go.dev/github.com/BlakeLiAFK/akshare)
- [Python 版本文档](https://akshare.akfamily.xyz/) (接口设计保持一致)

## 贡献

欢迎提交 Issue 和 Pull Request：

- 报告或修复 Bug
- 请求或发布新接口
- 改进文档
- 添加测试用例

## 声明

1. 本项目提供的所有数据仅供学术研究使用
2. 数据仅供参考，不构成任何投资建议
3. 任何基于本项目的投资决策，请注意数据风险
4. 本项目将持续提供开源金融数据
5. 由于不可控因素，部分接口可能会失效
6. 请遵守本项目使用的开源协议
7. 本项目是 [Python AKShare](https://github.com/akfamily/akshare) 的 Go 语言实现

## 许可证

[MIT License](LICENSE)

## 致谢

特别感谢 [AKShare](https://github.com/akfamily/akshare) 项目提供的数据接口设计和文档支持。

感谢以下数据提供方：

- 东方财富网
- 新浪财经
- 同花顺
- 上海证券交易所
- 深圳证券交易所
- 中国金融期货交易所
- 上海期货交易所
- 大连商品交易所
- 郑州商品交易所
- 以及其他所有数据源

## 引用

如需在论文中引用本项目：

```bibtex
@misc{akshare-go,
    author = {Blake Li},
    title = {AKShare-Go: Go Implementation of AKShare},
    year = {2024},
    publisher = {GitHub},
    journal = {GitHub repository},
    howpublished = {\url{https://github.com/BlakeLiAFK/akshare}},
}
```

## 展示你的项目

在你的项目中使用徽章：

```markdown
[![Data: akshare](https://img.shields.io/badge/Data%20Science-AKShare-green)](https://github.com/akfamily/akshare)
```

效果:

[![Data: akshare](https://img.shields.io/badge/Data%20Science-AKShare-green)](https://github.com/akfamily/akshare)
