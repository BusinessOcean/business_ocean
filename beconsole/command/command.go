package command

import (
	"becore/belogger"
	"context"
	"fmt"

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
	Run() CommandRunner
}

type CmdSlice map[string]Command

type BeCommand struct {
	cmdSlice   map[string]Command
	dependency fx.Option
}

// new Becommand
func NewCommand(cmdSlice CmdSlice, dependency fx.Option) BeCommand {
	return BeCommand{cmdSlice: cmdSlice, dependency: dependency}
}

// GetSubCommands gives a list of sub commands
func GetSubCommands(becmd BeCommand) []*cobra.Command {
	subCommands := make([]*cobra.Command, 0)
	for name, cmd := range becmd.cmdSlice {
		subCommands = append(subCommands, WrapSubCommand(name, cmd, becmd.dependency))
	}
	return subCommands
}

// WrapSubCommand wraps a subcommand with additional functionality and returns the wrapped command.
// It takes the name of the command, the command itself, and an fx.Option as input parameters.
// The wrapped command is created with the provided name and the short description from the command.
// When the wrapped command is executed, it sets up a logger, invokes the command's Run function using fx.Invoke,
// and starts an fx application with the provided fx.Option and fx.Options.
// The application is then started and stopped using the provided context.
// If any errors occur during the execution, they are logged using the logger.
// Finally, the wrapped command is returned.
func WrapSubCommand(name string, cmd Command, opt fx.Option) *cobra.Command {
	// Wait for termination signal (SIGINT or SIGTERM)
	// c := make(chan os.Signal, 1)
	// signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	// <-c

	wrappedCmd := &cobra.Command{
		Use:   name,
		Short: cmd.Short(),
		Run: func(c *cobra.Command, args []string) {
			// TODO: Use the logger from betools
			logger := belogger.NewBeLogger()

			opts := fx.Options(
				// TAKE CARE: Need to use the logger from betools
				// TODO: fx.WithLogger(logger.GetFxLogger),
				fx.Invoke(cmd.Run()),
			)
			ctx := context.Background()
			//  Hide uber fx logs : fx.NopLogger
			// app := fx.New(opt, opts, fx.NopLogger)
			app := fx.New(opt, opts)
			err := app.Start(ctx)
			defer func() {
				fmt.Println("Stopping the application")
				err = app.Stop(ctx)
				if err != nil {
					logger.Fatal(err)
				}
			}()
			if err != nil {
				logger.Fatal(err)
			}
			// Wait for the application to shutdown
			<-app.Done()
		},
	}
	cmd.Setup(wrappedCmd)
	return wrappedCmd
}
