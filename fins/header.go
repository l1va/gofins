package fins

type Header struct {
	icf byte
	rsv byte
	gct byte
	dst FinsAddr
	src FinsAddr
	sid byte
}

const (
	icfBridgesBit          byte = 7
	icfMessageTypeBit      byte = 6
	icfResponseRequiredBit byte = 0
)

func (h *Header) IsResponseRequired() bool {
	return h.icf&1<<icfResponseRequiredBit == 0
}

func (h *Header) FrameIsCommand() bool {
	return h.icf&1<<icfMessageTypeBit == 0
}

func (h *Header) FrameIsResponse() bool {
	return !h.FrameIsCommand()
}

func (h *Header) SetToRequireResponse() {
	h.icf |= 1 << icfResponseRequiredBit
}

func (h *Header) SetToRequireNoResponse() {
	h.icf &^= 1 << icfResponseRequiredBit
}

func defaultHeader(dst FinsAddr, src FinsAddr, sid byte) *Header {
	h := new(Header)
	h.icf = 0x80
	h.rsv = 0x00
	h.gct = 0x02
	h.dst = dst
	h.src = src
	h.sid = sid
	return h
}

func newHeaderNoResponse(dst FinsAddr, src FinsAddr, sid byte) *Header {
	h := defaultHeader(dst, src, sid)
	h.SetToRequireNoResponse()
	return h
}
