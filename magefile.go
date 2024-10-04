//go:build mage
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
	mageextras "github.com/mcandre/mage-extras"
	"github.com/mcandre/stank"
)

// artifactsPath describes where artifacts are produced.
var artifactsPath = "bin"

// Default references the default build task.
var Default = Test

// Govulncheck runs govulncheck.
func Govulncheck() error { return mageextras.Govulncheck("-scan", "package", "./...") }

// SnykTest runs Snyk SCA.
func Snyk() error { return mageextras.SnykTest() }

// Audit runs a security audit.
func Audit() error {
	mg.Deps(Govulncheck)
	return Snyk()
}

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

	cmdFunk := exec.Command("funk", examplesDir)
	cmdFunk.Stdout = os.Stdout
	cmdFunk.Stderr = os.Stderr

	if err := cmdFunk.Run(); err == nil {
		return errors.New("Expected non-zero exit status from funk")
	}

	return nil
}

// Test runs unit and integration tests.
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

// Deadcode runs deadcode.
func Deadcode() error { return mageextras.Deadcode("./...") }

// Gofmt runs gofmt.
func GoFmt() error { return mageextras.GoFmt("-s", "-w") }

// GoImports runs goimports.
func GoImports() error { return mageextras.GoImports("-w") }

// GoVet runs default go vet analyzers.
func GoVet() error { return mageextras.GoVet() }

// Errcheck runs errcheck.
func Errcheck() error { return mageextras.Errcheck("-blank") }

// Nakedret runs nakedret.
func Nakedret() error { return mageextras.Nakedret("-l", "0") }

// Revive runs revive.
func Revive() error { return mageextras.Revive() }

// Shadow runs go vet with shadow checks enabled.
func Shadow() error { return mageextras.GoVetShadow() }

// Staticcheck runs staticcheck.
func Staticcheck() error { return mageextras.Staticcheck() }

// Unmake runs unmake.
func Unmake() error {
	err := mageextras.Unmake(".")

	if err != nil {
		return err
	}

	return mageextras.Unmake("-n", ".")
}

// Lint runs the lint suite.
func Lint() error {
	mg.Deps(Deadcode)
	mg.Deps(GoFmt)
	mg.Deps(GoImports)
	mg.Deps(GoVet)
	mg.Deps(Errcheck)
	mg.Deps(Nakedret)
	mg.Deps(Revive)
	mg.Deps(Shadow)
	mg.Deps(Staticcheck)
	mg.Deps(Unmake)
	return nil
}

// portBasename labels the artifact basename.
var portBasename = fmt.Sprintf("stank-%s", stank.Version)

// repoNamespace identifies the Go namespace for this project.
var repoNamespace = "github.com/mcandre/stank"

// Factorio cross-compiles Go binaries for a multitude of platforms.
func Factorio() error { return mageextras.Factorio(portBasename) }

// Port builds and compresses artifacts.
func Port() error { mg.Deps(Factorio); return mageextras.Archive(portBasename, artifactsPath) }

// Install builds and installs Go applications.
func Install() error { return mageextras.Install() }

// Uninstall deletes installed Go applications.
func Uninstall() error { return mageextras.Uninstall("stink", "stank", "funk") }

// CleanCoverage deletes coverage data.
func CleanCoverage() error {
	if err := os.RemoveAll(CoverHTML); err != nil {
		return err
	}

	return os.RemoveAll(CoverProfile)
}

// Clean deletes artifacts.
func Clean() error { mg.Deps(CleanCoverage); return os.RemoveAll(artifactsPath) }
