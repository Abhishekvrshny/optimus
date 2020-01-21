package topic

import (
	"encoding/json"
	"fmt"
	"github.com/Abhishekvrshny/optimus/internal/util"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	core *core
}

func NewServer(c *core) *Server {
	return &Server{core:c}
}

func (s *Server) CreateTopic(w http.ResponseWriter, r *http.Request) {
	topic := mux.Vars(r)["topic"]
	body, err := util.ReadBody(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
	}
	req := CreateTopicRequest{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
	}
	fmt.Println(req)
	s.core.createTopic(topic)
	w.WriteHeader(http.StatusOK)
}

func (s *Server) GetTopic(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars)
	w.WriteHeader(http.StatusOK)
}

func (s *Server) Publish(w http.ResponseWriter, r *http.Request) {
	fmt.Println("here")
}