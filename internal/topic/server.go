package topic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Abhishekvrshny/optimus/internal/util"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	core *Core
}

func NewServer(c *Core) *Server {
	return &Server{core:c}
}

func (s *Server) CreateTopic(w http.ResponseWriter, r *http.Request) {
	topic := mux.Vars(r)["topic"]
	body, err := util.ReadBody(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
	}
	req := Topic{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
	}
	req.name = topic
	err = s.core.createTopic(req)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (s *Server) GetTopic(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars)
	w.WriteHeader(http.StatusOK)
}

func (s *Server) Publish(w http.ResponseWriter, r *http.Request) {
	topic := mux.Vars(r)["topic"]
	if !s.core.TopicExists(topic) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("topic doesn't exists"))
	}

	var b bytes.Buffer
	b.ReadFrom(r.Body)
	s.core.Publish(topic, b)
}