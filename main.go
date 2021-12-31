package main

import (
	"fmt"
	"log"
	"net/http"

	"felixwie.com/savannah/client"
	"felixwie.com/savannah/config"
	q "felixwie.com/savannah/queue"
	"felixwie.com/savannah/router"
)

func main() {
	queue := q.GetQueue()
	queue.Start()
	defer queue.Stop()

	cfg := config.GetConfig()

	r := router.GetRouter()

	if cfg.Ui {
		log.Println("serving ui from './client/out'")
		spa := client.SpaHandler{
			StaticPath: "./client/out",
			IndexPath:  "index.html",
		}

		r.PathPrefix("/ui").Handler(spa)
	}

	log.Printf("api listening on port %d", cfg.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), r))
}
