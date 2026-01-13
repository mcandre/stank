package mageextras

import (
	"os"
	"os/exec"
)

// Revive runs revive.
func Revive(args ...string) error {
	cmd := exec.Command("revive")
	cmd.Args = append(cmd.Args, "-exclude", "vendor/...")
	cmd.Args = append(cmd.Args, args...)
	cmd.Args = append(cmd.Args, "./...")
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
