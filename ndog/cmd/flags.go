package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

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
	silentPort := flag.String("z", "8080", "Scan provided Host & Port without sending data.")

	flag.Parse()

	log.Printf("arg listen: %t\n", *listen)
	log.Printf("arg udp: %t\n", *u)
	log.Printf("arg hex: %t\n", *hex)
	log.Printf("arg port: %d\n", *port)
	log.Printf("arg exe: %s\n", *exe)
	log.Printf("arg silentPort: %s\n", *silentPort)

	if *silentPort != "" {
		h, p, err := checkPorts(silentPort)
		if err != nil {
			return fmt.Errorf("invalid -z flag &/or arguments received: %w", err)
		}
		log.Printf("received host: %s\n", h)
		for _, port := range p {
			log.Printf("received port: %d\n", port)
		}
	}

	if *listen && !*u {
		tcp.ListenTCP(*port)
	}
	if *listen && *u {
		udp.ListenUDP(*port)
	}

	return nil
}

func checkPorts(silentPort *string) (host string, ports []int, err error) {
	var portSlice []int
	hostPortParts := strings.Split(*silentPort, " ")
	if len(hostPortParts) < 2 {
		return "", nil, fmt.Errorf("received invalid format for -z")
	}

	h := hostPortParts[0]
	if h == "" {
		return "", nil, fmt.Errorf("received invalid host")
	}
	p := strings.Split(hostPortParts[1], "-")

	for _, p := range p {
		portInt, err := strconv.Atoi(p)
		if err != nil {
			return "", nil, fmt.Errorf("failed to parse port %s: %s", p, err)
		}
		portSlice = append(portSlice, portInt)
	}

	return h, portSlice, nil
}
