package redis

import (
	"fmt"
	"github.com/go-redis/redis"
)

type Redis struct {
	client *redis.Client
}

// NewRedisQueue inits a RedisQueue driver
func NewRedisQueue() *Redis {
	options := &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", "http://localhost", 26379),
		Password: "",
		DB:       int(10),
	}
	client := redis.NewClient(options)
	_, err := client.Ping().Result()
	if err == nil {
		return &Redis{client:client}
	}
	return nil
}

