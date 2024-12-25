package auth

import (
	"auth/adapters"
	"auth/repository"
	"auth/routes"
	"auth/service"
	"becommon/fxutil"

	"go.uber.org/fx"
)

// type HealthCheckModuleParams struct {
// 	fx.In
// 	Lifecycle fx.Lifecycle
// 	Logger    *belogger.BeLogger

// 	AuthRoutes []*beroutes.Route    `name:"authroutes"`
// 	AuthDomain []bedomain.IBeDomain `group:"domains"`
// }

var AuthModules = fx.Options(
	fx.Provide(repository.NewAuthRepository),
	fx.Provide(adapters.NewAuthAdapter),
	fx.Provide(service.NewAuthService),
	fxutil.AnnotatedProvide(routes.NewAuthRoutes, `name:"authroutes"`),
	fxutil.AnnotatedProvide(NewBeAuthDomain, `group:"domains"`),
	// fx.Invoke(registerAuthLifecycleHooks),
)

// func registerAuthLifecycleHooks(params HealthCheckModuleParams) {
// 	params.Lifecycle.Append(fx.Hook{
// 		OnStart: func(ctx context.Context) error {
// 			params.Logger.Info("Starting HealthCheckModules...")
// 			// Example: Process the collected domains
// 			// for _, domain := range *&params.AuthDomain {
// 			// 	params.Logger.Info("Setting up domain:", domain)
// 			// 	// domain.Setup()
// 			// }
// 			// params.AuthDomain.Setup()
// 			// params.AuthDomain.Register(
// 			// 	authApis.AuthService_ServiceDesc,
// 			// 	params.AuthDomain.service,
// 			// )
// 			// params.AuthDomain.RegisterRoutes(
// 			// 	authApis.AuthService_ServiceDesc,
// 			// 	params.AuthRoutes,
// 			// )
// 			// params.AuthDomain.Run()
// 			// go func() { params.AuthDomain.Run() }()
// 			// handle and return error here
// 			return nil
// 		},
// 		OnStop: func(ctx context.Context) error {
// 			params.Logger.Info("Stopping HealthCheckModules...")
// 			return nil
// 		},
// 	})
// }
