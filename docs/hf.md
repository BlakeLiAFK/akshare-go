# hf 模块

> 本模块共有 1 个接口

## 目录

- [HfSp500](#hfsp500)

---

## HfSp500

**描述**: HfSp500 获取S&P500高频分钟数据

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| year | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/hf"
)

func main() {
	data, err := hf.HfSp500("")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

