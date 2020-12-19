package whttp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// 自定义参数map
type CMap map[string]interface{}

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

// Get请求
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

// 自定义Get请求
func (r *Request) CustomGet(URL string, params CMap, headers ...CMap) *Response {
	request, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return &Response{
			Resp:  nil,
			Error: err,
		}
	}

	// 设置请求参数
	q := request.URL.Query()
	for key, val := range params {
		q.Add(key, ToStrType(val))
	}

	// 构造请求参数赋值给请求url
	request.URL.RawQuery = q.Encode()

	// 添加请求头
	request.Header.Add("cache-control", "no-cache")
	// 设置请求头
	request.Header.Set("content-type", "application/x-www-form-urlencoded")

	for _, hs := range headers {
		for key, val := range hs {
			request.Header.Add(key, ToStrType(val))
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


// Post请求
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

// 自定义Post请求
func (r *Request) CustomPost(URL string, params CMap, headers ...CMap) *Response {

	u := url.Values{}
	for key, val := range params {
		u.Add(key, ToStrType(val))
	}

	payload := strings.NewReader(u.Encode())
	request, err := http.NewRequest("POST", URL, payload)
	if err != nil {
		return &Response{
			Resp:  nil,
			Error: err,
		}
	}

	// 设置请求头
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for _, hs := range headers {
		for key, val := range hs {
			request.Header.Add(key, ToStrType(val))
		}
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return &Response{
			Resp:  nil,
			Error: err,
		}
	}
	defer resp.Body.Close()

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


// 自定义Put请求
func (r *Request) CustomPut(URL string, params CMap, headers ...CMap) *Response {

	u := url.Values{}
	for key, val := range params {
		u.Add(key, ToStrType(val))
	}

	payload := strings.NewReader(u.Encode())
	request, err := http.NewRequest("PUT", URL, payload)
	if err != nil {
		return &Response{
			Resp:  nil,
			Error: err,
		}
	}

	// 设置请求头
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for _, hs := range headers {
		for key, val := range hs {
			request.Header.Add(key, ToStrType(val))
		}
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return &Response{
			Resp:  nil,
			Error: err,
		}
	}
	defer resp.Body.Close()

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


// 自定义Delete请求
func (r *Request) CustomDelete(URL string, params CMap, headers ...CMap) *Response {

	u := url.Values{}
	for key, val := range params {
		u.Add(key, ToStrType(val))
	}

	payload := strings.NewReader(u.Encode())
	request, err := http.NewRequest("DELETE", URL, payload)
	if err != nil {
		return &Response{
			Resp:  nil,
			Error: err,
		}
	}

	// 设置请求头
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for _, hs := range headers {
		for key, val := range hs {
			request.Header.Add(key, ToStrType(val))
		}
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return &Response{
			Resp:  nil,
			Error: err,
		}
	}
	defer resp.Body.Close()

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

// 自定义PATCH请求
func (r *Request) CustomPatch(URL string, params CMap, headers ...CMap) *Response {

	u := url.Values{}
	for key, val := range params {
		u.Add(key, ToStrType(val))
	}

	payload := strings.NewReader(u.Encode())
	request, err := http.NewRequest("PATCH", URL, payload)
	if err != nil {
		return &Response{
			Resp:  nil,
			Error: err,
		}
	}

	// 设置请求头
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for _, hs := range headers {
		for key, val := range hs {
			request.Header.Add(key, ToStrType(val))
		}
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return &Response{
			Resp:  nil,
			Error: err,
		}
	}
	defer resp.Body.Close()

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
