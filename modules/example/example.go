package example

import (
	"beservice/example/todo"
	"context"
	"example/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

// ExampleService starts a gRPC server.

type Example struct {
	service *service.TodoService
}

// new example service
func NewExampleService(todoService *service.TodoService) *Example {
	return &Example{service: todoService}
}

func (e Example) RunExampleService() {
	listener, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		log.Fatalln(err)
	}

	grpcOpts := GrpcInterceptor()
	grpcServer := grpc.NewServer(grpcOpts)
	todo.RegisterTodoServiceServer(grpcServer, e.service)
	// Use the service

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalln(err)
	}

}

func GrpcInterceptor() grpc.ServerOption {
	grpcServerOptions := grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
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
