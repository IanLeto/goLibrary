package utils

import (
	"context"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"goLibrary/config"
)

// redisgo 版本
func NewRedisClient(ctx context.Context) (redis.Conn, error) {
	client, err := redis.DialContext(ctx, "tcp", fmt.Sprintf(config.BaseConfig.RedisConfig.Address+":"+config.BaseConfig.RedisConfig.Port))
	return client, err
}

