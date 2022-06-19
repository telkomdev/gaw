package gaw

import (
	"context"
	"sync"
)

// Function base function type of gaw
type Function[R any] func() (R, error)

// invokeFunction will invoke the param function
// and set the value of value and err
func invokeFunction[R any](
	value *R,
	err *error,
	function Function[R],
	mx sync.RWMutex,
	done chan<- struct{}) {

	defer func() { close(done) }()

	mx.Lock()
	*value, *err = function()
	mx.Unlock()

	done <- struct{}{}
}

// handle will handle function invocation
func handle[R any](ctx context.Context,
	function Function[R]) *Result[R] {
	r := NewResult[R]()
	mx := sync.RWMutex{}
	go func() {

		var (
			value R
			err   error
			done  = make(chan struct{}, 1)
		)

		defer func() { close(r.awaitDone) }()

		go invokeFunction(&value, &err, function, mx, done)

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

		r.awaitDone <- true
	}()

	return r
}
