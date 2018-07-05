package fins

// IOAddress A FINS IO address representing some type of data or work area within the PLC
type IOAddress struct {
	memoryArea byte
	address    uint16
	bitOffset  byte
}

func NewIOAddress(memoryArea byte, address uint16) *IOAddress {
	return NewIOAddressWithBitOffset(memoryArea, address, 0)
}

func NewIOAddressWithBitOffset(memoryArea byte, address uint16, bitOffset byte) *IOAddress {
	ioAddr := new(IOAddress)
	ioAddr.memoryArea = memoryArea
	ioAddr.address = address
	ioAddr.bitOffset = bitOffset
	return ioAddr
}

func (ioAddr *IOAddress) MemoryArea() byte {
	return ioAddr.memoryArea
}

func (ioAddr *IOAddress) Address() uint16 {
	return ioAddr.address
}

func (ioAddr *IOAddress) BitOffset() byte {
	return ioAddr.bitOffset
}
