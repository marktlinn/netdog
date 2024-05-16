package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	f := &flags{}
	if err := f.Parse(flag.CommandLine, os.Args[1:]); err != nil {
		log.Fatalf("failed to process flags: %s\n", err)
	}
}
