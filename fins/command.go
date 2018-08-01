package fins

// Command A FINS command
type Command struct {
	payloadImpl
}

func NewCommand(commandCode uint16, data []byte) *Command {
	c := new(Command)
	c.commandCode = commandCode
	c.data = data
	return c
}
