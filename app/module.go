package main

import (
	"becore"
	"bedatabase"
	"example"

	"go.uber.org/fx"
)

var AppModule = fx.Options(
	becore.BecoreModule,
	example.ExampleModule,
	bedatabase.DatabaseModule,
)
