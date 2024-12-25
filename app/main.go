package main

import (
	"beconsole"
	"beconsole/command"
	"businessocean/bego"
)

func main() {
	appCmd := []command.BeCommand{}

	var cmds = command.CmdSlice{
		"bego": bego.NewBegoCommand(),
	}
	// run Bego App Command
	appCmd = append(appCmd, command.NewCommand(cmds, AppModule))

	var appConsole = beconsole.NewBeConsole(appCmd)
	appConsole.Execute()
}
