package gaw

import (
	"context"
	"errors"
	"testing"
	// "sync"
)

func TestAsyncShouldReturnValue(t *testing.T) {
	async := Async[string](context.Background(), func() (string, error) {
		return "hello 1", nil
	})

	// the test cases
	expected := "hello 1"
	<-async.Await()

	val := async.Get()
	if val != expected {
		t.Error("error: async val should match want")
	}
}

func TestAsyncFireForget(t *testing.T) {
	// mx := sync.RWMutex{}
	_ = Async[string](context.Background(), func() (string, error) {
		return "hello 1", nil
	})

	// // the test cases
	// expected := "hello 1"

	// // emit the Await
	// // <-async.Await()

	// mx.RLock()
	// val := async.Get()

	// if val == expected {
	// 	t.Error("error: async val should match want")
	// }
	// mx.RUnlock()
}

func TestAsyncShouldReturnError(t *testing.T) {
	async := Async[string](context.Background(), func() (string, error) {
		return "", errors.New("error: async")
	})

	<-async.Await()

	err := async.Err()
	if err == nil {
		t.Error("error: async Err should return error")
	}
}
