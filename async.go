package gaw

import (
	"context"
)

// Async will invoke Function in async mode
func Async[R any](ctx context.Context, function Function[R]) *Result[R] {
	return handle[R](ctx, function)
}

// AsyncAll will invoke multiple Function in async mode
func AsyncAll[R any](ctx context.Context, functions ...Function[R]) Results[R] {

	var (
		results Results[R]
	)

	for _, function := range functions {
		results = append(results, Async[R](ctx, function))
	}
	return results
}
