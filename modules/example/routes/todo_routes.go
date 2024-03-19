package routes

import (
	"becore/beroutes"
	"becore/beserver"
	"fmt"
)

type todo struct {
}

// routes = append(routes, beroutes.NewBeRoute(bego, todos{}))

func NewTodo(server *beserver.BeServer) beroutes.BeRoute {
	return beroutes.NewBeRoute(server, todo{})
}

// RegisterAPI implements beroutes.RegisterRouteAPI.
func (t todo) RegisterAPI(server *beserver.BeServer) {
	fmt.Println("Registering API")
	server.Get("/todos", func(ctx beserver.BeContext) {
		ctx.JSON([]string{"Write a blog post", "Write a book", "Write a song"})
	})

}
