package fins

import (
	"encoding/binary"
	"errors"
)

func readCommand(ioAddr *IOAddress, itemCount uint16) *Command {
	commandData := make([]byte, 0, 6)
	commandData = append(commandData, encodeIOAddress(ioAddr)...)
	commandData = append(commandData, []byte{0, 0}...)
	binary.BigEndian.PutUint16(commandData[4:6], itemCount)
	return NewCommand(CommandCodeMemoryAreaRead, commandData)
}

func writeCommand(ioAddr *IOAddress, itemCount uint16, bytes []byte) *Command {
	commandData := make([]byte, 0, 6+len(bytes))
	commandData = append(commandData, encodeIOAddress(ioAddr)...)
	commandData = append(commandData, []byte{0, 0}...)
	binary.BigEndian.PutUint16(commandData[4:6], itemCount)
	commandData = append(commandData, bytes...)
	return NewCommand(CommandCodeMemoryAreaWrite, commandData)
}

func encodeIOAddress(ioAddr *IOAddress) []byte {
	bytes := make([]byte, 4, 4)
	bytes[0] = ioAddr.MemoryArea()
	binary.BigEndian.PutUint16(bytes[1:3], ioAddr.Address())
	bytes[3] = ioAddr.BitOffset()
	return bytes
}

func decodeFrame(bytes []byte) *Frame {
	header := decodeHeader(bytes[0:10])
	var payload Payload
	if header.FrameIsCommand() {
		payload = decodeCommand(bytes[10:])
	} else if header.FrameIsResponse() {
		payload = decodeResponse(bytes[10:])
	}
	frame := NewFrame(header, payload)
	return frame
}

func encodeFrame(f *Frame) []byte {
	bytes := encodeHeader(f.Header())
	var payloadData []byte
	if f.Header().FrameIsCommand() {
		payloadData = encodeCommand(f.Payload().(*Command))
	} else if f.Header().FrameIsResponse() {
		payloadData = encodeResponse(f.Payload().(*Response))
	}
	bytes = append(bytes, payloadData...)
	return bytes
}

const (
	icfBridgesBit          byte = 7
	icfMessageTypeBit      byte = 6
	icfResponseRequiredBit byte = 0
)

func decodeHeader(bytes []byte) *Header {
	header := new(Header)
	icf := bytes[0]
	if icf&1<<icfResponseRequiredBit == 0 {
		header.responseRequired = true
	}
	if icf&1<<icfMessageTypeBit == 0 {
		header.messgeType = MessageTypeCommand
	} else {
		header.messgeType = MessageTypeResponse
	}
	header.gatewayCount = bytes[2]
	header.dst = NewAddress(bytes[3], bytes[4], bytes[5])
	header.src = NewAddress(bytes[6], bytes[7], bytes[8])
	header.serviceID = bytes[9]

	return header
}

func encodeHeader(h *Header) []byte {
	var icf byte
	icf = 0x80
	if h.responseRequired == false {
		icf |= 1 << icfResponseRequiredBit
	}
	if h.messgeType == MessageTypeResponse {
		icf |= 1 << icfMessageTypeBit
	}
	bytes := []byte{
		icf, 0x00, h.gatewayCount,
		h.dst.Network(), h.dst.Node(), h.dst.Unit(),
		h.src.Network(), h.src.Node(), h.src.Unit(),
		h.serviceID}
	return bytes
}

func decodeCommand(bytes []byte) *Command {
	return NewCommand(
		binary.BigEndian.Uint16(bytes[0:2]),
		bytes[2:])
}

func encodeCommand(command *Command) []byte {
	bytes := make([]byte, 2, 2+len(command.Data()))
	binary.BigEndian.PutUint16(bytes[0:2], command.CommandCode())
	bytes = append(bytes, command.Data()...)
	return bytes
}

func decodeResponse(bytes []byte) *Response {
	return NewResponse(
		binary.BigEndian.Uint16(bytes[0:2]),
		binary.BigEndian.Uint16(bytes[2:4]),
		bytes[4:])
}

func encodeResponse(response *Response) []byte {
	bytes := make([]byte, 4, 4+len(response.Data()))
	binary.BigEndian.PutUint16(bytes[0:2], response.CommandCode())
	binary.BigEndian.PutUint16(bytes[2:4], response.EndCode())
	bytes = append(bytes, response.Data()...)
	return bytes
}

var errBCDBadDigit = errors.New("Bad digit in BCD decoding")
var errBCDOverflow = errors.New("Overflow occurred in BCD decoding")

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
		return 0, errBCDOverflow
	}
	return x5<<1 + digit, nil
}

func decodeBCD(bcd []byte) (x uint64, err error) {
	for i, b := range bcd {
		hi, lo := uint64(b>>4), uint64(b&0x0f)
		if hi > 9 {
			return 0, errBCDBadDigit
		}
		x, err = timesTenPlusCatchingOverflow(x, hi)
		if err != nil {
			return 0, err
		}
		if lo == 0x0f && i == len(bcd)-1 {
			return x, nil
		}
		if lo > 9 {
			return 0, errBCDBadDigit
		}
		x, err = timesTenPlusCatchingOverflow(x, lo)
		if err != nil {
			return 0, err
		}
	}
	return x, nil
}
