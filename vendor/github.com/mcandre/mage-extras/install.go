package mageextras

import (
	"os"
	"os/exec"
	"path"

	"github.com/magefile/mage/mg"
)

// Install builds and installs Go applications.
func Install(args ...string) error {
	cmd := exec.Command("go")
	cmd.Args = append(cmd.Args, "install")
	cmd.Args = append(cmd.Args, args...)
	cmd.Args = append(cmd.Args, AllPackagesPath)
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Uninstall deletes installed Go applications.
func Uninstall(applications ...string) error {
	mg.Deps(LoadGoBinariesPath)

	for _, application := range applications {
		if err := os.RemoveAll(path.Join(LoadedGoBinariesPath, application)); err != nil {
			return err
		}
	}

	return nil
}
