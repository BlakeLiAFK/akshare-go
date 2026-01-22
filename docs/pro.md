# pro 模块

> 本模块共有 14 个接口

## 目录

- [BrokerPositions](#brokerpositions)
- [IndexInfo](#indexinfo)
- [IndexKline](#indexkline)
- [IndexMember](#indexmember)
- [IndexWeights](#indexweights)
- [LongPool](#longpool)
- [NewDataApi](#newdataapi)
- [ProApi](#proapi)
- [Query](#query)
- [QueryList](#querylist)
- [SetToken](#settoken)
- [ShortPool](#shortpool)
- [VarietyAll](#varietyall)
- [VarietyPositions](#varietypositions)

---

## BrokerPositions

**描述**: BrokerPositions 获取席位持仓数据

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| broker | string | - |
| date | string | 日期，格式：YYYYMMDD |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/pro"
)

func main() {
	data, err := pro.BrokerPositions("", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexInfo

**描述**: IndexInfo 获取指数信息

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| indexID | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/pro"
)

func main() {
	data, err := pro.IndexInfo("")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexKline

**描述**: IndexKline 获取指数行情数据

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| indexID | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/pro"
)

func main() {
	data, err := pro.IndexKline("")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexMember

**描述**: IndexMember 获取指数沉淀资金数据

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| indexID | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/pro"
)

func main() {
	data, err := pro.IndexMember("")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexWeights

**描述**: IndexWeights 获取指数权重数据

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| indexID | string | - |
| date | string | 日期，格式：YYYYMMDD |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/pro"
)

func main() {
	data, err := pro.IndexWeights("", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## LongPool

**描述**: LongPool 获取龙虎牛熊多头合约池

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
	"github.com/BlakeLiAFK/akshare/pro"
)

func main() {
	data, err := pro.LongPool("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## NewDataApi

**描述**: NewDataApi 创建新的数据API客户端

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| token | string | - |
| timeout | int | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/pro"
)

func main() {
	data, err := pro.NewDataApi("", 1)
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## ProApi

**描述**: ProApi 初始化 pro API

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| token | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/pro"
)

func main() {
	data, err := pro.ProApi("")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## Query

**描述**: Query 通用查询接口

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| apiName | string | - |
| kwargs | unknown | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/pro"
)

func main() {
	data, err := pro.Query("", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## QueryList

**描述**: QueryList 查询返回列表数据

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| apiName | string | - |
| kwargs | unknown | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/pro"
)

func main() {
	data, err := pro.QueryList("", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## SetToken

**描述**: SetToken 设置token到环境变量

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| token | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/pro"
)

func main() {
	data, err := pro.SetToken("")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## ShortPool

**描述**: ShortPool 获取龙虎牛熊空头合约池

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
	"github.com/BlakeLiAFK/akshare/pro"
)

func main() {
	data, err := pro.ShortPool("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## VarietyAll

**描述**: VarietyAll 获取商品列表数据

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
	"github.com/BlakeLiAFK/akshare/pro"
)

func main() {
	data, err := pro.VarietyAll()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## VarietyPositions

**描述**: VarietyPositions 获取合约持仓数据

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
	"github.com/BlakeLiAFK/akshare/pro"
)

func main() {
	data, err := pro.VarietyPositions("000001", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

