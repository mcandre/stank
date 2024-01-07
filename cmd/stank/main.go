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

var flagSh = flag.Bool("sh", false, "Limit results to specifically bare POSIX sh scripts")
var flagAlt = flag.Bool("alt", false, "Limit results to specifically alternative, non-POSIX lowlevel shell scripts")
var flagExcludeInterpreters = flag.String("exInterp", "", "Remove results with the given interpreter(s) (Comma separated)")
var flagHelp = flag.Bool("help", false, "Show usage information")
var flagVersion = flag.Bool("version", false, "Show version information")

// StankMode controls stank rule behavior.
type StankMode int

const (
	// ModePOSIXy matches POSIX-like shell scripts.
	ModePOSIXy StankMode = iota

	// ModePureSh matches specifically sh-interpreted scripts.
	ModePureSh

	// ModeAltShellScript matches certain non-POSIX shell scripts.
	ModeAltShellScript
)

// Stanker holds configuration for a stanky walk
type Stanker struct {
	// Mode is scan type.
	Mode StankMode

	// InterpreterExclusions remove results from scan report.
	InterpreterExclusions []string
}

// Walk sniffs a file system node for POSIXyness.
// If the file smells sufficiently POSIXy, the path is printed.
// Otherwise, the path is omitted.
func (o Stanker) Walk(pth string, info os.FileInfo, err error) error {
	smell, err := stank.Sniff(pth, stank.SniffConfig{})

	if err != nil && err != io.EOF {
		log.Print(err)
	}

	if smell.MachineGenerated {
		return nil
	}

	for _, interpreterExclusion := range o.InterpreterExclusions {
		if smell.Interpreter == interpreterExclusion {
			return nil
		}
	}

	switch o.Mode {
	case ModePureSh:
		if smell.POSIXy && (smell.Interpreter == "sh" || smell.Interpreter == "generic-sh") {
			fmt.Println(smell.Path)
		}
	case ModeAltShellScript:
		if smell.AltShellScript {
			fmt.Println(smell.Path)
		}
	default:
		if smell.POSIXy {
			fmt.Println(smell.Path)
		}
	}

	return nil
}

func main() {
	flag.Parse()

	stanker := Stanker{Mode: ModePOSIXy}

	if *flagSh {
		stanker.Mode = ModePureSh
	}

	if *flagAlt {
		stanker.Mode = ModeAltShellScript
	}

	stanker.InterpreterExclusions = strings.Split(*flagExcludeInterpreters, ",")

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
