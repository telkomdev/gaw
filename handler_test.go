package gaw

import (
	"context"
	"errors"
	"testing"
)

func TestNewHandler(t *testing.T) {
	h := newHandler[string](context.Background(), func() (string, error) {
		return "hello", nil
	})

	if h == nil {
		t.Error("error: newHandler should return non nil")
	}
}

func TestMultiHandlerHandleShouldReturnValue(t *testing.T) {
	h1 := newHandler[string](context.Background(), func() (string, error) {
		return "hello 1", nil
	})

	h2 := newHandler[string](context.Background(), func() (string, error) {
		return "hello 2", nil
	})

	// the test cases
	testCases := []struct {
		handler handler[string]
		want    string
	}{
		{
			handler: h1,
			want:    "hello 1",
		},

		{
			handler: h2,
			want:    "hello 2",
		},
	}

	for _, tc := range testCases {
		h := tc.handler.handle()
		<-h.Await()

		val := h.Get()
		if val != tc.want {
			t.Error("error: handle val should match want")
		}
	}
}

func TestMultiHandlerHandleShouldReturnError(t *testing.T) {
	h1 := newHandler[string](context.Background(), func() (string, error) {
		return "hello 1", errors.New("error: h1")
	})

	h2 := newHandler[string](context.Background(), func() (string, error) {
		return "", errors.New("error: h2")
	})

	// the test cases
	testCases := []struct {
		handler handler[string]
		want    bool
	}{
		{
			handler: h1,
			want:    true,
		},

		{
			handler: h2,
			want:    true,
		},
	}

	for _, tc := range testCases {
		h := tc.handler.handle()
		<-h.Await()

		err := h.Err()
		if (err != nil) != tc.want {
			t.Error("error: handle Err should return err")
		}
	}
}

func TestOneHandlerHandleShouldReturnValue(t *testing.T) {
	h1 := newHandler[string](context.Background(), func() (string, error) {
		return "hello 1", nil
	})

	// the test cases
	tc := struct {
		handler handler[string]
		want    string
	}{
		handler: h1,
		want:    "hello 1",
	}

	h := tc.handler.handle()
	<-h.Await()

	val := h.Get()
	if val != tc.want {
		t.Error("error: handle val should match want")
	}
}
