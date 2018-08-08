package fins

import (
	"fmt"
	"net"
)

// Server Omron FINS server (PLC emulator)
type Server struct {
	addr    Address
	conn    *net.UDPConn
	handler CommandHandler
}

type CommandHandler func(req request) response

func NewServer(plcAddr Address, handler CommandHandler) (*Server, error) {
	s := new(Server)
	s.addr = plcAddr

	conn, err := net.ListenUDP("udp", plcAddr.udpAddress)
	if err != nil {
		return nil, err
	}
	s.conn = conn

	s.handler = handler

	if handler == nil {
		s.handler = defaultHandler
	}

	go func() {
		var buf [1024]byte
		for {
			rlen, remote, err := conn.ReadFromUDP(buf[:])
			req := decodeRequest(buf[:rlen])
			resp := s.handler(req)

			_, err = conn.WriteToUDP(encodeResponse(resp), &net.UDPAddr{IP: remote.IP, Port: remote.Port})
			if err != nil {
				panic(err)
			}
		}
	}()

	return s, nil
}

func defaultHandler(r request) response {
	fmt.Printf("Null command handler: 0x%04x\n", r.commandCode)

	response := response{defaultResponseHeader(r.header), r.commandCode, EndCodeNotSupportedByModelVersion, []byte{}}
	return response
}

// Close Closes the FINS server
func (s *Server) Close() {
	s.conn.Close()
}
