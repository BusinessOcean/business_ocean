package main

import (
	"auth"
	"becommon"
	"becommon/bectx"
	"becommon/bedomain"
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
	fx.Provide(bego.NewBegoApp),
	fx.Provide(bego.NewBegoCommand),
	becommon.BeCommonModule,
	bedatabase.DatabaseModules,
	fx.Provide(bedomain.NewBaseDomain),
	auth.AuthModules,
	healthcheck.HealthCheckModules,
	// fx.Provide(bego.NewBeAppServiceDomain),
)
