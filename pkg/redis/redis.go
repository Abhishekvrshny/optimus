package redis

import (
	"fmt"
	"github.com/go-redis/redis"
)

type Redis struct {
	Client *redis.Client
}

// NewRedisQueue inits a RedisQueue driver
func NewRedisQueue() *Redis {
	options := &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", "localhost", 26379),
		Password: "",
		DB:       int(10),
	}
	client := redis.NewClient(options)
	_, err := client.Ping().Result()
	if err == nil {
		return &Redis{Client: client}
	}
	return nil
}

func (r *Redis) Subscribe(topic string) *redis.PubSub {
	pubsub := r.Client.Subscribe(topic)
	_, err := pubsub.Receive()
	if err != nil {
		panic(err)
	}
	return pubsub
}


