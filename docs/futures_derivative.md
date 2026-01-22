# futures_derivative 模块

> 本模块共有 15 个接口

## 目录

- [FuturesContractInfoCffex](#futurescontractinfocffex)
- [FuturesContractInfoCzce](#futurescontractinfoczce)
- [FuturesContractInfoDce](#futurescontractinfodce)
- [FuturesContractInfoGfex](#futurescontractinfogfex)
- [FuturesContractInfoIne](#futurescontractinfoine)
- [FuturesContractInfoShfe](#futurescontractinfoshfe)
- [FuturesDisplayMainSina](#futuresdisplaymainsina)
- [FuturesHogCore](#futureshogcore)
- [FuturesHogCost](#futureshogcost)
- [FuturesHogSupply](#futureshogsupply)
- [FuturesHoldPosSina](#futuresholdpossina)
- [FuturesMainSina](#futuresmainsina)
- [FuturesSpotSys](#futuresspotsys)
- [MatchMainContract](#matchmaincontract)
- [ZhSubscribeExchangeSymbol](#zhsubscribeexchangesymbol)

---

## FuturesContractInfoCffex

**描述**: FuturesContractInfoCffex 中国金融期货交易所-数据-交易参数

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
	"github.com/BlakeLiAFK/akshare/futures_derivative"
)

func main() {
	data, err := futures_derivative.FuturesContractInfoCffex("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesContractInfoCzce

**描述**: FuturesContractInfoCzce 郑州商品交易所-交易数据-参考数据

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
	"github.com/BlakeLiAFK/akshare/futures_derivative"
)

func main() {
	data, err := futures_derivative.FuturesContractInfoCzce("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesContractInfoDce

**描述**: FuturesContractInfoDce 大连商品交易所-业务/服务-业务参数-交易参数-合约信息查询

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
	"github.com/BlakeLiAFK/akshare/futures_derivative"
)

func main() {
	data, err := futures_derivative.FuturesContractInfoDce()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesContractInfoGfex

**描述**: FuturesContractInfoGfex 广州期货交易所-业务/服务-合约信息

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
	"github.com/BlakeLiAFK/akshare/futures_derivative"
)

func main() {
	data, err := futures_derivative.FuturesContractInfoGfex()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesContractInfoIne

**描述**: FuturesContractInfoIne 上海国际能源交易中心-业务指南-交易参数汇总(期货)

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
	"github.com/BlakeLiAFK/akshare/futures_derivative"
)

func main() {
	data, err := futures_derivative.FuturesContractInfoIne("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesContractInfoShfe

**描述**: FuturesContractInfoShfe 上海期货交易所-交易所服务-业务数据-交易参数汇总查询

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
	"github.com/BlakeLiAFK/akshare/futures_derivative"
)

func main() {
	data, err := futures_derivative.FuturesContractInfoShfe("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesDisplayMainSina

**描述**: FuturesDisplayMainSina 新浪主力连续合约品种一览表

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
	"github.com/BlakeLiAFK/akshare/futures_derivative"
)

func main() {
	data, err := futures_derivative.FuturesDisplayMainSina()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesHogCore

**描述**: FuturesHogCore 玄田数据-核心数据

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
	"github.com/BlakeLiAFK/akshare/futures_derivative"
)

func main() {
	data, err := futures_derivative.FuturesHogCore("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesHogCost

**描述**: FuturesHogCost 玄田数据-成本维度

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
	"github.com/BlakeLiAFK/akshare/futures_derivative"
)

func main() {
	data, err := futures_derivative.FuturesHogCost("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesHogSupply

**描述**: FuturesHogSupply 玄田数据-供应维度

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
	"github.com/BlakeLiAFK/akshare/futures_derivative"
)

func main() {
	data, err := futures_derivative.FuturesHogSupply("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesHoldPosSina

**描述**: FuturesHoldPosSina 新浪财经-期货-成交持仓

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbolType | string | 股票/基金代码 |
| contract | string | - |
| date | string | 日期，格式：YYYYMMDD |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/futures_derivative"
)

func main() {
	data, err := futures_derivative.FuturesHoldPosSina("000001", "", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesMainSina

**描述**: FuturesMainSina 新浪期货主力连续日数据

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
	"github.com/BlakeLiAFK/akshare/futures_derivative"
)

func main() {
	data, err := futures_derivative.FuturesMainSina("000001", "20230101", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FuturesSpotSys

**描述**: FuturesSpotSys 生意社-商品与期货-现期图

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |
| indicator | string | 指标类型 |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/futures_derivative"
)

func main() {
	data, err := futures_derivative.FuturesSpotSys("000001", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MatchMainContract

**描述**: MatchMainContract 获取指定交易所所有可提供数据的合约

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
	"github.com/BlakeLiAFK/akshare/futures_derivative"
)

func main() {
	data, err := futures_derivative.MatchMainContract("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## ZhSubscribeExchangeSymbol

**描述**: ZhSubscribeExchangeSymbol 订阅指定交易所品种代码

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
	"github.com/BlakeLiAFK/akshare/futures_derivative"
)

func main() {
	data, err := futures_derivative.ZhSubscribeExchangeSymbol("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

