package fins

// IOAddress A FINS IO address representing some type of data or work area within the PLC
type IOAddress struct {
	MemoryArea byte
	Address    uint16
	BitOffset  byte
}
