package fins

// Payload A FINS frame payload
type Payload struct {
	CommandCode uint16
	Data        []byte
}
