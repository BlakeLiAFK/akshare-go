# akshare-go 工作流设计

> 文档: Claude Code 工作流配置
> 创建日期: 2026-01-18

---

## 一、工作流概述

### 1.1 核心理念

```
┌─────────────────────────────────────────────────────────────┐
│                      Opus (规划/审查)                        │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐          │
│  │ 读取计划    │→ │ 分解任务    │→ │ 下发命令    │          │
│  └─────────────┘  └─────────────┘  └─────────────┘          │
│                            ↓                                 │
│  ┌───────────────────────────────────────────────────────┐  │
│  │                  Sonnet 子代理                         │  │
│  │  ┌─────────┐  ┌─────────┐  ┌─────────┐  ┌─────────┐  │  │
│  │  │go-coder │  │go-tester│  │go-coder │  │go-tester│  │  │
│  │  └────┬────┘  └────┬────┘  └────┬────┘  └────┬────┘  │  │
│  │       ↓            ↓            ↓            ↓        │  │
│  │  [实现函数1]   [测试函数1]  [实现函数2]   [测试函数2]  │  │
│  └───────────────────────────────────────────────────────┘  │
│                            ↓                                 │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐          │
│  │ 验证结果    │← │ code-review │← │ 汇总报告    │          │
│  └─────────────┘  └─────────────┘  └─────────────┘          │
└─────────────────────────────────────────────────────────────┘
```

### 1.2 模型分工

| 角色 | 模型 | 职责 |
|------|------|------|
| 主控 | Opus | 规划、命令下发、验收审查 |
| 编码 | Sonnet | 编写 Go 代码 |
| 测试 | Sonnet | 编写测试用例 |
| 审查 | Opus | 代码质量审查 |

### 1.3 Token 节省策略

- 使用 `opusplan` 模式: 规划用 Opus，执行自动切 Sonnet
- 设置 `CLAUDE_CODE_SUBAGENT_MODEL=sonnet`
- 大量代码工作通过子代理完成

---

## 二、目录结构

```
.claude/
├── settings.json          # 主配置 (hooks, 权限, 模型)
├── agents/
│   ├── go-coder.md        # 代码编写代理 (sonnet)
│   ├── go-tester.md       # 测试编写代理 (sonnet)
│   └── code-reviewer.md   # 代码审查代理 (opus)
├── commands/
│   ├── implement.md       # /implement <module>
│   ├── review.md          # /review [path]
│   ├── batch.md           # /batch <module> <count>
│   └── init-module.md     # /init-module <module>
├── hooks/
│   ├── check-completion.sh   # Stop hook - 检查任务完成
│   ├── check-subagent.sh     # SubagentStop hook
│   └── post-edit.sh          # 编辑后自动格式化
└── skills/
    ├── go-patterns.md     # Go 最佳实践
    └── data-sources.md    # 数据源参考
```

---

## 三、Hooks 机制

### 3.1 Stop Hook (持续工作)

当 Claude 准备停止时触发，检查:
1. 是否有未完成的 TODO
2. 代码是否编译通过
3. 测试是否通过

如果有未完成项，返回 `{"decision": "block", "reason": "..."}` 强制继续。

### 3.2 SubagentStop Hook

子代理完成时触发，检查:
1. 代码格式是否规范
2. 是否有明显错误

### 3.3 PostToolUse Hook

编辑 Go 文件后自动:
1. 运行 `gofmt` 格式化
2. 运行 `goimports` 整理导入

---

## 四、工作流命令

### 4.1 /implement <module>

实现指定模块的所有函数。

```bash
/implement stock
```

流程:
1. Opus 读取 `plans/modules/stock.md`
2. 按优先级排序函数列表
3. 调用 Sonnet 子代理逐个实现
4. 每个函数完成后验证编译和测试
5. 全部完成后生成报告

### 4.2 /batch <module> <count>

批量实现指定数量的函数 (持续工作模式)。

```bash
/batch stock 10
```

特点:
- Stop hook 确保完成所有函数才停止
- 遇到困难可跳过，继续下一个
- 定期保存进度

### 4.3 /review [path]

代码质量审查。

```bash
/review stock/
```

检查:
- 静态分析 (`go vet`)
- 格式规范 (`gofmt`)
- 测试覆盖率
- 代码规范

### 4.4 /init-module <module>

初始化模块目录结构。

```bash
/init-module stock
```

创建:
- `doc.go` - 包文档
- `types.go` - 类型定义
- `client.go` - HTTP 客户端
- 主文件

---

## 五、使用示例

### 5.1 开始新模块

```bash
# 1. 初始化模块
/init-module stock

# 2. 实现模块
/implement stock

# 3. 审查代码
/review stock/
```

### 5.2 批量实现

```bash
# 实现 stock 模块的 20 个函数
/batch stock 20
```

Claude 会:
1. 读取计划文档
2. 选择 20 个高优先级函数
3. 循环调用子代理实现
4. Stop hook 确保全部完成

### 5.3 手动触发子代理

```
请使用 go-coder 代理实现 stock_zh_a_spot_em 函数
```

---

## 六、配置说明

### 6.1 settings.json 关键配置

```json
{
  "model": "opusplan",        // Opus规划, Sonnet执行
  "env": {
    "CLAUDE_CODE_SUBAGENT_MODEL": "sonnet"  // 子代理用 Sonnet
  },
  "hooks": {
    "Stop": [...],            // 检查任务完成
    "SubagentStop": [...],    // 检查子代理完成
    "PostToolUse": [...]      // 自动格式化
  }
}
```

### 6.2 代理模型配置

```yaml
# go-coder.md
model: sonnet    # 编码用 Sonnet (省 token)

# code-reviewer.md
model: opus      # 审查用 Opus (高质量)
```

---

## 七、注意事项

1. **防止无限循环**: Stop hook 检查 `stop_hook_active` 字段
2. **错误恢复**: 遇到编译错误时自动修复
3. **进度保存**: 每完成一个函数就更新 TODO
4. **代码质量**: 子代理完成后由 Opus 审查

---

*文档版本: v1.0*
*最后更新: 2026-01-18*
