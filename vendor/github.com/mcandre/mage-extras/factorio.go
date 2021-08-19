package mageextras

import (
	"os"

	"github.com/mcandre/factorio"
)

// Factorio cross-compiles Go binaries for a multitude of platforms.
func Factorio(banner string, args ...string) error {
	if err := os.Setenv("FACTORIO_BANNER", banner); err != nil {
		return err
	}

	return factorio.Port(args)
}
