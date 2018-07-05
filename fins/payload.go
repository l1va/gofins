package fins

// Payload A FINS frame payload
type Payload interface {
	CommandCode() uint16
	Data() []byte
}

type payloadImpl struct {
	commandCode uint16
	data        []byte
}

func (p *payloadImpl) CommandCode() uint16 {
	return p.commandCode
}

func (p *payloadImpl) Data() []byte {
	return p.data
}
