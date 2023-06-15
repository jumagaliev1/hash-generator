package transport

import (
	pb "github.com/jumagaliev1/hash-generator/hasher/proto"
	"github.com/jumagaliev1/hash-generator/internal/service"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

var addr = ":50051"

type Server struct {
	pb.HashServiceServer
	service service.Service
}

func NewServer(service service.Service) *Server {
	return &Server{service: service}
}

func Run(service *service.Service) error {
	log.Info("Starting server...")
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	pb.RegisterHashServiceServer(s, &Server{service: *service})
	reflection.Register(s)

	return s.Serve(lis)
}
