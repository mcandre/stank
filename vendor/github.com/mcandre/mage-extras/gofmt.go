package mageextras

import (
	"os"
	"os/exec"

	"github.com/magefile/mage/mg"
)

// GoFmt runs gofmt.
func GoFmt(args ...string) error {
	mg.Deps(CollectGoFiles)

	for pth := range CollectedGoFiles {
		cmd := exec.Command("gofmt")
		cmd.Args = append(cmd.Args, args...)
		cmd.Args = append(cmd.Args, pth)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return err
		}
	}

	return nil
}
