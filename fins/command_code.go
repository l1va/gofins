package fins

const (
	// CommandCodeMemoryAreaRead Command code: Memory area read
	CommandCodeMemoryAreaRead uint16 = 0x0101

	// CommandCodeMemoryAreaWrite Command code: Memory area write
	CommandCodeMemoryAreaWrite uint16 = 0x0102

	// CommandCodeMemoryAreaFill Command code: Memory area fill
	CommandCodeMemoryAreaFill uint16 = 0x0103

	// CommandCodeMultipleMemoryAreaRead Command code: Memory area multiple read
	CommandCodeMultipleMemoryAreaRead uint16 = 0x0104

	// CommandCodeMemoryAreaTransfer Command code: Memory area transfer
	CommandCodeMemoryAreaTransfer uint16 = 0x0105

	// CommandCodeParameterAreaRead Command code: Parameter area read
	CommandCodeParameterAreaRead uint16 = 0x0201

	// CommandCodeParameterAreaWrite Command code: Parameter area write
	CommandCodeParameterAreaWrite uint16 = 0x0202

	// CommandCodeParameterAreaClear Command code: Parameter area clear
	CommandCodeParameterAreaClear uint16 = 0x0203

	// CommandCodeProgramAreaRead Command code: Program area read
	CommandCodeProgramAreaRead uint16 = 0x0301

	// CommandCodeProgramAreaWrite Command code: Program area write
	CommandCodeProgramAreaWrite uint16 = 0x0302

	// CommandCodeProgramAreaClear Command code: Program area clear
	CommandCodeProgramAreaClear uint16 = 0x0303

	// CommandCodeRun Command code: Run
	CommandCodeRun uint16 = 0x0401

	// CommandCodeStop Command code: Stop
	CommandCodeStop uint16 = 0x0402

	// CommandCodeCPUUnitDataRead Command code: CPU unit data read
	CommandCodeCPUUnitDataRead uint16 = 0x0501

	// CommandCodeConnectionDataRead Command code: Connection data read
	CommandCodeConnectionDataRead uint16 = 0x0502

	// CommandCodeCPUUnitStatusRead Command code: CPU unit status read
	CommandCodeCPUUnitStatusRead uint16 = 0x0601

	// CommandCodeCycleTimeRead Command code: Cycle time read
	CommandCodeCycleTimeRead uint16 = 0x0620
)
