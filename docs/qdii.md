# qdii 模块

> 本模块共有 3 个接口

## 目录

- [QdiiAIndexJsl](#qdiiaindexjsl)
- [QdiiECommJsl](#qdiiecommjsl)
- [QdiiEIndexJsl](#qdiieindexjsl)

---

## QdiiAIndexJsl

**描述**: QdiiAIndexJsl 获取集思录-T+0 QDII-亚洲市场-亚洲指数数据

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
	"github.com/BlakeLiAFK/akshare/qdii"
)

func main() {
	data, err := qdii.QdiiAIndexJsl()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## QdiiECommJsl

**描述**: QdiiECommJsl 获取集思录-T+0 QDII-欧美市场-欧美商品数据

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
	"github.com/BlakeLiAFK/akshare/qdii"
)

func main() {
	data, err := qdii.QdiiECommJsl()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## QdiiEIndexJsl

**描述**: QdiiEIndexJsl 获取集思录-T+0 QDII-欧美市场-欧美指数数据

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
	"github.com/BlakeLiAFK/akshare/qdii"
)

func main() {
	data, err := qdii.QdiiEIndexJsl()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

