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

var flagSh = flag.Bool("sh", false, "Limit results to specifically bare POSIX sh scripts")
var flagHelp = flag.Bool("help", false, "Show usage information")
var flagVersion = flag.Bool("version", false, "Show version information")

// Stanker holds configuration for a stanky walk
type Stanker struct {
	PureSh bool
}

// StankWalk sniffs a file system node for POSIXyness.
// If the file smells sufficiently POSIXy, the path is printed.
// Otherwise, the path is omitted.
func (o Stanker) Walk(pth string, info os.FileInfo, err error) error {
	smell, err := stank.Sniff(pth, false)
	if err != nil && err != io.EOF {
		log.Print(err)
	}

	if smell.POSIXy && (!o.PureSh || smell.Interpreter == "sh" || smell.Interpreter == "generic-sh") {
		fmt.Println(smell.Path)
	}

	return nil
}

func main() {
	flag.Parse()

	stanker := Stanker{}

	if *flagSh {
		stanker.PureSh = true
	}

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
		err = filepath.Walk(pth, stanker.Walk)

		if err != nil {
			log.Print(err)
			observedError = true
		}
	}

	if observedError {
		os.Exit(1)
	}
}
