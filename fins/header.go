package fins

// Header A FINS frame header
type Header struct {
	icf byte
	rsv byte
	gct byte
	dst Address
	src Address
	sid byte
}

const (
	icfBridgesBit          byte = 7
	icfMessageTypeBit      byte = 6
	icfResponseRequiredBit byte = 0
)

// IsResponseRequired Returns true if this header indicates that a response should be required
func (h *Header) IsResponseRequired() bool {
	return h.icf&1<<icfResponseRequiredBit == 0
}

// FrameIsCommand Returns true if the frame this header was contained within was a command
func (h *Header) FrameIsCommand() bool {
	return h.icf&1<<icfMessageTypeBit == 0
}

// FrameIsResponse Returns true if the frame this header was contained within was a response
func (h *Header) FrameIsResponse() bool {
	return !h.FrameIsCommand()
}

// SetToRequireResponse Will set this header to indicate that a response is required
func (h *Header) SetToRequireResponse() {
	h.icf |= 1 << icfResponseRequiredBit
}

// SetToRequireNoResponse Will set this header to indicate that a response is not required
func (h *Header) SetToRequireNoResponse() {
	h.icf &^= 1 << icfResponseRequiredBit
}

func defaultHeader(dst Address, src Address, sid byte) *Header {
	h := new(Header)
	h.icf = 0x80
	h.rsv = 0x00
	h.gct = 0x02
	h.dst = dst
	h.src = src
	h.sid = sid
	return h
}

func newHeaderNoResponse(dst Address, src Address, sid byte) *Header {
	h := defaultHeader(dst, src, sid)
	h.SetToRequireNoResponse()
	return h
}
