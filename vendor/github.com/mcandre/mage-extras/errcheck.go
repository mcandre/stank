package mageextras

import (
	"os"
	"os/exec"
)

// Errcheck runs errcheck.
func Errcheck(args ...string) error {
	cmd := exec.Command("errcheck")
	cmd.Args = append(cmd.Args, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
