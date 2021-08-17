package mageextras

import (
	"fmt"
	"path"

	"github.com/mcandre/zipc"
)

// Archive compresses build artifacts.
func Archive(portBasename string, artifactsPath string) error {
	archiveFilename := fmt.Sprintf("%s.zip", portBasename)
	return zipc.Compress(path.Join("bin", archiveFilename), []string{portBasename}, artifactsPath)
}
