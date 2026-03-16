//go:build mage

package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/mcandre/mx"
)

// Default references the default build task.
var Default = Test

// CoverHTML denotes the HTML formatted coverage filename.
const CoverHTML = "cover.html"

// CoverProfile denotes the raw coverage data filename.
const CoverProfile = "cover.out"

// Audit runs a security audit.
func Audit() error { return Govulncheck() }

// Clean deletes artifacts.
func Clean() error { return CleanCoverage() }

// CleanCoverage deletes coverage data.
func CleanCoverage() error {
	if err := sh.Rm(CoverHTML); err != nil {
		return err
	}

	return sh.Rm(CoverProfile)
}

// CoverageHTML generates HTML formatted coverage data.
func CoverageHTML() error {
	mg.Deps(CoverageProfile)
	return mx.CoverageHTML(CoverHTML, CoverProfile)
}

// CoverageProfile generates raw coverage data.
func CoverageProfile() error { return mx.CoverageProfile(CoverProfile) }

// Deadcode runs deadcode.
func Deadcode() error { return sh.RunV("deadcode", "./...") }

// Errcheck runs errcheck.
func Errcheck() error { return sh.RunV("errcheck", "-blank") }

// GoImports runs goimports.
func GoImports() error { return mx.GoImports("-w") }

// GoVet runs default go vet analyzers.
func GoVet() error { return mx.GoVet() }

// Govulncheck runs govulncheck.
func Govulncheck() error { return sh.RunV("govulncheck", "-scan", "package", "./...") }

// Install builds and installs Go applications.
func Install() error { return mx.Install() }

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

// Nakedret runs nakedret.
func Nakedret() error { return mx.Nakedret("-l", "0") }

// Shadow runs go vet with shadow checks enabled.
func Shadow() error { return mx.GoVetShadow() }

// Staticcheck runs staticcheck.
func Staticcheck() error { return sh.RunV("staticcheck", "./...") }

// Test runs a unit test.
func Test() error { return mx.UnitTest() }

// Uninstall deletes installed Go applications.
func Uninstall() error { return mx.Uninstall("stink", "stank", "funk") }
