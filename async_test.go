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

func TestAsyncAllShouldReturnValue(t *testing.T) {
	f1 := func() (string, error) {
		return "hello 1", nil
	}

	f2 := func() (string, error) {
		return "hello 2", nil
	}

	f3 := func() (string, error) {
		return "hello 3", nil
	}
	asyncAll := AsyncAll[string](context.Background(), []Function[string]{f1, f2, f3}...)

	// await
	asyncAll.Await()

	for _, v := range asyncAll.Get() {
		if v == "" {
			t.Error("error: asyncAll Get should return non empty string")
		}
	}
}

func TestAsyncAllShouldReturnError(t *testing.T) {
	f1 := func() (string, error) {
		return "", errors.New("error: f1")
	}

	f2 := func() (string, error) {
		return "", errors.New("error: f2")
	}

	f3 := func() (string, error) {
		return "", errors.New("error: f3")
	}

	asyncAll := AsyncAll[string](context.Background(), []Function[string]{f1, f2, f3}...)

	// await
	asyncAll.Await()

	for _, r := range asyncAll {
		err := r.Err()
		if err == nil {
			t.Error("error: asyncAll Err should return error")
		}
	}
}

func TestAsyncAllFireForget(t *testing.T) {
	f1 := func() (string, error) {
		return "hello 1", nil
	}

	f2 := func() (string, error) {
		return "hello 2", nil
	}

	f3 := func() (string, error) {
		return "hello 3", nil
	}
	asyncAll := AsyncAll[string](context.Background(), []Function[string]{f1, f2, f3}...)

	// omit the await
	// asyncAll.Await()

	// the test cases
	expected := ""

	for _, v := range asyncAll.Get() {
		if v != expected {
			t.Error("error: asyncAll Get should return empty string")
		}
	}
}

func TestAsyncFireForget(t *testing.T) {
	async := Async[string](context.Background(), func() (string, error) {
		// simulate heavy work that takes 1 seconds to finish
		time.Sleep(time.Second * 1)

		return "hello 1", nil
	})

	// the test cases
	expected := ""

	// omit the Await
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
