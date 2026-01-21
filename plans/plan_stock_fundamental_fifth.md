# stock_fundamental 模块第五批接口实现计划

> 任务: 实现同花顺、经济通、雪球等数据源的5个基本面接口
> 创建时间: 2025-01-18
> 状态: 进行中

## 一、任务概述

实现以下5个股票基本面数据接口：

1. **StockZyjsThs** - 同花顺主营介绍
2. **StockProfitForecastThs** - 同花顺盈利预测
3. **StockHkProfitForecastEt** - 经济通港股盈利预测
4. **StockIndividualBasicInfoXq** - 雪球A股公司简介
5. **StockIndividualBasicInfoUsXq** - 雪球美股公司简介

## 二、技术要点

### 2.1 数据源分析

| 接口 | 数据源 | URL | 编码 | 解析方式 |
|------|--------|-----|------|----------|
| StockZyjsThs | 同花顺 | basic.10jqka.com.cn | GB2312 | HTML列表 |
| StockProfitForecastThs | 同花顺 | basic.10jqka.com.cn | GBK | HTML表格 |
| StockHkProfitForecastEt | 经济通 | etnet.com.hk | UTF-8 | HTML表格 |
| StockIndividualBasicInfoXq | 雪球 | stock.xueqiu.com | UTF-8 | JSON |
| StockIndividualBasicInfoUsXq | 雪球 | stock.xueqiu.com | UTF-8 | JSON |

### 2.2 返回值类型

| 接口 | 返回类型 | 说明 |
|------|----------|------|
| StockZyjsThs | map[string]string | 键值对形式 |
| StockProfitForecastThs | dataframe.DataFrame | 动态列表格 |
| StockHkProfitForecastEt | []map[string]interface{} | 动态结构 |
| StockIndividualBasicInfoXq | []StockIpoInfoItem | 复用类型 |
| StockIndividualBasicInfoUsXq | []StockIpoInfoItem | 复用类型 |

## 三、实现步骤

### 步骤1: 添加类型定义
- [ ] 在 types.go 中添加需要的结构体类型

### 步骤2: 实现同花顺主营介绍
- [ ] 创建 stock_fundamental_zyjs_ths.go
- [ ] 实现 StockZyjsThs 函数
- [ ] 处理 GB2312 编码

### 步骤3: 实现同花顺盈利预测
- [ ] 创建 stock_fundamental_profit_forecast_ths.go
- [ ] 实现 StockProfitForecastThs 函数
- [ ] 支持4种 indicator 类型
- [ ] 处理动态表格索引

### 步骤4: 实现经济通港股盈利预测
- [ ] 创建 stock_fundamental_profit_forecast_hk_etnet.go
- [ ] 实现 StockHkProfitForecastEt 函数
- [ ] 支持4种 indicator 类型

### 步骤5: 实现雪球公司简介
- [ ] 创建 stock_fundamental_basic_info_xq.go
- [ ] 实现 StockIndividualBasicInfoXq 函数
- [ ] 实现 StockIndividualBasicInfoUsXq 函数
- [ ] 处理 cookie 认证

### 步骤6: 编写测试用例
- [ ] 创建 stock_fundamental_fifth_test.go
- [ ] 测试所有5个接口

### 步骤7: 编译验证
- [ ] 运行 go build 确保编译通过
- [ ] 运行 go test 确保测试通过

## 四、注意事项

1. **编码处理**: GB2312/GBK 需要正确解码
2. **HTML解析**: 使用 goquery 处理中文编码
3. **动态返回**: 不同 indicator 返回不同结构
4. **错误处理**: 提供有意义的错误信息
5. **注释规范**: 使用中文注释

## 五、文件清单

| 文件 | 说明 |
|------|------|
| stock_fundamental_zyjs_ths.go | 同花顺主营介绍 |
| stock_fundamental_profit_forecast_ths.go | 同花顺盈利预测 |
| stock_fundamental_profit_forecast_hk_etnet.go | 经济通港股盈利预测 |
| stock_fundamental_basic_info_xq.go | 雪球公司简介 |
| stock_fundamental_fifth_test.go | 测试用例 |

## 六、进度跟踪

- [x] 任务规划完成
- [ ] 类型定义添加
- [ ] StockZyjsThs 实现
- [ ] StockProfitForecastThs 实现
- [ ] StockHkProfitForecastEt 实现
- [ ] StockIndividualBasicInfoXq 实现
- [ ] StockIndividualBasicInfoUsXq 实现
- [ ] 测试用例编写
- [ ] 编译测试通过
