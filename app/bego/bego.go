package bego

import (
	"beconsole/command"
	"fmt"

	"github.com/spf13/cobra"
)

type bego struct{}

// NewBego returns a new instance of Bego
func NewBegoServer() *bego {
	return &bego{}
}

func (b *bego) Short() string {
	return " Run Bego Server "
}

func (b *bego) Setup(cmd *cobra.Command) error {
	return nil
}

func (b *bego) Run() command.CommandRunner {

	return func() {

		fmt.Println("I am working")

	}
}
