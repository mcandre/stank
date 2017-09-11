package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/mcandre/stank"
)

var flagHelp = flag.Bool("help", false, "Show usage information")
var flagVersion = flag.Bool("version", false, "Show version information")

// Rose holds configuration for a rosy walk.
type Rose struct {
	FoundPOSIXy bool
}

// Walk is a callback for filepath.Walk to scan for shell scripts.
func (o *Rose) Walk(pth string, info os.FileInfo, err error) error {
	smell, err := stank.Sniff(pth, stank.SniffConfig{})

	if err != nil && err != io.EOF {
		log.Print(err)
	}

	if smell.POSIXy && !(stank.LOWEREXTENSIONS2CONFIG[strings.ToLower(smell.Extension)] || stank.LOWERFILENAMES2CONFIG[strings.ToLower(smell.Filename)]) {
		fmt.Printf("Rewrite POSIX script in Ruby or other safer general purpose scripting language: %s\n", pth)
		o.FoundPOSIXy = true
	}

	return nil
}

func main() {
	flag.Parse()

	switch {
	case *flagVersion:
		fmt.Println(stank.Version)
		os.Exit(0)
	case *flagHelp:
		flag.PrintDefaults()
		os.Exit(0)
	}

	paths := flag.Args()

	rose := Rose{}

	for _, pth := range paths {
		filepath.Walk(pth, rose.Walk)
	}

	if rose.FoundPOSIXy {
		os.Exit(1)
	}
}
