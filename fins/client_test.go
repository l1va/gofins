package fins

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFinsClient(t *testing.T) {
	clientAddr := NewAddress("", 9600, 0, 2, 0)
	plcAddr := NewAddress("", 9601, 0, 10, 0)

	toWrite := []uint16{5, 4, 3, 2, 1}

	s, e := NewPLCSimulator(plcAddr)
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
