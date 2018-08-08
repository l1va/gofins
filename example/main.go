package main

import (
	"fmt"

	"github.com/l1va/gofins/fins"
)

func main() {

	clientAddr := fins.NewAddress("192.168.250.2", 9600, 0, 2, 0)
	plcAddr := fins.NewAddress("192.168.250.10", 9600, 0, 10, 0)

	//s, e := fins.NewServer(plcAddr, nil)
	//if e != nil {
	//	panic(e)
	//}
	//defer s.Close()

	c, e := fins.NewClient(clientAddr, plcAddr)
	if e != nil {
		panic(e)
	}
	defer c.Close()

	z, _ := c.ReadWords(fins.MemoryAreaDMWord, 10000, 500)
	fmt.Println(z)

	// s, _ := c.ReadString(fins.MemoryAreaDMWord, 10000, 10)
	// fmt.Println(s)
	// fmt.Println(len(s))

	// b, _ := c.ReadBits(fins.MemoryAreaDMWord, 10473, 2, 1)
	// fmt.Println(b)
	// fmt.Println(len(b))

	// c.WriteWords(fins.MemoryAreaDMWord, 24000, []uint16{z[0] + 1, z[1] - 1})
	// c.WriteBits(fins.MemoryAreaDMBit, 24002, 0, []bool{false, false, false, true,
	// 	true, false, false, true,
	// 	false, false, false, false,
	// 	true, true, true, true})
	// c.SetBit(fins.MemoryAreaDMBit, 24003, 1)
	// c.ResetBit(fins.MemoryAreaDMBit, 24003, 0)
	// c.ToggleBit(fins.MemoryAreaDMBit, 24003, 2)

	// cron := cron.New()
	// s := rasc.NewShelter()
	// cron.AddFunc("*/5 * * * * *", func() {
	// 	t, _ := c.ReadClock()
	// 	fmt.Printf("Setting PLC time to: %s\n", t.Format(time.RFC3339))
	// 	c.WriteString(fins.MemoryAreaDMWord, 10000, 10, t.Format(time.RFC3339))
	// })
	// cron.Start()
}
