package example

import (
	"bedatabase"
	"beservice/example/todo"
	"context"
	"example/repository"
	"example/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

// ExampleService starts a gRPC server.

type Example struct {
	database bedatabase.BeDatabase
}

// new example service
func NewExampleService(database bedatabase.BeDatabase) *Example {
	return &Example{database: database}
}

func (e Example) RunExampleService() {
	listener, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		log.Fatalln(err)
	}
	repo := repository.NewTodoRepository()

	// Use the repository
	todoService := service.NewTodoService(repo, e.database)

	grpcOpts := GrpcInterceptor()
	grpcServer := grpc.NewServer(grpcOpts)
	todo.RegisterTodoServiceServer(grpcServer, todoService)
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
