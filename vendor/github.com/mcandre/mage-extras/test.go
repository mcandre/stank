package mageextras

import (
	"os"
	"os/exec"
)

// UnitTest executes the Go unit test suite.
func UnitTest(args ...string) error {
	cmd := exec.Command("go")
	cmd.Args = append(cmd.Args, "test")
	cmd.Args = append(cmd.Args, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
