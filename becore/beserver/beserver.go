package beserver

import (
	"betypes/beerrors"
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"google.golang.org/grpc"
)

type BegoServer struct {
	http *BeHTTPServer
	grpc *BeGRPCServer
}

func NewBegoServer(
	http *BeHTTPServer,
	grpc *BeGRPCServer,
) *BegoServer {

	server := BegoServer{
		http: http,
		grpc: grpc,
	}
	return &server
}

func (s *BegoServer) Register(desc grpc.ServiceDesc, handler interface{}) error {
	// return error if interface is nil
	if handler == nil {
		return beerrors.ErrRegisterNilHandler
	}

	fmt.Println("Registering service: ", desc.ServiceName)
	fmt.Println("Handler: ", handler)

	s.grpc.RegisterService(&desc, handler)
	appServer := mvc.New(s.http)

	appServer.Handle(handler, mvc.GRPC{
		Server:      s.grpc,
		ServiceName: desc.ServiceName,
		Strict:      false,
	})

	return nil
}

func (s *BegoServer) GetHttpServer() *BeHTTPServer {
	return s.http
}

func (s *BegoServer) GetGrpcServer() *BeGRPCServer {
	return s.grpc
}
func (s *BegoServer) RunServer() error {

	// run http server
	// s.http.Run(iris.TLS(":2525", "./cert/server.crt", "./cert/server.key"))
	err := s.http.Run(iris.TLS(":50051", "./cert/server.crt", "./cert/server.key"))
	if err != nil {
		s.http.Logger().Errorf("Failed to run HTTP server: %v", err)
		return err
	}
	s.http.Logger().Infof("HTTP server is running on port 50051")
	return nil
}
