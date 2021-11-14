package main

import (
	"context"
	"os/exec"
)

type LocalShell struct{}

func (_ LocalShell) Execute(ctx context.Context, cmd string) ([]byte, error) {
	// wrap cmd into "sh -c '...command...'",
	// so that "special features" like output redirection work as expected.
	wrapperCmd := exec.CommandContext(ctx, "sh", "-c", cmd)
	output, err := wrapperCmd.CombinedOutput()

	return output, err
}
