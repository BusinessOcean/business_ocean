package healthcheck

import (
	"becommon/bedomain"
	"becore/beroutes"
	"healthcheck/service"

	"go.uber.org/fx"
)

var _ bedomain.IBeDomain = (*BeHealthCheckDomain)(nil)

type DomainParams struct {
	fx.In

	BaseModules *bedomain.BaseDomain `name:"domain"`
	Service     *service.HealthCheckService
	Routes      []beroutes.IRegisterRouteAPI
}

type BeHealthCheckDomain struct {
	*bedomain.BaseDomain
	service *service.HealthCheckService
	routes  []beroutes.IRegisterRouteAPI
}

func NewBeHealthCheckDomain(params DomainParams) *BeHealthCheckDomain {
	return &BeHealthCheckDomain{
		BaseDomain: params.BaseModules,
		service:    params.Service,
		routes:     params.Routes,
	}
}

// func (d *BeHealthCheckDomain) Register(interface{}) error {
// 	return nil
// }

// func (d *BeHealthCheckDomain) Setup() error {
// 	// d.GetServer().Register("healthcheck.apis.HealthCheckService", d.service)
// 	// apis.HealthCheckAPI(d.service)
// 	// fmt.Println("HealthCheckDomain setups")
// 	// apis.RegisterHealthCheckServiceServer(d.GetServer().GetGrpcServer(), d.service)

// 	return nil
// }

// func (d *BeHealthCheckDomain) Run() error {
// 	app := iris.New()
// 	app.Logger().SetLevel("debug")

// 	app.Get("/", func(ctx iris.Context) {
// 		ctx.HTML("<h1>Index Page</h1>")
// 	})

// 	// Register gRPC server.
// 	// grpcServer := d.GetServer().GetGrpcServer()
// 	// apis.RegisterHealthCheckServiceServer(grpcServer, d.service)

// 	serviceName := apis.File_healthcheck_apis_health_service_proto.Services().Get(0).FullName()
// 	fmt.Println("Service Name: --->", serviceName)

// 	// Register MVC application controller for gRPC services.
// 	// You can bind as many mvc gRpc services in the same Party or app,
// 	// as the ServiceName differs.
// 	mvc.New(app).
// 		Register(d.service).
// 		Handle(d.service, mvc.GRPC{
// 			Server:      d.GetServer().GetGrpcServer(),
// 			ServiceName: string(serviceName),
// 			Strict:      false,
// 		})

// 	err := app.Run(iris.TLS(":50051", "./cert/server.crt", "./cert/server.key"))
// 	if err != nil {
// 		app.Logger().Errorf("Failed to run HTTP server: %v", err)
// 		return err
// 	}
// 	return nil
// }
