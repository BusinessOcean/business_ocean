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
	becore.BecoreModule,
	fxutil.AnnotatedProvide(bectx.NewBeCtx, `name:"bectx"`),
	fxutil.AnnotatedProvide(bectx.NewBeAppCtx, `name:"beappctx"`),
	becommon.BeCommonModule,
	bedatabase.DatabaseModules,
	fx.Provide(bego.NewBegoApp),
	fx.Provide(bego.NewBegoCommand),
	healthcheck.HealthCheckModules,
	fx.Invoke(healthcheck.RegisterHealthLifecycleHooks),
	auth.AuthModules,
	fx.Invoke(auth.RegisterAuthLifecycleHooks),
)
