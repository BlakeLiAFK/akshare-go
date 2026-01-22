# nlp 模块

> 本模块共有 6 个接口

## 目录

- [NlpAnswer](#nlpanswer)
- [NlpOwnthink](#nlpownthink)
- [NlpOwnthinkAvp](#nlpownthinkavp)
- [NlpOwnthinkDesc](#nlpownthinkdesc)
- [NlpOwnthinkEntity](#nlpownthinkentity)
- [NlpOwnthinkTag](#nlpownthinktag)

---

## NlpAnswer

**描述**: NlpAnswer 智能问答

**数据源**: https://ownthink.com/robot.html

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| question | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/nlp"
)

func main() {
	data, err := nlp.NlpAnswer("")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## NlpOwnthink

**描述**: NlpOwnthink 知识图谱接口

**数据源**: https://ownthink.com/

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| word | string | - |
| indicator | string | 指标类型 |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/nlp"
)

func main() {
	data, err := nlp.NlpOwnthink("", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## NlpOwnthinkAvp

**描述**: NlpOwnthinkAvp 知识图谱-获取属性值对

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| word | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/nlp"
)

func main() {
	data, err := nlp.NlpOwnthinkAvp("")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## NlpOwnthinkDesc

**描述**: NlpOwnthinkDesc 知识图谱-获取描述

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| word | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/nlp"
)

func main() {
	data, err := nlp.NlpOwnthinkDesc("")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## NlpOwnthinkEntity

**描述**: NlpOwnthinkEntity 知识图谱-获取实体名称

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| word | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/nlp"
)

func main() {
	data, err := nlp.NlpOwnthinkEntity("")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## NlpOwnthinkTag

**描述**: NlpOwnthinkTag 知识图谱-获取标签

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| word | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/nlp"
)

func main() {
	data, err := nlp.NlpOwnthinkTag("")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

