# Futures 模块实现

> 创建: 2026-01-19
> 状态: 进行中

## 目标

实现 futures 模块 28 个 Python 文件的 Go 版本

## Python 文件清单 (按字母顺序)

| # | 文件 | 大小 | 状态 |
|---|------|------|------|
| 1 | cons.py | 18KB | 待实现 |
| 2 | cot.py | 58KB | 待实现 |
| 3 | futures_basis.py | 16KB | 待实现 |
| 4 | futures_comex_em.py | 3KB | 待实现 |
| 5 | futures_comm_ctp.py | 1KB | 待实现 |
| 6 | futures_comm_qihuo.py | 12KB | 待实现 |
| 7 | futures_contract_detail.py | 3KB | 待实现 |
| 8 | futures_daily_bar.py | 25KB | 待实现 |
| 9 | futures_foreign.py | 2KB | 待实现 |
| 10 | futures_hf_em.py | 8KB | 待实现 |
| 11 | futures_hist_em.py | 7KB | 待实现 |
| 12 | futures_hq_sina.py | 10KB | 待实现 |
| 13 | futures_index_ccidx.py | 2KB | 待实现 |
| 14 | futures_inventory_99.py | 3KB | 待实现 |
| 15 | futures_inventory_em.py | 3KB | 待实现 |
| 16 | futures_news_shmet.py | 2KB | 待实现 |
| 17 | futures_roll_yield.py | 6KB | 待实现 |
| 18 | futures_rule_em.py | 1KB | 待实现 |
| 19 | futures_rule.py | 2KB | 待实现 |
| 20 | futures_settlement_price_sgx.py | 3KB | 待实现 |
| 21 | futures_spot_stock_em.py | 4KB | 待实现 |
| 22 | futures_stock_js.py | 2KB | 待实现 |
| 23 | futures_to_spot.py | 13KB | 待实现 |
| 24 | futures_warehouse_receipt.py | 9KB | 待实现 |
| 25 | futures_zh_sina.py | 26KB | 待实现 |
| 26 | receipt.py | 24KB | 待实现 |
| 27 | requests_fun.py | 3KB | 待实现 |
| 28 | symbol_var.py | 6KB | 待实现 |

## 实现策略

1. 先实现基础工具函数 (cons.py, symbol_var.py, requests_fun.py)
2. 再实现核心数据接口
3. 最后实现复杂功能

## 完成标准

- [ ] 所有 28 个文件实现
- [ ] 单元测试覆盖
- [ ] 编译通过
- [ ] STATUS.md 更新
