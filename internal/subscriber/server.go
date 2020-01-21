package subscriber

import (
	"encoding/json"
	"github.com/Abhishekvrshny/optimus/internal/topic"
	http2 "github.com/Abhishekvrshny/optimus/pkg/http"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	topicCore *topic.Core
	subscriberCore *Core
}

func NewServer(topicCore *topic.Core, subscriberCore *Core) *Server{
	return &Server{topicCore:topicCore,subscriberCore:subscriberCore}
}

func (s *Server) CreateSubscriber(w http.ResponseWriter, r *http.Request) {
	topic := mux.Vars(r)["topic"]
	subs := mux.Vars(r)["name"]

	if !s.topicCore.TopicExists(topic) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("topic doesn't exists"))
		return
	}

	if !s.topicCore.SubscriberExists(topic, subs) {
		body, err := http2.ReadBody(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		req := Subscriber{}
		err = json.Unmarshal(body, &req)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		req.topic = topic
		req.name = subs
		log.Printf("received request to create subscriber %s on topic %s\n",subs,topic)

		s.subscriberCore.CreateSubscriber(req)
	}
}
