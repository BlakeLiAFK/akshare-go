# interest_rate 模块

> 本模块共有 1 个接口

## 目录

- [RateInterbank](#rateinterbank)

---

## RateInterbank

**描述**: RateInterbank 获取东方财富-银行间拆借利率数据

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| market | string | - |
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
	"github.com/BlakeLiAFK/akshare/interest_rate"
)

func main() {
	data, err := interest_rate.RateInterbank("", "000001", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

