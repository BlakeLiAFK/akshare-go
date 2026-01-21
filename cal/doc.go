// Package cal 提供已实现波动率计算相关接口
//
// 主要功能：
//   - Yang-Zhang 已实现波动率计算 (volatility_yz_rv)
//   - 从股票分钟数据计算波动率 (rv_from_stock_zh_a_hist_min_em)
//   - 从期货分钟数据计算波动率 (rv_from_futures_zh_minute_sina)
//
// 理论基础：
//   - Yang-Zhang 已实现波动率论文: https://www.jstor.org/stable/10.1086/209650
//   - 参考实现: https://github.com/hugogobato/Yang-Zhang-s-Realized-Volatility-Automated-Estimation-in-Python
//
// 数据源：
//   - 东方财富网 (股票分钟数据): https://quote.eastmoney.com
//   - 新浪财经 (期货分钟数据): https://vip.stock.finance.sina.com.cn
//
// 依赖说明：
//   - rv_from_stock_zh_a_hist_min_em 依赖 stock_feature 模块
//   - rv_from_futures_zh_minute_sina 依赖 futures 模块
//   - 上述依赖模块将在后续按字母顺序实现时补充
package cal
