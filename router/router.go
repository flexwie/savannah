package router

import (
	"felixwie.com/savannah/router/routes"
	"github.com/gorilla/mux"
)

var router *mux.Router

func init() {
	router = mux.NewRouter()

	router.HandleFunc("/webhook/github/{id}", routes.ReceiveGithubWebhook)
	router.HandleFunc("/api/sync", routes.SyncRepository)
	router.HandleFunc("/api/repositories", routes.GetRepositories)
}

func GetRouter() *mux.Router {
	return router
}
