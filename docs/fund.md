# fund 模块

> 本模块共有 86 个接口

## 目录

- [AmacAicInfo](#amacaicinfo)
- [AmacFundAccount](#amacfundaccount)
- [AmacFundAccountSubFund](#amacfundaccountsubfund)
- [AmacFundInfo](#amacfundinfo)
- [AmacFundSub](#amacfundsub)
- [AmacFuturesInfo](#amacfuturesinfo)
- [AmacManagerClassifyInfo](#amacmanagerclassifyinfo)
- [AmacManagerInfo](#amacmanagerinfo)
- [AmacManagerSecuritiesInfo](#amacmanagersecuritiesinfo)
- [AmacMemberInfo](#amacmemberinfo)
- [AmacMemberRosterInfo](#amacmemberrosterinfo)
- [AmacMemberSubInfo](#amacmembersubinfo)
- [AmacPersonFund](#amacpersonfund)
- [AmacPersonFundOrg](#amacpersonfundorg)
- [AmacSecuritiesInfo](#amacsecuritiesinfo)
- [FundAnnouncementDividendEm](#fundannouncementdividendem)
- [FundAnnouncementPersonnelEm](#fundannouncementpersonnelem)
- [FundAnnouncementReportEm](#fundannouncementreportem)
- [FundAumEm](#fundaumem)
- [FundAumHistEm](#fundaumhistem)
- [FundAumTrendEm](#fundaumtrendem)
- [FundBalancePositionLg](#fundbalancepositionlg)
- [FundCfEm](#fundcfem)
- [FundEtfCategorySina](#fundetfcategorysina)
- [FundEtfDividendSina](#fundetfdividendsina)
- [FundEtfFundDailyEm](#fundetffunddailyem)
- [FundEtfFundInfoEm](#fundetffundinfoem)
- [FundEtfHistEm](#fundetfhistem)
- [FundEtfHistMinEm](#fundetfhistminem)
- [FundEtfHistSina](#fundetfhistsina)
- [FundEtfSpotEm](#fundetfspotem)
- [FundEtfSpotThs](#fundetfspotths)
- [FundExchangeRankEm](#fundexchangerankem)
- [FundFeeEm](#fundfeeem)
- [FundFhEm](#fundfhem)
- [FundFhRankEm](#fundfhrankem)
- [FundFinancialFundDailyEm](#fundfinancialfunddailyem)
- [FundFinancialFundInfoEm](#fundfinancialfundinfoem)
- [FundGradedFundDailyEm](#fundgradedfunddailyem)
- [FundGradedFundInfoEm](#fundgradedfundinfoem)
- [FundHistXq](#fundhistxq)
- [FundHkFundHistEm](#fundhkfundhistem)
- [FundHkRankEm](#fundhkrankem)
- [FundHoldStructureEm](#fundholdstructureem)
- [FundHqTimeEm](#fundhqtimeem)
- [FundIndividualDetailInfoEm](#fundindividualdetailinfoem)
- [FundInfoIndexEm](#fundinfoindexem)
- [FundInfoXq](#fundinfoxq)
- [FundLcxRankEm](#fundlcxrankem)
- [FundLinghuoPositionLg](#fundlinghuopositionlg)
- [FundLofHistEm](#fundlofhistem)
- [FundLofHistMinEm](#fundlofhistminem)
- [FundLofSpotEm](#fundlofspotem)
- [FundManagerEm](#fundmanagerem)
- [FundMarketXq](#fundmarketxq)
- [FundMoneyFundDailyEm](#fundmoneyfunddailyem)
- [FundMoneyFundInfoEm](#fundmoneyfundinfoem)
- [FundMoneyRankEm](#fundmoneyrankem)
- [FundNameEm](#fundnameem)
- [FundNetValueXq](#fundnetvaluexq)
- [FundNewFoundEm](#fundnewfoundem)
- [FundOpenFundDailyEm](#fundopenfunddailyem)
- [FundOpenFundInfoEm](#fundopenfundinfoem)
- [FundOpenFundRankEm](#fundopenfundrankem)
- [FundOverviewEm](#fundoverviewem)
- [FundPortfolioBondHoldEm](#fundportfoliobondholdem)
- [FundPortfolioChangeEm](#fundportfoliochangeem)
- [FundPortfolioHoldEm](#fundportfolioholdem)
- [FundPortfolioIndustryAllocationEm](#fundportfolioindustryallocationem)
- [FundPortfolioXq](#fundportfolioxq)
- [FundPurchaseEm](#fundpurchaseem)
- [FundRatingAll](#fundratingall)
- [FundRatingJa](#fundratingja)
- [FundRatingSh](#fundratingsh)
- [FundRatingZs](#fundratingzs)
- [FundReportAssetAllocationCninfo](#fundreportassetallocationcninfo)
- [FundReportCombineCninfo](#fundreportcombinecninfo)
- [FundReportIndustryAllocationCninfo](#fundreportindustryallocationcninfo)
- [FundReportStockCninfo](#fundreportstockcninfo)
- [FundScaleChangeEm](#fundscalechangeem)
- [FundScaleCloseSina](#fundscaleclosesina)
- [FundScaleOpenSina](#fundscaleopensina)
- [FundScaleStructuredSina](#fundscalestructuredsina)
- [FundSearchXq](#fundsearchxq)
- [FundStockPositionLg](#fundstockpositionlg)
- [FundValueEstimationEm](#fundvalueestimationem)

---

## AmacAicInfo

**描述**: AmacAicInfo 获取中国证券投资基金业协会-信息公示-基金产品-证券公司直投基金

**数据源**: https://gs.amac.org.cn/amac-infodisc/res/aoin/product/index.html

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.AmacAicInfo()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## AmacFundAccount

**描述**: AmacFundAccount 获取中国证券投资基金业协会-信息公示-基金产品-基金公司及子公司集合资管产品公示

**数据源**: https://gs.amac.org.cn/amac-infodisc/res/fund/account/index.html

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.AmacFundAccount()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## AmacFundAccountSubFund

**描述**: AmacFundAccountSubFund 获取中国证券投资基金业协会-信息公示-基金产品-资产支持专项计划

**数据源**: https://gs.amac.org.cn/amac-infodisc/res/fund/abs/index.html

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.AmacFundAccountSubFund()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## AmacFundInfo

**描述**: AmacFundInfo 获取中国证券投资基金业协会-信息公示-基金产品-私募基金管理人基金产品

**数据源**: https://gs.amac.org.cn/amac-infodisc/res/pof/fund/index.html

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| startPage | string | - |
| endPage | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.AmacFundInfo("", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## AmacFundSub

**描述**: AmacFundSub 获取中国证券投资基金业协会-信息公示-基金产品-证券公司私募投资基金

**数据源**: https://gs.amac.org.cn/amac-infodisc/res/pof/subfund/index.html

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.AmacFundSub()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## AmacFuturesInfo

**描述**: AmacFuturesInfo 获取中国证券投资基金业协会-信息公示-基金产品-期货公司集合资管产品公示

**数据源**: https://gs.amac.org.cn/amac-infodisc/res/pof/futures/index.html

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.AmacFuturesInfo()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## AmacManagerClassifyInfo

**描述**: AmacManagerClassifyInfo 获取中国证券投资基金业协会-信息公示-私募基金管理人公示-私募基金管理人分类公示

**数据源**: https://gs.amac.org.cn/amac-infodisc/res/pof/manager/index.html

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.AmacManagerClassifyInfo()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## AmacManagerInfo

**描述**: AmacManagerInfo 获取中国证券投资基金业协会-信息公示-私募基金管理人公示-私募基金管理人综合查询

**数据源**: https://gs.amac.org.cn/amac-infodisc/res/pof/manager/index.html

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.AmacManagerInfo()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## AmacManagerSecuritiesInfo

**描述**: AmacManagerSecuritiesInfo 获取中国证券投资基金业协会-信息公示-诚信信息-已注销私募基金管理人名单

**数据源**: https://gs.amac.org.cn/amac-infodisc/res/cancelled/manager/index.html

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.AmacManagerSecuritiesInfo()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## AmacMemberInfo

**描述**: AmacMemberInfo 获取中国证券投资基金业协会-信息公示-会员信息-会员机构综合查询

**数据源**: https://gs.amac.org.cn/amac-infodisc/res/pof/member/index.html

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.AmacMemberInfo()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## AmacMemberRosterInfo

**描述**: AmacMemberRosterInfo 此函数在Python源码中未找到对应实现

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.AmacMemberRosterInfo()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## AmacMemberSubInfo

**描述**: AmacMemberSubInfo 获取中国证券投资基金业协会-信息公示-私募基金管理人公示-证券公司私募基金子公司管理人信息公示

**数据源**: https://gs.amac.org.cn/amac-infodisc/res/pof/member/index.html?primaryInvestType=private

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.AmacMemberSubInfo()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## AmacPersonFund

**描述**: AmacPersonFund 获取中国证券投资基金业协会-信息公示-从业人员信息-债券投资交易相关人员公示

**数据源**: https://gs.amac.org.cn/amac-infodisc/res/person/bond/index.html

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.AmacPersonFund()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## AmacPersonFundOrg

**描述**: AmacPersonFundOrg 获取中国证券投资基金业协会-信息公示-从业人员信息-基金从业人员资格注册信息

**数据源**: https://gs.amac.org.cn/amac-infodisc/res/person/fund/index.html

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.AmacPersonFundOrg("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## AmacSecuritiesInfo

**描述**: AmacSecuritiesInfo 获取中国证券投资基金业协会-信息公示-基金产品-证券公司集合资管产品公示

**数据源**: https://gs.amac.org.cn/amac-infodisc/res/pof/securities/index.html

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.AmacSecuritiesInfo()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundAnnouncementDividendEm

**描述**: FundAnnouncementDividendEm 获取基金分红配送公告

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundAnnouncementDividendEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundAnnouncementPersonnelEm

**描述**: FundAnnouncementPersonnelEm 获取基金人事调整公告

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundAnnouncementPersonnelEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundAnnouncementReportEm

**描述**: FundAnnouncementReportEm 获取基金定期报告公告

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundAnnouncementReportEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundAumEm

**描述**: FundAumEm 东方财富-基金-基金公司排名列表

**数据源**: https://fund.eastmoney.com/Company/lsgm.html

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundAumEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundAumHistEm

**描述**: FundAumHistEm 东方财富-基金-基金公司历年管理规模排行列表

**数据源**: https://fund.eastmoney.com/Company/lsgm.html

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundAumHistEm("")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundAumTrendEm

**描述**: FundAumTrendEm 东方财富-基金-基金市场管理规模走势图

**数据源**: https://fund.eastmoney.com/Company/default.html

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundAumTrendEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundBalancePositionLg

**描述**: FundBalancePositionLg 获取乐咕乐股平衡混合型基金仓位数据

**数据源**: https://legulegu.com/stockdata/fund-position/pos-pingheng

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundBalancePositionLg()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundCfEm

**描述**: FundCfEm 获取天天基金网-基金数据-分红送配-基金拆分

**数据源**: https://fund.eastmoney.com/data/fundchaifen.html

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundCfEm("")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundEtfCategorySina

**描述**: FundEtfCategorySina 获取新浪财经基金列表

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundEtfCategorySina("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundEtfDividendSina

**描述**: FundEtfDividendSina 获取ETF基金的累计分红信息

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundEtfDividendSina("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundEtfFundDailyEm

**描述**: FundEtfFundDailyEm 东方财富网-天天基金网-基金数据-场内交易基金

**数据源**: https://fund.eastmoney.com/cnjy_dwjz.html

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundEtfFundDailyEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundEtfFundInfoEm

**描述**: FundEtfFundInfoEm 东方财富网-天天基金网-基金数据-场内交易基金-历史净值明细

**数据源**: https://fundf10.eastmoney.com/jjjz_511280.html

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundEtfFundInfoEm("000001", "20230101", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundEtfHistEm

**描述**: FundEtfHistEm 获取 ETF 历史行情数据

**数据源**: https://quote.eastmoney.com/concept/sz000062.html

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundEtfHistEm("000001", "daily", "20230101", "20230101", "qfq")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundEtfHistMinEm

**描述**: FundEtfHistMinEm 获取 ETF 分钟级别行情数据

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundEtfHistMinEm("000001", "20230101", "20230101", "daily", "qfq")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundEtfHistSina

**描述**: FundEtfHistSina 获取ETF基金的日行情数据

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundEtfHistSina("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundEtfSpotEm

**描述**: FundEtfSpotEm 获取沪深京三大市场实时 ETF 行情数据

**数据源**: https://quote.eastmoney.com/center/gridlist.html#fund_etf

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundEtfSpotEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundEtfSpotThs

**描述**: FundEtfSpotThs 获取同花顺理财-基金数据-每日净值-ETF-实时行情

**数据源**: https://fund.10jqka.com.cn/datacenter/jz/kfs/etf/

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundEtfSpotThs("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundExchangeRankEm

**描述**: FundExchangeRankEm 获取场内交易基金排名

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundExchangeRankEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundFeeEm

**描述**: FundFeeEm 获取天天基金-基金档案-购买信息

**数据源**: https://fundf10.eastmoney.com/jjfl_015641.html

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundFeeEm("000001", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundFhEm

**描述**: FundFhEm 获取天天基金网-基金数据-分红送配-基金分红

**数据源**: https://fund.eastmoney.com/data/fundfenhong.html

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundFhEm("")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundFhRankEm

**描述**: FundFhRankEm 获取天天基金网-基金数据-分红送配-基金分红排行

**数据源**: https://fund.eastmoney.com/data/fundleijifenhong.html

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundFhRankEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundFinancialFundDailyEm

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundFinancialFundDailyEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundFinancialFundInfoEm

**描述**: FundFinancialFundInfoEm 东方财富网-天天基金网-基金数据-理财型基金-历史净值明细

**数据源**: https://fundf10.eastmoney.com/jjjz_000791.html

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundFinancialFundInfoEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundGradedFundDailyEm

**描述**: FundGradedFundDailyEm 东方财富网-天天基金网-基金数据-分级基金净值

**数据源**: https://fund.eastmoney.com/fjjj.html#1_1__0__zdf,desc_1

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundGradedFundDailyEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundGradedFundInfoEm

**描述**: FundGradedFundInfoEm 东方财富网-天天基金网-基金数据-分级基金-历史净值明细

**数据源**: https://fundf10.eastmoney.com/jjjz_150232.html

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundGradedFundInfoEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundHistXq

**描述**: FundHistXq 获取雪球基金历史业绩数据

**数据源**: https://danjuanfunds.com/djapi/fundx/base/fund/achievement/000001

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundHistXq("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundHkFundHistEm

**描述**: FundHkFundHistEm 东方财富网-天天基金网-基金数据-香港基金-历史净值明细(分红送配详情)

**数据源**: https://overseas.1234567.com.cn/f10/FundJz/968092#FHPS

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| code | string | 股票/基金代码 |
| indicator | string | 指标类型 |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundHkFundHistEm("000001", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundHkRankEm

**描述**: FundHkRankEm 获取香港基金排名

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundHkRankEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundHoldStructureEm

**描述**: FundHoldStructureEm 获取基金持有人结构数据

**数据源**: https://fund.eastmoney.com/data/cyrjglist.html

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundHoldStructureEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundHqTimeEm

**描述**: FundHqTimeEm Python源码中不存在此函数，可能是命名错误或已废弃

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundHqTimeEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundIndividualDetailInfoEm

**描述**: FundIndividualDetailInfoEm Python源码中不存在此函数，可能是命名错误或已废弃

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundIndividualDetailInfoEm("000001", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundInfoIndexEm

**描述**: FundInfoIndexEm 获取指数型基金信息

**数据源**: https://fund.eastmoney.com/trade/zs.html

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundInfoIndexEm("000001", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundInfoXq

**描述**: FundInfoXq 获取雪球基金基本信息

**数据源**: https://danjuanfunds.com/djapi/fund/000001

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundInfoXq("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundLcxRankEm

**描述**: FundLcxRankEm 获取理财型基金排名

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundLcxRankEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundLinghuoPositionLg

**描述**: FundLinghuoPositionLg 获取乐咕乐股灵活配置型基金仓位数据

**数据源**: https://legulegu.com/stockdata/fund-position/pos-linghuo

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundLinghuoPositionLg()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundLofHistEm

**描述**: FundLofHistEm 获取LOF历史行情数据

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundLofHistEm("000001", "daily", "20230101", "20230101", "qfq")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundLofHistMinEm

**描述**: FundLofHistMinEm 获取LOF分钟级别行情数据

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundLofHistMinEm("000001", "20230101", "20230101", "daily", "qfq")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundLofSpotEm

**描述**: FundLofSpotEm 获取LOF实时行情数据

**数据源**: https://quote.eastmoney.com/center/gridlist.html#fund_lof

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundLofSpotEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundManagerEm

**描述**: FundManagerEm 获取天天基金网-基金数据-基金经理大全

**数据源**: https://fund.eastmoney.com/manager/default.html

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundManagerEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundMarketXq

**描述**: FundMarketXq 获取雪球基金数据分析

**数据源**: https://danjuanfunds.com/djapi/fund/base/quote/data/index/analysis/000001

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundMarketXq("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundMoneyFundDailyEm

**描述**: FundMoneyFundDailyEm 东方财富网-天天基金网-基金数据-货币型基金收益

**数据源**: https://fund.eastmoney.com/HBJJ_pjsyl.html

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundMoneyFundDailyEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundMoneyFundInfoEm

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundMoneyFundInfoEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundMoneyRankEm

**描述**: FundMoneyRankEm 获取货币型基金排名

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundMoneyRankEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundNameEm

**描述**: FundNameEm 获取所有基金名称和类型

**数据源**: https://fund.eastmoney.com/manager/default.html

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundNameEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundNetValueXq

**描述**: FundNetValueXq 获取雪球基金盈利概率数据

**数据源**: https://danjuanfunds.com/djapi/fundx/base/fund/profit/ratio/000001

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundNetValueXq("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundNewFoundEm

**描述**: FundNewFoundEm 获取基金数据-新发基金-新成立基金

**数据源**: https://fund.eastmoney.com/data/xinfound.html

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundNewFoundEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundOpenFundDailyEm

**描述**: FundOpenFundDailyEm 获取当前交易日所有开放式基金净值

**数据源**: https://fund.eastmoney.com/fund.html

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundOpenFundDailyEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundOpenFundInfoEm

**描述**: FundOpenFundInfoEm 东方财富网-天天基金网-基金数据-开放式基金净值

**数据源**: https://fund.eastmoney.com/fund.html

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundOpenFundInfoEm("000001", "", "daily")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundOpenFundRankEm

**描述**: FundOpenFundRankEm 获取开放式基金排名

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundOpenFundRankEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundOverviewEm

**描述**: FundOverviewEm 获取天天基金-基金档案-基本概况

**数据源**: https://fundf10.eastmoney.com/jbgk_015641.html

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundOverviewEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundPortfolioBondHoldEm

**描述**: FundPortfolioBondHoldEm 获取基金债券持仓

**数据源**: https://fundf10.eastmoney.com/ccmx1_000001.html

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundPortfolioBondHoldEm("000001", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundPortfolioChangeEm

**描述**: FundPortfolioChangeEm 获取基金持仓变动

**数据源**: https://fundf10.eastmoney.com/ccbd_000001.html

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundPortfolioChangeEm("000001", "", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundPortfolioHoldEm

**描述**: FundPortfolioHoldEm 获取基金股票持仓

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundPortfolioHoldEm("000001", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundPortfolioIndustryAllocationEm

**描述**: FundPortfolioIndustryAllocationEm 获取基金行业配置

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundPortfolioIndustryAllocationEm("000001", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundPortfolioXq

**描述**: FundPortfolioXq 获取雪球基金持仓数据

**数据源**: https://danjuanfunds.com/djapi/fundx/base/fund/record/asset/percent

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundPortfolioXq("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundPurchaseEm

**描述**: FundPurchaseEm 获取基金申购状态

**数据源**: https://fund.eastmoney.com/Fund_sgzt_bzdm.html

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundPurchaseEm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundRatingAll

**描述**: FundRatingAll 获取全部基金评级数据

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundRatingAll()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundRatingJa

**描述**: FundRatingJa 获取济安金信评级

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundRatingJa("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundRatingSh

**描述**: FundRatingSh 获取上海证券评级

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundRatingSh("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundRatingZs

**描述**: FundRatingZs 获取招商证券评级

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundRatingZs("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundReportAssetAllocationCninfo

**描述**: FundReportAssetAllocationCninfo 巨潮资讯-数据中心-专题统计-基金报表-基金资产配置

**数据源**: https://webapi.cninfo.com.cn/#/thematicStatistics

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundReportAssetAllocationCninfo()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundReportCombineCninfo

**描述**: FundReportCombineCninfo Python源码中不存在此函数，可能是命名错误或已废弃

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundReportCombineCninfo("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundReportIndustryAllocationCninfo

**描述**: FundReportIndustryAllocationCninfo 巨潮资讯-数据中心-专题统计-基金报表-基金行业配置

**数据源**: https://webapi.cninfo.com.cn/#/thematicStatistics

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundReportIndustryAllocationCninfo("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundReportStockCninfo

**描述**: FundReportStockCninfo 巨潮资讯-数据中心-专题统计-基金报表-基金重仓股

**数据源**: https://webapi.cninfo.com.cn/#/thematicStatistics

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundReportStockCninfo("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundScaleChangeEm

**描述**: FundScaleChangeEm 获取基金规模变动数据

**数据源**: https://fund.eastmoney.com/data/gmbdlist.html

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundScaleChangeEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundScaleCloseSina

**描述**: FundScaleCloseSina 获取新浪封闭式基金规模数据

**数据源**: https://vip.stock.finance.sina.com.cn/fund_center/index.html#jjhqetf

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundScaleCloseSina()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundScaleOpenSina

**描述**: FundScaleOpenSina 获取新浪开放式基金规模数据

**数据源**: https://vip.stock.finance.sina.com.cn/fund_center/index.html#jjhqetf

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundScaleOpenSina("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundScaleStructuredSina

**描述**: FundScaleStructuredSina 获取新浪分级基金规模数据

**数据源**: https://vip.stock.finance.sina.com.cn/fund_center/index.html#jjgmfjall

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundScaleStructuredSina()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundSearchXq

**描述**: FundSearchXq 获取雪球基金交易规则

**数据源**: https://danjuanfunds.com/djapi/fund/detail/000001

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| keyword | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundSearchXq("")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundStockPositionLg

**描述**: FundStockPositionLg 获取乐咕乐股基金股票型基金仓位数据

**数据源**: https://legulegu.com/stockdata/fund-position/pos-stock

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundStockPositionLg("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## FundValueEstimationEm

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
	"github.com/BlakeLiAFK/akshare/fund"
)

func main() {
	data, err := fund.FundValueEstimationEm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

