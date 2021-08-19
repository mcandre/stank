//go:build mage
// +build mage

package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/magefile/mage/mg"
	"github.com/mcandre/mage-extras"
	"github.com/mcandre/stank"
)

// artifactsPath describes where artifacts are produced.
var artifactsPath = "bin"

// Default references the default build task.
var Default = Test

// UnitTests runs the unit test suite.
func UnitTest() error { return mageextras.UnitTest() }

// IntegrationTest executes the integration test suite.
func IntegrationTest() error {
	mg.Deps(Install)

	examplesDir := "examples"

	var stinkOut bytes.Buffer

	cmdStink := exec.Command("stink", path.Join(examplesDir, "hello.sh"))
	cmdStink.Stdout = bufio.NewWriter(&stinkOut)
	cmdStink.Stderr = os.Stderr

	if err := cmdStink.Run(); err != nil {
		return err
	}

	stinkOutString := stinkOut.String()

	if !strings.Contains(stinkOutString, "\"POSIXy\":true") {
		return fmt.Errorf("Expected stink output to treat hello.sh as POSIXy: true, got %s\n", stinkOutString)
	}

	cmdStank := exec.Command("stank", examplesDir)
	cmdStank.Stdout = os.Stdout
	cmdStank.Stderr = os.Stderr

	if err := cmdStank.Run(); err != nil {
		return err
	}

	cmdRosy := exec.Command("rosy", "-kame", examplesDir)
	cmdRosy.Stdout = os.Stdout
	cmdRosy.Stderr = os.Stderr
	err := cmdRosy.Run()

	if err == nil {
		return errors.New("Expected non-zero exit status from rosy")
	}

	cmdFunk := exec.Command("funk", examplesDir)
	cmdFunk.Stdout = os.Stdout
	cmdFunk.Stderr = os.Stderr
	err = cmdFunk.Run()

	if err == nil {
		return errors.New("Expected non-zero exit status from funk")
	}

	// ...

	return nil
}

// Text runs unit and integration tests.
func Test() error { mg.Deps(UnitTest); mg.Deps(IntegrationTest); return nil }

// CoverHTML denotes the HTML formatted coverage filename.
var CoverHTML = "cover.html"

// CoverProfile denotes the raw coverage data filename.
var CoverProfile = "cover.out"

// CoverageHTML generates HTML formatted coverage data.
func CoverageHTML() error {
	mg.Deps(CoverageProfile)
	return mageextras.CoverageHTML(CoverHTML, CoverProfile)
}

// CoverageProfile generates raw coverage data.
func CoverageProfile() error { return mageextras.CoverageProfile(CoverProfile) }

// GoVet runs go vet with shadow checks enabled.
func GoVet() error { return mageextras.GoVetShadow() }

// GoLint runs golint.
func GoLint() error { return mageextras.GoLint() }

// Gofmt runs gofmt.
func GoFmt() error { return mageextras.GoFmt("-s", "-w") }

// GoImports runs goimports.
func GoImports() error { return mageextras.GoImports("-w") }

// Errcheck runs errcheck.
func Errcheck() error { return mageextras.Errcheck("-blank") }

// Nakedret runs nakedret.
func Nakedret() error { return mageextras.Nakedret("-l", "0") }

// Lint runs the lint suite.
func Lint() error {
	mg.Deps(GoVet)
	mg.Deps(GoLint)
	mg.Deps(GoFmt)
	mg.Deps(GoImports)
	mg.Deps(Errcheck)
	mg.Deps(Nakedret)
	return nil
}

// portBasename labels the artifact basename.
var portBasename = fmt.Sprintf("stank-%s", stank.Version)

// repoNamespace identifies the Go namespace for this project.
var repoNamespace = "github.com/mcandre/stank"

// Xgo cross-compiles (c)Go binaries with additional targets enabled.
func Xgo() error {
	artifactsPathDist := path.Join(artifactsPath, portBasename)

	for _, name := range []string{"stink", "stank", "funk", "rosy"} {
		if err := mageextras.Xgo(
			artifactsPathDist,
			fmt.Sprintf("github.com/mcandre/stank/cmd/%v", name),
		); err != nil {
			return err
		}
	}

	return nil
}

// Port builds and compresses artifacts.
func Port() error { mg.Deps(Xgo); return mageextras.Archive(portBasename, artifactsPath) }

// Install builds and installs Go applications.
func Install() error { return mageextras.Install() }

// Uninstall deletes installed Go applications.
func Uninstall() error { return mageextras.Uninstall("stink", "stank", "funk", "rosy") }

// CleanCoverage deletes coverage data.
func CleanCoverage() error {
	if err := os.RemoveAll(CoverHTML); err != nil {
		return err
	}

	return os.RemoveAll(CoverProfile)
}

// Clean deletes artifacts.
func Clean() error { mg.Deps(CleanCoverage); return os.RemoveAll(artifactsPath) }
