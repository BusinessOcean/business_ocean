package main

import (
	"beconsole"
	"beconsole/commands"
	"businessocean/bego"

	"go.uber.org/fx"
)

func main() {
	appCmd := make([]commands.BeCommand, 0)

	var cmds = commands.CmdSlice{
		"bego": bego.NewBegoServer(),
	}

	// Bego service for microservice
	appCmd = append(appCmd, commands.NewCommand(cmds, Module))

	var businessocean = beconsole.NewBOConsole(appCmd)
	businessocean.Execute()
}

var Module = fx.Options()
