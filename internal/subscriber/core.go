package subscriber

import (
	"bytes"
	"encoding/json"
	"github.com/Abhishekvrshny/optimus/internal/message"
	"github.com/Abhishekvrshny/optimus/internal/topic"
	"github.com/Abhishekvrshny/optimus/pkg/redis"
	redis2 "github.com/go-redis/redis"
	"log"
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
	go c.readSubs( req.name,req.Url, pubsub.Channel())
}

func (c *Core) readSubs(name string,url string, channel <-chan *redis2.Message) {
	for msg := range channel {
		rec := message.Message{}
		json.Unmarshal([]byte(msg.Payload), &rec)
		res, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(rec.Body)))
		if err != nil {
			continue
		}
		log.Printf("received response for subsriber: %s code: %d message id: %s", name, res.StatusCode, rec.Id)
		c.q.Client.Publish(rec.Id, []byte("done"))
	}
}


