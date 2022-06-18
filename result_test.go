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

	go func() {
		f.Await() <- true
		close(f.Await())
	}()

	if !<-f.Await() {
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
