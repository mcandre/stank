package mageextras

import (
	"os"
	"os/exec"
)

// Xgo cross-compiles (c)Go binaries with additional targets enabled.
func Xgo(outputPath string, args ...string) error {
	if err := os.MkdirAll(outputPath, os.ModeDir|0775); err != nil {
		return err
	}

	var xgoParts []string
	xgoParts = append(xgoParts, "-dest")
	xgoParts = append(xgoParts, outputPath)
	xgoParts = append(xgoParts, args...)

	cmd := exec.Command(
		"xgo",
		xgoParts...,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
