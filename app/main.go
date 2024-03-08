package main

import (
	"beconsole"
	"beconsole/command"
	"businessocean/bego"

	"go.uber.org/fx"
)

func main() {
	appCmd := make([]command.BeCommand, 0)

	var cmds = command.CmdSlice{
		"bego": bego.NewBegoServer(),
	}

	// Bego service for microservice
	appCmd = append(appCmd, command.NewCommand(cmds, Module))

	var businessocean = beconsole.NewBOConsole(appCmd)
	businessocean.Execute()
}

var Module = fx.Options()
