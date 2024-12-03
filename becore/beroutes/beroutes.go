package beroutes

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
)

type BeApiBuiler struct{ *router.Route }

func NewRoutes(method string,
	path string,
	handler iris.Handler,
	middlewares ...iris.Handler) *Route {
	return &Route{
		Method:      method,
		Path:        path,
		Handler:     handler,
		Middlewares: middlewares,
	}
}

// Route represents a single route with its method, path, and handlers
type Route struct {
	Method      string
	Path        string
	Handler     iris.Handler
	Middlewares []iris.Handler
}
