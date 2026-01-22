# currency 模块

> 本模块共有 7 个接口

## 目录

- [CurrencyBocSafe](#currencybocsafe)
- [CurrencyBocSina](#currencybocsina)
- [CurrencyConvert](#currencyconvert)
- [CurrencyCurrencies](#currencycurrencies)
- [CurrencyHistory](#currencyhistory)
- [CurrencyLatest](#currencylatest)
- [CurrencyTimeSeries](#currencytimeseries)

---

## CurrencyBocSafe

**描述**: CurrencyBocSafe 人民币汇率中间价

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
	"github.com/BlakeLiAFK/akshare/currency"
)

func main() {
	data, err := currency.CurrencyBocSafe()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## CurrencyBocSina

**描述**: CurrencyBocSina 中国银行人民币牌价历史数据查询

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
	"github.com/BlakeLiAFK/akshare/currency"
)

func main() {
	data, err := currency.CurrencyBocSina("000001", "20230101", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## CurrencyConvert

**描述**: CurrencyConvert 货币转换

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| from | string | - |
| to | string | - |
| amount | float64 | - |
| apiKey | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/currency"
)

func main() {
	data, err := currency.CurrencyConvert("", "", 1.0, "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## CurrencyCurrencies

**描述**: CurrencyCurrencies 获取货币列表

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| cType | string | - |
| apiKey | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/currency"
)

func main() {
	data, err := currency.CurrencyCurrencies("", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## CurrencyHistory

**描述**: CurrencyHistory 获取历史汇率数据

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| base | string | - |
| date | string | 日期，格式：YYYYMMDD |
| symbols | string | 股票/基金代码 |
| apiKey | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/currency"
)

func main() {
	data, err := currency.CurrencyHistory("", "20230101", "000001", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## CurrencyLatest

**描述**: CurrencyLatest 获取最新汇率数据

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| base | string | - |
| symbols | string | 股票/基金代码 |
| apiKey | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/currency"
)

func main() {
	data, err := currency.CurrencyLatest("", "000001", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## CurrencyTimeSeries

**描述**: CurrencyTimeSeries 获取汇率时间序列数据

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| base | string | - |
| startDate | string | 开始日期，格式：YYYYMMDD |
| endDate | string | 结束日期，格式：YYYYMMDD |
| symbols | string | 股票/基金代码 |
| apiKey | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/currency"
)

func main() {
	data, err := currency.CurrencyTimeSeries("", "20230101", "20230101", "000001", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

