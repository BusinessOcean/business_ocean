package becore

import (
	"becore/beconfig"
	"becore/belogger"
	"becore/beserver"

	"go.uber.org/fx"
)

// Module exports dependency to container
var BecoreModule = fx.Options(
	fx.Provide(NewBaseApp),
	fx.Provide(beconfig.NewAppConfig),
	fx.Provide(belogger.NewBeLogger),
	fx.Provide(beserver.NewBeHttpServer),
	fx.Provide(beserver.NewBeGRPCServer),
	fx.Provide(beserver.NewBegoServer),
)
