package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/marktlinn/netdog/ndog/tcp"
	"github.com/marktlinn/netdog/ndog/udp"
)

const usage = `
Usage:
	nd [options]
Options:`

// flags defines the exepcted cli flags the program expects.
type flags struct {
	e          string
	p          int
	u, x, l, z bool
}

func (f *flags) Parse(flags *flag.FlagSet, args []string) error {
	flag.Usage = func() {
		fmt.Fprintln(flags.Output(), usage[1:])
		flags.PrintDefaults()
	}

	listen := flag.Bool("l", false, "Listen for TCP connections.")
	u := flag.Bool("u", false, "Listen for UDP connections.")
	hex := flag.Bool("x", false, "Enables Hex dump between server & client.")
	port := flag.Int("p", 8080, "Port to be targetted.")
	exe := flag.String("e", "", "Execute and pipe process between server and client.")
	silentPort := flag.String("z", "", "Scan provided Host & Port without sending data.")

	flag.Parse()

	log.Printf("arg listen: %t\n", *listen)
	log.Printf("arg udp: %t\n", *u)
	log.Printf("arg hex: %t\n", *hex)
	log.Printf("arg port: %d\n", *port)
	log.Printf("arg exe: %s\n", *exe)
	log.Printf("arg silentPort: %s\n", *silentPort)

	if *listen && !*u {
		tcp.ListenTCP(*port)
	}
	if *listen && *u {
		udp.ListenUDP(*port)
	}

	return nil
}
