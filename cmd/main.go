package main

import (
	"github.com/Abhishekvrshny/optimus/internal/topic"
	"github.com/Abhishekvrshny/optimus/pkg/redis"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	q := redis.NewRedisQueue()
	topicCore := topic.NewCore(q)
	topicServer := topic.NewServer(topicCore)

	router := mux.NewRouter()
	router.HandleFunc("/topics/{topic}", topicServer.CreateTopic).Methods("POST")
	router.HandleFunc("/topics/{topic}", topicServer.GetTopic).Methods("GET")
	router.HandleFunc("/topics/{topic}/publish", topicServer.Publish).Methods("POST")


	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8008",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
