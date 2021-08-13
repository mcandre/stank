package mageextras

import (
	"fmt"
	"os"
	"os/exec"
)

// Archive compresses build artifacts.
func Archive(portBasename string, artifactsPath string) error {
	archiveFilename := fmt.Sprintf("%s.zip", portBasename)

	cmd := exec.Command("zipc", "-chdir", artifactsPath, archiveFilename, portBasename)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
