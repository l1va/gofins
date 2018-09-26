package fins

// Header A FINS frame header
type Header struct {
	messageType      uint8
	responseRequired bool
	dst              finsAddress
	src              finsAddress
	serviceID        byte
	gatewayCount     uint8
}

const (
	// MessageTypeCommand Command message type
	MessageTypeCommand uint8 = iota

	// MessageTypeResponse Response message type
	MessageTypeResponse uint8 = iota
)

func defaultHeader(messageType uint8, responseRequired bool, dst finsAddress, src finsAddress, serviceID byte) Header {
	h := Header{}
	h.messageType = messageType
	h.responseRequired = responseRequired
	h.gatewayCount = 2
	h.dst = dst
	h.src = src
	h.serviceID = serviceID
	return h
}

func defaultCommandHeader(dst finsAddress, src finsAddress, serviceID byte) Header {
	h := defaultHeader(MessageTypeCommand, true, src, dst, serviceID)
	return h
}

func defaultResponseHeader(commandHeader Header) Header {
	h := defaultHeader(MessageTypeResponse, false, commandHeader.src, commandHeader.dst, commandHeader.serviceID)
	return h
}
