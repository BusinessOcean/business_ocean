package healthcheck

import (
	"becommon/fxutil"
	"becore/belogger"
	healthcheck "beservice/healthcheck/apis"
	"context"
	"healthcheck/adapters"
	"healthcheck/repository"
	"healthcheck/routes"
	"healthcheck/service"

	"go.uber.org/fx"
)

type HealthCheckModuleParams struct {
	fx.In
	Lifecycle   fx.Lifecycle
	Logger      *belogger.BeLogger
	HealthCheck *BeHealthCheckDomain `name:"healthcheck"`
}

var HealthCheckModules = fx.Options(
	fx.Provide(repository.NewHealthCheckRepository),
	fx.Provide(adapters.NewHealthCheckAdapter),
	fx.Provide(service.NewHealthCheckService),
	fx.Provide(routes.NewHealthCheckRoutes),
	fxutil.AnnotatedProvide(NewBeHealthCheckDomain, `name:"healthcheck"`),
	fx.Invoke(registerLifecycleHooks),
)

func registerLifecycleHooks(params HealthCheckModuleParams) {
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
				params.HealthCheck.routes,
			)
			params.HealthCheck.Run()
			// handle and return error here
			return nil
		},
		OnStop: func(ctx context.Context) error {
			params.Logger.Info("Stopping HealthCheckModules...")
			return params.HealthCheck.OnTerminate()
		},
	})
}
