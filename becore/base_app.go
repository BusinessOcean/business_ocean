package becore

import (
	"becommon/bectx"
	"becommon/beevent"
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

	eventBus *beevent.BeEventBus
}

func NewBaseApp(
	appCtx *bectx.BeAppCtx,
	config *beconfig.AppConfig,
	logger *belogger.BeLogger,
	server *beserver.BegoServer,
	eventBus *beevent.BeEventBus,
) *BaseApp {
	return &BaseApp{
		appCtx:   appCtx,
		config:   config,
		logger:   logger,
		server:   server,
		eventBus: eventBus,
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

func (app *BaseApp) EventBus() *beevent.BeEventBus {
	// Implementation of the OnTerminate method

	return app.eventBus
}

func (app *BaseApp) IsDev() bool {
	if app.config.Info.Mode == "development" {
		return true
	}
	return false
}
