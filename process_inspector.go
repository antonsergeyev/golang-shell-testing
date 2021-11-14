package main

import (
	"context"
	"strings"
)

// ProcessInspector lists currently running processes in the OS.
// It depends on the `Shell` interface and can be used both locally and over SSH.
type ProcessInspector struct {
	shell Shell
}

func (p *ProcessInspector) GetProcesses(all bool) ([]string, error) {
	cmd := "ps"

	if all {
		cmd += " -a"
	}

	output, err := p.shell.Execute(context.Background(), cmd)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	processes := []string{}

	for i, line := range lines {
		if i == 0 {
			// skip the header line
			continue
		}

		processes = append(processes, strings.TrimSpace(line))
	}

	return processes, nil
}
