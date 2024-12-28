package healthcheck

import (
	"becommon/fxutil"
	"becore/belogger"
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
	Lifecycle   fx.Lifecycle
	Logger      *belogger.BeLogger
	HealthCheck *BeHealthCheckDomain
}

var HealthCheckModules = fx.Options(
	fx.Provide(repository.NewHealthCheckRepository),
	fx.Provide(adapters.NewHealthCheckAdapter),
	fx.Provide(service.NewHealthCheckService),
	fxutil.AnnotatedProvide(routes.NewHealthCheckRoutes, `name:"healthcheckroutes"`),
	fx.Provide(NewBeHealthCheckDomain),
)

func RegisterHealthLifecycleHooks(params HealthCheckModuleParams) {
	params.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			params.Logger.Info("Starting HealthCheckModules...")
			params.HealthCheck.Setup()

			fmt.Printf("Service Name : %v", healthcheck.HealthCheckService_ServiceDesc)
			fmt.Printf("Service Service : %v", params.HealthCheck.service)

			params.HealthCheck.Register(
				healthcheck.HealthCheckService_ServiceDesc,
				params.HealthCheck.service,
			)
			params.HealthCheck.RegisterRoutes(
				healthcheck.HealthCheckService_ServiceDesc,
				params.HealthCheck.routes,
			)
			go func() {
				params.HealthCheck.Run(":5002")
			}()
			fmt.Println("HealthCheckModules is running")
			// handle and return error here
			return nil
		},
		OnStop: func(ctx context.Context) error {
			params.Logger.Info("Stopping HealthCheckModules...")
			return params.HealthCheck.OnTerminate()
		},
	})
}
