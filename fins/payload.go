package fins

import (
	"encoding/binary"
)

type Payload struct {
	commandCode uint16
	data        []byte
}

func parsePayload(bytes []byte) *Payload {
	p := &Payload{}
	p.commandCode = binary.BigEndian.Uint16(bytes[:2])
	p.data = bytes[2:]
	return p
}

const (
	// Normal
	EndCodeNormalCompletion uint16 = 0x0000

	//  Local Node Error
	EndCodeLocalNodeNotInNetwork       uint16 = 0x0101
	EndCodeTokenTimeout                uint16 = 0x0102
	EndCodeRetriesFailed               uint16 = 0x0103
	EndCodeTooManySendFrames           uint16 = 0x0104
	EndCodeNodeAddressRangeError       uint16 = 0x0105
	EndCodeNodeAddressRangeDuplication uint16 = 0x0106

	//  Destination Node Error
	EndCodeDestinationNodeNotInNetwork uint16 = 0x0201
	EndCodeUnitMissing                 uint16 = 0x0202
	EndCodeThirdNodeMissing            uint16 = 0x0203
	EndCodeDestinationNodeBusy         uint16 = 0x0204
	EndCodeResponseTimeout             uint16 = 0x0205

	//  Controller Error
	EndCodeCommunicationsControllerError uint16 = 0x0301
	EndCodeCpuUnitError                  uint16 = 0x0302
	EndCodeControllerError               uint16 = 0x0303
	EndCodeUnitNumberError               uint16 = 0x0304

	// Service Unsupported
	EndCodeUndefinedCommand           uint16 = 0x0401
	EndCodeNotSupportedByModelVersion uint16 = 0x0402

	// Routing Table Error
	EndCodeDestinationAddressSettingError uint16 = 0x0501
	EndCodeNoRoutingTables                uint16 = 0x0502
	EndCodeRoutingTableError              uint16 = 0x0503
	EndCodeTooManyRelays                  uint16 = 0x0504

	// Command Format Error
	EndCodeCommandTooLong        uint16 = 0x1001
	EndCodeCommandTooShort       uint16 = 0x1002
	EndCodeElementsDataDontMatch uint16 = 0x1003
	EndCodeCommandFormatError    uint16 = 0x1004
	EndCodeHeaderError           uint16 = 0x1005

	// Parameter Error
	EndCodeAreaClassificationMissing uint16 = 0x1101
	EndCodeAccessSizeError           uint16 = 0x1102
	EndCodeAddressRangeError         uint16 = 0x1103
	EndCodeAddressRangeExceeded      uint16 = 0x1104
	EndCodeProgramMissing            uint16 = 0x1106
	EndCodeRelationalError           uint16 = 0x1109
	EndCodeDuplicateDataAccess       uint16 = 0x110a
	EndCodeResponseTooBig            uint16 = 0x110b
	EndCodeParameterError            uint16 = 0x110c

	// Todo Finish entering all the end codes from Omron document W342 section 5-1-3
	EndCodeWriteNotPossibleReadOnly              uint16 = 0x2101
	EndCodeWriteNotPossibleProtected             uint16 = 0x2102
	EndCodeWriteNotPossibleCannotRegister        uint16 = 0x2103
	EndCodeWriteNotPossibleProgramMissing        uint16 = 0x2105
	EndCodeWriteNotPossibleFileMissing           uint16 = 0x2106
	EndCodeWriteNotPossibleFileNameAlreadyExists uint16 = 0x2107
	EndCodeWriteNotPossibleCannotChange          uint16 = 0x2108

	EndCodeNotExecutableInCurrentModeNotPossibleDuringExecution  uint16 = 0x2201
	EndCodeNotExecutableInCurrentModeNotPossibleWhileRunning     uint16 = 0x2202
	EndCodeNotExecutableInCurrentModeWrongPlcModeInProgram       uint16 = 0x2203
	EndCodeNotExecutableInCurrentModeWrongPlcModeInDebug         uint16 = 0x2204
	EndCodeNotExecutableInCurrentModeWrongPlcModeInMonitor       uint16 = 0x2205
	EndCodeNotExecutableInCurrentModeWrongPlcModeInRun           uint16 = 0x2206
	EndCodeNotExecutableInCurrentModeSpecifiedNodeNotPollingNode uint16 = 0x2207
	EndCodeNotExecutableInCurrentModeStepCannotBeExecuted        uint16 = 0x2208

	EndCodeNoSuchDeviceFileDeviceMissing uint16 = 0x2301
	EndCodeNoSuchDeviceMemoryMissing     uint16 = 0x2302
	EndCodeNoSuchDeviceClockMissing      uint16 = 0x2303

	EndCodeCannotStartStopTableMissing uint16 = 0x2401
)

const (
	CommandCodeMemoryAreaRead         uint16 = 0x0101
	CommandCodeMemoryAreaWrite        uint16 = 0x0102
	CommandCodeMemoryAreaFill         uint16 = 0x0103
	CommandCodeMultipleMemoryAreaRead uint16 = 0x0104
	CommandCodeMemoryAreaTransfer     uint16 = 0x0105
	CommandCodeParameterAreaRead      uint16 = 0x0201
	CommandCodeParameterAreaWrite     uint16 = 0x0202
	CommandCodeParameterAreaClear     uint16 = 0x0203
	CommandCodeProgramAreaRead        uint16 = 0x0301
	CommandCodeProgramAreaWrite       uint16 = 0x0302
	CommandCodeProgramAreaClear       uint16 = 0x0303
	CommandCodeRun                    uint16 = 0x0401
	CommandCodeStop                   uint16 = 0x0402
	CommandCodeCpuUnitDataRead        uint16 = 0x0501
	CommandCodeConnectionDataRead     uint16 = 0x0502
	CommandCodeCpuUnitStatusRead      uint16 = 0x0601
	CommandCodeCycleTimeRead          uint16 = 0x0620
)
