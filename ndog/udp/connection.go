package udp

import (
	"fmt"
	"net"
)

// UDPConnection represents a UDP connection.
type UDPConnection struct {
	conn *net.UDPConn
}

// NewUDPConnection creates a new UDP connection.
// It listens for UDP connections on the provided UDP address.
func NewUDPConnection(updAddr *net.UDPAddr) (*UDPConnection, error) {
	conn, err := net.ListenUDP("udp", updAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to listen to UDP address %s: %w", updAddr, err)
	}
	return &UDPConnection{conn: conn}, nil
}

// Write writes data to the UDP connection.
// It implements the io.Writer interface.
func (u *UDPConnection) Write(p []byte) (n int, err error) {
	return u.conn.Write(p)
}
