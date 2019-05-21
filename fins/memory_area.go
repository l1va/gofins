package fins

const (
	// MemoryAreaCIOBit Memory area: CIO area; bit
	MemoryAreaCIOBit byte = 0x30

	// MemoryAreaWRBit Memory area: work area; bit
	MemoryAreaWRBit byte = 0x31

	// MemoryAreaHRBit Memory area: holding area; bit
	MemoryAreaHRBit byte = 0x32

	// MemoryAreaARBit Memory area: axuillary area; bit
	MemoryAreaARBit byte = 0x33

	// MemoryAreaCIOWord Memory area: CIO area; word
	MemoryAreaCIOWord byte = 0xb0

	// MemoryAreaWRWord Memory area: work area; word
	MemoryAreaWRWord byte = 0xb1

	// MemoryAreaHRWord Memory area: holding area; word
	MemoryAreaHRWord byte = 0xb2

	// MemoryAreaARWord Memory area: auxillary area; word
	MemoryAreaARWord byte = 0xb3

	// MemoryAreaTimerCounterCompletionFlag Memory area: counter completion flag
	MemoryAreaTimerCounterCompletionFlag byte = 0x09

	// MemoryAreaTimerCounterPV Memory area: counter PV
	MemoryAreaTimerCounterPV byte = 0x89

	// MemoryAreaDMBit Memory area: data area; bit
	MemoryAreaDMBit byte = 0x02

	// MemoryAreaDMWord Memory area: data area; word
	MemoryAreaDMWord byte = 0x82

	// MemoryAreaTaskBit Memory area: task flags; bit
	MemoryAreaTaskBit byte = 0x06

	// MemoryAreaTaskStatus Memory area: task flags; status
	MemoryAreaTaskStatus byte = 0x46

	// MemoryAreaIndexRegisterPV Memory area: CIO bit
	MemoryAreaIndexRegisterPV byte = 0xdc

	// MemoryAreaDataRegisterPV Memory area: CIO bit
	MemoryAreaDataRegisterPV byte = 0xbc

	// MemoryAreaClockPulsesConditionFlagsBit Memory area: CIO bit
	MemoryAreaClockPulsesConditionFlagsBit byte = 0x07
)
