package mageextras

import (
	"os"
	"os/exec"

	"github.com/magefile/mage/mg"
)

// Nakedret runs nakedret.
func Nakedret(args ...string) error {
	mg.Deps(CollectGoFiles)

	for pth := range CollectedGoFiles {
		cmd := exec.Command("nakedret")
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
