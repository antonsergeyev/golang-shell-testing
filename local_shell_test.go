package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"
)

func TestLocalShell_Execute_ls(t *testing.T) {
	shell := LocalShell{}
	output, err := shell.Execute(context.Background(), "ls")
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(output), "local_shell_test.go") {
		t.Errorf("expected output to contain a go file, got:\n%s", string(output))
	}
}

func TestLocalShell_Execute_output_redirection(t *testing.T) {
	// create a temporary sandbox directory for our test
	tempDir, err := ioutil.TempDir(os.TempDir(), "go-shell-test*")
	if err != nil {
		t.Fatal(err)
	}

	shell := LocalShell{}

	// write to a text file using output redirection
	_, err = shell.Execute(
		context.Background(),
		fmt.Sprintf(
			"echo hello > %s/greeting.txt",
			tempDir,
		),
	)
	if err != nil {
		t.Fatal(err)
	}

	// read a file we just created
	output, err := shell.Execute(
		context.Background(),
		fmt.Sprintf(
			"cat %s/greeting.txt",
			tempDir,
		),
	)

	if err != nil {
		t.Fatal(err)
	}

	if strings.TrimSpace(string(output)) != "hello" {
		t.Errorf("unexpected output: %s", string(output))
	}
}

// Ensure LocalShell supports context cancellation.
// We expect that "sleep 5" fails quickly because of a context with timeout.
func TestLocalShell_Execute_context_cancel(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*10)
	defer cancel()

	shell := LocalShell{}
	_, err := shell.Execute(ctx, "sleep 5")
	if err == nil {
		t.Fatal("expected an error because context has been cancelled")
	}
}
