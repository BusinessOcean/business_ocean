package healthcheck

import (
	"becommon/bedomain"
	"becore/beroutes"
	"becore/beserver"
	"beservice/healthcheck/apis"
	"healthcheck/service"
	"log"
	"net"

	"github.com/kataras/iris/v12"
)

var _ bedomain.IBeDomainModule = (*BeHealthCheckDomain)(nil)

type BeHealthCheckDomain struct {
	*bedomain.BaseDomainModule
	service *service.HealthCheckService
	routes  []beroutes.IRegisterRouteAPI
}

func NewBeHealthCheckDomain(
	baseDomainModule *bedomain.BaseDomainModule,
	service *service.HealthCheckService,
	routes []beroutes.IRegisterRouteAPI,
) bedomain.IBeDomainModule {
	return &BeHealthCheckDomain{
		baseDomainModule,
		service,
		routes,
	}
}

func (app *BeHealthCheckDomain) Run() error {

	go app.RunGRPCServer(app.GetConfig(), app.GetGrpcServer())
	go app.RunHTTPServer(app.GetConfig(), app.GetHTTPServer())

	return nil
}

func (app *BeHealthCheckDomain) RunGRPCServer(bedomain.BeDomainConfig, *beserver.BeGrpcServer) error {
	app.GetLogger().Errorf("Run GRPCServer is not implemented")

	beroutes.NewRegisterRouteAPI(app.GetHTTPServer(), app.routes)

	app.GetHTTPServer().Run(iris.Addr(":8080"))
	return nil
}

func (app *BeHealthCheckDomain) RunHTTPServer(bedomain.BeDomainConfig, *beserver.BeHTTPServer) error {

	app.GetLogger().Errorf("Run HTTPServer is not implemented")

	listener, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		log.Fatalln(err)
	}
	apis.RegisterHealthCheckServiceServer(app.GetGrpcServer(), app.service)
	// Use the service

	err = app.GetGrpcServer().Serve(listener)

	if err != nil {
		log.Fatalln(err)
	}
	return nil
}
