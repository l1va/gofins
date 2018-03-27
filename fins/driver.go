package fins

import (
	"fmt"
	"errors"
)

/*
module.exports.Commands = {
    CONTROLLER_STATUS_READ : [0x06,0x01],
    MEMORY_AREA_READ       : [0x01,0x01],
    MEMORY_AREA_WRITE      : [0x01,0x02],
    MEMORY_AREA_FILL       : [0x01,0x03],
    RUN                    : [0x04,0x01],
    STOP                   : [0x04,0x02]
};
module.exports.MemoryAreas = {
    'E' : 0xA0,//Extended Memories
    'C' : 0xB0,//CIO
    'W' : 0xB1,//Work Area
    'H' : 0xB2,//Holding Bit
    'A' : 0xB3,//Auxiliary Bit
    'D' : 0x82//Data Memories
};
*/
//TODO: implement all areas and commands
var CMD_MEMORY_AREA_READ = []byte{0x01, 0x01}
var CMD_MEMORY_AREA_WRITE = []byte{0x01, 0x02}

var MEMORY_AREA_DATA = byte(0x82)
var MEMORY_AREA_WORK = byte(0xB1)

func readDCommand(header *Header, startAddress uint16, readCount uint16) []byte {
	return readCommand(MEMORY_AREA_DATA, header, startAddress, readCount)
}

func readWCommand(header *Header, startAddress uint16, readCount uint16) []byte {
	return readCommand(MEMORY_AREA_WORK, header, startAddress, readCount)
}

func readCommand(memoryArea byte, header *Header, startAddress uint16, readCount uint16) []byte {
	var addressBit byte = 0
	addressLower := byte(startAddress)
	addressUpper := byte(startAddress >> 8)
	countLower := byte(readCount)
	countUpper := byte(readCount >> 8)

	paramsBytes := []byte{
		memoryArea,
		addressUpper, addressLower,
		addressBit,
		countUpper, countLower}

	bytes1 := append(header.Format(), CMD_MEMORY_AREA_READ...)
	bytes2 := append(bytes1, paramsBytes...)
	return bytes2
}

func writeDCommand(header *Header, startAddress uint16, data []uint16) []byte {
	return writeCommand(MEMORY_AREA_DATA, header, startAddress, data)
}

func writeWCommand(header *Header, startAddress uint16, data []uint16) []byte {
	return writeCommand(MEMORY_AREA_WORK, header, startAddress, data)
}

func writeCommand(memoryArea byte, header *Header, startAddress uint16, data []uint16) []byte {
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

	bytes1 := append(header.Format(), CMD_MEMORY_AREA_WRITE...)
	bytes2 := append(bytes1, paramsBytes...)
	bytes3 := append(bytes2, toBytes(data)...)
	return bytes3
}

func toBytes(data []uint16) []byte {
	res := make([]byte, len(data)*2)
	for i := 0; i < len(data); i += 1 {
		res[2*i] = byte(data[i] >> 8)
		res[2*i+1] = byte(data[i])
	}
	return res
}

func parseResponse(bytes []byte) (*Response, error) {
	finishCode1 := bytes[12]
	finishCode2 := bytes[13]

	if finishCode1 != 0 || finishCode2 != 0 {
		msg := fmt.Sprintln("failure code: ", finishCode1, ": ", finishCode2)
		return nil, errors.New(msg)
	}

	return &Response{
		sid:  bytes[9],
		Data: toUint16(bytes[14:]),
	}, nil
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
