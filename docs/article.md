# article 模块

> 本模块共有 7 个接口

## 目录

- [ArticleEPUIndex](#articleepuindex)
- [ArticleFFCRR](#articleffcrr)
- [ArticleOmanRV](#articleomanrv)
- [ArticleOmanRVShort](#articleomanrvshort)
- [ArticleRlabRV](#articlerlabrv)
- [FredMD](#fredmd)
- [FredQD](#fredqd)

---

## ArticleEPUIndex

**描述**: ArticleEPUIndex 获取主要国家和地区的经济政策不确定性(EPU)指数

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
	"github.com/BlakeLiAFK/akshare/article"
)

func main() {
	data, err := article.ArticleEPUIndex("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## ArticleFFCRR

**描述**: ArticleFFCRR 获取 Fama-French Current Research Returns 多因子数据

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
	"github.com/BlakeLiAFK/akshare/article"
)

func main() {
	data, err := article.ArticleFFCRR()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## ArticleOmanRV

**描述**: ArticleOmanRV 获取 Oxford-Man 已实现波动率数据

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |
| index | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/article"
)

func main() {
	data, err := article.ArticleOmanRV("000001", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## ArticleOmanRVShort

**描述**: ArticleOmanRVShort 获取 Oxford-Man 简化版已实现波动率数据

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
	"github.com/BlakeLiAFK/akshare/article"
)

func main() {
	data, err := article.ArticleOmanRVShort("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## ArticleRlabRV

**描述**: ArticleRlabRV 获取 Risk-Lab 已实现波动率数据

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
	"github.com/BlakeLiAFK/akshare/article"
)

func main() {
	data, err := article.ArticleRlabRV("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FredMD

**描述**: FredMD 获取FRED-MD月度宏观经济数据集

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
	"github.com/BlakeLiAFK/akshare/article"
)

func main() {
	data, err := article.FredMD("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FredQD

**描述**: FredQD 获取FRED-QD季度宏观经济数据集

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
	"github.com/BlakeLiAFK/akshare/article"
)

func main() {
	data, err := article.FredQD("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

