package routes

import (
	"becore/beroutes"
	"becore/beserver"
	"fmt"

	"github.com/kataras/iris/v12"
)

var _ beroutes.IRegisterRouteAPI = (*healthCheck)(nil)

type healthCheck struct {
}

// RegisterAPI implements beroutes.RegisterRouteAPI.
func (t healthCheck) RegisterAPI(server *beserver.BeHTTPServer) {
	fmt.Println("Registering API")
	server.Get("/health", func(ctx beserver.BeContext) {
		ctx.JSON(iris.Map{"status": "Healthy"})
	})

}
