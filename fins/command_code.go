package fins

const (
	// CommandCodeMemoryAreaRead Command code: IO memory area read
	CommandCodeMemoryAreaRead uint16 = 0x0101

	// CommandCodeMemoryAreaWrite Command code: IO memory area write
	CommandCodeMemoryAreaWrite uint16 = 0x0102

	// CommandCodeMemoryAreaFill Command code: IO memory area fill
	CommandCodeMemoryAreaFill uint16 = 0x0103

	// CommandCodeMultipleMemoryAreaRead Command code: IO memory area multiple read
	CommandCodeMultipleMemoryAreaRead uint16 = 0x0104

	// CommandCodeMemoryAreaTransfer Command code: IO memory area transfer
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

	// CommandCodeRun Command code: Set operating mode to run
	CommandCodeRun uint16 = 0x0401

	// CommandCodeStop Command code: Set operating mode to stop
	CommandCodeStop uint16 = 0x0402

	// CommandCodeCPUUnitDataRead Command code: CPU unit data read
	CommandCodeCPUUnitDataRead uint16 = 0x0501

	// CommandCodeConnectionDataRead Command code: connection data read
	CommandCodeConnectionDataRead uint16 = 0x0502

	// CommandCodeCPUUnitStatusRead Command code: CPU unit status read
	CommandCodeCPUUnitStatusRead uint16 = 0x0601

	// CommandCodeCycleTimeRead Command code: cycle time read
	CommandCodeCycleTimeRead uint16 = 0x0620

	// CommandCodeClockRead Command code: clock read
	CommandCodeClockRead uint16 = 0x701

	// CommandCodeClockWrite Command code: clock write
	CommandCodeClockWrite uint16 = 0x702

	// CommandCodeMessageReadClear Command code: message read/clear
	CommandCodeMessageReadClear uint16 = 0x0920

	// CommandCodeAccessRightAcquire Command code: access right acquire
	CommandCodeAccessRightAcquire uint16 = 0x0c01

	// CommandCodeAccessRightForcedAcquire Command code: accress right forced acquire
	CommandCodeAccessRightForcedAcquire uint16 = 0x0c02

	// CommandCodeAccessRightRelease Command code: access right release
	CommandCodeAccessRightRelease uint16 = 0x0c03

	// CommandCodeErrorClear Command code: error clear
	CommandCodeErrorClear uint16 = 0x2101

	// CommandCodeErrorLogRead Command code: error log read
	CommandCodeErrorLogRead uint16 = 0x2102

	// CommandCodeErrorLogClear Command code: error log clear
	CommandCodeErrorLogClear uint16 = 0x2103

	// CommandCodeFINSWriteAccessLogRead Command code: FINS write access log read
	CommandCodeFINSWriteAccessLogRead uint16 = 0x2140

	// CommandCodeFINSWriteAccessLogWrite Command code: FINS write access log write
	CommandCodeFINSWriteAccessLogWrite uint16 = 0x2141

	// CommandCodeFileNameRead Command code: file name read
	CommandCodeFileNameRead uint16 = 0x2101

	// CommandCodeSingleFileRead Command code: file read
	CommandCodeSingleFileRead uint16 = 0x2102

	// CommandCodeSingleFileWrite Command code: file write
	CommandCodeSingleFileWrite uint16 = 0x2103

	// CommandCodeFileMemoryFormat Command code: file memory format
	CommandCodeFileMemoryFormat uint16 = 0x2104

	// CommandCodeFileDelete Command code: file delete
	CommandCodeFileDelete uint16 = 0x2105

	// CommandCodeFileCopy Command code: file copy
	CommandCodeFileCopy uint16 = 0x2107

	// CommandCodeFileNameChange Command code: file name change
	CommandCodeFileNameChange uint16 = 0x2108

	// CommandCodeMemoryAreaFileTransfer Command code: memory area file transfer
	CommandCodeMemoryAreaFileTransfer uint16 = 0x210a

	// CommandCodeParameterAreaFileTransfer Command code: parameter area file transfer
	CommandCodeParameterAreaFileTransfer uint16 = 0x210b

	// CommandCodeProgramAreaFileTransfer Command code: program area file transfer
	CommandCodeProgramAreaFileTransfer uint16 = 0x210b

	// CommandCodeDirectoryCreateDelete Command code: directory create/delete
	CommandCodeDirectoryCreateDelete uint16 = 0x2115

	// CommandCodeMemoryCassetteTransfer Command code: memory cassette transfer (CP1H and CP1L CPU units only)
	CommandCodeMemoryCassetteTransfer uint16 = 0x2120

	// CommandCodeForcedSetReset Command code: forced set/reset
	CommandCodeForcedSetReset uint16 = 0x2301

	// CommandCodeForcedSetResetCancel Command code: forced set/reset cancel
	CommandCodeForcedSetResetCancel uint16 = 0x2302

	// CommandCodeConvertToCompoWayFCommand Command code: convert to CompoWay/F command
	CommandCodeConvertToCompoWayFCommand uint16 = 0x2803

	// CommandCodeConvertToModbusRTUCommand Command code: convert to Modbus-RTU command
	CommandCodeConvertToModbusRTUCommand uint16 = 0x2804

	// CommandCodeConvertToModbusASCIICommand Command code: convert to Modbus-ASCII command
	CommandCodeConvertToModbusASCIICommand uint16 = 0x2805
)
