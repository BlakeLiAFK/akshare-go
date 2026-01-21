package utils

import (
	"context"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
)

// 默认User-Agent
const DefaultUserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"

// HTTPClient 全局 HTTP 客户端
var HTTPClient *resty.Client

func init() {
	HTTPClient = resty.New().
		SetTimeout(30*time.Second).
		SetRetryCount(3).
		SetRetryWaitTime(1*time.Second).
		SetRetryMaxWaitTime(10*time.Second).
		SetHeader("User-Agent", DefaultUserAgent).
		SetHeader("Accept", "application/json, text/plain, */*").
		SetHeader("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")

	// 设置重试条件
	HTTPClient.AddRetryCondition(func(r *resty.Response, err error) bool {
		if err != nil {
			return true
		}
		// 429 Too Many Requests 需要重试
		if r.StatusCode() == http.StatusTooManyRequests {
			return true
		}
		// 5xx 服务器错误需要重试
		if r.StatusCode() >= 500 {
			return true
		}
		return false
	})
}

// Get 发起 GET 请求
func Get(url string, params map[string]string) (*resty.Response, error) {
	req := HTTPClient.R()
	if params != nil {
		req.SetQueryParams(params)
	}
	return req.Get(url)
}

// GetWithHeaders 发起带自定义请求头的 GET 请求
func GetWithHeaders(url string, params, headers map[string]string) (*resty.Response, error) {
	req := HTTPClient.R()
	if params != nil {
		req.SetQueryParams(params)
	}
	if headers != nil {
		req.SetHeaders(headers)
	}
	return req.Get(url)
}

// Post 发起 POST 请求
func Post(url string, data interface{}) (*resty.Response, error) {
	return HTTPClient.R().SetBody(data).Post(url)
}

// PostForm 发起 POST 表单请求
func PostForm(url string, data map[string]string) (*resty.Response, error) {
	return HTTPClient.R().SetFormData(data).Post(url)
}

// PostWithHeaders 发起带自定义请求头的 POST 请求
func PostWithHeaders(url string, data interface{}, headers map[string]string) (*resty.Response, error) {
	req := HTTPClient.R().SetBody(data)
	if headers != nil {
		req.SetHeaders(headers)
	}
	return req.Post(url)
}

// GetJSON 发起 GET 请求并解析 JSON
func GetJSON(url string, params map[string]string) (gjson.Result, error) {
	resp, err := Get(url, params)
	if err != nil {
		return gjson.Result{}, err
	}
	return gjson.ParseBytes(resp.Body()), nil
}

// PostJSON 发起 POST 请求并解析 JSON
func PostJSON(url string, data interface{}) (gjson.Result, error) {
	resp, err := Post(url, data)
	if err != nil {
		return gjson.Result{}, err
	}
	return gjson.ParseBytes(resp.Body()), nil
}

// GetText 发起 GET 请求并返回文本
func GetText(url string, params map[string]string) (string, error) {
	resp, err := Get(url, params)
	if err != nil {
		return "", err
	}
	return resp.String(), nil
}

// GetHTML 发起 GET 请求并返回 HTML（用于后续 goquery 解析）
func GetHTML(url string, params map[string]string) (io.Reader, error) {
	resp, err := Get(url, params)
	if err != nil {
		return nil, err
	}
	return strings.NewReader(resp.String()), nil
}

// SetProxy 设置代理
func SetProxy(proxy string) {
	if proxy == "" {
		HTTPClient.RemoveProxy()
	} else {
		HTTPClient.SetProxy(proxy)
	}
}

// SetTimeout 设置超时
func SetTimeout(timeout time.Duration) {
	HTTPClient.SetTimeout(timeout)
}

// GetWithContext 发起带 context 的 GET 请求
func GetWithContext(ctx context.Context, url string, params map[string]string) (*resty.Response, error) {
	req := HTTPClient.R().SetContext(ctx)
	if params != nil {
		req.SetQueryParams(params)
	}
	return req.Get(url)
}

// PostWithContext 发起带 context 的 POST 请求
func PostWithContext(ctx context.Context, url string, data interface{}) (*resty.Response, error) {
	return HTTPClient.R().SetContext(ctx).SetBody(data).Post(url)
}

// PostJSONWithHeaders 发起带自定义请求头和查询参数的 POST JSON 请求
func PostJSONWithHeaders(url string, params map[string]string, data interface{}, headers map[string]string) (*resty.Response, error) {
	req := HTTPClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(data)

	if params != nil {
		req.SetQueryParams(params)
	}
	if headers != nil {
		req.SetHeaders(headers)
	}

	return req.Post(url)
}

// PostFormWithHeaders 发起带自定义请求头的 POST 表单请求
func PostFormWithHeaders(url string, formData map[string]string, headers map[string]string) (*resty.Response, error) {
	req := HTTPClient.R().SetFormData(formData)
	if headers != nil {
		req.SetHeaders(headers)
	}
	return req.Post(url)
}

// PostFormWithParams 发起带查询参数、表单数据、请求头和多值字段的 POST 表单请求
func PostFormWithParams(url string, params map[string]string, formData map[string]string, headers map[string]string, multiValueFields map[string][]string) (*resty.Response, error) {
	req := HTTPClient.R()

	if params != nil {
		req.SetQueryParams(params)
	}
	if formData != nil {
		req.SetFormData(formData)
	}
	if headers != nil {
		req.SetHeaders(headers)
	}
	if multiValueFields != nil {
		req.SetFormDataFromValues(multiValueFields)
	}

	return req.Post(url)
}

// PostWithCninfoHeaders 发起带巨潮资讯请求头的 POST 请求
func PostWithCninfoHeaders(url string, params map[string]string) (*resty.Response, error) {
	headers := map[string]string{
		"Accept":          "*/*",
		"Accept-Encoding": "gzip, deflate",
		"Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8",
		"Cache-Control":   "no-cache",
		"Content-Length":  "0",
		"Host":            "webapi.cninfo.com.cn",
		"Origin":          "https://webapi.cninfo.com.cn",
		"Pragma":          "no-cache",
		"Referer":         "https://webapi.cninfo.com.cn/",
		"User-Agent":      DefaultUserAgent,
	}

	req := HTTPClient.R().SetHeaders(headers)
	if params != nil {
		req.SetQueryParams(params)
	}
	return req.Post(url)
}

// GetWithSSEHeaders 发起带上交所请求头的 GET 请求
func GetWithSSEHeaders(url string, params map[string]string) (*resty.Response, error) {
	headers := map[string]string{
		"Referer":    "https://www.sse.com.cn/",
		"User-Agent": DefaultUserAgent,
	}

	req := HTTPClient.R().SetHeaders(headers)
	if params != nil {
		req.SetQueryParams(params)
	}
	return req.Get(url)
}

// GetWithSzseHeaders 发起带深交所请求头的 GET 请求
func GetWithSzseHeaders(url string, params map[string]string) (*resty.Response, error) {
	headers := map[string]string{
		"Referer":    "https://www.szse.cn/",
		"User-Agent": DefaultUserAgent,
	}

	req := HTTPClient.R().SetHeaders(headers)
	if params != nil {
		req.SetQueryParams(params)
	}
	return req.Get(url)
}
