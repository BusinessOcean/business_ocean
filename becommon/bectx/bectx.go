package bectx

import "context"

type BeCtx context.Context

// NewBeCtx creates a new instance of BeCtx.
func NewBeCtx() BeCtx {
	return BeCtx(context.Background())
}
