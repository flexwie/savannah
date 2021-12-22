package routes

import (
	"log"
	"net/http"

	q "felixwie.com/savannah/queue"
	"github.com/gorilla/mux"
)

func ReceiveWebhook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)

	queue := q.GetQueue()
	queue.Submit(&q.WebhookJob{ID: vars["id"]})

	log.Printf("ID: %#v", vars)
}
