package beserver

import (
	"becore/beroutes"
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

func (s *BegoServer) RegisterRoutes(group string, routes []*beroutes.Route) error {

	fmt.Println("Registering routes for group: ----->", group)
	fmt.Println("Registering routes for group: ----->", routes)

	party := s.http.Party(group)
	for _, route := range routes {
		// Apply middlewares if present
		if len(route.Middlewares) > 0 {
			party.Use(route.Middlewares...)
		}
		// Register the route
		party.Handle(route.Method, route.Path, route.Handler)
		fmt.Println("Registering route: ---->", route.Path)
	}

	return nil

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
	err := s.http.Run(iris.TLS(":50051", "./cert/server.crt", "./cert/server.key"))
	if err != nil {
		s.http.Logger().Errorf("Failed to run HTTP server: %v", err)
		return err
	}
	return nil
}
