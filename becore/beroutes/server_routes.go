package beroutes

import (
	"becore/beserver"
)

// Routes contains multiple routes
type BeRoute struct {
	*beserver.BeServer
	RegisterRouteAPI
}

func NewBeRoute(server *beserver.BeServer, route RegisterRouteAPI) BeRoute {
	return BeRoute{server, route}
}

// Route interface
type RegisterRouteAPI interface {
	RegisterAPI(*beserver.BeServer)
}

// NewRoutes sets up routes
func NewRegisterRouteAPI(routes []BeRoute) []RegisterRouteAPI {
	var convertedRoutes []RegisterRouteAPI

	for _, route := range routes {
		route.RegisterAPI(route.BeServer)
		convertedRoutes = append(convertedRoutes, route)
	}
	return convertedRoutes
}
