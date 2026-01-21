# crypto 模块实现计划

> 创建: 2026-01-18
> 状态: ✅ 已完成（2/2 接口）

## 目标

实现 akshare crypto/ 模块的加密货币数据接口。

## 接口清单

### 已实现（2个）

1. ✅ **CryptoBitcoinCME** - 芝加哥商业交易所比特币成交量报告
   - 函数：`CryptoBitcoinCME(date string) (dataframe.DataFrame, error)`
   - 参数：date - 日期格式"20230830"
   - 数据源：金十数据 (https://datacenter.jin10.com)
   - Python源码：`akshare/crypto/crypto_bitcoin_cme.py:13`
   - 状态：✅ 已完成并测试通过

2. ✅ **CryptoBitcoinHoldReport** - 比特币持仓报告
   - 函数：`CryptoBitcoinHoldReport() (dataframe.DataFrame, error)`
   - 数据源：金十数据 (https://datacenter.jin10.com)
   - Python源码：`akshare/crypto/crypto_hold.py:13`
   - 状态：✅ 已完成并测试通过

## 实现细节

### CryptoBitcoinCME

**API信息**:
- URL: `https://datacenter-api.jin10.com/reports/list`
- 参数: `category=cme`, `date=YYYY-MM-DD`, `attr_id=4`
- 需要特定 Headers: `X-App-Id`, `X-Version` 等

**数据字段**:
- 商品
- 类型
- 电子交易合约
- 场内成交合约
- 场外成交合约
- 成交量
- 未平仓合约
- 持仓变化

**实现要点**:
- 日期格式转换：`20230830` → `2023-08-30`
- 动态解析列名（从 `data.keys` 获取）
- 动态解析数据（从 `data.values` 获取）

### CryptoBitcoinHoldReport

**API信息**:
- URL: `https://datacenter-api.jin10.com/bitcoin_treasuries/list`
- Headers: `X-App-Id: lnFP5lxse24wPgtY`, `X-Version: 1.0.0`

**数据字段**（14个）:
- 代码
- 公司名称-英文
- 公司名称-中文
- 国家/地区
- 市值
- 比特币占市值比重
- 持仓成本
- 持仓占比
- 持仓量
- 当日持仓市值
- 查询日期
- 公告链接
- 分类
- 倍数

**实现要点**:
- 原始数据有16个字段，需要映射到14个有效字段
- 列索引映射：`[0, 1, 15, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 13]`
- 跳过索引 11 和 14（占位符字段）

## 文件结构

```
crypto/
├── doc.go                  # 包文档
├── crypto_bitcoin.go       # 比特币数据接口实现
└── crypto_bitcoin_test.go  # 测试文件
```

## 测试结果

```bash
=== RUN   TestCryptoBitcoinCME
    CME比特币成交量报告数据行数: 5, 列数: 8
    列名: [商品 类型 电子交易合约 场内成交合约 场外成交合约 成交量 未平仓合约 持仓变化]
--- PASS: TestCryptoBitcoinCME (1.65s)

=== RUN   TestCryptoBitcoinHoldReport
    比特币持仓报告数据行数: 59, 列数: 14
    列名: [代码 公司名称-英文 公司名称-中文 国家/地区 市值 比特币占市值比重 持仓成本 持仓占比 持仓量 当日持仓市值 查询日期 公告链接 分类 倍数]
--- PASS: TestCryptoBitcoinHoldReport (0.34s)

=== RUN   TestCryptoBitcoinCME_InvalidDate
    正确处理无效日期格式: 日期格式错误，应为YYYYMMDD格式，如：20230830
--- PASS: TestCryptoBitcoinCME_InvalidDate (0.00s)

PASS
ok  	github.com/BlakeLiAFK/akshare-go/crypto	2.232s
```

## 测试覆盖

- ✅ 正常请求测试（2个）
- ✅ 数据验证测试（列名、行数）
- ✅ 错误处理测试（无效日期格式）

总计：3个测试用例，全部通过

## 完成标准

- [x] 实现 2 个接口函数
- [x] 编写测试用例
- [x] 所有测试通过
- [x] 代码编译通过
- [x] 代码规范检查通过

## 备注

**数据源特点**:
- 金十数据提供实时更新的加密货币数据
- 需要设置特定的 Headers（X-App-Id, X-Version）才能访问
- API 返回 JSON 格式数据

**Go实现优化**:
- 使用 `utils.GetWithHeaders` 统一处理 HTTP 请求
- 使用 `gjson` 高效解析 JSON 数据
- 使用 `gota/dataframe` 构建表格数据

**下一步**:
继续按字母顺序实现 currency/ 模块
