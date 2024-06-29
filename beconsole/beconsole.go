package beconsole

import (
	"beconsole/command"

	"github.com/spf13/cobra"
)

// _rootConsole variable for holding the root command
var _rootConsole = &cobra.Command{
	Use:   "Business Ocean  Console Engine",
	Short: "Business Ocean Console Engine is a command runner for api architecture in golang.",
	Long: `
         Business Ocean server for centeralized logging and monitoring of the system. It also provides
		 microservices architecture for the system. It is a command runner for api architecture in golang.
		 migration, seed, server, and other commands are available to run the system.
		 backend and frontend are the two main parts of the system. Backend is the server and frontend is the client.
		 divided into smaller parts and each part is a microservice. Each microservice is a separate module and consumed
		 by app.
		  
		 Naming convention:

		 [BE] stand for Bego foundation block of the system.
		 
		 [BO] stand for Business Ocean. Business Ocean is the name of the system.

		 boConsole is the command runner for the system. It is a command runner for api architecture in golang and services.
		 all are connected to each other and work together to run the system.
		 Golang work is used to run the system. It is a command runner for api architecture in golang.
		 Flutter is used for frontend and golang is used for backend. 
	`,
	TraverseChildren: true,
}

type BeConsole struct {
	*cobra.Command
}

// NewBeConsole returns a new instance of BeConsole
func NewBeConsole(beCommands []command.BeCommand) BeConsole {
	bc := BeConsole{
		Command: _rootConsole,
	}

	for _, cmd := range beCommands {
		bc.AddCommand(command.GetSubCommands(cmd)...)
	}
	return bc
}
