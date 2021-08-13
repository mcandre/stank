package mageextras

import (
	"os"
	"os/exec"
)

// Goxcart cross-compiles Go binaries with additional targets enabled.
func Goxcart(outputPath string, args ...string) error {
	if err := os.MkdirAll(outputPath, os.ModeDir|0775); err != nil {
		return err
	}

	var goxcartParts []string
	goxcartParts = append(goxcartParts, args...)
	goxcartParts = append(goxcartParts, "-output")
	goxcartParts = append(goxcartParts, outputPath)
	goxcartParts = append(goxcartParts, "-commands")
	goxcartParts = append(goxcartParts, AllCommandsPath)

	cmd := exec.Command(
		"goxcart",
		goxcartParts...,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
