package utils

import (
	"context"
	"fmt"
	goredis "github.com/go-redis/redis"
	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
	"goLibrary/config"
	"math/rand"
	"time"
)

// redisgo 版本
func NewRedisClient(ctx context.Context) (redis.Conn, error) {
	client, err := redis.DialContext(ctx, "tcp", fmt.Sprintf(config.BaseConfig.RedisConfig.Address+":"+config.BaseConfig.RedisConfig.Port))
	return client, err
}

// goredis 版本
func NewRedis() *goredis.Client {
	return goredis.NewClient(&goredis.Options{
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

// 如何运行一个后台cache 周期任务
type CacheTask struct {
	Interval   time.Duration
	Pipe       goredis.Pipeliner
	Worker     chan struct{}
	Expiration time.Duration
}

func InitCacheTask() *CacheTask {
	return &CacheTask{
		Interval: 60 * time.Second,
		Pipe:     NewRedisPipeline(*NewRedis()).Pipeline(),
		Worker:   make(chan struct{}, 20),
	}
}

func (c *CacheTask) Run() error {
	// 生成测试数据
	trick := time.NewTicker(c.Interval)
	for {
		select {
		case <-trick.C:
			// 生成测试数据
			demoData := map[int]int{}
			for i := 0; i < 50; i++ {
				demoData[i] = rand.Intn(100)
			}
		}
	}

	//

}

func (c *CacheTask) Stop() error {
	panic("implement me")
}

func SendMsg() {

}

func CacheCron() {

}
