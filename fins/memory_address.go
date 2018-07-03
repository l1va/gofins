package fins

type IoAddress struct {
	MemoryArea byte
	Address    uint16
	BitOffset  byte
}

const (
	MemoryAreaCioBit                       byte = 0x30
	MemoryAreaWrBit                        byte = 0x31
	MemoryAreaHrBit                        byte = 0x32
	MemoryAreaArBit                        byte = 0x33
	MemoryAreaCioWord                      byte = 0xb0
	MemoryAreaWrWord                       byte = 0xb1
	MemoryAreaHrWord                       byte = 0xb2
	MemoryAreaArWord                       byte = 0xb3
	MemoryAreaTimerCounterCompletionFlag   byte = 0x09
	MemoryAreaTimerCounterPv               byte = 0x89
	MemoryAreaDmBit                        byte = 0x02
	MemoryAreaDmWord                       byte = 0x82
	MemoryAreaTaskBit                      byte = 0x06
	MemoryAreaTaskStatus                   byte = 0x46
	MemoryAreaIndexRegisterPv              byte = 0xdc
	MemoryAreaDataRegisterPv               byte = 0xbc
	MemoryAreaClockPulsesConditionFlagsBit byte = 0x07
)
