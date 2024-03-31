package beroutes

import (
	"becore/beserver"
)

// Route interface
type IRegisterRouteAPI interface {
	RegisterAPI(*beserver.BeHTTPServer)
}

// NewRoutes sets up routes
func NewRegisterRouteAPI(server *beserver.BeHTTPServer, routes []IRegisterRouteAPI) {
	for _, route := range routes {
		route.RegisterAPI(server)
	}
}
