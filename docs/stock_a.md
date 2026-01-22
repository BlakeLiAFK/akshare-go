# stock_a 模块

> 本模块共有 3 个接口

## 目录

- [StockBoardConceptNameEm](#stockboardconceptnameem)
- [StockIndividualFundFlowRank](#stockindividualfundflowrank)
- [StockZhASpotEm](#stockzhaspotem)

---

## StockBoardConceptNameEm

**描述**: StockBoardConceptNameEm 东方财富网-行情中心-沪深京板块-概念板块-名称

**数据源**: https://quote.eastmoney.com/center/boardlist.html#concept_board

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
	"github.com/BlakeLiAFK/akshare/stock_a"
)

func main() {
	data, err := stock_a.StockBoardConceptNameEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockIndividualFundFlowRank

**描述**: StockIndividualFundFlowRank 东方财富网-数据中心-资金流向-排名

**数据源**: https://data.eastmoney.com/zjlx/detail.html

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| indicator | string | 指标类型 |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/stock_a"
)

func main() {
	data, err := stock_a.StockIndividualFundFlowRank("")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockZhASpotEm

**描述**: StockZhASpotEm 东方财富网-沪深京 A 股-实时行情

**数据源**: https://quote.eastmoney.com/center/gridlist.html#hs_a_board

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
	"github.com/BlakeLiAFK/akshare/stock_a"
)

func main() {
	data, err := stock_a.StockZhASpotEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

