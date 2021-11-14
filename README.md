## Testing local and remote shell commands in Go.

This is an (intentionally simplified) example of how unix shell commands can be unit-tested in Go.

The idea is to extract a standard `exec.Command` function into an interface and provide several implementations:

* `LocalShell` - a basic implementation based on standard library. Can be used to run shell commands locally.
* `MockShell` - an implementation for testing. Can be used to write deterministic unit-tests for shell commands.
* `RemoteShell` - an implementation based on [melbahja/goph](github.com/melbahja/goph) package for running shell commands on remote machine over SSH.

To show how `Shell` interface can be applied, there is a `ProcessInspector` service that lists processes in the operating system using `ps`. It does not need to know whether it is executed locally or over SSH, and can be covered with unit tests.

Using this approach, it is relatively easy to test shell commands execution in Go.
