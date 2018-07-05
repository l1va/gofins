package fins

import (
	"encoding/hex"
	"fmt"
	"net"
)

// Server Omron FINS server (PLC emulator)
type Server struct {
	conn *net.UDPConn
	addr *Address
}

func NewServer(udpAddr *net.UDPAddr, addr *Address) (*Server, error) {
	s := new(Server)

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		return nil, err
	}
	s.conn = conn
	s.addr = addr

	go func() {
		for {
			var buf [1024]byte
			for {
				rlen, remote, err := conn.ReadFromUDP(buf[:])
				// s := string(buf[:rlen])
				fmt.Printf("Received %d bytes from %s\n%s", rlen, remote.IP, hex.Dump(buf[:rlen]))
				if err != nil {
					panic(err)
				}
			}
		}
	}()

	return s, nil
}

// Close Closes the FINS server
func (s *Server) Close() {
	s.conn.Close()
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	fmt.Printf("Received %d bytes\n", reqLen)
	// Send a response back to person contacting us.
	conn.Write([]byte("Message received."))
	// Close the connection when you're done with it.
	conn.Close()
}
