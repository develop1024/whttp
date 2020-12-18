# whttp

[![Production Ready](https://img.shields.io/badge/production-ready-blue.svg)](https://github.com/develop1024/whttp)

`whttp` 封装了原生的`http`包的请求方法，发送 `http` 请求更加简洁明了，可将得到返回值的字符串，也可将返回值绑定到结构体上。


#### 安装
```
go get github.com/develop1024/whttp
```

#### 使用方法
发送 `Get` 请求

```go
package main

import (
	"fmt"
	"github.com/develop1024/whttp"
)

func main() {
	request := whttp.Request{}
	response := request.Get("http://httpbin.org/get").Resp
	fmt.Printf("%s\n", response)
}
```
---

发送 `Post` 请求
```go
package main

import (
	"fmt"
	"github.com/develop1024/whttp"
)

func main() {
	request := whttp.Request{}
	response := request.Post("http://httpbin.org/post").Resp
	fmt.Printf("%s\n", response)
}
```

---


错误捕获
```go
package main

import (
	"fmt"
	"github.com/develop1024/whttp"
	"log"
)

func main() {
	request := whttp.Request{}
	resp := request.Get("httpsdf://httpbin.org")
	// 捕获错误
	if resp.Error != nil {
		log.Fatal(resp.Error)
	}
	fmt.Println(resp.Resp)
}
```


---

响应结果绑定到 `结构体` 上
```go
package main

import (
	"fmt"
	"github.com/develop1024/whttp"
	"log"
)

type Resp struct {
	Args interface{} `json:"args"`
	Headers Header `json:"headers"`
	Origin string `json:"origin"`
	Url string `json:"url"`
}

type Header struct {
	AcceptEncoding string `json:"Accept-Encoding"`
	Host string `json:"host"`
	UserAgent string `json:"User-Agent"`
	XAmznTraceId string `json:"X-Amzn-Trace-Id"`
}

func main() {
	request := whttp.Request{}
	var resp Resp
	err := request.Get("http://httpbin.org/get").Parse(&resp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.Url)
	fmt.Println(resp.Headers)
	fmt.Println(resp.Origin)
}
```
---

自定义请求参数和请求头
```go
package main

import (
	"fmt"
	"github.com/develop1024/whttp"
)

func main() {
	request := whttp.Request{}

	// 自定义请求参数和请求头
	response := request.CustomGet("http://httpbin.org/get", whttp.CMap{
		"name": "wanghaha",
		"age": "24",
	}, whttp.CMap{
		"Content-Type": "application/json",
	}).Resp
	fmt.Printf("%s\n", response)
}
```