package main

import (
	"becommon/bectx"
	"becore"
	"bedatabase"
	"healthcheck"

	"go.uber.org/fx"
)

var AppModule = fx.Options(
	fx.Provide(bectx.NewBeCtx),
	becore.BecoreModule,
	healthcheck.HealthCheckModules,
	bedatabase.DatabaseModule,
)
