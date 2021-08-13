package mageextras

import (
	"os"
	"os/user"
	"path"
)

// LoadedGoBinariesPath denotes the path to the Go binaries directory.
// Populated with LoadGoBinariesPath().
var LoadedGoBinariesPath = ""

// LoadGoBinariesPath populates LoadedGoBinariesPath.
func LoadGoBinariesPath() error {
	goPath := os.Getenv("GOPATH")

	if goPath == "" {
		user, err := user.Current()

		if err != nil {
			return err
		}

		goPath = path.Join(user.HomeDir, "go")
	}

	LoadedGoBinariesPath = path.Join(goPath, "bin")

	return nil
}
