package http

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"time"
)

type GetIanOriRequest struct {
}
type GetIanOriResponse struct {
}

func main() {
	// 创建一个 Resty 客户端
	client := resty.New()

	// 设置全局超时时间
	client.SetTimeout(5 * time.Second)

	// 设置重试策略：最多重试3次，每次重试间隔500毫秒
	client.SetRetryCount(3)
	client.SetRetryWaitTime(500 * time.Millisecond)
	client.AddRetryCondition(
		// 定义重试条件，这里是当返回5xx错误时重试
		func(r *resty.Response, err error) bool {
			return r.StatusCode() >= 500
		},
	)

	// 设置全局 Headers
	client.SetHeaders(map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	})

	// 发送 GET 请求
	resp, err := client.R().
		SetQueryParam("key", "value").
		SetHeader("Custom-Header", "value").
		Get("https://httpbin.org/get")
	if err != nil {
		log.Fatalf("Error on GET request: %v", err)
	}
	fmt.Println("GET Response Status:", resp.Status())
	fmt.Println("GET Response Body:", string(resp.Body()))

	// 发送 POST 请求
	resp, err = client.R().
		SetBody(`{"name": "John", "age": 30}`).
		Post("https://httpbin.org/post")
	if err != nil {
		log.Fatalf("Error on POST request: %v", err)
	}
	fmt.Println("POST Response Status:", resp.Status())
	fmt.Println("POST Response Body:", string(resp.Body()))

	// 演示客户端级别的自动 JSON & XML 解析
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	var user User

	// JSON 自动解析
	_, err = client.R().
		SetResult(&user).
		Get("https://httpbin.org/anything")
	if err != nil {
		log.Fatalf("Error on parsing JSON: %v", err)
	}
	fmt.Printf("Parsed User: %+v\n", user)
}
