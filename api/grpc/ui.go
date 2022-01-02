package grpc

import (
	"context"
	"log"

	"felixwie.com/savannah/config"
	pb "felixwie.com/savannah/internal/proto/api"
	"felixwie.com/savannah/queue"
)

type Server struct {
	pb.UnimplementedProjectsServer
}

var Cfg config.Config

func init() {
	Cfg = config.GetConfig()
}

func (s *Server) GetAll(ctx context.Context, in *pb.Empty) (*pb.GetAllMessage, error) {
	var data []*pb.Project
	for _, s := range Cfg.Source {
		log.Printf("adding %s to response", s.Name)
		data = append(data, &pb.Project{Name: s.Name, Repository: s.URL})
	}

	return &pb.GetAllMessage{Projects: data}, nil
}

func (s *Server) Sync(ctx context.Context, in *pb.SyncMessage) (*pb.Empty, error) {
	queue := queue.GetQueue()
	queue.Submit(nil)

	return &pb.Empty{}, nil
}
