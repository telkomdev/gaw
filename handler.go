package gaw

import (
	"context"
	"sync"
)

// Function base function type of gaw
type Function[R any] func() (R, error)

// handler base type of handle operation
type handler[R any] interface {
	handle() *Result[R]
}

// defaultHandler represent default handler
type defaultHandler[R any] struct {
	f func() *Result[R]
}

// handle will invoke the defaultHandler's f
func (d defaultHandler[R]) handle() *Result[R] {
	return d.f()
}

// newHandler the default handler constructor
func newHandler[R any](ctx context.Context, function Function[R]) handler[R] {
	return defaultHandler[R]{
		f: func() *Result[R] {
			r := NewResult[R]()
			mx := sync.RWMutex{}
			go func() {

				var (
					value R
					err   error
					done  = make(chan struct{}, 1)
				)

				defer func() { close(r.Await()) }()

				go func() {
					defer func() { close(done) }()

					mx.Lock()
					value, err = function()
					mx.Unlock()

					done <- struct{}{}
				}()

				select {
				case <-ctx.Done():
					mx.Lock()
					err = ctx.Err()
					mx.Unlock()
				case <-done:
				}

				mx.Lock()
				r.setValue(value)
				r.setErr(err)
				mx.Unlock()

				r.Await() <- true
			}()

			return r
		},
	}
}
