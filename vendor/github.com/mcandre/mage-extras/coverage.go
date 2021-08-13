package mageextras

import (
	"fmt"
	"os"
	"os/exec"
)

// CoverageHTML generates HTML formatted coverage data.
func CoverageHTML(htmlFilename string, profileFilename string) error {
	cmd := exec.Command(
		"go",
		"tool",
		"cover",
		fmt.Sprintf("-html=%s", profileFilename),
		"-o",
		htmlFilename,
	)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// CoverageProfile generates raw coverage data.
func CoverageProfile(profileFilename string) error {
	cmd := exec.Command(
		"go",
		"test",
		fmt.Sprintf("-coverprofile=%s", profileFilename),
	)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
