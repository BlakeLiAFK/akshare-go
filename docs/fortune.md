# fortune 模块

> 本模块共有 6 个接口

## 目录

- [ForbesRank](#forbesrank)
- [FortuneRank](#fortunerank)
- [HurunRank](#hurunrank)
- [IndexBloombergBillionaires](#indexbloombergbillionaires)
- [IndexBloombergBillionairesHist](#indexbloombergbillionaireshist)
- [XincaifuRank](#xincaifurank)

---

## ForbesRank

**描述**: ForbesRank 获取福布斯中国榜单

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
	"github.com/BlakeLiAFK/akshare/fortune"
)

func main() {
	data, err := fortune.ForbesRank("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FortuneRank

**描述**: FortuneRank 获取历年世界500强榜单数据

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| year | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/fortune"
)

func main() {
	data, err := fortune.FortuneRank("")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## HurunRank

**描述**: HurunRank 获取胡润排行榜

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| indicator | string | 指标类型 |
| year | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/fortune"
)

func main() {
	data, err := fortune.HurunRank("", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexBloombergBillionaires

**描述**: IndexBloombergBillionaires 获取彭博亿万富豪指数

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
	"github.com/BlakeLiAFK/akshare/fortune"
)

func main() {
	data, err := fortune.IndexBloombergBillionaires()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexBloombergBillionairesHist

**描述**: IndexBloombergBillionairesHist 获取彭博亿万富豪指数历史数据

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| year | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/fortune"
)

func main() {
	data, err := fortune.IndexBloombergBillionairesHist("")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## XincaifuRank

**描述**: XincaifuRank 获取新财富500人富豪榜

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| year | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/fortune"
)

func main() {
	data, err := fortune.XincaifuRank("")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

