package bego

import (
	"beconsole/command"

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
		domains BegoDomains,
	) {

		for _, domain := range domains {
			domain.Bootstrap()
			domain.Run()
		}

		// domainModules.Bootstrap()
	}
}
