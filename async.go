package gaw

import (
	"context"
)

// Async will invoke Function in async mode
func Async[R any](ctx context.Context, opFn Function[R]) *Result[R] {
	h := newHandler[R](ctx, opFn)
	return h.handle()
}
