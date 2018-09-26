package fins

import "net"

// finsAddress A FINS device address
type finsAddress struct {
	network byte
	node    byte
	unit    byte
}

// Address A full device address
type Address struct {
	finsAddress finsAddress
	udpAddress  *net.UDPAddr
}

func NewAddress(ip string, port int, network, node, unit byte) Address {
	return Address{
		udpAddress: &net.UDPAddr{
			IP:   net.ParseIP(ip),
			Port: port,
		},
		finsAddress: finsAddress{
			network: network,
			node:    node,
			unit:    unit,
		},
	}
}
