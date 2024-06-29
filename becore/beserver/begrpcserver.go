package beserver

import (
	"google.golang.org/grpc"
)

type BeGRPCServer struct {
	*grpc.Server
}

func NewBeGRPCServer(options ...grpc.ServerOption) *BeGRPCServer {
	grpcServer := grpc.NewServer(options...)
	return &BeGRPCServer{grpcServer}
}
