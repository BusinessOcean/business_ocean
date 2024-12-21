package routes

import (
	"auth/service"
	"becore/beroutes"

	"github.com/kataras/iris/v12"
)

func NewAuthRoutes(service *service.AuthService) []*beroutes.Route {

	authRoute := beroutes.NewRoutes(iris.MethodGet, "/auth", service.AuthApiRoute)
	authRoute2 := beroutes.NewRoutes(iris.MethodGet, "/auth1", service.AuthApiRoute2)
	routes := []*beroutes.Route{authRoute, authRoute2}

	return routes
}
