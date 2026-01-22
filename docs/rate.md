# rate 模块

> 本模块共有 4 个接口

## 目录

- [RateLpr](#ratelpr)
- [RateShibor](#rateshibor)
- [RepoRateHist](#reporatehist)
- [RepoRateQuery](#reporatequery)

---

## RateLpr

**描述**: RateLpr 获取LPR利率数据

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
	"github.com/BlakeLiAFK/akshare/rate"
)

func main() {
	data, err := rate.RateLpr()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## RateShibor

**描述**: RateShibor 获取Shibor利率数据

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
	"github.com/BlakeLiAFK/akshare/rate"
)

func main() {
	data, err := rate.RateShibor()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## RepoRateHist

**描述**: RepoRateHist 获取回购定盘利率历史数据

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| startDate | string | 开始日期，格式：YYYYMMDD |
| endDate | string | 结束日期，格式：YYYYMMDD |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/rate"
)

func main() {
	data, err := rate.RepoRateHist("20230101", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## RepoRateQuery

**描述**: RepoRateQuery 获取回购定盘利率查询数据

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
	"github.com/BlakeLiAFK/akshare/rate"
)

func main() {
	data, err := rate.RepoRateQuery("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

