# movie 模块

> 本模块共有 12 个接口

## 目录

- [BusinessValueArtist](#businessvalueartist)
- [MovieBoxofficeCinemaDaily](#movieboxofficecinemadaily)
- [MovieBoxofficeCinemaWeekly](#movieboxofficecinemaweekly)
- [MovieBoxofficeDaily](#movieboxofficedaily)
- [MovieBoxofficeMonthly](#movieboxofficemonthly)
- [MovieBoxofficeRealtime](#movieboxofficerealtime)
- [MovieBoxofficeWeekly](#movieboxofficeweekly)
- [MovieBoxofficeYearly](#movieboxofficeyearly)
- [MovieBoxofficeYearlyFirstWeek](#movieboxofficeyearlyfirstweek)
- [OnlineValueArtist](#onlinevalueartist)
- [VideoTv](#videotv)
- [VideoVarietyShow](#videovarietyshow)

---

## BusinessValueArtist

**描述**: BusinessValueArtist 获取艺恩-艺人-艺人商业价值

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
	"github.com/BlakeLiAFK/akshare/movie"
)

func main() {
	data, err := movie.BusinessValueArtist()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MovieBoxofficeCinemaDaily

**描述**: MovieBoxofficeCinemaDaily 获取电影票房-影院票房-日票房排行

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
	"github.com/BlakeLiAFK/akshare/movie"
)

func main() {
	data, err := movie.MovieBoxofficeCinemaDaily("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MovieBoxofficeCinemaWeekly

**描述**: MovieBoxofficeCinemaWeekly 获取电影票房-影院票房-周票房排行

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
	"github.com/BlakeLiAFK/akshare/movie"
)

func main() {
	data, err := movie.MovieBoxofficeCinemaWeekly("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MovieBoxofficeDaily

**描述**: MovieBoxofficeDaily 获取电影票房-单日票房

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
	"github.com/BlakeLiAFK/akshare/movie"
)

func main() {
	data, err := movie.MovieBoxofficeDaily("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MovieBoxofficeMonthly

**描述**: MovieBoxofficeMonthly 获取电影票房-单月票房

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
	"github.com/BlakeLiAFK/akshare/movie"
)

func main() {
	data, err := movie.MovieBoxofficeMonthly("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MovieBoxofficeRealtime

**描述**: MovieBoxofficeRealtime 获取电影票房-实时票房

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
	"github.com/BlakeLiAFK/akshare/movie"
)

func main() {
	data, err := movie.MovieBoxofficeRealtime()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MovieBoxofficeWeekly

**描述**: MovieBoxofficeWeekly 获取电影票房-单周票房

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
	"github.com/BlakeLiAFK/akshare/movie"
)

func main() {
	data, err := movie.MovieBoxofficeWeekly("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MovieBoxofficeYearly

**描述**: MovieBoxofficeYearly 获取电影票房-年度票房

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
	"github.com/BlakeLiAFK/akshare/movie"
)

func main() {
	data, err := movie.MovieBoxofficeYearly("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MovieBoxofficeYearlyFirstWeek

**描述**: MovieBoxofficeYearlyFirstWeek 获取电影票房-年度票房-年度首周票房

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
	"github.com/BlakeLiAFK/akshare/movie"
)

func main() {
	data, err := movie.MovieBoxofficeYearlyFirstWeek("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## OnlineValueArtist

**描述**: OnlineValueArtist 获取艺恩-艺人-艺人流量价值

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
	"github.com/BlakeLiAFK/akshare/movie"
)

func main() {
	data, err := movie.OnlineValueArtist()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## VideoTv

**描述**: VideoTv 获取艺恩-视频放映-电视剧集

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
	"github.com/BlakeLiAFK/akshare/movie"
)

func main() {
	data, err := movie.VideoTv()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## VideoVarietyShow

**描述**: VideoVarietyShow 获取艺恩-视频放映-综艺节目

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
	"github.com/BlakeLiAFK/akshare/movie"
)

func main() {
	data, err := movie.VideoVarietyShow()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

