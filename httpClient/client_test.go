package httpClient_test

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	url2 "net/url"
	"os"
	"testing"
	"time"
)

// HttpClientSuite :
type HttpClient2 struct {
	suite.Suite
	client http.Client
}

func (s *HttpClient2) SetupTest() {
	s.client = http.Client{}
}

// http: 任何类型的请求都可以发送的http 示例
func (s *HttpClient2) TestAnyClient() {
	request, err := http.NewRequest("GET", "http://example.com", nil)
	s.NoError(err)
	resp, err := s.client.Do(request)
	s.NoError(err)
	s.Equal(200, resp.StatusCode)
	res, err := httputil.DumpResponse(resp, true)
	fmt.Printf("%s\n", res)

}

// TestMarshal : 测试有http keepalive 的速度 30s
func (s *HttpClient2) TestKeepAlive() {
	s.client = http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 100,
		},
	}
	start := time.Now()
	for i := 0; i < 100; i++ {
		resp, err := s.client.Get("http://example.com")
		if err != nil {
			fmt.Println("Request failed:", err)
			return
		}

		_, err = ioutil.ReadAll(resp.Body)
		_ = resp.Body.Close()
		if err != nil {
			fmt.Println("Failed to read response:", err)
			return
		}
	}
	elapsed := time.Since(start)
	fmt.Println("Total time:", elapsed)
}

// TestMarshal : 测试无http keepalive 的速度 52s
func (s *HttpClient2) TestNoKeepAlive() {

	start := time.Now()
	for i := 0; i < 100; i++ {
		resp, err := s.client.Get("http://example.com")
		if err != nil {
			fmt.Println("Request failed:", err)
			return
		}

		_, err = ioutil.ReadAll(resp.Body)
		_ = resp.Body.Close()
		if err != nil {
			fmt.Println("Failed to read response:", err)
			return
		}
	}
	elapsed := time.Since(start)
	fmt.Println("Total time:", elapsed)
}

// http: 设置代理
func (s *HttpClient2) TestProxy() {
	proxyUrl, err := url2.Parse("localhost:8080")
	s.NoError(err)
	s.client.Transport = &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}
	resp, err := s.client.Get("http://example.com")
	s.NoError(err)
	s.Equal(200, resp.StatusCode)
	s.NoError(resp.Write(os.Stdout))
}

// http:
func (s *HttpClient2) TestRedirect() {

}

// TestHttpClient :
func TestHttpClient2(t *testing.T) {
	suite.Run(t, new(HttpClient2))
}
