package mageextras

import (
	"fmt"
	"os"
	"os/exec"
)

// Archive compresses build artifacts.
func Archive(portBasename string, artifactsPath string) error {
	archiveFilename := fmt.Sprintf("%s.tgz", portBasename)
	cmd := exec.Command("tar")
	cmd.Args = append(cmd.Args, "czf", archiveFilename, portBasename)
	cmd.Dir = artifactsPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
