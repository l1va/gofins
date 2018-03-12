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

	err := c.WriteDAsync(100, []uint16{5, 4, 3, 2, 1}, func(fins.Response) {
		log.Println("writing done!")
	})
	if err != nil {
		log.Println("writing request failed:", err)
	}

	err = c.ReadDAsync(100, 5, func(r fins.Response) {
		log.Println("readed values: ", r.Data)
	})
	if err != nil {
		log.Println("reading request failed:", err)
	}

	for i := 0; i < 10; i += 1 {
		t := uint16(i * 5)
		err := c.WriteD(200+t, []uint16{t, t + 1, t + 2, t + 3, t + 4})
		if err != nil {
			log.Fatal(err)
		}
	}

	err = c.WriteDNoResponse(200, []uint16{5, 4, 3, 2, 1})
	if err != nil {
		log.Println("writing request without response failed:", err)
	}

	vals, err := c.ReadD(200, 50)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(vals)
}
