package fins

type Frame struct {
	Header  *Header
	Payload *Payload
}

func NewFrame(header Header, payload Payload) *Frame {
	f := &Frame{
		Header:  header,
		Payload: payload,
	}
	return f
}
