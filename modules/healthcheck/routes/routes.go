package routes

import (
	"becore/beroutes"
	"becore/beserver"
)

func NewHealthCheckRoutes(server *beserver.BeHTTPServer) []beroutes.IRegisterRouteAPI {
	return []beroutes.IRegisterRouteAPI{
		&healthCheck{},
	}
}
