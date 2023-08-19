package httpClient_test

import (
	"github.com/stretchr/testify/suite"
	"goLibrary/httpClient"
	"testing"
)

// HttpClientSuite :
type HttpClientSuite struct {
	suite.Suite
	c *httpClient.LocalhostClient
}

func (s *HttpClientSuite) SetupTest() {
	// 直接初始化
	// 我们的base 路径已经写好了
	s.c = httpClient.NewLocalhostClient(nil)
}

// TestMarshal :
func (s *HttpClientSuite) TestConfig() {
	res, err := s.c.GetHelloWorld()
	s.NoError(err)
	s.Equal(0, res.Code)
}

// TestHttpClient :
func TestViperConfiguration(t *testing.T) {
	suite.Run(t, new(HttpClientSuite))
}
