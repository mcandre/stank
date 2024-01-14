package mageextras

import (
	"os"
	"os/exec"
)

// Staticcheck runs staticcheck.
func Staticcheck(args ...string) error {
	cmd := exec.Command("staticcheck")
	cmd.Args = append(cmd.Args, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
