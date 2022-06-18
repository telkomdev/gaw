package gaw

import (
	"context"
)

// Async will invoke Function in async mode
func Async[R any](ctx context.Context, function Function[R]) *Result[R] {
	h := newHandler[R](ctx, function)
	return h.handle()
}
