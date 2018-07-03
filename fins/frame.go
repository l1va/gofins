package fins

// Frame A FINS frame
type Frame struct {
	Header  *Header
	Payload *Payload
}

// NewFrame Creates a new FINS frame
func NewFrame(header *Header, payload *Payload) *Frame {
	f := &Frame{
		Header:  header,
		Payload: payload,
	}
	return f
}
