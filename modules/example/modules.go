package example

import (
	"example/adapters"
	"example/repository"
	"example/service"

	"go.uber.org/fx"
)

// Service is a function that starts a gRPC server.
var ExampleModule = fx.Options(
	fx.Provide(adapters.NewExampleAdapter),
	fx.Provide(NewExampleService),
	fx.Provide(service.NewTodoService),
	fx.Provide(repository.NewTodoRepository),
	fx.Invoke(func(example *Example) {
		go func() {
			example.RunExampleService()
		}()
	}),
)
