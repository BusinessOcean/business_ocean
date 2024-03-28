package bego

import (
	"becommon/bectx"
	"beconsole/command"
	"becore/belogger"
	"becore/beroutes"
	"becore/beserver"
	"fmt"
	"healthcheck"

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
	return func(
		logger *belogger.BeLogger,
		bego *beserver.BeServer,
		example *healthcheck.HealthAPI,
		ctx bectx.BeCtx,
		healthRoutes beroutes.BeRoute,
	) {
		logger.Info(`+-----------------------+`)
		logger.Info(`| Bego App ARCHITECTURE |`)
		logger.Info(`+-----------------------+`)
		fmt.Println("Bego Server is running....")
		routes := []beroutes.BeRoute{}
		routes = append(routes, healthRoutes)
		logger.Error("Bego Server is running....from beconsole")
		beroutes.NewRegisterRouteAPI(routes)
		bego.Run(iris.Addr(":8080"))

	}
}
