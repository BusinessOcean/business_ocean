package healthcheck

import (
	"becommon/fxutil"
	"healthcheck/adapters"
	"healthcheck/repository"
	"healthcheck/routes"
	"healthcheck/service"

	"go.uber.org/fx"
)

// type HealthCheckModuleParams struct {
// 	fx.In
// 	Lifecycle fx.Lifecycle
// 	Logger    *belogger.BeLogger

// 	Routes []*beroutes.Route `name:"healthcheckroutes"`

// 	HealthCheck *BeHealthCheckDomain `name:"healthcheckdomain"`
// }

var HealthCheckModules = fx.Options(
	fx.Provide(repository.NewHealthCheckRepository),
	fx.Provide(adapters.NewHealthCheckAdapter),
	fx.Provide(service.NewHealthCheckService),
	fxutil.AnnotatedProvide(routes.NewHealthCheckRoutes, `name:"healthcheckroutes"`),
	fxutil.AnnotatedProvide(NewBeHealthCheckDomain, `group:"domains"`),
	// fx.Invoke(registerHealthCheckLifecycleHooks),
)

// func registerHealthCheckLifecycleHooks(params HealthCheckModuleParams) {
// 	params.Lifecycle.Append(fx.Hook{
// 		OnStart: func(ctx context.Context) error {
// 			params.Logger.Info("Starting HealthCheckModules...")
// 			params.HealthCheck.Setup()
// 			params.HealthCheck.Register(
// 				healthcheck.HealthCheckService_ServiceDesc,
// 				params.HealthCheck.service,
// 			)
// 			params.HealthCheck.RegisterRoutes(
// 				healthcheck.HealthCheckService_ServiceDesc,
// 				params.Routes,
// 			)

// 			// params.HealthCheck.Run()
// 			// go func() { params.HealthCheck.Run() }()
// 			// handle and return error here
// 			return nil
// 		},
// 		OnStop: func(ctx context.Context) error {
// 			params.Logger.Info("Stopping HealthCheckModules...")
// 			return params.HealthCheck.OnTerminate()
// 		},
// 	})
// }
