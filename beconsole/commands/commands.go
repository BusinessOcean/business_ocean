package commands

import (
	"betools"
	"context"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

type CommandRunner any

// Command interface is used to implement sub-commands in the system.
type Command interface {
	// Short returns string about short description of the command
	// the string is shown in help screen of cobra command
	Short() string

	// Setup is used to setup flags or pre-run steps for the command.
	//
	// For example,
	//  cmd.Flags().IntVarP(&r.num, "num", "n", 5, "description")
	//
	Setup(cmd *cobra.Command) error

	// Run runs the command runner
	// run returns command runner which is a function with dependency
	// injected arguments.
	//
	// For example,
	//  Command{
	//   Run: func(l lib.Logger) {
	// 	   l.Info("i am working")
	// 	 },
	//  }
	//
	Run() (CommandRunner, error)
}

type CmdSlice map[string]Command

type BeCommand struct {
	commands CmdSlice
	opt      fx.Option
}

// GetSubCommands gives a list of sub commands
func GetSubCommands(becmd BeCommand) []*cobra.Command {
	subCommands := make([]*cobra.Command, 0)
	for name, cmd := range becmd.commands {
		subCommands = append(subCommands, WrapSubCommand(name, cmd, becmd.opt))
	}
	return subCommands
}

func WrapSubCommand(name string, cmd Command, opt fx.Option) *cobra.Command {
	wrappedCmd := &cobra.Command{
		Use:   name,
		Short: cmd.Short(),
		Run: func(c *cobra.Command, args []string) {
			// TODO: Use the logger from betools
			logger := betools.NewLogger()

			opts := fx.Options(
				// TAKE CARE: Need to use the logger from betools
				// TODO: fx.WithLogger(logger.GetFxLogger),
				fx.Invoke(cmd.Run()),
			)
			ctx := context.Background()
			app := fx.New(opt, opts)
			err := app.Start(ctx)
			defer func() {
				err = app.Stop(ctx)
				if err != nil {
					logger.Fatal(err)
				}
			}()
			if err != nil {
				logger.Fatal(err)
			}
		},
	}
	cmd.Setup(wrappedCmd)
	return wrappedCmd
}
