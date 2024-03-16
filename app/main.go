package main

import (
	"beconsole"
	"beconsole/command"
	"beservice/example/todo"
	"businessocean/bego"
)

func main() {
	appCmd := make([]command.BeCommand, 0)

	var cmds = command.CmdSlice{
		"bego": bego.NewBegoServer(),
	}
	t := todo.Todo{}
	t.ValidateAll()

	// Bego service for microservice
	appCmd = append(appCmd, command.NewCommand(cmds, AppModule))

	var businessocean = beconsole.NewBOConsole(appCmd)
	businessocean.Execute()
}
