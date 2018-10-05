package fins

// Header A FINS frame header
type Header struct {
	messageType      uint8
	responseRequired bool
	src              finsAddress
	dst              finsAddress
	serviceID        byte
	gatewayCount     uint8
}

const (
	// MessageTypeCommand Command message type
	MessageTypeCommand uint8 = iota

	// MessageTypeResponse Response message type
	MessageTypeResponse uint8 = iota
)

func defaultHeader(messageType uint8, responseRequired bool, src finsAddress, dst finsAddress, serviceID byte) Header {
	h := Header{}
	h.messageType = messageType
	h.responseRequired = responseRequired
	h.gatewayCount = 2
	h.src = src
	h.dst = dst
	h.serviceID = serviceID
	return h
}

func defaultCommandHeader(src finsAddress, dst finsAddress, serviceID byte) Header {
	h := defaultHeader(MessageTypeCommand, true, src, dst, serviceID)
	return h
}

func defaultResponseHeader(commandHeader Header) Header {
	h := defaultHeader(MessageTypeResponse, false, commandHeader.dst, commandHeader.src, commandHeader.serviceID)
	return h
}
