package gaw

import (
	"context"
	"errors"
	"testing"
	"time"
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
	async := Async[string](context.Background(), func() (string, error) {
		// simulate heavy work that takes 3 seconds to finish
		time.Sleep(time.Second * 1)

		return "hello 1", nil
	})

	// the test cases
	expected := ""

	// emit the Await
	// <-async.Await()

	val := async.Get()

	if val != expected {
		t.Error("error: async val should return empty")
	}
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
