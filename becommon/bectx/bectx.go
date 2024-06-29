package bectx

import "context"

type BeCtx context.Context

type BeAppCtx context.Context

// NewBeCtx creates a new instance of BeCtx.
func NewBeCtx() BeCtx {
	return BeCtx(context.Background())
}

// NewBeAppCtx creates a new instance of BeAppCtx.
func NewBeAppCtx() *BeAppCtx {
	ctx := BeAppCtx(context.Background())
	return &ctx
}
