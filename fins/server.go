package fins

import (
	"fmt"
	"net"
)

// Server Omron FINS server (PLC emulator)
type Server struct {
	conn    *net.UDPConn
	addr    *Address
	handler CommandHandler
}

type CommandHandler func(*Command) *Response

func NewServer(udpAddr *net.UDPAddr, addr *Address, handler CommandHandler) (*Server, error) {
	s := new(Server)

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		return nil, err
	}
	s.conn = conn
	s.addr = addr
	if handler == nil {
		s.handler = func(command *Command) *Response {
			fmt.Printf("Null command handler: 0x%04x\n", command.CommandCode())

			response := NewResponse(command.CommandCode(), EndCodeNotSupportedByModelVersion, []byte{})
			return response
		}
	} else {
		s.handler = handler
	}

	go func() {
		var buf [1024]byte
		for {
			//rlen
			rlen, remote, err := conn.ReadFromUDP(buf[:])
			cmdFrame := decodeFrame(buf[:rlen])
			cmd := cmdFrame.Payload().(*Command)
			rsp := s.handler(cmd)
			if err != nil {
				panic(err)
			}

			rspFrame := NewFrame(defaultResponseHeader(cmdFrame.Header()), rsp)
			_, err = conn.WriteToUDP(encodeFrame(rspFrame), &net.UDPAddr{IP: remote.IP, Port: remote.Port})
			if err != nil {
				panic(err)
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
