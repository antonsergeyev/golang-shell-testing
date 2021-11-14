package main

import "testing"

func TestProcessInspector_GetProcesses_all(t *testing.T) {
	// prepare a mock shell with an example output from "ps -a"
	shell := MockShell{
		Output: []byte(`
    PID TTY          TIME CMD
   1755 tty2     00:04:38 Xorg
   1788 tty2     00:00:00 gnome-session-b
`),
		Err: nil,
	}

	// instantiate process inspector with a mock shell
	inspector := ProcessInspector{shell: &shell}

	processes, err := inspector.GetProcesses(true)
	if err != nil {
		t.Fatal(err)
	}

	if shell.LastCommand != "ps -a" {
		t.Errorf("expected the executed command to be 'ps -a', got '%s'", shell.LastCommand)
	}

	if len(processes) != 2 {
		t.Errorf("expected 2 processes from ps -a, got %d", len(processes))
	}

	// ...more tests
}
