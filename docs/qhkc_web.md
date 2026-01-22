# qhkc_web 模块

> 本模块共有 3 个接口

## 目录

- [QhkcWebBasis](#qhkcwebbasis)
- [QhkcWebInventory](#qhkcwebinventory)
- [QhkcWebProfit](#qhkcwebprofit)

---

## QhkcWebBasis

**描述**: QhkcWebBasis 获取奇货可查网页版基差数据

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
	"github.com/BlakeLiAFK/akshare/qhkc_web"
)

func main() {
	data, err := qhkc_web.QhkcWebBasis("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## QhkcWebInventory

**描述**: QhkcWebInventory 获取奇货可查网页版库存数据

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
	"github.com/BlakeLiAFK/akshare/qhkc_web"
)

func main() {
	data, err := qhkc_web.QhkcWebInventory("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## QhkcWebProfit

**描述**: QhkcWebProfit 获取奇货可查网页版利润数据

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
	"github.com/BlakeLiAFK/akshare/qhkc_web"
)

func main() {
	data, err := qhkc_web.QhkcWebProfit("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

