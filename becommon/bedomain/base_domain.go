package bedomain

import (
	"becore/belogger"
	"becore/beroutes"
	"becore/beserver"

	"google.golang.org/grpc"
)

var _ IBeDomain = (*BaseDomain)(nil)

type BaseDomain struct {
	// config *BeDomainConfig
	server *beserver.BegoServer
	logger *belogger.BeLogger
}

// New BaseDomainModule
func NewBaseDomain(
	// config *BeDomainConfig,
	logger *belogger.BeLogger,
	server *beserver.BegoServer,
) *BaseDomain {
	return &BaseDomain{
		// config: config,
		logger: logger,
		server: server,
	}
}

// -------------------------------------------------------------------
// Implementation of base event hooks
// -------------------------------------------------------------------

func (d *BaseDomain) Setup() error {

	// set configuration for domain
	return nil
}

func (d *BaseDomain) GetServer() *beserver.BegoServer {
	return d.server
}

func (d *BaseDomain) Register(desc grpc.ServiceDesc, service interface{}) error {
	d.server.Register(desc, service)

	return nil
}

func (d *BaseDomain) RegisterRoutes(desc grpc.ServiceDesc, routes []*beroutes.Route) error {
	d.server.RegisterRoutes(desc.ServiceName, routes)
	return nil
}

func (d *BaseDomain) Run() error {
	d.server.RunServer()
	return nil
}

func (d *BaseDomain) OnTerminate() error {
	return nil
}
