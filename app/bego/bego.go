package bego

import (
	"beconsole/command"
	"becore/beconfig"
	"becore/beroutes"
	"becore/beserver"
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/spf13/cobra"
)

type bego struct{}

// NewBego returns a new instance of Bego
func NewBegoServer() *bego {
	return &bego{}
}

func (b *bego) Short() string {
	return " Run Bego Server "
}

func (b *bego) Setup(cmd *cobra.Command) error {
	return nil
}

func (b *bego) Run() command.CommandRunner {

	return func() {
		fmt.Println("Bego Server is running....")
		_bego := beserver.NewBeServer(beconfig.ServerConfig{AppName: "Bego"})
		_ = _bego
		routes := []beroutes.BeRoute{}
		routes = append(routes, beroutes.NewBeRoute(_bego, todos{}))

		beroutes.NewRegisterRouteAPI(routes)

		_bego.Run(iris.Addr(":8080"))

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
