package main

import (
	"auth"
	"becommon"
	"becommon/bectx"
	"becommon/fxutil"
	"becore"
	"bedatabase"
	"businessocean/bego"
	"healthcheck"

	"go.uber.org/fx"
)

var AppModule = fx.Options(
	fxutil.AnnotatedProvide(bectx.NewBeCtx, `name:"bectx"`),
	fxutil.AnnotatedProvide(bectx.NewBeAppCtx, `name:"beappctx"`),
	becore.BecoreModule,
	becommon.BeCommonModule,
	bedatabase.DatabaseModules,
	fx.Provide(bego.NewBegoApp),
	fx.Provide(bego.NewBegoCommand),
	healthcheck.HealthCheckModules,
	auth.AuthModules,
)
