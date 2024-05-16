package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/marktlinn/netdog/ndog/tcp"
	"github.com/marktlinn/netdog/ndog/udp"
	zlisten "github.com/marktlinn/netdog/ndog/z_listen"
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
	zHost      string
	zPorts     []int
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

	f.l = *listen
	f.p = *port
	f.u = *u
	f.x = *hex
	f.e = *exe

	log.Printf("arg listen: %t\n", f.l)
	log.Printf("arg udp: %t\n", f.u)
	log.Printf("arg hex: %t\n", f.x)
	log.Printf("arg port: %d\n", f.p)
	log.Printf("arg exe: %s\n", f.e)
	log.Printf("arg silentPort: %s\n", *silentPort)

	if *silentPort != "" {
		if err := f.checkZFlag(); err != nil {
			return fmt.Errorf("failed to parse -z flag: %w", err)
		}
	}

	if f.z {
		log.Printf("received host: %s\n", f.zHost)
		for _, port := range f.zPorts {
			log.Printf("received port: %d\n", port)
		}

		zlisten.ZListenPorts(f.zHost, f.zPorts)
	}
	if *listen && !*u {
		tcp.ListenTCP(*port)
	}
	if *listen && *u {
		udp.ListenUDP(*port)
	}

	return nil
}

func (f *flags) checkZFlag() error {
	zIndex := -1

	for i, arg := range os.Args {
		if arg == "-z" {
			zIndex = i
			break
		}
	}

	if zIndex != -1 && len(os.Args) > zIndex+2 {
		f.zHost = os.Args[zIndex+1]
		portRange := os.Args[zIndex+2]

		var err error
		f.zPorts, err = parsePorts(portRange)
		if err != nil {
			return fmt.Errorf("invalid -z flag &/or arguments received: %w", err)
		}
	} else {
		return fmt.Errorf("invalid -z flag &/or arguments received")
	}
	f.z = true

	return nil
}

func parsePorts(portRange string) ([]int, error) {
	portParts := strings.Split(portRange, "-")

	var portSlice []int
	if len(portParts) == 1 {
		port, err := strconv.Atoi(portParts[0])
		if err != nil {
			return nil, fmt.Errorf("failed to parse port %s: %s", portParts[0], err)
		}
		portSlice = append(portSlice, port)
	} else if len(portParts) == 2 {
		startPort, err := strconv.Atoi(portParts[0])
		if err != nil {
			return nil, fmt.Errorf("failed to parse first port %s: %s", portParts[0], err)
		}

		endPort, err := strconv.Atoi(portParts[1])
		if err != nil {
			return nil, fmt.Errorf("failed to parse end port %s: %s", portParts[1], err)
		}
		numOfPorts := endPort - startPort

		portSlice = make([]int, numOfPorts+1)

		portIndex := 0
		for port := startPort; port <= endPort; port++ {
			portSlice[portIndex] = port
			portIndex++
		}
	} else {
		return nil, fmt.Errorf("invalid port range format: %s", portRange)
	}
	return portSlice, nil
}
