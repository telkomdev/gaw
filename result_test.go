package gaw

import (
	"errors"
	"testing"
)

func TestNewResult(t *testing.T) {
	f := NewResult[string]()

	if f == nil {
		t.Error("error: NewResult should return non nil")
	}
}

func TestResultIsDone(t *testing.T) {
	f := NewResult[string]()
	f.setValue("hello world")

	go func() {
		f.awaitDone <- true
		close(f.awaitDone)
	}()

	f.Await()

	expected := "hello world"

	if f.Get() != expected {
		t.Error("error: Await should return true")
	}

	if f.IsErr() {
		t.Error("error: IsErr should return false")
	}
}

func TestResultIsErr(t *testing.T) {
	f := NewResult[string]()
	f.setErr(errors.New("error returned"))

	if !f.IsErr() {
		t.Error("error: IsErr should return true")
	}
}
