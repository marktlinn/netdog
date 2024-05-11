package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	log.Println("Running Main...")
	f := &flags{}
	if err := f.Parse(flag.CommandLine, os.Args[1:]); err != nil {
		os.Exit(1)
	}
}
