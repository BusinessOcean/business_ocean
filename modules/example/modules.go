package example

import (
	"go.uber.org/fx"
)

// Service is a function that starts a gRPC server.
var ExampleModule = fx.Option(
	fx.Provide(NewExampleService),
)
