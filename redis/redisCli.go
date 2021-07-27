package redis

import goredis "github.com/go-redis/redis"

func NewRedisClient() *goredis.Client {
	return goredis.NewClient(&goredis.Options{
		Network:    "tcp",
		Addr:       "localhost:6379",
		Password:   "123456",
		DB:         0,
		MaxRetries: 3,
	})
}

