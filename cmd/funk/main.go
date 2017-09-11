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

var flagEOL = flag.Bool("eol", true, "Report presence/absence of final end of line sequence")
var flagCR = flag.Bool("cr", true, "Report presence/absence of final end of line sequence")
var flagHelp = flag.Bool("help", false, "Show usage information")
var flagVersion = flag.Bool("version", false, "Show version information")

// Funk holds configuration for a funky walk.
type Funk struct {
	EOLCheck  bool
	CRCheck   bool
	FoundOdor bool
}

// CheckEOL analyzes POSIXy scripts for the presence/absence of a final end of line sequence such as \n at the end of a file, \r\n, etc.
func CheckEOL(smell stank.Smell) bool {
	if !smell.FinalEOL {
		fmt.Printf("Missing final end of line sequence: %s\n", smell.Path)
		return true
	}

	return false
}

// CheckCR analyzes POSIXy scripts for the presence/absence of a CR/CRLF line ending sequence.
func CheckCR(smell stank.Smell) bool {
	if smell.ContainsCR {
		fmt.Printf("NonPOSIX CR/CRLF line ending detected: %s\n", smell.Path)
		return true
	}

	return false
}

// CheckBOMs analyzes POSIXy scripts for byte order markers. If a BOM is found, CheckBOMs prints a warning and returns true.
// Otherwise, CheckBOMs returns false.
func CheckBOMs(smell stank.Smell) bool {
	if smell.BOM {
		fmt.Printf("Leading BOM reduces portability: %s\n", smell.Path)

		return true
	}

	return false
}

// CheckShebangs analyzes POSIXy scripts for some shebang oddities. If an oddity is found, CheckShebangs prints a warning and returns true.
// Otherwise, CheckShebangs returns false.
func CheckShebangs(smell stank.Smell) bool {
	// Empty extension and .sh are valid for POSIX scripts.
	// .envrc is also common for direnv-triggered shell scripts.
	if smell.Extension == "" || smell.Extension == ".sh" || smell.Extension == ".envrc" {
		return false
	}

	// Shebangs are ill advised for configuration files.
	if stank.LOWEREXTENSIONS2CONFIG[strings.ToLower(smell.Extension)] || stank.LOWERFILENAMES2CONFIG[strings.ToLower(smell.Filename)] {
		if smell.Shebang != "" {
			fmt.Printf("Configuration features shebang: %s\n", smell.Path)
			return true
		}

		return false
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

	// .ksh is valid for ksh derivatives, even the nonPOSIX lksh.
	if strings.Contains(smell.Interpreter, "ksh") && extensionSansDot == "ksh" {
		return false
	}

	// Mismatched shebangs and extensions result in a script being sent to the wrong parser depending on whether it is loaded as `<interpreter> <path>` vs. `./<path>`.
	if smell.Interpreter != extensionSansDot {
		fmt.Printf("Interpreter mismatch between shebang and extension: %s\n", smell.Path)
		return true
	}

	return false
}

// CheckPermissions analyzes POSIXy scripts for some file permission oddities. If an oddity is found, CheckPermissions prints a warning and returns true.
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

// CheckModulino warns when a smell features some aspects of an application, such as executable bits, and simultaneously some aspects of a library, such as a non-empty file extension.
// If the file is a pure application or library, CheckModulino returns false.
// Otherwise, CheckModulino returns true.
func CheckModulino(smell stank.Smell) bool {
	if stank.LOWEREXTENSIONS2CONFIG[strings.ToLower(smell.Extension)] || stank.LOWERFILENAMES2CONFIG[strings.ToLower(smell.Filename)] {
		return false
	}

	if (smell.Extension == "" && !smell.OwnerExecutable) || (smell.Extension != "" && (smell.Permissions&0100 != 0 || smell.Permissions&0010 != 0 || smell.Permissions&0001 != 0)) {
		fmt.Printf("Modulino ambiguity, either add owner executable permission or else rename with a .lib.sh extension: %s\n", smell.Path)
		return true
	}

	return false
}

// FunkyCheck analyzes POSIXy scripts for some oddities. If an oddity is found, FunkyCheck prints a warning and returns true.
// Otherwise, FunkyCheck returns false.
func (o Funk) FunkyCheck(smell stank.Smell) bool {
	var resEOL bool

	if o.EOLCheck {
		resEOL = CheckEOL(smell)
	}

	var resCR bool

	if o.CRCheck {
		resCR = CheckCR(smell)
	}

	resBOM := CheckBOMs(smell)
	resShebang := CheckShebangs(smell)
	resPerms := CheckPermissions(smell)
	resModulino := CheckModulino(smell)

	return resEOL || resCR || resBOM || resShebang || resPerms || resModulino
}

// Walk is a callback for filepath.Walk to lint shell scripts.
func (o *Funk) Walk(pth string, info os.FileInfo, err error) error {
	smell, err := stank.Sniff(pth, stank.SniffConfig{EOLCheck: o.EOLCheck, CRCheck: o.CRCheck})

	if err != nil && err != io.EOF {
		log.Print(err)
	}

	if smell.POSIXy {
		if o.FunkyCheck(smell) {
			o.FoundOdor = true
		}
	}

	return nil
}

func main() {
	flag.Parse()

	funk := Funk{}

	if *flagEOL {
		funk.EOLCheck = true
	}

	if *flagCR {
		funk.CRCheck = true
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

	for _, pth := range paths {
		filepath.Walk(pth, funk.Walk)
	}

	if funk.FoundOdor {
		os.Exit(1)
	}
}
