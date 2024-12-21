package auth

import (
	"auth/adapters"
	"auth/repository"
	"auth/routes"
	"auth/service"
	"becommon/fxutil"
	"becore/belogger"
	"context"

	authApi "beservice/auth/apis"

	"go.uber.org/fx"
)

type AuthModuleParams struct {
	fx.In
	Lifecycle fx.Lifecycle
	Logger    *belogger.BeLogger
	Auth      *BeAuthkDomain `name:"auth"`
}

var AuthModules = fx.Options(
	fx.Provide(repository.NewHealthCheckRepository),
	fx.Provide(adapters.NewAuthAdapter),
	fx.Provide(service.NewAuthService),
	fx.Provide(routes.NewAuthRoutes),
	fxutil.AnnotatedProvide(NewBeAuthDomain, `name:"auth"`),
	fx.Invoke(registerLifecycleHooks),
)

func registerLifecycleHooks(params AuthModuleParams) {
	params.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			params.Logger.Info("Starting AuthModule...")
			params.Auth.Setup()
			params.Auth.Register(
				authApi.AuthService_ServiceDesc,
				params.Auth.service,
			)
			params.Auth.RegisterRoutes(
				authApi.AuthService_ServiceDesc,
				params.Auth.routes,
			)
			params.Auth.Run()
			// handle and return error here
			return nil
		},
		OnStop: func(ctx context.Context) error {
			params.Logger.Info("Stopping AuthModule...")
			return params.Auth.OnTerminate()
		},
	})
}
