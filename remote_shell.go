package main

import (
	"context"
	"github.com/melbahja/goph"
	"golang.org/x/crypto/ssh"
	"time"
)

// RemoteShell wraps a melbahja/goph SSH client and implements our Shell interface
type RemoteShell struct {
	conn *goph.Client
	opts RemoteShellOptions
}

type RemoteShellOptions struct {
	Host        string
	Port        uint
	User        string
	Password    string
	KeyFile     string
	KeyPassword string
}

// Creates an SSH client using melbahja/goph library.
func NewRemoteShell(opts RemoteShellOptions) (*RemoteShell, error) {
	var err error
	var auth goph.Auth
	shell := RemoteShell{opts: opts}

	if len(opts.KeyFile) > 0 {
		auth, err = goph.Key(opts.KeyFile, opts.KeyPassword)
		if err != nil {
			return nil, err
		}
	} else {
		auth = goph.Password(opts.Password)
	}

	shell.conn, err = goph.NewConn(&goph.Config{
		Auth:     auth,
		User:     opts.User,
		Addr:     opts.Host,
		Port:     opts.Port,
		Timeout:  time.Second * 2,
		Callback: ssh.InsecureIgnoreHostKey(),
	})
	if err != nil {
		return nil, err
	}

	return &shell, nil
}

// Execute executes a command over SSH.
// `Goph` library provides exactly the function we need, so we just wrap it here.
func (r *RemoteShell) Execute(ctx context.Context, cmd string) ([]byte, error) {
	sshCmd, err := r.conn.CommandContext(ctx, cmd)
	if err != nil {
		return nil, err
	}

	return sshCmd.CombinedOutput()
}
