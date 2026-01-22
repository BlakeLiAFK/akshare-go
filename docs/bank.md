# bank 模块

> 本模块共有 4 个接口

## 目录

- [BankFjcfPageUrl](#bankfjcfpageurl)
- [BankFjcfTableDetail](#bankfjcftabledetail)
- [BankFjcfTotalNum](#bankfjcftotalnum)
- [BankFjcfTotalPage](#bankfjcftotalpage)

---

## BankFjcfPageUrl

**描述**: BankFjcfPageUrl 获取银保监分局行政处罚的分页数据列表

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| item | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/bank"
)

func main() {
	data, err := bank.BankFjcfPageUrl("")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## BankFjcfTableDetail

**描述**: BankFjcfTableDetail 获取银保监分局行政处罚的详细表格数据

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| url | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/bank"
)

func main() {
	data, err := bank.BankFjcfTableDetail("")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## BankFjcfTotalNum

**描述**: BankFjcfTotalNum 获取银保监分局行政处罚的总记录数

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| item | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/bank"
)

func main() {
	data, err := bank.BankFjcfTotalNum("")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## BankFjcfTotalPage

**描述**: BankFjcfTotalPage 计算银保监分局行政处罚的总页数

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| item | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/bank"
)

func main() {
	data, err := bank.BankFjcfTotalPage("")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

