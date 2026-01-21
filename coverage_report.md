# 代码覆盖率报告

生成时间: 2026-01-21
测试命令: `go test -coverprofile=coverage.out -covermode=count ./...`

---

## 总体覆盖率

**24.1%** (总计所有模块的语句覆盖率)

---

## 模块覆盖率详情

### 🟢 高覆盖率模块 (>= 80%)

| 模块 | 覆盖率 | 状态 |
|------|--------|------|
| cal | 89.0% | ✅ PASS |
| crypto | 86.2% | ✅ PASS |

### 🟡 中等覆盖率模块 (40% - 80%)

| 模块 | 覆盖率 | 状态 |
|------|--------|------|
| currency | 51.3% | ✅ PASS |

### 🔴 低覆盖率模块 (< 40%)

| 模块 | 覆盖率 | 状态 | 备注 |
|------|--------|------|------|
| bond | 26.2% | ❌ FAIL | 3个测试失败 |
| air | 15.5% | ✅ PASS | - |
| article | 19.6% | ✅ PASS | - |
| bank | 18.9% | ✅ PASS | - |

### 📋 无测试文件模块

| 模块 | 状态 |
|------|------|
| akshare-go (根目录) | 无测试文件 |
| data | 无测试文件 |

---

## 测试失败模块分析

### bond 模块 (26.2% 覆盖率, FAIL)

**失败测试**:
- `TestBondSpotQuote` - 未获取到有效数据 (0行0列)
- `TestBondSpotDeal` - 未获取到有效数据 (0行0列)
- `TestBondChinaCloseReturn` - 未获取到有效数据 (0行0列)

**原因**: API 返回空数据,可能是网络问题或 API 变更

**建议**:
- 检查 API 是否正常可访问
- 添加重试机制
- 或标记为 `t.Skip()` 跳过不稳定的网络测试

---

## 覆盖率分布

```
89%-86%  ██ 2个模块 (cal, crypto)
51%      █  1个模块 (currency)
26%-19%  ████ 4个模块 (bond, air, article, bank)
无测试   ██ 2个模块 (根目录, data)
```

---

## 改进建议

### 优先级 P0 (必须)
1. **提升核心模块覆盖率** - air, article, bank, bond 覆盖率低于 30%
   - 为现有函数添加更多测试用例
   - 覆盖边界条件和错误处理

2. **修复失败测试** - bond 模块的 3 个失败测试
   - 调查 API 数据源问题
   - 添加 mock 数据作为备选

### 优先级 P1 (重要)
3. **添加缺失的测试文件**
   - 根目录需要集成测试
   - data 模块需要单元测试

4. **提高测试质量**
   - 当前很多测试只验证非空,应该验证数据正确性
   - 添加表驱动测试 (table-driven tests)
   - 增加边界条件测试

### 优先级 P2 (建议)
5. **网络测试优化**
   - 为网络依赖的测试添加 `testing.Short()` 支持
   - 使用 `-short` flag 跳过慢速网络测试
   - 考虑添加 mock 服务器

6. **CI/CD 集成**
   - 设置覆盖率阈值 (建议 >= 60%)
   - 自动生成覆盖率徽章
   - PR 中显示覆盖率变化

---

## 文件清单

- `coverage.out` - 覆盖率原始数据
- `coverage.html` - HTML 可视化报告 (使用 `go tool cover -html=coverage.out` 生成)
- `coverage_run.log` - 完整测试输出日志
- `coverage_report.md` - 本报告

---

## 查看详细覆盖率

### HTML 报告
```bash
# 在浏览器中打开可视化报告
open coverage.html
```

### 命令行查看
```bash
# 查看总体覆盖率
go tool cover -func=coverage.out | grep total

# 查看特定模块覆盖率
go tool cover -func=coverage.out | grep bond

# 查看特定文件覆盖率
go tool cover -func=coverage.out | grep bond_jsl.go
```

---

## 结论

项目当前总体覆盖率为 **24.1%**,仍有较大提升空间。主要问题:

1. ✅ **优势**: cal 和 crypto 模块覆盖率良好 (>85%)
2. ⚠️ **警告**: 大部分模块覆盖率偏低 (<30%)
3. ❌ **问题**: bond 模块存在测试失败,需要修复

**下一步行动**:
1. 优先修复 bond 模块的 3 个失败测试
2. 为 air, article, bank 模块补充测试用例
3. 目标: 将总体覆盖率提升至 **60%** 以上
