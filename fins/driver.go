package fins

import (
	"encoding/binary"
)

func readCommand(ioAddr IoAddress, readCount uint16) *Payload {
	p := &Payload{
		CommandCode: CommandCodeMemoryAreaRead,
		Data:        make([]byte, 0, 6),
	}
	p.Data = append(p.Data, encodeIoAddress(ioAddr)...)
	p.Data = append(p.Data, []byte{0, 0}...)
	binary.BigEndian.PutUint16(p.Data[4:6], readCount)
	return p
}

func encodeIoAddress(ioAddr IoAddress) []byte {
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
