package bedomain

import (
	"becore/beconfig"
	"becore/belogger"
	"becore/beroutes"
	"becore/beserver"

	"google.golang.org/grpc"
)

var _ IBeDomain = (*BaseDomain)(nil)

type BaseDomain struct {
	config *beconfig.AppConfig
	server *beserver.BegoServer
	logger *belogger.BeLogger
}

// New BaseDomainModule
func NewBaseDomain(
	config *beconfig.AppConfig,
	logger *belogger.BeLogger,
	server *beserver.BegoServer,
) *BaseDomain {
	return &BaseDomain{
		config: config,
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

func (d *BaseDomain) Run(port string) error {
	// port, err := getRandomPort()
	// if err != nil {
	// 	return fmt.Errorf("Error: %v", err)
	// }

	// fmt.Println("Random available port:", port)

	d.server.RunServer(port)
	return nil
}

func (d *BaseDomain) OnTerminate() error {
	return nil
}

// func getRandomPort() (string, error) {
// 	listener, err := net.Listen("tcp", ":0") // Use ":0" to let the OS pick an available port
// 	if err != nil {
// 		return "", fmt.Errorf("could not find an available port: %v", err)
// 	}
// 	defer listener.Close()

// 	// Get the port from the listener
// 	addr := listener.Addr().(*net.TCPAddr)
// 	return fmt.Sprintf(":%d", addr.Port), nil
// }
