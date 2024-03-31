package healthcheck

import (
	"becommon/bedomain"
	"healthcheck/adapters"
	"healthcheck/repository"
	"healthcheck/routes"
	"healthcheck/service"

	"go.uber.org/fx"
)

// const HEALTH_ROUTES_TAG = `health_routes`

// Service is a function that starts a gRPC server.
var HealthCheckModules = fx.Options(
	fx.Provide(repository.NewHealthCheckRepository),
	fx.Provide(adapters.NewHealthCheckAdapter),
	fx.Provide(service.NewHealthCheckService),
	fx.Provide(routes.NewHealthCheckRoutes),
	fx.Provide(
		fx.Annotate(
			NewBeHealthCheckDomain,
			fx.As(new(bedomain.IBeDomainModule)),
			fx.ResultTags(`group:"bego"`),
		),
	),
)
