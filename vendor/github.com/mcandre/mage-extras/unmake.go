package mageextras

import (
	"os"
	"os/exec"
)

// Unmake runs unmake.
func Unmake(args ...string) error {
	cmd := exec.Command("unmake")
	cmd.Args = append(cmd.Args, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
