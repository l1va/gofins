package fins

import (
	"encoding/binary"
)

func readDataCommand(startAddress uint16, readCount uint16) *Payload {
	return readCommand(IoAddress{
		MemoryArea: MemoryAreaDmWord,
		Address:    startAddress,
		BitOffset:  0,
	}, readCount)
}

func readWorkCommand(startAddress uint16, readCount uint16) *Payload {
	return readCommand(IoAddress{
		MemoryArea: MemoryAreaWrWord,
		Address:    startAddress,
		BitOffset:  0,
	}, readCount)
}

func readCommand(address IoAddress, readCount uint16) *Payload {
	addressLower := byte(startAddress)
	addressUpper := byte(startAddress >> 8)
	countLower := byte(readCount)
	countUpper := byte(readCount >> 8)

	paramsBytes := []byte{
		memoryArea,
		addressUpper, addressLower,
		addressBit,
		countUpper, countLower}

	bytes1 := append(header.Format(), CommandCodeMemoryAreaRead...)
	bytes2 := append(bytes1, paramsBytes...)
	return bytes2
}

func writeDataCommand(header *Header, startAddress uint16, data []uint16) *Payload {
	return writeCommand(MEMORY_AREA_DATA, header, startAddress, data)
}

func writeWorkCommand(header *Header, startAddress uint16, data []uint16) *Payload {
	return writeCommand(MEMORY_AREA_WORK, header, startAddress, data)
}

func writeCommand(memoryArea byte, header *Header, startAddress uint16, data []uint16) *Payload {
	var addressBit byte = 0
	addressLower := byte(startAddress)
	addressUpper := byte(startAddress >> 8)
	dataLen := len(data)
	lenLower := byte(dataLen)
	lenUpper := byte(dataLen >> 8)

	paramsBytes := []byte{
		memoryArea,
		addressUpper, addressLower,
		addressBit,
		lenUpper, lenLower}

	bytes1 := append(header.Format(), CommandCodeMemoryAreaWrite...)
	bytes2 := append(bytes1, paramsBytes...)
	bytes3 := append(bytes2, toBytes(data)...)
	return bytes3
}

func parseFrame(bytes []byte) (*Frame, error) {
	h := decodeHeader(bytes[:10])
	p := decodePayload(bytes[10:])
	return NewFrame(h, p), nil
	// if endCode != EndCodeNormalCompletion {
	// 	msg := fmt.Sprintf("Failed with end code: 0x%x", endCode)
	// 	return nil, errors.New(msg)
	// }
}

func decodeHeader(bytes []byte) *Header {
	h := &Header{
		icf: bytes[0],
		rsv: bytes[1],
		gct: bytes[2],
		dst: &FinsAddr{
			Network: bytes[3],
			Node:    bytes[4],
			Unit:    bytes[5],
		},
		src: &FinsAddr{
			Network: bytes[6],
			Node:    bytes[7],
			Unit:    bytes[8],
		},
		sid: bytes[9],
	}
	return h
}

func (h *Header) encodeHeader() []byte {
	bytes := []byte{
		h.icf, h.rsv, h.gct,
		h.dst.Network, h.dst.Node, h.dst.Unit,
		h.src.Network, h.src.Node, h.src.Unit,
		h.sid}
	return bytes
}

func decodePayload(bytes []byte) *Payload {
	p := &Payload{
		CommandCode: binary.BigEndian.Uint16(bytes[:2]),
		Data:        bytes[2:],
	}
	return p
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
