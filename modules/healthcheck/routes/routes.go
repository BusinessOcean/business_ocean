package routes

import (
	"becore/beroutes"
	"healthcheck/service"

	"github.com/kataras/iris/v12"
)

func NewHealthCheckRoutes(service *service.HealthCheckService) []*beroutes.Route {

	healthRoute := beroutes.NewRoutes(iris.MethodGet, "/health", service.HealthCheckApiRoute)
	healthRoute2 := beroutes.NewRoutes(iris.MethodGet, "/health1", service.HealthCheckApiRoute2)
	routes := []*beroutes.Route{healthRoute, healthRoute2}

	return routes
}
