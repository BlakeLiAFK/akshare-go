# index 模块

> 本模块共有 105 个接口

## 目录

- [DrewryWCIIndex](#drewrywciindex)
- [GetZHIndexPageCount](#getzhindexpagecount)
- [IndexAiCx](#indexaicx)
- [IndexAllCNI](#indexallcni)
- [IndexAnalysisDailySW](#indexanalysisdailysw)
- [IndexAnalysisMonthlySW](#indexanalysismonthlysw)
- [IndexAnalysisWeekMonthSW](#indexanalysisweekmonthsw)
- [IndexAnalysisWeeklySW](#indexanalysisweeklysw)
- [IndexAwprCx](#indexawprcx)
- [IndexBeiCx](#indexbeicx)
- [IndexBiCx](#indexbicx)
- [IndexCSIndexAll](#indexcsindexall)
- [IndexCciCx](#indexccicx)
- [IndexCiCx](#indexcicx)
- [IndexCodeIDMapEM](#indexcodeidmapem)
- [IndexComponentSW](#indexcomponentsw)
- [IndexDeiCx](#indexdeicx)
- [IndexDetailCNI](#indexdetailcni)
- [IndexDetailHistAdjustCNI](#indexdetailhistadjustcni)
- [IndexDetailHistCNI](#indexdetailhistcni)
- [IndexERI](#indexeri)
- [IndexEri](#indexeri)
- [IndexFiCx](#indexficx)
- [IndexGlobalHistEM](#indexglobalhistem)
- [IndexGlobalHistSina](#indexglobalhistsina)
- [IndexGlobalNameTable](#indexglobalnametable)
- [IndexGlobalSpotEM](#indexglobalspotem)
- [IndexHistCNI](#indexhistcni)
- [IndexHistFundSW](#indexhistfundsw)
- [IndexHistSW](#indexhistsw)
- [IndexHogSpotPrice](#indexhogspotprice)
- [IndexIiCx](#indexiicx)
- [IndexInnerQuoteSugarMsweet](#indexinnerquotesugarmsweet)
- [IndexKqFashion](#indexkqfashion)
- [IndexKqFzForeignTrade](#indexkqfzforeigntrade)
- [IndexKqFzPrice](#indexkqfzprice)
- [IndexKqFzProsperity](#indexkqfzprosperity)
- [IndexLiCx](#indexlicx)
- [IndexMinSW](#indexminsw)
- [IndexNeawCx](#indexneawcx)
- [IndexNeeiCx](#indexneeicx)
- [IndexNeiCx](#indexneicx)
- [IndexNewsSentimentScope](#indexnewssentimentscope)
- [IndexOption1000indexMinQvix](#indexoption1000indexminqvix)
- [IndexOption1000indexQvix](#indexoption1000indexqvix)
- [IndexOption100etfMinQvix](#indexoption100etfminqvix)
- [IndexOption100etfQvix](#indexoption100etfqvix)
- [IndexOption300etfMinQvix](#indexoption300etfminqvix)
- [IndexOption300etfQvix](#indexoption300etfqvix)
- [IndexOption300indexMinQvix](#indexoption300indexminqvix)
- [IndexOption300indexQvix](#indexoption300indexqvix)
- [IndexOption500etfMinQvix](#indexoption500etfminqvix)
- [IndexOption500etfQvix](#indexoption500etfqvix)
- [IndexOption50etfMinQvix](#indexoption50etfminqvix)
- [IndexOption50etfQvix](#indexoption50etfqvix)
- [IndexOption50indexMinQvix](#indexoption50indexminqvix)
- [IndexOption50indexQvix](#indexoption50indexqvix)
- [IndexOptionCybMinQvix](#indexoptioncybminqvix)
- [IndexOptionCybQvix](#indexoptioncybqvix)
- [IndexOptionKcbMinQvix](#indexoptionkcbminqvix)
- [IndexOptionKcbQvix](#indexoptionkcbqvix)
- [IndexOuterQuoteSugarMsweet](#indexouterquotesugarmsweet)
- [IndexPmiComCx](#indexpmicomcx)
- [IndexPmiManCx](#indexpmimancx)
- [IndexPmiSerCx](#indexpmisercx)
- [IndexPriceCFLP](#indexpricecflp)
- [IndexPriceCflp](#indexpricecflp)
- [IndexQliCx](#indexqlicx)
- [IndexRealtimeFundSW](#indexrealtimefundsw)
- [IndexRealtimeSW](#indexrealtimesw)
- [IndexSiCx](#indexsicx)
- [IndexStockConsCSIndex](#indexstockconscsindex)
- [IndexStockConsSina](#indexstockconssina)
- [IndexStockConsWeightCSI](#indexstockconsweightcsi)
- [IndexStockInfo](#indexstockinfo)
- [IndexSugarMSweet](#indexsugarmsweet)
- [IndexSugarMsweet](#indexsugarmsweet)
- [IndexTiCx](#indexticx)
- [IndexUSSpotSina](#indexusspotsina)
- [IndexUSStockSina](#indexusstocksina)
- [IndexVolumeCFLP](#indexvolumecflp)
- [IndexVolumeCflp](#indexvolumecflp)
- [IndexYW](#indexyw)
- [IndexYwMonthPrice](#indexywmonthprice)
- [IndexYwProsperity](#indexywprosperity)
- [IndexYwWeekPrice](#indexywweekprice)
- [IndexZHAHist](#indexzhahist)
- [IndexZHAHistMinEM](#indexzhahistminem)
- [MustParseDate](#mustparsedate)
- [SWIndexFirstLevel](#swindexfirstlevel)
- [SWIndexSecondLevel](#swindexsecondlevel)
- [SWIndexThirdCons](#swindexthirdcons)
- [SWIndexThirdLevel](#swindexthirdlevel)
- [SpotGoods](#spotgoods)
- [StockHKIndexDailyEM](#stockhkindexdailyem)
- [StockHKIndexDailySina](#stockhkindexdailysina)
- [StockHKIndexSpotEM](#stockhkindexspotem)
- [StockHKIndexSpotSina](#stockhkindexspotsina)
- [StockZHIndexDaily](#stockzhindexdaily)
- [StockZHIndexSpotEM](#stockzhindexspotem)
- [StockZHIndexSpotSina](#stockzhindexspotsina)
- [StockZhIndexDailyEM](#stockzhindexdailyem)
- [StockZhIndexHistCSIndex](#stockzhindexhistcsindex)
- [StockZhIndexSpotEM](#stockzhindexspotem)
- [StockZhIndexValueCSIndex](#stockzhindexvaluecsindex)

---

## DrewryWCIIndex

**描述**: DrewryWCIIndex Drewry 集装箱指数

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.DrewryWCIIndex("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## GetZHIndexPageCount

**描述**: GetZHIndexPageCount 获取新浪指数总页数

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.GetZHIndexPageCount()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexAiCx

**描述**: IndexAiCx 财新数据-指数报告-AI策略指数

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexAiCx()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexAllCNI

**描述**: IndexAllCNI 国证指数-最近交易日的所有指数

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexAllCNI()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexAnalysisDailySW

**描述**: IndexAnalysisDailySW 申万指数分析-日报告

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |
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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexAnalysisDailySW("000001", "20230101", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexAnalysisMonthlySW

**描述**: IndexAnalysisMonthlySW 申万指数分析-月报告

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |
| date | string | 日期，格式：YYYYMMDD |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexAnalysisMonthlySW("000001", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexAnalysisWeekMonthSW

**描述**: IndexAnalysisWeekMonthSW 申万周/月报表日期序列

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexAnalysisWeekMonthSW("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexAnalysisWeeklySW

**描述**: IndexAnalysisWeeklySW 申万指数分析-周报告

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |
| date | string | 日期，格式：YYYYMMDD |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexAnalysisWeeklySW("000001", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexAwprCx

**描述**: IndexAwprCx 财新数据-指数报告-新经济入职工资溢价水平

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexAwprCx()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexBeiCx

**描述**: IndexBeiCx 财新数据-指数报告-基石经济指数

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexBeiCx()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexBiCx

**描述**: IndexBiCx 财新数据-指数报告-基础指数

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexBiCx()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexCSIndexAll

**描述**: IndexCSIndexAll 中证指数网站-指数列表

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexCSIndexAll()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexCciCx

**描述**: IndexCciCx 财新数据-指数报告-大宗商品指数

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexCciCx()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexCiCx

**描述**: IndexCiCx 财新数据-指数报告-资本投入指数

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexCiCx()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexCodeIDMapEM

**描述**: IndexCodeIDMapEM 获取东方财富指数代码ID映射

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexCodeIDMapEM()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexComponentSW

**描述**: IndexComponentSW 申万指数成份股

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexComponentSW("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexDeiCx

**描述**: IndexDeiCx 财新数据-指数报告-数字经济指数

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexDeiCx()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexDetailCNI

**描述**: IndexDetailCNI 国证指数-样本详情-指定日期的样本成份

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexDetailCNI("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexDetailHistAdjustCNI

**描述**: IndexDetailHistAdjustCNI 国证指数-样本详情-历史调样

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexDetailHistAdjustCNI("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexDetailHistCNI

**描述**: IndexDetailHistCNI 国证指数-样本详情-历史样本

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexDetailHistCNI("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexERI

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexERI("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexEri

**描述**: IndexEri 浙江省排污权交易指数

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexEri("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexFiCx

**描述**: IndexFiCx 财新数据-指数报告-融合指数

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexFiCx()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexGlobalHistEM

**描述**: IndexGlobalHistEM 东方财富网-全球指数历史行情

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexGlobalHistEM("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexGlobalHistSina

**描述**: IndexGlobalHistSina 新浪财经-全球指数历史行情

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexGlobalHistSina("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexGlobalNameTable

**描述**: IndexGlobalNameTable 获取新浪全球指数名称代码映射表

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexGlobalNameTable()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexGlobalSpotEM

**描述**: IndexGlobalSpotEM 东方财富网-全球指数实时行情

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexGlobalSpotEM()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexHistCNI

**描述**: IndexHistCNI 国证指数历史行情数据

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |
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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexHistCNI("000001", "20230101", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexHistFundSW

**描述**: IndexHistFundSW 申万基金指数历史行情

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexHistFundSW("000001", "daily")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexHistSW

**描述**: IndexHistSW 申万指数历史行情

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexHistSW("000001", "daily")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexHogSpotPrice

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexHogSpotPrice()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexIiCx

**描述**: IndexIiCx 财新数据-指数报告-产业指数

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexIiCx()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexInnerQuoteSugarMsweet

**描述**: IndexInnerQuoteSugarMsweet 配额内进口糖估算指数

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexInnerQuoteSugarMsweet()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexKqFashion

**描述**: IndexKqFashion 柯桥时尚指数

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexKqFashion("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexKqFzForeignTrade

**描述**: IndexKqFzForeignTrade 柯桥纺织外贸指数

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexKqFzForeignTrade()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexKqFzPrice

**描述**: IndexKqFzPrice 柯桥纺织价格指数

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexKqFzPrice()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexKqFzProsperity

**描述**: IndexKqFzProsperity 柯桥纺织景气指数

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexKqFzProsperity()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexLiCx

**描述**: IndexLiCx 财新数据-指数报告-劳动力投入指数

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexLiCx()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexMinSW

**描述**: IndexMinSW 申万指数分时数据

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexMinSW("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexNeawCx

**描述**: IndexNeawCx 财新数据-指数报告-新经济行业入职平均工资水平

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexNeawCx()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexNeeiCx

**描述**: IndexNeeiCx 财新数据-指数报告-新动能指数

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexNeeiCx()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexNeiCx

**描述**: IndexNeiCx 财新数据-指数报告-中国新经济指数

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexNeiCx()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexNewsSentimentScope

**描述**: IndexNewsSentimentScope 数库-A股新闻情绪指数

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexNewsSentimentScope()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexOption1000indexMinQvix

**描述**: IndexOption1000indexMinQvix 中证1000股指 期权波动率指数 QVIX 分时

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexOption1000indexMinQvix()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexOption1000indexQvix

**描述**: IndexOption1000indexQvix 中证1000股指 期权波动率指数 QVIX

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexOption1000indexQvix()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexOption100etfMinQvix

**描述**: IndexOption100etfMinQvix 深证100ETF 期权波动率指数 QVIX 分时

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexOption100etfMinQvix()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexOption100etfQvix

**描述**: IndexOption100etfQvix 深证100ETF 期权波动率指数 QVIX

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexOption100etfQvix()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexOption300etfMinQvix

**描述**: IndexOption300etfMinQvix 300ETF 期权波动率指数 QVIX 分时

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexOption300etfMinQvix()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexOption300etfQvix

**描述**: IndexOption300etfQvix 300ETF 期权波动率指数 QVIX

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexOption300etfQvix()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexOption300indexMinQvix

**描述**: IndexOption300indexMinQvix 中证300股指 期权波动率指数 QVIX 分时

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexOption300indexMinQvix()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexOption300indexQvix

**描述**: IndexOption300indexQvix 中证300股指 期权波动率指数 QVIX

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexOption300indexQvix()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexOption500etfMinQvix

**描述**: IndexOption500etfMinQvix 500ETF 期权波动率指数 QVIX 分时

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexOption500etfMinQvix()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexOption500etfQvix

**描述**: IndexOption500etfQvix 500ETF 期权波动率指数 QVIX

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexOption500etfQvix()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexOption50etfMinQvix

**描述**: IndexOption50etfMinQvix 50ETF 期权波动率指数 QVIX 分时

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexOption50etfMinQvix()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexOption50etfQvix

**描述**: IndexOption50etfQvix 50ETF 期权波动率指数 QVIX

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexOption50etfQvix()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexOption50indexMinQvix

**描述**: IndexOption50indexMinQvix 上证50股指 期权波动率指数 QVIX 分时

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexOption50indexMinQvix()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexOption50indexQvix

**描述**: IndexOption50indexQvix 上证50股指 期权波动率指数 QVIX

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexOption50indexQvix()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexOptionCybMinQvix

**描述**: IndexOptionCybMinQvix 创业板 期权波动率指数 QVIX 分时

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexOptionCybMinQvix()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexOptionCybQvix

**描述**: IndexOptionCybQvix 创业板 期权波动率指数 QVIX

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexOptionCybQvix()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexOptionKcbMinQvix

**描述**: IndexOptionKcbMinQvix 科创板 期权波动率指数 QVIX 分时

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexOptionKcbMinQvix()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexOptionKcbQvix

**描述**: IndexOptionKcbQvix 科创板 期权波动率指数 QVIX

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexOptionKcbQvix()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexOuterQuoteSugarMsweet

**描述**: IndexOuterQuoteSugarMsweet 配额外进口糖估算指数

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexOuterQuoteSugarMsweet()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexPmiComCx

**描述**: IndexPmiComCx 财新数据-指数报告-综合PMI

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexPmiComCx()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexPmiManCx

**描述**: IndexPmiManCx 财新数据-指数报告-制造业PMI

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexPmiManCx()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexPmiSerCx

**描述**: IndexPmiSerCx 财新数据-指数报告-服务业PMI

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexPmiSerCx()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexPriceCFLP

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexPriceCFLP("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexPriceCflp

**描述**: IndexPriceCflp 中国公路物流运价指数

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexPriceCflp("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexQliCx

**描述**: IndexQliCx 财新数据-指数报告-高质量因子指数

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexQliCx()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexRealtimeFundSW

**描述**: IndexRealtimeFundSW 申万基金指数实时行情

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexRealtimeFundSW("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexRealtimeSW

**描述**: IndexRealtimeSW 申万指数实时行情

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexRealtimeSW("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexSiCx

**描述**: IndexSiCx 财新数据-指数报告-溢出指数

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexSiCx()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexStockConsCSIndex

**描述**: IndexStockConsCSIndex 中证指数网站-成份股目录

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexStockConsCSIndex("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexStockConsSina

**描述**: IndexStockConsSina 新浪-股票指数成份股

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexStockConsSina("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexStockConsWeightCSI

**描述**: IndexStockConsWeightCSI 中证指数网站-样本权重

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexStockConsWeightCSI("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexStockInfo

**描述**: IndexStockInfo 聚宽-指数数据-指数列表

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexStockInfo()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexSugarMSweet

**描述**: IndexSugarMSweet 沐甜科技数据中心-中国食糖指数

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexSugarMSweet()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexSugarMsweet

**描述**: IndexSugarMsweet 中国食糖指数

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexSugarMsweet()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexTiCx

**描述**: IndexTiCx 财新数据-指数报告-科技投入指数

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexTiCx()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexUSSpotSina

**描述**: IndexUSSpotSina 新浪财经-美股指数实时行情

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexUSSpotSina()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexUSStockSina

**描述**: IndexUSStockSina 新浪财经-美股指数历史行情

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexUSStockSina("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexVolumeCFLP

**描述**: IndexVolumeCFLP 中国公路物流运量指数

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexVolumeCFLP("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexVolumeCflp

**描述**: IndexVolumeCflp 中国公路物流运量指数

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexVolumeCflp("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexYW

**描述**: IndexYW 义乌小商品指数

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexYW("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexYwMonthPrice

**描述**: IndexYwMonthPrice 义乌小商品月价格指数

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexYwMonthPrice()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexYwProsperity

**描述**: IndexYwProsperity 义乌小商品月景气指数

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexYwProsperity()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexYwWeekPrice

**描述**: IndexYwWeekPrice 义乌小商品周价格指数

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexYwWeekPrice()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexZHAHist

**描述**: IndexZHAHist 获取中国股票指数历史行情数据

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |
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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexZHAHist("000001", "daily", "20230101", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IndexZHAHistMinEM

**描述**: IndexZHAHistMinEM 获取中国指数分时行情数据

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |
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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.IndexZHAHistMinEM("000001", "daily", "20230101", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MustParseDate

**描述**: MustParseDate 解析日期字符串，支持多种格式

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| s | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.MustParseDate("")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## SWIndexFirstLevel

**描述**: SWIndexFirstLevel 乐咕乐股-申万一级行业分类

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.SWIndexFirstLevel()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## SWIndexSecondLevel

**描述**: SWIndexSecondLevel 乐咕乐股-申万二级行业分类

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.SWIndexSecondLevel()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## SWIndexThirdCons

**描述**: SWIndexThirdCons 乐咕乐股-申万三级行业成份股

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.SWIndexThirdCons("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## SWIndexThirdLevel

**描述**: SWIndexThirdLevel 乐咕乐股-申万三级行业分类

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.SWIndexThirdLevel()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## SpotGoods

**描述**: SpotGoods 新浪财经-商品现货价格指数

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.SpotGoods("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHKIndexDailyEM

**描述**: StockHKIndexDailyEM 东方财富网-港股指数历史行情

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.StockHKIndexDailyEM("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHKIndexDailySina

**描述**: StockHKIndexDailySina 新浪财经-港股指数历史行情

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.StockHKIndexDailySina("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHKIndexSpotEM

**描述**: StockHKIndexSpotEM 东方财富网-港股指数实时行情

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.StockHKIndexSpotEM()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHKIndexSpotSina

**描述**: StockHKIndexSpotSina 新浪财经-港股指数实时行情

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.StockHKIndexSpotSina()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockZHIndexDaily

**描述**: StockZHIndexDaily 东方财富网-指数日K线数据

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |
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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.StockZHIndexDaily("000001", "20230101", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockZHIndexSpotEM

**描述**: StockZHIndexSpotEM 东方财富网-中国指数实时行情

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.StockZHIndexSpotEM("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockZHIndexSpotSina

**描述**: StockZHIndexSpotSina 新浪财经-中国指数实时行情

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.StockZHIndexSpotSina()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockZhIndexDailyEM

**描述**: StockZhIndexDailyEM 东方财富-股票指数日线数据

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |
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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.StockZhIndexDailyEM("000001", "20230101", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockZhIndexHistCSIndex

**描述**: StockZhIndexHistCSIndex 中证指数历史行情

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |
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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.StockZhIndexHistCSIndex("000001", "20230101", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockZhIndexSpotEM

**描述**: StockZhIndexSpotEM 东方财富-沪深京指数实时行情

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.StockZhIndexSpotEM("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockZhIndexValueCSIndex

**描述**: StockZhIndexValueCSIndex 中证指数估值数据

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
	"github.com/BlakeLiAFK/akshare/index"
)

func main() {
	data, err := index.StockZhIndexValueCSIndex("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

