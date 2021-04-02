package utils_test

import (
	"context"
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/cast"
	"github.com/stretchr/testify/suite"
	"goLibrary/config"
	"goLibrary/testsuites"
	"goLibrary/utils"
	"testing"
)

// HttpClientSuite :
type RedisClientSuite struct {
	suite.Suite
	testsuites.ConfigSuit
	client redis.Conn
}

func (s *RedisClientSuite) SetupTest() {
	var (
		err error
	)
	config.InitConfig("../config/config.yaml")
	ctx := context.Background()
	s.client, err = utils.NewRedisClient(ctx)
	s.NoError(err)
}

// TestMarshal :
func (s *RedisClientSuite) TestPing() {
	res, err := s.client.Do("PING")
	s.NoError(err)
	s.Equal("PONG", cast.ToString(res))
}

// TestHttpClient :
func TestRedisClient(t *testing.T) {
	suite.Run(t, new(RedisClientSuite))
}
