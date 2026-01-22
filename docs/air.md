# air 模块

> 本模块共有 8 个接口

## 目录

- [AirCityTable](#aircitytable)
- [AirQualityHebei](#airqualityhebei)
- [AirQualityHist](#airqualityhist)
- [AirQualityRank](#airqualityrank)
- [AirQualityWatchPoint](#airqualitywatchpoint)
- [SunriseCityList](#sunrisecitylist)
- [SunriseDaily](#sunrisedaily)
- [SunriseMonthly](#sunrisemonthly)

---

## AirCityTable

**描述**: AirCityTable 获取真气网所有城市列表及其空气质量

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
	"github.com/BlakeLiAFK/akshare/air"
)

func main() {
	data, err := air.AirCityTable()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## AirQualityHebei

**描述**: AirQualityHebei 获取河北省空气质量数据

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
	"github.com/BlakeLiAFK/akshare/air"
)

func main() {
	data, err := air.AirQualityHebei()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## AirQualityHist

**描述**: AirQualityHist 获取指定城市的历史空气质量数据

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| city | string | - |
| period | string | 周期：daily/weekly/monthly |
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
	"github.com/BlakeLiAFK/akshare/air"
)

func main() {
	data, err := air.AirQualityHist("", "daily", "20230101", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## AirQualityRank

**描述**: AirQualityRank 获取空气质量排名

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
	"github.com/BlakeLiAFK/akshare/air"
)

func main() {
	data, err := air.AirQualityRank()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## AirQualityWatchPoint

**描述**: AirQualityWatchPoint 获取指定城市监测点的空气质量数据

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| city | string | - |
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
	"github.com/BlakeLiAFK/akshare/air"
)

func main() {
	data, err := air.AirQualityWatchPoint("", "20230101", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## SunriseCityList

**描述**: SunriseCityList 获取支持查询日出日落数据的城市列表

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
	"github.com/BlakeLiAFK/akshare/air"
)

func main() {
	data, err := air.SunriseCityList()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## SunriseDaily

**描述**: SunriseDaily 获取指定城市指定日期的日出日落数据

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| date | string | 日期，格式：YYYYMMDD |
| city | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/air"
)

func main() {
	data, err := air.SunriseDaily("20230101", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## SunriseMonthly

**描述**: SunriseMonthly 获取指定月份的每日日出日落数据

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| date | string | 日期，格式：YYYYMMDD |
| city | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/air"
)

func main() {
	data, err := air.SunriseMonthly("20230101", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

