package fins

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"encoding/binary"
)

func TestFinsClient(t *testing.T) {
	clientAddr := NewAddress("", 9600, 0, 2, 0)
	plcAddr := NewAddress("", 9601, 0, 10, 0)

	toWrite := []uint16{5, 4, 3, 2, 1}
	handler := func(req request, mem []byte) response {
		l := uint16(len(toWrite))
		bts := make([]byte, 2*l, 2*l)
		for i := 0; i < int(l); i++ {
			binary.BigEndian.PutUint16(bts[i*2:i*2+2], toWrite[i])
		}
		return response{
			header:      defaultResponseHeader(req.header),
			commandCode: req.commandCode,
			endCode:     EndCodeNormalCompletion,
			data:        bts,
		}
	}

	s, e := NewServer(plcAddr, handler)
	if e != nil {
		panic(e)
	}
	defer s.Close()

	c, e := NewClient(clientAddr, plcAddr)
	if e != nil {
		panic(e)
	}
	defer c.Close()

	err := c.WriteWords(MemoryAreaDMWord, 100, toWrite)
	assert.Nil(t, err)

	vals, err := c.ReadWords(MemoryAreaDMWord, 100, 5)
	assert.Nil(t, err)
	assert.Equal(t, toWrite, vals)

}
