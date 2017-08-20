package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/mcandre/stank"
)

var flagHelp = flag.Bool("help", false, "Show usage information")
var flagVersion = flag.Bool("version", false, "Show version information")

// Rose holds configuration for a rosy walk.
type Rose struct {
	FoundPOSIXy bool
}

func (o Rose) Walk(pth string, info os.FileInfo, err error) error {
	smell, err := stank.Sniff(pth)

	if err != nil && err != io.EOF {
		log.Print(err)
	}

	if smell.POSIXy {
		fmt.Printf("Rewrite POSIX script in Ruby or other safer general purpose scripting language: %s\n", pth)
		o.FoundPOSIXy = true
	}

	return err
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
