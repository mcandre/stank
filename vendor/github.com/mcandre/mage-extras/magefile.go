//go:build mage

package main

import (
	"fmt"
	"os"

	"github.com/magefile/mage/mg"
	mageextras "github.com/mcandre/mage-extras"
)

// Default references the default build task.
var Default = CoverageHTML

// CoverHTML denotes the HTML formatted coverage filename.
var CoverHTML = "cover.html"

// CoverProfile denotes the raw coverage data filename.
var CoverProfile = "cover.out"

// Govulncheck runs govulncheck.
func Govulncheck() error { return mageextras.Govulncheck("-scan", "package", "./...") }

// Snyk runs Snyk SCA.
func Snyk() error { return mageextras.SnykTest() }

// Audit runs a security audit.
func Audit() error {
	mg.Deps(Govulncheck)
	return Snyk()
}

// CoverageHTML generates HTML formatted coverage data.
func CoverageHTML() error {
	mg.Deps(CoverageProfile)
	return mageextras.CoverageHTML(CoverHTML, CoverProfile)
}

// CoverageProfile generates raw coverage data.
func CoverageProfile() error { return mageextras.CoverageProfile(CoverProfile) }

// Test executes the unit test suite.
func Test() error { return mageextras.UnitTest() }

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
func Revive() error { return mageextras.Revive("-set_exit_status") }

// Shadow runs go vet with shadow checks enabled.
func Shadow() error { return mageextras.GoVetShadow() }

// Staticcheck runs staticcheck.
func Staticcheck() error { return mageextras.Staticcheck("./...") }

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

// NoVendor lists non-vendored Go source files.
func NoVendor() error {
	mg.Deps(mageextras.CollectGoFiles)

	for pth, _ := range mageextras.CollectedGoFiles {
		fmt.Println(pth)
	}

	return nil
}

// CleanCoverage deletes coverage data.
func CleanCoverage() error {
	if err := os.RemoveAll(CoverHTML); err != nil {
		return err
	}

	return os.RemoveAll(CoverProfile)
}

// Clean deletes build artifacts.
func Clean() error { mg.Deps(CleanCoverage); return nil }
