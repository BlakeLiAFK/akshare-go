# event 模块

> 本模块共有 6 个接口

## 目录

- [GetCityCode](#getcitycode)
- [GetCityName](#getcityname)
- [GetProvinceCode](#getprovincecode)
- [GetProvinceName](#getprovincename)
- [MigrationAreaBaidu](#migrationareabaidu)
- [MigrationScaleBaidu](#migrationscalebaidu)

---

## GetCityCode

**描述**: GetCityCode 根据城市名称获取城市代码

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| name | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/event"
)

func main() {
	data, err := event.GetCityCode("")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## GetCityName

**描述**: GetCityName 根据城市代码获取城市名称

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| code | string | 股票/基金代码 |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/event"
)

func main() {
	data, err := event.GetCityName("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## GetProvinceCode

**描述**: GetProvinceCode 根据省份名称获取省份代码

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| name | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/event"
)

func main() {
	data, err := event.GetProvinceCode("")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## GetProvinceName

**描述**: GetProvinceName 根据省份代码获取省份名称

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| code | string | 股票/基金代码 |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/event"
)

func main() {
	data, err := event.GetProvinceName("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MigrationAreaBaidu

**描述**: MigrationAreaBaidu 百度地图慧眼-百度迁徙-迁入/迁出地详情

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| area | string | - |
| indicator | string | 指标类型 |
| date | string | 日期，格式：YYYYMMDD |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/event"
)

func main() {
	data, err := event.MigrationAreaBaidu("", "", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MigrationScaleBaidu

**描述**: MigrationScaleBaidu 百度地图慧眼-百度迁徙-迁徙规模

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| area | string | - |
| indicator | string | 指标类型 |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/event"
)

func main() {
	data, err := event.MigrationScaleBaidu("", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

