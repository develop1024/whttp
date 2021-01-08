package whttp

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// URL参数别名
type Params map[string]interface{}

// Body参数别名
type Data map[string]interface{}

// 请求头别名
type Headers map[string]interface{}

type Cookies []http.Cookie

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

// 将json格式的相应内容转为map类型
func (resp *Response) GetJsonToMap(key string) gjson.Result {
	result := gjson.GetBytes(resp.Resp, key)
	return result
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

// 普通GET请求
func (r *Request) Get(URL string) *Response {
	return r.CustomRequest(URL, "GET")
}

// 普通POST请求
func (r *Request) Post(URL string) *Response {
	return r.CustomRequest(URL, "POST")
}

// GET请求支持自定义参数和请求头
func (r *Request) GetRequest(URL string, v ...interface{}) *Response {
	return r.CustomRequest(URL, "GET", v...)
}

// POST请求支持自定义参数和请求头
func (r *Request) PostRequest(URL string, v ...interface{}) *Response {
	return r.CustomRequest(URL, "POST", v...)
}

// PUT请求支持自定义参数和请求头
func (r *Request) PutRequest(URL string, v ...interface{}) *Response {
	return r.CustomRequest(URL, "PUT", v...)
}

// DELETE请求支持自定义参数和请求头
func (r *Request) DeleteRequest(URL string, v ...interface{}) *Response {
	return r.CustomRequest(URL, "DELETE", v...)
}

// PATCH请求支持自定义参数和请求头
func (r *Request) PatchRequest(URL string, v ...interface{}) *Response {
	return r.CustomRequest(URL, "PATCH", v...)
}

// 自定义请求支持自定义参数和请求头
func (r *Request) CustomRequest(URL string, METHOD string,  v ...interface{}) *Response {
	u := url.Values{}
	var request *http.Request

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
		case time.Duration:
			// 设置超时时间
			http.DefaultClient.Timeout = item.(time.Duration)
		case Cookies:
			// 设置cookie
			for _, cookie := range item.(Cookies) {
				request.AddCookie(&cookie)
			}
		default:

		}
	}

	// 构建参数
	payload := strings.NewReader(u.Encode())

	// 数据重新赋值
	request, err := http.NewRequest(METHOD, URL, payload)
	if err != nil {
		return &Response{
			Resp:  nil,
			Error: err,
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
