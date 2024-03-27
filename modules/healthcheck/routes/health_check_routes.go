package routes

import (
	"becore/beroutes"
	"becore/beserver"
	"fmt"

	"github.com/kataras/iris/v12"
)

type healthCheck struct {
}

// routes = append(routes, beroutes.NewBeRoute(bego, healthCheck{}))

func NewHealthCheckRoute(server *beserver.BeServer) beroutes.BeRoute {
	return beroutes.NewBeRoute(server, healthCheck{})
}

// RegisterAPI implements beroutes.RegisterRouteAPI.
func (t healthCheck) RegisterAPI(server *beserver.BeServer) {
	fmt.Println("Registering API")
	server.Get("/health", func(ctx beserver.BeContext) {
		ctx.JSON(iris.Map{"status": "Healthy"})
	})

}
