package becore

import (
	"becommon/bectx"
	"becore/beconfig"
	"becore/belogger"
	"becore/beserver"
)

var _ IBegoApp = (*BaseApp)(nil)

type BaseApp struct {
	appCtx *bectx.BeAppCtx
	config *beconfig.AppConfig
	logger *belogger.BeLogger
	server *beserver.BegoServer
}

func NewBaseApp(
	appCtx *bectx.BeAppCtx,
	config *beconfig.AppConfig,
	logger *belogger.BeLogger,
	server *beserver.BegoServer,
) *BaseApp {
	return &BaseApp{
		appCtx: appCtx,
		config: config,
		logger: logger,
		server: server,
	}
}

// Bootstrap is a method to initialize the application
func (app *BaseApp) Bootstrap() error {
	// Implementation of the Bootstrap method
	return nil
}

func (app *BaseApp) OnTerminate() error {
	// Implementation of the OnTerminate method

	return nil
}
func (app *BaseApp) Run() error {
	app.server.RunServer()
	return nil
}

func (app *BaseApp) IsDev() bool {
	if app.config.Info.Mode == "development" {
		return true
	}
	return false
}
