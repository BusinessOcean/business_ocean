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

// type HealthCheckModuleParams struct {
// 	fx.In
// 	Lifecycle fx.Lifecycle
// 	Logger    *belogger.BeLogger

// 	AuthRoutes []*beroutes.Route    `name:"authroutes"`
// 	AuthDomain []bedomain.IBeDomain `group:"domains"`
// }

var AppModule = fx.Options(
	fxutil.AnnotatedProvide(bectx.NewBeCtx, `name:"bectx"`),
	fxutil.AnnotatedProvide(bectx.NewBeAppCtx, `name:"beappctx"`),
	fx.Provide(bego.NewBegoApp),
	fx.Provide(bego.NewBegoCommand),
	fx.Provide(bedomain.NewBaseDomain),
	becore.BecoreModule,
	becommon.BeCommonModule,
	bedatabase.DatabaseModules,
	auth.AuthModules,
	healthcheck.HealthCheckModules,
	// fx.Provide(bego.NewBeAppServiceDomain),
)
