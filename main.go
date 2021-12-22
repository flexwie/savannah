package main

import (
	"log"
	"net/http"
	"time"

	"felixwie.com/savannah/config"
	q "felixwie.com/savannah/queue"
	"felixwie.com/savannah/router"
)

func main() {
	config := config.GetConfig()
	log.Printf("config: %#v", config)

	queue := q.GetQueue()
	queue.Start()
	defer queue.Stop()

	r := router.GetRouter()
	srv := &http.Server{
		Handler:      r,
		WriteTimeout: 2 * time.Second,
		ReadTimeout:  2 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
