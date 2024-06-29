package healthcheck

import (
	"becommon/fxutil"
	"becore/belogger"
	"beservice/healthcheck/apis"
	"context"
	"healthcheck/adapters"
	"healthcheck/repository"
	"healthcheck/routes"
	"healthcheck/service"

	"go.uber.org/fx"
)

type HealthCheckDomainParams struct {
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
	fx.Invoke(func(domain HealthCheckDomainParams) {
		// Lifecycle hooks
		domain.Lifecycle.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				domain.Logger.Info("HealthCheckModules started")
				domain.HealthCheck.Setup()
				domain.HealthCheck.Register(apis.HealthCheckService_ServiceDesc, domain.HealthCheck.service)
				domain.HealthCheck.Run()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				domain.Logger.Info("HealthCheckModules stopped")
				return domain.HealthCheck.OnTerminate()

			},
		})
	}),
)
