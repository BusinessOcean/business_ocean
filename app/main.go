package app

import (
	"fmt"

	"go.uber.org/fx"
)

// Define a generic type
type Item[T any] struct {
	Value T
}

func main() {
	// Create a new UberFx application
	app := fx.New(
		fx.Provide(NewItem),
		fx.Invoke(UseItem),
	)

	// Run the application
	app.Run()
}

// Create a new Item
func NewItem() Item[string] {
	return Item[string]{Value: "Hello, World!"}
}

// Use the Item
func UseItem(item Item[string]) {
	fmt.Println(item.Value)
}
