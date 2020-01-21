package subscriber

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Abhishekvrshny/optimus/internal/message"
	"github.com/Abhishekvrshny/optimus/internal/topic"
	"github.com/Abhishekvrshny/optimus/pkg/redis"
	redis2 "github.com/go-redis/redis"
	"net/http"
)

type Core struct {
	q *redis.Redis
	topicCore *topic.Core
}

func NewCore(q *redis.Redis, topicCore *topic.Core) *Core{
	return &Core{q, topicCore}
}

func (c *Core) CreateSubscriber(req Subscriber) {
	c.topicCore.CreateSubscriber(req.topic, req.name)
	pubsub := c.q.Subscribe(req.topic)
	go c.readSubs(req.Url, pubsub.Channel())
}

func (c *Core) readSubs(url string, channel <-chan *redis2.Message) {
	for msg := range channel {
		fmt.Println("subscription read")
		rec := message.Message{}
		json.Unmarshal([]byte(msg.Payload), &rec)
		fmt.Println(url)
		_, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(rec.Body)))
		if err != nil {
			fmt.Println(err)
		}
		c.q.Client.Publish(rec.Id, []byte("done"))
	}
}


