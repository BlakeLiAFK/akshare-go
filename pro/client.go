package pro

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// DataApi 奇货可查数据接口客户端
type DataApi struct {
	token   string
	httpURL string
	timeout time.Duration
	client  *http.Client
}

// NewDataApi 创建新的数据API客户端
//
// 参数:
//   - token: API接口TOKEN，用于用户认证
//   - timeout: 超时设置（秒），默认10秒
//
// 返回:
//   - *DataApi: 数据API客户端
func NewDataApi(token string, timeout int) *DataApi {
	if timeout <= 0 {
		timeout = 10
	}
	return &DataApi{
		token:   token,
		httpURL: QhkcAPIURL,
		timeout: time.Duration(timeout) * time.Second,
		client: &http.Client{
			Timeout: time.Duration(timeout) * time.Second,
		},
	}
}

// Query 通用查询接口
//
// 参数:
//   - apiName: 需要调取的接口名称
//   - kwargs: 请求参数
//
// 返回:
//   - map[string]interface{}: 返回的JSON数据
//   - error: 错误信息
func (d *DataApi) Query(apiName string, kwargs ...string) (map[string]interface{}, error) {
	// 构建URL
	pathParts := []string{apiName}
	pathParts = append(pathParts, kwargs...)
	url := d.httpURL + "/" + strings.Join(pathParts, "/")

	// 创建请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	// 设置请求头
	req.Header.Set("X-Token", d.token)

	// 发送请求
	resp, err := d.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}
	defer resp.Body.Close()

	// 检查状态码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("连接异常, 请检查您的Token是否过期和输入的参数是否正确, 状态码: %d", resp.StatusCode)
	}

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %w", err)
	}

	// 解析JSON
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("解析JSON失败: %w", err)
	}

	return result, nil
}

// QueryList 查询返回列表数据
//
// 参数:
//   - apiName: 需要调取的接口名称
//   - kwargs: 请求参数
//
// 返回:
//   - []map[string]interface{}: 返回的列表数据
//   - error: 错误信息
func (d *DataApi) QueryList(apiName string, kwargs ...string) ([]map[string]interface{}, error) {
	// 构建URL
	pathParts := []string{apiName}
	pathParts = append(pathParts, kwargs...)
	url := d.httpURL + "/" + strings.Join(pathParts, "/")

	// 创建请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	// 设置请求头
	req.Header.Set("X-Token", d.token)

	// 发送请求
	resp, err := d.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}
	defer resp.Body.Close()

	// 检查状态码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("连接异常, 请检查您的Token是否过期和输入的参数是否正确, 状态码: %d", resp.StatusCode)
	}

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %w", err)
	}

	// 解析JSON
	var result []map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("解析JSON失败: %w", err)
	}

	return result, nil
}

// IndexInfo 获取指数信息
func (d *DataApi) IndexInfo(indexID string) ([]map[string]interface{}, error) {
	return d.QueryList("index_info", indexID)
}

// IndexWeights 获取指数权重数据
func (d *DataApi) IndexWeights(indexID, date string) ([]map[string]interface{}, error) {
	return d.QueryList("index_weights", indexID, date)
}

// IndexKline 获取指数行情数据
func (d *DataApi) IndexKline(indexID string) ([]map[string]interface{}, error) {
	return d.QueryList("index_kline", indexID)
}

// IndexMember 获取指数沉淀资金数据
func (d *DataApi) IndexMember(indexID string) ([]map[string]interface{}, error) {
	return d.QueryList("index_member", indexID)
}

// VarietyAll 获取商品列表数据
func (d *DataApi) VarietyAll() ([]map[string]interface{}, error) {
	return d.QueryList("variety_all")
}

// VarietyPositions 获取合约持仓数据
func (d *DataApi) VarietyPositions(symbol, date string) ([]map[string]interface{}, error) {
	return d.QueryList("variety_positions", symbol, date)
}

// BrokerPositions 获取席位持仓数据
func (d *DataApi) BrokerPositions(broker, date string) ([]map[string]interface{}, error) {
	return d.QueryList("broker_positions", broker, date)
}

// LongPool 获取龙虎牛熊多头合约池
func (d *DataApi) LongPool(date string) ([]map[string]interface{}, error) {
	return d.QueryList("long_pool", date)
}

// ShortPool 获取龙虎牛熊空头合约池
func (d *DataApi) ShortPool(date string) ([]map[string]interface{}, error) {
	return d.QueryList("short_pool", date)
}
