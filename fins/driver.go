package fins

import (
	"encoding/binary"
)

// request A FINS command request
type request struct {
	header Header
	commandCode uint16
	data   []byte
}

// response A FINS command response
type response struct {
	header      Header
	commandCode uint16
	endCode     uint16
	data        []byte
}

// memoryAddress A plc memory address to do a work
type memoryAddress struct {
	memoryArea byte
	address    uint16
	bitOffset  byte
}

func memAddr(memoryArea byte, address uint16) memoryAddress {
	return memAddrWithBitOffset(memoryArea, address, 0)
}

func memAddrWithBitOffset(memoryArea byte, address uint16, bitOffset byte) memoryAddress {
	return memoryAddress{memoryArea, address, bitOffset}
}

func readCommand(memoryAddr memoryAddress, itemCount uint16) []byte {
	commandData := make([]byte, 2, 8)
	binary.BigEndian.PutUint16(commandData[0:2], CommandCodeMemoryAreaRead)
	commandData = append(commandData, encodeMemoryAddress(memoryAddr)...)
	commandData = append(commandData, []byte{0, 0}...)
	binary.BigEndian.PutUint16(commandData[6:8], itemCount)
	return commandData
}

func writeCommand(memoryAddr memoryAddress, itemCount uint16, bytes []byte) []byte {
	commandData := make([]byte, 2, 8+len(bytes))
	binary.BigEndian.PutUint16(commandData[0:2], CommandCodeMemoryAreaWrite)
	commandData = append(commandData, encodeMemoryAddress(memoryAddr)...)
	commandData = append(commandData, []byte{0, 0}...)
	binary.BigEndian.PutUint16(commandData[6:8], itemCount)
	commandData = append(commandData, bytes...)
	return commandData
}

func clockReadCommand() []byte {
	commandData := make([]byte, 2, 2)
	binary.BigEndian.PutUint16(commandData[0:2], CommandCodeClockRead)
	return commandData
}

func encodeMemoryAddress(memoryAddr memoryAddress) []byte {
	bytes := make([]byte, 4, 4)
	bytes[0] = memoryAddr.memoryArea
	binary.BigEndian.PutUint16(bytes[1:3], memoryAddr.address)
	bytes[3] = memoryAddr.bitOffset
	return bytes
}

func decodeMemoryAddress(data []byte) memoryAddress {
	return memoryAddress{data[0], binary.BigEndian.Uint16(data[1:3]), data[3]}
}

func decodeRequest(bytes []byte) request {
	return request{
		decodeHeader(bytes[0:10]),
		binary.BigEndian.Uint16(bytes[10:12]),
		bytes[12:],
	}
}

func decodeResponse(bytes []byte) response {
	return response{
		decodeHeader(bytes[0:10]),
		binary.BigEndian.Uint16(bytes[10:12]),
		binary.BigEndian.Uint16(bytes[12:14]),
		bytes[14:],
	}
}
func encodeResponse(resp response) []byte {
	bytes := make([]byte, 4, 4+len(resp.data))
	binary.BigEndian.PutUint16(bytes[0:2], resp.commandCode)
	binary.BigEndian.PutUint16(bytes[2:4], resp.endCode)
	bytes = append(bytes, resp.data...)
	bh := encodeHeader(resp.header)
	bh = append(bh, bytes...)
	return bh
}

const (
	icfBridgesBit          byte = 7
	icfMessageTypeBit      byte = 6
	icfResponseRequiredBit byte = 0
)

func decodeHeader(bytes []byte) Header {
	header := Header{}
	icf := bytes[0]
	if icf&1<<icfResponseRequiredBit == 0 {
		header.responseRequired = true
	}
	if icf&1<<icfMessageTypeBit == 0 {
		header.messageType = MessageTypeCommand
	} else {
		header.messageType = MessageTypeResponse
	}
	header.gatewayCount = bytes[2]
	header.dst = finsAddress{bytes[3], bytes[4], bytes[5]}
	header.src = finsAddress{bytes[6], bytes[7], bytes[8]}
	header.serviceID = bytes[9]

	return header
}

func encodeHeader(h Header) []byte {
	var icf byte
	icf = 1 << icfBridgesBit
	if h.responseRequired == false {
		icf |= 1 << icfResponseRequiredBit
	}
	if h.messageType == MessageTypeResponse {
		icf |= 1 << icfMessageTypeBit
	}
	bytes := []byte{
		icf, 0x00, h.gatewayCount,
		h.dst.network, h.dst.node, h.dst.unit,
		h.src.network, h.src.node, h.src.unit,
		h.serviceID}
	return bytes
}

func encodeBCD(x uint64) []byte {
	if x == 0 {
		return []byte{0x0f}
	}
	var n int
	for xx := x; xx > 0; n++ {
		xx = xx / 10
	}
	bcd := make([]byte, (n+1)/2)
	if n%2 == 1 {
		hi, lo := byte(x%10), byte(0x0f)
		bcd[(n-1)/2] = hi<<4 | lo
		x = x / 10
		n--
	}
	for i := n/2 - 1; i >= 0; i-- {
		hi, lo := byte((x/10)%10), byte(x%10)
		bcd[i] = hi<<4 | lo
		x = x / 100
	}
	return bcd
}

func timesTenPlusCatchingOverflow(x uint64, digit uint64) (uint64, error) {
	x5 := x<<2 + x
	if int64(x5) < 0 || x5<<1 > ^digit {
		return 0, BCDOverflowError{}
	}
	return x5<<1 + digit, nil
}

func decodeBCD(bcd []byte) (x uint64, err error) {
	for i, b := range bcd {
		hi, lo := uint64(b>>4), uint64(b&0x0f)
		if hi > 9 {
			return 0, BCDBadDigitError{"hi", hi}
		}
		x, err = timesTenPlusCatchingOverflow(x, hi)
		if err != nil {
			return 0, err
		}
		if lo == 0x0f && i == len(bcd)-1 {
			return x, nil
		}
		if lo > 9 {
			return 0, BCDBadDigitError{"lo", lo}
		}
		x, err = timesTenPlusCatchingOverflow(x, lo)
		if err != nil {
			return 0, err
		}
	}
	return x, nil
}
