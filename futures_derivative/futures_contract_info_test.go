package futures_derivative

import (
	"testing"
	"time"
)

func TestFuturesContractInfoCffex(t *testing.T) {
	// 使用最近的交易日
	date := time.Now().AddDate(0, 0, -3).Format("20060102")

	records, err := FuturesContractInfoCffex(date)
	if err != nil {
		t.Logf("FuturesContractInfoCffex warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("FuturesContractInfoCffex returned %d records", len(records))

	// 如果有数据，验证结构
	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"合约代码", "合约月份", "挂盘基准价", "上市日", "最后交易日"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}

func TestFuturesContractInfoCzce(t *testing.T) {
	// 使用最近的交易日
	date := time.Now().AddDate(0, 0, -3).Format("20060102")

	records, err := FuturesContractInfoCzce(date)
	if err != nil {
		t.Logf("FuturesContractInfoCzce warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("FuturesContractInfoCzce returned %d records", len(records))

	// 如果有数据，验证结构
	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"产品名称", "合约代码", "产品代码", "交易单位", "第一交易日"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}

func TestFuturesContractInfoDce(t *testing.T) {
	records, err := FuturesContractInfoDce()
	if err != nil {
		t.Logf("FuturesContractInfoDce warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("FuturesContractInfoDce returned %d records", len(records))

	// 如果有数据，验证结构
	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"品种名称", "合约", "交易单位", "最小变动价位"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}

func TestFuturesContractInfoGfex(t *testing.T) {
	records, err := FuturesContractInfoGfex()
	if err != nil {
		t.Logf("FuturesContractInfoGfex warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("FuturesContractInfoGfex returned %d records", len(records))

	// 如果有数据，验证结构
	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"品种", "合约代码", "交易单位", "最小变动单位"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}

func TestFuturesContractInfoIne(t *testing.T) {
	// 使用最近的交易日
	date := time.Now().AddDate(0, 0, -3).Format("20060102")

	records, err := FuturesContractInfoIne(date)
	if err != nil {
		t.Logf("FuturesContractInfoIne warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("FuturesContractInfoIne returned %d records", len(records))

	// 如果有数据，验证结构
	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"合约代码", "上市日", "到期日", "开始交割日", "最后交割日"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}

func TestFuturesContractInfoShfe(t *testing.T) {
	// 使用最近的交易日
	date := time.Now().AddDate(0, 0, -3).Format("20060102")

	records, err := FuturesContractInfoShfe(date)
	if err != nil {
		t.Logf("FuturesContractInfoShfe warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("FuturesContractInfoShfe returned %d records", len(records))

	// 如果有数据，验证结构
	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"合约代码", "上市日", "到期日", "开始交割日", "最后交割日"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}
