//go:build mage

package main

import (
	"fmt"
	"os"

	"github.com/magefile/mage/mg"
	"github.com/mcandre/factorio"
	mageextras "github.com/mcandre/mage-extras"
)

// artifactsPath describes where artifacts are produced.
var artifactsPath = "bin"

// Default references the default build task.
var Default = Test

// Govulncheck runs govulncheck.
func Govulncheck() error { return mageextras.Govulncheck("-scan", "package", "./...") }

// Audit runs a security audit.
func Audit() error { return Govulncheck() }

// Deadcode runs deadcode.
func Deadcode() error { return mageextras.Deadcode("./...") }

// DockerBuild creates local Docker buildx images.
func DockerBuild() error {
	return mageextras.Tuggy(
		"-t", fmt.Sprintf("n4jm4/factorio:%s", factorio.Version),
		"--load",
	)
}

// DockerPush creates and tag aliases remote Docker buildx images.
func DockerPush() error {
	return mageextras.Tuggy(
		"-t", fmt.Sprintf("n4jm4/factorio:%s", factorio.Version),
		"-a", "n4jm4/factorio",
		"--push",
	)
}

// DockerTest creates and tag aliases remote test Docker buildx images.
func DockerTest() error {
	return mageextras.Tuggy(
		"-t", "n4jm4/factorio:test",
		"--load",
		"--push",
	)
}

// GoImports runs goimports.
func GoImports() error { return mageextras.GoImports("-w") }

// GoVet runs default go vet analyzers.
func GoVet() error { return mageextras.GoVet() }

// Errcheck runs errcheck.
func Errcheck() error { return mageextras.Errcheck("-blank") }

// Nakedret runs nakedret.
func Nakedret() error { return mageextras.Nakedret("-l", "0") }

// Shadow runs go vet with shadow checks enabled.
func Shadow() error { return mageextras.GoVetShadow() }

// Staticcheck runs staticcheck.
func Staticcheck() error { return mageextras.Staticcheck("./...") }

// Lint runs the lint suite.
func Lint() error {
	mg.Deps(Deadcode)
	mg.Deps(GoImports)
	mg.Deps(GoVet)
	mg.Deps(Errcheck)
	mg.Deps(Nakedret)
	mg.Deps(Shadow)
	mg.Deps(Staticcheck)
	return nil
}

// portBasename labels the artifact basename.
var portBasename = fmt.Sprintf("factorio-%s", factorio.Version)

// repoNamespace identifies the Go namespace for this project.
var repoNamespace = "github.com/mcandre/factorio"

// Factorio cross-compiles Go binaries for a multitude of platforms.
func Factorio() error { mg.Deps(Install); return mageextras.Factorio(portBasename) }

// Port builds and compresses artifacts.
func Port() error {
	mg.Deps(Factorio);

	return mageextras.Chandler(
		"-C",
		artifactsPath,
		"-czf",
		fmt.Sprintf("%s.tgz", portBasename),
		portBasename,
	)
}

// Test runs a test suite.
func Test() error { return mageextras.UnitTest() }

// Install builds and installs Go applications.
func Install() error { return mageextras.Install() }

// Uninstall deletes installed Go applications.
func Uninstall() error { return mageextras.Uninstall("factorio") }

// Clean deletes artifacts.
func Clean() error { return os.RemoveAll(artifactsPath) }
