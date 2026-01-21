package qhkc

import (
	"github.com/BlakeLiAFK/akshare/pro"
)

// 奇货可查指数数据
// 注意：此模块需要token才能访问，请通过 pro.ProApi() 获取客户端

// GetProClient 获取奇货可查API客户端
//
// 参数:
//   - token: API token
//
// 返回:
//   - *pro.DataApi: 数据API客户端
//   - error: 错误信息
//
// 示例:
//
//	client, err := qhkc.GetProClient("your_token")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	data, err := client.IndexInfo("index0070c0eb-93ba-2da9-6633-fa70cb90e959")
func GetProClient(token string) (*pro.DataApi, error) {
	return pro.ProApi(token)
}

// IndexInfo 获取指数信息
//
// 通过奇货可查API获取指数信息
//
// 参数:
//   - client: 奇货可查API客户端
//   - indexID: 指数ID
//
// 返回:
//   - []map[string]interface{}: 指数信息数据
//   - error: 错误信息
func IndexInfo(client *pro.DataApi, indexID string) ([]map[string]interface{}, error) {
	return client.IndexInfo(indexID)
}

// IndexWeights 获取指数权重数据
//
// 通过奇货可查API获取指数权重数据
//
// 参数:
//   - client: 奇货可查API客户端
//   - indexID: 指数ID
//   - date: 日期，格式 "2018-08-08"
//
// 返回:
//   - []map[string]interface{}: 指数权重数据
//   - error: 错误信息
func IndexWeights(client *pro.DataApi, indexID, date string) ([]map[string]interface{}, error) {
	return client.IndexWeights(indexID, date)
}

// IndexKline 获取指数行情数据
//
// 通过奇货可查API获取指数行情数据
//
// 参数:
//   - client: 奇货可查API客户端
//   - indexID: 指数ID
//
// 返回:
//   - []map[string]interface{}: 指数行情数据
//   - error: 错误信息
func IndexKline(client *pro.DataApi, indexID string) ([]map[string]interface{}, error) {
	return client.IndexKline(indexID)
}

// IndexMember 获取指数沉淀资金数据
//
// 通过奇货可查API获取指数沉淀资金数据
//
// 参数:
//   - client: 奇货可查API客户端
//   - indexID: 指数ID
//
// 返回:
//   - []map[string]interface{}: 指数沉淀资金数据
//   - error: 错误信息
func IndexMember(client *pro.DataApi, indexID string) ([]map[string]interface{}, error) {
	return client.IndexMember(indexID)
}
