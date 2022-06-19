package gaw

import (
	"context"
	"errors"
	"testing"
)

func TestNewHandler(t *testing.T) {
	r := handle[string](context.Background(), func() (string, error) {
		return "hello", nil
	})

	if r == nil {
		t.Error("error: newHandler should return non nil")
	}
}

func TestMultiHandlerHandleShouldReturnValue(t *testing.T) {
	r1 := handle[string](context.Background(), func() (string, error) {
		return "hello 1", nil
	})

	r2 := handle[string](context.Background(), func() (string, error) {
		return "hello 2", nil
	})

	// the test cases
	testCases := []struct {
		result *Result[string]
		want   string
	}{
		{
			result: r1,
			want:   "hello 1",
		},

		{
			result: r2,
			want:   "hello 2",
		},
	}

	for _, tc := range testCases {
		tc.result.Await()

		val := tc.result.Get()
		if val != tc.want {
			t.Error("error: handle val should match want")
		}
	}
}

func TestMultiHandlerHandleShouldReturnError(t *testing.T) {
	r1 := handle[string](context.Background(), func() (string, error) {
		return "hello 1", errors.New("error: r1")
	})

	r2 := handle[string](context.Background(), func() (string, error) {
		return "", errors.New("error: r2")
	})

	// the test cases
	testCases := []struct {
		result *Result[string]
		want   bool
	}{
		{
			result: r1,
			want:   true,
		},

		{
			result: r2,
			want:   true,
		},
	}

	for _, tc := range testCases {
		tc.result.Await()

		err := tc.result.Err()
		if (err != nil) != tc.want {
			t.Error("error: handle Err should return err")
		}
	}
}

func TestOneHandlerHandleShouldReturnValue(t *testing.T) {
	r1 := handle[string](context.Background(), func() (string, error) {
		return "hello 1", nil
	})

	// the test cases
	tc := struct {
		result *Result[string]
		want   string
	}{
		result: r1,
		want:   "hello 1",
	}

	tc.result.Await()

	val := tc.result.Get()
	if val != tc.want {
		t.Error("error: handle val should match want")
	}
}
