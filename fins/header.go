package fins

// Header A FINS frame header
type Header struct {
	messgeType       uint8
	responseRequired bool
	dst              *Address
	src              *Address
	serviceID        byte
	gatewayCount     uint8
}

const (
	// MessageTypeCommand Command message type
	MessageTypeCommand uint8 = iota

	// MessageTypeResponse Response message type
	MessageTypeResponse uint8 = iota
)

// IsResponseRequired Returns true if this header indicates that a response should be required
func (h *Header) IsResponseRequired() bool {
	return h.responseRequired
}

// FrameIsCommand Returns true if the frame this header was contained within was a command
func (h *Header) FrameIsCommand() bool {
	return h.messgeType == MessageTypeCommand
}

// FrameIsResponse Returns true if the frame this header was contained within was a response
func (h *Header) FrameIsResponse() bool {
	return h.messgeType == MessageTypeResponse
}

// SetToRequireResponse Will set this header to indicate that a response is required
// func (h *Header) SetToRequireResponse() {
// 	h.responseRequired = true
// }

// SetToRequireNoResponse Will set this header to indicate that a response is not required
// func (h *Header) SetToRequireNoResponse() {
// 	h.responseRequired = false
// }

// SetToCommandMessageType Will set this header to indicate that the message is a command
// func (h *Header) SetToCommandMessageType() {
// 	h.messgeType = MessageTypeCommand
// }

// SetToResponseMessageType Will set this header to indicate that the message is a response
// func (h *Header) SetToResponseMessageType() {
// 	h.messgeType = MessageTypeResponse
// }

// GatewayCount Gets the gateway count
func (h *Header) GatewayCount() byte {
	return h.gatewayCount
}

// SourceAddress Gets the source address
func (h *Header) SourceAddress() Address {
	return *h.src
}

// DestinationAddress Gets the destination address
func (h *Header) DestinationAddress() Address {
	return *h.dst
}

// ServiceID Gets the service id
func (h *Header) ServiceID() byte {
	return h.serviceID
}

func defaultHeader(messageType uint8, responseRequired bool, dst *Address, src *Address, serviceID byte) *Header {
	h := new(Header)
	h.messgeType = messageType
	h.responseRequired = responseRequired
	h.gatewayCount = 2
	h.dst = dst
	h.src = src
	h.serviceID = serviceID
	return h
}

func defaultCommandHeader(dst *Address, src *Address, serviceID byte) *Header {
	h := defaultHeader(MessageTypeCommand, true, src, dst, serviceID)
	return h
}

func defaultResponseHeader(commandHeader *Header) *Header {
	h := defaultHeader(MessageTypeResponse, false, commandHeader.src, commandHeader.dst, commandHeader.serviceID)
	return h
}
