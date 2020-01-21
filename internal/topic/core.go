package topic

import (
	"bytes"
	"fmt"
	"github.com/Abhishekvrshny/optimus/pkg/redis"
)

type Core struct {
	q *redis.Redis
	topicSubscriberMap map[string]map[string]string
}

func NewCore(q *redis.Redis) *Core {
	return &Core{q, make(map[string]map[string]string)}
}

func (c *Core) createTopic(req Topic) error{
	fmt.Println(req)
	if _, ok := c.topicSubscriberMap[req.name]; ok {
		return fmt.Errorf("topic already exists")
	}
	c.topicSubscriberMap[req.name] = make(map[string]string)
	return nil
}


func (c *Core) getTopic(name string) {

}

func (c *Core) TopicExists(topic string) bool {
	if _, ok := c.topicSubscriberMap[topic]; ok {
		return true
	}
	return false
}

func (c *Core) SubscriberExists(topic string, subs string) bool {
	if _, ok := c.topicSubscriberMap[topic][subs]; ok {
		return true
	}
	return false
}

func (c *Core) Publish(s string, buffer bytes.Buffer) {
	c.q.Client.Publish(s, buffer.String())
}
