package auth

import (
	"auth/adapters"
	"auth/repository"
	"auth/routes"
	"auth/service"
	"becommon/fxutil"
	"becore/belogger"
	"becore/beroutes"
	authApis "beservice/auth/apis"
	"context"
	"fmt"

	"go.uber.org/fx"
)

type HealthCheckModuleParams struct {
	fx.In
	Lifecycle fx.Lifecycle
	Logger    *belogger.BeLogger

	AuthRoutes []*beroutes.Route `name:"authroutes"`
	AuthDomain *BeAuthDomain     `name:"authdomain"`
}

var AuthModules = fx.Options(
	fx.Provide(repository.NewAuthRepository),
	fx.Provide(adapters.NewAuthAdapter),
	fx.Provide(service.NewAuthService),
	fxutil.AnnotatedProvide(routes.NewAuthRoutes, `name:"authroutes"`),
	fxutil.AnnotatedProvide(NewBeAuthDomain, `name:"authdomain"`),
)

func RegisterAuthLifecycleHooks(params HealthCheckModuleParams) {
	params.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			params.Logger.Info("Starting HealthCheckModules...")
			params.AuthDomain.Setup()
			params.AuthDomain.Register(
				authApis.AuthService_ServiceDesc,
				params.AuthDomain.service,
			)
			params.AuthDomain.RegisterRoutes(
				authApis.AuthService_ServiceDesc,
				params.AuthRoutes,
			)
			fmt.Println("### MODULE AuthModules Run")
			params.AuthDomain.Run()
			// go func() { params.AuthDomain.Run() }()
			// handle and return error here
			return nil
		},
		OnStop: func(ctx context.Context) error {
			params.Logger.Info("Stopping HealthCheckModules...")
			return params.AuthDomain.OnTerminate()
		},
	})
}
