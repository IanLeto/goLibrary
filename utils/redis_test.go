package utils_test

import (
	"context"
	"fmt"
	goredis "github.com/go-redis/redis"
	"github.com/gomodule/redigo/redis"
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
	client       redis.Conn
	client2      *goredis.Client
	remoteClient *goredis.Client
}

func (s *RedisClientSuite) SetupTest() {
	var (
		err error
	)
	config.InitConfig("../config/config.yaml")
	ctx := context.Background()
	s.client, err = utils.NewRedisClient(ctx)
	s.NoError(err)

	s.client2 = utils.NewRedis()
}

// TestMarshal :
func (s *RedisClientSuite) TestPing() {
	var xxx = struct {
		x int
	}{
		x: 1,
	}
	var a = struct {
		a int
	}{

	}
	if &xxx == nil {
		fmt.Println(1)
	}
	if &a == nil {
		fmt.Println(2)
	}
}

func (s *RedisClientSuite) TestPing2() {
	res := s.client2.Ping()
	s.NoError(res.Err())
	s.Equal(res.Val(), "PONG")
}
func (s *RedisClientSuite) TestPing3() {
	config.BaseConfig.RedisConfig.Address = "10.68.132.168"
	cli := utils.NewRedis()
	res := cli.Ping()
	s.NoError(res.Err())
	s.Equal(res.Val(), "PONG")
}

// TestHttpClient :
func TestRedisClient(t *testing.T) {
	suite.Run(t, new(RedisClientSuite))
}
