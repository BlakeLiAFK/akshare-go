# stock 模块

> 本模块共有 167 个接口

## 目录

- [BuildURL](#buildurl)
- [FormatStockCode](#formatstockcode)
- [GetPreviousTradingDay](#getprevioustradingday)
- [GetStockName](#getstockname)
- [IsTradingDay](#istradingday)
- [ParseStockCode](#parsestockcode)
- [StockAllotmentCninfo](#stockallotmentcninfo)
- [StockAskBidEm](#stockaskbidem)
- [StockBidAskEm](#stockbidaskem)
- [StockBoardConceptConsEm](#stockboardconceptconsem)
- [StockBoardConceptHistEm](#stockboardconcepthistem)
- [StockBoardConceptHistMinEm](#stockboardconcepthistminem)
- [StockBoardConceptNameEm](#stockboardconceptnameem)
- [StockBoardConceptSpotEm](#stockboardconceptspotem)
- [StockBoardIndustryConsEm](#stockboardindustryconsem)
- [StockBoardIndustryHistEm](#stockboardindustryhistem)
- [StockBoardIndustryHistMinEm](#stockboardindustryhistminem)
- [StockBoardIndustryNameEm](#stockboardindustrynameem)
- [StockBoardIndustrySpotEm](#stockboardindustryspotem)
- [StockBusinessEm](#stockbusinessem)
- [StockCgEquityMortgageCninfo](#stockcgequitymortgagecninfo)
- [StockCgEquityMortgageEm](#stockcgequitymortgageem)
- [StockCgGuaranteeCninfo](#stockcgguaranteecninfo)
- [StockCgGuaranteeEm](#stockcgguaranteeem)
- [StockCgLawsuitCninfo](#stockcglawsuitcninfo)
- [StockCgLawsuitEm](#stockcglawsuitem)
- [StockDividendCninfo](#stockdividendcninfo)
- [StockDtPoolEm](#stockdtpoolem)
- [StockDzjyGgMx](#stockdzjyggmx)
- [StockDzjyHygtj](#stockdzjyhygtj)
- [StockDzjyHyyybtj](#stockdzjyhyyybtj)
- [StockDzjyMrTj](#stockdzjymrtj)
- [StockDzjyMrmx](#stockdzjymrmx)
- [StockDzjySctj](#stockdzjysctj)
- [StockDzjyYybPm](#stockdzjyyybpm)
- [StockFundEm](#stockfundem)
- [StockFundFlowEm](#stockfundflowem)
- [StockFundHoldEm](#stockfundholdem)
- [StockFundHoldRankEm](#stockfundholdrankem)
- [StockGsrlEm](#stockgsrlem)
- [StockGsrlGsdtEm](#stockgsrlgsdtem)
- [StockHkCompanyProfileEm](#stockhkcompanyprofileem)
- [StockHkComparisonEm](#stockhkcomparisonem)
- [StockHkDailyEm](#stockhkdailyem)
- [StockHkDividendPayoutEm](#stockhkdividendpayoutem)
- [StockHkFamousEm](#stockhkfamousem)
- [StockHkFamousSina](#stockhkfamoussina)
- [StockHkFhpxDetailThs](#stockhkfhpxdetailths)
- [StockHkFhpxThs](#stockhkfhpxths)
- [StockHkFinancialIndicatorEm](#stockhkfinancialindicatorem)
- [StockHkGrowthComparisonEm](#stockhkgrowthcomparisonem)
- [StockHkHistSina](#stockhkhistsina)
- [StockHkHotRankEm](#stockhkhotrankem)
- [StockHkMinuteSina](#stockhkminutesina)
- [StockHkScaleComparisonEm](#stockhkscalecomparisonem)
- [StockHkSecurityProfileEm](#stockhksecurityprofileem)
- [StockHkSpotEm](#stockhkspotem)
- [StockHkSpotSina](#stockhkspotsina)
- [StockHkValuationComparisonEm](#stockhkvaluationcomparisonem)
- [StockHoldChangeCninfo](#stockholdchangecninfo)
- [StockHoldControlCninfo](#stockholdcontrolcninfo)
- [StockHoldControlEm](#stockholdcontrolem)
- [StockHoldManagementDetailCninfo](#stockholdmanagementdetailcninfo)
- [StockHoldManagementPersonEm](#stockholdmanagementpersonem)
- [StockHoldNumCninfo](#stockholdnumcninfo)
- [StockHoldNumCninfoDetail](#stockholdnumcninfodetail)
- [StockHotKeywordEm](#stockhotkeywordem)
- [StockHotRankDetailEm](#stockhotrankdetailem)
- [StockHotRankDetailRealtimeEm](#stockhotrankdetailrealtimeem)
- [StockHotRankEm](#stockhotrankem)
- [StockHotRankLatestEm](#stockhotranklatestem)
- [StockHotRankRelateEm](#stockhotrankrelateem)
- [StockHotSearchBaidu](#stockhotsearchbaidu)
- [StockHotUpEm](#stockhotupem)
- [StockHsgtBoardRank](#stockhsgtboardrank)
- [StockHsgtHoldStock](#stockhsgtholdstock)
- [StockHsgtNorthNetFlowIn](#stockhsgtnorthnetflowin)
- [StockIndividualFundFlowEm](#stockindividualfundflowem)
- [StockIndividualFundFlowRankEm](#stockindividualfundflowrankem)
- [StockIndividualInfoEm](#stockindividualinfoem)
- [StockIndustryCategory](#stockindustrycategory)
- [StockIndustryCninfo](#stockindustrycninfo)
- [StockIndustryConst](#stockindustryconst)
- [StockIndustryPe](#stockindustrype)
- [StockIndustryPeCninfo](#stockindustrypecninfo)
- [StockIndustryPeRatioCninfo](#stockindustryperatiocninfo)
- [StockIndustrySw](#stockindustrysw)
- [StockIndustrySwConstituent](#stockindustryswconstituent)
- [StockIndustrySwDaily](#stockindustryswdaily)
- [StockIndustrySwIndex](#stockindustryswindex)
- [StockIndustrySwSpot](#stockindustryswspot)
- [StockInfoACodeNameEm](#stockinfoacodenameem)
- [StockInfoBjNameCode](#stockinfobjnamecode)
- [StockInfoChangeName](#stockinfochangename)
- [StockInfoEm](#stockinfoem)
- [StockInfoEmList](#stockinfoemlist)
- [StockInfoShDelist](#stockinfoshdelist)
- [StockInfoShNameCode](#stockinfoshnamecode)
- [StockInfoSzChangeName](#stockinfoszchangename)
- [StockInfoSzDelist](#stockinfoszdelist)
- [StockInfoSzNameCode](#stockinfosznamecode)
- [StockIntradayEm](#stockintradayem)
- [StockIntradaySina](#stockintradaysina)
- [StockIpoSummaryCninfo](#stockiposummarycninfo)
- [StockJsWeiboNlpTime](#stockjsweibonlptime)
- [StockJsWeiboReport](#stockjsweiboreport)
- [StockKcbReport](#stockkcbreport)
- [StockMainFundFlow](#stockmainfundflow)
- [StockMarketActivity](#stockmarketactivity)
- [StockMarketFundFlow](#stockmarketfundflow)
- [StockMarketFundFlowEm](#stockmarketfundflowem)
- [StockMarketSummary](#stockmarketsummary)
- [StockNewCninfo](#stocknewcninfo)
- [StockNewGhCninfo](#stocknewghcninfo)
- [StockNewIpoCninfo](#stocknewipocninfo)
- [StockNewsCx](#stocknewscx)
- [StockNewsEm](#stocknewsem)
- [StockProfileCninfo](#stockprofilecninfo)
- [StockProfileDetailEm](#stockprofiledetailem)
- [StockProfileEm](#stockprofileem)
- [StockQsJyEm](#stockqsjyem)
- [StockRankForecast](#stockrankforecast)
- [StockRankForecastCninfo](#stockrankforecastcninfo)
- [StockRepurchaseEm](#stockrepurchaseem)
- [StockSectorCodeNameEm](#stocksectorcodenameem)
- [StockShareChangeCninfo](#stocksharechangecninfo)
- [StockShareChangesEm](#stocksharechangesem)
- [StockShareHoldEm](#stockshareholdem)
- [StockSseSummary](#stockssesummary)
- [StockStopEm](#stockstopem)
- [StockSzseSummary](#stockszsesummary)
- [StockUsDailyEm](#stockusdailyem)
- [StockUsFamousEm](#stockusfamousem)
- [StockUsFamousSina](#stockusfamoussina)
- [StockUsHistSina](#stockushistsina)
- [StockUsJs](#stockusjs)
- [StockUsJsSina](#stockusjssina)
- [StockUsPink](#stockuspink)
- [StockUsPinkSina](#stockuspinksina)
- [StockUsSpotEm](#stockusspotem)
- [StockUsSpotSina](#stockusspotsina)
- [StockWeiboNlp](#stockweibonlp)
- [StockXqComments](#stockxqcomments)
- [StockXqHot](#stockxqhot)
- [StockZbPoolEm](#stockzbpoolem)
- [StockZhAHist](#stockzhahist)
- [StockZhAHistSina](#stockzhahistsina)
- [StockZhAMinuteSina](#stockzhaminutesina)
- [StockZhASpotEm](#stockzhaspotem)
- [StockZhASpotSina](#stockzhaspotsina)
- [StockZhATickSina](#stockzhaticksina)
- [StockZhATickTx](#stockzhaticktx)
- [StockZhAhDaily](#stockzhahdaily)
- [StockZhAhName](#stockzhahname)
- [StockZhAhSpot](#stockzhahspot)
- [StockZhAhTx](#stockzhahtx)
- [StockZhBHistSina](#stockzhbhistsina)
- [StockZhBSpotSina](#stockzhbspotsina)
- [StockZhComparisonEm](#stockzhcomparisonem)
- [StockZhDupontComparisonEm](#stockzhdupontcomparisonem)
- [StockZhGrowthComparisonEm](#stockzhgrowthcomparisonem)
- [StockZhKcbHistSina](#stockzhkcbhistsina)
- [StockZhKcbSpotSina](#stockzhkcbspotsina)
- [StockZhScaleComparisonEm](#stockzhscalecomparisonem)
- [StockZhValuationComparisonEm](#stockzhvaluationcomparisonem)
- [StockZtPoolEm](#stockztpoolem)
- [ValidateStockCode](#validatestockcode)

---

## BuildURL

**描述**: BuildURL 构建完整URL

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| baseURL | string | - |
| params | map[string]string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.BuildURL("", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FormatStockCode

**描述**: FormatStockCode 格式化股票代码

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |
| market | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.FormatStockCode("000001", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## GetPreviousTradingDay

**描述**: GetPreviousTradingDay 获取前一个交易日

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.GetPreviousTradingDay("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## GetStockName

**描述**: GetStockName 获取股票名称（简化实现）

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |
| market | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.GetStockName("000001", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## IsTradingDay

**描述**: IsTradingDay 判断是否为交易日

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.IsTradingDay("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## ParseStockCode

**描述**: ParseStockCode 解析股票代码

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| code | string | 股票/基金代码 |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.ParseStockCode("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockAllotmentCninfo

**描述**: StockAllotmentCninfo 巨潮资讯-配股

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockAllotmentCninfo()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockAskBidEm

**描述**: StockAskBidEm 东方财富-买卖盘数据

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockAskBidEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockBidAskEm

**描述**: StockBidAskEm 东方财富-行情报价

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockBidAskEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockBoardConceptConsEm

**描述**: StockBoardConceptConsEm 获取概念板块成分股（东方财富数据源）

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| boardCode | string | 股票/基金代码 |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockBoardConceptConsEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockBoardConceptHistEm

**描述**: StockBoardConceptHistEm 东方财富-概念板块-历史行情

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |
| period | string | 周期：daily/weekly/monthly |
| startDate | string | 开始日期，格式：YYYYMMDD |
| endDate | string | 结束日期，格式：YYYYMMDD |
| adjust | string | 复权类型：qfq(前复权)/hfq(后复权)/空字符串(不复权) |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockBoardConceptHistEm("000001", "daily", "20230101", "20230101", "qfq")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockBoardConceptHistMinEm

**描述**: StockBoardConceptHistMinEm 东方财富-概念板块-分钟行情

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockBoardConceptHistMinEm("000001", "daily")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockBoardConceptNameEm

**描述**: StockBoardConceptNameEm 获取概念板块名称列表（东方财富数据源）

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockBoardConceptNameEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockBoardConceptSpotEm

**描述**: StockBoardConceptSpotEm 获取概念板块实时行情（东方财富数据源）

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockBoardConceptSpotEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockBoardIndustryConsEm

**描述**: StockBoardIndustryConsEm 获取行业板块成分股（东方财富数据源）

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| boardCode | string | 股票/基金代码 |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockBoardIndustryConsEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockBoardIndustryHistEm

**描述**: StockBoardIndustryHistEm 东方财富-行业板块-历史行情

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |
| period | string | 周期：daily/weekly/monthly |
| startDate | string | 开始日期，格式：YYYYMMDD |
| endDate | string | 结束日期，格式：YYYYMMDD |
| adjust | string | 复权类型：qfq(前复权)/hfq(后复权)/空字符串(不复权) |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockBoardIndustryHistEm("000001", "daily", "20230101", "20230101", "qfq")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockBoardIndustryHistMinEm

**描述**: StockBoardIndustryHistMinEm 东方财富-行业板块-分钟行情

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockBoardIndustryHistMinEm("000001", "daily")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockBoardIndustryNameEm

**描述**: StockBoardIndustryNameEm 获取行业板块名称列表（东方财富数据源）

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockBoardIndustryNameEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockBoardIndustrySpotEm

**描述**: StockBoardIndustrySpotEm 获取行业板块实时行情（东方财富数据源）

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockBoardIndustrySpotEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockBusinessEm

**描述**: StockBusinessEm 东方财富-主营业务

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockBusinessEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockCgEquityMortgageCninfo

**描述**: StockCgEquityMortgageCninfo 巨潮资讯-数据中心-专题统计-公司治理-股权质押

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockCgEquityMortgageCninfo("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockCgEquityMortgageEm

**描述**: StockCgEquityMortgageEm 东方财富-股权质押

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockCgEquityMortgageEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockCgGuaranteeCninfo

**描述**: StockCgGuaranteeCninfo 巨潮资讯-数据中心-专题统计-公司治理-对外担保

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockCgGuaranteeCninfo("000001", "20230101", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockCgGuaranteeEm

**描述**: StockCgGuaranteeEm 东方财富-对外担保

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockCgGuaranteeEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockCgLawsuitCninfo

**描述**: StockCgLawsuitCninfo 巨潮资讯-数据中心-专题统计-公司治理-公司诉讼

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockCgLawsuitCninfo("000001", "20230101", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockCgLawsuitEm

**描述**: StockCgLawsuitEm 东方财富-法律诉讼

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockCgLawsuitEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockDividendCninfo

**描述**: StockDividendCninfo 巨潮资讯-分红配送

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockDividendCninfo("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockDtPoolEm

**描述**: StockDtPoolEm 东方财富-跌停板池

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockDtPoolEm("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockDzjyGgMx

**描述**: StockDzjyGgMx 东方财富-大宗交易-个股明细

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockDzjyGgMx("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockDzjyHygtj

**描述**: StockDzjyHygtj 东方财富-大宗交易-活跃股统计

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockDzjyHygtj()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockDzjyHyyybtj

**描述**: StockDzjyHyyybtj 东方财富-大宗交易-活跃营业部统计

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockDzjyHyyybtj()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockDzjyMrTj

**描述**: StockDzjyMrTj 东方财富-大宗交易-每日统计

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockDzjyMrTj("20230101", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockDzjyMrmx

**描述**: StockDzjyMrmx 东方财富-大宗交易-每日明细

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockDzjyMrmx("000001", "20230101", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockDzjySctj

**描述**: StockDzjySctj 东方财富-大宗交易-市场统计

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockDzjySctj()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockDzjyYybPm

**描述**: StockDzjyYybPm 东方财富-大宗交易-营业部排名

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockDzjyYybPm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockFundEm

**描述**: StockFundEm 东方财富-股票资金流向

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockFundEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockFundFlowEm

**描述**: StockFundFlowEm 东方财富-板块资金流向

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| sector | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockFundFlowEm("")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockFundHoldEm

**描述**: StockFundHoldEm 东方财富-基金持股

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockFundHoldEm("000001", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockFundHoldRankEm

**描述**: StockFundHoldRankEm 东方财富-基金持股排名

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockFundHoldRankEm("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockGsrlEm

**描述**: StockGsrlEm 东方财富-高送转

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockGsrlEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockGsrlGsdtEm

**描述**: StockGsrlGsdtEm 东方财富网-数据中心-股市日历-公司动态

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockGsrlGsdtEm("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHkCompanyProfileEm

**描述**: StockHkCompanyProfileEm 东方财富-港股-公司概况

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockHkCompanyProfileEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHkComparisonEm

**描述**: StockHkComparisonEm 东方财富-港股对比

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbols | string | 股票/基金代码 |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockHkComparisonEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHkDailyEm

**描述**: StockHkDailyEm 获取港股历史K线（东方财富数据源）

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| code | string | 股票/基金代码 |
| startDate | string | 开始日期，格式：YYYYMMDD |
| endDate | string | 结束日期，格式：YYYYMMDD |
| adjust | string | 复权类型：qfq(前复权)/hfq(后复权)/空字符串(不复权) |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockHkDailyEm("000001", "20230101", "20230101", "qfq")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHkDividendPayoutEm

**描述**: StockHkDividendPayoutEm 东方财富-港股-派息记录

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockHkDividendPayoutEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHkFamousEm

**描述**: StockHkFamousEm 东方财富-港股知名股票

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockHkFamousEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHkFamousSina

**描述**: StockHkFamousSina 新浪港股知名股票

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockHkFamousSina("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHkFhpxDetailThs

**描述**: StockHkFhpxDetailThs 同花顺-港股-分红派息详情

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockHkFhpxDetailThs("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHkFhpxThs

**描述**: StockHkFhpxThs 同花顺-港股分红派息

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockHkFhpxThs()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHkFinancialIndicatorEm

**描述**: StockHkFinancialIndicatorEm 东方财富-港股-财务指标

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockHkFinancialIndicatorEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHkGrowthComparisonEm

**描述**: StockHkGrowthComparisonEm 东方财富-港股-成长能力对比

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockHkGrowthComparisonEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHkHistSina

**描述**: StockHkHistSina 新浪港股历史数据

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |
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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockHkHistSina("000001", "daily", "qfq")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHkHotRankEm

**描述**: StockHkHotRankEm 东方财富-港股热门排名

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockHkHotRankEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHkMinuteSina

**描述**: StockHkMinuteSina 新浪港股分时数据

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockHkMinuteSina("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHkScaleComparisonEm

**描述**: StockHkScaleComparisonEm 东方财富-港股-规模对比

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockHkScaleComparisonEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHkSecurityProfileEm

**描述**: StockHkSecurityProfileEm 东方财富-港股-证券资料

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockHkSecurityProfileEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHkSpotEm

**描述**: StockHkSpotEm 获取港股实时行情（东方财富数据源）

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockHkSpotEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHkSpotSina

**描述**: StockHkSpotSina 新浪港股实时行情

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockHkSpotSina()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHkValuationComparisonEm

**描述**: StockHkValuationComparisonEm 东方财富-港股-估值对比

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockHkValuationComparisonEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHoldChangeCninfo

**描述**: StockHoldChangeCninfo 巨潮资讯-持股变动

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockHoldChangeCninfo("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHoldControlCninfo

**描述**: StockHoldControlCninfo 巨潮资讯-实际控制人

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockHoldControlCninfo("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHoldControlEm

**描述**: StockHoldControlEm 东方财富-实际控制人

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockHoldControlEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHoldManagementDetailCninfo

**描述**: StockHoldManagementDetailCninfo 巨潮资讯-高管持股明细

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockHoldManagementDetailCninfo("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHoldManagementPersonEm

**描述**: StockHoldManagementPersonEm 东方财富-高管人员持股

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockHoldManagementPersonEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHoldNumCninfo

**描述**: StockHoldNumCninfo 巨潮资讯-股东人数

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockHoldNumCninfo("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHoldNumCninfoDetail

**描述**: StockHoldNumCninfoDetail 巨潮资讯-股东人数详细

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockHoldNumCninfoDetail("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHotKeywordEm

**描述**: StockHotKeywordEm 东方财富-热门关键词

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockHotKeywordEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHotRankDetailEm

**描述**: StockHotRankDetailEm 东方财富-热度排名详情

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockHotRankDetailEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHotRankDetailRealtimeEm

**描述**: StockHotRankDetailRealtimeEm 东方财富-热度排名实时详情

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockHotRankDetailRealtimeEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHotRankEm

**描述**: StockHotRankEm 东方财富-热门排名

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockHotRankEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHotRankLatestEm

**描述**: StockHotRankLatestEm 东方财富-最新热度排名

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockHotRankLatestEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHotRankRelateEm

**描述**: StockHotRankRelateEm 东方财富-相关热度排名

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockHotRankRelateEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHotSearchBaidu

**描述**: StockHotSearchBaidu 百度股市通热搜

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockHotSearchBaidu()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHotUpEm

**描述**: StockHotUpEm 东方财富-飙升榜

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockHotUpEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHsgtBoardRank

**描述**: StockHsgtBoardRank 沪深港通板块排行

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockHsgtBoardRank()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHsgtHoldStock

**描述**: StockHsgtHoldStock 沪深港通持股

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| market | string | - |
| indicator | string | 指标类型 |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockHsgtHoldStock("", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHsgtNorthNetFlowIn

**描述**: StockHsgtNorthNetFlowIn 沪股通/深股通资金流向-北向资金净流入

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockHsgtNorthNetFlowIn()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockIndividualFundFlowEm

**描述**: StockIndividualFundFlowEm 获取个股资金流向（东方财富数据源）

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| code | string | 股票/基金代码 |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockIndividualFundFlowEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockIndividualFundFlowRankEm

**描述**: StockIndividualFundFlowRankEm 获取资金流向排行（东方财富数据源）

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockIndividualFundFlowRankEm("")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockIndividualInfoEm

**描述**: StockIndividualInfoEm 获取个股详细信息（东方财富数据源）

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| code | string | 股票/基金代码 |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockIndividualInfoEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockIndustryCategory

**描述**: StockIndustryCategory 行业分类

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockIndustryCategory("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockIndustryCninfo

**描述**: StockIndustryCninfo 巨潮资讯-行业分类

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockIndustryCninfo("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockIndustryConst

**描述**: StockIndustryConst 行业成分股

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockIndustryConst("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockIndustryPe

**描述**: StockIndustryPe 行业市盈率

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockIndustryPe("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockIndustryPeCninfo

**描述**: StockIndustryPeCninfo 巨潮资讯-行业市盈率

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockIndustryPeCninfo("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockIndustryPeRatioCninfo

**描述**: StockIndustryPeRatioCninfo 巨潮资讯-数据中心-行业市盈率

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockIndustryPeRatioCninfo("000001", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockIndustrySw

**描述**: StockIndustrySw 申万行业分类

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockIndustrySw()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockIndustrySwConstituent

**描述**: StockIndustrySwConstituent 申万行业成分股

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockIndustrySwConstituent("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockIndustrySwDaily

**描述**: StockIndustrySwDaily 申万行业日行情

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockIndustrySwDaily("000001", "20230101", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockIndustrySwIndex

**描述**: StockIndustrySwIndex 申万行业指数列表

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockIndustrySwIndex()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockIndustrySwSpot

**描述**: StockIndustrySwSpot 申万行业实时行情

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockIndustrySwSpot()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockInfoACodeNameEm

**描述**: StockInfoACodeNameEm 获取 A 股代码名称列表（东方财富数据源）

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockInfoACodeNameEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockInfoBjNameCode

**描述**: StockInfoBjNameCode 北交所股票代码

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockInfoBjNameCode()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockInfoChangeName

**描述**: StockInfoChangeName 股票更名历史

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockInfoChangeName("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockInfoEm

**描述**: StockInfoEm 东方财富-股票信息

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockInfoEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockInfoEmList

**描述**: StockInfoEmList 东方财富-股票列表

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| market | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockInfoEmList("")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockInfoShDelist

**描述**: StockInfoShDelist 上交所退市股票

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockInfoShDelist()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockInfoShNameCode

**描述**: StockInfoShNameCode 上交所股票代码

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockInfoShNameCode("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockInfoSzChangeName

**描述**: StockInfoSzChangeName 深交所更名历史

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockInfoSzChangeName("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockInfoSzDelist

**描述**: StockInfoSzDelist 深交所退市股票

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockInfoSzDelist()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockInfoSzNameCode

**描述**: StockInfoSzNameCode 深交所股票代码

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockInfoSzNameCode("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockIntradayEm

**描述**: StockIntradayEm 获取个股日内分时数据（东方财富数据源）

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| code | string | 股票/基金代码 |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockIntradayEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockIntradaySina

**描述**: StockIntradaySina 新浪-分时数据

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockIntradaySina("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockIpoSummaryCninfo

**描述**: StockIpoSummaryCninfo 巨潮资讯-IPO汇总

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockIpoSummaryCninfo()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockJsWeiboNlpTime

**描述**: StockJsWeiboNlpTime 微博舆情时间数据

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockJsWeiboNlpTime()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockJsWeiboReport

**描述**: StockJsWeiboReport 微博舆情报告

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| timePeriod | string | 周期：daily/weekly/monthly |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockJsWeiboReport("daily")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockKcbReport

**描述**: StockKcbReport 科创板研报

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockKcbReport("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockMainFundFlow

**描述**: StockMainFundFlow 主力资金流向

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockMainFundFlow()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockMarketActivity

**描述**: StockMarketActivity 市场活跃度统计

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockMarketActivity()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockMarketFundFlow

**描述**: StockMarketFundFlow 市场资金流向

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockMarketFundFlow()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockMarketFundFlowEm

**描述**: StockMarketFundFlowEm 获取大盘资金流向（东方财富数据源）

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockMarketFundFlowEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockMarketSummary

**描述**: StockMarketSummary 市场汇总数据（沪深合计）

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockMarketSummary()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockNewCninfo

**描述**: StockNewCninfo 巨潮资讯-新股发行

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockNewCninfo()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockNewGhCninfo

**描述**: StockNewGhCninfo 巨潮资讯-新股过会数据

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockNewGhCninfo()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockNewIpoCninfo

**描述**: StockNewIpoCninfo 巨潮资讯-新股IPO数据

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockNewIpoCninfo()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockNewsCx

**描述**: StockNewsCx 财联社-电报

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockNewsCx("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockNewsEm

**描述**: StockNewsEm 东方财富-股票新闻

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockNewsEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockProfileCninfo

**描述**: StockProfileCninfo 巨潮资讯-公司档案

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockProfileCninfo("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockProfileDetailEm

**描述**: StockProfileDetailEm 东方财富-股票详细档案

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockProfileDetailEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockProfileEm

**描述**: StockProfileEm 东方财富-股票档案

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockProfileEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockQsJyEm

**描述**: StockQsJyEm 东方财富-强势股票

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockQsJyEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockRankForecast

**描述**: StockRankForecast 机构评级预测

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockRankForecast("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockRankForecastCninfo

**描述**: StockRankForecastCninfo 巨潮资讯-数据中心-评级预测

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockRankForecastCninfo("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockRepurchaseEm

**描述**: StockRepurchaseEm 东方财富-股票回购

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockRepurchaseEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockSectorCodeNameEm

**描述**: StockSectorCodeNameEm 获取板块代码名称列表（东方财富数据源）

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| sectorType | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockSectorCodeNameEm("")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockShareChangeCninfo

**描述**: StockShareChangeCninfo 巨潮资讯-股本变动

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockShareChangeCninfo("000001", "20230101", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockShareChangesEm

**描述**: StockShareChangesEm 东方财富-股本变动

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockShareChangesEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockShareHoldEm

**描述**: StockShareHoldEm 东方财富-股东持股

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockShareHoldEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockSseSummary

**描述**: StockSseSummary 上交所市场汇总数据

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockSseSummary()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockStopEm

**描述**: StockStopEm 东方财富-停牌股票

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockStopEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockSzseSummary

**描述**: StockSzseSummary 深交所市场汇总数据

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockSzseSummary()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockUsDailyEm

**描述**: StockUsDailyEm 获取美股历史K线（东方财富数据源）

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| code | string | 股票/基金代码 |
| startDate | string | 开始日期，格式：YYYYMMDD |
| endDate | string | 结束日期，格式：YYYYMMDD |
| adjust | string | 复权类型：qfq(前复权)/hfq(后复权)/空字符串(不复权) |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockUsDailyEm("000001", "20230101", "20230101", "qfq")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockUsFamousEm

**描述**: StockUsFamousEm 东方财富-美股知名股票

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockUsFamousEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockUsFamousSina

**描述**: StockUsFamousSina 新浪美股知名股票

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockUsFamousSina("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockUsHistSina

**描述**: StockUsHistSina 新浪美股历史数据

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |
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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockUsHistSina("000001", "daily", "qfq")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockUsJs

**描述**: StockUsJs 新浪美股JavaScript数据

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockUsJs("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockUsJsSina

**描述**: StockUsJsSina 新浪美股JavaScript数据

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockUsJsSina("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockUsPink

**描述**: StockUsPink 美股粉单市场

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockUsPink()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockUsPinkSina

**描述**: StockUsPinkSina 新浪美股粉单市场

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockUsPinkSina()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockUsSpotEm

**描述**: StockUsSpotEm 获取美股实时行情（东方财富数据源）

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockUsSpotEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockUsSpotSina

**描述**: StockUsSpotSina 新浪美股实时行情

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockUsSpotSina("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockWeiboNlp

**描述**: StockWeiboNlp 微博股票情感分析

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockWeiboNlp("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockXqComments

**描述**: StockXqComments 雪球-股票评论

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockXqComments("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockXqHot

**描述**: StockXqHot 雪球-热门股票

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockXqHot()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockZbPoolEm

**描述**: StockZbPoolEm 东方财富-炸板池

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockZbPoolEm("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockZhAHist

**描述**: StockZhAHist 获取 A 股历史 K 线数据

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| code | string | 股票/基金代码 |
| period | string | 周期：daily/weekly/monthly |
| startDate | string | 开始日期，格式：YYYYMMDD |
| endDate | string | 结束日期，格式：YYYYMMDD |
| adjust | string | 复权类型：qfq(前复权)/hfq(后复权)/空字符串(不复权) |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockZhAHist("000001", "daily", "20230101", "20230101", "qfq")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockZhAHistSina

**描述**: StockZhAHistSina 新浪A股历史数据

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |
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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockZhAHistSina("000001", "daily", "qfq")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockZhAMinuteSina

**描述**: StockZhAMinuteSina 新浪A股分时数据

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockZhAMinuteSina("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockZhASpotEm

**描述**: StockZhASpotEm 获取 A 股实时行情（东方财富数据源）

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockZhASpotEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockZhASpotSina

**描述**: StockZhASpotSina 新浪A股实时行情

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockZhASpotSina()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockZhATickSina

**描述**: StockZhATickSina 新浪A股逐笔数据

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockZhATickSina("000001", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockZhATickTx

**描述**: StockZhATickTx 腾讯-A股逐笔数据

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockZhATickTx("000001", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockZhAhDaily

**描述**: StockZhAhDaily A+H股历史行情

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |
| startDate | string | 开始日期，格式：YYYYMMDD |
| endDate | string | 结束日期，格式：YYYYMMDD |
| adjust | string | 复权类型：qfq(前复权)/hfq(后复权)/空字符串(不复权) |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockZhAhDaily("000001", "20230101", "20230101", "qfq")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockZhAhName

**描述**: StockZhAhName A+H股名称列表

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockZhAhName()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockZhAhSpot

**描述**: StockZhAhSpot A+H股实时行情

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockZhAhSpot()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockZhAhTx

**描述**: StockZhAhTx 腾讯-AH股对比

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockZhAhTx()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockZhBHistSina

**描述**: StockZhBHistSina 新浪B股历史数据

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |
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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockZhBHistSina("000001", "daily", "qfq")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockZhBSpotSina

**描述**: StockZhBSpotSina 新浪B股实时行情

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockZhBSpotSina()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockZhComparisonEm

**描述**: StockZhComparisonEm 东方财富-A股对比

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbols | string | 股票/基金代码 |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockZhComparisonEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockZhDupontComparisonEm

**描述**: StockZhDupontComparisonEm 东方财富-A股-个股-杜邦分析对比

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockZhDupontComparisonEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockZhGrowthComparisonEm

**描述**: StockZhGrowthComparisonEm 东方财富-A股-个股-成长性对比

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockZhGrowthComparisonEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockZhKcbHistSina

**描述**: StockZhKcbHistSina 新浪科创板历史数据

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |
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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockZhKcbHistSina("000001", "daily", "qfq")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockZhKcbSpotSina

**描述**: StockZhKcbSpotSina 新浪科创板实时行情

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockZhKcbSpotSina()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockZhScaleComparisonEm

**描述**: StockZhScaleComparisonEm 东方财富-A股-个股-规模对比

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockZhScaleComparisonEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockZhValuationComparisonEm

**描述**: StockZhValuationComparisonEm 东方财富-A股-个股-估值对比

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockZhValuationComparisonEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockZtPoolEm

**描述**: StockZtPoolEm 东方财富-涨停板池

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
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.StockZtPoolEm("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## ValidateStockCode

**描述**: ValidateStockCode 验证股票代码格式

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| code | string | 股票/基金代码 |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	data, err := stock.ValidateStockCode("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

