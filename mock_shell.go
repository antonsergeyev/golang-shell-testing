package main

import (
	"context"
)

// MockShell A shell implementation for testing.
// It always returns determinitistic results.
type MockShell struct {
	// an output and error to be returned when command is executed
	Output []byte
	Err    error
	// store the last executed command to be inspected
	LastCommand string
}

func (t *MockShell) Execute(ctx context.Context, cmd string) ([]byte, error) {
	t.LastCommand = cmd

	return t.Output, t.Err
}
