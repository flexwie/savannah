package router

import (
	"felixwie.com/savannah/router/routes"
	"github.com/gorilla/mux"
)

var router *mux.Router

func init() {
	router = mux.NewRouter()

	router.HandleFunc("/webhook/{id}", routes.ReceiveWebhook)
}

func GetRouter() *mux.Router {
	return router
}
