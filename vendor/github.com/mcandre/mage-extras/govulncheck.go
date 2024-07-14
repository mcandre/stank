package mageextras

import (
	"os"
	"os/exec"
)

// Govulncheck runs govulncheck.
func Govulncheck(args ...string) error {
	cmd := exec.Command("govulncheck")
	cmd.Args = append(cmd.Args, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
