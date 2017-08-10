package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/mcandre/stank"
)

var flagPrettyPrint = flag.Bool("pp", false, "Prettyprint smell records")
var flagHelp = flag.Bool("help", false, "Show usage information")
var flagVersion = flag.Bool("version", false, "Show version information")

// StinkWalkFair sniffs a path,
// printing the smell of the script.
//
// If pp is true, the smell record is prettyprinted.
func StinkWalkFair(pth string, info os.FileInfo, err error, pp bool) error {
	smell, err := stank.Sniff(pth)

	var smellBytes []byte

	if pp {
		smellBytes, _ = json.MarshalIndent(smell, "", "  ")
	} else {
		smellBytes, _ = json.Marshal(smell)
	}

	smellJSON := string(smellBytes)

	fmt.Println(smellJSON)

	return err
}

// StinkWalk sniffs a file system node for POSIXyness and prints the smell record to STDOUT.
func StinkWalk(pth string, info os.FileInfo, err error) error {
	return StinkWalkFair(pth, info, err, false)
}

// StinkWalkPretty sniffs a file system node for POSIXyness and prettyprints the smell record to STDOUT.
func StinkWalkPretty(pth string, info os.FileInfo, err error) error {
	return StinkWalkFair(pth, info, err, true)
}

func main() {
	flag.Parse()

	var pp bool

	switch {
	case *flagPrettyPrint:
		pp = true
	case *flagVersion:
		fmt.Println(stank.Version)
		os.Exit(0)
	case *flagHelp:
		flag.PrintDefaults()
		os.Exit(0)
	}

	paths := flag.Args()

	var err error

	for _, pth := range paths {
		if pp {
			err = filepath.Walk(pth, StinkWalkPretty)
		} else {
			err = filepath.Walk(pth, StinkWalk)
		}

		if err != nil {
			log.Print(err)
		}
	}
}
