package mageextras

import (
	"os"
	"os/exec"
)

// Compile runs go build recursively.
func Compile(args ...string) error {
	cmdName := "go"

	cmdParameters := []string{cmdName}
	cmdParameters = append(cmdParameters, "build")
	cmdParameters = append(cmdParameters, args...)
	cmdParameters = append(cmdParameters, "./...")

	cmd := exec.Command(cmdName)
	cmd.Args = cmdParameters
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
