package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/mcandre/stank"
)

var flagHelp = flag.Bool("help", false, "Show usage information")
var flagVersion = flag.Bool("version", false, "Show version information")

// StankWalk sniffs a file system node for POSIXyness.
// If the file smells sufficiently POSIXy, the path is printed.
// Otherwise, the path is omitted.
func StankWalk(pth string, info os.FileInfo, err error) error {
	smell, err := stank.Sniff(pth)
	if err != nil {
		log.Print(err)
	}

	if smell.POSIXy {
		fmt.Println(smell.Path)
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

	var observedError bool
	var err error

	for _, pth := range paths {
		err = filepath.Walk(pth, StankWalk)

		if err != nil {
			log.Print(err)
			observedError = true
		}
	}

	if observedError {
		os.Exit(1)
	}
}
