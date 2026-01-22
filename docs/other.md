# other 模块

> 本模块共有 7 个接口

## 目录

- [CarMarketCateCPCA](#carmarketcatecpca)
- [CarMarketCountryCPCA](#carmarketcountrycpca)
- [CarMarketFuelCPCA](#carmarketfuelcpca)
- [CarMarketManRankCPCA](#carmarketmanrankcpca)
- [CarMarketSegmentCPCA](#carmarketsegmentcpca)
- [CarMarketTotalCPCA](#carmarkettotalcpca)
- [CarSaleRankGasgoo](#carsalerankgasgoo)

---

## CarMarketCateCPCA

**描述**: CarMarketCateCPCA 乘联会-统计数据-车型大类

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |
| indicator | string | 指标类型 |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/other"
)

func main() {
	data, err := other.CarMarketCateCPCA("000001", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## CarMarketCountryCPCA

**描述**: CarMarketCountryCPCA 乘联会-统计数据-国别细分市场

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
	"github.com/BlakeLiAFK/akshare/other"
)

func main() {
	data, err := other.CarMarketCountryCPCA()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## CarMarketFuelCPCA

**描述**: CarMarketFuelCPCA 乘联会-统计数据-新能源细分市场

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
	"github.com/BlakeLiAFK/akshare/other"
)

func main() {
	data, err := other.CarMarketFuelCPCA("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## CarMarketManRankCPCA

**描述**: CarMarketManRankCPCA 乘联会-统计数据-厂商排名

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |
| indicator | string | 指标类型 |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/other"
)

func main() {
	data, err := other.CarMarketManRankCPCA("000001", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## CarMarketSegmentCPCA

**描述**: CarMarketSegmentCPCA 乘联会-统计数据-级别细分市场

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
	"github.com/BlakeLiAFK/akshare/other"
)

func main() {
	data, err := other.CarMarketSegmentCPCA("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## CarMarketTotalCPCA

**描述**: CarMarketTotalCPCA 乘联会-统计数据-总体市场

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |
| indicator | string | 指标类型 |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/other"
)

func main() {
	data, err := other.CarMarketTotalCPCA("000001", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## CarSaleRankGasgoo

**描述**: CarSaleRankGasgoo 盖世汽车-汽车行业制造企业数据库-销量数据

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |
| date | string | 日期，格式：YYYYMMDD |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/other"
)

func main() {
	data, err := other.CarSaleRankGasgoo("000001", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

