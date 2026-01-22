# crypto 模块

> 本模块共有 2 个接口

## 目录

- [CryptoBitcoinCME](#cryptobitcoincme)
- [CryptoBitcoinHoldReport](#cryptobitcoinholdreport)

---

## CryptoBitcoinCME

**描述**: CryptoBitcoinCME 芝加哥商业交易所-比特币成交量报告

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
	"github.com/BlakeLiAFK/akshare/crypto"
)

func main() {
	data, err := crypto.CryptoBitcoinCME("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## CryptoBitcoinHoldReport

**描述**: CryptoBitcoinHoldReport 金十数据-比特币持仓报告

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
	"github.com/BlakeLiAFK/akshare/crypto"
)

func main() {
	data, err := crypto.CryptoBitcoinHoldReport()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

