package zlisten

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

// ZListenPorts listens for and manages TCP connections on the specified ports.
func ZListenPorts(host string, ports []int) error {
	for _, port := range ports {
		go func(port int) {
			address := fmt.Sprintf("%s:%d", host, port)
			listener, err := net.Listen("tcp", address)
			if err != nil {
				log.Printf("failed to listen to TCP port %d: %s\n", port, err)
				return
			}
			defer listener.Close()

			log.Printf("Listening on %s:%d\n", host, port)

			for {
				conn, err := listener.Accept()
				if err != nil {
					log.Printf("failed to accept connection: %s\n", err)
					continue
				}
				go handleConnection(conn, &port)
			}
		}(port)
	}
	select {}
}

// handleConnection handles the incoming and outgoing data for a TCP connection. It ensure all data is writen to stdin.
func handleConnection(conn net.Conn, port *int) {
	defer conn.Close()

	log.Printf("Connection to localhost port %d [tcp] succeeded\n", *port)

	go func() {
		defer log.Printf("Connection to port %d [tcp] closed\n", *port)

		_, err := io.Copy(os.Stdout, conn)
		if err != nil {
			log.Printf("failed to send connection output to Stdout: %s\n", err)
		}
	}()
	select {}
}
