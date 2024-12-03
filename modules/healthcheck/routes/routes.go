package routes

import (
	"becore/beroutes"
	"healthcheck/service"

	"github.com/kataras/iris/v12"
)

func NewHealthCheckRoutes(service *service.HealthCheckService) []*beroutes.Route {

	healthRoute := beroutes.NewRoutes(iris.MethodGet, "/health", service.HealthCheckApiRoute)
	routes := []*beroutes.Route{healthRoute}

	return routes
}
