package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/mcandre/stank"
)

var flagPrettyPrint = flag.Bool("pp", false, "Prettyprint smell records")
var flagHelp = flag.Bool("help", false, "Show usage information")
var flagVersion = flag.Bool("version", false, "Show version information")

// Stinker holds configuration for a stinky walk.
type Stinker struct {
	PrettyPrint bool
}

// Walk sniffs a path,
// printing the smell of the script.
//
// If PrettyPrint is false, then the smell is minified.
func (o Stinker) Walk(pth string, info os.FileInfo, err error) error {
	smell, err := stank.Sniff(pth)

	if err != nil && err != io.EOF {
		log.Print(err)

		return err
	}

	if smell.Directory {
		return nil
	}

	var smellBytes []byte

	if o.PrettyPrint {
		smellBytes, _ = json.MarshalIndent(smell, "", "  ")
	} else {
		smellBytes, _ = json.Marshal(smell)
	}

	smellJSON := string(smellBytes)

	fmt.Println(smellJSON)

	return err
}

func main() {
	flag.Parse()

	stinker := Stinker{}

	switch {
	case *flagPrettyPrint:
		stinker.PrettyPrint = true
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
		err = filepath.Walk(pth, stinker.Walk)

		if err != nil && err != io.EOF {
			log.Print(err)
		}
	}
}
