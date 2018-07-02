package fins

type Frame struct {
	header  *Header
	payload *Payload
}

func NewFrame(header *Header, payload *Payload) *Frame {
	f := &Frame{}
	f.header = header
	f.payload = payload
	return f
}
