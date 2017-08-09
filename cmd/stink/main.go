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

var flagHelp = flag.Bool("help", false, "Show usage information")
var flagVersion = flag.Bool("version", false, "Show version information")

func StinkWalk(pth string, info os.FileInfo, err error) error {
	smell, err := stank.Sniff(pth)

	smellBytes, _ := json.Marshal(smell)
	smellJSON := string(smellBytes)

	fmt.Println(smellJSON)

	return err
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

	for _, pth := range paths {
		err := filepath.Walk(pth, StinkWalk)

		if err != nil {
			log.Print(err)
		}
	}
}
