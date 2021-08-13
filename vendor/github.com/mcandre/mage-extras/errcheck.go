package mageextras

import (
	"os"
	"os/exec"
)

// Errcheck runs errcheck.
func Errcheck(args ...string) error {
	cmdName := "errcheck"

	cmdParameters := []string{cmdName}
	cmdParameters = append(cmdParameters, args...)

	cmd := exec.Command(cmdName)
	cmd.Args = cmdParameters
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
