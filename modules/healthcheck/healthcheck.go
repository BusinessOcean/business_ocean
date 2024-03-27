package healthcheck

import (
	"beservice/healthcheck/apis"
	"context"
	"fmt"
	"healthcheck/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

// ExampleService starts a gRPC server.

type HealthAPI struct {
	service *service.HealthCheckService
}

// new healthcheck service
func NewHealthCheckService(service *service.HealthCheckService) *HealthAPI {
	return &HealthAPI{service: service}
}

func (e HealthAPI) RunHealthCheckService() {
	listener, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		log.Fatalln(err)
	}

	grpcOpts := GrpcInterceptor()
	grpcServer := grpc.NewServer(grpcOpts)
	apis.RegisterHealthCheckServiceServer(grpcServer, e.service)
	// Use the service

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalln(err)
	}

}
func (e HealthAPI) StopHealthCheckService() {
	fmt.Println("Stopping Example Service application...")
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
