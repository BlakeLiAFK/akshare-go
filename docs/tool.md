# tool 模块

> 本模块共有 1 个接口

## 目录

- [ToolTradeDateHistSina](#tooltradedatehistsina)

---

## ToolTradeDateHistSina

**描述**: ToolTradeDateHistSina 新浪财经-交易日历-历史数据

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
	"github.com/BlakeLiAFK/akshare/tool"
)

func main() {
	data, err := tool.ToolTradeDateHistSina()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

