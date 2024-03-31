package beserver

import (
	"context"

	"google.golang.org/grpc"
)

type BeGrpcServer struct {
	*grpc.Server
}

// NewBeGrpcServer returns a new instance of BeGrpcServer
func NewBeGrpcServer() *BeGrpcServer {
	grpcOpts := GrpcInterceptor()
	grpcServer := grpc.NewServer(grpcOpts)
	return &BeGrpcServer{Server: grpcServer}
}

func GrpcInterceptor() grpc.ServerOption {
	grpcServerOptions := grpc.UnaryInterceptor(
		func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
		) (resp interface{}, err error) {
			if v, ok := req.(validatable); ok {
				err := v.Validate()
				if err != nil {
					return nil, err
				}
			}
			_, _ = handler(ctx, req)
			return handler(ctx, req)
		})
	return grpcServerOptions
}

type validatable interface {
	Validate() error
}
