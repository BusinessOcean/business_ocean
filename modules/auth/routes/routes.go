package routes

import (
	"auth/service"
	"becore/beroutes"

	"github.com/kataras/iris/v12"
)

func NewAuthRoutes(service *service.AuthService) []*beroutes.Route {

	healthRoute := beroutes.NewRoutes(iris.MethodGet, "/health", service.AuthHealthCheck)
	routes := []*beroutes.Route{healthRoute}

	return routes
}
