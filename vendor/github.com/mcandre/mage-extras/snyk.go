package mageextras

import (
	"os"
	"os/exec"
)

// SnykTest executes a snyk security audit.
func SnykTest(args ...string) error {
	cmdName := "snyk"

	cmdParameters := []string{cmdName}
	cmdParameters = append(cmdParameters, "test")
	cmdParameters = append(cmdParameters, args...)

	cmd := exec.Command(cmdName)
	cmd.Args = cmdParameters
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
