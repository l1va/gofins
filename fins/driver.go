package fins

import (
	"encoding/binary"
	"errors"
)

func readCommand(ioAddr IOAddress, readCount uint16) *Payload {
	p := &Payload{
		CommandCode: CommandCodeMemoryAreaRead,
		Data:        make([]byte, 0, 6),
	}
	p.Data = append(p.Data, encodeIOAddress(ioAddr)...)
	p.Data = append(p.Data, []byte{0, 0}...)
	binary.BigEndian.PutUint16(p.Data[4:6], readCount)
	return p
}

func encodeIOAddress(ioAddr IOAddress) []byte {
	bytes := make([]byte, 4, 4)
	bytes[0] = ioAddr.MemoryArea
	binary.BigEndian.PutUint16(bytes[1:3], ioAddr.Address)
	bytes[3] = ioAddr.BitOffset
	return bytes
}

// func writeDataCommand(header *Header, startAddress uint16, data []uint16) *Payload {
// 	return writeCommand(MEMORY_AREA_DATA, header, startAddress, data)
// }

// func writeWorkCommand(header *Header, startAddress uint16, data []uint16) *Payload {
// 	return writeCommand(MEMORY_AREA_WORK, header, startAddress, data)
// }

// func writeCommand(memoryArea byte, header *Header, startAddress uint16, data []uint16) *Payload {
// 	var addressBit byte = 0
// 	addressLower := byte(startAddress)
// 	addressUpper := byte(startAddress >> 8)
// 	dataLen := len(data)
// 	lenLower := byte(dataLen)
// 	lenUpper := byte(dataLen >> 8)

// 	paramsBytes := []byte{
// 		memoryArea,
// 		addressUpper, addressLower,
// 		addressBit,
// 		lenUpper, lenLower}

// 	bytes1 := append(header.Format(), CommandCodeMemoryAreaWrite...)
// 	bytes2 := append(bytes1, paramsBytes...)
// 	bytes3 := append(bytes2, toBytes(data)...)
// 	return bytes3
// }

func decodeFrame(bytes []byte) *Frame {
	frame := &Frame{
		Header:  decodeHeader(bytes[:10]),
		Payload: decodePayload(bytes[10:]),
	}
	return frame
}

func encodeFrame(f *Frame) []byte {
	bytes := encodeHeader(f.Header)
	bytes = append(bytes, encodePayload(f.Payload)...)
	return bytes
}

func decodeHeader(bytes []byte) *Header {
	header := &Header{
		icf: bytes[0],
		rsv: bytes[1],
		gct: bytes[2],
		dst: Address{
			Network: bytes[3],
			Node:    bytes[4],
			Unit:    bytes[5],
		},
		src: Address{
			Network: bytes[6],
			Node:    bytes[7],
			Unit:    bytes[8],
		},
		sid: bytes[9],
	}
	return header
}

func encodeHeader(h *Header) []byte {
	bytes := []byte{
		h.icf, h.rsv, h.gct,
		h.dst.Network, h.dst.Node, h.dst.Unit,
		h.src.Network, h.src.Node, h.src.Unit,
		h.sid}
	return bytes
}

func decodePayload(bytes []byte) *Payload {
	payload := &Payload{
		CommandCode: binary.BigEndian.Uint16(bytes[:2]),
		Data:        bytes[2:],
	}
	return payload
}

func encodePayload(payload *Payload) []byte {
	bytes := make([]byte, 2, 2+len(payload.Data))
	binary.BigEndian.PutUint16(bytes, payload.CommandCode)
	bytes = append(bytes, payload.Data...)
	return bytes
}

func toUint16(data []byte) []uint16 {
	res := make([]uint16, len(data)/2)
	for i := 0; i < len(data); i += 2 {
		upper := uint16(data[i]) << 8
		lower := uint16(data[i+1])
		res[i/2] = (upper | lower)
	}
	return res
}

func toBytes(data []uint16) []byte {
	res := make([]byte, len(data)*2)
	for i := 0; i < len(data); i++ {
		res[2*i] = byte(data[i] >> 8)
		res[2*i+1] = byte(data[i])
	}
	return res
}

// ErrBCDBadDigit BCD bad digit error
var ErrBCDBadDigit = errors.New("Bad digit in BCD decoding")

// ErrBCDOverflow BCD overflow error
var ErrBCDOverflow = errors.New("Overflow occurred in BCD decoding")

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
		return 0, ErrBCDOverflow
	}
	return x5<<1 + digit, nil
}

func decodeBCD(bcd []byte) (x uint64, err error) {
	for i, b := range bcd {
		hi, lo := uint64(b>>4), uint64(b&0x0f)
		if hi > 9 {
			return 0, ErrBCDBadDigit
		}
		x, err = timesTenPlusCatchingOverflow(x, hi)
		if err != nil {
			return 0, err
		}
		if lo == 0x0f && i == len(bcd)-1 {
			return x, nil
		}
		if lo > 9 {
			return 0, ErrBCDBadDigit
		}
		x, err = timesTenPlusCatchingOverflow(x, lo)
		if err != nil {
			return 0, err
		}
	}
	return x, nil
}
