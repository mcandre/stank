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

var flagKame = flag.Bool("kame", false, "Recommend faster shells")
var flagUsagi = flag.Bool("usagi", false, "Recommend more robust shells")
var flagAhiru = flag.Bool("ahiru", false, "Recommend sh for portability")
var flagHelp = flag.Bool("help", false, "Show usage information")
var flagVersion = flag.Bool("version", false, "Show version information")

// RoseMode controls rosy rule behavior.
type RoseMode int

const (
	// ModeRose encourages developers to rewrite scripts in different languages.
	ModeRose RoseMode = iota

	// ModeKame recommends faster shells.
	ModeKame

	// ModeUsagi recommends more robust shells.
	ModeUsagi

	// ModeAhiru recommends POSIX sh for portability.
	ModeAhiru
)

// Rose holds configuration for a rosy walk.
type Rose struct {
	Mode         RoseMode
	FoundWarning bool
}

// KAMEINTERPRETERS catalogues relatively more performant POSIX shells.
var KAMEINTERPRETERS = map[string]bool{
	"sh":    true,
	"ksh":   true,
	"mksh":  true,
	"pdksh": true,
	"ksh93": true,
	"ksh88": true,
	"oksh":  true,
	"dash":  true,
	"posh":  true,
}

// USAGIINTERPRETERS catalogues POSIX shells with basic debugging `set` features, such as set -euf.
var USAGIINTERPRETERS = map[string]bool{
	"ksh":   true,
	"mksh":  true,
	"pdksh": true,
	"ksh93": true,
	"ksh88": true,
	"bash":  true,
	"zsh":   true,
}

// CheckShebang warns on POSIXy scripts that lack a shebang line to distinguish which interpreter should be used.
func CheckShebang(smell stank.Smell) bool {
	if smell.Interpreter == "generic-sh" {
		fmt.Printf("Clarify interpreter with a shebang line: %s\n", smell.Path)
		return true
	}

	return false
}

// CheckRose accepts a POSIXy smell and warns the user to rewrite in a safer language.
func CheckRose(smell stank.Smell) bool {
	fmt.Printf("Rewrite POSIX script in Ruby or other safer general purpose scripting language: %s\n", smell.Path)
	return true
}

// CheckKame accepts a POSIXy smell and warns the user to rewrite in faster shells.
func CheckKame(smell stank.Smell) bool {
	if _, ok := KAMEINTERPRETERS[smell.Interpreter]; !ok {
		fmt.Printf("Rewrite script in sh, ksh, posh, dash, etc. for performance boost: %s\n", smell.Path)
		return true
	}

	return false
}

// CheckUsagi accepts a POSIXy smell and warns the user to rewrite in more robust shells.
func CheckUsagi(smell stank.Smell) bool {
	if _, ok := USAGIINTERPRETERS[smell.Interpreter]; !ok {
		fmt.Printf("Rewrite script in ksh, bash, zsh, etc., and enable debugging flags for robustness: %s\n", smell.Path)
		return true
	}

	return false
}

// CheckAhiru analyzes the interpreter of a POSIXy smell.
// If the interpreter is not pure sh, CheckAhiru prints an warning and returns true.
// Otherwise, CheckAhiru returns false.
func CheckAhiru(smell stank.Smell) bool {
	if smell.Interpreter != "sh" {
		fmt.Printf("Rewrite in pure POSIX sh for portability: %s\n", smell.Path)
		return true
	}

	return false
}

// Walk is a callback for filepath.Walk to scan for shell scripts.
func (o *Rose) Walk(pth string, info os.FileInfo, err error) error {
	smell, err := stank.Sniff(pth, stank.SniffConfig{})

	if err != nil && err != io.EOF {
		log.Print(err)
	}

	if smell.MachineGenerated {
		return nil
	}

	if (smell.POSIXy || smell.AltShellScript) &&
		!(stank.LOWEREXTENSIONS2CONFIG[strings.ToLower(smell.Extension)] || stank.LOWERFILENAMES2CONFIG[strings.ToLower(smell.Filename)]) &&
		!CheckShebang(smell) {

		switch o.Mode {
		case ModeRose:
			o.FoundWarning = CheckRose(smell)
		case ModeKame:
			o.FoundWarning = CheckKame(smell)
		case ModeUsagi:
			o.FoundWarning = CheckUsagi(smell)
		case ModeAhiru:
			o.FoundWarning = CheckAhiru(smell)
		}
	}

	return nil
}

func main() {
	flag.Parse()

	rose := Rose{Mode: ModeRose}

	switch {
	case *flagKame:
		rose.Mode = ModeKame
	case *flagUsagi:
		rose.Mode = ModeUsagi
	case *flagAhiru:
		rose.Mode = ModeAhiru
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
		filepath.Walk(pth, rose.Walk)
	}

	if rose.FoundWarning {
		os.Exit(1)
	}
}
