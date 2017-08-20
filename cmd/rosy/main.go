package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/mcandre/stank"
)

var flagHelp = flag.Bool("help", false, "Show usage information")
var flagVersion = flag.Bool("version", false, "Show version information")

// PointWalkFunc is any function that operates on a path and boolean pointer.
// For example, a linting function may have the pointer point to true to signify
// the tripping of a linter warning.
type PointWalkFunc func(pth string, p *bool)

// PointWalk recursivly applies the given PointWalkFunc to each
// nondirectory node in a given path.
func PointWalk(pth string, fn PointWalkFunc, p *bool) {
	fi, err := os.Stat(pth)

	if err != nil {
		log.Panic(err)
	}

	switch mode := fi.Mode(); {
	case mode.IsDir():
		fis, err := ioutil.ReadDir(pth)

		if err != nil {
			log.Panic(err)
		}

		for _, fi := range fis {
			PointWalk(path.Join(pth, fi.Name()), fn, p)
		}
	default:
		fn(pth, p)
	}
}

func RosyWalk(pth string, p *bool) {
	smell, err := stank.Sniff(pth)

	if err != nil && err != io.EOF {
		log.Print(err)
	}

	if smell.POSIXy {
		fmt.Printf("Rewrite POSIX script in Ruby or other safer general purpose scripting language: %s\n", pth)
		*p = true
	}
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

	var foundPOSIX bool

	for _, pth := range paths {
		PointWalk(pth, RosyWalk, &foundPOSIX)
	}

	if foundPOSIX {
		os.Exit(1)
	}
}
