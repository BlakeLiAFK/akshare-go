# spot 模块

> 本模块共有 7 个接口

## 目录

- [SpotGoldenBenchmarkSge](#spotgoldenbenchmarksge)
- [SpotHistSge](#spothistsge)
- [SpotPriceQh](#spotpriceqh)
- [SpotPriceTableQh](#spotpricetableqh)
- [SpotQuotationsSge](#spotquotationssge)
- [SpotSilverBenchmarkSge](#spotsilverbenchmarksge)
- [SpotSymbolTableSge](#spotsymboltablesge)

---

## SpotGoldenBenchmarkSge

**描述**: SpotGoldenBenchmarkSge 获取上海黄金交易所-数据资讯-上海金基准价-历史数据

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
	"github.com/BlakeLiAFK/akshare/spot"
)

func main() {
	data, err := spot.SpotGoldenBenchmarkSge()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## SpotHistSge

**描述**: SpotHistSge 获取上海黄金交易所-数据资讯-行情走势-历史数据

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
	"github.com/BlakeLiAFK/akshare/spot"
)

func main() {
	data, err := spot.SpotHistSge("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## SpotPriceQh

**描述**: SpotPriceQh 获取99期货-数据-期现-现货走势

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
	"github.com/BlakeLiAFK/akshare/spot"
)

func main() {
	data, err := spot.SpotPriceQh("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## SpotPriceTableQh

**描述**: SpotPriceTableQh 获取99期货现货品种表

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
	"github.com/BlakeLiAFK/akshare/spot"
)

func main() {
	data, err := spot.SpotPriceTableQh()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## SpotQuotationsSge

**描述**: SpotQuotationsSge 获取上海黄金交易所-数据资讯-行情走势-实时数据

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
	"github.com/BlakeLiAFK/akshare/spot"
)

func main() {
	data, err := spot.SpotQuotationsSge("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## SpotSilverBenchmarkSge

**描述**: SpotSilverBenchmarkSge 获取上海黄金交易所-数据资讯-上海银基准价-历史数据

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
	"github.com/BlakeLiAFK/akshare/spot"
)

func main() {
	data, err := spot.SpotSilverBenchmarkSge()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## SpotSymbolTableSge

**描述**: SpotSymbolTableSge 获取上海黄金交易所品种表

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
	"github.com/BlakeLiAFK/akshare/spot"
)

func main() {
	data, err := spot.SpotSymbolTableSge()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

