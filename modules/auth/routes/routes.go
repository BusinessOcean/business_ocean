package routes

import (
	"auth/service"
	"becore/beroutes"

	"github.com/kataras/iris/v12"
)

func NewAuthRoutes(service *service.AuthService) []*beroutes.Route {

	healthRoute := beroutes.NewRoutes(iris.MethodGet, "/authhealth", service.AuthHealthCheck)
	userCreatedRoute := beroutes.NewRoutes(iris.MethodGet, "/userCreated", service.CreateUserEvent)
	routes := []*beroutes.Route{healthRoute, userCreatedRoute}

	return routes
}
