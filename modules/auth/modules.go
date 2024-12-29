package auth

import (
	"auth/adapters"
	"auth/repository"
	"auth/routes"
	"auth/service"
	"becommon/fxutil"
	"becore/belogger"
	"context"
	"fmt"

	auth "beservice/auth/apis"

	"go.uber.org/fx"
)

type AuthModuleParams struct {
	fx.In
	Lifecycle fx.Lifecycle
	Logger    *belogger.BeLogger
	Auth      *BeAuthkDomain
}

var AuthModules = fx.Options(
	fxutil.AnnotatedProvide(routes.NewAuthRoutes, `name:"authroutes"`),
	fx.Provide(repository.NewHealthCheckRepository),
	fx.Provide(adapters.NewAuthAdapter),
	fx.Provide(service.NewAuthService),
	fx.Provide(NewBeAuthDomain),
	fx.Invoke(registerAuthLifecycleHooks),
)

func registerAuthLifecycleHooks(params AuthModuleParams) {
	params.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			params.Logger.Info("Starting AuthModule...")
			params.Auth.Setup()

			fmt.Printf("Service Name : %v", auth.AuthService_ServiceDesc)
			fmt.Printf("Service Service : %v", params.Auth.service)

			params.Auth.Register(
				auth.AuthService_ServiceDesc,
				params.Auth.service,
			)
			params.Auth.RegisterRoutes(
				auth.AuthService_ServiceDesc,
				params.Auth.routes,
			)
			go func() {
				params.Auth.Run(":5003")
			}()

			params.Logger.Println("AuthModule is running")

			// handle and return error here
			return nil
		},
		OnStop: func(ctx context.Context) error {
			params.Logger.Info("Stopping AuthModule...")
			return params.Auth.OnTerminate()
		},
	})
}
