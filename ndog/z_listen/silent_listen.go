package zlisten

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

// ZListenPorts listens for and manages TCP connections on the specified ports.
func ZListenPorts(port []int) error {
	address := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen to TCP port :%d: %w", port, err)
	}
	defer listener.Close()

	log.Printf("Listening on port %d\n", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error: listener failed to accept connection: %s\n", err)
			continue
		}

		go handleConnection(conn)

	}
}

// handleConnection handles the incoming and outgoing data for a TCP connection. It ensure all data is writen to stdin or stdout respectively.
func handleConnection(conn net.Conn) {
	defer conn.Close()

	log.Printf("Connection established to %s\n", conn.RemoteAddr().String())

	go func() {
		defer log.Printf("Connection to %s closed\n", conn.RemoteAddr().String())

		_, err := io.Copy(os.Stdout, conn)
		if err != nil {
			log.Printf("failed to send connection output to Stdout: %s\n", err)
		}
	}()
}
