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

	// 打印response输出string类型
	fmt.Println(resp)
	// 打印response输出string类型
	fmt.Println(resp.ToString())
	// 打印response输出[]byte类型
	fmt.Println(resp.Resp)
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

	// 打印response输出string类型
	fmt.Println(resp)
	// 打印response输出string类型
	fmt.Println(resp.ToString())
	// 打印response输出[]byte类型
	fmt.Println(resp.Resp)
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
	resp := request.PutRequest("http://httpbin.org/put")
	if resp.Error != nil {
		log.Fatal(resp.Error)
	}

	// 打印response输出string类型
	fmt.Println(resp)
	// 打印response输出string类型
	fmt.Println(resp.ToString())
	// 打印response输出[]byte类型
	fmt.Println(resp.Resp)
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
	resp := request.DeleteRequest("http://httpbin.org/delete")
	if resp.Error != nil {
		log.Fatal(resp.Error)
	}

	// 打印response输出string类型
	fmt.Println(resp)
	// 打印response输出string类型
	fmt.Println(resp.ToString())
	// 打印response输出[]byte类型
	fmt.Println(resp.Resp)
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
	resp := request.PatchRequest("http://httpbin.org/patch")
	if resp.Error != nil {
		log.Fatal(resp.Error)
	}

	// 打印response输出string类型
	fmt.Println(resp)
	// 打印response输出string类型
	fmt.Println(resp.ToString())
	// 打印response输出[]byte类型
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
	resp := request.Get("http://httpbin.org/get")
	if resp.Error != nil {
		log.Fatal(resp.Error)
	}

	var data Resp
	err := resp.Parse(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data.Headers.Host)
	fmt.Println(data.Url)
	fmt.Println(data.Origin)
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
	// whttp.Params 里传入URL参数
	// whttp.Data 里传入Body参数
	request := whttp.Request{}
	resp := request.GetRequest("http://httpbin.org/get", whttp.Params{
		"limit": 100,
	},
	whttp.Data{
		"name": "wanghaha",
		"age": 24,
	},
	whttp.Headers{
		"Content-Type": "application/json",
	})
	if resp.Error != nil {
		log.Fatal(resp.Error)
	}

	fmt.Println(resp)
	fmt.Println(resp.ToString())
}
```

设置 `cookie` 
```go
package main

import (
	"fmt"
	"github.com/develop1024/whttp"
	"log"
	"net/http"
)

func main() {
	request := whttp.Request{}
	cookies := whttp.Cookies {
		http.Cookie {
			Name: "name",
			Value: "xxxx",
		},
	}
	resp := request.GetRequest("http://httpbin.org/get", cookies)
	if resp.Error != nil {
		log.Fatal(resp.Error)
	}
	fmt.Println(resp)
}
```


设置请求超时时间
```go
package main

import (
	"fmt"
	"github.com/develop1024/whttp"
	"log"
	"time"
)

func main() {
	request := whttp.Request{}

	resp := request.GetRequest("http://httpbin.org/get", time.Microsecond * 100)
	if resp.Error != nil {
		log.Fatal(resp.Error)
	}
	fmt.Println(resp)
}

```