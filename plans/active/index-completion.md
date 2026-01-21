# Index 模块补全

> 创建: 2026-01-18
> 状态: 进行中

## 目标

将 index 模块从 42% 完成度提升到 100%

## 现状分析

### Python 模块 (24个)
| 模块 | Go 状态 | 说明 |
|------|---------|------|
| index_cflp | ❌ | 中物联钢铁物流专业委员会指数 |
| index_cni | ✅ | 国证指数 |
| index_csindex | ✅ | 中证指数 |
| index_cx | ❌ | 财新指数 |
| index_drewry | ❌ | 德路里航运指数 |
| index_eri | ❌ | 义乌小商品指数 |
| index_global_em | ✅ | 全球指数(东财) |
| index_global_sina | ✅ | 全球指数(新浪) |
| index_hog | ❌ | 生猪数据指数 |
| index_kq_fz | ❌ | 柯桥纺织指数 |
| index_kq_ss | ❌ | 柯桥时尚指数 |
| index_option_qvix | ❌ | 期权波动率指数 |
| index_research_fund_sw | ❌ | 申万基金研究 |
| index_research_sw | ❌ | 申万行业研究 |
| index_spot | ✅ | 现货指数 |
| index_stock_hk | ✅ | 港股指数 |
| index_stock_us_sina | ✅ | 美股指数(新浪) |
| index_stock_zh | ❌ | A股指数行情 |
| index_stock_zh_csindex | ❌ | A股指数(中证) |
| index_sugar | ❌ | 糖果指数 |
| index_sw | ✅ | 申万指数(部分) |
| index_yw | ❌ | 义乌指数 |
| index_zh_a_scope | ❌ | A股指数范围 |
| index_zh_em | ❌ | A股指数(东财) |

### 待实现 (15个)
1. index_cflp
2. index_cx
3. index_drewry
4. index_eri
5. index_hog
6. index_kq_fz
7. index_kq_ss
8. index_option_qvix
9. index_research_fund_sw
10. index_research_sw
11. index_stock_zh
12. index_stock_zh_csindex
13. index_sugar
14. index_yw
15. index_zh_a_scope

## 实现步骤

- [ ] 实现 index_cflp.go
- [ ] 实现 index_cx.go
- [ ] 实现 index_drewry.go
- [ ] 实现 index_eri.go
- [ ] 实现 index_hog.go
- [ ] 实现 index_kq_fz.go
- [ ] 实现 index_kq_ss.go
- [ ] 实现 index_option_qvix.go
- [ ] 实现 index_research_fund_sw.go
- [ ] 实现 index_research_sw.go
- [ ] 实现 index_stock_zh.go
- [ ] 实现 index_stock_zh_csindex.go
- [ ] 实现 index_sugar.go
- [ ] 实现 index_yw.go
- [ ] 实现 index_zh_a_scope.go
- [ ] 补充单元测试
- [ ] 更新 STATUS.md

## 完成标准

- [ ] 所有 15 个缺失模块已实现
- [ ] 每个函数都有对应的单元测试
- [ ] go test ./index/... 全部通过
- [ ] STATUS.md 更新为 100%
