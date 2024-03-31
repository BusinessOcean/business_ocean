package becore

import (
	"becore/beconfig"
	"becore/belogger"
	"becore/beserver"
	"bedatabase/memory"

	"go.uber.org/fx"
)

// Module exports dependency to container
var BecoreModule = fx.Options(
	fx.Provide(belogger.NewBeLogger),
	fx.Provide(beconfig.NewConfig),
	fx.Provide(beserver.NewBeHttpServer),
	fx.Provide(beserver.NewBeGrpcServer),
	fx.Provide(memory.NewInMemoryDatabase),
)
