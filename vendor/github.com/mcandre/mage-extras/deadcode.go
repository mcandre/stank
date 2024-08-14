package mageextras

import (
	"os"
	"os/exec"
)

// Deadcode runs deadcode.
func Deadcode(args ...string) error {
	cmd := exec.Command("deadcode")
	cmd.Args = append(cmd.Args, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
