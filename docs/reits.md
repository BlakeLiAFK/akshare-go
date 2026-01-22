# reits 模块

> 本模块共有 3 个接口

## 目录

- [ReitsHistEm](#reitshistem)
- [ReitsHistMinEm](#reitshistminem)
- [ReitsRealtimeEm](#reitsrealtimeem)

---

## ReitsHistEm

**描述**: ReitsHistEm 获取东方财富网-REITs-沪深REITs-历史行情

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
	"github.com/BlakeLiAFK/akshare/reits"
)

func main() {
	data, err := reits.ReitsHistEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## ReitsHistMinEm

**描述**: ReitsHistMinEm 获取东方财富网-REITs-沪深REITs-分钟行情

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
	"github.com/BlakeLiAFK/akshare/reits"
)

func main() {
	data, err := reits.ReitsHistMinEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## ReitsRealtimeEm

**描述**: ReitsRealtimeEm 获取东方财富网-REITs-沪深REITs-实时行情

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
	"github.com/BlakeLiAFK/akshare/reits"
)

func main() {
	data, err := reits.ReitsRealtimeEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

