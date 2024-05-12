package main

import (
	"flag"
	"os"
)

func main() {
	f := &flags{}
	if err := f.Parse(flag.CommandLine, os.Args[1:]); err != nil {
		os.Exit(1)
	}
}
