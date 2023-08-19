package httpClient

import "github.com/dghubble/sling"

type APIClient interface {
	GetHelloWorld()
}

// sling 存路经用
type XXClient struct {
	commonArgs *CommonArgs
	agent      *sling.Sling
}

type CommonArgs struct {
	Address   string `json:"address"`
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
}

// Agent :
func (c *XXClient) Agent() *sling.Sling {
	return c.agent.New()
}

// Post :
func (c *XXClient) Post(path string) *sling.Sling {
	return c.Agent().Post(path)
}

// Get :
func (c *XXClient) Get(path string) *sling.Sling {
	return c.Agent().Get(path)
}

func NewXXClient(doer sling.Doer) *XXClient {
	return &XXClient{
		commonArgs: nil,
		agent:      sling.New().Base("baseUrl").Path("路径前缀").Doer(doer),
	}
}
