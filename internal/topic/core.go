package topic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/Abhishekvrshny/optimus/internal/message"
	"github.com/Abhishekvrshny/optimus/pkg/redis"
	guuid "github.com/google/uuid"
)

type Core struct {
	q                  *redis.Redis
	topicSubscriberMap map[string]map[string]bool
	topicMap           map[string]Topic
}

func NewCore(q *redis.Redis) *Core {
	return &Core{q,
		make(map[string]map[string]bool),
		make(map[string]Topic),
	}
}

func (c *Core) createTopic(req Topic) error {
	if _, ok := c.topicSubscriberMap[req.name]; ok {
		return fmt.Errorf("topic already exists")
	}
	c.topicSubscriberMap[req.name] = make(map[string]bool)
	c.topicMap[req.name] = req
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

func (c *Core) Publish(topic string, body bytes.Buffer, header map[string][]string) error {
	reqId := guuid.New().String()
	msg := message.Message{
		body.String(),
		header,
		reqId,
	}

	b, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	pubsub := c.q.Subscribe(reqId)
	channel := pubsub.Channel()

	c.q.Client.Publish(topic, string(b))

	timer := time.NewTimer(time.Duration(c.topicMap[topic].TimeoutInMs) * time.Millisecond)
	count := 0
	for {
		if count == len(c.topicSubscriberMap[topic]) {
			log.Println("Success: Consumed by all subscribers")
			return nil
		}
		select {
		case <-channel:
			count += 1
		case <-timer.C:
			log.Println("Fallback: Write to Kafka and return")
			return nil
		}
	}

	return nil
}

func (c *Core) CreateSubscriber(topicName string, subsName string) {
	c.topicSubscriberMap[topicName][subsName] = true
}
