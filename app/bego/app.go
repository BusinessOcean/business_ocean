package bego

import (
	"becommon/bectx"
	"becore"
	"becore/beconfig"
	"becore/belogger"
	"becore/beserver"

	"go.uber.org/fx"
)

type BegoAppParams struct {
	fx.In
	Beappctx *bectx.BeAppCtx `name:"beappctx"`
	Config   *beconfig.AppConfig
	Logger   *belogger.BeLogger
	Server   *beserver.BegoServer
}

// BegoApp is the main application struct
// Add more fields for synchronization, configuration, etc.
type BegoApp struct {
	*becore.BaseApp
}

func NewBegoApp(p BegoAppParams) *BegoApp {
	baseApp := becore.NewBaseApp(p.Beappctx, p.Config, p.Logger, p.Server)
	app := &BegoApp{BaseApp: baseApp}
	return app
}
