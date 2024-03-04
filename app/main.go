package main

import (
	"beconsole"
	"beconsole/commands"
)

func main() {
	commands := []commands.BeCommand{}

	var BusinessOcean = beconsole.NewBOConsole(commands)
	_ = BusinessOcean
}
