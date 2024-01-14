package mageextras

import (
	"os"
	"os/exec"
)

// SnykTest executes a snyk security audit.
func SnykTest(args ...string) error {
	cmd := exec.Command("snyk")
	cmd.Args = append(cmd.Args, "test")
	cmd.Args = append(cmd.Args, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
