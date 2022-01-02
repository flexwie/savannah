package router

import (
	"felixwie.com/savannah/api/routes"
	"github.com/gorilla/mux"
)

var router *mux.Router

func init() {
	router = mux.NewRouter()

	router.HandleFunc("/webhook/github/{id}", routes.ReceiveGithubWebhook)
}

func GetRouter() *mux.Router {
	return router
}
