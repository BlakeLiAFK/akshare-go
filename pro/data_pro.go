package pro

import (
	"fmt"
	"os"
	"strings"
)

// ProApi 初始化 pro API
//
// 第一次可以通过设置环境变量 AKSHARE_TOKEN 来记录自己的token凭证
// 临时token可以通过本参数传入
//
// 参数:
//   - token: API token，为空则从环境变量或文件读取
//
// 返回:
//   - *DataApi: 数据API客户端
//   - error: 错误信息
//
// 示例:
//
//	pro, err := pro.ProApi("your_token")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	data, err := pro.VarietyAll()
func ProApi(token string) (*DataApi, error) {
	if token == "" {
		token = getToken()
	}

	if token == "" {
		return nil, fmt.Errorf(TokenErrMsg)
	}

	return NewDataApi(token, 10), nil
}

// getToken 从环境变量或文件获取token
func getToken() string {
	// 优先从环境变量获取
	if token := os.Getenv("AKSHARE_TOKEN"); token != "" {
		return token
	}

	// 尝试从文件读取
	if data, err := os.ReadFile(TokenFilePath); err == nil {
		token := strings.TrimSpace(string(data))
		if token != "" {
			return token
		}
	}

	return ""
}

// SetToken 设置token到环境变量
func SetToken(token string) error {
	return os.Setenv("AKSHARE_TOKEN", token)
}
