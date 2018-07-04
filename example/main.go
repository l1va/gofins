package main

import (
	"fmt"
	"time"

	"github.com/siyka-au/gofins/fins"
)

func main() {
	c := fins.NewClient("192.168.250.10:9600", fins.Address{
		Network: 0,
		Node:    10,
		Unit:    0,
	},
		fins.Address{
			Network: 0,
			Node:    2,
			Unit:    0,
		})
	defer c.CloseConnection()

	// z, _ := c.ReadWords(fins.MemoryAreaDMWord, 10470, 4)
	// fmt.Println(z)

	// s, _ := c.ReadString(fins.MemoryAreaDMWord, 10000, 10)
	// fmt.Println(s)
	// fmt.Println(len(s))

	t, _ := c.ReadClock()
	fmt.Println(t.Format(time.RFC3339))
}
