package fins

// IOAddress A FINS IO address representing some type of data or work area within the PLC
type IOAddress struct {
	MemoryArea byte
	Address    uint16
	BitOffset  byte
}

const (
	// MemoryAreaCIOBit Memory area: CIO bit
	MemoryAreaCIOBit byte = 0x30

	// MemoryAreaWRBit Memory area: WR bit
	MemoryAreaWRBit byte = 0x31

	// MemoryAreaHRBit Memory area: HR bit
	MemoryAreaHRBit byte = 0x32

	// MemoryAreaARBit Memory area: CIO bit
	MemoryAreaARBit byte = 0x33

	// MemoryAreaCIOWord Memory area: CIO bit
	MemoryAreaCIOWord byte = 0xb0

	// MemoryAreaWRWord Memory area: CIO bit
	MemoryAreaWRWord byte = 0xb1

	// MemoryAreaHRWord Memory area: CIO bit
	MemoryAreaHRWord byte = 0xb2

	// MemoryAreaARWord Memory area: CIO bit
	MemoryAreaARWord byte = 0xb3

	// MemoryAreaTimerCounterCompletionFlag Memory area: CIO bit
	MemoryAreaTimerCounterCompletionFlag byte = 0x09

	// MemoryAreaTimerCounterPV Memory area: CIO bit
	MemoryAreaTimerCounterPV byte = 0x89

	// MemoryAreaDMBit Memory area: CIO bit
	MemoryAreaDMBit byte = 0x02

	// MemoryAreaDMWord Memory area: CIO bit
	MemoryAreaDMWord byte = 0x82

	// MemoryAreaTaskBit Memory area: CIO bit
	MemoryAreaTaskBit byte = 0x06

	// MemoryAreaTaskStatus Memory area: CIO bit
	MemoryAreaTaskStatus byte = 0x46

	// MemoryAreaIndexRegisterPV Memory area: CIO bit
	MemoryAreaIndexRegisterPV byte = 0xdc

	// MemoryAreaDataRegisterPV Memory area: CIO bit
	MemoryAreaDataRegisterPV byte = 0xbc

	// MemoryAreaClockPulsesConditionFlagsBit Memory area: CIO bit
	MemoryAreaClockPulsesConditionFlagsBit byte = 0x07
)
