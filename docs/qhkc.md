# qhkc 模块

> 本模块共有 5 个接口

## 目录

- [GetProClient](#getproclient)
- [IndexInfo](#indexinfo)
- [IndexKline](#indexkline)
- [IndexMember](#indexmember)
- [IndexWeights](#indexweights)

---

## GetProClient

**描述**: GetProClient 获取奇货可查API客户端

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
	"github.com/BlakeLiAFK/akshare/qhkc"
)

func main() {
	data, err := qhkc.GetProClient("")
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
| client | *pro.DataApi | - |
| indexID | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/qhkc"
)

func main() {
	data, err := qhkc.IndexInfo(nil, "")
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
| client | *pro.DataApi | - |
| indexID | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/qhkc"
)

func main() {
	data, err := qhkc.IndexKline(nil, "")
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
| client | *pro.DataApi | - |
| indexID | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/qhkc"
)

func main() {
	data, err := qhkc.IndexMember(nil, "")
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
| client | *pro.DataApi | - |
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
	"github.com/BlakeLiAFK/akshare/qhkc"
)

func main() {
	data, err := qhkc.IndexWeights(nil, "", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

