# fx 模块

> 本模块共有 6 个接口

## 目录

- [CurrencyPairMap](#currencypairmap)
- [FxCSwapCm](#fxcswapcm)
- [FxPairQuote](#fxpairquote)
- [FxQuoteBaidu](#fxquotebaidu)
- [FxSpotQuote](#fxspotquote)
- [FxSwapQuote](#fxswapquote)

---

## CurrencyPairMap

**描述**: CurrencyPairMap 获取指定货币的所有可获取货币对的数据

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
	"github.com/BlakeLiAFK/akshare/fx"
)

func main() {
	data, err := fx.CurrencyPairMap("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FxCSwapCm

**描述**: FxCSwapCm 获取外汇掉期C-Swap定盘曲线

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
	"github.com/BlakeLiAFK/akshare/fx"
)

func main() {
	data, err := fx.FxCSwapCm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FxPairQuote

**描述**: FxPairQuote 获取外币对即期报价

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
	"github.com/BlakeLiAFK/akshare/fx"
)

func main() {
	data, err := fx.FxPairQuote()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FxQuoteBaidu

**描述**: FxQuoteBaidu 获取百度股市通-外汇-行情榜单

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
	"github.com/BlakeLiAFK/akshare/fx"
)

func main() {
	data, err := fx.FxQuoteBaidu("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FxSpotQuote

**描述**: FxSpotQuote 获取人民币外汇即期报价

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
	"github.com/BlakeLiAFK/akshare/fx"
)

func main() {
	data, err := fx.FxSpotQuote()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FxSwapQuote

**描述**: FxSwapQuote 获取人民币外汇远掉报价

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
	"github.com/BlakeLiAFK/akshare/fx"
)

func main() {
	data, err := fx.FxSwapQuote()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

