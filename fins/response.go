package fins

// Response A FINS command response
type Response struct {
	payloadImpl
	endCode uint16
}

func NewResponse(commandCode uint16, endCode uint16, data []byte) *Response {
	r := new(Response)
	r.commandCode = commandCode
	r.endCode = endCode
	r.data = data
	return r
}

func (r *Response) EndCode() uint16 {
	return r.endCode
}
