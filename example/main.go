package main

import (
	"fmt"

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

	z, _ := c.ReadWords(fins.MemoryAreaDmWord, 10470, 2)

	fmt.Println(z)
}
