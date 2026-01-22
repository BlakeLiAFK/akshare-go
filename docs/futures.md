# futures 模块

> 本模块共有 53 个接口

## 目录

- [ChineseToEnglish](#chinesetoenglish)
- [FindChinese](#findchinese)
- [FuturesCFFEXRankTable](#futurescffexranktable)
- [FuturesCZCERankTable](#futuresczceranktable)
- [FuturesComexInventoryFunc](#futurescomexinventoryfunc)
- [FuturesContractDetailEMFunc](#futurescontractdetailemfunc)
- [FuturesContractDetailFunc](#futurescontractdetailfunc)
- [FuturesDCEPositionRank](#futuresdcepositionrank)
- [FuturesDeliveryCZCE](#futuresdeliveryczce)
- [FuturesDeliveryDCE](#futuresdeliverydce)
- [FuturesDeliverySHFE](#futuresdeliveryshfe)
- [FuturesForeignCommodityRealtimeFunc](#futuresforeigncommodityrealtimefunc)
- [FuturesForeignCommoditySubscribeExchangeSymbol](#futuresforeigncommoditysubscribeexchangesymbol)
- [FuturesForeignHistEM](#futuresforeignhistem)
- [FuturesGFEXPositionRank](#futuresgfexpositionrank)
- [FuturesGFEXWarehouseReceipt](#futuresgfexwarehousereceipt)
- [FuturesHFMinuteEMFunc](#futureshfminuteemfunc)
- [FuturesHFSpotEMFunc](#futureshfspotemfunc)
- [FuturesHQSubscribeExchangeSymbol](#futureshqsubscribeexchangesymbol)
- [FuturesHistEMFunc](#futureshistemfunc)
- [FuturesHistTableEMFunc](#futureshisttableemfunc)
- [FuturesIndexCCIDX](#futuresindexccidx)
- [FuturesInventoryEMFunc](#futuresinventoryemfunc)
- [FuturesNewsSHMET](#futuresnewsshmet)
- [FuturesRankSum](#futuresranksum)
- [FuturesRuleEMFunc](#futuresruleemfunc)
- [FuturesRuleFunc](#futuresrulefunc)
- [FuturesSHFERankTable](#futuresshferanktable)
- [FuturesSHFEWarehouseReceipt](#futuresshfewarehousereceipt)
- [FuturesSpotPrice](#futuresspotprice)
- [FuturesSpotPriceDaily](#futuresspotpricedaily)
- [FuturesSpotPricePrevious](#futuresspotpriceprevious)
- [FuturesSymbolMarkFunc](#futuressymbolmarkfunc)
- [FuturesToSpotCZCE](#futurestospotczce)
- [FuturesToSpotDCE](#futurestospotdce)
- [FuturesToSpotSHFE](#futurestospotshfe)
- [FuturesWarehouseReceiptCZCE](#futureswarehousereceiptczce)
- [FuturesWarehouseReceiptDCE](#futureswarehousereceiptdce)
- [FuturesZhDailySina](#futureszhdailysina)
- [FuturesZhMinuteSina](#futureszhminutesina)
- [FuturesZhRealtimeFunc](#futureszhrealtimefunc)
- [FuturesZhSpotFunc](#futureszhspotfunc)
- [GetCFFEXDaily](#getcffexdaily)
- [GetCZCEDaily](#getczcedaily)
- [GetDCEDaily](#getdcedaily)
- [GetFuturesDaily](#getfuturesdaily)
- [GetGFEXDaily](#getgfexdaily)
- [GetINEDaily](#getinedaily)
- [GetRollYield](#getrollyield)
- [GetRollYieldBar](#getrollyieldbar)
- [GetSHFEDaily](#getshfedaily)
- [SymbolMarket](#symbolmarket)
- [SymbolVarieties](#symbolvarieties)

---

## ChineseToEnglish

**描述**: ChineseToEnglish 将期货品种中文名称映射为英文缩写

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| chineseVar | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.ChineseToEnglish("")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FindChinese

**描述**: FindChinese 提取字符串中的中文字符

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| s | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.FindChinese("")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesCFFEXRankTable

**描述**: FuturesCFFEXRankTable 中国金融期货交易所前20会员持仓排名数据

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| date | string | 日期，格式：YYYYMMDD |
| varsList | []string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.FuturesCFFEXRankTable("20230101", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesCZCERankTable

**描述**: FuturesCZCERankTable 郑州商品交易所前20会员持仓排名数据

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
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.FuturesCZCERankTable("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesComexInventoryFunc

**描述**: FuturesComexInventoryFunc 东方财富-COMEX库存数据

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
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.FuturesComexInventoryFunc("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesContractDetailEMFunc

**描述**: FuturesContractDetailEMFunc 查询期货合约详情-东方财富

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
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.FuturesContractDetailEMFunc("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesContractDetailFunc

**描述**: FuturesContractDetailFunc 查询期货合约详情-新浪

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
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.FuturesContractDetailFunc("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesDCEPositionRank

**描述**: FuturesDCEPositionRank 大连商品交易所-每日持仓排名-具体合约

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| date | string | 日期，格式：YYYYMMDD |
| varsList | []string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.FuturesDCEPositionRank("20230101", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesDeliveryCZCE

**描述**: FuturesDeliveryCZCE 郑州商品交易所-月度交割查询

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
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.FuturesDeliveryCZCE("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesDeliveryDCE

**描述**: FuturesDeliveryDCE 大连商品交易所-交割统计

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
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.FuturesDeliveryDCE("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesDeliverySHFE

**描述**: FuturesDeliverySHFE 上海期货交易所-交割情况表

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
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.FuturesDeliverySHFE("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesForeignCommodityRealtimeFunc

**描述**: FuturesForeignCommodityRealtime 获取外盘期货实时行情

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
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.FuturesForeignCommodityRealtimeFunc("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesForeignCommoditySubscribeExchangeSymbol

**描述**: FuturesForeignCommoditySubscribeExchangeSymbol 获取需要订阅的外盘期货代码列表

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
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.FuturesForeignCommoditySubscribeExchangeSymbol()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesForeignHistEM

**描述**: FuturesForeignHistEM 东方财富-外盘期货历史数据

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |
| period | string | 周期：daily/weekly/monthly |
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
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.FuturesForeignHistEM("000001", "daily", "20230101", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesGFEXPositionRank

**描述**: FuturesGFEXPositionRank 广州期货交易所-日成交持仓排名

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| date | string | 日期，格式：YYYYMMDD |
| varsList | []string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.FuturesGFEXPositionRank("20230101", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesGFEXWarehouseReceipt

**描述**: FuturesGFEXWarehouseReceipt 广州期货交易所-行情数据-仓单日报

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
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.FuturesGFEXWarehouseReceipt("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesHFMinuteEMFunc

**描述**: FuturesHFMinuteEMFunc 东方财富-高频期货分钟数据

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
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.FuturesHFMinuteEMFunc("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesHFSpotEMFunc

**描述**: FuturesHFSpotEMFunc 东方财富-高频期货实时行情

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
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.FuturesHFSpotEMFunc("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesHQSubscribeExchangeSymbol

**描述**: FuturesHQSubscribeExchangeSymbol 获取外盘期货品种对照表

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
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.FuturesHQSubscribeExchangeSymbol()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesHistEMFunc

**描述**: FuturesHistEM 东方财富-期货行情-历史数据

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |
| period | string | 周期：daily/weekly/monthly |
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
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.FuturesHistEMFunc("000001", "daily", "20230101", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesHistTableEMFunc

**描述**: FuturesHistTableEM 东方财富-期货行情-交易所品种对照表

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
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.FuturesHistTableEMFunc()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesIndexCCIDX

**描述**: FuturesIndexCCIDX 中证商品指数-商品指数-日频率

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
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.FuturesIndexCCIDX("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesInventoryEMFunc

**描述**: FuturesInventoryEMFunc 东方财富-期货库存数据

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
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.FuturesInventoryEMFunc("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesNewsSHMET

**描述**: FuturesNewsSHMET 上海金属网-快讯

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
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.FuturesNewsSHMET("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesRankSum

**描述**: FuturesRankSum 获取指定交易日五个期货交易所前5/10/15/20会员持仓排名汇总数据

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| date | string | 日期，格式：YYYYMMDD |
| varsList | []string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.FuturesRankSum("20230101", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesRuleEMFunc

**描述**: FuturesRuleEMFunc 东方财富网-期货行情-品种及交易规则

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
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.FuturesRuleEMFunc()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesRuleFunc

**描述**: FuturesRuleFunc 国泰君安期货-交易日历数据表

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
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.FuturesRuleFunc("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesSHFERankTable

**描述**: FuturesSHFERankTable 上海期货交易所会员成交及持仓排名表

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| date | string | 日期，格式：YYYYMMDD |
| varsList | []string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.FuturesSHFERankTable("20230101", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesSHFEWarehouseReceipt

**描述**: FuturesSHFEWarehouseReceipt 上海期货交易所指定交割仓库期货仓单日报

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
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.FuturesSHFEWarehouseReceipt("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesSpotPrice

**描述**: FuturesSpotPrice 获取指定交易日大宗商品现货价格及相应基差

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| date | string | 日期，格式：YYYYMMDD |
| varsList | []string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.FuturesSpotPrice("20230101", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesSpotPriceDaily

**描述**: FuturesSpotPriceDaily 获取指定时间段内大宗商品现货价格及相应基差

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| startDay | string | - |
| endDay | string | - |
| varsList | []string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.FuturesSpotPriceDaily("", "", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesSpotPricePrevious

**描述**: FuturesSpotPricePrevious 获取具体交易日大宗商品现货价格及相应基差(历史格式)

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
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.FuturesSpotPricePrevious("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesSymbolMarkFunc

**描述**: FuturesSymbolMarkFunc 获取期货品种代码映射

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
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.FuturesSymbolMarkFunc()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesToSpotCZCE

**描述**: FuturesToSpotCZCE 郑州商品交易所-期转现统计

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
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.FuturesToSpotCZCE("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesToSpotDCE

**描述**: FuturesToSpotDCE 大连商品交易所-期转现

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
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.FuturesToSpotDCE("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesToSpotSHFE

**描述**: FuturesToSpotSHFE 上海期货交易所-期转现

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
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.FuturesToSpotSHFE("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesWarehouseReceiptCZCE

**描述**: FuturesWarehouseReceiptCZCE 郑州商品交易所-交易数据-仓单日报

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
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.FuturesWarehouseReceiptCZCE("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesWarehouseReceiptDCE

**描述**: FuturesWarehouseReceiptDCE 大连商品交易所-行情数据-统计数据-日统计-仓单日报

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
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.FuturesWarehouseReceiptDCE("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesZhDailySina

**描述**: FuturesZhDailySina 获取期货日线数据

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
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.FuturesZhDailySina("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesZhMinuteSina

**描述**: FuturesZhMinuteSina 获取期货分钟数据

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |
| period | string | 周期：daily/weekly/monthly |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.FuturesZhMinuteSina("000001", "daily")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesZhRealtimeFunc

**描述**: FuturesZhRealtimeFunc 获取期货品种当前时刻所有可交易的合约实时数据

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
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.FuturesZhRealtimeFunc("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesZhSpotFunc

**描述**: FuturesZhSpotFunc 获取期货实时行情

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |
| market | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.FuturesZhSpotFunc("000001", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## GetCFFEXDaily

**描述**: GetCFFEXDaily 中国金融期货交易所-日频率交易数据

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
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.GetCFFEXDaily("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## GetCZCEDaily

**描述**: GetCZCEDaily 郑州商品交易所-日频率-量价数据

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
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.GetCZCEDaily("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## GetDCEDaily

**描述**: GetDCEDaily 大连商品交易所日交易数据

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
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.GetDCEDaily("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## GetFuturesDaily

**描述**: GetFuturesDaily 交易所日交易数据

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| startDate | string | 开始日期，格式：YYYYMMDD |
| endDate | string | 结束日期，格式：YYYYMMDD |
| market | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.GetFuturesDaily("20230101", "20230101", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## GetGFEXDaily

**描述**: GetGFEXDaily 广州期货交易所-日频率-量价数据

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
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.GetGFEXDaily("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## GetINEDaily

**描述**: GetINEDaily 上海国际能源交易中心-日频率-量价数据

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
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.GetINEDaily("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## GetRollYield

**描述**: GetRollYield 指定交易日指定品种（主力和次主力）或任意两个合约的展期收益率

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| date | string | 日期，格式：YYYYMMDD |
| variety | string | - |
| symbol1 | string | 股票/基金代码 |
| symbol2 | string | 股票/基金代码 |
| data | []FuturesDailyBar | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.GetRollYield("20230101", "", "000001", "000001", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## GetRollYieldBar

**描述**: GetRollYieldBar 展期收益率

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| typeMethod | string | - |
| variety | string | - |
| date | string | 日期，格式：YYYYMMDD |
| startDay | string | - |
| endDay | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.GetRollYieldBar("", "", "20230101", "", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## GetSHFEDaily

**描述**: GetSHFEDaily 上海期货交易所-日频率-量价数据

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
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.GetSHFEDaily("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## SymbolMarket

**描述**: SymbolMarket 根据品种代码获取交易所代码

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbolDetail | string | 股票/基金代码 |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.SymbolMarket("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## SymbolVarieties

**描述**: SymbolVarieties 从合约代码中提取品种代码

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| contractCode | string | 股票/基金代码 |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/futures"
)

func main() {
	data, err := futures.SymbolVarieties("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

