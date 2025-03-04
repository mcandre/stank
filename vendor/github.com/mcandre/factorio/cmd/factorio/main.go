// Package main implements a tool for automating Go cross-compilation.
package main

import (
	"github.com/mcandre/factorio"

	"flag"
	"fmt"
	"log"
	"os"
)

var flagVersion = flag.Bool("version", false, "show version")
var flagHelp = flag.Bool("help", false, "show usage menu")

func usage() {
	program, err := os.Executable()

	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("Usage: %v [OPTION] [-- <go build flags>]\n", program)
	flag.PrintDefaults()
}

func main() {
	flag.Parse()

	switch {
	case *flagVersion:
		fmt.Println(factorio.Version)
		os.Exit(0)
	case *flagHelp:
		usage()
		os.Exit(0)
	}

	args := flag.Args()

	if err := factorio.Port(args); err != nil {
		log.Fatal(err)
	}
}
