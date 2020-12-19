package whttp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)
// URL参数别名
type Params map[string]interface{}

// Body参数别名
type Data map[string]interface{}

// 请求头别名
type Headers map[string]interface{}

// 请求结构体
type Request struct {}

// 响应结构体
type Response struct {
	Resp []byte
	Error error
}

func (resp *Response) String () string {
	return resp.ToString()
}

// 将响应的结果解析到结构体上
func (resp *Response) Parse(StructData interface{}) error {
	err := json.Unmarshal(resp.Resp, StructData)
	if err != nil {
		return err
	}
	return nil
}

// 将响应结果转为string
func (resp *Response) ToString() string {
	return string(resp.Resp)
}

// 普通GET请求
func (r *Request) Get(URL string) *Response {
	resp, err := http.Get(URL)
	if err != nil {
		return &Response{
			Resp:  nil,
			Error: err,
		}
	}

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &Response{
			Resp:  nil,
			Error: err,
		}
	}

	defer resp.Body.Close()

	return &Response{
		Resp: response,
		Error: err,
	}
}

// 普通POST请求
func (r *Request) Post(URL string) *Response {
	resp, err := http.Post(URL, "application/json", nil)
	if err != nil {
		return &Response{
			Resp:  nil,
			Error: err,
		}
	}

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &Response{
			Resp:  nil,
			Error: err,
		}
	}
	defer resp.Body.Close()

	return &Response{
		Resp:  response,
		Error: nil,
	}
}

// 类型转换为 string
func ToStrType(v interface{}) string {
	switch v.(type) {
	case string:
		return v.(string)
	case int:
		return strconv.Itoa(v.(int))
	case float32:
		return strconv.FormatFloat(float64(v.(float32)), 'f', 6, 64)
	case float64:
		return strconv.FormatFloat(v.(float64), 'f', 6, 64)
	case bool:
		return strconv.FormatBool(v.(bool))
	default:
		return v.(string)
	}
}

// GET请求支持自定义参数和请求头
func (r *Request) GetRequest(URL string, v ...interface{}) *Response {
	request, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return &Response{
			Resp:  nil,
			Error: err,
		}
	}

	for _, item := range v {
		switch item.(type) {
		case Params:
			// 设置请求参数
			q := request.URL.Query()
			for key, val := range item.(Params) {
				q.Add(key, ToStrType(val))
			}

			// 构造请求参数赋值给请求url
			request.URL.RawQuery = q.Encode()

			for key, child := range item.(Params) {
				q.Add(key, ToStrType(child))
			}
		case Headers:
			// 添加请求头
			request.Header.Add("cache-control", "no-cache")
			// 设置请求头
			request.Header.Set("content-type", "application/x-www-form-urlencoded")

			for key, val := range item.(Headers) {
				request.Header.Add(key, ToStrType(val))
			}
		default:

		}
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return &Response{
			Resp:  nil,
			Error: err,
		}
	}

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &Response{
			Resp:  nil,
			Error: err,
		}
	}
	defer resp.Body.Close()

	return &Response{
		Resp:  response,
		Error: nil,
	}
}

// POST请求支持自定义参数和请求头
func (r *Request) PostRequest(URL string, v ...interface{}) *Response {
	u := url.Values{}
	payload := strings.NewReader(u.Encode())
	request, err := http.NewRequest("POST", URL, payload)
	if err != nil {
		return &Response{
			Resp:  nil,
			Error: err,
		}
	}

	for _, item := range v {
		switch item.(type) {
		case Params:
			// 设置URL请求参数
			q := request.URL.Query()
			for key, val := range item.(Params) {
				q.Add(key, ToStrType(val))
			}

			// 构造请求参数赋值给请求url
			request.URL.RawQuery = q.Encode()

			for key, child := range item.(Params) {
				q.Add(key, ToStrType(child))
			}
		case Data:
			// 添加 body 请求参数
			for key, val := range item.(Data) {
				u.Add(key, ToStrType(val))
			}
		case Headers:
			// 设置请求头
			request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			for key, val := range item.(Headers) {
				request.Header.Add(key, ToStrType(val))
			}
		default:

		}
	}

	// 发送请求
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return &Response{
			Resp:  nil,
			Error: err,
		}
	}
	defer resp.Body.Close()

	// 读取请求
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &Response{
			Resp:  nil,
			Error: err,
		}
	}

	return &Response{
		Resp:  response,
		Error: nil,
	}
}

