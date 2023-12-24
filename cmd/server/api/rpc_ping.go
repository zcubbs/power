package api

import (
	"context"
	pb "github.com/zcubbs/power/proto/gen/v1"
)

func (s *Server) Ping(_ context.Context, _ *pb.Empty) (*pb.PingResponse, error) {
	return &pb.PingResponse{
		Message:   "Pong",
		Version:   s.cfg.Version,
		Commit:    s.cfg.Commit,
		BuildTime: s.cfg.Date,
	}, nil
}
