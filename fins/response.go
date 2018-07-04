package fins

// Response A FINS command response
type Response struct {
	CommandCode uint16
	EndCode     uint16
	Data        []byte
}
