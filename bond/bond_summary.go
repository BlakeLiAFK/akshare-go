package bond

import (
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

// BondCashSummarySSE 上交所债券现券汇总
//
// 获取上登债券信息网的债券现券市场概览数据
//
// 参数:
//   - date: 指定日期，格式: "20210111"
//
// 返回:
//   - dataframe.DataFrame: 包含债券现券汇总数据
//   - error: 错误信息
//
// 数据源: http://bond.sse.com.cn/data/statistics/overview/bondow/
func BondCashSummarySSE(date string) (dataframe.DataFrame, error) {
	if len(date) != 8 {
		return dataframe.DataFrame{}, fmt.Errorf("日期格式错误，应为YYYYMMDD")
	}

	url := "http://query.sse.com.cn/commonExcelDd.do"
	headers := map[string]string{
		"Referer":    "http://bond.sse.com.cn/",
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
	}

	params := map[string]string{
		"sqlId":      "COMMON_SSEBOND_SCSJ_SCTJ_SCGL_ZQXQSCGL_CX_L",
		"TRADE_DATE": fmt.Sprintf("%s-%s-%s", date[:4], date[4:6], date[6:]),
	}

	resp, err := utils.GetWithHeaders(url, params, headers)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	// 解析Excel数据
	df, err := utils.ReadExcelFromBytes(resp.Body())
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("解析Excel失败: %w", err)
	}

	// 重命名列
	df = df.Rename("债券现货", "债券现货")
	df = df.Rename("托管只数", "托管只数")
	df = df.Rename("托管市值", "托管市值")
	df = df.Rename("托管面值", "托管面值")
	df = df.Rename("数据日期", "数据日期")

	// 转换数据类型
	df = df.Mutate(series.New(df.Col("托管只数").Float(), series.Float, "托管只数"))
	df = df.Mutate(series.New(df.Col("托管市值").Float(), series.Float, "托管市值"))
	df = df.Mutate(series.New(df.Col("托管面值").Float(), series.Float, "托管面值"))

	return df, nil
}

// BondDealSummarySSE 上交所债券成交汇总
//
// 获取上登债券信息网的债券成交概览数据
//
// 参数:
//   - date: 指定日期，格式: "20210104"
//
// 返回:
//   - dataframe.DataFrame: 包含债券成交汇总数据
//   - error: 错误信息
//
// 数据源: http://bond.sse.com.cn/data/statistics/overview/turnover/
func BondDealSummarySSE(date string) (dataframe.DataFrame, error) {
	if len(date) != 8 {
		return dataframe.DataFrame{}, fmt.Errorf("日期格式错误，应为YYYYMMDD")
	}

	url := "http://query.sse.com.cn/commonExcelDd.do"
	headers := map[string]string{
		"Referer":    "http://bond.sse.com.cn/",
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
	}

	params := map[string]string{
		"sqlId":      "COMMON_SSEBOND_SCSJ_SCTJ_SCGL_ZQCJGL_CX_L",
		"TRADE_DATE": fmt.Sprintf("%s-%s-%s", date[:4], date[4:6], date[6:]),
	}

	resp, err := utils.GetWithHeaders(url, params, headers)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	// 解析Excel数据
	df, err := utils.ReadExcelFromBytes(resp.Body())
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("解析Excel失败: %w", err)
	}

	// 重命名列
	df = df.Rename("债券类型", "债券类型")
	df = df.Rename("当日成交笔数", "当日成交笔数")
	df = df.Rename("当日成交金额", "当日成交金额")
	df = df.Rename("当年成交笔数", "当年成交笔数")
	df = df.Rename("当年成交金额", "当年成交金额")
	df = df.Rename("数据日期", "数据日期")

	// 转换数据类型
	df = df.Mutate(series.New(df.Col("当日成交笔数").Float(), series.Float, "当日成交笔数"))
	df = df.Mutate(series.New(df.Col("当日成交金额").Float(), series.Float, "当日成交金额"))
	df = df.Mutate(series.New(df.Col("当年成交笔数").Float(), series.Float, "当年成交笔数"))
	df = df.Mutate(series.New(df.Col("当年成交金额").Float(), series.Float, "当年成交金额"))

	return df, nil
}
