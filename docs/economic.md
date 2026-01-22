# economic 模块

> 本模块共有 230 个接口

## 目录

- [CryptoJSSpot](#cryptojsspot)
- [MacroAustraliaBankRate](#macroaustraliabankrate)
- [MacroAustraliaCPIQuarterly](#macroaustraliacpiquarterly)
- [MacroAustraliaCPIYearly](#macroaustraliacpiyearly)
- [MacroAustraliaPPIQuarterly](#macroaustraliappiquarterly)
- [MacroAustraliaRetailRateMonthly](#macroaustraliaretailratemonthly)
- [MacroAustraliaTrade](#macroaustraliatrade)
- [MacroAustraliaUnemploymentRate](#macroaustraliaunemploymentrate)
- [MacroBankAustraliaInterestRate](#macrobankaustraliainterestrate)
- [MacroBankBrazilInterestRate](#macrobankbrazilinterestrate)
- [MacroBankChinaInterestRate](#macrobankchinainterestrate)
- [MacroBankEnglishInterestRate](#macrobankenglishinterestrate)
- [MacroBankEuroInterestRate](#macrobankeurointerestrate)
- [MacroBankIndiaInterestRate](#macrobankindiainterestrate)
- [MacroBankJapanInterestRate](#macrobankjapaninterestrate)
- [MacroBankNewzealandInterestRate](#macrobanknewzealandinterestrate)
- [MacroBankRussiaInterestRate](#macrobankrussiainterestrate)
- [MacroBankSwitzerlandInterestRate](#macrobankswitzerlandinterestrate)
- [MacroBankUsaInterestRate](#macrobankusainterestrate)
- [MacroCNBS](#macrocnbs)
- [MacroCanadaBankRate](#macrocanadabankrate)
- [MacroCanadaCPIMonthly](#macrocanadacpimonthly)
- [MacroCanadaCPIYearly](#macrocanadacpiyearly)
- [MacroCanadaCoreCPIMonthly](#macrocanadacorecpimonthly)
- [MacroCanadaCoreCPIYearly](#macrocanadacorecpiyearly)
- [MacroCanadaGDPMonthly](#macrocanadagdpmonthly)
- [MacroCanadaNewHouseRate](#macrocanadanewhouserate)
- [MacroCanadaRetailRateMonthly](#macrocanadaretailratemonthly)
- [MacroCanadaTrade](#macrocanadatrade)
- [MacroCanadaUnemploymentRate](#macrocanadaunemploymentrate)
- [MacroChinaAgriculturalIndex](#macrochinaagriculturalindex)
- [MacroChinaAgriculturalProduct](#macrochinaagriculturalproduct)
- [MacroChinaAuReport](#macrochinaaureport)
- [MacroChinaBDTIIndex](#macrochinabdtiindex)
- [MacroChinaBSIIndex](#macrochinabsiindex)
- [MacroChinaBankFinancing](#macrochinabankfinancing)
- [MacroChinaCPI](#macrochinacpi)
- [MacroChinaCPIMonthly](#macrochinacpimonthly)
- [MacroChinaCPIYearly](#macrochinacpiyearly)
- [MacroChinaCentralBankBalance](#macrochinacentralbankbalance)
- [MacroChinaCommodityPriceIndex](#macrochinacommoditypriceindex)
- [MacroChinaConstructionIndex](#macrochinaconstructionindex)
- [MacroChinaConstructionPriceIndex](#macrochinaconstructionpriceindex)
- [MacroChinaConsumerGoodsRetail](#macrochinaconsumergoodsretail)
- [MacroChinaCxPMIYearly](#macrochinacxpmiyearly)
- [MacroChinaCxServicesPMIYearly](#macrochinacxservicespmiyearly)
- [MacroChinaCzsr](#macrochinaczsr)
- [MacroChinaDailyEnergy](#macrochinadailyenergy)
- [MacroChinaEnergyIndex](#macrochinaenergyindex)
- [MacroChinaEnterpriseBoomIndex](#macrochinaenterpriseboomindex)
- [MacroChinaExportsYoY](#macrochinaexportsyoy)
- [MacroChinaFDI](#macrochinafdi)
- [MacroChinaForeignExchangeGold](#macrochinaforeignexchangegold)
- [MacroChinaFreightIndex](#macrochinafreightindex)
- [MacroChinaFxGold](#macrochinafxgold)
- [MacroChinaFxReservesYearly](#macrochinafxreservesyearly)
- [MacroChinaGDP](#macrochinagdp)
- [MacroChinaGDPYearly](#macrochinagdpyearly)
- [MacroChinaGdzctz](#macrochinagdzctz)
- [MacroChinaGyzjz](#macrochinagyzjz)
- [MacroChinaHKBuildingAmount](#macrochinahkbuildingamount)
- [MacroChinaHKBuildingVolume](#macrochinahkbuildingvolume)
- [MacroChinaHKCPI](#macrochinahkcpi)
- [MacroChinaHKCPIRatio](#macrochinahkcpiratio)
- [MacroChinaHKCore](#macrochinahkcore)
- [MacroChinaHKGBP](#macrochinahkgbp)
- [MacroChinaHKGBPRatio](#macrochinahkgbpratio)
- [MacroChinaHKMarketInfo](#macrochinahkmarketinfo)
- [MacroChinaHKPPI](#macrochinahkppi)
- [MacroChinaHKRateOfUnemployment](#macrochinahkrateofunemployment)
- [MacroChinaHKTradeDiffRatio](#macrochinahktradediffratio)
- [MacroChinaHgjck](#macrochinahgjck)
- [MacroChinaImportsYoY](#macrochinaimportsyoy)
- [MacroChinaIndustrialProductionYoY](#macrochinaindustrialproductionyoy)
- [MacroChinaInsurance](#macrochinainsurance)
- [MacroChinaInsuranceIncome](#macrochinainsuranceincome)
- [MacroChinaInternationalTourismFx](#macrochinainternationaltourismfx)
- [MacroChinaLPIIndex](#macrochinalpiindex)
- [MacroChinaLPR](#macrochinalpr)
- [MacroChinaM2Yearly](#macrochinam2yearly)
- [MacroChinaMarketMarginSH](#macrochinamarketmarginsh)
- [MacroChinaMarketMarginSZ](#macrochinamarketmarginsz)
- [MacroChinaMobileNumber](#macrochinamobilenumber)
- [MacroChinaMoneySupply](#macrochinamoneysupply)
- [MacroChinaNBSNation](#macrochinanbsnation)
- [MacroChinaNBSRegion](#macrochinanbsregion)
- [MacroChinaNationalTaxReceipts](#macrochinanationaltaxreceipts)
- [MacroChinaNewFinancialCredit](#macrochinanewfinancialcredit)
- [MacroChinaNewHousePrice](#macrochinanewhouseprice)
- [MacroChinaNonManPMI](#macrochinanonmanpmi)
- [MacroChinaPMI](#macrochinapmi)
- [MacroChinaPMIYearly](#macrochinapmiyearly)
- [MacroChinaPPI](#macrochinappi)
- [MacroChinaPPIYearly](#macrochinappiyearly)
- [MacroChinaPassengerLoadFactor](#macrochinapassengerloadfactor)
- [MacroChinaPostalTelecommunicational](#macrochinapostaltelecommunicational)
- [MacroChinaQyspjg](#macrochinaqyspjg)
- [MacroChinaRMB](#macrochinarmb)
- [MacroChinaRealEstate](#macrochinarealestate)
- [MacroChinaReserveRequirementRatio](#macrochinareserverequirementratio)
- [MacroChinaRetailPriceIndex](#macrochinaretailpriceindex)
- [MacroChinaShiborAll](#macrochinashiborall)
- [MacroChinaShrzgm](#macrochinashrzgm)
- [MacroChinaSocietyElectricity](#macrochinasocietyelectricity)
- [MacroChinaSocietyTrafficVolume](#macrochinasocietytrafficvolume)
- [MacroChinaStockMarketCap](#macrochinastockmarketcap)
- [MacroChinaSupplyOfMoney](#macrochinasupplyofmoney)
- [MacroChinaTradeBalance](#macrochinatradebalance)
- [MacroChinaUrbanUnemployment](#macrochinaurbanunemployment)
- [MacroChinaVegetableBasket](#macrochinavegetablebasket)
- [MacroChinaWbck](#macrochinawbck)
- [MacroChinaWhxd](#macrochinawhxd)
- [MacroChinaXfzxx](#macrochinaxfzxx)
- [MacroChinaYwElectronicIndex](#macrochinaywelectronicindex)
- [MacroConsGold](#macroconsgold)
- [MacroConsOpecMonth](#macroconsopecmonth)
- [MacroConsOpecMonthWithLimit](#macroconsopecmonthwithlimit)
- [MacroConsSilver](#macroconssilver)
- [MacroEuroCPIMoM](#macroeurocpimom)
- [MacroEuroCPIYoY](#macroeurocpiyoy)
- [MacroEuroCurrentAccountMoM](#macroeurocurrentaccountmom)
- [MacroEuroEmploymentChangeQoQ](#macroeuroemploymentchangeqoq)
- [MacroEuroGDPYoY](#macroeurogdpyoy)
- [MacroEuroIndustrialProductionMoM](#macroeuroindustrialproductionmom)
- [MacroEuroLMEHolding](#macroeurolmeholding)
- [MacroEuroLMEStock](#macroeurolmestock)
- [MacroEuroManufacturingPMI](#macroeuromanufacturingpmi)
- [MacroEuroPPIMoM](#macroeuroppimom)
- [MacroEuroRetailSalesMoM](#macroeuroretailsalesmom)
- [MacroEuroSentixInvestorConfidence](#macroeurosentixinvestorconfidence)
- [MacroEuroServicesPMI](#macroeuroservicespmi)
- [MacroEuroTradeBalance](#macroeurotradebalance)
- [MacroEuroUnemploymentRateMoM](#macroeurounemploymentratemom)
- [MacroEuroZEWEconomicSentiment](#macroeurozeweconomicsentiment)
- [MacroFXSentiment](#macrofxsentiment)
- [MacroGermanyCPIMonthly](#macrogermanycpimonthly)
- [MacroGermanyCPIYearly](#macrogermanycpiyearly)
- [MacroGermanyCore](#macrogermanycore)
- [MacroGermanyGDP](#macrogermanygdp)
- [MacroGermanyIFO](#macrogermanyifo)
- [MacroGermanyRetailSaleMonthly](#macrogermanyretailsalemonthly)
- [MacroGermanyRetailSaleYearly](#macrogermanyretailsaleyearly)
- [MacroGermanyTradeAdjusted](#macrogermanytradeadjusted)
- [MacroGermanyZEW](#macrogermanyzew)
- [MacroInfoWS](#macroinfows)
- [MacroJapanBankRate](#macrojapanbankrate)
- [MacroJapanCPIYearly](#macrojapancpiyearly)
- [MacroJapanCore](#macrojapancore)
- [MacroJapanCoreCPIYearly](#macrojapancorecpiyearly)
- [MacroJapanHeadIndicator](#macrojapanheadindicator)
- [MacroJapanUnemploymentRate](#macrojapanunemploymentrate)
- [MacroRMBDeposit](#macrormbdeposit)
- [MacroRMBLoan](#macrormbloan)
- [MacroShippingBCI](#macroshippingbci)
- [MacroShippingBCTI](#macroshippingbcti)
- [MacroShippingBDI](#macroshippingbdi)
- [MacroShippingBPI](#macroshippingbpi)
- [MacroStockFinance](#macrostockfinance)
- [MacroSwissBankRate](#macroswissbankrate)
- [MacroSwissCPIYearly](#macroswisscpiyearly)
- [MacroSwissCore](#macroswisscore)
- [MacroSwissGDPQuarterly](#macroswissgdpquarterly)
- [MacroSwissGDPYearly](#macroswissgdpyearly)
- [MacroSwissSVME](#macroswisssvme)
- [MacroSwissTrade](#macroswisstrade)
- [MacroUKBankRate](#macroukbankrate)
- [MacroUKCPIMonthly](#macroukcpimonthly)
- [MacroUKCPIYearly](#macroukcpiyearly)
- [MacroUKCore](#macroukcore)
- [MacroUKCoreCPIMonthly](#macroukcorecpimonthly)
- [MacroUKCoreCPIYearly](#macroukcorecpiyearly)
- [MacroUKGDPQuarterly](#macroukgdpquarterly)
- [MacroUKGDPYearly](#macroukgdpyearly)
- [MacroUKHalifaxMonthly](#macroukhalifaxmonthly)
- [MacroUKHalifaxYearly](#macroukhalifaxyearly)
- [MacroUKRetailMonthly](#macroukretailmonthly)
- [MacroUKRetailYearly](#macroukretailyearly)
- [MacroUKRightmoveMonthly](#macroukrightmovemonthly)
- [MacroUKRightmoveYearly](#macroukrightmoveyearly)
- [MacroUKTrade](#macrouktrade)
- [MacroUKUnemploymentRate](#macroukunemploymentrate)
- [MacroUsaAdpEmployment](#macrousaadpemployment)
- [MacroUsaApiCrudeStock](#macrousaapicrudestock)
- [MacroUsaBuildingPermits](#macrousabuildingpermits)
- [MacroUsaBusinessInventories](#macrousabusinessinventories)
- [MacroUsaCbConsumerConfidence](#macrousacbconsumerconfidence)
- [MacroUsaCftcCHolding](#macrousacftccholding)
- [MacroUsaCftcMerchantCurrencyHolding](#macrousacftcmerchantcurrencyholding)
- [MacroUsaCftcMerchantGoodsHolding](#macrousacftcmerchantgoodsholding)
- [MacroUsaCftcNcHolding](#macrousacftcncholding)
- [MacroUsaCmeMerchantGoodsHolding](#macrousacmemerchantgoodsholding)
- [MacroUsaCoreCpiMonthly](#macrousacorecpimonthly)
- [MacroUsaCorePcePrice](#macrousacorepceprice)
- [MacroUsaCorePpi](#macrousacoreppi)
- [MacroUsaCpiMonthly](#macrousacpimonthly)
- [MacroUsaCpiYoy](#macrousacpiyoy)
- [MacroUsaCrudeInner](#macrousacrudeinner)
- [MacroUsaCurrentAccount](#macrousacurrentaccount)
- [MacroUsaDurableGoodsOrders](#macrousadurablegoodsorders)
- [MacroUsaEiaCrudeRate](#macrousaeiacruderate)
- [MacroUsaExistHomeSales](#macrousaexisthomesales)
- [MacroUsaExportPrice](#macrousaexportprice)
- [MacroUsaFactoryOrders](#macrousafactoryorders)
- [MacroUsaGdpMonthly](#macrousagdpmonthly)
- [MacroUsaHousePriceIndex](#macrousahousepriceindex)
- [MacroUsaHouseStarts](#macrousahousestarts)
- [MacroUsaImportPrice](#macrousaimportprice)
- [MacroUsaIndustrialProduction](#macrousaindustrialproduction)
- [MacroUsaInitialJobless](#macrousainitialjobless)
- [MacroUsaIsmNonPmi](#macrousaismnonpmi)
- [MacroUsaIsmPmi](#macrousaismpmi)
- [MacroUsaJobCuts](#macrousajobcuts)
- [MacroUsaLmci](#macrousalmci)
- [MacroUsaMichiganConsumerSentiment](#macrousamichiganconsumersentiment)
- [MacroUsaNahbHouseMarketIndex](#macrousanahbhousemarketindex)
- [MacroUsaNewHomeSales](#macrousanewhomesales)
- [MacroUsaNfibSmallBusiness](#macrousanfibsmallbusiness)
- [MacroUsaNonFarm](#macrousanonfarm)
- [MacroUsaPendingHomeSales](#macrousapendinghomesales)
- [MacroUsaPersonalSpending](#macrousapersonalspending)
- [MacroUsaPhs](#macrousaphs)
- [MacroUsaPmi](#macrousapmi)
- [MacroUsaPpi](#macrousappi)
- [MacroUsaRealConsumerSpending](#macrousarealconsumerspending)
- [MacroUsaRetailSales](#macrousaretailsales)
- [MacroUsaRigCount](#macrousarigcount)
- [MacroUsaServicesPmi](#macrousaservicespmi)
- [MacroUsaSpcs20](#macrousaspcs20)
- [MacroUsaTradeBalance](#macrousatradebalance)
- [MacroUsaUnemploymentRate](#macrousaunemploymentrate)

---

## CryptoJSSpot

**描述**: CryptoJSSpot 主流加密货币的实时行情数据

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.CryptoJSSpot()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroAustraliaBankRate

**描述**: MacroAustraliaBankRate 东方财富-经济数据-澳大利亚-央行公布利率决议

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroAustraliaBankRate()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroAustraliaCPIQuarterly

**描述**: MacroAustraliaCPIQuarterly 东方财富-经济数据-澳大利亚-消费者物价指数季率

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroAustraliaCPIQuarterly()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroAustraliaCPIYearly

**描述**: MacroAustraliaCPIYearly 东方财富-经济数据-澳大利亚-消费者物价指数年率

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroAustraliaCPIYearly()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroAustraliaPPIQuarterly

**描述**: MacroAustraliaPPIQuarterly 东方财富-经济数据-澳大利亚-生产者物价指数季率

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroAustraliaPPIQuarterly()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroAustraliaRetailRateMonthly

**描述**: MacroAustraliaRetailRateMonthly 东方财富-经济数据-澳大利亚-零售销售月率

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroAustraliaRetailRateMonthly()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroAustraliaTrade

**描述**: MacroAustraliaTrade 东方财富-经济数据-澳大利亚-贸易帐

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroAustraliaTrade()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroAustraliaUnemploymentRate

**描述**: MacroAustraliaUnemploymentRate 东方财富-经济数据-澳大利亚-失业率

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroAustraliaUnemploymentRate()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroBankAustraliaInterestRate

**描述**: MacroBankAustraliaInterestRate 澳洲联储利率决议报告

**数据源**: https://datacenter.jin10.com/reportType/dc_australia_interest_rate_decision

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroBankAustraliaInterestRate()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroBankBrazilInterestRate

**描述**: MacroBankBrazilInterestRate 巴西央行利率决议报告

**数据源**: https://datacenter.jin10.com/reportType/dc_brazil_interest_rate_decision

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroBankBrazilInterestRate()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroBankChinaInterestRate

**描述**: MacroBankChinaInterestRate 中国央行利率决议报告

**数据源**: https://datacenter.jin10.com/reportType/dc_china_interest_rate_decision

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroBankChinaInterestRate()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroBankEnglishInterestRate

**描述**: MacroBankEnglishInterestRate 英国央行利率决议报告

**数据源**: https://datacenter.jin10.com/reportType/dc_english_interest_rate_decision

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroBankEnglishInterestRate()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroBankEuroInterestRate

**描述**: MacroBankEuroInterestRate 欧洲央行利率决议报告

**数据源**: https://datacenter.jin10.com/reportType/dc_interest_rate_decision

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroBankEuroInterestRate()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroBankIndiaInterestRate

**描述**: MacroBankIndiaInterestRate 印度央行利率决议报告

**数据源**: https://datacenter.jin10.com/reportType/dc_india_interest_rate_decision

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroBankIndiaInterestRate()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroBankJapanInterestRate

**描述**: MacroBankJapanInterestRate 日本央行利率决议报告

**数据源**: https://datacenter.jin10.com/reportType/dc_japan_interest_rate_decision

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroBankJapanInterestRate()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroBankNewzealandInterestRate

**描述**: MacroBankNewzealandInterestRate 新西兰联储利率决议报告

**数据源**: https://datacenter.jin10.com/reportType/dc_newzealand_interest_rate_decision

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroBankNewzealandInterestRate()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroBankRussiaInterestRate

**描述**: MacroBankRussiaInterestRate 俄罗斯央行利率决议报告

**数据源**: https://datacenter.jin10.com/reportType/dc_russia_interest_rate_decision

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroBankRussiaInterestRate()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroBankSwitzerlandInterestRate

**描述**: MacroBankSwitzerlandInterestRate 瑞士央行利率决议报告

**数据源**: https://datacenter.jin10.com/reportType/dc_switzerland_interest_rate_decision

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroBankSwitzerlandInterestRate()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroBankUsaInterestRate

**描述**: MacroBankUsaInterestRate 美联储利率决议报告

**数据源**: https://datacenter.jin10.com/reportType/dc_usa_interest_rate_decision

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroBankUsaInterestRate()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroCNBS

**描述**: MacroCNBS 国家金融与发展实验室-中国宏观杠杆率数据

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroCNBS()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroCanadaBankRate

**描述**: MacroCanadaBankRate 东方财富-经济数据-加拿大-央行公布利率决议

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroCanadaBankRate()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroCanadaCPIMonthly

**描述**: MacroCanadaCPIMonthly 东方财富-经济数据-加拿大-消费者物价指数月率

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroCanadaCPIMonthly()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroCanadaCPIYearly

**描述**: MacroCanadaCPIYearly 东方财富-经济数据-加拿大-消费者物价指数年率

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroCanadaCPIYearly()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroCanadaCoreCPIMonthly

**描述**: MacroCanadaCoreCPIMonthly 东方财富-经济数据-加拿大-核心消费者物价指数月率

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroCanadaCoreCPIMonthly()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroCanadaCoreCPIYearly

**描述**: MacroCanadaCoreCPIYearly 东方财富-经济数据-加拿大-核心消费者物价指数年率

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroCanadaCoreCPIYearly()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroCanadaGDPMonthly

**描述**: MacroCanadaGDPMonthly 东方财富-经济数据-加拿大-GDP月率

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroCanadaGDPMonthly()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroCanadaNewHouseRate

**描述**: MacroCanadaNewHouseRate 东方财富-经济数据-加拿大-新屋开工

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroCanadaNewHouseRate()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroCanadaRetailRateMonthly

**描述**: MacroCanadaRetailRateMonthly 东方财富-经济数据-加拿大-零售销售月率

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroCanadaRetailRateMonthly()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroCanadaTrade

**描述**: MacroCanadaTrade 东方财富-经济数据-加拿大-贸易帐

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroCanadaTrade()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroCanadaUnemploymentRate

**描述**: MacroCanadaUnemploymentRate 东方财富-经济数据-加拿大-失业率

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroCanadaUnemploymentRate()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaAgriculturalIndex

**描述**: MacroChinaAgriculturalIndex 农副指数

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaAgriculturalIndex()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaAgriculturalProduct

**描述**: MacroChinaAgriculturalProduct 农产品批发价格200指数

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaAgriculturalProduct()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaAuReport

**描述**: MacroChinaAuReport 上海黄金交易所报告

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaAuReport()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaBDTIIndex

**描述**: MacroChinaBDTIIndex 原油运输指数

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaBDTIIndex()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaBSIIndex

**描述**: MacroChinaBSIIndex 超灵便型船运价指数

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaBSIIndex()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaBankFinancing

**描述**: MacroChinaBankFinancing 银行理财产品发行数量

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaBankFinancing()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaCPI

**描述**: MacroChinaCPI 东方财富-中国居民消费价格指数

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaCPI()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaCPIMonthly

**描述**: MacroChinaCPIMonthly 金十数据中心-中国CPI月率报告

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaCPIMonthly()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaCPIYearly

**描述**: MacroChinaCPIYearly 金十数据中心-中国CPI年率报告

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaCPIYearly()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaCentralBankBalance

**描述**: MacroChinaCentralBankBalance 获取央行货币当局资产负债

**数据源**: https://finance.sina.com.cn/mac/#fininfo-8-0-31-2

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaCentralBankBalance()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaCommodityPriceIndex

**描述**: MacroChinaCommodityPriceIndex 大宗商品价格指数

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaCommodityPriceIndex()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaConstructionIndex

**描述**: MacroChinaConstructionIndex 建材指数

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaConstructionIndex()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaConstructionPriceIndex

**描述**: MacroChinaConstructionPriceIndex 建材价格指数

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaConstructionPriceIndex()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaConsumerGoodsRetail

**描述**: MacroChinaConsumerGoodsRetail 东方财富-社会消费品零售总额

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaConsumerGoodsRetail()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaCxPMIYearly

**描述**: MacroChinaCxPMIYearly 金十数据中心-中国财新制造业PMI终值报告

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaCxPMIYearly()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaCxServicesPMIYearly

**描述**: MacroChinaCxServicesPMIYearly 金十数据中心-中国财新服务业PMI报告

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaCxServicesPMIYearly()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaCzsr

**描述**: MacroChinaCzsr 东方财富-财政收入

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaCzsr()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaDailyEnergy

**描述**: MacroChinaDailyEnergy 金十数据-日度能源数据

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaDailyEnergy()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaEnergyIndex

**描述**: MacroChinaEnergyIndex 能源指数

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaEnergyIndex()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaEnterpriseBoomIndex

**描述**: MacroChinaEnterpriseBoomIndex 中国-企业景气及企业家信心指数

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaEnterpriseBoomIndex()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaExportsYoY

**描述**: MacroChinaExportsYoY 金十数据中心-中国以美元计算出口年率报告

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaExportsYoY()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaFDI

**描述**: MacroChinaFDI 东方财富-经济数据一览-中国-外商直接投资数据

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaFDI()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaForeignExchangeGold

**描述**: MacroChinaForeignExchangeGold 获取央行黄金和外汇储备

**数据源**: https://finance.sina.com.cn/mac/#fininfo-5-0-31-2

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaForeignExchangeGold()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaFreightIndex

**描述**: MacroChinaFreightIndex 获取航贸运价指数

**数据源**: https://finance.sina.com.cn/mac/#industry-22-0-31-2

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaFreightIndex()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaFxGold

**描述**: MacroChinaFxGold 东方财富-外汇和黄金储备

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaFxGold()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaFxReservesYearly

**描述**: MacroChinaFxReservesYearly 金十数据中心-中国外汇储备报告

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaFxReservesYearly()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaGDP

**描述**: MacroChinaGDP 东方财富-中国国内生产总值

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaGDP()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaGDPYearly

**描述**: MacroChinaGDPYearly 金十数据中心-中国GDP年率报告

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaGDPYearly()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaGdzctz

**描述**: MacroChinaGdzctz 东方财富-中国城镇固定资产投资

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaGdzctz()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaGyzjz

**描述**: MacroChinaGyzjz 东方财富-工业增加值增长

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaGyzjz()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaHKBuildingAmount

**描述**: MacroChinaHKBuildingAmount 东方财富-经济数据一览-中国香港-香港楼宇买卖合约成交金额

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaHKBuildingAmount()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaHKBuildingVolume

**描述**: MacroChinaHKBuildingVolume 东方财富-经济数据一览-中国香港-香港楼宇买卖合约数量

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaHKBuildingVolume()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaHKCPI

**描述**: MacroChinaHKCPI 东方财富-经济数据一览-中国香港-消费者物价指数

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaHKCPI()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaHKCPIRatio

**描述**: MacroChinaHKCPIRatio 东方财富-经济数据一览-中国香港-消费者物价指数年率

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaHKCPIRatio()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaHKCore

**描述**: MacroChinaHKCore 东方财富-经济数据一览-中国香港-核心指标数据

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaHKCore("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaHKGBP

**描述**: MacroChinaHKGBP 东方财富-经济数据一览-中国香港-香港GDP

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaHKGBP()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaHKGBPRatio

**描述**: MacroChinaHKGBPRatio 东方财富-经济数据一览-中国香港-香港GDP同比

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaHKGBPRatio()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaHKMarketInfo

**描述**: MacroChinaHKMarketInfo 东方财富-沪港通资金流向

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaHKMarketInfo()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaHKPPI

**描述**: MacroChinaHKPPI 东方财富-经济数据一览-中国香港-香港制造业PPI年率

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaHKPPI()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaHKRateOfUnemployment

**描述**: MacroChinaHKRateOfUnemployment 东方财富-经济数据一览-中国香港-失业率

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaHKRateOfUnemployment()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaHKTradeDiffRatio

**描述**: MacroChinaHKTradeDiffRatio 东方财富-经济数据一览-中国香港-香港商品贸易差额年率

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaHKTradeDiffRatio()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaHgjck

**描述**: MacroChinaHgjck 东方财富-海关进出口增减情况一览表

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaHgjck()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaImportsYoY

**描述**: MacroChinaImportsYoY 金十数据中心-中国以美元计算进口年率报告

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaImportsYoY()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaIndustrialProductionYoY

**描述**: MacroChinaIndustrialProductionYoY 金十数据中心-中国规模以上工业增加值年率报告

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaIndustrialProductionYoY()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaInsurance

**描述**: MacroChinaInsurance 获取保险业经营情况

**数据源**: https://finance.sina.com.cn/mac/#fininfo-19-0-31-3

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaInsurance()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaInsuranceIncome

**描述**: MacroChinaInsuranceIncome 原保险保费收入

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaInsuranceIncome()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaInternationalTourismFx

**描述**: MacroChinaInternationalTourismFx 获取国际旅游外汇收入构成

**数据源**: https://finance.sina.com.cn/mac/#industry-15-0-31-3

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaInternationalTourismFx()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaLPIIndex

**描述**: MacroChinaLPIIndex 物流景气指数

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaLPIIndex()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaLPR

**描述**: MacroChinaLPR LPR品种详细数据

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaLPR()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaM2Yearly

**描述**: MacroChinaM2Yearly 金十数据中心-中国M2货币供应年率报告

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaM2Yearly()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaMarketMarginSH

**描述**: MacroChinaMarketMarginSH 上海融资融券数据

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaMarketMarginSH()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaMarketMarginSZ

**描述**: MacroChinaMarketMarginSZ 深圳融资融券数据

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaMarketMarginSZ()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaMobileNumber

**描述**: MacroChinaMobileNumber 移动电话用户数

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaMobileNumber()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaMoneySupply

**描述**: MacroChinaMoneySupply 东方财富-货币供应量

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaMoneySupply()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaNBSNation

**描述**: MacroChinaNBSNation 国家统计局全国数据通用接口

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| kind | string | - |
| path | string | - |
| period | string | 周期：daily/weekly/monthly |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaNBSNation("", "", "daily")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaNBSRegion

**描述**: MacroChinaNBSRegion 国家统计局地区数据通用接口

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| kind | string | - |
| path | string | - |
| indicator | string | 指标类型 |
| region | string | - |
| period | string | 周期：daily/weekly/monthly |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaNBSRegion("", "", "", "", "daily")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaNationalTaxReceipts

**描述**: MacroChinaNationalTaxReceipts 中国-全国税收收入

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaNationalTaxReceipts()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaNewFinancialCredit

**描述**: MacroChinaNewFinancialCredit 中国-新增信贷数据

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaNewFinancialCredit()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaNewHousePrice

**描述**: MacroChinaNewHousePrice 中国-新房价指数

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| cityFirst | string | - |
| citySecond | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaNewHousePrice("", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaNonManPMI

**描述**: MacroChinaNonManPMI 金十数据中心-中国官方非制造业PMI报告

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaNonManPMI()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaPMI

**描述**: MacroChinaPMI 东方财富-中国采购经理人指数

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaPMI()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaPMIYearly

**描述**: MacroChinaPMIYearly 金十数据中心-中国官方制造业PMI报告

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaPMIYearly()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaPPI

**描述**: MacroChinaPPI 东方财富-中国工业品出厂价格指数

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaPPI()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaPPIYearly

**描述**: MacroChinaPPIYearly 金十数据中心-中国PPI年率报告

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaPPIYearly()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaPassengerLoadFactor

**描述**: MacroChinaPassengerLoadFactor 获取民航客座率及载运率

**数据源**: https://finance.sina.com.cn/mac/#industry-20-0-31-1

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaPassengerLoadFactor()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaPostalTelecommunicational

**描述**: MacroChinaPostalTelecommunicational 获取邮电业务基本情况

**数据源**: https://finance.sina.com.cn/mac/#industry-11-0-31-1

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaPostalTelecommunicational()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaQyspjg

**描述**: MacroChinaQyspjg 东方财富-经济数据一览-中国-企业商品价格指数

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaQyspjg()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaRMB

**描述**: MacroChinaRMB 人民币汇率中间价

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaRMB()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaRealEstate

**描述**: MacroChinaRealEstate 获取国房景气指数

**数据源**: https://data.eastmoney.com/cjsj/hyzs_list_EMM00121987.html

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaRealEstate()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaReserveRequirementRatio

**描述**: MacroChinaReserveRequirementRatio 存款准备金率

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaReserveRequirementRatio()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaRetailPriceIndex

**描述**: MacroChinaRetailPriceIndex 获取商品零售价格指数

**数据源**: https://finance.sina.com.cn/mac/#price-12-0-31-1

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaRetailPriceIndex()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaShiborAll

**描述**: MacroChinaShiborAll 上海银行同业拆借利率(SHIBOR)

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaShiborAll()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaShrzgm

**描述**: MacroChinaShrzgm 商务数据中心-国内贸易-社会融资规模增量统计

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaShrzgm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaSocietyElectricity

**描述**: MacroChinaSocietyElectricity 获取全社会用电分类情况表

**数据源**: https://finance.sina.com.cn/mac/#industry-6-0-31-1

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaSocietyElectricity()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaSocietyTrafficVolume

**描述**: MacroChinaSocietyTrafficVolume 获取全社会客货运输量

**数据源**: https://finance.sina.com.cn/mac/#industry-10-0-31-1

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaSocietyTrafficVolume()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaStockMarketCap

**描述**: MacroChinaStockMarketCap 东方财富-全国股票交易统计表

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaStockMarketCap()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaSupplyOfMoney

**描述**: MacroChinaSupplyOfMoney 获取货币供应量

**数据源**: https://finance.sina.com.cn/mac/#fininfo-1-0-31-1

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaSupplyOfMoney()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaTradeBalance

**描述**: MacroChinaTradeBalance 金十数据中心-中国以美元计算贸易帐报告

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaTradeBalance()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaUrbanUnemployment

**描述**: MacroChinaUrbanUnemployment 国家统计局-月度数据-城镇调查失业率

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaUrbanUnemployment()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaVegetableBasket

**描述**: MacroChinaVegetableBasket 菜篮子产品批发价格指数

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaVegetableBasket()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaWbck

**描述**: MacroChinaWbck 东方财富-本外币存款

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaWbck()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaWhxd

**描述**: MacroChinaWhxd 东方财富-外汇贷款数据

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaWhxd()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaXfzxx

**描述**: MacroChinaXfzxx 东方财富-消费者信心指数

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaXfzxx()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaYwElectronicIndex

**描述**: MacroChinaYwElectronicIndex 义乌小商品指数-电子元器件

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroChinaYwElectronicIndex()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroConsGold

**描述**: MacroConsGold 全球最大黄金ETF-SPDR Gold Trust持仓报告

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroConsGold()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroConsOpecMonth

**描述**: MacroConsOpecMonth 欧佩克报告-月度

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroConsOpecMonth()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroConsOpecMonthWithLimit

**描述**: MacroConsOpecMonthWithLimit 欧佩克报告-月度（限制获取数量）

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| limit | int | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroConsOpecMonthWithLimit(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroConsSilver

**描述**: MacroConsSilver 全球最大白银ETF-iShares Silver Trust持仓报告

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroConsSilver()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroEuroCPIMoM

**描述**: MacroEuroCPIMoM 金十数据中心-欧元区CPI月率报告

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroEuroCPIMoM()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroEuroCPIYoY

**描述**: MacroEuroCPIYoY 金十数据中心-欧元区CPI年率报告

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroEuroCPIYoY()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroEuroCurrentAccountMoM

**描述**: MacroEuroCurrentAccountMoM 金十数据中心-欧元区经常帐报告

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroEuroCurrentAccountMoM()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroEuroEmploymentChangeQoQ

**描述**: MacroEuroEmploymentChangeQoQ 金十数据中心-欧元区季调后就业人数季率报告

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroEuroEmploymentChangeQoQ()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroEuroGDPYoY

**描述**: MacroEuroGDPYoY 金十数据中心-欧元区季度GDP年率报告

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroEuroGDPYoY()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroEuroIndustrialProductionMoM

**描述**: MacroEuroIndustrialProductionMoM 金十数据中心-欧元区工业产出月率报告

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroEuroIndustrialProductionMoM()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroEuroLMEHolding

**描述**: MacroEuroLMEHolding 金十数据中心-伦敦金属交易所(LME)-持仓报告

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroEuroLMEHolding()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroEuroLMEStock

**描述**: MacroEuroLMEStock 金十数据中心-伦敦金属交易所(LME)-库存报告

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroEuroLMEStock()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroEuroManufacturingPMI

**描述**: MacroEuroManufacturingPMI 金十数据中心-欧元区制造业PMI初值报告

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroEuroManufacturingPMI()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroEuroPPIMoM

**描述**: MacroEuroPPIMoM 金十数据中心-欧元区PPI月率报告

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroEuroPPIMoM()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroEuroRetailSalesMoM

**描述**: MacroEuroRetailSalesMoM 金十数据中心-欧元区零售销售月率报告

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroEuroRetailSalesMoM()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroEuroSentixInvestorConfidence

**描述**: MacroEuroSentixInvestorConfidence 金十数据中心-欧元区Sentix投资者信心指数报告

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroEuroSentixInvestorConfidence()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroEuroServicesPMI

**描述**: MacroEuroServicesPMI 金十数据中心-欧元区服务业PMI终值报告

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroEuroServicesPMI()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroEuroTradeBalance

**描述**: MacroEuroTradeBalance 金十数据中心-欧元区未季调贸易帐报告

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroEuroTradeBalance()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroEuroUnemploymentRateMoM

**描述**: MacroEuroUnemploymentRateMoM 金十数据中心-欧元区失业率报告

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroEuroUnemploymentRateMoM()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroEuroZEWEconomicSentiment

**描述**: MacroEuroZEWEconomicSentiment 金十数据中心-欧元区ZEW经济景气指数报告

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroEuroZEWEconomicSentiment()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroFXSentiment

**描述**: MacroFXSentiment 金十数据-外汇-投机情绪报告

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroFXSentiment("20230101", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroGermanyCPIMonthly

**描述**: MacroGermanyCPIMonthly 东方财富-经济数据-德国-消费者物价指数月率终值

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroGermanyCPIMonthly()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroGermanyCPIYearly

**描述**: MacroGermanyCPIYearly 东方财富-经济数据-德国-消费者物价指数年率终值

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroGermanyCPIYearly()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroGermanyCore

**描述**: MacroGermanyCore 东方财富-经济数据一览-德国-核心指标数据

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroGermanyCore("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroGermanyGDP

**描述**: MacroGermanyGDP 东方财富-经济数据-德国-GDP

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroGermanyGDP()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroGermanyIFO

**描述**: MacroGermanyIFO 东方财富-经济数据-德国-IFO商业景气指数

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroGermanyIFO()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroGermanyRetailSaleMonthly

**描述**: MacroGermanyRetailSaleMonthly 东方财富-经济数据-德国-实际零售销售月率

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroGermanyRetailSaleMonthly()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroGermanyRetailSaleYearly

**描述**: MacroGermanyRetailSaleYearly 东方财富-经济数据-德国-实际零售销售年率

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroGermanyRetailSaleYearly()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroGermanyTradeAdjusted

**描述**: MacroGermanyTradeAdjusted 东方财富-经济数据-德国-贸易帐(季调后)

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroGermanyTradeAdjusted()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroGermanyZEW

**描述**: MacroGermanyZEW 东方财富-经济数据-德国-ZEW经济景气指数

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroGermanyZEW()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroInfoWS

**描述**: MacroInfoWS 华尔街见闻-日历-宏观

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroInfoWS("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroJapanBankRate

**描述**: MacroJapanBankRate 东方财富-经济数据-日本-央行公布利率决议

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroJapanBankRate()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroJapanCPIYearly

**描述**: MacroJapanCPIYearly 东方财富-经济数据-日本-全国消费者物价指数年率

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroJapanCPIYearly()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroJapanCore

**描述**: MacroJapanCore 东方财富-经济数据一览-日本-核心指标数据

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroJapanCore("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroJapanCoreCPIYearly

**描述**: MacroJapanCoreCPIYearly 东方财富-经济数据-日本-全国核心消费者物价指数年率

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroJapanCoreCPIYearly()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroJapanHeadIndicator

**描述**: MacroJapanHeadIndicator 东方财富-经济数据-日本-领先指标终值

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroJapanHeadIndicator()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroJapanUnemploymentRate

**描述**: MacroJapanUnemploymentRate 东方财富-经济数据-日本-失业率

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroJapanUnemploymentRate()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroRMBDeposit

**描述**: MacroRMBDeposit 同花顺-数据中心-宏观数据-人民币存款余额

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroRMBDeposit()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroRMBLoan

**描述**: MacroRMBLoan 同花顺-数据中心-宏观数据-新增人民币贷款

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroRMBLoan()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroShippingBCI

**描述**: MacroShippingBCI 海岬型运费指数(BCI)

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroShippingBCI()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroShippingBCTI

**描述**: MacroShippingBCTI 成品油运输指数(BCTI)

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroShippingBCTI()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroShippingBDI

**描述**: MacroShippingBDI 波罗的海干散货指数(BDI)

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroShippingBDI()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroShippingBPI

**描述**: MacroShippingBPI 巴拿马型运费指数(BPI)

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroShippingBPI()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroStockFinance

**描述**: MacroStockFinance 同花顺-数据中心-宏观数据-股票筹资

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroStockFinance()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroSwissBankRate

**描述**: MacroSwissBankRate 东方财富-经济数据-瑞士-央行公布利率决议

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroSwissBankRate()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroSwissCPIYearly

**描述**: MacroSwissCPIYearly 东方财富-经济数据-瑞士-消费者物价指数年率

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroSwissCPIYearly()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroSwissCore

**描述**: MacroSwissCore 东方财富-经济数据一览-瑞士-核心指标数据

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroSwissCore("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroSwissGDPQuarterly

**描述**: MacroSwissGDPQuarterly 东方财富-经济数据-瑞士-GDP季率

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroSwissGDPQuarterly()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroSwissGDPYearly

**描述**: MacroSwissGDPYearly 东方财富-经济数据-瑞士-GDP年率

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroSwissGDPYearly()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroSwissSVME

**描述**: MacroSwissSVME 东方财富-经济数据-瑞士-SVME采购经理人指数

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroSwissSVME()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroSwissTrade

**描述**: MacroSwissTrade 东方财富-经济数据-瑞士-贸易帐

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroSwissTrade()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUKBankRate

**描述**: MacroUKBankRate 东方财富-经济数据-英国-央行公布利率决议

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUKBankRate()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUKCPIMonthly

**描述**: MacroUKCPIMonthly 东方财富-经济数据-英国-消费者物价指数月率

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUKCPIMonthly()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUKCPIYearly

**描述**: MacroUKCPIYearly 东方财富-经济数据-英国-消费者物价指数年率

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUKCPIYearly()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUKCore

**描述**: MacroUKCore 东方财富-经济数据一览-英国-核心指标数据

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUKCore("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUKCoreCPIMonthly

**描述**: MacroUKCoreCPIMonthly 东方财富-经济数据-英国-核心消费者物价指数月率

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUKCoreCPIMonthly()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUKCoreCPIYearly

**描述**: MacroUKCoreCPIYearly 东方财富-经济数据-英国-核心消费者物价指数年率

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUKCoreCPIYearly()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUKGDPQuarterly

**描述**: MacroUKGDPQuarterly 东方财富-经济数据-英国-GDP季率初值

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUKGDPQuarterly()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUKGDPYearly

**描述**: MacroUKGDPYearly 东方财富-经济数据-英国-GDP年率初值

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUKGDPYearly()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUKHalifaxMonthly

**描述**: MacroUKHalifaxMonthly 东方财富-经济数据-英国-Halifax房价指数月率

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUKHalifaxMonthly()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUKHalifaxYearly

**描述**: MacroUKHalifaxYearly 东方财富-经济数据-英国-Halifax房价指数年率

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUKHalifaxYearly()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUKRetailMonthly

**描述**: MacroUKRetailMonthly 东方财富-经济数据-英国-零售销售月率

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUKRetailMonthly()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUKRetailYearly

**描述**: MacroUKRetailYearly 东方财富-经济数据-英国-零售销售年率

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUKRetailYearly()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUKRightmoveMonthly

**描述**: MacroUKRightmoveMonthly 东方财富-经济数据-英国-Rightmove房价指数月率

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUKRightmoveMonthly()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUKRightmoveYearly

**描述**: MacroUKRightmoveYearly 东方财富-经济数据-英国-Rightmove房价指数年率

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUKRightmoveYearly()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUKTrade

**描述**: MacroUKTrade 东方财富-经济数据-英国-贸易帐

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUKTrade()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUKUnemploymentRate

**描述**: MacroUKUnemploymentRate 东方财富-经济数据-英国-失业率

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUKUnemploymentRate()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUsaAdpEmployment

**描述**: MacroUsaAdpEmployment 美国ADP就业人数报告

**数据源**: https://datacenter.jin10.com/reportType/dc_usa_adp_employment

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUsaAdpEmployment()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUsaApiCrudeStock

**描述**: MacroUsaApiCrudeStock 美国API原油库存报告

**数据源**: https://datacenter.jin10.com/reportType/dc_usa_api_crude_stock

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUsaApiCrudeStock()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUsaBuildingPermits

**描述**: MacroUsaBuildingPermits 美国营建许可总数报告

**数据源**: https://datacenter.jin10.com/reportType/dc_usa_building_permits

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUsaBuildingPermits()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUsaBusinessInventories

**描述**: MacroUsaBusinessInventories 美国商业库存月率报告

**数据源**: https://datacenter.jin10.com/reportType/dc_usa_business_inventories

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUsaBusinessInventories()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUsaCbConsumerConfidence

**描述**: MacroUsaCbConsumerConfidence 美国谘商会消费者信心指数报告

**数据源**: https://datacenter.jin10.com/reportType/dc_usa_cb_consumer_confidence

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUsaCbConsumerConfidence()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUsaCftcCHolding

**描述**: MacroUsaCftcCHolding 美国商品期货交易委员会CFTC商品类非商业持仓报告

**数据源**: https://datacenter.jin10.com/reportType/dc_cftc_c_report

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUsaCftcCHolding()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUsaCftcMerchantCurrencyHolding

**描述**: MacroUsaCftcMerchantCurrencyHolding 美国商品期货交易委员会CFTC外汇类商业持仓报告

**数据源**: https://datacenter.jin10.com/reportType/dc_cftc_merchant_currency

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUsaCftcMerchantCurrencyHolding()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUsaCftcMerchantGoodsHolding

**描述**: MacroUsaCftcMerchantGoodsHolding 美国商品期货交易委员会CFTC商品类商业持仓报告

**数据源**: https://datacenter.jin10.com/reportType/dc_cftc_merchant_goods

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUsaCftcMerchantGoodsHolding()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUsaCftcNcHolding

**描述**: MacroUsaCftcNcHolding 美国商品期货交易委员会CFTC外汇类非商业持仓报告

**数据源**: https://datacenter.jin10.com/reportType/dc_cftc_nc_report

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUsaCftcNcHolding()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUsaCmeMerchantGoodsHolding

**描述**: MacroUsaCmeMerchantGoodsHolding CME贵金属报告

**数据源**: https://datacenter.jin10.com/org

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUsaCmeMerchantGoodsHolding()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUsaCoreCpiMonthly

**描述**: MacroUsaCoreCpiMonthly 美国核心CPI月率报告

**数据源**: https://datacenter.jin10.com/reportType/dc_usa_core_cpi

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUsaCoreCpiMonthly()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUsaCorePcePrice

**描述**: MacroUsaCorePcePrice 美国核心PCE物价指数年率报告

**数据源**: https://datacenter.jin10.com/reportType/dc_usa_core_pce_price

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUsaCorePcePrice()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUsaCorePpi

**描述**: MacroUsaCorePpi 美国核心生产者物价指数报告

**数据源**: https://datacenter.jin10.com/reportType/dc_usa_core_ppi

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUsaCorePpi()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUsaCpiMonthly

**描述**: MacroUsaCpiMonthly 美国CPI月率报告

**数据源**: https://datacenter.jin10.com/reportType/dc_usa_cpi

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUsaCpiMonthly()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUsaCpiYoy

**描述**: MacroUsaCpiYoy 美国CPI年率报告（东方财富）

**数据源**: https://data.eastmoney.com/cjsj/foreign_0_12.html

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUsaCpiYoy()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUsaCrudeInner

**描述**: MacroUsaCrudeInner 美国原油产量报告

**数据源**: https://datacenter.jin10.com/reportType/dc_eia_crude_oil_produce

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUsaCrudeInner()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUsaCurrentAccount

**描述**: MacroUsaCurrentAccount 美国经常账报告

**数据源**: https://datacenter.jin10.com/reportType/dc_usa_current_account

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUsaCurrentAccount()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUsaDurableGoodsOrders

**描述**: MacroUsaDurableGoodsOrders 美国耐用品订单月率报告

**数据源**: https://datacenter.jin10.com/reportType/dc_usa_durable_goods_orders

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUsaDurableGoodsOrders()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUsaEiaCrudeRate

**描述**: MacroUsaEiaCrudeRate 美国EIA原油库存报告

**数据源**: https://datacenter.jin10.com/reportType/dc_usa_eia_crude_rate

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUsaEiaCrudeRate()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUsaExistHomeSales

**描述**: MacroUsaExistHomeSales 美国成屋销售总数年化报告

**数据源**: https://datacenter.jin10.com/reportType/dc_usa_exist_home_sales

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUsaExistHomeSales()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUsaExportPrice

**描述**: MacroUsaExportPrice 美国出口价格指数报告

**数据源**: https://datacenter.jin10.com/reportType/dc_usa_export_price

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUsaExportPrice()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUsaFactoryOrders

**描述**: MacroUsaFactoryOrders 美国工厂订单月率报告

**数据源**: https://datacenter.jin10.com/reportType/dc_usa_factory_orders

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUsaFactoryOrders()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUsaGdpMonthly

**描述**: MacroUsaGdpMonthly 美国国内生产总值(GDP)报告

**数据源**: https://datacenter.jin10.com/reportType/dc_usa_gdp

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUsaGdpMonthly()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUsaHousePriceIndex

**描述**: MacroUsaHousePriceIndex 美国FHFA房价指数月率报告

**数据源**: https://datacenter.jin10.com/reportType/dc_usa_house_price_index

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUsaHousePriceIndex()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUsaHouseStarts

**描述**: MacroUsaHouseStarts 美国新屋开工总数年化报告

**数据源**: https://datacenter.jin10.com/reportType/dc_usa_house_starts

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUsaHouseStarts()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUsaImportPrice

**描述**: MacroUsaImportPrice 美国进口物价指数报告

**数据源**: https://datacenter.jin10.com/reportType/dc_usa_import_price

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUsaImportPrice()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUsaIndustrialProduction

**描述**: MacroUsaIndustrialProduction 美国工业产出月率报告

**数据源**: https://datacenter.jin10.com/reportType/dc_usa_industrial_production

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUsaIndustrialProduction()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUsaInitialJobless

**描述**: MacroUsaInitialJobless 美国初请失业金人数报告

**数据源**: https://datacenter.jin10.com/reportType/dc_usa_initial_jobless

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUsaInitialJobless()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUsaIsmNonPmi

**描述**: MacroUsaIsmNonPmi 美国ISM非制造业PMI报告

**数据源**: https://datacenter.jin10.com/reportType/dc_usa_ism_non_pmi

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUsaIsmNonPmi()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUsaIsmPmi

**描述**: MacroUsaIsmPmi 美国ISM制造业PMI报告

**数据源**: https://datacenter.jin10.com/reportType/dc_usa_ism_pmi

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUsaIsmPmi()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUsaJobCuts

**描述**: MacroUsaJobCuts 美国挑战者企业裁员人数报告

**数据源**: https://datacenter.jin10.com/reportType/dc_usa_job_cuts

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUsaJobCuts()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUsaLmci

**描述**: MacroUsaLmci 美联储劳动力市场状况指数报告

**数据源**: https://datacenter.jin10.com/reportType/dc_usa_lmci

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUsaLmci()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUsaMichiganConsumerSentiment

**描述**: MacroUsaMichiganConsumerSentiment 美国密歇根大学消费者信心指数初值报告

**数据源**: https://datacenter.jin10.com/reportType/dc_usa_michigan_consumer_sentiment

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUsaMichiganConsumerSentiment()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUsaNahbHouseMarketIndex

**描述**: MacroUsaNahbHouseMarketIndex 美国NAHB房产市场指数报告

**数据源**: https://datacenter.jin10.com/reportType/dc_usa_nahb_house_market_index

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUsaNahbHouseMarketIndex()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUsaNewHomeSales

**描述**: MacroUsaNewHomeSales 美国新屋销售总数年化报告

**数据源**: https://datacenter.jin10.com/reportType/dc_usa_new_home_sales

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUsaNewHomeSales()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUsaNfibSmallBusiness

**描述**: MacroUsaNfibSmallBusiness 美国NFIB小型企业信心指数报告

**数据源**: https://datacenter.jin10.com/reportType/dc_usa_nfib_small_business

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUsaNfibSmallBusiness()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUsaNonFarm

**描述**: MacroUsaNonFarm 美国非农就业人数报告

**数据源**: https://datacenter.jin10.com/reportType/dc_usa_non_farm

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUsaNonFarm()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUsaPendingHomeSales

**描述**: MacroUsaPendingHomeSales 美国成屋签约销售指数月率报告

**数据源**: https://datacenter.jin10.com/reportType/dc_usa_pending_home_sales

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUsaPendingHomeSales()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUsaPersonalSpending

**描述**: MacroUsaPersonalSpending 美国个人支出月率报告

**数据源**: https://datacenter.jin10.com/reportType/dc_usa_personal_spending

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUsaPersonalSpending()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUsaPhs

**描述**: MacroUsaPhs 美国未决房屋销售月率报告（东方财富）

**数据源**: https://data.eastmoney.com/cjsj/foreign_0_5.html

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUsaPhs()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUsaPmi

**描述**: MacroUsaPmi 美国Markit制造业PMI报告

**数据源**: https://datacenter.jin10.com/reportType/dc_usa_pmi

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUsaPmi()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUsaPpi

**描述**: MacroUsaPpi 美国生产者物价指数报告

**数据源**: https://datacenter.jin10.com/reportType/dc_usa_ppi

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUsaPpi()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUsaRealConsumerSpending

**描述**: MacroUsaRealConsumerSpending 美国实际个人消费支出季率初值报告

**数据源**: https://datacenter.jin10.com/reportType/dc_usa_real_consumer_spending

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUsaRealConsumerSpending()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUsaRetailSales

**描述**: MacroUsaRetailSales 美国零售销售月率报告

**数据源**: https://datacenter.jin10.com/reportType/dc_usa_retail_sales

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUsaRetailSales()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUsaRigCount

**描述**: MacroUsaRigCount 贝克休斯钻井平台报告

**数据源**: https://datacenter.jin10.com/reportType/dc_rig_count_summary

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUsaRigCount()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUsaServicesPmi

**描述**: MacroUsaServicesPmi 美国Markit服务业PMI初值报告

**数据源**: https://datacenter.jin10.com/reportType/dc_usa_services_pmi

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUsaServicesPmi()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUsaSpcs20

**描述**: MacroUsaSpcs20 美国S&P/CS20座大城市房价指数年率报告

**数据源**: https://datacenter.jin10.com/reportType/dc_usa_spcs20

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUsaSpcs20()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUsaTradeBalance

**描述**: MacroUsaTradeBalance 美国贸易帐报告

**数据源**: https://datacenter.jin10.com/reportType/dc_usa_trade_balance

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUsaTradeBalance()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroUsaUnemploymentRate

**描述**: MacroUsaUnemploymentRate 美国失业率报告

**数据源**: https://datacenter.jin10.com/reportType/dc_usa_unemployment_rate

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
	"github.com/BlakeLiAFK/akshare/economic"
)

func main() {
	data, err := economic.MacroUsaUnemploymentRate()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

