package healthcheck

import (
	"becommon/fxutil"
	"becore/belogger"
	"becore/beroutes"
	healthcheck "beservice/healthcheck/apis"
	"context"
	"fmt"
	"healthcheck/adapters"
	"healthcheck/repository"
	"healthcheck/routes"
	"healthcheck/service"

	"go.uber.org/fx"
)

type HealthCheckModuleParams struct {
	fx.In
	Lifecycle fx.Lifecycle
	Logger    *belogger.BeLogger

	Routes []*beroutes.Route `name:"healthcheckroutes"`

	HealthCheck *BeHealthCheckDomain `name:"healthcheck"`
}

var HealthCheckModules = fx.Options(
	fx.Provide(repository.NewHealthCheckRepository),
	fx.Provide(adapters.NewHealthCheckAdapter),
	fx.Provide(service.NewHealthCheckService),
	fxutil.AnnotatedProvide(routes.NewHealthCheckRoutes, `name:"healthcheckroutes"`),
	fxutil.AnnotatedProvide(NewBeHealthCheckDomain, `name:"healthcheck"`),
)

func RegisterHealthCheckLifecycleHooks(params HealthCheckModuleParams) {
	params.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			params.Logger.Info("Starting HealthCheckModules...")
			params.HealthCheck.Setup()
			params.HealthCheck.Register(
				healthcheck.HealthCheckService_ServiceDesc,
				params.HealthCheck.service,
			)
			params.HealthCheck.RegisterRoutes(
				healthcheck.HealthCheckService_ServiceDesc,
				params.Routes,
			)

			fmt.Println("### MODULE HealthCheck Run")
			params.HealthCheck.Run()
			// go func() { params.HealthCheck.Run() }()
			// handle and return error here
			return nil
		},
		OnStop: func(ctx context.Context) error {
			params.Logger.Info("Stopping HealthCheckModules...")
			return params.HealthCheck.OnTerminate()
		},
	})
}
