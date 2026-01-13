package mageextras

import (
	"os"
	"os/exec"
)

// Tuggy runs tuggy.
func Tuggy(args ...string) error {
	cmd := exec.Command("tuggy")
	cmd.Args = append(cmd.Args, args...)
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
