package cal

import (
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

// RVFromStockZhAHistMinEM 从东方财富网获取股票的分钟级历史行情数据并格式化
//
// 参数:
//
//	symbol: 股票代码,如"000001"
//	startDate: 开始日期时间,格式"2021-10-20 09:30:00"
//	endDate: 结束日期时间,格式"2024-11-01 15:00:00"
//	period: 时间周期,可选{"1","5","15","30","60"}分钟
//	adjust: 复权方式,可选{""(不复权),"qfq"(前复权),"hfq"(后复权)}
//
// 返回:
//   - dataframe.DataFrame: 包含Date(索引),Open,High,Low,Close列的数据
//   - error: 错误信息
//
// 数据源: https://quote.eastmoney.com
//
// 依赖说明:
//
//	此函数依赖 stock_feature 模块的 StockZhAHistMinEM 函数
//	该模块将在后续按字母顺序实现时补充
func RVFromStockZhAHistMinEM(symbol, startDate, endDate, period, adjust string) (dataframe.DataFrame, error) {
	return dataframe.DataFrame{}, errors.New("此函数依赖 stock_feature 模块，该模块尚未实现，将在后续补充")
}

// RVFromFuturesZhMinuteSina 从新浪财经获取期货的分钟级历史行情数据并格式化
//
// 参数:
//
//	symbol: 期货合约代码,如"IF2008"代表沪深300期货2020年8月合约
//	period: 时间周期,可选{"1","5","15","30","60"}分钟
//
// 返回:
//   - dataframe.DataFrame: 包含Date(索引),Open,High,Low,Close列的数据
//   - error: 错误信息
//
// 数据源: https://vip.stock.finance.sina.com.cn
//
// 依赖说明:
//
//	此函数依赖 futures 模块的 FuturesZhMinuteSina 函数
//	该模块将在后续按字母顺序实现时补充
func RVFromFuturesZhMinuteSina(symbol, period string) (dataframe.DataFrame, error) {
	return dataframe.DataFrame{}, errors.New("此函数依赖 futures 模块，该模块尚未实现，将在后续补充")
}

// VolatilityYZRV 计算 Yang-Zhang 已实现波动率
//
// 参数:
//
//	data: 包含 Open, High, Low, Close 列的 DataFrame，且 Date 列为索引
//
// 返回:
//   - dataframe.DataFrame: 包含 date, rv 列的数据
//   - error: 错误信息
//
// 理论基础:
//
//	论文地址: https://www.jstor.org/stable/10.1086/209650
//	参考实现: https://github.com/hugogobato/Yang-Zhang-s-Realized-Volatility-Automated-Estimation-in-Python
//
// 公式:
//
//	RV^2 = Vo + k*Vc + (1-k)*Vrs
//	其中:
//	  - Vo: 隔夜波动率, Vo = 1/(n-1)*sum(Oi-Obar)^2
//	  - Vc: 收盘波动率, Vc = 1/(n-1)*sum(ci-Cbar)^2
//	  - k: 权重系数, k = 0.34/(1.34+(n+1)/(n-1))
//	  - Vrs: Rogers-Satchell波动率, Vrs = ui(ui-ci)+di(di-ci)
//	    其中 ui = ln(Hi/Oi), ci = ln(Ci/Oi), di = ln(Li/Oi), oi = ln(Oi/Ci-1)
func VolatilityYZRV(data dataframe.DataFrame) (dataframe.DataFrame, error) {
	// 检查必需的列
	requiredCols := []string{"Open", "High", "Low", "Close"}
	for _, col := range requiredCols {
		if data.Col(col).Err != nil {
			return dataframe.DataFrame{}, fmt.Errorf("缺少必需列: %s", col)
		}
	}

	nRows, _ := data.Dims()
	if nRows < 2 {
		return dataframe.DataFrame{}, errors.New("数据行数不足，至少需要2行")
	}

	// 提取列数据
	opens := data.Col("Open").Float()
	highs := data.Col("High").Float()
	lows := data.Col("Low").Float()
	closes := data.Col("Close").Float()

	// 计算中间变量 (从第二行开始)
	ui := make([]float64, nRows-1)
	ci := make([]float64, nRows-1)
	di := make([]float64, nRows-1)
	oi := make([]float64, nRows-1)
	rs := make([]float64, nRows-1)
	dates := make([]time.Time, nRows-1)

	// 获取日期列（假设有Date列或索引）
	dateCol := data.Col("Date")
	if dateCol.Err != nil {
		// 如果没有Date列，创建虚拟日期
		baseDate := time.Now().AddDate(0, 0, -(nRows - 1))
		for i := 0; i < nRows-1; i++ {
			dates[i] = baseDate.AddDate(0, 0, i+1)
		}
	} else {
		// 解析日期
		for i := 1; i < nRows; i++ {
			dateStr := dateCol.Elem(i).String()
			parsedDate, err := time.Parse("2006-01-02 15:04:05", dateStr)
			if err != nil {
				// 尝试其他格式
				parsedDate, err = time.Parse("2006-01-02", dateStr)
				if err != nil {
					return dataframe.DataFrame{}, fmt.Errorf("解析日期失败: %v", err)
				}
			}
			dates[i-1] = parsedDate
		}
	}

	// 计算各项指标
	for i := 1; i < nRows; i++ {
		idx := i - 1
		// ui = ln(Hi/Oi)
		ui[idx] = math.Log(highs[i] / opens[i])
		// ci = ln(Ci/Oi)
		ci[idx] = math.Log(closes[i] / opens[i])
		// di = ln(Li/Oi)
		di[idx] = math.Log(lows[i] / opens[i])
		// oi = ln(Oi/Ci-1)
		oi[idx] = math.Log(opens[i] / closes[i-1])
		// RS = ui(ui-ci) + di(di-ci)
		rs[idx] = ui[idx]*(ui[idx]-ci[idx]) + di[idx]*(di[idx]-ci[idx])
	}

	// 按日分组计算
	// 先将数据按日期分组
	type DayData struct {
		Date time.Time
		RS   []float64
		OI   []float64
		CI   []float64
	}

	dayMap := make(map[string]*DayData)
	for i := 0; i < len(dates); i++ {
		dateKey := dates[i].Format("2006-01-02")
		if _, exists := dayMap[dateKey]; !exists {
			dayMap[dateKey] = &DayData{
				Date: dates[i],
				RS:   []float64{},
				OI:   []float64{},
				CI:   []float64{},
			}
		}
		dayMap[dateKey].RS = append(dayMap[dateKey].RS, rs[i])
		dayMap[dateKey].OI = append(dayMap[dateKey].OI, oi[i])
		dayMap[dateKey].CI = append(dayMap[dateKey].CI, ci[i])
	}

	// 计算每日的统计量
	type DayResult struct {
		Date   string
		RsMean float64
		OiVar  float64
		CiVar  float64
	}

	var dayResults []DayResult
	for dateKey, dayData := range dayMap {
		// 计算 RS 均值
		rsMean := mean(dayData.RS)
		// 计算 OI 方差
		oiVar := variance(dayData.OI)
		// 计算 CI 方差
		ciVar := variance(dayData.CI)

		dayResults = append(dayResults, DayResult{
			Date:   dateKey,
			RsMean: rsMean,
			OiVar:  oiVar,
			CiVar:  ciVar,
		})
	}

	if len(dayResults) == 0 {
		return dataframe.DataFrame{}, errors.New("没有有效的日数据")
	}

	// 计算参数
	n := float64(len(dates)) / float64(len(dayResults))
	k := 0.34 / (1.34 + (n+1)/(n-1))

	// 计算 Yang-Zhang RV
	var resultDates []string
	var resultRVs []float64

	for _, dayResult := range dayResults {
		// RV^2 = (1-k)*Vrs + Vo + k*Vc
		// 其中 Vrs = rs_mean, Vo = oi_var, Vc = ci_var
		rvSquared := (1-k)*dayResult.RsMean + dayResult.OiVar + k*dayResult.CiVar
		rv := math.Sqrt(math.Abs(rvSquared)) // 使用绝对值避免负数

		resultDates = append(resultDates, dayResult.Date)
		resultRVs = append(resultRVs, rv)
	}

	// 构建返回的 DataFrame
	df := dataframe.New(
		series.New(resultDates, series.String, "date"),
		series.New(resultRVs, series.Float, "rv"),
	)

	return df, nil
}

// mean 计算切片的均值
func mean(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	sum := 0.0
	for _, v := range data {
		sum += v
	}
	return sum / float64(len(data))
}

// variance 计算切片的方差
func variance(data []float64) float64 {
	if len(data) <= 1 {
		return 0
	}
	m := mean(data)
	sumSquaredDiff := 0.0
	for _, v := range data {
		diff := v - m
		sumSquaredDiff += diff * diff
	}
	return sumSquaredDiff / float64(len(data)-1)
}
