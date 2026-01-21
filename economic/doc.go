// Package economic 提供宏观经济数据接口
//
// 本包实现了以下数据源的宏观经济相关数据获取：
//   - 金十数据中心：全球主要国家和地区的宏观经济指标
//   - 国家统计局：中国官方宏观经济数据
//   - 东方财富：部分中国宏观经济指标
//   - 同花顺：财经宏观数据
//
// 主要功能模块：
//   - 央行利率：美联储、欧洲央行、中国央行等主要央行利率决议
//   - 美国宏观数据：GDP、CPI、PPI、失业率等49个指标
//   - 欧元区宏观数据：欧元区主要经济指标
//   - 英国宏观数据：英国主要经济指标
//   - 中国宏观数据：GDP、CPI、PPI、进出口、PMI等78个指标
//   - 其他国家：加拿大、澳大利亚、日本、德国、瑞士等国家数据
//   - 国家统计局：中国官方统计数据
//
// 数据来源：
//   - 金十数据中心 (https://datacenter.jin10.com/)
//   - 国家统计局 (https://data.stats.gov.cn/)
//   - 东方财富 (https://www.eastmoney.com/)
//   - 同花顺 (https://www.10jqka.com.cn/)
//
// 对应 Python akshare 源码：
//   - akshare/economic/macro_bank.py
//   - akshare/economic/macro_usa.py
//   - akshare/economic/macro_euro.py
//   - akshare/economic/macro_china.py
//   - akshare/economic/macro_china_nbs.py
//   - 其他 macro_*.py 文件
package economic
