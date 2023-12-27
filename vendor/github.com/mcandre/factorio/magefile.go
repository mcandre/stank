//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/magefile/mage/mg"
	"github.com/mcandre/factorio"
	mageextras "github.com/mcandre/mage-extras"
)

// artifactsPath describes where artifacts are produced.
var artifactsPath = "bin"

// Default references the default build task.
var Default = Test

// Audit runs a security audit.
func Audit() error { return mageextras.SnykTest() }

// GoVet runs go vet with shadow checks enabled.
func GoVet() error { return mageextras.GoVetShadow() }

// Gofmt runs gofmt.
func GoFmt() error { return mageextras.GoFmt("-s", "-w") }

// GoImports runs goimports.
func GoImports() error { return mageextras.GoImports("-w") }

// Errcheck runs errcheck.
func Errcheck() error { return mageextras.Errcheck("-blank") }

// Nakedret runs nakedret.
func Nakedret() error { return mageextras.Nakedret("-l", "0") }

// Staticcheck runs staticcheck.
func Staticcheck() error { return mageextras.Staticcheck() }

// Unmake runs unmake.
func Unmake() error { return mageextras.Unmake(".") }

// Lint runs the lint suite.
func Lint() error {
	mg.Deps(GoVet)
	mg.Deps(GoFmt)
	mg.Deps(GoImports)
	mg.Deps(Errcheck)
	mg.Deps(Nakedret)
	mg.Deps(Staticcheck)
	mg.Deps(Unmake)
	return nil
}

// portBasename labels the artifact basename.
var portBasename = fmt.Sprintf("factorio-%s", factorio.Version)

// repoNamespace identifies the Go namespace for this project.
var repoNamespace = "github.com/mcandre/factorio"

// Factorio cross-compiles Go binaries for a multitude of platforms.
func Factorio() error { return mageextras.Factorio(portBasename) }

// Port builds and compresses artifacts.
func Port() error {
	cmd := exec.Command("factorio")
	cmd.Env = append(
		cmd.Environ(),
		fmt.Sprintf("FACTORIO_BANNER=factorio-%v", factorio.Version),
	)

	if err := cmd.Run(); err != nil {
		return err
	}

	return mageextras.Archive(portBasename, artifactsPath)
}

// Test runs a test suite.
func Test() error { mg.Deps(Port); return nil }

// Install builds and installs Go applications.
func Install() error { return mageextras.Install() }

// Uninstall deletes installed Go applications.
func Uninstall() error { return mageextras.Uninstall("factorio") }

// Clean deletes artifacts.
func Clean() error { return os.RemoveAll(artifactsPath) }
