package httpClient_test

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
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

// TestHttpClient :
func TestHttpClient2(t *testing.T) {
	suite.Run(t, new(HttpClient2))
}
