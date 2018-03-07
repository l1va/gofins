package fins

// For now we have only one PLC - it means we do not need in any networks settings -
// almost everywhere is 0.
type Header struct {
	icf byte
	rsv byte
	gct byte
	dna byte
	da1 byte
	da2 byte
	sna byte
	sa1 byte
	sa2 byte
	sid byte
}

func newHeader(sid byte) *Header {
	h := new(Header)
	h.icf = icf()
	h.rsv = rsv()
	h.gct = gct()
	h.dna = dstNetwork()
	h.da1 = dstNode()
	h.da2 = dstUnit()
	h.sna = srcNetwork()
	h.sa1 = srcNode()
	h.sa2 = srcUnit()
	h.sid = sid
	return h
}

func (f *Header) Format() []byte {

	return []byte{
		f.icf, f.rsv, f.gct,
		f.dna, f.da1, f.da2,
		f.sna, f.sa1, f.sa2,
		f.sid}
}

func icf() byte {
	return 0x80 //128
}

func rsv() byte {
	return 0
}

func gct() byte {
	return 0x02
}

func dstNetwork() byte {
	return 0
}

func dstNode() byte {
	return 0
}

func dstUnit() byte {
	return 0
}

func srcNetwork() byte {
	return 0
}

func srcNode() byte {
	return 0x22
}

func srcUnit() byte {
	return 0
}
