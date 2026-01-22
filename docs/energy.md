# energy 模块

> 本模块共有 8 个接口

## 目录

- [EnergyCarbonBJ](#energycarbonbj)
- [EnergyCarbonDomestic](#energycarbondomestic)
- [EnergyCarbonEU](#energycarboneu)
- [EnergyCarbonGZ](#energycarbongz)
- [EnergyCarbonHB](#energycarbonhb)
- [EnergyCarbonSZ](#energycarbonsz)
- [EnergyOilDetail](#energyoildetail)
- [EnergyOilHist](#energyoilhist)

---

## EnergyCarbonBJ

**描述**: EnergyCarbonBJ 北京市碳排放权电子交易平台-北京市碳排放权公开交易行情

**数据源**: https://www.bjets.com.cn/article/jyxx/

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
	"github.com/BlakeLiAFK/akshare/energy"
)

func main() {
	data, err := energy.EnergyCarbonBJ()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## EnergyCarbonDomestic

**描述**: EnergyCarbonDomestic 碳交易网-行情信息

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
	"github.com/BlakeLiAFK/akshare/energy"
)

func main() {
	data, err := energy.EnergyCarbonDomestic("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## EnergyCarbonEU

**描述**: EnergyCarbonEU 深圳碳排放交易所-国际碳情

**数据源**: http://www.cerx.cn/dailynewsOuter/index.htm

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
	"github.com/BlakeLiAFK/akshare/energy"
)

func main() {
	data, err := energy.EnergyCarbonEU()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## EnergyCarbonGZ

**描述**: EnergyCarbonGZ 广州碳排放权交易中心-行情信息

**数据源**: http://www.cnemission.com/article/hqxx/

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
	"github.com/BlakeLiAFK/akshare/energy"
)

func main() {
	data, err := energy.EnergyCarbonGZ()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## EnergyCarbonHB

**描述**: EnergyCarbonHB 湖北碳排放权交易中心-现货交易数据-配额-每日概况

**数据源**: http://www.hbets.cn/

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
	"github.com/BlakeLiAFK/akshare/energy"
)

func main() {
	data, err := energy.EnergyCarbonHB()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## EnergyCarbonSZ

**描述**: EnergyCarbonSZ 深圳碳排放交易所-国内碳情

**数据源**: http://www.cerx.cn/dailynewsCN/index.htm

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
	"github.com/BlakeLiAFK/akshare/energy"
)

func main() {
	data, err := energy.EnergyCarbonSZ()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## EnergyOilDetail

**描述**: EnergyOilDetail 全国各地区的汽油和柴油油价

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
	"github.com/BlakeLiAFK/akshare/energy"
)

func main() {
	data, err := energy.EnergyOilDetail("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## EnergyOilHist

**描述**: EnergyOilHist 汽柴油历史调价信息

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
	"github.com/BlakeLiAFK/akshare/energy"
)

func main() {
	data, err := energy.EnergyOilHist()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

