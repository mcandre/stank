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
var flagHelp = flag.Bool("help", false, "Show usage information")
var flagVersion = flag.Bool("version", false, "Show version information")

type RoseMode int

const (
	ModeRose RoseMode = iota
	ModeKame
	ModeUsagi
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

// USAGIINTERPRETERS catalogues POSIX shells with rich debugging `set` features, such as set -euo pipefail.
var USAGIINTERPRETERS = map[string]bool{
	"ksh":   true,
	"mksh":  true,
	"pdksh": true,
	"ksh93": true,
	"ksh88": true,
	"bash":  true,
	"zsh":   true,
}

// CheckRose accepts a POSIXy smell and warns the user to rewrite in a safer language.
func CheckRose(smell stank.Smell) bool {
	fmt.Printf("Rewrite POSIX script in Ruby or other safer general purpose scripting language: %s\n", smell.Path)
	return true
}

// CheckShebang warns on POSIXy scripts that lack a shebang line to distinguish which interpreter should be used.
func CheckShebang(smell stank.Smell) bool {
	if smell.Interpreter == "generic-sh" {
		fmt.Printf("Clarify interpreter with a shebang line: %s\n", smell.Path)
		return true
	}

	return false
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

// Walk is a callback for filepath.Walk to scan for shell scripts.
func (o *Rose) Walk(pth string, info os.FileInfo, err error) error {
	smell, err := stank.Sniff(pth, stank.SniffConfig{})

	if err != nil && err != io.EOF {
		log.Print(err)
	}

	if smell.POSIXy && !(stank.LOWEREXTENSIONS2CONFIG[strings.ToLower(smell.Extension)] || stank.LOWERFILENAMES2CONFIG[strings.ToLower(smell.Filename)]) && !CheckShebang(smell) {
		switch o.Mode {
		case ModeRose:
			o.FoundWarning = CheckRose(smell)
		case ModeKame:
			o.FoundWarning = CheckKame(smell)
		case ModeUsagi:
			o.FoundWarning = CheckUsagi(smell)
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
