package main

import (
	"becommon"
	"becommon/bectx"
	"becore"
	"bedatabase"
	"businessocean/bego"
	"healthcheck"

	"go.uber.org/fx"
)

var AppModule = fx.Options(
	fx.Provide(bectx.NewBeCtx),
	becommon.BeCommonModule,
	becore.BecoreModule,
	bedatabase.DatabaseModule,
	healthcheck.HealthCheckModules,
	fx.Provide(
		fx.Annotate(
			bego.NewBegoDomains,
			fx.ParamTags(`group:"bego"`),
		),
	),
)
