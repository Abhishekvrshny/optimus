package subscriber

import (
	"bytes"
	"fmt"
	"github.com/Abhishekvrshny/optimus/pkg/redis"
	redis2 "github.com/go-redis/redis"
	"net/http"
)

type Core struct {
	q *redis.Redis
}

func NewCore(q *redis.Redis) *Core{
	return &Core{q}
}

func (c *Core) CreateSubscriber(req Subscriber) {
	pubsub := c.q.Subscribe(req.topic)
	go c.readSubs(req.url, pubsub.Channel())
}

func (c *Core) readSubs(url string, channel <-chan *redis2.Message) {
	for msg := range channel {
		fmt.Println("subscription read")
		res, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(msg.Payload)))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res)
		fmt.Println(msg.Channel, msg.Payload)
	}
}


