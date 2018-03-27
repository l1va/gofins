package fins

import (
	"testing"
	"log"
	"net"
	"github.com/stretchr/testify/assert"
)

//TODO: implement me

func TestFinsClient(t *testing.T) {
	plcAddr := ":9600"

	toWrite := []uint16{5, 4, 3, 2, 1}

	answers := map[byte]response{
		1: makeWriteAnswer(1, true),
		2: makeReadAnswer(2, toWrite),
		3: makeWriteAnswer(3, false),
		4: makeReadAnswer(4, toWrite),
	}
	plc := NewPLCMock(plcAddr, answers)
	defer plc.CloseConnection()

	c := NewClient(plcAddr)
	defer c.CloseConnection()
	err := c.WriteD(100, toWrite)
	assert.Nil(t, err)
	vals, err := c.ReadD(100, 5)
	assert.Nil(t, err)
	assert.Equal(t, toWrite, vals)
	err = c.WriteDNoResponse(200, toWrite)
	assert.Nil(t, err)
	vals, err = c.ReadD(200, 5)
	assert.Nil(t, err)
	assert.Equal(t, toWrite, vals)
}

type response struct {
	data   []byte
	needed bool
}

func makeWriteAnswer(sid byte, respNeeded bool) response {
	ans := make([]byte, 14)
	ans[9] = sid
	return response{data: ans, needed: respNeeded}
}
func makeReadAnswer(sid byte, data []uint16) response {
	ans := make([]byte, 14)
	ans[9] = sid
	return response{data: append(ans, toBytes(data)...), needed: true}
}

type PLCMock struct {
	answers map[byte]response
	pc      net.PacketConn
}

func NewPLCMock(plcAddr string, answers map[byte]response) *PLCMock {
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
		buf := make([]byte, 2048)
		n, addr, err := c.pc.ReadFrom(buf)
		if err != nil {
			log.Fatal(err)
		}

		if n > 0 {
			sid := buf[9]
			ans, exist := c.answers[sid]
			if exist {
				if ans.needed {
					c.pc.WriteTo(c.answers[sid].data, addr)
				}
			} else {
				log.Fatal("There is no answer for sid =", sid)
			}
		} else {
			log.Fatal("Cannot read request: ", buf)
		}
	}
}
