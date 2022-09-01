package httpClientDemo

import (
	"github.com/dghubble/sling"
	"github.com/sirupsen/logrus"
)

var args = CommonArgs{
	Address:   "localhost:8080",
	AppKey:    "1",
	AppSecret: "2",
}

// sling 存路经用
type LocalhostClient struct {
	commonArgs *CommonArgs
	agent      *sling.Sling
}

// Agent :
func (c *LocalhostClient) Agent() *sling.Sling {
	return c.agent.New()
}

// Post :
func (c *LocalhostClient) Post(path string) *sling.Sling {
	return c.Agent().Post(path)
}

// Get :
func (c *LocalhostClient) Get(path string) *sling.Sling {
	return c.Agent().Get(path)
}

func NewLocalhostClient(doer sling.Doer) *LocalhostClient {
	return &LocalhostClient{
		commonArgs: nil,
		// base 路径
		agent: sling.New().Base("http://localhost:8080").Path("").Doer(doer),
	}
}

// 定义baseclient
// 手把手 写一个http get 请求
// step1 定义 req and res

type LocalhostHelloRequestInfo struct {
}

type LocalhostResponseInfo struct {
	Msg  string      `json:"msg"`
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

//step2 new http cli

func (c *LocalhostClient) GetHelloWorld() (*LocalhostResponseInfo, error) {
	// 定义返回结构
	var result = &LocalhostResponseInfo{}
	_, err := c.Get("hello").QueryStruct(LocalhostHelloRequestInfo{}).Receive(result, result)
	if err != nil {
		logrus.Errorf("请求错误err:%s", err)
		return nil, err
	}
	return result, err
}
