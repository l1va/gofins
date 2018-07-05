package fins

// Address A FINS device address
type Address struct {
	network byte
	node    byte
	unit    byte
}

func NewAddress(network byte, node byte, unit byte) *Address {
	a := new(Address)
	a.network = network
	a.node = node
	a.unit = unit
	return a
}

func (a *Address) Network() byte {
	return a.network
}

func (a *Address) Node() byte {
	return a.node
}

func (a *Address) Unit() byte {
	return a.unit
}
