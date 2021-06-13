package utils

import (
	"context"
	"fmt"
	goredis "github.com/go-redis/redis"
	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
	"goLibrary/config"
)

// redisgo 版本
func NewRedisClient(ctx context.Context) (redis.Conn, error) {
	client, err := redis.DialContext(ctx, "tcp", fmt.Sprintf(config.BaseConfig.RedisConfig.Address+":"+config.BaseConfig.RedisConfig.Port))
	return client, err
}

// goredis 版本
func NewRedis() {
	goredis.NewClient(&goredis.Options{
		Network:   "tcp",
		Addr:      fmt.Sprintf("%s:%s", config.BaseConfig.RedisConfig.Address, config.BaseConfig.RedisConfig.Port),
		Dialer:    nil,
		OnConnect: nil,
	})
}

// pipeline 建议使用backend 因为我们要统计
type RedisBackend struct {
	count     int
	pipe      goredis.Pipeliner
	batchSize int
}

func NewRedisPipeline(client goredis.Client) goredis.Pipeliner {
	return client.Pipeline()
}

func (b RedisBackend) batchFull() bool {
	return false
}
func (b RedisBackend) Push(data chan string) {
loop:
	for {
		select {
		case v, ok := <-data:
			if !ok {
				break loop
			}
			// 执行命令
			b.pipe.Set("key", v, 100)
			if b.batchFull() {
				res, err := b.pipe.Exec()
				if err != nil {
					logrus.Errorf("执行pipeline error:%s", err)
					return
				}
				logrus.Debugf("%v : %d data had been pushed to %v totally", b, len(res), "key")
			}
		}
	}
}

func SendMsg() {

}
