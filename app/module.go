package main

import (
	"becore"
	"example"

	"go.uber.org/fx"
)

var AppModule = fx.Options(
	becore.BecoreModule,
	example.ExampleModule,
)
