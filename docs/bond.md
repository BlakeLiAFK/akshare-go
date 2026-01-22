# bond 模块

> 本模块共有 42 个接口

## 目录

- [BondBuyBackHistEM](#bondbuybackhistem)
- [BondCBAdjLogsJSL](#bondcbadjlogsjsl)
- [BondCBIndexJSL](#bondcbindexjsl)
- [BondCBJSL](#bondcbjsl)
- [BondCBProfileSina](#bondcbprofilesina)
- [BondCBRedeemJSL](#bondcbredeemjsl)
- [BondCBSummarySina](#bondcbsummarysina)
- [BondCashSummarySSE](#bondcashsummarysse)
- [BondChinaCloseReturn](#bondchinaclosereturn)
- [BondChinaYield](#bondchinayield)
- [BondCompositeIndexCBond](#bondcompositeindexcbond)
- [BondCorporateIssueCninfo](#bondcorporateissuecninfo)
- [BondCovComparison](#bondcovcomparison)
- [BondCovIssueCninfo](#bondcovissuecninfo)
- [BondCovStockIssueCninfo](#bondcovstockissuecninfo)
- [BondDealSummarySSE](#bonddealsummarysse)
- [BondDebtNAFMII](#bonddebtnafmii)
- [BondInfoCm](#bondinfocm)
- [BondInfoCmQuery](#bondinfocmquery)
- [BondInfoDetailCm](#bondinfodetailcm)
- [BondLocalGovernmentIssueCninfo](#bondlocalgovernmentissuecninfo)
- [BondNewCompositeIndexCBond](#bondnewcompositeindexcbond)
- [BondSHBuyBackEM](#bondshbuybackem)
- [BondSZBuyBackEM](#bondszbuybackem)
- [BondSpotDeal](#bondspotdeal)
- [BondSpotQuote](#bondspotquote)
- [BondTreasureIssueCninfo](#bondtreasureissuecninfo)
- [BondZHCovInfoTHS](#bondzhcovinfoths)
- [BondZHHSDaily](#bondzhhsdaily)
- [BondZHHSSpot](#bondzhhsspot)
- [BondZHUSRate](#bondzhusrate)
- [BondZhCov](#bondzhcov)
- [BondZhCovInfo](#bondzhcovinfo)
- [BondZhCovInfoThs](#bondzhcovinfoths)
- [BondZhCovValueAnalysis](#bondzhcovvalueanalysis)
- [BondZhHsCovDaily](#bondzhhscovdaily)
- [BondZhHsCovMin](#bondzhhscovmin)
- [BondZhHsCovPreMin](#bondzhhscovpremin)
- [BondZhHsCovSpot](#bondzhhscovspot)
- [BondZhUsRate](#bondzhusrate)
- [MacroChinaBondPublic](#macrochinabondpublic)
- [MacroChinaSwapRate](#macrochinaswaprate)

---

## BondBuyBackHistEM

**描述**: BondBuyBackHistEM 质押式回购历史数据

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
	"github.com/BlakeLiAFK/akshare/bond"
)

func main() {
	data, err := bond.BondBuyBackHistEM("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## BondCBAdjLogsJSL

**描述**: BondCBAdjLogsJSL 集思录可转债转股价调整记录

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
	"github.com/BlakeLiAFK/akshare/bond"
)

func main() {
	data, err := bond.BondCBAdjLogsJSL("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## BondCBIndexJSL

**描述**: BondCBIndexJSL 集思录可转债等权指数

### 输入参数

无参数

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/bond"
)

func main() {
	data, err := bond.BondCBIndexJSL()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## BondCBJSL

**描述**: BondCBJSL 集思录可转债列表

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| cookie | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/bond"
)

func main() {
	data, err := bond.BondCBJSL("")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## BondCBProfileSina

**描述**: BondCBProfileSina 新浪财经可转债详情资料

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
	"github.com/BlakeLiAFK/akshare/bond"
)

func main() {
	data, err := bond.BondCBProfileSina("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## BondCBRedeemJSL

**描述**: BondCBRedeemJSL 集思录可转债强赎数据

### 输入参数

无参数

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/bond"
)

func main() {
	data, err := bond.BondCBRedeemJSL()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## BondCBSummarySina

**描述**: BondCBSummarySina 新浪财经可转债债券概况

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
	"github.com/BlakeLiAFK/akshare/bond"
)

func main() {
	data, err := bond.BondCBSummarySina("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## BondCashSummarySSE

**描述**: BondCashSummarySSE 上交所债券现券汇总

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
	"github.com/BlakeLiAFK/akshare/bond"
)

func main() {
	data, err := bond.BondCashSummarySSE("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## BondChinaCloseReturn

**描述**: BondChinaCloseReturn 收盘收益率曲线历史数据

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
	"github.com/BlakeLiAFK/akshare/bond"
)

func main() {
	data, err := bond.BondChinaCloseReturn("000001", "daily", "20230101", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## BondChinaYield

**描述**: BondChinaYield 中国债券信息网-国债及其他债券收益率曲线

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
	"github.com/BlakeLiAFK/akshare/bond"
)

func main() {
	data, err := bond.BondChinaYield("20230101", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## BondCompositeIndexCBond

**描述**: BondCompositeIndexCBond 中债综合指数

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
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
	"github.com/BlakeLiAFK/akshare/bond"
)

func main() {
	data, err := bond.BondCompositeIndexCBond("", "daily")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## BondCorporateIssueCninfo

**描述**: BondCorporateIssueCninfo 巨潮资讯-数据中心-专题统计-债券报表-债券发行-企业债发行

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
	"github.com/BlakeLiAFK/akshare/bond"
)

func main() {
	data, err := bond.BondCorporateIssueCninfo("20230101", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## BondCovComparison

**描述**: BondCovComparison 东方财富网-行情中心-债券市场-可转债比价表

**数据源**: https://quote.eastmoney.com/center/fullscreenlist.html#convertible_comparison

### 输入参数

无参数

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/bond"
)

func main() {
	data, err := bond.BondCovComparison()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## BondCovIssueCninfo

**描述**: BondCovIssueCninfo 巨潮资讯-数据中心-专题统计-债券报表-债券发行-可转债发行

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
	"github.com/BlakeLiAFK/akshare/bond"
)

func main() {
	data, err := bond.BondCovIssueCninfo("20230101", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## BondCovStockIssueCninfo

**描述**: BondCovStockIssueCninfo 巨潮资讯-数据中心-专题统计-债券报表-债券发行-可转债转股

### 输入参数

无参数

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/bond"
)

func main() {
	data, err := bond.BondCovStockIssueCninfo()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## BondDealSummarySSE

**描述**: BondDealSummarySSE 上交所债券成交汇总

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
	"github.com/BlakeLiAFK/akshare/bond"
)

func main() {
	data, err := bond.BondDealSummarySSE("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## BondDebtNAFMII

**描述**: BondDebtNAFMII 银行间市场债务融资工具

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| page | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/bond"
)

func main() {
	data, err := bond.BondDebtNAFMII("")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## BondInfoCm

**描述**: BondInfoCm 中国外汇交易中心暨全国银行间同业拆借中心-数据-债券信息-信息查询

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| bondName | string | - |
| bondCode | string | 股票/基金代码 |
| bondIssue | string | - |
| bondType | string | - |
| couponType | string | - |
| issueYear | string | - |
| underwriter | string | - |
| grade | string | - |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/bond"
)

func main() {
	data, err := bond.BondInfoCm("", "000001", "", "", "", "", "", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## BondInfoCmQuery

**描述**: BondInfoCmQuery 中国外汇交易中心暨全国银行间同业拆借中心-查询相关指标的参数

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
	"github.com/BlakeLiAFK/akshare/bond"
)

func main() {
	data, err := bond.BondInfoCmQuery("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## BondInfoDetailCm

**描述**: BondInfoDetailCm 中国外汇交易中心暨全国银行间同业拆借中心-数据-债券信息-信息查询-债券详情

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
	"github.com/BlakeLiAFK/akshare/bond"
)

func main() {
	data, err := bond.BondInfoDetailCm("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## BondLocalGovernmentIssueCninfo

**描述**: BondLocalGovernmentIssueCninfo 巨潮资讯-数据中心-专题统计-债券报表-债券发行-地方债发行

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
	"github.com/BlakeLiAFK/akshare/bond"
)

func main() {
	data, err := bond.BondLocalGovernmentIssueCninfo("20230101", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## BondNewCompositeIndexCBond

**描述**: BondNewCompositeIndexCBond 中债新综合指数

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
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
	"github.com/BlakeLiAFK/akshare/bond"
)

func main() {
	data, err := bond.BondNewCompositeIndexCBond("", "daily")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## BondSHBuyBackEM

**描述**: BondSHBuyBackEM 上证质押式回购

### 输入参数

无参数

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/bond"
)

func main() {
	data, err := bond.BondSHBuyBackEM()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## BondSZBuyBackEM

**描述**: BondSZBuyBackEM 深证质押式回购

### 输入参数

无参数

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/bond"
)

func main() {
	data, err := bond.BondSZBuyBackEM()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## BondSpotDeal

**描述**: BondSpotDeal 中国外汇交易中心-现券市场成交行情

### 输入参数

无参数

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/bond"
)

func main() {
	data, err := bond.BondSpotDeal()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## BondSpotQuote

**描述**: BondSpotQuote 中国外汇交易中心-现券市场做市报价

### 输入参数

无参数

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/bond"
)

func main() {
	data, err := bond.BondSpotQuote()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## BondTreasureIssueCninfo

**描述**: BondTreasureIssueCninfo 巨潮资讯-数据中心-专题统计-债券报表-债券发行-国债发行

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
	"github.com/BlakeLiAFK/akshare/bond"
)

func main() {
	data, err := bond.BondTreasureIssueCninfo("20230101", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## BondZHCovInfoTHS

**描述**: BondZHCovInfoTHS 同花顺可转债数据

### 输入参数

无参数

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/bond"
)

func main() {
	data, err := bond.BondZHCovInfoTHS()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## BondZHHSDaily

**描述**: BondZHHSDaily 新浪财经-债券-沪深债券-历史行情数据

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
	"github.com/BlakeLiAFK/akshare/bond"
)

func main() {
	data, err := bond.BondZHHSDaily("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## BondZHHSSpot

**描述**: BondZHHSSpot 新浪财经-债券-沪深债券-实时行情数据

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
	"github.com/BlakeLiAFK/akshare/bond"
)

func main() {
	data, err := bond.BondZHHSSpot("", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## BondZHUSRate

**描述**: BondZHUSRate 中美国债收益率

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| startDate | string | 开始日期，格式：YYYYMMDD |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/bond"
)

func main() {
	data, err := bond.BondZHUSRate("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## BondZhCov

**描述**: BondZhCov 东方财富网-数据中心-新股数据-可转债数据

**数据源**: https://data.eastmoney.com/kzz/default.html

### 输入参数

无参数

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/bond"
)

func main() {
	data, err := bond.BondZhCov()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## BondZhCovInfo

**描述**: BondZhCovInfo 东方财富网-数据中心-新股数据-可转债详情

**数据源**: https://data.eastmoney.com/kzz/detail/123121.html

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
	"github.com/BlakeLiAFK/akshare/bond"
)

func main() {
	data, err := bond.BondZhCovInfo("000001", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## BondZhCovInfoThs

**描述**: BondZhCovInfoThs 同花顺-数据中心-可转债

**数据源**: https://data.10jqka.com.cn/ipo/bond/

### 输入参数

无参数

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/bond"
)

func main() {
	data, err := bond.BondZhCovInfoThs()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## BondZhCovValueAnalysis

**描述**: BondZhCovValueAnalysis 东方财富网-数据中心-新股数据-可转债数据-价值分析-溢价率分析

**数据源**: https://data.eastmoney.com/kzz/detail/113527.html

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
	"github.com/BlakeLiAFK/akshare/bond"
)

func main() {
	data, err := bond.BondZhCovValueAnalysis("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## BondZhHsCovDaily

**描述**: BondZhHsCovDaily 新浪财经-债券-沪深可转债的历史行情数据

**数据源**: https://vip.stock.finance.sina.com.cn/mkt/#hskzz_z

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
	"github.com/BlakeLiAFK/akshare/bond"
)

func main() {
	data, err := bond.BondZhHsCovDaily("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## BondZhHsCovMin

**描述**: BondZhHsCovMin 东方财富网-可转债-分时行情

**数据源**: https://quote.eastmoney.com/concept/sz128039.html

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
	"github.com/BlakeLiAFK/akshare/bond"
)

func main() {
	data, err := bond.BondZhHsCovMin("000001", "daily", "qfq", "20230101", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## BondZhHsCovPreMin

**描述**: BondZhHsCovPreMin 东方财富网-可转债-分时行情-盘前

**数据源**: https://quote.eastmoney.com/concept/sz128039.html

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
	"github.com/BlakeLiAFK/akshare/bond"
)

func main() {
	data, err := bond.BondZhHsCovPreMin("000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## BondZhHsCovSpot

**描述**: BondZhHsCovSpot 新浪财经-债券-沪深可转债的实时行情数据

**数据源**: https://vip.stock.finance.sina.com.cn/mkt/#hskzz_z

### 输入参数

无参数

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/bond"
)

func main() {
	data, err := bond.BondZhHsCovSpot()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## BondZhUsRate

**描述**: BondZhUsRate 东方财富网-数据中心-经济数据-中美国债收益率

### 输入参数

| 参数名 | 类型 | 说明 |
|------|------|------|
| startDate | string | 开始日期，格式：YYYYMMDD |

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/bond"
)

func main() {
	data, err := bond.BondZhUsRate("20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaBondPublic

**描述**: MacroChinaBondPublic 中国-债券信息披露-债券发行

### 输入参数

无参数

### 输出参数

返回 `[]map[string]interface{}` 类型的数据，每个map包含以下字段：

具体字段请参考示例输出。

### 代码示例

```go
package main

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/bond"
)

func main() {
	data, err := bond.MacroChinaBondPublic()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

## MacroChinaSwapRate

**描述**: MacroChinaSwapRate FR007利率互换曲线历史数据

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
	"github.com/BlakeLiAFK/akshare/bond"
)

func main() {
	data, err := bond.MacroChinaSwapRate("20230101", "20230101")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
```


---

