package fins

import (
	"encoding/binary"
	"math"
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

	// ------------- Test Words
	err := c.WriteWords(MemoryAreaDMWord, 100, toWrite)
	assert.Nil(t, err)

	vals, err := c.ReadWords(MemoryAreaDMWord, 100, 5)
	assert.Nil(t, err)
	assert.Equal(t, toWrite, vals)

	// test setting response timeout
	c.SetTimeoutMs(50)
	_, err = c.ReadWords(MemoryAreaDMWord, 100, 5)
	assert.Nil(t, err)

	// ------------- Test Strings
	err = c.WriteString(MemoryAreaDMWord, 10, "ф1234")
	assert.Nil(t, err)

	v, err := c.ReadString(MemoryAreaDMWord, 12, 1)
	assert.Nil(t, err)
	assert.Equal(t, "12", v)

	v, err = c.ReadString(MemoryAreaDMWord, 10, 3)
	assert.Nil(t, err)
	assert.Equal(t, "ф1234", v)

	v, err = c.ReadString(MemoryAreaDMWord, 10, 5)
	assert.Nil(t, err)
	assert.Equal(t, "ф1234", v)

	// ------------- Test Bytes
	err = c.WriteBytes(MemoryAreaDMWord, 10, []byte{0x00, 0x00 ,0xC1 , 0xA0})
	assert.Nil(t, err)

	b, err := c.ReadBytes(MemoryAreaDMWord, 10, 2)
	assert.Nil(t, err)
	assert.Equal(t, []byte{0x00, 0x00 ,0xC1 , 0xA0}, b)

	buf := make([]byte, 8, 8)
	binary.LittleEndian.PutUint64(buf[:], math.Float64bits(-20))
	err = c.WriteBytes(MemoryAreaDMWord, 10, buf)
	assert.Nil(t, err)

	b, err = c.ReadBytes(MemoryAreaDMWord, 10, 4)
	assert.Nil(t, err)
	assert.Equal(t, []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x34, 0xc0}, b)


	// ------------- Test Bits
	err = c.WriteBits(MemoryAreaDMBit, 10, 2, []bool{true, false, true})
	assert.Nil(t, err)

	bs, err := c.ReadBits(MemoryAreaDMBit, 10, 2, 3)
	assert.Nil(t, err)
	assert.Equal(t, []bool{true, false, true}, bs)

	bs, err = c.ReadBits(MemoryAreaDMBit, 10, 1, 5)
	assert.Nil(t, err)
	assert.Equal(t, []bool{false, true, false, true, false}, bs)

}
