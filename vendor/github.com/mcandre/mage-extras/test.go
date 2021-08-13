package mageextras

import (
	"os"
	"os/exec"
)

// UnitTest executes the Go unit test suite.
func UnitTest(args ...string) error {
	cmdName := "go"

	cmdParameters := []string{cmdName}
	cmdParameters = append(cmdParameters, "test")
	cmdParameters = append(cmdParameters, args...)

	cmd := exec.Command(cmdName)
	cmd.Args = cmdParameters
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
