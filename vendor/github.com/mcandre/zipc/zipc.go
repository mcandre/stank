// Package zipc provides utilities for ZIP archiving files with a specified current working directory.
package zipc

import (
	"log"
	"path"

	"github.com/jhoonb/archivex"
)

// Version is semver.
const Version = "0.0.4"

// Compress creates a compressed archive.
//
// archivePath specifies the target archive
// paths specifies the source files
// root specifies an optional top-level directory for the source files
//
// If root is blank (""), then path resolution is informed by the current working directory.
// Otherwise, path resolution is informed by root.
func Compress(archivePath string, paths []string, root string) error {
	if root != "" {
		for i, pth := range paths {
			paths[i] = path.Join(root, pth)
		}
	}

	archive := new(archivex.ZipFile)

	if err := archive.Create(archivePath); err != nil {
		return err
	}

	for _, source := range paths {
		if err := archive.AddAll(source, true); err != nil {
			log.Panic(err)
		}
	}

	return archive.Close()
}
