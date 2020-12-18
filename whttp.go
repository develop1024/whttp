package whttp

import (
	"encoding/json"
	"io/ioutil"
"net/http"
"net/url"
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
func (r *Request) CustomGet(URL string, params CMap, headers CMap) *Response {
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
		q.Add(key, val.(string))
	}

	// 构造请求参数赋值给请求url
	request.URL.RawQuery = q.Encode()

	// 添加请求头
	request.Header.Add("cache-control", "no-cache")
	// 设置请求头
	request.Header.Set("content-type", "application/x-www-form-urlencoded")

	for key, val := range headers {
		request.Header.Add(key, val.(string))
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
func (r *Request) CustomPost(URL string, params CMap, headers CMap) *Response {

	u := url.Values{}
	for key, val := range params {
		u.Add(key, val.(string))
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
	for key, val := range headers {
		request.Header.Add(key, val.(string))
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
