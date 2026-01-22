# cal 模块

> 本模块共有 3 个接口

## 目录

- [RVFromFuturesZhMinuteSina](#rvfromfutureszhminutesina)
- [RVFromStockZhAHistMinEM](#rvfromstockzhahistminem)
- [VolatilityYZRV](#volatilityyzrv)

---

## RVFromFuturesZhMinuteSina

**描述**: RVFromFuturesZhMinuteSina 从新浪财经获取期货的分钟级历史行情数据并格式化

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |
| period | string | 周期：daily/weekly/monthly |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/cal"
)

func main() {
	data, err := cal.RVFromFuturesZhMinuteSina("000001", "daily")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## RVFromStockZhAHistMinEM

**描述**: RVFromStockZhAHistMinEM 从东方财富网获取股票的分钟级历史行情数据并格式化

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |
| startDate | string | 开始日期，格式：YYYYMMDD |
| endDate | string | 结束日期，格式：YYYYMMDD |
| period | string | 周期：daily/weekly/monthly |
| adjust | string | 复权类型：qfq(前复权)/hfq(后复权)/空字符串(不复权) |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/cal"
)

func main() {
	data, err := cal.RVFromStockZhAHistMinEM("000001", "20230101", "20230101", "daily", "qfq")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## VolatilityYZRV

**描述**: VolatilityYZRV 计算 Yang-Zhang 已实现波动率

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| data | dataframe.DataFrame | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/cal"
)

func main() {
	data, err := cal.VolatilityYZRV(nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

