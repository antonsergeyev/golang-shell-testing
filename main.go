package main

import (
	"fmt"
	"log"
)

func main() {
	shell := LocalShell{}

	// Uncomment to run a process inspector over SSH
	/*shell, err := NewRemoteShell(RemoteShellOptions{
		Host:        "server-address",
		Port:        22,
		User:        "username",
		KeyFile:     "/home/username/.ssh/id_rsa",
	})
	if err != nil {
		log.Fatalln(err)
	}*/

	processInspector := ProcessInspector{shell: shell}
	processes, err := processInspector.GetProcesses(true)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%+v\n", processes)
}
