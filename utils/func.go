package utils

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/tidwall/gjson"
)

// FetchPaginatedData 分页获取数据并合并结果
// 适用于东方财富等分页接口
// url: 请求地址
// baseParams: 基础请求参数
// timeout: 超时时间（秒）
// 返回合并后的 JSON 数组
func FetchPaginatedData(url string, baseParams map[string]string, timeout int) ([]gjson.Result, error) {
	if timeout <= 0 {
		timeout = 15
	}

	// 复制参数避免修改原始参数
	params := make(map[string]string)
	for k, v := range baseParams {
		params[k] = v
	}

	// 设置临时超时
	oldTimeout := HTTPClient.GetClient().Timeout
	HTTPClient.SetTimeout(time.Duration(timeout) * time.Second)
	defer HTTPClient.SetTimeout(oldTimeout)

	// 获取第一页数据
	resp, err := Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求第一页失败: %w", err)
	}

	text := resp.String()
	diffArr := gjson.Get(text, "data.diff").Array()
	if len(diffArr) == 0 {
		return nil, nil
	}

	// 计算分页信息
	perPageNum := len(diffArr)
	total := gjson.Get(text, "data.total").Int()
	totalPage := int(math.Ceil(float64(total) / float64(perPageNum)))

	// 存储所有数据
	allData := make([]gjson.Result, 0, total)
	allData = append(allData, diffArr...)

	// 获取剩余页面
	for page := 2; page <= totalPage; page++ {
		params["pn"] = fmt.Sprintf("%d", page)

		// 添加随机延迟避免请求过于频繁
		time.Sleep(time.Duration(500+rand.Intn(1000)) * time.Millisecond)

		resp, err := Get(url, params)
		if err != nil {
			continue // 跳过失败的页面
		}

		pageData := gjson.Get(resp.String(), "data.diff").Array()
		allData = append(allData, pageData...)
	}

	return allData, nil
}

// FetchPaginatedDataWithCallback 分页获取数据，每页数据通过回调处理
// 适用于需要逐页处理的场景
func FetchPaginatedDataWithCallback(
	url string,
	baseParams map[string]string,
	timeout int,
	callback func(pageData []gjson.Result, page int) error,
) error {
	if timeout <= 0 {
		timeout = 15
	}

	params := make(map[string]string)
	for k, v := range baseParams {
		params[k] = v
	}

	oldTimeout := HTTPClient.GetClient().Timeout
	HTTPClient.SetTimeout(time.Duration(timeout) * time.Second)
	defer HTTPClient.SetTimeout(oldTimeout)

	// 获取第一页
	resp, err := Get(url, params)
	if err != nil {
		return fmt.Errorf("请求第一页失败: %w", err)
	}

	text := resp.String()
	diffArr := gjson.Get(text, "data.diff").Array()
	if len(diffArr) == 0 {
		return nil
	}

	// 处理第一页
	if err := callback(diffArr, 1); err != nil {
		return err
	}

	// 计算分页
	perPageNum := len(diffArr)
	total := gjson.Get(text, "data.total").Int()
	totalPage := int(math.Ceil(float64(total) / float64(perPageNum)))

	// 获取剩余页面
	for page := 2; page <= totalPage; page++ {
		params["pn"] = fmt.Sprintf("%d", page)
		time.Sleep(time.Duration(500+rand.Intn(1000)) * time.Millisecond)

		resp, err := Get(url, params)
		if err != nil {
			continue
		}

		pageData := gjson.Get(resp.String(), "data.diff").Array()
		if err := callback(pageData, page); err != nil {
			return err
		}
	}

	return nil
}

// RetryFunc 重试执行函数
// fn: 需要重试的函数
// maxRetries: 最大重试次数
// baseDelay: 基础延迟（毫秒）
func RetryFunc(fn func() error, maxRetries int, baseDelay int) error {
	var lastErr error

	for attempt := 0; attempt < maxRetries; attempt++ {
		if err := fn(); err != nil {
			lastErr = err
			if attempt < maxRetries-1 {
				// 指数退避 + 随机抖动
				delay := baseDelay * (1 << attempt)
				jitter := rand.Intn(500)
				time.Sleep(time.Duration(delay+jitter) * time.Millisecond)
			}
		} else {
			return nil
		}
	}

	return lastErr
}

// Chunk 将切片分块
func Chunk[T any](slice []T, size int) [][]T {
	if size <= 0 {
		return nil
	}

	var chunks [][]T
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

// Contains 检查切片是否包含元素
func Contains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// Unique 去重切片
func Unique[T comparable](slice []T) []T {
	seen := make(map[T]bool)
	result := make([]T, 0, len(slice))
	for _, v := range slice {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

// Map 对切片应用函数
func Map[T any, R any](slice []T, fn func(T) R) []R {
	result := make([]R, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// Filter 过滤切片
func Filter[T any](slice []T, fn func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range slice {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}
