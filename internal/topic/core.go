package topic

import "github.com/Abhishekvrshny/optimus/pkg/redis"

type core struct {
	q *redis.Redis
}

func NewCore(q *redis.Redis) *core {
	return &core{q}
}

func (c *core) createTopic(name string) {

}


func (c *core) getTopic(name string) {

}
