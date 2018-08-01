package fins

// Frame A FINS frame
type Frame struct {
	header  *Header
	payload Payload
}

// NewFrame Creates a new FINS frame
func NewFrame(header *Header, payload Payload) *Frame {
	f := new(Frame)
	f.header = header
	f.payload = payload
	return f
}

func (f *Frame) Header() *Header {
	return f.header
}

func (f *Frame) Payload() Payload {
	return f.payload
}
