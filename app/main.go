package main

import (
	"beconsole"
	"beconsole/command"
	"businessocean/bego"
)

func main() {
	appCmd := []command.BeCommand{}

	var cmds = command.CmdSlice{
		"bego": bego.NewBegoServer(),
	}
	// Bego service for microservice
	appCmd = append(appCmd, command.NewCommand(cmds, AppModule))

	var businessocean = beconsole.NewBOConsole(appCmd)
	businessocean.Execute()
}
