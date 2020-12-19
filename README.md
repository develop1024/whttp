# whttp

[![Production Ready](https://img.shields.io/badge/production-ready-blue.svg)](https://github.com/develop1024/whttp)

`whttp` 封装了原生的`http`包的请求方法，发送 `http` 请求更加简洁明了，可将得到返回值的字符串，也可将返回值绑定到结构体上。


#### 安装
```
go get github.com/develop1024/whttp
```

#### 使用方法
发送 `GET` 请求

```go
package main

import (
	"fmt"
	"github.com/develop1024/whttp"
	"log"
)

func main() {
	request := whttp.Request{}
	resp := request.Get("http://httpbin.org/get")

	if resp.Error != nil {
		log.Fatal(resp.Error)
	}

	fmt.Println(resp)
	fmt.Println(resp.ToString())
}

```
---

发送 `POST` 请求
```go
package main

import (
	"fmt"
	"github.com/develop1024/whttp"
	"log"
)

func main() {
	request := whttp.Request{}
	resp := request.Post("http://httpbin.org/post")

	if resp.Error != nil {
		log.Fatal(resp.Error)
	}

	fmt.Println(resp)
	fmt.Println(resp.ToString())
}

```

---

发送 `PUT` 请求
```go
package main

import (
	"fmt"
	"github.com/develop1024/whttp"
	"log"
)

func main() {
	request := whttp.Request{}
	resp := request.CustomPut("http://httpbin.org/put", nil)

	if resp.Error != nil {
		log.Fatal(resp.Error)
	}

	fmt.Println(resp)
	fmt.Println(resp.ToString())
}

```

---
发送 `DELETE` 请求
```go
package main

import (
	"fmt"
	"github.com/develop1024/whttp"
	"log"
)

func main() {
	request := whttp.Request{}
	resp := request.CustomDelete("http://httpbin.org/delete", nil)

	if resp.Error != nil {
		log.Fatal(resp.Error)
	}

	fmt.Println(resp)
	fmt.Println(resp.ToString())
}
```

---
发送 `PATCH` 请求
```go
package main

import (
	"fmt"
	"github.com/develop1024/whttp"
	"log"
)

func main() {
	request := whttp.Request{}
	resp := request.CustomPatch("http://httpbin.org/patch", nil)

	if resp.Error != nil {
		log.Fatal(resp.Error)
	}

	fmt.Println(resp)
	fmt.Println(resp.ToString())
}
```

---


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

	var entityResp Resp

	err := request.Get("http://httpbin.org/get").Parse(&entityResp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(entityResp.Url)
	fmt.Println(entityResp.Headers.Host)
}

```
---

自定义请求参数和请求头
```go
package main

import (
	"fmt"
	"github.com/develop1024/whttp"
	"log"
)

func main() {
	request := whttp.Request{}

	resp := request.CustomGet("http://httpbin.org/get", whttp.CMap{
		"name": "wangahah",
		"age": 24,
	}, whttp.CMap{
		"Content-Type": "application/json",
	})

	if resp.Error != nil {
		log.Fatal(resp.Error)
	}

	fmt.Println(resp)
}
```