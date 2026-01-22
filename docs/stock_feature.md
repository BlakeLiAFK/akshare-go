# stock_feature 模块

> 本模块共有 148 个接口

## 目录

- [GetCookieCsrf](#getcookiecsrf)
- [GetTokenLg](#gettokenlg)
- [StockAAllPb](#stockaallpb)
- [StockABelowNetAssetStatistics](#stockabelownetassetstatistics)
- [StockACongestionLg](#stockacongestionlg)
- [StockAGxlLg](#stockagxllg)
- [StockAHighLow](#stockahighlow)
- [StockAHighLowStatistics](#stockahighlowstatistics)
- [StockAIndicator](#stockaindicator)
- [StockAPb](#stockapb)
- [StockAPe](#stockape)
- [StockAPeAndPb](#stockapeandpb)
- [StockATtmLyr](#stockattmlyr)
- [StockAccountEm](#stockaccountem)
- [StockAccountStatisticsEm](#stockaccountstatisticsem)
- [StockAnalystDetailEm](#stockanalystdetailem)
- [StockAnalystEm](#stockanalystem)
- [StockBjASpotEm](#stockbjaspotem)
- [StockBoardChangeEm](#stockboardchangeem)
- [StockBoardConceptConstThs](#stockboardconceptconstths)
- [StockBoardConceptThs](#stockboardconceptths)
- [StockBoardIndustryConstThs](#stockboardindustryconstths)
- [StockBoardIndustryThs](#stockboardindustryths)
- [StockBuffettIndexLg](#stockbuffettindexlg)
- [StockChangesEm](#stockchangesem)
- [StockClassifyConstSina](#stockclassifyconstsina)
- [StockClassifySina](#stockclassifysina)
- [StockCommentEm](#stockcommentem)
- [StockConceptConstFutu](#stockconceptconstfutu)
- [StockConceptFutu](#stockconceptfutu)
- [StockCybASpotEm](#stockcybaspotem)
- [StockCyqDetailEm](#stockcyqdetailem)
- [StockCyqEm](#stockcyqem)
- [StockDisclosureCninfo](#stockdisclosurecninfo)
- [StockDtPoolEm](#stockdtpoolem)
- [StockDxsylEm](#stockdxsylem)
- [StockEbsLg](#stockebslg)
- [StockEsgMsciSina](#stockesgmscisina)
- [StockEsgRateSina](#stockesgratesina)
- [StockEsgRftSina](#stockesgrftsina)
- [StockEsgZdSina](#stockesgzdsina)
- [StockFhpsDetailEm](#stockfhpsdetailem)
- [StockFhpsDetailThs](#stockfhpsdetailths)
- [StockFhpsEm](#stockfhpsem)
- [StockFundFlow](#stockfundflow)
- [StockFundFlowIndividual](#stockfundflowindividual)
- [StockGddhEm](#stockgddhem)
- [StockGdfxFreeHoldingAnalyseEm](#stockgdfxfreeholdinganalyseem)
- [StockGdfxFreeHoldingDetailEm](#stockgdfxfreeholdingdetailem)
- [StockGdfxFreeHoldingStatisticsEm](#stockgdfxfreeholdingstatisticsem)
- [StockGdfxHoldingAnalyseEm](#stockgdfxholdinganalyseem)
- [StockGdfxHoldingDetailEm](#stockgdfxholdingdetailem)
- [StockGdfxHoldingStatisticsEm](#stockgdfxholdingstatisticsem)
- [StockGgcgEm](#stockggcgem)
- [StockGpzyDetailEm](#stockgpzydetailem)
- [StockGpzyEm](#stockgpzyem)
- [StockGpzyPledgeRatioEm](#stockgpzypledgeratioem)
- [StockGpzyProfileEm](#stockgpzyprofileem)
- [StockHistEm](#stockhistem)
- [StockHkGgtComponentsEm](#stockhkggtcomponentsem)
- [StockHkIndicatorEniu](#stockhkindicatoreniu)
- [StockHkValuationBaidu](#stockhkvaluationbaidu)
- [StockHotDealXq](#stockhotdealxq)
- [StockHotFollowXq](#stockhotfollowxq)
- [StockHotTweetXq](#stockhottweetxq)
- [StockHotXq](#stockhotxq)
- [StockHsgtEmFlow](#stockhsgtemflow)
- [StockHsgtExchangeRate](#stockhsgtexchangerate)
- [StockHsgtFundFlowSummaryEm](#stockhsgtfundflowsummaryem)
- [StockHsgtFundMinEm](#stockhsgtfundminem)
- [StockHsgtHoldStockEm](#stockhsgtholdstockem)
- [StockHsgtMinEm](#stockhsgtminem)
- [StockHsgtSpotEm](#stockhsgtspotem)
- [StockInfoChange](#stockinfochange)
- [StockInfoCjzcEm](#stockinfocjzcem)
- [StockInfoGlobalCls](#stockinfoglobalcls)
- [StockInfoGlobalEm](#stockinfoglobalem)
- [StockInfoGlobalFutu](#stockinfoglobalfutu)
- [StockInfoGlobalSina](#stockinfoglobalsina)
- [StockInfoGlobalThs](#stockinfoglobalths)
- [StockInnerTradeXq](#stockinnertradexq)
- [StockIrmCninfo](#stockirmcninfo)
- [StockJgdyDetailEm](#stockjgdydetailem)
- [StockJgdyEm](#stockjgdyem)
- [StockKcbASpotEm](#stockkcbaspotem)
- [StockLhYybCapital](#stocklhyybcapital)
- [StockLhYybControl](#stocklhyybcontrol)
- [StockLhYybMost](#stocklhyybmost)
- [StockLhbDetailEm](#stocklhbdetailem)
- [StockLhbGgtjSina](#stocklhbggtjsina)
- [StockLhbJgmmtjEm](#stocklhbjgmmtjem)
- [StockLhbJgstatisticEm](#stocklhbjgstatisticem)
- [StockLhbStockDetailEm](#stocklhbstockdetailem)
- [StockLhbStockStatisticEm](#stocklhbstockstatisticem)
- [StockLhbTraderstatisticEm](#stocklhbtraderstatisticem)
- [StockLhbYybDetailEm](#stocklhbyybdetailem)
- [StockLhbYybphEm](#stocklhbyybphem)
- [StockMarginAccountInfo](#stockmarginaccountinfo)
- [StockMarginEm](#stockmarginem)
- [StockMarginSse](#stockmarginsse)
- [StockMarginSzse](#stockmarginszse)
- [StockMarketActivityLegu](#stockmarketactivitylegu)
- [StockNewSpotEm](#stocknewspotem)
- [StockPankouEm](#stockpankouem)
- [StockPgEm](#stockpgem)
- [StockQbzfEm](#stockqbzfem)
- [StockQsjyEm](#stockqsjyem)
- [StockRankCxdThs](#stockrankcxdths)
- [StockRankCxflThs](#stockrankcxflths)
- [StockRankCxgThs](#stockrankcxgths)
- [StockRankCxslThs](#stockrankcxslths)
- [StockRankLjqdThs](#stockrankljqdths)
- [StockRankLjqsThs](#stockrankljqsths)
- [StockRankLxszThs](#stockranklxszths)
- [StockRankLxxdThs](#stockranklxxdths)
- [StockRankXstpThs](#stockrankxstpths)
- [StockRankXxtpThs](#stockrankxxtpths)
- [StockRankXzjpThs](#stockrankxzjpths)
- [StockReportDisclosure](#stockreportdisclosure)
- [StockReportEm](#stockreportem)
- [StockResearchReportEm](#stockresearchreportem)
- [StockSgtReferenceExchangeRateSse](#stocksgtreferenceexchangeratesse)
- [StockSgtReferenceExchangeRateSzse](#stocksgtreferenceexchangerateszse)
- [StockSgtSettlementExchangeRateSse](#stocksgtsettlementexchangeratesse)
- [StockSgtSettlementExchangeRateSzse](#stocksgtsettlementexchangerateszse)
- [StockShASpotEm](#stockshaspotem)
- [StockSnsSseinfo](#stocksnssseinfo)
- [StockSyEm](#stocksyem)
- [StockSzASpotEm](#stockszaspotem)
- [StockTfpEm](#stocktfpem)
- [StockThreeReportEm](#stockthreereportem)
- [StockUsValuationBaidu](#stockusvaluationbaidu)
- [StockValueEm](#stockvalueem)
- [StockYjbbEm](#stockyjbbem)
- [StockYjkbEm](#stockyjkbem)
- [StockYjygEm](#stockyjygem)
- [StockYzxdrEm](#stockyzxdrem)
- [StockZbPoolEm](#stockzbpoolem)
- [StockZdhtmxEm](#stockzdhtmxem)
- [StockZhAGdhs](#stockzhagdhs)
- [StockZhAGdhsDetailEm](#stockzhagdhsdetailem)
- [StockZhAHistTx](#stockzhahisttx)
- [StockZhASpotEm](#stockzhaspotem)
- [StockZhBSpotEm](#stockzhbspotem)
- [StockZhValuationBaidu](#stockzhvaluationbaidu)
- [StockZhVoteBaidu](#stockzhvotebaidu)
- [StockZtPoolEm](#stockztpoolem)
- [StockZtPoolPreviousEm](#stockztpoolpreviousem)

---

## GetCookieCsrf

**描述**: GetCookieCsrf 获取乐咕的 Cookie 和 CSRF Token

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| pageURL | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.GetCookieCsrf("")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## GetTokenLg

**描述**: GetTokenLg 生成乐咕的 token

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.GetTokenLg()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockAAllPb

**描述**: StockAAllPb 全部A股-等权重市净率、中位数市净率

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockAAllPb()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockABelowNetAssetStatistics

**描述**: StockABelowNetAssetStatistics A股破净股统计

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockABelowNetAssetStatistics()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockACongestionLg

**描述**: StockACongestionLg 乐估乐股-筹码拥挤度

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockACongestionLg()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockAGxlLg

**描述**: StockAGxlLg 乐估乐股-股息率

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockAGxlLg()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockAHighLow

**描述**: StockAHighLow A股新高新低统计

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockAHighLow("")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockAHighLowStatistics

**描述**: StockAHighLowStatistics 乐咕乐股-创新高、新低的股票数量

**数据源**: https://www.legulegu.com/stockdata/high-low-statistics

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockAHighLowStatistics("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockAIndicator

**描述**: StockAIndicator A股指标数据

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockAIndicator("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockAPb

**描述**: StockAPb A股市净率

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockAPb("")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockAPe

**描述**: StockAPe A股市盈率

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockAPe("")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockAPeAndPb

**描述**: StockAPeAndPb A股市盈率和市净率

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockAPeAndPb("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockATtmLyr

**描述**: StockATtmLyr A股滚动市盈率

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockATtmLyr()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockAccountEm

**描述**: StockAccountEm 东方财富-股票开户数

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockAccountEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockAccountStatisticsEm

**描述**: StockAccountStatisticsEm 东方财富网-数据中心-特色数据-股票账户统计

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockAccountStatisticsEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockAnalystDetailEm

**描述**: StockAnalystDetailEm 东方财富-分析师详情

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| analystId | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockAnalystDetailEm("")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockAnalystEm

**描述**: StockAnalystEm 东方财富-分析师排名

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockAnalystEm("")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockBjASpotEm

**描述**: StockBjASpotEm 东方财富网-北 A 股-实时行情

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockBjASpotEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockBoardChangeEm

**描述**: StockBoardChangeEm 东方财富-行情中心-当日板块异动详情

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockBoardChangeEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockBoardConceptConstThs

**描述**: StockBoardConceptConstThs 同花顺概念板块成分股

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockBoardConceptConstThs("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockBoardConceptThs

**描述**: StockBoardConceptThs 同花顺概念板块

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockBoardConceptThs()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockBoardIndustryConstThs

**描述**: StockBoardIndustryConstThs 同花顺行业板块成分股

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockBoardIndustryConstThs("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockBoardIndustryThs

**描述**: StockBoardIndustryThs 同花顺行业板块

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockBoardIndustryThs()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockBuffettIndexLg

**描述**: StockBuffettIndexLg 乐估乐股-底部研究-巴菲特指标

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockBuffettIndexLg()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockChangesEm

**描述**: StockChangesEm 东方财富-行情中心-盘口异动

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockChangesEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockClassifyConstSina

**描述**: StockClassifyConstSina 新浪-行业成分股

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockClassifyConstSina("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockClassifySina

**描述**: StockClassifySina 新浪-行业分类

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockClassifySina()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockCommentEm

**描述**: StockCommentEm 东方财富-千股千评

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockCommentEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockConceptConstFutu

**描述**: StockConceptConstFutu 富途-概念板块成分股

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockConceptConstFutu("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockConceptFutu

**描述**: StockConceptFutu 富途-概念板块

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockConceptFutu()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockCybASpotEm

**描述**: StockCybASpotEm 东方财富网-创业板-实时行情

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockCybASpotEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockCyqDetailEm

**描述**: StockCyqDetailEm 东方财富-筹码分布详情

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockCyqDetailEm("000001", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockCyqEm

**描述**: StockCyqEm 东方财富-筹码分布

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockCyqEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockDisclosureCninfo

**描述**: StockDisclosureCninfo 巨潮资讯-信息披露

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |
| category | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockDisclosureCninfo("000001", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockDtPoolEm

**描述**: StockDtPoolEm 东方财富网-行情中心-涨停板行情-跌停股池

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockDtPoolEm("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockDxsylEm

**描述**: StockDxsylEm 东方财富-打新收益率

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockDxsylEm("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockEbsLg

**描述**: StockEbsLg 乐估乐股-情绪择时指标

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockEbsLg()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockEsgMsciSina

**描述**: StockEsgMsciSina 新浪财经-ESG-MSCI评级

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockEsgMsciSina()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockEsgRateSina

**描述**: StockEsgRateSina 新浪财经-ESG评级

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockEsgRateSina()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockEsgRftSina

**描述**: StockEsgRftSina 新浪财经-ESG-RFT评级

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockEsgRftSina()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockEsgZdSina

**描述**: StockEsgZdSina 新浪财经-ESG-中登评级

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockEsgZdSina()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockFhpsDetailEm

**描述**: StockFhpsDetailEm 东方财富-分红配送详情

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockFhpsDetailEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockFhpsDetailThs

**描述**: StockFhpsDetailThs 同花顺-分红派息详情

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockFhpsDetailThs("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockFhpsEm

**描述**: StockFhpsEm 东方财富-分红配送

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockFhpsEm("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockFundFlow

**描述**: StockFundFlow 板块资金流向

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockFundFlow("")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockFundFlowIndividual

**描述**: StockFundFlowIndividual 个股资金流向

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockFundFlowIndividual("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockGddhEm

**描述**: StockGddhEm 东方财富网-数据中心-股东大会

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockGddhEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockGdfxFreeHoldingAnalyseEm

**描述**: StockGdfxFreeHoldingAnalyseEm 东方财富-股东分析-流通股持股分析

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockGdfxFreeHoldingAnalyseEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockGdfxFreeHoldingDetailEm

**描述**: StockGdfxFreeHoldingDetailEm 东方财富-股东分析-流通股持股明细

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockGdfxFreeHoldingDetailEm("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockGdfxFreeHoldingStatisticsEm

**描述**: StockGdfxFreeHoldingStatisticsEm 东方财富网-数据中心-股东分析-股东持股统计-十大流通股东

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockGdfxFreeHoldingStatisticsEm("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockGdfxHoldingAnalyseEm

**描述**: StockGdfxHoldingAnalyseEm 东方财富-股东分析-持股分析

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockGdfxHoldingAnalyseEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockGdfxHoldingDetailEm

**描述**: StockGdfxHoldingDetailEm 东方财富-股东分析-持股明细

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockGdfxHoldingDetailEm("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockGdfxHoldingStatisticsEm

**描述**: StockGdfxHoldingStatisticsEm 东方财富网-数据中心-股东分析-股东持股统计-十大股东

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockGdfxHoldingStatisticsEm("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockGgcgEm

**描述**: StockGgcgEm 东方财富网-数据中心-特色数据-高管持股

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockGgcgEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockGpzyDetailEm

**描述**: StockGpzyDetailEm 东方财富-股权质押详情

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockGpzyDetailEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockGpzyEm

**描述**: StockGpzyEm 东方财富-股权质押

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockGpzyEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockGpzyPledgeRatioEm

**描述**: StockGpzyPledgeRatioEm 东方财富网-数据中心-特色数据-股权质押-上市公司质押比例

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockGpzyPledgeRatioEm("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockGpzyProfileEm

**描述**: StockGpzyProfileEm 东方财富网-数据中心-特色数据-股权质押-股权质押市场概况

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockGpzyProfileEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHistEm

**描述**: StockHistEm 东方财富-历史行情

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |
| period | string | 周期：daily/weekly/monthly |
| adjust | string | 复权类型：qfq(前复权)/hfq(后复权)/空字符串(不复权) |
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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockHistEm("000001", "daily", "qfq", "20230101", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHkGgtComponentsEm

**描述**: StockHkGgtComponentsEm 东方财富网-行情中心-港股市场-港股通成份股

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockHkGgtComponentsEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHkIndicatorEniu

**描述**: StockHkIndicatorEniu 亿牛网-港股指标

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockHkIndicatorEniu("000001", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHkValuationBaidu

**描述**: StockHkValuationBaidu 百度股市通-港股估值

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |
| indicator | string | 指标类型 |
| period | string | 周期：daily/weekly/monthly |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockHkValuationBaidu("000001", "", "daily")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHotDealXq

**描述**: StockHotDealXq 雪球-交易热度

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockHotDealXq("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHotFollowXq

**描述**: StockHotFollowXq 雪球-关注人数

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockHotFollowXq("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHotTweetXq

**描述**: StockHotTweetXq 雪球-讨论数

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockHotTweetXq("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHotXq

**描述**: StockHotXq 雪球-热门股票

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockHotXq()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHsgtEmFlow

**描述**: StockHsgtEmFlow 东方财富-沪深港通资金流向

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockHsgtEmFlow("")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHsgtExchangeRate

**描述**: StockHsgtExchangeRate 港股通汇率

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockHsgtExchangeRate()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHsgtFundFlowSummaryEm

**描述**: StockHsgtFundFlowSummaryEm 东方财富网-数据中心-资金流向-沪深港通资金流向

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockHsgtFundFlowSummaryEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHsgtFundMinEm

**描述**: StockHsgtFundMinEm 东方财富-沪深港通资金流向-分钟数据

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockHsgtFundMinEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHsgtHoldStockEm

**描述**: StockHsgtHoldStockEm 东方财富-数据中心-沪深港通持股-个股排行

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockHsgtHoldStockEm("", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHsgtMinEm

**描述**: StockHsgtMinEm 东方财富-沪深港通分时资金流向

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockHsgtMinEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockHsgtSpotEm

**描述**: StockHsgtSpotEm 东方财富网-沪深港通持股-实时行情

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockHsgtSpotEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockInfoChange

**描述**: StockInfoChange 股票信息变更

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockInfoChange()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockInfoCjzcEm

**描述**: StockInfoCjzcEm 东方财富-财经早餐

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockInfoCjzcEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockInfoGlobalCls

**描述**: StockInfoGlobalCls 财联社-全球财经快讯

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockInfoGlobalCls("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockInfoGlobalEm

**描述**: StockInfoGlobalEm 东方财富-全球财经快讯

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockInfoGlobalEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockInfoGlobalFutu

**描述**: StockInfoGlobalFutu 富途牛牛-全球财经快讯

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockInfoGlobalFutu()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockInfoGlobalSina

**描述**: StockInfoGlobalSina 新浪财经-全球财经快讯

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockInfoGlobalSina()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockInfoGlobalThs

**描述**: StockInfoGlobalThs 同花顺-全球财经快讯

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockInfoGlobalThs()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockInnerTradeXq

**描述**: StockInnerTradeXq 雪球-内部交易

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockInnerTradeXq("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockIrmCninfo

**描述**: StockIrmCninfo 巨潮资讯-互动易

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockIrmCninfo("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockJgdyDetailEm

**描述**: StockJgdyDetailEm 东方财富-机构调研详情

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockJgdyDetailEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockJgdyEm

**描述**: StockJgdyEm 东方财富-机构调研

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockJgdyEm("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockKcbASpotEm

**描述**: StockKcbASpotEm 东方财富网-科创板-实时行情

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockKcbASpotEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockLhYybCapital

**描述**: StockLhYybCapital 龙虎榜-营业部资金排名

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockLhYybCapital()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockLhYybControl

**描述**: StockLhYybControl 龙虎榜-营业部统计

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockLhYybControl()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockLhYybMost

**描述**: StockLhYybMost 龙虎榜-营业部上榜次数

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockLhYybMost()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockLhbDetailEm

**描述**: StockLhbDetailEm 东方财富网-数据中心-龙虎榜单-龙虎榜详情

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockLhbDetailEm("20230101", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockLhbGgtjSina

**描述**: StockLhbGgtjSina 新浪-龙虎榜-个股统计

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockLhbGgtjSina("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockLhbJgmmtjEm

**描述**: StockLhbJgmmtjEm 东方财富-龙虎榜-机构买卖统计

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockLhbJgmmtjEm("20230101", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockLhbJgstatisticEm

**描述**: StockLhbJgstatisticEm 东方财富-龙虎榜-机构席位统计

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockLhbJgstatisticEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockLhbStockDetailEm

**描述**: StockLhbStockDetailEm 东方财富-龙虎榜-个股龙虎榜详情

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockLhbStockDetailEm("000001", "20230101", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockLhbStockStatisticEm

**描述**: StockLhbStockStatisticEm 东方财富网-数据中心-龙虎榜单-个股上榜统计

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockLhbStockStatisticEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockLhbTraderstatisticEm

**描述**: StockLhbTraderstatisticEm 东方财富-龙虎榜-营业部统计

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockLhbTraderstatisticEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockLhbYybDetailEm

**描述**: StockLhbYybDetailEm 东方财富-龙虎榜-营业部详情

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockLhbYybDetailEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockLhbYybphEm

**描述**: StockLhbYybphEm 东方财富-龙虎榜-营业部排行

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockLhbYybphEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockMarginAccountInfo

**描述**: StockMarginAccountInfo 东方财富网-数据中心-融资融券-融资融券账户统计-两融账户信息

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockMarginAccountInfo()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockMarginEm

**描述**: StockMarginEm 东方财富-融资融券数据

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockMarginEm("20230101", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockMarginSse

**描述**: StockMarginSse 上交所融资融券汇总

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockMarginSse("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockMarginSzse

**描述**: StockMarginSzse 深交所融资融券汇总

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockMarginSzse("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockMarketActivityLegu

**描述**: StockMarketActivityLegu 乐估乐股-市场活跃度

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockMarketActivityLegu()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockNewSpotEm

**描述**: StockNewSpotEm 东方财富网-新股-实时行情

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockNewSpotEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockPankouEm

**描述**: StockPankouEm 东方财富-盘口异动

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockPankouEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockPgEm

**描述**: StockPgEm 东方财富-配股

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockPgEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockQbzfEm

**描述**: StockQbzfEm 东方财富-全部增发

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockQbzfEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockQsjyEm

**描述**: StockQsjyEm 东方财富-强势股统计

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockQsjyEm("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockRankCxdThs

**描述**: StockRankCxdThs 同花顺-技术选股-创新低

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockRankCxdThs("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockRankCxflThs

**描述**: StockRankCxflThs 同花顺-技术选股-持续放量

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockRankCxflThs()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockRankCxgThs

**描述**: StockRankCxgThs 同花顺-技术选股-创新高

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockRankCxgThs("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockRankCxslThs

**描述**: StockRankCxslThs 同花顺-技术选股-持续缩量

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockRankCxslThs()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockRankLjqdThs

**描述**: StockRankLjqdThs 同花顺-技术选股-量价齐跌

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockRankLjqdThs()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockRankLjqsThs

**描述**: StockRankLjqsThs 同花顺-技术选股-量价齐升

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockRankLjqsThs()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockRankLxszThs

**描述**: StockRankLxszThs 同花顺-技术选股-连续上涨

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockRankLxszThs()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockRankLxxdThs

**描述**: StockRankLxxdThs 同花顺-技术选股-连续下跌

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockRankLxxdThs()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockRankXstpThs

**描述**: StockRankXstpThs 同花顺-技术选股-向上突破

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockRankXstpThs("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockRankXxtpThs

**描述**: StockRankXxtpThs 同花顺-技术选股-向下突破

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockRankXxtpThs("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockRankXzjpThs

**描述**: StockRankXzjpThs 同花顺-技术选股-险资举牌

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockRankXzjpThs()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockReportDisclosure

**描述**: StockReportDisclosure 巨潮资讯-业绩预告披露时间

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockReportDisclosure("000001", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockReportEm

**描述**: StockReportEm 东方财富-研究报告

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockReportEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockResearchReportEm

**描述**: StockResearchReportEm 东方财富-研究报告详情

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockResearchReportEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockSgtReferenceExchangeRateSse

**描述**: StockSgtReferenceExchangeRateSse 上交所-港股通参考汇率

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockSgtReferenceExchangeRateSse()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockSgtReferenceExchangeRateSzse

**描述**: StockSgtReferenceExchangeRateSzse 深交所-港股通参考汇率

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockSgtReferenceExchangeRateSzse()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockSgtSettlementExchangeRateSse

**描述**: StockSgtSettlementExchangeRateSse 上交所-港股通结算汇率

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockSgtSettlementExchangeRateSse()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockSgtSettlementExchangeRateSzse

**描述**: StockSgtSettlementExchangeRateSzse 深交所-港股通结算汇率

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockSgtSettlementExchangeRateSzse()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockShASpotEm

**描述**: StockShASpotEm 东方财富网-沪 A 股-实时行情

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockShASpotEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockSnsSseinfo

**描述**: StockSnsSseinfo 上证e互动

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockSnsSseinfo("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockSyEm

**描述**: StockSyEm 东方财富-数据中心-商誉

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockSyEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockSzASpotEm

**描述**: StockSzASpotEm 东方财富网-深 A 股-实时行情

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockSzASpotEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockTfpEm

**描述**: StockTfpEm 东方财富网-数据中心-特色数据-停复牌信息

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockTfpEm("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockThreeReportEm

**描述**: StockThreeReportEm 东方财富-三大报表

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |
| reportType | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockThreeReportEm("000001", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockUsValuationBaidu

**描述**: StockUsValuationBaidu 百度股市通-美股估值

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |
| indicator | string | 指标类型 |
| period | string | 周期：daily/weekly/monthly |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockUsValuationBaidu("000001", "", "daily")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockValueEm

**描述**: StockValueEm 东方财富网-数据中心-估值分析

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockValueEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockYjbbEm

**描述**: StockYjbbEm 东方财富-数据中心-年报季报-业绩快报-业绩报表

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockYjbbEm("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockYjkbEm

**描述**: StockYjkbEm 东方财富-业绩快报

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockYjkbEm("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockYjygEm

**描述**: StockYjygEm 东方财富-业绩预告

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockYjygEm("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockYzxdrEm

**描述**: StockYzxdrEm 东方财富-数据中心-一致性预期-盈利预测

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockYzxdrEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockZbPoolEm

**描述**: StockZbPoolEm 东方财富网-行情中心-涨停板行情-炸板股池

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockZbPoolEm("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockZdhtmxEm

**描述**: StockZdhtmxEm 东方财富-数据中心-重大合同明细

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockZdhtmxEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockZhAGdhs

**描述**: StockZhAGdhs A股股东户数

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockZhAGdhs("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockZhAGdhsDetailEm

**描述**: StockZhAGdhsDetailEm 东方财富-股东户数详情

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockZhAGdhsDetailEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockZhAHistTx

**描述**: StockZhAHistTx 腾讯财经-A股历史行情

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockZhAHistTx("000001", "20230101", "20230101", "qfq")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockZhASpotEm

**描述**: StockZhASpotEm 东方财富网-沪深京 A 股-实时行情

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockZhASpotEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockZhBSpotEm

**描述**: StockZhBSpotEm 东方财富网-沪深 B 股-实时行情

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockZhBSpotEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockZhValuationBaidu

**描述**: StockZhValuationBaidu 百度股市通-A股估值

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| symbol | string | 股票/基金代码 |
| indicator | string | 指标类型 |
| period | string | 周期：daily/weekly/monthly |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockZhValuationBaidu("000001", "", "daily")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockZhVoteBaidu

**描述**: StockZhVoteBaidu 百度股市通-A股投票

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockZhVoteBaidu("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockZtPoolEm

**描述**: StockZtPoolEm 东方财富网-行情中心-涨停板行情-涨停股池

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockZtPoolEm("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## StockZtPoolPreviousEm

**描述**: StockZtPoolPreviousEm 东方财富网-行情中心-涨停板行情-昨日涨停股池

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
	"github.com/BlakeLiAFK/akshare/stock_feature"
)

func main() {
	data, err := stock_feature.StockZtPoolPreviousEm("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

