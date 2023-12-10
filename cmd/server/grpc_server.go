package main

import (
	"context"
	pb "github.com/zcubbs/power/proto/gen"
)

type server struct {
	pb.UnimplementedBlueprintServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) GenerateProject(ctx context.Context, in *pb.GenerateRequest) (*pb.GenerateResponse, error) {
	// Implement your logic here
	// For example: Generate a project blueprint and return a URL or status

	return &pb.GenerateResponse{ /* fields */ }, nil
}
