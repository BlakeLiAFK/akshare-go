# stock_fundamental 模块

> 本模块共有 55 个接口

## 目录

- [GetSupportedSymbolList](#getsupportedsymbollist)
- [StockAddStock](#stockaddstock)
- [StockCirculateStockHolder](#stockcirculatestockholder)
- [StockFinancialAbstract](#stockfinancialabstract)
- [StockFinancialAbstractNewThs](#stockfinancialabstractnewths)
- [StockFinancialAbstractThs](#stockfinancialabstractths)
- [StockFinancialAnalysisIndicatorEm](#stockfinancialanalysisindicatorem)
- [StockFinancialBenefitNewThs](#stockfinancialbenefitnewths)
- [StockFinancialBenefitThs](#stockfinancialbenefitths)
- [StockFinancialCashNewThs](#stockfinancialcashnewths)
- [StockFinancialCashThs](#stockfinancialcashths)
- [StockFinancialDebtNewThs](#stockfinancialdebtnewths)
- [StockFinancialDebtThs](#stockfinancialdebtths)
- [StockFinancialHkAnalysisIndicatorEm](#stockfinancialhkanalysisindicatorem)
- [StockFinancialHkReportEm](#stockfinancialhkreportem)
- [StockFinancialReportSina](#stockfinancialreportsina)
- [StockFinancialUsAnalysisIndicatorEm](#stockfinancialusanalysisindicatorem)
- [StockFinancialUsReportEm](#stockfinancialusreportem)
- [StockFundStockHolder](#stockfundstockholder)
- [StockHistoryDividend](#stockhistorydividend)
- [StockHistoryDividendDetail](#stockhistorydividenddetail)
- [StockHkProfitForecastEt](#stockhkprofitforecastet)
- [StockIndividualBasicInfoHkXq](#stockindividualbasicinfohkxq)
- [StockIndividualBasicInfoUsXq](#stockindividualbasicinfousxq)
- [StockIndividualBasicInfoXq](#stockindividualbasicinfoxq)
- [StockInstituteHold](#stockinstitutehold)
- [StockInstituteHoldDetail](#stockinstituteholddetail)
- [StockInstituteRecommend](#stockinstituterecommend)
- [StockInstituteRecommendDetail](#stockinstituterecommenddetail)
- [StockIpoDeclareEm](#stockipodeclareem)
- [StockIpoInfo](#stockipoinfo)
- [StockIpoReviewEm](#stockiporeviewem)
- [StockIpoTutorEm](#stockipotutorem)
- [StockKcbSse](#stockkcbsse)
- [StockMainStockHolder](#stockmainstockholder)
- [StockManagementChangeThs](#stockmanagementchangeths)
- [StockNoticeReport](#stocknoticereport)
- [StockProfitForecastEm](#stockprofitforecastem)
- [StockProfitForecastThs](#stockprofitforecastths)
- [StockProfitForecastThsDetail](#stockprofitforecastthsdetail)
- [StockRegisterAllEm](#stockregisterallem)
- [StockRegisterBj](#stockregisterbj)
- [StockRegisterCyb](#stockregistercyb)
- [StockRegisterDb](#stockregisterdb)
- [StockRegisterKcb](#stockregisterkcb)
- [StockRegisterSh](#stockregistersh)
- [StockRegisterSz](#stockregistersz)
- [StockRestrictedReleaseDetailEm](#stockrestrictedreleasedetailem)
- [StockRestrictedReleaseQueueEm](#stockrestrictedreleasequeueem)
- [StockRestrictedReleaseStockholderEm](#stockrestrictedreleasestockholderem)
- [StockRestrictedReleaseSummaryEm](#stockrestrictedreleasesummaryem)
- [StockShareholderChangeThs](#stockshareholderchangeths)
- [StockZhAGbjgEm](#stockzhagbjgem)
- [StockZygcEm](#stockzygcem)
- [StockZyjsThs](#stockzyjsths)

---

## GetSupportedSymbolList

**描述**: GetSupportedSymbolList 获取支持的symbol列表

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
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.GetSupportedSymbolList()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockAddStock

**描述**: StockAddStock 增发股票

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
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockAddStock("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockCirculateStockHolder

**描述**: StockCirculateStockHolder 流通股股东

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
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockCirculateStockHolder("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockFinancialAbstract

**描述**: StockFinancialAbstract 新浪财经-财务报表-关键指标

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
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockFinancialAbstract("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockFinancialAbstractNewThs

**描述**: StockFinancialAbstractNewThs 同花顺-财务指标-重要指标（新版API）

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
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
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockFinancialAbstractNewThs("000001", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockFinancialAbstractThs

**描述**: StockFinancialAbstractThs 同花顺-财务指标-主要指标（旧版）

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
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
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockFinancialAbstractThs("000001", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockFinancialAnalysisIndicatorEm

**描述**: StockFinancialAnalysisIndicatorEm 东方财富-A股-财务分析-主要指标

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
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
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockFinancialAnalysisIndicatorEm("000001", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockFinancialBenefitNewThs

**描述**: StockFinancialBenefitNewThs 同花顺-财务指标-利润表（新版API）

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
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
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockFinancialBenefitNewThs("000001", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockFinancialBenefitThs

**描述**: StockFinancialBenefitThs 同花顺-财务指标-利润表（旧版）

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
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
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockFinancialBenefitThs("000001", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockFinancialCashNewThs

**描述**: StockFinancialCashNewThs 同花顺-财务指标-现金流量表（新版API）

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
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
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockFinancialCashNewThs("000001", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockFinancialCashThs

**描述**: StockFinancialCashThs 同花顺-财务指标-现金流量表（旧版）

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
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
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockFinancialCashThs("000001", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockFinancialDebtNewThs

**描述**: StockFinancialDebtNewThs 同花顺-财务指标-资产负债表（新版API）

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
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
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockFinancialDebtNewThs("000001", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockFinancialDebtThs

**描述**: StockFinancialDebtThs 同花顺-财务指标-资产负债表（旧版）

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
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
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockFinancialDebtThs("000001", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockFinancialHkAnalysisIndicatorEm

**描述**: StockFinancialHkAnalysisIndicatorEm 东方财富-港股-财务分析-主要指标

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
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
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockFinancialHkAnalysisIndicatorEm("000001", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockFinancialHkReportEm

**描述**: StockFinancialHkReportEm 东方财富-港股-财务报表-三大报表

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| stock | string | - |
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
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockFinancialHkReportEm("", "000001", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockFinancialReportSina

**描述**: StockFinancialReportSina 新浪财经-财务报表-三大报表

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| stock | string | - |
| symbol | string | 股票/基金代码 |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockFinancialReportSina("", "000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockFinancialUsAnalysisIndicatorEm

**描述**: StockFinancialUsAnalysisIndicatorEm 东方财富-美股-财务分析-主要指标

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
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
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockFinancialUsAnalysisIndicatorEm("000001", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockFinancialUsReportEm

**描述**: StockFinancialUsReportEm 东方财富-美股-财务分析-三大报表

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| stock | string | - |
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
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockFinancialUsReportEm("", "000001", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockFundStockHolder

**描述**: StockFundStockHolder 新浪财经-股本股东-基金持股

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
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockFundStockHolder("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHistoryDividend

**描述**: StockHistoryDividend 新浪财经-历史分红

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
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockHistoryDividend()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHistoryDividendDetail

**描述**: StockHistoryDividendDetail 新浪财经-分红配股详情

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |
| indicator | string | 指标类型 |
| date | string | 日期，格式：YYYYMMDD |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockHistoryDividendDetail("000001", "", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHkProfitForecastEt

**描述**: StockHkProfitForecastEt 经济通-港股盈利预测

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
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
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockHkProfitForecastEt("000001", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockIndividualBasicInfoHkXq

**描述**: StockIndividualBasicInfoHkXq 雪球-港股公司简介

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |
| token | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockIndividualBasicInfoHkXq("000001", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockIndividualBasicInfoUsXq

**描述**: StockIndividualBasicInfoUsXq 雪球-美股公司简介

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |
| token | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockIndividualBasicInfoUsXq("000001", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockIndividualBasicInfoXq

**描述**: StockIndividualBasicInfoXq 雪球-A股公司简介

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |
| token | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockIndividualBasicInfoXq("000001", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockInstituteHold

**描述**: StockInstituteHold 新浪财经-股票-机构持股一览表

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
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockInstituteHold("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockInstituteHoldDetail

**描述**: StockInstituteHoldDetail 新浪财经-股票-机构持股详情

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| stock | string | - |
| quarter | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockInstituteHoldDetail("", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockInstituteRecommend

**描述**: StockInstituteRecommend 新浪财经-机构推荐池

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
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockInstituteRecommend("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockInstituteRecommendDetail

**描述**: StockInstituteRecommendDetail 新浪财经-机构推荐池-股票评级记录

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
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockInstituteRecommendDetail("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockIpoDeclareEm

**描述**: StockIpoDeclareEm 东方财富网-数据中心-新股申购-首发申报企业信息

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
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockIpoDeclareEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockIpoInfo

**描述**: StockIpoInfo 新浪财经-新股发行

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| stock | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockIpoInfo("")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockIpoReviewEm

**描述**: StockIpoReviewEm 东方财富网-数据中心-新股申购-过会企业信息

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
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockIpoReviewEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockIpoTutorEm

**描述**: StockIpoTutorEm 东方财富网-数据中心-新股申购-辅导备案信息

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
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockIpoTutorEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockKcbSse

**描述**: StockKcbSse 上交所科创板-股票发行

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
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockKcbSse()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockMainStockHolder

**描述**: StockMainStockHolder 新浪财经-股本股东-主要股东

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| stock | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockMainStockHolder("")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockManagementChangeThs

**描述**: StockManagementChangeThs 同花顺-公司大事-高管持股变动

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
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockManagementChangeThs("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockNoticeReport

**描述**: StockNoticeReport 东方财富网-数据中心-公告大全-沪深京 A 股公告

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
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockNoticeReport("000001", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockProfitForecastEm

**描述**: StockProfitForecastEm 东方财富网-数据中心-研究报告-盈利预测

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
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockProfitForecastEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockProfitForecastThs

**描述**: StockProfitForecastThs 同花顺-盈利预测

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
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
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockProfitForecastThs("000001", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockProfitForecastThsDetail

**描述**: StockProfitForecastThsDetail 同花顺-盈利预测（详细版）

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
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
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockProfitForecastThsDetail("000001", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockRegisterAllEm

**描述**: StockRegisterAllEm 东方财富网-IPO审核信息-全部

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
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockRegisterAllEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockRegisterBj

**描述**: StockRegisterBj 北交所注册制股票

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
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockRegisterBj()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockRegisterCyb

**描述**: StockRegisterCyb 东方财富网-IPO审核信息-创业板

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
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockRegisterCyb()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockRegisterDb

**描述**: StockRegisterDb 主板注册制股票

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
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockRegisterDb()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockRegisterKcb

**描述**: StockRegisterKcb 东方财富网-IPO审核信息-科创板

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
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockRegisterKcb()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockRegisterSh

**描述**: StockRegisterSh 上交所注册制股票

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
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockRegisterSh()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockRegisterSz

**描述**: StockRegisterSz 深交所注册制股票

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
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockRegisterSz()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockRestrictedReleaseDetailEm

**描述**: StockRestrictedReleaseDetailEm 东方财富网-限售股解禁详情

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
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
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockRestrictedReleaseDetailEm("20230101", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockRestrictedReleaseQueueEm

**描述**: StockRestrictedReleaseQueueEm 东方财富网-个股限售解禁批次

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
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockRestrictedReleaseQueueEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockRestrictedReleaseStockholderEm

**描述**: StockRestrictedReleaseStockholderEm 限售股解禁-股东详情

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
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockRestrictedReleaseStockholderEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockRestrictedReleaseSummaryEm

**描述**: StockRestrictedReleaseSummaryEm 东方财富网-限售股解禁汇总

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
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockRestrictedReleaseSummaryEm("000001", "20230101", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockShareholderChangeThs

**描述**: StockShareholderChangeThs 同花顺-公司大事-股东持股变动

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
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockShareholderChangeThs("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockZhAGbjgEm

**描述**: StockZhAGbjgEm 东方财富-A股数据-股本结构

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
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockZhAGbjgEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockZygcEm

**描述**: StockZygcEm 东方财富网-个股-主营构成

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
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockZygcEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockZyjsThs

**描述**: StockZyjsThs 同花顺-主营介绍

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
	"github.com/BlakeLiAFK/akshare/stock_fundamental"
)

func main() {
	data, err := stock_fundamental.StockZyjsThs("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

