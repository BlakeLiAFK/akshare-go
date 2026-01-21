# akshare-go 项目说明

## 项目概述

用 Go 语言完整复刻 Python akshare 库 (1154 个金融数据接口)。

## 技术栈

- **语言**: Go 1.21+
- **HTTP**: `net/http` + `resty/v2`
- **解析**: `goquery` (HTML), `gjson` (JSON)
- **JS引擎**: `goja` (解密加密数据)
- **数据结构**: `gota/dataframe`

## 项目结构

```
akshare-go/
├── internal/               # 内部工具包
│   ├── httpclient/         # HTTP 客户端
│   ├── parser/             # 数据解析器
│   └── types/              # 公共类型
├── stock/                  # 股票模块 (128 函数)
├── stock_feature/          # 股票特色 (197 函数)
├── fund/                   # 基金模块 (82 函数)
├── futures/                # 期货模块 (85 函数)
├── bond/                   # 债券模块 (40 函数)
├── index/                  # 指数模块 (95 函数)
├── option/                 # 期权模块 (50 函数)
├── economic/               # 宏观经济 (215 函数)
└── plans/                  # 计划文档
```

## 计划文档

- **主索引**: `plans/plan_akshare-go.md`
- **模块计划**: `plans/modules/*.md`
- **数据源**: `plans/modules/datasources.md`
- **架构设计**: `plans/modules/architecture.md`

## 代码规范

1. **命名**: camelCase, 导出用 PascalCase
2. **注释**: 全部使用中文
3. **文件**: 每个文件不超过 500 行
4. **函数**: 每个函数不超过 50 行
5. **错误**: 使用 `fmt.Errorf` 包装，提供有意义信息

## 常用命令

```bash
# 编译
go build ./...

# 测试
go test ./...

# 格式化
go fmt ./...

# 静态检查
go vet ./...
```

## 工作流

0. 遵循 A-Z 的方式,逐个目录参考与实现
1. **规划**: Opus 负责 (读取计划、分解任务)
2. **编码**: Sonnet 子代理 (使用 go-coder 代理)
3. **测试**: Sonnet 子代理 (使用 go-tester 代理)
4. **审查**: Opus 负责 (使用 code-reviewer 代理)

## 数据源

参考原版 Python 实现: `_akshare_source/akshare/`

## 编码规范

1. **测试**: 必须每个接口都必须有测试用例
2. **数据源**: 严格参考原版 Python 实现

主要数据源:

- 东方财富 (eastmoney.com)
- 新浪财经 (sina.com.cn)
- 同花顺 (10jqka.com.cn)
- 交易所官网

## 优先级

- **P0**: 核心接口，必须实现
- **P1**: 重要接口，第二批
- **P2**: 一般接口，第三批
- **P3**: 可选接口，最后实现

# 完整模块清单（按字母顺序）

✅ air/ (9) 🔜 article/ (7) 🔜 bank/ (4)
🔜 bond/ (42) 🔜 cal/ (3) 🔜 crypto/ (2)
🔜 currency/ (7) 🔜 economic/ (230) 🔜 energy/ (8)
🔜 event/ (2) 🔜 forex/ (2) 🔜 fortune/ (6)
🔜 fund/ (84) 🔜 futures/ (85) 🔜 futures_derivative/ (15)
🔜 fx/ (6) 🔜 hf/ (1) 🔜 interest_rate/ (1)
🔜 movie/ (16) 🔜 news/ (6) 🔜 nlp/ (2)
🔜 option/ (52) 🔜 pro/ (1) 🔜 qdii/ (3)
🔜 qhkc_web/ (10) 🔜 rate/ (2) 🔜 reits/ (3)
🔜 spot/ (15) 🔜 stock_a/ (6) 🔜 stock_feature/ (207)
🔜 stock_fundamental/ (56) 🔜 tool/ (1)

严格遵循akshare的设计, 文件名, 函数, 函数参数, 函数返回值都要保持一致. 并且最后要写单元测试,覆盖每一个函数.
有两组开发者同时实现, 一组是正序, 一组是逆序.
