package mageextras

import (
	"os"
	"os/exec"
)

// Chandler runs chandler.
func Chandler(args ...string) error {
	cmd := exec.Command("chandler")
	cmd.Args = append(cmd.Args, args...)
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
