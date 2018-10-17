package fins

import (
	"encoding/binary"
	"fmt"
	"net"
)

// Server Omron FINS server (PLC emulator)
type Server struct {
	addr    Address
	conn    *net.UDPConn
	handler CommandHandler
	dmarea []byte
}

type CommandHandler func(req request, dmarea []byte) response

const DM_AREA_SIZE = 32768

func NewServer(plcAddr Address, handler CommandHandler) (*Server, error) {
	s := new(Server)
	s.addr = plcAddr
	s.dmarea = make([]byte, DM_AREA_SIZE)

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
			if rlen > 0 {
				req := decodeRequest(buf[:rlen])
				resp := s.handler(req, s.dmarea)

				_, err = conn.WriteToUDP(encodeResponse(resp), &net.UDPAddr{IP: remote.IP, Port: remote.Port})
				if err != nil {
					panic(err)
				}
			}
		}
	}()

	return s, nil
}

func defaultHandler(r request, dmarea []byte) response {
	fmt.Printf("Null command handler: 0x%04x\n", r.commandCode)

	response := response{defaultResponseHeader(r.header), r.commandCode, EndCodeNotSupportedByModelVersion, []byte{}}
	return response
}

// Works with only DM area, 2 byte integers
func DMAreaHandler(r request, dmarea []byte) response {
	var endCode uint16
	data := []byte{}
	switch r.commandCode {
	case CommandCodeMemoryAreaRead, CommandCodeMemoryAreaWrite:
		memAddr := decodeMemoryAddress(r.data[:4])
		ic := binary.BigEndian.Uint16(r.data[4:6]) // Item count

		if memAddr.memoryArea != MemoryAreaDMWord {
			endCode = EndCodeNotSupportedByModelVersion
		} else if memAddr.address+ic*2 > DM_AREA_SIZE { // Check address boundary
			endCode = EndCodeAddressRangeExceeded
		} else {
			if r.commandCode == CommandCodeMemoryAreaRead {
				data = dmarea[memAddr.address:memAddr.address+ic*2]
			} else { // Write command
				copy(dmarea[memAddr.address:memAddr.address+ic*2], r.data[6:6+ic*2])
			}
			endCode = EndCodeNormalCompletion
		}
	default:
		return defaultHandler(r, dmarea)
	}
	return response{defaultResponseHeader(r.header), r.commandCode, endCode, data}
}

// Close Closes the FINS server
func (s *Server) Close() {
	s.conn.Close()
}
