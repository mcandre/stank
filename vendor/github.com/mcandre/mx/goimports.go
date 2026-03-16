package mx

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// GoImports runs goimports.
func GoImports(args ...string) error {
	mg.Deps(CollectGoFiles)

	for pth := range CollectedGoFiles {
		var as []string
		as = append(as, args...)
		as = append(as, pth)

		if err := sh.RunV("goimports", as...); err != nil {
			return err
		}
	}

	return nil
}
