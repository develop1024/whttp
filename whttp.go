package whttp

import (
"io/ioutil"
"net/http"
"net/url"
"strings"
)

// 自定义参数map
type CMap map[string]interface{}

// Get请求
func Get(URL string) (response []byte, err error) {
	resp, err := http.Get(URL)
	if err != nil {
		return nil, err
	}

	response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return response, nil
}

// 自定义Get请求
func CustomGet(URL string, params CMap, headers CMap) (response []byte, err error) {
	request, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return response, nil
}



// Post请求
func Post(URL string) (response []byte, err error) {
	resp, err := http.Post(URL, "application/json", nil)
	if err != nil {
		return nil, err
	}

	response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return response, nil
}

// 自定义Post请求
func CustomPost(URL string, params CMap, headers CMap) (response []byte, err error) {

	u := url.Values{}
	for key, val := range params {
		u.Add(key, val.(string))
	}

	payload := strings.NewReader(u.Encode())
	request, err := http.NewRequest("POST", URL, payload)
	if err != nil {
		return nil, err
	}

	// 设置请求头
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for key, val := range headers {
		request.Header.Add(key, val.(string))
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return response, nil
}
