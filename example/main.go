package main

import (
	"fmt"
	"log"
	"github.com/l1va/gofins/fins"
)

func main() {
	plcAddr := "192.168.250.1:9600"

	c := fins.NewClient(plcAddr)
	defer c.CloseConnection()

	for i := 0; i < 10; i += 1 {
		t := uint16(i * 5)
		err := c.WriteDM(200, []uint16{t, t + 1, t + 2, t + 3, t + 4})
		if err != nil {
			log.Fatal(err)
		}
	}

	vals, err := c.ReadDM(200, 50)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(vals)
}
