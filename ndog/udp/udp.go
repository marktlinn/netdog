package udp

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

// ListenUDP listens for and manages UDP connections on the specified port.
func ListenUDP(port int) error {
	address := fmt.Sprintf(":%d", port)
	updAddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return fmt.Errorf("failed to resolve udp address: %s", err)
	}

	udp, err := NewUDPConnection(updAddr)
	if err != nil {
		return fmt.Errorf("failed to listen to UDP port :%d: %w", port, err)
	}
	defer udp.conn.Close()

	log.Printf("Listening for UDP connections on port %s\n", address)

	handleConnection(udp.conn)

	return nil
}

// handleConnection handles the incoming and outgoing data for a UDP connection. It ensure all data is writen to stdin or stdout respectively.
func handleConnection(conn *net.UDPConn) {
	go func() {
		for {
			_, err := io.Copy(conn, os.Stdin)
			if err != nil {
				log.Printf("Error writing to UDP connection: %s\n", err)
			}
		}
	}()

	go func() {
		buffer := make([]byte, 1024)
		for {
			n, _, err := conn.ReadFromUDP(buffer)
			if err != nil {
				log.Printf("Error reading from UDP connection: %s\n", err)
				return
			}

			_, err = os.Stdout.Write(buffer[:n])
			if err != nil {
				log.Printf("Error writing to stdout; %s\n", err)
				return
			}
		}
	}()

	select {}
}