// PUT请求支持自定义参数和请求头
func (r *Request) PutRequest(URL string, v ...interface{}) *Response {
	u := url.Values{}
	payload := strings.NewReader(u.Encode())
	request, err := http.NewRequest("PUT", URL, payload)
	if err != nil {
		return &Response{
			Resp:  nil,
			Error: err,
		}
	}

	for _, item := range v {
		switch item.(type) {
		case Params:
			// 设置URL请求参数
			q := request.URL.Query()
			for key, val := range item.(Params) {
				q.Add(key, ToStrType(val))
			}

			// 构造请求参数赋值给请求url
			request.URL.RawQuery = q.Encode()

			for key, child := range item.(Params) {
				q.Add(key, ToStrType(child))
			}
		case Data:
			// 添加 body 请求参数
			for key, val := range item.(Data) {
				u.Add(key, ToStrType(val))
			}
		case Headers:
			// 设置请求头
			request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			for key, val := range item.(Headers) {
				request.Header.Add(key, ToStrType(val))
			}
		default:

		}
	}

	// 发送请求
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return &Response{
			Resp:  nil,
			Error: err,
		}
	}
	defer resp.Body.Close()

	// 读取请求
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &Response{
			Resp:  nil,
			Error: err,
		}
	}

	return &Response{
		Resp:  response,
		Error: nil,
	}
}

// DELETE请求支持自定义参数和请求头
func (r *Request) DeleteRequest(URL string, v ...interface{}) *Response {
	u := url.Values{}
	payload := strings.NewReader(u.Encode())
	request, err := http.NewRequest("DELETE", URL, payload)
	if err != nil {
		return &Response{
			Resp:  nil,
			Error: err,
		}
	}

	for _, item := range v {
		switch item.(type) {
		case Params:
			// 设置URL请求参数
			q := request.URL.Query()
			for key, val := range item.(Params) {
				q.Add(key, ToStrType(val))
			}

			// 构造请求参数赋值给请求url
			request.URL.RawQuery = q.Encode()

			for key, child := range item.(Params) {
				q.Add(key, ToStrType(child))
			}
		case Data:
			// 添加 body 请求参数
			for key, val := range item.(Data) {
				u.Add(key, ToStrType(val))
			}
		case Headers:
			// 设置请求头
			request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			for key, val := range item.(Headers) {
				request.Header.Add(key, ToStrType(val))
			}
		default:

		}
	}

	// 发送请求
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return &Response{
			Resp:  nil,
			Error: err,
		}
	}
	defer resp.Body.Close()

	// 读取请求
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &Response{
			Resp:  nil,
			Error: err,
		}
	}

	return &Response{
		Resp:  response,
		Error: nil,
	}
}

// PATCH请求支持自定义参数和请求头
func (r *Request) PatchRequest(URL string, v ...interface{}) *Response {
	u := url.Values{}
	payload := strings.NewReader(u.Encode())
	request, err := http.NewRequest("PATCH", URL, payload)
	if err != nil {
		return &Response{
			Resp:  nil,
			Error: err,
		}
	}

	for _, item := range v {
		switch item.(type) {
		case Params:
			// 设置URL请求参数
			q := request.URL.Query()
			for key, val := range item.(Params) {
				q.Add(key, ToStrType(val))
			}

			// 构造请求参数赋值给请求url
			request.URL.RawQuery = q.Encode()

			for key, child := range item.(Params) {
				q.Add(key, ToStrType(child))
			}
		case Data:
			// 添加 body 请求参数
			for key, val := range item.(Data) {
				u.Add(key, ToStrType(val))
			}
		case Headers:
			// 设置请求头
			request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			for key, val := range item.(Headers) {
				request.Header.Add(key, ToStrType(val))
			}
		default:

		}
	}

	// 发送请求
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return &Response{
			Resp:  nil,
			Error: err,
		}
	}
	defer resp.Body.Close()

	// 读取请求
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &Response{
			Resp:  nil,
			Error: err,
		}
	}

	return &Response{
		Resp:  response,
		Error: nil,
	}
}

// 自定义请求支持自定义参数和请求头
func (r *Request) CustomRequest(URL string, METHOD string,  v ...interface{}) *Response {
	u := url.Values{}
	payload := strings.NewReader(u.Encode())
	request, err := http.NewRequest(METHOD, URL, payload)
	if err != nil {
		return &Response{
			Resp:  nil,
			Error: err,
		}
	}

	for _, item := range v {
		switch item.(type) {
		case Params:
			// 设置URL请求参数
			q := request.URL.Query()
			for key, val := range item.(Params) {
				q.Add(key, ToStrType(val))
			}

			// 构造请求参数赋值给请求url
			request.URL.RawQuery = q.Encode()

			for key, child := range item.(Params) {
				q.Add(key, ToStrType(child))
			}
		case Data:
			// 添加 body 请求参数
			for key, val := range item.(Data) {
				u.Add(key, ToStrType(val))
			}
		case Headers:
			// 设置请求头
			request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			for key, val := range item.(Headers) {
				request.Header.Add(key, ToStrType(val))
			}
		default:

		}
	}

	// 发送请求
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return &Response{
			Resp:  nil,
			Error: err,
		}
	}
	defer resp.Body.Close()

	// 读取请求
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &Response{
			Resp:  nil,
			Error: err,
		}
	}

	return &Response{
		Resp:  response,
		Error: nil,
	}
}
