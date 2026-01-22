# option 模块

> 本模块共有 39 个接口

## 目录

- [ConvertDate](#convertdate)
- [GetCalendar](#getcalendar)
- [GetCzceOptionSymbols](#getczceoptionsymbols)
- [GetDceOptionSymbols](#getdceoptionsymbols)
- [GetFinanceOptionSymbols](#getfinanceoptionsymbols)
- [GetJSONPath](#getjsonpath)
- [GetLatestDataDate](#getlatestdatadate)
- [GetSinaFinanceOptionSymbols](#getsinafinanceoptionsymbols)
- [GetSinaOptionSymbols](#getsinaoptionsymbols)
- [LastTradingDay](#lasttradingday)
- [OptionCommInfo](#optioncomminfo)
- [OptionCommSymbol](#optioncommsymbol)
- [OptionCommoditySina](#optioncommoditysina)
- [OptionCommoditySinaAll](#optioncommoditysinaall)
- [OptionContractInfoCtp](#optioncontractinfoctp)
- [OptionCurrentCffexEm](#optioncurrentcffexem)
- [OptionCurrentDaySse](#optioncurrentdaysse)
- [OptionCurrentDaySzse](#optioncurrentdayszse)
- [OptionCurrentEm](#optioncurrentem)
- [OptionDailyStatsSse](#optiondailystatssse)
- [OptionDailyStatsSzse](#optiondailystatsszse)
- [OptionFinance](#optionfinance)
- [OptionFinanceDaily](#optionfinancedaily)
- [OptionFinanceSSE](#optionfinancesse)
- [OptionFinanceSZSE](#optionfinanceszse)
- [OptionFinanceSina](#optionfinancesina)
- [OptionFinanceSinaChain](#optionfinancesinachain)
- [OptionFinanceSinaDetail](#optionfinancesinadetail)
- [OptionFinanceSinaVolatility](#optionfinancesinavolatility)
- [OptionHistDce](#optionhistdce)
- [OptionHistShfe](#optionhistshfe)
- [OptionHistYearlyCzce](#optionhistyearlyczce)
- [OptionLhbEm](#optionlhbem)
- [OptionMargin](#optionmargin)
- [OptionMarginSymbol](#optionmarginsymbol)
- [OptionPremiumAnalysisEm](#optionpremiumanalysisem)
- [OptionRiskAnalysisEm](#optionriskanalysisem)
- [OptionRiskIndicatorSse](#optionriskindicatorsse)
- [OptionValueAnalysisEm](#optionvalueanalysisem)

---

## ConvertDate

**描述**: ConvertDate 转换日期字符串为time.Time对象

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| dateStr | string | 日期，格式：YYYYMMDD |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/option"
)

func main() {
	data, err := option.ConvertDate("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## GetCalendar

**描述**: GetCalendar 获取交易日历

### 输入参数

无参数

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/option"
)

func main() {
	data, err := option.GetCalendar()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## GetCzceOptionSymbols

**描述**: GetCzceOptionSymbols 获取郑商所期权品种列表

### 输入参数

无参数

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/option"
)

func main() {
	data, err := option.GetCzceOptionSymbols()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## GetDceOptionSymbols

**描述**: GetDceOptionSymbols 获取大连商品期权品种列表

### 输入参数

无参数

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/option"
)

func main() {
	data, err := option.GetDceOptionSymbols()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## GetFinanceOptionSymbols

**描述**: GetFinanceOptionSymbols 获取金融期权品种列表

### 输入参数

无参数

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/option"
)

func main() {
	data, err := option.GetFinanceOptionSymbols()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## GetJSONPath

**描述**: GetJSONPath 获取JSON配置文件的路径

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| name | string | - |
| moduleFile | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/option"
)

func main() {
	data, err := option.GetJSONPath("", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## GetLatestDataDate

**描述**: GetLatestDataDate 获取最新的有数据的交易日

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| day | time.Time | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/option"
)

func main() {
	data, err := option.GetLatestDataDate(nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## GetSinaFinanceOptionSymbols

**描述**: GetSinaFinanceOptionSymbols 获取新浪金融期权品种列表

### 输入参数

无参数

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/option"
)

func main() {
	data, err := option.GetSinaFinanceOptionSymbols()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## GetSinaOptionSymbols

**描述**: GetSinaOptionSymbols 获取新浪期权品种列表

### 输入参数

无参数

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/option"
)

func main() {
	data, err := option.GetSinaOptionSymbols()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## LastTradingDay

**描述**: LastTradingDay 获取前一个交易日

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| day | interface{} | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/option"
)

func main() {
	data, err := option.LastTradingDay(nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## OptionCommInfo

**描述**: OptionCommInfo 获取商品期权手续费信息

**数据源**: https://www.9qihuo.com/qiquanshouxufei

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/option"
)

func main() {
	data, err := option.OptionCommInfo("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## OptionCommSymbol

**描述**: OptionCommSymbol 获取商品期权品种列表

**数据源**: https://www.9qihuo.com/qiquanshouxufei

### 输入参数

无参数

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/option"
)

func main() {
	data, err := option.OptionCommSymbol()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## OptionCommoditySina

**描述**: OptionCommoditySina 新浪商品期权-实时行情

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/option"
)

func main() {
	data, err := option.OptionCommoditySina("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## OptionCommoditySinaAll

**描述**: OptionCommoditySinaAll 新浪商品期权-所有品种实时行情

### 输入参数

无参数

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/option"
)

func main() {
	data, err := option.OptionCommoditySinaAll()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## OptionContractInfoCtp

**描述**: OptionContractInfoCtp openctp-合约信息接口-期权合约

### 输入参数

无参数

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/option"
)

func main() {
	data, err := option.OptionContractInfoCtp()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## OptionCurrentCffexEm

**描述**: OptionCurrentCffexEm 东方财富网-中金所期权行情

### 输入参数

无参数

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/option"
)

func main() {
	data, err := option.OptionCurrentCffexEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## OptionCurrentDaySse

**描述**: OptionCurrentDaySse 上海证券交易所-产品-股票期权-信息披露-当日合约

**数据源**: http://www.sse.com.cn/assortment/options/disclo/preinfo/

### 输入参数

无参数

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/option"
)

func main() {
	data, err := option.OptionCurrentDaySse()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## OptionCurrentDaySzse

**描述**: OptionCurrentDaySzse 深圳证券交易所-期权子网-行情数据-当日合约

### 输入参数

无参数

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/option"
)

func main() {
	data, err := option.OptionCurrentDaySzse()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## OptionCurrentEm

**描述**: OptionCurrentEm 东方财富网-行情中心-期权市场

**数据源**: https://quote.eastmoney.com/center/qqsc.html

### 输入参数

无参数

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/option"
)

func main() {
	data, err := option.OptionCurrentEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## OptionDailyStatsSse

**描述**: OptionDailyStatsSse 上海证券交易所-产品-股票期权-每日统计

**数据源**: https://www.sse.com.cn/assortment/options/date/

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| date | string | 日期，格式：YYYYMMDD |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/option"
)

func main() {
	data, err := option.OptionDailyStatsSse("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## OptionDailyStatsSzse

**描述**: OptionDailyStatsSzse 深圳证券交易所-市场数据-期权数据-日度概况

**数据源**: https://investor.szse.cn/market/option/day/index.html

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| date | string | 日期，格式：YYYYMMDD |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/option"
)

func main() {
	data, err := option.OptionDailyStatsSzse("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## OptionFinance

**描述**: OptionFinance 金融期权-实时行情数据

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| exchange | string | - |
| symbol | string | 股票/基金代码 |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/option"
)

func main() {
	data, err := option.OptionFinance("", "000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## OptionFinanceDaily

**描述**: OptionFinanceDaily 金融期权-日频历史数据

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |
| startDate | string | 开始日期，格式：YYYYMMDD |
| endDate | string | 结束日期，格式：YYYYMMDD |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/option"
)

func main() {
	data, err := option.OptionFinanceDaily("000001", "20230101", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## OptionFinanceSSE

**描述**: OptionFinanceSSE 上交所金融期权-实时行情

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/option"
)

func main() {
	data, err := option.OptionFinanceSSE("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## OptionFinanceSZSE

**描述**: OptionFinanceSZSE 深交所金融期权-实时行情

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/option"
)

func main() {
	data, err := option.OptionFinanceSZSE("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## OptionFinanceSina

**描述**: OptionFinanceSina 新浪金融期权-实时行情

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/option"
)

func main() {
	data, err := option.OptionFinanceSina("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## OptionFinanceSinaChain

**描述**: OptionFinanceSinaChain 新浪金融期权-期权链数据

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |
| date | string | 日期，格式：YYYYMMDD |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/option"
)

func main() {
	data, err := option.OptionFinanceSinaChain("000001", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## OptionFinanceSinaDetail

**描述**: OptionFinanceSinaDetail 新浪金融期权-详细行情数据

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |
| contract | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/option"
)

func main() {
	data, err := option.OptionFinanceSinaDetail("000001", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## OptionFinanceSinaVolatility

**描述**: OptionFinanceSinaVolatility 新浪金融期权-波动率数据

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/option"
)

func main() {
	data, err := option.OptionFinanceSinaVolatility("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## OptionHistDce

**描述**: OptionHistDce 大连商品交易所-期权-日频行情数据

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |
| tradeDate | string | 日期，格式：YYYYMMDD |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/option"
)

func main() {
	data, err := option.OptionHistDce("000001", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## OptionHistShfe

**描述**: OptionHistShfe 上海期货交易所-期权-日频行情数据

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |
| tradeDate | string | 日期，格式：YYYYMMDD |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/option"
)

func main() {
	data, err := option.OptionHistShfe("000001", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## OptionHistYearlyCzce

**描述**: OptionHistYearlyCzce 郑州商品交易所-交易数据-历史行情下载-期权历史行情下载

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |
| year | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/option"
)

func main() {
	data, err := option.OptionHistYearlyCzce("000001", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## OptionLhbEm

**描述**: OptionLhbEm 东方财富网-数据中心-期货期权-期权龙虎榜单

**数据源**: https://data.eastmoney.com/other/qqlhb.html

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |
| indicator | string | 指标类型 |
| tradeDate | string | 日期，格式：YYYYMMDD |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/option"
)

func main() {
	data, err := option.OptionLhbEm("000001", "", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## OptionMargin

**描述**: OptionMargin 获取商品期权保证金

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/option"
)

func main() {
	data, err := option.OptionMargin("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## OptionMarginSymbol

**描述**: OptionMarginSymbol 获取商品期权品种代码和名称

### 输入参数

无参数

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/option"
)

func main() {
	data, err := option.OptionMarginSymbol()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## OptionPremiumAnalysisEm

**描述**: OptionPremiumAnalysisEm 东方财富网-数据中心-特色数据-期权折溢价

**数据源**: https://data.eastmoney.com/other/premium.html

### 输入参数

无参数

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/option"
)

func main() {
	data, err := option.OptionPremiumAnalysisEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## OptionRiskAnalysisEm

**描述**: OptionRiskAnalysisEm 东方财富网-数据中心-特色数据-期权风险分析

**数据源**: https://data.eastmoney.com/other/riskanal.html

### 输入参数

无参数

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/option"
)

func main() {
	data, err := option.OptionRiskAnalysisEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## OptionRiskIndicatorSse

**描述**: OptionRiskIndicatorSse 上海证券交易所-产品-股票期权-期权风险指标

**数据源**: http://www.sse.com.cn/assortment/options/risk/

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| date | string | 日期，格式：YYYYMMDD |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/option"
)

func main() {
	data, err := option.OptionRiskIndicatorSse("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## OptionValueAnalysisEm

**描述**: OptionValueAnalysisEm 东方财富网-数据中心-特色数据-期权价值分析

**数据源**: https://data.eastmoney.com/other/valueAnal.html

### 输入参数

无参数

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/option"
)

func main() {
	data, err := option.OptionValueAnalysisEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

