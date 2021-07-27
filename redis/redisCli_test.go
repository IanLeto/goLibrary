package redis_test

import (
	goredis "github.com/go-redis/redis"
	"github.com/stretchr/testify/suite"
	"testing"
)

type RedisSuite struct {
	suite.Suite
	redisCli *goredis.Client
}

func (s *RedisSuite) SetupTest() {
	s.redisCli = goredis.NewClient(&goredis.Options{
		Network:   "tcp",
		Addr:      "localhost:6379",
		Dialer:    nil,
		OnConnect: nil,
		Password:  "123456",
		DB:        0,
	})
}

// TestMarshal :
func (s *RedisSuite) TestHelloWorld() {
	v1 := s.redisCli.Exists("root")
	v2 := s.redisCli.Get("root")
	v3 := s.redisCli.Get("root2")
	v1.Val()
	v2.Val()
	v3.Val()


}

func TestRedis(t *testing.T) {
	suite.Run(t, new(RedisSuite))
}
