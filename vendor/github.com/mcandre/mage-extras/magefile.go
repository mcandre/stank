// +build mage

package main

import (
	"fmt"
	"os"

	"github.com/magefile/mage/mg"
	"github.com/mcandre/mage-extras"
)

// Default references the default build task.
var Default = CoverageHTML

// CoverHTML denotes the HTML formatted coverage filename.
var CoverHTML = "cover.html"

// CoverProfile denotes the raw coverage data filename.
var CoverProfile = "cover.out"

// CoverageHTML generates HTML formatted coverage data.
func CoverageHTML() error { mg.Deps(CoverageProfile); return mageextras.CoverageHTML(CoverHTML, CoverProfile) }

// CoverageProfile generates raw coverage data.
func CoverageProfile() error { return mageextras.CoverageProfile(CoverProfile) }

// Test executes the unit test suite.
func Test() error { return mageextras.UnitTest() }

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
