package transport

import (
	"context"
	pb "github.com/jumagaliev1/hash-generator/hasher/proto"
)

func (s *Server) Get(ctx context.Context, in *pb.GetHashRequest) (*pb.GetHashResponse, error) {
	hash, err := s.service.Hash.Get(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.GetHashResponse{Hash: hash}, nil
}
