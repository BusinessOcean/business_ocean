package bego

import (
	"beconsole/command"
	"becore/belogger"
	"becore/beroutes"
	"becore/beserver"
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/spf13/cobra"
)

type begoCommand struct{}

// NewBego returns a new instance of Bego
func NewBegoServer() *begoCommand {
	return &begoCommand{}
}

func (b *begoCommand) Short() string {
	return " Run Bego Server "
}

func (b *begoCommand) Setup(cmd *cobra.Command) error {
	return nil
}

func (b *begoCommand) Run() command.CommandRunner {

	return func(logger *belogger.BeLogger, bego *beserver.BeServer) {
		fmt.Println("Bego Server is running....")
		routes := []beroutes.BeRoute{}
		routes = append(routes, beroutes.NewBeRoute(bego, todos{}))
		logger.Error("Bego Server is running....from beconsole")
		beroutes.NewRegisterRouteAPI(routes)

		bego.Run(iris.Addr(":8080"))

	}
}

type todos struct {
}

// RegisterAPI implements beroutes.RegisterRouteAPI.
func (t todos) RegisterAPI(server *beserver.BeServer) {
	fmt.Println("Registering API")
	server.Get("/todos", func(ctx beserver.BeContext) {
		ctx.JSON([]string{"Write a blog post", "Write a book", "Write a song"})
	}, nil)

}
