package becore

import (
	"becore/beconfig"
	"becore/belogger"
	"becore/beserver"
	"becore/infrastructure"

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
	fx.Provide(infrastructure.NewFirebaseApp),
	fx.Provide(infrastructure.NewFirebaseAuth),
	fx.Provide(infrastructure.NewFirebaseMessagingClient),
	fx.Provide(infrastructure.NewFirestoreClient),
)
