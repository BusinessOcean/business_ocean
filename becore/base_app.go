package becore

import (
	"becommon/bectx"
	"becommon/bedomain"
	"becore/beconfig"
	"becore/belogger"
	"becore/beserver"
	"fmt"
)

var _ IBegoApp = (*BaseApp)(nil)

type BaseApp struct {
	appCtx *bectx.BeAppCtx
	config *beconfig.AppConfig
	logger *belogger.BeLogger
	server *beserver.BegoServer
	domain []bedomain.IBeDomain
}

func NewBaseApp(
	appCtx *bectx.BeAppCtx,
	config *beconfig.AppConfig,
	logger *belogger.BeLogger,
	server *beserver.BegoServer,
	domain []bedomain.IBeDomain,
) *BaseApp {
	return &BaseApp{
		appCtx: appCtx,
		config: config,
		logger: logger,
		server: server,
		domain: domain,
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

func (app *BaseApp) GetModules() bool {
	fmt.Println("Modules: ", app.domain)

	for _, domain := range app.domain {
		fmt.Println("Setting up domain:", len(app.domain))
		fmt.Printf("Domain type: %T\n", domain)
		domain.Setup()
	}
	return false
}
