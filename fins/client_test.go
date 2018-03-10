package fins

import (
	"testing"
	"log"
	"net"
	"github.com/stretchr/testify/assert"
)

//TODO: implement me

func TestFinsClient(t *testing.T) {
	plcAddr := "127.0.0.1:8000"

	toWrite := []uint16{5, 4, 3, 2, 1}

	answers := map[byte][]byte{
		1: makeWriteAnswer(1),
		2: makeReadAnswer(2, toWrite),
	}
	plc := NewPLCMock(plcAddr, answers)
	defer plc.CloseConnection()

	c := NewClient(plcAddr)
	err := c.WriteD(100, toWrite)
	assert.Nil(t, err)
	vals, err := c.ReadD(100, 5)
	assert.Nil(t, err)
	assert.Equal(t, toWrite, vals)
}

func makeWriteAnswer(sid byte) []byte {
	ans := make([]byte, 14)
	ans[9] = sid
	return ans
}
func makeReadAnswer(sid byte, data []uint16) []byte {
	ans := make([]byte, 14)
	ans[9] = sid
	return append(ans, toBytes(data)...)
}

type PLCMock struct {
	answers map[byte][]byte
	pc      net.PacketConn
}

func NewPLCMock(plcAddr string, answers map[byte][]byte) *PLCMock {
	c := new(PLCMock)
	c.answers = answers

	pc, err := net.ListenPacket("udp", plcAddr)
	if err != nil {
		log.Fatal(err)
	}
	c.pc = pc

	go c.listenLoop()

	return c

}

func (c *PLCMock) CloseConnection() {
	c.pc.Close()
}

func (c *PLCMock) listenLoop() {
	for {
		buf := make([]byte, 1024)
		n, addr, err := c.pc.ReadFrom(buf)
		if err != nil {
			log.Fatal(err)
		}

		if n > 0 {
			sid := buf[9]
			c.pc.WriteTo(c.answers[sid], addr)
		} else {
			log.Println("cannot read request: ", buf)
		}
	}
}
