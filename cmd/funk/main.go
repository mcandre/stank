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

// Funk holds configuration for a funky walk.
type Funk struct {
	FoundOdor bool
}

// CheckShebangs analyzes POSIXy scripts for some shebang oddities. If an oddity is found, CheckShebangs prints a warning and returns true.
// Otherwise, CheckShebangs returns false.
func CheckShebangs(smell stank.Smell) bool {
	// Empty extension and .sh are valid for POSIX scripts.
	if smell.Extension == "" || smell.Extension == ".sh" {
		return false
	}

	// Shebangs are ill advised for configuration files.
	if stank.LOWEREXTENSIONS2CONFIG[strings.ToLower(smell.Extension)] || stank.LOWERFILENAMES2CONFIG[strings.ToLower(smell.Filename)] {
		if smell.Shebang != "" {
			fmt.Printf("Configuration features shebang: %s\n", smell.Path)
			return true
		} else {
			return false
		}
	}

	if smell.Shebang == "" {
		fmt.Printf("Missing shebang: %s\n", smell.Path)
		return true
	}

	extensionSansDot := smell.Extension[1:]

	// .bash is valid for bash4 scripts.
	if smell.Interpreter == "bash4" && extensionSansDot == "bash" {
		return false
	}

	// Mismatched shebangs and extensions result in a script being sent to the wrong parser depending on whether it is loaded as `<interpreter> <path>` vs. `./<path>`.
	if smell.Interpreter != extensionSansDot {
		fmt.Printf("Interpreter mismatch between shebang and extension: %s\n", smell.Path)
		return true
	}

	return false
}

// CheckExecutableBits analyzes POSIXy scripts for some file permission oddities. If an oddity is found, CheckPermissions prints a warning and returns true.
// Otherwise, CheckPermissions returns false.
func CheckPermissions(smell stank.Smell) bool {
	if smell.Permissions&0100 == 0 && smell.Permissions&0010 == 0 && smell.Permissions&0001 == 0 {
		return false
	}

	if stank.LOWEREXTENSIONS2CONFIG[strings.ToLower(smell.Extension)] || stank.LOWERFILENAMES2CONFIG[strings.ToLower(smell.Filename)] {
		fmt.Printf("Configuration features executable permissions: %s\n", smell.Path)
		return true
	}

	return false
}

// FunkyCheck analyzes POSIXy scripts for some oddities. If an oddity is found, FunkyCheck prints a warning and returns true.
// Otherwise, FunkyCheck returns false.
func FunkyCheck(smell stank.Smell) bool {
	res1 := CheckShebangs(smell)
	res2 := CheckPermissions(smell)

	return res1 || res2
}

func (o *Funk) Walk(pth string, info os.FileInfo, err error) error {
	smell, err := stank.Sniff(pth)

	if err != nil && err != io.EOF {
		log.Print(err)
	}

	if smell.POSIXy {
		if !FunkyCheck(smell) {
			o.FoundOdor = true
		}
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

	funk := Funk{}

	for _, pth := range paths {
		filepath.Walk(pth, funk.Walk)
	}

	if funk.FoundOdor {
		os.Exit(1)
	}
}
