package bedomain

import (
	"becommon/beevents"
	"becore/behook"
	"becore/belogger"
	"becore/beserver"
)

var _ IBeDomainModule = (*BaseDomainModule)(nil)

type BaseDomainModule struct {
	config     *BeDomainConfig
	grpcServer *beserver.BeGrpcServer
	httpServer *beserver.BeHTTPServer
	logger     belogger.BeLogger
}

// New BaseDomainModule
func NewBaseDomainModule(
	logger belogger.BeLogger,
	grpcServer *beserver.BeGrpcServer,
	httpServer *beserver.BeHTTPServer,
) *BaseDomainModule {
	return &BaseDomainModule{
		logger:     logger,
		grpcServer: grpcServer,
		httpServer: httpServer,
	}
}

// -------------------------------------------------------------------
// Implementation of base event hooks
// -------------------------------------------------------------------

// Bootstrap implements IBaseDomainModule.

func (app *BaseDomainModule) Bootstrap() error {
	config, err := app.LoadConfig()
	if err != nil {
		return err
	}
	app.config = config
	return nil
}
func (app *BaseDomainModule) Run() error {
	panic("implement Run in Domain Module")
}

func (app *BaseDomainModule) OnBeforeBootstrap() *behook.Hook[*beevents.ExampleHookEvent] {
	return nil
}

func (app *BaseDomainModule) OnTerminate() *behook.Hook[*beevents.ExampleHookEvent] {
	return nil
}

// LoadConfig implements IBaseDomainModule.
func (app *BaseDomainModule) LoadConfig() (*BeDomainConfig, error) {
	app.logger.Errorf("LoadConfig is not implemented")

	config := BeDomainConfig{Port: "8080", DomainName: "localhost"}

	return &config, nil
}

// GetGrpcServer returns the grpcServer variable.
func (app *BaseDomainModule) GetGrpcServer() *beserver.BeGrpcServer {
	return app.grpcServer
}

// GetHTTPServer returns the httpServer variable.
func (app *BaseDomainModule) GetHTTPServer() *beserver.BeHTTPServer {
	return app.httpServer
}

// GetLogger returns the logger variable.
func (app *BaseDomainModule) GetLogger() belogger.BeLogger {
	return app.logger
}

// GetLogger returns the logger variable.
func (app *BaseDomainModule) GetConfig() *BeDomainConfig {
	return app.config
}
