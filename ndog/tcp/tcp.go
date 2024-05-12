package tcp

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func ListenTcp(port int) error {
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

	_, err := io.Copy(conn, os.Stdin)
	if err != nil {
		log.Printf("failed to send input to conncetion client: %s\n", err)
	}
}
