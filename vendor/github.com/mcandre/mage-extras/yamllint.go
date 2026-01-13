package mageextras

import (
	"os"
	"os/exec"
)

// Yamllint runs yamllint.
func Yamllint(args ...string) error {
	cmd := exec.Command("yamllint")
	cmd.Args = append(cmd.Args, args...)
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
