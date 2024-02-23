package io

import (
	"errors"
	"testing"
)

func TestNopCloser(t *testing.T) {
	err := NopCloser.Close()
	if err != nil {
		t.Errorf("NopCloser.Close() returned an error: %v", err)
	}
}

func TestCustomCloserSuccess(t *testing.T) {
	closer := NewCloser(func() error {
		return nil // Simulate successful close
	})

	err := closer.Close()
	if err != nil {
		t.Errorf("CustomCloser.Close() returned an error on success path: %v", err)
	}
}

func TestCustomCloserFailure(t *testing.T) {
	expectedErr := errors.New("close error")
	closer := NewCloser(func() error {
		return expectedErr // Simulate failure to close
	})

	err := closer.Close()
	if !errors.Is(err, expectedErr) {
		t.Errorf("CustomCloser.Close() did not return expected error on failure path: got %v, want %v", err, expectedErr)
	}
}

func TestCloseFunction(t *testing.T) {
	called := false
	closer := NewCloser(func() error {
		called = true
		return nil // or return an error to simulate failure
	})

	Close(closer)

	if !called {
		t.Errorf("Close function did not call the Closer's Close method")
	}
}
