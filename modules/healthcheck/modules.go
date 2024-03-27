package healthcheck

import (
	"healthcheck/adapters"
	"healthcheck/repository"
	"healthcheck/service"

	"go.uber.org/fx"
)

// Service is a function that starts a gRPC server.
var HealthCheckModules = fx.Options(
	fx.Provide(adapters.NewHealthCheckAdapter),
	fx.Provide(NewHealthCheckService),
	fx.Provide(service.NewHealthCheckService),
	fx.Provide(repository.NewHealthCheckRepository),
	fx.Invoke(func(service *HealthAPI) {
		// go RunHealthCheckAPI(lifecycle, logger, service)
		go func() {
			service.RunHealthCheckService() // Start gRPC server
		}()
	}),
)

// // Define the function to be invoked
// func RunHealthCheckAPI(lifecycle fx.Lifecycle, logger *belogger.BeLogger, service *HealthAPI) {

// 	lifecycle.Append(fx.Hook{
// 		OnStart: func(ctx context.Context) error {
// 			logger.Println("Starting application...")
// 			service.RunHealthCheckService() // Start gRPC server
// 			// err :=
// 			// if err != nil {
// 			// 	logger.Printf("Failed to start gRPC server: %v", err)
// 			// 	return err
// 			// }
// 			return nil
// 		},
// 		OnStop: func(ctx context.Context) error {
// 			logger.Println("Stopping application...")
// 			service.StopHealthCheckService() // Stop gRPC server gracefully
// 			// err :=
// 			// if err != nil {
// 			// 	logger.Printf("Failed to stop gRPC server: %v", err)
// 			// 	return err
// 			// }
// 			return nil
// 		},
// 	})

// }
