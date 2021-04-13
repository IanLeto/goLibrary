package utils_test

import (
	"context"
	"fmt"
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

// TestHttpClient :
func TestRedisClient(t *testing.T) {
	suite.Run(t, new(RedisClientSuite))
}
