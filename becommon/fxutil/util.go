package fxutil

import "go.uber.org/fx"

// AnnotatedProvide wraps fx.Annotate to make it simpler to use.
func AnnotatedProvide(constructor interface{}, resultTag string) fx.Option {
	return fx.Provide(
		fx.Annotate(
			constructor,
			fx.ResultTags(resultTag),
		),
	)
}
