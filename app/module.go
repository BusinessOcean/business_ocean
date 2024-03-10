package main

import (
	"becore"

	"go.uber.org/fx"
)

var AppModule = fx.Options(becore.BecoreModule)
