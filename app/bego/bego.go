package bego

import (
	"beconsole/command"

	"github.com/spf13/cobra"
)

type begoCommand struct{}

// NewBego returns a new instance of Bego
func NewBegoCommand() *begoCommand {
	return &begoCommand{}
}

func (b *begoCommand) Short() string {
	return "Run Bego Server "
}

func (b *begoCommand) Setup(cmd *cobra.Command) error {
	return nil
}

func (b *begoCommand) Run() command.CommandRunner {
	return func(app *BegoApp) {

		// fmt.Println("Bego is runnings", app.BeCtx, app.BeAppCtx)
		// fmt.Println("Bego is runnings")
		// fmt.Printf("Config: %v\n", &app.BaseApp)
		// app.logger.Info("Bego is running with logger")
		app.Bootstrap()
		app.IsDev()
		// app.Run()
		// app.Logger.Info("Bego is running with logger")

		// logger.Print("Bego is running with logger")
		// // for _, domain := range domains {
		// domain.Bootstrap()
		// domain.Run()
		// }

		// domainModules.Bootstrap()
	}
}
