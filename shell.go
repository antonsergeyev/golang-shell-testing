package main

import (
	"context"
)

// Shell An interface for running shell commands in the OS
type Shell interface {
	Execute(ctx context.Context, cmd string) ([]byte, error)
}
