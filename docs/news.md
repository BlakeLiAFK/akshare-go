# news 模块

> 本模块共有 6 个接口

## 目录

- [NewsCctv](#newscctv)
- [NewsEconomicBaidu](#newseconomicbaidu)
- [NewsReportTimeBaidu](#newsreporttimebaidu)
- [NewsTradeNotifyDividendBaidu](#newstradenotifydividendbaidu)
- [NewsTradeNotifySuspendBaidu](#newstradenotifysuspendbaidu)
- [StockNewsEm](#stocknewsem)

---

## NewsCctv

**描述**: NewsCctv 新闻联播文字稿

**数据源**: https://tv.cctv.com/lm/xwlb

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
	"github.com/BlakeLiAFK/akshare/news"
)

func main() {
	data, err := news.NewsCctv("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## NewsEconomicBaidu

**描述**: NewsEconomicBaidu 百度股市通-经济数据

**数据源**: https://gushitong.baidu.com/calendar

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| date | string | 日期，格式：YYYYMMDD |
| cookie | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/news"
)

func main() {
	data, err := news.NewsEconomicBaidu("20230101", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## NewsReportTimeBaidu

**描述**: NewsReportTimeBaidu 百度股市通-财报发行

**数据源**: https://gushitong.baidu.com/calendar

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| date | string | 日期，格式：YYYYMMDD |
| cookie | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/news"
)

func main() {
	data, err := news.NewsReportTimeBaidu("20230101", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## NewsTradeNotifyDividendBaidu

**描述**: NewsTradeNotifyDividendBaidu 百度股市通-交易提醒-分红派息

**数据源**: https://gushitong.baidu.com/calendar

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| date | string | 日期，格式：YYYYMMDD |
| cookie | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/news"
)

func main() {
	data, err := news.NewsTradeNotifyDividendBaidu("20230101", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## NewsTradeNotifySuspendBaidu

**描述**: NewsTradeNotifySuspendBaidu 百度股市通-交易提醒-停复牌

**数据源**: https://gushitong.baidu.com/calendar

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| date | string | 日期，格式：YYYYMMDD |
| cookie | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/news"
)

func main() {
	data, err := news.NewsTradeNotifySuspendBaidu("20230101", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockNewsEm

**描述**: StockNewsEm 东方财富-个股新闻-最近 100 条新闻

**数据源**: https://so.eastmoney.com/news/s?keyword=603777

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
	"github.com/BlakeLiAFK/akshare/news"
)

func main() {
	data, err := news.StockNewsEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

