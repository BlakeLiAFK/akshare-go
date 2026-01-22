# forex 模块

> 本模块共有 3 个接口

## 目录

- [ForexHistEm](#forexhistem)
- [ForexSpotEm](#forexspotem)
- [GetMarketCode](#getmarketcode)

---

## ForexHistEm

**描述**: ForexHistEm 东方财富网-行情中心-外汇市场-所有汇率-历史行情数据

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
	"github.com/BlakeLiAFK/akshare/forex"
)

func main() {
	data, err := forex.ForexHistEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## ForexSpotEm

**描述**: ForexSpotEm 东方财富网-行情中心-外汇市场-所有汇率-实时行情数据

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
	"github.com/BlakeLiAFK/akshare/forex"
)

func main() {
	data, err := forex.ForexSpotEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## GetMarketCode

**描述**: GetMarketCode 根据货币对获取市场代码

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
	"github.com/BlakeLiAFK/akshare/forex"
)

func main() {
	data, err := forex.GetMarketCode("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

