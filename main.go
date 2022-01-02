package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	api "felixwie.com/savannah/api/grpc"
	"felixwie.com/savannah/client"
	"felixwie.com/savannah/config"
	q "felixwie.com/savannah/queue"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"

	pb "felixwie.com/savannah/internal/proto/api"
)

func grpcHandler(server *grpc.Server, other http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			server.ServeHTTP(rw, r)
		} else {
			other.ServeHTTP(rw, r)
		}
	}), &http2.Server{})
}

func ServeGrpcWithGateway[T any](
	s T,
	registerServer func(grpc.ServiceRegistrar, T),
	registerHandlerFromEndpoint func(context.Context, *runtime.ServeMux, string, []grpc.DialOption,
	) error) http.Handler {
	cfg := config.GetConfig()
	ctx := context.Background()
	mux := runtime.NewServeMux()
	addr := fmt.Sprintf("localhost:%d", cfg.Port)

	srv := grpc.NewServer()

	registerServer(srv, s)
	registerHandlerFromEndpoint(ctx, mux, addr, []grpc.DialOption{grpc.WithInsecure()})

	return grpcHandler(srv, mux)
}

func main() {
	// configure queue
	queue := q.GetQueue()
	queue.Start()
	defer queue.Stop()

	// get config
	cfg := config.GetConfig()

	// get routers
	router := mux.NewRouter()
	grpcRouter := ServeGrpcWithGateway[pb.ProjectsServer](&api.Server{}, pb.RegisterProjectsServer, pb.RegisterProjectsHandlerFromEndpoint)
	router.PathPrefix("/api").Handler(grpcRouter)
	router.PathPrefix("/test.Test").Handler(grpcRouter)

	if cfg.Ui {
		log.Println("serving ui from \"./client/out\"")
		spa := client.SpaHandler{
			StaticPath: "./client/out",
			IndexPath:  "index.html",
		}

		router.PathPrefix("/ui").Handler(spa)
		router.Path("/").HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			http.Redirect(rw, r, "/ui", http.StatusMovedPermanently)
		})
	}

	wrapper := h2c.NewHandler(router, &http2.Server{})

	log.Printf("api listening on port %d", cfg.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), wrapper))
}
