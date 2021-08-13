package mageextras

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Gox cross-compiles Go binaries.
func Gox(outputPath string, artifactStructure string) error {
	if err := os.MkdirAll(outputPath, os.ModeDir|0775); err != nil {
		return err
	}

	artifactStructureAnchored := strings.Join(
		[]string{outputPath, artifactStructure},
		PathSeparatorString,
	)

	cmd := exec.Command(
		"gox",
		fmt.Sprintf("-output=%s", artifactStructureAnchored),
		AllCommandsPath,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
