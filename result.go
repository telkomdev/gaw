package gaw

import (
	"sync"
)

// Result represent things that eventually available
type Result[R any] struct {
	value R
	err   error
	await chan bool
	mx    sync.RWMutex
}

// NewResult the Result constructor
func NewResult[R any]() *Result[R] {
	return &Result[R]{
		await: make(chan bool, 1),
	}
}

// Get will return Result value
func (f *Result[R]) Get() R {
	f.mx.RLock()
	defer func() { f.mx.RUnlock() }()
	return f.value
}

// setValue will new value to Result's value
func (f *Result[R]) setValue(value R) {
	f.mx.Lock()
	defer func() { f.mx.Unlock() }()
	f.value = value
}

// Await will determine wheter Result is done or not done
func (f *Result[R]) Await() chan bool {
	return f.await
}

// Err will return error
func (f *Result[R]) Err() error {
	return f.err
}

// IsErr will return true if err not nil
func (r *Result[R]) IsErr() bool {
	return r.err != nil
}

// setErr will new error to Result's error
func (f *Result[R]) setErr(err error) {
	f.err = err
}
